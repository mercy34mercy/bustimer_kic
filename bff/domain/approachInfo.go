package domain

type ApproachInfo struct {
	MoreMin        string `json:"more_min"`
	RealARivalTime string `json:"real_arrival_time"`
	Direction      string `json:"direction"`
	ScheduledTime  string `json:"scheduled_time"`
	Delay          string `json:"delay"`
	BusStop        string `json:"bus_stop"`
	BusName        string `json:"bus_name"`
	RequiredTime   int    `json:"required_time"`
}

type ApproachInfos struct {
	ApproachInfo []*ApproachInfo `json:"approach_infos"`
}
