package jsonapi

// https://jsonapi.org/format/#document-resource-object-relationships
type Relationships map[string]*Relationship

func (o Relationships) Validate() error {
	for _, v := range o {
		if err := v.Validate(); err != nil {
			return err
		}
	}

	return nil
}

// https://jsonapi.org/format/#document-resource-object-relationships
type Relationship struct {
	Links Links `json:"links,omitempty"`
	Data  any   `json:"data,omitempty"`
	Meta  Meta  `json:"meta,omitempty"`
}

func (o *Relationship) Validate() error {
	if o.Links == nil && o.Data == nil && o.Meta == nil {
		return ErrRelationshipMustContainLinksDataOrMeta
	}

	if o.Links != nil {
		if err := o.Links.Validate(); err != nil {
			return err
		}
	}
	// TODO: Validate data
	if o.Meta != nil {
		if err := o.Meta.Validate(); err != nil {
			return err
		}
	}

	return nil
}
