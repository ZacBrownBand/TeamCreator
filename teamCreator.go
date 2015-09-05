package main

/**
  * This is the main module.
  * It is responsiple for driving the program.
  */

import (
  "fmt"
  "os"
)

func main() {
  // load the player data from the disk
  var players map[string]Player = loadData()

  // while the user is not done
  for {
    printMenu()

    switch getInput() {
    case 0:
        quit()
    case 1:
        printPlayers(players)
    case 2:
        runSimulation(players)
    }
  }
}

// prints the options to the user
func printMenu() {
  fmt.Print("1) View all players.\n")
  fmt.Print("2) Run team creator.\n")
  fmt.Print("0) Exit\n")
}

func quit() {
  fmt.Print("Good bye\n")
  os.Exit(0)
}

// prompts and returns a users input
func getInput() (i int) {
  fmt.Print("> ")
  fmt.Scan(&i)
  return
}

// prints all the players and there information
func printPlayers(players map[string]Player) {
  fmt.Println(getPlayerTableHeader())
  for _, p := range players {
    fmt.Println(p.toString())
  }
  fmt.Print("\n")
}
