package jinfeng

import (
	"fmt"
)

func Spec(dim string) (string, error) {
	switch dim[:2] {
	case "PP":
		switch dim {
		case "PP100*6.0":
			return "P0100x06.0", nil
		case "PP120*6.5":
			return "P0120x06.5", nil
		case "PP140*7.0":
			return "P0140x07.0", nil
		case "PP140*9.0":
			return "P0140x09.0", nil
		case "PP160*8.0":
			return "P0160x08.0", nil
		case "PP160*10.0":
			return "P0160x10.0", nil
		case "PP180*9.0":
			return "P0180x09.0", nil
		case "PP180*11.0":
			return "P0180x11.0", nil
		case "PP200*10.0":
			return "P0200x10.0", nil
		case "PP200*11.0":
			return "P0200x11.0", nil
		case "PP200*12.0":
			return "P0200x12.0", nil
		case "PP220*11.0":
			return "P0220x11.0", nil
		case "PP220*13.0":
			return "P0220x13.0", nil
		case "PP240*12.0":
			return "P0240x12.0", nil
		case "PP240*14.0":
			return "P0240x14.0", nil
		default:
			return "", fmt.Errorf("unknown pp spec in JinFeng machine: %s", dim)
		}

	default:
		return "", fmt.Errorf("unknown profile spec in JinFeng machine: %s", dim)
	}
}
