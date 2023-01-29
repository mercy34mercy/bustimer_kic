package localcache

import (
	"fmt"
	"practice-colly/config"
	"practice-colly/controller"
	"practice-colly/domain/model"
	"time"

	"github.com/patrickmn/go-cache"
)

var (
	c *cache.Cache
)

func Init() *cache.Cache {
	c = cache.New(12*time.Hour, 24*time.Hour)
	return c
}

func GetGoChache() *cache.Cache {
	return c
}

func CreateCachefromTimetable(busstop string,destination string,timetable model.TimeTable) {
	c.Set(busstop+destination, timetable, cache.DefaultExpiration)
}

func CreateTimetableCache() {
	for {
		for _, busname := range config.BusNameList {
			fmt.Printf(busname)
			//バスの名前から、そのバスの停留所と行き先のセットを取得
			busstoplistCtrl := controller.FindBusstopListController{}
			busstoplist, err := busstoplistCtrl.FindBusstopList(busname)
			if err != nil {
			}

			for _, businfo := range busstoplist {
				//
				timetableCtrl := controller.CacheTimetableController{}
				timetable,busstop,destination := timetableCtrl.FindCacheTimetable(businfo)
				c.Set(busstop+destination, timetable, cache.DefaultExpiration)
			}
		}
		fmt.Println("キャッシュ化完了")
		time.Sleep(config.TimeTableCacheUpdateDuration)

	}

}
