package functions

import (
	"context"
	"fmt"
	"time"

	"github.com/mt1976/frantic-core/logHandler"
	baseline "github.com/mt1976/frantic-mass/app/dao/baseline"
	"github.com/mt1976/frantic-mass/app/dao/dateIndex"
	"github.com/mt1976/frantic-mass/app/dao/goal"
	user "github.com/mt1976/frantic-mass/app/dao/user"
	"github.com/mt1976/frantic-mass/app/dao/weightProjection"
	"github.com/mt1976/frantic-mass/app/dao/weightProjectionHistory"
	"github.com/mt1976/frantic-mass/app/types"
)

func BuildWeightGoalProjection(user user.User, weight types.Weight, goal goal.Goal) error {
	// Projection function to create weight projections based on user baseline and goal.
	// This function will create projections for the next n months based on the user's baseline and goal.
	// It assumes that the user has a baseline set up and a goal defined.

	// Get Today's date
	today := Today()
	logHandler.InfoLogger.Printf("Today's date is: %v", today)
	// Get the Today DateIndex record
	_, todayRecord, err := dateIndex.GetToday()
	if err != nil {
		logHandler.ErrorLogger.Println("Error getting today's DateIndex record:", err)
		return fmt.Errorf("error getting today's DateIndex: %v", err)
	}

	// Ensure the user has a baseline set up
	if user.ID == 0 {
		logHandler.ErrorLogger.Println("User ID is not set, cannot create projections.")
		return fmt.Errorf("user ID is not set")
	}

	if weight.IsZero() {
		logHandler.ErrorLogger.Println("Weight is zero, cannot create projections.")
		return fmt.Errorf("weight is zero")
	}

	// Calculate the number of weeks in the next n months.

	// Get the basline for the user
	userID := user.ID
	thisBaseline, err := baseline.GetByUserID(userID)
	if err != nil {
		logHandler.ErrorLogger.Println("Error getting baseline for user:", userID, err)
		return fmt.Errorf("error getting baseline for user %d: %v", userID, err)
	}
	weeksInNextMonths := thisBaseline.ProjectionPeriod * 4 // Assuming 4 weeks per month
	logHandler.InfoLogger.Printf("Creating %d projections for UserID: %d", weeksInNextMonths, userID)

	weeklyWeightLoss := goal.LossPerWeek
	if goal.AverageWeightLoss.IsTrue() {
		// If the goal is an average weight loss goal, we calculate the weekly weight loss based on the target weight and the number of weeks.
		avgWeightLoss, err := AverageWeightLoss(userID)
		if err != nil {
			logHandler.ErrorLogger.Println("Error calculating average weight loss for user:", userID, err)
			return fmt.Errorf("error calculating average weight loss for user %d: %v", userID, err)
		}
		weeklyWeightLoss = *avgWeightLoss
		logHandler.InfoLogger.Printf("Calculated weekly weight loss for user %d: %v", userID, weeklyWeightLoss.String())
	}

	if weeklyWeightLoss.IsZero() {
		logHandler.WarningLogger.Println("Weekly weight loss is zero")
	}

	logHandler.InfoLogger.Printf("Projection Period: %d weeks, UserID: %d, Weight: %v, Goal Loss Per Week: %v", weeksInNextMonths, userID, weight, weeklyWeightLoss)

	trackingWeight := weight

	for j := 1; j < weeksInNextMonths; j++ {

		//quick int to float
		//fj := float64(j)
		trackingWeight = trackingWeight.Minus(weeklyWeightLoss)
		when := time.Now().Add(time.Duration(j) * (time.Hour * 24) * 7) // Add j weeks to the current date
		var vstarget string
		switch {
		case trackingWeight.LT(goal.TargetWeight.Value):
			logHandler.WarningLogger.Printf("Tracking weight %v is less than target weight %v for UserID: %d", trackingWeight, goal.TargetWeight, userID)
			vstarget = "Below " + goal.TargetWeight.String() + " (Target)"
		case trackingWeight.GT(goal.TargetWeight.Value):
			logHandler.WarningLogger.Printf("Tracking weight %v is greater than target weight %v for UserID: %d", trackingWeight, goal.TargetWeight, userID)
			vstarget = "Above " + goal.TargetWeight.String() + " (Target)"
		default:
			logHandler.InfoLogger.Printf("Tracking weight %v is equal to target weight %v for UserID: %d", trackingWeight, goal.TargetWeight, userID)
			vstarget = "At " + goal.TargetWeight.String() + " (Target)"
		}

		toGoal := goal.TargetWeight.Minus(trackingWeight) // Calculate the total weight loss needed to reach the goal
		toGoal = *toGoal.Invert()

		np, newProjectionErr := weightProjection.New(context.TODO(), userID, goal.ID, j, trackingWeight, weeklyWeightLoss, when, fmt.Sprintf("Projection For_%v/%v @ %v", userID, j, weeklyWeightLoss.Value), vstarget, toGoal)
		if newProjectionErr != nil {
			logHandler.ErrorLogger.Println(newProjectionErr)
			return fmt.Errorf("error creating projection for user %d: %v", userID, newProjectionErr)
		} else {
			logHandler.InfoLogger.Printf("Projection Created:[%v]", np.CompositeID)
		}
		// Update the weight projection history
		_, err = weightProjectionHistory.New(context.TODO(), todayRecord, np)
		if err != nil {
			logHandler.ErrorLogger.Println("Error creating weight projection history:", err)
			return fmt.Errorf("error creating weight projection history for user %d: %v", userID, err)
		} else {
			logHandler.InfoLogger.Printf("Weight Projection History Created:[%v]", np.CompositeID)
		}

	}
	// Log the successful creation of projections
	logHandler.InfoLogger.Printf("Successfully created %d projections for UserID: %d", weeksInNextMonths, userID)
	return nil
}

func BuildWeightGoalsProjections(user user.User, weight types.Weight) error {
	// Projections function to create weight projections based on user baseline and goal.
	// This function will create projections for the next n months based on the user's baseline and goal.
	// It assumes that the user has a baseline set up and a goal defined.

	// Get the user's goals
	goals, err := goal.GetAllWhere(goal.FIELD_UserID, user.ID)
	if err != nil {
		logHandler.ErrorLogger.Println("Error getting goals for user:", user.ID, err)
		return fmt.Errorf("error getting goals for user %d: %v", user.ID, err)
	}
	if len(goals) == 0 {
		logHandler.ErrorLogger.Println("No goals found for user:", user.ID)
		return fmt.Errorf("no goals found for user %d", user.ID)
	}

	// Loop through each goal and create projections
	for _, g := range goals {
		logHandler.InfoLogger.Printf("Creating projections for UserID: %d, GoalID: %d", user.ID, g.ID)
		err = BuildWeightGoalProjection(user, weight, g)
		if err != nil {
			logHandler.ErrorLogger.Println("Error creating projections for user:", user.ID, "goal:", g.ID, err)
			return fmt.Errorf("error creating projections for user %d, goal %d: %v", user.ID, g.ID, err)
		}
	}

	logHandler.InfoLogger.Printf("Successfully created projections for UserID: %d", user.ID)
	return nil
}
