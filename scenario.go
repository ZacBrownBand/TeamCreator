package main

/**
  * Defines a a collection of teams that can co-exsits.
  */

import (
  "strconv"
  "fmt"
)

type Scenario struct {
  average int
  high int
  id string
  low int
  teams []Team
}

// creates a senerio object given the teams
func newScenario(teams []Team) Scenario{
  return Scenario{
    id: UUID(),
    teams: teams,
  }
}

// builds a string for displaying the players information
func (s *Scenario) toString() (result string) {
  for i, t := range s.teams {
    result += "Team "  + strconv.Itoa(i) + "\n"
    for _, p := range(t.players) {
      result += fmt.Sprintf("%-10s", p.name)
    }
    result += "\n"
  }
  low := strconv.Itoa(int(s.low))
  high := strconv.Itoa(int(s.high))
  average := strconv.Itoa(int(s.average))
  result += fmt.Sprintf("Low:%-8sHigh:%-8sAve:%-8s", low, high, average)
  result += "\n"
  return
}
