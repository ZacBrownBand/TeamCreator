package main

/**
  * Module to do the heavy lifting. Computes all the different
  * combinations.
  */

import "fmt"

func runSimulation(players map[string]Player) {
  var playerIds []string
  for _, p := range(players) {
    playerIds = append(playerIds, p.id)
  }

  ppt := 3                     // sqr ceil of players
  workSpace := [][]string{}    // the layout of players
  offsets := make([]int, ppt)  // for shitting the layout
  sims := []Scenario{}          // all the team layouts

  // init the original configufration {{a,b,c},{d,e,f},{h,i,g}}
  initSimulation(&workSpace, playerIds, ppt)

  // calculate the 2n + 2 posible team senerios
  addTeams(&workSpace, &sims, &players, ppt)


  /**  UNDER WORK **/

  for nextOffest(&offsets, ppt) {
      // increment
      addTeams(&workSpace, &sims, &players, ppt)
  }

  for i := 0; i < len(sims); i++ {
      fmt.Println(sims[i].toString())
    }
}

func nextOffest(offsets *[]int, ppt int) bool {
  // for i := 1; i < ppt; i++ {
  //     if (*offsets)[i] < ppt {
  //       return true
  //     }
  // }
  return false
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
