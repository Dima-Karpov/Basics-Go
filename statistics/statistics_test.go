package statistics

import "testing"

func TestAvg(t *testing.T) {
	type args struct {
		nums []float64
	}
	var tests []struct {
		name string
		args args
		want float64
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Avg(tt.args.nums); got != tt.want {
				t.Errorf("Avg() = %v, want %v", got, tt.want)
			}
		})
	}

	nums := []float64{1, 2, 3}
	want := 2.0
	got := Avg(nums)
	if got != want {
		t.Errorf("получено %f, ожидалось %f\n", got, want)
	}
	// ошибка
	nums = []float64{1, 2, 3}
	want = 3.0
	got = Avg(nums)
	if got == want {
		t.Errorf("получено %f, ожидалось %f\n", got, want)
	}
}
