package controllers

/*
Utilisation du module
- https://github.com/MichaelS11/go-scheduler

	The Cron format is in the form of:

	Seconds, Minutes, Hours, Day of month, Month, Day of week, Year

	Field name     Mandatory?   Allowed values    Allowed special characters
	----------     ----------   --------------    --------------------------
	Seconds        No           0-59              * / , -
	Minutes        Yes          0-59              * / , -
	Hours          Yes          0-23              * / , -
	Day of month   Yes          1-31              * / , - L W
	Month          Yes          1-12 or JAN-DEC   * / , -
	Day of week    Yes          0-6 or SUN-SAT    * / , - L #
	Year           No           1970-2099         * / , -
*/
import (
	"github.com/MichaelS11/go-scheduler"
	"github.com/astaxie/beego"
)

func init() {

	s := scheduler.NewScheduler()
	err := s.Make("myJob", "5 5 * * * * *", job, "myData")
	if err != nil {
		beego.Error("Make error:", err)
	}
	// Starts the job schedule. Job will run at it's next run time.
	beego.Debug("scheduler", "Start myJob")
	s.Start("myJob")
}

func job(dataInterface interface{}) {
	data := dataInterface.(string)
	beego.Debug("job", data)
}
