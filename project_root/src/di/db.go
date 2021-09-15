package di

import (
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "os"
)

func CreateDb() *gorm.DB {
    
    host    := "localhost" // os.Getenv("POSTGRES_HOST")
    user    := os.Getenv("POSTGRES_USER")
    pass    := os.Getenv("POSTGRES_PASSWORD")
    dbname  := os.Getenv("POSTGRES_DB")
    port    := "5432" // os.Getenv("POSTGRES_PORT")
    sslMode := os.Getenv("POSTGRES_DB_SSL_MODE")
    tz      := "Asia/Shanghai" // os.Getenv("POSTGRES_DB_SSL_MODE")
    
    
    dsn :=  "host="     + host      + " " +
            "user="     + user      + " " +
            "password=" + pass      + " " +
            "dbname="   + dbname    + " " +
            "port="     + port      + " " +
            "sslmode="  + sslMode   + " " +
            "TimeZone=" + tz
    
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
        // TODO: add adapter for logger
        // Logger: di.CreateLogger(),
    })
    if err != nil {
        panic("failed to connect database")
    }
    
    return db
}
