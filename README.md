# scanner

[![go][go-version-src]][go-version-href]
[![tests][tests-src]][tests-href]
[![license][license-src]][license-href]

CLI to parse files for a list or to get metadata. Powered by Go.

Metadata parsed by [`mediainfo`](https://mediaarea.net/en/MediaInfo) (v23.10).

## Install

```bash
go install github.com/ewilan-riviere/scanner@latest
```

You have to install [`mediainfo`](https://mediaarea.net/en/MediaInfo) with `brew`, `apt` or `scoop` if you want to use `metadata` command.

## Usage

Options:

- `-o|--output`: output file path, if not set, output is printed in stdout

### Parse

```bash
scanner parse /path/to/directory
```

Output example

```bash
.gitignore
LICENSE
README.md
go.mod
go.sum
main.go
main_test.go
pkg/files/files.go
pkg/mediainfo/convert/convert.go
pkg/mediainfo/mediainfo.go
pkg/mediainfo/output/output.go
pkg/mediainfo/types/types.go
pkg/printer/printer.go
pkg/scan/scan.go
scanner
```

You can save the output in a file

```bash
scanner parse -o=path/to/output.json /path/to/directory
```

### Metadata

```bash
scanner metadata /path/to/file
```

Output example

```bash
UniqueID: empty
CompleteName: ./test/media/test.mp3
Format: MPEG Audio
FormatVersion: empty
FileSize: 266 KiB
Duration: 11 s 49 ms
OverallBitRate: 128 kb/s
Album: P1PDD Le conclave de Troie
AlbumPerformer: P1PDD & Mr Piouf
PartPosition: 1
TrackName: Introduction
TrackNamePosition: 1
Performer: Mr Piouf
Genre: Roleplaying game
RecordedDate: 2016
FrameRate: empty
MovieName: empty
EncodedDate: empty
WritingApplication: empty
WritingLibrary: LAME3.100
Cover: Yes
CoverType: Cover (front)
CoverMime: image/jpeg
Comment: http://www.p1pdd.com
Attachments: empty
Videos: []
Audios:
[ID: , Format: MPEG Audio, FormatInfo: , FormatProfile: Layer 3, FormatVersion: Version 1, CommercialName: , CodecID: , Duration: 11 s 50 ms, BitRateMode: Constant, BitRate: 128 kb/s, Channel: 2 channels, ChannelLayout: , SamplingRate: 44.1 kHz, FrameRate: 38.281 FPS (1152 SPF), Compression: Lossy, StreamSize: 173 KiB (65%), WritingLibrary: LAME3.100, EncodingSettings: -m j -V 4 -q 2 -lowpass 17 -b 128, Title: , Language: , ServiceKind: , Default: , Forced: ]
Texts: []
Images:
[ID: , Format: JPEG, Width: 640 pixels, Height: 640 pixels, ColorSpace: YUV, ChromaSubsampling: 4:2:0, BitDepth: 8 bits, CompressionMode: Lossy, StreamSize: 91.0 KiB (34%)]
Menu: []
```

You can save the output in a file

```bash
scanner metadata -o=path/to/output.json /path/to/file
```

## License

[MIT](LICENSE) © Ewilan Rivière

[go-version-src]: https://img.shields.io/static/v1?style=flat&label=Go&message=v1.21&color=00ADD8&logo=go&logoColor=ffffff&labelColor=18181b
[go-version-href]: https://go.dev/
[tests-src]: https://img.shields.io/github/actions/workflow/status/ewilan-riviere/files/run-tests.yml?branch=main&label=tests&style=flat&colorA=18181B
[tests-href]: https://github.com/ewilan-riviere/scanner/actions
[license-src]: https://img.shields.io/github/license/ewilan-riviere/files.svg?style=flat&colorA=18181B&colorB=00ADD8
[license-href]: https://github.com/ewilan-riviere/scanner/blob/main/LICENSE
