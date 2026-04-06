package model

import "encoding/xml"

type Request struct {
	XMLName                xml.Name                `xml:"Request"`
	DeploymentMode         string                  `xml:"deploymentMode,attr,omitempty"`
	OrderRequest           *OrderRequest           `xml:"OrderRequest,omitempty"`
	OrderChangeRequest     *OrderChangeRequest     `xml:"OrderChangeRequest,omitempty"`
	PunchOutOrderMessage   *PunchOutOrderMessage   `xml:"PunchOutOrderMessage,omitempty"`
	ReceivingAdviceRequest *ReceivingAdviceRequest `xml:"ReceivingAdviceRequest,omitempty"`
}

func (r *Request) PayloadType() string {
	switch {
	case r.OrderRequest != nil:
		return "OrderRequest"
	case r.OrderChangeRequest != nil:
		return "OrderChangeRequest"
	case r.PunchOutOrderMessage != nil:
		return "PunchOutOrderMessage"
	case r.ReceivingAdviceRequest != nil:
		return "ReceivingAdviceRequest"
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
	XMLName               xml.Name            `xml:"OrderChangeRequest"`
	OrderRequestReference *OrderRequestHeader `xml:"OrderRequestReference,omitempty"`
	ItemChange            []*ItemOut          `xml:"ItemChange,omitempty"`
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

// ReceivingAdviceRequest represents an inbound EDI 861 (RECADV) document,
// sent by the buyer to confirm goods have been received.
type ReceivingAdviceRequest struct {
	XMLName xml.Name               `xml:"ReceivingAdviceRequest"`
	Header  *ReceivingAdviceHeader `xml:"ReceivingAdviceHeader,omitempty"`
	Orders  []*ReceivingOrder      `xml:"ReceivingOrder,omitempty"`
}

type ReceivingAdviceHeader struct {
	XMLName          xml.Name          `xml:"ReceivingAdviceHeader"`
	ID               string            `xml:"receivingAdviceID,attr,omitempty"`
	Date             string            `xml:"receivingAdviceDate,attr,omitempty"`
	ShipNoticeIDInfo *ShipNoticeIDInfo `xml:"ShipNoticeIDInfo,omitempty"`
}

type ShipNoticeIDInfo struct {
	XMLName xml.Name `xml:"ShipNoticeIDInfo"`
	ID      string   `xml:"shipNoticeID,attr,omitempty"`
	Date    string   `xml:"shipNoticeDate,attr,omitempty"`
}

type ReceivingOrder struct {
	XMLName  xml.Name                 `xml:"ReceivingOrder"`
	OrderRef *ReceivingOrderReference `xml:"OrderReference,omitempty"`
	Details  []*ReceivingOrderDetail  `xml:"ReceivingOrderDetail,omitempty"`
}

type ReceivingOrderReference struct {
	XMLName   xml.Name `xml:"OrderReference"`
	OrderID   string   `xml:"orderID,attr,omitempty"`
	OrderDate string   `xml:"orderDate,attr,omitempty"`
}

type ReceivingOrderDetail struct {
	XMLName       xml.Name            `xml:"ReceivingOrderDetail"`
	LineNumber    int                 `xml:"lineNumber,attr,omitempty"`
	Quantity      float64             `xml:"quantity,attr,omitempty"`
	ReceivingDate string              `xml:"receivingDate,attr,omitempty"`
	UnitOfMeasure string              `xml:"UnitOfMeasure,omitempty"`
	Condition     *ReceivingCondition `xml:"ReceivingCondition,omitempty"`
	ItemOut       *ReceivingItemOut   `xml:"ItemOut,omitempty"`
}

type ReceivingCondition struct {
	XMLName xml.Name `xml:"ReceivingCondition"`
	Code    string   `xml:"receivingConditionCode,attr,omitempty"`
}

type ReceivingItemOut struct {
	XMLName        xml.Name `xml:"ItemOut"`
	SupplierPartID string   `xml:"ItemID>SupplierPartID,omitempty"`
	BuyerPartID    string   `xml:"ItemID>BuyerPartID,omitempty"`
}
