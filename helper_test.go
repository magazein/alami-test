package main

import (
	"io/ioutil"
	"math/rand"
	"os"
	"reflect"
	"testing"
)

func Test_readCsvFile(t *testing.T) {
	tmpdir := t.TempDir()

	type args struct {
		filePath string
	}
	type setup struct {
		isCreateCsv bool
		csvName     string
		content     string
	}
	tests := []struct {
		name    string
		args    args
		setup   setup
		want    [][]string
		wantErr bool
	}{
		{
			name: "error open",
			args: args{
				filePath: tmpdir + "/name.csv",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "error read",
			args: args{
				filePath: tmpdir + "/name.csv",
			},
			setup: setup{
				isCreateCsv: true,
				csvName:     "name.csv",
				content:     "\"test",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "no error",
			args: args{
				filePath: tmpdir + "/name.csv",
			},
			setup: setup{
				isCreateCsv: true,
				csvName:     "name.csv",
				content:     "test",
			},
			want: [][]string{
				{"test"},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setup.isCreateCsv {
				file, err := ioutil.TempFile(tmpdir, tt.setup.csvName)
				if err != nil {
					t.Error(err)
				}

				file.WriteString(tt.setup.content)
				tt.args.filePath = file.Name()
				defer os.Remove(file.Name())
			}

			got, err := readCsvFile(tt.args.filePath)
			if (err != nil) != tt.wantErr {
				t.Errorf("readCsvFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("readCsvFile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_writeCsvFile(t *testing.T) {
	tmpdir := t.TempDir()

	type args struct {
		name string
		data [][]string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "error create",
			args: args{
				name: tmpdir + "",
			},
			wantErr: true,
		},
		{
			name: "success",
			args: args{
				name: tmpdir + "/test.csv",
				data: [][]string{
					{"test"},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := writeCsvFile(tt.args.name, tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("writeCsvFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

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
			addToChan(tt.args.channel, tt.args.data)
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
			if got := average(tt.args.data...); got != tt.want {
				t.Errorf("average() = %v, want %v", got, tt.want)
			}
		})
	}
}

func RandStringRunes(n int) string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
