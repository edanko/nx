package img

import (
	"fmt"
	"strings"
)

func spec(dim string) (string, error) {
	switch dim[:2] {
	case "PP":
		var b, c, s, r float64

		switch dim {
		case "PP100*6.0":
			b = 100
			c = 20
			s = 6
			r = 5
		case "PP120*6.5":
			b = 120
			c = 23.5
			s = 7 // 6.5
			r = 5
		case "PP140*7.0":
			b = 140
			c = 26
			s = 7
			r = 6
		case "PP140*9.0":
			b = 140
			c = 26
			s = 9
			r = 6
		case "PP160*8.0":
			b = 160
			c = 28
			s = 8
			r = 7
		case "PP160*10.0":
			b = 160
			c = 28
			s = 10
			r = 7
		case "PP180*9.0":
			b = 180
			c = 31
			s = 9
			r = 7
		case "PP180*11.0":
			b = 180
			c = 31
			s = 11
			r = 7
		case "PP200*10.0":
			b = 200
			c = 34
			s = 10
			r = 8
		case "PP200*11.0":
			b = 200
			c = 34
			s = 11
			r = 8
		case "PP200*12.0":
			b = 200
			c = 34
			s = 12
			r = 8
		case "PP220*11.0":
			b = 220
			c = 37
			s = 11
			r = 8.5
		case "PP220*13.0":
			b = 220
			c = 37
			s = 13
			r = 8.5
		case "PP240*12.0":
			b = 240
			c = 40
			s = 12
			r = 9
		case "PP240*14.0":
			b = 240
			c = 40
			s = 14
			r = 9
		default:
			return "", fmt.Errorf("unknown pp spec in IMG machine: %s", dim)
		}
		return fmt.Sprintf("SHAPE=HP\r\nSTART_OF_PARAMS\r\nNO_OF_PARAMS=5\r\nNORM=%s\r\nB=%0.1f\r\nC=%0.1f\r\nS=%0.1f\r\nR=%0.1f\r\nEND_OF_PARAMS", dim, b, c, s, r), nil

	case "FB":
		t := strings.Split(dim[2:], "*")
		return fmt.Sprintf("SHAPE=FB\r\nSTART_OF_PARAMS\r\nNO_OF_PARAMS=3\r\nNORM=%s\r\nB=%s\r\nS=%s\r\nEND_OF_PARAMS", dim, t[0], t[1]), nil

	default:
		return "", fmt.Errorf("unknown profile spec in IMG machine: %s", dim)
	}
}
