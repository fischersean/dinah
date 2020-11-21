package statcast

import (
	"io"
	"time"
)

type Pitch struct {
	PitchType                    string    `csv:"pitch_type"`
	RawGameDate                  string    `csv:"game_date"`
	GameDate                     time.Time `csv:"-"`
	ReleaseSpeed                 float64   `csv:"release_speed"`
	ReleasePosX                  float64   `csv:"release_pos_x"`
	ReleasePosZ                  float64   `csv:"release_pos_z"`
	PlayerName                   string    `csv:"player_name"`
	Batter                       int       `csv:"batter"`
	Pitcher                      int       `csv:"pitcher"`
	Events                       string    `csv:"events"`
	Description                  string    `csv:"description"`
	Zone                         int       `csv:"zone"`
	Des                          string    `csv:"des"`
	GameType                     string    `csv:"game_type"`
	Stand                        string    `csv:"stand"`
	PThrows                      string    `csv:"p_throws"`
	HomeTeam                     string    `csv:"home_team"`
	AwayTeam                     string    `csv:"away_team"`
	Type                         string    `csv:"type"`
	HitLocation                  int       `csv:"hit_location"`
	BbType                       string    `csv:"bb_type"`
	Balls                        int       `csv:"balls"`
	Strikes                      int       `csv:"strikes"`
	GameYear                     int       `csv:"game_year"` // Might need to change this to datetime obj
	PfxX                         float64   `csv:"pfx_x"`
	PfxZ                         float64   `csv:"pfx_z"`
	PlateX                       float64   `csv:"plate_x"`
	PlateZ                       float64   `csv:"plate_z"`
	On3B                         int       `csv:"on_3b"`
	On2B                         int       `csv:"on_2b"`
	On1B                         int       `csv:"on_1b"`
	OutsWhenUp                   int       `csv:"outs_when_up"`
	Inning                       int       `csv:"inning"`
	InningTopbot                 string    `csv:"inning_topbot"`
	HcX                          float64   `csv:"hc_x"`
	HcY                          float64   `csv:"hc_y"`
	Fielder2                     int       `csv:"fielder_2"`
	Umpire                       int       `csv:"umpire"`
	Vx0                          float64   `csv:"vx0"`
	Vy0                          float64   `csv:"vy0"`
	Vz0                          float64   `csv:"vz0"`
	Ax                           float64   `csv:"ax"`
	Ay                           float64   `csv:"ay"`
	Az                           float64   `csv:"az"`
	SzTop                        float64   `csv:"sz_top"`
	SzBot                        float64   `csv:"sz_bot"`
	HitDistanceSc                int       `csv:"hit_distance_sc"`
	LaunchSpeed                  float64   `csv:"launch_speed"`
	LaunchAngle                  float64   `csv:"launch_angle"`
	EffectiveSpeed               float64   `csv:"effective_speed"`
	ReleaseSpinRate              int       `csv:"release_spin_rate"`
	ReleaseExtension             float64   `csv:"release_extension"`
	GamePk                       int       `csv:"game_pk"`
	Pitcher1                     int       `csv:"pitcher.1"`
	Fielder21                    int       `csv:"fielder_2.1"`
	Fielder3                     int       `csv:"fielder_3"`
	Fielder4                     int       `csv:"fielder_4"`
	Fielder5                     int       `csv:"fielder_5"`
	Fielder6                     int       `csv:"fielder_6"`
	Fielder7                     int       `csv:"fielder_7"`
	Fielder8                     int       `csv:"fielder_8"`
	Fielder9                     int       `csv:"fielder_9"`
	ReleasePosY                  float64   `csv:"release_pos_y"`
	EstimatedBaUsingSpeedangle   float64   `csv:"estimated_ba_using_speedangle"`
	EstimatedWobaUsingSpeedangle float64   `csv:"estimated_woba_using_speedangle"`
	WobaValue                    float64   `csv:"woba_value"`
	WobaDenom                    float64   `csv:"woba_denom"`
	BabipValue                   float64   `csv:"babip_value"`
	IsoValue                     float64   `csv:"iso_value"`
	LaunchSpeedAngle             int       `csv:"launch_speed_angle"`
	AtBatNumber                  int       `csv:"at_bat_number"`
	PitchNumber                  int       `csv:"pitch_number"`
	PitchName                    string    `csv:"pitch_name"`
	HomeScore                    int       `csv:"home_score"`
	AwayScore                    int       `csv:"away_score"`
	BatScore                     int       `csv:"bat_score"`
	FldScore                     int       `csv:"fld_score"`
	PostAwayScore                int       `csv:"post_away_score"`
	PostHomeScore                int       `csv:"post_home_score"`
	PostBatScore                 int       `csv:"post_bat_score"`
	PostFldScore                 int       `csv:"post_fld_score"`
	IfFieldingAlignment          string    `csv:"if_fielding_alignment"`
	OfFieldingAlignment          string    `csv:"of_fielding_alignment"`
}

type Dataset struct {
	Pitches   []Pitch
	DateBegin time.Time
	DateEnd   time.Time
}

func findMaxDate(p []Pitch) (d time.Time, err error) {

	maxDate := time.Now()

	for _, v := range p {
		tmpDate := v.GameDate

		if tmpDate.After(maxDate) {
			maxDate = tmpDate
		}
	}

	return d, nil

}

func findMinDate(p []Pitch) (d time.Time, err error) {

	minDate := time.Now()

	for _, v := range p {
		tmpDate := v.GameDate

		if tmpDate.Before(minDate) {
			minDate = tmpDate
		}
	}

	return d, nil

}

// FromCsv parses a local csv and returns a Statcast dataset
func FromCsv(r io.Reader) (ds Dataset, err error) {

	ds.Pitches, err = marshalPitches(r)

	if err != nil {
		return ds, err
	}

	ds.DateBegin, err = findMinDate(ds.Pitches)

	if err != nil {
		return ds, err
	}

	ds.DateEnd, err = findMaxDate(ds.Pitches)

	if err != nil {
		return ds, err
	}

	return ds, nil
}

//// FromHttp downloads and parses a Statcast dataset spanning the states d0 and d1
//func SeasonFromHttp(season int) (ds Dataset, err error) {
//ds.Pitches, err = downloadStatcastSeason(season)

//if err != nil {
//return ds, err
//}

//ds.DateBegin, err = findMinDate(ds.Pitches)

//if err != nil {
//return ds, err
//}

//ds.DateEnd, err = findMaxDate(ds.Pitches)

//if err != nil {
//return ds, err
//}

//return ds, nil
//}

// FromHttp downloads and parses a Statcast dataset spanning the states d0 and d1
func FromHttp(d0 time.Time, d1 time.Time) (ds Dataset, err error) {
	ds.Pitches, err = downloadStatcastRange(d0, d1)

	if err != nil {
		return ds, err
	}

	ds.DateBegin = d0
	ds.DateEnd = d1

	return ds, nil
}

// Append combines two statcast datasets into a single struct
func (ds *Dataset) Append(nds Dataset) {
	ds.Pitches = append(ds.Pitches, nds.Pitches...)

	if nds.DateBegin.Before(ds.DateBegin) {
		ds.DateBegin = nds.DateBegin
	}

	if nds.DateEnd.After(ds.DateEnd) {
		ds.DateEnd = nds.DateEnd
	}

}
