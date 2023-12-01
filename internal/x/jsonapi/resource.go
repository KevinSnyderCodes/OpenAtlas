package jsonapi

type Resource[Attributes any] struct {
	ID   string `json:"id,omitempty"`
	Type string `json:"type,omitempty"`
	// TODO: Support lid

	Attributes    Attributes    `json:"attributes,omitempty"`
	Relationships Relationships `json:"relationships,omitempty"`
	Links         Links         `json:"links,omitempty"`
	Meta          Meta          `json:"meta,omitempty"`
}

func (o *Resource[Attributes]) Validate() error {
	if o.ID == "" {
		return ErrResourceMustContainID
	}
	if o.Type == "" {
		return ErrResourceMustContainType
	}

	if o.Relationships != nil {
		if err := o.Relationships.Validate(); err != nil {
			return err
		}
	}
	if o.Links != nil {
		if err := o.Links.Validate(); err != nil {
			return err
		}
	}
	if o.Meta != nil {
		if err := o.Meta.Validate(); err != nil {
			return err
		}
	}

	return nil
}
