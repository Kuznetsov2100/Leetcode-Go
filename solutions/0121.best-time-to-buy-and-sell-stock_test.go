package solutions

import "testing"

func TestMaxProfit(t *testing.T) {
	type args struct {
		prices []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "test1", args: args{prices: []int{7, 1, 5, 3, 6, 4}}, want: 5},
		{name: "test2", args: args{prices: []int{}}, want: 0},
		{name: "test3", args: args{prices: []int{7, 6, 4, 3, 1}}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxProfit0121(tt.args.prices); got != tt.want {
				t.Errorf("MaxProfit() = %v, want %v", got, tt.want)
			}
		})
	}
}