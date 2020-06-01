package team

import (
	"github.com/saksham2410/getmega/models"
	"gonum.org/v1/gonum/stat/combin"
)

func MakePair(ch chan models.CanMatch, n int, k int) {
	totalPairs := combin.Combinations(n, k)

	for i := 0; i < len(totalPairs)-1; i++ {
		for j := i + 1; j < len(totalPairs); j++ {
			if checkDuplicate(totalPairs[j], totalPairs[i]) {
				cM := models.CanMatch{
					Team1: totalPairs[i],
					Team2: totalPairs[j],
				}
				ch <- cM
			}
		}
	}
	close(ch)
}

func checkDuplicate(Team1 []int, Team2 []int) bool {
	for _, each1 := range Team1 {
		for _, each2 := range Team2 {
			if each1 == each2 {
				return false
			}
		}
	}
	return true
}
