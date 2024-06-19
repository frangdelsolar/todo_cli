package auth

import (
	"fmt"
)

// GetUserById retrieves a user from the database by their ID.
//
// Parameters:
// - id: the ID of the user to retrieve.
//
// Returns:
// - *User
func GetUserById(id string) (*User, error) {
	var u User
	
	db.First(&u, "id = ?", id)
	if u == (User{}) {
		return nil, fmt.Errorf("user with ID %s not found", fmt.Sprint(id))
	}
	
	return &u, nil
}


// GetAllUsers retrieves all users from the database.
//
// Returns:
// - []User
func GetAllUsers() []User {
	var users []User
	
	db.Find(&users)

	if len(users) == 0 {
		log.Warn().Msg("No users found")
	}
	
	return users
}

// CreateUser creates a new user with the given name and email.
//
// Parameters:
// - name: the name of the user.
// - email: the email of the user.
// Returns:
// - *User: the newly created user.
// - error: an error if the user creation fails.
func CreateUser(name string, email string) (*User, error) {

	u, err := NewUser(name, email)
	if err != nil {
		log.Err(err).Msg("Error creating user")
		return u, err
	}

	db.Create(&u)

	return u, nil
}

// UpdateUser updates the name and email of a user with the given ID.
//
// Parameters:
// - id: the ID of the user to update.
// - name: the new name of the user.
// - email: the new email of the user.
// - requestedBy: the ID of the user requesting the update.
//
// Returns:
// - *User
func UpdateUser(id string, name string, email string, requestedBy string) (*User, error) {

	u, err := GetUserById(id)
	if err != nil {
		log.Err(err).Msg("Error retrieving User")
		return nil, err
	}

	if canUpdate, err := CanUpdateUser(id, requestedBy); !canUpdate {
		return nil, fmt.Errorf("user %s is not authorized to delete user %s", requestedBy, id)
	} else if err != nil {
		return nil, err
	}

	err = u.Update(name, email)
	if err != nil {
		return u, err
	}

	db.Save(&u)
	
	return u, nil
}

// DeleteUser deletes a user with the given ID if the user making the request is authorized to do so.
//
// Parameters:
// - id: the ID of the user to delete.
// - requestedBy: the ID of the user making the request.
//
// Returns:
// - error: an error if the user making the request is not authorized to delete the user or if there was an error retrieving the user.
func DeleteUser(id string, requestedBy string) error {	
	u, err := GetUserById(id)
	if err != nil {
		log.Err(err).Msg("Error retrieving User")
		return err
	}

	if canDelete, err := CanUpdateUser(id, requestedBy); !canDelete {
		return fmt.Errorf("user %s is not authorized to delete user %s", requestedBy, id)
	} else if err != nil {
		return err
	}


	db.Delete(&u)
	
	return nil
}

// CanUpdateUser checks if the user with the given userId can be updated by the user with the given requestedBy.
//
// Parameters:
// - userId: the ID of the user to update.
// - requestedBy: the ID of the user requesting the update.
//
// Returns:
// - bool: true if the user can be updated, false otherwise.
// - error: an error if there was an issue retrieving the users or if the user is not authorized to update the user.
func CanUpdateUser(userId string, requestedBy string) (bool, error) {
	rBy, err := GetUserById(requestedBy)
	if err != nil {
		log.Err(err).Msg("Error retrieving User")
		return false, err
	}

	u, err := GetUserById(userId)
	if err != nil {
		log.Err(err).Msg("Error retrieving User")
		return false, err
	}

	if rBy.ID != u.ID {
		return false, fmt.Errorf("user %s is not authorized to delete user %s", requestedBy, userId)
	}
	
	return true, nil
}