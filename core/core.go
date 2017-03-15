package core

import (
	"encoding/csv"
	"fmt"
	"github.com/etowett/fiesta/utils"
	"github.com/scorredoira/email"
	"log"
	"net/mail"
	"net/smtp"
	"os"
	"time"
)

type CostData struct {
	Username, Amount string
}

func CalcUsage() {

	nw := time.Now()
	dt := nw.Format("2006-01-02")
	start := dt + " 00:00:00"
	stop := dt + " 23:59:59"

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

	loc := "/tmp/data.csv"

	createCsv(costs, loc)

	var total string

	err = utils.DbCon.QueryRow("select sum(cost) as cost from bsms_smsrecipient where time_sent>=? and time_sent<=?", start, stop).Scan(&total)

	subl := fmt.Sprintf("Day's stats for %v", dt)
	body := fmt.Sprintf("Hi,\nTotal (<b style='background: red;'>Summary</b>) Usage: %v", total)

	// compose the message
	// m := email.NewMessage(subl, body)
	m := email.NewHTMLMessage(subl, body)
	m.From = mail.Address{
		Name: "SMSLeopard NoReply", Address: "noreply@smsleopard.com",
	}
	m.To = []string{"etowett@focusmobile.co"}

	err = m.Attach(loc)
	utils.CheckError("Cannot attach file", err)

	auth := smtp.PlainAuth("", "noreply@smsleopard.com", "autocook25#", "smtp.gmail.com")
	err = email.Send("smtp.gmail.com:587", auth, m)
	utils.CheckError("Cannot send mail:", err)

	err = os.Remove(loc)
	utils.CheckError("cannot delete file:", err)

	return
}
