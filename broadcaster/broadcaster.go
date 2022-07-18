package broadcaster

import (
	"encoding/json"
	"github.com/xXNurioXx/event-superintendent/v2/connections"
)

func Broadcast(sender connections.Connection, id string, message string) {
	event, _ := json.Marshal(Event{
		Sender:  sender.Connection.RemoteAddr().String(),
		Id:      id,
		Content: message,
	})
	messageToEmit := string(event) + "\n"

	for _, connection := range connections.Connections {
		connection.Writer.WriteString(messageToEmit) // add line end
		connection.Writer.Flush()
	}
}
