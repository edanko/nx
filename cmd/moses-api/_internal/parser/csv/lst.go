package csv

import (
	"bufio"
	"io"
	"regexp"
	"strings"
)

var (
	reProfType = regexp.MustCompile(`PP\/(\d+)x(\d+\.?\d+?)`)
	reQuality  = regexp.MustCompile(`Qual :\s+(\w+)`)
)

func DimAndQualityFromLst(rd io.Reader) (dim, quality string, err error) {
	s := bufio.NewScanner(rd)

	for s.Scan() {
		if dim != "" && quality != "" {
			break
		}

		str := s.Text()

		if strings.Contains(str, "Type/Dim.") {
			tmp := reProfType.FindAllStringSubmatch(str, -1)
			dim = "RP" + tmp[0][1] + "*" + tmp[0][2]
			continue
		}

		if strings.Contains(str, "Qual") {
			tmp := reQuality.FindAllStringSubmatch(str, -1)
			quality = tmp[0][1]
			continue
		}
	}
	if err = s.Err(); err != nil {
		return
	}

	return
}
