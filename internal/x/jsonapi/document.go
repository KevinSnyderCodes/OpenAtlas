package jsonapi

// TODO: Replace all instances of `any` with strong types

type Document[Data *Resource[*Attributes] | []*Resource[*Attributes], Attributes any] struct {
	Data   Data     `json:"data,omitempty"`
	Errors []*Error `json:"errors,omitempty"`
	Meta   Meta     `json:"meta,omitempty"`

	JSONAPI  any   `json:"jsonapi,omitempty"`
	Links    Links `json:"links,omitempty"`
	Included any   `json:"included,omitempty"`
}

func (o *Document[Data, Attributes]) Validate() error {
	if o.Data == nil && o.Errors == nil && o.Meta == nil {
		return ErrDocumentMustContainTopLevelMember
	}
	if o.Data != nil && o.Errors != nil {
		return ErrDocumentMustNotContainBothDataAndErrors
	}
	if o.Data == nil && o.Included != nil {
		return ErrDocumentMustNotContainIncludedWithoutData
	}

	if o.Links != nil {
		if err := o.Links.Validate(); err != nil {
			return err
		}

		for k, _ := range o.Links {
			switch k {
			case "self":
				// TODO: Validate
			case "related":
				// TODO: Validate
			case "describedby":
				// TODO: Validate
			case "first":
				// TODO: Validate
			case "last":
				// TODO: Validate
			case "prev":
				// TODO: Validate
			case "next":
				// TODO: Validate
			}
		}
	}

	switch v := any(o.Data).(type) {
	case *Resource[*Attributes]:
		if err := v.Validate(); err != nil {
			return err
		}
	case []*Resource[*Attributes]:
		for _, vv := range v {
			if err := vv.Validate(); err != nil {
				return err
			}
		}
	case nil:
	default:
		return ErrDocumentDataInvalidType
	}

	return nil
}
