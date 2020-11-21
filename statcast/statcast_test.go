package statcast

import (
	"os"
	"testing"
	"time"
)

//func TestSeasonFromHttp(t *testing.T) {
//ds, err := SeasonFromHttp(2020)

//if err != nil {
//t.Fatalf(err.Error())
//}

//if len(ds.Pitches) == 0 {
//t.Fatalf("No pitch data found")
//}
//}

func TestFromHttp(t *testing.T) {
	d0 := time.Date(2020, time.July, 30, 0, 0, 0, 0, time.UTC)
	d1 := time.Date(2020, time.July, 31, 0, 0, 0, 0, time.UTC)

	ds, err := FromHttp(d0, d1)

	if err != nil {
		t.Fatalf(err.Error())
	}

	if len(ds.Pitches) == 0 {
		t.Fatalf("No pitch data found")
	}
}

func TestFromCsv(t *testing.T) {

	tfile, err := os.Open("../resources/sctestdata.csv")

	if err != nil {
		t.Fatalf(err.Error())
	}

	ds, err := FromCsv(tfile)

	if err != nil {
		t.Fatalf(err.Error())
	}
	if len(ds.Pitches) == 0 {
		t.Fatalf("Could not parse csv")
	}

}

func TestParseDate(t *testing.T) {
	d, err := parseDate("2020-09-21")
	testDate := time.Date(2020, time.Month(9), 21, 0, 0, 0, 0, time.UTC)

	if err != nil {
		t.Fatalf(err.Error())
	}

	if !d.Equal(testDate) {
		t.Logf("%+v", d)
		t.Logf("%+v", testDate)
		t.Fatalf("Date not parsed correctly")
	}
}

func TestAppend(t *testing.T) {

	tfile, err := os.Open("../resources/sctestdata.csv")

	if err != nil {
		t.Fatalf(err.Error())
	}

	ds, err := FromCsv(tfile)

	if err != nil {
		t.Fatalf(err.Error())
	}

	initLen := len(ds.Pitches)
	ds.Append(ds)

	if 2*initLen != len(ds.Pitches) {
		t.Fatalf("Append not successful")
	}

}
