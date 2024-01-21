package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/ibmdb/go_ibm_db"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load() // load .env file
	if err != nil {
		log.Fatal(err)
	}

	db2_host := os.Getenv("DB2HOST")
	db2_name := os.Getenv("DBNAME")
	db2_uid := os.Getenv("DB2UID")
	db2_pwd := os.Getenv("DB2INST1_PASSWORD")

	con := "HOSTNAME=" + db2_host +
		";DATABASE=" + db2_name +
		";PORT=50000" +
		";UID=" + db2_uid +
		";PWD=" + db2_pwd

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
