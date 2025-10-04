// SECRETSEEKER BY GYROSCOPICURINAL
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"runtime/debug"
	"time"
)

const (
	RATE_LIMIT_PS = 5
)

func ERRHANDLE(ERR any) { // pinnacle of go coding ngl
	if ERR != nil {
		log.Fatal(ERR)
	}
}

func CONVERT_BYTES_TO_MAP(BYTES []byte) map[string]interface{} {
	RESULT := map[string]interface{}{}
	ERR := json.Unmarshal(BYTES, &RESULT)
	ERRHANDLE(ERR)
	return RESULT
}

func POST_REQUEST(URL string, POST_BODY map[string]interface{}) map[string]interface{} {
	JDATA, ERR := json.Marshal(POST_BODY)
	ERRHANDLE(ERR)

	REQ, ERR := http.NewRequest("POST", URL, bytes.NewBuffer(JDATA))
	ERRHANDLE(ERR)
	REQ.Header.Set("Content-Type", "application/json")

	CLIENT := &http.Client{}
	RESP, ERR := CLIENT.Do(REQ)
	ERRHANDLE(ERR)
	defer RESP.Body.Close()

	RESPBODY, ERR := io.ReadAll(RESP.Body)
	ERRHANDLE(ERR)

	BODY := CONVERT_BYTES_TO_MAP(RESPBODY)
	return BODY
}

func GET_REQUEST(URL string) map[string]interface{} {
	RESP, ERR := http.Get(URL)
	ERRHANDLE(ERR)
	defer RESP.Body.Close()

	BBODY, ERR := io.ReadAll(RESP.Body)
	ERRHANDLE(ERR)

	BODY := CONVERT_BYTES_TO_MAP(BBODY)
	return BODY
}

func EXTRACT_PARAMS_FROM_DATA(DATA map[string]interface{}, PARAMS []string) any {
	DATA_ARRAY := make(map[int]map[string]interface{})
	SUB_TBL := 0
	for _, PARAM := range PARAMS {
		if DATA_SLICE, OK := DATA["data"].([]interface{}); OK {
			for _, TABLE := range DATA_SLICE {
				SUB_TBL = SUB_TBL + 1
				for KEY, VALUE := range TABLE.(map[string]interface{}) {
					if KEY == PARAM {
						DATA_ARRAY[SUB_TBL][KEY] = VALUE
					}
				}
			}
		}
	}
	return DATA_ARRAY
}

func ADD_DATA_TO_TBL(DATA map[string]map[string]interface{}, PARAMS map[string]interface{}, FOLDERN string) map[string]map[string]interface{} {
	for KEY, VALUE := range PARAMS {
		DATA[FOLDERN][KEY] = VALUE
	}
	fmt.Println(DATA)
	time.Sleep(time.Second * 5)
	return DATA
}

func main() {
	// if you want to keep previous crash logs, remove "|os.O_TRUNC" in LG_CRASH
	defer func() {
		if r := recover(); r != nil {
			stackTrace := string(debug.Stack())
			LG_CRASH(r, stackTrace)

			fmt.Println("App crashed: ", r)
			fmt.Println("Stack trace: ", stackTrace)
		}
	}()

	INIT_UI()
}

func LG_CRASH(err interface{}, stackTrace string) {
	F, E := os.OpenFile("crash.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if E != nil {
		fmt.Println("Error opening log file:", E)
		return
	}
	defer F.Close()

	log.SetOutput(F)
	log.Printf("Time: %s\nError: %v\nStack Trace:\n%s\n\n", time.Now().Format(time.RFC3339), err, stackTrace)
}
