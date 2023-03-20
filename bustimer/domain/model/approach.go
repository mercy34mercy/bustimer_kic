package model

import (
	"fmt"
	"strconv"
)

type ApproachInfo struct {
	MoreMin         string `json:"more_min"`
	RealArrivalTime string `json:"real_arrival_time"`
	Direction       string `json:"direction"`
	ScheduledTime   string `json:"scheduled_time"`
	Delay           string `json:"delay"`
	BusStop         string `json:"bus_stop"`
	BusName         string `json:"bus_name"`
	RequiredTime    int    `json:"required_time"`
}

type ApproachInfos struct {
	ApproachInfo []ApproachInfo `json:"approach_infos"`
}

type compareFastInfo struct {
	index int
	hour  int
	min   int
}

func CreateApproachInfos() ApproachInfos {
	return ApproachInfos{
		ApproachInfo: make([]ApproachInfo, 0),
	}
}

func (infos ApproachInfos) GetFastThree() ApproachInfos {
	approachInfos := CreateApproachInfos()
	// 3つ以上データがない場合はそのまま返す
	if len(infos.ApproachInfo) <= 3 {
		return infos
	}

	// １番、２番、３番と早い順番にインデックスを保持しておく
	hour := 99
	min := 99
	first := compareFastInfo{
		index: 0,
		hour:  hour,
		min:   min,
	}
	second := compareFastInfo{
		index: 1,
		hour:  hour,
		min:   min,
	}
	third := compareFastInfo{
		index: 2,
		hour:  hour,
		min:   min,
	}
	fmt.Printf("first: %v, %v, %v\n", first.index, first.hour, first.min)
	fmt.Printf("second: %v, %v, %v\n", second.index, second.hour, second.min)
	fmt.Printf("third: %v, %v, %v\n", third.index, third.hour, third.min)

	for i, v := range infos.ApproachInfo {
		hour, _ = strconv.Atoi(v.RealArrivalTime[:2])
		min, _ = strconv.Atoi(v.RealArrivalTime[3:])
		fmt.Printf("%02d:%02d\n", hour, min)
		if hour < first.hour || (hour == first.hour && min < first.min) {
			third = second
			second = first
			first.index = i
			first.hour = hour
			first.min = min
			fmt.Println("firstを入れ替えます。")
			fmt.Printf("first: %v, %v, %v\n", first.index, first.hour, first.min)
			fmt.Printf("second: %v, %v, %v\n", second.index, second.hour, second.min)
			fmt.Printf("third: %v, %v, %v\n", third.index, third.hour, third.min)
			continue
		}
		if hour < second.hour || (hour == second.hour && min < second.min) {
			third = second
			second.index = i
			second.hour = hour
			second.min = min
			fmt.Println("secondを入れ替えます。")
			fmt.Printf("first: %v, %v, %v\n", first.index, first.hour, first.min)
			fmt.Printf("second: %v, %v, %v\n", second.index, second.hour, second.min)
			fmt.Printf("third: %v, %v, %v\n", third.index, third.hour, third.min)
			continue
		}
		if hour < third.hour || (hour == third.hour && min < third.min) {
			third.index = i
			third.hour = hour
			third.min = min
			fmt.Println("thirdを入れ替えます。")
			fmt.Printf("first: %v, %v, %v\n", first.index, first.hour, first.min)
			fmt.Printf("second: %v, %v, %v\n", second.index, second.hour, second.min)
			fmt.Printf("third: %v, %v, %v\n", third.index, third.hour, third.min)
			continue
		}
	}
	approachInfos.ApproachInfo = append(approachInfos.ApproachInfo, infos.ApproachInfo[first.index])
	approachInfos.ApproachInfo = append(approachInfos.ApproachInfo, infos.ApproachInfo[second.index])
	approachInfos.ApproachInfo = append(approachInfos.ApproachInfo, infos.ApproachInfo[third.index])
	return approachInfos
}
