package application

import (
	"strings"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func getTotalForYear(result []primitive.M) (total float64, nightTotal float64) {
	var first, last float64 = 0, 0
	var nfirst, nlast float64 = 0, 0
	firstDay, firstNight := false, false

	for _, val := range result {
		if strings.ToLower(val["rate"].(string)) == "day" {
			if !firstDay {
				first = val["reading"].(float64)
				firstDay = true
			}

			last = val["reading"].(float64)
		}

		if strings.ToLower(val["rate"].(string)) == "night" {
			if !firstNight {
				nfirst = val["reading"].(float64)
				firstNight = true
			}

			nlast = val["reading"].(float64)
		}
	}

	return (last - first), (nlast - nfirst)
}
