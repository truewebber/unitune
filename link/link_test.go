package link

import (
	"testing"
)

func TestIsYandexMusic(t *testing.T) {
	type args struct {
		link string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsYandexMusic(tt.args.link); got != tt.want {
				t.Errorf("IsYandexMusic() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsSpotify(t *testing.T) {
	type args struct {
		link string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test1",
			args: args{link: "https://open.spotify.com/track/2OSsMwCiZPC44tJ4OPVSZF"},
			want: true,
		},
		{
			name: "test2",
			args: args{link: "https://open.spotify.com/track/2OSsMwCiZPC44tJ4OPVSZF?si=ftK-TWUNQtOCuDGC3p7t5Q"},
			want: true,
		},
		{
			name: "test3",
			args: args{link: "http://open.spotify.com/track/2OSsMwCiZPC44tJ4OPVSZF?si=ftK-TWUNQtOCuDGC3p7t5Q"},
			want: true,
		},
		{
			name: "test4",
			args: args{link: "https://open.spotify1.com/track/2OSsMwCiZPC44tJ4OPVSZF?si=ftK-TWUNQtOCuDGC3p7t5Q"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsSpotify(tt.args.link); got != tt.want {
				t.Errorf("IsSpotify() = %v, want %v", got, tt.want)
			}
		})
	}
}
