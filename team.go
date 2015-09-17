package main

/**
  * Defines a team.
  * @module team
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

/**
  * Creates a team object given a slice of players.
  * @constructs
  * @param   {Array.<Player>} players
  * @returns {Team}
  */
func newTeam(players []Player) Team{
  team := Team{
    id: "",
    numPlayers: len(players),
    players: players,
    score: 0,
  }

  team.sortPlayers()

  for _, p := range(team.players) {
      team.id += p.id
      team.score += int(p.score)
  }

  return team
}

/**
  * Sorts the players on the team by the players ids. This is used
  * as a hash for avoiding using the same scenario more than once.
  */
func (t Team) sortPlayers() {
  p := t.players
  n:= len(p)
  var iMin int = 0
  for j := 0; j < n-1; j++ {
    iMin = j
    for i := j + 1; i < n; i++ {
      if p[i].id < p[iMin].id {
        iMin = i
      }
    }
    if iMin != j {
      p[j], p[iMin] = p[iMin], p[j]
    }
  }
}

/**
  * Creates a formatted string sumerizing the team.
  * @param {String} result
  */
func (t *Team) toString() (result string) {
  result += fmt.Sprintf("Team: %s\n", t.name)
  for i := 0; i < t.numPlayers; i++ {
    result += fmt.Sprintf("%-10s", t.players[i].name)
  }
  result += fmt.Sprintf("\nscore: %4s", strconv.Itoa(int(t.score)))
  return
}
