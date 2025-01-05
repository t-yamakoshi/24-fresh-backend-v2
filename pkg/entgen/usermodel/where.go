// Code generated by ent, DO NOT EDIT.

package usermodel

import (
	"entgo.io/ent/dialect/sql"
	"github.com/t-yamakoshi/24-fresh-backend-v2/pkg/entgen/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.UserModel {
	return predicate.UserModel(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.UserModel {
	return predicate.UserModel(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.UserModel {
	return predicate.UserModel(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.UserModel {
	return predicate.UserModel(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.UserModel {
	return predicate.UserModel(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.UserModel {
	return predicate.UserModel(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.UserModel {
	return predicate.UserModel(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.UserModel {
	return predicate.UserModel(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.UserModel {
	return predicate.UserModel(sql.FieldLTE(FieldID, id))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.UserModel {
	return predicate.UserModel(sql.FieldEQ(FieldName, v))
}

// UserName applies equality check predicate on the "user_name" field. It's identical to UserNameEQ.
func UserName(v string) predicate.UserModel {
	return predicate.UserModel(sql.FieldEQ(FieldUserName, v))
}

// Email applies equality check predicate on the "email" field. It's identical to EmailEQ.
func Email(v string) predicate.UserModel {
	return predicate.UserModel(sql.FieldEQ(FieldEmail, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.UserModel {
	return predicate.UserModel(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.UserModel {
	return predicate.UserModel(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.UserModel {
	return predicate.UserModel(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.UserModel {
	return predicate.UserModel(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.UserModel {
	return predicate.UserModel(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.UserModel {
	return predicate.UserModel(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.UserModel {
	return predicate.UserModel(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.UserModel {
	return predicate.UserModel(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.UserModel {
	return predicate.UserModel(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.UserModel {
	return predicate.UserModel(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.UserModel {
	return predicate.UserModel(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.UserModel {
	return predicate.UserModel(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.UserModel {
	return predicate.UserModel(sql.FieldContainsFold(FieldName, v))
}

// UserNameEQ applies the EQ predicate on the "user_name" field.
func UserNameEQ(v string) predicate.UserModel {
	return predicate.UserModel(sql.FieldEQ(FieldUserName, v))
}

// UserNameNEQ applies the NEQ predicate on the "user_name" field.
func UserNameNEQ(v string) predicate.UserModel {
	return predicate.UserModel(sql.FieldNEQ(FieldUserName, v))
}

// UserNameIn applies the In predicate on the "user_name" field.
func UserNameIn(vs ...string) predicate.UserModel {
	return predicate.UserModel(sql.FieldIn(FieldUserName, vs...))
}

// UserNameNotIn applies the NotIn predicate on the "user_name" field.
func UserNameNotIn(vs ...string) predicate.UserModel {
	return predicate.UserModel(sql.FieldNotIn(FieldUserName, vs...))
}

// UserNameGT applies the GT predicate on the "user_name" field.
func UserNameGT(v string) predicate.UserModel {
	return predicate.UserModel(sql.FieldGT(FieldUserName, v))
}

// UserNameGTE applies the GTE predicate on the "user_name" field.
func UserNameGTE(v string) predicate.UserModel {
	return predicate.UserModel(sql.FieldGTE(FieldUserName, v))
}

// UserNameLT applies the LT predicate on the "user_name" field.
func UserNameLT(v string) predicate.UserModel {
	return predicate.UserModel(sql.FieldLT(FieldUserName, v))
}

// UserNameLTE applies the LTE predicate on the "user_name" field.
func UserNameLTE(v string) predicate.UserModel {
	return predicate.UserModel(sql.FieldLTE(FieldUserName, v))
}

// UserNameContains applies the Contains predicate on the "user_name" field.
func UserNameContains(v string) predicate.UserModel {
	return predicate.UserModel(sql.FieldContains(FieldUserName, v))
}

// UserNameHasPrefix applies the HasPrefix predicate on the "user_name" field.
func UserNameHasPrefix(v string) predicate.UserModel {
	return predicate.UserModel(sql.FieldHasPrefix(FieldUserName, v))
}

// UserNameHasSuffix applies the HasSuffix predicate on the "user_name" field.
func UserNameHasSuffix(v string) predicate.UserModel {
	return predicate.UserModel(sql.FieldHasSuffix(FieldUserName, v))
}

// UserNameEqualFold applies the EqualFold predicate on the "user_name" field.
func UserNameEqualFold(v string) predicate.UserModel {
	return predicate.UserModel(sql.FieldEqualFold(FieldUserName, v))
}

// UserNameContainsFold applies the ContainsFold predicate on the "user_name" field.
func UserNameContainsFold(v string) predicate.UserModel {
	return predicate.UserModel(sql.FieldContainsFold(FieldUserName, v))
}

// EmailEQ applies the EQ predicate on the "email" field.
func EmailEQ(v string) predicate.UserModel {
	return predicate.UserModel(sql.FieldEQ(FieldEmail, v))
}

// EmailNEQ applies the NEQ predicate on the "email" field.
func EmailNEQ(v string) predicate.UserModel {
	return predicate.UserModel(sql.FieldNEQ(FieldEmail, v))
}

// EmailIn applies the In predicate on the "email" field.
func EmailIn(vs ...string) predicate.UserModel {
	return predicate.UserModel(sql.FieldIn(FieldEmail, vs...))
}

// EmailNotIn applies the NotIn predicate on the "email" field.
func EmailNotIn(vs ...string) predicate.UserModel {
	return predicate.UserModel(sql.FieldNotIn(FieldEmail, vs...))
}

// EmailGT applies the GT predicate on the "email" field.
func EmailGT(v string) predicate.UserModel {
	return predicate.UserModel(sql.FieldGT(FieldEmail, v))
}

// EmailGTE applies the GTE predicate on the "email" field.
func EmailGTE(v string) predicate.UserModel {
	return predicate.UserModel(sql.FieldGTE(FieldEmail, v))
}

// EmailLT applies the LT predicate on the "email" field.
func EmailLT(v string) predicate.UserModel {
	return predicate.UserModel(sql.FieldLT(FieldEmail, v))
}

// EmailLTE applies the LTE predicate on the "email" field.
func EmailLTE(v string) predicate.UserModel {
	return predicate.UserModel(sql.FieldLTE(FieldEmail, v))
}

// EmailContains applies the Contains predicate on the "email" field.
func EmailContains(v string) predicate.UserModel {
	return predicate.UserModel(sql.FieldContains(FieldEmail, v))
}

// EmailHasPrefix applies the HasPrefix predicate on the "email" field.
func EmailHasPrefix(v string) predicate.UserModel {
	return predicate.UserModel(sql.FieldHasPrefix(FieldEmail, v))
}

// EmailHasSuffix applies the HasSuffix predicate on the "email" field.
func EmailHasSuffix(v string) predicate.UserModel {
	return predicate.UserModel(sql.FieldHasSuffix(FieldEmail, v))
}

// EmailEqualFold applies the EqualFold predicate on the "email" field.
func EmailEqualFold(v string) predicate.UserModel {
	return predicate.UserModel(sql.FieldEqualFold(FieldEmail, v))
}

// EmailContainsFold applies the ContainsFold predicate on the "email" field.
func EmailContainsFold(v string) predicate.UserModel {
	return predicate.UserModel(sql.FieldContainsFold(FieldEmail, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.UserModel) predicate.UserModel {
	return predicate.UserModel(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.UserModel) predicate.UserModel {
	return predicate.UserModel(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.UserModel) predicate.UserModel {
	return predicate.UserModel(sql.NotPredicates(p))
}