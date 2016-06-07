package main

import (
	"fmt"
    "database/sql"
   // "github.com/go-sql-driver/mysql"
     _ "github.com/go-sql-driver/mysql"
    "github.com/ant0ine/go-json-rest/rest"
    "log"
    "net/http"
    "sync"

)

func main() {
    api := rest.NewApi()
    api.Use(rest.DefaultDevStack...)
    router, err := rest.MakeRouter(
        rest.Get("/getData", GetData),      
    )
    if err != nil {
        log.Fatal(err)
    }
    api.SetApp(router)
    log.Fatal(http.ListenAndServe(":8080", api.MakeHandler()))
}

type Country struct {
    Code string
    Name string
}

var store = map[string]*Country{}

var lock = sync.RWMutex{}

func GetData(w rest.ResponseWriter, r *rest.Request) {    
    
        getDataFromMysql()
        
        result := "True"
        w.WriteJson(result)
}


func getDataFromMysql() {
    //con, err := sql.Open("mysql", store.user+":"+store.password+"@/"+store.database)
    con, err := sql.Open("mysql", "root:novell@/mydb")
    if err != nil {
        panic(err.Error())  // Just for example purpose. You should use proper error handling instead of panic
    }else{}
    defer con.Close()

	//rows, err := con.Query("select password from salondb where email='salona'")
	//if err != nil { /* error handling */panic(err.Error())}

    rows, _ := con.Query("SELECT * FROM salondb;")

    columns, _ := rows.Columns()


//	fmt.Println(rows)

   count := len(columns)
    values := make([]interface{}, count)
    valuePtrs := make([]interface{}, count)

    for rows.Next() {

        for i, _ := range columns {
            valuePtrs[i] = &values[i]
        }

        rows.Scan(valuePtrs...)

        for i, col := range columns {

            var v interface{}

            val := values[i]

            b, ok := val.([]byte)

            if (ok) {
                v = string(b)
            } else {
                v = val
            }

            fmt.Println(col, v)
        }
    }


	// items := make([]*SomeStruct, 0, 10)
	// var ida, idb uint
	// for rows.Next() {
 //    	err = rows.Scan(&ida, &idb)
 //    	if err != nil { /* error handling */}
 //    	items = append(items, &SomeStruct{ida, idb})
 //    }
}
/* 
Install MySQL for windows;

C:\Program Files\MySQL\MySQL Server 5.7\bin>mysql -u root -p
Enter password: ******


mysql> CREATE DATABASE mydb;

mysql> USE mydb;

mysql> CREATE TABLE salondb(
 email VARCHAR(100) NOT NULL,
 password VARCHAR(40) NOT NULL,
 PRIMARY KEY ( email )
 );

INSERT INTO salondb (email, password) VALUES ("salona", "mypass");
INSERT INTO salondb (email, password) VALUES ("salonb", "mypass");
INSERT INTO salondb (email, password) VALUES ("salonc", "mypass");

SELECT * salondb;

*/