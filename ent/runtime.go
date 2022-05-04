// Code generated by entc, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/shifty11/cosmos-gov/ent/chain"
	"github.com/shifty11/cosmos-gov/ent/discordchannel"
	"github.com/shifty11/cosmos-gov/ent/grant"
	"github.com/shifty11/cosmos-gov/ent/lenschaininfo"
	"github.com/shifty11/cosmos-gov/ent/proposal"
	"github.com/shifty11/cosmos-gov/ent/rpcendpoint"
	"github.com/shifty11/cosmos-gov/ent/schema"
	"github.com/shifty11/cosmos-gov/ent/telegramchat"
	"github.com/shifty11/cosmos-gov/ent/user"
	"github.com/shifty11/cosmos-gov/ent/wallet"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	chainMixin := schema.Chain{}.Mixin()
	chainMixinFields0 := chainMixin[0].Fields()
	_ = chainMixinFields0
	chainFields := schema.Chain{}.Fields()
	_ = chainFields
	// chainDescCreateTime is the schema descriptor for create_time field.
	chainDescCreateTime := chainMixinFields0[0].Descriptor()
	// chain.DefaultCreateTime holds the default value on creation for the create_time field.
	chain.DefaultCreateTime = chainDescCreateTime.Default.(func() time.Time)
	// chainDescUpdateTime is the schema descriptor for update_time field.
	chainDescUpdateTime := chainMixinFields0[1].Descriptor()
	// chain.DefaultUpdateTime holds the default value on creation for the update_time field.
	chain.DefaultUpdateTime = chainDescUpdateTime.Default.(func() time.Time)
	// chain.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	chain.UpdateDefaultUpdateTime = chainDescUpdateTime.UpdateDefault.(func() time.Time)
	// chainDescIsEnabled is the schema descriptor for is_enabled field.
	chainDescIsEnabled := chainFields[4].Descriptor()
	// chain.DefaultIsEnabled holds the default value on creation for the is_enabled field.
	chain.DefaultIsEnabled = chainDescIsEnabled.Default.(bool)
	discordchannelMixin := schema.DiscordChannel{}.Mixin()
	discordchannelMixinFields0 := discordchannelMixin[0].Fields()
	_ = discordchannelMixinFields0
	discordchannelFields := schema.DiscordChannel{}.Fields()
	_ = discordchannelFields
	// discordchannelDescCreateTime is the schema descriptor for create_time field.
	discordchannelDescCreateTime := discordchannelMixinFields0[0].Descriptor()
	// discordchannel.DefaultCreateTime holds the default value on creation for the create_time field.
	discordchannel.DefaultCreateTime = discordchannelDescCreateTime.Default.(func() time.Time)
	// discordchannelDescUpdateTime is the schema descriptor for update_time field.
	discordchannelDescUpdateTime := discordchannelMixinFields0[1].Descriptor()
	// discordchannel.DefaultUpdateTime holds the default value on creation for the update_time field.
	discordchannel.DefaultUpdateTime = discordchannelDescUpdateTime.Default.(func() time.Time)
	// discordchannel.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	discordchannel.UpdateDefaultUpdateTime = discordchannelDescUpdateTime.UpdateDefault.(func() time.Time)
	// discordchannelDescRoles is the schema descriptor for roles field.
	discordchannelDescRoles := discordchannelFields[3].Descriptor()
	// discordchannel.DefaultRoles holds the default value on creation for the roles field.
	discordchannel.DefaultRoles = discordchannelDescRoles.Default.(string)
	grantMixin := schema.Grant{}.Mixin()
	grantMixinFields0 := grantMixin[0].Fields()
	_ = grantMixinFields0
	grantFields := schema.Grant{}.Fields()
	_ = grantFields
	// grantDescCreateTime is the schema descriptor for create_time field.
	grantDescCreateTime := grantMixinFields0[0].Descriptor()
	// grant.DefaultCreateTime holds the default value on creation for the create_time field.
	grant.DefaultCreateTime = grantDescCreateTime.Default.(func() time.Time)
	// grantDescUpdateTime is the schema descriptor for update_time field.
	grantDescUpdateTime := grantMixinFields0[1].Descriptor()
	// grant.DefaultUpdateTime holds the default value on creation for the update_time field.
	grant.DefaultUpdateTime = grantDescUpdateTime.Default.(func() time.Time)
	// grant.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	grant.UpdateDefaultUpdateTime = grantDescUpdateTime.UpdateDefault.(func() time.Time)
	lenschaininfoMixin := schema.LensChainInfo{}.Mixin()
	lenschaininfoMixinFields0 := lenschaininfoMixin[0].Fields()
	_ = lenschaininfoMixinFields0
	lenschaininfoFields := schema.LensChainInfo{}.Fields()
	_ = lenschaininfoFields
	// lenschaininfoDescCreateTime is the schema descriptor for create_time field.
	lenschaininfoDescCreateTime := lenschaininfoMixinFields0[0].Descriptor()
	// lenschaininfo.DefaultCreateTime holds the default value on creation for the create_time field.
	lenschaininfo.DefaultCreateTime = lenschaininfoDescCreateTime.Default.(func() time.Time)
	// lenschaininfoDescUpdateTime is the schema descriptor for update_time field.
	lenschaininfoDescUpdateTime := lenschaininfoMixinFields0[1].Descriptor()
	// lenschaininfo.DefaultUpdateTime holds the default value on creation for the update_time field.
	lenschaininfo.DefaultUpdateTime = lenschaininfoDescUpdateTime.Default.(func() time.Time)
	// lenschaininfo.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	lenschaininfo.UpdateDefaultUpdateTime = lenschaininfoDescUpdateTime.UpdateDefault.(func() time.Time)
	proposalMixin := schema.Proposal{}.Mixin()
	proposalMixinFields0 := proposalMixin[0].Fields()
	_ = proposalMixinFields0
	proposalFields := schema.Proposal{}.Fields()
	_ = proposalFields
	// proposalDescCreateTime is the schema descriptor for create_time field.
	proposalDescCreateTime := proposalMixinFields0[0].Descriptor()
	// proposal.DefaultCreateTime holds the default value on creation for the create_time field.
	proposal.DefaultCreateTime = proposalDescCreateTime.Default.(func() time.Time)
	// proposalDescUpdateTime is the schema descriptor for update_time field.
	proposalDescUpdateTime := proposalMixinFields0[1].Descriptor()
	// proposal.DefaultUpdateTime holds the default value on creation for the update_time field.
	proposal.DefaultUpdateTime = proposalDescUpdateTime.Default.(func() time.Time)
	// proposal.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	proposal.UpdateDefaultUpdateTime = proposalDescUpdateTime.UpdateDefault.(func() time.Time)
	rpcendpointMixin := schema.RpcEndpoint{}.Mixin()
	rpcendpointMixinFields0 := rpcendpointMixin[0].Fields()
	_ = rpcendpointMixinFields0
	rpcendpointFields := schema.RpcEndpoint{}.Fields()
	_ = rpcendpointFields
	// rpcendpointDescCreateTime is the schema descriptor for create_time field.
	rpcendpointDescCreateTime := rpcendpointMixinFields0[0].Descriptor()
	// rpcendpoint.DefaultCreateTime holds the default value on creation for the create_time field.
	rpcendpoint.DefaultCreateTime = rpcendpointDescCreateTime.Default.(func() time.Time)
	// rpcendpointDescUpdateTime is the schema descriptor for update_time field.
	rpcendpointDescUpdateTime := rpcendpointMixinFields0[1].Descriptor()
	// rpcendpoint.DefaultUpdateTime holds the default value on creation for the update_time field.
	rpcendpoint.DefaultUpdateTime = rpcendpointDescUpdateTime.Default.(func() time.Time)
	// rpcendpoint.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	rpcendpoint.UpdateDefaultUpdateTime = rpcendpointDescUpdateTime.UpdateDefault.(func() time.Time)
	telegramchatMixin := schema.TelegramChat{}.Mixin()
	telegramchatMixinFields0 := telegramchatMixin[0].Fields()
	_ = telegramchatMixinFields0
	telegramchatFields := schema.TelegramChat{}.Fields()
	_ = telegramchatFields
	// telegramchatDescCreateTime is the schema descriptor for create_time field.
	telegramchatDescCreateTime := telegramchatMixinFields0[0].Descriptor()
	// telegramchat.DefaultCreateTime holds the default value on creation for the create_time field.
	telegramchat.DefaultCreateTime = telegramchatDescCreateTime.Default.(func() time.Time)
	// telegramchatDescUpdateTime is the schema descriptor for update_time field.
	telegramchatDescUpdateTime := telegramchatMixinFields0[1].Descriptor()
	// telegramchat.DefaultUpdateTime holds the default value on creation for the update_time field.
	telegramchat.DefaultUpdateTime = telegramchatDescUpdateTime.Default.(func() time.Time)
	// telegramchat.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	telegramchat.UpdateDefaultUpdateTime = telegramchatDescUpdateTime.UpdateDefault.(func() time.Time)
	userMixin := schema.User{}.Mixin()
	userMixinFields0 := userMixin[0].Fields()
	_ = userMixinFields0
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescCreateTime is the schema descriptor for create_time field.
	userDescCreateTime := userMixinFields0[0].Descriptor()
	// user.DefaultCreateTime holds the default value on creation for the create_time field.
	user.DefaultCreateTime = userDescCreateTime.Default.(func() time.Time)
	// userDescUpdateTime is the schema descriptor for update_time field.
	userDescUpdateTime := userMixinFields0[1].Descriptor()
	// user.DefaultUpdateTime holds the default value on creation for the update_time field.
	user.DefaultUpdateTime = userDescUpdateTime.Default.(func() time.Time)
	// user.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	user.UpdateDefaultUpdateTime = userDescUpdateTime.UpdateDefault.(func() time.Time)
	walletMixin := schema.Wallet{}.Mixin()
	walletMixinFields0 := walletMixin[0].Fields()
	_ = walletMixinFields0
	walletFields := schema.Wallet{}.Fields()
	_ = walletFields
	// walletDescCreateTime is the schema descriptor for create_time field.
	walletDescCreateTime := walletMixinFields0[0].Descriptor()
	// wallet.DefaultCreateTime holds the default value on creation for the create_time field.
	wallet.DefaultCreateTime = walletDescCreateTime.Default.(func() time.Time)
	// walletDescUpdateTime is the schema descriptor for update_time field.
	walletDescUpdateTime := walletMixinFields0[1].Descriptor()
	// wallet.DefaultUpdateTime holds the default value on creation for the update_time field.
	wallet.DefaultUpdateTime = walletDescUpdateTime.Default.(func() time.Time)
	// wallet.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	wallet.UpdateDefaultUpdateTime = walletDescUpdateTime.UpdateDefault.(func() time.Time)
}
