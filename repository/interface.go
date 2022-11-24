package repository

type AfterEodRepoItf interface {
	Find(idx int) ([]string, error)
	Update(idx int, row []string) error
	Replace(data [][]string)
	WriteCSV() error
}

type BeforeEodRepoItf interface {
	Get() ([][]string, error)
}
