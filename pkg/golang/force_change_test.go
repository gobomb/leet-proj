package golang

import (
	"testing"
)

func TestSortFloat64Fast(t *testing.T) {
	type args struct {
		a []float64
	}

	tests := []struct {
		name string
		args args
	}{
		{
			name: "",
			args: args{
				a: []float64{},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SortFloat64FastV1(a)
			// log.Panicln(a)
			SortFloat64FastV2(b)
		})
	}
}
