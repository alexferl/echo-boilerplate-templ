package testing

import (
	"os"
	"path"
	"runtime"

	"github.com/alexferl/echo-boilerplate-templ/config"
)

func init() {
	_, filename, _, _ := runtime.Caller(0)
	dir := path.Join(path.Dir(filename), "..")
	err := os.Chdir(dir)
	if err != nil {
		panic(err)
	}

	c := config.New()
	c.BindFlags()
}
