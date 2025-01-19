package matchmaker

import (
	"sync"

	"golang.org/x/net/websocket"
)

type Match map[*websocket.Conn]*websocket.Conn

func (m *Match) Prepare(first, second *websocket.Conn) {
	(*m)[first] = second
}

func (m *Match) Play() {
}

type MatchMaker struct {
	mu      sync.Mutex
	Queue   []*websocket.Conn
	Matches chan *Match
}

func NewMatchMaker() *MatchMaker {
	return &MatchMaker{
		mu:      sync.Mutex{},
		Queue:   make([]*websocket.Conn, 0),
		Matches: make(chan *Match),
	}
}

func (mm *MatchMaker) tryPair() {
	if len(mm.Queue) >= 2 {
		f := mm.Queue[0]
		s := mm.Queue[1]
		mm.Queue = mm.Queue[2:]
		match := make(Match)
		match.Prepare(f, s)
		mm.Matches <- &match
	}
}

func (mm *MatchMaker) IncomingConn(wc *websocket.Conn) {
	mm.mu.Lock()
	mm.Queue = append(mm.Queue, wc)
	mm.tryPair()
	mm.mu.Unlock()
}
