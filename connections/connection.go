package connections

import (
	"bufio"
	"net"
)

type Connection struct {
	Connection net.Conn
	Reader     bufio.Reader
	Writer     bufio.Writer
}
