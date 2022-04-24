// Code generated by entc, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/shifty11/cosmos-gov/ent/chain"
	"github.com/shifty11/cosmos-gov/ent/lenschaininfo"
	"github.com/shifty11/cosmos-gov/ent/proposal"
	"github.com/shifty11/cosmos-gov/ent/rpcendpoint"
	"github.com/shifty11/cosmos-gov/ent/schema"
	"github.com/shifty11/cosmos-gov/ent/user"
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
	rpcendpointFields := schema.RpcEndpoint{}.Fields()
	_ = rpcendpointFields
	// rpcendpointDescCreatedAt is the schema descriptor for created_at field.
	rpcendpointDescCreatedAt := rpcendpointFields[0].Descriptor()
	// rpcendpoint.DefaultCreatedAt holds the default value on creation for the created_at field.
	rpcendpoint.DefaultCreatedAt = rpcendpointDescCreatedAt.Default.(func() time.Time)
	// rpcendpointDescUpdatedAt is the schema descriptor for updated_at field.
	rpcendpointDescUpdatedAt := rpcendpointFields[1].Descriptor()
	// rpcendpoint.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	rpcendpoint.DefaultUpdatedAt = rpcendpointDescUpdatedAt.Default.(func() time.Time)
	// rpcendpoint.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	rpcendpoint.UpdateDefaultUpdatedAt = rpcendpointDescUpdatedAt.UpdateDefault.(func() time.Time)
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
}
