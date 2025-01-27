// Code generated by ent, DO NOT EDIT.

package entgen

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/t-yamakoshi/24-fresh-backend-v2/pkg/adapter/entgen/usermodel"
)

// UserModel is the model entity for the UserModel schema.
type UserModel struct {
	config `json:"-"`
	// ID of the ent.
	ID int64 `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// CognitoID holds the value of the "cognito_id" field.
	CognitoID string `json:"cognito_id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// UserName holds the value of the "user_name" field.
	UserName string `json:"user_name,omitempty"`
	// BirthDate holds the value of the "birth_date" field.
	BirthDate time.Time `json:"birth_date,omitempty"`
	// Gender holds the value of the "gender" field.
	Gender usermodel.Gender `json:"gender,omitempty"`
	// SelfIntroduction holds the value of the "self_introduction" field.
	SelfIntroduction string `json:"self_introduction,omitempty"`
	// ProfileImage holds the value of the "profile_image" field.
	ProfileImage string `json:"profile_image,omitempty"`
	// Email holds the value of the "email" field.
	Email string `json:"email,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserModelQuery when eager-loading is set.
	Edges        UserModelEdges `json:"edges"`
	selectValues sql.SelectValues
}

// UserModelEdges holds the relations/edges for other nodes in the graph.
type UserModelEdges struct {
	// Followers holds the value of the followers edge.
	Followers []*FollowsModel `json:"followers,omitempty"`
	// Followees holds the value of the followees edge.
	Followees []*FollowsModel `json:"followees,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// FollowersOrErr returns the Followers value or an error if the edge
// was not loaded in eager-loading.
func (e UserModelEdges) FollowersOrErr() ([]*FollowsModel, error) {
	if e.loadedTypes[0] {
		return e.Followers, nil
	}
	return nil, &NotLoadedError{edge: "followers"}
}

// FolloweesOrErr returns the Followees value or an error if the edge
// was not loaded in eager-loading.
func (e UserModelEdges) FolloweesOrErr() ([]*FollowsModel, error) {
	if e.loadedTypes[1] {
		return e.Followees, nil
	}
	return nil, &NotLoadedError{edge: "followees"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*UserModel) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case usermodel.FieldID:
			values[i] = new(sql.NullInt64)
		case usermodel.FieldCognitoID, usermodel.FieldName, usermodel.FieldUserName, usermodel.FieldGender, usermodel.FieldSelfIntroduction, usermodel.FieldProfileImage, usermodel.FieldEmail:
			values[i] = new(sql.NullString)
		case usermodel.FieldCreatedAt, usermodel.FieldUpdatedAt, usermodel.FieldBirthDate:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the UserModel fields.
func (um *UserModel) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case usermodel.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			um.ID = int64(value.Int64)
		case usermodel.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				um.CreatedAt = value.Time
			}
		case usermodel.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				um.UpdatedAt = value.Time
			}
		case usermodel.FieldCognitoID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field cognito_id", values[i])
			} else if value.Valid {
				um.CognitoID = value.String
			}
		case usermodel.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				um.Name = value.String
			}
		case usermodel.FieldUserName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field user_name", values[i])
			} else if value.Valid {
				um.UserName = value.String
			}
		case usermodel.FieldBirthDate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field birth_date", values[i])
			} else if value.Valid {
				um.BirthDate = value.Time
			}
		case usermodel.FieldGender:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field gender", values[i])
			} else if value.Valid {
				um.Gender = usermodel.Gender(value.String)
			}
		case usermodel.FieldSelfIntroduction:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field self_introduction", values[i])
			} else if value.Valid {
				um.SelfIntroduction = value.String
			}
		case usermodel.FieldProfileImage:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field profile_image", values[i])
			} else if value.Valid {
				um.ProfileImage = value.String
			}
		case usermodel.FieldEmail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field email", values[i])
			} else if value.Valid {
				um.Email = value.String
			}
		default:
			um.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the UserModel.
// This includes values selected through modifiers, order, etc.
func (um *UserModel) Value(name string) (ent.Value, error) {
	return um.selectValues.Get(name)
}

// QueryFollowers queries the "followers" edge of the UserModel entity.
func (um *UserModel) QueryFollowers() *FollowsModelQuery {
	return NewUserModelClient(um.config).QueryFollowers(um)
}

// QueryFollowees queries the "followees" edge of the UserModel entity.
func (um *UserModel) QueryFollowees() *FollowsModelQuery {
	return NewUserModelClient(um.config).QueryFollowees(um)
}

// Update returns a builder for updating this UserModel.
// Note that you need to call UserModel.Unwrap() before calling this method if this UserModel
// was returned from a transaction, and the transaction was committed or rolled back.
func (um *UserModel) Update() *UserModelUpdateOne {
	return NewUserModelClient(um.config).UpdateOne(um)
}

// Unwrap unwraps the UserModel entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (um *UserModel) Unwrap() *UserModel {
	_tx, ok := um.config.driver.(*txDriver)
	if !ok {
		panic("entgen: UserModel is not a transactional entity")
	}
	um.config.driver = _tx.drv
	return um
}

// String implements the fmt.Stringer.
func (um *UserModel) String() string {
	var builder strings.Builder
	builder.WriteString("UserModel(")
	builder.WriteString(fmt.Sprintf("id=%v, ", um.ID))
	builder.WriteString("created_at=")
	builder.WriteString(um.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(um.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("cognito_id=")
	builder.WriteString(um.CognitoID)
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(um.Name)
	builder.WriteString(", ")
	builder.WriteString("user_name=")
	builder.WriteString(um.UserName)
	builder.WriteString(", ")
	builder.WriteString("birth_date=")
	builder.WriteString(um.BirthDate.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("gender=")
	builder.WriteString(fmt.Sprintf("%v", um.Gender))
	builder.WriteString(", ")
	builder.WriteString("self_introduction=")
	builder.WriteString(um.SelfIntroduction)
	builder.WriteString(", ")
	builder.WriteString("profile_image=")
	builder.WriteString(um.ProfileImage)
	builder.WriteString(", ")
	builder.WriteString("email=")
	builder.WriteString(um.Email)
	builder.WriteByte(')')
	return builder.String()
}

// UserModels is a parsable slice of UserModel.
type UserModels []*UserModel
