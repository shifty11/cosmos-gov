package database

import (
	"context"
	"github.com/shifty11/cosmos-gov/ent"
	"github.com/shifty11/cosmos-gov/ent/chain"
	"github.com/shifty11/cosmos-gov/ent/grant"
	"github.com/shifty11/cosmos-gov/ent/wallet"
	"time"
)

type WalletManager struct {
	client       *ent.Client
	ctx          context.Context
	chainManager *ChainManager
}

func NewWalletManager(client *ent.Client, ctx context.Context, chainManager *ChainManager) *WalletManager {
	return &WalletManager{client: client, ctx: ctx, chainManager: chainManager}
}

func (manager *WalletManager) ByAddress(address string) (*ent.Wallet, error) {
	return manager.client.Wallet.
		Query().
		Where(wallet.AddressEQ(address)).
		Only(manager.ctx)
}

func (manager *WalletManager) ByUser(entUser *ent.User) ([]*ent.Wallet, error) {
	return entUser.
		QueryWallets().
		WithGrants().
		WithChain().
		All(manager.ctx)
}

func (manager *WalletManager) ByUserAndChain(entUser *ent.User, entChain *ent.Chain) ([]*ent.Wallet, error) {
	return entUser.
		QueryWallets().
		WithGrants().
		WithChain().
		Where(wallet.HasChainWith(chain.ChainID(entChain.ChainID))).
		All(manager.ctx)
}

type GrantData struct {
	Granter   string
	Grantee   string
	Type      string
	ExpiresAt time.Time
}

func (manager *WalletManager) SaveGrant(entUser *ent.User, chainName string, g *GrantData) (*ent.Grant, error) {
	w, err := entUser.
		QueryWallets().
		Where(wallet.And(
			wallet.AddressEQ(g.Granter),
			wallet.HasChainWith(chain.NameEQ(chainName)),
		)).
		First(manager.ctx)
	if err != nil { // if wallet doesn't exist -> create wallet and grant (inside a transaction)
		entChain, err := manager.chainManager.ByName(chainName)
		if err != nil {
			return nil, err
		}
		if err := WithTx(manager.ctx, manager.client, func(tx *ent.Tx) error {
			w, err = manager.client.Wallet.
				Create().
				SetAddress(g.Granter).
				SetChain(entChain).
				AddUsers(entUser).
				Save(manager.ctx)
			if err != nil {
				return err
			}
			return manager.client.Grant.
				Create().
				SetGrantee(g.Grantee).
				SetType(g.Type).
				SetExpiresAt(g.ExpiresAt).
				SetGranter(w).
				Exec(manager.ctx)
		}); err != nil {
			return nil, err
		}
		return entUser.
			QueryWallets().
			Where(wallet.And(
				wallet.AddressEQ(g.Granter),
				wallet.HasChainWith(chain.NameEQ(chainName)),
			)).
			QueryGrants().
			Where(grant.And(
				grant.GranteeEQ(g.Grantee),
				grant.TypeEQ(g.Type),
			)).
			First(manager.ctx)
	} else {
		entGrant, err := w.
			QueryGrants().
			Where(grant.And(
				grant.GranteeEQ(g.Grantee),
				grant.TypeEQ(g.Type),
			)).
			First(manager.ctx)
		if err != nil { // create grant if it doesn't exist
			return manager.client.Grant.
				Create().
				SetGrantee(g.Grantee).
				SetType(g.Type).
				SetExpiresAt(g.ExpiresAt).
				SetGranter(w).
				Save(manager.ctx)
		} else { // update grant if it exists
			return entGrant.
				Update().
				SetExpiresAt(g.ExpiresAt).
				Save(manager.ctx)
		}
	}
}

func (manager *WalletManager) DeleteGrant(chainName string, granter string, grantee string) (int, error) {
	return manager.client.Grant.
		Delete().
		Where(grant.And(
			grant.GranteeEQ(grantee),
			grant.HasGranterWith(
				wallet.AddressEQ(granter),
				wallet.HasChainWith(chain.NameEQ(chainName)),
			))).
		Exec(manager.ctx)
}
