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

You have to install [`mediainfo`](https://mediaarea.net/en/MediaInfo) with brew, apt or scoop if you want to use `metadata` command.

## Usage

Options:

- `-o|--output`: output file path, if not set, output is printed in stdout

### Parse

```bash
files parse /path/to/directory
```

### Metadata

```bash
files metadata /path/to/file
```

## License

[MIT](LICENSE) © Ewilan Rivière

[go-version-src]: https://img.shields.io/static/v1?style=flat&label=Go&message=v1.21&color=00ADD8&logo=go&logoColor=ffffff&labelColor=18181b
[go-version-href]: https://go.dev/
[tests-src]: https://img.shields.io/github/actions/workflow/status/ewilan-riviere/files/run-tests.yml?branch=main&label=tests&style=flat&colorA=18181B
[tests-href]: https://github.com/ewilan-riviere/scanner/actions
[license-src]: https://img.shields.io/github/license/ewilan-riviere/files.svg?style=flat&colorA=18181B&colorB=00ADD8
[license-href]: https://github.com/ewilan-riviere/scanner/blob/main/LICENSE
