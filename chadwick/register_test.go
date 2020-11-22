package chadwick

import (
	"testing"
)

func TestFromCsv(t *testing.T) {
	p, err := PeopleFromCsv("../resources/cwickpeopletest.csv")

	if err != nil {
		t.Fatalf(err.Error())
	}

	if len(p) != 49 {
		t.Fatalf("Count in sample not expected: %d", len(p))
	}

	c, err := CountriesFromCsv("../resources/cwcountrytest.csv")

	if err != nil {
		t.Fatalf(err.Error())
	}

	if len(c) != 20 {
		t.Fatalf("Count in sample not expected: %d", len(c))
	}

}

func TestFromHttp(t *testing.T) {

	c, err := CountriesFromHttp()

	if err != nil {
		t.Fatalf(err.Error())
	}

	if len(c) == 0 {
		t.Fatalf("Count in sample not expected: %d", len(c))
	}

	oldUrl := peopleRegister

	peopleRegister = "https://raw.githubusercontent.com/fischersean/dinah/main/resources/cwickpeopletest.csv"
	//t.SkipNow() // We already know this works so we'll skipp to prevent downloading 45+MB file
	p, err := PeopleFromHttp()

	if err != nil {
		t.Fatalf(err.Error())
	}

	if len(p) == 0 {
		t.Fatalf("Count in sample not expected: %d", len(p))
	}

	peopleRegister = oldUrl

}
