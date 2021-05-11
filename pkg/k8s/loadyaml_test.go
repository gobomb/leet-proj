package k8s

import (
	"log"
	"testing"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
)

var (
	scheme = runtime.NewScheme()
)

func init() {
	log.SetFlags(log.Lshortfile)
	utilruntime.Must(clientgoscheme.AddToScheme(scheme))
	utilruntime.Must(corev1.AddToScheme(scheme))
}

func Test_load(t *testing.T) {
	type args struct {
		src []byte
	}
	tests := []struct {
		name    string
		args    args
		want    interface{}
		wantErr bool
	}{
		{
			"1",
			args{
				[]byte(`kind: Deployment
apiVersion: apps/v1
spec:
  template:
    spec:
      containers:
      - name: patch
        tty: true `),
			},
			nil,
			true,
		},
	}
	patchbyte := []byte(`kind: Deployment
apiVersion: apps/v1
spec:
  template:
    spec:
      containers:
      - name: patch-bb
        tty: true `)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := load(tt.args.src, scheme)
			if err != nil {
				t.Errorf("%v\n", err)
				return
			}
			obj, err := patch(got, patchbyte, "", scheme)
			if err != nil {
				t.Errorf("%+v\n", err)
				return
			}
			deploy, err := checkdeploy(obj)
			if err != nil {
				t.Errorf("%+v\n", err)
				return
			}
			if deploy.Spec.Template.Spec.Containers[0].Name != "patch-bb" {
				t.Errorf("faild:%v\n", deploy)
			}
			// if (err != nil) != tt.wantErr {
			// 	t.Errorf("load() error = %v, wantErr %v", err, tt.wantErr)
			// 	return
			// }
			// if !reflect.DeepEqual(got, tt.want) {
			// 	t.Errorf("load() = %v, want %v", got, tt.want)
			// }
		})
	}
}

func TestPatchOffline(t *testing.T) {
	patchbyte := []byte(`spec:
  template:
    spec:
      containers:
      - name: patch
        tty: false `)

	got, err := load([]byte(`kind: Deployment
apiVersion: apps/v1
spec:
  template:
    spec:
      containers:
      - name: patch
        image: quay.io/kubernetes_incubator/nfs-provisioner:1.0.9
        ports:
            - name: nfs
              containerPort: 2049
            - name: mountd
              containerPort: 20048
            - name: rpcbind
              containerPort: 111
            - name: rpcbind-udp
              containerPort: 111
              protocol: UDP
        securityContext:
            capabilities:
              add:
                - DAC_READ_SEARCH
                - SYS_RESOURCE
        tty: true `), scheme)

	if err != nil {
		t.Errorf("%+v\n", err)
		return
	}
	deploy, err := checkdeploy(got)
	if err != nil {
		t.Errorf("%v\n", err)
		return
	}
	// log.Printf("before patch %#v", deploy)
	err = PatchOffline(deploy, patchbyte, "", scheme)
	if err != nil {
		t.Errorf("%+v\n", err)
		return
	}
	// log.Printf("after patch %+v", deploy)
	deploy2, err := checkdeploy(got)
	if err != nil {
		t.Errorf("%+v\n", err)
		return
	}
	log.Printf(" patched deploy: %v\n", deploy2.Spec.Template.Spec.Containers[0].Ports)
	log.Printf(" patched deploy's name: %v\n", deploy2.Spec.Template.Spec.Containers[0].Name)

	// if deploy2.Spec.Template.Spec.Containers[0].Name != "patch-bb" {
	// 	t.Errorf("faild:%v\n", deploy)
	// }
}

func Test_getGVK(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			getGVK(scheme)
		})
	}
}
