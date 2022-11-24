package repository

import (
	"encoding/csv"
	"os"

	"github.com/magazein/alami-test/constant"
)

type BeforeEodRepo struct {
	filepath string
}

func NewBeforeEodRepo(filepath string) *BeforeEodRepo {
	return &BeforeEodRepo{
		filepath: filepath,
	}
}

func (r *BeforeEodRepo) Get() ([][]string, error) {
	// open file
	f, err := os.Open(r.filepath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	// create file reader
	reader := csv.NewReader(f)
	reader.Comma = constant.Comma

	// read all file data
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	return records, nil
}
