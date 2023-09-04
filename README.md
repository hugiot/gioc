# GIOC

A simple go IOC container.


## Authors

- [@xiaocailc](https://github.com/xcocx)


## Installation

### Prerequisites

- **[Go](https://go.dev/)** >= 1.20

With [Go module](https://github.com/golang/go/wiki/Modules) support, simply add the following import

```bash
import "github.com/hugiot/gioc/src/container"
```

### Demo

**NOTICE:** 

All gioc examples have been moved as standalone repository to [here](https://github.com/hugiot/gioc-examples).

```go
package main

import (
	"github.com/hugiot/gioc/examples/ioc/provider"
	"github.com/hugiot/gioc/examples/ioc/service"
	"github.com/hugiot/gioc/src/container"
	"go.uber.org/zap"
)

func main() {
	container.AddServerProvider(&provider.AppServiceProvider{})
	container.Boot()

	logger := container.Make(service.Logger).(*zap.Logger)
	defer func() {
		_ = logger.Sync()
	}()

	logger.Debug("this is debug")
	logger.Info("this is info")
	logger.Warn("this is warn")
	logger.Error("this is error")
}
```

## License

[MIT](https://github.com/hugiot/gioc/blob/master/LICENSE)

