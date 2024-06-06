package taskcmd

import (
	"strconv"
	"todo_cli/data"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var AddEffectivePeriodCmd = &cobra.Command{
	Use:   "period",
	Short: "Add Efective Period",
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
	AddEffectivePeriodCmd.Flags().IntP("taskId", "t", 0, "The id of the task")
	AddEffectivePeriodCmd.Flags().StringP("startDate", "s", "", "The start date of the effective period. Default is the current date.")
	AddEffectivePeriodCmd.Flags().StringP("endDate", "e", "", "The end date of the effective period. If not provided, the task will be active until it is an end date is provided.")
	AddEffectivePeriodCmd.Flags().StringP("frequency", "f", "", "The frequency of the task within the effective period (daily, weekly, monthly, yearly). Default is monthly.")
}
