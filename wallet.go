package monero

import (
	"errors"
)

// updated to include methods as of 0.14.0.2
const supportedMajorVersion = 1

// WalletClient is an rpc client
type WalletClient struct {
	*CallClient
}

// NewWalletClient creates a new wallet client
func NewWalletClient(endpoint, username, password string) *WalletClient {
	return &WalletClient{NewCallClient(endpoint, username, password)}
}

// GetBalances will fetch balances for all addresses and subaddress in a wallet
func (c *WalletClient) GetBalances() (Balance, error) {
	var rep Balance
	if err := c.Wallet("getbalance", nil, &rep); err != nil {
		return rep, err
	}
	return rep, nil
}

// GetBalanceForAccount fetches the balance for a given account
func (c *WalletClient) GetBalanceForAccount(accountIndex uint32) (Balance, error) {
	var rep Balance
	request := GetBalance{accountIndex, []uint32{}}
	if err := c.Wallet("getbalance", request, &rep); err != nil {
		return rep, err
	}
	return rep, nil
}

// GetAddresses will get information about all addresses in wallet
// filtered by the given parameters
func (c *WalletClient) GetAddresses(filters *AddressFilters) ([]AddressInfo, error) {
	var rep Address
	if err := c.Wallet("getaddress", filters, &rep); err != nil {
		return rep.Addresses, err
	}
	return rep.Addresses, nil
}

// GetAddressesByAccount will get all addresses for a given account
func (c *WalletClient) GetAddressesByAccount(accountIndex uint32) ([]AddressInfo, error) {
	var rep Address
	request := struct {
		AccountIndex uint32 `json:"account_index"`
	}{accountIndex}
	if err := c.Wallet("getaddress", &request, &rep); err != nil {
		return rep.Addresses, err
	}
	return rep.Addresses, nil
}

// GetAddressIndex will get the account index of a given address
func (c *WalletClient) GetAddressIndex() (uint32, error) {
	var rep AddressIndex
	err := c.Wallet("get_address_index", nil, &rep)
	if err != nil {
		return rep.Index, err
	}
	return rep.Index, nil
}

// CreateAddress will create a new subaddress for a given account, with a given label
func (c *WalletClient) CreateAddress(accountIndex uint32, label string) (CreatedAddress, error) {
	var rep CreatedAddress
	request := CreateAddress{accountIndex, label}
	err := c.Wallet("create_address", request, &rep)

	if err != nil {
		return rep, err
	}

	return rep, nil
}

// LabelAddress will set a label for a given accountIndex
func (c *WalletClient) LabelAddress(accountIndex uint32, label string) error {
	request := LabelAddress{accountIndex, label}
	err := c.Wallet("label_address", request, nil)

	if err != nil {
		return err
	}

	return nil
}

// GetAccounts will fetch information about accounts
func (c *WalletClient) GetAccounts(tag string) (Accounts, error) {
	var rep Accounts
	request := GetAccounts{tag}

	err := c.Wallet("get_accounts", request, &rep)

	if err != nil {
		return rep, err
	}

	return rep, nil
}

// CreateAccount will create a new account with a given label
func (c *WalletClient) CreateAccount(label string) (CreatedAccount, error) {
	var rep CreatedAccount
	request := CreateAccount{label}

	err := c.Wallet("create_account", request, &rep)

	if err != nil {
		return rep, err
	}

	return rep, nil
}

// LabelAccount will set a label for a given accountIndex
func (c *WalletClient) LabelAccount(accountIndex uint32, label string) error {
	request := LabelAccount{accountIndex, label}
	err := c.Wallet("label_account", request, nil)

	if err != nil {
		return err
	}

	return nil
}

// GetAccountTags will get tags for all accounts
func (c *WalletClient) GetAccountTags() ([]AccountTagInfo, error) {
	var response AccountTags
	err := c.Wallet("get_account_tags", nil, &response)

	if err != nil {
		return response.Tags, err
	}

	return response.Tags, nil
}

// TagAccounts will set a tag for a given set of accounts
func (c *WalletClient) TagAccounts(tag string, accounts []uint32) error {
	request := TagAccounts{tag, accounts}
	err := c.Wallet("tag_accounts", &request, nil)

	if err != nil {
		return err
	}

	return nil
}

// UntagAccounts will remove a tag for a given set of accounts
func (c *WalletClient) UntagAccounts(accounts []uint32) error {
	request := UntagAccounts{accounts}
	err := c.Wallet("untag_accounts", &request, nil)

	if err != nil {
		return err
	}

	return nil
}

// SetAccountTagDescription Set a description for a an account given by a supplied tag
func (c *WalletClient) SetAccountTagDescription(tag string, description string) error {
	request := SetAccountTagDescription{tag, description}
	err := c.Wallet("set_account_tag_description", &request, nil)

	if err != nil {
		return err
	}

	return nil
}

// GetHeight will get the currently synced height of the blockchain
func (c *WalletClient) GetHeight() (Height, error) {
	var rep Height
	if err := c.Wallet("getheight", nil, &rep); err != nil {
		return rep, err
	}
	return rep, nil
}

// Transfer will transfer funds to a given address
func (c *WalletClient) Transfer(req TransferInput) (Transfer, error) {
	var rep Transfer
	if err := c.Wallet("transfer", req, &rep); err != nil {
		return rep, err
	}
	return rep, nil
}

// TransferSplit will split a given transaction into multiple transfers
func (c *WalletClient) TransferSplit(req TransferInput) (TransferSplit, error) {
	var rep TransferSplit
	if err := c.Wallet("transfer_split", req, &rep); err != nil {
		return rep, err
	}
	return rep, nil
}

// SignTransfer will sign a given prebuilt transfer
func (c *WalletClient) SignTransfer(req SignTransfer) (SignedTransfer, error) {
	var response SignedTransfer
	err := c.Wallet("sign_transfer", req, &response)

	if err != nil {
		return response, err
	}

	return response, nil
}

// SubmitTransfer will submit a prebuilt and signed transfer to the network
func (c *WalletClient) SubmitTransfer(req SubmitTransfer) ([]string, error) {
	var response SubmittedTransfer
	err := c.Wallet("submit_transfer", req, &response)

	if err != nil {
		return response.TxHashList, err
	}

	return response.TxHashList, nil
}

// SweepDust will sweep dust inputs from the wallet
func (c *WalletClient) SweepDust(req SweepDust) (TransferSplit, error) {
	var response TransferSplit
	if err := c.Wallet("sweep_dust", nil, &response); err != nil {
		return response, err
	}
	return response, nil
}

// SweepAll will sweep all dust matching the given parameters
func (c *WalletClient) SweepAll(req SweepAllDust) (TransferSplit, error) {
	var response TransferSplit
	if err := c.Wallet("sweep_all", nil, &response); err != nil {
		return response, err
	}
	return response, nil
}

// SweepSingle will sweep all dust matching the given parameters
func (c *WalletClient) SweepSingle(req SweepSingle) (Transfer, error) {
	var response Transfer
	if err := c.Wallet("sweep_single", nil, &response); err != nil {
		return response, err
	}
	return response, nil
}

// RelayTx will relay a given transaction to the Monero network
func (c *WalletClient) RelayTx(hexEncodedTx string) (string, error) {
	var response RelayedTransaction
	request := RelayTransaction{hexEncodedTx}

	err := c.Wallet("relay_tx", request, &response)

	if err != nil {
		return response.TxHash, err
	}

	return response.TxHash, nil
}

// Store will save the current wallet state to the wallet file
func (c *WalletClient) Store() error {
	err := c.Wallet("store", nil, nil)

	if err != nil {
		return err
	}

	return nil
}

// GetPayments will fetch all payments to the currently opened wallet
func (c *WalletClient) GetPayments(paymentID string) ([]Payment, error) {
	var rep Payments
	req := GetPayments{paymentID}
	err := c.Wallet("get_payments", req, &rep)

	if err != nil {
		return rep.Payments, err
	}

	return rep.Payments, nil
}

// GetBulkPayments will fetch all payments using a given paymentId
func (c *WalletClient) GetBulkPayments(paymentIds []string, minBlockHeight uint64) ([]Payment, error) {
	req := GetBulkPayments{
		paymentIds,
		minBlockHeight,
	}
	rep := Payments{}
	err := c.Wallet("get_bulk_payments", req, &rep)

	if err != nil {
		return rep.Payments, err
	}

	return rep.Payments, nil
}

// IncomingTransfers will fetch information about incoming transactions
func (c *WalletClient) IncomingTransfers(req IncomingTransfers) ([]TransferDetails, error) {
	var rep IncomingTransfersData
	err := c.Wallet("incoming_transfers", req, &rep)

	if err != nil {
		return rep.Transfers, err
	}

	return rep.Transfers, nil
}

// QueryKey ...
func (c *WalletClient) QueryKey(keyType string) (string, error) {
	req := struct {
		KeyType string `json:"key_type"`
	}{keyType}
	var rep struct {
		Key string `json:"key"`
	}

	err := c.Wallet("query_key", req, &rep)

	if err != nil {
		return rep.Key, err
	}

	return rep.Key, nil
}

// MakeIntegratedAddress will created an integrated address
func (c *WalletClient) MakeIntegratedAddress(req MakeIntegratedAddress) (IntegratedAddress, error) {
	var rep IntegratedAddress
	err := c.Wallet("make_integrated_address", req, &rep)

	if err != nil {
		return rep, err
	}
	return rep, nil
}

// SplitIntegratedAddress will split an integrated up into its parts
func (c *WalletClient) SplitIntegratedAddress(integratedAddress string) (SplitAddress, error) {
	req := struct {
		IntegratedAddress string `json:"integrated_address,omitempty"`
	}{
		integratedAddress,
	}
	var rep SplitAddress
	err := c.Wallet("split_integrated_address", req, &rep)

	if err != nil {
		return rep, err
	}

	return rep, nil
}

// StopWallet will save the currently opened wallet and stop the daemon(?)
func (c *WalletClient) StopWallet() error {
	err := c.Wallet("stop_wallet", nil, nil)
	if err != nil {
		return err
	}
	return nil
}

// RescanBlockchain will rescan the entire blockchain (this takes a long time)
func (c *WalletClient) RescanBlockchain() error {
	err := c.Wallet("rescan_blockchain", nil, nil)
	if err != nil {
		return err
	}
	return nil
}

// SetTransactionNotes add a set of notes for a given set of transaction ids
func (c *WalletClient) SetTransactionNotes(transactionIDs []string, notes []string) error {
	if len(transactionIDs) != len(notes) {
		return errors.New("each transaction id needs to have a note specified")
	}

	req := struct {
		TxIDs []string `json:"txids"`
		Notes []string `json:"notes"`
	}{transactionIDs, notes}

	err := c.Wallet("set_transaction_notes", &req, nil)
	if err != nil {
		return err
	}
	return nil
}

// GetTransactionNotes get notes for a set of given transaction ids
func (c *WalletClient) GetTransactionNotes(transactionIDs []string) ([]string, error) {
	req := struct {
		TxIDs []string `json:"txids"`
	}{transactionIDs}
	var rep struct {
		Notes []string `json:"notes"`
	}
	err := c.Wallet("get_transaction_notes", req, &rep)

	if err != nil {
		return []string{}, err
	}

	return rep.Notes, nil
}

// SetAttribute will set an attribute in the wallet
func (c *WalletClient) SetAttribute(key string, value string) error {
	request := struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	}{key, value}
	err := c.Wallet("set_attribute", request, nil)

	if err != nil {
		return err
	}

	return nil
}

// GetAttribute will set an attribute in the wallet
func (c *WalletClient) GetAttribute(key string) (string, error) {
	request := struct {
		Key string `json:"key"`
	}{key}
	var response struct {
		Value string `json:"value"`
	}
	err := c.Wallet("get_attribute", request, &response)

	if err != nil {
		return response.Value, err
	}

	return response.Value, nil
}

// GetTransactionKey will set an attribute in the wallet
func (c *WalletClient) GetTransactionKey(txid string) (string, error) {
	request := struct {
		TransactionID string `json:"txid"`
	}{txid}
	var response struct {
		Key string `json:"tx_key"`
	}
	err := c.Wallet("get_tx_key", request, &response)

	if err != nil {
		return response.Key, err
	}

	return response.Key, nil
}

// CheckTransactionKey ...
func (c *WalletClient) CheckTransactionKey(txid string, txKey string, address string) (TransactionKey, error) {
	var response TransactionKey
	request := struct {
		TxID    string `json:"txid"`
		TxKey   string `json:"tx_key"`
		Address string `json:"address"`
	}{txid, txKey, address}

	err := c.Wallet("check_tx_key", request, &response)

	if err != nil {
		return response, err
	}

	return response, nil
}

// GetTransactionProof will fetch a proof for the given transaction
func (c *WalletClient) GetTransactionProof(txid string, address string, message string) (string, error) {
	var response struct {
		Signature string `json:"signature"`
	}

	request := struct {
		TxID    string `json:"txid"`
		Address string `json:"address"`
		Message string `json:"message"`
	}{txid, address, message}
	err := c.Wallet("get_tx_proof", request, &response)

	if err != nil {
		return response.Signature, err
	}

	return response.Signature, nil
}

// CheckTransactionProof will return the result of check a given transaction proof
func (c *WalletClient) CheckTransactionProof(req CheckTransactionProof) (CheckedProof, error) {
	var response CheckedProof
	err := c.Wallet("get_tx_proof", req, &response)

	if err != nil {
		return response, err
	}

	return response, nil
}

// GetSpendProof ...
func (c *WalletClient) GetSpendProof(transactionID string, message string) (string, error) {
	var response struct {
		Signature string `json:"signature"`
	}
	request := struct {
		TxID    string `json:"txid"`
		Message string `json:"message"`
	}{transactionID, message}

	err := c.Wallet("get_spend_proof", request, &response)

	if err != nil {
		return response.Signature, err
	}

	return response.Signature, nil
}

// CheckSpendProof ...
func (c *WalletClient) CheckSpendProof(transactionID string, message string, signature string) (bool, error) {
	var response struct {
		Good bool `json:"good"`
	}
	request := struct {
		TxID      string `json:"txid"`
		Signature string `json:"signature"`
		Message   string `json:"message"`
	}{transactionID, message, signature}

	err := c.Wallet("check_spend_proof", request, &response)

	if err != nil {
		return response.Good, err
	}

	return response.Good, nil
}

// GetReserveProof ...
func (c *WalletClient) GetReserveProof(accountIndex uint32, amount uint64, message string, all bool) (string, error) {
	var response struct {
		Signature string `json:"signature"`
	}
	request := struct {
		All          bool   `json:"all"`
		AccountIndex uint32 `json:"account_index"`
		Amount       uint64 `json:"amount"`
		Message      string `json:"message"`
	}{all, accountIndex, amount, message}

	err := c.Wallet("check_spend_proof", request, &response)

	if err != nil {
		return response.Signature, err
	}

	return response.Signature, nil
}

// CheckReserveProof ...
func (c *WalletClient) CheckReserveProof(address string, message string, signature string) (CheckedReserveProof, error) {
	var response CheckedReserveProof
	request := struct {
		Address   string `json:"address"`
		Signature string `json:"signature"`
		Message   string `json:"message"`
	}{address, message, signature}

	err := c.Wallet("check_spend_proof", request, &response)

	if err != nil {
		return response, err
	}

	return response, nil
}

// GetTransfers will fetch transfers involving the current wallet
func (c *WalletClient) GetTransfers(req GetTransfersFilter) (Transfers, error) {
	var rep Transfers
	if err := c.Wallet("get_transfers", req, &rep); err != nil {
		return rep, err
	}
	return rep, nil
}

// GetTransferByTxID will fetch a transaction given by a transaction id
func (c *WalletClient) GetTransferByTxID(txid string) (Transfer, error) {
	req := struct {
		Txid string `json:"txid"`
	}{
		txid,
	}
	var rep struct {
		Trade Transfer `json:"transfer"`
	}
	if err := c.Wallet("get_transfer_by_txid", req, &rep); err != nil {
		return rep.Trade, err
	}
	return rep.Trade, nil
}

// Sign will create a signature for the given data
func (c *WalletClient) Sign(data string) (string, error) {
	req := struct {
		Data string `json:"data"`
	}{
		data,
	}
	var rep struct {
		Signature string `json:"signature"`
	}
	if err := c.Wallet("sign", req, &rep); err != nil {
		return rep.Signature, err
	}
	return rep.Signature, nil
}

// Verify the given signature
func (c *WalletClient) Verify(data string, address string, signature string) (bool, error) {
	req := struct {
		Data      string `json:"data"`
		Address   string `json:"address"`
		Signature string `json:"signature"`
	}{data, address, signature}
	var rep struct {
		Good bool `json:"good"`
	}
	if err := c.Wallet("verify", req, &rep); err != nil {
		return rep.Good, err
	}
	return rep.Good, nil
}

// ExportOutputs ...
func (c *WalletClient) ExportOutputs() (string, error) {
	var rep struct {
		Outputs string `json:"outputs_data_hex"`
	}
	if err := c.Wallet("export_outputs", nil, &rep); err != nil {
		return rep.Outputs, err
	}
	return rep.Outputs, nil
}

// ImportOutputs ...
func (c *WalletClient) ImportOutputs(outputs string) (uint64, error) {
	req := struct {
		Outputs string `json:"outputs_data_hex"`
	}{outputs}
	var rep struct {
		NumImported uint64 `json:"num_imported"`
	}
	if err := c.Wallet("import_outputs", req, &rep); err != nil {
		return rep.NumImported, err
	}
	return rep.NumImported, nil
}

// ExportKeyImages ...
func (c *WalletClient) ExportKeyImages() ([]SignedKeyImage, error) {
	var rep struct {
		SignedKeyImages []SignedKeyImage `json:"signed_key_images"`
	}
	if err := c.Wallet("export_key_images", nil, &rep); err != nil {
		return rep.SignedKeyImages, err
	}
	return rep.SignedKeyImages, nil
}

// ImportKeyImages ...
func (c *WalletClient) ImportKeyImages(images []SignedKeyImage) (ImportedKeyImages, error) {
	req := struct {
		SignedKeyImages []SignedKeyImage `json:"signed_key_images"`
	}{images}
	var rep ImportedKeyImages
	if err := c.Wallet("import_key_images", req, &rep); err != nil {
		return rep, err
	}
	return rep, nil
}

// MakeURI ...
func (c *WalletClient) MakeURI(req URISpec) (string, error) {
	var rep struct {
		URI string `json:"uri"`
	}
	if err := c.Wallet("make_uri", req, &rep); err != nil {
		return rep.URI, err
	}
	return rep.URI, nil
}

// ParseURI ...
func (c *WalletClient) ParseURI(uri string) (ParsedURI, error) {
	var rep ParsedURI
	req := struct {
		URI string `json:"uri"`
	}{uri}
	if err := c.Wallet("parse_uri", req, &rep); err != nil {
		return rep, err
	}
	return rep, nil
}

// GetAddressBook will fetch the current wallets address book
func (c *WalletClient) GetAddressBook(entries []uint) ([]AddressBookEntry, error) {
	req := struct {
		Entries []uint `json:"entries,omitempty"`
	}{
		entries,
	}
	rep := struct {
		Entries []AddressBookEntry `json:"entries,omitempty"`
	}{}
	if err := c.Wallet("get_address_book", req, &rep); err != nil {
		return rep.Entries, err
	}
	return rep.Entries, nil
}

// AddAddressBookEntry will add a given entry to the addressbook
func (c *WalletClient) AddAddressBookEntry(address, paymentID, description string) (uint64, error) {
	req := struct {
		Address     string `json:"address,omitempty"`
		PaymentID   string `json:"payment_id,omitempty"`
		Description string `json:"description,omitempty"`
	}{
		address,
		paymentID,
		description,
	}
	var rep struct {
		Index uint64 `json:"index"`
	}
	if err := c.Wallet("add_address_book", req, &rep); err != nil {
		return rep.Index, err
	}
	return rep.Index, nil
}

// GetAddressBookEntries gets a given entry from the addressbook
func (c *WalletClient) GetAddressBookEntries(entries []uint64) ([]AddressBookEntry, error) {
	req := struct {
		Entries []uint64 `json:"entries"`
	}{entries}
	var rep struct {
		Entries []AddressBookEntry `json:"entries"`
	}
	if err := c.Wallet("get_address_book", req, &rep); err != nil {
		return rep.Entries, err
	}
	return rep.Entries, nil
}

// DeleteAddressBookEntry removes a given entry from the addressbook
func (c *WalletClient) DeleteAddressBookEntry(index uint64) error {
	req := struct {
		Index uint64 `json:"index"`
	}{
		index,
	}
	if err := c.Wallet("delete_address_book", req, nil); err != nil {
		return err
	}
	return nil
}

// RescanSpent will rescan spent inputs
func (c *WalletClient) RescanSpent() error {
	if err := c.Wallet("rescan_spent", nil, nil); err != nil {
		return err
	}
	return nil
}

// Refresh ...
func (c *WalletClient) Refresh(startHeight uint64) (RefreshResult, error) {
	var rep RefreshResult
	req := struct {
		StartHeight uint64 `json:"start_height"`
	}{startHeight}

	if err := c.Wallet("refresh", &req, &rep); err != nil {
		return rep, err
	}
	return rep, nil
}

// StartMining ...
func (c *WalletClient) StartMining(req StartMining) error {
	if err := c.Wallet("start_mining", &req, nil); err != nil {
		return err
	}
	return nil
}

// StopMining ...
func (c *WalletClient) StopMining() error {
	if err := c.Wallet("stop_mining", nil, nil); err != nil {
		return err
	}
	return nil
}

// GetLanguages fetches a list of the available wallet languages
func (c *WalletClient) GetLanguages() ([]string, error) {
	var rep struct {
		Languages []string `json:"languages"`
	}
	if err := c.Wallet("get_languages", nil, &rep); err != nil {
		return rep.Languages, err
	}
	return rep.Languages, nil
}

// CreateWallet will create a new wallet
func (c *WalletClient) CreateWallet(filename string, password string, language string) error {
	req := struct {
		Filename string `json:"filename"`
		Password string `json:"password"`
		Language string `json:"language"`
	}{filename, password, language}
	if err := c.Wallet("create_wallet", req, nil); err != nil {
		return err
	}
	return nil
}

// OpenWallet ...
func (c *WalletClient) OpenWallet(filename string, password string) error {
	req := struct {
		Filename string `json:"filename"`
		Password string `json:"password"`
	}{filename, password}
	if err := c.Wallet("open_wallet", req, nil); err != nil {
		return err
	}
	return nil
}

// CloseWallet ...
func (c *WalletClient) CloseWallet() error {
	if err := c.Wallet("close_wallet", nil, nil); err != nil {
		return err
	}
	return nil
}

// ChangeWalletPassword ...
func (c *WalletClient) ChangeWalletPassword(oldPassword string, newPassword string) error {
	req := struct {
		OldPassword string `json:"old_password"`
		NewPassword string `json:"new_password"`
	}{oldPassword, newPassword}
	if err := c.Wallet("change_wallet_password", req, nil); err != nil {
		return err
	}
	return nil
}

// IsMultisig fetch information whether current wallet is multisig
func (c *WalletClient) IsMultisig() (MultisigInfo, error) {
	var rep MultisigInfo
	if err := c.Wallet("is_multisig", nil, &rep); err != nil {
		return rep, err
	}
	return rep, nil
}

// PrepareMultisig ...
func (c *WalletClient) PrepareMultisig() (string, error) {
	var rep struct {
		MultisigInfo string `json:"multisig_info"`
	}
	if err := c.Wallet("prepare_multisig", nil, &rep); err != nil {
		return rep.MultisigInfo, err
	}
	return rep.MultisigInfo, nil
}

// MakeMultisig ...
func (c *WalletClient) MakeMultisig(multisigInfo []string, threshold uint32, password string) (Multisig, error) {
	var rep Multisig
	req := struct {
		MultisigInfo []string `json:"multisig_info"`
		Threshold    uint32   `json:"threshold"`
		Password     string   `json:"password"`
	}{multisigInfo, threshold, password}
	if err := c.Wallet("make_multisig", req, &rep); err != nil {
		return rep, err
	}
	return rep, nil
}

// ExportMultisigInfo ...
func (c *WalletClient) ExportMultisigInfo() (string, error) {
	var rep struct {
		Info string `json:"info"`
	}
	if err := c.Wallet("export_multisig_info", nil, &rep); err != nil {
		return rep.Info, err
	}
	return rep.Info, nil
}

// ImportMultisigInfo ...
func (c *WalletClient) ImportMultisigInfo(info []string) (uint64, error) {
	var rep struct {
		NumberOutputs uint64 `json:"n_outputs"`
	}
	if err := c.Wallet("import_multisig_info", nil, &rep); err != nil {
		return rep.NumberOutputs, err
	}
	return rep.NumberOutputs, nil
}

// FinalizeMultisig finalize a partially created multisig address
func (c *WalletClient) FinalizeMultisig(password string, multisigInfo []string) (string, error) {
	req := struct {
		Password     string   `json:"password"`
		MultisigInfo []string `json:"multisig_info"`
	}{password, multisigInfo}
	var rep struct {
		Address string `json:"address"`
	}
	if err := c.Wallet("finalize_multisig", req, &rep); err != nil {
		return rep.Address, err
	}
	return rep.Address, nil
}

// ExchangeMultisigKeys ...
func (c *WalletClient) ExchangeMultisigKeys(password string, multisigInfo []string) (Multisig, error) {
	req := struct {
		Password     string   `json:"password"`
		MultisigInfo []string `json:"multisig_info"`
	}{password, multisigInfo}
	var rep Multisig
	if err := c.Wallet("exchange_multisig_keys", req, &rep); err != nil {
		return rep, err
	}
	return rep, nil
}

// SignMultisig will sign a given multisig transaction
func (c *WalletClient) SignMultisig(txDataHex string) (SignedMultisigTransaction, error) {
	req := struct {
		TxDataHex string `json:"tx_data_hex"`
	}{txDataHex}
	var rep SignedMultisigTransaction
	if err := c.Wallet("sign_multisig", req, &rep); err != nil {
		return rep, err
	}
	return rep, nil
}

// SubmitMultisig will submit a given signed multisig transaction to the network
func (c *WalletClient) SubmitMultisig(txDataHex string) ([]string, error) {
	req := struct {
		TxDataHex string `json:"tx_data_hex"`
	}{txDataHex}
	var rep struct {
		TxHashList []string `json:"tx_hash_list"`
	}
	if err := c.Wallet("submit_multisig", req, &rep); err != nil {
		return rep.TxHashList, err
	}
	return rep.TxHashList, nil
}

// GetVersion get current version for wallet rpc daemon
func (c *WalletClient) GetVersion() (uint32, error) {
	var rep struct {
		Version uint32 `json:"version"`
	}
	if err := c.Wallet("get_version", nil, &rep); err != nil {
		return rep.Version, err
	}
	return rep.Version, nil
}
