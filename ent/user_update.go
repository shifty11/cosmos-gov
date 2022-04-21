// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/shifty11/cosmos-gov/ent/discordchannel"
	"github.com/shifty11/cosmos-gov/ent/predicate"
	"github.com/shifty11/cosmos-gov/ent/telegramchat"
	"github.com/shifty11/cosmos-gov/ent/user"
	"github.com/shifty11/cosmos-gov/ent/wallet"
)

// UserUpdate is the builder for updating User entities.
type UserUpdate struct {
	config
	hooks    []Hook
	mutation *UserMutation
}

// Where appends a list predicates to the UserUpdate builder.
func (uu *UserUpdate) Where(ps ...predicate.User) *UserUpdate {
	uu.mutation.Where(ps...)
	return uu
}

// SetUpdatedAt sets the "updated_at" field.
func (uu *UserUpdate) SetUpdatedAt(t time.Time) *UserUpdate {
	uu.mutation.SetUpdatedAt(t)
	return uu
}

// SetName sets the "name" field.
func (uu *UserUpdate) SetName(s string) *UserUpdate {
	uu.mutation.SetName(s)
	return uu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (uu *UserUpdate) SetNillableName(s *string) *UserUpdate {
	if s != nil {
		uu.SetName(*s)
	}
	return uu
}

// SetLoginToken sets the "login_token" field.
func (uu *UserUpdate) SetLoginToken(s string) *UserUpdate {
	uu.mutation.SetLoginToken(s)
	return uu
}

// SetNillableLoginToken sets the "login_token" field if the given value is not nil.
func (uu *UserUpdate) SetNillableLoginToken(s *string) *UserUpdate {
	if s != nil {
		uu.SetLoginToken(*s)
	}
	return uu
}

// AddTelegramChatIDs adds the "telegram_chats" edge to the TelegramChat entity by IDs.
func (uu *UserUpdate) AddTelegramChatIDs(ids ...int64) *UserUpdate {
	uu.mutation.AddTelegramChatIDs(ids...)
	return uu
}

// AddTelegramChats adds the "telegram_chats" edges to the TelegramChat entity.
func (uu *UserUpdate) AddTelegramChats(t ...*TelegramChat) *UserUpdate {
	ids := make([]int64, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return uu.AddTelegramChatIDs(ids...)
}

// AddDiscordChannelIDs adds the "discord_channels" edge to the DiscordChannel entity by IDs.
func (uu *UserUpdate) AddDiscordChannelIDs(ids ...int64) *UserUpdate {
	uu.mutation.AddDiscordChannelIDs(ids...)
	return uu
}

// AddDiscordChannels adds the "discord_channels" edges to the DiscordChannel entity.
func (uu *UserUpdate) AddDiscordChannels(d ...*DiscordChannel) *UserUpdate {
	ids := make([]int64, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return uu.AddDiscordChannelIDs(ids...)
}

// AddWalletIDs adds the "wallets" edge to the Wallet entity by IDs.
func (uu *UserUpdate) AddWalletIDs(ids ...int) *UserUpdate {
	uu.mutation.AddWalletIDs(ids...)
	return uu
}

// AddWallets adds the "wallets" edges to the Wallet entity.
func (uu *UserUpdate) AddWallets(w ...*Wallet) *UserUpdate {
	ids := make([]int, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return uu.AddWalletIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uu *UserUpdate) Mutation() *UserMutation {
	return uu.mutation
}

// ClearTelegramChats clears all "telegram_chats" edges to the TelegramChat entity.
func (uu *UserUpdate) ClearTelegramChats() *UserUpdate {
	uu.mutation.ClearTelegramChats()
	return uu
}

// RemoveTelegramChatIDs removes the "telegram_chats" edge to TelegramChat entities by IDs.
func (uu *UserUpdate) RemoveTelegramChatIDs(ids ...int64) *UserUpdate {
	uu.mutation.RemoveTelegramChatIDs(ids...)
	return uu
}

// RemoveTelegramChats removes "telegram_chats" edges to TelegramChat entities.
func (uu *UserUpdate) RemoveTelegramChats(t ...*TelegramChat) *UserUpdate {
	ids := make([]int64, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return uu.RemoveTelegramChatIDs(ids...)
}

// ClearDiscordChannels clears all "discord_channels" edges to the DiscordChannel entity.
func (uu *UserUpdate) ClearDiscordChannels() *UserUpdate {
	uu.mutation.ClearDiscordChannels()
	return uu
}

// RemoveDiscordChannelIDs removes the "discord_channels" edge to DiscordChannel entities by IDs.
func (uu *UserUpdate) RemoveDiscordChannelIDs(ids ...int64) *UserUpdate {
	uu.mutation.RemoveDiscordChannelIDs(ids...)
	return uu
}

// RemoveDiscordChannels removes "discord_channels" edges to DiscordChannel entities.
func (uu *UserUpdate) RemoveDiscordChannels(d ...*DiscordChannel) *UserUpdate {
	ids := make([]int64, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return uu.RemoveDiscordChannelIDs(ids...)
}

// ClearWallets clears all "wallets" edges to the Wallet entity.
func (uu *UserUpdate) ClearWallets() *UserUpdate {
	uu.mutation.ClearWallets()
	return uu
}

// RemoveWalletIDs removes the "wallets" edge to Wallet entities by IDs.
func (uu *UserUpdate) RemoveWalletIDs(ids ...int) *UserUpdate {
	uu.mutation.RemoveWalletIDs(ids...)
	return uu
}

// RemoveWallets removes "wallets" edges to Wallet entities.
func (uu *UserUpdate) RemoveWallets(w ...*Wallet) *UserUpdate {
	ids := make([]int, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return uu.RemoveWalletIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (uu *UserUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	uu.defaults()
	if len(uu.hooks) == 0 {
		affected, err = uu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			uu.mutation = mutation
			affected, err = uu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(uu.hooks) - 1; i >= 0; i-- {
			if uu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = uu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, uu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (uu *UserUpdate) SaveX(ctx context.Context) int {
	affected, err := uu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uu *UserUpdate) Exec(ctx context.Context) error {
	_, err := uu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uu *UserUpdate) ExecX(ctx context.Context) {
	if err := uu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uu *UserUpdate) defaults() {
	if _, ok := uu.mutation.UpdatedAt(); !ok {
		v := user.UpdateDefaultUpdatedAt()
		uu.mutation.SetUpdatedAt(v)
	}
}

func (uu *UserUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   user.Table,
			Columns: user.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: user.FieldID,
			},
		},
	}
	if ps := uu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: user.FieldUpdatedAt,
		})
	}
	if value, ok := uu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldName,
		})
	}
	if value, ok := uu.mutation.LoginToken(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldLoginToken,
		})
	}
	if uu.mutation.TelegramChatsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   user.TelegramChatsTable,
			Columns: []string{user.TelegramChatsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: telegramchat.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedTelegramChatsIDs(); len(nodes) > 0 && !uu.mutation.TelegramChatsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   user.TelegramChatsTable,
			Columns: []string{user.TelegramChatsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: telegramchat.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.TelegramChatsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   user.TelegramChatsTable,
			Columns: []string{user.TelegramChatsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: telegramchat.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uu.mutation.DiscordChannelsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   user.DiscordChannelsTable,
			Columns: []string{user.DiscordChannelsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: discordchannel.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedDiscordChannelsIDs(); len(nodes) > 0 && !uu.mutation.DiscordChannelsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   user.DiscordChannelsTable,
			Columns: []string{user.DiscordChannelsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: discordchannel.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.DiscordChannelsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   user.DiscordChannelsTable,
			Columns: []string{user.DiscordChannelsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: discordchannel.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uu.mutation.WalletsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.WalletsTable,
			Columns: user.WalletsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: wallet.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedWalletsIDs(); len(nodes) > 0 && !uu.mutation.WalletsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.WalletsTable,
			Columns: user.WalletsPrimaryKey,
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.WalletsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.WalletsTable,
			Columns: user.WalletsPrimaryKey,
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, uu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// UserUpdateOne is the builder for updating a single User entity.
type UserUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UserMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (uuo *UserUpdateOne) SetUpdatedAt(t time.Time) *UserUpdateOne {
	uuo.mutation.SetUpdatedAt(t)
	return uuo
}

// SetName sets the "name" field.
func (uuo *UserUpdateOne) SetName(s string) *UserUpdateOne {
	uuo.mutation.SetName(s)
	return uuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableName(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetName(*s)
	}
	return uuo
}

// SetLoginToken sets the "login_token" field.
func (uuo *UserUpdateOne) SetLoginToken(s string) *UserUpdateOne {
	uuo.mutation.SetLoginToken(s)
	return uuo
}

// SetNillableLoginToken sets the "login_token" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableLoginToken(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetLoginToken(*s)
	}
	return uuo
}

// AddTelegramChatIDs adds the "telegram_chats" edge to the TelegramChat entity by IDs.
func (uuo *UserUpdateOne) AddTelegramChatIDs(ids ...int64) *UserUpdateOne {
	uuo.mutation.AddTelegramChatIDs(ids...)
	return uuo
}

// AddTelegramChats adds the "telegram_chats" edges to the TelegramChat entity.
func (uuo *UserUpdateOne) AddTelegramChats(t ...*TelegramChat) *UserUpdateOne {
	ids := make([]int64, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return uuo.AddTelegramChatIDs(ids...)
}

// AddDiscordChannelIDs adds the "discord_channels" edge to the DiscordChannel entity by IDs.
func (uuo *UserUpdateOne) AddDiscordChannelIDs(ids ...int64) *UserUpdateOne {
	uuo.mutation.AddDiscordChannelIDs(ids...)
	return uuo
}

// AddDiscordChannels adds the "discord_channels" edges to the DiscordChannel entity.
func (uuo *UserUpdateOne) AddDiscordChannels(d ...*DiscordChannel) *UserUpdateOne {
	ids := make([]int64, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return uuo.AddDiscordChannelIDs(ids...)
}

// AddWalletIDs adds the "wallets" edge to the Wallet entity by IDs.
func (uuo *UserUpdateOne) AddWalletIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.AddWalletIDs(ids...)
	return uuo
}

// AddWallets adds the "wallets" edges to the Wallet entity.
func (uuo *UserUpdateOne) AddWallets(w ...*Wallet) *UserUpdateOne {
	ids := make([]int, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return uuo.AddWalletIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uuo *UserUpdateOne) Mutation() *UserMutation {
	return uuo.mutation
}

// ClearTelegramChats clears all "telegram_chats" edges to the TelegramChat entity.
func (uuo *UserUpdateOne) ClearTelegramChats() *UserUpdateOne {
	uuo.mutation.ClearTelegramChats()
	return uuo
}

// RemoveTelegramChatIDs removes the "telegram_chats" edge to TelegramChat entities by IDs.
func (uuo *UserUpdateOne) RemoveTelegramChatIDs(ids ...int64) *UserUpdateOne {
	uuo.mutation.RemoveTelegramChatIDs(ids...)
	return uuo
}

// RemoveTelegramChats removes "telegram_chats" edges to TelegramChat entities.
func (uuo *UserUpdateOne) RemoveTelegramChats(t ...*TelegramChat) *UserUpdateOne {
	ids := make([]int64, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return uuo.RemoveTelegramChatIDs(ids...)
}

// ClearDiscordChannels clears all "discord_channels" edges to the DiscordChannel entity.
func (uuo *UserUpdateOne) ClearDiscordChannels() *UserUpdateOne {
	uuo.mutation.ClearDiscordChannels()
	return uuo
}

// RemoveDiscordChannelIDs removes the "discord_channels" edge to DiscordChannel entities by IDs.
func (uuo *UserUpdateOne) RemoveDiscordChannelIDs(ids ...int64) *UserUpdateOne {
	uuo.mutation.RemoveDiscordChannelIDs(ids...)
	return uuo
}

// RemoveDiscordChannels removes "discord_channels" edges to DiscordChannel entities.
func (uuo *UserUpdateOne) RemoveDiscordChannels(d ...*DiscordChannel) *UserUpdateOne {
	ids := make([]int64, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return uuo.RemoveDiscordChannelIDs(ids...)
}

// ClearWallets clears all "wallets" edges to the Wallet entity.
func (uuo *UserUpdateOne) ClearWallets() *UserUpdateOne {
	uuo.mutation.ClearWallets()
	return uuo
}

// RemoveWalletIDs removes the "wallets" edge to Wallet entities by IDs.
func (uuo *UserUpdateOne) RemoveWalletIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.RemoveWalletIDs(ids...)
	return uuo
}

// RemoveWallets removes "wallets" edges to Wallet entities.
func (uuo *UserUpdateOne) RemoveWallets(w ...*Wallet) *UserUpdateOne {
	ids := make([]int, len(w))
	for i := range w {
		ids[i] = w[i].ID
	}
	return uuo.RemoveWalletIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (uuo *UserUpdateOne) Select(field string, fields ...string) *UserUpdateOne {
	uuo.fields = append([]string{field}, fields...)
	return uuo
}

// Save executes the query and returns the updated User entity.
func (uuo *UserUpdateOne) Save(ctx context.Context) (*User, error) {
	var (
		err  error
		node *User
	)
	uuo.defaults()
	if len(uuo.hooks) == 0 {
		node, err = uuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			uuo.mutation = mutation
			node, err = uuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(uuo.hooks) - 1; i >= 0; i-- {
			if uuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = uuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, uuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (uuo *UserUpdateOne) SaveX(ctx context.Context) *User {
	node, err := uuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (uuo *UserUpdateOne) Exec(ctx context.Context) error {
	_, err := uuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uuo *UserUpdateOne) ExecX(ctx context.Context) {
	if err := uuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uuo *UserUpdateOne) defaults() {
	if _, ok := uuo.mutation.UpdatedAt(); !ok {
		v := user.UpdateDefaultUpdatedAt()
		uuo.mutation.SetUpdatedAt(v)
	}
}

func (uuo *UserUpdateOne) sqlSave(ctx context.Context) (_node *User, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   user.Table,
			Columns: user.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: user.FieldID,
			},
		},
	}
	id, ok := uuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "User.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := uuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, user.FieldID)
		for _, f := range fields {
			if !user.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != user.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := uuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: user.FieldUpdatedAt,
		})
	}
	if value, ok := uuo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldName,
		})
	}
	if value, ok := uuo.mutation.LoginToken(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldLoginToken,
		})
	}
	if uuo.mutation.TelegramChatsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   user.TelegramChatsTable,
			Columns: []string{user.TelegramChatsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: telegramchat.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedTelegramChatsIDs(); len(nodes) > 0 && !uuo.mutation.TelegramChatsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   user.TelegramChatsTable,
			Columns: []string{user.TelegramChatsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: telegramchat.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.TelegramChatsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   user.TelegramChatsTable,
			Columns: []string{user.TelegramChatsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: telegramchat.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uuo.mutation.DiscordChannelsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   user.DiscordChannelsTable,
			Columns: []string{user.DiscordChannelsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: discordchannel.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedDiscordChannelsIDs(); len(nodes) > 0 && !uuo.mutation.DiscordChannelsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   user.DiscordChannelsTable,
			Columns: []string{user.DiscordChannelsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: discordchannel.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.DiscordChannelsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   user.DiscordChannelsTable,
			Columns: []string{user.DiscordChannelsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: discordchannel.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uuo.mutation.WalletsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.WalletsTable,
			Columns: user.WalletsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: wallet.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedWalletsIDs(); len(nodes) > 0 && !uuo.mutation.WalletsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.WalletsTable,
			Columns: user.WalletsPrimaryKey,
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.WalletsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.WalletsTable,
			Columns: user.WalletsPrimaryKey,
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &User{config: uuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, uuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
