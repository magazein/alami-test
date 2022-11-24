package helper

import (
	"strconv"
	"testing"
)

func Test_proceedAsync(t *testing.T) {
	callback := func(i int, data []string) error {
		// get row number
		_, err := strconv.Atoi(data[0])
		if err != nil {
			return err
		}

		return nil
	}

	type args struct {
		concurrency int
		callback    func(i int, data []string) error
		source      [][]string
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
				callback: callback,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			errs := ProceedAsync(tt.args.concurrency, tt.args.callback, tt.args.source)
			if len(errs) > 0 != tt.wantErr {
				t.Errorf("proceedAsync() error = %v, wantErr %v", errs, tt.wantErr)
				return
			}
		})
	}
}
