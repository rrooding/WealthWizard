package csv

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_ProcessFile(t *testing.T) {
	tests := []struct {
		name         string
		csvString    string
		wantMapSlice []map[string]string
	}{
		{"Valid data", "COL1,COL2,COL3\n1,2,3\n4,5,6\n", []map[string]string{{"COL1": "1", "COL2": "2", "COL3": "3"}, {"COL1": "4", "COL2": "5", "COL3": "6"}}},
		{"Missing header", "COL1,,COL3\n1,2,3\n4,5,6\n", []map[string]string{{"COL1": "1", "1": "2", "COL3": "3"}, {"COL1": "4", "1": "5", "COL3": "6"}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Creating a CSV temp file for testing
			tmpfile, err := ioutil.TempFile("", "test*.csv")
			if err != nil {
				t.Fatal(err)
			}
			defer os.Remove(tmpfile.Name())

			_, err = tmpfile.WriteString(tt.csvString)
			tmpfile.Sync() // Persist data on disk

			filepath := tmpfile.Name()
			writerChannel := make(chan map[string]string)

			go ProcessFile(filepath, writerChannel)

			for _, wantMap := range tt.wantMapSlice {
				record := <-writerChannel
				if diff := cmp.Diff(wantMap, record); diff != "" {
					t.Errorf("ProcessFile() mismatch (-want +got):\n%s", diff)
				}
			}
		})
	}
}

func Test_processLine(t *testing.T) {
	tests := []struct {
		name    string
		headers []string
		data    []string
		want    map[string]string
		wantErr bool
	}{
		{"Correct data", []string{"A", "B"}, []string{"1", "2"}, map[string]string{"A": "1", "B": "2"}, false},
		{"Mismatched data", []string{"A", "B"}, []string{"1"}, nil, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := processLine(tt.headers, tt.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("processLine() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("processLine() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
