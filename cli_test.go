package main

import "testing"

func Test_contains(t *testing.T) {
	type args struct {
		str     string
		substrs []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "",
			args: args{
				str:     "abc def ghi",
				substrs: []string{"abc"},
			},
			want: true,
		},
		{
			name: "",
			args: args{
				str:     "abc def ghi",
				substrs: []string{"ab", "de", "gh"},
			},
			want: true,
		},
		{
			name: "",
			args: args{
				str:     "abc def ghi",
				substrs: []string{"abcdefghi"},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := contains(tt.args.str, tt.args.substrs); got != tt.want {
				t.Errorf("contains() = %v, want %v", got, tt.want)
			}
		})
	}
}
