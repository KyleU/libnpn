package npnconnection

import (
	"time"

	"github.com/gofrs/uuid"
)

type Channel struct {
	Key          string      `json:"key"`
	MemberIDs    []uuid.UUID `json:"memberIDs"`
	LastUpdate   time.Time   `json:"lastUpdate"`
	MessageCount int         `json:"messageCount"`
}

type channels []*Channel

func newChannel(key string) *Channel {
	return &Channel{Key: key, MemberIDs: []uuid.UUID{}, LastUpdate: time.Now()}
}

// Adds a Connection to this Channel
func (s *Service) Join(connID uuid.UUID, ch string) error {
	conn, ok := s.connections[connID]
	if !ok {
		return invalidConnection(connID)
	}
	if !chanContains(conn.Channels, ch) {
		conn.Channels = append(conn.Channels, ch)
	}

	s.channelsMu.Lock()
	defer s.channelsMu.Unlock()

	curr, ok := s.channels[ch]
	if !ok {
		curr = newChannel(ch)
		s.channels[ch] = curr
	}
	if !containsUUID(curr.MemberIDs, connID) {
		curr.MemberIDs = append(curr.MemberIDs, connID)
	}
	return nil
}

// Removes a Connection from this Channel
func (s *Service) Leave(connID uuid.UUID, ch string) error {
	conn, ok := s.connections[connID]
	if !ok {
		return invalidConnection(connID)
	}
	conn.Channels = chanWithout(conn.Channels, ch)

	s.channelsMu.Lock()
	defer s.channelsMu.Unlock()

	curr, ok := s.channels[ch]
	if !ok {
		curr = newChannel(ch)
	}
	filtered := make([]uuid.UUID, 0)
	for _, i := range curr.MemberIDs {
		if i != connID {
			filtered = append(filtered, i)
		}
	}

	if len(filtered) == 0 {
		delete(s.channels, ch)
		return nil
	}

	if len(filtered) == 0 {
		delete(s.channels, ch)
		return nil
	}
	s.channels[ch].MemberIDs = filtered
	return s.sendOnlineUpdate(ch, conn.ID, conn.Profile.UserID, false)
}

func containsString(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func containsUUID(s []uuid.UUID, e uuid.UUID) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func chanContains(c []string, id string) bool {
	for _, x := range c {
		if x == id {
			return true
		}
	}
	return false
}

func chanWithout(c []string, ch string) []string {
	ret := make([]string, 0, len(c))
	for _, x := range c {
		if x != ch {
			ret = append(ret, x)
		}
	}
	return ret
}
