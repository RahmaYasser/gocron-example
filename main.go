package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-co-op/gocron"
	_ "github.com/jackc/pgx/v4/stdlib"
)
func main(){

	// Deletes all notifications every start of the day
	s := gocron.NewScheduler(time.UTC)

	s.Every(1).Day().At("00:00").Do(DeleteJob)

	s.StartAsync()

	http.ListenAndServe(":8090", nil)
}

func PrintJob(){
	fmt.Println("hello rahma")
}
func DeleteJob(){
	db:=GetDbConnetion()
	sqlStatement := `DELETE FROM notifications`
	_, err := db.Exec(sqlStatement)
	if err != nil {
  		fmt.Println(err)
	}
}


func GetDbConnetion() *sql.DB {
	dataSourceName := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s",
	os.Getenv("DB_HOST"),
	os.Getenv("DB_PORT"),
	os.Getenv("DB_NAME"),
	os.Getenv("DB_USER"),
	os.Getenv("DB_PASSWORD"))
	db, err := sql.Open("pgx", dataSourceName)
	if err != nil {
		log.Fatal(fmt.Sprintf("unable to conect to db"))
		panic(err)
	}
	log.Println("connected to db ")

	//test connection
	err = db.Ping()
	if err != nil {
		fmt.Println(err)
		log.Fatal("cannot ping db")
		panic(err)
	}
	log.Println("pinged db")
	return db
}
