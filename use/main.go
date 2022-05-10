package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Info struct {
	Type     string `json:"type"`
	Features []struct {
		Type     string `json:"type"`
		Geometry struct {
			Type        string      `json:"type"`
			Coordinates [][]float64 `json:"coordinates"`
		} `json:"geometry"`
		Properties struct {
			ID int `json:"id"`
		} `json:"properties"`
	} `json:"features"`
}

func main() {
	data, err := os.Open("links.geojson")
	if err != nil {
		fmt.Println(err)
	}
	byteValue, _ := ioutil.ReadAll(data)

	var db_info Info

	json.Unmarshal(byteValue, &db_info)

	fmt.Println(db_info)
}
