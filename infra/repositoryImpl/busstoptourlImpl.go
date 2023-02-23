package repositoryimpl

import (
	"bustimerkic/config"
	"bustimerkic/domain/model"
	"bustimerkic/domain/repository"
	"bustimerkic/infra"
	"bustimerkic/sqlc/gen"
	"context"
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
)

type BusstopToTimetableRepositoryImpl struct{}

func NewBusstopToUrlRepositoryImpl() repository.BusstopToTimetableRepository {
	return &BusstopToTimetableRepositoryImpl{}
}

func (repository *BusstopToTimetableRepositoryImpl) CreateMultiTimetable(timetable []model.TimeTableandDestination, destinationlist []string) model.MultiTimeTable {
	multitimetable := model.CreateNewMultiTimeTable(destinationlist)

	for _, time := range timetable {
		multitimetable.TimeTable[time.Destination] = time.TimeTable
	}
	return multitimetable
}

func (repository *BusstopToTimetableRepositoryImpl) FindURLFromBusstop(busstop string, destination string) ([]string,error) {
	//立命館大学からどこかへ行くとき
	ctx := context.Background()
	db := infra.GetDB()
	queries := bustimersqlc.New(db)
	var busstopurl []string

	destination = strings.Replace(destination, "行き", "", -1)

	//山越中町専用
	if busstop == "山越中町" && destination == "立命館大学" {
		busstopinfo,err := queries.GetBusinfoFromDestination(ctx,bustimersqlc.GetBusinfoFromDestinationParams{Busstop: busstop,Destination:"四条河原町・三条京阪行き" })
		if err != nil {
			return nil,err
		}
		for _, info := range busstopinfo {
			busstopurl = append(busstopurl, info.Url)
		}
	} else if destination == "山越中町" && busstop == "立命館大学前" {
		busstopinfo,err := queries.GetBusinfoFromDestination(ctx,bustimersqlc.GetBusinfoFromDestinationParams{Busstop: busstop,Destination:"宇多野･山越行き"})
		if err != nil {
			return nil,err
		}

		for _, info := range busstopinfo {
			busstopurl = append(busstopurl, info.Url)
		}
	}

	//M1と12番のコンフリクト問題解消
	for _, bus := range config.M1or12or59busstopList {
		if busstop == bus && (destination == "立命館大学") {
			var destinationList [7]string = [7]string{"原谷行き", "金閣寺・立命館大学行き", "立命館大学行き", "金閣寺･竜安寺・山越行き", "山越中町行き", "竜安寺・山越行き", "宇多野･山越行き"}
			for _, des := range destinationList {
				busstopinfo,err := queries.GetBusinfoFromDestination(ctx,bustimersqlc.GetBusinfoFromDestinationParams{Busstop: busstop,Destination: des})
				if err != nil {
					return nil,err
				}
				for _, info := range busstopinfo {
					busstopurl = append(busstopurl, info.Url)
				}
			}
			if err := validateUrl(busstopurl); err != nil {
				return nil,err
			}
			return busstopurl,nil
		} else if busstop == bus && destination == "三条京阪前" {
			var destinationList [1]string = [1]string{"四条河原町・三条京阪行き"}
			for _, des := range destinationList {
				busstopinfo,err := queries.GetBusinfoFromDestination(ctx,bustimersqlc.GetBusinfoFromDestinationParams{Busstop: busstop,Destination: des})
				if err != nil {
					return nil,err
				}
				for _, info := range busstopinfo {
					busstopurl = append(busstopurl, info.Url)
				}
			}
			if err := validateUrl(busstopurl); err != nil {
				return nil,err
			}
			return busstopurl,nil
		} else if busstop == "立命館大学前" && destination == bus {
			var destinationList [7]string = [7]string{"原谷行き", "金閣寺・立命館大学行き", "立命館大学行き", "金閣寺･竜安寺・山越行き", "山越中町行き", "竜安寺・山越行き", "宇多野･山越行き"}
			for _, des := range destinationList {
				busstopList,err := queries.GetBusinfoFromDestination(ctx,bustimersqlc.GetBusinfoFromDestinationParams{Busstop: destination,Destination: des})
				if err != nil {
					return nil,err
				}

				for _, info := range busstopList {
					if info.Busname == "52・55号系統" {
						for _, splitbus := range config.Busname52and55 {
							busstopinfo,err := queries.GetBusinfoFromBusname(ctx,bustimersqlc.GetBusinfoFromBusnameParams{Busname: splitbus,Busstop: busstop})
							if err != nil {
								return nil,err
							}
							for _, info := range busstopinfo {
								flag := true
								for _, des := range destinationList {
									if info.Destination == des {
										flag = false
									}
								}
								if flag {
									busstopurl = append(busstopurl, info.Url)
								}
							}
						}
					} else if info.Busname == "15・50号系統" {
						for _, splitbus := range config.Busname15and50 {
							busstopinfo,err := queries.GetBusinfoFromBusname(ctx,bustimersqlc.GetBusinfoFromBusnameParams{Busname: splitbus,Busstop: busstop})
							if err != nil {
								//エラーハンドリング
								fmt.Printf("db select Error!!!! err:%v\n", err)
							}
							for _, info := range busstopinfo {
								flag := true
								for _, des := range destinationList {
									if info.Destination == des {
										flag = false
									}
								}
								if flag {
									busstopurl = append(busstopurl, info.Url)
								}
							}
						}

					} else {
						busstopinfo,err := queries.GetBusinfoFromBusname(ctx,bustimersqlc.GetBusinfoFromBusnameParams{Busname: info.Busname,Busstop: busstop})
						if err  != nil {
							return nil,err
						}

						for _, businfo := range busstopinfo {
							flag := true
							for _, des := range destinationList {
								if des == businfo.Destination {
									flag = false
								}
							}
							if flag {
								busstopurl = append(busstopurl, info.Url)
							}
						}
					}

				}
			}
			if err := validateUrl(busstopurl); err != nil {
				return nil,err
			}
			return busstopurl,nil
		}
	}

	if destination == "立命館大学" {
		busstopinfo,err := queries.GetBusinfoFromDestination(ctx,bustimersqlc.GetBusinfoFromDestinationParams{Busstop: busstop,Destination: destination+"行き"})
		if err != nil {
			//エラーハンドリング
			fmt.Printf("db select Error!!!! err:%v\n", err)
		}
		for _, url := range busstopinfo {
			busstopurl = append(busstopurl, url.Url)
		}
	} else {
		destinationList := config.ChangedestinationList
		busstopList,err := queries.GetBusinfoFromDestination(ctx,bustimersqlc.GetBusinfoFromDestinationParams{Busstop: destination,Destination:"立命館大学行き" })
		if err != nil {
			return nil,err
		}

		for _, bus := range busstopList {
			//方向によって行き先が変わるバスの処理
			busname := bus.Busname
			if bus.Busname == "快速立命館号系統" {
				busname = "快速205号系統"
			} else if bus.Busname == "快速205号系統" {
				busname = "快速立命館号系統"
			}

			if busname == "52・55号系統" {
				for _, splitbus := range config.Busname52and55 {
					busstopinfo,err := queries.GetBusinfoFromBusname(ctx,bustimersqlc.GetBusinfoFromBusnameParams{Busname: splitbus,Busstop: busstop})
					if err != nil {
						return nil,err
					}
					for _, info := range busstopinfo {
						flag := true
						for _, des := range destinationList {
							if info.Destination == des {
								flag = false
							}
						}
						if flag {
							busstopurl = append(busstopurl, info.Url)
						}
					}
				}
			} else if busname == "15・50号系統" {
				for _, splitbus := range config.Busname15and50 {
					busstopinfo,err := queries.GetBusinfoFromBusname(ctx,bustimersqlc.GetBusinfoFromBusnameParams{Busname: splitbus,Busstop: busstop})
					if err != nil {
						return nil,err
					}
					for _, info := range busstopinfo {
						flag := true
						for _, des := range destinationList {
							if info.Destination == des {
								flag = false
							}
						}
						if flag {
							busstopurl = append(busstopurl, info.Url)
						}
					}
				}

			} else {
				busstopinfo,err :=  queries.GetBusinfoFromBusname(ctx,bustimersqlc.GetBusinfoFromBusnameParams{Busname: busname,Busstop: busstop})
				if err != nil {
					return nil,err
				}
				for _, info := range busstopinfo {
					flag := true
					for _, des := range destinationList {
						if info.Destination == des {
							flag = false
						}
					}
					if flag {
						busstopurl = append(busstopurl, info.Url)
					}
				}
			}
		}
	}
	if err := validateUrl(busstopurl); err != nil {
		return nil,err
	}
	return busstopurl,nil
}

func (repository *BusstopToTimetableRepositoryImpl) EncodeDestination(destination string) (wrapdestination string) {
	//M1と12番のコンフリクト問題解消
	var destinationList [6]string = [6]string{"原谷行き", "金閣寺・立命館大学行き", "金閣寺･竜安寺・山越行き", "山越中町行き", "竜安寺・山越行き", "宇多野･山越行き"}
	// for _, bus := range config.M1or12or59busstopList {
	// 	if busstop == bus {
			for _, des := range destinationList {
				if destination == des {
					return "立命館大学行き"
				}
		// 	}
		// }
	}
	return destination
}

func (repository *BusstopToTimetableRepositoryImpl) FindBusstopList(busname string) ([]bustimersqlc.GetBusstopAndDestinationRow, error) {
	var err error
	ctx := context.Background()
	db := infra.GetDB()
	queries := bustimersqlc.New(db)
	// busstoplist := []model.Busstop{}

	busstoplist,err:= queries.GetBusstopAndDestination(ctx,busname)
	if err != nil {
		//エラーハンドリング
		fmt.Printf("db select Error!!!! err:%v\n", err)
	}
	return busstoplist, err

}

func (repository *BusstopToTimetableRepositoryImpl) FindURL(busstop string, destination string) ([]string, error) {
	//立命館大学に行く時
	ctx := context.Background()
	db := infra.GetDB()
	queries := bustimersqlc.New(db)
	var busstopurl []string

	//M1と12番のコンフリクト問題解消
	for _, bus := range config.M1or12or59busstopList {
		if busstop == bus && destination == "立命館大学行き" {
			var destinationList [3]string = [3]string{"原谷行き", "金閣寺・立命館大学行き", "立命館大学行き"}
			for _, des := range destinationList {
				busstopinfo,err := queries.GetBusinfoFromDestination(ctx,bustimersqlc.GetBusinfoFromDestinationParams{Busstop: busstop,Destination: des})
				if err != nil {
					fmt.Printf("db select Error!!!! err:%v\n", err)
				}
				for _, bus := range busstopinfo {
					busstopurl = append(busstopurl, bus.Url)
				}
			}
			return busstopurl,nil
		}
	}

	busstopinfo,err := queries.GetBusinfoFromDestination(ctx,bustimersqlc.GetBusinfoFromDestinationParams{Busstop: busstop,Destination: destination})
	if err != nil {
		fmt.Printf("db select Error!!!! err:%v\n", err)
	}
	for _, bus := range busstopinfo {
		busstopurl = append(busstopurl, bus.Url)
	}
	return busstopurl, err
}

func (repository *BusstopToTimetableRepositoryImpl) FindTimetable(url []string) (model.TimeTable,error) {
	timetable := model.CreateNewTimeTable()

	var err error

	if(len(url) == 0){
		return timetable,errors.New("urls not found")
	}
	for _, u := range url {
		scrapdata, via, busstop := scrapHTML(u)
		getTimeTable(timetable, scrapdata, via, busstop)
	}
	timetable.SortOneBusTime()

	return timetable,err
}

func scrapHTML(url string) (scrapData []string, via string, busstop string) {
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

		for i := 0; i < len(result); i++ {
			if strings.Contains(result[i], "号") {
				Via = result[i]
			}

			if result[i] == "【快速】"{
				Via += result[i]
			}
		}
	})

	c.OnHTML(".tt-time", func(e *colly.HTMLElement) {
		if Via == "M1号系統" {
			for i, busstop := range config.M1BusstopList {
				if busstop == Busstop {
					if strings.Contains(e.Text, "(") {
						var firstresult []string
						result := ""

						time := strings.Split(e.Text, " ")
						for _, j := range time {
							if strings.Contains(j, "(") {
								firstresult = append(firstresult, j)
							}
						}

						for k, t := range firstresult {
							if k == 0 {
								result += t
							} else {
								result += " "
								result += t
							}
						}

						result = strings.Replace(result, "(", "", -1)
						result = strings.Replace(result, ")", "", -1)
						re := regexp.MustCompile(`.*台`)
						result = re.ReplaceAllString(result, "")
						// timelist := strings.Split(result, "\n")
						scrapData = append(scrapData, result)
						break
					} else {
						result := ""
						scrapData = append(scrapData, result)
						break
					}
				}
				if i == len(config.M1BusstopList)-1 {
					result := strings.Replace(e.Text, "(", "", -1)
					result = strings.Replace(result, ")", "", -1)
					re := regexp.MustCompile(`.*台`)
					result = re.ReplaceAllString(result, "")
					// timelist := strings.Split(result, "\n")
					scrapData = append(scrapData, result)
				}
			}
		} else {
			result := strings.Replace(e.Text, "(", "", -1)
			result = strings.Replace(result, ")", "", -1)
			re := regexp.MustCompile(`.*台`)
			result = re.ReplaceAllString(result, "")
			// timelist := strings.Split(result, "\n")
			scrapData = append(scrapData, result)
		}
	})

	c.OnRequest(func(r *colly.Request) {
	})
	c.Visit(url)
	return scrapData, Via, Busstop
}

func getTimeTable(timetable model.TimeTable, scrapedata []string, via string, busstop string) model.TimeTable {

	timelist := []string{}

	for i := 0; i < len(scrapedata); i++ {
		timelist = append(timelist, scrapedata[i])
	}

	Via := via
	for i := 0; i < len(timelist)/3; i++ {

		weekdaylist := strings.Split(timelist[i*3], " ")

		for a := 0; a < len(weekdaylist); a++ {
			if weekdaylist[a] != "" {
				timetable.Weekdays[i+5] = append(timetable.Weekdays[i+5], model.OneBusTime{
					BusName: Via,
					Min:     strconv.FormatInt(toInt64(weekdaylist[a]), 10),
					BusStop: "1番乗り場",
				})
			}
		}

		holidaylist := strings.Split(timelist[i*3+2], " ")
		for b := 0; b < len(holidaylist); b++ {
			if holidaylist[b] != "" {
				timetable.Holidays[i+5] = append(timetable.Holidays[i+5], model.OneBusTime{
					BusName: Via,
					Min:     strconv.FormatInt(toInt64(holidaylist[b]), 10),
					BusStop: "1番乗り場",
				})
			}

		}
	}
	return timetable
}

func toInt64(strVal string) int64 {
	rex := regexp.MustCompile("[0-9]+")
	strVal = rex.FindString(strVal)
	intVal, err := strconv.ParseInt(strVal, 10, 64)
	if err != nil {
		fmt.Println(err)
	}
	return intVal
}

func validateUrl(url []string) error {
	length := len(url)
	if length == 0 {
		return fmt.Errorf("length must be greater than 0, length = %d", length)
	}
	return nil
}