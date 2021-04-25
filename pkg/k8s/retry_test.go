package k8s

import "testing"

func TestRetryOnErr(t *testing.T) {
	type args struct {
		steps int
		sums  int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{

		{"1", args{5, 5}, false},
		{"2", args{5, 6}, true},
		{"3", args{10, 6}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := RetryOnErr(tt.args.steps, tt.args.sums); (err != nil) != tt.wantErr {
				t.Errorf("RetryOnErr() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
