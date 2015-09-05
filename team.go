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
  sortPlayers(&players)
  
  var score int = 0
  var id string = ""
  for _, p := range(players) {
      score += int(p.score)
      id += p.id
  }

  team := Team{
    id: id,
    numPlayers: len(players),
    players: players,
    score: score,
  }
  return team
}

// sort the scenarios by the difference between the worst and best teams scores
func sortPlayers(p *[]Player) {
  n:= len(*p)
  var iMin int = 0
  for j := 0; j < n-1; j++ {
    iMin = j
    for i := j + 1; i < n; i++ {
      if (*p)[i].id < (*p)[iMin].id {
        iMin = i
      }
    }
    if iMin != j {
      (*p)[j], (*p)[iMin] = (*p)[iMin], (*p)[j]
    }
  }
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
