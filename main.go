package main

import (
	"fmt"
	"log"

	"github.com/magazein/alami-test/constant"
	"github.com/magazein/alami-test/repository"
	"github.com/magazein/alami-test/usecase"
)

func main() {
	// init repository
	beforeEodRepo := repository.NewBeforeEodRepo(constant.BeforeEodCsvName)
	afterEodRepo := repository.NewAfterEodRepo(constant.AfterEodCsvName)

	// init usecase
	eodUc := usecase.NewEndOfDayUC(beforeEodRepo, afterEodRepo)

	// execute usecase
	err := eodUc.Proceed()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Done")
}
