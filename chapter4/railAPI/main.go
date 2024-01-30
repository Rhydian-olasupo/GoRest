package main

import (
	"database/sql"
	"encoding/json"
	"go_trial/gorest/chapter4/dbutils"
	"log"
	"net/http"
	"time"

	"github.com/emicklei/go-restful"
	_ "modernc.org/sqlite"
)

/*func main() {
	//Connect to the database
	db, err := sql.Open("sqlite", "./rail.db")

	if err != nil {
		log.Println("Driver creation failed")
	}
	//Create Tables
	dbutils.Initialize(db)
}*/

// Database driver visible to the whole program
var DB *sql.DB

//Train Resource is the model holding rail infrormation

type TrainResource struct {
	ID              int
	DriverName      string
	OperatingStatus bool
}

type StationResource struct {
	ID          int
	Name        string
	OpeningTime time.Time
	ClosingTime time.Time
}

type ScheduleResource struct {
	ID           int
	TrainID      int
	StationID    int
	ArrrivalTime time.Time
}

//Using go-restful
//Register adds path and routes to a new service instance

func (t *TrainResource) Register(container *restful.Container) {
	ws := new(restful.WebService)
	ws.Path("/v1/trains").Consumes(restful.MIME_JSON).Produces(restful.MIME_JSON)
	ws.Route(ws.GET("/{train-id}").To(t.getTrain))
	ws.Route(ws.POST("").To(t.createTrain))
	ws.Route(ws.DELETE("/{train-id}").To(t.removeTrain))
	container.Add(ws)

}

func (s *StationResource) Register(container *restful.Container) {
	ws := new(restful.WebService)
	ws.Path("/v1/stations").Consumes(restful.MIME_JSON).Produces(restful.MIME_JSON)
	ws.Route(ws.GET("/{station-name}").To(s.getStation))
	ws.Route(ws.POST("").To(s.createStation))
	ws.Route(ws.DELETE("/{station-name}").To(s.removeStation))
	container.Add(ws)
}

//GET http://localhost:8000/v1/stations/ibadan

func (s *StationResource) getStation(request *restful.Request, response *restful.Response) {
	name := request.PathParameter("station-name")
	err := DB.QueryRow("select ID,NAME,OPENING_TIME,CLOSING_TIME FROM station where name =?", name).Scan(&s.ID, &s.Name, &s.OpeningTime, &s.ClosingTime)
	if err != nil {
		log.Println(err)
		response.AddHeader("Content-Type", "text/plain")
		response.WriteErrorString(http.StatusNotFound, "Station could not be found")
	} else {
		response.WriteEntity(s)
	}
}

//POST hhtp://localhost:8000/v1/stations

func (s StationResource) createStation(request *restful.Request, response *restful.Response) {
	decoder := json.NewDecoder(request.Request.Body)
	var z StationResource
	err := decoder.Decode(&z)
	if err != nil {
		log.Println("Error decoding payload:", err)
		response.AddHeader("Content-Type", "text/plain")
		response.WriteErrorString(http.StatusBadRequest, "Invalid payload")
		return
	}

	// No need to parse time again, it's already a time.Time value
	// Use z.OpeningTime and z.ClosingTime directly

	statement, err := DB.Prepare("insert into station (NAME, OPENING_TIME, CLOSING_TIME) values (?, ?, ?)")
	if err != nil {
		log.Println("Error preparing statement:", err)
		response.AddHeader("Content-Type", "text/plain")
		response.WriteErrorString(http.StatusInternalServerError, "Database error")
		return
	}

	defer statement.Close()

	result, err := statement.Exec(z.Name, z.OpeningTime, z.ClosingTime)
	if err != nil {
		log.Println("Error executing statement:", err)
		response.AddHeader("Content-Type", "text/plain")
		response.WriteErrorString(http.StatusInternalServerError, "Database error")
		return
	}

	newID, err := result.LastInsertId()
	if err != nil {
		log.Println("Error getting last insert ID:", err)
		response.AddHeader("Content-Type", "text/plain")
		response.WriteErrorString(http.StatusInternalServerError, "Database error")
		return
	}

	z.ID = int(newID)
	response.WriteHeaderAndEntity(http.StatusCreated, z)
}

//DELETE http://localhost:8000/v1/stations/ibadan

func (s StationResource) removeStation(request *restful.Request, response *restful.Response) {
	name := request.PathParameter("station-name")
	statement, _ := DB.Prepare("delete from station where name =?")
	_, err := statement.Exec(name)
	if err != nil {
		response.AddHeader("Content-Type", "text/plain")
		response.WriteErrorString(http.StatusInternalServerError, err.Error())
	} else {
		response.WriteHeader(http.StatusOK)
	}
}

// Get http://localhost:8000/v1/trains/1
func (t *TrainResource) getTrain(request *restful.Request,
	response *restful.Response) {
	id := request.PathParameter("train-id")
	err := DB.QueryRow("select ID, DRIVER_NAME,OPERATING_STATUS FROM train where id =?", id).Scan(&t.ID, &t.DriverName, &t.OperatingStatus)
	if err != nil {
		log.Println(err)
		response.AddHeader("Content-Type", "text/plain")
		response.WriteErrorString(http.StatusNotFound, "Train could not be found")
	} else {
		response.WriteEntity(t)
	}
}

//Post http://localhost:8000/v1/trains

func (t TrainResource) createTrain(request *restful.Request, response *restful.Response) {
	log.Println(request.Request.Body)
	decoder := json.NewDecoder(request.Request.Body)
	var b TrainResource
	err := decoder.Decode(&b)
	log.Println(b.DriverName, b.OperatingStatus)
	if err != nil {
		log.Println("Cant decode Json")
	}
	statement, _ := DB.Prepare("insert into train (DRIVER_NAME,OPERATING_STATUS) values (?,?)")
	result, err := statement.Exec(b.DriverName, b.OperatingStatus)
	if err == nil {
		newID, _ := result.LastInsertId()
		b.ID = int(newID)
		response.WriteHeaderAndEntity(http.StatusCreated, b)
	} else {
		response.AddHeader("Content-Type", "text/plain")
		response.WriteErrorString(http.StatusInternalServerError, err.Error())
	}
}

//DELTE http://localhost:8000/v1/trains/1

func (t TrainResource) removeTrain(request *restful.Request, response *restful.Response) {
	id := request.PathParameter("train-id")
	statement, _ := DB.Prepare("delete from train where id =?")
	_, err := statement.Exec(id)
	if err != nil {
		response.AddHeader("Content-Type", "text/plain")
		response.WriteErrorString(http.StatusInternalServerError, err.Error())
	} else {
		response.WriteHeader(http.StatusOK)
	}
}

func main() {
	var err error
	DB, err = sql.Open("sqlite", "./railapi.db")
	if err != nil {
		log.Println("Driver creation failed!")
	}

	dbutils.Initialize(DB)
	wsContainer := restful.NewContainer()
	wsContainer.Router(restful.CurlyRouter{})
	t := TrainResource{}
	s := StationResource{}
	s.Register(wsContainer)
	t.Register(wsContainer)
	log.Printf("start listeninig on localhost:8000")
	server := &http.Server{Addr: ":8000", Handler: wsContainer}
	log.Fatal(server.ListenAndServe())
}
