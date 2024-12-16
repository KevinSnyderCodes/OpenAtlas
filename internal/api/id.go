package api

import (
	"fmt"
	"math/rand"
	"strings"
)

var _ WeakInternalID = (*InternalID)(nil)
var _ WeakExternalID = (*ExternalID)(nil)

var _ StrongInternalID = (*ConfigurationVersionInternalID)(nil)
var _ StrongExternalID = (*ConfigurationVersionExternalID)(nil)
var _ StrongInternalID = (*PlanInternalID)(nil)
var _ StrongExternalID = (*PlanExternalID)(nil)
var _ StrongInternalID = (*RunInternalID)(nil)
var _ StrongExternalID = (*RunExternalID)(nil)

const idBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func GenerateID() string {
	b := make([]byte, 16)
	for i := range b {
		b[i] = idBytes[rand.Intn(len(idBytes))]
	}
	return string(b)
}

type Validator interface {
	Validate() error
}

type WeakInternalID interface {
	fmt.Stringer
	Validator

	ExternalID(prefix string) WeakExternalID
}

type WeakExternalID interface {
	fmt.Stringer
	Validator

	Prefix() string
	InternalID() WeakInternalID
}

type StrongInternalID interface {
	fmt.Stringer
	Validator

	ExternalID() StrongExternalID
}

type StrongExternalID interface {
	fmt.Stringer
	Validator

	Prefix() string
	InternalID() StrongInternalID
}

type ExternalID string

func (o ExternalID) Validate() error {
	if o == "" {
		return fmt.Errorf("id must not be empty")
	}

	ss := strings.Split(string(o), "-")
	if len(ss) != 2 {
		return fmt.Errorf("invalid format")
	}

	if err := o.InternalID().Validate(); err != nil {
		return err
	}

	return nil
}

func (o ExternalID) Prefix() string {
	ss := strings.Split(string(o), "-")
	if len(ss) != 2 {
		return ""
	}
	return ss[0]
}

func (o ExternalID) InternalID() WeakInternalID {
	ss := strings.Split(string(o), "-")
	if len(ss) != 2 {
		return nil
	}
	return InternalID(ss[1])
}

func (o ExternalID) String() string {
	return string(o)
}

type InternalID string

func (o InternalID) Validate() error {
	if o == "" {
		return fmt.Errorf("id must not be empty")
	}

	if len(o) != 16 {
		return fmt.Errorf("invalid id: %s", o)
	}

	return nil
}

func (o InternalID) ExternalID(prefix string) WeakExternalID {
	return ExternalID(fmt.Sprintf("%s-%s", prefix, o))
}

func (o InternalID) String() string {
	return string(o)
}

type ConfigurationVersionInternalID InternalID

func (o ConfigurationVersionInternalID) String() string {
	return InternalID(o).String()
}

func (o ConfigurationVersionInternalID) Validate() error {
	internalID := InternalID(o)

	if err := internalID.Validate(); err != nil {
		return err
	}

	return nil
}

func (o ConfigurationVersionInternalID) ExternalID() StrongExternalID {
	return ConfigurationVersionExternalID(InternalID(o).ExternalID("cv").String())
}

type ConfigurationVersionExternalID ExternalID

func (o ConfigurationVersionExternalID) String() string {
	return ExternalID(o).String()
}

func (o ConfigurationVersionExternalID) Validate() error {
	externalID := ExternalID(o)

	if err := externalID.Validate(); err != nil {
		return err
	}
	if externalID.Prefix() != "cv" {
		return fmt.Errorf("invalid configuration version id prefix: %s", externalID.Prefix())
	}

	return nil
}

func (o ConfigurationVersionExternalID) Prefix() string {
	return ExternalID(o).Prefix()
}

func (o ConfigurationVersionExternalID) InternalID() StrongInternalID {
	return ConfigurationVersionInternalID(ExternalID(o).InternalID().String())
}

type PlanInternalID InternalID

func (o PlanInternalID) String() string {
	return InternalID(o).String()
}

func (o PlanInternalID) Validate() error {
	internalID := InternalID(o)

	if err := internalID.Validate(); err != nil {
		return err
	}

	return nil
}

func (o PlanInternalID) ExternalID() StrongExternalID {
	return PlanExternalID(InternalID(o).ExternalID("plan").String())
}

type PlanExternalID ExternalID

func (o PlanExternalID) String() string {
	return ExternalID(o).String()
}

func (o PlanExternalID) Validate() error {
	externalID := ExternalID(o)

	if err := externalID.Validate(); err != nil {
		return err
	}
	if externalID.Prefix() != "plan" {
		return fmt.Errorf("invalid plan id prefix: %s", externalID.Prefix())
	}

	return nil
}

func (o PlanExternalID) Prefix() string {
	return ExternalID(o).Prefix()
}

func (o PlanExternalID) InternalID() StrongInternalID {
	return PlanInternalID(ExternalID(o).InternalID().String())
}

type RunInternalID InternalID

func (o RunInternalID) String() string {
	return InternalID(o).String()
}

func (o RunInternalID) Validate() error {
	internalID := InternalID(o)

	if err := internalID.Validate(); err != nil {
		return err
	}

	return nil
}

func (o RunInternalID) ExternalID() StrongExternalID {
	return RunExternalID(InternalID(o).ExternalID("run").String())
}

type RunExternalID ExternalID

func (o RunExternalID) String() string {
	return ExternalID(o).String()
}

func (o RunExternalID) Validate() error {
	externalID := ExternalID(o)

	if err := externalID.Validate(); err != nil {
		return err
	}
	if externalID.Prefix() != "run" {
		return fmt.Errorf("invalid run id prefix: %s", externalID.Prefix())
	}

	return nil
}

func (o RunExternalID) Prefix() string {
	return ExternalID(o).Prefix()
}

func (o RunExternalID) InternalID() StrongInternalID {
	return RunInternalID(ExternalID(o).InternalID().String())
}

type WorkspaceExternalID ExternalID

func (o WorkspaceExternalID) String() string {
	return ExternalID(o).String()
}

func (o WorkspaceExternalID) Validate() error {
	externalID := ExternalID(o)

	if err := externalID.Validate(); err != nil {
		return err
	}
	if externalID.Prefix() != "ws" {
		return fmt.Errorf("invalid workspace id prefix: %s", externalID.Prefix())
	}

	return nil
}

func (o WorkspaceExternalID) Prefix() string {
	return ExternalID(o).Prefix()
}

func (o WorkspaceExternalID) InternalID() StrongInternalID {
	return WorkspaceInternalID(ExternalID(o).InternalID().String())
}

type WorkspaceInternalID InternalID

func (o WorkspaceInternalID) String() string {
	return InternalID(o).String()
}

func (o WorkspaceInternalID) Validate() error {
	internalID := InternalID(o)

	if err := internalID.Validate(); err != nil {
		return err
	}

	return nil
}

func (o WorkspaceInternalID) ExternalID() StrongExternalID {
	return WorkspaceExternalID(InternalID(o).ExternalID("ws").String())
}
