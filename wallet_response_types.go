package monero

// Balance represent the response data for a get balance request
type Balance struct {
	Balance               uint64              `json:"balance"`
	UnlockedBalance       uint64              `json:"unlocked_balance"`
	MultigsigImportNeeded bool                `json:"multisig_import_needed"`
	PerSubaddress         []PerSubAddressInfo `json:"per_subaddress"`
}

// SubAddressInfo contains information about a set of subaddresses in a get
// balance request
type SubAddressInfo struct {
	AddressIndex      uint32 `json:"address_index"`
	Address           string `json:"address"`
	Balance           uint64 `json:"balance"`
	UnlockedBalance   uint64 `json:"unlocked_balance"`
	Label             string `json:"label"`
	NumUnspentOutputs uint64 `json:"num_unspent_outputs"`
}

// PerSubAddressInfo ...
type PerSubAddressInfo struct {
	AccountIndex      uint32 `json:"account_index"`
	AddressIndex      uint32 `json:"address_index"`
	Address           string `json:"address"`
	Balance           uint64 `json:"balance"`
	UnlockedBalance   uint64 `json:"unlocked_balance"`
	Label             string `json:"label"`
	NumUnspentOutputs uint64 `json:"num_unspent_outputs"`
}

// AddressInfo contains information about a wallet address
type AddressInfo struct {
	Address      string `json:"address"`
	Label        string `json:"label"`
	AddressIndex uint32 `json:"address_index"`
	Used         bool   `json:"used"`
}

// Address represents the response data for a getaddress request
type Address struct {
	Addresses []AddressInfo `json:"addresses"`
}

// SubAddressIndex represents an index for a subaddress
type SubAddressIndex struct {
	Major uint32 `json:"major"`
	Minor uint32 `json:"minor"`
}

// AddressIndex represents the response data for a getaddressindex request
type AddressIndex struct {
	Index SubAddressIndex `json:"index"`
}

// CreatedAddress represents the response data for a create_address request
type CreatedAddress struct {
	Address      string `json:"address"`
	AddressIndex uint32 `json:"address_index"`
}

// SubAddressAccountInfo represents account information for a subaddress
type SubAddressAccountInfo struct {
	AccountIndex    uint32 `json:"account_index"`
	BaseAddress     string `json:"base_address"`
	Balance         uint64 `json:"balance"`
	UnlockedBalance uint64 `json:"unlocked_balance"`
	Label           string `json:"label"`
	Tag             string `json:"tag"`
}

// Accounts represents information returned from a getaccounts request
type Accounts struct {
	TotalBalance         uint64                  `json:"total_balance"`
	TotalUnlockedBalance uint64                  `json:"total_unlocked_balance"`
	Accounts             []SubAddressAccountInfo `json:"subaddress_accounts"`
}

// CreatedAccount represents information about a newly created account
type CreatedAccount struct {
	AccountIndex uint32 `json:"account_index"`
	Address      string `json:"address"`
}

// AccountTagInfo represents information about an account tag
type AccountTagInfo struct {
	Tag      string   `json:"tag"`
	Label    string   `json:"label"`
	Accounts []uint32 `json:"accounts"`
}

// AccountTags represents data returned by a get_account_tags request
type AccountTags struct {
	Tags []AccountTagInfo `json:"account_tags"`
}

// Height represents the latest (synced) blockchain height
type Height struct {
	Height uint64 `json:"height"`
}

// Transfer represents details about a particular transfer response
type Transfer struct {
	TxHash        string `json:"tx_hash"`
	TxKey         string `json:"tx_key"`
	Amount        uint64 `json:"amount"`
	Fee           uint64 `json:"fee"`
	TxBlob        string `json:"tx_blob"`
	TxMetadata    string `json:"tx_metadata"`
	MultisigTxSet string `json:"multisig_txset"`
	UnsignedTxSet string `json:"unsigned_txset"`
}

// TransferSplit represents a response from a transfer_split request
type TransferSplit struct {
	TxHashList     []string `json:"tx_hash_list"`
	TxKeyList      []string `json:"tx_key_list"`
	AmountList     []uint64 `json:"amount_list"`
	FeeList        []uint64 `json:"fee_list"`
	TxBlobList     []string `json:"tx_blob_list"`
	TxMetadataList []string `json:"tx_metadata_list"`
	MultisigTxSet  string   `json:"multisig_txset"`
	UnsignedTxSet  string   `json:"unsigned_txset"`
}

// SignedTransfer represents a response from a sign_transfer request
type SignedTransfer struct {
	SignedTxSet string   `json:"signed_txset"`
	TxHashList  []string `json:"tx_hash_list"`
	TxRawList   []string `json:"tx_raw_list"`
}

// SubmittedTransfer represents a response from a submit_transfer request
type SubmittedTransfer struct {
	TxHashList []string `json:"tx_hash_list"`
}

// RelayedTransaction represents a response from a relay_tx request
type RelayedTransaction struct {
	TxHash string `json:"tx_hash"`
}

// Payment contains information about a payment
type Payment struct {
	PaymentID       string `json:"payment_id"`
	TxHash          string `json:"tx_hash"`
	Amount          uint64 `json:"amount"`
	BlockHeight     uint64 `json:"block_height"`
	UnlockTime      uint64 `json:"unlock_time"`
	SubAddressIndex uint32 `json:"subaddr_index"`
}

// Payments represents a response from a get_payments request
type Payments struct {
	Payments []Payment `json:"payments"`
}

// TransferDetails contains information about an incoming transfer
type TransferDetails struct {
	Amount          uint64          `json:"amount"`
	Spent           bool            `json:"spent"`
	GlobalIndex     uint64          `json:"global_index"`
	TxHash          string          `json:"tx_hash"`
	SubAddressIndex SubAddressIndex `json:"subaddr_index"`
	KeyImage        string          `json:"key_image"`
}

// IncomingTransfersData represents a response from a incoming_transfers request
type IncomingTransfersData struct {
	Transfers []TransferDetails `json:"transfers"`
}

// IntegratedAddress represents an integrated address returned by a make_integrated_address request
type IntegratedAddress struct {
	IntegratedAddress string `json:"integrated_address"`
	PaymentID         string `json:"payment_id"`
}

// SplitAddress represents the split form of an integrated address returned by split_integrated_address
type SplitAddress struct {
	StandardAddress string `json:"standard_address"`
	PaymentID       string `json:"payment_id"`
	IsSubaddress    bool   `json:"is_subaddress"`
}

// TransactionKey represents a check_tx_key response
type TransactionKey struct {
	Received      uint64 `json:"received"`
	InPool        bool   `json:"in_pool"`
	Confirmations uint64 `json:"confirmations"`
}

// CheckedProof represents a check_tx_proof result
type CheckedProof struct {
	Good          bool   `json:"good"`
	Received      uint64 `json:"received"`
	InPool        bool   `json:"in_pool"`
	Confirmations uint64 `json:"confirmations"`
}

// TransferEntry represents a transaction entry from a get_transfers request
type TransferEntry struct {
	TransactionID                   string          `json:"txid"`
	PaymentID                       string          `json:"payment_id"`
	Height                          uint64          `json:"height"`
	Timestamp                       uint64          `json:"timestamp"`
	Amount                          uint64          `json:"amount"`
	Fee                             uint64          `json:"Fee"`
	Note                            string          `json:"note"`
	Destinations                    []Destination   `json:"destinations"`
	Type                            string          `json:"type"`
	UnlockTime                      uint64          `json:"unlock_time"`
	SubAddressIndex                 SubAddressIndex `json:"subaddr_index"`
	Address                         string          `json:"address"`
	DoubleSpendSeen                 bool            `json:"double_spend_seen"`
	Confirmations                   uint64          `json:"confirmations"`
	SuggestedConfirmationsThreshold uint64          `json:"suggested_confirmations_threshold"`
}

// Transfers represents transfers data from a get_transfers request
type Transfers struct {
	In      []TransferEntry `json:"in,omitempty"`
	Out     []TransferEntry `json:"out,omitempty"`
	Pending []TransferEntry `json:"pending,omitempty"`
	Failed  []TransferEntry `json:"failed,omitempty"`
	Pool    []TransferEntry `json:"pool,omitempty"`
}

// CheckedReserveProof represents the result of a reverse proof verification
type CheckedReserveProof struct {
	Good  bool   `json:"good"`
	Total uint64 `json:"total"`
	Spent uint64 `json:"spent"`
}

// SignedKeyImage represents the result of a export_key_images request
type SignedKeyImage struct {
	KeyImage  string `json:"key_image"`
	Signature string `json:"signature"`
}

// ImportedKeyImages represents the result of a import_key_images request
type ImportedKeyImages struct {
	Height  uint64 `json:"height"`
	Spent   uint64 `json:"spent"`
	Unspent uint64 `json:"unspent"`
}

// ParsedURI represents a parsed URI string
type ParsedURI struct {
	URI               URISpec  `json:"uri"`
	UnknownParameters []string `json:"unknown_parameters"`
}

// AddressBookEntry represents an entry from an addressbook
type AddressBookEntry struct {
	Index       uint64 `json:"index"`
	Address     string `json:"address"`
	PaymentID   string `json:"payment_id"`
	Description string `json:"description"`
}

// RefreshResult represents the result of a blockchain refresh
type RefreshResult struct {
	BlocksFetch   uint64 `json:"blocks_fetched"`
	ReceivedMoney bool   `json:"received_money"`
}

// MultisigInfo ...
type MultisigInfo struct {
	Multisig  bool   `json:"multisig"`
	Ready     bool   `json:"ready"`
	Threshold uint32 `json:"threshold"`
	Total     uint32 `json:"total"`
}

// Multisig represents a partially created multisig address
type Multisig struct {
	Address      string `json:"address"`
	MultisigInfo string `json:"multisig_info"`
}

// SignedMultisigTransaction represents a signed multisig transaction
type SignedMultisigTransaction struct {
	DataHex  string   `json:"tx_data_hex"`
	HashList []string `json:"tx_hash_list"`
}

/// GetTransferByTxIDResponse ...
type GetTransferByTxIDResponse struct {
	Transfer  TransferEntry   `json:"transfer"`
	Transfers []TransferEntry `json:"transfers"`
}
