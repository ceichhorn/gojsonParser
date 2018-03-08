ipackage main

import (
	"encoding/json"
	"fmt"
	"strings"
	"io/ioutil"
	"os"
	//"os/exec"
	"regexp"
	"log"
	"strconv"
)

// Apis struct which contains
// an array of users
type Apis struct {
	Rewrites []Rewrites `json:"url_rewrites"`


}

type Rewrites struct {
  Path string  `json:"path"`
	Match string  `json:"match_pattern"`
}

func main() {

// Open our jsonFile
	jsonFile, err := os.Open("out3.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}


	fmt.Println("Successfully Opened out3.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()


	// read our opened json as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)


