# spotlight

[![go][go-version-src]][go-version-href]
[![tests][tests-src]][tests-href]
[![license][license-src]][license-href]

Built for `mediainfo` `stable` `23.10`.

## Install

```bash
go install github.com/ewilan-riviere/spotlight@latest
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
[tests-src]: https://img.shields.io/github/actions/workflow/status/ewilan-riviere/notifier/run-tests.yml?branch=main&label=tests&style=flat&colorA=18181B
[tests-href]: https://packagist.org/packages/ewilan-riviere/notifier
[license-src]: https://img.shields.io/github/license/ewilan-riviere/spotlight.svg?style=flat&colorA=18181B&colorB=00ADD8
[license-href]: https://github.com/ewilan-riviere/spotlight/blob/main/LICENSE
