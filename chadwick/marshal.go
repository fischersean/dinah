package chadwick

import (
	"encoding/csv"
	"github.com/gocarina/gocsv"
	"io"
)

func marshalPeople(r io.Reader) (p []Person, err error) {

	pReader := csv.NewReader(r)

	if err = gocsv.UnmarshalCSV(pReader, &p); err != nil {
		return nil, err
	}

	return p, nil
}

func marshalContries(r io.Reader) (c []Country, err error) {

	cReader := csv.NewReader(r)

	if err = gocsv.UnmarshalCSV(cReader, &c); err != nil {
		return nil, err
	}

	return c, nil
}
