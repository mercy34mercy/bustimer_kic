package repository

import (
	"encoding/json"
	"fmt"
	"github.com/mercy34mercy/bustimer_kic/bff/domain"
	"github.com/mercy34mercy/bustimer_kic/bff/domain/dto"
	"github.com/mercy34mercy/bustimer_kic/bff/graph/model"
	"net/http"
	"net/url"
)

type TimetableRepository interface {
	FindTimetable(fr string, to string) (*model.TimeTable, error)
}

type TimetableRepositoryImpl struct{}

func NewTimetableRepository() TimetableRepository {
	return &TimetableRepositoryImpl{}
}

func (tr *TimetableRepositoryImpl) FindTimetable(fr string, to string) (*model.TimeTable, error) {
	HostName := "https://busdes-kic.com"
	path := "/timetable"

	urls, err := url.Parse(HostName + path)
	if err != nil {
		return nil, err
	}

	// クエリパラメータを設定
	queryParams := url.Values{}
	queryParams.Set("fr", fr)
	queryParams.Set("to", to)

	urls.RawQuery = queryParams.Encode()

	fmt.Print(urls.String())

	req, err := http.NewRequest("GET", urls.String(), nil)

	client := &http.Client{}

	// HTTPリクエストを送信し、レスポンスを受信
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var timeTable domain.TimeTable

	err = json.NewDecoder(resp.Body).Decode(&timeTable)
	if err != nil {
		return nil, err
	}

	modelTimeTable := dto.TimetableDto(timeTable)

	return &modelTimeTable, nil
}
