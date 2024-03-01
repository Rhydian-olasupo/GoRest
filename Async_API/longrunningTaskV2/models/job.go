package models

import (
	"time"

	"github.com/google/uuid"
)

//Job represents UUID of a job

type Job struct {
	ID        uuid.UUID   `json:"uuid"`
	Type      string      `json:"type"`
	ExtraData interface{} `json:"extradata"`
}

//Wokrer A data

type Log struct {
	ClientTime time.Time `json:"client_time"`
}

//Callback (webhook) data

type CallBack struct {
	CallBackUrl string `json:"callback _url"`
}

//Email Job

type Mail struct {
	EmailAddress string `json:"email_address"`
}
