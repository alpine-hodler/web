package scalar

type TransferMethod string

const (
	TransferMethodDeposit  TransferMethod = "deposit"
	TransferMethodWithdraw TransferMethod = "withdraw"
)

func (transferMethod *TransferMethod) String() (str string) {
	if transferMethod != nil {
		str = string(*transferMethod)
	}
	return
}
