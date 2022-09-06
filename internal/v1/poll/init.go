package poll

import "time"

var LastPollTime time.Time

func init() {
	LastPollTime = time.Now().Add(-time.Minute)
}
