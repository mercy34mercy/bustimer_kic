package config

import (
	"fmt"
	"strings"
)

func GetRequiredeTime(fr string,to string,busname string) int {
	fmt.Println(fr,to,busname)
	to = strings.Replace(to, "行き", "", -1)
	switch fr {
	case "立命館大学前":
		switch busname{
		case "快速立命館号系統":
			return int(float64(36.0/12.0) * float64(GetBusstop(fr,to,busname)))
		case "50号系統":
			return int(float64(42.0/27.0) * float64(GetBusstop(fr,to,busname)))
		}
	case "京都駅前":
		switch busname{
		case "快速立命館号系統":
			return int(float64(36.0/12.0) * float64(12-GetBusstop(fr,to,busname)))
		case "50号系統":
			return int(float64(42.0/27.0) * float64(27-GetBusstop(fr,to,busname)))
		}
	}
	return 10
}


func GetBusstop (fr string,to string,busname string) int {
	switch fr {
	case "立命館大学前","京都駅前":
		switch busname{
		case "快速立命館号系統":
			switch to{
			case "立命館大学":
				return 0
			case "衣笠校前":
				return 1
			case "北野白梅町":
				return 2
			case "西ノ京円町":
				return 3
			case "西大路御池":
				return 4
			case "西大路三条":
				return 5
			case "西大路四条":
				return 6
			case "西大路五条":
				return 7
			case "西大路七条":
				return 8
			case "七条千本":
				return 9
			case "七条大宮・京都水族館前":
				return 10
			case "烏丸七条":
				return 11
			case "京都駅":
				return 12
			}
		case "50号系統":
			switch to{
			case "立命館大学":
				return 0
			case  "桜木町":
				return 1
			case "わら天神前":
				return 2
			case "衣笠校前":
				return 3
			case "北野白梅町":
				return 4
			case "北野天満宮":
				return 5
			case "上七軒":
				return 6
			case "千本今出川":
				return 7
			case "千本中立売":
				return 8
			case "知恵光院中立売":
				return 9
			case "大宮中立売":
				return 10
			case "堀川中立売":
				return 11
			case "堀川下長者町":
				return 12
			case "堀川下立売":
				return 13
			case "堀川丸太町":
				return 14
			case "二条城前":
				return 15
			case "堀川御池":
				return 16
			case "堀川三条":
				return 17
			case "堀川蛸薬師":
				return 18
			case "四条堀川":
				return 19
			case "四条西洞院":
				return 20
			case "西洞院佛光寺":
				return 21
			case "西洞院松原":
				return 22
			case "五条西洞院":
				return 23
			case "西洞院六条":
				return 24
			case "西洞院正面":
				return 25
			case "七条西洞院":
				return 26
			case "京都駅":
				return 27
			}
		}
	}
	return 0
}