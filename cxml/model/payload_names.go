package model

// RequestPayloadName implementations for all Request payload types.
// PunchOutOrderMessage and OrderRequest are defined in punchout.go.

func (r *OrderChangeRequest) RequestPayloadName() string     { return "OrderChangeRequest" }
func (r *ConfirmationRequest) RequestPayloadName() string    { return "ConfirmationRequest" }
func (r *ProfileRequest) RequestPayloadName() string         { return "ProfileRequest" }
func (r *StatusUpdateRequest) RequestPayloadName() string    { return "StatusUpdateRequest" }
func (r *PunchOutSetupRequest) RequestPayloadName() string   { return "PunchOutSetupRequest" }
func (r *ReceivingAdviceRequest) RequestPayloadName() string { return "ReceivingAdviceRequest" }
func (r *ShipNoticeRequest) RequestPayloadName() string      { return "ShipNoticeRequest" }
func (r *InvoiceDetailRequest) RequestPayloadName() string   { return "InvoiceDetailRequest" }
