# gif
[![godoc](https://godoc.org/github.com/ItsJimi/gif/pkg/convert?status.svg)](https://godoc.org/github.com/ItsJimi/gif/pkg/convert)

🏞 The simplest way to create gif from video

## Install
Install [ffmpeg](https://ffmpeg.org/) before running `gif`.

Use precompiled versions in [releases page](https://github.com/ItsJimi/gif/releases)

or

```shell
go get -u github.com/itsjimi/gif
```

## Usage
### CLI
#### Help
```shell
gif help
```
#### Convert
```shell
gif convert --help
```
### Code
#### Convert from folder
```go
package main

import "github.com/ItsJimi/gif/pkg/convert"

func main() {
  options := convert.Options{
    FPS:   30,
    Scale: -1,
    Crop:  "",
  }

  convert.FromFolder("./videos", "./gifs", options)
}
```

## Contribute
Feel free to fork and make pull requests

## License
[MIT](https://github.com/ItsJimi/gif/blob/master/LICENSE)
