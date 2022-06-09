/*
 * Copyright (c)  by Saurav from 2022
 */

package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"foss/app/constants"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func HandlePostStart(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		log.Print("GET method called")
		json.NewEncoder(w).Encode("Hello World")
	}
}

type Salary struct {
	Basic, HRA, TA float64
}

type Employee struct {
	FirstName, LastName, Email string
	Age                        int
	MonthlySalary              []Salary
}

// HandleDefault method to get invoked if '/' are called with url.
func HandleDefault(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		http.ServeFile(w, r, "resources/index.html")
	}
}

const DOCKER_PERSISTENT_DIR = "/usr/storage"

//HandleStore to store json data to docker persistent memory
func HandleStore(w http.ResponseWriter, r *http.Request) {
	app1 := AppInfo{
		AppId:       "FossApp-Inst",
		AppName:     "FossApp",
		FileVersion: "1.0",
		LicenseData: "Hello storing this info inside storage",
	}
	app2 := AppInfo{
		AppId:       "EvRatio-Inst",
		AppName:     "Ev Ratio",
		FileVersion: "1.0",
		LicenseData: "Hello storing EvRatio info inside storage",
	}
	app3 := AppInfo{
		AppId:       "Jio-Inst",
		AppName:     "Jio",
		FileVersion: "1.0",
		LicenseData: "Hello storing Jio info inside storage",
	}
	var appInfo Apps
	appInfo = Apps{Infos: []AppInfo{
		app1, app2, app3,
	}}
	/*file, err := json.Marshal(appInfo)
	if err != nil {
		log.Print(err)
	}*/
	file, err := json.MarshalIndent(appInfo, "", " ")
	if err != nil {
		log.Print(err)
	}
	if _, err := os.Stat(DOCKER_PERSISTENT_DIR); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(DOCKER_PERSISTENT_DIR, os.ModePerm)
		if err != nil {
			log.Println(err)
		}
	}
	if err != nil {
		json.NewEncoder(w).Encode("error while creating file")
	}
	err = ioutil.WriteFile(DOCKER_PERSISTENT_DIR+"/data.json", file, 0644)

	if _, err := os.Stat(DOCKER_PERSISTENT_DIR + "/data.json"); err == nil {
		log.Printf("file exist %v", true)

	} else if errors.Is(err, os.ErrNotExist) {
		log.Printf("file exist %v", false)

	} else {
		log.Print("entered in else condition")
	}
	if err != nil {
		log.Print(err)
	}
	/*file, err := os.Create(DOCKER_PERSISTENT_DIR + "/data.txt")
	count, err := file.WriteString(string(byte))
	if err != nil {
		log.Print(err)
	}
	if count > 0 {
		log.Printf("count %v", count)
	}*/
	/*data := Employee{
		FirstName: "Mark",
		LastName:  "Jones",
		Email:     "mark@gmail.com",
		Age:       25,
		MonthlySalary: []Salary{
			Salary{
				Basic: 15000.00,
				HRA:   5000.00,
				TA:    2000.00,
			},
			Salary{
				Basic: 16000.00,
				HRA:   5000.00,
				TA:    2100.00,
			},
			Salary{
				Basic: 17000.00,
				HRA:   5000.00,
				TA:    2200.00,
			},
		},
	}

	file, _ := json.MarshalIndent(data, "", " ")

	_ = ioutil.WriteFile("test.json", file, 0644)*/
	fmt.Fprint(w, "successfully file written")
}

func HandleRead(w http.ResponseWriter, r *http.Request) {
	//dir, _ := os.Getwd()
	//log.Print(dir)
	jsonFile, err := os.Open(DOCKER_PERSISTENT_DIR + "/data.json")
	if err != nil {
		json.NewEncoder(w).Encode("error while opening file")
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.NewEncoder(w).Encode(string(byteValue))
	/*fmt.Printf("data files as json %v", prettyJson(byteValue))
	var apps Apps
	err = json.Unmarshal(byteValue, &apps)
	if err != nil {
		json.NewEncoder(w).Encode("error with unmarshing 1")
	}
	log.Print(apps)
	err = json.NewEncoder(w).Encode(apps)
	if err != nil {
		json.NewEncoder(w).Encode("error with unmarshing 2")
	}
	jsonFile.Close()*/
}

func prettyJson(app interface{}) interface{} {
	marshal, err := json.Marshal(app)
	if err != nil {
		log.Print(err)
	}
	log.Printf("pretty string %v", marshal)
	return string(marshal)
}

//HandleLogin test container calls
func HandleLogin(w http.ResponseWriter, r *http.Request) {
	log.Print("GET :: HandleLogin")
	baseUrl, ok := os.LookupEnv("APP_ADDRESS")
	if !ok {
		log.Print("Address not found")
		_ = fmt.Sprintf("baseurl not found")
	}
	log.Printf("Api Req: %v", baseUrl+constants.AppUrl)
	request, err := http.NewRequest(http.MethodGet, baseUrl+constants.AppUrl, nil)
	if err != nil {
		_ = fmt.Sprintf("preparing request error")
	}
	client := &http.Client{}
	res, err := client.Do(request)
	if err != nil || res == nil {
		_ = fmt.Sprintf("response error")
	}
	bytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		_ = fmt.Sprintf("bytes conversion error with response")
	}
	err = json.NewEncoder(w).Encode(string(bytes))
	if err != nil {
		return
	}
}
