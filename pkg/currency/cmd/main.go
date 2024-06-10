package main

import (
	c "github.com/frangdelsolar/todo_cli/pkg/currency"
)

func main(){
	c.InitCurrency()

	c.DownloadRates()
	// now:= time.Date(2021, 8, 30, 0, 0, 0, 0, time.UTC)
	// rate, err :=c.GetRateByDate(now, c.Blue)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(rate)
}





