package model

import "encoding/xml"

// InvoiceDetailRequest is the supplier-to-buyer invoice document.
// DTD: InvoiceDetail.dtd 1.2.069 — <!ELEMENT InvoiceDetailRequest>
type InvoiceDetailRequest struct {
	XMLName                    xml.Name                    `xml:"InvoiceDetailRequest"`
	InvoiceDetailRequestHeader *InvoiceDetailRequestHeader `xml:"InvoiceDetailRequestHeader"`
	InvoiceDetailOrder         []*InvoiceDetailOrder       `xml:"InvoiceDetailOrder,omitempty"`
	InvoiceDetailHeaderOrder   []*InvoiceDetailHeaderOrder `xml:"InvoiceDetailHeaderOrder,omitempty"`
	InvoiceDetailSummary       *InvoiceDetailSummary       `xml:"InvoiceDetailSummary"`
}

// InvoiceDetailRequestHeader contains invoice-level metadata.
type InvoiceDetailRequestHeader struct {
	XMLName                      xml.Name                      `xml:"InvoiceDetailRequestHeader"`
	InvoiceID                    string                        `xml:"invoiceID,attr"`                   // REQUIRED
	IsInformationOnly            string                        `xml:"isInformationOnly,attr,omitempty"` // (yes)
	Purpose                      string                        `xml:"purpose,attr,omitempty"`           // standard|creditMemo|debitMemo|lineLevelCreditMemo|lineLevelDebitMemo
	Operation                    string                        `xml:"operation,attr,omitempty"`         // new|delete
	InvoiceDate                  string                        `xml:"invoiceDate,attr"`                 // REQUIRED
	InvoiceOrigin                string                        `xml:"invoiceOrigin,attr,omitempty"`     // supplier|buyer
	IsERS                        string                        `xml:"isERS,attr,omitempty"`             // (yes)
	InvoiceDetailHeaderIndicator *InvoiceDetailHeaderIndicator `xml:"InvoiceDetailHeaderIndicator"`
	InvoiceDetailLineIndicator   *InvoiceDetailLineIndicator   `xml:"InvoiceDetailLineIndicator"`
	InvoicePartner               []*InvoicePartner             `xml:"InvoicePartner,omitempty"`
	DocumentReference            *DocumentReference            `xml:"DocumentReference,omitempty"`
	InvoiceIDInfo                *InvoiceIDInfo                `xml:"InvoiceIDInfo,omitempty"`
	PaymentProposalIDInfo        *PaymentProposalIDInfo        `xml:"PaymentProposalIDInfo,omitempty"`
	InvoiceDetailShipping        *InvoiceDetailShipping        `xml:"InvoiceDetailShipping,omitempty"`
	ShipNoticeIDInfo             *ShipNoticeIDInfo             `xml:"ShipNoticeIDInfo,omitempty"`
	InvoiceDetailPaymentTerm     []*InvoiceDetailPaymentTerm   `xml:"InvoiceDetailPaymentTerm,omitempty"`
	PaymentTerm                  []*PaymentTerm                `xml:"PaymentTerm,omitempty"`
	PaymentInformation           *PaymentInformation           `xml:"PaymentInformation,omitempty"`
	Period                       *Period                       `xml:"Period,omitempty"`
	Comments                     *Comments                     `xml:"Comments,omitempty"`
	IdReference                  []*IdReference                `xml:"IdReference,omitempty"`
	Extrinsic                    []*Extrinsic                  `xml:"Extrinsic,omitempty"`
}

// InvoiceDetailHeaderIndicator contains boolean flags for the invoice header.
type InvoiceDetailHeaderIndicator struct {
	XMLName                         xml.Name `xml:"InvoiceDetailHeaderIndicator"`
	IsHeaderInvoice                 string   `xml:"isHeaderInvoice,attr,omitempty"`
	IsPriceBasedLineLevelCreditMemo string   `xml:"isPriceBasedLineLevelCreditMemo,attr,omitempty"`
	IsVatRecoverable                string   `xml:"isVatRecoverable,attr,omitempty"`
}

// InvoiceDetailLineIndicator carries flags about which details are embedded in line items.
type InvoiceDetailLineIndicator struct {
	XMLName                 xml.Name `xml:"InvoiceDetailLineIndicator"`
	IsTaxInLine             string   `xml:"isTaxInLine,attr,omitempty"`             // (yes)
	IsSpecialHandlingInLine string   `xml:"isSpecialHandlingInLine,attr,omitempty"` // (yes)
	IsShippingInLine        string   `xml:"isShippingInLine,attr,omitempty"`        // (yes)
	IsDiscountInLine        string   `xml:"isDiscountInLine,attr,omitempty"`        // (yes)
	IsAccountingInLine      string   `xml:"isAccountingInLine,attr,omitempty"`      // (yes)
	IsPriceAdjustmentInLine string   `xml:"isPriceAdjustmentInLine,attr,omitempty"` // (yes)
}

// InvoicePartner identifies a party involved in the invoice (e.g. from, to, remitTo).
type InvoicePartner struct {
	XMLName     xml.Name       `xml:"InvoicePartner"`
	Contact     *Contact       `xml:"Contact"`
	IdReference []*IdReference `xml:"IdReference,omitempty"`
}

// InvoiceIDInfo is an alternate reference to a prior InvoiceDetailRequest.
type InvoiceIDInfo struct {
	XMLName     xml.Name `xml:"InvoiceIDInfo"`
	InvoiceID   string   `xml:"invoiceID,attr,omitempty"`
	InvoiceDate string   `xml:"invoiceDate,attr,omitempty"`
}

// PaymentProposalIDInfo references a PaymentProposalRequest.
type PaymentProposalIDInfo struct {
	XMLName           xml.Name `xml:"PaymentProposalIDInfo"`
	PaymentProposalID string   `xml:"paymentProposalID,attr,omitempty"`
}

// InvoiceDetailShipping contains shipping details within an invoice.
type InvoiceDetailShipping struct {
	XMLName            xml.Name              `xml:"InvoiceDetailShipping"`
	ShippingDate       string                `xml:"shippingDate,attr,omitempty"`
	Contact            []*Contact            `xml:"Contact"`
	CarrierIdentifier  []*CarrierIdentifier  `xml:"CarrierIdentifier,omitempty"`
	ShipmentIdentifier []*ShipmentIdentifier `xml:"ShipmentIdentifier,omitempty"`
	DocumentReference  *DocumentReference    `xml:"DocumentReference,omitempty"`
}

// InvoiceDetailPaymentTerm is the deprecated per-invoice payment term element.
// Use PaymentTerm instead in new implementations.
type InvoiceDetailPaymentTerm struct {
	XMLName           xml.Name `xml:"InvoiceDetailPaymentTerm"`
	PayInNumberOfDays string   `xml:"payInNumberOfDays,attr"` // REQUIRED
	PercentageRate    string   `xml:"percentageRate,attr"`    // REQUIRED
}

// PaymentInformation holds the payment net due date.
type PaymentInformation struct {
	XMLName           xml.Name `xml:"PaymentInformation"`
	PaymentNetDueDate string   `xml:"paymentNetDueDate,attr,omitempty"`
}

// ─── Order-level grouping ─────────────────────────────────────────────────────

// InvoiceDetailOrder groups invoice items for a particular order (line-level invoice).
type InvoiceDetailOrder struct {
	XMLName                     xml.Name                     `xml:"InvoiceDetailOrder"`
	InvoiceDetailOrderInfo      *InvoiceDetailOrderInfo      `xml:"InvoiceDetailOrderInfo"`
	InvoiceDetailReceiptInfo    *InvoiceDetailReceiptInfo    `xml:"InvoiceDetailReceiptInfo,omitempty"`
	InvoiceDetailShipNoticeInfo *InvoiceDetailShipNoticeInfo `xml:"InvoiceDetailShipNoticeInfo,omitempty"`
	InvoiceDetailItem           []*InvoiceDetailItem         `xml:"InvoiceDetailItem,omitempty"`
	InvoiceDetailServiceItem    []*InvoiceDetailServiceItem  `xml:"InvoiceDetailServiceItem,omitempty"`
}

// InvoiceDetailHeaderOrder groups header-level summary for a particular order (header invoice).
type InvoiceDetailHeaderOrder struct {
	XMLName                   xml.Name                   `xml:"InvoiceDetailHeaderOrder"`
	InvoiceDetailOrderInfo    *InvoiceDetailOrderInfo    `xml:"InvoiceDetailOrderInfo"`
	InvoiceDetailOrderSummary *InvoiceDetailOrderSummary `xml:"InvoiceDetailOrderSummary"`
}

// InvoiceDetailOrderInfo provides reference to the purchase order being invoiced.
type InvoiceDetailOrderInfo struct {
	XMLName                  xml.Name                  `xml:"InvoiceDetailOrderInfo"`
	OrderReference           *OrderReference           `xml:"OrderReference,omitempty"`
	MasterAgreementReference *MasterAgreementReference `xml:"MasterAgreementReference,omitempty"`
	MasterAgreementIDInfo    *MasterAgreementIDInfo    `xml:"MasterAgreementIDInfo,omitempty"`
	OrderIDInfo              *OrderIDInfo              `xml:"OrderIDInfo,omitempty"`
	SupplierOrderInfo        *SupplierOrderInfo        `xml:"SupplierOrderInfo,omitempty"`
}

// OrderIDInfo is the buyer-system order identifier (used when OrderReference is not available).
type OrderIDInfo struct {
	XMLName     xml.Name       `xml:"OrderIDInfo"`
	OrderID     string         `xml:"orderID,attr,omitempty"`
	OrderDate   string         `xml:"orderDate,attr,omitempty"`
	IdReference []*IdReference `xml:"IdReference,omitempty"`
}

// InvoiceDetailReceiptInfo cross-references a buyer receipt document.
type InvoiceDetailReceiptInfo struct {
	XMLName          xml.Name          `xml:"InvoiceDetailReceiptInfo"`
	ReceiptReference *ReceiptReference `xml:"ReceiptReference,omitempty"`
	ReceiptIDInfo    *ReceiptIDInfo    `xml:"ReceiptIDInfo,omitempty"`
}

// ReceiptReference is a pointer to a prior ReceiptRequest document.
type ReceiptReference struct {
	XMLName           xml.Name           `xml:"ReceiptReference"`
	ReceiptID         string             `xml:"receiptID,attr,omitempty"`
	ReceiptDate       string             `xml:"receiptDate,attr,omitempty"`
	DocumentReference *DocumentReference `xml:"DocumentReference"`
}

// ReceiptIDInfo identifies a receipt by the buyer's system ID.
type ReceiptIDInfo struct {
	XMLName     xml.Name       `xml:"ReceiptIDInfo"`
	ReceiptID   string         `xml:"receiptID,attr,omitempty"`
	ReceiptDate string         `xml:"receiptDate,attr,omitempty"`
	IdReference []*IdReference `xml:"IdReference,omitempty"`
}

// InvoiceDetailShipNoticeInfo cross-references a ship notice document.
type InvoiceDetailShipNoticeInfo struct {
	XMLName             xml.Name             `xml:"InvoiceDetailShipNoticeInfo"`
	ShipNoticeReference *ShipNoticeReference `xml:"ShipNoticeReference,omitempty"`
	ShipNoticeIDInfo    *ShipNoticeIDInfo    `xml:"ShipNoticeIDInfo,omitempty"`
}

// ─── Line items ───────────────────────────────────────────────────────────────

// InvoiceDetailItem represents a goods line item in a line-level invoice.
type InvoiceDetailItem struct {
	XMLName                          xml.Name                          `xml:"InvoiceDetailItem"`
	InvoiceLineNumber                string                            `xml:"invoiceLineNumber,attr"` // REQUIRED
	Quantity                         string                            `xml:"quantity,attr"`          // REQUIRED
	ReferenceDate                    string                            `xml:"referenceDate,attr,omitempty"`
	InspectionDate                   string                            `xml:"inspectionDate,attr,omitempty"`
	ParentInvoiceLineNumber          string                            `xml:"parentInvoiceLineNumber,attr,omitempty"`
	ItemType                         string                            `xml:"itemType,attr,omitempty"`
	CompositeItemType                string                            `xml:"compositeItemType,attr,omitempty"`
	Reason                           string                            `xml:"reason,attr,omitempty"`  // (return)
	IsAdHoc                          string                            `xml:"isAdHoc,attr,omitempty"` // (yes)
	UnitOfMeasure                    *UnitOfMeasure                    `xml:"UnitOfMeasure"`
	UnitPrice                        *UnitPrice                        `xml:"UnitPrice"`
	PriceBasisQuantity               *PriceBasisQuantity               `xml:"PriceBasisQuantity,omitempty"`
	InvoiceDetailItemReference       *InvoiceDetailItemReference       `xml:"InvoiceDetailItemReference"`
	ReceiptLineItemReference         *ReceiptLineItemReference         `xml:"ReceiptLineItemReference,omitempty"`
	ShipNoticeLineItemReference      *ShipNoticeLineItemReference      `xml:"ShipNoticeLineItemReference,omitempty"`
	SubtotalAmount                   *SubtotalAmount                   `xml:"SubtotalAmount,omitempty"`
	Tax                              *Tax                              `xml:"Tax,omitempty"`
	InvoiceDetailLineSpecialHandling *InvoiceDetailLineSpecialHandling `xml:"InvoiceDetailLineSpecialHandling,omitempty"`
	InvoiceDetailLineShipping        *InvoiceDetailLineShipping        `xml:"InvoiceDetailLineShipping,omitempty"`
	ShipNoticeIDInfo                 *ShipNoticeIDInfo                 `xml:"ShipNoticeIDInfo,omitempty"`
	GrossAmount                      *GrossAmount                      `xml:"GrossAmount,omitempty"`
	InvoiceDetailDiscount            *InvoiceDetailDiscount            `xml:"InvoiceDetailDiscount,omitempty"`
	InvoiceItemModifications         *InvoiceItemModifications         `xml:"InvoiceItemModifications,omitempty"`
	TotalCharges                     *TotalCharges                     `xml:"TotalCharges,omitempty"`
	TotalAllowances                  *TotalAllowances                  `xml:"TotalAllowances,omitempty"`
	TotalAmountWithoutTax            *TotalAmountWithoutTax            `xml:"TotalAmountWithoutTax,omitempty"`
	NetAmount                        *NetAmount                        `xml:"NetAmount,omitempty"`
	Distribution                     []*Distribution                   `xml:"Distribution,omitempty"`
	Packaging                        []*Packaging                      `xml:"Packaging,omitempty"`
	ServiceEntryItemReference        *ServiceEntryItemReference        `xml:"ServiceEntryItemReference,omitempty"`
	ServiceEntryItemIDInfo           *ServiceEntryItemIDInfo           `xml:"ServiceEntryItemIDInfo,omitempty"`
	ProductMovementItemIDInfo        *ProductMovementItemIDInfo        `xml:"ProductMovementItemIDInfo,omitempty"`
	InvoiceDetailItemIndustry        *InvoiceDetailItemIndustry        `xml:"InvoiceDetailItemIndustry,omitempty"`
	CustomsInfo                      *CustomsInfo                      `xml:"CustomsInfo,omitempty"`
	Comments                         *Comments                         `xml:"Comments,omitempty"`
	Extrinsic                        []*Extrinsic                      `xml:"Extrinsic,omitempty"`
}

// InvoiceDetailServiceItem represents a services line item in a line-level invoice.
type InvoiceDetailServiceItem struct {
	XMLName                           xml.Name                           `xml:"InvoiceDetailServiceItem"`
	InvoiceLineNumber                 string                             `xml:"invoiceLineNumber,attr"` // REQUIRED
	Quantity                          string                             `xml:"quantity,attr,omitempty"`
	ReferenceDate                     string                             `xml:"referenceDate,attr,omitempty"`
	InspectionDate                    string                             `xml:"inspectionDate,attr,omitempty"`
	ParentInvoiceLineNumber           string                             `xml:"parentInvoiceLineNumber,attr,omitempty"`
	ItemType                          string                             `xml:"itemType,attr,omitempty"`
	IsAdHoc                           string                             `xml:"isAdHoc,attr,omitempty"` // (yes)
	InvoiceDetailServiceItemReference *InvoiceDetailServiceItemReference `xml:"InvoiceDetailServiceItemReference"`
	SubtotalAmount                    *SubtotalAmount                    `xml:"SubtotalAmount"`
	Period                            *Period                            `xml:"Period,omitempty"`
	UnitOfMeasure                     *UnitOfMeasure                     `xml:"UnitOfMeasure,omitempty"`
	UnitPrice                         *UnitPrice                         `xml:"UnitPrice,omitempty"`
	Tax                               *Tax                               `xml:"Tax,omitempty"`
	GrossAmount                       *GrossAmount                       `xml:"GrossAmount,omitempty"`
	InvoiceDetailDiscount             *InvoiceDetailDiscount             `xml:"InvoiceDetailDiscount,omitempty"`
	InvoiceItemModifications          *InvoiceItemModifications          `xml:"InvoiceItemModifications,omitempty"`
	TotalCharges                      *TotalCharges                      `xml:"TotalCharges,omitempty"`
	TotalAllowances                   *TotalAllowances                   `xml:"TotalAllowances,omitempty"`
	TotalAmountWithoutTax             *TotalAmountWithoutTax             `xml:"TotalAmountWithoutTax,omitempty"`
	NetAmount                         *NetAmount                         `xml:"NetAmount,omitempty"`
	Distribution                      []*Distribution                    `xml:"Distribution,omitempty"`
	ServiceEntryItemReference         *ServiceEntryItemReference         `xml:"ServiceEntryItemReference,omitempty"`
	ServiceEntryItemIDInfo            *ServiceEntryItemIDInfo            `xml:"ServiceEntryItemIDInfo,omitempty"`
	UnitRate                          *UnitRate                          `xml:"UnitRate,omitempty"`
	PriceBasisQuantity                *PriceBasisQuantity                `xml:"PriceBasisQuantity,omitempty"`
	InvoiceLaborDetail                *InvoiceLaborDetail                `xml:"InvoiceLaborDetail,omitempty"`
	Comments                          *Comments                          `xml:"Comments,omitempty"`
	Extrinsic                         []*Extrinsic                       `xml:"Extrinsic,omitempty"`
}

// InvoiceDetailItemReference identifies the PO line being invoiced.
type InvoiceDetailItemReference struct {
	XMLName                            xml.Name                            `xml:"InvoiceDetailItemReference"`
	LineNumber                         string                              `xml:"lineNumber,attr"`             // REQUIRED
	SerialNumber                       string                              `xml:"serialNumber,attr,omitempty"` // DEPRECATED
	ItemID                             *ItemID                             `xml:"ItemID,omitempty"`
	Description                        *Description                        `xml:"Description,omitempty"`
	Classification                     []*Classification                   `xml:"Classification,omitempty"`
	ManufacturerPartID                 *ManufacturerPartID                 `xml:"ManufacturerPartID,omitempty"`
	ManufacturerName                   *ManufacturerName                   `xml:"ManufacturerName,omitempty"`
	Country                            *Country                            `xml:"Country,omitempty"`
	SerialNumbers                      []*SerialNumber                     `xml:"SerialNumber,omitempty"`
	SupplierBatchID                    *SupplierBatchID                    `xml:"SupplierBatchID,omitempty"`
	InvoiceDetailItemReferenceIndustry *InvoiceDetailItemReferenceIndustry `xml:"InvoiceDetailItemReferenceIndustry,omitempty"`
}

// InvoiceDetailServiceItemReference identifies the master agreement line being invoiced.
type InvoiceDetailServiceItemReference struct {
	XMLName        xml.Name          `xml:"InvoiceDetailServiceItemReference"`
	LineNumber     string            `xml:"lineNumber,attr,omitempty"`
	Classification []*Classification `xml:"Classification,omitempty"`
	ItemID         *ItemID           `xml:"ItemID,omitempty"`
	Description    *Description      `xml:"Description,omitempty"`
}

// ReceiptLineItemReference cross-references a receipt document line.
type ReceiptLineItemReference struct {
	XMLName           xml.Name `xml:"ReceiptLineItemReference"`
	ReceiptLineNumber string   `xml:"receiptLineNumber,attr"` // REQUIRED
}

// ─── Line-level financial helpers ────────────────────────────────────────────

// SubtotalAmount wraps a Money element for a line subtotal.
type SubtotalAmount struct {
	XMLName xml.Name `xml:"SubtotalAmount"`
	Money   *Money   `xml:"Money"`
}

// GrossAmount is the subtotal plus taxes and charges before discounts.
type GrossAmount struct {
	XMLName xml.Name `xml:"GrossAmount"`
	Money   *Money   `xml:"Money"`
}

// NetAmount is the gross amount minus discounts.
type NetAmount struct {
	XMLName xml.Name `xml:"NetAmount"`
	Money   *Money   `xml:"Money"`
}

// TotalCharges is the total of all charges on an invoice.
type TotalCharges struct {
	XMLName xml.Name `xml:"TotalCharges"`
	Money   *Money   `xml:"Money"`
}

// TotalAllowances is the total of all allowances on an invoice.
type TotalAllowances struct {
	XMLName xml.Name `xml:"TotalAllowances"`
	Money   *Money   `xml:"Money"`
}

// TotalAmountWithoutTax is the sum of subtotal, charges, and allowances excluding tax.
type TotalAmountWithoutTax struct {
	XMLName xml.Name `xml:"TotalAmountWithoutTax"`
	Money   *Money   `xml:"Money"`
}

// InvoiceDetailDiscount represents a discount applied at line or header level.
type InvoiceDetailDiscount struct {
	XMLName        xml.Name        `xml:"InvoiceDetailDiscount"`
	PercentageRate string          `xml:"percentageRate,attr,omitempty"`
	Money          *Money          `xml:"Money"`
	Distribution   []*Distribution `xml:"Distribution,omitempty"`
}

// InvoiceItemModifications holds per-line allowances and charges.
type InvoiceItemModifications struct {
	XMLName      xml.Name        `xml:"InvoiceItemModifications"`
	Modification []*Modification `xml:"Modification"`
}

// InvoiceHeaderModifications holds header-level allowances and charges.
type InvoiceHeaderModifications struct {
	XMLName      xml.Name        `xml:"InvoiceHeaderModifications"`
	Modification []*Modification `xml:"Modification"`
}

// InvoiceDetailLineSpecialHandling contains special handling info for a line.
type InvoiceDetailLineSpecialHandling struct {
	XMLName      xml.Name        `xml:"InvoiceDetailLineSpecialHandling"`
	Description  *Description    `xml:"Description,omitempty"`
	Money        *Money          `xml:"Money"`
	Distribution []*Distribution `xml:"Distribution,omitempty"`
}

// InvoiceDetailLineShipping contains per-line shipping details.
type InvoiceDetailLineShipping struct {
	XMLName               xml.Name               `xml:"InvoiceDetailLineShipping"`
	InvoiceDetailShipping *InvoiceDetailShipping `xml:"InvoiceDetailShipping"`
	Money                 *Money                 `xml:"Money"`
	Distribution          []*Distribution        `xml:"Distribution,omitempty"`
}

// ─── Header-level order summary ───────────────────────────────────────────────

// InvoiceDetailOrderSummary is an order-level summary used in header invoices.
type InvoiceDetailOrderSummary struct {
	XMLName                          xml.Name                          `xml:"InvoiceDetailOrderSummary"`
	InvoiceLineNumber                string                            `xml:"invoiceLineNumber,attr"` // REQUIRED
	InspectionDate                   string                            `xml:"inspectionDate,attr,omitempty"`
	SubtotalAmount                   *SubtotalAmount                   `xml:"SubtotalAmount"`
	Period                           *Period                           `xml:"Period,omitempty"`
	Tax                              *Tax                              `xml:"Tax,omitempty"`
	InvoiceDetailLineSpecialHandling *InvoiceDetailLineSpecialHandling `xml:"InvoiceDetailLineSpecialHandling,omitempty"`
	InvoiceDetailLineShipping        *InvoiceDetailLineShipping        `xml:"InvoiceDetailLineShipping,omitempty"`
	GrossAmount                      *GrossAmount                      `xml:"GrossAmount,omitempty"`
	InvoiceDetailDiscount            *InvoiceDetailDiscount            `xml:"InvoiceDetailDiscount,omitempty"`
	NetAmount                        *NetAmount                        `xml:"NetAmount,omitempty"`
	Comments                         *Comments                         `xml:"Comments,omitempty"`
	Extrinsic                        []*Extrinsic                      `xml:"Extrinsic,omitempty"`
}

// InvoiceDetailSummaryLineItemModifications is a summary of all line modifications.
type InvoiceDetailSummaryLineItemModifications struct {
	XMLName      xml.Name        `xml:"InvoiceDetailSummaryLineItemModifications"`
	Modification []*Modification `xml:"Modification"`
}

// ─── Invoice summary ──────────────────────────────────────────────────────────

// InvoiceDetailSummary is the invoice totals section.
type InvoiceDetailSummary struct {
	XMLName                                   xml.Name                                   `xml:"InvoiceDetailSummary"`
	SubtotalAmount                            *SubtotalAmount                            `xml:"SubtotalAmount"`
	Tax                                       *Tax                                       `xml:"Tax"`
	SpecialHandlingAmount                     *SpecialHandlingAmount                     `xml:"SpecialHandlingAmount,omitempty"`
	ShippingAmount                            *ShippingAmount                            `xml:"ShippingAmount,omitempty"`
	GrossAmount                               *GrossAmount                               `xml:"GrossAmount,omitempty"`
	InvoiceDetailDiscount                     *InvoiceDetailDiscount                     `xml:"InvoiceDetailDiscount,omitempty"`
	InvoiceHeaderModifications                *InvoiceHeaderModifications                `xml:"InvoiceHeaderModifications,omitempty"`
	InvoiceDetailSummaryLineItemModifications *InvoiceDetailSummaryLineItemModifications `xml:"InvoiceDetailSummaryLineItemModifications,omitempty"`
	TotalCharges                              *TotalCharges                              `xml:"TotalCharges,omitempty"`
	TotalAllowances                           *TotalAllowances                           `xml:"TotalAllowances,omitempty"`
	TotalAmountWithoutTax                     *TotalAmountWithoutTax                     `xml:"TotalAmountWithoutTax,omitempty"`
	NetAmount                                 *NetAmount                                 `xml:"NetAmount"`
	DepositAmount                             *DepositAmount                             `xml:"DepositAmount,omitempty"`
	DueAmount                                 *DueAmount                                 `xml:"DueAmount,omitempty"`
	InvoiceDetailSummaryIndustry              *InvoiceDetailSummaryIndustry              `xml:"InvoiceDetailSummaryIndustry,omitempty"`
}

// SpecialHandlingAmount is a special handling surcharge.
type SpecialHandlingAmount struct {
	XMLName      xml.Name        `xml:"SpecialHandlingAmount"`
	Description  *Description    `xml:"Description,omitempty"`
	Money        *Money          `xml:"Money"`
	Distribution []*Distribution `xml:"Distribution,omitempty"`
}

// ShippingAmount is the total shipping charge on an invoice.
type ShippingAmount struct {
	XMLName      xml.Name        `xml:"ShippingAmount"`
	Money        *Money          `xml:"Money"`
	Distribution []*Distribution `xml:"Distribution,omitempty"`
}

// DepositAmount is the prepayment/deposit deducted from the total due.
type DepositAmount struct {
	XMLName xml.Name `xml:"DepositAmount"`
	Money   *Money   `xml:"Money"`
}

// DueAmount is the net amount payable after deposits and credits.
type DueAmount struct {
	XMLName xml.Name `xml:"DueAmount"`
	Money   *Money   `xml:"Money"`
}

// ─── Service entry / labor types ──────────────────────────────────────────────

// ServiceEntryItemReference cross-references a service entry document line.
type ServiceEntryItemReference struct {
	XMLName           xml.Name           `xml:"ServiceEntryItemReference"`
	ServiceLineNumber string             `xml:"serviceLineNumber,attr"`
	ServiceEntryID    string             `xml:"serviceEntryID,attr,omitempty"`
	ServiceEntryDate  string             `xml:"serviceEntryDate,attr,omitempty"`
	DocumentReference *DocumentReference `xml:"DocumentReference"`
}

// ServiceEntryItemIDInfo references a line in a prior ServiceEntryRequest by ID.
type ServiceEntryItemIDInfo struct {
	XMLName           xml.Name       `xml:"ServiceEntryItemIDInfo"`
	ServiceLineNumber string         `xml:"serviceLineNumber,attr"`
	ServiceEntryID    string         `xml:"serviceEntryID,attr"`
	ServiceEntryDate  string         `xml:"serviceEntryDate,attr,omitempty"`
	IdReference       []*IdReference `xml:"IdReference,omitempty"`
}

// ProductMovementItemIDInfo references a product movement document line.
type ProductMovementItemIDInfo struct {
	XMLName            xml.Name       `xml:"ProductMovementItemIDInfo"`
	MovementLineNumber string         `xml:"movementLineNumber,attr"`
	MovementID         string         `xml:"movementID,attr"`
	MovementDate       string         `xml:"movementDate,attr"`
	IdReference        []*IdReference `xml:"IdReference,omitempty"`
}

// InvoiceDetailItemIndustry groups industry-specific invoice item extensions.
type InvoiceDetailItemIndustry struct {
	XMLName                 xml.Name                 `xml:"InvoiceDetailItemIndustry"`
	InvoiceDetailItemRetail *InvoiceDetailItemRetail `xml:"InvoiceDetailItemRetail,omitempty"`
}

// CustomsInfo holds customs document references for an invoice line.
type CustomsInfo struct {
	XMLName      xml.Name      `xml:"CustomsInfo"`
	Issuer       *Issuer       `xml:"Issuer"`
	DocumentInfo *DocumentInfo `xml:"DocumentInfo"`
}

// InvoiceLaborDetail provides labor-related details for a service invoice line.
type InvoiceLaborDetail struct {
	XMLName               xml.Name               `xml:"InvoiceLaborDetail"`
	Contractor            *Contractor            `xml:"Contractor,omitempty"`
	JobDescription        *JobDescription        `xml:"JobDescription,omitempty"`
	Supervisor            *Supervisor            `xml:"Supervisor,omitempty"`
	WorkLocation          *WorkLocation          `xml:"WorkLocation,omitempty"`
	InvoiceTimeCardDetail *InvoiceTimeCardDetail `xml:"InvoiceTimeCardDetail,omitempty"`
}

// Contractor identifies the contracting party for a service.
type Contractor struct {
	XMLName              xml.Name              `xml:"Contractor"`
	ContractorIdentifier *ContractorIdentifier `xml:"ContractorIdentifier"`
	Contact              *Contact              `xml:"Contact"`
}

// ContractorIdentifier is the identifier string for a contractor.
type ContractorIdentifier struct {
	XMLName xml.Name `xml:"ContractorIdentifier"`
	Domain  string   `xml:"domain,attr"` // (supplierReferenceID|buyerReferenceID) REQUIRED
	Value   string   `xml:",chardata"`
}

// JobDescription wraps a textual description of the job performed.
type JobDescription struct {
	XMLName     xml.Name     `xml:"JobDescription"`
	Description *Description `xml:"Description"`
}

// Supervisor identifies the supervisor of the service work.
type Supervisor struct {
	XMLName xml.Name `xml:"Supervisor"`
	Contact *Contact `xml:"Contact"`
}

// WorkLocation identifies the physical location where work was performed.
type WorkLocation struct {
	XMLName xml.Name `xml:"WorkLocation"`
	Address *Address `xml:"Address"`
}

// InvoiceTimeCardDetail references the time card for a service invoice line.
type InvoiceTimeCardDetail struct {
	XMLName           xml.Name           `xml:"InvoiceTimeCardDetail"`
	TimeCardReference *TimeCardReference `xml:"TimeCardReference,omitempty"`
	TimeCardIDInfo    *TimeCardIDInfo    `xml:"TimeCardIDInfo,omitempty"`
}

// TimeCardReference points to a prior time card document.
type TimeCardReference struct {
	XMLName           xml.Name           `xml:"TimeCardReference"`
	TimeCardID        string             `xml:"timeCardID,attr,omitempty"`
	DocumentReference *DocumentReference `xml:"DocumentReference"`
}

// TimeCardIDInfo identifies a time card by its system ID.
type TimeCardIDInfo struct {
	XMLName    xml.Name `xml:"TimeCardIDInfo"`
	TimeCardID string   `xml:"timeCardID,attr"`
}

// InvoiceDetailItemReferenceIndustry groups industry-specific item reference extensions.
type InvoiceDetailItemReferenceIndustry struct {
	XMLName                          xml.Name                          `xml:"InvoiceDetailItemReferenceIndustry"`
	InvoiceDetailItemReferenceRetail *InvoiceDetailItemReferenceRetail `xml:"InvoiceDetailItemReferenceRetail,omitempty"`
}

// InvoiceDetailSummaryIndustry groups industry-specific invoice summary extensions.
type InvoiceDetailSummaryIndustry struct {
	XMLName                    xml.Name                    `xml:"InvoiceDetailSummaryIndustry"`
	InvoiceDetailSummaryRetail *InvoiceDetailSummaryRetail `xml:"InvoiceDetailSummaryRetail,omitempty"`
}
