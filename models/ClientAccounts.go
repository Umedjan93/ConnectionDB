package models

import (
	"encoding/xml"
	"time"
)

type xmlForClientAccounts struct {
	XMLName        xml.Name         `xml:"xml_name"`
	ClientAccounts []ClientAccounts `xml:"client_accounts"`
}

type ClientAccounts struct {
	XMLName       xml.Name  `xml:"xml_name"`
	id            int       `json:"id,omitempty" xml:"id"`
	accNumber     string    `json:"acc_number,omitempty" xml:"acc_number"`
	clientID      int       `json:"client_id,omitempty" xml:"client_id"`
	balance       int       `json:"balance,omitempty" xml:"balance"`
	swift         string    `json:"swift,omitempty" xml:"swift"`
	LastVisitDate time.Time `json:"last_visit_date" xml:"last_visit_date"`
	locked        bool      `json:"locked,omitempty" xml:"locked"`
	creationDate  time.Time `json:"creation_date" xml:"creation_date"`
	upDate        time.Time `json:"up_date" xml:"up_date"`
	deleted       bool      `json:"deleted,omitempty" xml:"deleted"`
}
