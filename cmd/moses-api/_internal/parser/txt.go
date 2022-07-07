package parser

import (
	"bytes"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/edanko/moses/internal/entities"
)

func ProcessTxt(fname string) (map[string]*entities.Profile, error) {
	f, err := os.Open(fname)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	b, err := io.ReadAll(f)
	if err != nil {
		return nil, err
	}
	b = bytes.ReplaceAll(b, []byte("\r"), []byte(""))

	fbase := strings.Split(strings.TrimSuffix(filepath.Base(fname), ".txt"), "_")

	ps := strings.Split(fbase[0], "-")

	var project, section string

	switch len(ps) {
	case 1:
		return nil, entities.ErrProject
	case 2:
		project = ps[0]
		section = ps[1]
	}

	dimension := renameDim(fbase[1])
	quality := fbase[2]

	bars := bytes.Split(b, []byte("\n.\n"))

	allProfiles := make(map[string]*entities.Profile)

	for _, bar := range bars {
		profs := bytes.Split(bar, []byte("\n\n"))

		for _, prof := range profs[1:] {
			p := new(entities.Profile)

			p.Project = project
			p.Section = section

			ls := strings.Split(string(prof), "\n")

			pos := strings.Split(strings.TrimSpace(ls[0]), "-")

			switch len(pos) {
			case 1:
				p.PosNo = pos[0]
			case 2:
				p.Section = pos[0]
				p.PosNo = pos[1]
			}

			if found, exists := allProfiles[p.Project+p.Section+p.PosNo]; exists {
				found.Quantity++
				continue
			}

			p.Length, err = strconv.ParseFloat(ls[1], 64)
			if err != nil {
				return nil, err
			}

			p.Quality = quality
			p.Dim = dimension
			p.Quantity = 1

			p.L, err = parseEnd(ls[2])
			if err != nil {
				return nil, err
			}
			p.R, err = parseEnd(ls[3])
			if err != nil {
				return nil, err
			}

			if len(ls) > 4 {
				var inv bool

				for _, h := range ls[5:] {
					h = strings.TrimSpace(h)

					if h == "inv" {
						inv = true
						continue
					}

					hs := strings.Fields(h)
					h := new(entities.Hole)

					h.Name = hs[0]
					h.Params = make(map[string]float64)

					for _, param := range hs[1:] {
						spl := strings.Split(strings.ToUpper(param), "=")

						k := spl[0]
						v, err := strconv.ParseFloat(spl[1], 64)
						if err != nil {
							return nil, err
						}

						h.Params[k] = v
					}

					p.Holes = append(p.Holes, h)
				}

				if inv {
					p.InvertHolesX()
				}
			}

			allProfiles[p.Project+p.Section+p.PosNo] = p
		}
	}
	return allProfiles, nil
}

func parseEnd(s string) (*entities.End, error) {
	fs := strings.Fields(s)
	e := new(entities.End)
	e.Params = make(map[string]float64)

	nb := strings.Split(fs[0], "-")
	e.Name = nb[0]
	if len(nb) > 1 {
		e.WebBevel = getBevel(nb[1])
	}

	for _, p := range fs[1:] {
		spl := strings.Split(strings.ToUpper(p), "=")

		k := spl[0]
		v, err := strconv.ParseFloat(spl[1], 64)
		if err != nil {
			return nil, err
		}

		switch k {
		case "ANGLE":
			k = "V1"
			v += 90
		case "V1":
			v += 90
		case "V2":
			v = 90 - v
		case "SNIPE":
			k = "Ks"
		}

		e.Params[k] = v
	}

	return e, nil
}

func getBevel(s string) *entities.Bevel {
	b := new(entities.Bevel)

	return b
}

func renameDim(dim string) string {
	dim = strings.ToUpper(dim)
	switch dim {
	case "10":
		dim = "PP100*6.0"
	case "12":
		dim = "PP120*6.5"
	case "14a":
		dim = "PP140*7.0"
	case "16a":
		dim = "PP160*8.0"
	case "16b":
		dim = "PP160*10.0"
	case "18a":
		dim = "PP180*9.0"
	case "18b":
		dim = "PP180*11.0"
	case "20a":
		dim = "PP200*10.0"
	case "20b":
		dim = "PP200*11.0"
	case "22a":
		dim = "PP220*11.0"
	case "22b":
		dim = "PP220*13.0"
	case "24a":
		dim = "PP240*12.0"
	case "24b":
		dim = "PP240*14.0"
	default:
		if strings.Contains(dim, "RP") {
			dim = strings.ReplaceAll(dim, "RP", "PP")
		}
		if strings.Contains(dim, "HP") {
			dim = strings.ReplaceAll(dim, "HP", "PP")
		}
		if strings.ContainsRune(dim, 'X') {
			dim = strings.ReplaceAll(dim, "X", "*")
		}
		if !strings.ContainsRune(dim, '.') {
			dim += ".0"
		}
	}
	return dim
}
