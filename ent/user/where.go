// Code generated by ent, DO NOT EDIT.

package user

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/tiagoposse/connect/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.User {
	return predicate.User(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.User {
	return predicate.User(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldID, id))
}

// IDEqualFold applies the EqualFold predicate on the ID field.
func IDEqualFold(id string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldID, id))
}

// IDContainsFold applies the ContainsFold predicate on the ID field.
func IDContainsFold(id string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldID, id))
}

// Email applies equality check predicate on the "email" field. It's identical to EmailEQ.
func Email(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldEmail, v))
}

// Firstname applies equality check predicate on the "firstname" field. It's identical to FirstnameEQ.
func Firstname(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldFirstname, v))
}

// Lastname applies equality check predicate on the "lastname" field. It's identical to LastnameEQ.
func Lastname(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldLastname, v))
}

// Provider applies equality check predicate on the "provider" field. It's identical to ProviderEQ.
func Provider(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldProvider, v))
}

// Password applies equality check predicate on the "password" field. It's identical to PasswordEQ.
func Password(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldPassword, v))
}

// Salt applies equality check predicate on the "salt" field. It's identical to SaltEQ.
func Salt(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldSalt, v))
}

// PhotoURL applies equality check predicate on the "photo_url" field. It's identical to PhotoURLEQ.
func PhotoURL(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldPhotoURL, v))
}

// Disabled applies equality check predicate on the "disabled" field. It's identical to DisabledEQ.
func Disabled(v bool) predicate.User {
	return predicate.User(sql.FieldEQ(FieldDisabled, v))
}

// DisabledReason applies equality check predicate on the "disabled_reason" field. It's identical to DisabledReasonEQ.
func DisabledReason(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldDisabledReason, v))
}

// GroupID applies equality check predicate on the "group_id" field. It's identical to GroupIDEQ.
func GroupID(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldGroupID, v))
}

// EmailEQ applies the EQ predicate on the "email" field.
func EmailEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldEmail, v))
}

// EmailNEQ applies the NEQ predicate on the "email" field.
func EmailNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldEmail, v))
}

// EmailIn applies the In predicate on the "email" field.
func EmailIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldEmail, vs...))
}

// EmailNotIn applies the NotIn predicate on the "email" field.
func EmailNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldEmail, vs...))
}

// EmailGT applies the GT predicate on the "email" field.
func EmailGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldEmail, v))
}

// EmailGTE applies the GTE predicate on the "email" field.
func EmailGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldEmail, v))
}

// EmailLT applies the LT predicate on the "email" field.
func EmailLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldEmail, v))
}

// EmailLTE applies the LTE predicate on the "email" field.
func EmailLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldEmail, v))
}

// EmailContains applies the Contains predicate on the "email" field.
func EmailContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldEmail, v))
}

// EmailHasPrefix applies the HasPrefix predicate on the "email" field.
func EmailHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldEmail, v))
}

// EmailHasSuffix applies the HasSuffix predicate on the "email" field.
func EmailHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldEmail, v))
}

// EmailEqualFold applies the EqualFold predicate on the "email" field.
func EmailEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldEmail, v))
}

// EmailContainsFold applies the ContainsFold predicate on the "email" field.
func EmailContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldEmail, v))
}

// FirstnameEQ applies the EQ predicate on the "firstname" field.
func FirstnameEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldFirstname, v))
}

// FirstnameNEQ applies the NEQ predicate on the "firstname" field.
func FirstnameNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldFirstname, v))
}

// FirstnameIn applies the In predicate on the "firstname" field.
func FirstnameIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldFirstname, vs...))
}

// FirstnameNotIn applies the NotIn predicate on the "firstname" field.
func FirstnameNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldFirstname, vs...))
}

// FirstnameGT applies the GT predicate on the "firstname" field.
func FirstnameGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldFirstname, v))
}

// FirstnameGTE applies the GTE predicate on the "firstname" field.
func FirstnameGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldFirstname, v))
}

// FirstnameLT applies the LT predicate on the "firstname" field.
func FirstnameLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldFirstname, v))
}

// FirstnameLTE applies the LTE predicate on the "firstname" field.
func FirstnameLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldFirstname, v))
}

// FirstnameContains applies the Contains predicate on the "firstname" field.
func FirstnameContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldFirstname, v))
}

// FirstnameHasPrefix applies the HasPrefix predicate on the "firstname" field.
func FirstnameHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldFirstname, v))
}

// FirstnameHasSuffix applies the HasSuffix predicate on the "firstname" field.
func FirstnameHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldFirstname, v))
}

// FirstnameEqualFold applies the EqualFold predicate on the "firstname" field.
func FirstnameEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldFirstname, v))
}

// FirstnameContainsFold applies the ContainsFold predicate on the "firstname" field.
func FirstnameContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldFirstname, v))
}

// LastnameEQ applies the EQ predicate on the "lastname" field.
func LastnameEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldLastname, v))
}

// LastnameNEQ applies the NEQ predicate on the "lastname" field.
func LastnameNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldLastname, v))
}

// LastnameIn applies the In predicate on the "lastname" field.
func LastnameIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldLastname, vs...))
}

// LastnameNotIn applies the NotIn predicate on the "lastname" field.
func LastnameNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldLastname, vs...))
}

// LastnameGT applies the GT predicate on the "lastname" field.
func LastnameGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldLastname, v))
}

// LastnameGTE applies the GTE predicate on the "lastname" field.
func LastnameGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldLastname, v))
}

// LastnameLT applies the LT predicate on the "lastname" field.
func LastnameLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldLastname, v))
}

// LastnameLTE applies the LTE predicate on the "lastname" field.
func LastnameLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldLastname, v))
}

// LastnameContains applies the Contains predicate on the "lastname" field.
func LastnameContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldLastname, v))
}

// LastnameHasPrefix applies the HasPrefix predicate on the "lastname" field.
func LastnameHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldLastname, v))
}

// LastnameHasSuffix applies the HasSuffix predicate on the "lastname" field.
func LastnameHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldLastname, v))
}

// LastnameEqualFold applies the EqualFold predicate on the "lastname" field.
func LastnameEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldLastname, v))
}

// LastnameContainsFold applies the ContainsFold predicate on the "lastname" field.
func LastnameContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldLastname, v))
}

// ProviderEQ applies the EQ predicate on the "provider" field.
func ProviderEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldProvider, v))
}

// ProviderNEQ applies the NEQ predicate on the "provider" field.
func ProviderNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldProvider, v))
}

// ProviderIn applies the In predicate on the "provider" field.
func ProviderIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldProvider, vs...))
}

// ProviderNotIn applies the NotIn predicate on the "provider" field.
func ProviderNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldProvider, vs...))
}

// ProviderGT applies the GT predicate on the "provider" field.
func ProviderGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldProvider, v))
}

// ProviderGTE applies the GTE predicate on the "provider" field.
func ProviderGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldProvider, v))
}

// ProviderLT applies the LT predicate on the "provider" field.
func ProviderLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldProvider, v))
}

// ProviderLTE applies the LTE predicate on the "provider" field.
func ProviderLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldProvider, v))
}

// ProviderContains applies the Contains predicate on the "provider" field.
func ProviderContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldProvider, v))
}

// ProviderHasPrefix applies the HasPrefix predicate on the "provider" field.
func ProviderHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldProvider, v))
}

// ProviderHasSuffix applies the HasSuffix predicate on the "provider" field.
func ProviderHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldProvider, v))
}

// ProviderEqualFold applies the EqualFold predicate on the "provider" field.
func ProviderEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldProvider, v))
}

// ProviderContainsFold applies the ContainsFold predicate on the "provider" field.
func ProviderContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldProvider, v))
}

// PasswordEQ applies the EQ predicate on the "password" field.
func PasswordEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldPassword, v))
}

// PasswordNEQ applies the NEQ predicate on the "password" field.
func PasswordNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldPassword, v))
}

// PasswordIn applies the In predicate on the "password" field.
func PasswordIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldPassword, vs...))
}

// PasswordNotIn applies the NotIn predicate on the "password" field.
func PasswordNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldPassword, vs...))
}

// PasswordGT applies the GT predicate on the "password" field.
func PasswordGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldPassword, v))
}

// PasswordGTE applies the GTE predicate on the "password" field.
func PasswordGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldPassword, v))
}

// PasswordLT applies the LT predicate on the "password" field.
func PasswordLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldPassword, v))
}

// PasswordLTE applies the LTE predicate on the "password" field.
func PasswordLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldPassword, v))
}

// PasswordContains applies the Contains predicate on the "password" field.
func PasswordContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldPassword, v))
}

// PasswordHasPrefix applies the HasPrefix predicate on the "password" field.
func PasswordHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldPassword, v))
}

// PasswordHasSuffix applies the HasSuffix predicate on the "password" field.
func PasswordHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldPassword, v))
}

// PasswordIsNil applies the IsNil predicate on the "password" field.
func PasswordIsNil() predicate.User {
	return predicate.User(sql.FieldIsNull(FieldPassword))
}

// PasswordNotNil applies the NotNil predicate on the "password" field.
func PasswordNotNil() predicate.User {
	return predicate.User(sql.FieldNotNull(FieldPassword))
}

// PasswordEqualFold applies the EqualFold predicate on the "password" field.
func PasswordEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldPassword, v))
}

// PasswordContainsFold applies the ContainsFold predicate on the "password" field.
func PasswordContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldPassword, v))
}

// SaltEQ applies the EQ predicate on the "salt" field.
func SaltEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldSalt, v))
}

// SaltNEQ applies the NEQ predicate on the "salt" field.
func SaltNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldSalt, v))
}

// SaltIn applies the In predicate on the "salt" field.
func SaltIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldSalt, vs...))
}

// SaltNotIn applies the NotIn predicate on the "salt" field.
func SaltNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldSalt, vs...))
}

// SaltGT applies the GT predicate on the "salt" field.
func SaltGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldSalt, v))
}

// SaltGTE applies the GTE predicate on the "salt" field.
func SaltGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldSalt, v))
}

// SaltLT applies the LT predicate on the "salt" field.
func SaltLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldSalt, v))
}

// SaltLTE applies the LTE predicate on the "salt" field.
func SaltLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldSalt, v))
}

// SaltContains applies the Contains predicate on the "salt" field.
func SaltContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldSalt, v))
}

// SaltHasPrefix applies the HasPrefix predicate on the "salt" field.
func SaltHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldSalt, v))
}

// SaltHasSuffix applies the HasSuffix predicate on the "salt" field.
func SaltHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldSalt, v))
}

// SaltIsNil applies the IsNil predicate on the "salt" field.
func SaltIsNil() predicate.User {
	return predicate.User(sql.FieldIsNull(FieldSalt))
}

// SaltNotNil applies the NotNil predicate on the "salt" field.
func SaltNotNil() predicate.User {
	return predicate.User(sql.FieldNotNull(FieldSalt))
}

// SaltEqualFold applies the EqualFold predicate on the "salt" field.
func SaltEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldSalt, v))
}

// SaltContainsFold applies the ContainsFold predicate on the "salt" field.
func SaltContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldSalt, v))
}

// PhotoURLEQ applies the EQ predicate on the "photo_url" field.
func PhotoURLEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldPhotoURL, v))
}

// PhotoURLNEQ applies the NEQ predicate on the "photo_url" field.
func PhotoURLNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldPhotoURL, v))
}

// PhotoURLIn applies the In predicate on the "photo_url" field.
func PhotoURLIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldPhotoURL, vs...))
}

// PhotoURLNotIn applies the NotIn predicate on the "photo_url" field.
func PhotoURLNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldPhotoURL, vs...))
}

// PhotoURLGT applies the GT predicate on the "photo_url" field.
func PhotoURLGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldPhotoURL, v))
}

// PhotoURLGTE applies the GTE predicate on the "photo_url" field.
func PhotoURLGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldPhotoURL, v))
}

// PhotoURLLT applies the LT predicate on the "photo_url" field.
func PhotoURLLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldPhotoURL, v))
}

// PhotoURLLTE applies the LTE predicate on the "photo_url" field.
func PhotoURLLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldPhotoURL, v))
}

// PhotoURLContains applies the Contains predicate on the "photo_url" field.
func PhotoURLContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldPhotoURL, v))
}

// PhotoURLHasPrefix applies the HasPrefix predicate on the "photo_url" field.
func PhotoURLHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldPhotoURL, v))
}

// PhotoURLHasSuffix applies the HasSuffix predicate on the "photo_url" field.
func PhotoURLHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldPhotoURL, v))
}

// PhotoURLIsNil applies the IsNil predicate on the "photo_url" field.
func PhotoURLIsNil() predicate.User {
	return predicate.User(sql.FieldIsNull(FieldPhotoURL))
}

// PhotoURLNotNil applies the NotNil predicate on the "photo_url" field.
func PhotoURLNotNil() predicate.User {
	return predicate.User(sql.FieldNotNull(FieldPhotoURL))
}

// PhotoURLEqualFold applies the EqualFold predicate on the "photo_url" field.
func PhotoURLEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldPhotoURL, v))
}

// PhotoURLContainsFold applies the ContainsFold predicate on the "photo_url" field.
func PhotoURLContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldPhotoURL, v))
}

// DisabledEQ applies the EQ predicate on the "disabled" field.
func DisabledEQ(v bool) predicate.User {
	return predicate.User(sql.FieldEQ(FieldDisabled, v))
}

// DisabledNEQ applies the NEQ predicate on the "disabled" field.
func DisabledNEQ(v bool) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldDisabled, v))
}

// DisabledReasonEQ applies the EQ predicate on the "disabled_reason" field.
func DisabledReasonEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldDisabledReason, v))
}

// DisabledReasonNEQ applies the NEQ predicate on the "disabled_reason" field.
func DisabledReasonNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldDisabledReason, v))
}

// DisabledReasonIn applies the In predicate on the "disabled_reason" field.
func DisabledReasonIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldDisabledReason, vs...))
}

// DisabledReasonNotIn applies the NotIn predicate on the "disabled_reason" field.
func DisabledReasonNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldDisabledReason, vs...))
}

// DisabledReasonGT applies the GT predicate on the "disabled_reason" field.
func DisabledReasonGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldDisabledReason, v))
}

// DisabledReasonGTE applies the GTE predicate on the "disabled_reason" field.
func DisabledReasonGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldDisabledReason, v))
}

// DisabledReasonLT applies the LT predicate on the "disabled_reason" field.
func DisabledReasonLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldDisabledReason, v))
}

// DisabledReasonLTE applies the LTE predicate on the "disabled_reason" field.
func DisabledReasonLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldDisabledReason, v))
}

// DisabledReasonContains applies the Contains predicate on the "disabled_reason" field.
func DisabledReasonContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldDisabledReason, v))
}

// DisabledReasonHasPrefix applies the HasPrefix predicate on the "disabled_reason" field.
func DisabledReasonHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldDisabledReason, v))
}

// DisabledReasonHasSuffix applies the HasSuffix predicate on the "disabled_reason" field.
func DisabledReasonHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldDisabledReason, v))
}

// DisabledReasonIsNil applies the IsNil predicate on the "disabled_reason" field.
func DisabledReasonIsNil() predicate.User {
	return predicate.User(sql.FieldIsNull(FieldDisabledReason))
}

// DisabledReasonNotNil applies the NotNil predicate on the "disabled_reason" field.
func DisabledReasonNotNil() predicate.User {
	return predicate.User(sql.FieldNotNull(FieldDisabledReason))
}

// DisabledReasonEqualFold applies the EqualFold predicate on the "disabled_reason" field.
func DisabledReasonEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldDisabledReason, v))
}

// DisabledReasonContainsFold applies the ContainsFold predicate on the "disabled_reason" field.
func DisabledReasonContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldDisabledReason, v))
}

// GroupIDEQ applies the EQ predicate on the "group_id" field.
func GroupIDEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldGroupID, v))
}

// GroupIDNEQ applies the NEQ predicate on the "group_id" field.
func GroupIDNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldGroupID, v))
}

// GroupIDIn applies the In predicate on the "group_id" field.
func GroupIDIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldGroupID, vs...))
}

// GroupIDNotIn applies the NotIn predicate on the "group_id" field.
func GroupIDNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldGroupID, vs...))
}

// GroupIDGT applies the GT predicate on the "group_id" field.
func GroupIDGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldGroupID, v))
}

// GroupIDGTE applies the GTE predicate on the "group_id" field.
func GroupIDGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldGroupID, v))
}

// GroupIDLT applies the LT predicate on the "group_id" field.
func GroupIDLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldGroupID, v))
}

// GroupIDLTE applies the LTE predicate on the "group_id" field.
func GroupIDLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldGroupID, v))
}

// GroupIDContains applies the Contains predicate on the "group_id" field.
func GroupIDContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldGroupID, v))
}

// GroupIDHasPrefix applies the HasPrefix predicate on the "group_id" field.
func GroupIDHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldGroupID, v))
}

// GroupIDHasSuffix applies the HasSuffix predicate on the "group_id" field.
func GroupIDHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldGroupID, v))
}

// GroupIDEqualFold applies the EqualFold predicate on the "group_id" field.
func GroupIDEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldGroupID, v))
}

// GroupIDContainsFold applies the ContainsFold predicate on the "group_id" field.
func GroupIDContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldGroupID, v))
}

// HasDevices applies the HasEdge predicate on the "devices" edge.
func HasDevices() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, DevicesTable, DevicesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasDevicesWith applies the HasEdge predicate on the "devices" edge with a given conditions (other predicates).
func HasDevicesWith(preds ...predicate.Device) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := newDevicesStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasKeys applies the HasEdge predicate on the "keys" edge.
func HasKeys() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, KeysTable, KeysColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasKeysWith applies the HasEdge predicate on the "keys" edge with a given conditions (other predicates).
func HasKeysWith(preds ...predicate.ApiKey) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := newKeysStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasAudit applies the HasEdge predicate on the "audit" edge.
func HasAudit() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, AuditTable, AuditColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasAuditWith applies the HasEdge predicate on the "audit" edge with a given conditions (other predicates).
func HasAuditWith(preds ...predicate.Audit) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := newAuditStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasGroup applies the HasEdge predicate on the "group" edge.
func HasGroup() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, GroupTable, GroupColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasGroupWith applies the HasEdge predicate on the "group" edge with a given conditions (other predicates).
func HasGroupWith(preds ...predicate.Group) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := newGroupStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.User) predicate.User {
	return predicate.User(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.User) predicate.User {
	return predicate.User(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.User) predicate.User {
	return predicate.User(sql.NotPredicates(p))
}
