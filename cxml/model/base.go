// Package model contains cXML data structures for all cXML document types.
// Spec: cXML Reference Guide, DTD version 1.2.069.
package model

import (
	"encoding/xml"
	"strconv"
	"strings"
)

// ─── Money ───────────────────────────────────────────────────────────────────

// Money represents a monetary amount with currency code.
// The textual content is the numeric amount as a string (e.g. "123.45").
type Money struct {
	XMLName           xml.Name `xml:"Money"`
	Currency          string   `xml:"currency,attr"`                    // ISO 4217 currency code, REQUIRED
	AlternateAmount   string   `xml:"alternateAmount,attr,omitempty"`   // amount in alternate currency
	AlternateCurrency string   `xml:"alternateCurrency,attr,omitempty"` // alternate currency code
	Value             string   `xml:",chardata"`                        // numeric amount
	Amount            float64  `xml:"-"`                                // backward-compat helper for existing consumers
}

// UnmarshalXML keeps both Value and Amount populated for compatibility.
func (m *Money) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type moneyAlias struct {
		Currency          string `xml:"currency,attr"`
		AlternateAmount   string `xml:"alternateAmount,attr,omitempty"`
		AlternateCurrency string `xml:"alternateCurrency,attr,omitempty"`
		Value             string `xml:",chardata"`
	}
	var aux moneyAlias
	if err := d.DecodeElement(&aux, &start); err != nil {
		return err
	}
	m.XMLName = start.Name
	m.Currency = aux.Currency
	m.AlternateAmount = aux.AlternateAmount
	m.AlternateCurrency = aux.AlternateCurrency
	m.Value = strings.TrimSpace(aux.Value)
	if m.Value != "" {
		if f, err := strconv.ParseFloat(m.Value, 64); err == nil {
			m.Amount = f
		}
	}
	return nil
}

// MarshalXML emits Value if provided; otherwise falls back to Amount.
func (m Money) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type moneyAlias struct {
		Currency          string `xml:"currency,attr"`
		AlternateAmount   string `xml:"alternateAmount,attr,omitempty"`
		AlternateCurrency string `xml:"alternateCurrency,attr,omitempty"`
		Value             string `xml:",chardata"`
	}
	value := strings.TrimSpace(m.Value)
	if value == "" {
		value = strconv.FormatFloat(m.Amount, 'f', -1, 64)
	}
	return e.EncodeElement(moneyAlias{
		Currency:          m.Currency,
		AlternateAmount:   m.AlternateAmount,
		AlternateCurrency: m.AlternateCurrency,
		Value:             value,
	}, start)
}

// Total wraps a single Money element.
type Total struct {
	XMLName       xml.Name       `xml:"Total"`
	Money         *Money         `xml:"Money"`
	Modifications *Modifications `xml:"Modifications,omitempty"`
}

// UnitPrice wraps Money with optional price modifications.
type UnitPrice struct {
	XMLName           xml.Name           `xml:"UnitPrice"`
	Money             *Money             `xml:"Money"`
	Modifications     *Modifications     `xml:"Modifications,omitempty"`
	PricingConditions *PricingConditions `xml:"PricingConditions,omitempty"`
}

// ─── Naming / text ───────────────────────────────────────────────────────────

// Name is a localised name element.
type Name struct {
	XMLName xml.Name `xml:"Name"`
	Lang    string   `xml:"xml:lang,attr,omitempty"`
	Value   string   `xml:",chardata"`
}

// ShortName is a short label embedded inside Description mixed content.
type ShortName struct {
	XMLName xml.Name `xml:"ShortName"`
	Value   string   `xml:",chardata"`
}

// Description contains a localised human-readable text.
// Mixed content (chardata + optional ShortName children) —
// Value holds all text nodes; ShortName is extracted separately.
type Description struct {
	XMLName   xml.Name   `xml:"Description"`
	Lang      string     `xml:"-"`
	ShortName string     `xml:"xml:lang,attr,omitempty"` // backward-compatible alias for lang
	Type      string     `xml:"type,attr,omitempty"`
	Short     *ShortName `xml:"ShortName,omitempty"`
	Value     string     `xml:",chardata"`
}

// Comments is a localised note, optionally containing Attachment children.
type Comments struct {
	XMLName     xml.Name      `xml:"Comments"`
	Lang        string        `xml:"xml:lang,attr,omitempty"`
	Type        string        `xml:"type,attr,omitempty"`
	Attachments []*Attachment `xml:"Attachment,omitempty"`
	Value       string        `xml:",chardata"`
}

// ─── Attachments / URLs ──────────────────────────────────────────────────────

// URL holds a hyperlink reference.
type URL struct {
	XMLName xml.Name `xml:"URL"`
	Name    string   `xml:"name,attr,omitempty"`
	Value   string   `xml:",chardata"`
}

// Attachment references an attached file via URL.
type Attachment struct {
	XMLName    xml.Name `xml:"Attachment"`
	Visibility string   `xml:"visibility,attr,omitempty"` // (internal)
	URL        *URL     `xml:"URL"`
}

// AttachmentReference refers to a stored attachment by name.
type AttachmentReference struct {
	XMLName    xml.Name    `xml:"AttachmentReference"`
	Length     string      `xml:"length,attr,omitempty"`
	Version    string      `xml:"version,attr,omitempty"`
	Name       *Name       `xml:"Name"`
	InternalID *InternalID `xml:"InternalID"`
	URL        *URL        `xml:"URL,omitempty"`
}

// ─── Identifiers ─────────────────────────────────────────────────────────────

// Extrinsic is an extensible name/value element for custom data.
type Extrinsic struct {
	XMLName xml.Name `xml:"Extrinsic"`
	Name    string   `xml:"name,attr"` // REQUIRED
	Content string   `xml:",innerxml"`
}

// InternalID is a buyer-side document identifier.
type InternalID struct {
	XMLName xml.Name `xml:"InternalID"`
	Domain  string   `xml:"domain,attr,omitempty"`
	Value   string   `xml:",chardata"`
}

// IdReference links to another document or external system record.
type IdReference struct {
	XMLName     xml.Name     `xml:"IdReference"`
	Identifier  string       `xml:"identifier,attr"` // REQUIRED
	Domain      string       `xml:"domain,attr"`     // REQUIRED
	Creator     *Creator     `xml:"Creator,omitempty"`
	Description *Description `xml:"Description,omitempty"`
}

// Creator records who created a reference.
type Creator struct {
	XMLName xml.Name `xml:"Creator"`
	Lang    string   `xml:"xml:lang,attr,omitempty"`
	Value   string   `xml:",chardata"`
}

// DocumentReference points to another cXML document by its payloadID.
type DocumentReference struct {
	XMLName   xml.Name `xml:"DocumentReference"`
	PayloadID string   `xml:"payloadID,attr"` // REQUIRED
}

// DocumentInfo is an EMPTY element identifying an external document.
type DocumentInfo struct {
	XMLName      xml.Name `xml:"DocumentInfo"`
	DocumentID   string   `xml:"documentID,attr"`   // REQUIRED
	DocumentType string   `xml:"documentType,attr"` // REQUIRED
	DocumentDate string   `xml:"documentDate,attr,omitempty"`
}

// ─── Units / measures ────────────────────────────────────────────────────────

// UnitOfMeasure is a UN/CEFACT unit code (e.g. "EA", "KG").
type UnitOfMeasure struct {
	XMLName xml.Name `xml:"UnitOfMeasure"`
	Value   string   `xml:",chardata"`
}

// Classification assigns an item to a commodity category.
type Classification struct {
	XMLName xml.Name `xml:"Classification"`
	Domain  string   `xml:"domain,attr,omitempty"` // e.g. "UNSPSC"
	Value   string   `xml:",chardata"`
}

// Percentage holds a numeric percentage as an attribute.
type Percentage struct {
	XMLName xml.Name `xml:"Percentage"`
	Percent string   `xml:"percent,attr"` // REQUIRED
}

// PriceBasisQuantity indicates the quantity basis for a unit price.
type PriceBasisQuantity struct {
	XMLName       xml.Name       `xml:"PriceBasisQuantity"`
	Quantity      string         `xml:"quantity,attr"` // REQUIRED
	UnitOfMeasure *UnitOfMeasure `xml:"UnitOfMeasure"`
	Description   *Description   `xml:"Description,omitempty"`
}

// ─── Manufacturers ───────────────────────────────────────────────────────────

// ManufacturerPartID is the manufacturer's own part number.
type ManufacturerPartID struct {
	XMLName xml.Name `xml:"ManufacturerPartID"`
	Value   string   `xml:",chardata"`
}

// ManufacturerName is the manufacturer's company name.
type ManufacturerName struct {
	XMLName xml.Name `xml:"ManufacturerName"`
	Lang    string   `xml:"xml:lang,attr,omitempty"`
	Value   string   `xml:",chardata"`
}

// ─── Time / period ───────────────────────────────────────────────────────────

// Period defines a date range.
type Period struct {
	XMLName   xml.Name `xml:"Period"`
	StartDate string   `xml:"startDate,attr"` // REQUIRED
	EndDate   string   `xml:"endDate,attr"`   // REQUIRED
}

// ValidityPeriod is a pricing validity range.
type ValidityPeriod struct {
	XMLName   xml.Name `xml:"ValidityPeriod"`
	StartDate string   `xml:"startDate,attr,omitempty"`
	EndDate   string   `xml:"endDate,attr,omitempty"`
}

// ─── Pricing conditions ──────────────────────────────────────────────────────

// PricingConditions holds a set of price scales.
type PricingConditions struct {
	XMLName xml.Name `xml:"PricingConditions"`
	Scales  *Scales  `xml:"Scales,omitempty"`
}

// Scales wraps multiple Scale elements.
type Scales struct {
	XMLName xml.Name `xml:"Scales"`
	Scale   []*Scale `xml:"Scale"`
}

// Scale defines a quantity-based price tier.
type Scale struct {
	XMLName        xml.Name        `xml:"Scale"`
	Quantity       string          `xml:"quantity,attr,omitempty"`
	ConditionType  *ConditionType  `xml:"ConditionType,omitempty"`
	ValidityPeriod *ValidityPeriod `xml:"ValidityPeriod,omitempty"`
	CostTermValue  *CostTermValue  `xml:"CostTermValue,omitempty"`
}

// ConditionType classifies a pricing condition.
type ConditionType struct {
	XMLName xml.Name `xml:"ConditionType"`
	Value   string   `xml:",chardata"`
}

// CostTermValue holds a monetary or percentage cost term.
type CostTermValue struct {
	XMLName    xml.Name    `xml:"CostTermValue"`
	Money      *Money      `xml:"Money,omitempty"`
	Percentage *Percentage `xml:"Percentage,omitempty"`
}

// ─── Modifications ───────────────────────────────────────────────────────────

// Modifications holds a list of price modifications.
type Modifications struct {
	XMLName      xml.Name        `xml:"Modifications"`
	Modification []*Modification `xml:"Modification"`
}

// Modification represents a single price adjustment.
type Modification struct {
	XMLName             xml.Name             `xml:"Modification"`
	Level               string               `xml:"level,attr,omitempty"`
	OriginalPrice       *OriginalPrice       `xml:"OriginalPrice,omitempty"`
	AdditionalDeduction *AdditionalDeduction `xml:"AdditionalDeduction,omitempty"`
	AdditionalCost      *AdditionalCost      `xml:"AdditionalCost,omitempty"`
	Tax                 *Tax                 `xml:"Tax,omitempty"`
	ModificationDetail  *ModificationDetail  `xml:"ModificationDetail,omitempty"`
}

// OriginalPrice holds the pre-modification price.
type OriginalPrice struct {
	XMLName xml.Name `xml:"OriginalPrice"`
	Type    string   `xml:"type,attr,omitempty"`
	Money   *Money   `xml:"Money"`
}

// ModificationDetail describes the details of a modification.
type ModificationDetail struct {
	XMLName     xml.Name     `xml:"ModificationDetail"`
	Name        string       `xml:"name,attr"` // REQUIRED
	Code        string       `xml:"code,attr,omitempty"`
	StartDate   string       `xml:"startDate,attr,omitempty"`
	EndDate     string       `xml:"endDate,attr,omitempty"`
	Scope       string       `xml:"scope,attr,omitempty"`
	Description *Description `xml:"Description,omitempty"`
	IdReference *IdReference `xml:"IdReference,omitempty"`
}

// AdditionalDeduction represents a price reduction.
type AdditionalDeduction struct {
	XMLName          xml.Name          `xml:"AdditionalDeduction"`
	Type             string            `xml:"type,attr,omitempty"`
	DeductionAmount  *DeductionAmount  `xml:"DeductionAmount,omitempty"`
	DeductionPercent *DeductionPercent `xml:"DeductionPercent,omitempty"`
	DeductedPrice    *DeductedPrice    `xml:"DeductedPrice,omitempty"`
}

// DeductionAmount is a fixed monetary deduction. Contains Money.
type DeductionAmount struct {
	XMLName xml.Name `xml:"DeductionAmount"`
	Money   *Money   `xml:"Money"`
}

// DeductionPercent is a percentage-based deduction.
type DeductionPercent struct {
	XMLName xml.Name `xml:"DeductionPercent"`
	Percent string   `xml:"percent,attr"` // REQUIRED
}

// DeductedPrice is the post-deduction price.
type DeductedPrice struct {
	XMLName xml.Name `xml:"DeductedPrice"`
	Money   *Money   `xml:"Money"`
}

// AdditionalCost represents an added cost (money or percentage).
type AdditionalCost struct {
	XMLName    xml.Name    `xml:"AdditionalCost"`
	Money      *Money      `xml:"Money,omitempty"`
	Percentage *Percentage `xml:"Percentage,omitempty"`
}

// ─── Batch / tracking ────────────────────────────────────────────────────────

// Batch holds manufacturing batch/lot information.
type Batch struct {
	XMLName        xml.Name        `xml:"Batch"`
	BatchID        string          `xml:"batchID,attr,omitempty"`
	ExpirationDate string          `xml:"expirationDate,attr,omitempty"`
	ProductionDate string          `xml:"productionDate,attr,omitempty"`
	BestBeforeDate *BestBeforeDate `xml:"BestBeforeDate,omitempty"`
	SerialNumber   []*SerialNumber `xml:"SerialNumber,omitempty"`
}

// SupplierBatchID is the supplier's batch reference.
type SupplierBatchID struct {
	XMLName xml.Name `xml:"SupplierBatchID"`
	Value   string   `xml:",chardata"`
}

// BestBeforeDate is an expiry indicator.
type BestBeforeDate struct {
	XMLName xml.Name `xml:"BestBeforeDate"`
	Value   string   `xml:",chardata"`
}

// SerialNumber is an item serial tracking value.
type SerialNumber struct {
	XMLName xml.Name `xml:"SerialNumber"`
	Value   string   `xml:",chardata"`
}

// AssetInfo holds asset tracking data for a line item.
type AssetInfo struct {
	XMLName xml.Name `xml:"AssetInfo"`
	AssetID string   `xml:"assetID,attr,omitempty"`
	Serial  string   `xml:"serial,attr,omitempty"`
}

// ─── Hazardous material ───────────────────────────────────────────────────────

// Hazard identifies hazardous material classification.
type Hazard struct {
	XMLName     xml.Name     `xml:"Hazard"`
	Code        string       `xml:"code,attr,omitempty"`
	Description *Description `xml:"Description,omitempty"`
}

// ServiceLevel indicates transport service level.
type ServiceLevel struct {
	XMLName xml.Name `xml:"ServiceLevel"`
	Value   string   `xml:",chardata"`
}

// ─── Routing ─────────────────────────────────────────────────────────────────

// Path represents a routing path for multi-hop delivery.
type Path struct {
	XMLName xml.Name `xml:"Path"`
	Node    []*Node  `xml:"Node,omitempty"`
}

// Node is a waypoint in a routing path.
type Node struct {
	XMLName     xml.Name     `xml:"Node"`
	Type        string       `xml:"type,attr,omitempty"`
	Description *Description `xml:"Description,omitempty"`
}

// Charge is the monetary charge for a distribution line.
type Charge struct {
	XMLName xml.Name `xml:"Charge"`
	Money   *Money   `xml:"Money"`
}

// ─── Forward declarations for types defined in financial.go ─────────────────
// (Tax, Shipping, Distribution are defined in financial.go but referenced here)
