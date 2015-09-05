package main

/**
  * This module is responsiple for reading data into the
  * aplication and deserializing it.
  */

import (
  "bufio"
  "fmt"
  "os"
  "strings"
)

// load data from the file specified in config and return the collection of
// players
func loadData() map[string]Player {
    // load the info from the file
    data, err := readData()

    if (err != nil) {
      fmt.Println("There was an issue reading the file " +
        fileName +
        ". Please ensure that this file exsists in the following directory: " +
        filePath +
        "\nPlease see the config.go file to configure these settings." )

      os.Exit(0)
    }

    // convet the raw data into typed data
    return processData(data)
}

// open the file and read in the data, storing the data in a collection
// of lines
func readData() ([]string, error){
  root, _ := os.Getwd()
  file, err := os.Open(root + filePath + fileName)

  if err != nil {
    return nil, err
  }

  defer file.Close()

  var lines []string
  scanner := bufio.NewScanner(file)

  for scanner.Scan() {
    lines = append(lines, scanner.Text())
  }

  return lines, scanner.Err()
}

// converts a collection of strings into player objects
func processData(data []string) map[string]Player {
    players := make(map[string]Player)
    for _, d := range data {
      // allows for comments
      if !strings.Contains(d, "//") {
        var newPlayer = newPlayer(d)

        players[newPlayer.id] = newPlayer
      }
    }

    return players
}
