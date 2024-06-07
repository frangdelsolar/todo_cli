package period

import (
	"strconv"
	"todo_cli/data"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var CreateEffectivePeriodCmd = &cobra.Command{
	Use:   "add",
	Short: "Add Efective Periods",
	Run: func(cmd *cobra.Command, args []string) {
		log.Info().Msg("Adding Effective Period to task")

		taskId := cmd.Flag("taskId").Value.String()
		taskIdInt, err := strconv.Atoi(taskId)
		if err != nil {
			log.Err(err).Msg("Error parsing taskId")
			return
		}

		startDate := cmd.Flag("startDate").Value.String()
		endDate := cmd.Flag("endDate").Value.String()
		frequency := cmd.Flag("frequency").Value.String()
		_, err = data.CreateEffectivePeriod(uint(taskIdInt), startDate, endDate, frequency)
		if err != nil {
			log.Err(err).Msg("Error adding effective period to task")
			return
		}

		log.Info().Msg("Effective Period added to task")
	},
}

func init() {
	CreateEffectivePeriodCmd.Flags().IntP("taskId", "t", 0, "The id of the task")
	CreateEffectivePeriodCmd.Flags().StringP("startDate", "s", "", "The start date of the effective period. Default is the current date.")
	CreateEffectivePeriodCmd.Flags().StringP("endDate", "e", "", "The end date of the effective period. If not provided, the task will be active until it is an end date is provided.")
	CreateEffectivePeriodCmd.Flags().StringP("frequency", "f", "", "The frequency of the task within the effective period (daily, weekly, monthly, yearly). Default is monthly.")
}
