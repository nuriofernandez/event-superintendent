package connections

import (
	"bufio"
	"net"
)

var Connections = make([]Connection, 0) // Create an empty list of Connections

// AddConnection adds the provided connection to the list
func AddConnection(conn net.Conn) Connection {
	newConnection := Connection{
		Connection: conn,
		Reader:     *bufio.NewReader(conn),
		Writer:     *bufio.NewWriter(conn),
	}
	Connections = append(Connections, newConnection)
	return newConnection
}

func RemoveConnection(connection Connection) {
	index := findConnection(connection)

	if index == -1 {
		return // ignore not found
	}

	Connections = append(Connections[:index], Connections[index+1:]...)
}

func findConnection(connection Connection) int {
	for index, element := range Connections {
		if connection.Connection == element.Connection {
			return index
		}
	}
	return -1
}
