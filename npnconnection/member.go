package npnconnection

import (
	"github.com/kyleu/libnpn/npncore"
	"github.com/gofrs/uuid"
)

// Returns membership details for the provided Channel
func (s *Service) GetOnline(ch string) []uuid.UUID {
	connections, ok := s.channels[ch]
	if !ok {
		connections = newChannel(ch)
	}
	online := make([]uuid.UUID, 0)
	for _, cID := range connections.MemberIDs {
		c, ok := s.connections[cID]
		if ok && c != nil && (!containsUUID(online, c.Profile.UserID)) {
			online = append(online, c.Profile.UserID)
		}
	}

	return online
}

func (s *Service) sendOnlineUpdate(ch string, connID uuid.UUID, userID uuid.UUID, connected bool) error {
	p := OnlineUpdate{UserID: userID, Connected: connected}
	onlineMsg := NewMessage(npncore.KeySystem, "online-update", p)
	return s.WriteChannel(ch, onlineMsg, connID)
}
