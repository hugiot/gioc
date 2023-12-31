# GIOC

A simple go IOC container.


## Authors

- [@xiaocailc](https://github.com/xcocx)


## Installation

### Prerequisites

- **[Go](https://go.dev/)** >= 1.20

With [Go module](https://github.com/golang/go/wiki/Modules) support, simply add the following import

```bash
import "github.com/hugiot/gioc"
```

### Demo

```go
package main

import (
	"github.com/hugiot/gioc"
	"log"
	"os"
)

const (
	LogService string = "log"
)

type AppServiceProvider struct {
}

func (a AppServiceProvider) Register(c gioc.Container) {
	c.Single(LogService, func() any {
		return log.New(os.Stderr, "custom | ", log.LstdFlags)
	})
}

func (a AppServiceProvider) Boot(c gioc.Container) {
}

func main() {
	gioc.AddServerProvider(&AppServiceProvider{})
	gioc.Boot()
	logger := gioc.Make(LogService).(*log.Logger)
	logger.Println("this is content")
}

```

## License

[MIT](https://github.com/hugiot/gioc/blob/master/LICENSE)

