package main

/**
  * Config settings for the teamCreator program.
  *
  * Using the Configuartion allows for a data file's location to
  * customizable and allows for the format of the file to easily
  * be updated.
  */

const (
    // The file name to read data in from
    fileName string = "data.txt"
    // The directory path for teh file
    filePath string = "/projects/src/github.com/zacbrownband/teamCreator/"
    // The position of the name in data file
    nameField int = 0
    // The position of the sex in the data file
    sexField = 1
    // The position of the age in the data file
    ageField int = 2
    // The postion of the scsore in the data file
    scoreField int = 3
)
