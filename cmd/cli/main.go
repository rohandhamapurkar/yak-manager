package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"upwork-assessment-1/pkg"
	"upwork-assessment-1/structs"
)

func main() {
	var config structs.Config
	jsonFilePath := flag.String("file", "", "json file path")
	elapsedDays := flag.Int("days", 0, "number of elapsed days, default value 0")

	flag.Parse()

	config.ElapsedDays = *elapsedDays

	bytes := pkg.ReadJsonFileBytes(*jsonFilePath)

	err := json.Unmarshal(bytes, &config.Herd)

	if err != nil {
		fmt.Println(err)
		log.Fatalln("Error unmarshaling json file")
	}

	var herd structs.Herd

	// making a copy as nested struct to avoid nested struct to be passed as reference
	herd.Herd = make([]structs.Yak, len(config.Herd.Herd))
	copy(herd.Herd, config.Herd.Herd)

	str := pkg.PrintStocks(herd, config.ElapsedDays)

	fmt.Println(str)

}
