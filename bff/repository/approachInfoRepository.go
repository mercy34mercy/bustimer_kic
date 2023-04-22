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

type ApproachInfoRepository interface {
	FindApproachInfo(fr string, to string) (*model.ApproachInfos, error)
}

type ApproachInfoRepositoryImpl struct{}

func NewApproachInfoRepository() ApproachInfoRepository {
	return &ApproachInfoRepositoryImpl{}
}

func (air *ApproachInfoRepositoryImpl) FindApproachInfo(fr string, to string) (*model.ApproachInfos, error) {
	HostName := "https://busdes-kic.com"
	path := "/nextbus"

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
		panic(err)
	}
	defer resp.Body.Close()

	var approachInfos domain.ApproachInfos
	err = json.NewDecoder(resp.Body).Decode(&approachInfos)
	if err != nil {
		return nil, err
	}

	modelApproachInfo := dto.ApproachInfosDto(approachInfos)

	return &modelApproachInfo, nil
}
