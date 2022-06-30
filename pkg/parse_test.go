package pkg

import (
	"bytes"
	"io/ioutil"
	"os"
	"testing"
)

func TestJsonParser(t *testing.T) {
	// Open our jsonFile
	jsonFile, _ := os.Open("../test.json")
	defer jsonFile.Close()

	expectedBytes, _ := ioutil.ReadAll(jsonFile)

	resultBytes := ReadJsonFileBytes("../test.json")

	res := bytes.Compare(resultBytes, expectedBytes)

	if res != 0 {
		t.Errorf("Expected and resultant bytes do not match")
	}

}
