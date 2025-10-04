package main

import (
	"fmt"
	"strconv"
	"time"
)

// API PLACE_HOLDERS
var PARAMS map[string]interface{}
var SPARAMS map[string]interface{}

// API PLACE_HOLDERS

var GS_URL_MAP = make(map[string]map[string]interface{})

var GS_API_PARAMS = []any{ // me when i uh golang
	"gewqgeuqw",
	"wgfeqgegwq",
	123,
	456,
}

func GAME_SEARCH_UPDATE() {
	GS_URL_MAP = map[string]map[string]interface{}{}
}

func GET_UNIVERSE_ID(ID int) int {
	URL_ID := strconv.Itoa(ID)
	PARAMS = GET_REQUEST("https://apis.roblox.com/universes/v1/places/" + URL_ID + "/universe")
	UID := F64_TO_INT(PARAMS["universeId"].(float64))
	return UID
}

func SEARCH_GAME(ID int) map[string]map[string]interface{} {
	INFO_ARRAY := make(map[string]map[string]interface{})
	UID := GET_UNIVERSE_ID(ID)
	GUID := strconv.Itoa(UID)
	// SUBPLACES
	SPARAMS = GET_REQUEST("https://develop.roblox.com/v1/universes/" + GUID + "/places")
	SDATA := EXTRACT_PARAMS_FROM_DATA(SPARAMS, []string{"id", "name"})
	S_EXT_DATA := ADD_DATA_TO_TBL(INFO_ARRAY, SDATA.(map[string]interface{}), "SUBPLACES")
	INFO_ARRAY = S_EXT_DATA
	fmt.Println(S_EXT_DATA)
	time.Sleep(time.Second * 5)
	// SUBPLACES
	return INFO_ARRAY
}
