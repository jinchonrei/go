package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type users struct {
	list []string
}

func main() {
	fmt.Println("fdas")
	jsonFile, err := os.Open("users.json")
	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var result = []byte(byteValue)

	var f interface{}
	json.Unmarshal(result, &f)

	m := f.(map[string]interface{})

	for k, v := range m {
		switch vv := v.(type) {
		case string:
			fmt.Println(k, vv)
		case []interface{}:
			for i, u := range vv {
				fmt.Println(k, i, u)
			}
		}
	}

}
