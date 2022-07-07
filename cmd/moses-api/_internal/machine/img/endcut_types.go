package img

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"

	"github.com/edanko/moses/internal/entities"
)

func params(dim string, e *entities.End) (string, string) {
	if dim[:2] == "rp" {
		spl := strings.Split(dim[2:], "*")
		t, _ := strconv.ParseFloat(spl[1], 64)
		h, _ := strconv.ParseFloat(spl[0], 64)

		switch e.Name {
		case "0":
			return e21(e.Params["angle"], 0, 0, 0, 0, 0, 0), ""

		case "bk":
			if e.Params["angle"] > 0 {
				return e21(e.Params["angle"], 0, 0, 0, 0, 0, 0), sc12(e.Params["h"], e.Params["h"], 0, 0, e.Params["angle"])
			} else {
				return e21(e.Params["angle"], 0, e.Params["h"], 0, 0, 0, 0), ""
			}

		case "0sc":
			return e21r(0, 0, e.Params["r"], 0, 0, 0, 0), ""

		case "e23f":
			return e23f(e.Params["a"], e.Params["b"]), ""

		case "e23f_sc":
			return e23f(e.Params["a"], e.Params["b"]), sc12(e.Params["r"], e.Params["r"], 0, e.Params["r"], e.Params["angle"])

		case "fl":
			return e21r(e.Params["angle"], 45, e.Params["r"], 0, 0, 0, 0), ""

		case "fv":
			return e21r(e.Params["angle"], 45, e.Params["r"], e.Params["a"], t-2, 0, 0), ""

		case "fv.1":
			return e21r(e.Params["angle"], 45, e.Params["r"], 0, 0, 0, 0), ""

		case "sb1":
			return e213rs(0, 30, 30, e.Params["r"], e.Params["r"], 0, e.Params["r"], -1, 30, t-2, 0, 0), ""

		case "sb1.1":
			return e21r(e.Params["angle"], 0, e.Params["r"], 0, 0, 0, 0), ""

		case "sb2":
			return e213rs(0, 45, 45, e.Params["r"], e.Params["r"], 0, e.Params["r"], -1, 45, t-2, 0, 0), ""

		case "sb3", "sb20":
			return e21r(e.Params["angle"], 45, e.Params["r"], 30, t-2, 0, 0), ""

		case "sb3.1", "sb20.1":
			return e21r(e.Params["angle"], 45, e.Params["r"], 0, 0, 0, 0), ""

		case "sb3v":
			return e22(e.Params["angle"], 45, e.Params["r"], e.Params["r"]+e.Params["u"], e.Params["u"], e.Params["r"], 0, 30, t-2, 0, 0), ""

		case "sb4", "sb4-01":
			return e27r(0, 0, e.Params["down"], e.Params["up"], e.Params["up"], 0, 45, t-2, 0, 0), ""

		case "sb4_2":
			return e27(0, 0, e.Params["e"], e.Params["up"], e.Params["up"], 0, 45, t-2, 0, 0), ""

		case "sb5", "sb5-01":
			return e21r(e.Params["angle"], 45, e.Params["r"], 45, t-2, 0, 0), ""

		case "sb5_2":
			return e21(e.Params["angle"], 45, e.Params["e"], 45, t-2, 0, 0), ""

		case "sb5v":
			return e22(e.Params["angle"], 45, e.Params["r"], e.Params["r"]+e.Params["u"], e.Params["u"], e.Params["r"], 0, 45, t-2, 0, 0), ""

		case "sb6":
			return e213rs(0, 45, 45, e.Params["r"], e.Params["r"], 0, e.Params["r"], -1, 45, t-2, 0, 0), ""

		case "sb6.1":
			return e21r(e.Params["angle"], 0, e.Params["r"], 0, 0, 0, 0), ""

		case "sb8":
			return e210rs(e.Params["angle"], 30, 30, 0, 0, 0, 0, 13, 30, t-2, 0, 0), ""

		case "sb8.1":
			return e210rs(e.Params["angle"], 30, 30, 0, 0, 0, 0, 13, 0, 0, 0, 0), ""

		case "sb9":
			return e213rs(e.Params["angle"], 45, 45, e.Params["r"], e.Params["r"], 0, e.Params["r"], 15, 0, 0, 0, 0), ""

		case "sb9-01":
			return e213rs(e.Params["angle"], 45, 45, e.Params["r"], e.Params["r"], 0, e.Params["r"], 15, 45, t-2, 0, 0), ""

		case "sb10":
			return e28r(0, 0, e.Params["r"], 1, 0, e.Params["a"], h-20, 0, 0, 0, 0), ""

		case "sb11":
			return e21r(0, 45, e.Params["r"], 0, 0, 0, 0), ""

		case "sb11_2":
			return e21(0, 45, e.Params["e"], 0, 0, 0, 0), ""

		case "sc50":
			return e22(0, 0, e.Params["r"]+50, e.Params["r"], 0, e.Params["r"], 0, 0, 0, 0, 0), ""

		case "bl":
			return e21r(e.Params["angle"], 0, e.Params["r"], 0, 0, 0, 0), ""

		case "blg":
			return e21r(e.Params["angle"], 0, e.Params["r"], 45, t-2, 0, 0), ""

		case "br":
			if e.Params["h"] != 0 {
				return e29(e.Params["angle"], 0, 5, 5, 0, 30, d(h, e.Params["bulb"]), e.Params["bulb"], 0, 0, 0, 0), sc12(e.Params["h"], e.Params["h"], 0, 0, e.Params["angle"])
			} else {
				return e29(e.Params["angle"], 0, 5, 5, 0, 30, d(h, e.Params["bulb"]), e.Params["bulb"], 0, 0, 0, 0), ""
			}

		case "brbz":
			if e.Params["h"] != 0 {
				return e29(e.Params["angle"], e.Params["angle2"], 5, 5, 0, 30, d(h, e.Params["bulb"]), e.Params["bulb"], 0, 0, 0, 0), sc12(e.Params["h"], e.Params["h"], 0, 0, e.Params["angle"])
			} else {
				return e29(e.Params["angle"], e.Params["angle2"], 5, 5, 0, 30, d(h, e.Params["bulb"]), e.Params["bulb"], 0, 0, 0, 0), ""
			}

		case "bvbz":
			return e221rs(e.Params["angle"], e.Params["angle2"], e.Params["b"], e.Params["h"], e.Params["h"]-e.Params["r"], e.Params["r"], e.Params["h2"]-t, -1, 0, 0, 0, 0), ""

		case "ba":
			return e21(e.Params["angle"], 0, 0, 0, 0, 0, 0), ""

		case "bc":
			return e23(0, 90-e.Params["angle"], 0, 0, 0, 0, 0, 0, 0), ""

		case "bck":
			return e23(e.Params["angle"], 90-e.Params["angle2"], 0, e.Params["b"], e.Params["h"], 0, 0, 0, 0), ""

		case "bcr":
			return e23r(e.Params["angle"], 90-e.Params["angle2"], 0, e.Params["r"], e.Params["h"], 0, 0, 0, 0), ""

		case "bac":
			return e23(0, bac(h, e.Params["3"], e.Params["2"]), 0, 0, e.Params["2"], 0, 0, 0, 0), sc12(e.Params["1"], e.Params["2"], 0, 0, e.Params["angle"])

		case "bv":
			return e22(e.Params["angle"], 0, e.Params["b"], e.Params["h"], e.Params["h"]-e.Params["r"], e.Params["r"], 0, 0, 0, 0, 0), ""

		case "bvf":
			if e.Params["r"] == 0 {
				return e21(e.Params["angle"], 25, 0, 0, 0, 0, 0), ""
			} else {
				return e22(e.Params["angle"], 25, e.Params["b"], e.Params["h"], e.Params["h"]-e.Params["r"], e.Params["r"], 0, 0, 0, 0, 0), ""
			}

		case "bvg":
			return e22(e.Params["angle"], 0, e.Params["b"], e.Params["h"], e.Params["h"]-e.Params["r"], e.Params["r"], 0, 0, 0, 0, 0), ""

		case "bw":
			return e29(e.Params["angle"], 0, 5, 5, 0, 30, d(h, e.Params["bulb"]), e.Params["bulb"], 0, 0, 0, 0), sc12(e.Params["r"], e.Params["r"], 0, e.Params["r"], e.Params["angle"])

		case "btn":
			if e.Params["h"] != 0 {
				return e29(e.Params["angle"], 0, 5, 5, 0, 30, d(h, e.Params["bulb"]), e.Params["bulb"], 0, 0, 0, 0), sc12(e.Params["l"], e.Params["h"], 0, e.Params["r"], e.Params["angle"])
			} else {
				return e29(e.Params["angle"], 0, 5, 5, 0, 30, d(h, e.Params["bulb"]), e.Params["bulb"], 0, 0, 0, 0), ""
			}

		case "bmn":
			return e22(e.Params["angle"], 0, e.Params["b"], e.Params["r"], 0, e.Params["r"], 0, 0, 0, 0, 0), ""

		case "bmnbz":
			return e221rs(e.Params["angle"], e.Params["angle2"], e.Params["b"], e.Params["r"], 0, e.Params["r"], e.Params["h2"]-t, -1, 0, 0, 0, 0), ""

		case "blgbz":
			return e221rs(e.Params["angle"], e.Params["angle2"], e.Params["r"], e.Params["r"], 0, e.Params["r"], e.Params["h2"]-t, -1, 45, t-2, 0, 0), ""

		case "blbz":
			return e221rs(e.Params["angle"], e.Params["angle2"], e.Params["r"], e.Params["r"], 0, e.Params["r"], e.Params["h2"]-t, -1, 0, 0, 0, 0), ""

		case "bkbz":
			return e221rs(e.Params["angle"], e.Params["angle2"], e.Params["h"], e.Params["h"], 0, 0, e.Params["h2"]-t, -1, 0, 0, 0, 0), ""

		case "babz":
			return e221rs(e.Params["angle"], e.Params["angle2"], 0, 0, 0, 0, e.Params["h2"]-t, -1, 0, 0, 0, 0), ""

		case "bxaf", "bxag":
			var c, r2 float64

			if e.Params["r"] > 0 {
				c = e.Params["h"] - e.Params["r"]
			}

			if t/2 < 1 {
				r2 = 1
			} else {
				r2 = t / 2
			}

			return e372rs(25, e.Params["l"], e.Params["h"], c, e.Params["r"], h/4, l(h, e.Params["hprof"]), h-e.Params["hprof"], 1, r2, 0, 0, 0, 0, 0, 0, 0), ""

		case "0_45":
			return e21(0, 45, 0, 0, 0, 0, 0), ""

		case "1":
			return e27r(0, 0, e.Params["down"], e.Params["up"], e.Params["up"], 0, 0, 0, 0, 0), ""

		case "b-102":
			return e213rs(e.Params["angle"], 30, 30, e.Params["r"], e.Params["r"], 0, e.Params["r"], 12, 30, t-2, 0, 0), ""

		case "21":
			if e.Params["snipe"] != 0 {
				var e1 float64
				if e.Params["angle"] > 90 {
					e1 = 0
				} else {
					e1 = e.Params["snipe"]
				}
				return e21(e.Params["v1"], e.Params["v2"], e1, 0, 0, 0, 0), sc12(e.Params["snipe"], e.Params["snipe"], 0, 0, e.Params["v1"])
			} else {
				return e21r(e.Params["v1"], e.Params["v2"], e.Params["r1"], 0, 0, 0, 0), ""
			}

		case "21-116", "21-216", "21-145", "21--145", "21-173", "21--173", "21-245", "21-273", "21-275", "21-346":
			if e.Params["snipe"] != 0 {
				var e1 float64
				if e.Params["angle"] > 90 {
					e1 = 0
				} else {
					e1 = e.Params["snipe"]
				}
				return e21(e.Params["v1"], e.Params["v2"], e1, 45, t-2, 0, 0), sc12(e.Params["snipe"], e.Params["snipe"], 0, 0, e.Params["v1"])
			} else {
				return e21r(e.Params["v1"], e.Params["v2"], e.Params["r1"], 45, t-2, 0, 0), ""
			}

		case "21-131", "21-132":
			return e213rs(e.Params["v1"], 25, 25, e.Params["r1"], e.Params["r1"], 0, e.Params["r1"], -1, 25, t-2, 0, 0), ""

		case "21--132":
			return e213rs(e.Params["v1"], 25, 25, e.Params["r1"], e.Params["r1"], 0, e.Params["r1"], -1, 0, 0, 25, t-2), ""

		case "23":
			if e.Params["r"] > 0 {
				return e23r(e.Params["v1"], e.Params["v3"], e.Params["v2"], e.Params["r1"], e.Params["b"], 0, 0, 0, 0), ""
			} else {
				return e23(e.Params["v1"], e.Params["v3"], e.Params["v2"], 0, e.Params["b"], 0, 0, 0, 0), ""
			}

		case "27":
			return e27r(e.Params["v1"], e.Params["v2"], e.Params["r1"], e.Params["r2"], e.Params["a"], 0, 0, 0, 0, 0), ""

		default:
			fmt.Println("unknown endcut:", e.Name)
			// return "", ""
		}
	}
	if dim[:2] == "fb" {
		switch e.Name {
		case "ba":
			return e11(e.Params["angle"], 0, 0, 0, 0, 0), ""

		case "bk":
			return e11(e.Params["angle"], e.Params["h"], 0, 0, 0, 0), ""

		case "baosv":
			return e11(e.Params["angle"], 0, 0, 0, e.Params["deg"], 2), ""

		case "batsv":
			return e11(e.Params["angle"], 0, e.Params["deg"], 2, 0, 0), ""

		case "bc":
			return e13(90, e.Params["angle"], 0, e.Params["h"], 0, 0, 0, 0), ""

		case "bck":
			return e13(e.Params["angle"], e.Params["angle2"], e.Params["b"], e.Params["h"], 0, 0, 0, 0), ""

		default:
			return "", ""
		}
	}

	return "", ""

}

func holeToString(h *entities.Hole) string {
	switch h.Name {
	case "ae_half_hole", "hole", "i1":
		return i1(h.Params["x"], h.Params["r"])

	case "notch":
		return i1(h.Params["u"], h.Params["r"])

	case "ae_hole", "i2":
		return i2(h.Params["x"], h.Params["y"], h.Params["r"])

	case "hc":
		return i4(h.Params["x"], h.Params["a"], h.Params["b"], h.Params["r"])

	case "hin":
		return i11(h.Params["x"], h.Params["l"], h.Params["b"])
	default:
		return ""
	}
}

func d(height, bulb float64) float64 {
	var sub float64
	switch bulb {
	default:
		log.Fatalln("unknown bulb", bulb, "in br or bw template")
	case 21:
		sub = 20.01
	case 24:
		sub = 21.74
	case 27:
		sub = 25.05
	case 29:
		sub = 27.78
	case 32:
		sub = 29.52
	case 35:
		sub = 32.83
	case 38:
		sub = 35.35
	}

	return height - sub
}

func bac(h, arg1, arg2 float64) float64 {
	return 90 - math.Atan(arg1/(h-arg2))*180/math.Pi
}

func l(height, hprof float64) float64 {
	switch height - hprof {
	case 20, 24:
		return 100
	case 40:
		return 188
	case 60:
		return 260
	case 80:
		return 290
	case 100:
		return 374
	case 120:
		return 449
	default:
		log.Fatalln("unknown height and hprof difference:", height-hprof)
	}
	return 0
}

// get angle for cx template
func cx_doc(h, t, legUp float64) float64 {
	if legUp == 0 {
		legUp = h
	}
	legSide := (h - (3 * t)) / 2

	return 90 - math.Atan(legUp/legSide)*180/math.Pi
}

// for fl bv template
func a(h, angle float64) float64 {
	if angle == 0 {
		return 0
	}

	var sign float64

	// just inverted direction
	if angle < 0 {
		sign = 1
	} else {
		sign = -1
	}

	return sign * h * math.Sin(math.Abs(angle)*math.Pi/180)
}
