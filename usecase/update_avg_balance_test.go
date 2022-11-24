package usecase

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/magazein/alami-test/repository"
)

func TestEndOfDayUC_UpdateAvgBalance(t *testing.T) {
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
			name: "error read balance",
			args: args{
				data: []string{"", "", "", "-"},
			},
			wantErr: true,
		},
		{
			name: "error read previous balance",
			args: args{
				data: []string{"", "", "", "10", "-"},
			},
			wantErr: true,
		},
		{
			name: "error read id",
			args: args{
				data: []string{"-", "", "", "10", "20"},
			},
			wantErr: true,
		},
		{
			name: "error get after eod data",
			args: args{
				data: []string{"1", "", "", "10", "20"},
			},
			fields: fields{
				afterEodRepo: func() repository.AfterEodRepoItf {
					repo := NewMockAfterEodRepoItf(ctrl)
					repo.EXPECT().Find(1).Return(nil, errors.New("foo"))
					return repo
				}(),
			},
			wantErr: true,
		},
		{
			name: "error update after eod data",
			args: args{
				data: []string{"1", "", "", "10", "20"},
				i:    5,
			},
			fields: fields{
				afterEodRepo: func() repository.AfterEodRepoItf {
					repo := NewMockAfterEodRepoItf(ctrl)
					repo.EXPECT().Find(1).Return(
						[]string{"0", "", "", "10", "20", "", "", "", ""},
						nil,
					)
					repo.EXPECT().Update(1, []string{"0", "", "", "10", "20", "", "", "15", "5"}).Return(errors.New("foo"))
					return repo
				}(),
			},
			wantErr: true,
		},
		{
			name: "success",
			args: args{
				data: []string{"1", "", "", "10", "20"},
				i:    5,
			},
			fields: fields{
				afterEodRepo: func() repository.AfterEodRepoItf {
					repo := NewMockAfterEodRepoItf(ctrl)
					repo.EXPECT().Find(1).Return(
						[]string{"0", "", "", "10", "20", "", "", "", ""},
						nil,
					)
					repo.EXPECT().Update(1, []string{"0", "", "", "10", "20", "", "", "15", "5"}).Return(nil)
					return repo
				}(),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := &EndOfDayUC{
				beforeEodRepo: tt.fields.beforeEodRepo,
				afterEodRepo:  tt.fields.afterEodRepo,
			}
			if err := uc.UpdateAvgBalance(tt.args.i, tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("EndOfDayUC.UpdateAvgBalance() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
