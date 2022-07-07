package jinfeng

import (
	"context"

	"github.com/edanko/moses/internal/entities"
)

func Normalize(ctx context.Context, p *entities.Profile) {
	p.Quality = "D40"
}
