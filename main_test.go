package main

// import (
// 	"io/ioutil"
// 	"net/http/httptest"
// 	"testing"
// )

// func TestHelloHandler(t *testing.T) {
// 	router := Routing()

// 	req := httptest.NewRequest("GET", "/timetable?fr=北大路新町&to=立命館大学前", nil)
// 	rec := httptest.NewRecorder()

// 	router.ServeHTTP(rec,req)

// 	resp := rec.Result()

// 	_, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		t.Errorf("cannot read test response: %v", err)
// 	}

// 	if resp.StatusCode != 200 {
// 		t.Errorf("got = %d, want = 200", resp.StatusCode)
// 	}


// }