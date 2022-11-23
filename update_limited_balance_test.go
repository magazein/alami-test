package main

import (
	"reflect"
	"testing"
)

func Test_updateLimitedBalance(t *testing.T) {
	type args struct {
		i      int
		data   []string
		target [][]string
	}
	tests := []struct {
		name    string
		args    args
		want    [][]string
		wantErr bool
	}{
		{
			name: "error read row number",
			args: args{
				data: []string{"-", "", "", "10", "20"},
			},
			wantErr: true,
			want:    nil,
		},
		{
			name: "error read updated balance",
			args: args{
				data: []string{"0", "", "", "10"},
				target: [][]string{
					{"0", "", "", "-", "", "", "", "", "", "", ""},
				},
			},
			wantErr: true,
			want:    nil,
		},
		{
			name: "success",
			args: args{
				data: []string{"0", "", "", "10", ""},
				target: [][]string{
					{"0", "", "", "10", "", "", "", "", "", "", ""},
				},
				i: 2,
			},
			wantErr: false,
			want: [][]string{
				{"0", "", "", "20", "", "2", "", "", "", "", ""},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := updateLimitedBalance(tt.args.i, tt.args.data, tt.args.target)
			if (err != nil) != tt.wantErr {
				t.Errorf("updateLimitedBalance() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("updateLimitedBalance() = %v, want %v", got, tt.want)
			}
		})
	}
}
