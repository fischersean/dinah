package statcast

import (
	"testing"
	"time"
)

func TestDownloadScRange(t *testing.T) {
	t.SkipNow()
	d0 := time.Date(2020, time.July, 30, 23, 0, 0, 0, time.UTC)
	d1 := time.Date(2020, time.July, 31, 23, 0, 0, 0, time.UTC)
	b, err := downloadStatcastRange(d0, d1)

	if err != nil {
		t.Fatalf(err.Error())
	}

	if len(b) == 0 {
		t.Fatalf("No pitches found")
	}

}

func TestDownloadScDay(t *testing.T) {
	d0 := time.Date(2020, time.July, 23, 23, 0, 0, 0, time.UTC)
	b, err := downloadStatcastDay(d0)

	if err != nil {
		t.Fatalf(err.Error())
	}

	if len(b) == 0 {
		t.Fatalf("No pitches found")
	}

	d1 := time.Date(2020, time.February, 23, 23, 0, 0, 0, time.UTC)
	b, err = downloadStatcastDay(d1)

	if err != nil {
		t.Fatalf(err.Error())
	}

	if len(b) != 0 {
		t.Fatalf("Pitches found in day where no games played")
	}

	//t.Logf("%+v", b)

}

func TestUrlEncode(t *testing.T) {
	url := constructUrl(time.Date(2020, time.July, 30, 23, 0, 0, 0, time.UTC), time.Date(2020, time.July, 31, 23, 0, 0, 0, time.UTC))
	//t.Log(url)

	if url != "https://baseballsavant.mlb.com/statcast_search/csv?all=true&game_date_gt=2020-07-30&game_date_lt=2020-07-31&group_by=name&min_pass=0&player_event_sort=api_h_launch_speed&sort_col=pitches&sort_order=desc&type=details" {
		t.Fatalf("Url does not match expected value")
	}
}
