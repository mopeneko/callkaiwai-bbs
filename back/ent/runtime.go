// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/mopeneko/callkaiwai-bbs/back/ent/ipaddresslog"
	"github.com/mopeneko/callkaiwai-bbs/back/ent/post"
	"github.com/mopeneko/callkaiwai-bbs/back/ent/schema"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	ipaddresslogFields := schema.IPAddressLog{}.Fields()
	_ = ipaddresslogFields
	// ipaddresslogDescIPAddress is the schema descriptor for ip_address field.
	ipaddresslogDescIPAddress := ipaddresslogFields[0].Descriptor()
	// ipaddresslog.IPAddressValidator is a validator for the "ip_address" field. It is called by the builders before save.
	ipaddresslog.IPAddressValidator = ipaddresslogDescIPAddress.Validators[0].(func(string) error)
	// ipaddresslogDescCreatedAt is the schema descriptor for created_at field.
	ipaddresslogDescCreatedAt := ipaddresslogFields[1].Descriptor()
	// ipaddresslog.DefaultCreatedAt holds the default value on creation for the created_at field.
	ipaddresslog.DefaultCreatedAt = ipaddresslogDescCreatedAt.Default.(func() time.Time)
	// ipaddresslogDescUpdatedAt is the schema descriptor for updated_at field.
	ipaddresslogDescUpdatedAt := ipaddresslogFields[2].Descriptor()
	// ipaddresslog.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	ipaddresslog.DefaultUpdatedAt = ipaddresslogDescUpdatedAt.Default.(func() time.Time)
	// ipaddresslog.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	ipaddresslog.UpdateDefaultUpdatedAt = ipaddresslogDescUpdatedAt.UpdateDefault.(func() time.Time)
	postFields := schema.Post{}.Fields()
	_ = postFields
	// postDescGender is the schema descriptor for gender field.
	postDescGender := postFields[1].Descriptor()
	// post.GenderValidator is a validator for the "gender" field. It is called by the builders before save.
	post.GenderValidator = postDescGender.Validators[0].(func(int) error)
	// postDescTweetURL is the schema descriptor for tweet_url field.
	postDescTweetURL := postFields[3].Descriptor()
	// post.TweetURLValidator is a validator for the "tweet_url" field. It is called by the builders before save.
	post.TweetURLValidator = postDescTweetURL.Validators[0].(func(string) error)
	// postDescTiktokURL is the schema descriptor for tiktok_url field.
	postDescTiktokURL := postFields[4].Descriptor()
	// post.TiktokURLValidator is a validator for the "tiktok_url" field. It is called by the builders before save.
	post.TiktokURLValidator = postDescTiktokURL.Validators[0].(func(string) error)
	// postDescContactURL is the schema descriptor for contact_url field.
	postDescContactURL := postFields[5].Descriptor()
	// post.ContactURLValidator is a validator for the "contact_url" field. It is called by the builders before save.
	post.ContactURLValidator = postDescContactURL.Validators[0].(func(string) error)
	// postDescCreatedAt is the schema descriptor for created_at field.
	postDescCreatedAt := postFields[7].Descriptor()
	// post.DefaultCreatedAt holds the default value on creation for the created_at field.
	post.DefaultCreatedAt = postDescCreatedAt.Default.(func() time.Time)
	// postDescUpdatedAt is the schema descriptor for updated_at field.
	postDescUpdatedAt := postFields[8].Descriptor()
	// post.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	post.DefaultUpdatedAt = postDescUpdatedAt.Default.(func() time.Time)
	// post.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	post.UpdateDefaultUpdatedAt = postDescUpdatedAt.UpdateDefault.(func() time.Time)
}
