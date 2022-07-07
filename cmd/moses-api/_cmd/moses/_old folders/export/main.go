package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/edanko/moses/internal/mongodb"

	"github.com/edanko/moses/internal/service/nest"
	"github.com/edanko/moses/internal/service/profile"
	"github.com/edanko/moses/internal/service/remnant"
	"github.com/edanko/moses/internal/service/spacing"

	"github.com/joho/godotenv"
)

func main() {
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

	nestCollection := db.Collection("nests")
	nestRepo := mongodb.NewNestRepo(nestCollection)
	nestService := nest.NewService(nestRepo, remnantService, profileService, spacingService)

	nests, err := nestService.GetAll(context.Background())
	if err != nil {
		panic(err)
	}

	for _, nest := range nests {
		fmt.Println(nest.ID)
	}
}
