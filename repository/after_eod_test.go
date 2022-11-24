package repository

import (
	"reflect"
	"testing"
)

func TestNewAfterEodRepo(t *testing.T) {
	got := NewAfterEodRepo("test")
	if got == nil {
		t.Error("NewAfterEodRepo got nil")
	}
}

func TestAfterEodRepo_Find(t *testing.T) {
	type fields struct {
		data     [][]string
		filepath string
	}
	type args struct {
		idx int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []string
		wantErr bool
	}{
		{
			name: "negative idx",
			fields: fields{
				data: [][]string{
					{"test"},
				},
			},
			args: args{
				idx: -1,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "idx out of range",
			fields: fields{
				data: [][]string{
					{"test"},
				},
			},
			args: args{
				idx: 2,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "success",
			fields: fields{
				data: [][]string{
					{"test"},
				},
			},
			args: args{
				idx: 0,
			},
			want:    []string{"test"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &AfterEodRepo{
				data:     tt.fields.data,
				filepath: tt.fields.filepath,
			}
			got, err := r.Find(tt.args.idx)
			if (err != nil) != tt.wantErr {
				t.Errorf("AfterEodRepo.Find() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AfterEodRepo.Find() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAfterEodRepo_Update(t *testing.T) {
	type fields struct {
		data     [][]string
		filepath string
	}
	type args struct {
		idx int
		row []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "negative idx",
			fields: fields{
				data: [][]string{
					{"test"},
				},
			},
			args: args{
				idx: -1,
			},
			wantErr: true,
		},
		{
			name: "idx out of range",
			fields: fields{
				data: [][]string{
					{"test"},
				},
			},
			args: args{
				idx: 2,
			},
			wantErr: true,
		},
		{
			name: "success",
			fields: fields{
				data: [][]string{
					{"test"},
				},
			},
			args: args{
				idx: 0,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &AfterEodRepo{
				data:     tt.fields.data,
				filepath: tt.fields.filepath,
			}
			if err := r.Update(tt.args.idx, tt.args.row); (err != nil) != tt.wantErr {
				t.Errorf("AfterEodRepo.Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAfterEodRepo_Replace(t *testing.T) {
	type fields struct {
		data     [][]string
		filepath string
	}
	type args struct {
		data [][]string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "success",
			fields: fields{
				data: [][]string{},
			},
			args: args{
				data: [][]string{
					{"test"},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &AfterEodRepo{
				data:     tt.fields.data,
				filepath: tt.fields.filepath,
			}
			r.Replace(tt.args.data)
		})
	}
}

func TestAfterEodRepo_WriteCSV(t *testing.T) {
	tmpdir := t.TempDir()

	type fields struct {
		data     [][]string
		filepath string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "error create",
			fields: fields{
				filepath: tmpdir + "",
			},
			wantErr: true,
		},
		{
			name: "success",
			fields: fields{
				filepath: tmpdir + "/test.csv",
				data: [][]string{
					{"test"},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &AfterEodRepo{
				data:     tt.fields.data,
				filepath: tt.fields.filepath,
			}
			if err := r.WriteCSV(); (err != nil) != tt.wantErr {
				t.Errorf("AfterEodRepo.WriteCSV() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
