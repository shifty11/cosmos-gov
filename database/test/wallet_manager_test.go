package db_testing_base_test

import (
	_ "github.com/mattn/go-sqlite3"
	"github.com/shifty11/cosmos-gov/database"
	"github.com/shifty11/cosmos-gov/database/test"
	"github.com/shifty11/cosmos-gov/ent/grant"
	"github.com/shifty11/cosmos-gov/ent/wallet"
	"testing"
	"time"
)

func Test_SaveGrant_New(t *testing.T) {
	client, ctx, m := db_testing_base.GetBase(t)
	//goland:noinspection GoUnhandledErrorResult
	defer client.Close()

	entChain := m.ChainManager.Create("testchain", []string{"test-rpc"})
	entUser := m.TelegramUserManager.GetOrCreateUser(100, "Roland")

	grantData := &database.GrantData{
		Granter:   "testchainadsfasdf",
		Grantee:   "testchainkljsadflkasjdf",
		Type:      "/cosmos.gov.v1beta1.MsgVote",
		ExpiresAt: time.Now(),
	}

	g, err := m.WalletManager.SaveGrant(entUser, entChain.Name, grantData)
	if err != nil {
		t.Errorf("could not create g")
	}
	if g.Grantee != grantData.Grantee {
		t.Errorf("value error")
	}
	if g.Type != grantData.Type {
		t.Errorf("value error")
	}

	w, err := m.WalletManager.ByUser(entUser)
	if err != nil {
		t.Errorf("no create wallet")
	}
	if len(w) != 1 {
		t.Errorf("wrong number of wallets: %v", len(w))
	}

	sameWallet, err := g.QueryGranter().Only(ctx)
	if err != nil {
		t.Errorf("error")
	}
	if w[0].ID != sameWallet.ID {
		t.Errorf("not the same wallet")
	}
}

func Test_SaveGrant_ExistingWallet(t *testing.T) {
	client, ctx, m := db_testing_base.GetBase(t)
	//goland:noinspection GoUnhandledErrorResult
	defer client.Close()

	entChain := m.ChainManager.Create("testchain", []string{"test-rpc"})
	entUser := m.TelegramUserManager.GetOrCreateUser(100, "Roland")

	grantData := &database.GrantData{
		Granter:   "testchainadsfasdf",
		Grantee:   "testchainkljsadflkasjdf",
		Type:      "/cosmos.gov.v1beta1.MsgVote",
		ExpiresAt: time.Now(),
	}

	client.Wallet.
		Create().
		SetAddress(grantData.Granter).
		AddUsers(entUser).
		SetChain(entChain).
		ExecX(ctx)

	g, err := m.WalletManager.SaveGrant(entUser, entChain.Name, grantData)
	if err != nil {
		t.Errorf("could not create g")
	}
	if g.Grantee != grantData.Grantee {
		t.Errorf("value error")
	}
	if g.Type != grantData.Type {
		t.Errorf("value error")
	}

	w, err := m.WalletManager.ByUser(entUser)
	if err != nil {
		t.Errorf("no create wallet")
	}
	if len(w) != 1 {
		t.Errorf("wrong number of wallets: %v", len(w))
	}

	sameWallet, err := g.QueryGranter().Only(ctx)
	if err != nil {
		t.Errorf("error")
	}
	if w[0].ID != sameWallet.ID {
		t.Errorf("not the same wallet")
	}
}

func Test_SaveGrant_ExistingGrant_Update(t *testing.T) {
	client, ctx, m := db_testing_base.GetBase(t)
	//goland:noinspection GoUnhandledErrorResult
	defer client.Close()

	entChain := m.ChainManager.Create("testchain", []string{"test-rpc"})
	entUser := m.TelegramUserManager.GetOrCreateUser(100, "Roland")

	grantData := &database.GrantData{
		Granter:   "testchainadsfasdf",
		Grantee:   "testchainkljsadflkasjdf",
		Type:      "/cosmos.gov.v1beta1.MsgVote",
		ExpiresAt: time.Now(),
	}

	newW := client.Wallet.
		Create().
		SetAddress(grantData.Granter).
		AddUsers(entUser).
		SetChain(entChain).
		SaveX(ctx)
	client.Grant.
		Create().
		SetGrantee(grantData.Grantee).
		SetGranter(newW).
		SetType(grantData.Type).
		SetExpiresAt(grantData.ExpiresAt.Add(time.Hour * 24)).
		SaveX(ctx)

	g, err := m.WalletManager.SaveGrant(entUser, entChain.Name, grantData)
	if err != nil {
		t.Errorf("could not create g")
	}
	if g.Grantee != grantData.Grantee {
		t.Errorf("value error")
	}
	if g.Type != grantData.Type {
		t.Errorf("value error")
	}
	if g.ExpiresAt.Unix() != grantData.ExpiresAt.Unix() {
		t.Errorf("value error")
	}

	w, err := m.WalletManager.ByUser(entUser)
	if err != nil {
		t.Errorf("no create wallet")
	}
	if len(w) != 1 {
		t.Errorf("wrong number of wallets: %v", len(w))
	}

	sameWallet, err := g.QueryGranter().Only(ctx)
	if err != nil {
		t.Errorf("error")
	}
	if w[0].ID != sameWallet.ID {
		t.Errorf("not the same wallet")
	}
}

func Test_SaveGrant_ExistingGrant_Create(t *testing.T) {
	client, ctx, m := db_testing_base.GetBase(t)
	//goland:noinspection GoUnhandledErrorResult
	defer client.Close()

	entChain := m.ChainManager.Create("testchain", []string{"test-rpc"})
	entUser := m.TelegramUserManager.GetOrCreateUser(100, "Roland")

	grantData := &database.GrantData{
		Granter:   "testchainadsfasdf",
		Grantee:   "testchainkljsadflkasjdf",
		Type:      "/cosmos.gov.v1beta1.MsgVote",
		ExpiresAt: time.Now(),
	}

	newW := client.Wallet.
		Create().
		SetAddress(grantData.Granter).
		AddUsers(entUser).
		SetChain(entChain).
		SaveX(ctx)
	client.Grant.
		Create().
		SetGrantee(grantData.Grantee).
		SetGranter(newW).
		SetType("x").
		SetExpiresAt(grantData.ExpiresAt).
		SaveX(ctx)

	g, err := m.WalletManager.SaveGrant(entUser, entChain.Name, grantData)
	if err != nil {
		t.Errorf("could not create g")
	}
	if g.Grantee != grantData.Grantee {
		t.Errorf("value error")
	}
	if g.Type != grantData.Type {
		t.Errorf("value error")
	}

	w, err := m.WalletManager.ByUser(entUser)
	if err != nil {
		t.Errorf("no create wallet")
	}
	if len(w) != 1 {
		t.Errorf("wrong number of wallets: %v", len(w))
	}

	sameWallet, err := g.QueryGranter().Only(ctx)
	if err != nil {
		t.Errorf("error")
	}
	if w[0].ID != sameWallet.ID {
		t.Errorf("not the same wallet")
	}

	grants, err := client.Grant.
		Query().
		Where(grant.HasGranterWith(wallet.AddressEQ(grantData.Granter))).
		All(ctx)
	if err != nil {
		t.Errorf("no create wallet")
	}
	if len(grants) != 2 {
		t.Errorf("wrong number of wallets: %v", len(grants))
	}
}
