package invert

import "testing"

func TestInvert(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "test0", args: args{"test"}, want: "tset"},
		{name: "test1", args: args{"MamA"}, want: "AmaM"},
		{name: "test2", args: args{"test_test"}, want: "tset_tset"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Invert(tt.args.s); got != tt.want {
				t.Errorf("Invert() = %v, want %v", got, tt.want)
			}
		})
	}
}
