package main

import (
	"fmt"
	"net"
	"strconv"
	"strings"
	"sync"

	"github.com/toastsandwich/networking-tic-tac-toe/service"
)

type grid [3][3]int

func NewGrid() *grid {
	g := &grid{
		{-1, -1, -1},
		{-1, -1, -1},
		{-1, -1, -1},
	}
	return g
}

func (g *grid) Check() int {
	diagonal := (g[0][0] == g[1][1] && g[1][1] == g[2][2]) ||
		(g[0][2] == g[1][1] && g[1][1] == g[2][0]) && g[1][1] != -1
	if diagonal {
		return g[1][1]
	}
	for i := 0; i < 3; i++ {
		if g[i][0] == g[i][1] && g[i][1] == g[i][2] && g[i][0] != -1 {
			return g[i][0]
		}
		if g[0][i] == g[1][i] && g[1][i] == g[2][i] && g[0][i] != -1 {
			return g[0][i]
		}
	}
	return -1
}

func (g *grid) Edit(_x, _y string, who int) error {
	x, err := strconv.Atoi(_x)
	if err != nil {
		return err
	}
	y, err := strconv.Atoi(_y)
	if err != nil {
		return err
	}
	g[x][y] = who
	return nil
}

type ConnPair struct {
	X net.Conn
	O net.Conn
}

func NewConnPair(x, o net.Conn) ConnPair {
	return ConnPair{
		X: x,
		O: o,
	}
}

func (cp *ConnPair) Close() (err error) {
	if err = cp.X.Close(); err != nil {
		return err
	}
	if err = cp.O.Close(); err != nil {
		return err
	}
	return nil
}

func (cp *ConnPair) Write(b []byte) (n int, err error) {
	if n, err = cp.X.Write(b); err != nil {
		return n, err
	}
	if n, err = cp.O.Write(b); err != nil {
		return n, err
	}
	return
}

type GameServer struct {
	mu sync.Mutex

	Addr     string
	Services *service.Service
}

func NewGameServer(host, port string, service *service.Service) *GameServer {
	return &GameServer{
		Addr:     net.JoinHostPort(host, port),
		Services: service,
	}
}

func (g *GameServer) Start() error {
	fmt.Println("game server starting")
	ln, err := net.Listen("tcp", g.Addr)
	if err != nil {
		return err
	}
	defer ln.Close()
	var waitConn net.Conn
	for {
		conn, err := ln.Accept()
		if err != nil {
			panic(err)
		}
		g.mu.Lock()
		if waitConn == nil {
			waitConn = conn
			waitConn.Write([]byte("X"))
		} else {
			cp := NewConnPair(waitConn, conn)
			conn.Write([]byte("O"))
			waitConn = nil
			go g.HandleConnPairs(cp)
		}
		g.mu.Unlock()
	}
}

func (g *GameServer) HandleConnPairs(cp ConnPair) {
	defer func() {
		if err := cp.Close(); err != nil {
			panic(err)
		}
	}()
	grid := NewGrid()
	currentPlayer := cp.X // This will be the player who will make the move.

	// read
	// edit
	// check
	// if win
	// send the resps
	// else next play

	buf := make([]byte, 1024) // this is read buffer
	for i := 0; i < 9; i++ {
		// reading the move
		n, err := currentPlayer.Read(buf)
		if err != nil {
			fmt.Println(err)
		}
		data := string(buf[:n])

		data = strings.TrimSpace(data)
		move := strings.Split(data, ",")
		if len(move) != 2 {
			currentPlayer.Write([]byte("malformed move"))
			i--
			continue
		}
		for i, m := range move {
			move[i] = strings.TrimSpace(m)
		}
		if err := grid.Edit(move[0], move[1], i%2); err != nil {
			fmt.Println(err)
			currentPlayer.Write([]byte(err.Error()))
			i--
			continue
		}
		win := grid.Check()
		if win == 0 {
			cp.Write([]byte("winner is X"))
			return
		} else if win == 1 {
			cp.Write([]byte("winner is O"))
			return
		}
		if i%2 == 0 {
			currentPlayer = cp.O
		} else {
			currentPlayer = cp.X
		}
	}
	cp.Write([]byte("game is draw"))
}
