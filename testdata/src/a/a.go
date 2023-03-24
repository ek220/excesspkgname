package a

import (
	_ "database/sql"
	"fmt"
	http "net/http" // want "excess package name: http \"net/http\""
	os2 "os"
	sync "sync" // want "excess package name: sync \"sync\""
)

func Foo() {
	fmt.Println(os2.Args, http.StatusOK, &sync.Mutex{})

}
