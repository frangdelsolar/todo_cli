package period

import (
	"github.com/spf13/cobra"
)

var EffectivePeriodCmd = &cobra.Command{
	Use:   "period",
	Short: "Manage Efective Periods",
}

func init() {
	EffectivePeriodCmd.AddCommand(CreateEffectivePeriodCmd)

}
