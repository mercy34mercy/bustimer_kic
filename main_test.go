package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http/httptest"
	"practice-colly/config"
	"practice-colly/domain/model"
	"practice-colly/infra"
	"practice-colly/infra/localcache"
	"testing"
)

func TestHandler(t *testing.T) {
	infra.Init()
	localcache.Init()
	router := Routing()

	for i, busstoplist := range config.BusstoptoRitsList {
		for _, busstop := range busstoplist {
			req := httptest.NewRequest("GET", "/timetable?fr="+busstop+"&to=立命館大学", nil)
			rec := httptest.NewRecorder()

			router.ServeHTTP(rec, req)

			resp := rec.Result()

			timetable := model.CreateNewTimeTable()
			if resp.StatusCode != 404 {
				if err := json.NewDecoder(rec.Body).Decode(&timetable); err != nil {
					t.Fatal(err)
				}
			}

			_, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				t.Errorf("cannot read test response: %v", err)
			}

			if busstop == "立命館大学前" {
				if resp.StatusCode != 404 {
					t.Errorf("got = %d, want = 404,  %s  %s → %s", resp.StatusCode, config.BusnameToRits[i], busstop, "立命館大学")
				}
			} else {
				flag := true

				for _, bustime := range timetable.Weekdays {
					for _, time := range bustime {
						if(config.BusnameToRits[i] == "52号系統" || config.BusnameToRits[i] == "55号系統"){
							if time.BusName == config.BusnameToRits[i] || time.BusName == "52・55号系統"{
								flag = false
							}
						}
						
						if(config.BusnameToRits[i] == "50号系統" || config.BusnameToRits[i] == "15号系統"){
							if time.BusName == config.BusnameToRits[i] || time.BusName == "15・50号系統"{
								flag = false
							}
						}
						if time.BusName == config.BusnameToRits[i] {
							flag = false
						}
					}
				}
	
				if flag {
					t.Errorf("notfound %s  %s → %s", config.BusnameToRits[i], busstop, "立命館大学")
				}
				if resp.StatusCode != 200 {
					t.Errorf("got = %d, want = 200,  %s  %s → %s", resp.StatusCode, config.BusnameToRits[i], busstop, "立命館大学")
				}
			}
		}
	}

	for i, busstoplist := range config.BusstopfromRitsList {
		for _, busstop := range busstoplist {
			req := httptest.NewRequest("GET", "/timetable?fr=立命館大学前&to="+busstop, nil)
			rec := httptest.NewRecorder()

			router.ServeHTTP(rec, req)

			resp := rec.Result()

			timetable := model.CreateNewTimeTable()
			if resp.StatusCode != 404 {
				if err := json.NewDecoder(rec.Body).Decode(&timetable); err != nil {
					t.Fatal(err)
				}
			}

			_, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				t.Errorf("cannot read test response: %v", err)
			}



			if busstop == "立命館大学前" {
				if resp.StatusCode != 404 {
					t.Errorf("got = %d, want = 404,  %s  %s → %s", resp.StatusCode, config.BusnameFromRits[i], "立命館大学前", busstop)
				}
			} else {
				if resp.StatusCode != 200 {
					t.Errorf("got = %d, want = 200,  %s  %s → %s", resp.StatusCode, config.BusnameFromRits[i], "立命館大学前", busstop)
				}
				flag := true

				for _, bustime := range timetable.Weekdays {
					for _, time := range bustime {
						if time.BusName == config.BusnameFromRits[i] {
							flag = false
						}
					}
				}
	
				if flag {
					t.Errorf("notfound %s  %s → %s", config.BusnameFromRits[i], "立命館大学", busstop)
				}
			}
		}
	}
}
