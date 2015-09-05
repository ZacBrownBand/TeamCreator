package main

/**
  * Defines a player and provides some helper methods.
  */

import (
  "strconv"
  "strings"
  "fmt"
)

type Player struct {
  age int
  id string
  name string
  score int
  sex string
}

// creates a player object based on serialized string
func newPlayer(data string) Player{
  parts := strings.Split(data, " ")

  age, _ := strconv.ParseInt(parts[ageField], 0, 64);
  score, _ := strconv.ParseInt(parts[scoreField], 0, 64)

  return Player {
    age: int(age),
    id: UUID(),
    name: parts[nameField],
    score: int(score),
    sex: parts[sexField],
  }
}

// returns a formated table head for the player data
func getPlayerTableHeader() string {
  result := fmt.Sprintf("%-12s%-8s%-8s%8s", "Name", "Age", "Sex", "Skill")
  return result + "\n" + strings.Repeat("_", len(result));
}

// builds a string for displaying the players information
func (p *Player) toString() (result string) {
  age := strconv.Itoa(int(p.age))
  result += fmt.Sprintf("%-12s%-8s%-8s%8s", p.name, age, p.sex, strconv.Itoa(int(p.score)))
  return
}
