package models

import "encoding/xml"

type xmlForATMs struct {
	XMLName xml.Name `xml:"xml_name"`
	ATMs []ATMs `xml:"atms"`
}

type ATMs struct {
	XMLName xml.Name `xml:"xml_name"`
	id            int    `json:"id" xml:"id"`
	address       string `json:"address,omitempty" xml:"address"`
	balance       int    `json:"balance,omitempty" xml:"balance"`
	status        string `json:"status,omitempty" xml:"status"`
	maxCashLimit  int    `json:"max_cash_limit,omitempty" xml:"max_cash_limit"`
	commissionFee bool   `json:"commission_fee,omitempty" xml:"commission_fee"`
}
