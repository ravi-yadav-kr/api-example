package db

import (
    "database/sql"
    _ "github.com/lib/pq"
    log "github.com/sirupsen/logrus"
)

const (
    Host     = "localhost"
    Port     = 5432
    User     = "myuser"
    Password = "mypassword"
    Dbname   = "mydb"
)

func InitialiseDb(connStr string) *sql.DB {
    db, err := sql.Open("postgres", connStr)
    if err != nil {
        log.Errorf("Could not Initialise Db")
        return nil
    }

    query := `CREATE TABLE IF NOT EXISTS employee(
    id varchar(36) NOT NULL PRIMARY KEY,
    name varchar(50), 
    age char(4), 
    dept_name varchar(10),
    phone_number varchar(12),
    address varchar(256),
    joining_date Date
    );`

    _, err = db.Exec(query)
    if err != nil {
        log.Error("Could not create table employee", err)
        return nil
    }

    log.Info("Db Initialised Successfully")

    log.WithFields(log.Fields{
        "animal": "walrus",
        "number": 1,
        "size":   10,
    }).Info("A walrus appears")

    return db
}
