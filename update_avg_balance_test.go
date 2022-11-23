package main

import (
	"reflect"
	"testing"
)

func Test_updateAvgBalance(t *testing.T) {
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
			name: "error read balance",
			args: args{
				data: []string{"", "", "", "-"},
			},
			wantErr: true,
			want:    nil,
		},
		{
			name: "error read previous balance",
			args: args{
				data: []string{"", "", "", "10", "-"},
			},
			wantErr: true,
			want:    nil,
		},
		{
			name: "error read row number",
			args: args{
				data: []string{"-", "", "", "10", "20"},
			},
			wantErr: true,
			want:    nil,
		},
		{
			name: "success",
			args: args{
				data: []string{"0", "", "", "10", "20"},
				target: [][]string{
					{"0", "", "", "10", "20", "", "", "", ""},
				},
				i: 8,
			},
			wantErr: false,
			want: [][]string{
				{"0", "", "", "10", "20", "", "", "15", "8"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := updateAvgBalance(tt.args.i, tt.args.data, tt.args.target)
			if (err != nil) != tt.wantErr {
				t.Errorf("updateAvgBalance() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("updateAvgBalance() = %v, want %v", got, tt.want)
			}
		})
	}
}
