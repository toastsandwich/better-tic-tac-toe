package matchmaker

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/toastsandwich/networking-tic-tac-toe/model"
)

type Match struct {
	X, O *model.User // both players
	// G    [3][3]rune // grid Grid will be added later, right now lets add duration to test the button and match
	Duration time.Duration
}

func (m *Match) Start() {
	fmt.Printf("match started between %s and %s.\n", m.X.Username, m.O.Username)
	time.Sleep(m.Duration)
	fmt.Printf("match ended between %s and %s.\n", m.X.Username, m.O.Username)
}

type MatchMaker struct {
	Q       []*model.User
	Matches chan Match
	mu      sync.Mutex
}

func NewMatchMaker() *MatchMaker {
	return &MatchMaker{
		Q:       make([]*model.User, 0),
		Matches: make(chan Match),
		mu:      sync.Mutex{},
	}
}

func (mm *MatchMaker) AddToQ(u *model.User) {
	mm.mu.Lock()
	mm.Q = append(mm.Q, u)
	mm.tryMatch()
	mm.mu.Unlock()
}

func (mm *MatchMaker) tryMatch() {
	if len(mm.Q) >= 2 {
		t := rand.Intn(15) + 10
		match := Match{
			X:        mm.Q[0],
			O:        mm.Q[1],
			Duration: time.Duration(t) * time.Second,
		}
		mm.Q = mm.Q[2:]
		mm.Matches <- match
	}
}

func (mm *MatchMaker) GetMatches() <-chan Match {
	return mm.Matches
}
