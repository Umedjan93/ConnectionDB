package models

import (
	"encoding/xml"
	"time"
)

type xmlForServices struct {
	XMLName xml.Name  `xml:"xml_name"`
	Service []Service `xml:"service"`
}

type Services struct {
	XMLName      xml.Name  `xml:"xml_name"`
	id           int       `json:"id,omitempty" xml:"id"`
	serviceName  string    `json:"service_name,omitempty" xml:"service_name"`
	merchantID   int       `json:"merchant_id,omitempty" xml:"merchant_id"`
	merchantName string    `json:"merchant_name,omitempty" xml:"merchant_name"`
	price        int       `json:"price,omitempty" xml:"price"`
	creationDate time.Time `json:"creation_date" xml:"creation_date"`
	upDate       time.Time `json:"up_date" xml:"up_date"`
	deleted      bool      `json:"deleted,omitempty" xml:"deleted"`
}
