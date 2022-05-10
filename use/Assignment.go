package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func Assignment() {

	jsonFile, err := os.Open("users.json")
	if err != nil {
		fmt.Println(err)
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var result map[string]interface{}
	json.Unmarshal([]byte(byteValue), &result)

	fmt.Println(result["geojson"])

	// b, err := json.Marshal(geojson)
	// if err != nil {
	// 	log.Fatalf("JSON marshaling failed: %s", err)
	// }

	// fmt.Println(string(b))

	// // jsonFileRead, err := ioutil.ReadFile("links.geojson")
	// byteValue, _ := ioutil.ReadAll(jsonFile)
	// fc, _ := geojson.UnmarshalFeatureCollection(byteValue)
	// fmt.Print(fc)
	// // rawFeatureJSON := []byte(jsonFile)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// // fc, _ := geojson.UnmarshalFeatureCollection(rawFeatureJSON)
	// // fmt.Print(fc)

	// rawFeatureJSON, err := os.Open("links.geojson")
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// // byteValue, _ := ioutil.ReadAll(jsonFile)

	// // var db_info Info

	// // json.Unmarshal(byteValue, &db_info)

	// // fmt.Println(db_info)

	// // defer jsonFile.Close()

	// // rawFeatureJSON := []byte(`
	// // { "type": "FeatureCollection",
	// //   "features": [
	// //     { "type": "Feature",
	// //       "geometry": {"type": "Point", "coordinates": [102.0, 0.5]},
	// //       "properties": {"prop0": "value0"}
	// //     }
	// //   ]
	// // }`)

	// fc1, err := geojson.UnmarshalFeatureCollection(rawFeatureJSON)

	// fc2 := geojson.NewFeatureCollection()
	// err := json.Unmarshal(rawJSON, fc2)

	// // Geometry
	// rawGeometryJSON := []byte(`{"type": "Point", "coordinates": [102.0, 0.5]}`)
	// g, err := geojson.UnmarshalGeometry(rawGeometryJSON)

	// g.IsPoint() == true
	// g.Point == []float64{102.0, 0.5}
}
