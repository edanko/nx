package cmd

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/edanko/moses/internal/entities"
	"github.com/edanko/moses/internal/repo/mongodb"
	"github.com/edanko/moses/internal/report"
	"github.com/edanko/moses/internal/service/nest"
	"github.com/edanko/moses/internal/service/nester"
	"github.com/edanko/moses/internal/service/profile"
	"github.com/edanko/moses/internal/service/remnant"
	"github.com/edanko/moses/internal/service/spacing"
	"github.com/edanko/moses/internal/service/stock"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	"golang.org/x/sync/errgroup"
)

var nestsCmd = &cobra.Command{
	Use: "nests (launch)",
	// Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			log.Fatalln("No launch specified")
		}

		err := godotenv.Load()
		if err != nil {
			fmt.Println(err)
		}

		db, err := mongodb.NewMongoDB(os.Getenv("MONGODB_URI"), os.Getenv("MONGODB_DATABASE"))
		if err != nil {
			log.Fatalln(err)
		}

		profileCollection := db.Collection("profiles")
		profileRepo := mongodb.NewProfileRepo(profileCollection)
		profileService := profile.NewService(profileRepo)

		remnantCollection := db.Collection("remnants")
		remnantRepo := mongodb.NewRemnantRepo(remnantCollection)
		remnantService := remnant.NewService(remnantRepo)

		spacingCollection := db.Collection("spacing")
		spacingRepo := mongodb.NewSpacingRepo(spacingCollection)
		spacingService := spacing.NewService(spacingRepo)

		stockCollection := db.Collection("stock")
		stockRepo := mongodb.NewStockRepo(stockCollection)
		stockService := stock.NewService(stockRepo)

		nestCollection := db.Collection("nests")
		nestRepo := mongodb.NewNestRepo(nestCollection)
		nestService := nest.NewService(nestRepo, remnantService, profileService, spacingService)

		_ = stockService
		/* for _, v := range []string{"PP160*10.0", "PP140*7.0","PP140*9.0","PP200*10.0", "PP180*11.0", "PP220*11.0", "PP120*6.5", "PP200*12.0"} {

			s := new(entities.Stock)
			s.Dim = v
			s.Quality = "A40"
			s.Length = 12000

			_, err = stockService.Create(context.Background(), s)
			if err != nil {
				log.Fatalln(err)
			}
		} */

		/*s := new(entities.Spacing)
		s.Dim = "PP140*7.0"
		s.Length = 7
		s.Name = "21"
		s.HasBevel = true
		s.HasScallop = false
		s.Machine = entities.IMG
		_, err = spacingService.Create(context.Background(), s)
		if err != nil {
			log.Fatalln(err)
		}

		/* s = new(entities.Spacing)
		s.Dim = "PP180*11.0"
		s.Length = 10
		s.Name = "27"
		s.HasBevel = false
		s.HasScallop = false
		s.Machine = entities.IMG
		_, err = spacingService.Create(context.Background(), s)
		if err != nil {
			log.Fatalln(err)
		}

		s = new(entities.Spacing)
		s.Dim = "PP200*10.0"
		s.Length = 10
		s.Name = "23"
		s.HasBevel = false
		s.HasScallop = false
		s.Machine = entities.IMG

		_, err = spacingService.Create(context.Background(), s)
		if err != nil {
			log.Fatalln(err)
		}

		s = new(entities.Spacing)
		s.Dim = "PP220*11.0"
		s.Length = 11
		s.Name = "21"
		s.HasBevel = false
		s.HasScallop = false
		s.Machine = entities.IMG

		_, err = spacingService.Create(context.Background(), s)
		if err != nil {
			log.Fatalln(err)
		}

		s = new(entities.Spacing)
		s.Dim = "PP220*11.0"
		s.Length = 11
		s.Name = "23"
		s.HasBevel = false
		s.HasScallop = false
		s.Machine = entities.IMG

		_, err = spacingService.Create(context.Background(), s)
		if err != nil {
			log.Fatalln(err)
		}

		s = new(entities.Spacing)
		s.Dim = "PP120*6.5"
		s.Length = 6.5
		s.Name = "23"
		s.HasBevel = false
		s.HasScallop = false
		s.Machine = entities.IMG

		_, err = spacingService.Create(context.Background(), s)
		if err != nil {
			log.Fatalln(err)
		}

		s = new(entities.Spacing)
		s.Dim = "PP220*11.0"
		s.Length = 11
		s.Name = "27"
		s.HasBevel = false
		s.HasScallop = false
		s.Machine = entities.IMG

		_, err = spacingService.Create(context.Background(), s)
		if err != nil {
			log.Fatalln(err)
		}

		s = new(entities.Spacing)
		s.Dim = "PP180*11.0"
		s.Length = 12
		s.Name = "21"
		s.HasBevel = true
		s.HasScallop = false
		s.Machine = entities.IMG

		_, err = spacingService.Create(context.Background(), s)
		if err != nil {
			log.Fatalln(err)
		}

		s = new(entities.Spacing)
		s.Dim = "PP200*12.0"
		s.Length = 12
		s.Name = "21"
		s.HasBevel = true
		s.HasScallop = false
		s.Machine = entities.IMG

		_, err = spacingService.Create(context.Background(), s)
		if err != nil {
			log.Fatalln(err)
		}

		s = new(entities.Spacing)
		s.Dim = "PP200*12.0"
		s.Length = 12
		s.Name = "27"
		s.HasBevel = false
		s.HasScallop = false
		s.Machine = entities.IMG

		_, err = spacingService.Create(context.Background(), s)
		if err != nil {
			log.Fatalln(err)
		}

		s = new(entities.Spacing)
		s.Dim = "PP220*12.0"
		s.Length = 8
		s.Name = "21"
		s.HasBevel = false
		s.HasScallop = false
		s.Machine = entities.IMG

		_, err = spacingService.Create(context.Background(), s)
		if err != nil {
			log.Fatalln(err)
		}

		s = new(entities.Spacing)
		s.Dim = "PP180*11.0"
		s.Length = 10
		s.Name = "23"
		s.HasBevel = false
		s.HasScallop = false
		s.Machine = entities.IMG

		_, err = spacingService.Create(context.Background(), s)
		if err != nil {
			log.Fatalln(err)
		} */

		nst := nester.NewService(nestService, remnantService, spacingService, stockService)

		profiles, err := profileService.GetAllLaunch(context.Background(), args[0])
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(len(profiles))

		byDim := make(map[string]map[string][]*entities.Profile)

		for _, p := range profiles {

			/* 		if _, exist := byDim[p.Dim]; !exist {
				byDim = make(map[string]map[string][]*entities.Profile)
			} */
			if _, exist := byDim[p.Dim][p.Quality]; !exist {
				byDim[p.Dim] = make(map[string][]*entities.Profile)
			}

			byDim[p.Dim][p.Quality] = append(byDim[p.Dim][p.Quality], p)
		}

		machine := entities.IMG

		g := new(errgroup.Group)

		for _, dims := range byDim {
			for _, ps := range dims {
				ps := ps
				g.Go(func() error {
					nests, err := nst.Nest(context.Background(), machine, ps)
					if err != nil {
						return err
					}

					for i, nest := range nests {

						sp := make([]float64, 0, len(nest.Profiles))

						for _, p := range nest.Profiles {
							nest.ProfilesIds = append(nest.ProfilesIds, p.ID)
							sp = append(sp, p.Spacing...)
						}

						nest.Machine = machine
						nest.Project = nest.Profiles[0].Project
						nest.Launch = args[0]
						nest.NestName = "nest_" + strconv.Itoa(i)
						nest.Spacings = sp

						// save nest
						_, err = nestService.Create(context.Background(), nest)
						if err != nil {
							return err
						}

						// use remnant
						if nest.Bar.RemnantID != nil {
							rem, err := remnantService.GetOne(context.Background(), *nest.Bar.RemnantID)
							if err != nil {
								return err
							}

							rem.Used = true
							rem.UsedIn = nest.NestName

							_, err = remnantService.Update(context.Background(), rem)
							if err != nil {
								return err
							}
						}

						// save remnant
						if nest.HasUsefulRemnant() {
							rem := nest.GetRemnant()

							_, err = remnantService.Create(context.Background(), rem)
							if err != nil {
								return err
							}
						}
					}
					return nil
				})
			}
		}

		if err := g.Wait(); err != nil {
			log.Fatalln(err)
		}

		_ = nestService

		nests, err := nestService.GetAll(context.Background())
		if err != nil {
			fmt.Println(err)
		}

		/* fmt.Println("barlist")
		s, err := report.BarListString(nests)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(s)

		fmt.Println() */

		fmt.Println("nesting list")
		s, err := report.NestingListString(nests)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(s)
	},
}

func init() {
	rootCmd.AddCommand(nestsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// nestCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// nestCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
