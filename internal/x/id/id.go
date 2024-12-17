package id

import (
	"fmt"
	"strings"
)

func parseExternalID(id string) (string, string, error) {
	ss := strings.Split(id, "-")
	if len(ss) != 2 {
		return "", "", fmt.Errorf("invalid format")
	}

	return ss[0], ss[1], nil
}

type id struct {
	prefix string

	value struct {
		prefix string
		id     string
	}
}

func newIDWithPrefix(prefix string) id {
	return id{
		prefix: prefix,
	}
}

func newIDWithPrefixFromInternalID(prefix, id string) (id, error) {
	v := newIDWithPrefix(prefix)

	v.value.prefix = prefix
	v.value.id = id

	if err := v.validate(); err != nil {
		return v, fmt.Errorf("error validating: %w", err)
	}

	return v, nil
}

func newIDWithPrefixFromExternalID(prefix, id string) (id, error) {
	v := newIDWithPrefix(prefix)

	{
		prefix, id, err := parseExternalID(id)
		if err != nil {
			return v, fmt.Errorf("error parsing external id: %w", err)
		}

		v.value.prefix = prefix
		v.value.id = id
	}

	if err := v.validate(); err != nil {
		return v, fmt.Errorf("error validating: %w", err)
	}

	return v, nil
}

func (o id) validate() error {
	if o.value.prefix != o.prefix {
		return fmt.Errorf("invalid id prefix: %s", o.value.prefix)
	}
	if len(o.value.id) != 16 {
		return fmt.Errorf("invalid id: %s", o.value.id)
	}

	return nil
}

func (o id) InternalID() string {
	return o.value.id
}

func (o id) ExternalID() string {
	return fmt.Sprintf("%s-%s", o.prefix, o.value.id)
}

func makeIDTypeFromInternalID[T ~struct{ id }](prefix, id string) (T, error) {
	vv, err := newIDWithPrefixFromInternalID(prefix, id)
	if err != nil {
		return T{}, nil
	}

	return T{vv}, nil
}

func makeIDTypeFromExternalID[T ~struct{ id }](prefix, id string) (T, error) {
	vv, err := newIDWithPrefixFromExternalID(prefix, id)
	if err != nil {
		return T{}, nil
	}

	return T{vv}, nil
}

type ConfigurationVersionID struct{ id }

func NewConfigurationVersionIDFromInternalID(id string) (ConfigurationVersionID, error) {
	return makeIDTypeFromInternalID[ConfigurationVersionID]("cv", id)
}

func NewConfigurationVersionIDFromExternalID(id string) (ConfigurationVersionID, error) {
	return makeIDTypeFromExternalID[ConfigurationVersionID]("cv", id)
}

type PlanID struct{ id }

func NewPlanIDFromInternalID(id string) (PlanID, error) {
	return makeIDTypeFromInternalID[PlanID]("plan", id)
}

func NewPlanIDFromExternalID(id string) (PlanID, error) {
	return makeIDTypeFromExternalID[PlanID]("plan", id)
}

type RunID struct{ id }

func NewRunIDFromInternalID(id string) (RunID, error) {
	return makeIDTypeFromInternalID[RunID]("run", id)
}

func NewRunIDFromExternalID(id string) (RunID, error) {
	return makeIDTypeFromExternalID[RunID]("run", id)
}

type WorkspaceID struct{ id }

func NewWorkspaceIDFromInternalID(id string) (WorkspaceID, error) {
	return makeIDTypeFromInternalID[WorkspaceID]("ws", id)
}

func NewWorkspaceIDFromExternalID(id string) (WorkspaceID, error) {
	return makeIDTypeFromExternalID[WorkspaceID]("ws", id)
}
