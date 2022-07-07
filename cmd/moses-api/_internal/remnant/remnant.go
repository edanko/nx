package remnant

import (
	"bufio"
	"context"
	"encoding/csv"
	"errors"
	"io"
	"os"
	"strings"

	"github.com/edanko/moses/internal/entities"
	"github.com/edanko/moses/internal/repo/mongodb"
	"github.com/edanko/moses/internal/service/remnant"

	"github.com/joho/godotenv"
	"github.com/jszwec/csvutil"
)

type csvRemnant struct {
	Id          int     `csv:"id"`
	Project     string  `csv:"prj"`
	Section     string  `csv:"sec"`
	Dim         string  `csv:"dim"`
	Quality     string  `csv:"quality"`
	Length      float64 `csv:"rem_len"`
	MarkingText string  `csv:"marking_text"`
	UsedIn      string  `csv:"used_in"`
}

func Load(filename string) (int64, error) {
	err := godotenv.Load()
	if err != nil {
		return 0, err
	}

	db, err := mongodb.NewMongoDB(os.Getenv("MONGODB_URI"), os.Getenv("MONGODB_DATABASE"))
	if err != nil {
		return 0, err
	}

	remnantCollection := db.Collection("remnants")
	remnantRepo := mongodb.NewRemnantRepo(remnantCollection)
	remnantService := remnant.NewService(remnantRepo)

	f, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	csvReader := csv.NewReader(bufio.NewReader(f))
	csvReader.Comma = ';'

	dec, err := csvutil.NewDecoder(csvReader)
	if err != nil {
		return 0, err
	}

	var total int64
	for {
		var r csvRemnant
		if err := dec.Decode(&r); errors.Is(err, io.EOF) {
			break
		} else if err != nil {
			return 0, err
		}

		rem := new(entities.Remnant)
		rem.Project = r.Project
		rem.From = r.Section
		rem.Dim = strings.Replace(r.Dim, "X", "*", 1)

		if !strings.Contains(rem.Dim, ".") {
			rem.Dim += ".0"
		}

		rem.Quality = r.Quality
		rem.Length = r.Length
		rem.Marking = r.MarkingText
		rem.UsedIn = r.UsedIn
		rem.Used = r.UsedIn != ""

		_, err = remnantService.Create(context.Background(), rem)
		if err != nil {
			return 0, err
		}
		total++
	}

	return total, nil
}
