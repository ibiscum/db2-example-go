package main

import (
	"database/sql"
	"fmt"

	_ "github.com/ibmdb/go_ibm_db"
)

func main() {
	con := "HOSTNAME=localhost;DATABASE=;PORT=50000;UID=db2inst1;PWD="
	db, err := sql.Open("go_ibm_db", con)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	query := "SELECT STRINGID, SUBSTR(STRING,1,60) AS STRING FROM SYSIBM.SYSXMLSTRINGS ORDER BY STRINGID FETCH FIRST 3 ROWS ONLY"
	st, err := db.Prepare(query)
	if err != nil {
		fmt.Println(err)
	}

	rows, err := st.Query()
	if err != nil {
		fmt.Println(err)
	}

	cols, _ := rows.Columns()
	fmt.Printf("%s    %s\n", cols[0], cols[1])
	fmt.Println("-------------------------------------")
	defer rows.Close()
	for rows.Next() {
		var strid, str string
		err = rows.Scan(&strid, &str)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("%v        %v\n", strid, str)
	}
	db.Close()
}
