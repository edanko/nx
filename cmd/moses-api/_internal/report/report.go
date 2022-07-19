package report

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/edanko/moses/internal/entities"
)

const (
	newLine = "\n"
)

func NestingListString(nn []*entities.Nest) (string, error) {
	o := strings.Builder{}
	o.Grow(2048)

	byDim := make(map[entities.MachineType]map[string]map[string][]*entities.Nest)

	for _, n := range nn {
		if _, ok := byDim[n.Machine]; !ok {
			byDim[n.Machine] = make(map[string]map[string][]*entities.Nest)
		}

		if _, ok := byDim[n.Machine][n.Bar.Dim]; !ok {
			byDim[n.Machine][n.Bar.Dim] = make(map[string][]*entities.Nest)
		}

		byDim[n.Machine][n.Bar.Dim][n.Bar.Quality] = append(byDim[n.Machine][n.Bar.Dim][n.Bar.Quality], n)
	}

	bl := strings.Builder{}
	bl.Grow(1024)

	for m, dd := range byDim {
		//fmt.Printf("bars for %s\n", m)
		bl.WriteString("bars for ")
		bl.WriteString(m.String())
		bl.WriteRune('\n')

		for d, qq := range dd {
			//fmt.Printf("dim: %s \n", d)
			bl.WriteString(d)
			bl.WriteRune(' ')

			for q, nn := range qq {
				//fmt.Printf("quality: %s\n", q)

				bl.WriteString(q)
				bl.WriteRune(' ')

				var stockBars, remnants int

				for _, n := range nn {
					if n.Bar.Remnant != nil {
						remnants++
					} else {
						stockBars++
					}
				}

				if stockBars > 0 {
					// fmt.Printf("stock bars: %d\n", stockBars)
					bl.WriteString("stock bars: ")
					bl.WriteString(strconv.Itoa(stockBars))
				}

				if stockBars > 0 && remnants > 0 {
					bl.WriteString(" and ")
				}

				if remnants > 0 {
					//fmt.Printf("remnants: %d\n", remnants)
					bl.WriteString("remnants: ")
					bl.WriteString(strconv.Itoa(remnants))
				}
				bl.WriteRune('\n')
				//fmt.Println()
			}

		}

	}

	fmt.Println(bl.String())

	for m, dd := range byDim {
		o.WriteString(fmt.Sprintf("nests for %s\n\n", m))

		for d, qq := range dd {
			o.WriteString(fmt.Sprintf("---- %s", d))

			for q, nn := range qq {
				o.WriteString(fmt.Sprintf(" / %s -----\n", q))

				for _, n := range nn {
					o.WriteString("Имя программы: ")
					o.WriteString(n.NestName)
					o.WriteString(newLine)

					if n.Bar.Remnant != nil {
						o.WriteString("Использовать ДМО: ")
						o.WriteString(n.Bar.Remnant.From)
						o.WriteString(newLine)
					}

					o.WriteString("Длина: ")
					o.WriteString(strconv.FormatFloat(n.Bar.Length, 'g', -1, 64))
					o.WriteString(" / Использовано: ")
					o.WriteString(strconv.FormatFloat(n.Bar.UsedLength, 'g', -1, 64))
					o.WriteString(" / Лом: ")
					o.WriteString(strconv.FormatFloat(n.Scrap(), 'g', -1, 64))
					o.WriteString(fmt.Sprintf(" (%.2f %%)", n.NestingPercent()))
					o.WriteString(newLine)

					if r := n.GetRemnant(); r != nil {
						o.WriteString("ДМО: ")
						o.WriteString(strconv.FormatFloat(r.Length, 'g', -1, 64))
						o.WriteString(" / Маркировка ДМО: ")
						o.WriteString(r.Marking)
						o.WriteString(newLine)
					}

					o.WriteString("------------------------------\r\n| Секция | Позиция |  Длина  |\r\n------------------------------\r\n")

					for _, p := range n.Profiles {

						o.WriteString(fmt.Sprintf("| %6s | %7s | %7g |\r\n", p.Section, p.PosNo, p.Length))
					}
					o.WriteString("------------------------------\r\n")

				}
				o.WriteString(newLine)

			}
			o.WriteString(newLine)

		}
		o.WriteString(newLine)

	}

	return o.String(), nil
}

// TODO: move to report service someday
/*barlist := "Ведомость расхода материала"
partlist := "Партлист"

f := excelize.NewFile()

f.SetSheetName("Sheet1", barlist)

_ = f.SetCellValue(barlist, "A1", "Карта раскроя")       // 1
_ = f.SetCellValue(barlist, "B1", "Тип оборудования")    // 2
_ = f.SetCellValue(barlist, "C1", "Исп. ДМО")            // 3
_ = f.SetCellValue(barlist, "D1", "Типоразмер")          // 4
_ = f.SetCellValue(barlist, "E1", "Марка материала")     // 5
_ = f.SetCellValue(barlist, "F1", "Длина, мм")           // 6
_ = f.SetCellValue(barlist, "G1", "Масса заготовки, кг") // 7
_ = f.SetCellValue(barlist, "H1", "Коэф. раскроя, %")    // 8
_ = f.SetCellValue(barlist, "I1", "Маркировка ДМО")      // 9
_ = f.SetCellValue(barlist, "J1", "Длина ДМО, мм")       // 10
_ = f.SetCellValue(barlist, "K1", "Масса ДМО, кг")       // 11

row := 2

for _, b := range allBars {
	cell, _ := excelize.CoordinatesToCellName(1, row)
	_ = f.SetCellValue(barlist, cell, b.NestName())

	cell, _ = excelize.CoordinatesToCellName(2, row)
	_ = f.SetCellValue(barlist, cell, "Камера резки профиля")

	cell, _ = excelize.CoordinatesToCellName(3, row)
	_ = f.SetCellValue(barlist, cell, b.UsedRemnant)

	cell, _ = excelize.CoordinatesToCellName(4, row)
	_ = f.SetCellValue(barlist, cell, b.Dim())

	cell, _ = excelize.CoordinatesToCellName(5, row)
	_ = f.SetCellValue(barlist, cell, b.Quality())

	cell, _ = excelize.CoordinatesToCellName(6, row)
	_ = f.SetCellValue(barlist, cell, b.Capacity)

	cell, _ = excelize.CoordinatesToCellName(7, row)
	_ = f.SetCellValue(barlist, cell, mass(b.Dim(), b.Capacity))

	cell, _ = excelize.CoordinatesToCellName(8, row)
	_ = f.SetCellValue(barlist, cell, b.Length/b.Capacity*100)

	if b.Capacity-b.Length > 1000 {
		cell, _ = excelize.CoordinatesToCellName(9, row)
		_ = f.SetCellValue(barlist, cell, b.NestName()+"R01")

		cell, _ = excelize.CoordinatesToCellName(10, row)
		_ = f.SetCellValue(barlist, cell, b.RemnantLength())

		cell, _ = excelize.CoordinatesToCellName(11, row)
		_ = f.SetCellValue(barlist, cell, mass(b.Dim(), b.RemnantLength()))
	}

	row++
}

f.SetActiveSheet(f.NewSheet(partlist))

_ = f.SetCellValue(partlist, "A1", "Чертеж")             // 1
_ = f.SetCellValue(partlist, "B1", "Заказ")              // 2
_ = f.SetCellValue(partlist, "C1", "Секция")             // 3
_ = f.SetCellValue(partlist, "D1", "Позиция")            // 4
_ = f.SetCellValue(partlist, "E1", "Карта раскроя")      // 5
_ = f.SetCellValue(partlist, "F1", "Типоразмер детали")  // 6
_ = f.SetCellValue(partlist, "G1", "Марка материала")    // 7
_ = f.SetCellValue(partlist, "H1", "Кол-во, шт")         // 8
_ = f.SetCellValue(partlist, "I1", "Длина, мм")          // 9
_ = f.SetCellValue(partlist, "J1", "Масса 1 детали, кг") // 10
_ = f.SetCellValue(partlist, "K1", "Общая масса, кг")    // 11
_ = f.SetCellValue(partlist, "L1", "Маршрут обработки")  // 12
_ = f.SetCellValue(partlist, "M1", "Примечание")         // 13

row = 2

for _, b := range allBars {
	for _, p := range b.Parts {
		//cell, _ := excelize.CoordinatesToCellName(1, row)
		//_ = f.SetCellValue(partlist, cell, b.NestName())

		cell, _ := excelize.CoordinatesToCellName(2, row)
		_ = f.SetCellValue(partlist, cell, p.Project)

		cell, _ = excelize.CoordinatesToCellName(3, row)
		_ = f.SetCellValue(partlist, cell, p.Section)

		cell, _ = excelize.CoordinatesToCellName(4, row)

		pos, _ := strconv.Atoi(p.PosNo)
		_ = f.SetCellValue(partlist, cell, pos)

		cell, _ = excelize.CoordinatesToCellName(5, row)
		_ = f.SetCellValue(partlist, cell, b.NestName())

		cell, _ = excelize.CoordinatesToCellName(6, row)
		_ = f.SetCellValue(partlist, cell, p.Dim)

		cell, _ = excelize.CoordinatesToCellName(7, row)
		_ = f.SetCellValue(partlist, cell, p.Quality)

		cell, _ = excelize.CoordinatesToCellName(8, row)
		_ = f.SetCellValue(partlist, cell, p.Quantity)

		cell, _ = excelize.CoordinatesToCellName(9, row)
		_ = f.SetCellValue(partlist, cell, p.Length)

		cell, _ = excelize.CoordinatesToCellName(10, row)
		_ = f.SetCellValue(partlist, cell, "m")

		cell, _ = excelize.CoordinatesToCellName(11, row)
		_ = f.SetCellValue(partlist, cell, "m2")

		//cell, _ = excelize.CoordinatesToCellName(12, row)
		//_ = f.SetCellValue(partlist, cell, mass(b.Dim(), b.RemnantLength()))

		//cell, _ = excelize.CoordinatesToCellName(13, row)
		//_ = f.SetCellValue(partlist, cell, mass(b.Dim(), b.RemnantLength()))

		row++
	}
}

if err := f.SaveAs(path.Join("out", nm+".xlsx")); err != nil {
	fmt.Println(err)
}*/

// gost 21937-76
/* func mass(dim string, length float64) float64 {

	length /= 1000

	var mass float64

	switch dim {
	case "RP100X6":
		mass = 6.76
	case "RP120X6.5":
		mass = 8.75
	case "RP140X7":
		mass = 11.05
	case "RP140X9":
		mass = 13.23
	case "RP160X8":
		mass = 14.08
	case "RP160X10":
		mass = 16.60
	case "RP180X9":
		mass = 17.41
	case "RP180X11":
		mass = 20.24
	case "RP200X10":
		mass = 21.47
	case "RP200X12":
		mass = 24.60
	case "RP220X11":
		mass = 25.75
	case "RP220X13":
		mass = 29.20
	case "RP240X12":
		mass = 30.42
	case "RP240X14":
		mass = 34.18
	}

	return mass * length
} */
