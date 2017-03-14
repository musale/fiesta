package core

import (
	"log"
	"time"
)

func CalcUsage() {

	nw := time.Now()
	start := nw.Format("2006-01-02") + " 00:00:00"
	stop := nw.Format("2006-01-02") + " 23:59:59"

	var username, amount string

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

	for rows.Next() {
		var out OutboxData
		var tme string
		var cost sql.NullString
		err := rows.Scan(&out.MessageID, &out.SendType, &out.SenderID, &out.Currency, &tme, &out.Message, &cost, &out.RecCount)
		if err != nil {
			utils.Logger.Println("error scan out", err)
			return outbox
		}

		if cost.Valid {
			out.Cost = cost.String
		} else {
			out.Cost = "0.00"
		}

		layout := "2006-01-02 15:04:05"
		t, err := time.Parse(layout, tme)

		if err != nil {
			fmt.Println(err)
		}
		out.TimeSent = t
		outbox = append(outbox, out)
	}

	return outbox
}
