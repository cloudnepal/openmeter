// Code generated by ent, DO NOT EDIT.

package db

import (
	"time"

	"github.com/openmeterio/openmeter/internal/credit/postgres_connector/ent/db/creditentry"
	"github.com/openmeterio/openmeter/internal/credit/postgres_connector/ent/db/feature"
	"github.com/openmeterio/openmeter/internal/credit/postgres_connector/ent/db/ledger"
	"github.com/openmeterio/openmeter/internal/credit/postgres_connector/ent/schema"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	creditentryMixin := schema.CreditEntry{}.Mixin()
	creditentryMixinFields0 := creditentryMixin[0].Fields()
	_ = creditentryMixinFields0
	creditentryMixinFields1 := creditentryMixin[1].Fields()
	_ = creditentryMixinFields1
	creditentryFields := schema.CreditEntry{}.Fields()
	_ = creditentryFields
	// creditentryDescCreatedAt is the schema descriptor for created_at field.
	creditentryDescCreatedAt := creditentryMixinFields1[0].Descriptor()
	// creditentry.DefaultCreatedAt holds the default value on creation for the created_at field.
	creditentry.DefaultCreatedAt = creditentryDescCreatedAt.Default.(func() time.Time)
	// creditentryDescUpdatedAt is the schema descriptor for updated_at field.
	creditentryDescUpdatedAt := creditentryMixinFields1[1].Descriptor()
	// creditentry.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	creditentry.DefaultUpdatedAt = creditentryDescUpdatedAt.Default.(func() time.Time)
	// creditentry.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	creditentry.UpdateDefaultUpdatedAt = creditentryDescUpdatedAt.UpdateDefault.(func() time.Time)
	// creditentryDescNamespace is the schema descriptor for namespace field.
	creditentryDescNamespace := creditentryFields[0].Descriptor()
	// creditentry.NamespaceValidator is a validator for the "namespace" field. It is called by the builders before save.
	creditentry.NamespaceValidator = creditentryDescNamespace.Validators[0].(func(string) error)
	// creditentryDescPriority is the schema descriptor for priority field.
	creditentryDescPriority := creditentryFields[6].Descriptor()
	// creditentry.DefaultPriority holds the default value on creation for the priority field.
	creditentry.DefaultPriority = creditentryDescPriority.Default.(uint8)
	// creditentryDescEffectiveAt is the schema descriptor for effective_at field.
	creditentryDescEffectiveAt := creditentryFields[7].Descriptor()
	// creditentry.DefaultEffectiveAt holds the default value on creation for the effective_at field.
	creditentry.DefaultEffectiveAt = creditentryDescEffectiveAt.Default.(func() time.Time)
	// creditentryDescID is the schema descriptor for id field.
	creditentryDescID := creditentryMixinFields0[0].Descriptor()
	// creditentry.DefaultID holds the default value on creation for the id field.
	creditentry.DefaultID = creditentryDescID.Default.(func() string)
	featureMixin := schema.Feature{}.Mixin()
	featureMixinFields0 := featureMixin[0].Fields()
	_ = featureMixinFields0
	featureMixinFields1 := featureMixin[1].Fields()
	_ = featureMixinFields1
	featureFields := schema.Feature{}.Fields()
	_ = featureFields
	// featureDescCreatedAt is the schema descriptor for created_at field.
	featureDescCreatedAt := featureMixinFields1[0].Descriptor()
	// feature.DefaultCreatedAt holds the default value on creation for the created_at field.
	feature.DefaultCreatedAt = featureDescCreatedAt.Default.(func() time.Time)
	// featureDescUpdatedAt is the schema descriptor for updated_at field.
	featureDescUpdatedAt := featureMixinFields1[1].Descriptor()
	// feature.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	feature.DefaultUpdatedAt = featureDescUpdatedAt.Default.(func() time.Time)
	// feature.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	feature.UpdateDefaultUpdatedAt = featureDescUpdatedAt.UpdateDefault.(func() time.Time)
	// featureDescNamespace is the schema descriptor for namespace field.
	featureDescNamespace := featureFields[0].Descriptor()
	// feature.NamespaceValidator is a validator for the "namespace" field. It is called by the builders before save.
	feature.NamespaceValidator = featureDescNamespace.Validators[0].(func(string) error)
	// featureDescName is the schema descriptor for name field.
	featureDescName := featureFields[1].Descriptor()
	// feature.NameValidator is a validator for the "name" field. It is called by the builders before save.
	feature.NameValidator = featureDescName.Validators[0].(func(string) error)
	// featureDescMeterSlug is the schema descriptor for meter_slug field.
	featureDescMeterSlug := featureFields[2].Descriptor()
	// feature.MeterSlugValidator is a validator for the "meter_slug" field. It is called by the builders before save.
	feature.MeterSlugValidator = featureDescMeterSlug.Validators[0].(func(string) error)
	// featureDescArchived is the schema descriptor for archived field.
	featureDescArchived := featureFields[4].Descriptor()
	// feature.DefaultArchived holds the default value on creation for the archived field.
	feature.DefaultArchived = featureDescArchived.Default.(bool)
	// featureDescID is the schema descriptor for id field.
	featureDescID := featureMixinFields0[0].Descriptor()
	// feature.DefaultID holds the default value on creation for the id field.
	feature.DefaultID = featureDescID.Default.(func() string)
	ledgerMixin := schema.Ledger{}.Mixin()
	ledgerMixinFields0 := ledgerMixin[0].Fields()
	_ = ledgerMixinFields0
	ledgerMixinFields1 := ledgerMixin[1].Fields()
	_ = ledgerMixinFields1
	ledgerFields := schema.Ledger{}.Fields()
	_ = ledgerFields
	// ledgerDescCreatedAt is the schema descriptor for created_at field.
	ledgerDescCreatedAt := ledgerMixinFields0[0].Descriptor()
	// ledger.DefaultCreatedAt holds the default value on creation for the created_at field.
	ledger.DefaultCreatedAt = ledgerDescCreatedAt.Default.(func() time.Time)
	// ledgerDescUpdatedAt is the schema descriptor for updated_at field.
	ledgerDescUpdatedAt := ledgerMixinFields0[1].Descriptor()
	// ledger.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	ledger.DefaultUpdatedAt = ledgerDescUpdatedAt.Default.(func() time.Time)
	// ledger.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	ledger.UpdateDefaultUpdatedAt = ledgerDescUpdatedAt.UpdateDefault.(func() time.Time)
	// ledgerDescNamespace is the schema descriptor for namespace field.
	ledgerDescNamespace := ledgerFields[0].Descriptor()
	// ledger.NamespaceValidator is a validator for the "namespace" field. It is called by the builders before save.
	ledger.NamespaceValidator = ledgerDescNamespace.Validators[0].(func(string) error)
	// ledgerDescSubject is the schema descriptor for subject field.
	ledgerDescSubject := ledgerFields[1].Descriptor()
	// ledger.SubjectValidator is a validator for the "subject" field. It is called by the builders before save.
	ledger.SubjectValidator = ledgerDescSubject.Validators[0].(func(string) error)
	// ledgerDescHighwatermark is the schema descriptor for highwatermark field.
	ledgerDescHighwatermark := ledgerFields[3].Descriptor()
	// ledger.DefaultHighwatermark holds the default value on creation for the highwatermark field.
	ledger.DefaultHighwatermark = ledgerDescHighwatermark.Default.(func() time.Time)
	// ledgerDescID is the schema descriptor for id field.
	ledgerDescID := ledgerMixinFields1[0].Descriptor()
	// ledger.DefaultID holds the default value on creation for the id field.
	ledger.DefaultID = ledgerDescID.Default.(func() string)
}
