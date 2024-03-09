package models

type TimelineEntry interface {
	Type() TimelineEntryType
}

type TimelineEntryType string

const (
	TimelineEntryTypeJob       TimelineEntryType = "job"
	TimelineEntryTypeEducation TimelineEntryType = "education"
)

type Job struct {
	Start        string
	End          string
	Company      string
	Title        string
	ChangeReason ChangeReason
	Toolbox      Toolbox
	Details      []string
}

var _ TimelineEntry = (*Job)(nil)

func (j Job) Type() TimelineEntryType {
	return TimelineEntryTypeJob
}

type ChangeReason string

const (
	ChangeReasonNew       ChangeReason = "new"
	ChangeReasonPromotion ChangeReason = "promotion"
)

type Toolbox struct {
	Plan    []string
	Code    []string
	Model   []string
	View    []string
	Build   []string
	Run     []string
	Persist []string
	Move    []string
}

type Education struct {
	Start  string
	End    string
	School string
	Degree string
	Field  string
}

var _ TimelineEntry = (*Education)(nil)

func (e Education) Type() TimelineEntryType {
	return TimelineEntryTypeEducation
}

type Project struct {
	Name        string
	Description string
	URL         string
	Languages   []string
}
