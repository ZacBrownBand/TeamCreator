package main

/**
  * Defines a team and provides some helper methods.
  */

import (
  "fmt"
  "strconv"
)

type Team struct {
  id string
  name string
  numPlayers int
  players []Player
  score int
}

// creates a team object given a slice of players
func newTeam(players []Player) Team{
  var score int = 0
  for _, p := range(players) {
      score += int(p.score)
  }
  team := Team{
    id: UUID(),
    numPlayers: len(players),
    players: players,
    score: score,
  }
  return team
}

// creates a formatted string sumerizing the team
func (t *Team) toString() (result string) {
  result += "Team: " + t.name + "\n"
  for i := 0; i < t.numPlayers; i++ {
    result += fmt.Sprintf("%-10s", t.players[i].name)
  }
  result += "\nscore: " + fmt.Sprintf("%4s", strconv.Itoa(int(t.score)))
  return
}
