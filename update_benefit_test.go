package main

import (
	"reflect"
	"testing"
)

func Test_updateBenefit(t *testing.T) {
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
			name: "error read row number",
			args: args{
				data: []string{"-", "", "", "10", "20"},
			},
			wantErr: true,
			want:    nil,
		},
		{
			name: "balance > 150",
			args: args{
				data: []string{"0", "", "", "200", "20"},
				target: [][]string{
					{"0", "", "", "200", "", "", "", "", ""},
				},
				i: 8,
			},
			wantErr: false,
			want: [][]string{
				{"0", "", "", "225", "8", "", "", "", ""},
			},
		},
		{
			name: "balance between 100-150",
			args: args{
				data: []string{"0", "", "", "130", "20"},
				target: [][]string{
					{"0", "", "", "130", "", "", "", "", "", "2", ""},
				},
				i: 2,
			},
			wantErr: false,
			want: [][]string{
				{"0", "", "", "130", "", "", "", "", "", "5", "2"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := updateBenefit(tt.args.i, tt.args.data, tt.args.target)
			if (err != nil) != tt.wantErr {
				t.Errorf("updateBenefit() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("updateBenefit() = %v, want %v", got, tt.want)
			}
		})
	}
}
