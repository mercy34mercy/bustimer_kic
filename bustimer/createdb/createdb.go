package createdb

import (
	"context"
	"regexp"
	"strconv"
	"strings"

	"github.com/mercy34mercy/bustimer_kic/bustimer/infra"
	bustimersqlc "github.com/mercy34mercy/bustimer_kic/bustimer/sqlc/gen"

	"github.com/gocolly/colly"
)

func dbcreate() {
	ctx := context.Background()
	db := infra.GetDB()
	queries := bustimersqlc.New(db)
	for i := 3300; i < 999999; i++ {
		index := strconv.Itoa(i)
		length := len(index)
		idx := ""

		switch length {
		case 1:
			idx = "00000" + index
		case 2:
			idx = "0000" + index
		case 3:
			idx = "000" + index
		case 4:
			idx = "00" + index
		case 5:
			idx = "0" + index
		case 6:
			idx = index
		}
		busname, busstop, destination, url := getViaandBusstop(idx)
		queries.CreateBusstopUrl(ctx, bustimersqlc.CreateBusstopUrlParams{Busstop: busstop, Busname: busname, Destination: destination, Url: url})
	}

}

func getViaandBusstop(index string) (busname string, busstop string, destination string, url string) {
	c := colly.NewCollector()

	Busname := ""
	Busstop := ""
	Destination := ""
	URL := "https://www2.city.kyoto.lg.jp/kotsu/busdia/hyperdia/" + index + ".htm"

	// Extract title element
	c.OnHTML("h1", func(e *colly.HTMLElement) {
		rep := regexp.MustCompile(` |　|\[|\]`)
		e.Text = strings.Replace(e.Text, "-", "", -1)
		result := rep.Split(e.Text, -1)

		for i := 0; i < len(result); i++ {
			if strings.Contains(result[i], "行き") {
				Destination = result[i]
			}
		}
		Busname = result[2]
		Busstop = result[0]
	})
	// Before making a request print "Visiting URL: https://XXX"
	c.OnRequest(func(r *colly.Request) {
	})
	c.Visit(URL)

	return Busname, Busstop, Destination, URL
}