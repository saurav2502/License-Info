/*
 * Copyright (c)  by Saurav from 2022
 */

package main

import (
	"foss/app/service"
	"log"
	"net/http"
)

func main() {
	log.Print("starting container initiated ...")
	http.HandleFunc("/hello", service.HandlePostStart)
	http.HandleFunc("/store", service.HandleStore)
	http.HandleFunc("/read", service.HandleRead)
	http.HandleFunc("/", service.HandleDefault)
	log.Print("container started from port")
	log.Fatal(http.ListenAndServe(":8083", nil))
}
