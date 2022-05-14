package scalar

type Format string

const (
	FormatPDF Format = "pdf"
	FormatCSV Format = "csv"
)

// String converts a pointer reference to Format to a string
func (format *Format) String() (str string) {
	if format != nil {
		str = string(*format)
	}
	return
}
