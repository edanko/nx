package profiles

import (
	"context"
	"io/fs"
	"log"
	"os"
	"path"
	"strings"

	"github.com/edanko/moses/internal/parser/csv"
	"github.com/edanko/moses/internal/repo/mongodb"
	"github.com/edanko/moses/internal/service/profile"

	"github.com/joho/godotenv"
	"golang.org/x/sync/errgroup"
)

func ImportCsv(launch string) {
	files, err := fs.Glob(os.DirFS(path.Join("data", "in-conv")), "*.csv")
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

			parts, err := csv.ProcessCsv(f)
			if err != nil {
				return err
			}

			lst := strings.ReplaceAll(file, ".csv", ".lst")
			f2, err := os.Open(lst)
			if err != nil {
				return err
			}
			defer f.Close()

			dim, quality, err := csv.DimAndQualityFromLst(f2)
			if err != nil {
				return err
			}

			ctx := context.Background()
			for _, p := range parts {
				p.Launch = launch
				p.Dim = dim
				p.Quality = quality
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
