package main

/**
  * Defines a team and provides some helper methods
  */

import "fmt"

type Team struct {
  id string
  numPlayers int
  players []string // the ids of the playes for this team
  name string
}

// creates a team object given a slice of players ids
func newTeam(ids []string) Team{
  id := UUID()
  return Team{id: id, players: ids, numPlayers: len(ids)}
}

func (t *Team) toString() (result string) {
  result += "Team: " + t.name + "\n"
  for i := 0; i < t.numPlayers; i++ {
    result += fmt.Sprintf("%-10s", t.players[i])
  }
  return
}
