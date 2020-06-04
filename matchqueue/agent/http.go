package agent

import (
	"sync"
	"time"

	"github.com/saksham2410/matchqueue/matcher"
)

type HttpMatchingServer struct {
	Matcher       *matcher.Matcher
	mu            sync.Mutex
	Stats         HttpMatchingServerStats
	lastStats     HttpMatchingServerStats
	lastStatsTime time.Time
}

func NewHttpMatchingServer(maxTime matcher.Time, maxScore matcher.PlayerScore, scoreGroupLen int) *HttpMatchingServer {
	s := &HttpMatchingServer{
		Matcher: matcher.NewMatcher(maxTime, maxScore, scoreGroupLen),
		mu:      sync.Mutex{},
	}
	s.Stats.ServerStartTime = time.Now()
	s.lastStatsTime = s.Stats.ServerStartTime
	return s
}

func (s *HttpMatchingServer) Match(currentTime matcher.Time, count int) {
	s.mu.Lock()
	s.Matcher.Match(currentTime, count)
	s.mu.Unlock()
}

func (s *HttpMatchingServer) Sweep(before matcher.Time) {
	s.mu.Lock()
	s.Matcher.Sweep(before)
	s.mu.Unlock()
}
