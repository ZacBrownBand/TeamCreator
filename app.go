package main

/**
  * This is the main module. It is responsiple for driving the program.
  * @module app
  */

import (
  "fmt"
  "os"
)

/**
  * The entry point for the application.
  */
func main() {
  // Load the player data from the disk
  var players map[string]Player = loadData()

  // Main application loop
  for {
    _printMenu()

    switch _getInput() {
    case 0: // The user wants to quit
        _quit()
    case 1: // The user wants to see all the player data
        _printPlayers(players)
    case 2: // The user wants to get teams
        // IDEA : This should be instanciating an object and running the sim 
        runSimulation(players)
    }
  }
}

/**
  * Prompts the user for input and returns there input.
  * @returns {Int} i The input from the user
  */
func _getInput() (i int) {
  fmt.Print("> ")
  fmt.Scan(&i)
  return
}

/**
  * Prints the menu.
  * @private
  */
func _printMenu() {
  fmt.Print("1) View all players.\n")
  fmt.Print("2) Run team creator.\n")
  fmt.Print("0) Exit\n")
}

/**
  * Prints all the players and there information.
  * @param {Array.<String, Player>} The players to print.
  */
func _printPlayers(players map[string]Player) {
  fmt.Println(getPlayerTableHeader())
  for _, p := range players {
    fmt.Println(p.toString())
  }
  fmt.Print("\n")
}

/**
  * Quits the application
  */
func _quit() {
  os.Exit(0)
}
