// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// ChainsColumns holds the columns for the "chains" table.
	ChainsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "name", Type: field.TypeString, Unique: true},
		{Name: "display_name", Type: field.TypeString, Unique: true},
	}
	// ChainsTable holds the schema information for the "chains" table.
	ChainsTable = &schema.Table{
		Name:       "chains",
		Columns:    ChainsColumns,
		PrimaryKey: []*schema.Column{ChainsColumns[0]},
	}
	// ProposalsColumns holds the columns for the "proposals" table.
	ProposalsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "proposal_id", Type: field.TypeInt},
		{Name: "title", Type: field.TypeString},
		{Name: "description", Type: field.TypeString},
		{Name: "voting_start_time", Type: field.TypeTime},
		{Name: "voting_end_time", Type: field.TypeTime},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"PROPOSAL_STATUS_UNSPECIFIED", "PROPOSAL_STATUS_DEPOSIT_PERIOD", "PROPOSAL_STATUS_VOTING_PERIOD", "PROPOSAL_STATUS_PASSED", "PROPOSAL_STATUS_REJECTED", "PROPOSAL_STATUS_FAILED"}},
		{Name: "chain_proposals", Type: field.TypeInt, Nullable: true},
	}
	// ProposalsTable holds the schema information for the "proposals" table.
	ProposalsTable = &schema.Table{
		Name:       "proposals",
		Columns:    ProposalsColumns,
		PrimaryKey: []*schema.Column{ProposalsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "proposals_chains_proposals",
				Columns:    []*schema.Column{ProposalsColumns[9]},
				RefColumns: []*schema.Column{ChainsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "chat_id", Type: field.TypeInt64, Unique: true},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// UserChainsColumns holds the columns for the "user_chains" table.
	UserChainsColumns = []*schema.Column{
		{Name: "user_id", Type: field.TypeInt},
		{Name: "chain_id", Type: field.TypeInt},
	}
	// UserChainsTable holds the schema information for the "user_chains" table.
	UserChainsTable = &schema.Table{
		Name:       "user_chains",
		Columns:    UserChainsColumns,
		PrimaryKey: []*schema.Column{UserChainsColumns[0], UserChainsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "user_chains_user_id",
				Columns:    []*schema.Column{UserChainsColumns[0]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "user_chains_chain_id",
				Columns:    []*schema.Column{UserChainsColumns[1]},
				RefColumns: []*schema.Column{ChainsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		ChainsTable,
		ProposalsTable,
		UsersTable,
		UserChainsTable,
	}
)

func init() {
	ProposalsTable.ForeignKeys[0].RefTable = ChainsTable
	UserChainsTable.ForeignKeys[0].RefTable = UsersTable
	UserChainsTable.ForeignKeys[1].RefTable = ChainsTable
}
