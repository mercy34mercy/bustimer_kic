package repositoryimpl

import (
	"fmt"
	"practice-colly/domain/model"
	"practice-colly/domain/repository"
	"practice-colly/infra"
	"regexp"
	"strings"

	"github.com/gocolly/colly"
)

type BusstopToTimetableRepositoryImpl struct{}

func NewBusstopToUrlRepositoryImpl() repository.BusstopToTimetableRepository{
	return &BusstopToTimetableRepositoryImpl{}
}

func (repository *BusstopToTimetableRepositoryImpl) FindURL(busstop string,destination string)([]string,error){
	var err error
	db := infra.GetDB()
	busstopinfo := []model.BusstopUrl{}
	var busstopurl []string
	if err = db.Where("destination = ? AND busstop LIKE ?",destination,"%" +busstop + "%").Find(&busstopinfo).Error; err!= nil {
		//エラーハンドリング
		fmt.Printf("db select Error!!!! err:%v\n", err)
	}
	// fmt.Printf("%v",busstopinfo)
	for _,bus := range busstopinfo{
		// fmt.Println(bus.URL)
		busstopurl = append(busstopurl, bus.URL)
	}
	return busstopurl,err
}

func (repository *BusstopToTimetableRepositoryImpl) FindTimetable(url []string)(model.TimeTable){
	timetable := model.CreateNewTimeTable()

	for _,u := range url {
		scrapdata,via,busstop := scrapHTML(u)
		getTimeTable(timetable,scrapdata,via,busstop)
	}
	timetable.SortOneBusTime()

	return timetable
}

func scrapHTML(url string) (scrapData []string,via string,busstop string) {
	c := colly.NewCollector()

	scrapData = []string{}

	Via := ""
	Busstop := ""

	// Extract title element
	c.OnHTML("h1", func(e *colly.HTMLElement) {
		rep := regexp.MustCompile(` |　|\[|\]`)
		e.Text = strings.Replace(e.Text, "-", "", -1)
		result := rep.Split(e.Text, -1)

		Busstop = result[0]

		for i:=0;i<len(result);i++ {
			if (strings.Contains(result[i], "号")){
				Via = result[i]
			}
		}

		fmt.Println("駅名 : ", result[0], "\nバス名 : ", result[2], "\n行き先 : ", result[5], "\nURL : ", result[6])
	})

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
	c.Visit(url)
	return scrapData,Via,Busstop
}

func getTimeTable(timetable model.TimeTable,scrapedata []string,via string,busstop string)(model.TimeTable) {
	

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
	return timetable
}