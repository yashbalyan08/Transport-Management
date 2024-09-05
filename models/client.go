package modules

import (
	_ "github.com/go-sql-driver/mysql"
)

type Client struct {
	ClientId   string `json:"clientid,omitempty"`
	ClientName string `json:"clientname,omitempty"`
}
