package database

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalculationHistoryRepository_Save(t *testing.T) {
	t.Run("saves calculation to history", func(t *testing.T) {
		var (
			userID      = "someUserID"
			calculation = Calculation{Expression: "2+2", Result: "4"}
			repository  = newCalculationHistoryRepository()
		)

		err := repository.Save(userID, calculation)

		assert.NoError(t, err)
		history, err := repository.GetHistory(userID)
		assert.NoError(t, err)
		assert.Contains(t, history, calculation)
	})

	t.Run("resets history when max history size is exceeded", func(t *testing.T) {
		var (
			userID      = "someUserID"
			calculation = Calculation{Expression: "2+2", Result: "4"}
			repository  = newCalculationHistoryRepository()
		)

		repository.maxHistory = 1

		err := repository.Save(userID, calculation)
		assert.NoError(t, err)

		err = repository.Save(userID, calculation)
		assert.NoError(t, err)

		history, err := repository.GetHistory(userID)
		assert.NoError(t, err)
		assert.Equal(t, 1, len(history))
		assert.Contains(t, history, calculation)
	})
}

func TestCalculationHistoryRepository_GetHistory(t *testing.T) {
	t.Run("returns empty history for new user", func(t *testing.T) {
		var (
			userID     = "newUserID"
			repository = newCalculationHistoryRepository()
		)

		history, err := repository.GetHistory(userID)

		assert.NoError(t, err)
		assert.Empty(t, history)
	})
}

func TestCalculationHistoryRepository_GetCalculationHistoryRepository(t *testing.T) {
	t.Run("returns same instance on multiple calls", func(t *testing.T) {
		var (
			userID      = "someUserID"
			calculation = Calculation{Expression: "2+2", Result: "4"}
			repository1 = GetCalculationHistoryRepository()
			repository2 = GetCalculationHistoryRepository()
		)
		assert.NoError(t, repository1.Save(userID, calculation))

		history, err := repository2.GetHistory(userID)

		assert.NoError(t, err)
		assert.Equal(t, repository1, repository2)
		assert.Equal(t, calculation, history[0])
	})
}
