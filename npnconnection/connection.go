package npnconnection

import (
	"github.com/kyleu/libnpn/npncore"
	"github.com/kyleu/libnpn/npnuser"

	"emperror.dev/errors"
	"github.com/gofrs/uuid"
	"github.com/gorilla/websocket"
)

// Registers a new Connection for this Service using the provided npnuser.Profile and websocket.Conn
func (s *Service) Register(profile *npnuser.Profile, c *websocket.Conn) (uuid.UUID, error) {
	conn := &Connection{
		ID:      npncore.UUID(),
		Profile: profile,
		Channels: nil,
		socket:  c,
	}

	s.connectionsMu.Lock()
	defer s.connectionsMu.Unlock()

	s.connections[conn.ID] = conn
	return conn.ID, nil
}

// Removes a Connection from this Service
func (s *Service) Disconnect(connID uuid.UUID) (bool, error) {
	conn, ok := s.connections[connID]
	if !ok {
		return false, errors.New(npncore.IDErrorString(npncore.KeyConnection, connID.String()))
	}
	left := false

	if conn.Channels != nil {
		left = true
		for _, x := range conn.Channels {
			err := s.Leave(connID, x)
			if err != nil {
				return left, errors.Wrap(err, "error leaving channel ["+x+"]")
			}
		}
	}

	s.connectionsMu.Lock()
	defer s.connectionsMu.Unlock()

	delete(s.connections, connID)
	return left, nil
}

func invalidConnection(id uuid.UUID) error {
	return errors.New(npncore.IDErrorString(npncore.KeyConnection, id.String()))
}
