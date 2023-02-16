// Code generated by ent, DO NOT EDIT.

package post

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/mopeneko/callkaiwai-bbs/back/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.Post {
	return predicate.Post(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.Post {
	return predicate.Post(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.Post {
	return predicate.Post(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.Post {
	return predicate.Post(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.Post {
	return predicate.Post(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.Post {
	return predicate.Post(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.Post {
	return predicate.Post(sql.FieldLTE(FieldID, id))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldName, v))
}

// Gender applies equality check predicate on the "gender" field. It's identical to GenderEQ.
func Gender(v int) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldGender, v))
}

// Introduction applies equality check predicate on the "introduction" field. It's identical to IntroductionEQ.
func Introduction(v string) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldIntroduction, v))
}

// TweetURL applies equality check predicate on the "tweet_url" field. It's identical to TweetURLEQ.
func TweetURL(v string) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldTweetURL, v))
}

// TiktokURL applies equality check predicate on the "tiktok_url" field. It's identical to TiktokURLEQ.
func TiktokURL(v string) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldTiktokURL, v))
}

// ContactURL applies equality check predicate on the "contact_url" field. It's identical to ContactURLEQ.
func ContactURL(v string) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldContactURL, v))
}

// ContactID applies equality check predicate on the "contact_id" field. It's identical to ContactIDEQ.
func ContactID(v string) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldContactID, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldUpdatedAt, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Post {
	return predicate.Post(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Post {
	return predicate.Post(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Post {
	return predicate.Post(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Post {
	return predicate.Post(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Post {
	return predicate.Post(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Post {
	return predicate.Post(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Post {
	return predicate.Post(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Post {
	return predicate.Post(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Post {
	return predicate.Post(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Post {
	return predicate.Post(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Post {
	return predicate.Post(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Post {
	return predicate.Post(sql.FieldContainsFold(FieldName, v))
}

// GenderEQ applies the EQ predicate on the "gender" field.
func GenderEQ(v int) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldGender, v))
}

// GenderNEQ applies the NEQ predicate on the "gender" field.
func GenderNEQ(v int) predicate.Post {
	return predicate.Post(sql.FieldNEQ(FieldGender, v))
}

// GenderIn applies the In predicate on the "gender" field.
func GenderIn(vs ...int) predicate.Post {
	return predicate.Post(sql.FieldIn(FieldGender, vs...))
}

// GenderNotIn applies the NotIn predicate on the "gender" field.
func GenderNotIn(vs ...int) predicate.Post {
	return predicate.Post(sql.FieldNotIn(FieldGender, vs...))
}

// GenderGT applies the GT predicate on the "gender" field.
func GenderGT(v int) predicate.Post {
	return predicate.Post(sql.FieldGT(FieldGender, v))
}

// GenderGTE applies the GTE predicate on the "gender" field.
func GenderGTE(v int) predicate.Post {
	return predicate.Post(sql.FieldGTE(FieldGender, v))
}

// GenderLT applies the LT predicate on the "gender" field.
func GenderLT(v int) predicate.Post {
	return predicate.Post(sql.FieldLT(FieldGender, v))
}

// GenderLTE applies the LTE predicate on the "gender" field.
func GenderLTE(v int) predicate.Post {
	return predicate.Post(sql.FieldLTE(FieldGender, v))
}

// IntroductionEQ applies the EQ predicate on the "introduction" field.
func IntroductionEQ(v string) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldIntroduction, v))
}

// IntroductionNEQ applies the NEQ predicate on the "introduction" field.
func IntroductionNEQ(v string) predicate.Post {
	return predicate.Post(sql.FieldNEQ(FieldIntroduction, v))
}

// IntroductionIn applies the In predicate on the "introduction" field.
func IntroductionIn(vs ...string) predicate.Post {
	return predicate.Post(sql.FieldIn(FieldIntroduction, vs...))
}

// IntroductionNotIn applies the NotIn predicate on the "introduction" field.
func IntroductionNotIn(vs ...string) predicate.Post {
	return predicate.Post(sql.FieldNotIn(FieldIntroduction, vs...))
}

// IntroductionGT applies the GT predicate on the "introduction" field.
func IntroductionGT(v string) predicate.Post {
	return predicate.Post(sql.FieldGT(FieldIntroduction, v))
}

// IntroductionGTE applies the GTE predicate on the "introduction" field.
func IntroductionGTE(v string) predicate.Post {
	return predicate.Post(sql.FieldGTE(FieldIntroduction, v))
}

// IntroductionLT applies the LT predicate on the "introduction" field.
func IntroductionLT(v string) predicate.Post {
	return predicate.Post(sql.FieldLT(FieldIntroduction, v))
}

// IntroductionLTE applies the LTE predicate on the "introduction" field.
func IntroductionLTE(v string) predicate.Post {
	return predicate.Post(sql.FieldLTE(FieldIntroduction, v))
}

// IntroductionContains applies the Contains predicate on the "introduction" field.
func IntroductionContains(v string) predicate.Post {
	return predicate.Post(sql.FieldContains(FieldIntroduction, v))
}

// IntroductionHasPrefix applies the HasPrefix predicate on the "introduction" field.
func IntroductionHasPrefix(v string) predicate.Post {
	return predicate.Post(sql.FieldHasPrefix(FieldIntroduction, v))
}

// IntroductionHasSuffix applies the HasSuffix predicate on the "introduction" field.
func IntroductionHasSuffix(v string) predicate.Post {
	return predicate.Post(sql.FieldHasSuffix(FieldIntroduction, v))
}

// IntroductionIsNil applies the IsNil predicate on the "introduction" field.
func IntroductionIsNil() predicate.Post {
	return predicate.Post(sql.FieldIsNull(FieldIntroduction))
}

// IntroductionNotNil applies the NotNil predicate on the "introduction" field.
func IntroductionNotNil() predicate.Post {
	return predicate.Post(sql.FieldNotNull(FieldIntroduction))
}

// IntroductionEqualFold applies the EqualFold predicate on the "introduction" field.
func IntroductionEqualFold(v string) predicate.Post {
	return predicate.Post(sql.FieldEqualFold(FieldIntroduction, v))
}

// IntroductionContainsFold applies the ContainsFold predicate on the "introduction" field.
func IntroductionContainsFold(v string) predicate.Post {
	return predicate.Post(sql.FieldContainsFold(FieldIntroduction, v))
}

// TweetURLEQ applies the EQ predicate on the "tweet_url" field.
func TweetURLEQ(v string) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldTweetURL, v))
}

// TweetURLNEQ applies the NEQ predicate on the "tweet_url" field.
func TweetURLNEQ(v string) predicate.Post {
	return predicate.Post(sql.FieldNEQ(FieldTweetURL, v))
}

// TweetURLIn applies the In predicate on the "tweet_url" field.
func TweetURLIn(vs ...string) predicate.Post {
	return predicate.Post(sql.FieldIn(FieldTweetURL, vs...))
}

// TweetURLNotIn applies the NotIn predicate on the "tweet_url" field.
func TweetURLNotIn(vs ...string) predicate.Post {
	return predicate.Post(sql.FieldNotIn(FieldTweetURL, vs...))
}

// TweetURLGT applies the GT predicate on the "tweet_url" field.
func TweetURLGT(v string) predicate.Post {
	return predicate.Post(sql.FieldGT(FieldTweetURL, v))
}

// TweetURLGTE applies the GTE predicate on the "tweet_url" field.
func TweetURLGTE(v string) predicate.Post {
	return predicate.Post(sql.FieldGTE(FieldTweetURL, v))
}

// TweetURLLT applies the LT predicate on the "tweet_url" field.
func TweetURLLT(v string) predicate.Post {
	return predicate.Post(sql.FieldLT(FieldTweetURL, v))
}

// TweetURLLTE applies the LTE predicate on the "tweet_url" field.
func TweetURLLTE(v string) predicate.Post {
	return predicate.Post(sql.FieldLTE(FieldTweetURL, v))
}

// TweetURLContains applies the Contains predicate on the "tweet_url" field.
func TweetURLContains(v string) predicate.Post {
	return predicate.Post(sql.FieldContains(FieldTweetURL, v))
}

// TweetURLHasPrefix applies the HasPrefix predicate on the "tweet_url" field.
func TweetURLHasPrefix(v string) predicate.Post {
	return predicate.Post(sql.FieldHasPrefix(FieldTweetURL, v))
}

// TweetURLHasSuffix applies the HasSuffix predicate on the "tweet_url" field.
func TweetURLHasSuffix(v string) predicate.Post {
	return predicate.Post(sql.FieldHasSuffix(FieldTweetURL, v))
}

// TweetURLIsNil applies the IsNil predicate on the "tweet_url" field.
func TweetURLIsNil() predicate.Post {
	return predicate.Post(sql.FieldIsNull(FieldTweetURL))
}

// TweetURLNotNil applies the NotNil predicate on the "tweet_url" field.
func TweetURLNotNil() predicate.Post {
	return predicate.Post(sql.FieldNotNull(FieldTweetURL))
}

// TweetURLEqualFold applies the EqualFold predicate on the "tweet_url" field.
func TweetURLEqualFold(v string) predicate.Post {
	return predicate.Post(sql.FieldEqualFold(FieldTweetURL, v))
}

// TweetURLContainsFold applies the ContainsFold predicate on the "tweet_url" field.
func TweetURLContainsFold(v string) predicate.Post {
	return predicate.Post(sql.FieldContainsFold(FieldTweetURL, v))
}

// TiktokURLEQ applies the EQ predicate on the "tiktok_url" field.
func TiktokURLEQ(v string) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldTiktokURL, v))
}

// TiktokURLNEQ applies the NEQ predicate on the "tiktok_url" field.
func TiktokURLNEQ(v string) predicate.Post {
	return predicate.Post(sql.FieldNEQ(FieldTiktokURL, v))
}

// TiktokURLIn applies the In predicate on the "tiktok_url" field.
func TiktokURLIn(vs ...string) predicate.Post {
	return predicate.Post(sql.FieldIn(FieldTiktokURL, vs...))
}

// TiktokURLNotIn applies the NotIn predicate on the "tiktok_url" field.
func TiktokURLNotIn(vs ...string) predicate.Post {
	return predicate.Post(sql.FieldNotIn(FieldTiktokURL, vs...))
}

// TiktokURLGT applies the GT predicate on the "tiktok_url" field.
func TiktokURLGT(v string) predicate.Post {
	return predicate.Post(sql.FieldGT(FieldTiktokURL, v))
}

// TiktokURLGTE applies the GTE predicate on the "tiktok_url" field.
func TiktokURLGTE(v string) predicate.Post {
	return predicate.Post(sql.FieldGTE(FieldTiktokURL, v))
}

// TiktokURLLT applies the LT predicate on the "tiktok_url" field.
func TiktokURLLT(v string) predicate.Post {
	return predicate.Post(sql.FieldLT(FieldTiktokURL, v))
}

// TiktokURLLTE applies the LTE predicate on the "tiktok_url" field.
func TiktokURLLTE(v string) predicate.Post {
	return predicate.Post(sql.FieldLTE(FieldTiktokURL, v))
}

// TiktokURLContains applies the Contains predicate on the "tiktok_url" field.
func TiktokURLContains(v string) predicate.Post {
	return predicate.Post(sql.FieldContains(FieldTiktokURL, v))
}

// TiktokURLHasPrefix applies the HasPrefix predicate on the "tiktok_url" field.
func TiktokURLHasPrefix(v string) predicate.Post {
	return predicate.Post(sql.FieldHasPrefix(FieldTiktokURL, v))
}

// TiktokURLHasSuffix applies the HasSuffix predicate on the "tiktok_url" field.
func TiktokURLHasSuffix(v string) predicate.Post {
	return predicate.Post(sql.FieldHasSuffix(FieldTiktokURL, v))
}

// TiktokURLIsNil applies the IsNil predicate on the "tiktok_url" field.
func TiktokURLIsNil() predicate.Post {
	return predicate.Post(sql.FieldIsNull(FieldTiktokURL))
}

// TiktokURLNotNil applies the NotNil predicate on the "tiktok_url" field.
func TiktokURLNotNil() predicate.Post {
	return predicate.Post(sql.FieldNotNull(FieldTiktokURL))
}

// TiktokURLEqualFold applies the EqualFold predicate on the "tiktok_url" field.
func TiktokURLEqualFold(v string) predicate.Post {
	return predicate.Post(sql.FieldEqualFold(FieldTiktokURL, v))
}

// TiktokURLContainsFold applies the ContainsFold predicate on the "tiktok_url" field.
func TiktokURLContainsFold(v string) predicate.Post {
	return predicate.Post(sql.FieldContainsFold(FieldTiktokURL, v))
}

// ContactURLEQ applies the EQ predicate on the "contact_url" field.
func ContactURLEQ(v string) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldContactURL, v))
}

// ContactURLNEQ applies the NEQ predicate on the "contact_url" field.
func ContactURLNEQ(v string) predicate.Post {
	return predicate.Post(sql.FieldNEQ(FieldContactURL, v))
}

// ContactURLIn applies the In predicate on the "contact_url" field.
func ContactURLIn(vs ...string) predicate.Post {
	return predicate.Post(sql.FieldIn(FieldContactURL, vs...))
}

// ContactURLNotIn applies the NotIn predicate on the "contact_url" field.
func ContactURLNotIn(vs ...string) predicate.Post {
	return predicate.Post(sql.FieldNotIn(FieldContactURL, vs...))
}

// ContactURLGT applies the GT predicate on the "contact_url" field.
func ContactURLGT(v string) predicate.Post {
	return predicate.Post(sql.FieldGT(FieldContactURL, v))
}

// ContactURLGTE applies the GTE predicate on the "contact_url" field.
func ContactURLGTE(v string) predicate.Post {
	return predicate.Post(sql.FieldGTE(FieldContactURL, v))
}

// ContactURLLT applies the LT predicate on the "contact_url" field.
func ContactURLLT(v string) predicate.Post {
	return predicate.Post(sql.FieldLT(FieldContactURL, v))
}

// ContactURLLTE applies the LTE predicate on the "contact_url" field.
func ContactURLLTE(v string) predicate.Post {
	return predicate.Post(sql.FieldLTE(FieldContactURL, v))
}

// ContactURLContains applies the Contains predicate on the "contact_url" field.
func ContactURLContains(v string) predicate.Post {
	return predicate.Post(sql.FieldContains(FieldContactURL, v))
}

// ContactURLHasPrefix applies the HasPrefix predicate on the "contact_url" field.
func ContactURLHasPrefix(v string) predicate.Post {
	return predicate.Post(sql.FieldHasPrefix(FieldContactURL, v))
}

// ContactURLHasSuffix applies the HasSuffix predicate on the "contact_url" field.
func ContactURLHasSuffix(v string) predicate.Post {
	return predicate.Post(sql.FieldHasSuffix(FieldContactURL, v))
}

// ContactURLIsNil applies the IsNil predicate on the "contact_url" field.
func ContactURLIsNil() predicate.Post {
	return predicate.Post(sql.FieldIsNull(FieldContactURL))
}

// ContactURLNotNil applies the NotNil predicate on the "contact_url" field.
func ContactURLNotNil() predicate.Post {
	return predicate.Post(sql.FieldNotNull(FieldContactURL))
}

// ContactURLEqualFold applies the EqualFold predicate on the "contact_url" field.
func ContactURLEqualFold(v string) predicate.Post {
	return predicate.Post(sql.FieldEqualFold(FieldContactURL, v))
}

// ContactURLContainsFold applies the ContainsFold predicate on the "contact_url" field.
func ContactURLContainsFold(v string) predicate.Post {
	return predicate.Post(sql.FieldContainsFold(FieldContactURL, v))
}

// ContactIDEQ applies the EQ predicate on the "contact_id" field.
func ContactIDEQ(v string) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldContactID, v))
}

// ContactIDNEQ applies the NEQ predicate on the "contact_id" field.
func ContactIDNEQ(v string) predicate.Post {
	return predicate.Post(sql.FieldNEQ(FieldContactID, v))
}

// ContactIDIn applies the In predicate on the "contact_id" field.
func ContactIDIn(vs ...string) predicate.Post {
	return predicate.Post(sql.FieldIn(FieldContactID, vs...))
}

// ContactIDNotIn applies the NotIn predicate on the "contact_id" field.
func ContactIDNotIn(vs ...string) predicate.Post {
	return predicate.Post(sql.FieldNotIn(FieldContactID, vs...))
}

// ContactIDGT applies the GT predicate on the "contact_id" field.
func ContactIDGT(v string) predicate.Post {
	return predicate.Post(sql.FieldGT(FieldContactID, v))
}

// ContactIDGTE applies the GTE predicate on the "contact_id" field.
func ContactIDGTE(v string) predicate.Post {
	return predicate.Post(sql.FieldGTE(FieldContactID, v))
}

// ContactIDLT applies the LT predicate on the "contact_id" field.
func ContactIDLT(v string) predicate.Post {
	return predicate.Post(sql.FieldLT(FieldContactID, v))
}

// ContactIDLTE applies the LTE predicate on the "contact_id" field.
func ContactIDLTE(v string) predicate.Post {
	return predicate.Post(sql.FieldLTE(FieldContactID, v))
}

// ContactIDContains applies the Contains predicate on the "contact_id" field.
func ContactIDContains(v string) predicate.Post {
	return predicate.Post(sql.FieldContains(FieldContactID, v))
}

// ContactIDHasPrefix applies the HasPrefix predicate on the "contact_id" field.
func ContactIDHasPrefix(v string) predicate.Post {
	return predicate.Post(sql.FieldHasPrefix(FieldContactID, v))
}

// ContactIDHasSuffix applies the HasSuffix predicate on the "contact_id" field.
func ContactIDHasSuffix(v string) predicate.Post {
	return predicate.Post(sql.FieldHasSuffix(FieldContactID, v))
}

// ContactIDIsNil applies the IsNil predicate on the "contact_id" field.
func ContactIDIsNil() predicate.Post {
	return predicate.Post(sql.FieldIsNull(FieldContactID))
}

// ContactIDNotNil applies the NotNil predicate on the "contact_id" field.
func ContactIDNotNil() predicate.Post {
	return predicate.Post(sql.FieldNotNull(FieldContactID))
}

// ContactIDEqualFold applies the EqualFold predicate on the "contact_id" field.
func ContactIDEqualFold(v string) predicate.Post {
	return predicate.Post(sql.FieldEqualFold(FieldContactID, v))
}

// ContactIDContainsFold applies the ContainsFold predicate on the "contact_id" field.
func ContactIDContainsFold(v string) predicate.Post {
	return predicate.Post(sql.FieldContainsFold(FieldContactID, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Post {
	return predicate.Post(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Post {
	return predicate.Post(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Post {
	return predicate.Post(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Post {
	return predicate.Post(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Post {
	return predicate.Post(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Post {
	return predicate.Post(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Post {
	return predicate.Post(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Post {
	return predicate.Post(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Post {
	return predicate.Post(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Post {
	return predicate.Post(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Post {
	return predicate.Post(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Post {
	return predicate.Post(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Post {
	return predicate.Post(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Post {
	return predicate.Post(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Post {
	return predicate.Post(sql.FieldLTE(FieldUpdatedAt, v))
}

// HasIPAddressLog applies the HasEdge predicate on the "ip_address_log" edge.
func HasIPAddressLog() predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, IPAddressLogTable, IPAddressLogColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasIPAddressLogWith applies the HasEdge predicate on the "ip_address_log" edge with a given conditions (other predicates).
func HasIPAddressLogWith(preds ...predicate.IPAddressLog) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(IPAddressLogInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, IPAddressLogTable, IPAddressLogColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Post) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Post) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Post) predicate.Post {
	return predicate.Post(func(s *sql.Selector) {
		p(s.Not())
	})
}
