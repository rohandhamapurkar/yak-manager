package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"upwork-assessment-1/pkg"
	"upwork-assessment-1/structs"

	"github.com/gin-gonic/gin"
)

func main() {

	var config structs.Config
	jsonFilePath := flag.String("file", "", "json file path")

	flag.Parse()

	bytes := pkg.ReadJsonFileBytes(*jsonFilePath)

	err := json.Unmarshal(bytes, &config.Herd)

	if err != nil {
		fmt.Println(err)
		log.Fatalln("Error unmarshaling json file")
	}

	router := gin.Default()

	router.GET("/yak-shop/stock/:T", func(c *gin.Context) {
		elapsedDays, err := strconv.Atoi(c.Param("T"))
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid elapsed days value",
			})
		}
		var herd structs.Herd

		// making a copy as nested struct to avoid nested struct to be passed as reference
		herd.Herd = make([]structs.Yak, len(config.Herd.Herd))
		copy(herd.Herd, config.Herd.Herd)

		milk, wool := pkg.ComputeStockAndHerd(herd, elapsedDays)

		c.JSON(http.StatusOK, gin.H{
			"milk":  milk,
			"skins": wool,
		})
	})

	router.GET("/yak-shop/herd/:T", func(c *gin.Context) {
		elapsedDays, err := strconv.Atoi(c.Param("T"))
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid elapsed days value",
			})
		}
		var herd structs.Herd

		// making a copy as nested struct to avoid nested struct to be passed as reference
		herd.Herd = make([]structs.Yak, len(config.Herd.Herd))
		copy(herd.Herd, config.Herd.Herd)

		pkg.ComputeStockAndHerd(herd, elapsedDays)

		c.JSON(http.StatusOK, gin.H{
			"herd": herd.Herd,
		})
	})

	err = router.Run(":8000")

	if err != nil {
		fmt.Println(err)
		log.Fatalln("Error running server")
	}

	fmt.Println("Running server on port :8000")

}
