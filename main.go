package main

import (
	"fmt"
	"net/http"
	"os"
	"practice-colly/controller"
	"practice-colly/domain/model"
	"practice-colly/infra"
	"regexp"
	"strings"
	"github.com/gocolly/colly"
	"github.com/labstack/echo"
)

var e = echo.New()


func main() {
	infra.Init()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	Routing()
	e.Debug = true
	e.Logger.Fatal(e.Start(":" + port))

}



func Routing() {
	e.GET("/", func(c echo.Context) error {
		return c.HTML(http.StatusOK, "<h1>Busdes! Clean Architecture API</h1>")
	})

	e.GET("/timetable/:busstop/:destination",func(c echo.Context) error{
		
		busstop := c.Param("busstop")
		destination := c.Param("destination")
		busstoptourlCtrl := controller.BusstoToUrlController{}
		url,err := busstoptourlCtrl.FindURL(busstop,destination)
		if err != nil {
		}
		timetablecontroller := controller.TimetableController{}
		timetable := timetablecontroller.FindTimetable(url)
		return c.JSON(http.StatusOK,timetable)
	})
}



func handler(w http.ResponseWriter, r *http.Request) {
	busstop := "円町"
	destination := "立命館大学行き"
	busstoptourlCtrl := controller.BusstoToUrlController{}
	url,err := busstoptourlCtrl.FindURL(busstop,destination)
	if err != nil {

	}
	timetablecontroller := controller.TimetableController{}
	timetable := timetablecontroller.FindTimetable(url)

	fmt.Printf("%v",timetable)
	fmt.Fprint(w, timetable)
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

func getTimeTable(scrapedata []string,via string,busstop string) {
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
					Via:     Via,
					Min:     weekdaylist[a],
					BusStop: "1番乗り場",
				})
			}
		}

		holidaylist := strings.Split(timelist[i*3+1], " ")
		for b := 0; b < len(holidaylist); b++ {
			if holidaylist[b] != "" {
				timetable.Holidays[i+5] =  append(timetable.Holidays[i+5], model.OneBusTime{
					Via:     Via,
					Min:     holidaylist[b],
					BusStop: "1番乗り場",
				})
			}

		}
		saturdaylist := strings.Split(timelist[i*3+2], " ")
		for c := 0; c < len(saturdaylist); c++ {
			if saturdaylist[c] != "" {
				timetable.Saturdays[i+5] = append(timetable.Saturdays[i+5], model.OneBusTime{
					Via:     Via,
					Min:     saturdaylist[c],
					BusStop: "1番乗り場",
				})
			}
		}
	}
	fmt.Printf("%v\n", timetable)
}


// func dbcreate(){
// 	db := infra.GetDB()
// 	for i:=3300;i<999999;i++{
// 		index := strconv.Itoa(i)
// 		length := len(index)
// 		idx := ""

// 		switch(length){
// 		case 1:
// 			idx = "00000" + index
// 		case 2:
// 			idx = "0000"  + index
// 		case 3:
// 			idx = "000"   + index
// 		case 4:
// 			idx = "00"    + index
// 		case 5:
// 			idx = "0"     + index
// 		case 6:
// 			idx = index
// 		}
// 		busname,busstop,destination,url := getViaandBusstops(idx)
// 		if len(busname) != 0{
// 			db.Create(&model.BusstopUrl{Busstop: busstop,Busname: busname,Destination: destination,URL: url})
// 		}
		
// 	}

// }



// func getViaandBusstops(index string) (busname string, busstop string, destination string, url string) {
// 	c := colly.NewCollector()

// 	Busname := ""
// 	Busstop := ""
// 	Destination := ""
// 	URL := "https://www2.city.kyoto.lg.jp/kotsu/busdia/hyperdia/" + index + ".htm"

// 	// Extract title element
// 	c.OnHTML("h1", func(e *colly.HTMLElement) {
// 		rep := regexp.MustCompile(` |　|\[|\]`)
// 		e.Text = strings.Replace(e.Text, "-", "", -1)
// 		result := rep.Split(e.Text, -1)

// 		for i:=0;i<len(result);i++ {
// 			if (strings.Contains(result[i], "行き")){
// 				Destination = result[i]
// 			}
// 		}

// 		Busname = result[2]
// 		Busstop = result[0]

// 		fmt.Println("駅名 : ", Busstop, "\nバス名 : ", Busname, "\n行き先 : ", Destination)
// 	})
// 	// Before making a request print "Visiting URL: https://XXX"
// 	c.OnRequest(func(r *colly.Request) {
// 		// fmt.Println("Visiting URL:", r.URL.String())
// 	})
// 	c.Visit(URL)

// 	return Busname,Busstop,Destination,URL
// }