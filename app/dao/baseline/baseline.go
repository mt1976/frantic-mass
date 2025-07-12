package baseline

import (
	"fmt"

	"github.com/mt1976/frantic-core/logHandler"
)

func CmToFeetInches(cm float64) (int, int) {
	inchesTotal := cm / 2.54
	feet := int(inchesTotal) / 12
	inches := int(inchesTotal) % 12
	return feet, inches
}

func FeetInchesToCm(feet, inches int) float64 {
	inchesTotal := float64(feet*12 + inches)
	cm := inchesTotal * 2.54
	return cm
}

func (record *baseline_Store) HeightAsFeetInches() string {
	return record.HeightAsFeetInches()
}

func (record *baseline_Store) HeightAsCm() string {
	return fmt.Sprintf("%.2f cm", record.Height)
}

func (record *baseline_Store) HeightAsInches() string {
	return record.HeightAsInches()
}

func (record *baseline_Store) HeightAsMetres() string {
	return record.HeightAsMetres()
}
func (record *baseline_Store) HeightAsString() string {
	return record.HeightAsString()
}

func GetByUserID(userID int) (*baseline_Store, error) {
	// This function should implement the logic to retrieve a baseline record by userID
	// For now, we return nil and nil to avoid compilation errors
	if userID <= 0 {
		return nil, fmt.Errorf("invalid userID: %d", userID)
	}

	records, err := GetAllWhere(FIELD_UserID, userID)
	if err != nil {
		logHandler.ErrorLogger.Println("Error retrieving baseline by userID:", userID, "Error:", err)
		return nil, err
	}
	if len(records) == 0 {
		logHandler.ErrorLogger.Println("No baseline found for userID:", userID)
		return nil, fmt.Errorf("no baseline found for userID: %d", userID)
	}
	if len(records) > 1 {
		logHandler.ErrorLogger.Println("Multiple baselines found for userID:", userID, "using the first one")
	}
	// Assuming records is a slice of baseline_Store, return the first record
	return &records[0], nil
}
