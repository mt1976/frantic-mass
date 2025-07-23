package functions

import (
	"github.com/mt1976/frantic-core/logHandler"
	"github.com/mt1976/frantic-mass/app/dao/goal"
)

func GetGoals(userID int) ([]goal.Goal, error) {
	// This function retrieves all goals for a user
	// It returns a slice of types.Goal and an error if any occurs

	goals, err := goal.GetAllWhere(goal.FIELD_UserID, userID)
	if err != nil {
		return nil, err
	}

	// Filter out goals that are not active
	var userGoals []goal.Goal
	for _, g := range goals {
		if g.Audit.DeletedBy != "" {
			logHandler.WarningLogger.Printf("Skipping deleted goal ID %d for user %d", g.ID, userID)
			continue // Skip deleted goals
		}
		userGoals = append(userGoals, g)
	}
	if len(userGoals) == 0 {
		logHandler.WarningLogger.Printf("No active goals found for user %d", userID)
		return nil, nil // No active goals found
	}

	return userGoals, nil
}
