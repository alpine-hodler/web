package scalar

// Granularity is the level of granularity needed for examining historical data.
type Granularity uint8

const (
	Seconds60 Granularity = iota
	Seconds300
	Seconds900
	Seconds3600
	Seconds21600
	Seoncds86400
)

// Int will convert the historic granularity into an integer
func (granularity Granularity) Int() (i int) {
	switch granularity {
	case Seconds60:
		i = 60
	case Seconds300:
		i = 300
	case Seconds900:
		i = 900
	case Seconds3600:
		i = 3600
	case Seconds21600:
		i = 21600
	case Seoncds86400:
		i = 86400
	}
	return
}
