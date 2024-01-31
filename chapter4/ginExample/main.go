package main

import (
	"database/sql"
	"go_trial/gorest/chapter4/dbutils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "modernc.org/sqlite"
)

var DB *sql.DB

type StationResource struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	OpeningTime string `json:"opening_time"`
	ClosingTime string `json:"closing_time"`
}

//Get station retunr the station detail

func GetStation(c *gin.Context) {
	var station StationResource

	id := c.Param("station_id")
	err := DB.QueryRow("select ID,NAME,CAST(OPENING_TIME as CHAR),CAST(CLOSING_TIME as CHAR) from station where id =?", id).Scan(&station.ID, &station.Name, station.OpeningTime, &station.ClosingTime)
	if err != nil {
		log.Panicln(err)
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(200, gin.H{
			"result": station,
		})
	}
}

func CreateStation(c *gin.Context) {
	var station StationResource

	if err := c.BindJSON(&station); err != nil {
		c.String(http.StatusBadRequest, "Invalid JSON payload")
		return
	}

	statement, err := DB.Prepare("INSERT INTO station (NAME, OPENING_TIME, CLOSING_TIME) VALUES (?, ?, ?)")
	if err != nil {
		c.String(http.StatusInternalServerError, "Error preparing SQL statement")
		return
	}
	defer statement.Close()

	result, err := statement.Exec(station.Name, station.OpeningTime, station.ClosingTime)
	if err != nil {
		c.String(http.StatusInternalServerError, "Error executing SQL statement")
		return
	}

	newID, err := result.LastInsertId()
	if err != nil {
		c.String(http.StatusInternalServerError, "Error getting last insert ID")
		return
	}

	station.ID = int(newID)
	c.JSON(http.StatusOK, gin.H{
		"result": station,
	})
}

func RemoveStation(c *gin.Context) {
	id := c.Param("station-id")
	statement, _ := DB.Prepare("delete from station where id=?")
	_, err := statement.Exec(id)
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
	} else {
		c.String(http.StatusOK, "")
	}
}

func main() {
	var err error
	DB, err = sql.Open("sqlite", "./railapi.db")
	if err != nil {
		log.Println("Driver creation failed!")
	}
	dbutils.Initialize(DB)
	r := gin.Default()
	// Add routes to REST verbs
	r.GET("/v1/stations/:station_id", GetStation)
	r.POST("/v1/stations", CreateStation)
	r.DELETE("/v1/stations/:station_id", RemoveStation)
	r.Run(":8000") // Default listen and serve on 0.0.0.0:8080
}
