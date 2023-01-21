// Code generated by ent, DO NOT EDIT.

package ent

import (
	"sample/ent/schema"
	"sample/ent/user"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescAge is the schema descriptor for age field.
	userDescAge := userFields[2].Descriptor()
	// user.DefaultAge holds the default value on creation for the age field.
	user.DefaultAge = userDescAge.Default.(int)
}
