package normalize

import (
	"context"

	"github.com/edanko/moses/internal/entities"
)

func Normalize(ctx context.Context, p *entities.Profile) {
	if p.Project == "056" {
		if p.Quality[:1] == "A" {
			p.Quality = "D" + p.Quality[1:]
		}
	}
}
