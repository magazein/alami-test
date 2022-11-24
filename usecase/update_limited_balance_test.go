package usecase

import (
	"errors"
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/magazein/alami-test/repository"
)

func TestEndOfDayUC_UpdateLimitedBalance(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type fields struct {
		beforeEodRepo repository.BeforeEodRepoItf
		afterEodRepo  repository.AfterEodRepoItf
	}
	type args struct {
		i    int
		data []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "error read id",
			args: args{
				data: []string{"-", "", "", "10", "20"},
			},
			wantErr: true,
		},
		{
			name: "error get after eod data",
			fields: fields{
				afterEodRepo: func() repository.AfterEodRepoItf {
					repo := NewMockAfterEodRepoItf(ctrl)
					repo.EXPECT().Find(1).Return(nil, errors.New("foo"))
					return repo
				}(),
			},
			args: args{
				data: []string{"1", "", "", "200", "20"},
				i:    8,
			},
			wantErr: true,
		},
		{
			name: "error read updated balance",
			fields: fields{
				afterEodRepo: func() repository.AfterEodRepoItf {
					repo := NewMockAfterEodRepoItf(ctrl)
					repo.EXPECT().Find(1).Return(
						[]string{"1", "", "", "-", "", "", "", "", "", "", ""},
						nil,
					)
					return repo
				}(),
			},
			args: args{
				data: []string{"1", "", "", "200", "20"},
				i:    8,
			},
			wantErr: true,
		},
		{
			name: "id <= 100",
			fields: fields{
				afterEodRepo: func() repository.AfterEodRepoItf {
					repo := NewMockAfterEodRepoItf(ctrl)
					repo.EXPECT().Find(1).Return(
						[]string{"1", "", "", "10", "", "", "", "", "", "", ""},
						nil,
					)
					repo.EXPECT().Update(1, []string{"1", "", "", "20", "", "8", "", "", "", "", ""}).Return(nil)
					return repo
				}(),
			},
			args: args{
				data: []string{"1", "", "", "200", "20"},
				i:    8,
			},
			wantErr: false,
		},
		{
			name: "id > 100",
			fields: fields{
				afterEodRepo: func() repository.AfterEodRepoItf {
					repo := NewMockAfterEodRepoItf(ctrl)
					repo.EXPECT().Find(200).Return(
						[]string{"200", "", "", "10", "", "", "", "", "", "", ""},
						nil,
					)
					repo.EXPECT().Update(200, []string{"200", "", "", "10", "", "8", "", "", "", "", ""}).Return(nil)
					return repo
				}(),
			},
			args: args{
				data: []string{"200", "", "", "200", "20"},
				i:    8,
			},
			wantErr: false,
		},
		{
			name: "error update",
			fields: fields{
				afterEodRepo: func() repository.AfterEodRepoItf {
					repo := NewMockAfterEodRepoItf(ctrl)
					repo.EXPECT().Find(200).Return(
						[]string{"200", "", "", "10", "", "", "", "", "", "", ""},
						nil,
					)
					repo.EXPECT().Update(200, []string{"200", "", "", "10", "", "8", "", "", "", "", ""}).Return(errors.New("foo"))
					return repo
				}(),
			},
			args: args{
				data: []string{"200", "", "", "200", "20"},
				i:    8,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := &EndOfDayUC{
				beforeEodRepo: tt.fields.beforeEodRepo,
				afterEodRepo:  tt.fields.afterEodRepo,
			}
			if err := uc.UpdateLimitedBalance(tt.args.i, tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("EndOfDayUC.UpdateLimitedBalance() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
