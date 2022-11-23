package main

import (
	"os"
	"testing"
)

func Test_main(t *testing.T) {
	tests := []struct {
		name     string
		preMain  func()
		postMain func()
	}{
		{
			name: "error read before eod csv",
			preMain: func() {
				os.Rename("BeforeEod.csv", "BeforeEod.csv.bak")
			},
			postMain: func() {
				os.Rename("BeforeEod.csv.bak", "BeforeEod.csv")
			},
		},
		{
			name: "error read after eod csv",
			preMain: func() {
				os.Rename("BeforeEod.csv", "BeforeEod.csv.bak")
				writeCsvFile("BeforeEod.csv", [][]string{
					{"id", "Nama", "Age", "Balanced", "Previous Balanced", "Average Balanced", "Free Transfer"},
					{"1", "Liam", "36", "197", "164", "193", "2"},
					{"2", "Liam", "36", "197", "164", "193", "2"},
					{"3", "Liam", "36", "197", "164", "193", "2"},
				})

				os.Rename("AfterEod.csv", "AfterEod.csv.bak")
			},
			postMain: func() {
				os.Rename("BeforeEod.csv.bak", "BeforeEod.csv")
				os.Rename("AfterEod.csv.bak", "AfterEod.csv")
			},
		},
		{
			name: "len data is not same",
			preMain: func() {
				os.Rename("BeforeEod.csv", "BeforeEod.csv.bak")
				writeCsvFile("BeforeEod.csv", [][]string{
					{"id", "Nama", "Age", "Balanced", "Previous Balanced", "Average Balanced", "Free Transfer"},
					{"1", "Liam", "36", "197", "164", "193", "2"},
					{"2", "Liam", "36", "197", "164", "193", "2"},
					{"3", "Liam", "36", "197", "164", "193", "2"},
				})

				os.Rename("AfterEod.csv", "AfterEod.csv.bak")
				writeCsvFile("AfterEod.csv", [][]string{})
			},
			postMain: func() {
				os.Rename("BeforeEod.csv.bak", "BeforeEod.csv")
				os.Rename("AfterEod.csv.bak", "AfterEod.csv")
			},
		},
		{
			name: "data is just header",
			preMain: func() {
				os.Rename("BeforeEod.csv", "BeforeEod.csv.bak")
				writeCsvFile("BeforeEod.csv", [][]string{
					{"id", "Nama", "Age", "Balanced", "Previous Balanced", "Average Balanced", "Free Transfer"},
				})

				os.Rename("AfterEod.csv", "AfterEod.csv.bak")
				writeCsvFile("AfterEod.csv", [][]string{
					{"id", "Nama", "Age", "Balanced", "No 2b Thread-No", "No 3 Thread-No", "Previous Balanced", "Average Balanced", "No 1 Thread-No", "Free Transfer", "No 2a Thread-No"},
				})
			},
			postMain: func() {
				os.Rename("BeforeEod.csv.bak", "BeforeEod.csv")
				os.Rename("AfterEod.csv.bak", "AfterEod.csv")
			},
		},
		{
			name: "error proceed data",
			preMain: func() {
				os.Rename("BeforeEod.csv", "BeforeEod.csv.bak")
				writeCsvFile("BeforeEod.csv", [][]string{
					{"id", "Nama", "Age", "Balanced", "Previous Balanced", "Average Balanced", "Free Transfer"},
					{"1", "Liam", "36", "197", "164", "193", "2"},
					{"-", "Liam", "36", "197", "164", "193", "2"},
				})

				os.Rename("AfterEod.csv", "AfterEod.csv.bak")
				writeCsvFile("AfterEod.csv", [][]string{
					{"id", "Nama", "Age", "Balanced", "No 2b Thread-No", "No 3 Thread-No", "Previous Balanced", "Average Balanced", "No 1 Thread-No", "Free Transfer", "No 2a Thread-No"},
					{"1", "Liam", "36", "197", "", "", "164", "193", "", "2", ""},
					{"2", "Liam", "36", "197", "", "", "164", "193", "", "2", ""},
				})
			},
			postMain: func() {
				os.Rename("BeforeEod.csv.bak", "BeforeEod.csv")
				os.Rename("AfterEod.csv.bak", "AfterEod.csv")
			},
		},
		{
			name: "error write csv file",
			preMain: func() {
				os.Rename("BeforeEod.csv", "BeforeEod.csv.bak")
				writeCsvFile("BeforeEod.csv", [][]string{
					{"id", "Nama", "Age", "Balanced", "Previous Balanced", "Average Balanced", "Free Transfer"},
					{"1", "Liam", "36", "197", "164", "193", "2"},
					{"2", "Liam", "36", "197", "164", "193", "2"},
				})

				os.Rename("AfterEod.csv", "AfterEod.csv.bak")
				writeCsvFile("AfterEod.csv", [][]string{
					{"id", "Nama", "Age", "Balanced", "No 2b Thread-No", "No 3 Thread-No", "Previous Balanced", "Average Balanced", "No 1 Thread-No", "Free Transfer", "No 2a Thread-No"},
					{"1", "Liam", "36", "197", "", "", "164", "193", "", "2", ""},
					{"2", "Liam", "36", "197", "", "", "164", "193", "", "2", ""},
				})

				os.Chmod("AfterEod.csv", 0444)
			},
			postMain: func() {
				os.Chmod("AfterEod.csv", 0644)
				os.Rename("BeforeEod.csv.bak", "BeforeEod.csv")
				os.Rename("AfterEod.csv.bak", "AfterEod.csv")
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.preMain()
			main()
			tt.postMain()
		})
	}
}
