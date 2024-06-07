package game

import (
	"net"
)

type Game struct {
	Board         [3][3]string
	Players       []Player
	CurrentPlayer Player
	Watchers      []net.Conn
}
