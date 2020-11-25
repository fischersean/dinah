package chadwick

import (
	"net/http"
)

var peopleRegister = "https://raw.githubusercontent.com/chadwickbureau/register/master/data/people.csv"

var countryRegister = "https://raw.githubusercontent.com/chadwickbureau/register/master/data/names.csv"

func downloadPeopleRegister() (p []Person, err error) {

	res, err := http.Get(peopleRegister)

	if err != nil {
		return nil, err
	}

	p, err = marshalPeople(res.Body)

	if err != nil {
		return p, err
	}

	return p, nil
}

func downloadCountryRegister() (c []Country, err error) {

	res, err := http.Get(countryRegister)

	if err != nil {
		return nil, err
	}

	c, err = marshalContries(res.Body)

	if err != nil {
		return c, err
	}

	return c, nil
}
