package chadwick

import (
	"github.com/fischersean/dinah/common"
)

var peopleRegister = "https://raw.githubusercontent.com/chadwickbureau/register/master/data/people.csv"

var countryRegister = "https://raw.githubusercontent.com/chadwickbureau/register/master/data/names.csv"

func downloadPeopleRegister() (p []Person, err error) {
	r, err := common.DownloadFileUrl(peopleRegister)

	if err != nil {
		return p, err
	}

	p, err = marshalPeople(r)

	if err != nil {
		return p, err
	}

	return p, nil
}

func downloadCountryRegister() (c []Country, err error) {
	r, err := common.DownloadFileUrl(countryRegister)

	if err != nil {
		return c, err
	}

	c, err = marshalContries(r)

	if err != nil {
		return c, err
	}

	return c, nil
}
