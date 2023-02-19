package repositoryimpl

import (
	"practice-colly/config"
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
				t.Errorf("BusstopToTimetableRepositoryImpl.EncodeDestination() = %v, want %v", gotWrapdestination,"立命館大学行き")
			}
		})
	}
}
