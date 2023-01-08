package localcache

import (
	"fmt"
	"practice-colly/config"
	"practice-colly/controller"
	"time"

	"github.com/patrickmn/go-cache"
)

var (
	c *cache.Cache
)

func Init() *cache.Cache{
	c = cache.New(12*time.Hour, 24*time.Hour)
	return c
}

func GetGoChache() *cache.Cache{
	return c
}



func CreateTimetableCache() {
	for _,busname := range config.BusNameList{
	fmt.Printf(busname)
	//バスの名前から、そのバスの停留所と行き先のセットを取得
	busstoplistCtrl := controller.FindBusstopListController{}
	busstoplist, err := busstoplistCtrl.FindBusstopList(busname)
	if err != nil {
	}

	for _,businfo:= range busstoplist {
		fmt.Printf(businfo.Busstop,businfo.Destination)
		//停留所と行き先のセットからURLを取得
		busstoptourlCtrl := controller.BusstoToUrlController{}
		url, err := busstoptourlCtrl.FindURL(businfo.Busstop, businfo.Destination)
		if err != nil {
		}

		//URLから時刻表を取得
		timetablecontroller := controller.TimetableController{}
		timetable := timetablecontroller.FindTimetable(url)

		c.Set(businfo.Busstop+businfo.Destination, timetable, cache.DefaultExpiration)
	}

	}

	fmt.Println("キャッシュ化完了")


}
