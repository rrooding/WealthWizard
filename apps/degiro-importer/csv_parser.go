package main

import (
  "encoding/csv"
  "errors"
  "fmt"
  "io"
  "os"
)

func processCsvFile(input cliInput, writerChannel chan<- map[string]string) {
  // Open the file
  file, err := os.Open(input.filepath)
  check(err)
  defer file.Close()

  // Define headers and line slice
  var headers, line []string

  // Initialize the CSV reader
  reader := csv.NewReader(file)

  // Reading the first line, where we will find our headers
  headers, err = reader.Read()
  check(err)

  // Iterate over each line
  for {
    line, err = reader.Read()
    if err == io.EOF {
      close(writerChannel)
      break
    } else if err != nil {
      exitGracefully(err)
    }

    record, err := processCsvLine(headers, line)

    if err != nil {
      fmt.Printf("Line: %sError: %s\n", line, err)
      continue
    }

    writerChannel <- record
  }
}

func processCsvLine(headers []string, dataList []string) (map[string]string, error) {
  if len(dataList) != len(headers) {
    return nil, errors.New("Line doesn't match headers format")
  }

  recordMap := make(map[string]string)
  for i, name := range headers {
    recordMap[name] = dataList[i]
  }

  return recordMap, nil
}
