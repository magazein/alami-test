package helper

import (
	"testing"
)

func Test_addToChan(t *testing.T) {
	type args struct {
		channel chan []string
		data    [][]string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "success",
			args: args{
				channel: make(chan []string, 1),
				data: [][]string{
					{"test"},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			AddToChan(tt.args.channel, tt.args.data)
		})
	}
}

func Test_average(t *testing.T) {
	type args struct {
		data []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "success",
			args: args{
				data: []int{10, 30},
			},
			want: 20,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Average(tt.args.data...); got != tt.want {
				t.Errorf("average() = %v, want %v", got, tt.want)
			}
		})
	}
}
