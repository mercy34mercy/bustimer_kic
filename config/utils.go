package config

import (
	"fmt"
	"strings"
)

func GetRequiredeTime(fr string, to string, busname string) int {
	fmt.Println(fr, to, busname)
	to = strings.Replace(to, "行き", "", -1)
	switch busname {
	case "快速立命館号系統", "快速205号系統":
		switch to {
		case "立命館大学":
			return int(float64(RequiredtimeRitsumei/12.0) * float64(12-GetBusstop(fr, busname)))
		default:
			return int(float64(RequiredtimeRitsumei/12.0) * float64(GetBusstop(to, busname)))
		}
	case "50号系統":
		switch to {
		case "立命館大学":
			return int(float64(Requiredtime50/27.0) * float64(27-GetBusstop(fr, busname)))
		default:
			return int(float64(Requiredtime50/27.0) * float64(GetBusstop(to, busname)))
		}
	case "12号系統":
		switch to{
		case "立命館大学":
			return int(float64(Requiredtime12/30.0) * float64(30-GetBusstop(fr, busname)))
		default:
			return int(float64(Requiredtime12/30.0) * float64(30-GetBusstop(to, busname)))
		}
	case "15号系統":
	case "51号系統":
	case "59号系統":
	}
	return 0
}

func GetBusstop(notritsumei string, busname string) int {
	switch busname {
	case "快速立命館号系統":
		switch notritsumei {
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
		switch notritsumei {
		case "立命館大学":
			return 0
		case "桜木町":
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
	case "12号系統号":
		switch notritsumei {
		case "三条京阪前":
			return 30
		case "四条京阪前":
			return 29
		case "四条河原町":
			return 28
		case "四条高倉":
			return 27
		case "四条烏丸《地下鉄四条駅》":
			return 26
		case "四条烏丸":
			return 25
		case "四条西洞院":
			return 24
		case "四条堀川":
			return 23
		case "堀川蛸薬師":
			return 22
		case "堀川三条":
			return 21
		case "堀川御池":
			return 20
		case "二条城前":
			return 19
		case "堀川丸太町":
			return 18
		case "堀川下立売":
			return 17
		case "堀川下長者町":
			return 16
		case "堀川中立売":
			return 15
		case "一条戻橋・晴明神社前":
			return 14
		case "堀川今出川":
			return 13
		case "堀川上立売":
			return 12
		case "堀川寺ノ内":
			return 11
		case "天神公園前":
			return 10
		case "堀川鞍馬口":
			return 9
		case "北大路堀川":
			return 8
		case "大徳寺前":
			return 7
		case "建勲神社前":
			return 6
		case "船岡山":
			return 5
		case "千本北大路":
			return 4
		case "金閣寺道":
			return 3
		case "わら天神前":
			return 2
		case "桜木町":
			return 1
		case "立命館大学前":
			return 0
		}
	}
	return 0
}
