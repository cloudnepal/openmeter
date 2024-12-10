package recurrence_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/openmeterio/openmeter/openmeter/testutils"
	"github.com/openmeterio/openmeter/pkg/recurrence"
)

func TestPeriod(t *testing.T) {
	startTime := testutils.GetRFC3339Time(t, "2021-01-01T01:00:00Z")
	endTime := testutils.GetRFC3339Time(t, "2021-01-01T02:00:00Z")

	period := recurrence.Period{
		From: startTime,
		To:   endTime,
	}

	t.Run("Should be inclusive of start time", func(t *testing.T) {
		assert.True(t, period.Contains(startTime))
	})

	t.Run("Should be inclusive of end time", func(t *testing.T) {
		assert.True(t, period.Contains(endTime))
	})

	t.Run("Should be true for value in between", func(t *testing.T) {
		assert.True(t, period.Contains(startTime.Add(time.Second)))
	})

	t.Run("Should be false for earlier time", func(t *testing.T) {
		assert.False(t, period.Contains(startTime.Add(-time.Second)))
	})

	t.Run("Should be false for later time", func(t *testing.T) {
		assert.False(t, period.Contains(endTime.Add(time.Second)))
	})
}
