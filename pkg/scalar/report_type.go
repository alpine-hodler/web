package scalar

type ReportType string

const (
	ReportTypeFills                   ReportType = "fills"
	ReportTypeAcccount                ReportType = "account"
	ReportTypeOTCFills                ReportType = "otc_fills"
	ReportType1099kTransactionHistory ReportType = "type_1099k_transaction_history"
	ReportTypeTaxInvoice              ReportType = "tax_invoice"
)

// String converts a pointer reference to ReportType to a string
func (reportType *ReportType) String() (str string) {
	if reportType != nil {
		str = string(*reportType)
	}
	return
}
