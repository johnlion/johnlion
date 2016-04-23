package exec

import (
	"github.com/johnlion/sites/web"
)

func DefaultRun() {
	web.SetFlag()
	web.Run()
}
