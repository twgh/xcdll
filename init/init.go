package init

import (
	"github.com/twgh/xcdll"
	"os"
)

func init() {
	_ = os.WriteFile("xcgui.dll", xcdll.DLL, 0777)
}
