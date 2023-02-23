package utils

import (
	"bustimerkic/controller"
	"bustimerkic/domain/model"
	"bustimerkic/infra"
	bustimersqlc "bustimerkic/sqlc/gen"
	"context"
	"fmt"
	"net/http"
	"regexp"
	"strings"

	"github.com/gocolly/colly"
)



func handler(w http.ResponseWriter, r *http.Request) {
	busstop := "円町"
	destination := "立命館大学行き"
	busstoptourlCtrl := controller.BusstoToUrlController{}
	url, err := busstoptourlCtrl.FindURL(busstop, destination)
	if err != nil {

	}
	timetablecontroller := controller.TimetableController{}
	for _,u := range url {
		timetable,_ := timetablecontroller.FindTimetable(u,destination)
		fmt.Printf("%v", timetable)
		fmt.Fprint(w, timetable)
	}
	


	
}

func scrapHTML() (scrapData []string) {
	c := colly.NewCollector()

	scrapData = []string{}

	c.OnHTML(".tt-time", func(e *colly.HTMLElement) {
		result := strings.Replace(e.Text, "(", "", -1)
		result = strings.Replace(result, ")", "", -1)
		re := regexp.MustCompile(`.*台`)
		result = re.ReplaceAllString(result, "")
		// timelist := strings.Split(result, "\n")
		scrapData = append(scrapData, result)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting URL:", r.URL.String())
	})
	c.Visit("https://www2.city.kyoto.lg.jp/kotsu/busdia/hyperdia/091011.htm")
	return scrapData
}

func getViaandBusstop() (via string, busstop string) {
	c := colly.NewCollector()

	Via := ""
	Busstop := ""

	// Extract title element
	c.OnHTML("h1", func(e *colly.HTMLElement) {
		rep := regexp.MustCompile(` |　|\[|\]`)
		e.Text = strings.Replace(e.Text, "-", "", -1)
		result := rep.Split(e.Text, -1)

		Via = result[5]
		Busstop = result[0]

		fmt.Println("駅名 : ", result[0], "\nバス名 : ", result[2], "\n行き先 : ", result[5], "\nURL : ", result[6])
	})
	// Before making a request print "Visiting URL: https://XXX"
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting URL:", r.URL.String())
	})
	c.Visit("https://www2.city.kyoto.lg.jp/kotsu/busdia/hyperdia/091011.htm")

	return Via, Busstop
}

func getTimeTable(scrapedata []string, via string, busstop string) {
	timetable := model.CreateNewTimeTable()

	timelist := []string{}

	for i := 0; i < len(scrapedata); i++ {
		timelist = append(timelist, scrapedata[i])
	}

	fmt.Println(len(timelist))

	Via := via
	for i := 0; i < len(timelist)/3; i++ {

		weekdaylist := strings.Split(timelist[i*3], " ")

		for a := 0; a < len(weekdaylist); a++ {
			if weekdaylist[a] != "" {
				timetable.Weekdays[i+5] = append(timetable.Weekdays[i+5], model.OneBusTime{
					BusName:     Via,
					Min:     weekdaylist[a],
					BusStop: "1番乗り場",
				})
			}
		}

		holidaylist := strings.Split(timelist[i*3+1], " ")
		for b := 0; b < len(holidaylist); b++ {
			if holidaylist[b] != "" {
				timetable.Holidays[i+5] = append(timetable.Holidays[i+5], model.OneBusTime{
					BusName:     Via,
					Min:     holidaylist[b],
					BusStop: "1番乗り場",
				})
			}

		}
	}
	fmt.Printf("%v\n", timetable)
}

func Dbcreate() {
	ctx := context.Background()
	db := infra.GetDB()
	queries := bustimersqlc.New(db)
	idx := "4030410"
		fmt.Println(idx)
		busname, busstop, destination, url := getViaandBusstops(idx)
		if len(busname) != 0 {
			queries.CreateBusstopUrl(ctx,bustimersqlc.CreateBusstopUrlParams{Busstop: busstop, Busname: busname, Destination: destination, Url: url})
		}
}

func getViaandBusstops(index string) (busname string, busstop string, destination string, url string) {
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
		Via := ""

		for i := 0; i < len(result); i++ {
			if strings.Contains(result[i], "行き") {
				Destination = result[i]
			}
		}
		for i := 0; i < len(result); i++ {
			if strings.Contains(result[i], "号") {
				Via = result[i]
			}
		}

		Busname = Via
		Busstop = result[0]

		fmt.Println("駅名 : ", Busstop, "\nバス名 : ", Busname, "\n行き先 : ", Destination, "\nurl : ", URL)
	})
	// Before making a request print "Visiting URL: https://XXX"
	c.OnRequest(func(r *colly.Request) {
		// fmt.Println("Visiting URL:", r.URL.String())
	})
	c.Visit(URL)

	return Busname, Busstop, Destination, URL
}