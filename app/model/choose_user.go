package model

// ChooseUser represents the model for selecting a user.
type ChooseUser struct {
	Users []User `json:"users"` // List of users to choose from
}

// User represents a user in the system.
type User struct {
	ID   string `json:"id"`   // Unique identifier for the user
	Name string `json:"name"` // Name of the user
}

// NewChooseUser creates a new ChooseUser model with the provided users.
func NewChooseUser(users []User) *ChooseUser {
	return &ChooseUser{
		Users: users,
	}
}

// GetUserByID retrieves a user by their ID from the ChooseUser model.
func (c *ChooseUser) GetUserByID(id string) *User {
	for _, user := range c.Users {
		if user.ID == id {
			return &user
		}
	}
	return nil // Return nil if no user found with the given ID
}

// GetAllUsers returns all users in the ChooseUser model.
func (c *ChooseUser) GetAllUsers() []User {
	return c.Users
}

// AddUser adds a new user to the ChooseUser model.
func (c *ChooseUser) AddUser(user User) {
	c.Users = append(c.Users, user)
}

// RemoveUser removes a user from the ChooseUser model by their ID.
func (c *ChooseUser) RemoveUser(id string) {
	for i, user := range c.Users {
		if user.ID == id {
			c.Users = append(c.Users[:i], c.Users[i+1:]...)
			break
		}
	}
}

// UpdateUser updates an existing user in the ChooseUser model.
func (c *ChooseUser) UpdateUser(updatedUser User) {
	for i, user := range c.Users {
		if user.ID == updatedUser.ID {
			c.Users[i] = updatedUser
			break
		}
	}
}

// ClearUsers clears all users from the ChooseUser model.
func (c *ChooseUser) ClearUsers() {
	c.Users = []User{}
}
