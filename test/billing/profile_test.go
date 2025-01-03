package billing

import (
	"context"
	"testing"
	"time"

	"github.com/samber/lo"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	appentitybase "github.com/openmeterio/openmeter/openmeter/app/entity/base"
	"github.com/openmeterio/openmeter/openmeter/billing"
	"github.com/openmeterio/openmeter/pkg/datex"
	"github.com/openmeterio/openmeter/pkg/models"
)

type ProfileTestSuite struct {
	BaseSuite
}

func TestProfile(t *testing.T) {
	suite.Run(t, new(ProfileTestSuite))
}

func (s *ProfileTestSuite) TestProfileLifecycle() {
	ctx := context.Background()
	ns := "test_create_billing_profile"

	_ = s.InstallSandboxApp(s.T(), ns)

	s.T().Run("missing default profile", func(t *testing.T) {
		defaultProfile, err := s.BillingService.GetDefaultProfile(ctx, billing.GetDefaultProfileInput{
			Namespace: ns,
		})
		require.NoError(t, err)
		require.Nil(t, defaultProfile)
	})

	var profile *billing.Profile
	var err error

	minimalCreateProfileInput := MinimalCreateProfileInputTemplate
	minimalCreateProfileInput.Namespace = ns

	s.T().Run("create default profile", func(t *testing.T) {
		profile, err = s.BillingService.CreateProfile(ctx, minimalCreateProfileInput)

		require.NoError(t, err)
		require.NotNil(t, profile)
	})

	profile.CreatedAt = profile.CreatedAt.Truncate(time.Microsecond)
	profile.UpdatedAt = profile.UpdatedAt.Truncate(time.Microsecond)
	profile.WorkflowConfig.CreatedAt = profile.WorkflowConfig.CreatedAt.Truncate(time.Microsecond)
	profile.WorkflowConfig.UpdatedAt = profile.WorkflowConfig.UpdatedAt.Truncate(time.Microsecond)

	s.T().Run("fetching the default profile is possible", func(t *testing.T) {
		defaultProfile, err := s.BillingService.GetDefaultProfile(ctx, billing.GetDefaultProfileInput{
			Namespace: ns,
		})
		require.NoError(t, err)
		require.NotNil(t, defaultProfile)
		require.Equal(t, profile, defaultProfile)
	})

	s.T().Run("creating a second default profile fails", func(t *testing.T) {
		_, err := s.BillingService.CreateProfile(ctx, minimalCreateProfileInput)
		require.Error(t, err)
		require.ErrorIs(t, err, billing.ErrDefaultProfileAlreadyExists)
	})

	s.T().Run("fetching the profile by id", func(t *testing.T) {
		fetchedProfile, err := s.BillingService.GetProfile(ctx, billing.GetProfileInput{
			Profile: models.NamespacedID{
				Namespace: ns,
				ID:        profile.ID,
			},
			Expand: billing.ProfileExpand{
				Apps: true,
			},
		})

		require.NoError(t, err)
		require.Equal(t, profile, fetchedProfile)
	})

	s.T().Run("deleted profile handling", func(t *testing.T) {
		require.NoError(t, s.BillingService.DeleteProfile(ctx, billing.DeleteProfileInput{
			Namespace: ns,
			ID:        profile.ID,
		}))

		t.Run("deleting a profile twice yields an error", func(t *testing.T) {
			require.ErrorIs(t, s.BillingService.DeleteProfile(ctx, billing.DeleteProfileInput{
				Namespace: ns,
				ID:        profile.ID,
			}), billing.ErrProfileAlreadyDeleted)
		})

		t.Run("fetching a deleted profile by id returns the profile", func(t *testing.T) {
			fetchedProfile, err := s.BillingService.GetProfile(ctx, billing.GetProfileInput{
				Profile: models.NamespacedID{
					Namespace: ns,
					ID:        profile.ID,
				},
			})

			require.NoError(t, err)
			require.Equal(t, profile.ID, fetchedProfile.ID)
		})

		t.Run("same profile can be created after the previous one was deleted", func(t *testing.T) {
			profile, err = s.BillingService.CreateProfile(ctx, minimalCreateProfileInput)

			require.NoError(t, err)
			require.NotNil(t, profile)
		})
	})
}

func (s *ProfileTestSuite) TestProfileFieldSetting() {
	ctx := context.Background()
	t := s.T()
	ns := "test_profile_field_setting"

	app := s.InstallSandboxApp(s.T(), ns)

	input := billing.CreateProfileInput{
		Namespace: ns,
		Default:   true,
		Name:      "Awesome Default Profile",

		Metadata: map[string]string{
			"key": "value",
		},

		WorkflowConfig: billing.WorkflowConfig{
			Collection: billing.CollectionConfig{
				Alignment: billing.AlignmentKindSubscription,
				Interval:  datex.MustParse(t, "PT30M"),
			},
			Invoicing: billing.InvoicingConfig{
				AutoAdvance: true,
				DraftPeriod: datex.MustParse(t, "PT1H"),
				DueAfter:    datex.MustParse(t, "PT24H"),
			},
			Payment: billing.PaymentConfig{
				CollectionMethod: billing.CollectionMethodSendInvoice,
			},
		},

		Supplier: billing.SupplierContact{
			Name: "Awesome Supplier",
			Address: models.Address{
				Country:     lo.ToPtr(models.CountryCode("US")),
				PostalCode:  lo.ToPtr("12345"),
				City:        lo.ToPtr("City"),
				State:       lo.ToPtr("State"),
				Line1:       lo.ToPtr("Line 1"),
				Line2:       lo.ToPtr("Line 2"),
				PhoneNumber: lo.ToPtr("1234567890"),
			},
		},

		Apps: billing.CreateProfileAppsInput{
			Invoicing: billing.AppReference{
				Type: appentitybase.AppTypeSandbox,
			},
			Payment: billing.AppReference{
				Type: appentitybase.AppTypeSandbox,
			},
			Tax: billing.AppReference{
				Type: appentitybase.AppTypeSandbox,
			},
		},
	}

	profile, err := s.BillingService.CreateProfile(ctx, input)
	require.NoError(s.T(), err)
	require.NotNil(s.T(), profile)

	profile.CreatedAt = profile.CreatedAt.Truncate(time.Microsecond)
	profile.UpdatedAt = profile.UpdatedAt.Truncate(time.Microsecond)
	profile.WorkflowConfig.CreatedAt = profile.WorkflowConfig.CreatedAt.Truncate(time.Microsecond)
	profile.WorkflowConfig.UpdatedAt = profile.WorkflowConfig.UpdatedAt.Truncate(time.Microsecond)

	// Let's fetch the profile again
	fetchedProfile, err := s.BillingService.GetProfile(ctx, billing.GetProfileInput{
		Profile: models.NamespacedID{
			Namespace: ns,
			ID:        profile.ID,
		},
		Expand: billing.ProfileExpandAll,
	})

	// Sanity check db conversion & fetching
	require.NoError(s.T(), err)
	require.Equal(s.T(), profile, fetchedProfile)

	// let's add the db derived fields to the input
	expectedProfile := billing.Profile{
		BaseProfile: billing.BaseProfile{
			ID: profile.ID,

			Namespace:   input.Namespace,
			Name:        input.Name,
			Description: input.Description,
			Default:     input.Default,

			CreatedAt: fetchedProfile.CreatedAt,
			UpdatedAt: fetchedProfile.UpdatedAt,
			DeletedAt: fetchedProfile.DeletedAt,

			WorkflowConfig: input.WorkflowConfig,
			Supplier:       input.Supplier,

			Metadata:      input.Metadata,
			AppReferences: fetchedProfile.AppReferences,
		},
		Apps: fetchedProfile.Apps,
	}

	expectedProfile.WorkflowConfig.ID = fetchedProfile.WorkflowConfig.ID
	expectedProfile.WorkflowConfig.CreatedAt = fetchedProfile.WorkflowConfig.CreatedAt
	expectedProfile.WorkflowConfig.UpdatedAt = fetchedProfile.WorkflowConfig.UpdatedAt
	expectedProfile.WorkflowConfig.DeletedAt = fetchedProfile.WorkflowConfig.DeletedAt

	// Let's check if the fields are set correctly
	require.Equal(s.T(), expectedProfile, *fetchedProfile)
	require.Equal(s.T(), app.GetID(), fetchedProfile.Apps.Tax.GetID())
	require.Equal(s.T(), app.GetID(), fetchedProfile.Apps.Invoicing.GetID())
	require.Equal(s.T(), app.GetID(), fetchedProfile.Apps.Payment.GetID())
}

func (s *ProfileTestSuite) TestProfileUpdates() {
	// Given a profile
	ctx := context.Background()
	ns := "test_profile_updates"

	_ = s.InstallSandboxApp(s.T(), ns)

	input := billing.CreateProfileInput{
		Namespace: ns,
		Default:   true,

		Name: "Awesome Default Profile",

		Apps: MinimalCreateProfileInputTemplate.Apps,

		WorkflowConfig: billing.WorkflowConfig{
			Collection: billing.CollectionConfig{
				Alignment: billing.AlignmentKindSubscription,
				Interval:  datex.MustParse(s.T(), "PT30M"),
			},
			Invoicing: billing.InvoicingConfig{
				AutoAdvance: true,
				DraftPeriod: datex.MustParse(s.T(), "PT1H"),
				DueAfter:    datex.MustParse(s.T(), "PT24H"),
			},
			Payment: billing.PaymentConfig{
				CollectionMethod: billing.CollectionMethodSendInvoice,
			},
		},

		Supplier: billing.SupplierContact{
			Name: "Awesome Supplier",
			Address: models.Address{
				Country:     lo.ToPtr(models.CountryCode("US")),
				PostalCode:  lo.ToPtr("12345"),
				City:        lo.ToPtr("City"),
				State:       lo.ToPtr("State"),
				Line1:       lo.ToPtr("Line 1"),
				Line2:       lo.ToPtr("Line 2"),
				PhoneNumber: lo.ToPtr("1234567890"),
			},
		},
	}

	profile, err := s.BillingService.CreateProfile(ctx, input)
	require.NoError(s.T(), err)
	require.NotNil(s.T(), profile)

	profile.CreatedAt = profile.CreatedAt.Truncate(time.Microsecond)
	profile.UpdatedAt = profile.UpdatedAt.Truncate(time.Microsecond)
	profile.WorkflowConfig.CreatedAt = profile.WorkflowConfig.CreatedAt.Truncate(time.Microsecond)
	profile.WorkflowConfig.UpdatedAt = profile.WorkflowConfig.UpdatedAt.Truncate(time.Microsecond)

	// Let's fetch the profile again
	fetchedProfile, err := s.BillingService.GetProfile(ctx, billing.GetProfileInput{
		Profile: models.NamespacedID{
			Namespace: ns,
			ID:        profile.ID,
		},
		Expand: billing.ProfileExpandAll,
	})

	// Sanity check db conversion & fetching
	require.NoError(s.T(), err)
	require.Equal(s.T(), profile, fetchedProfile)

	s.T().Run("update profile", func(t *testing.T) {
		// When updating the profile
		updateInput := billing.UpdateProfileInput{
			ID:          profile.ID,
			Namespace:   ns,
			Default:     true,
			Name:        "Awesome Default Profile [update]",
			Description: lo.ToPtr("Updated description"),
			CreatedAt:   profile.CreatedAt,

			UpdatedAt: profile.UpdatedAt,

			WorkflowConfig: billing.WorkflowConfig{
				CreatedAt: profile.WorkflowConfig.CreatedAt,
				Collection: billing.CollectionConfig{
					Alignment: billing.AlignmentKindSubscription,
					Interval:  datex.MustParse(s.T(), "PT30M"),
				},
				Invoicing: billing.InvoicingConfig{
					AutoAdvance: true,
					DraftPeriod: datex.MustParse(s.T(), "PT2H"),
					DueAfter:    datex.MustParse(s.T(), "PT48H"),
				},
				Payment: billing.PaymentConfig{
					CollectionMethod: billing.CollectionMethodChargeAutomatically,
				},
			},

			Supplier: billing.SupplierContact{
				Name: "Awesome Supplier [update]",
				Address: models.Address{
					Country:     lo.ToPtr(models.CountryCode("HU")),
					PostalCode:  lo.ToPtr("12345 [update]"),
					City:        lo.ToPtr("City [update]"),
					State:       lo.ToPtr("State [update]"),
					Line1:       lo.ToPtr("Line 1 [update]"),
					Line2:       lo.ToPtr("Line 2 [update]"),
					PhoneNumber: lo.ToPtr("1234567890 [update]"),
				},
			},
		}

		updatedProfile, err := s.BillingService.UpdateProfile(ctx, updateInput)

		// Then the profile is updated

		require.NoError(t, err)
		require.NotNil(t, updatedProfile)

		require.NotEqual(t, fetchedProfile.UpdatedAt, updatedProfile.UpdatedAt, "the new updated at is returned")

		// Set up DB only fields
		expectedOutput := billing.Profile{
			BaseProfile: billing.BaseProfile{
				ID:          updateInput.ID,
				Namespace:   updateInput.Namespace,
				CreatedAt:   fetchedProfile.CreatedAt,
				Name:        updateInput.Name,
				Description: updateInput.Description,
				Metadata:    updateInput.Metadata,
				Supplier:    updateInput.Supplier,
				Default:     updateInput.Default,

				WorkflowConfig: updateInput.WorkflowConfig,
				AppReferences:  fetchedProfile.AppReferences,
			},
		}
		expectedOutput.WorkflowConfig.ID = fetchedProfile.WorkflowConfig.ID
		expectedOutput.UpdatedAt = updatedProfile.UpdatedAt                               // This is checked by the previous assertion
		expectedOutput.WorkflowConfig.UpdatedAt = updatedProfile.WorkflowConfig.UpdatedAt // This is checked by the previous assertion
		expectedOutput.Apps = fetchedProfile.Apps

		require.Equal(t, expectedOutput, *updatedProfile)
	})
}
