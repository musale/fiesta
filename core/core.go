package core

import (
	"log"
	"time"

	"github.com/etowett/fiesta/utils"
)

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

	// var data []map[string]string

	for rows.Next() {
		// var usage map[string]string
		var username, amount string
		err := rows.Scan(&username, &amount)
		if err != nil {
			log.Println("error scan", err)
		}

		// usage["username"] = username
		// usage["amount"] = amount

		// data = append(data, usage)
		log.Println(username, ": ", amount)
	}

	// log.Println(data)
}
