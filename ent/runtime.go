// Code generated by entc, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/shifty11/cosmos-gov/ent/chain"
	"github.com/shifty11/cosmos-gov/ent/discordchannel"
	"github.com/shifty11/cosmos-gov/ent/lenschaininfo"
	"github.com/shifty11/cosmos-gov/ent/migrationinfo"
	"github.com/shifty11/cosmos-gov/ent/proposal"
	"github.com/shifty11/cosmos-gov/ent/schema"
	"github.com/shifty11/cosmos-gov/ent/telegramchat"
	"github.com/shifty11/cosmos-gov/ent/user"
	"github.com/shifty11/cosmos-gov/ent/wallet"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	chainFields := schema.Chain{}.Fields()
	_ = chainFields
	// chainDescCreatedAt is the schema descriptor for created_at field.
	chainDescCreatedAt := chainFields[0].Descriptor()
	// chain.DefaultCreatedAt holds the default value on creation for the created_at field.
	chain.DefaultCreatedAt = chainDescCreatedAt.Default.(func() time.Time)
	// chainDescUpdatedAt is the schema descriptor for updated_at field.
	chainDescUpdatedAt := chainFields[1].Descriptor()
	// chain.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	chain.DefaultUpdatedAt = chainDescUpdatedAt.Default.(func() time.Time)
	// chain.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	chain.UpdateDefaultUpdatedAt = chainDescUpdatedAt.UpdateDefault.(func() time.Time)
	// chainDescIsEnabled is the schema descriptor for is_enabled field.
	chainDescIsEnabled := chainFields[4].Descriptor()
	// chain.DefaultIsEnabled holds the default value on creation for the is_enabled field.
	chain.DefaultIsEnabled = chainDescIsEnabled.Default.(bool)
	discordchannelFields := schema.DiscordChannel{}.Fields()
	_ = discordchannelFields
	// discordchannelDescCreatedAt is the schema descriptor for created_at field.
	discordchannelDescCreatedAt := discordchannelFields[0].Descriptor()
	// discordchannel.DefaultCreatedAt holds the default value on creation for the created_at field.
	discordchannel.DefaultCreatedAt = discordchannelDescCreatedAt.Default.(func() time.Time)
	// discordchannelDescUpdatedAt is the schema descriptor for updated_at field.
	discordchannelDescUpdatedAt := discordchannelFields[1].Descriptor()
	// discordchannel.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	discordchannel.DefaultUpdatedAt = discordchannelDescUpdatedAt.Default.(func() time.Time)
	// discordchannel.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	discordchannel.UpdateDefaultUpdatedAt = discordchannelDescUpdatedAt.UpdateDefault.(func() time.Time)
	// discordchannelDescRoles is the schema descriptor for roles field.
	discordchannelDescRoles := discordchannelFields[5].Descriptor()
	// discordchannel.DefaultRoles holds the default value on creation for the roles field.
	discordchannel.DefaultRoles = discordchannelDescRoles.Default.(string)
	lenschaininfoFields := schema.LensChainInfo{}.Fields()
	_ = lenschaininfoFields
	// lenschaininfoDescCreatedAt is the schema descriptor for created_at field.
	lenschaininfoDescCreatedAt := lenschaininfoFields[0].Descriptor()
	// lenschaininfo.DefaultCreatedAt holds the default value on creation for the created_at field.
	lenschaininfo.DefaultCreatedAt = lenschaininfoDescCreatedAt.Default.(func() time.Time)
	// lenschaininfoDescUpdatedAt is the schema descriptor for updated_at field.
	lenschaininfoDescUpdatedAt := lenschaininfoFields[1].Descriptor()
	// lenschaininfo.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	lenschaininfo.DefaultUpdatedAt = lenschaininfoDescUpdatedAt.Default.(func() time.Time)
	// lenschaininfo.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	lenschaininfo.UpdateDefaultUpdatedAt = lenschaininfoDescUpdatedAt.UpdateDefault.(func() time.Time)
	migrationinfoFields := schema.MigrationInfo{}.Fields()
	_ = migrationinfoFields
	// migrationinfoDescCreatedAt is the schema descriptor for created_at field.
	migrationinfoDescCreatedAt := migrationinfoFields[0].Descriptor()
	// migrationinfo.DefaultCreatedAt holds the default value on creation for the created_at field.
	migrationinfo.DefaultCreatedAt = migrationinfoDescCreatedAt.Default.(func() time.Time)
	// migrationinfoDescUpdatedAt is the schema descriptor for updated_at field.
	migrationinfoDescUpdatedAt := migrationinfoFields[1].Descriptor()
	// migrationinfo.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	migrationinfo.DefaultUpdatedAt = migrationinfoDescUpdatedAt.Default.(func() time.Time)
	// migrationinfo.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	migrationinfo.UpdateDefaultUpdatedAt = migrationinfoDescUpdatedAt.UpdateDefault.(func() time.Time)
	// migrationinfoDescIsMigrated is the schema descriptor for is_migrated field.
	migrationinfoDescIsMigrated := migrationinfoFields[2].Descriptor()
	// migrationinfo.DefaultIsMigrated holds the default value on creation for the is_migrated field.
	migrationinfo.DefaultIsMigrated = migrationinfoDescIsMigrated.Default.(bool)
	proposalFields := schema.Proposal{}.Fields()
	_ = proposalFields
	// proposalDescCreatedAt is the schema descriptor for created_at field.
	proposalDescCreatedAt := proposalFields[0].Descriptor()
	// proposal.DefaultCreatedAt holds the default value on creation for the created_at field.
	proposal.DefaultCreatedAt = proposalDescCreatedAt.Default.(func() time.Time)
	// proposalDescUpdatedAt is the schema descriptor for updated_at field.
	proposalDescUpdatedAt := proposalFields[1].Descriptor()
	// proposal.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	proposal.DefaultUpdatedAt = proposalDescUpdatedAt.Default.(func() time.Time)
	// proposal.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	proposal.UpdateDefaultUpdatedAt = proposalDescUpdatedAt.UpdateDefault.(func() time.Time)
	telegramchatFields := schema.TelegramChat{}.Fields()
	_ = telegramchatFields
	// telegramchatDescCreatedAt is the schema descriptor for created_at field.
	telegramchatDescCreatedAt := telegramchatFields[0].Descriptor()
	// telegramchat.DefaultCreatedAt holds the default value on creation for the created_at field.
	telegramchat.DefaultCreatedAt = telegramchatDescCreatedAt.Default.(func() time.Time)
	// telegramchatDescUpdatedAt is the schema descriptor for updated_at field.
	telegramchatDescUpdatedAt := telegramchatFields[1].Descriptor()
	// telegramchat.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	telegramchat.DefaultUpdatedAt = telegramchatDescUpdatedAt.Default.(func() time.Time)
	// telegramchat.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	telegramchat.UpdateDefaultUpdatedAt = telegramchatDescUpdatedAt.UpdateDefault.(func() time.Time)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescCreatedAt is the schema descriptor for created_at field.
	userDescCreatedAt := userFields[0].Descriptor()
	// user.DefaultCreatedAt holds the default value on creation for the created_at field.
	user.DefaultCreatedAt = userDescCreatedAt.Default.(func() time.Time)
	// userDescUpdatedAt is the schema descriptor for updated_at field.
	userDescUpdatedAt := userFields[1].Descriptor()
	// user.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	user.DefaultUpdatedAt = userDescUpdatedAt.Default.(func() time.Time)
	// user.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	user.UpdateDefaultUpdatedAt = userDescUpdatedAt.UpdateDefault.(func() time.Time)
	// userDescName is the schema descriptor for name field.
	userDescName := userFields[3].Descriptor()
	// user.DefaultName holds the default value on creation for the name field.
	user.DefaultName = userDescName.Default.(string)
	// userDescLoginToken is the schema descriptor for login_token field.
	userDescLoginToken := userFields[5].Descriptor()
	// user.DefaultLoginToken holds the default value on creation for the login_token field.
	user.DefaultLoginToken = userDescLoginToken.Default.(string)
	walletFields := schema.Wallet{}.Fields()
	_ = walletFields
	// walletDescCreatedAt is the schema descriptor for created_at field.
	walletDescCreatedAt := walletFields[0].Descriptor()
	// wallet.DefaultCreatedAt holds the default value on creation for the created_at field.
	wallet.DefaultCreatedAt = walletDescCreatedAt.Default.(func() time.Time)
	// walletDescUpdatedAt is the schema descriptor for updated_at field.
	walletDescUpdatedAt := walletFields[1].Descriptor()
	// wallet.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	wallet.DefaultUpdatedAt = walletDescUpdatedAt.Default.(func() time.Time)
	// wallet.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	wallet.UpdateDefaultUpdatedAt = walletDescUpdatedAt.UpdateDefault.(func() time.Time)
}
