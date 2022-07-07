package nest

import (
	"time"

	"github.com/google/uuid"
)

type Nest struct {
	id        uuid.UUID
	createdAt time.Time
	updatedAt time.Time
	name      string
	length    float64
}
