package leetcode

import (
	"reflect"
	"testing"
)

func Test_restoreIpAddresses(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "1",
			args: args{
				s: "101023",
			},
			want: []string{"1.0.10.23", "1.0.102.3", "10.1.0.23", "10.10.2.3", "101.0.2.3"},
		},
		{
			name: "2",
			args: args{
				s: "25525511135",
			},
			want: []string{"255.255.11.135", "255.255.111.35"},
		},
		{
			name: "3",
			args: args{
				s: "0000",
			},
			want: []string{"0.0.0.0"},
		},
		{
			name: "4",
			args: args{
				s: "1111",
			},
			want: []string{"1.1.1.1"},
		},
		{
			name: "5",
			args: args{
				s: "010010",
			},
			want: []string{"0.10.0.10", "0.100.1.0"},
		},
		{
			name: "6",
			args: args{
				s: "101",
			},
			want: []string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := restoreIpAddresses(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("restoreIpAddresses() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_checkIpSep(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1",
			args: args{
				s: "01",
			},
			want: false,
		},
		{
			name: "2",
			args: args{
				s: "101",
			},
			want: true,
		},
		{
			name: "3",
			args: args{
				s: "256",
			},
			want: false,
		},
		{
			name: "4",
			args: args{
				s: "1",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkIpSep(tt.args.s); got != tt.want {
				t.Errorf("checkIpSep() = %v, want %v", got, tt.want)
			}
		})
	}
}
