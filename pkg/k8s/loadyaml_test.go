package k8s

import (
	"testing"
)

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
			got, err := load(tt.args.src)
			if err != nil {
				t.Errorf("%v\n", err)
				return
			}
			obj, err := patch(got, patchbyte, "")
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
