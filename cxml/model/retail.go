package model

import "encoding/xml"

// retail.go — Retail and industry-extension types shared across multiple DTDs.
// These are referenced by types in shipnotice.go, invoice.go, and logistics.go
// but live here so they are not accidentally validated against a single DTD.

// Characteristic is a generic name/value attribute pair (EMPTY element).
// DTD: cXML.dtd <!ELEMENT Characteristic EMPTY>
type Characteristic struct {
	XMLName xml.Name `xml:"Characteristic"`
	Domain  string   `xml:"domain,attr"`
	Value   string   `xml:"value,attr"`
	Code    string   `xml:"code,attr,omitempty"`
}

// EANID is the EAN/IAN barcode for an item (DEPRECATED — prefer IdReference).
type EANID struct {
	XMLName xml.Name `xml:"EANID"`
	Value   string   `xml:",chardata"`
}

// EuropeanWasteCatalogID is the EU waste catalogue identifier for an item.
type EuropeanWasteCatalogID struct {
	XMLName xml.Name `xml:"EuropeanWasteCatalogID"`
	Value   string   `xml:",chardata"`
}

// ItemIndicator is a generic domain/value flag element (EMPTY).
type ItemIndicator struct {
	XMLName xml.Name `xml:"ItemIndicator"`
	Domain  string   `xml:"domain,attr"`
	Value   string   `xml:"value,attr"`
}

// PromotionDealID identifies a promotional deal applied to a line.
type PromotionDealID struct {
	XMLName xml.Name `xml:"PromotionDealID"`
	Value   string   `xml:",chardata"`
}

// PromotionVariantID identifies the variant within a promotional deal.
type PromotionVariantID struct {
	XMLName xml.Name `xml:"PromotionVariantID"`
	Value   string   `xml:",chardata"`
}

// TotalRetailAmount wraps a Money element for the total retail amount.
type TotalRetailAmount struct {
	XMLName xml.Name `xml:"TotalRetailAmount"`
	Money   *Money   `xml:"Money"`
}

// AdditionalPrices holds supplementary price information for a line.
// Content is captured as raw XML for flexibility.
type AdditionalPrices struct {
	XMLName xml.Name `xml:"AdditionalPrices"`
	Content string   `xml:",innerxml"`
}

// AdditionalAmounts holds supplementary amount information for a summary.
// Content is captured as raw XML for flexibility.
type AdditionalAmounts struct {
	XMLName xml.Name `xml:"AdditionalAmounts"`
	Content string   `xml:",innerxml"`
}

// ItemDetailRetail contains retail-specific item detail fields.
// DTD: cXML.dtd <!ELEMENT ItemDetailRetail (EANID?, EuropeanWasteCatalogID?, Characteristic*)>
type ItemDetailRetail struct {
	XMLName                xml.Name                `xml:"ItemDetailRetail"`
	EANID                  *EANID                  `xml:"EANID,omitempty"`
	EuropeanWasteCatalogID *EuropeanWasteCatalogID `xml:"EuropeanWasteCatalogID,omitempty"`
	Characteristic         []*Characteristic       `xml:"Characteristic,omitempty"`
}

// InvoiceDetailItemRetail contains retail-specific invoice item fields.
type InvoiceDetailItemRetail struct {
	XMLName            xml.Name            `xml:"InvoiceDetailItemRetail"`
	AdditionalPrices   *AdditionalPrices   `xml:"AdditionalPrices,omitempty"`
	TotalRetailAmount  *TotalRetailAmount  `xml:"TotalRetailAmount,omitempty"`
	ItemIndicator      []*ItemIndicator    `xml:"ItemIndicator,omitempty"`
	PromotionDealID    *PromotionDealID    `xml:"PromotionDealID,omitempty"`
	PromotionVariantID *PromotionVariantID `xml:"PromotionVariantID,omitempty"`
}

// InvoiceDetailItemReferenceRetail contains retail-specific item reference fields.
type InvoiceDetailItemReferenceRetail struct {
	XMLName                xml.Name                `xml:"InvoiceDetailItemReferenceRetail"`
	EANID                  *EANID                  `xml:"EANID,omitempty"`
	EuropeanWasteCatalogID *EuropeanWasteCatalogID `xml:"EuropeanWasteCatalogID,omitempty"`
	Characteristic         []*Characteristic       `xml:"Characteristic,omitempty"`
}

// InvoiceDetailSummaryRetail contains retail-specific invoice summary fields.
type InvoiceDetailSummaryRetail struct {
	XMLName           xml.Name           `xml:"InvoiceDetailSummaryRetail"`
	AdditionalAmounts *AdditionalAmounts `xml:"AdditionalAmounts,omitempty"`
}
