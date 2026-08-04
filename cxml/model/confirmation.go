package model

import "encoding/xml"

// ConfirmationRequest confirms acceptance/rejection or detailed commit status for an order.
type ConfirmationRequest struct {
	XMLName                     xml.Name                     `xml:"ConfirmationRequest"`
	ConfirmationHeader          *ConfirmationHeader          `xml:"ConfirmationHeader,omitempty"`
	OrderReference              *OrderReference              `xml:"OrderReference,omitempty"`
	OrderStatusRequestReference *OrderStatusRequestReference `xml:"OrderStatusRequestReference,omitempty"`
	OrderStatusRequestIDInfo    *OrderStatusRequestIDInfo    `xml:"OrderStatusRequestIDInfo,omitempty"`
	ConfirmationItem            []*ConfirmationItem          `xml:"ConfirmationItem,omitempty"`
}

// ConfirmationHeader contains request-level confirmation metadata.
type ConfirmationHeader struct {
	XMLName           xml.Name           `xml:"ConfirmationHeader"`
	ConfirmID         string             `xml:"confirmID,attr,omitempty"`
	Operation         string             `xml:"operation,attr,omitempty"`
	Type              string             `xml:"type,attr,omitempty"`
	NoticeDate        string             `xml:"noticeDate,attr,omitempty"`
	InvoiceID         string             `xml:"invoiceID,attr,omitempty"`
	IncoTerms         string             `xml:"incoTerms,attr,omitempty"`
	Version           string             `xml:"version,attr,omitempty"`
	DocumentReference *DocumentReference `xml:"DocumentReference,omitempty"`
	Total             *Total             `xml:"Total,omitempty"`
	Shipping          *Shipping          `xml:"Shipping,omitempty"`
	Tax               *Tax               `xml:"Tax,omitempty"`
	Contact           []*Contact         `xml:"Contact,omitempty"`
	Hazard            []*Hazard          `xml:"Hazard,omitempty"`
	Comments          []*Comments        `xml:"Comments,omitempty"`
	IdReference       []*IdReference     `xml:"IdReference,omitempty"`
	Extrinsic         []*Extrinsic       `xml:"Extrinsic,omitempty"`
}

// ConfirmationItem is the line-level item confirmation.
type ConfirmationItem struct {
	XMLName            xml.Name              `xml:"ConfirmationItem"`
	Quantity           string                `xml:"quantity,attr,omitempty"`
	LineNumber         string                `xml:"lineNumber,attr,omitempty"`
	ParentLineNumber   string                `xml:"parentLineNumber,attr,omitempty"`
	ItemType           string                `xml:"itemType,attr,omitempty"`
	CompositeItemType  string                `xml:"compositeItemType,attr,omitempty"`
	UnitOfMeasure      *UnitOfMeasure        `xml:"UnitOfMeasure,omitempty"`
	Contact            []*Contact            `xml:"Contact,omitempty"`
	Hazard             []*Hazard             `xml:"Hazard,omitempty"`
	ConfirmationStatus []*ConfirmationStatus `xml:"ConfirmationStatus,omitempty"`
}

// ConfirmationStatus contains status details for a confirmation line.
type ConfirmationStatus struct {
	XMLName                     xml.Name                      `xml:"ConfirmationStatus"`
	Quantity                    string                        `xml:"quantity,attr,omitempty"`
	Type                        string                        `xml:"type,attr,omitempty"`
	ShipmentDate                string                        `xml:"shipmentDate,attr,omitempty"`
	DeliveryDate                string                        `xml:"deliveryDate,attr,omitempty"`
	UnitOfMeasure               *UnitOfMeasure                `xml:"UnitOfMeasure,omitempty"`
	ItemIn                      *ItemIn                       `xml:"ItemIn,omitempty"`
	UnitPrice                   *UnitPrice                    `xml:"UnitPrice,omitempty"`
	Tax                         *Tax                          `xml:"Tax,omitempty"`
	Shipping                    *Shipping                     `xml:"Shipping,omitempty"`
	SupplierBatchID             *SupplierBatchID              `xml:"SupplierBatchID,omitempty"`
	ScheduleLineReference       *ScheduleLineReference        `xml:"ScheduleLineReference,omitempty"`
	ComponentConsumptionDetails []*ComponentConsumptionDetail `xml:"ComponentConsumptionDetails,omitempty"`
	Comments                    []*Comments                   `xml:"Comments,omitempty"`
	Extrinsic                   []*Extrinsic                  `xml:"Extrinsic,omitempty"`
}

// OrderReference identifies the order being referenced.
type OrderReference struct {
	XMLName           xml.Name           `xml:"OrderReference"`
	OrderID           string             `xml:"orderID,attr,omitempty"`
	OrderDate         string             `xml:"orderDate,attr,omitempty"`
	DocumentReference *DocumentReference `xml:"DocumentReference,omitempty"`
}

type OrderStatusRequestReference struct {
	XMLName           xml.Name           `xml:"OrderStatusRequestReference"`
	PayloadID         string             `xml:"payloadID,attr,omitempty"`
	DocumentReference *DocumentReference `xml:"DocumentReference,omitempty"`
}

type OrderStatusRequestIDInfo struct {
	XMLName xml.Name `xml:"OrderStatusRequestIDInfo"`
	ID      string   `xml:"orderStatusRequestID,attr,omitempty"`
}

type ScheduleLineReference struct {
	XMLName               xml.Name     `xml:"ScheduleLineReference"`
	Quantity              string       `xml:"quantity,attr,omitempty"`
	RequestedDeliveryDate string       `xml:"requestedDeliveryDate,attr,omitempty"`
	LineNumber            string       `xml:"lineNumber,attr,omitempty"`
	Extrinsic             []*Extrinsic `xml:"Extrinsic,omitempty"`
}

type ComponentConsumptionDetail struct {
	XMLName xml.Name `xml:"ComponentConsumptionDetails"`
	Content string   `xml:",innerxml"`
}
