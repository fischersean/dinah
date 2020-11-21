package statcast

import (
	"encoding/csv"
	"fmt"
	"github.com/gocarina/gocsv"
	"io"
	"strconv"
	"strings"
	"time"
)

type CSVReader struct {
	stdReader *csv.Reader
}

// parseDate parses the string date format present in the Statcast database
func parseDate(dstring string) (d time.Time, err error) {

	comps := strings.Split(dstring, "-")

	if len(comps) == 2 {
		return d, fmt.Errorf("Could not parse date components: %v", comps)
	}

	year, err := strconv.Atoi(comps[0])

	if err != nil {
		return d, err
	}

	month, err := strconv.Atoi(comps[1])

	if err != nil {
		return d, err
	}

	day, err := strconv.Atoi(comps[2])

	if err != nil {
		return d, err
	}

	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC), nil

}

// Read returns a single recrod from a statcast dataset
func (sr *CSVReader) Read() (retSlice []string, err error) {

	retSlice, err = sr.stdReader.Read()

	if err != nil {
		return nil, err
	}

	for i, v := range retSlice {
		if v == "null" {
			retSlice[i] = ""
		}
	}

	return retSlice, err

}

// ReadAll reads all records from a statcast csv dataset
func (sr *CSVReader) ReadAll() (retSlice [][]string, err error) {

	for {
		tmpSlice, err := sr.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			return nil, err
		}

		retSlice = append(retSlice, tmpSlice)
	}

	return retSlice, err

}

func newCSVReader(r io.Reader) *CSVReader {

	csvread := csv.NewReader(r)
	return &CSVReader{
		stdReader: csvread,
	}

}

func marshalPitches(r io.Reader) (pitches []Pitch, err error) {

	scReader := newCSVReader(r)

	if err = gocsv.UnmarshalCSV(scReader, &pitches); err != nil {
		return nil, err
	}

	// Cast all of the date fields to time.Time

	for i := range pitches {
		pitches[i].GameDate, err = parseDate(pitches[i].RawGameDate)
		if err != nil {
			return pitches, err
		}
	}

	return pitches, nil
}
