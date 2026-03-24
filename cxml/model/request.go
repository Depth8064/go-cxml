package model

import "encoding/xml"

type Request struct {
	XMLName              xml.Name              `xml:"Request"`
	DeploymentMode       string                `xml:"deploymentMode,attr,omitempty"`
	OrderRequest         *OrderRequest         `xml:"OrderRequest,omitempty"`
	OrderChangeRequest   *OrderChangeRequest   `xml:"OrderChangeRequest,omitempty"`
	PunchOutOrderMessage *PunchOutOrderMessage `xml:"PunchOutOrderMessage,omitempty"`
}

func (r *Request) PayloadType() string {
	switch {
	case r.OrderRequest != nil:
		return "OrderRequest"
	case r.OrderChangeRequest != nil:
		return "OrderChangeRequest"
	case r.PunchOutOrderMessage != nil:
		return "PunchOutOrderMessage"
	default:
		return ""
	}
}

type OrderRequest struct {
	XMLName            xml.Name            `xml:"OrderRequest"`
	OrderRequestHeader *OrderRequestHeader `xml:"OrderRequestHeader,omitempty"`
	ItemOut            []*ItemOut          `xml:"ItemOut,omitempty"`
}

type OrderChangeRequest struct {
	XMLName                xml.Name            `xml:"OrderChangeRequest"`
	OrderRequestReference  *OrderRequestHeader `xml:"OrderRequestReference,omitempty"`
	ItemChange             []*ItemOut          `xml:"ItemChange,omitempty"`
}

type OrderRequestHeader struct {
	XMLName   xml.Name `xml:"OrderRequestHeader"`
	OrderID   string   `xml:"orderID,attr,omitempty"`
	OrderDate string   `xml:"orderDate,attr,omitempty"`
	Total     *Money   `xml:"Total>Money,omitempty"`
	ShipTo    *Party   `xml:"ShipTo,omitempty"`
	BillTo    *Party   `xml:"BillTo,omitempty"`
}

type ItemOut struct {
	XMLName    xml.Name    `xml:"ItemOut"`
	Quantity   float64     `xml:"quantity,attr,omitempty"`
	LineNumber int         `xml:"lineNumber,attr,omitempty"`
	ItemDetail *ItemDetail `xml:"ItemDetail,omitempty"`
}

type ItemDetail struct {
	XMLName        xml.Name        `xml:"ItemDetail"`
	UnitPrice      *Money          `xml:"UnitPrice>Money,omitempty"`
	Description    *Description    `xml:"Description,omitempty"`
	UnitOfMeasure  string          `xml:"UnitOfMeasure,omitempty"`
	Classification *Classification `xml:"Classification,omitempty"`
}

type Description struct {
	XMLName   xml.Name `xml:"Description"`
	ShortName string   `xml:"xml:lang,attr,omitempty"`
	Value     string   `xml:",chardata"`
}

type Classification struct {
	XMLName xml.Name `xml:"Classification"`
	Domain  string   `xml:"domain,attr,omitempty"`
	Value   string   `xml:",chardata"`
}

type Money struct {
	XMLName  xml.Name `xml:"Money"`
	Currency string   `xml:"currency,attr,omitempty"`
	Amount   float64  `xml:",chardata"`
}
