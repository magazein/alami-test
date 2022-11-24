package repository

import (
	"encoding/csv"
	"os"

	"github.com/magazein/alami-test/constant"
)

type AfterEodRepo struct {
	data     [][]string
	filepath string
}

func NewAfterEodRepo(filepath string) *AfterEodRepo {
	return &AfterEodRepo{
		filepath: filepath,
	}
}

func (r *AfterEodRepo) Find(idx int) ([]string, error) {
	if idx < 0 || idx >= len(r.data) {
		return nil, constant.ErrOutOfRange
	}

	return r.data[idx], nil
}

func (r *AfterEodRepo) Update(idx int, row []string) error {
	if idx < 0 || idx >= len(r.data) {
		return constant.ErrOutOfRange
	}

	r.data[idx] = row

	return nil
}

func (r *AfterEodRepo) Replace(data [][]string) {
	r.data = data
}

func (r *AfterEodRepo) WriteCSV() error {
	// create file
	f, err := os.Create(r.filepath)
	if err != nil {
		return err
	}
	defer f.Close()

	// create file writter
	w := csv.NewWriter(f)
	w.Comma = constant.Comma

	// write data to file
	if err = w.WriteAll(r.data); err != nil {
		return err
	}

	return nil
}
