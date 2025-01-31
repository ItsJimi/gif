# gif
[![go reference](https://pkg.go.dev/badge/github.com/ItsJimi/gif)](https://pkg.go.dev/github.com/ItsJimi/gif/pkg/convert)

🏞 The simplest way to create gif from video

## Install
Install [ffmpeg](https://ffmpeg.org/) before running `gif`.

### CLI

Use precompiled versions in [releases page](https://github.com/ItsJimi/gif/releases)

or

```shell
go install github.com/ItsJimi/gif@latest
```

### Code

```shell
go get -u github.com/ItsJimi/gif
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
gif convert my-video.mp4 my-gif.gif
gif convert my-video.mp4 my-gif.gif --fps 60 --scale 1280
gif convert ./folder ./another-folder --crop "1280:720:30:60"
```
### Code
#### Convert from folder
```go
package main

import (
  "fmt"
  "github.com/ItsJimi/gif/pkg/convert"
)

func main() {
  options := convert.Options{
    FPS:   30,
    Scale: -1,
    Crop:  "",
  }

  err := convert.FromFolder("./videos", "./gifs", options)
  if err != nil {
    fmt.Println(err)
  }
}
```

## Contribute
Feel free to fork and make pull requests

## License
[MIT](https://github.com/ItsJimi/gif/blob/master/LICENSE)
