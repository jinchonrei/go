package MyJson

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Tests struct {
	List  []string
	Field string
}

func ReadFile() (test Tests) {
	jsonFile, err := os.Open("../MyJson/users.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal([]byte(byteValue), &test)

	for i := 0; i < len(test.List); i++ {
		///	fmt.Println(test.List[i])
	}
	return
}

func main() {
	jsonFile, err := os.Open("users.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var test Tests
	json.Unmarshal([]byte(byteValue), &test)

	for i := 0; i < len(test.List); i++ {
		fmt.Println(test.List[i])
	}
	fmt.Println(test)
	fmt.Println(test.Field)
}
