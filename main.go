package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type queryParams struct {
	entityformdate string
}

func main() {
	query := queryParams{
		entityformdate: "2018-11-08",
	}

	endpoint := "https://data.colorado.gov/resource/4eit-nuxn.json?$where=entityformdate>'" + query.entityformdate + "'"

	response, error := http.Get(endpoint)

	body, error := ioutil.ReadAll(response.Body)

	if error != nil {
		log.Fatal(error)
	}

	fmt.Printf("%s", body)
}
