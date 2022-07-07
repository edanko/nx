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

func ImportDxf(launch string) {
	files, err := filepath.Glob(filepath.Join("data", "in-conv", "*", "*.dxf"))
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
			f, err := os.Open(file)
			if err != nil {
				return err
			}
			defer f.Close()

			p, err := parser.ProcessMssDxf(f)
			if err != nil {
				return err
			}

			p.Launch = launch
			_, err = profileService.Create(context.Background(), p)
			if err != nil {
				return err
			}

			return nil
		})
	}

	if err := g.Wait(); err != nil {
		log.Fatalln(err)
	}
}
