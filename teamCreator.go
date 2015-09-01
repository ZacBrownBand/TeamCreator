package main

/**
  * This is the main module.
  * It is responsiple for driving the program
  */
  
import (
  "fmt"
  "os"
)

func main() {
  //var x map[string]int
  //x["key"] = 10

  var players map[string]Player = loadData()

  // enter program loop
  for {
    printMenu()
    switch getInput() {
    case 0:
      fmt.Print("Good bye\n")
      os.Exit(0)
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
