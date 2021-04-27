package leetcode

import (
	"testing"
)

func Test_countAndSay(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"1",
			args{
				4,
			},
			"1211",
		},
		{
			"2",
			args{
				1,
			},
			"1",
		},
		{
			"3",
			args{
				2,
			},
			"11",
		},
		{
			"1",
			args{
				3,
			},
			"21",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countAndSay(tt.args.n); got != tt.want {
				t.Errorf("countAndSay() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_countAndSayStr(t *testing.T) {
	type args struct {
		n string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"1",
			args{
				"21",
			},
			"1211",
		},
		{
			"2",
			args{
				"33344133",
			},
			"33241123",
		},
		{
			"3",
			args{
				"1",
			},
			"11",
		},
		{
			"4",
			args{
				"2",
			},
			"12",
		},
	}

	tt := struct {
		name string
		args args
		want string
	}{
		"2",
		args{
			"33344133",
		},
		"33241123",
	}
	t.Run(tt.name, func(t *testing.T) {
		if got := countAndSayStr(tt.args.n); got != tt.want {
			t.Errorf("countAndSayStr() = %v, want %v", got, tt.want)
		}
	})
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countAndSayStr(tt.args.n); got != tt.want {
				t.Errorf("countAndSayStr() = %v, want %v", got, tt.want)
			}
		})
		_ = tt
	}
}
