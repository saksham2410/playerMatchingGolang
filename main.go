package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"

	utils "github.com/saksham2410/getmega/helpers"
	"github.com/saksham2410/getmega/models"
	"github.com/saksham2410/getmega/team"
	"github.com/wangjia184/sortedset"
)

// var playersOnline int
var wg sync.WaitGroup

type Player struct {
	Name  string
	Level float32
	// ReceiveChan chan
}
type Players *sortedset.SortedSet
type Team struct {
	TeamID int
	// ReceiveChan chan
}

// func numPlayers(nm []models.Player) bool {
// 	if len(pl) == 0 {
// 		// w := sync.WaitGroup{}
// 		// w.Add(1)
// 		return false
// 	}
// 	return true
// }

// type allPlayers struct := []models.Player{}
func addPlayer(players *sortedset.SortedSet, online *int, mm *[]models.Player, name string, level float32, c chan int) {
	*online++
	newPlayer := models.Player{
		Name:  name,
		Level: level,
	}
	fmt.Println(*online)
	players.AddOrUpdate(strconv.Itoa(*online), sortedset.SCORE(level), newPlayer)
	*mm = append(*mm, newPlayer)
	fmt.Printf("New Player: %+v \n", newPlayer)
	fmt.Println(*mm)
	// if *online != 1 {
	c <- *online
	fmt.Println("Reached here")
	// }
}

func getAllPlayers(mm []models.Player) {
	for i, player := range mm {
		fmt.Printf("%d %+v \n", i, player)
	}
}

func clearAllPlayers(mm []models.Player) []models.Player {
	mm = nil
	mm = []models.Player{}
	fmt.Println("Cleared all player list")
	return mm
}

func makeMatch(mm []models.Player, size int) {
	if size > len(mm)/2 {
		fmt.Println("Please enter a value from 0 to", len(mm)/2)
		return
	}
	for i, match := range team.CreateMatches(mm, size) {
		rankDiff := utils.Abs(match.Team1.AverageRank - match.Team2.AverageRank)
		fmt.Printf("%d > ", i)
		for _, player := range match.Team1.Players {
			fmt.Printf("%v:%v ", player.Name, player.Level)
		}
		fmt.Print(" -VS- ")
		for _, player := range match.Team2.Players {
			fmt.Printf("%v:%v ", player.Name, player.Level)
		}
		fmt.Printf("-> Diff: %v \n", rankDiff)
	}
}
func makeMatchsingle(mm []models.Player, name string, size int) {

	if size > len(mm)-1 {
		fmt.Println("Please enter a value from 0 to", len(mm)-1)
		return
	}
	// out:
	for _, match := range team.CreateMatches(mm, size) {
		for _, player := range mm {

			if player.Name == name {
				// fmt.Println(player.Name)
				rankDiff1 := utils.Abs(match.Team1.AverageRank - player.Level)
				// fmt.Printf("%d > ", i)
				for _, player1 := range match.Team1.Players {
					// fmt.Println(player1)
					if player.Name == player1.Name {
						continue
					} else {
						fmt.Printf("%v:%v ", player1.Name, player1.Level)
					}
				}
				fmt.Print(" -VS- ")
				fmt.Println(player.Name, player.Level)
				fmt.Printf("-> Diff: %v \n", rankDiff1)
				rankDiff2 := utils.Abs(match.Team2.AverageRank - player.Level)
				// fmt.Printf("%d > ", i)
				for _, player1 := range match.Team2.Players {
					if player.Name == player1.Name {
						continue
					} else {
						fmt.Printf("%v:%v ", player1.Name, player1.Level)
					}
				}
				fmt.Print(" -VS- ")
				fmt.Println(player.Name, player.Level)
				fmt.Printf("-> Diff: %v \n", rankDiff2)
			}
		}

	}
}

func main() {
	playersOnline := 0
	c := make(chan int)
	mm := &[]models.Player{}
	Players := sortedset.New()
	go addPlayer(Players, &playersOnline, mm, "saksham", 5, c)
	fmt.Println(<-c)
	Players.AddOrUpdate("2", 10, "saksham")
	fmt.Println(*Players.GetByKey(strconv.Itoa(1)))
	fmt.Println(*Players.GetByKey(strconv.Itoa(2)))
	fmt.Println(Players.GetByScoreRange(0, 11, &sortedset.GetByScoreRangeOptions{
		Limit: 2,
	}), "These values will be returned")
	// c1 := make(chan int)
	// go addPlayer(&playersOnline, mm, "saksham2", 50, c1)
	// fmt.Println(<-c1)
	// go addPlayer(&playersOnline, mm, "saksham3", 10, c)
	// fmt.Println(<-c)
	// go addPlayer(&playersOnline, mm, "saksham4", 20, c)
	// fmt.Println(<-c)
	getAllPlayers(*mm)
	time.Sleep(time.Second * 2)
}
