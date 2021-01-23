package npnconnection

import (
	"encoding/json"
	"sync"

	"emperror.dev/errors"

	"github.com/kyleu/libnpn/npncore"
	"github.com/kyleu/libnpn/npnuser"

	"github.com/gofrs/uuid"
	"github.com/gorilla/websocket"
)

// Represents a user's WebSocket session
type Connection struct {
	ID       uuid.UUID
	Profile  *npnuser.Profile
	Channels []string
	socket   *websocket.Conn
	mu       sync.Mutex
}

// Creates a new Connection
func NewConnection(svc string, profile *npnuser.Profile, socket *websocket.Conn) *Connection {
	return &Connection{
		ID:       npncore.UUID(),
		Profile:  profile,
		Channels: nil,
		socket:   socket,
	}
}

// Transforms this Connection to a serializable Status object
func (c *Connection) ToStatus() *Status {
	if c.Channels == nil {
		return &Status{ID: c.ID, UserID: c.Profile.UserID, Username: c.Profile.Name, Channels: nil}
	}
	return &Status{ID: c.ID, UserID: c.Profile.UserID, Username: c.Profile.Name, Channels: c.Channels}
}

// Writes bytes to this Connection, you should probably use a helper method
func (c *Connection) Write(b []byte) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	err := c.socket.WriteMessage(websocket.TextMessage, b)
	return errors.Wrap(err, "unable to write to websocket")
}

// Reads bytes from this Connection, you should probably use a helper method
func (c *Connection) Read() ([]byte, error) {
	_, message, err := c.socket.ReadMessage()
	return message, errors.Wrap(err, "unable to write to websocket")
}

// Closes the backing socket
func (c *Connection) Close() error {
	return c.socket.Close()
}

// Serializable representation of a Connection
type Status struct {
	ID       uuid.UUID `json:"id"`
	UserID   uuid.UUID `json:"userID"`
	Username string    `json:"username"`
	Channels []string  `json:"channels"`
}

// Array helper
type Statuses = []*Status

// Common message struct for passing a service, command and parameter
type Message struct {
	Channel string          `json:"channel,omitempty"`
	Cmd     string          `json:"cmd"`
	Param   json.RawMessage `json:"param"`
}

// Constructor
func NewMessage(ch string, cmd string, param interface{}) *Message {
	return &Message{Channel: ch, Cmd: cmd, Param: json.RawMessage(npncore.ToJSON(param, nil))}
}

// Returns a string in "cmd" format, ignoring the param
func (m *Message) String() string {
	return m.Channel + ":" + m.Cmd
}

// Message for updates of a user's online status
type OnlineUpdate struct {
	UserID    uuid.UUID `json:"userID"`
	Connected bool      `json:"connected"`
}
