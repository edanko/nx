package parser

import (
	"fmt"
	"os"
	"strconv"

	"github.com/edanko/moses/internal/entities"

	"github.com/edanko/gen"
)

const (
	prj = "056"
)

func bev(e *gen.End) *entities.Bevel {
	if e.AngleTs == 0 && e.AngleOs == 0 {
		return nil
	}
	return &entities.Bevel{
		AngleTs: e.AngleTs,
		AngleOs: e.AngleOs,
		DepthTs: e.DepthTs,
		DepthOs: e.DepthOs,
	}
}

func ProcessGen(fname string) (map[string]*entities.Profile, error) {
	f, err := os.Open(fname)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	g := gen.ReadListedProfile(f)

	profs := make(map[string]*entities.Profile, len(g))

	for _, prof := range g {
		p := new(entities.Profile)

		p.Project = prj
		p.Section = prof.BlockNo
		p.PosNo = prof.PosNo

		if prof, exists := profs[p.Project+p.Section+p.PosNo]; exists {
			prof.Quantity++
			continue
		}

		p.Quantity = 1
		p.Length = prof.TlengthManual

		if prof.Form != "STRAIGHT" {
			fmt.Println(prof.Form)
		}

		// if g.CommonData.Shape[:2] == "PP" {
		// 	switch g.CommonData.Dimension {
		// 	case "203*10.0":
		// 		g.CommonData.Dimension = "200*10.0"
		// 	case "183*11.0":
		// 		g.CommonData.Dimension = "180*11.0"
		// 	case "143*9.0":
		// 		g.CommonData.Dimension = "140*9.0"
		// 	case "163*10.0":
		// 		g.CommonData.Dimension = "160*10.0"
		// 	case "143*7.0":
		// 		g.CommonData.Dimension = "140*7.0"
		// 	case "223*11.0":
		// 		g.CommonData.Dimension = "220*11.0"
		// 	}
		// }

		// p.Dim = g.CommonData.Shape[:2] + g.CommonData.Dimension
		// p.Quality = g.CommonData.Quality

		p.L = new(entities.End)
		p.L.Name = strconv.Itoa(prof.LeftEnd.EndcutType)

		params := make(map[string]float64, 10)

		if prof.LeftEnd.A != 0 {
			params["A"] = prof.LeftEnd.A
		}
		if prof.LeftEnd.B != 0 {
			params["B"] = prof.LeftEnd.B
		}
		if prof.LeftEnd.C != 0 {
			params["C"] = prof.LeftEnd.C
		}
		if prof.LeftEnd.Ks != 0 {
			params["Ks"] = prof.LeftEnd.Ks
		}
		if prof.LeftEnd.R1 != 0 {
			params["R1"] = prof.LeftEnd.R1
		}
		if prof.LeftEnd.R2 != 0 {
			params["R2"] = prof.LeftEnd.R2
		}
		if prof.LeftEnd.V1 != 0 {
			params["V1"] = prof.LeftEnd.V1
		}
		if prof.LeftEnd.V2 != 0 {
			params["V2"] = prof.LeftEnd.V2
		}
		if prof.LeftEnd.V3 != 0 {
			params["V3"] = prof.LeftEnd.V3
		}
		if prof.LeftEnd.V4 != 0 {
			params["V4"] = prof.LeftEnd.V4
		}

		p.L.Params = params

		p.L.WebBevel = bev(prof.LeftEnd)

		p.R = new(entities.End)
		p.R.Name = strconv.Itoa(prof.RightEnd.EndcutType)

		params = make(map[string]float64, 10)

		if prof.RightEnd.A != 0 {
			params["A"] = prof.RightEnd.A
		}
		if prof.RightEnd.B != 0 {
			params["B"] = prof.RightEnd.B
		}
		if prof.RightEnd.C != 0 {
			params["C"] = prof.RightEnd.C
		}
		if prof.RightEnd.Ks != 0 {
			params["Ks"] = prof.RightEnd.Ks
		}
		if prof.RightEnd.R1 != 0 {
			params["R1"] = prof.RightEnd.R1
		}
		if prof.RightEnd.R2 != 0 {
			params["R2"] = prof.RightEnd.R2
		}
		if prof.RightEnd.V1 != 0 {
			params["V1"] = prof.RightEnd.V1
		}
		if prof.RightEnd.V2 != 0 {
			params["V2"] = prof.RightEnd.V2
		}
		if prof.RightEnd.V3 != 0 {
			params["V3"] = prof.RightEnd.V3
		}
		if prof.RightEnd.V4 != 0 {
			params["V4"] = prof.RightEnd.V4
		}

		p.R.Params = params

		p.R.WebBevel = bev(prof.RightEnd)

		for _, h := range prof.Holes {
			c := new(entities.Hole)
			c.Name = h.Name
			c.Params = make(map[string]float64)

			if h.DistOrigin != 0 {
				c.Params["X"] = h.DistOrigin
			}
			if h.DistOriginV != 0 {
				c.Params["Y"] = h.DistOriginV
			}

			p.Holes = append(p.Holes, c)
		}

		profs[p.Project+p.Section+p.PosNo] = p
	}
	return profs, nil
}
