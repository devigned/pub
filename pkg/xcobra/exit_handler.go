//+build !noexit

package xcobra

import (
	"os"
)

func exitWithCode(err error) {
	if e, ok := err.(ErrorWithCode); ok {
		os.Exit(e.Code)
	}
	os.Exit(1)
}
