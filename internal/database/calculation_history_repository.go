package database

import (
	"log"
	"sync"
)

type Calculation struct {
	Expression string
	Result     string
}

type CalculationHistoryRepository struct {
	mutex                   sync.Mutex
	history                 map[string][]Calculation
	maxHistory              int
	currentTotalHistorySize int // Counter for the total number of entries
}

var (
	instance *CalculationHistoryRepository
	once     sync.Once
)

func newCalculationHistoryRepository() *CalculationHistoryRepository {
	return &CalculationHistoryRepository{
		mutex:                   sync.Mutex{},
		history:                 make(map[string][]Calculation),
		maxHistory:              1000, // Default max entries across all users
		currentTotalHistorySize: 0,    // counter to ensure we reset get_history after a total of maxHistory entries
	}
}

// Exposed Singleton of the CalculationHistoryRepository
func GetCalculationHistoryRepository() *CalculationHistoryRepository {
	once.Do(func() {
		instance = newCalculationHistoryRepository()
	})
	return instance
}

// This is just a simple way for me to avoid this small project from crashing in case of
// something spamming the container for some reason
func (r *CalculationHistoryRepository) resetHistoryIfTooLarge() {
	// and in theory we could be missing a lock on a mutex here, again it's just for demonstration purposes.
	// A simple solution would just be putting this into the Add function, I just moved it here to not pollute right now
	if r.currentTotalHistorySize >= r.maxHistory {
		log.Println("Total get_history size exceeded max limit, resetting get_history map")
		r.history = make(map[string][]Calculation) // Reinitialize the map
		r.currentTotalHistorySize = 0              // Reset the counter
	}
}

func (r *CalculationHistoryRepository) Save(userID string, calculation Calculation) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	// Call resetHistoryIfTooLarge to check and reset get_history if needed
	r.resetHistoryIfTooLarge()

	// Add new entry to the get_history for the user
	r.history[userID] = append(r.history[userID], calculation)

	// Update the total get_history size counter to ensure we can clear the get_history once it grows too large
	r.currentTotalHistorySize++
	return nil
}

func (r *CalculationHistoryRepository) GetHistory(userID string) ([]Calculation, error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	return r.history[userID], nil
}
