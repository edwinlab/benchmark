package benchmark

import (
	"time"
)

type Payment struct {
	Id             int
	SynchronizedAt time.Time
	CreatedAt      time.Time
}
