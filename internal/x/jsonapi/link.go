package jsonapi

type Links map[string]any

func (o Links) Validate() error {
	for _, v := range o {
		switch v.(type) {
		case string:
			// TODO: Validate URI-reference [RFC3986 Section 4.1]
		case *LinkObject:
			// TODO: Validate
		case nil:
		default:
			return ErrLinkInvalidType
		}
	}

	return nil
}

type Link interface {
	string | *LinkObject
}

type LinkObject struct {
	Href string `json:"href,omitempty"`

	Rel         string `json:"rel,omitempty"`
	DescribedBy string `json:"describedby,omitempty"`
	Title       string `json:"title,omitempty"`
	Type        string `json:"type,omitempty"`
	HrefLang    any    `json:"hreflang,omitempty"`
	Meta        any    `json:"meta,omitempty"`
}

func (o *LinkObject) Validate() error {
	if o.Href == "" {
		return ErrLinkObjectMustContainHref
	}
	if o.Rel != "" {
		if err := validateURIReference(o.Rel); err != nil {
			return err
		}
	}
	if o.DescribedBy != "" {
		// TODO: Validate
	}
	if o.HrefLang != nil {
		switch v := o.HrefLang.(type) {
		case string:
			if err := validateLanguageTag(v); err != nil {
				return err
			}
		case []string:
			for _, vv := range v {
				if err := validateLanguageTag(vv); err != nil {
					return err
				}
			}
		default:
			return ErrLinkObjectHrefLangInvalidType
		}
	}

	return nil
}
