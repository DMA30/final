package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Ошибка при установлении соединения:", err)
		return
	}
	defer conn.Close()

	status, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		fmt.Println("Get response error", err)
		return
	}
	fmt.Println(status)

	fmt.Println("Подключение установлено!")

	fmt.Printf("Введите ход игрока O (строка,столбец): ")
	var x, y int
	_, err = fmt.Scanf("%d,%d\n", &x, &y)
	if err != nil {
		fmt.Println("Не удалось получить ввод:", err)
		return
	}
	conn.Write([]byte(fmt.Sprintf("%d,%d\n", x, y)))
	fmt.Printf("Введите ход игрока Х (строка,столбец): ")
	_, err = fmt.Scanf("%d,%d\n", &x, &y)
	if err != nil {
		fmt.Println("Не удалось получить ввод:", err)
		return
	}
	conn.Write([]byte(fmt.Sprintf("%d,%d\n", x, y)))

}
