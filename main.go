package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type queryParams struct {
	startDate string
	endDate   string
	limit     string
}

func parseFlags() queryParams {
	now := time.Now()
	oneWeekAgo := now.AddDate(0, 0, -7)
	dateFormat := "2006-01-02T00:00:00"

	var startDate, endDate, limit string
	flag.StringVar(&startDate, "startDate", oneWeekAgo.Format(dateFormat), "Set the earliest entity formation date to return.")
	flag.StringVar(&endDate, "endDate", now.Format(dateFormat), "Set the latest entity formation date to return.")
	flag.StringVar(&limit, "limit", "10", "Number of records to retrieve.")
	flag.Parse()

	params := queryParams{
		startDate: startDate,
		endDate:   endDate,
		limit:     limit,
	}
	return params
}

func main() {

	params := parseFlags()
	endpoint := "https://data.colorado.gov/resource/4eit-nuxn.json?$where=entityformdate>'" + params.startDate + "'%20AND%20entityformdate<'" + params.endDate + "'&$limit=" + params.limit
	response, error := http.Get(endpoint)

	body, error := ioutil.ReadAll(response.Body)

	if error != nil {
		log.Fatal(error)
	}

	fmt.Printf("%s", body)
	fmt.Println(endpoint)
}
