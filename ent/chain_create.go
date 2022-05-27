// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/shifty11/cosmos-gov/ent/chain"
	"github.com/shifty11/cosmos-gov/ent/discordchannel"
	"github.com/shifty11/cosmos-gov/ent/draftproposal"
	"github.com/shifty11/cosmos-gov/ent/proposal"
	"github.com/shifty11/cosmos-gov/ent/rpcendpoint"
	"github.com/shifty11/cosmos-gov/ent/telegramchat"
	"github.com/shifty11/cosmos-gov/ent/wallet"
)

// ChainCreate is the builder for creating a Chain entity.
type ChainCreate struct {
	config
	mutation *ChainMutation
	hooks    []Hook
}

// SetCreateTime sets the "create_time" field.
func (cc *ChainCreate) SetCreateTime(t time.Time) *ChainCreate {
	cc.mutation.SetCreateTime(t)
	return cc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (cc *ChainCreate) SetNillableCreateTime(t *time.Time) *ChainCreate {
	if t != nil {
		cc.SetCreateTime(*t)
	}
	return cc
}

// SetUpdateTime sets the "update_time" field.
func (cc *ChainCreate) SetUpdateTime(t time.Time) *ChainCreate {
	cc.mutation.SetUpdateTime(t)
	return cc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (cc *ChainCreate) SetNillableUpdateTime(t *time.Time) *ChainCreate {
	if t != nil {
		cc.SetUpdateTime(*t)
	}
	return cc
}

// SetChainID sets the "chain_id" field.
func (cc *ChainCreate) SetChainID(s string) *ChainCreate {
	cc.mutation.SetChainID(s)
	return cc
}

// SetAccountPrefix sets the "account_prefix" field.
func (cc *ChainCreate) SetAccountPrefix(s string) *ChainCreate {
	cc.mutation.SetAccountPrefix(s)
	return cc
}

// SetName sets the "name" field.
func (cc *ChainCreate) SetName(s string) *ChainCreate {
	cc.mutation.SetName(s)
	return cc
}

// SetDisplayName sets the "display_name" field.
func (cc *ChainCreate) SetDisplayName(s string) *ChainCreate {
	cc.mutation.SetDisplayName(s)
	return cc
}

// SetIsEnabled sets the "is_enabled" field.
func (cc *ChainCreate) SetIsEnabled(b bool) *ChainCreate {
	cc.mutation.SetIsEnabled(b)
	return cc
}

// SetNillableIsEnabled sets the "is_enabled" field if the given value is not nil.
func (cc *ChainCreate) SetNillableIsEnabled(b *bool) *ChainCreate {
	if b != nil {
		cc.SetIsEnabled(*b)
	}
	return cc
}

// SetIsVotingEnabled sets the "is_voting_enabled" field.
func (cc *ChainCreate) SetIsVotingEnabled(b bool) *ChainCreate {
	cc.mutation.SetIsVotingEnabled(b)
	return cc
}

// SetNillableIsVotingEnabled sets the "is_voting_enabled" field if the given value is not nil.
func (cc *ChainCreate) SetNillableIsVotingEnabled(b *bool) *ChainCreate {
	if b != nil {
		cc.SetIsVotingEnabled(*b)
	}
	return cc
}

// SetIsFeegrantUsed sets the "is_feegrant_used" field.
func (cc *ChainCreate) SetIsFeegrantUsed(b bool) *ChainCreate {
	cc.mutation.SetIsFeegrantUsed(b)
	return cc
}

// SetNillableIsFeegrantUsed sets the "is_feegrant_used" field if the given value is not nil.
func (cc *ChainCreate) SetNillableIsFeegrantUsed(b *bool) *ChainCreate {
	if b != nil {
		cc.SetIsFeegrantUsed(*b)
	}
	return cc
}

// AddProposalIDs adds the "proposals" edge to the Proposal entity by IDs.
func (cc *ChainCreate) AddProposalIDs(ids ...int) *ChainCreate {
	cc.mutation.AddProposalIDs(ids...)
	return cc
}

// AddProposals adds the "proposals" edges to the Proposal entity.
func (cc *ChainCreate) AddProposals(p ...*Proposal) *ChainCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return cc.AddProposalIDs(ids...)
}

// AddDraftProposalIDs adds the "draft_proposals" edge to the DraftProposal entity by IDs.
func (cc *ChainCreate) AddDraftProposalIDs(ids ...int) *ChainCreate {
	cc.mutation.AddDraftProposalIDs(ids...)
	return cc
}

// AddDraftProposals adds the "draft_proposals" edges to the DraftProposal entity.
func (cc *ChainCreate) AddDraftProposals(d ...*DraftProposal) *ChainCreate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return cc.AddDraftProposalIDs(ids...)
}

// AddTelegramChatIDs adds the "telegram_chats" edge to the TelegramChat entity by IDs.
func (cc *ChainCreate) AddTelegramChatIDs(ids ...int) *ChainCreate {
	cc.mutation.AddTelegramChatIDs(ids...)
	return cc
}

// AddTelegramChats adds the "telegram_chats" edges to the TelegramChat entity.
func (cc *ChainCreate) AddTelegramChats(t ...*TelegramChat) *ChainCreate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return cc.AddTelegramChatIDs(ids...)
}

// AddDiscordChannelIDs adds the "discord_channels" edge to the DiscordChannel entity by IDs.
func (cc *ChainCreate) AddDiscordChannelIDs(ids ...int) *ChainCreate {
	cc.mutation.AddDiscordChannelIDs(ids...)
	return cc
}

// AddDiscordChannels adds the "discord_channels" edges to the DiscordChannel entity.
func (cc *ChainCreate) AddDiscordChannels(d ...*DiscordChannel) *ChainCreate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return cc.AddDiscordChannelIDs(ids...)
}

// AddRPCEndpointIDs adds the "rpc_endpoints" edge to the RpcEndpoint entity by IDs.
func (cc *ChainCreate) AddRPCEndpointIDs(ids ...int) *ChainCreate {
	cc.mutation.AddRPCEndpointIDs(ids...)
	return cc
}

// AddRPCEndpoints adds the "rpc_endpoints" edges to the RpcEndpoint entity.
func (cc *ChainCreate) AddRPCEndpoints(r ...*RpcEndpoint) *ChainCreate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return cc.AddRPCEndpointIDs(ids...)
}

// AddWalletIDs adds the "wallets" edge to the Wallet entity by IDs.
func (cc *ChainCreate) AddWalletIDs(ids ...int) *ChainCreate {
	cc.mutation.AddWalletIDs(ids...)
	return cc
}

// AddWallets adds the "wallets" edges to the Wallet entity.
func (cc *ChainCreate) AddWallets(w ...*Wallet) *ChainCreate {
	ids := make([]int, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return cc.AddWalletIDs(ids...)
}

// Mutation returns the ChainMutation object of the builder.
func (cc *ChainCreate) Mutation() *ChainMutation {
	return cc.mutation
}

// Save creates the Chain in the database.
func (cc *ChainCreate) Save(ctx context.Context) (*Chain, error) {
	var (
		err  error
		node *Chain
	)
	cc.defaults()
	if len(cc.hooks) == 0 {
		if err = cc.check(); err != nil {
			return nil, err
		}
		node, err = cc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ChainMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cc.check(); err != nil {
				return nil, err
			}
			cc.mutation = mutation
			if node, err = cc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(cc.hooks) - 1; i >= 0; i-- {
			if cc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (cc *ChainCreate) SaveX(ctx context.Context) *Chain {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cc *ChainCreate) Exec(ctx context.Context) error {
	_, err := cc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cc *ChainCreate) ExecX(ctx context.Context) {
	if err := cc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cc *ChainCreate) defaults() {
	if _, ok := cc.mutation.CreateTime(); !ok {
		v := chain.DefaultCreateTime()
		cc.mutation.SetCreateTime(v)
	}
	if _, ok := cc.mutation.UpdateTime(); !ok {
		v := chain.DefaultUpdateTime()
		cc.mutation.SetUpdateTime(v)
	}
	if _, ok := cc.mutation.IsEnabled(); !ok {
		v := chain.DefaultIsEnabled
		cc.mutation.SetIsEnabled(v)
	}
	if _, ok := cc.mutation.IsVotingEnabled(); !ok {
		v := chain.DefaultIsVotingEnabled
		cc.mutation.SetIsVotingEnabled(v)
	}
	if _, ok := cc.mutation.IsFeegrantUsed(); !ok {
		v := chain.DefaultIsFeegrantUsed
		cc.mutation.SetIsFeegrantUsed(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cc *ChainCreate) check() error {
	if _, ok := cc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New(`ent: missing required field "Chain.create_time"`)}
	}
	if _, ok := cc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New(`ent: missing required field "Chain.update_time"`)}
	}
	if _, ok := cc.mutation.ChainID(); !ok {
		return &ValidationError{Name: "chain_id", err: errors.New(`ent: missing required field "Chain.chain_id"`)}
	}
	if _, ok := cc.mutation.AccountPrefix(); !ok {
		return &ValidationError{Name: "account_prefix", err: errors.New(`ent: missing required field "Chain.account_prefix"`)}
	}
	if _, ok := cc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Chain.name"`)}
	}
	if _, ok := cc.mutation.DisplayName(); !ok {
		return &ValidationError{Name: "display_name", err: errors.New(`ent: missing required field "Chain.display_name"`)}
	}
	if _, ok := cc.mutation.IsEnabled(); !ok {
		return &ValidationError{Name: "is_enabled", err: errors.New(`ent: missing required field "Chain.is_enabled"`)}
	}
	if _, ok := cc.mutation.IsVotingEnabled(); !ok {
		return &ValidationError{Name: "is_voting_enabled", err: errors.New(`ent: missing required field "Chain.is_voting_enabled"`)}
	}
	if _, ok := cc.mutation.IsFeegrantUsed(); !ok {
		return &ValidationError{Name: "is_feegrant_used", err: errors.New(`ent: missing required field "Chain.is_feegrant_used"`)}
	}
	return nil
}

func (cc *ChainCreate) sqlSave(ctx context.Context) (*Chain, error) {
	_node, _spec := cc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (cc *ChainCreate) createSpec() (*Chain, *sqlgraph.CreateSpec) {
	var (
		_node = &Chain{config: cc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: chain.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: chain.FieldID,
			},
		}
	)
	if value, ok := cc.mutation.CreateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: chain.FieldCreateTime,
		})
		_node.CreateTime = value
	}
	if value, ok := cc.mutation.UpdateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: chain.FieldUpdateTime,
		})
		_node.UpdateTime = value
	}
	if value, ok := cc.mutation.ChainID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: chain.FieldChainID,
		})
		_node.ChainID = value
	}
	if value, ok := cc.mutation.AccountPrefix(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: chain.FieldAccountPrefix,
		})
		_node.AccountPrefix = value
	}
	if value, ok := cc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: chain.FieldName,
		})
		_node.Name = value
	}
	if value, ok := cc.mutation.DisplayName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: chain.FieldDisplayName,
		})
		_node.DisplayName = value
	}
	if value, ok := cc.mutation.IsEnabled(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: chain.FieldIsEnabled,
		})
		_node.IsEnabled = value
	}
	if value, ok := cc.mutation.IsVotingEnabled(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: chain.FieldIsVotingEnabled,
		})
		_node.IsVotingEnabled = value
	}
	if value, ok := cc.mutation.IsFeegrantUsed(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: chain.FieldIsFeegrantUsed,
		})
		_node.IsFeegrantUsed = value
	}
	if nodes := cc.mutation.ProposalsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   chain.ProposalsTable,
			Columns: []string{chain.ProposalsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: proposal.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cc.mutation.DraftProposalsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   chain.DraftProposalsTable,
			Columns: []string{chain.DraftProposalsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: draftproposal.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cc.mutation.TelegramChatsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   chain.TelegramChatsTable,
			Columns: chain.TelegramChatsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: telegramchat.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cc.mutation.DiscordChannelsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   chain.DiscordChannelsTable,
			Columns: chain.DiscordChannelsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: discordchannel.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cc.mutation.RPCEndpointsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   chain.RPCEndpointsTable,
			Columns: []string{chain.RPCEndpointsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: rpcendpoint.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cc.mutation.WalletsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   chain.WalletsTable,
			Columns: []string{chain.WalletsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: wallet.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ChainCreateBulk is the builder for creating many Chain entities in bulk.
type ChainCreateBulk struct {
	config
	builders []*ChainCreate
}

// Save creates the Chain entities in the database.
func (ccb *ChainCreateBulk) Save(ctx context.Context) ([]*Chain, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ccb.builders))
	nodes := make([]*Chain, len(ccb.builders))
	mutators := make([]Mutator, len(ccb.builders))
	for i := range ccb.builders {
		func(i int, root context.Context) {
			builder := ccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ChainMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ccb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ccb *ChainCreateBulk) SaveX(ctx context.Context) []*Chain {
	v, err := ccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ccb *ChainCreateBulk) Exec(ctx context.Context) error {
	_, err := ccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccb *ChainCreateBulk) ExecX(ctx context.Context) {
	if err := ccb.Exec(ctx); err != nil {
		panic(err)
	}
}
