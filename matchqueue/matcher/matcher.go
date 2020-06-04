package matcher

import (
	"fmt"
	"math"

	"github.com/wangjia184/sortedset"
)

const (
	TimeGroupCount = 60
)

type PlayerId string
type Time int
type PlayerScore uint
type OnGroupMatchedEventCallback func(group *Group)
type ScoreRadiusFunc func(deltaT Time) PlayerScore

type Group struct {
	Players      []*Player
	removed      []bool
	removedCount int
}

func NewGroup(count int) *Group {
	return &Group{
		Players:      make([]*Player, count),
		removed:      make([]bool, count),
		removedCount: 0,
	}
}

func (g *Group) isEmpty() bool {
	return g.removedCount >= len(g.Players)
}

func (g *Group) softRemove(p *Player) {
	for i, v := range g.Players {
		if v == p && !g.removed[i] {
			g.removed[i] = true
			g.removedCount++
		}
	}
}

func (g *Group) PlayerIds() []PlayerId {
	s := make([]PlayerId, len(g.Players))
	i := 0
	for _, p := range g.Players {
		s[i] = p.Id
		i++
	}
	return s
}

func (g *Group) StandardDeviation() float64 {
	sum := float64(0)
	count := 0
	for _, p := range g.Players {
		sum += float64(p.Score)
		count++
	}
	if count <= 0 {
		return 0
	}
	mean := sum / float64(count)

	sum = 0
	for _, p := range g.Players {
		sum += math.Pow(float64(p.Score)-mean, 2)
	}
	sd := math.Sqrt(sum / float64(count))
	return sd
}

func (g *Group) AverageWaitTime(currentTime Time) float64 {
	sum := 0
	count := 0
	for _, p := range g.Players {
		sum += int(currentTime - p.JoinTime)
		count++
	}
	if count <= 0 {
		return 0
	}
	mean := float64(sum) / float64(count)
	return mean
}


func (m *Matcher) MatchForPlayer(id PlayerId, currentTime Time, count int) error {
	p, ok := m.players[id]
	if !ok {
		return PlayerNotExistsError(id)
	}
	if p.Group != nil {
		return PlayerAlreadyMatchedError(id)
	}
	g := NewGroup(count)
	i := 0
	scoreRadius := m.ScoreRadiusFunc(currentTime - p.JoinTime)
	startTime := Time(m.playerQueue.GetByRank(1, false).Score())
	m.IterPlayerCandidates(p, startTime, currentTime, scoreRadius, func(v interface{}) bool {
		candidate := v.(*Player)
		if candidate.Group == nil {
			g.Players[i] = candidate
			i++
		}
		if i >= count {
			m.groups = append(m.groups, g)
			for _, matchedPlayer := range g.Players {
				m.playerQueue.Remove(string(matchedPlayer.Id))
				m.timeScoreGrid.Del(matchedPlayer.gridX, int(matchedPlayer.Score), matchedPlayer)
				matchedPlayer.Group = g
				m.waitTime.AddItem(m.timeScoreGrid.GetYGroupIndex(int(matchedPlayer.Score)), float64(currentTime-matchedPlayer.JoinTime))
			}
			if m.OnGroupMatchedEventCallback != nil {
				m.OnGroupMatchedEventCallback(g)
			}
			return true
		}
		return false
	})
	return nil
}

func (m *Matcher) getMinTime(currentTime Time) Time {
	fmt.Println(Time(currentTime-Time(180)), "You made a comment here")
	return currentTime - Time(180)
}

func (m *Matcher) AutoRemove(currentTime Time) {
	for _, v := range m.playerQueue.GetByScoreRange(sortedset.SCORE(0), sortedset.SCORE(m.getMinTime(currentTime)), &sortedset.GetByScoreRangeOptions{ExcludeEnd: true}) {
		m.Remove(PlayerId(v.Key()))
	}
}

func (m *Matcher) Match(currentTime Time, count int) {
	m.AutoRemove(currentTime)
	m.waitTime.AddTimeAuto(float64(currentTime))

	for _, v := range m.playerQueue.GetByScoreRange(sortedset.SCORE(0), sortedset.SCORE(currentTime), nil) {
		_ = m.MatchForPlayer(PlayerId(v.Key()), currentTime, count)
	}

	m.waitTime.Merge()
}