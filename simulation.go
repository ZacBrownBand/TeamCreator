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
  teams := [][]Team{}          // all the team layouts
  numSims := 0                 // the number of team layout, len of teams

  // init the original configufration {{a,b,c},{d,e,f},{h,i,g}}
  initSimulation(&workSpace, playerIds, ppt)
  numSims += addTeams(&workSpace, &teams, ppt)

  for nextOffest(&offsets, ppt) {
      // increment
      numSims += addTeams(&workSpace, &teams, ppt)
  }

  for i := 0; i < numSims; i++ {
    for j := 0; j < ppt; j++ {
      ids := teams[i][j].players
      for _, ids := range(ids) {
        fmt.Printf("%-10s", players[ids].name)
      }
      fmt.Print("\n")
      //fmt.Println(teams[i][j].toString())
    }
    fmt.Print("\n")
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
func addTeams(workSpace *[][]string, teams *[][]Team, ppt int) (numSims int){

  /** up and down **/
  row := []Team{}
  col := []Team{}
  for i := 0; i < ppt; i++ {
    teamRow := []string{}
    teamCol := []string{}
    for j := 0; j < ppt; j++ {
      teamRow = append(teamRow, (*workSpace)[i][j])
      teamCol = append(teamCol, (*workSpace)[j][i])
    }
    // create team, add team to row
    row = append(row, newTeam(teamRow))
    col = append(col, newTeam(teamCol))
  }
  // add the team row to the teams matrix
  *teams = append(*teams, row)
  numSims++

  *teams = append(*teams, col)
  numSims++

  /** Diagonals **/
  for i := 1; i < ppt; i++ {
    diag := []Team{}
    diag2 := []Team{}
    for j := 0; j < ppt; j++ {
      x := 0
      y := j
      teamDiag := []string{}
      teamDiag2 := []string{}
      for k := 0; k < ppt; k++ {
        teamDiag = append(teamDiag, (*workSpace)[x][y])
        teamDiag2 = append(teamDiag2, (*workSpace)[x][ppt - 1 - y])
        x = (x + 1) % ppt
        y = (y + i) % ppt
      }
      // create team, add team to row
      diag = append(diag, newTeam(teamDiag))
      diag2 = append(diag2, newTeam(teamDiag2))
    }
    // add the team row to the teams matrix
    *teams = append(*teams, diag)
    numSims++

    *teams = append(*teams, diag2)
    numSims++
  }

  return
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

/*
func runSimulation(players map[string]Player) {
  var ids []string
  for _, p := range(players) {
    ids = append(ids, p.id)
  }

	p, err:=NewPerm(ids, nil)
	if err!=nil{
		fmt.Println(err)
		return
	}

  for result, _ := p.Next(); err==nil; result, err = p.Next(){
    index := p.Index()
    strs := result.([]string)
    fmt.Printf("%14s", strconv.Itoa(int(index)))
    for i := 0; i < len(strs); i++ {
        person := players[strs[i]]
        fmt.Printf("%-10s", person.name)
    }
    fmt.Print("\n")
	}
}
*/
