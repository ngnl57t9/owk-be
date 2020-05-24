// Code generated by entc, DO NOT EDIT.

package ent

import (
	"owknight-be/ent/profile"
	"owknight-be/ent/role"
	"owknight-be/ent/schema"
	"owknight-be/ent/session"
	"owknight-be/ent/user"
	"time"
)

// The init function reads all schema descriptors with runtime
// code (default values, validators or hooks) and stitches it
// to their package variables.
func init() {
	profileFields := schema.Profile{}.Fields()
	_ = profileFields
	// profileDescUpdatedAt is the schema descriptor for updatedAt field.
	profileDescUpdatedAt := profileFields[2].Descriptor()
	// profile.DefaultUpdatedAt holds the default value on creation for the updatedAt field.
	profile.DefaultUpdatedAt = profileDescUpdatedAt.Default.(func() time.Time)
	// profileDescCreatedAt is the schema descriptor for createdAt field.
	profileDescCreatedAt := profileFields[3].Descriptor()
	// profile.UpdateDefaultCreatedAt holds the default value on update for the createdAt field.
	profile.UpdateDefaultCreatedAt = profileDescCreatedAt.UpdateDefault.(func() time.Time)
	roleFields := schema.Role{}.Fields()
	_ = roleFields
	// roleDescUpdatedAt is the schema descriptor for updatedAt field.
	roleDescUpdatedAt := roleFields[2].Descriptor()
	// role.DefaultUpdatedAt holds the default value on creation for the updatedAt field.
	role.DefaultUpdatedAt = roleDescUpdatedAt.Default.(func() time.Time)
	// roleDescCreatedAt is the schema descriptor for createdAt field.
	roleDescCreatedAt := roleFields[3].Descriptor()
	// role.UpdateDefaultCreatedAt holds the default value on update for the createdAt field.
	role.UpdateDefaultCreatedAt = roleDescCreatedAt.UpdateDefault.(func() time.Time)
	sessionFields := schema.Session{}.Fields()
	_ = sessionFields
	// sessionDescCreatedAt is the schema descriptor for createdAt field.
	sessionDescCreatedAt := sessionFields[2].Descriptor()
	// session.DefaultCreatedAt holds the default value on creation for the createdAt field.
	session.DefaultCreatedAt = sessionDescCreatedAt.Default.(func() time.Time)
	// sessionDescUpdatedAt is the schema descriptor for updatedAt field.
	sessionDescUpdatedAt := sessionFields[3].Descriptor()
	// session.UpdateDefaultUpdatedAt holds the default value on update for the updatedAt field.
	session.UpdateDefaultUpdatedAt = sessionDescUpdatedAt.UpdateDefault.(func() time.Time)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescName is the schema descriptor for name field.
	userDescName := userFields[0].Descriptor()
	// user.NameValidator is a validator for the "name" field. It is called by the builders before save.
	user.NameValidator = func() func(string) error {
		validators := userDescName.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(name string) error {
			for _, fn := range fns {
				if err := fn(name); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// userDescUsername is the schema descriptor for username field.
	userDescUsername := userFields[1].Descriptor()
	// user.UsernameValidator is a validator for the "username" field. It is called by the builders before save.
	user.UsernameValidator = func() func(string) error {
		validators := userDescUsername.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(username string) error {
			for _, fn := range fns {
				if err := fn(username); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// userDescUpdatedAt is the schema descriptor for updatedAt field.
	userDescUpdatedAt := userFields[4].Descriptor()
	// user.DefaultUpdatedAt holds the default value on creation for the updatedAt field.
	user.DefaultUpdatedAt = userDescUpdatedAt.Default.(func() time.Time)
	// userDescCreatedAt is the schema descriptor for createdAt field.
	userDescCreatedAt := userFields[5].Descriptor()
	// user.UpdateDefaultCreatedAt holds the default value on update for the createdAt field.
	user.UpdateDefaultCreatedAt = userDescCreatedAt.UpdateDefault.(func() time.Time)
}