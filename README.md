# spotlight

[![go][go-version-src]][go-version-href]
[![tests][tests-src]][tests-href]
[![license][license-src]][license-href]

CLI to parse files for a list or to get metadata. Powered by Go.

Metadata parsed by [`mediainfo.23.10`](https://mediaarea.net/en/MediaInfo) (hard requirement).

## Install

```bash
go install github.com/ewilan-riviere/files@latest
```

## Usage

```bash
go build -o files ; ./files parse /Volumes/data/video/animation
go build -o files ; ./files --output="output/files.json" parse /Volumes/data/video/animation

go build -o files ; ./files --output="output/metadata.json" metadata /Volumes/data/music/librairies/podcasts/F.Kermesse/FK.1_Le.Gore.Philippe.Bouvard.de.la.mort.mp3
```

## License

[MIT](LICENSE) © Ewilan Rivière

[go-version-src]: https://img.shields.io/static/v1?style=flat&label=Go&message=v1.21&color=00ADD8&logo=go&logoColor=ffffff&labelColor=18181b
[go-version-href]: https://go.dev/
[tests-src]: https://img.shields.io/github/actions/workflow/status/ewilan-riviere/files/run-tests.yml?branch=main&label=tests&style=flat&colorA=18181B
[tests-href]: https://github.com/ewilan-riviere/files/actions
[license-src]: https://img.shields.io/github/license/ewilan-riviere/files.svg?style=flat&colorA=18181B&colorB=00ADD8
[license-href]: https://github.com/ewilan-riviere/files/blob/main/LICENSE
