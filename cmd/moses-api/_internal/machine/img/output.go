package img

import (
	"context"
	"strconv"
	"strings"

	"github.com/edanko/moses/internal/entities"
)

const (
	newLine = "\r\n"
)

func Output(_ context.Context, n *entities.Nest) (string, error) {
	o := strings.Builder{}

	o.WriteString("TYPE_OF_GENERIC_FILE=MOSES_NESTED_PROFILE")
	o.WriteString(newLine)
	o.WriteString("VERSION=1.0")
	o.WriteString(newLine)
	o.WriteString("USAGE=PLASMA")
	o.WriteString(newLine)

	o.WriteString("COMMON_DATA")
	o.WriteString(newLine)
	o.WriteString("NEST_NAME=")
	o.WriteString(n.NestName)
	o.WriteString(newLine)

	s, err := spec(n.Bar.Dim)
	if err != nil {
		return "", err
	}
	o.WriteString(s)
	o.WriteString(newLine)

	o.WriteString("RAW_LENGTH=")
	o.WriteString(strconv.FormatFloat(n.Bar.Length, 'g', -1, 64))
	o.WriteString(newLine)

	o.WriteString("USED_LENGTH=")
	o.WriteString(strconv.FormatFloat(n.Bar.UsedLength, 'g', -1, 64))
	o.WriteString(newLine)

	rest := n.Bar.Length - n.Bar.UsedLength

	o.WriteString("REST_LENGTH=")
	o.WriteString(strconv.FormatFloat(rest, 'g', -1, 64))
	o.WriteString(newLine)

	o.WriteString("TEXT_HEIGHT=20")
	o.WriteString(newLine)

	o.WriteString("TEXT_WIDTH=16")
	o.WriteString(newLine)
	o.WriteString("TEXT_PLANE=0")
	o.WriteString(newLine)
	o.WriteString("TEXT_PLACING=3")
	o.WriteString(newLine)
	o.WriteString("TEXT_U=60")
	o.WriteString(newLine)
	o.WriteString("TEXT_V=40")
	o.WriteString(newLine)
	o.WriteString("MATERIAL=St37-2")
	o.WriteString(newLine)
	o.WriteString("NO_OF_PROFS=")
	o.WriteString(strconv.Itoa(len(n.Profiles)))
	o.WriteString(newLine)

	o.WriteString("END_OF_COMMON_DATA")
	o.WriteString(newLine)

	var leftPoint float64

	for i, p := range n.Profiles {
		o.WriteString("PROFILE_DATA")
		o.WriteString(newLine)

		o.WriteString("TLENGTH=")
		o.WriteString(strconv.FormatFloat(p.Length, 'g', -1, 64))
		o.WriteString(newLine)

		var ident string

		switch {
		case p.Length < 400:
			ident = p.PosNo

		case p.Length < 500:
			ident = p.Section + "-" + p.PosNo

		case p.Length > 500 && p.Length < 700:
			ident = p.Section + "-" + p.PosNo + " " + p.Quality

		default:
			ident = n.Project + "-" + p.Section + "-" + p.PosNo + " " + p.Quality
		}

		o.WriteString("IDENT_STRING=")
		o.WriteString(ident)
		o.WriteString(newLine)

		o.WriteString("NO_OF_MARKS=0")
		o.WriteString(newLine)
		o.WriteString("NO_OF_ICUTS=")
		o.WriteString(strconv.Itoa(len(p.Holes)))
		o.WriteString(newLine)

		o.WriteString("NO_OF_PARTS=")
		o.WriteString(strconv.Itoa(i + 1))
		o.WriteString(newLine)

		o.WriteString("END_OF_PROFILE_DATA")
		o.WriteString(newLine)

		o.WriteString("LEFT_END")
		o.WriteString(newLine)

		leftPoint += n.Spacings[i*2]

		o.WriteString("LEFT_CLOSEST_POINT=")
		o.WriteString(strconv.FormatFloat(leftPoint, 'g', -1, 64))
		o.WriteString(newLine)

		left, lscallop := params(n.Bar.Dim, p.L)

		o.WriteString(left)

		o.WriteString("END_OF_LEFT_END")
		o.WriteString(newLine)

		leftPoint += p.Length

		o.WriteString("RIGHT_END")
		o.WriteString(newLine)
		o.WriteString("LEFT_FARTHEST_POINT=")
		o.WriteString(strconv.FormatFloat(leftPoint, 'g', -1, 64))
		o.WriteString(newLine)

		right, rscallop := params(n.Bar.Dim, p.R)

		o.WriteString(right)

		o.WriteString("END_OF_RIGHT_END")
		o.WriteString(newLine)

		leftPoint += n.Spacings[i*2+1]

		if lscallop != "" {
			o.WriteString("LEFT_SCALLOP")
			o.WriteString(newLine)

			o.WriteString(lscallop)

			o.WriteString("END_OF_LEFT_SCALLOP")
			o.WriteString(newLine)
		}

		if rscallop != "" {
			o.WriteString("RIGHT_SCALLOP")
			o.WriteString(newLine)

			o.WriteString(rscallop)

			o.WriteString("END_OF_RIGHT_SCALLOP")
			o.WriteString(newLine)
		}

		if len(p.Holes) > 0 {
			for _, hole := range p.Holes {
				o.WriteString("START_OF_ICUT")
				o.WriteString(newLine)

				icutString := holeToString(hole)
				o.WriteString(icutString)

				o.WriteString("END_OF_ICUT")
				o.WriteString(newLine)
			}
		}
	}

	return o.String(), nil
}
