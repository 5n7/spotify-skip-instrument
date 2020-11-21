# spotify-skip-instrument

`spotify-skip-instrument` skips instrument songs on Spotify.

## Notice

Currently this application is only known to work on Linux.

## Installation

Download the binary from [GitHub Releases](https://github.com/skmatz/spotify-skip-instrument/releases).

Or, if you have Go, you can install `spotify-skip-instrument` with the following command.

```console
go get github.com/skmatz/spotify-skip-instrument/...
```

## Usage

```sh
# basic usage
spotify-skip-instrument

# run in the background
spotify-skip-instrument &>/dev/null &
```

## Configuration

You can set the keywords to skip.  
Put the following JSON file in `~/.config/spotify-skip-instrument/config.json`.  
If you do not put this file, the default is to set `[instrument]` as the keywords to skip.

```json
{
  "skip_keywords": ["instrument"]
}
```
