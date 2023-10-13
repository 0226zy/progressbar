# ProgressBar 

> simple progress bar at the terminal

## Installation

```shell
go mod init github.com/my/repo
```

Then install progressbar

```shell
go get github.com/0226zy/progressbar
```

## Quickstart

```go
package main

import (
        "fmt"
        "time"

        "github.com/0226zy/progressbar"
)

func main() {
        bar := progressbar.NewProgressBar(100)

        ch := bar.Start()

        for i := int64(1); i < 100; i++ {
                ch <- 1
                time.Sleep(100 * time.Millisecond)
        }
        bar.Close()

        fmt.Println("\n========== again ==========")
        // begin again
        ch = bar.Start()
        for i := int64(1); i < 100; i++ {
                ch <- 1
                time.Sleep(50 * time.Millisecond)
        }
        bar.Close()

```

## Contributors

