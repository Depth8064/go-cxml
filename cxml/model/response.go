package model

import "encoding/xml"

type Response struct {
	XMLName     xml.Name      `xml:"Response"`
	Status      *Status       `xml:"Status,omitempty"`
	OrderResponse *OrderResponse `xml:"OrderResponse,omitempty"`
}

// OrderResponse represents a cXML order response.
type OrderResponse struct {
	XMLName xml.Name `xml:"OrderResponse"`
	OrderResponseHeader *OrderResponseHeader `xml:"OrderResponseHeader,omitempty"`
}

type OrderResponseHeader struct {
	XMLName  xml.Name `xml:"OrderResponseHeader"`
	OrderID  string   `xml:"orderID,attr,omitempty"`
	OrderDate string  `xml:"orderDate,attr,omitempty"`
	Status   *Status  `xml:"Status,omitempty"`
}

// Status is used for both Response and top-level success/failure.
type Status struct {
	XMLName xml.Name `xml:"Status"`
	Code    string   `xml:"code,attr,omitempty"`
	Text    string   `xml:"text,attr,omitempty"`
}
