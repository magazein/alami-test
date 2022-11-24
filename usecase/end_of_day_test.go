package usecase

import (
	"errors"
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/magazein/alami-test/repository"
)

func TestNewEndOfDayUC(t *testing.T) {
	var beforeEodRepo repository.BeforeEodRepoItf
	var afterEodRepo repository.AfterEodRepoItf
	got := NewEndOfDayUC(beforeEodRepo, afterEodRepo)
	if got == nil {
		t.Error("NewEndOfDayUC got nil")
	}
}

func TestEndOfDayUC_Proceed(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type fields struct {
		beforeEodRepo repository.BeforeEodRepoItf
		afterEodRepo  repository.AfterEodRepoItf
		self          EndOfDayUCItf
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "error get before eod data",
			fields: fields{
				beforeEodRepo: func() repository.BeforeEodRepoItf {
					repo := NewMockBeforeEodRepoItf(ctrl)
					repo.EXPECT().Get().Return(nil, errors.New("foo"))
					return repo
				}(),
			},
			wantErr: true,
		},
		{
			name: "invalid data",
			fields: fields{
				beforeEodRepo: func() repository.BeforeEodRepoItf {
					repo := NewMockBeforeEodRepoItf(ctrl)
					repo.EXPECT().Get().Return(
						[][]string{
							{"id", "Nama", "Age", "Balanced", "Previous Balanced", "Average Balanced", "Free Transfer"},
						},
						nil,
					)
					return repo
				}(),
			},
			wantErr: true,
		},
		{
			name: "error update average balance",
			fields: fields{
				beforeEodRepo: func() repository.BeforeEodRepoItf {
					repo := NewMockBeforeEodRepoItf(ctrl)
					repo.EXPECT().Get().Return(
						[][]string{
							{"id", "Nama", "Age", "Balanced", "Previous Balanced", "Average Balanced", "Free Transfer"},
							{"1", "Liam", "36", "197", "164", "193", "2"},
						},
						nil,
					)
					return repo
				}(),
				afterEodRepo: func() repository.AfterEodRepoItf {
					repo := NewMockAfterEodRepoItf(ctrl)
					repo.EXPECT().Replace(
						[][]string{
							{"id", "Nama", "Age", "Balanced", "No 2b Thread-No", "No 3 Thread-No", "Previous Balanced", "Average Balanced", "No 1 Thread-No", "Free Transfer", "No 2a Thread-No"},
							{"1", "Liam", "36", "197", "", "", "164", "193", "", "2", ""},
						},
					)
					return repo
				}(),
				self: func() EndOfDayUCItf {
					uc := NewMockEndOfDayUCItf(ctrl)
					uc.EXPECT().UpdateAvgBalance(gomock.Any(), gomock.Any()).Return(errors.New("foo")).AnyTimes()
					return uc
				}(),
			},
			wantErr: true,
		},
		{
			name: "error update benefit",
			fields: fields{
				beforeEodRepo: func() repository.BeforeEodRepoItf {
					repo := NewMockBeforeEodRepoItf(ctrl)
					repo.EXPECT().Get().Return(
						[][]string{
							{"id", "Nama", "Age", "Balanced", "Previous Balanced", "Average Balanced", "Free Transfer"},
							{"1", "Liam", "36", "197", "164", "193", "2"},
						},
						nil,
					)
					return repo
				}(),
				afterEodRepo: func() repository.AfterEodRepoItf {
					repo := NewMockAfterEodRepoItf(ctrl)
					repo.EXPECT().Replace(gomock.Any())
					return repo
				}(),
				self: func() EndOfDayUCItf {
					uc := NewMockEndOfDayUCItf(ctrl)
					uc.EXPECT().UpdateAvgBalance(gomock.Any(), gomock.Any()).Return(nil).AnyTimes()
					uc.EXPECT().UpdateBenefit(gomock.Any(), gomock.Any()).Return(errors.New("foo")).AnyTimes()
					return uc
				}(),
			},
			wantErr: true,
		},
		{
			name: "error update limited balance",
			fields: fields{
				beforeEodRepo: func() repository.BeforeEodRepoItf {
					repo := NewMockBeforeEodRepoItf(ctrl)
					repo.EXPECT().Get().Return(
						[][]string{
							{"id", "Nama", "Age", "Balanced", "Previous Balanced", "Average Balanced", "Free Transfer"},
							{"1", "Liam", "36", "197", "164", "193", "2"},
						},
						nil,
					)
					return repo
				}(),
				afterEodRepo: func() repository.AfterEodRepoItf {
					repo := NewMockAfterEodRepoItf(ctrl)
					repo.EXPECT().Replace(gomock.Any())
					return repo
				}(),
				self: func() EndOfDayUCItf {
					uc := NewMockEndOfDayUCItf(ctrl)
					uc.EXPECT().UpdateAvgBalance(gomock.Any(), gomock.Any()).Return(nil).AnyTimes()
					uc.EXPECT().UpdateBenefit(gomock.Any(), gomock.Any()).Return(nil).AnyTimes()
					uc.EXPECT().UpdateLimitedBalance(gomock.Any(), gomock.Any()).Return(errors.New("foo")).AnyTimes()
					return uc
				}(),
			},
			wantErr: true,
		},
		{
			name: "error write csv",
			fields: fields{
				beforeEodRepo: func() repository.BeforeEodRepoItf {
					repo := NewMockBeforeEodRepoItf(ctrl)
					repo.EXPECT().Get().Return(
						[][]string{
							{"id", "Nama", "Age", "Balanced", "Previous Balanced", "Average Balanced", "Free Transfer"},
							{"1", "Liam", "36", "197", "164", "193", "2"},
						},
						nil,
					)
					return repo
				}(),
				afterEodRepo: func() repository.AfterEodRepoItf {
					repo := NewMockAfterEodRepoItf(ctrl)
					repo.EXPECT().Replace(gomock.Any())
					repo.EXPECT().WriteCSV().Return(errors.New("foo"))
					return repo
				}(),
				self: func() EndOfDayUCItf {
					uc := NewMockEndOfDayUCItf(ctrl)
					uc.EXPECT().UpdateAvgBalance(gomock.Any(), gomock.Any()).Return(nil).AnyTimes()
					uc.EXPECT().UpdateBenefit(gomock.Any(), gomock.Any()).Return(nil).AnyTimes()
					uc.EXPECT().UpdateLimitedBalance(gomock.Any(), gomock.Any()).Return(nil).AnyTimes()
					return uc
				}(),
			},
			wantErr: true,
		},
		{
			name: "success",
			fields: fields{
				beforeEodRepo: func() repository.BeforeEodRepoItf {
					repo := NewMockBeforeEodRepoItf(ctrl)
					repo.EXPECT().Get().Return(
						[][]string{
							{"id", "Nama", "Age", "Balanced", "Previous Balanced", "Average Balanced", "Free Transfer"},
							{"1", "Liam", "36", "197", "164", "193", "2"},
						},
						nil,
					)
					return repo
				}(),
				afterEodRepo: func() repository.AfterEodRepoItf {
					repo := NewMockAfterEodRepoItf(ctrl)
					repo.EXPECT().Replace(gomock.Any())
					repo.EXPECT().WriteCSV().Return(nil)
					return repo
				}(),
				self: func() EndOfDayUCItf {
					uc := NewMockEndOfDayUCItf(ctrl)
					uc.EXPECT().UpdateAvgBalance(gomock.Any(), gomock.Any()).Return(nil).AnyTimes()
					uc.EXPECT().UpdateBenefit(gomock.Any(), gomock.Any()).Return(nil).AnyTimes()
					uc.EXPECT().UpdateLimitedBalance(gomock.Any(), gomock.Any()).Return(nil).AnyTimes()
					return uc
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
				self:          tt.fields.self,
			}
			if err := uc.Proceed(); (err != nil) != tt.wantErr {
				t.Errorf("EndOfDayUC.Proceed() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
