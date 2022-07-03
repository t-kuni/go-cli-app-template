package system

import "time"

type ITimer interface {
	Now() time.Time
}
