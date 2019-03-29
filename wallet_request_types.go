package monero

// GetBalance describes parameters for a getbalance request
type GetBalance struct {
	AccountIndex   uint32   `json:"account_index"`
	AddressIndices []uint32 `json:"address_indices,omitempty"`
}

// AddressFilters describes filter parameters for a getaddress request
type AddressFilters struct {
	AccountIndex uint32   `json:"account_index"`
	AddressIndex []uint32 `json:"address_index,omitempty"`
}

// GetAddressIndex describes the request parameters for a getaddressindex request
type GetAddressIndex struct {
	Address string `json:"address"`
}

// CreateAddress describes the request parameters for a create_address request
type CreateAddress struct {
	AccountIndex uint32 `json:"account_index"`
	Label        string `json:"label,omitempty"`
}

// LabelAddress describes the request parameters for a label_address request
type LabelAddress struct {
	AccountIndex uint32 `json:"index"`
	Label        string `json:"label"`
}

// GetAccounts describes the request parameters for a get_accounts request
type GetAccounts struct {
	Tag string `json:"tag,omitempty"`
}

// CreateAccount describes the request parameters for a create_account request
type CreateAccount struct {
	Label string `json:"label"`
}

// LabelAccount describes the request parameters for a label_address request
type LabelAccount struct {
	AccountIndex uint32 `json:"account_index"`
	Label        string `json:"label"`
}

// TagAccounts describes the request parameters for a tag_accounts request
type TagAccounts struct {
	Tag      string   `json:"tag"`
	Accounts []uint32 `json:"accounts"`
}

// UntagAccounts describes the request parameters for a untag_accounts request
type UntagAccounts struct {
	Accounts []uint32 `json:"accounts"`
}

// SetAccountTagDescription describes the request parameters for a set_account_tag_description request
type SetAccountTagDescription struct {
	Tag         string `json:"tag"`
	Description string `json:"description"`
}

// Destination represents details about a destination for a transfer
type Destination struct {
	Amount  uint64 `json:"amount"`
	Address string `json:"address"`
}

// TransferInput represents details about parameters in a transaction request
type TransferInput struct {
	Destinations      []Destination `json:"destinations"`
	AccountIndex      uint32        `json:"account_index"`
	SubAddressIndices []uint32      `json:"subaddr_indices,omitempty"`
	Priority          *uint32       `json:"priority"`
	Mixin             uint64        `json:"mixin"`
	RingSize          uint64        `json:"ring_size"`
	UnlockTime        *uint64       `json:"unlock_time,omitempty"`
	PaymentID         *string       `json:"payment_id,omitempty"`
	GetTxKey          *bool         `json:"get_tx_key,omitempty"`
	DoNotRelay        *bool         `json:"do_not_relay,omitempty"`
	GetTxHex          *bool         `json:"get_tx_hex,omitempty"`
	TxMetadata        *bool         `json:"tx_metadata,omitempty"`
}

// SignTransfer represents details about parameters in a sign_transfer request
type SignTransfer struct {
	UnsignedTxSet string `json:"unsigned_txset"`
	ExportRaw     bool   `json:"export_raw"`
}

// SubmitTransfer represents details about parameters in a submit_transfer request
type SubmitTransfer struct {
	TxDataHex []string `json:"tx_data_hex"`
}

// SweepDust represents details about parameters in a sweep_dust request
type SweepDust struct {
	GetTxKeys     bool `json:"get_tx_keys"`
	DoNotRelay    bool `json:"do_not_relay"`
	GetTxHex      bool `json:"get_tx_hex"`
	GetTxMetadata bool `json:"get_tx_metadata"`
}

// SweepAllDust represents details about parameters in a sweep_all request
type SweepAllDust struct {
	Address           string   `json:"address"`
	AccountIndex      uint32   `json:"account_index"`
	SubAddressIndices []uint32 `json:"subaddr_indices,omitempty"`
	Priority          uint32   `json:"priority"`
	Mixin             uint64   `json:"mixin"`
	RingSize          uint64   `json:"ring_size"`
	UnlockTime        uint64   `json:"unlock_time"`
	PaymentID         string   `json:"payment_id"`
	GetTxKey          bool     `json:"get_tx_key"`
	BelowAmount       uint64   `json:"below_amount"`
	DoNotRelay        bool     `json:"do_not_relay"`
	GetTxHex          bool     `json:"get_tx_hex"`
	TxMetadata        bool     `json:"tx_metadata"`
}

// SweepSingle represents details about parameters in a sweep_single request
type SweepSingle struct {
	Address    string `json:"address"`
	Priority   uint32 `json:"priority"`
	Mixin      uint64 `json:"mixin"`
	RingSize   uint64 `json:"ring_size"`
	Outputs    uint64 `json:"outputs"`
	UnlockTime uint64 `json:"unlock_time"`
	PaymentID  string `json:"payment_id"`
	GetTxKey   bool   `json:"get_tx_key"`
	KeyImage   string `json:"key_image"`
	DoNotRelay bool   `json:"do_not_relay"`
	GetTxHex   bool   `json:"get_tx_hex"`
	TxMetadata bool   `json:"tx_metadata"`
}

// RelayTransaction represents details about parameters in a relay_tx request
type RelayTransaction struct {
	Hex string `json:"hex"`
}

// GetPayments represents details about parameters in a get_payments request
type GetPayments struct {
	PaymentID string `json:"payment_id"`
}

// GetBulkPayments represents details about parameters in a get_bulk_payments request
type GetBulkPayments struct {
	PaymentIDs     []string `json:"payment_ids,omitempty"`
	MinBlockHeight uint64   `json:"min_block_height"`
}

// IncomingTransfers represents details about parameters in a incoming_transfers request
type IncomingTransfers struct {
	TransferType      string   `json:"transfer_type,omitempty"`
	AccountIndex      uint32   `json:"account_index"`
	SubAddressIndices []uint32 `json:"subaddr_indices"`
}

// MakeIntegratedAddress request parameters
type MakeIntegratedAddress struct {
	StandardAddress string `json:"standard_address"`
	PaymentID       string `json:"payment_id"`
}

// CheckTransactionProof request parameters
type CheckTransactionProof struct {
	TransactionID string `json:"txid"`
	Address       string `json:"address"`
	Message       string `json:"message"`
	Signature     string `json:"signature"`
}

// GetTransfersFilter request parameters
type GetTransfersFilter struct {
	In                *bool    `json:"in,omitempty"`
	Out               *bool    `json:"out,omitempty"`
	Pending           *bool    `json:"pending,omitempty"`
	Failed            *bool    `json:"failed,omitempty"`
	Pool              *bool    `json:"pool,omitempty"`
	FilterByHeight    *bool    `json:"filter_by_height,omitempty"`
	MinHeight         *uint64  `json:"min_height,omitempty"`
	MaxHeight         *uint64  `json:"max_height,omitempty"`
	AccountIndex      *uint32  `json:"account_index,omitempty"`
	SubAddressIndices []uint32 `json:"subaddr_indices,omitempty"`
}

// URISpec make_uri request parameters
type URISpec struct {
	Address       string `json:"address"`
	PaymentID     string `json:"payment_id"`
	Amount        uint64 `json:"amount"`
	Description   string `json:"tx_description"`
	RecipientName string `json:"recipient_name"`
}

// StartMining start_mining request parameters
type StartMining struct {
	ThreadsCount     uint64 `json:"threads_count"`
	BackgroundMining bool   `json:"do_background_mining"`
	IgnoreBattery    bool   `json:"ignore_battery"`
}
