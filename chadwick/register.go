package chadwick

import (
	"os"
)

type Person struct {
	KeyPerson       string `csv:"key_person"`
	KeyUuid         string `csv:"key_uuid"`
	KeyMlbam        int    `csv:"key_mlbam"`
	KeyRetro        string `csv:"key_retro"`
	KeyBbref        string `csv:"key_bbref"`
	KeyBbrefMinors  string `csv:"key_bbref_minors"`
	KeyFangraphs    int    `csv:"key_fangraphs"`
	KeyNpb          string `csv:"key_npb"`
	KeySrNfl        string `csv:"key_sr_nfl"`
	KeySrNba        string `csv:"key_sr_nba"`
	KeySrNhl        string `csv:"key_sr_nhl"`
	KeyFindagrave   int    `csv:"key_findagrave"`
	NameLast        string `csv:"name_last"`
	NameFirst       string `csv:"name_first"`
	NameGiven       string `csv:"name_given"`
	NameSuffix      string `csv:"name_suffix"`
	NameMatrilineal string `csv:"name_matrilineal"`
	NameNick        string `csv:"name_nick"`
	BirthYear       int    `csv:"birth_year"`
	BirthMonth      int    `csv:"birth_month"`
	BirthDay        int    `csv:"birth_day"`
	DeathYear       int    `csv:"death_year"`
	DeathMonth      int    `csv:"death_month"`
	DeathDay        int    `csv:"death_day"`
	ProPlayedFirst  int    `csv:"pro_played_first"`
	ProPlayedLast   int    `csv:"pro_played_last"`
	MlbPlayedFirst  int    `csv:"mlb_played_first"`
	MlbPlayedLast   int    `csv:"mlb_played_last"`
	ColPlayedFirst  int    `csv:"col_played_first"`
	ColPlayedLast   int    `csv:"col_played_last"`
	ProManagedFirst int    `csv:"pro_managed_first"`
	ProManagedLast  int    `csv:"pro_managed_last"`
	MlbManagedFirst int    `csv:"mlb_managed_first"`
	MlbManagedLast  int    `csv:"mlb_managed_last"`
	ColManagedFirst int    `csv:"col_managed_first"`
	ColManagedLast  int    `csv:"col_managed_last"`
	ProUmpiredFirst int    `csv:"pro_umpired_first"`
	ProUmpiredLast  int    `csv:"pro_umpired_last"`
	MlbUmpiredFirst int    `csv:"mlb_umpired_first"`
	MlbUmpiredLast  int    `csv:"mlb_umpired_last"`
}

type Country struct {
	KeyIsoAlpha2 string `csv:"key_iso_alpha2"`
	KeyIsoAlpha3 string `csv:"key_iso_alpha3"`
	KeyIoc       string `csv:"key_ioc"`
	KeyFifa      string `csv:"key_fifa"`
	NameFull     string `csv:"name_full_len"`
}

// PeopleFromCsv loads a CSV in the schema of chadwick's people regiter from the path provided
func PeopleFromCsv(path string) (p []Person, err error) {

	f, err := os.Open(path)

	if err != nil {
		return p, err
	}

	p, err = marshalPeople(f)

	if err != nil {
		return p, err
	}

	return p, nil
}

// PeopleFromHttp downloads a full copy of the chadwick people register dataset
func PeopleFromHttp() (p []Person, err error) {

	p, err = downloadPeopleRegister()

	if err != nil {
		return p, err
	}

	return p, nil
}

// CountriesFromCsv loads a country dataset CSV from the provided path
func CountriesFromCsv(path string) (c []Country, err error) {

	f, err := os.Open(path)

	if err != nil {
		return c, err
	}

	c, err = marshalContries(f)

	if err != nil {
		return c, err
	}
	return c, nil
}

// CountriesFromHttp downloads a copy od chadwick's country name dataaset
func CountriesFromHttp() (c []Country, err error) {

	c, err = downloadCountryRegister()

	if err != nil {
		return c, err
	}
	return c, err
}
