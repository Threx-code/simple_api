package config

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/Threx-code/simple_api/utils"

	_ "github.com/go-sql-driver/mysql"
)

var (
	connection *sql.DB
)

var server, err = utils.LoadConfig("../../")

func dns(dbName string) string {
	if dbName == "" {
		return fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", server.DBRoot, server.DBRootPassword, server.DBHost, "")
	}

	return fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", server.DBUser, server.DBPassword, server.DBHost, server.DBName)
}

func CreateDatabase() {
	db, err := sql.Open("mysql", dns(""))

	if err != nil {
		fmt.Printf("unable to create connection %s", err.Error())
		return
	}
	defer db.Close()

	ctx, cancelfunc := context.WithTimeout(context.Background(), time.Second*5)
	defer cancelfunc()
	query := "CREATE DATABASE IF NOT EXISTS " + server.DBName

	res, err := db.ExecContext(ctx, query)

	if err != nil {
		fmt.Printf("unable to create database %s", err.Error())
		return
	}

	_, err = db.ExecContext(ctx, "CREATE USER IF NOT EXISTS "+server.DBUser+"@'localhost' IDENTIFIED BY '"+server.DBPassword+"'")

	if err != nil {
		fmt.Printf("unable to create user %s", err.Error())
		return
	}

	_, err = db.ExecContext(ctx, "GRANT ALL PRIVILEGES ON "+server.DBName+".* TO "+server.DBUser+"@localhost")

	if err != nil {
		fmt.Printf("unable to grant permission to user %s", err.Error())
		return
	}

	_, err = db.ExecContext(ctx, "FLUSH PRIVILEGES ")

	if err != nil {
		fmt.Printf("unable to flush privileges %s", err.Error())
		return
	}

	row, _ := res.RowsAffected()

	fmt.Printf("rows affected %d", row)
}

func Connect() {
	db, err := sql.Open("mysql", dns(server.DBName))

	if err != nil {
		fmt.Printf("unable to create connection %s", err.Error())
		return
	}
	defer db.Close()

	ctx, cancelfunc := context.WithTimeout(context.Background(), time.Second*5)
	defer cancelfunc()

	db.SetConnMaxIdleTime(100)
	db.SetConnMaxLifetime(100)
	db.SetMaxOpenConns(20)

	err = db.PingContext(ctx)
	if err != nil {
		fmt.Printf("unable to ping database %s", err.Error())
		return
	}

	connection = db
}

func GetDB() *sql.DB {
	return connection
}
