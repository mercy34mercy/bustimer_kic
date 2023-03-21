package config

import (
	"strings"
)

func GetRequiredeTime(fr string, to string, busname string) int {
	to = strings.Replace(to, "行き", "", -1)
	switch busname {
	case "立命館ダイレクト号系統":
		return GetBusstop(fr, busname)
	case "快速立命館号系統", "快速205号系統":
		switch to {
		case "立命館大学":
			return int(float64(RequiredtimeRitsumei/12.0) * float64(GetBusstop(fr, busname)))
		default:
			return int(float64(RequiredtimeRitsumei/12.0) * float64(GetBusstop(to, busname)))
		}
	case "50号系統", "15・50号系統":
		switch to {
		case "立命館大学":
			return int(float64(Requiredtime50/27.0) * float64(GetBusstop(fr, busname)))
		default:
			return int(float64(Requiredtime50/27.0) * float64(GetBusstop(to, busname)))
		}
	case "12号系統":
		switch to {
		case "立命館大学":
			return int(float64(Requiredtime12/30.0) * float64(GetBusstop(fr, busname)))
		default:
			return int(float64(Requiredtime12/30.0) * float64(GetBusstop(to, busname)))
		}
	case "15号系統":
		switch to {
		case "立命館大学":
			return int(float64(Requiredtime15/22.0) * float64(GetBusstop(fr, busname)))
		default:
			return int(float64(Requiredtime15/22.0) * float64(GetBusstop(to, busname)))
		}
	case "51号系統":
		switch to {
		case "立命館大学":
			return int(float64(Requiredtime51/23.0) * float64(GetBusstop(fr, busname)))
		default:
			return int(float64(Requiredtime51/23.0) * float64(GetBusstop(to, busname)))
		}
	case "59号系統":
		switch to {
		case "金閣寺･竜安寺・山越行き","竜安寺・山越行き":
			return int(float64(RequiredTime59/23.0) * float64(GetBusstop(fr, busname)))
		default:
			return int(float64(RequiredTime59/23.0) * float64(GetBusstop(to, busname)))
		}
	case "臨号系統":
		switch to{
		case "立命館大学":
			return int(float64(RequiredTimeRin/8.0) * float64(GetBusstop(fr, busname)))
		default:
			return int(float64(RequiredTimeRin/8.0) * float64(GetBusstop(to, busname)))
		}
	case "臨号系統【快速】":
		switch to{
		case "立命館大学":
			return int(float64(RequiredTimeRinRapid/4.0) * float64(GetBusstop(fr, busname)))
		default:
			return int(float64(RequiredTimeRinRapid/4.0) * float64(GetBusstop(to, busname)))
		}
	case "52号系統", "55号系統":
		switch to {
		case "原谷行き", "立命館大学":
			return int(float64(RequiredTime52/16.0) * float64(GetBusstop(fr, busname)))
		default:
			return int(float64(RequiredTime52/16.0) * float64(GetBusstop(to, busname)))
		}
	case "M1号系統":
		switch to {
		case "原谷行き", "立命館大学":
			return int(float64(RequiredTimeM1/10.0) * float64(GetBusstop(fr, busname)))
		default:
			return int(float64(RequiredTimeM1/10.0) * float64(GetBusstop(to, busname)))
		}
	case "205号系統":
	case "204号系統":
	}
	return 0
}

func GetBusstop(notritsumei string, busname string) int {
	switch busname {
	case "立命館ダイレクト号系統":
		switch notritsumei {
			case "西大路四条《阪急･嵐電西院駅》":
				return 15
			case "西大路御池":
				return 11
			case "西ノ京円町《ＪＲ円町駅》":
				return 8
		}
	case "臨号系統":
		switch notritsumei {
		case "立命館大学前":
			return 0
		case "等持院東町":
			return 1
		case "府立体育館前《島津アリーナ京都前》":
			return 2
		case "大将軍":
			return 3
		case "北野中学前":
			return 4
		case "西ノ京円町《ＪＲ円町駅》":
			return 5
		case "太子道":
			return 6
		case "西大路御池":
			return 7
		case "西大路四条《阪急･嵐電西院駅》":
			return 8
		}
	case "臨号系統【快速】":
		switch notritsumei {
		case "立命館大学前":
			return 0
		case "北野白梅町":
			return 1
		case "西ノ京円町《ＪＲ円町駅》":
			return 2
		case "二条駅前":
			return 3
		case "三条京阪前":
			return 4
		}
	case "快速立命館号系統", "快速205号系統":
		switch notritsumei {
		case "立命館大学":
			return 0
		case "衣笠校前":
			return 1
		case "北野白梅町":
			return 2
		case "西ノ京円町《ＪＲ円町駅》":
			return 3
		case "西大路御池":
			return 4
		case "西大路三条":
			return 5
		case "西大路四条《阪急･嵐電西院駅》":
			return 6
		case "西大路五条":
			return 7
		case "西大路七条":
			return 8
		case "西大路駅前":
			return 9
		case "七条千本":
			return 9
		case "七条大宮・京都水族館前":
			return 10
		case "烏丸七条":
			return 11
		case "京都駅前":
			return 12
		}
	case "50号系統", "15・50号系統":
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
		case "北野天満宮前":
			return 5
		case "上七軒":
			return 6
		case "千本今出川":
			return 7
		case "千本中立売":
			return 8
		case "智恵光院中立売":
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
		case "西洞院仏光寺":
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
		case "京都駅前":
			return 27
		}
	case "12号系統":
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
	case "15号系統":
		switch notritsumei {
		case "三条京阪前":
			return 22
		case "四条京阪前":
			return 21
		case "四条河原町":
			return 20
		case "河原町三条":
			return 19
		case "京都市役所前":
			return 18
		case "堺町御池":
			return 17
		case "烏丸御池":
			return 16
		case "新町御池":
			return 15
		case "堀川御池":
			return 14
		case "神泉苑前":
			return 13
		case "二条駅前":
			return 12
		case "千本旧二条":
			return 11
		case "千本丸太町":
			return 10
		case "丸太町七本松":
			return 9
		case "丸太町御前通":
			return 8
		case "西ノ京円町《ＪＲ円町駅》":
			return 7
		case "北野中学前":
			return 6
		case "大将軍":
			return 5
		case "北野白梅町":
			return 4
		case "わら天神前":
			return 3
		case "衣笠校前":
			return 2
		case "桜木町":
			return 1
		case "立命館大学前":
			return 0
		}
	case "51号系統":
		switch notritsumei {
		case "三条京阪前":
			return 23
		case "四条京阪前":
			return 22
		case "四条河原町":
			return 21
		case "河原町三条":
			return 20
		case "京都市役所前":
			return 19
		case "堺町御池":
			return 18
		case "烏丸二条":
			return 17
		case "烏丸御池":
			return 16
		case "烏丸丸太町《地下鉄丸太町駅》":
			return 15
		case "烏丸下立売":
			return 14
		case "烏丸下長者町":
			return 13
		case "烏丸一条":
			return 12
		case "烏丸今出川《地下鉄今出川駅》":
			return 11
		case "上京区総合庁舎前":
			return 10
		case "堀川今出川":
			return 9
		case "今出川大宮":
			return 8
		case "今出川浄福寺":
			return 7
		case "千本今出川":
			return 6
		case "上七軒":
			return 5
		case "北野天満宮前":
			return 4
		case "北野白梅町":
			return 3
		case "衣笠校前":
			return 2
		case "小松原児童公園前":
			return 1
		case "立命館大学前":
			return 0
		}
	case "59号系統":
		switch notritsumei {
		case "四条河原町":
			return 23
		case "四条京阪前":
			return 22
		case "三条京阪前":
			return 21
		case "河原町三条":
			return 20
		case "京都市役所前":
			return 19
		case "河原町丸太町":
			return 18
		case "荒神口":
			return 17
		case "府立医大病院前":
			return 16
		case "河原町今出川":
			return 15
		case "同志社前":
			return 14
		case "烏丸今出川《地下鉄今出川駅》":
			return 13
		case "上京区総合庁舎前":
			return 12
		case "堀川今出川":
			return 11
		case "今出川大宮":
			return 10
		case "今出川浄福寺":
			return 9
		case "千本今出川":
			return 8
		case "千本上立売":
			return 7
		case "乾隆校前":
			return 6
		case "千本鞍馬口":
			return 5
		case "ライトハウス前":
			return 4
		case "千本北大路":
			return 3
		case "金閣寺道":
			return 2
		case "桜木町":
			return 1
		case "立命館大学前":
			return 0
		case "竜安寺前":
			return 1
		case "塔ノ下町":
			return 2
		case "御室仁和寺":
			return 3
		case "福王子":
			return 4
		case "鳴滝本町":
			return 5
		case "宇多野病院前":
			return 6
		case "ユースホステル前":
			return 7
		case "山越":
			return 8
		case "広沢池・佛大広沢校前":
			return 9
		case "山越中町":
			return 10
		}
	case "M1号系統":
		switch notritsumei {
		case "北大路バスターミナル《地下鉄北大路駅》":
			return 10
		case "北大路新町":
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
	case "52号系統":
		switch notritsumei {
		case "四条烏丸《地下鉄四条駅》":
			return 18
		case "四条西洞院":
			return 17
		case "四条堀川":
			return 16
		case "四条大宮":
			return 15
		case "みぶ操車場前":
			return 14
		case "千本三条・朱雀立命館前":
			return 13
		case "二条駅前":
			return 12
		case "千本旧二条":
			return 11
		case "千本丸太町":
			return 10
		case "丸太町七本松":
			return 9
		case "七本松出水":
			return 8
		case "七本松仁和寺街道":
			return 7
		case "上七軒":
			return 6
		case "北野天満宮前":
			return 5
		case "北野白梅町":
			return 4
		case "衣笠校前":
			return 3
		case "わら天神前":
			return 2
		case "桜木町":
			return 1
		case "立命館大学前":
			return 0
		}

	case "55号系統":
		switch notritsumei {
		case "四条烏丸《地下鉄四条駅》":
			return 18
		case "四条西洞院":
			return 17
		case "四条堀川":
			return 16
		case "四条大宮":
			return 15
		case "みぶ操車場前":
			return 14
		case "千本三条・朱雀立命館前":
			return 13
		case "二条駅前":
			return 12
		case "千本旧二条":
			return 11
		case "千本丸太町":
			return 10
		case "千本出水":
			return 9
		case "千本中立売":
			return 8
		case "千本今出川":
			return 7
		case "上七軒":
			return 6
		case "北野天満宮前":
			return 5
		case "北野白梅町":
			return 4
		case "衣笠校前":
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
