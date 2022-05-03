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
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "chain_id", Type: field.TypeString, Unique: true},
		{Name: "account_prefix", Type: field.TypeString},
		{Name: "name", Type: field.TypeString, Unique: true},
		{Name: "display_name", Type: field.TypeString, Unique: true},
		{Name: "is_enabled", Type: field.TypeBool, Default: true},
	}
	// ChainsTable holds the schema information for the "chains" table.
	ChainsTable = &schema.Table{
		Name:       "chains",
		Columns:    ChainsColumns,
		PrimaryKey: []*schema.Column{ChainsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "chain_name",
				Unique:  true,
				Columns: []*schema.Column{ChainsColumns[5]},
			},
		},
	}
	// DiscordChannelsColumns holds the columns for the "discord_channels" table.
	DiscordChannelsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "channel_id", Type: field.TypeInt64, Unique: true},
		{Name: "name", Type: field.TypeString},
		{Name: "is_group", Type: field.TypeBool},
		{Name: "roles", Type: field.TypeString, Default: ""},
		{Name: "discord_channel_user", Type: field.TypeInt, Nullable: true},
	}
	// DiscordChannelsTable holds the schema information for the "discord_channels" table.
	DiscordChannelsTable = &schema.Table{
		Name:       "discord_channels",
		Columns:    DiscordChannelsColumns,
		PrimaryKey: []*schema.Column{DiscordChannelsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "discord_channels_users_user",
				Columns:    []*schema.Column{DiscordChannelsColumns[7]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "discordchannel_id",
				Unique:  true,
				Columns: []*schema.Column{DiscordChannelsColumns[0]},
			},
		},
	}
	// GrantsColumns holds the columns for the "grants" table.
	GrantsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "grantee", Type: field.TypeString},
		{Name: "type", Type: field.TypeString},
		{Name: "expires_at", Type: field.TypeTime},
		{Name: "wallet_grants", Type: field.TypeInt, Nullable: true},
	}
	// GrantsTable holds the schema information for the "grants" table.
	GrantsTable = &schema.Table{
		Name:       "grants",
		Columns:    GrantsColumns,
		PrimaryKey: []*schema.Column{GrantsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "grants_wallets_grants",
				Columns:    []*schema.Column{GrantsColumns[6]},
				RefColumns: []*schema.Column{WalletsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// LensChainInfosColumns holds the columns for the "lens_chain_infos" table.
	LensChainInfosColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "name", Type: field.TypeString, Unique: true},
		{Name: "cnt_errors", Type: field.TypeInt},
	}
	// LensChainInfosTable holds the schema information for the "lens_chain_infos" table.
	LensChainInfosTable = &schema.Table{
		Name:       "lens_chain_infos",
		Columns:    LensChainInfosColumns,
		PrimaryKey: []*schema.Column{LensChainInfosColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "lenschaininfo_name",
				Unique:  true,
				Columns: []*schema.Column{LensChainInfosColumns[3]},
			},
		},
	}
	// MigrationInfosColumns holds the columns for the "migration_infos" table.
	MigrationInfosColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "is_migrated", Type: field.TypeBool, Default: false},
	}
	// MigrationInfosTable holds the schema information for the "migration_infos" table.
	MigrationInfosTable = &schema.Table{
		Name:       "migration_infos",
		Columns:    MigrationInfosColumns,
		PrimaryKey: []*schema.Column{MigrationInfosColumns[0]},
	}
	// ProposalsColumns holds the columns for the "proposals" table.
	ProposalsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "proposal_id", Type: field.TypeUint64},
		{Name: "title", Type: field.TypeString},
		{Name: "description", Type: field.TypeString},
		{Name: "voting_start_time", Type: field.TypeTime},
		{Name: "voting_end_time", Type: field.TypeTime},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"PROPOSAL_STATUS_FAILED", "PROPOSAL_STATUS_UNSPECIFIED", "PROPOSAL_STATUS_DEPOSIT_PERIOD", "PROPOSAL_STATUS_VOTING_PERIOD", "PROPOSAL_STATUS_PASSED", "PROPOSAL_STATUS_REJECTED"}},
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
				OnDelete:   schema.Cascade,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "proposal_proposal_id_chain_proposals",
				Unique:  true,
				Columns: []*schema.Column{ProposalsColumns[3], ProposalsColumns[9]},
			},
		},
	}
	// RPCEndpointsColumns holds the columns for the "rpc_endpoints" table.
	RPCEndpointsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "endpoint", Type: field.TypeString, Unique: true},
		{Name: "chain_rpc_endpoints", Type: field.TypeInt, Nullable: true},
	}
	// RPCEndpointsTable holds the schema information for the "rpc_endpoints" table.
	RPCEndpointsTable = &schema.Table{
		Name:       "rpc_endpoints",
		Columns:    RPCEndpointsColumns,
		PrimaryKey: []*schema.Column{RPCEndpointsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "rpc_endpoints_chains_rpc_endpoints",
				Columns:    []*schema.Column{RPCEndpointsColumns[4]},
				RefColumns: []*schema.Column{ChainsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "rpcendpoint_endpoint",
				Unique:  true,
				Columns: []*schema.Column{RPCEndpointsColumns[3]},
			},
		},
	}
	// TelegramChatsColumns holds the columns for the "telegram_chats" table.
	TelegramChatsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "chat_id", Type: field.TypeInt64, Unique: true},
		{Name: "name", Type: field.TypeString},
		{Name: "is_group", Type: field.TypeBool},
		{Name: "telegram_chat_user", Type: field.TypeInt, Nullable: true},
	}
	// TelegramChatsTable holds the schema information for the "telegram_chats" table.
	TelegramChatsTable = &schema.Table{
		Name:       "telegram_chats",
		Columns:    TelegramChatsColumns,
		PrimaryKey: []*schema.Column{TelegramChatsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "telegram_chats_users_user",
				Columns:    []*schema.Column{TelegramChatsColumns[6]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "telegramchat_id",
				Unique:  true,
				Columns: []*schema.Column{TelegramChatsColumns[0]},
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "user_id", Type: field.TypeInt64, Default: 0},
		{Name: "name", Type: field.TypeString},
		{Name: "chat_id", Type: field.TypeInt64},
		{Name: "type", Type: field.TypeEnum, Enums: []string{"telegram", "discord"}},
		{Name: "login_token", Type: field.TypeString, Default: ""},
		{Name: "user_name", Type: field.TypeString, Default: "<not set>"},
		{Name: "chat_name", Type: field.TypeString, Default: "<not set>"},
		{Name: "is_group", Type: field.TypeBool, Default: false},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "user_user_id_type",
				Unique:  true,
				Columns: []*schema.Column{UsersColumns[3], UsersColumns[6]},
			},
		},
	}
	// WalletsColumns holds the columns for the "wallets" table.
	WalletsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "address", Type: field.TypeString},
		{Name: "chain_wallets", Type: field.TypeInt, Nullable: true},
	}
	// WalletsTable holds the schema information for the "wallets" table.
	WalletsTable = &schema.Table{
		Name:       "wallets",
		Columns:    WalletsColumns,
		PrimaryKey: []*schema.Column{WalletsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "wallets_chains_wallets",
				Columns:    []*schema.Column{WalletsColumns[4]},
				RefColumns: []*schema.Column{ChainsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "wallet_address",
				Unique:  false,
				Columns: []*schema.Column{WalletsColumns[3]},
			},
			{
				Name:    "wallet_chain_wallets",
				Unique:  false,
				Columns: []*schema.Column{WalletsColumns[4]},
			},
		},
	}
	// DiscordChannelChainsColumns holds the columns for the "discord_channel_chains" table.
	DiscordChannelChainsColumns = []*schema.Column{
		{Name: "discord_channel_id", Type: field.TypeInt},
		{Name: "chain_id", Type: field.TypeInt},
	}
	// DiscordChannelChainsTable holds the schema information for the "discord_channel_chains" table.
	DiscordChannelChainsTable = &schema.Table{
		Name:       "discord_channel_chains",
		Columns:    DiscordChannelChainsColumns,
		PrimaryKey: []*schema.Column{DiscordChannelChainsColumns[0], DiscordChannelChainsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "discord_channel_chains_discord_channel_id",
				Columns:    []*schema.Column{DiscordChannelChainsColumns[0]},
				RefColumns: []*schema.Column{DiscordChannelsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "discord_channel_chains_chain_id",
				Columns:    []*schema.Column{DiscordChannelChainsColumns[1]},
				RefColumns: []*schema.Column{ChainsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// TelegramChatChainsColumns holds the columns for the "telegram_chat_chains" table.
	TelegramChatChainsColumns = []*schema.Column{
		{Name: "telegram_chat_id", Type: field.TypeInt},
		{Name: "chain_id", Type: field.TypeInt},
	}
	// TelegramChatChainsTable holds the schema information for the "telegram_chat_chains" table.
	TelegramChatChainsTable = &schema.Table{
		Name:       "telegram_chat_chains",
		Columns:    TelegramChatChainsColumns,
		PrimaryKey: []*schema.Column{TelegramChatChainsColumns[0], TelegramChatChainsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "telegram_chat_chains_telegram_chat_id",
				Columns:    []*schema.Column{TelegramChatChainsColumns[0]},
				RefColumns: []*schema.Column{TelegramChatsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "telegram_chat_chains_chain_id",
				Columns:    []*schema.Column{TelegramChatChainsColumns[1]},
				RefColumns: []*schema.Column{ChainsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
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
	// UserWalletsColumns holds the columns for the "user_wallets" table.
	UserWalletsColumns = []*schema.Column{
		{Name: "user_id", Type: field.TypeInt},
		{Name: "wallet_id", Type: field.TypeInt},
	}
	// UserWalletsTable holds the schema information for the "user_wallets" table.
	UserWalletsTable = &schema.Table{
		Name:       "user_wallets",
		Columns:    UserWalletsColumns,
		PrimaryKey: []*schema.Column{UserWalletsColumns[0], UserWalletsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "user_wallets_user_id",
				Columns:    []*schema.Column{UserWalletsColumns[0]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "user_wallets_wallet_id",
				Columns:    []*schema.Column{UserWalletsColumns[1]},
				RefColumns: []*schema.Column{WalletsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		ChainsTable,
		DiscordChannelsTable,
		GrantsTable,
		LensChainInfosTable,
		MigrationInfosTable,
		ProposalsTable,
		RPCEndpointsTable,
		TelegramChatsTable,
		UsersTable,
		WalletsTable,
		DiscordChannelChainsTable,
		TelegramChatChainsTable,
		UserChainsTable,
		UserWalletsTable,
	}
)

func init() {
	DiscordChannelsTable.ForeignKeys[0].RefTable = UsersTable
	GrantsTable.ForeignKeys[0].RefTable = WalletsTable
	ProposalsTable.ForeignKeys[0].RefTable = ChainsTable
	RPCEndpointsTable.ForeignKeys[0].RefTable = ChainsTable
	TelegramChatsTable.ForeignKeys[0].RefTable = UsersTable
	WalletsTable.ForeignKeys[0].RefTable = ChainsTable
	DiscordChannelChainsTable.ForeignKeys[0].RefTable = DiscordChannelsTable
	DiscordChannelChainsTable.ForeignKeys[1].RefTable = ChainsTable
	TelegramChatChainsTable.ForeignKeys[0].RefTable = TelegramChatsTable
	TelegramChatChainsTable.ForeignKeys[1].RefTable = ChainsTable
	UserChainsTable.ForeignKeys[0].RefTable = UsersTable
	UserChainsTable.ForeignKeys[1].RefTable = ChainsTable
	UserWalletsTable.ForeignKeys[0].RefTable = UsersTable
	UserWalletsTable.ForeignKeys[1].RefTable = WalletsTable
}
