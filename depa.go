package depa

import (
	"fmt"

	cmn "github.com/rhuairahrighairidh/go-mod-test-common"
)

func GetVersion() string {
	return fmt.Sprintf("depA: %s, common: %s", "v0.1.0", cmn.GetVersion())
}
