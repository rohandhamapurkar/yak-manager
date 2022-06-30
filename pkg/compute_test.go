package pkg

import (
	"testing"
	"upwork-assessment-1/structs"
)

func TestPrintStocks(t *testing.T) {
	input := structs.Config{
		Herd: structs.Herd{Herd: []structs.Yak{
			{
				Name: "Betty-1",
				Age:  4,
				Sex:  "f",
			},
			{
				Name: "Betty-2",
				Age:  8,
				Sex:  "f",
			},
			{
				Name: "Betty-3",
				Age:  9.5,
				Sex:  "f",
			},
		}},
		ElapsedDays: 13,
	}

	str := PrintStocks(input.Herd, input.ElapsedDays)
	message := "Output for T = 13 :\nIn Stock:\n\t1104.480 liters of milk\n\t3 skins of wool\nHerd:\n\tBetty-1 4.13 years old\n\tBetty-2 8.13 years old\n\tBetty-3 9.63 years old\n"

	if str != message {
		t.Errorf("got %q, wanted %q", str, message)
	}

}

func TestComputeStockAndHerdWith13DaysElapsed(t *testing.T) {
	input := structs.Config{
		Herd: structs.Herd{Herd: []structs.Yak{
			{
				Name: "Betty-1",
				Age:  4,
				Sex:  "f",
			},
			{
				Name: "Betty-2",
				Age:  8,
				Sex:  "f",
			},
			{
				Name: "Betty-3",
				Age:  9.5,
				Sex:  "f",
			},
		}},
		ElapsedDays: 13,
	}

	expectedMilk := float32(1104.480)
	expectedWool := 3

	milk, wool := ComputeStockAndHerd(input.Herd, input.ElapsedDays)

	if milk != expectedMilk || wool != expectedWool {
		t.Errorf("got %f, wanted %f", milk, expectedMilk)
		t.Errorf("got %d, wanted %d", wool, expectedWool)
	}

}

func TestComputeStockAndHerdWith14DaysElapsed(t *testing.T) {
	input := structs.Config{
		Herd: structs.Herd{Herd: []structs.Yak{
			{
				Name: "Betty-1",
				Age:  4,
				Sex:  "f",
			},
			{
				Name: "Betty-2",
				Age:  8,
				Sex:  "f",
			},
			{
				Name: "Betty-3",
				Age:  9.5,
				Sex:  "f",
			},
		}},
		ElapsedDays: 14,
	}

	expectedMilk := float32(1188.8099)
	expectedWool := 4

	milk, wool := ComputeStockAndHerd(input.Herd, input.ElapsedDays)

	if milk != expectedMilk || wool != expectedWool {
		t.Errorf("got %f, wanted %f", milk, expectedMilk)
		t.Errorf("got %d, wanted %d", wool, expectedWool)
	}

}
