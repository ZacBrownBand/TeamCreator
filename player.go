package main

/**
  * Defines a player and provides some helper methods
  */
  
import (
  "strconv"
  "strings"
  "fmt"
)

type Player struct {
  age int64
  id string
  name string
  sex string
  scores []int64
}

// creates a player object based on serialized string
func newPlayer(data string) Player{
  parts := strings.Split(data, " ")

  name := parts[0]
  age, _ := strconv.ParseInt(parts[2], 0, 64);
  var sex string = "Male";
  if parts[1] == "F" {
    sex = "Female"
  }

  var scores []int64
  for _, s := range parts[3:] {
    score, _ := strconv.ParseInt(s, 0, 64)
    scores = append(scores, score)
  }

  id := UUID()
  return Player{name: name, id: id, sex: sex, age: age, scores: scores}
}

// returns a formated table head for the player data
func getPlayerTableHeader() string {
  result := fmt.Sprintf("%-12s%-8s%-8s%8s%8s%8s", "Name", "Age", "Sex", "Endurance", "Speed", "Skill")
  return result + "\n" + strings.Repeat("_", len(result));
}

// builds a string for displaying the players information
func (p *Player) toString() (result string) {
  age := strconv.Itoa(int(p.age))
  result += fmt.Sprintf("%-12s%-8s%-8s", p.name, age, p.sex)

  for _, score := range p.scores {
    result += fmt.Sprintf("%8s", strconv.Itoa(int(score)))
  }

  return
}
