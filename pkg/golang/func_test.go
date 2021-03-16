package golang

import "testing"

func TestBook_Pages(t *testing.T) {
	tests := []struct {
		name string
		b    Book
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.b.Pages(); got != tt.want {
				t.Errorf("Book.Pages() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBook_SetPages(t *testing.T) {
	type args struct {
		pages int
	}
	tests := []struct {
		name string
		b    *Book
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.b.SetPages(tt.args.pages)
		})
	}
}

func Test_call(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"call"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			call()
		})
	}
}
