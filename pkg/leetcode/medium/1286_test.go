package medium

import (
	"log"
	"reflect"
	"testing"
)

func TestConstructor(t *testing.T) {
	type args struct {
		characters        string
		combinationLength int
	}

	tests := []struct {
		name string
		args args
	}{
		{
			name: "",
			args: args{
				characters:        "abc",
				combinationLength: 2,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Constructor(tt.args.characters, tt.args.combinationLength)

			for {
				n1 := got.Next()
				n2 := got.HasNext()
				log.Printf("%v %v", n1, n2)
				if !n2 {
					break
				}
			}
		})
	}
}

func Test_makePermutation(t *testing.T) {
	type args struct {
		s      string
		ind    int
		length int
		bs     []byte
	}

	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "",
			args: args{
				s:      "abcd",
				ind:    0,
				length: 2,
				bs:     []byte{},
			},
			want: []string{
				"ab",
				"ac",
				"ad",
				"bc",
				"bd",
				"cd",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := makePermutation(tt.args.s, tt.args.ind, tt.args.length, tt.args.bs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("makePermutation() = %s, want %s", got, tt.want)
			}
		})
	}
}
