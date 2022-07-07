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

func ImportGen(launch string) {
	files, err := filepath.Glob(filepath.Join("data", "in-conv", "*.gen"))
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

	// allProfiles := make(map[string]map[string]map[string]*entities.Profile)

	g := new(errgroup.Group)

	for _, file := range files {
		file := file
		g.Go(func() error {
			/* f, err := os.Open(file)
			if err != nil {
				log.Fatalln(err)
			}
			defer f.Close() */

			// TODO: change "file" to io.Reader
			parts, err := parser.ProcessGen(file)
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

		/* if _, exist := allProfiles[dim][quality]; !exist {
			allProfiles[dim] = make(map[string]map[string]*entities.Profile)
		}

		allProfiles[dim][quality] = parts //= append(allProfiles[dim][quality], parts...)
		*/
	}

	if err := g.Wait(); err != nil {
		log.Fatalln(err)
	}

	/* for _, v := range allProfiles {
		for _, profParts := range v {

			n := nester.Nest()
			n.Parts = profParts
			n.Nest()

			filename := path.Join("data", "out", "999 "+n.Bars[0].Section(), "txt", n.TxtFileNameString()+".txt")

			err := utils.WriteStringToFile(filename, n.TxtOutputString())
			if err != nil {
				panic(err)
			}
			fmt.Println("[+]", n.TxtFileNameString()+".txt", "successfully created")
		}
	} */
}
