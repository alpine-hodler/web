package scalar

type Status string

const (
	StatusOnline     Status = "online"
	StatusOffline    Status = "offline"
	StatusInternal   Status = "internal"
	StatusDelisted   Status = "delisted"
	StatusPending    Status = "pending"
	StatusCreating   Status = "creating"
	StatusRead       Status = "ready"
	StatusCreated    Status = "created"
	StatusInProgress Status = "in_progress"
	StatusFailed     Status = "failed"
	STatusComplete   Status = "complete"
)

// String converts a pointer reference to Status to a string
func (status *Status) String() (str string) {
	if status != nil {
		str = string(*status)
	}
	return
}
