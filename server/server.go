package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"net"
	"strconv"
	"strings"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Ошибка подключения:", err.Error())
		return
	}
	defer listener.Close()

	fmt.Println("Сервер запущен...")
	rand.NewSource(time.Now().UnixNano())

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting:", err.Error())
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)

	for {
		expression, answer := generateExpression()
		fmt.Fprintln(conn, expression)

		response, _ := reader.ReadString('\n')
		response = strings.TrimSpace(response)
		responseInt, _ := strconv.Atoi(response)

		if responseInt == answer {
			fmt.Fprintln(conn, "Правильно!")
		} else {
			fmt.Fprintln(conn, "Ответ не верный. Правильный ответ:", answer)
		}
	}
}

func generateExpression() (string, int) {
	num1 := rand.Intn(100)
	num2 := rand.Intn(100)
	operator := rand.Intn(2)

	var expression string
	var answer int

	switch operator {
	case 0:
		expression = fmt.Sprintf("%d + %d", num1, num2)
		answer = num1 + num2
	case 1:
		expression = fmt.Sprintf("%d - %d", num1, num2)
		answer = num1 - num2
	}

	return expression, answer
}
