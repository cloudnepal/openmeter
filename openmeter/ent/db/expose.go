// Code generated by ent, DO NOT EDIT.

package db

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"github.com/openmeterio/openmeter/pkg/framework/entutils"
)

func (c *Client) GetConfig() *entutils.RawEntConfig {
	return &entutils.RawEntConfig{
		Driver: c.config.driver,
		Debug:  c.config.debug,
		Log:    c.config.log,
	}
}

type ExposedTxDriver struct {
	Driver *txDriver
}

// ignores hooks
func (d *ExposedTxDriver) Rollback() error {
	return d.Driver.tx.Rollback()
}

// ignores hooks
func (d *ExposedTxDriver) Commit() error {
	return d.Driver.tx.Commit()
}

// HijackTx returns a new transaction driver with the provided options.
// The returned transaction can later be used to instanciate new clients.
func (c *Client) HijackTx(ctx context.Context, opts *sql.TxOptions) (context.Context, *entutils.RawEntConfig, *ExposedTxDriver, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, nil, nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}

	driver := &txDriver{tx: tx, drv: c.driver}

	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return ctx, &entutils.RawEntConfig{
		Driver: cfg.driver,
		Debug:  cfg.debug,
		Log:    cfg.log,
	}, &ExposedTxDriver{Driver: driver}, nil
}

// NewTxClientFromConfig creates a new transactional client from a (hijacked) configuration.
func NewTxClientFromRawConfig(ctx context.Context, cfg entutils.RawEntConfig) *Tx {
	config := config{
		driver: cfg.Driver,
		debug:  cfg.Debug,
		log:    cfg.Log,
		hooks:  &hooks{},
		inters: &inters{},
	}

	return &Tx{
		ctx:    ctx,
		config: config,
		// Clients templated from defined schemas

		BalanceSnapshot: NewBalanceSnapshotClient(config),

		BillingInvoice: NewBillingInvoiceClient(config),

		BillingInvoiceItem: NewBillingInvoiceItemClient(config),

		BillingProfile: NewBillingProfileClient(config),

		BillingWorkflowConfig: NewBillingWorkflowConfigClient(config),

		Entitlement: NewEntitlementClient(config),

		Feature: NewFeatureClient(config),

		Grant: NewGrantClient(config),

		NotificationChannel: NewNotificationChannelClient(config),

		NotificationEvent: NewNotificationEventClient(config),

		NotificationEventDeliveryStatus: NewNotificationEventDeliveryStatusClient(config),

		NotificationRule: NewNotificationRuleClient(config),

		UsageReset: NewUsageResetClient(config),
	}
}
