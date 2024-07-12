package contractor

import (
	"fmt"
)

// CreateFrequency creates a new frequency based on the input and stores it in the database.
//
// Parameters:
// - input: a pointer to a NewFrequencyInput struct containing the necessary data for creating the frequency.
//
// Returns:
// - *Frequency: the newly created frequency.
// - error: an error if the creation process encounters any issues.
func CreateFrequency(input *NewFrequencyInput) (*Frequency, error) {
	freq, err := NewFrequency(input)
	if err != nil {
		log.Err(err).Msg("Error creating frequency")
		return nil, err
	}

	db.Create(&freq)

	log.Trace().Interface("frequency", freq).Msg("Created frequency")
	log.Info().Msg("Created frequency")

	return freq, nil
}

// GetFrequencyById retrieves a Frequency object by its ID and the ID of the user who requested it.
//
// Parameters:
// - id: The ID of the Frequency object to retrieve.
// - requestedBy: The ID of the user who is requesting the Frequency object.
//
// Returns:
// - *Frequency: A pointer to the retrieved Frequency object.
// - error: An error if the Frequency object with the specified ID is not found.
func GetFrequencyById(id string, requestedBy string) (*Frequency, error) {
	var freq Frequency

	db.First(&freq, "id = ? AND created_by = ?", id, requestedBy)

	if freq == (Frequency{}) {
		return nil, fmt.Errorf("frequency with ID %s not found", id)
	}

	return &freq, nil
}

// DeleteFrequency deletes a frequency from the database by its ID.
//
// Parameters:
// - id: the ID of the frequency to delete.
// - requestedBy: the ID of the user who is requesting the deletion.
//
// Returns:
// - error: an error if the frequency retrieval or deletion fails.
func DeleteFrequency(id string, requestedBy string) error {
	freq, err := GetFrequencyById(id, requestedBy)
	if err != nil {
		log.Err(err).Msg("Error getting frequency")
		return err
	}

	db.Delete(&freq)

	log.Info().Msgf("Deleted frequency with ID %s", id)

	return nil
}
