package models

import (
	"encoding/xml"
	"time"
)

type xmlForTransactions struct {
	XMLName xml.Name `xml:"xml_name"`
	Transactions []Transactions `xml:"transactions"`
}

type Transactions struct {
	XMLName        xml.Name  `xml:"xml_name"`
	id             int       `json:"id,omitempty" xml:"id"`
	payerAcc       string    `json:"payer_acc,omitempty" xml:"payer_acc"`
	beneficiaryAcc string    `json:"beneficiary_acc,omitempty" xml:"beneficiary_acc"`
	transferAmount int       `json:"transfer_amount,omitempty" xml:"transfer_amount"`
	creationDate   time.Time `json:"creation_date" xml:"creation_date"`
	completed      bool      `json:"completed,omitempty" xml:"completed"`
}