package utils

import (
	"net"
	"fmt"
	"bufio"
	"encoding/json"
	"strings"
)
type LOL struct {
	Data string
}

type SSRResponse struct {
	Data string
	Type string
}

func Connect() net.Conn {
	con, err := net.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
    	panic(err)
	}
	fmt.Println("Connected!")
	return con
}
func Send(con net.Conn, messageType string, data string) string {
	fmt.Fprintf(con, "{ \"type\": \"" + messageType + "\", \"data\": " + data + " }\f")
	resp, _ := bufio.NewReader(con).ReadString('\x0c')
	var ssrResponse SSRResponse
	json.Unmarshal([]byte(strings.Trim(resp, "\x0c")), &ssrResponse)
    return ssrResponse.Data
}

func Close(con net.Conn) {
	con.Close()
}