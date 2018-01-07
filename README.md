# spotify-cli
Spotify from the command-line.

An educational reimplementation of functionality from [spotify-local-control](https://github.com/f-mer/spotify-local-control) in [go](https://golang.org/).

## Installation
Via `go get`:

```sh
$ go get github.com/f-mer/spotify-cli/cmd/spotify
```

Then control spotify via `spotify <command>`.

## Usage

```sh
$ spotify --help
  Usage:

    spotify [<flags>] <command> [<args> ...]

  Flags:

    -h, --help     Output usage information.
        --version  Show application version.

  Commands:

    help                 Show help for a command.
    play                 Play a track by uri
    pause                Pause playback
    resume               Resume playback
    status               Display status information
```

## License
[MIT](https://tldrlegal.com/license/mit-license)
