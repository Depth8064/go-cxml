package model

import "encoding/xml"

// ShipNoticeRequest is the Advance Ship Notice (ASN) document,
// sent by the supplier to notify the buyer of a pending shipment.
// DTD: Fulfill.dtd 1.2.069 — <!ELEMENT ShipNoticeRequest>
type ShipNoticeRequest struct {
	XMLName           xml.Name             `xml:"ShipNoticeRequest"`
	ShipNoticeHeader  *ShipNoticeHeader    `xml:"ShipNoticeHeader"`
	ShipControl       []*ShipControl       `xml:"ShipControl,omitempty"`
	ShipNoticePortion []*ShipNoticePortion `xml:"ShipNoticePortion,omitempty"`
}

// ShipNoticeHeader contains shipment-level metadata common to all items.
type ShipNoticeHeader struct {
	XMLName               xml.Name                 `xml:"ShipNoticeHeader"`
	ShipmentID            string                   `xml:"shipmentID,attr"`          // REQUIRED
	Operation             string                   `xml:"operation,attr,omitempty"` // new|update|delete
	NoticeDate            string                   `xml:"noticeDate,attr"`          // REQUIRED
	ShipmentDate          string                   `xml:"shipmentDate,attr,omitempty"`
	DeliveryDate          string                   `xml:"deliveryDate,attr,omitempty"`
	ShipmentType          string                   `xml:"shipmentType,attr,omitempty"`    // actual|planned
	FulfillmentType       string                   `xml:"fulfillmentType,attr,omitempty"` // partial|complete
	RequestedDeliveryDate string                   `xml:"requestedDeliveryDate,attr,omitempty"`
	Reason                string                   `xml:"reason,attr,omitempty"`           // return
	ActivityStepType      string                   `xml:"activityStepType,attr,omitempty"` // stockTransfer|stockShippingAdvice
	ServiceLevel          []*ServiceLevel          `xml:"ServiceLevel,omitempty"`
	DocumentReference     *DocumentReference       `xml:"DocumentReference,omitempty"`
	Contact               []*Contact               `xml:"Contact,omitempty"`
	LegalEntity           *LegalEntity             `xml:"LegalEntity,omitempty"`
	OrganizationalUnit    []*OrganizationalUnit    `xml:"OrganizationalUnit,omitempty"`
	Hazard                []*Hazard                `xml:"Hazard,omitempty"`
	Comments              []*Comments              `xml:"Comments,omitempty"`
	TermsOfDelivery       *TermsOfDelivery         `xml:"TermsOfDelivery,omitempty"`
	TermsOfTransport      []*TermsOfTransport      `xml:"TermsOfTransport,omitempty"`
	Packaging             *Packaging               `xml:"Packaging,omitempty"`
	Extrinsic             []*Extrinsic             `xml:"Extrinsic,omitempty"`
	IdReference           []*IdReference           `xml:"IdReference,omitempty"`
	ReferenceDocumentInfo []*ReferenceDocumentInfo `xml:"ReferenceDocumentInfo,omitempty"`
}

// ShipControl identifies the carrier responsible for a portion of the shipment.
type ShipControl struct {
	XMLName               xml.Name                `xml:"ShipControl"`
	StartDate             string                  `xml:"startDate,attr,omitempty"`
	CarrierIdentifier     []*CarrierIdentifier    `xml:"CarrierIdentifier"`
	ShipmentIdentifier    []*ShipmentIdentifier   `xml:"ShipmentIdentifier"`
	PackageIdentification *PackageIdentification  `xml:"PackageIdentification,omitempty"`
	Route                 []*Route                `xml:"Route,omitempty"`
	TransportInformation  []*TransportInformation `xml:"TransportInformation,omitempty"`
	Contact               []*Contact              `xml:"Contact,omitempty"`
	Comments              []*Comments             `xml:"Comments,omitempty"`
	Extrinsic             []*Extrinsic            `xml:"Extrinsic,omitempty"`
}

// ShipmentIdentifier is a tracking number or bill of lading reference.
type ShipmentIdentifier struct {
	XMLName            xml.Name `xml:"ShipmentIdentifier"`
	Domain             string   `xml:"domain,attr,omitempty"` // e.g. "trackingNumberID", "billOfLadingID"
	TrackingNumberDate string   `xml:"trackingNumberDate,attr,omitempty"`
	TrackingURL        string   `xml:"trackingURL,attr,omitempty"`
	Value              string   `xml:",chardata"`
}

// PackageIdentification provides inclusive range of package numbers on the containers.
type PackageIdentification struct {
	XMLName    xml.Name `xml:"PackageIdentification"`
	RangeBegin string   `xml:"rangeBegin,attr"` // REQUIRED
	RangeEnd   string   `xml:"rangeEnd,attr"`   // REQUIRED
}

// ShipNoticePortion links a shipment to a specific purchase order.
type ShipNoticePortion struct {
	XMLName                  xml.Name                  `xml:"ShipNoticePortion"`
	OrderReference           *OrderReference           `xml:"OrderReference"`
	MasterAgreementReference *MasterAgreementReference `xml:"MasterAgreementReference,omitempty"`
	MasterAgreementIDInfo    *MasterAgreementIDInfo    `xml:"MasterAgreementIDInfo,omitempty"`
	Contact                  []*Contact                `xml:"Contact,omitempty"`
	Comments                 []*Comments               `xml:"Comments,omitempty"`
	Extrinsic                []*Extrinsic              `xml:"Extrinsic,omitempty"`
	ShipNoticeItem           []*ShipNoticeItem         `xml:"ShipNoticeItem,omitempty"`
	ReferenceDocumentInfo    []*ReferenceDocumentInfo  `xml:"ReferenceDocumentInfo,omitempty"`
}

// ShipNoticeItem describes one line item included in the shipment.
type ShipNoticeItem struct {
	XMLName                     xml.Name                      `xml:"ShipNoticeItem"`
	Quantity                    string                        `xml:"quantity,attr"`   // REQUIRED
	LineNumber                  string                        `xml:"lineNumber,attr"` // REQUIRED
	ParentLineNumber            string                        `xml:"parentLineNumber,attr,omitempty"`
	ShipNoticeLineNumber        string                        `xml:"shipNoticeLineNumber,attr,omitempty"`
	ItemType                    string                        `xml:"itemType,attr,omitempty"`
	CompositeItemType           string                        `xml:"compositeItemType,attr,omitempty"`
	StockTransferType           string                        `xml:"stockTransferType,attr,omitempty"` // intra|inter
	OutboundType                string                        `xml:"outboundType,attr,omitempty"`      // stockTransport
	ItemID                      *ItemID                       `xml:"ItemID,omitempty"`
	ShipNoticeItemDetail        *ShipNoticeItemDetail         `xml:"ShipNoticeItemDetail,omitempty"`
	UnitOfMeasure               *UnitOfMeasure                `xml:"UnitOfMeasure"`
	Packaging                   []*Packaging                  `xml:"Packaging,omitempty"`
	Hazard                      []*Hazard                     `xml:"Hazard,omitempty"`
	Batch                       []*Batch                      `xml:"Batch,omitempty"`
	SupplierBatchID             []*SupplierBatchID            `xml:"SupplierBatchID,omitempty"`
	AssetInfo                   []*AssetInfo                  `xml:"AssetInfo,omitempty"`
	TermsOfDelivery             *TermsOfDelivery              `xml:"TermsOfDelivery,omitempty"`
	OrderedQuantity             *OrderedQuantity              `xml:"OrderedQuantity,omitempty"`
	ShipNoticeItemIndustry      *ShipNoticeItemIndustry       `xml:"ShipNoticeItemIndustry,omitempty"`
	ComponentConsumptionDetails []*ComponentConsumptionDetail `xml:"ComponentConsumptionDetails,omitempty"`
	ReferenceDocumentInfo       []*ReferenceDocumentInfo      `xml:"ReferenceDocumentInfo,omitempty"`
	Comments                    []*Comments                   `xml:"Comments,omitempty"`
	Extrinsic                   []*Extrinsic                  `xml:"Extrinsic,omitempty"`
}

// ShipNoticeItemDetail holds additional descriptive data about a shipped item.
type ShipNoticeItemDetail struct {
	XMLName            xml.Name            `xml:"ShipNoticeItemDetail"`
	UnitPrice          *UnitPrice          `xml:"UnitPrice,omitempty"`
	Description        []*Description      `xml:"Description,omitempty"`
	UnitOfMeasure      *UnitOfMeasure      `xml:"UnitOfMeasure,omitempty"`
	PriceBasisQuantity *PriceBasisQuantity `xml:"PriceBasisQuantity,omitempty"`
	Classification     []*Classification   `xml:"Classification,omitempty"`
	ManufacturerPartID *ManufacturerPartID `xml:"ManufacturerPartID,omitempty"`
	ManufacturerName   *ManufacturerName   `xml:"ManufacturerName,omitempty"`
	Dimension          []*Dimension        `xml:"Dimension,omitempty"`
	ItemDetailIndustry *ItemDetailIndustry `xml:"ItemDetailIndustry,omitempty"`
	Extrinsic          []*Extrinsic        `xml:"Extrinsic,omitempty"`
}

// OrderedQuantity is the originally ordered amount, for comparison.
type OrderedQuantity struct {
	XMLName       xml.Name       `xml:"OrderedQuantity"`
	Quantity      string         `xml:"quantity,attr"` // REQUIRED
	UnitOfMeasure *UnitOfMeasure `xml:"UnitOfMeasure,omitempty"`
}

// MasterAgreementReference points to a master/blanket agreement document.
type MasterAgreementReference struct {
	XMLName           xml.Name           `xml:"MasterAgreementReference"`
	AgreementID       string             `xml:"agreementID,attr,omitempty"`
	AgreementDate     string             `xml:"agreementDate,attr,omitempty"`
	AgreementType     string             `xml:"agreementType,attr,omitempty"`
	DocumentReference *DocumentReference `xml:"DocumentReference,omitempty"`
}

// MasterAgreementIDInfo identifies a master agreement by its buyer-system ID.
type MasterAgreementIDInfo struct {
	XMLName       xml.Name       `xml:"MasterAgreementIDInfo"`
	AgreementID   string         `xml:"agreementID,attr,omitempty"`
	AgreementDate string         `xml:"agreementDate,attr,omitempty"`
	AgreementType string         `xml:"agreementType,attr,omitempty"`
	IdReference   []*IdReference `xml:"IdReference,omitempty"`
}

// ShipNoticeItemIndustry groups industry-specific ship notice item fields.
type ShipNoticeItemIndustry struct {
	XMLName                    xml.Name                    `xml:"ShipNoticeItemIndustry"`
	ShipNoticeItemRetail       *ShipNoticeItemRetail       `xml:"ShipNoticeItemRetail,omitempty"`
	ShipNoticeItemLifeSciences *ShipNoticeItemLifeSciences `xml:"ShipNoticeItemLifeSciences,omitempty"`
}

// ShipNoticeItemRetail contains retail-specific ship notice item fields.
type ShipNoticeItemRetail struct {
	XMLName            xml.Name        `xml:"ShipNoticeItemRetail"`
	BestBeforeDate     *BestBeforeDate `xml:"BestBeforeDate,omitempty"`
	ExpiryDate         string          `xml:"ExpiryDate,omitempty"`
	FreeGoodsQuantity  string          `xml:"FreeGoodsQuantity,omitempty"`
	PromotionDealID    string          `xml:"PromotionDealID,omitempty"`
	PromotionVariantID string          `xml:"PromotionVariantID,omitempty"`
}

// ShipNoticeItemLifeSciences contains life-sciences-specific item fields.
type ShipNoticeItemLifeSciences struct {
	XMLName            xml.Name              `xml:"ShipNoticeItemLifeSciences"`
	Study              []*Study              `xml:"Study,omitempty"`
	ProtocolID         *ProtocolID           `xml:"ProtocolID,omitempty"`
	PoolID             *PoolID               `xml:"PoolID,omitempty"`
	MedicationListInfo []*MedicationListInfo `xml:"MedicationListInfo,omitempty"`
	Extrinsic          []*Extrinsic          `xml:"Extrinsic,omitempty"`
}

// ItemDetailIndustry groups industry-specific item detail extensions.
type ItemDetailIndustry struct {
	XMLName                xml.Name          `xml:"ItemDetailIndustry"`
	IsConfigurableMaterial string            `xml:"isConfigurableMaterial,attr,omitempty"`
	ItemDetailRetail       *ItemDetailRetail `xml:"ItemDetailRetail,omitempty"`
}

// Study represents a clinical study reference.
type Study struct {
	XMLName     xml.Name     `xml:"Study"`
	StudyID     *StudyID     `xml:"StudyID"`
	Description *Description `xml:"Description,omitempty"`
}

// StudyID is the identifier for a clinical study.
type StudyID struct {
	XMLName xml.Name `xml:"StudyID"`
	Value   string   `xml:",chardata"`
}

// ProtocolID is a clinical protocol identifier.
type ProtocolID struct {
	XMLName xml.Name `xml:"ProtocolID"`
	Value   string   `xml:",chardata"`
}

// PoolID is a pooling identifier for life-sciences items.
type PoolID struct {
	XMLName xml.Name `xml:"PoolID"`
	Value   string   `xml:",chardata"`
}

// MedicationListInfo holds a list of named property values for medication.
type MedicationListInfo struct {
	XMLName       xml.Name         `xml:"MedicationListInfo"`
	Type          string           `xml:"type,attr,omitempty"`
	PropertyValue []*PropertyValue `xml:"PropertyValue,omitempty"`
}

// PropertyValue is a named property with optional characteristic children.
type PropertyValue struct {
	XMLName        xml.Name          `xml:"PropertyValue"`
	Name           string            `xml:"name,attr,omitempty"`
	Characteristic []*Characteristic `xml:"Characteristic,omitempty"`
}

// ShipNoticeReference points back to an earlier ShipNoticeRequest.
type ShipNoticeReference struct {
	XMLName           xml.Name           `xml:"ShipNoticeReference"`
	ShipNoticeID      string             `xml:"shipNoticeID,attr,omitempty"`
	ShipNoticeDate    string             `xml:"shipNoticeDate,attr,omitempty"`
	DocumentReference *DocumentReference `xml:"DocumentReference"`
}

// ShipNoticeLineItemReference cross-references a line in a prior ShipNoticeRequest.
type ShipNoticeLineItemReference struct {
	XMLName              xml.Name `xml:"ShipNoticeLineItemReference"`
	ShipNoticeLineNumber string   `xml:"shipNoticeLineNumber,attr"` // REQUIRED
}
