package models

import "time"

type Activity struct {
	UserID   string    `json:"user_id"`
	Username string    `json:"username"`
	Status   string    `json:"status"`
	Platform string    `json:"platform"`
	Activity string    `json:"activity"`
	Type     string    `json:"type"`
	Details  string    `json:"details"`
	State    string    `json:"state"`
	Updated  time.Time `json:"updated"`
}
