[![Go](https://github.com/ek220/excesspkgname/actions/workflows/go.yml/badge.svg)](https://github.com/ek220/excesspkgname/actions/workflows/go.yml)

# excesspkgname

A linter that checks excess package name

```
import (
	_ "database/sql" // ok
	"fmt"            // ok
	http "net/http"  // excess package name
	os2 "os"         // ok
	sync "sync"      // // excess package name
)

```

# Install

```
go install github.com/ek220/excesspkgname/cmd/excesspkgname@v1.0.0
```

# Run

```
excesspkgname ./...
```
