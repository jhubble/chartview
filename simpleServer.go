package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"time"
	"math/rand"
	"strings"
)

type DataPoint struct {
	Time  int64 `json:"timestamp"`
	Value int
}

type DataPoints struct {
	Data     []DataPoint
	Datatype string `json:"type"`
}

func main() {

	fmt.Println("Starting server on 8081")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var url = "public/" + r.URL.Path[1:]
		if strings.HasSuffix(url, "/") {
			url += "index.html"
		}
		fmt.Println("Serving: " + url)
		http.ServeFile(w, r, url)
	})

	http.HandleFunc("/data", func(w http.ResponseWriter, r *http.Request) {
		datatype := r.URL.Query().Get("type")
		timeRange := r.URL.Query().Get("range")
		result := ""
                number := 100
		if datatype == "storage" {
			result = generate_data(number,timeRange, -1, 2, datatype)
		} else if datatype == "cpu" {
			result = generate_data(number,timeRange, -5, 5, datatype)
		} else if datatype == "network" {
			result = generate_data(number,timeRange, -7, 7, datatype)
		} else if datatype == "memory" {
			result = generate_data(number,timeRange, -4, 5, datatype)
		} else {
			result = "ERROR. Chart type not recognized";
		}
		fmt.Fprintf(w, "%s", result)
	})
	http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hi")
	})

	log.Fatal(http.ListenAndServe(":8081", nil))

}

// interval is range to display, either "week" or "hour"
// min is the minimum value that a data point can differ from previous (e.g. -5)
// max is the maximum value that a data point can differ from previous (e.g. 10)
func generate_data(number int, interval string, min int, max int, datatype string) string {
	timeRange := 0 // time period to display (in milliseconds)

	if interval == "week" {
		timeRange = 60 * 60 * 24 * 7
	}
	if interval == "hour" {
		timeRange = 60 * 60
	}

	millisBetweenPoints := timeRange / number

	time := time.Date(2017, time.November, 10, 23, 0, 0, 0, time.UTC).Unix()

	// Initial value in lower 50% so there is room to grow
	data := rand.Intn(50)

	points := make([]DataPoint, number)
	for i := range points {
		points[i] = DataPoint{
			time,
			data,
		}
		time += int64(millisBetweenPoints)

		// get random number within a negative to positive range
		data += rand.Intn(max-min+1) + min
		if data > 100 {
			data = 100
		}
		if data < 0 {
			data = 0
		}
	}

	generatedData := DataPoints{points, datatype}

	pagesJson, err := json.Marshal(&generatedData)
	if err != nil {
		log.Fatal("Cannot encode to JSON ", err)
		return "ERROR returning data"
	}

	return string(pagesJson)
}

//CPU usage, disk capacity, memory usage, network traffic,
func generate_cpu() DataPoints {
	number := 200

	time := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC).Unix()

	cpu := rand.Intn(100)

	points := make([]DataPoint, number)
	for i := range points {
		points[i] = DataPoint{
			time,
			cpu,
		}
		time += 1 * 60 * 60 * 24
		cpu += (rand.Intn(20) - 10)
		if cpu > 100 {
			cpu = 100
		}
		if cpu < 0 {
			cpu = 0
		}
	}

	return DataPoints{points, "cpu"}
}
