package main

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

const (
	filename    = "Placement_keywords.csv"
	outfilename = "output.json"
)

type Placement struct {
	Domain             string  `json:"domain"`
	Url                string  `json:"url"`
	PlacementType      string  `json:"PlacementType"`
	Campaign           string  `json:"campaign"`
	AdGroup            string  `json:"adGroup"`
	Clicks             int     `json:"clicks"`
	Impressions        int     `json:"impressions"`
	Ctr                float64 `json:"ctr"`
	AvgCpc             float64 `json:"avgCpc"`
	Cost               int     `json:"cost"`
	ConversionRate     float64 `json:"conversionRate"`
	CostConversionRate float64 `json:"costConversionRate"`
}

// TODO: make a set of methods to clean up
// strconv shit in newJson func

// newJson takes a string array and maps the values to
// a Placement obj
func newJson(r []string) Placement {
	// convert strings to numbers, ignore errors for now
	// might want validate these later
	cl, _ := strconv.Atoi(r[5])
	im, _ := strconv.Atoi(r[6])
	ct, _ := strconv.ParseFloat(r[7], 64)
	av, _ := strconv.ParseFloat(r[8], 64)
	co, _ := strconv.Atoi(r[9])
	cr, _ := strconv.ParseFloat(r[10], 64)
	ccr, _ := strconv.ParseFloat(r[11], 64)
	p := Placement{
		Domain:             r[0],
		Url:                r[1],
		PlacementType:      r[2],
		Campaign:           r[3],
		AdGroup:            r[4],
		Clicks:             cl,
		Impressions:        im,
		Ctr:                ct,
		AvgCpc:             av,
		Cost:               co,
		ConversionRate:     cr,
		CostConversionRate: ccr,
	}
	return p
}

func csvToJson() []Placement {
	pp := []Placement{}

	// open csv and assign to variable `f`
	f, err := os.Open(filename)
	if err != nil {
		fmt.Print("ERROR: failed to open input file")
		panic(err)
	}
	defer f.Close()

	// parse csv records into a variable `r`, ignore octothorpe
	r := csv.NewReader(f)
	r.Comment = '#'

	// load the records from `r` into an array called `records`
	records, err := r.ReadAll()
	if err != nil {
		fmt.Print("ERROR: failed to read records from csv reader")
		panic(err)
	}

	// create a json obj from each item in the array `records` and append it to
	// the array of Placements
	for _, v := range records {
		j := newJson(v)
		pp = append(pp, j)
	}

	// print summary to screen to verify. the number will be one than the total
	// records in the input file if comments are embedded
	fmt.Printf("records converted: %v", len(pp))
	return pp
}

func writeJson(pp []Placement) {
	// open file for output and assign to variable `f`
	f, err := os.Create(outfilename)
	if err != nil {
		fmt.Print("ERROR: failed to create file")
	}
	writer := bufio.NewWriter(f)
	defer f.Close()

	// loop over array of Placements and write each json to the output file
	for _, v := range pp {
		bytes, err := json.Marshal(v)
		if err != nil {
			fmt.Print("ERROR: failed to marshal placement into json")
		}
		fmt.Fprintf(writer, "%+v\n", string(bytes))
		writer.Flush()
	}

}

func main() {
	writeJson(csvToJson())
}
