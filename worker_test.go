package main

import (
	"strconv"
	"testing"
)

func Test_proceedAsync(t *testing.T) {
	callback := func(i int, data []string, target [][]string) ([][]string, error) {
		// get row number
		row, err := strconv.Atoi(data[0])
		if err != nil {
			return nil, err
		}

		target[row][2] = "10"
		return target, nil
	}

	type args struct {
		concurrency int
		callback    func(i int, data []string, target [][]string) ([][]string, error)
		source      [][]string
		target      [][]string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "some error",
			args: args{
				concurrency: 2,
				source: [][]string{
					{"1", "test", "10"},
					{"-", "test", "0"},
				},
				target: [][]string{
					{"id", "name", "balance"},
					{"1", "test", "0"},
					{"2", "test", "0"},
				},
				callback: callback,
			},
			wantErr: true,
		},
		{
			name: "no error",
			args: args{
				concurrency: 2,
				source: [][]string{
					{"1", "test", "10"},
					{"2", "test", "10"},
				},
				target: [][]string{
					{"id", "name", "balance"},
					{"1", "test", "0"},
					{"2", "test", "0"},
				},
				callback: callback,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			errs := proceedAsync(tt.args.concurrency, tt.args.callback, tt.args.source, tt.args.target)
			if len(errs) > 0 != tt.wantErr {
				t.Errorf("proceedAsync() error = %v, wantErr %v", errs, tt.wantErr)
				return
			}
		})
	}
}
