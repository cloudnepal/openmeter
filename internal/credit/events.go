package credit

import (
	"errors"

	"github.com/openmeterio/openmeter/internal/event/models"
	"github.com/openmeterio/openmeter/internal/event/spec"
)

const (
	EventSubsystem spec.EventSubsystem = "credit"

	EventCreateGrant spec.EventName = "createGrant"
	EventVoidGrant   spec.EventName = "voidGrant"
)

type GrantEvent struct {
	Grant

	SubjectKey models.SubjectKeyAndID `json:"subject"`
	// Namespace from Grant cannot be used as it will never be serialized
	Namespace models.NamespaceID `json:"namespace"`
}

func (g GrantEvent) Validate() error {
	// Basic sanity on grant
	if g.Grant.ID == "" {
		return errors.New("GrantID must be set")
	}

	if g.Grant.OwnerID == "" {
		return errors.New("GrantOwnerID must be set")
	}

	if err := g.SubjectKey.Validate(); err != nil {
		return err
	}

	if err := g.Namespace.Validate(); err != nil {
		return err
	}

	return nil
}

type GrantCreatedEvent GrantEvent

var grantCreatedEventSpec = spec.EventTypeSpec{
	Subsystem:   EventSubsystem,
	Name:        EventCreateGrant,
	SpecVersion: "1.0",
	Version:     "v1",
}

func (e GrantCreatedEvent) Spec() *spec.EventTypeSpec {
	return &grantCreatedEventSpec
}

type GrantVoidedEvent GrantEvent

var grantVoidedEventSpec = spec.EventTypeSpec{
	Subsystem:   EventSubsystem,
	Name:        EventVoidGrant,
	SpecVersion: "1.0",
	Version:     "v1",
}

func (e GrantVoidedEvent) Spec() *spec.EventTypeSpec {
	return &grantVoidedEventSpec
}
