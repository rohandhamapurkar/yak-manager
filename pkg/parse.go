package pkg

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func ReadJsonFileBytes(jsonFilePath string) []byte {
	if jsonFilePath == "" {
		log.Fatalln("Json file path not provided")
	}

	// Open our jsonFile
	jsonFile, err := os.Open(jsonFilePath)
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
		log.Fatalln("Error opening json file")
	}
	// fmt.Println("Successfully Opened json")

	// close json file after function end
	defer jsonFile.Close()

	bytes, err := ioutil.ReadAll(jsonFile)

	if err != nil {
		fmt.Println(err)
		log.Fatalln("Error reading json file")
	}

	return bytes
}
