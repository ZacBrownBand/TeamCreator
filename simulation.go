package main

/**
  * Module to do the heavy lifting. Computes all the different
  * combinations.
  * @module simulation
  */

import "fmt"

func runSimulation(players map[string]Player) {
  // person per team should be a square number
  var ppt int = 1
  for (ppt * ppt < len(players)){
    ppt++
  }

  // fill in the empty slots with people with zero scores
  for i := len(players); i < ppt*ppt; i++ {
    uuid := UUID()
    players[uuid] = Player {name: "--", id: uuid, sex: "Male"}
  }

  // get the keys from the map to allow for itteration on the map
  var playerIds []string
  for _, p := range(players) {
    playerIds = append(playerIds, p.id)
  }

  workSpace := [][]string{}    // the layout of players
  offsets := make([]int, ppt)  // for shifting the layout
  sims := []Scenario{}         // all the team layouts

  // init the original configufration {{a,b,c},{d,e,f},{h,i,g}}
  initSimulation(&workSpace, playerIds, ppt)

  run := true
  for run == true {
    // calculate the 2n + 2 posible team senerios
    addTeams(&workSpace, &sims, &players, ppt)
    run = nextOffest(&workSpace, &offsets, ppt)
  }

  run = true
  for run == true {
    // calculate the 2n + 2 posible team senerios
    addTeams(&workSpace, &sims, &players, ppt)
    run = nextOffest2(&workSpace, &offsets, ppt)
  }

  // sort to get the scenarios with the lowest difference
  sortSims(&sims)

  numResults := 3
  top := make([]Scenario, numResults)
  index := 0
  i := 0
  // print top five
  for index < numResults && i < len(sims) {
    if !contains(&top, &(sims[i]), index) {
      top[index] = sims[i]
      index++
    }
    i++
  }

  for _, s := range top {
    fmt.Println(s.toString())
  }
}

func contains(sens *[]Scenario, sen *Scenario, index int) bool {
  for i := 0; i < index; i++ {
    if (*sens)[i].id == (*sen).id {
      return true
    }
  }
  return false
}

// sort the scenarios by the difference between the worst and best teams scores
func sortSims(sims *[]Scenario) {
  n:= len(*sims)
  var iMin int = 0
  for j := 0; j < n-1; j++ {
    iMin = j
    for i := j + 1; i < n; i++ {
      if (*sims)[i].delta < (*sims)[iMin].delta {
        iMin = i
      }
    }
    if iMin != j {
      (*sims)[j], (*sims)[iMin] = (*sims)[iMin], (*sims)[j]
    }
  }
}

func nextOffest(workSpace *[][]string, offsets *[]int, ppt int) bool {
  var complete bool = false
  var index int = ppt - 1
  for complete == false {
    if index == 0 {
      return false
    }
    shiftRow(index, workSpace);
    (*offsets)[index]++
    if (*offsets)[index] >= ppt {
       (*offsets)[index] = 0
       index--
    } else {
      complete = true
    }
  }

  return true
}

func nextOffest2(workSpace *[][]string, offsets *[]int, ppt int) bool {
  var complete bool = false
  var index int = ppt - 1
  for complete == false {
    if index == 0 {
      return false
    }
    shiftCol(index, workSpace);
    (*offsets)[index]++
    if (*offsets)[index] >= ppt {
       (*offsets)[index] = 0
       index--
    } else {
      complete = true
    }
  }

  return true
}

func shiftRow(index int, workSpace *[][]string) {
  var length = len((*workSpace)[index])
  var lastEl = (*workSpace)[index][length-1]

  for i := length - 1; i > 0; i-- {
    (*workSpace)[index][i] = (*workSpace)[index][i-1]
  }

  (*workSpace)[index][0] = lastEl
}

func shiftCol(index int, workSpace *[][]string) {
  var length = len((*workSpace)[index])
  var lastEl = (*workSpace)[length-1][index]

  for i := length - 1; i > 0; i-- {
    (*workSpace)[i][index] = (*workSpace)[i-1][index]
  }

  (*workSpace)[0][index] = lastEl
}

//add all the team combinations for the current orientaion of the workSpace
func addTeams(workSpace *[][]string, teams *[]Scenario, players *map[string]Player, ppt int) {
  /** up and down **/
  row := []Team{}
  col := []Team{}
  for i := 0; i < ppt; i++ {
    teamRow := []Player{}
    teamCol := []Player{}
    for j := 0; j < ppt; j++ {
      teamRow = append(teamRow, (*players)[(*workSpace)[i][j]])
      teamCol = append(teamCol, (*players)[(*workSpace)[j][i]])
    }
    // create team, add team to row
    row = append(row, newTeam(teamRow))
    col = append(col, newTeam(teamCol))
  }
  // add the team row to the teams matrix
  *teams = append(*teams, newScenario(row))
  *teams = append(*teams, newScenario(col))

  /** Diagonals **/
  for i := 1; i < ppt; i++ {
    diag := []Team{}
    diag2 := []Team{}
    for j := 0; j < ppt; j++ {
      x := 0
      y := j
      teamDiag := []Player{}
      teamDiag2 := []Player{}
      for k := 0; k < ppt; k++ {
        teamDiag = append(teamDiag, (*players)[(*workSpace)[x][y]])
        teamDiag2 = append(teamDiag2, (*players)[(*workSpace)[x][ppt - 1 - y]])
        x = (x + 1) % ppt
        y = (y + i) % ppt
      }
      // create team, add team to row
      diag = append(diag, newTeam(teamDiag))
      diag2 = append(diag2, newTeam(teamDiag2))
    }
    // add the team row to the teams matrix
    *teams = append(*teams, newScenario(diag))
    *teams = append(*teams, newScenario(diag2))
  }
}

// put all the people into the workspace in a grid
func initSimulation(workSpace *[][]string, playerIds []string, ppt int) {
  pIndex := 0
  for i := 0; i < ppt; i++ {
    row := []string{}
    for j := 0; j < ppt; j++ {
      row = append(row, playerIds[pIndex])
      pIndex++
    }
    *workSpace = append(*workSpace, row)
  }
}
