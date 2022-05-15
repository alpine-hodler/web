package scalar

type ComplianceJob string

const (
	ComplianceJobTweets ComplianceJob = "tweets"
	ComplianceJobUsers  ComplianceJob = "users"
)

// String returns the string value of a ComplianceJob type.
func (complianceJob *ComplianceJob) String() (str string) {
	if complianceJob != nil {
		str = string(*complianceJob)
	}
	return
}
