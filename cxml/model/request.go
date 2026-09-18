package model

import "encoding/xml"

type Request struct {
	XMLName                xml.Name                `xml:"Request"`
	DeploymentMode         string                  `xml:"deploymentMode,attr,omitempty"`
	OrderRequest           *OrderRequest           `xml:"OrderRequest,omitempty"`
	OrderChangeRequest     *OrderChangeRequest     `xml:"OrderChangeRequest,omitempty"`
	ConfirmationRequest    *ConfirmationRequest    `xml:"ConfirmationRequest,omitempty"`
	ProfileRequest         *ProfileRequest         `xml:"ProfileRequest,omitempty"`
	StatusUpdateRequest    *StatusUpdateRequest    `xml:"StatusUpdateRequest,omitempty"`
	PunchOutSetupRequest   *PunchOutSetupRequest   `xml:"PunchOutSetupRequest,omitempty"`
	PunchOutOrderMessage   *PunchOutOrderMessage   `xml:"PunchOutOrderMessage,omitempty"`
	ReceivingAdviceRequest *ReceivingAdviceRequest `xml:"ReceivingAdviceRequest,omitempty"`
	ShipNoticeRequest      *ShipNoticeRequest      `xml:"ShipNoticeRequest,omitempty"`
	InvoiceDetailRequest   *InvoiceDetailRequest   `xml:"InvoiceDetailRequest,omitempty"`
}

func (r *Request) PayloadType() string {
	switch {
	case r.OrderRequest != nil:
		return "OrderRequest"
	case r.OrderChangeRequest != nil:
		return "OrderChangeRequest"
	case r.ConfirmationRequest != nil:
		return "ConfirmationRequest"
	case r.ProfileRequest != nil:
		return "ProfileRequest"
	case r.StatusUpdateRequest != nil:
		return "StatusUpdateRequest"
	case r.PunchOutSetupRequest != nil:
		return "PunchOutSetupRequest"
	case r.PunchOutOrderMessage != nil:
		return "PunchOutOrderMessage"
	case r.ReceivingAdviceRequest != nil:
		return "ReceivingAdviceRequest"
	case r.ShipNoticeRequest != nil:
		return "ShipNoticeRequest"
	case r.InvoiceDetailRequest != nil:
		return "InvoiceDetailRequest"
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
	XMLName                  xml.Name `xml:"OrderRequestHeader"`
	OrderID                  string   `xml:"orderID,attr,omitempty"`
	OrderDate                string   `xml:"orderDate,attr,omitempty"`
	OrderType                string   `xml:"orderType,attr,omitempty"`
	Type                     string   `xml:"type,attr,omitempty"`
	OrderVersion             string   `xml:"orderVersion,attr,omitempty"`
	RequisitionID            string   `xml:"requisitionID,attr,omitempty"`
	ShipComplete             string   `xml:"shipComplete,attr,omitempty"`
	PickUpDate               string   `xml:"pickUpDate,attr,omitempty"`
	RequestedDeliveryDate    string   `xml:"requestedDeliveryDate,attr,omitempty"`
	EffectiveDate            string   `xml:"effectiveDate,attr,omitempty"`
	ExpirationDate           string   `xml:"expirationDate,attr,omitempty"`
	AgreementID              string   `xml:"agreementID,attr,omitempty"`
	AgreementPayloadID       string   `xml:"agreementPayloadID,attr,omitempty"`
	ParentAgreementID        string   `xml:"parentAgreementID,attr,omitempty"`
	ParentAgreementPayloadID string   `xml:"parentAgreementPayloadID,attr,omitempty"`
	IsInternalVersion        string   `xml:"isInternalVersion,attr,omitempty"`
	IsSTOOutbound            string   `xml:"isSTOOutbound,attr,omitempty"`

	Total               *Money                      `xml:"Total>Money,omitempty"`
	ShipTo              *Party                      `xml:"ShipTo,omitempty"`
	BillTo              *Party                      `xml:"BillTo,omitempty"`
	BusinessPartner     []*BusinessPartner          `xml:"BusinessPartner,omitempty"`
	LegalEntity         *LegalEntity                `xml:"LegalEntity,omitempty"`
	OrganizationalUnit  []*OrganizationalUnit       `xml:"OrganizationalUnit,omitempty"`
	Shipping            *Shipping                   `xml:"Shipping,omitempty"`
	Tax                 *Tax                        `xml:"Tax,omitempty"`
	Payment             *Payment                    `xml:"Payment,omitempty"`
	PaymentTerm         []*PaymentTerm              `xml:"PaymentTerm,omitempty"`
	Contact             []*Contact                  `xml:"Contact,omitempty"`
	Comments            []*Comments                 `xml:"Comments,omitempty"`
	Followup            *Followup                   `xml:"Followup,omitempty"`
	ControlKeys         *ControlKeys                `xml:"ControlKeys,omitempty"`
	DocumentReference   *DocumentReference          `xml:"DocumentReference,omitempty"`
	SupplierOrderInfo   *SupplierOrderInfo          `xml:"SupplierOrderInfo,omitempty"`
	TermsOfDelivery     *TermsOfDelivery            `xml:"TermsOfDelivery,omitempty"`
	DeliveryPeriod      *DeliveryPeriod             `xml:"DeliveryPeriod,omitempty"`
	IdReference         []*IdReference              `xml:"IdReference,omitempty"`
	OrderHeaderIndustry *OrderRequestHeaderIndustry `xml:"OrderRequestHeaderIndustry,omitempty"`
	Extrinsic           []*Extrinsic                `xml:"Extrinsic,omitempty"`
}

type ItemOut struct {
	XMLName                xml.Name `xml:"ItemOut"`
	Quantity               float64  `xml:"quantity,attr,omitempty"`
	LineNumber             int      `xml:"lineNumber,attr,omitempty"`
	RequisitionID          string   `xml:"requisitionID,attr,omitempty"`
	RequestedDeliveryDate  string   `xml:"requestedDeliveryDate,attr,omitempty"`
	RequestedShipmentDate  string   `xml:"requestedShipmentDate,attr,omitempty"`
	IsAdHoc                string   `xml:"isAdHoc,attr,omitempty"`
	ParentLineNumber       string   `xml:"parentLineNumber,attr,omitempty"`
	ItemType               string   `xml:"itemType,attr,omitempty"`
	CompositeItemType      string   `xml:"compositeItemType,attr,omitempty"`
	ItemClassification     string   `xml:"itemClassification,attr,omitempty"`
	ItemCategory           string   `xml:"itemCategory,attr,omitempty"`
	IsReturn               string   `xml:"isReturn,attr,omitempty"`
	ReturnAuthorizationNum string   `xml:"returnAuthorizationNumber,attr,omitempty"`

	ItemID            *ItemID            `xml:"ItemID,omitempty"`
	ItemDetail        *ItemDetail        `xml:"ItemDetail,omitempty"`
	BlanketItemDetail *BlanketItemDetail `xml:"BlanketItemDetail,omitempty"`
	ShipTo            *ShipTo            `xml:"ShipTo,omitempty"`
	Shipping          *Shipping          `xml:"Shipping,omitempty"`
	Tax               *Tax               `xml:"Tax,omitempty"`
	SpendDetail       *SpendDetail       `xml:"SpendDetail,omitempty"`
	Distribution      []*Distribution    `xml:"Distribution,omitempty"`
	Contact           []*Contact         `xml:"Contact,omitempty"`
	TermsOfDelivery   *TermsOfDelivery   `xml:"TermsOfDelivery,omitempty"`
	Comments          []*Comments        `xml:"Comments,omitempty"`
	Tolerances        *Tolerances        `xml:"Tolerances,omitempty"`
	ControlKeys       *ControlKeys       `xml:"ControlKeys,omitempty"`
	ScheduleLine      []*ScheduleLine    `xml:"ScheduleLine,omitempty"`
	ItemOutIndustry   *ItemOutIndustry   `xml:"ItemOutIndustry,omitempty"`
	Packaging         []*Packaging       `xml:"Packaging,omitempty"`
	ReleaseInfo       *ReleaseInfo       `xml:"ReleaseInfo,omitempty"`
	Batch             *Batch             `xml:"Batch,omitempty"`
}

type ItemDetail struct {
	XMLName            xml.Name            `xml:"ItemDetail"`
	UnitPrice          *Money              `xml:"UnitPrice>Money,omitempty"`
	Description        *Description        `xml:"Description,omitempty"`
	UnitOfMeasure      string              `xml:"UnitOfMeasure,omitempty"`
	Classification     *Classification     `xml:"Classification,omitempty"`
	Classifications    []*Classification   `xml:"-"`
	ManufacturerPartID *ManufacturerPartID `xml:"ManufacturerPartID,omitempty"`
	ManufacturerName   *ManufacturerName   `xml:"ManufacturerName,omitempty"`
	URL                *URL                `xml:"URL,omitempty"`
	LeadTime           string              `xml:"LeadTime,omitempty"`
	Dimension          []*Dimension        `xml:"Dimension,omitempty"`
	PriceBasisQuantity *PriceBasisQuantity `xml:"PriceBasisQuantity,omitempty"`
	Extrinsic          []*Extrinsic        `xml:"Extrinsic,omitempty"`
}

// ItemID identifies an ordered product.
type ItemID struct {
	XMLName                 xml.Name       `xml:"ItemID"`
	SupplierPartID          string         `xml:"SupplierPartID,omitempty"`
	SupplierPartRevision    string         `xml:"-"`
	SupplierPartAuxiliaryID string         `xml:"SupplierPartAuxiliaryID,omitempty"`
	BuyerPartID             string         `xml:"BuyerPartID,omitempty"`
	IdReference             []*IdReference `xml:"IdReference,omitempty"`
}

func (i *ItemID) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type supplierPartID struct {
		Value      string `xml:",chardata"`
		RevisionID string `xml:"revisionID,attr"`
	}
	type itemID struct {
		XMLName                 xml.Name       `xml:"ItemID"`
		SupplierPartID          supplierPartID `xml:"SupplierPartID"`
		SupplierPartAuxiliaryID string         `xml:"SupplierPartAuxiliaryID"`
		BuyerPartID             string         `xml:"BuyerPartID"`
		IdReference             []*IdReference `xml:"IdReference"`
	}

	var parsed itemID
	if err := d.DecodeElement(&parsed, &start); err != nil {
		return err
	}
	i.XMLName = parsed.XMLName
	i.SupplierPartID = parsed.SupplierPartID.Value
	i.SupplierPartRevision = parsed.SupplierPartID.RevisionID
	i.SupplierPartAuxiliaryID = parsed.SupplierPartAuxiliaryID
	i.BuyerPartID = parsed.BuyerPartID
	i.IdReference = parsed.IdReference
	return nil
}

func (i ItemID) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type supplierPartID struct {
		Value      string `xml:",chardata"`
		RevisionID string `xml:"revisionID,attr,omitempty"`
	}
	type itemID struct {
		XMLName                 xml.Name       `xml:"ItemID"`
		SupplierPartID          supplierPartID `xml:"SupplierPartID,omitempty"`
		SupplierPartAuxiliaryID string         `xml:"SupplierPartAuxiliaryID,omitempty"`
		BuyerPartID             string         `xml:"BuyerPartID,omitempty"`
		IdReference             []*IdReference `xml:"IdReference,omitempty"`
	}

	if start.Name.Local == "" {
		start.Name = xml.Name{Local: "ItemID"}
	}
	return e.EncodeElement(itemID{
		XMLName: start.Name,
		SupplierPartID: supplierPartID{
			Value:      i.SupplierPartID,
			RevisionID: i.SupplierPartRevision,
		},
		SupplierPartAuxiliaryID: i.SupplierPartAuxiliaryID,
		BuyerPartID:             i.BuyerPartID,
		IdReference:             i.IdReference,
	}, start)
}

// BusinessPartner associates an additional partner role with the order.
type BusinessPartner struct {
	XMLName     xml.Name       `xml:"BusinessPartner"`
	Type        string         `xml:"type,attr,omitempty"`
	Role        string         `xml:"role,attr,omitempty"`
	Address     *Address       `xml:"Address,omitempty"`
	IdReference []*IdReference `xml:"IdReference,omitempty"`
}

type LegalEntity struct {
	XMLName     xml.Name     `xml:"LegalEntity"`
	IdReference *IdReference `xml:"IdReference,omitempty"`
}

type OrganizationalUnit struct {
	XMLName     xml.Name     `xml:"OrganizationalUnit"`
	IdReference *IdReference `xml:"IdReference,omitempty"`
}

type SupplierOrderInfo struct {
	XMLName   xml.Name `xml:"SupplierOrderInfo"`
	OrderID   string   `xml:"orderID,attr,omitempty"`
	OrderDate string   `xml:"orderDate,attr,omitempty"`
}

type OrderRequestHeaderIndustry struct {
	XMLName               xml.Name                 `xml:"OrderRequestHeaderIndustry"`
	ReferenceDocumentInfo []*ReferenceDocumentInfo `xml:"ReferenceDocumentInfo,omitempty"`
	LifeSciences          *LifeSciences            `xml:"LifeSciences,omitempty"`
	AerospaceAndDefense   *AerospaceAndDefense     `xml:"AerospaceAndDefense,omitempty"`
}

type Followup struct {
	XMLName xml.Name `xml:"Followup"`
	URL     *URL     `xml:"URL,omitempty"`
}

type ScheduleLine struct {
	XMLName               xml.Name       `xml:"ScheduleLine"`
	Quantity              string         `xml:"quantity,attr,omitempty"`
	RequestedDeliveryDate string         `xml:"requestedDeliveryDate,attr,omitempty"`
	LineNumber            string         `xml:"lineNumber,attr,omitempty"`
	RequestedShipmentDate string         `xml:"requestedShipmentDate,attr,omitempty"`
	UnitOfMeasure         *UnitOfMeasure `xml:"UnitOfMeasure,omitempty"`
	ShipTo                *ShipTo        `xml:"ShipTo,omitempty"`
	Extrinsic             []*Extrinsic   `xml:"Extrinsic,omitempty"`
}

type BlanketItemDetail struct {
	XMLName            xml.Name            `xml:"BlanketItemDetail"`
	Description        []*Description      `xml:"Description,omitempty"`
	UnitPrice          *UnitPrice          `xml:"UnitPrice,omitempty"`
	UnitOfMeasure      *UnitOfMeasure      `xml:"UnitOfMeasure,omitempty"`
	PriceBasisQuantity *PriceBasisQuantity `xml:"PriceBasisQuantity,omitempty"`
	Classification     []*Classification   `xml:"Classification,omitempty"`
	Extrinsic          []*Extrinsic        `xml:"Extrinsic,omitempty"`
}

type ReleaseInfo struct {
	XMLName                    xml.Name       `xml:"ReleaseInfo"`
	ReleaseType                string         `xml:"releaseType,attr,omitempty"`
	CumulativeReceivedQuantity string         `xml:"cumulativeReceivedQuantity,attr,omitempty"`
	ReleaseNumber              string         `xml:"releaseNumber,attr,omitempty"`
	UnitOfMeasure              *UnitOfMeasure `xml:"UnitOfMeasure,omitempty"`
	Extrinsic                  []*Extrinsic   `xml:"Extrinsic,omitempty"`
}

type ItemOutIndustry struct {
	XMLName                    xml.Name                    `xml:"ItemOutIndustry"`
	ReferenceDocumentInfo      []*ReferenceDocumentInfo    `xml:"ReferenceDocumentInfo,omitempty"`
	ItemOutLifeSciences        *ItemOutLifeSciences        `xml:"ItemOutLifeSciences,omitempty"`
	ItemOutAerospaceAndDefense *ItemOutAerospaceAndDefense `xml:"ItemOutAerospaceAndDefense,omitempty"`
}

type ReferenceDocumentInfo struct {
	XMLName           xml.Name           `xml:"ReferenceDocumentInfo"`
	DocumentInfo      *DocumentInfo      `xml:"DocumentInfo,omitempty"`
	DocumentReference *DocumentReference `xml:"DocumentReference,omitempty"`
	DateInfo          []*DateInfo        `xml:"DateInfo,omitempty"`
	Contact           []*Contact         `xml:"Contact,omitempty"`
	Extrinsic         []*Extrinsic       `xml:"Extrinsic,omitempty"`
}

type DateInfo struct {
	XMLName xml.Name `xml:"DateInfo"`
	Type    string   `xml:"type,attr,omitempty"`
	Date    string   `xml:"date,attr,omitempty"`
}

type LifeSciences struct {
	XMLName       xml.Name `xml:"LifeSciences"`
	OrderCategory string   `xml:"orderCategory,attr,omitempty"`
}

type AerospaceAndDefense struct {
	XMLName xml.Name `xml:"AerospaceAndDefense"`
}

type ItemOutLifeSciences struct {
	XMLName xml.Name `xml:"ItemOutLifeSciences"`
}

type ItemOutAerospaceAndDefense struct {
	XMLName xml.Name `xml:"ItemOutAerospaceAndDefense"`
}

// ProfileRequest asks a supplier for supported transactions/options.
type ProfileRequest struct {
	XMLName xml.Name `xml:"ProfileRequest"`
}

// StatusUpdateRequest updates the status of a related document.
type StatusUpdateRequest struct {
	XMLName           xml.Name           `xml:"StatusUpdateRequest"`
	DocumentReference *DocumentReference `xml:"DocumentReference,omitempty"`
	Status            *Status            `xml:"Status,omitempty"`
	Extrinsic         []*Extrinsic       `xml:"Extrinsic,omitempty"`
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
