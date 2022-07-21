package error

import "time"

type Error struct {
	AppName     string
	Operation   string
	Code        int
	CreatedDate time.Time
}
