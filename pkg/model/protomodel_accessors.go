package model

// * This is a generated file, do not edit

// Details converts the native protomodel ProtoDetails to a local
// CoinbaseAccountLedgerDetails.
func (m *CoinbaseAccountLedger) Details() *CoinbaseAccountLedgerDetails {
	return &CoinbaseAccountLedgerDetails{CoinbaseAccountLedgerDetails: m.ProtoDetails}
}

// Details converts the native protomodel ProtoDetails to a local
// CoinbaseAccountTransferDetails.
func (m *CoinbaseAccountTransfer) Details() *CoinbaseAccountTransferDetails {
	return &CoinbaseAccountTransferDetails{CoinbaseAccountTransferDetails: m.ProtoDetails}
}

// AddressInfo converts the native protomodel ProtoAddressInfo to a local
// CoinbaseCryptoAddressInfo.
func (m *CoinbaseCryptoAddress) AddressInfo() *CoinbaseCryptoAddressInfo {
	return &CoinbaseCryptoAddressInfo{CoinbaseCryptoAddressInfo: m.ProtoAddressInfo}
}

// Warnings converts the native protomodel ProtoWarnings to a local
// []*CoinbaseCryptoAddressWarning.
func (m *CoinbaseCryptoAddress) Warnings() (slice []*CoinbaseCryptoAddressWarning) {
	for _, v := range m.ProtoWarnings {
		slice = append(slice, &CoinbaseCryptoAddressWarning{CoinbaseCryptoAddressWarning: *v})
	}
	return
}

// Details converts the native protomodel ProtoDetails to a local
// CoinbaseCurrencyDetails.
func (m *CoinbaseCurrency) Details() *CoinbaseCurrencyDetails {
	return &CoinbaseCurrencyDetails{CoinbaseCurrencyDetails: m.ProtoDetails}
}

// Limits converts the native protomodel ProtoLimits to a local CoinbaseLimits.
func (m *CoinbasePaymentMethod) Limits() *CoinbaseLimits {
	return &CoinbaseLimits{CoinbaseLimits: m.ProtoLimits}
}

// FiatAccount converts the native protomodel ProtoFiatAccount to a local
// CoinbaseFiatAccount.
func (m *CoinbasePaymentMethod) FiatAccount() *CoinbaseFiatAccount {
	return &CoinbaseFiatAccount{CoinbaseFiatAccount: m.ProtoFiatAccount}
}

// CryptoAccount converts the native protomodel ProtoCryptoAccount to a local
// CoinbaseCryptoAccount.
func (m *CoinbasePaymentMethod) CryptoAccount() *CoinbaseCryptoAccount {
	return &CoinbaseCryptoAccount{CoinbaseCryptoAccount: m.ProtoCryptoAccount}
}

// RecurringOptions converts the native protomodel ProtoRecurringOptions to a
// local []*CoinbaseRecurringOptions.
func (m *CoinbasePaymentMethod) RecurringOptions() (slice []*CoinbaseRecurringOptions) {
	for _, v := range m.ProtoRecurringOptions {
		slice = append(slice, &CoinbaseRecurringOptions{CoinbaseRecurringOptions: *v})
	}
	return
}

// AvailableBalance converts the native protomodel ProtoAvailableBalance to a
// local CoinbaseAvailableBalance.
func (m *CoinbasePaymentMethod) AvailableBalance() *CoinbaseAvailableBalance {
	return &CoinbaseAvailableBalance{CoinbaseAvailableBalance: m.ProtoAvailableBalance}
}

// PickerData converts the native protomodel ProtoPickerData to a local
// CoinbasePickerData.
func (m *CoinbasePaymentMethod) PickerData() *CoinbasePickerData {
	return &CoinbasePickerData{CoinbasePickerData: m.ProtoPickerData}
}

// Balance converts the native protomodel ProtoBalance to a local
// CoinbaseBalance.
func (m *CoinbasePickerData) Balance() *CoinbaseBalance {
	return &CoinbaseBalance{CoinbaseBalance: m.ProtoBalance}
}

// BankCountry converts the native protomodel ProtoBankCountry to a local
// CoinbaseBankCountry.
func (m *CoinbaseSepaDepositInformation) BankCountry() *CoinbaseBankCountry {
	return &CoinbaseBankCountry{CoinbaseBankCountry: m.ProtoBankCountry}
}

// BankCountry converts the native protomodel ProtoBankCountry to a local
// CoinbaseBankCountry.
func (m *CoinbaseSwiftDepositInformation) BankCountry() *CoinbaseBankCountry {
	return &CoinbaseBankCountry{CoinbaseBankCountry: m.ProtoBankCountry}
}

// BankCountry converts the native protomodel ProtoBankCountry to a local
// CoinbaseBankCountry.
func (m *CoinbaseUkDepositInformation) BankCountry() *CoinbaseBankCountry {
	return &CoinbaseBankCountry{CoinbaseBankCountry: m.ProtoBankCountry}
}

// WireDepositInformation converts the native protomodel
// ProtoWireDepositInformation to a local CoinbaseWireDepositInformation.
func (m *CoinbaseWallet) WireDepositInformation() *CoinbaseWireDepositInformation {
	return &CoinbaseWireDepositInformation{CoinbaseWireDepositInformation: m.ProtoWireDepositInformation}
}

// SwiftDepositInformation converts the native protomodel
// ProtoSwiftDepositInformation to a local CoinbaseSwiftDepositInformation.
func (m *CoinbaseWallet) SwiftDepositInformation() *CoinbaseSwiftDepositInformation {
	return &CoinbaseSwiftDepositInformation{CoinbaseSwiftDepositInformation: m.ProtoSwiftDepositInformation}
}

// SepaDepositInformation converts the native protomodel
// ProtoSepaDepositInformation to a local CoinbaseSepaDepositInformation.
func (m *CoinbaseWallet) SepaDepositInformation() *CoinbaseSepaDepositInformation {
	return &CoinbaseSepaDepositInformation{CoinbaseSepaDepositInformation: m.ProtoSepaDepositInformation}
}

// UkDepositInformation converts the native protomodel ProtoUkDepositInformation
// to a local CoinbaseUkDepositInformation.
func (m *CoinbaseWallet) UkDepositInformation() *CoinbaseUkDepositInformation {
	return &CoinbaseUkDepositInformation{CoinbaseUkDepositInformation: m.ProtoUkDepositInformation}
}

// BankCountry converts the native protomodel ProtoBankCountry to a local
// CoinbaseBankCountry.
func (m *CoinbaseWireDepositInformation) BankCountry() *CoinbaseBankCountry {
	return &CoinbaseBankCountry{CoinbaseBankCountry: m.ProtoBankCountry}
}

// Schema converts the native protomodel ProtoSchema to a local
// []*IexRulesScheme.
func (m *IexRulesSchema) Schema() (slice []*IexRulesScheme) {
	for _, v := range m.ProtoSchema {
		slice = append(slice, &IexRulesScheme{IexRulesScheme: *v})
	}
	return
}

// Result converts the native protomodel ProtoResult to a local
// KrakenServerTimeResult.
func (m *KrakenServerTime) Result() *KrakenServerTimeResult {
	return &KrakenServerTimeResult{KrakenServerTimeResult: m.ProtoResult}
}

// Result converts the native protomodel ProtoResult to a local
// KrakenSystemStatusResult.
func (m *KrakenSystemStatus) Result() *KrakenSystemStatusResult {
	return &KrakenSystemStatusResult{KrakenSystemStatusResult: m.ProtoResult}
}

// AssetContract converts the native protomodel ProtoAssetContract to a local
// OpenseaAssetContract.
func (m *OpenseaAsset) AssetContract() *OpenseaAssetContract {
	return &OpenseaAssetContract{OpenseaAssetContract: m.ProtoAssetContract}
}

// Collection converts the native protomodel ProtoCollection to a local
// OpenseaCollection.
func (m *OpenseaAsset) Collection() *OpenseaCollection {
	return &OpenseaCollection{OpenseaCollection: m.ProtoCollection}
}

// Owner converts the native protomodel ProtoOwner to a local OpenseaOwner.
func (m *OpenseaAsset) Owner() *OpenseaOwner { return &OpenseaOwner{OpenseaOwner: m.ProtoOwner} }

// Creator converts the native protomodel ProtoCreator to a local
// OpenseaCreator.
func (m *OpenseaAsset) Creator() *OpenseaCreator {
	return &OpenseaCreator{OpenseaCreator: m.ProtoCreator}
}

// Assets converts the native protomodel ProtoAssets to a local []*OpenseaAsset.
func (m *OpenseaAssets) Assets() (slice []*OpenseaAsset) {
	for _, v := range m.ProtoAssets {
		slice = append(slice, &OpenseaAsset{OpenseaAsset: *v})
	}
	return
}

// DisplayData converts the native protomodel ProtoDisplayData to a local
// OpenseaDisplayData.
func (m *OpenseaCollection) DisplayData() *OpenseaDisplayData {
	return &OpenseaDisplayData{OpenseaDisplayData: m.ProtoDisplayData}
}

// User converts the native protomodel ProtoUser to a local OpenseaUser.
func (m *OpenseaCreator) User() *OpenseaUser { return &OpenseaUser{OpenseaUser: m.ProtoUser} }

// User converts the native protomodel ProtoUser to a local OpenseaUser.
func (m *OpenseaOwner) User() *OpenseaUser { return &OpenseaUser{OpenseaUser: m.ProtoUser} }
