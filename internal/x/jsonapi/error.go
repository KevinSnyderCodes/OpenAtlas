package jsonapi

type Error struct {
	ID     string       `json:"id,omitempty"`
	Links  Links        `json:"links,omitempty"`
	Status string       `json:"status,omitempty"`
	Code   string       `json:"code,omitempty"`
	Title  string       `json:"title,omitempty"`
	Detail string       `json:"detail,omitempty"`
	Source *ErrorSource `json:"source,omitempty"`
	Meta   Meta         `json:"meta,omitempty"`
}

func (o *Error) Validate() error {
	if o.ID == "" &&
		o.Links == nil &&
		o.Status == "" &&
		o.Code == "" &&
		o.Title == "" &&
		o.Detail == "" &&
		o.Source == nil &&
		o.Meta == nil {
		return ErrErrorMustContainTopLevelMember
	}

	return nil
}

type ErrorSource struct {
	Pointer   string `json:"pointer,omitempty"`
	Parameter string `json:"parameter,omitempty"`
	Header    string `json:"header,omitempty"`
}
