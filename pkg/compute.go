package pkg

import (
	"fmt"
	"math"
	"upwork-assessment-1/structs"
)

func doublePointPrecision(num float32) float32 {
	return float32(math.Round(float64((num)*100)) / 100)
}
func singlePointPrecision(num float32) float32 {
	return float32(math.Round(float64((num)*10)) / 10)
}

// prints the stocks based on the herd
func PrintStocks(herd structs.Herd, elapsedDays int) string {

	var str string

	str += fmt.Sprintf("Output for T = %d :\n", elapsedDays)
	milk, wool := ComputeStockAndHerd(herd, elapsedDays)
	str += "In Stock:\n"
	str += fmt.Sprintf("\t%.3f liters of milk\n", milk)
	str += fmt.Sprintf("\t%d skins of wool\n", wool)
	str += "Herd:\n"
	for _, item := range herd.Herd {
		str += fmt.Sprintf("\t%s %.2f years old\n", item.Name, item.Age)
	}

	return str

}

func ComputeStockAndHerd(herd structs.Herd, elapsedDays int) (float32, int) {
	var milk float32
	var wool int

	for i := 0; i < elapsedDays; i++ {
		// iterate over all yaks in my herd
		for j := 0; j < len(herd.Herd); j++ {

			if herd.Herd[j].Age >= 10 {
				continue
			}

			// if day 0 then shave
			if i == 0 {
				herd.Herd[j].AgeLastShaved = singlePointPrecision(herd.Herd[j].Age)
				wool++
			} else if herd.Herd[j].Age == 1 {
				herd.Herd[j].AgeLastShaved = singlePointPrecision(herd.Herd[j].Age)
				wool++
			} else if herd.Herd[j].Age > 1 {
				eligible := math.Ceil(float64(8 + (herd.Herd[j].Age * 100 * 0.01)))
				if i%int(eligible) == 0 {
					herd.Herd[j].AgeLastShaved = singlePointPrecision(herd.Herd[j].Age)
					wool++
				}
			}

			// if female yak then milk
			if herd.Herd[j].Sex == "f" {
				milk += 50 - ((herd.Herd[j].Age * 100) * 0.03)
			}

			// increment yak age
			herd.Herd[j].Age = doublePointPrecision(herd.Herd[j].Age) + 0.01

		}

	}
	return milk, wool

}
