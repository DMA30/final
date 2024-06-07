package player

import (
	"fmt"
	"net"
)

type Player struct {
	Index      int
	Connection net.Conn
	// Symbol string
	Score int
}

func NewPlayer(index int) *Player {
	return &Player{
		Index:      index,
		Connection: nil,
		Score:      0,
	}
}

// устанавливает соединение игрока
func (p *Player) SetConnection(conn net.Conn) {
	p.Connection = conn
}

// возвращаем соединение
func (p *Player) GetConnection() net.Conn {
	return p.Connection
}

// увеличиваем счет
func (p *Player) UpdateScore(score int) {
	p.Score += score
}

// возвращаем счет
func (p Player) GetScore() int {
	return p.Score
}

// выводим счет
func (p Player) PrintScore() {
	fmt.Println("Счёт игрока:", p.GetScore())
}
