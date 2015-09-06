package main

/**
  * Defines a player.
  * @module player
  */

import (
  "fmt"
  "strconv"
  "strings"
)

type Player struct {
  age int
  id string
  name string
  score int
  sex string
}

/**
  * Creates a player object based on serialized string.
  * @constructs
  * @param   {String} data A serialised string to use to create the player
  * @returns {Player}
  */
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

/**
  * Creates a formated table head for the player data.
  * @returns {String} A string that can be used as a header
  *                   for a table of players
  * @static
  */
func getPlayerTableHeader() string {
  result := fmt.Sprintf("%-12s%-8s%-8s%8s\n", "Name", "Age", "Sex", "Skill")
  return result + strings.Repeat("-", len(result))
}

/**
  * Creates a user freindly string for displaying the players information.
  * @retusn {String} result
  */
func (p *Player) toString() (result string) {
  age := strconv.Itoa(int(p.age))
  result += fmt.Sprintf("%-12s%-8s%-8s%8s", p.name, age, p.sex, strconv.Itoa(int(p.score)))
  return
}
