package csv

import (
	"encoding/csv"
	"errors"
	"io"
	"strconv"
	"strings"

	"github.com/edanko/moses/internal/entities"
)

const (
	prj = "056"
)

func ProcessCsv(rd io.Reader) (map[string]*entities.Profile, error) {
	profs := make(map[string]*entities.Profile)

	r := csv.NewReader(rd)

	// skip header
	_, err := r.Read()
	if err != nil {
		return nil, err
	}

	for {
		l, err := r.Read()
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			return nil, err
		}

		p := new(entities.Profile)
		name := strings.Split(l[0], "-")

		p.Project = prj
		p.Section = name[0]
		p.PosNo = name[len(name)-1][:len(name[len(name)-1])-1]

		if p, exists := profs[p.Project+p.Section+p.PosNo]; exists {
			p.Quantity++
			continue
		}

		p.Quantity = 1

		p.Length, err = strconv.ParseFloat(l[3], 64)
		if err != nil {
			return nil, err
		}

		p.L = new(entities.End)
		p.L.Name = l[5]

		params := make(map[string]float64, 10)

		if n := stof(l[6]); n > 0 {
			params["A"] = n
		}
		if n := stof(l[7]); n > 0 {
			params["B"] = n
		}
		if n := stof(l[8]); n > 0 {
			params["C"] = n
		}
		if n := stof(l[9]); n > 0 {
			params["Ks"] = n
		}
		if n := stof(l[10]); n > 0 {
			params["R1"] = n
		}
		if n := stof(l[11]); n > 0 {
			params["R2"] = n
		}
		if n := stof(l[12]); n > 0 {
			params["V1"] = n
		}
		if n := stof(l[13]); n > 0 {
			params["V2"] = n
		}
		if n := stof(l[14]); n > 0 {
			params["V3"] = n
		}
		if n := stof(l[15]); n > 0 {
			params["V4"] = n
		}

		p.L.Params = params
		p.L.WebBevel = bev(l[16])

		p.R = new(entities.End)
		p.R.Name = l[19]

		params = make(map[string]float64, 10)

		if n := stof(l[20]); n > 0 {
			params["A"] = n
		}
		if n := stof(l[21]); n > 0 {
			params["B"] = n
		}
		if n := stof(l[22]); n > 0 {
			params["C"] = n
		}
		if n := stof(l[23]); n > 0 {
			params["Ks"] = n
		}
		if n := stof(l[24]); n > 0 {
			params["R1"] = n
		}
		if n := stof(l[25]); n > 0 {
			params["R2"] = n
		}
		if n := stof(l[26]); n > 0 {
			params["V1"] = n
		}
		if n := stof(l[27]); n > 0 {
			params["V2"] = n
		}
		if n := stof(l[28]); n > 0 {
			params["V3"] = n
		}
		if n := stof(l[29]); n > 0 {
			params["V4"] = n
		}

		p.R.Params = params

		p.R.WebBevel = bev(l[30])

		p.TraceBevel = bev(l[48])

		profs[p.Project+p.Section+p.PosNo] = p
	}
	return profs, err
}

func stof(s string) float64 {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0.0
	}
	return f
}
