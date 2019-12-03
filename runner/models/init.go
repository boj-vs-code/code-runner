package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetProblemById(id rune) Problem {
	url := fmt.Sprintf("http://boj-vs-code.appspot.com/problem/%d", id)
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var problem Problem
	err = json.Unmarshal(bytes, &problem)
	if err != nil {
		panic(err)
	}

	return problem
}
