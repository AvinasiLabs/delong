package schedule

type EventType string

const (
	EVENT_TYPE_RESOLVE = "resolve"
	EVENT_TYPE_RUN     = "run"
)

type SchedulerEvent struct {
	Type        EventType
	ExecutionId uint
	Payload     map[string]any // optional key-value pairs
}
