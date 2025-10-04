package main

var AS_URL_MAP = make(map[string]map[string]interface{})

var AS_API_PARAMS = []any{ // me when i uh golang
	"gewqgeuqw",
	"wgfeqgegwq",
	123,
	456,
}

func ACCOUNT_SEARCH_UPDATE() {
	AS_URL_MAP = map[string]map[string]interface{}{
		"USER_ID": {
			"usernames":          []string{AS_API_PARAMS[0].(string)},
			"excludeBannedUsers": false,
		},
	}
}

func GETUSERID(USER string) int {
	AS_API_PARAMS = []any{USER}
	ACCOUNT_SEARCH_UPDATE()
	PARAMS := POST_REQUEST("https://users.roblox.com/v1/usernames/users", AS_URL_MAP["USER_ID"])
	DATA := EXTRACT_PARAMS_FROM_DATA(PARAMS, []string{"id"}).(int)
	return DATA
}

func SEARCH_UNBANNED_ACCOUNT(USER string) {
	// USERID := 123
	// if reflect.TypeOf(USER).Kind() == reflect.String {
	// 	USERID = GETUSERID(USER)
	// }
}
