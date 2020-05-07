package depa

import (
	"fmt"

	cmn "github.com/rhuairahrighairidh/go-mod-test-common"
)

func GetVersion() string {
	cmn.NewFunctionality()
	return fmt.Sprintf("depA: %s, common: %s", "v0.1.1", cmn.GetVersion())
}
