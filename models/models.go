package models

type RoomMaker struct {
	//accessing db
}
type CanMatch struct {
	Team1 []int
	Team2 []int
}

type PossibleTeams struct {
	Team []int
}

//Defined 2 teams that will participate
type Match struct {
	Team1 Team
	Team2 Team
}

type SingleMatch struct {
	Team1 Team
}
type Player struct {
	Name  string
	Level float32
}
type Team struct {
	TeamID      int64
	AverageRank float32 //Average rank of the team
	Players     []Player
	minClients int
	maxClients int
}
