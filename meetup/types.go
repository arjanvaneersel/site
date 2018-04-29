package meetup

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

type MeetupTime struct {
	time.Time
}

func (m *MeetupTime) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), "\"")
	if s == "null" {
		return
	}

	i, err := strconv.ParseInt(string(b), 10, 64)
	if err != nil {
		return
	}

	m.Time = time.Unix(i/1000, 0)
	return nil
}

func (m *MeetupTime) MarshalJSON() ([]byte, error) {
	if m.IsZero() {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf("\"%s\"", m.Format(""))), nil
}
