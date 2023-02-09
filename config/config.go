package config

import "time"

const Requiredtime12 = 53.0
const Requiredtime50 = 42.0
const Requiredtime15 = 34.0
const Requiredtime51 = 36.0
const Requiredtime59 = 40.0 
const RequiredtimeRitsumei = 36.0
const RequiredTimeM1 = 16
const RequiredTime52 = 30
const RequiredTime55 = 30
const RequiredTime59 = 36
const RequiredTimeRin = 17.0
const RequiredTime204 = 0
const RequiredTime205 = 0

var M1BusstopList [8]string = [8]string{"北大路バスターミナル《地下鉄北大路駅》", "北大路新町", "北大路堀川", "大徳寺前", "建勲神社前", "船岡山", "千本北大路", "金閣寺道"}
var M1or12BusstopList [33]string = [33]string{
	"三条京阪前",
	"四条京阪前",
	"四条河原町",
	"四条高倉",
	"四条烏丸《地下鉄四条駅》",
	"四条烏丸",
	"四条西洞院",
	"四条堀川",
	"堀川蛸薬師",
	"堀川三条",
	"堀川御池",
	"二条城前",
	"堀川丸太町",
	"堀川下立売",
	"堀川下長者町",
	"堀川中立売",
	"一条戻橋・晴明神社前",
	"堀川今出川",
	"堀川上立売",
	"堀川寺ノ内",
	"天神公園前",
	"堀川鞍馬口",
	"北大路堀川",
	"大徳寺前",
	"建勲神社前",
	"船岡山",
	"千本北大路",
	"金閣寺道",
	"わら天神前",
	"桜木町",
	"立命館大学前","北大路バスターミナル《地下鉄北大路駅》", "北大路新町"}

const (
	TimeTableCacheUpdateDuration = 24 * 60 * 60 * time.Second
)

var DistinationList [1][]string = [1][]string{M1BusstopList[:]}

// var M1DestinationList[11] string = [11]string {"北大路バスターミナル《地下鉄北大路駅》","北大路新町","北大路堀川","大徳寺前","建勲神社前","船岡山","千本北大路","金閣寺道","わら天神前","桜木町","立命館大学前"}
// var N50DesitnationList[100] string = [100]string {"京都駅前","七条西洞院","西洞院正面","西洞院六条","五条西洞院","西洞院松原","西洞院佛光寺","四条西洞院","四条堀川","堀川蛸薬師","堀川三条","堀川御池","二条城前","堀川丸太町","堀川下立売","堀川下長者町","堀川中立売","大宮中立売","知恵光院中立売","千本中立売","千本今出川","上七軒","北野天満宮前","北野白梅町","衣笠校前","わら天神前","桜木町","立命館大学前"}

var BusNameList [11]string = [11]string{"50号系統", "快速立命館号系統", "205号系統", "52号系統", "55号系統", "15号系統", "臨号系統", "204号系統", "12号系統", "15号系統", "M1号系統"}
