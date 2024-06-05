package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

var board [3][3]string
var currentPlayer = "X"

func main() {
	listen, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Ошибка при запуске сервера:", err)
		return
	}
	defer listen.Close()

	fmt.Println("Ожидание подключений...")

	conn, err := listen.Accept()
	if err != nil {
		fmt.Println("Ошибка при установлении соединения:", err)
		return
	}
	defer conn.Close()

	fmt.Println("Подключение установлено!")

	setupBoard()

	for {
		displayBoard()
		if checkWinner() {
			fmt.Printf("Победитель: %s\n", currentPlayer)
			break
		}

		fmt.Printf("Ход игрока %s (строка,столбец): ", currentPlayer)

		reader := bufio.NewReader(conn)
		move, _ := reader.ReadString('\n')
		move = strings.TrimSpace(move)
		doMove(move)

		if currentPlayer == "X" {
			currentPlayer = "O"
		} else {
			currentPlayer = "X"
		}
	}

	displayBoard()
}

func setupBoard() {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			board[i][j] = "_"
		}
	}
}

func displayBoard() {
	for i := 0; i < 3; i++ {
		fmt.Println(board[i])
	}
}

func doMove(move string) {
	pos := strings.Split(move, ",")
	row := pos[0]
	col := pos[1]
	boardRow := int(row[0] - '0')
	boardCol := int(col[0] - '0')

	if board[boardRow][boardCol] == "_" {
		board[boardRow][boardCol] = currentPlayer
	} else {
		fmt.Println("Уже есть значение в этой ячейке. Попробуйте снова.")
	}
}

func checkWinner() bool {
	// Проверяем выигрыш по горизонталям
	for i := 0; i < 3; i++ {
		if board[i][0] != "_" && board[i][0] == board[i][1] && board[i][0] == board[i][2] {
			return true
		}
	}

	// Проверяем выигрыш по вертикалям
	for i := 0; i < 3; i++ {
		if board[0][i] != "_" && board[0][i] == board[1][i] && board[0][i] == board[2][i] {
			return true
		}
	}

	// Проверяем выигрыш по диагоналям
	if board[0][0] != "_" && board[0][0] == board[1][1] && board[0][0] == board[2][2] {
		return true
	}

	if board[0][2] != "_" && board[0][2] == board[1][1] && board[0][2] == board[2][0] {
		return true
	}

	return false
}
