package main

import (
	"io/ioutil"
	"net/http/httptest"
	"practice-colly/config"
	"practice-colly/infra"
	"practice-colly/infra/localcache"
	"testing"
)

func TestHandler(t *testing.T) {
	infra.Init()
	localcache.Init()
	router := Routing()

	for i, busstoplist := range config.AllBusstopList {
		for _, busstop := range busstoplist {
			req := httptest.NewRequest("GET", "/timetable?fr="+busstop+"&to=立命館大学", nil)
			rec := httptest.NewRecorder()

			router.ServeHTTP(rec, req)

			resp := rec.Result()

			_, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				t.Errorf("cannot read test response: %v", err)
			}
			if busstop == "立命館大学前" {
				if resp.StatusCode != 404 {
					t.Errorf("got = %d, want = 404,  %s号系統  %s → %s", resp.StatusCode,config.Busname[i],busstop,"立命館大学")
				}
			} else {
				if resp.StatusCode != 200 {
					t.Errorf("got = %d, want = 200,  %s号系統  %s → %s", resp.StatusCode,config.Busname[i],busstop,"立命館大学")
				}
			}
		}
	}

	for i, busstoplist := range config.AllBusstopList {
		for _, busstop := range busstoplist {
			req := httptest.NewRequest("GET", "/timetable?fr=立命館大学前&to="+busstop, nil)
			rec := httptest.NewRecorder()

			router.ServeHTTP(rec, req)

			resp := rec.Result()

			_, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				t.Errorf("cannot read test response: %v", err)
			}

			if busstop == "立命館大学前" {
				if resp.StatusCode != 404 {
					t.Errorf("got = %d, want = 404,  %s号系統  %s → %s", resp.StatusCode,config.Busname[i],"立命館大学前",busstop)
				}
			} else {
				if resp.StatusCode != 200 {
					t.Errorf("got = %d, want = 200,  %s号系統  %s → %s", resp.StatusCode,config.Busname[i],"立命館大学前",busstop)
				}
			}
		}
	}
}
