package statcast

import (
	"net/http"
	"net/url"
	"time"
)

func constructUrl(beginDate time.Time, endDate time.Time) (scUrl string) {

	u := &url.URL{
		Scheme: "https",
		Host:   "baseballsavant.mlb.com",
		Path:   "statcast_search/csv",
	}

	v := url.Values{}

	v.Add("all", "true")
	v.Add("game_date_gt", beginDate.Format("2006-01-02"))
	v.Add("game_date_lt", endDate.Format("2006-01-02"))
	v.Add("group_by", "name")
	v.Add("sort_col", "pitches")
	v.Add("player_event_sort", "api_h_launch_speed")
	v.Add("sort_order", "desc")
	v.Add("min_pass", "0")
	v.Add("type", "details")

	u.RawQuery = v.Encode()
	scUrl = u.String()

	return scUrl
}

func downloadStatcastDay(d time.Time) (pitches []Pitch, err error) {

	url := constructUrl(d, d)

	res, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	r := res.Body


	pitches, err = marshalPitches(r)

	if err != nil {
		return nil, err
	}

	return pitches, err
}

func downloadStatcastRange(d0 time.Time, d1 time.Time) (pitches []Pitch, err error) {

	// We can only do 5 days at a time

	increment, _ := time.ParseDuration("24h")
	for day := d0; !day.After(d1); day = day.Add(increment) {

		time.Sleep(time.Duration(500) * time.Millisecond)
		p, err := downloadStatcastDay(day)
		if err != nil {
			return nil, err
		}
		pitches = append(pitches, p...)
	}

	return pitches, err
}

//func downloadStatcastSeason(season int) (pitches []Pitch, err error) {

//pitches, err = downloadStatcastRange(time.Date(season, 1, 1, 0, 0, 0, 0, time.UTC), time.Date(season, 12, 31, 0, 0, 0, 0, time.UTC))

//if err != nil {
//return nil, err
//}

//return pitches, err
//}
