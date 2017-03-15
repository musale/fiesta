package core

import (
	"encoding/csv"
	"github.com/etowett/fiesta/utils"
	"log"
	"os"
	"time"
)

type CostData struct {
	Username, Amount string
}

type Result struct {
	Costs              []CostData
	Start, Stop, Total string
}

func CalcUsage() {

	nw := time.Now()
	start := nw.Format("2006-01-02") + " 00:00:00"
	stop := nw.Format("2006-01-02") + " 23:59:59"

	stmt, err := utils.DbCon.Prepare("select u.username, sum(r.cost) from auth_user u join bsms_smsrecipient r on u.id=r.user_id where r.time_sent>? and r.time_sent<? group by u.username")

	if err != nil {
		log.Fatal("prepare select out", err)
	}

	defer stmt.Close()

	rows, err := stmt.Query(start, stop)

	if err != nil {
		log.Fatal("query select out", err)
	}

	defer rows.Close()

	var costs []CostData
	for rows.Next() {
		var usage CostData
		err := rows.Scan(&usage.Username, &usage.Amount)
		utils.CheckError("Error scan", err)
		costs = append(costs, usage)
	}

	createCsv(costs)

	err = utils.DbCon.QueryRow("select sum(cost) as cost from bsms_smsrecipient where time_sent>=? and time_sent<=?", start, stop).Scan(&result.Total)

	log.Println("Data: ", result)

	subject := fmt.Sprintf("Day's stats for %v", nw.Format("2006-01-02"))
	message := fmt.Sprintf(
		"Hi,\nTotal Usage Today: %v\n.", total)
	to := "etowett@focusmobile.co"

	utils.SendMail(subject, message, to)

	return
}

func createCsv(costs []CostData) {
	file, err := os.Create("/tmp/data.csv")
	utils.CheckError("Cannot create file", err)
	defer file.Close()

	writer := csv.NewWriter(file)
	err := writer.Write([]string{"USERNAME", "COST"})
	utils.CheckError("Cannot write to file", err)

	for _, cost := range costs {
		err := writer.Write([]string{cost.Username, cost.Amount})
		utils.CheckError("Cannot write to file", err)
	}

	defer writer.Flush()
}
