package golang

import (
	"reflect"
	"testing"
)

var _ Man = &father{}

func TestNewMan(t *testing.T) {
	var man Man
	tests := []struct {
		name string
		want Man
	}{
		{
			name: "1",
			want: man,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewMan(); !reflect.DeepEqual(got, tt.want) {
				got.Talk()
				got.Walk()
				got.Play()
				got.Shop()

				got.(*son).father.Play()
				// t.Errorf("NewMan() = %v, want %v", got, tt.want)
			}
		})
	}
}
