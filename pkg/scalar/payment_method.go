package scalar

// PaymentMethod scalar represents the different kinds of payment methods that can be used when
// buying and selling bitcoin, bitcoin cash, litecoin or ethereum.

// As fiat accounts can be used for buying and selling, they have an associated payment method. This
// type of a payment method will also have a fiat_account reference to the actual account.
type PaymentMethod string

const (
	// PaymentMethodACHBankAccount is the PM for a regular US bank account
	PaymentMethodACHBankAccount PaymentMethod = "ach_bank_account"

	// PaymentMethodBankWire is the PM for a bank wire (US only)
	PaymentMethodBankWire PaymentMethod = "bank_wire"

	// PaymentMethodCreditCard is the PM for a credit card (can’t be used for buying/selling)
	PaymentMethodCreditCard PaymentMethod = "credit_card"

	// PaymentMethodCreditCard is the PM for cryptocurrency transactions
	PaymentMethodCrypto PaymentMethod = "crypto"

	// PaymentMethodEFTBankAccount is the PM for a Canadian EFT bank account
	PaymentMethodEFTBankAccount PaymentMethod = "eft_bank_account"

	PaymentMethodFedwire PaymentMethod = "fedwire"

	// PaymentMethodFiatAccount	is the PM for a fiat nominated Coinbase account
	PaymentMethodFiatAccount PaymentMethod = "fiat_account"

	// PaymentMethodInterac	is the PM for an Interac Online for Canadian bank accounts
	PaymentMethodInterac PaymentMethod = "interac"

	PaymentMethodIntraBankAccount PaymentMethod = "intra_bank_account"

	// PaymentMethodSEPABankAccount is the PM for a European SEPA bank account
	PaymentMethodSEPABankAccount PaymentMethod = "sepa_bank_account"

	// PaymentMethodSecure3D is the PM for a Secure3D verified payment card
	PaymentMethodSecure3D PaymentMethod = "secure3d_card"

	PaymentMethodSWIFTBankAccount PaymentMethod = "swift_bank_account"

	// PaymentMethodiDealBankAccount is the PM for a iDeal bank account (Europe)
	PaymentMethodiDealBankAccount PaymentMethod = "ideal_bank_account"
)
