package repositoryimpl

import (
	"bustimerkic/config"
	"bustimerkic/infra"
	"testing"
)

func Test_toInt64(t *testing.T) {
	type args struct {
		strVal string
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "toInt64test",
			args: args{"5"},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := toInt64(tt.args.strVal); got != tt.want {
				t.Errorf("toInt64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBusstopToTimetableRepositoryImpl_EncodeDestination(t *testing.T) {
	tests := config.ChangedestinationList
	for _, tt := range tests {
		t.Run(tt, func(t *testing.T) {
			repository := &BusstopToTimetableRepositoryImpl{}
			if gotWrapdestination := repository.EncodeDestination(tt); gotWrapdestination != "立命館大学行き" {
				t.Errorf("BusstopToTimetableRepositoryImpl.EncodeDestination() = %v, want %v", gotWrapdestination, "立命館大学行き")
			}
		})
	}
}

func TestFindUrlFromBusstopToRitsumei(t *testing.T) {
	t.Parallel()
	infra.Init("../../gorm.db")
	for _, busstoplist := range config.BusstoptoRitsList {
		for _, busstop := range busstoplist {
			repository := &BusstopToTimetableRepositoryImpl{}
			_, err := repository.FindURLFromBusstop(busstop, "立命館大学")
			if busstop != "立命館大学前" {
				if err != nil {
					t.Errorf("BusstopToTimetableRepositoryImpl.FindURLFromBusstop() %s → %s URL Not Found", busstop, "立命館大学")
				}
			}else{
				if err == nil{
					t.Errorf("BusstopToTimetableRepositoryImpl.FindURLFromBusstop() %s → %s This URL must not exist", busstop, "立命館大学")
				}
			}
		}
	}
}

func TestFindUrlFromBusstopFromRitsumei(t *testing.T) {
	t.Parallel()
	infra.Init("../../gorm.db")
	for _, busstoplist := range config.BusstopfromRitsList {
		for _, busstop := range busstoplist {
			repository := &BusstopToTimetableRepositoryImpl{}
			_, err := repository.FindURLFromBusstop("立命館大学前", busstop)
			if busstop != "立命館大学前" {
				if err != nil {
					t.Errorf("BusstopToTimetableRepositoryImpl.FindURLFromBusstop() %s → %s URL Not Found", "立命館大学前", busstop)
				}
			}else{
				if err == nil{
					t.Errorf("BusstopToTimetableRepositoryImpl.FindURLFromBusstop() %s → %s This URL must not exist", busstop, "立命館大学")
				}
			}
		}
	}
}
