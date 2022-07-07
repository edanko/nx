package csv

import (
	"log"

	"github.com/edanko/moses/internal/entities"
)

func bev(s string) *entities.Bevel {
	switch s {
	default:
		log.Println("add bevel", s)
		return nil
	}
}
