package utils

import (
	"github.com/prateek-rajdev/musicRecommendation/song/model"
	"reflect"
)

func getNumFields(obj interface{}) int {
	objType := reflect.TypeOf(obj)

	if objType.Kind() == reflect.Ptr {
		objType = objType.Elem()
	}

	return objType.NumField()
}

func CalculateSimilarityIndex(song1, song2 model.Song) float64 {
	attributes := getNumFields(song1)
	sameAttributes := 0

	if song1.Genre == song2.Genre {
		sameAttributes++
	}
	if song1.Tempo == song2.Tempo {
		sameAttributes++
	}
	if song1.Singer == song2.Singer {
		sameAttributes++
	}
	if song1.PopularityScore == song2.PopularityScore {
		sameAttributes++
	}
	if song1.ReleaseYear == song2.ReleaseYear {
		sameAttributes++
	}

	return float64(sameAttributes) / float64(attributes)
}
