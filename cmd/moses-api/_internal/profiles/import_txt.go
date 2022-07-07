package profiles

import (
	"context"
	"log"
	"os"
	"path/filepath"

	"github.com/edanko/moses/internal/parser"
	"github.com/edanko/moses/internal/repo/mongodb"
	"github.com/edanko/moses/internal/service/profile"

	"github.com/joho/godotenv"
	"golang.org/x/sync/errgroup"
)

func ImportTxt(launch string) {
	files, err := filepath.Glob(filepath.Join("data", "in-conv", "*.txt"))
	if err != nil {
		log.Fatalln(err)
	}
	if len(files) == 0 {
		log.Fatalln("No input files!")
	}

	err = godotenv.Load()
	if err != nil {
		log.Fatalln(err)
	}

	db, err := mongodb.NewMongoDB(os.Getenv("MONGODB_URI"), os.Getenv("MONGODB_DATABASE"))
	if err != nil {
		log.Fatalln(err)
	}

	profileCollection := db.Collection("profiles")
	profileRepo := mongodb.NewProfileRepo(profileCollection)
	profileService := profile.NewService(profileRepo)

	g := new(errgroup.Group)

	for _, file := range files {
		file := file
		g.Go(func() error {
			parts, err := parser.ProcessTxt(file)
			if err != nil {
				return err
			}

			ctx := context.Background()
			for _, p := range parts {
				p.Launch = launch
				_, err = profileService.Create(ctx, p)
				if err != nil {
					return err
				}
			}

			return nil
		})
	}

	if err := g.Wait(); err != nil {
		log.Fatalln(err)
	}
}
