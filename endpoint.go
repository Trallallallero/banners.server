package main

import (
	"encoding/json"
	"log"
	"strconv"

	database "banners.database"
	utils "banners.utils"
)

func GetBanners(language, pageId, deviceId string) string {

	if utils.Sanitize(language) {
		return "Illegal escape sequence found in language input query parameter"
	}

	pageIdValue, pageIdValueError := strconv.Atoi(pageId)
	if pageIdValueError != nil {
		return pageIdValueError.Error()
	}

	deviceIdValue, deviceIdValueError := strconv.Atoi(deviceId)
	if deviceIdValueError != nil {
		return deviceIdValueError.Error()
	}

	database.Seed()

	banners := database.SelectBanners(language, pageIdValue, deviceIdValue)

	zones := database.SelectZones(banners)

	s, marshalError := json.Marshal(zones)
	if marshalError != nil {
		log.Fatal(marshalError)
	}
	return string(s)
}
