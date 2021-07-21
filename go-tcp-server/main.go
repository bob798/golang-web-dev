package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func main() {
	list, err := net.Listen("tcp", ":8090")
	if err != nil {
		fmt.Println(err)

	}
	defer list.Close()

	for {
		conn, err := list.Accept()
		if err != nil {
			fmt.Println(err)
			return

		}

		defer conn.Close()

		go handle(conn)

	}
}

func handle(conn net.Conn) {
	m, u := request(conn)
	if len(m) < 1 && len(u) < 1 {
		return
	}

	if m == "GET" && u == "/hello" {
		respond(conn, "Hello Bob")
		return
	}
	respond(conn, fmt.Sprintf("m=%s,u=%s", m, u))
}

func request(conn net.Conn) (string, string) {
	//buf := make([]byte, 0, 4096)
	//len := 0
	//
	//for {
	//	n, err := conn.Read(buf[len:])
	//	if n > 0 {
	//		len += n
	//	}
	//	if err != nil {
	//		if err != io.EOF {
	//			//Error Handler
	//		}
	//
	//		break
	//	}
	//}

	i := 0
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		//输出内容
		//GET / HTTP/1.1
		if i == 0 {
			// request line
			m := strings.Fields(ln)[0] // method
			u := strings.Fields(ln)[1] // uri
			fmt.Println("***METHOD", m)
			fmt.Println("***URI", u)
			return m, u
		}
		i++
	}
	return "", ""
}

func respond(conn net.Conn, body string) {

	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}
