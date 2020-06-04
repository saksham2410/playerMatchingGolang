package team

import (
	"sort"

	utils "github.com/saksham2410/getmega/helpers"
	"github.com/saksham2410/getmega/models"
)

func CreateMatches(allPlayers []models.Player, teamSize int, typeofmatch int) []models.Match {
	allMatches := []models.Match{}

	if typeofmatch == 0 {
		matchedPlayersChannel := make(chan models.PossibleTeams)
		go MakePair(matchedPlayersChannel, len(allPlayers), teamSize, typeofmatch)
		for cM := range matchedPlayersChannel {
			newMatch := models.Match{}
			team1Avg := float32(0)
			for _, memberInTeam1 := range cM.Team1 {
				newMatch.Team1.Players = append(newMatch.Team1.Players, allPlayers[memberInTeam1])
				team1Avg += allPlayers[memberInTeam1].Level
			}
			team1Avg = team1Avg / float32(len(newMatch.Team1.Players))
			newMatch.Team1.AverageRank = team1Avg
			allMatches = append(allMatches, newMatch)
		}
	}
	if typeofmatch == 1 {
		matchedPlayersChannel := make(chan models.CanMatch)
		go MakePair(matchedPlayersChannel, len(allPlayers), teamSize, typeofmatch)
		for cM := range matchedPlayersChannel {
			newMatch := models.Match{}

			team1Avg := float32(0)
			for _, memberInTeam1 := range cM.Team1 {
				newMatch.Team1.Players = append(newMatch.Team1.Players, allPlayers[memberInTeam1])
				team1Avg += allPlayers[memberInTeam1].Level
			}
			team1Avg = team1Avg / float32(len(newMatch.Team1.Players))
			newMatch.Team1.AverageRank = team1Avg

			team2Avg := float32(0)
			for _, memberInTeam2 := range cM.Team2 {
				newMatch.Team2.Players = append(newMatch.Team2.Players, allPlayers[memberInTeam2])
				team2Avg += float32(allPlayers[memberInTeam2].Level)
			}
			team2Avg = team2Avg / float32(len(newMatch.Team2.Players))
			newMatch.Team2.AverageRank = team2Avg

			allMatches = append(allMatches, newMatch)
		}
	}
	sortMatchesByScoreDiffrence(allMatches)
	return allMatches
}

func CreateMatchesSingle(allPlayers []models.Player, teamSize int) []models.SingleMatch {
	allMatches := []models.SingleMatch{}
	matchedPlayersChannel := make(chan models.CanMatch)
	go MakePair(matchedPlayersChannel, len(allPlayers), teamSize)

	for cM := range matchedPlayersChannel {
		newMatch := models.SingleMatch{}
		TeamTotal := append(cM.Team1, cM.Team2...)
		team1Avg := float32(0)
		for _, memberInTeam1 := range TeamTotal {
			newMatch.Team1.Players = append(newMatch.Team1.Players, allPlayers[memberInTeam1])
			team1Avg += allPlayers[memberInTeam1].Level
		}
		team1Avg = team1Avg / float32(len(newMatch.Team1.Players))
		newMatch.Team1.AverageRank = team1Avg

		allMatches = append(allMatches, newMatch)
	}
	// sortMatchesByScoreDiffrence(allMatches)
	return allMatches
}

func sortMatchesByScoreDiffrence(allMatches []models.Match) {
	sort.Slice(allMatches, func(i, j int) bool {
		return utils.Abs(allMatches[i].Team1.AverageRank-allMatches[i].Team2.AverageRank) < utils.Abs(allMatches[j].Team1.AverageRank-allMatches[j].Team2.AverageRank)
	})
}
