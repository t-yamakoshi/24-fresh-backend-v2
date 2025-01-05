// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// UserModelsColumns holds the columns for the "user_models" table.
	UserModelsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "user_name", Type: field.TypeString},
		{Name: "email", Type: field.TypeString},
	}
	// UserModelsTable holds the schema information for the "user_models" table.
	UserModelsTable = &schema.Table{
		Name:       "user_models",
		Columns:    UserModelsColumns,
		PrimaryKey: []*schema.Column{UserModelsColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		UserModelsTable,
	}
)

func init() {
}