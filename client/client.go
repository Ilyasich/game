package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Ошибка соеденения:", err.Error())
		return
	}
	defer conn.Close()

	fmt.Println("Подключен к серверу. Введите 'exit', чтобы завершить работу.")
	reader := bufio.NewReader(os.Stdin)

	for {
		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("Server: " + message)

		if strings.Contains(message, "Правильно!") {
			continue
		}

		if strings.Contains(message, "Ответ не верный") {
			continue
		}

		fmt.Print("Ваш ответ: ")
		answer, _ := reader.ReadString('\n')
		answer = strings.TrimSpace(answer)

		if answer == "exit" {
			fmt.Println("Выход из игры.")
			break
		}

		conn.Write([]byte(answer + "\n"))
	}
}
