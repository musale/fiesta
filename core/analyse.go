package core

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/etowett/fiesta/utils"
)

// CostData
type CostData struct {
	Username string `json:"username"`
	Amount   string `json:"amount"`
}

type UsageData struct {
	Costs []CostData `json:"costs"`
	Total string     `json:"total"`
}

func RangePage(w http.ResponseWriter, r *http.Request) {
	start := r.FormValue("start")
	stop := r.FormValue("stop")
	mail := r.FormValue("mail")
	// dest := "etowett@gmail.com,kkk@kkk.com"

	data := getUsageData(start, stop)

	if mail == "True" {
		go mailData(data)
	}

	w.WriteHeader(200)
	w.Header().Set("Server", "G-Analytics")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func getUsageData(start, stop string) UsageData {

	utils.Logger.Println("Get usage from: ", start, " to: ", stop)

	var total string
	var usageData UsageData

	err := utils.DbCon.QueryRow("select sum(cost) as cost from bsms_smsrecipient where time_sent>=? and time_sent<=?", start, stop).Scan(&total)

	if err != nil {
		utils.Logger.Println("get total error: ", err)
		return usageData
	}

	usageData.Total = total

	stmt, err := utils.DbCon.Prepare("select u.username, sum(r.cost) from auth_user u join bsms_smsrecipient r on u.id=r.user_id where r.time_sent>? and r.time_sent<? group by u.username")

	if err != nil {
		utils.Logger.Println("slect costs error: ", err)
		return usageData
	}

	defer stmt.Close()

	rows, err := stmt.Query(start, stop)

	if err != nil {
		utils.Logger.Println("exec costs error: ", err)
		return usageData
	}

	defer rows.Close()

	var costs []CostData
	for rows.Next() {
		var usage CostData
		err := rows.Scan(&usage.Username, &usage.Amount)
		if err != nil {
			utils.Logger.Println("scan costs error: ", err)
			return usageData
		}
		costs = append(costs, usage)
	}

	usageData.Costs = costs

	return usageData
}

func mailData(data UsageData) {
	fileLoc := "/tmp/data.csv"

	createCsv(data.Costs, fileLoc)
	msg := `
Hi,\n

Total Summary Usage: KES %v.\n

Sincerely.
`
	subj := fmt.Sprintf("Summary Stat Usage")
	body := fmt.Sprintf(msg, fmt.Sprintf("%.2s", data.Total))
	dest := []string{"etowett@focusmobile.co"}

	utils.SendMail(subj, body, dest, fileLoc)

	return
}

func createCsv(costs []CostData, loc string) {
	file, err := os.Create(loc)
	utils.CheckError("Cannot create file", err)
	defer file.Close()

	writer := csv.NewWriter(file)
	err = writer.Write([]string{"USERNAME", "COST"})
	utils.CheckError("Cannot write to file", err)

	for _, cost := range costs {
		err := writer.Write([]string{cost.Username, cost.Amount})
		utils.CheckError("Cannot write to file", err)
	}

	defer writer.Flush()
}
