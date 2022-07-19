package remnant

import (
	"time"

	"github.com/google/uuid"
)

type Remnant struct {
	id        uuid.UUID
	createdAt time.Time
	updatedAt time.Time
	name      string
	quality   string
	typ       string
	length    float64
	width     float64
	thickness float64
}
