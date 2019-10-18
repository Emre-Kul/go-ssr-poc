package utils

import (
	"net"
	"fmt"
	"bufio"
	"encoding/json"
	"time"
)

var reader *bufio.Reader

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
	reader = bufio.NewReader(con)
	return con
}
func Send(con net.Conn, messageType string, data string) string {
	fmt.Fprintf(con, "{ \"type\": \"" + messageType + "\", \"data\": " + data + " }\f")
	con.SetReadDeadline(time.Now().Add(3 * time.Second))
	message := make([]byte, 4096)
	length, _ := con.Read(message)
	if(length > 0) {
		var ssrResponse SSRResponse
		// fmt.Println(json.Valid(message[:length-1]))
		json.Unmarshal(message[:length-1], &ssrResponse)
		// fmt.Println(string(message[:length-1]))
    	return ssrResponse.Data
	}
	return "<></>"
	
}

func Close(con net.Conn) {
	con.Close()
}