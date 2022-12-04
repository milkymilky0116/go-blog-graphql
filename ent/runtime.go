// Code generated by ent, DO NOT EDIT.

package ent

import (
	"github.com/MangoSteen0903/go-blog-graphql/ent/post"
	"github.com/MangoSteen0903/go-blog-graphql/ent/schema"
	"github.com/MangoSteen0903/go-blog-graphql/ent/user"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	postFields := schema.Post{}.Fields()
	_ = postFields
	// postDescLikes is the schema descriptor for Likes field.
	postDescLikes := postFields[2].Descriptor()
	// post.DefaultLikes holds the default value on creation for the Likes field.
	post.DefaultLikes = postDescLikes.Default.(int)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescIsAdmin is the schema descriptor for is_admin field.
	userDescIsAdmin := userFields[3].Descriptor()
	// user.DefaultIsAdmin holds the default value on creation for the is_admin field.
	user.DefaultIsAdmin = userDescIsAdmin.Default.(bool)
}
