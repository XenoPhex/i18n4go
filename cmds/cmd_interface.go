package cmds

import (
	"github.com/XenoPhex/i18n4go/common"
)

type CommandInterface interface {
	common.PrinterInterface
	Options() common.Options
	Run() error
}
