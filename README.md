# GoCam


### GetListDevices
```go
package main

import (
	"fmt"

	"github.com/AndrusGerman/gocam"
)

func main() {
	devices := gocam.GetListDevices()
	fmt.Println(devices)
}

```