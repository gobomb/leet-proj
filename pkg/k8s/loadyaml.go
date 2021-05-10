package k8s

import (
	"log"
	"reflect"
	"strings"

	jsonpatch "github.com/evanphx/json-patch"
	"github.com/pkg/errors"
	appv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/util/strategicpatch"
	"k8s.io/apimachinery/pkg/util/yaml"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"

	"k8s.io/apimachinery/pkg/types"

	"k8s.io/apimachinery/pkg/runtime"
)

var (
	scheme = runtime.NewScheme()
)

func init() {
	log.SetFlags(log.Lshortfile)
	utilruntime.Must(clientgoscheme.AddToScheme(scheme))
	utilruntime.Must(corev1.AddToScheme(scheme))
}

func load(src []byte) (runtime.Object, error) {
	decode := serializer.NewCodecFactory(scheme).UniversalDeserializer().Decode
	obj, gvk, err := decode(src, nil, nil)
	if err != nil {
		log.Printf("%v\n", err)
		return nil, errors.Wrap(err, "unable to unmarshal Moudule.Spec.DeploymentConfigYaml.")
	}
	log.Printf("gvk %v\n", gvk)

	return obj, nil
}

func checkdeploy(obj runtime.Object) (*appv1.Deployment, error) {
	deploy, ok := obj.(*appv1.Deployment)
	if !ok {
		return nil, errors.Errorf("could not reflect the deploy")
	}
	name := deploy.Spec.Template.Spec.Containers[0].Name
	log.Printf("name %v\n", name)
	return deploy, nil
}

// ref: https://github.com/kubernetes/kubectl/blob/master/pkg/cmd/patch/patch.go Latest commit 04f62ff
func patch(obj runtime.Object, patchBytes []byte, patchType types.PatchType) (runtime.Object, error) {
	if patchType == "" {
		patchType = types.MergePatchType
	}
	// 	var patchBytes []byte
	// if len(o.PatchFile) > 0 {
	// 	var err error
	// 	patchBytes, err = ioutil.ReadFile(o.PatchFile)
	// 	if err != nil {
	// 		return fmt.Errorf("unable to read patch file: %v", err)
	// 	}
	// } else {
	// 	patchBytes = []byte(o.Patch)
	// }

	patchBytes, err := yaml.ToJSON(patchBytes)
	if err != nil {
		return nil, errors.Errorf("unable to parse %q: %v", patchBytes, err)
	}

	originalObjJS, err := runtime.Encode(unstructured.UnstructuredJSONScheme, obj)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	originalPatchedObjJS, err := getPatchedJSON(patchType, originalObjJS, patchBytes, obj.GetObjectKind().GroupVersionKind(), scheme)
	if err != nil {
		return nil, err
	}

	decoder := serializer.NewCodecFactory(scheme).UniversalDeserializer()

	// targetObj, err := runtime.Decode(unstructured.UnstructuredJSONScheme, originalPatchedObjJS)
	targetObj, err := runtime.Decode(decoder, originalPatchedObjJS)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	didPatch := !reflect.DeepEqual(obj, targetObj)
	if !didPatch {
		return nil, errors.Errorf("didn't patch!")
	}
	log.Printf("gvk %v\n", targetObj.GetObjectKind().GroupVersionKind())
	return targetObj, nil
}

func getPatchedJSON(patchType types.PatchType, originalJS, patchJS []byte, gvk schema.GroupVersionKind, creater runtime.ObjectCreater) ([]byte, error) {
	switch patchType {
	case types.JSONPatchType:
		patchObj, err := jsonpatch.DecodePatch(patchJS)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		bytes, err := patchObj.Apply(originalJS)
		// TODO: This is pretty hacky, we need a better structured error from the json-patch
		if err != nil && strings.Contains(err.Error(), "doc is missing key") {
			msg := err.Error()
			ix := strings.Index(msg, "key:")
			key := msg[ix+5:]
			return bytes, errors.Errorf("Object to be patched is missing field (%s)", key)
		}
		return bytes, errors.WithStack(err)

	case types.MergePatchType:
		rs, err := jsonpatch.MergePatch(originalJS, patchJS)
		return rs, errors.WithStack(err)

	case types.StrategicMergePatchType:
		// get a typed object for this GVK if we need to apply a strategic merge patch
		obj, err := creater.New(gvk)
		if err != nil {
			return nil, errors.Errorf("strategic merge patch is not supported for %s locally, try --type merge", gvk.String())
		}
		return strategicpatch.StrategicMergePatch(originalJS, patchJS, obj)

	default:
		// only here as a safety net - go-restful filters content-type
		return nil, errors.Errorf("unknown Content-Type header for patch: %v", patchType)
	}
}
