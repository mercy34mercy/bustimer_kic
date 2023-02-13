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
var M1or12or59busstopList [61]string = [61]string{
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
	"桜木町", "北大路バスターミナル《地下鉄北大路駅》", "北大路新町",
	"河原町丸太町", "荒神口", "府立医大病院前", "河原町今出川", "同志社前", "烏丸今出川《地下鉄今出川駅》", "上京区総合庁舎前", "堀川今出川", "今出川大宮", "今出川浄福寺", "千本今出川", "千本上立売", "乾隆校前", "千本鞍馬口", "ライトハウス前", "千本北大路", "金閣寺道", "桜木町","竜安寺前", "塔ノ下町", "御室仁和寺", "福王子", "鳴滝本町", "宇多野病院前", "ユースホステル前", "山越", "広沢池・佛大広沢校前", "山越中町"}

const (
	TimeTableCacheUpdateDuration = 24 * 60 * 60 * time.Second
)

var DistinationList [1][]string = [1][]string{M1BusstopList[:]}

// var M1DestinationList[11] string = [11]string {"北大路バスターミナル《地下鉄北大路駅》","北大路新町","北大路堀川","大徳寺前","建勲神社前","船岡山","千本北大路","金閣寺道","わら天神前","桜木町","立命館大学前"}
// var N50DesitnationList[100] string = [100]string {"京都駅前","七条西洞院","西洞院正面","西洞院六条","五条西洞院","西洞院松原","西洞院佛光寺","四条西洞院","四条堀川","堀川蛸薬師","堀川三条","堀川御池","二条城前","堀川丸太町","堀川下立売","堀川下長者町","堀川中立売","大宮中立売","知恵光院中立売","千本中立売","千本今出川","上七軒","北野天満宮前","北野白梅町","衣笠校前","わら天神前","桜木町","立命館大学前"}

var BusNameList [11]string = [11]string{"50号系統", "快速立命館号系統", "205号系統", "52号系統", "55号系統", "15号系統", "臨号系統", "204号系統", "12号系統", "15号系統", "M1号系統"}

var BusRinNishioziList [9]string = [9]string{"立命館大学前", "等持院東町", "府立体育館前《島津アリーナ京都前》", "大将軍", "北野中学前", "西ノ京円町《ＪＲ円町駅》", "太子道", "西大路御池", "西大路四条《阪急･嵐電西院駅》"}
var BusRinRisumeiList [6]string = [6]string{"立命館大学前", "北野白梅町", "西ノ京円町《ＪＲ円町駅》", "四条大宮", "二条駅前", "三条京阪前"}
var Bus55List [19]string = [19]string{"四条烏丸《地下鉄四条駅》", "四条西洞院", "四条堀川", "四条大宮", "みぶ操車場前", "千本三条・朱雀立命館前", "二条城前", "千本旧二条", "千本丸太町", "丸太町七本松", "七本松出水", "七本松仁和寺街道", "上七軒", "北野天満宮前", "北野白梅町", "衣笠校前", "わら天神前", "桜木町", "立命館大学前"}
var Bus52List [19]string = [19]string{"四条烏丸《地下鉄四条駅》", "四条西洞院", "四条堀川", "四条大宮", "みぶ操車場前", "千本三条・朱雀立命館前", "二条城前", "千本旧二条", "千本丸太町", "丸太町七本松", "七本松出水", "七本松仁和寺街道", "上七軒", "北野天満宮前", "北野白梅町", "衣笠校前", "わら天神前", "桜木町", "立命館大学前"}
var Bus59List [35]string = [35]string{"河原町三条", "四条河原町", "四条京阪前", "三条京阪前", "河原町三条", "京都市役所前", "河原町丸太町", "荒神口", "府立医大病院前", "河原町今出川", "同志社前", "烏丸今出川《地下鉄今出川駅》", "上京区総合庁舎前", "堀川今出川", "今出川大宮", "今出川浄福寺", "千本今出川", "千本上立売", "乾隆校前", "千本鞍馬口", "ライトハウス前", "千本北大路", "金閣寺道", "桜木町", "立命館大学前", "竜安寺前", "塔ノ下町", "御室仁和寺", "福王子", "鳴滝本町", "宇多野病院前", "ユースホステル前", "山越", "広沢池・佛大広沢校前", "山越中町"}
var Bus51List [19]string = [19]string{"四条烏丸《地下鉄四条駅》", "四条西洞院", "四条堀川", "四条大宮", "みぶ操車場前", "千本三条・朱雀立命館前", "二条城前", "千本旧二条", "千本丸太町", "丸太町七本松", "七本松出水", "七本松仁和寺街道", "上七軒", "北野天満宮前", "北野白梅町", "衣笠校前", "わら天神前", "桜木町", "立命館大学前"}
var BusM1List [11]string = [11]string{"北大路バスターミナル《地下鉄北大路駅》", "北大路新町", "北大路堀川", "大徳寺前", "建勲神社前", "船岡山", "千本北大路", "金閣寺道", "わら天神前", "桜木町", "立命館大学前"}
var Bus15List [23]string = [23]string{"三条京阪前", "四条京阪前", "四条河原町", "河原町三条", "京都市役所前", "堺町御池", "烏丸御池", "新町御池", "堀川御池", "神泉苑前", "二条駅前", "千本旧二条", "千本丸太町", "丸太町七本松", "丸太町御前通", "西ノ京円町《ＪＲ円町駅》", "北野中学前", "大将軍", "北野白梅町", "わら天神前", "衣笠校前", "桜木町", "立命館大学前"}
var BusRitsumeiList [13]string = [13]string{"京都駅前", "烏丸七条", "七条大宮・京都水族館前", "七条千本", "西大路七条", "西大路五条", "西大路四条《阪急･嵐電西院駅》", "西大路三条", "西大路御池", "西ノ京円町《ＪＲ円町駅》", "北野白梅町", "衣笠校前", "西大路駅前"}
var Bus12List [30]string = [30]string{"三条京阪前", "四条京阪前", "四条河原町", "四条高倉", "四条烏丸《地下鉄四条駅》", "四条西洞院", "四条堀川", "堀川蛸薬師", "堀川三条", "堀川御池", "二条城前", "堀川丸太町", "堀川下立売", "堀川下長者町", "堀川中立売", "一条戻橋・晴明神社前", "堀川今出川", "堀川上立売", "堀川寺ノ内", "天神公園前", "堀川鞍馬口", "北大路堀川", "大徳寺前", "建勲神社前", "船岡山", "千本北大路", "金閣寺道", "わら天神前", "桜木町", "立命館大学前"}
var Bus50List [28]string = [28]string{"京都駅前", "七条西洞院", "西洞院正面", "西洞院六条", "五条西洞院", "西洞院松原", "西洞院仏光寺", "四条西洞院", "四条堀川", "堀川蛸薬師", "堀川三条", "堀川御池", "二条城前", "堀川丸太町", "堀川下立売", "堀川下長者町", "堀川中立売", "大宮中立売", "智恵光院中立売", "千本中立売", "千本今出川", "上七軒", "北野天満宮前", "北野白梅町", "衣笠校前", "わら天神前", "桜木町", "立命館大学前"}

var AllBusstopList [11][]string = [11][]string{Bus50List[:], Bus12List[:], Bus15List[:], BusRinNishioziList[:], BusRinRisumeiList[:], BusRitsumeiList[:], BusM1List[:], Bus51List[:], Bus59List[:], Bus52List[:], Bus55List[:]}
var Busname [11]string = [11]string{"50","12","15","臨西大路","臨立","快速立命館","M1","51","59","52","55"}