package main

/**
  * Defines a a collection of teams that can co-exsits.
  */

import (
  "fmt"
  "strings"
)

type Scenario struct {
  delta int
  high int
  id string
  low int
  teams []Team
}

// creates a senerio object given the teams
func newScenario(teams []Team) Scenario{
  sortTeams(&teams)

  numTeams := len(teams)
  lowIndex := 0
  highIndex := 0
  curScore := 0
  id := ""
  for i := 0; i < numTeams; i++ {
    id += teams[i].id
    curScore = teams[i].score
    if teams[lowIndex].score > curScore {
      lowIndex = i
    }
    if teams[highIndex].score < curScore {
      highIndex = i
    }
  }

  return Scenario{
    delta: teams[highIndex].score -  teams[lowIndex].score,
    high: teams[highIndex].score,
    id: id,
    teams: teams,
    low: teams[lowIndex].score,
  }
}

// sort the scenarios by the difference between the worst and best teams scores
func sortTeams(t *[]Team) {
  n:= len(*t)
  var iMin int = 0
  for j := 0; j < n-1; j++ {
    iMin = j
    for i := j + 1; i < n; i++ {
      if (*t)[i].id < (*t)[iMin].id {
        iMin = i
      }
    }
    if iMin != j {
      (*t)[j], (*t)[iMin] = (*t)[iMin], (*t)[j]
    }
  }
}

// builds a string for displaying the players information
func (s *Scenario) toString() (result string) {
  lineLength := len(s.teams[0].players) * 14 + 20
  dashLine := strings.Repeat("-", lineLength) + "\n"

  for i, t := range s.teams {
    sum := 0
    result += fmt.Sprintf("| Team %1d| ", i)
    for _, p := range(t.players) {
      sum += p.score
      result += fmt.Sprintf("%-10s |  ", p.name)
    }
    result += "Score: " + fmt.Sprint(sum) + "|\n" + dashLine
  }
  stats := "| Stats | " + fmt.Sprintf("Low:%-7dHigh:%-7dDelta:%-7d", s.low, s.high, s.delta)
  stats += strings.Repeat(" ", (lineLength - len(stats)) - 1) + "|\n"
  result = dashLine + stats + dashLine + result + "\n\n"
  return
}
