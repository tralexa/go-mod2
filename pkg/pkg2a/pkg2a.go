package pkg2a

import (
	"github.com/tralexa/go-mod1/pkg/pkg1b"
	"github.com/tralexa/go-mod2/pkg/pkg2b"
)

func Do() {
	pkg1b.Do()
	pkg2b.Do()
}