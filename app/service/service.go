/*
 * Copyright (c)  by Saurav from 2022
 */

package service

import (
	"encoding/json"
	"fmt"
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

const DOCKER_PERSISTENT_DIR = "/resources"

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
	err = ioutil.WriteFile("data.json", file, 0644)
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
	getwd, _ := os.Getwd()
	log.Print(getwd)
	jsonFile, err := os.Open(getwd + "/data.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	fmt.Printf("data files as json %v", prettyJson(byteValue))
	var apps Apps
	json.Unmarshal(byteValue, &apps)
	log.Print(apps)
	json.NewEncoder(w).Encode(apps)
}

func prettyJson(app interface{}) interface{} {
	marshal, err := json.Marshal(app)
	if err != nil {
		log.Print(err)
	}
	log.Printf("pretty string %v", marshal)
	return marshal
}
