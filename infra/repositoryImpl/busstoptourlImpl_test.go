package repositoryimpl

import "testing"

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
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := toInt64(tt.args.strVal); got != tt.want {
				t.Errorf("toInt64() = %v, want %v", got, tt.want)
			}
		})
	}
}
