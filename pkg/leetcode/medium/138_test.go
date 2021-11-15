package medium

import (
	"reflect"
	"testing"
)

func Test_copyRandomList(t *testing.T) {
	type args struct {
		head *RandomNode
	}

	tests := []struct {
		name string
		args args
		want *RandomNode
	}{
		{
			name: "1",
			args: args{
				head: makeRandomNode([][]int{{7, null}, {13, 0}, {11, 4}, {10, 2}, {1, 0}}),
			},
			want: makeRandomNode([][]int{{7, null}, {13, 0}, {11, 4}, {10, 2}, {1, 0}}),
		},
		{
			name: "2",
			args: args{
				head: makeRandomNode([][]int{{1, 1}, {2, 1}}),
			},
			want: makeRandomNode([][]int{{1, 1}, {2, 1}}),
		},
		{
			name: "3",
			args: args{
				head: makeRandomNode([][]int{{3, null}, {3, 0}, {3, null}}),
			},
			want: makeRandomNode([][]int{{3, null}, {3, 0}, {3, null}}),
		},
		{
			name: "",
			args: args{
				head: makeRandomNode([][]int{}),
			},
			want: makeRandomNode([][]int{}),
		},
	}

	toTest := []func(*RandomNode) *RandomNode{copyRandomList, copyRandomList1}

	for _, f := range toTest {
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				got := f(tt.args.head)

				switch {
				case tt.want == nil && got != nil:
					t.Errorf("copyRandomList() = %v, want %v failed at case 0", got, tt.want)
				case tt.want == nil && got == nil:
					t.Logf("ok")
				case got.Val != tt.want.Val:
					t.Errorf("copyRandomList() = %v, want %v failed at case 1", got, tt.want)

				case got.Next == tt.want.Next && got.Next != nil:
					t.Errorf("copyRandomList() = %v, want %v failed at case 2", got, tt.want)

				case !(got.Next == nil || tt.want.Next == nil) && got.Next.Val != tt.want.Next.Val:
					t.Errorf("copyRandomList() = %v, want %v failed at case 3", got, tt.want)

				case (got.Random != nil && tt.want.Random == nil) || (got.Random == nil && tt.want.Random != nil):
					t.Errorf("copyRandomList() = %v, want %v failed at case 4", got, tt.want)

				case !(got.Random == nil || tt.want.Random == nil) && got.Random.Val != tt.want.Random.Val:
					t.Errorf("copyRandomList() = %v, want %v failed at case 5", got, tt.want)

				default:
				}
			})
		}
	}
}

func Test_makeRandomNode(t *testing.T) {
	n1 := &RandomNode{Val: 1}
	n2 := &RandomNode{Val: 2}
	n1.Next = n2
	n1.Random = n2
	n2.Random = n2

	n21 := &RandomNode{Val: 3}
	n22 := &RandomNode{Val: 3}
	n23 := &RandomNode{Val: 3}
	n21.Next = n22
	n22.Next = n23
	n21.Random = nil
	n22.Random = n21
	n23.Random = nil

	type args struct {
		list [][]int
	}

	tests := []struct {
		name string
		args args
		want *RandomNode
	}{
		{
			name: "1",
			args: args{
				list: [][]int{{1, 1}, {2, 1}},
			},
			want: n1,
		},
		{
			name: "2",
			args: args{
				list: [][]int{{3, null}, {3, 0}, {3, null}},
			},
			want: n21,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := makeRandomNode(tt.args.list); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("makeRandomNode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_randomNodeMaker_makeRandomNode(t *testing.T) {
	n1 := &RandomNode{Val: 1}
	n2 := &RandomNode{Val: 2}
	n1.Next = n2
	n1.Random = n2
	n2.Random = n2

	type fields struct {
		list   [][]int
		set    []*RandomNode
		length int
	}

	type args struct {
		i int
	}

	tests := []struct {
		name   string
		fields fields
		args   args
		want   *RandomNode
	}{
		{
			name: "1",
			fields: fields{
				list:   [][]int{{1, 1}, {2, 1}},
				set:    []*RandomNode{n1, n2},
				length: 2,
			},
			args: args{
				i: 0,
			},
			want: n1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &randomNodeMaker{
				list:   tt.fields.list,
				set:    tt.fields.set,
				length: tt.fields.length,
			}
			if got := r.makeRandomNode(tt.args.i); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("randomNodeMaker.makeRandomNode() = %v, want %v", got, tt.want)
			}
		})
	}
}
