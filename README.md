# spotifyctrl
Spotify from the command-line.

An educational reimplementation of functionality from [spotify-local-control](https://github.com/f-mer/spotify-local-control) in [go](https://golang.org/).

## Installation
Via `go get`:

```sh
$ go get github.com/f-mer/spotifyctrl/cmd/spotify
```

Then control spotify via `spotifyctrl <command>`.

## Usage

```sh
$ spotifyctrl --help

  Usage:

    spotifyctrl [<flags>] <command> [<args> ...]

  Flags:

    -h, --help             Output usage information.
        --start-port=4370
        --end-port=4380
        --version          Show application version.

  Commands:

    help                 Show help for a command.
    play                 Play a track by uri.
    pause                Pause playback.
    resume               Resume playback.
    status               Display status information.
```

## License
[MIT](https://tldrlegal.com/license/mit-license)
