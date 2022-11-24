package repository

import (
	"io/ioutil"
	"os"
	"reflect"
	"testing"
)

func TestNewBeforeEodRepo(t *testing.T) {
	got := NewBeforeEodRepo("test")
	if got == nil {
		t.Error("NewBeforeEodRepo got nil")
	}
}

func TestBeforeEodRepo_Get(t *testing.T) {
	tmpdir := t.TempDir()

	type fields struct {
		filepath string
	}
	type setup struct {
		isCreateCsv bool
		csvName     string
		content     string
	}
	tests := []struct {
		name    string
		fields  fields
		setup   setup
		want    [][]string
		wantErr bool
	}{
		{
			name: "error open",
			fields: fields{
				filepath: tmpdir + "/name.csv",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "error read",
			fields: fields{
				filepath: tmpdir + "/name.csv",
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
			fields: fields{
				filepath: tmpdir + "/name.csv",
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
			r := &BeforeEodRepo{
				filepath: tt.fields.filepath,
			}

			if tt.setup.isCreateCsv {
				file, err := ioutil.TempFile(tmpdir, tt.setup.csvName)
				if err != nil {
					t.Error(err)
				}

				file.WriteString(tt.setup.content)
				r.filepath = file.Name()
				defer os.Remove(file.Name())
			}

			got, err := r.Get()
			if (err != nil) != tt.wantErr {
				t.Errorf("BeforeEodRepo.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BeforeEodRepo.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}
