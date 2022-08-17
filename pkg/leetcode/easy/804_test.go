package easy

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func Test_uniqueMorseRepresentations(t *testing.T) {
	type args struct {
		words []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				words: []string{"gin", "zen", "gig", "msg"},
			},
			want: 2,
		},
		{
			name: "",
			args: args{
				words: []string{"a"},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		convey.Convey(tt.name, t, convey.StackFail, func() {
			got := uniqueMorseRepresentations(tt.args.words)
			convey.So(got, ShouldEqualDiff, tt.want)
		})
	}
}
