package model

import "encoding/xml"

// ─── Terms of delivery ────────────────────────────────────────────────────────

// TermsOfDelivery specifies the delivery terms (Incoterms).
type TermsOfDelivery struct {
	XMLName               xml.Name               `xml:"TermsOfDelivery"`
	TermsOfDeliveryCode   *TermsOfDeliveryCode   `xml:"TermsOfDeliveryCode"`
	ShippingPaymentMethod *ShippingPaymentMethod `xml:"ShippingPaymentMethod"`
	TransportTerms        *TransportTerms        `xml:"TransportTerms,omitempty"`
	Address               *Address               `xml:"Address,omitempty"`
	Comments              []*Comments            `xml:"Comments,omitempty"`
}

// TermsOfDeliveryCode is the Incoterms code (e.g. "EXW", "DDP").
type TermsOfDeliveryCode struct {
	XMLName xml.Name `xml:"TermsOfDeliveryCode"`
	Value   string   `xml:"value,attr,omitempty"`
	Text    string   `xml:",chardata"`
}

// ShippingPaymentMethod specifies who pays for shipping.
type ShippingPaymentMethod struct {
	XMLName xml.Name `xml:"ShippingPaymentMethod"`
	Value   string   `xml:"value,attr,omitempty"` // e.g. "prepaid", "collect"
	Text    string   `xml:",chardata"`
}

// TransportTerms holds additional transport arrangement details.
type TransportTerms struct {
	XMLName xml.Name `xml:"TransportTerms"`
	Value   string   `xml:"value,attr,omitempty"`
	Text    string   `xml:",chardata"`
}

// ─── Terms of transport ───────────────────────────────────────────────────────

// TermsOfTransport describes equipment and sealing for transport.
type TermsOfTransport struct {
	XMLName                     xml.Name                     `xml:"TermsOfTransport"`
	SealID                      *SealID                      `xml:"SealID,omitempty"`
	SealingPartyCode            *SealingPartyCode            `xml:"SealingPartyCode,omitempty"`
	EquipmentIdentificationCode *EquipmentIdentificationCode `xml:"EquipmentIdentificationCode,omitempty"`
	TransportTerms              *TransportTerms              `xml:"TransportTerms,omitempty"`
	Dimension                   []*Dimension                 `xml:"Dimension,omitempty"`
	Extrinsic                   []*Extrinsic                 `xml:"Extrinsic,omitempty"`
}

// SealID is a container seal identifier.
type SealID struct {
	XMLName xml.Name `xml:"SealID"`
	Value   string   `xml:",chardata"`
}

// SealingPartyCode identifies the party responsible for sealing.
type SealingPartyCode struct {
	XMLName xml.Name `xml:"SealingPartyCode"`
	Value   string   `xml:",chardata"`
}

// EquipmentIdentificationCode identifies the transport equipment.
type EquipmentIdentificationCode struct {
	XMLName xml.Name `xml:"EquipmentIdentificationCode"`
	Value   string   `xml:",chardata"`
}

// ─── Dimensions ───────────────────────────────────────────────────────────────

// Dimension holds a measurement value with a unit of measure.
type Dimension struct {
	XMLName       xml.Name       `xml:"Dimension"`
	Type          string         `xml:"type,attr,omitempty"` // e.g. "weight", "volume", "length"
	Quantity      string         `xml:"quantity,attr,omitempty"`
	UnitOfMeasure *UnitOfMeasure `xml:"UnitOfMeasure,omitempty"`
}

// ─── Delivery timing ──────────────────────────────────────────────────────────

// DeliveryPeriod wraps a Period to specify a delivery window.
type DeliveryPeriod struct {
	XMLName xml.Name `xml:"DeliveryPeriod"`
	Period  *Period  `xml:"Period"`
}

// ─── Control keys ─────────────────────────────────────────────────────────────

// ControlKeys specifies operational controls for order processing.
type ControlKeys struct {
	XMLName            xml.Name            `xml:"ControlKeys"`
	OCInstruction      *OCInstruction      `xml:"OCInstruction,omitempty"`
	ASNInstruction     *ASNInstruction     `xml:"ASNInstruction,omitempty"`
	InvoiceInstruction *InvoiceInstruction `xml:"InvoiceInstruction,omitempty"`
	SESInstruction     *SESInstruction     `xml:"SESInstruction,omitempty"`
}

// OCInstruction controls whether order changes are allowed.
type OCInstruction struct {
	XMLName xml.Name `xml:"OCInstruction"`
	Value   string   `xml:"value,attr"` // (allowed|notAllowed|requiredBeforeASN) REQUIRED
	Lower   *Lower   `xml:"Lower,omitempty"`
	Upper   *Upper   `xml:"Upper,omitempty"`
}

// ASNInstruction controls advance ship notice requirements.
type ASNInstruction struct {
	XMLName xml.Name `xml:"ASNInstruction"`
	Value   string   `xml:"value,attr"` // (required|notRequired|optional) REQUIRED
	Lower   *Lower   `xml:"Lower,omitempty"`
	Upper   *Upper   `xml:"Upper,omitempty"`
}

// InvoiceInstruction controls invoice handling.
type InvoiceInstruction struct {
	XMLName           xml.Name        `xml:"InvoiceInstruction"`
	Value             string          `xml:"value,attr"`                      // (invoiceRequired|invoiceNotRequired|invoiceForbidden|evaluated) REQUIRED
	VerificationType  string          `xml:"verificationType,attr,omitempty"` // (2way|3way|4way)
	UnitPriceEditable string          `xml:"unitPriceEditable,attr,omitempty"`
	Lower             *Lower          `xml:"Lower,omitempty"`
	TemporaryPrice    *TemporaryPrice `xml:"TemporaryPrice,omitempty"`
	Upper             *Upper          `xml:"Upper,omitempty"`
}

// SESInstruction controls service entry sheet requirements.
type SESInstruction struct {
	XMLName           xml.Name `xml:"SESInstruction"`
	Value             string   `xml:"value,attr"` // (required|notRequired|optional) REQUIRED
	UnitPriceEditable string   `xml:"unitPriceEditable,attr,omitempty"`
	Lower             *Lower   `xml:"Lower,omitempty"`
	Upper             *Upper   `xml:"Upper,omitempty"`
}

// ─── Packaging ────────────────────────────────────────────────────────────────

// Packaging describes physical packaging for shipment.
type Packaging struct {
	XMLName                              xml.Name                              `xml:"Packaging"`
	PackagingCode                        []*PackagingCode                      `xml:"PackagingCode,omitempty"`
	Dimension                            []*Dimension                          `xml:"Dimension,omitempty"`
	Description                          *Description                          `xml:"Description,omitempty"`
	PackagingLevelCode                   *PackagingLevelCode                   `xml:"PackagingLevelCode,omitempty"`
	ShippingContainerSerialCode          *ShippingContainerSerialCode          `xml:"ShippingContainerSerialCode,omitempty"`
	ShippingContainerSerialCodeReference *ShippingContainerSerialCodeReference `xml:"ShippingContainerSerialCodeReference,omitempty"`
	AssetInfo                            []*AssetInfo                          `xml:"AssetInfo,omitempty"`
	BestBeforeDate                       *BestBeforeDate                       `xml:"BestBeforeDate,omitempty"`
	DispatchQuantity                     *DispatchQuantity                     `xml:"DispatchQuantity,omitempty"`
	Extrinsic                            []*Extrinsic                          `xml:"Extrinsic,omitempty"`
	FreeGoodsQuantity                    *FreeGoodsQuantity                    `xml:"FreeGoodsQuantity,omitempty"`
	OrderedQuantity                      *OrderedQuantity                      `xml:"OrderedQuantity,omitempty"`
	PackageID                            *PackageID                            `xml:"PackageID,omitempty"`
	PackageTypeCodeIdentifierCode        *PackageTypeCodeIdentifierCode        `xml:"PackageTypeCodeIdentifierCode,omitempty"`
	PackagingIndustry                    *PackagingIndustry                    `xml:"PackagingIndustry,omitempty"`
	QuantityVarianceNote                 *QuantityVarianceNote                 `xml:"QuantityVarianceNote,omitempty"`
	ShippingMark                         *ShippingMark                         `xml:"ShippingMark,omitempty"`
	StoreCode                            *StoreCode                            `xml:"StoreCode,omitempty"`
}

// PackagingCode is a packaging type code.
type PackagingCode struct {
	XMLName xml.Name `xml:"PackagingCode"`
	XMLLang string   `xml:"xml:lang,attr,omitempty"`
	Value   string   `xml:",chardata"`
}

// PackagingLevelCode is the level in a multi-level packaging hierarchy.
type PackagingLevelCode struct {
	XMLName xml.Name `xml:"PackagingLevelCode"`
	Value   string   `xml:",chardata"`
}

// ShippingContainerSerialCode is an SSCC barcode value.
type ShippingContainerSerialCode struct {
	XMLName xml.Name `xml:"ShippingContainerSerialCode"`
	Value   string   `xml:",chardata"`
}

// ShippingContainerSerialCodeReference refers to another container.
type ShippingContainerSerialCodeReference struct {
	XMLName xml.Name `xml:"ShippingContainerSerialCodeReference"`
	Value   string   `xml:",chardata"`
}

// ─── Packaging sub-types ──────────────────────────────────────────────────────

// TemporaryPrice is an EMPTY flag element indicating a temporary price applies.
type TemporaryPrice struct {
	XMLName xml.Name `xml:"TemporaryPrice"`
	Value   string   `xml:"value,attr"` // (yes|no) REQUIRED
}

// DispatchQuantity is the quantity dispatched for delivery.
type DispatchQuantity struct {
	XMLName       xml.Name       `xml:"DispatchQuantity"`
	Quantity      string         `xml:"quantity,attr,omitempty"`
	UnitOfMeasure *UnitOfMeasure `xml:"UnitOfMeasure,omitempty"`
}

// FreeGoodsQuantity is the quantity delivered at no cost (samples, promotions, etc.).
type FreeGoodsQuantity struct {
	XMLName       xml.Name       `xml:"FreeGoodsQuantity"`
	Quantity      string         `xml:"quantity,attr,omitempty"`
	UnitOfMeasure *UnitOfMeasure `xml:"UnitOfMeasure,omitempty"`
}

// PackageID provides identifiers for an individual package.
type PackageID struct {
	XMLName                 xml.Name                 `xml:"PackageID"`
	GlobalIndividualAssetID *GlobalIndividualAssetID `xml:"GlobalIndividualAssetID,omitempty"`
	ReturnablePackageID     *ReturnablePackageID     `xml:"ReturnablePackageID,omitempty"`
	PackageTrackingID       *PackageTrackingID       `xml:"PackageTrackingID,omitempty"`
}

// GlobalIndividualAssetID is a global unique asset identifier.
type GlobalIndividualAssetID struct {
	XMLName xml.Name `xml:"GlobalIndividualAssetID"`
	Value   string   `xml:",chardata"`
}

// ReturnablePackageID identifies a returnable/deposit package.
type ReturnablePackageID struct {
	XMLName xml.Name `xml:"ReturnablePackageID"`
	Value   string   `xml:",chardata"`
}

// PackageTrackingID is the carrier tracking ID for a package.
type PackageTrackingID struct {
	XMLName xml.Name `xml:"PackageTrackingID"`
	Value   string   `xml:",chardata"`
}

// PackageTypeCodeIdentifierCode identifies the package type code scheme.
type PackageTypeCodeIdentifierCode struct {
	XMLName xml.Name `xml:"PackageTypeCodeIdentifierCode"`
	Value   string   `xml:",chardata"`
}

// PackagingIndustry groups industry-specific packaging extensions.
type PackagingIndustry struct {
	XMLName               xml.Name               `xml:"PackagingIndustry"`
	PackagingLifeSciences *PackagingLifeSciences `xml:"PackagingLifeSciences,omitempty"`
}

// PackagingLifeSciences holds life-sciences-specific packaging data.
type PackagingLifeSciences struct {
	XMLName            xml.Name              `xml:"PackagingLifeSciences"`
	MedicationListInfo []*MedicationListInfo `xml:"MedicationListInfo,omitempty"`
}

// QuantityVarianceNote is a note explaining a quantity variance (partial delivery).
type QuantityVarianceNote struct {
	XMLName xml.Name `xml:"QuantityVarianceNote"`
	Value   string   `xml:",chardata"`
}

// ShippingMark is a mark/label on the shipping package.
type ShippingMark struct {
	XMLName xml.Name `xml:"ShippingMark"`
	Value   string   `xml:",chardata"`
}

// StoreCode identifies the destination store for retail deliveries.
type StoreCode struct {
	XMLName xml.Name `xml:"StoreCode"`
	Value   string   `xml:",chardata"`
}
