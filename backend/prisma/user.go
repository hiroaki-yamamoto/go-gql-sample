package prisma

// Functions for User model.

// GetID retuns the user's ID.
func (me User) GetID() string {
	return me.ID
}
