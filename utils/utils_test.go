package utils

import "testing"

func TestGetPercentage(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"case1", args{1, 3}, "33.33%"},
		{"case2", args{2, 3}, "66.67%"},
		{"case3", args{3, 3}, "100.00%"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetPercentage(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("GetPercentage() = %v, want %v", got, tt.want)
			}
		})
	}
}
