package main

import (
	"log"
	"os"
	"text/template"

	"github.com/tj/kingpin"

	"github.com/f-mer/spotifyctrl/internal/client"
)

var (
	version = "master"

	app          = kingpin.New("spotifyctrl", "Spotify from the command-line.")
	startPort    = app.Flag("start-port", "").Default("4370").Int()
	endPort      = app.Flag("end-port", "").Default("4380").Int()
	play         = app.Command("play", "Play a track by uri.")
	playUri      = play.Arg("uri", "The resource identifier that you can enter in the Desktop client's search box to locate an artist, album, or track. To find a Spotify URI right-click on the track's name.").Required().String()
	playContext  = play.Arg("context", "The base-62 identifier that you can find at the end of the Spotify URI (see above) for an artist, track, album, playlist, etc.").String()
	pause        = app.Command("pause", "Pause playback.")
	resume       = app.Command("resume", "Resume playback.")
	status       = app.Command("status", "Display status information.")
	statusFormat = status.Flag("format", "Go template").Default("{{.Track.ArtistResource.Name}} - {{.Track.TrackResource.Name}}").String()
)

func main() {
	app.Version(version)
	kingpin.MustParse(app.Parse(os.Args[1:]))

	spotify, err := client.New(*startPort, *endPort)
	if err != nil {
		log.Fatalf("error initializing client: %s", err)
	}

	switch kingpin.MustParse(app.Parse(os.Args[1:])) {
	case play.FullCommand():
		_, err := spotify.Play(&client.PlayInput{
			Uri:     *playUri,
			Context: *playContext,
		})
		if err != nil {
			log.Fatalf("error playing track: %s", err)
		}
	case pause.FullCommand():
		_, err := spotify.Pause(&client.PauseInput{
			Pause: true,
		})
		if err != nil {
			log.Fatalf("error pausing track: %s", err)
		}
	case resume.FullCommand():
		_, err := spotify.Pause(&client.PauseInput{
			Pause: false,
		})
		if err != nil {
			log.Fatalf("error resuming track: %s", err)
		}
	case status.FullCommand():
		tmpl, err := template.New("format").Parse(*statusFormat)
		if err != nil {
			log.Fatalf("error parsing template: %s", err)
		}

		status, _, err := spotify.Status()
		if err != nil {
			log.Fatalf("error retrieving status: %s", err)
		}

		tmpl.Execute(os.Stdout, status)
	}
}
