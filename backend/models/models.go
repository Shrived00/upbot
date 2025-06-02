package models

import "time"

type User struct {
	ID    string `json:"id" bson:"_id,omitempty"`
	Email string `json:"email" bson:"email"`
	Tasks []Task `json:"tasks" bson:"tasks,omitempty"`
}

type Task struct {
	ID        string `json:"id" bson:"_id,omitempty"`
	URL       string `json:"url" bson:"url"`
	IsActive  bool   `json:"isActive" bson:"isActive"`
	UserID    string `json:"userId" bson:"userId"`
	FailCount int    `json:"failCount" bson:"failCount"`
	Logs      []Log  `json:"logs" bson:"logs,omitempty"`
}

type Log struct {
	Time        time.Time `json:"time" bson:"time"`
	TimeTake    int64     `json:"timeTake" bson:"timeTake"`
	LogResponse string    `json:"logResponse" bson:"logResponse"`
	IsSuccess   bool      `json:"isSuccess" bson:"isSuccess"`
	RespCode    int       `json:"respCode" bson:"respCode"`
}
