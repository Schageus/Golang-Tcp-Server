package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Println(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()

	request(conn)
	respond(conn)
}

func request(conn net.Conn) {
	i := 0
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		print(ln)
		if i == 0 {
			sf := strings.Fields(ln)
			url := strings.ToLower(sf[1])
			fmt.Println("\n*** METHOD", sf[0])
			fmt.Println("*** URL", url)
			switch url {
			case "/contact":
				contact(conn)
			case "/info":
				info(conn)
			default:
				respond(conn)
			}
		}
		if ln == "" {
			break
		}
		i++
	}
}

func respond(conn net.Conn) {
	body := `<!DOCTYPE html>
	<html>
		<head>
			<title>Home</title>
		</head>
		<body>
			<strong>Home!</strong>
		</body>
	</html>`

	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}

func contact(conn net.Conn) {
	body := `<!DOCTYPE html>
	<html>
		<head>
			<title>Contact</title>
		</head>
		<body>
			<strong>Contact!</strong>
		</body>
	</html>`

	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}

func info(conn net.Conn) {
	body := `<!DOCTYPE html>
	<html>
		<head>
			<title>Info</title>
		</head>
		<body>
			<strong>Info!</strong>
		</body>
	</html>`

	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}
