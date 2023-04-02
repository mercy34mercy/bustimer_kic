package repositoryimpl

import (
	"testing"

	"github.com/mercy34mercy/bustimer_kic/bustimer/config"
	"github.com/mercy34mercy/bustimer_kic/bustimer/infra"
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
			} else {
				if err == nil {
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
			} else {
				if err == nil {
					t.Errorf("BusstopToTimetableRepositoryImpl.FindURLFromBusstop() %s → %s This URL must not exist", busstop, "立命館大学")
				}
			}
		}
	}
}

func BenchmarkFindTimetable(b *testing.B) {
	infra.Init("../../gorm.db")
	for i := 0; i < b.N; i++ {
		repository := &BusstopToTimetableRepositoryImpl{}
		urls, err := repository.FindURLFromBusstop("立命館大学前", "京都駅前")
		if err != nil {
			b.Errorf("BusstopToTimetableRepositoryImpl.FindURLFromBusstop() %s → %s URL Not Found", "立命館大学前", "京都駅前")
		}

		_, err = repository.FindTimetable(urls)
		if err != nil {
			b.Errorf("BusstopToTimetableRepositoryImpl.FindTimetable() %s Timetable Not Found", urls)
		}
	}
}

func BenchmarkFindTimetableParallel(b *testing.B) {
	infra.Init("../../gorm.db")
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			repository := &BusstopToTimetableRepositoryImpl{}
			urls, err := repository.FindURLFromBusstop("立命館大学前", "京都駅前")
			if err != nil {
				b.Errorf("BusstopToTimetableRepositoryImpl.FindURLFromBusstop() %s → %s URL Not Found", "立命館大学前", "京都駅前")
			}

			_, err = repository.FindTimetableParallel(urls)
			if err != nil {
				b.Errorf("BusstopToTimetableRepositoryImpl.FindTimetable() %s Timetable Not Found", urls)
			}
		}
	})
}