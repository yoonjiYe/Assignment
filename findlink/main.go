package main

import (
	"fmt"
	"io/ioutil"
	"math"

	"github.com/paulmach/orb"
	"github.com/paulmach/orb/geojson"
)

const (
	GEO_FILE = "links.geojson"
)

func main() {
	var longP float64
	var latP float64
	fmt.Print("경도와 위도를 입력하시오. (예) 127.027268062 37.499212063 => ")
	fmt.Scanf("%f %f", &longP, &latP)
	fmt.Print(longP, latP)
	// geojson file to a feature collection
	b, _ := ioutil.ReadFile(GEO_FILE)

	//go에서 사용 가능하게 unmarshal
	featureCollection, _ := geojson.UnmarshalFeatureCollection(b)

	// 최단거리와, 해당 좌표를 계산하기 위한 함수 호출

	DistanceFromCoordi(featureCollection, orb.Point{longP, latP})
}

// DistanceFromCoordi 각 LineString을 모두 탐색하면서 각 line별 입력좌표와의 최단거리 계신
func DistanceFromCoordi(fc *geojson.FeatureCollection, point orb.Point) {
	var minDistance float64 = 0
	var log float64 = 0
	var lat float64 = 0
	//LineString 모두 탐색
	for _, feature := range fc.Features {
		featureCoordi := feature.Geometry.(orb.LineString)
		// 각 LineString의 좌표를 하나씩 이동하면서 line생성
		// 생성된 line을 바탕으로 입력 좌표와의 최단 거리 및 접점좌표 계산
		for j := 0; j < len(featureCoordi)-1; j++ {
			result := Calculate(featureCoordi[j], featureCoordi[j+1], point)
			if minDistance == 0 {
				minDistance = result[0]
				log = result[1]
				lat = result[2]
			}
			if result[0] < minDistance {
				minDistance = result[0]
				log = result[1]
				lat = result[2]
			}
		}
	}
	fmt.Println("거리:", minDistance, "경도:", log, "위도:", lat)
}

// Calculate 두 점간의 거리와, 해당 좌표를 리턴
func Calculate(a, b, point orb.Point) []float64 {
	var result []float64
	x := a[0]
	y := a[1]
	dx := b[0] - x
	dy := b[1] - y

	if dx != 0 || dy != 0 {
		// 직선방정식을 통한 t값 도출
		// t = 닮음을 활용한 비율
		t := ((point[0]-x)*dx + (point[1]-y)*dy) / (dx*dx + dy*dy)

		if t > 1 {
			x = b[0]
			y = b[1]
		} else if t > 0 {
			x += dx * t
			y += dy * t
		}
	}

	// 위,경도 값으로 구한 거리(좌표간의 거리)
	// dx = point[0] - x
	// dy = point[1] - y
	//h := math.Sqrt(dx*dx + dy*dy)

	// 위, 경도 값으로 구한 미터단위 거리
	h := measure(point[1], y, point[0], x)

	result = append(result, h)
	result = append(result, x)
	result = append(result, y)
	return result
}

// 위,경도 값으로 거리를 구하는 함수(미터 단위)
func measure(lat1 float64, lon1 float64, lat2 float64, lon2 float64) float64 {
	var R = 6378.137 // Radius of earth in KM
	var dLat = lat2*math.Pi/180 - lat1*math.Pi/180
	var dLon = lon2*math.Pi/180 - lon1*math.Pi/180
	var a = math.Sin(dLat/2)*math.Sin(dLat/2) +
		math.Cos(lat1*math.Pi/180)*math.Cos(lat2*math.Pi/180)*
			math.Sin(dLon/2)*math.Sin(dLon/2)
	var c = 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	var d = R * c
	return d * 1000
}
