package vote_permission

import (
	"context"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	pb "github.com/shifty11/cosmos-gov/api/grpc/protobuf/go/vote_permission_service"
	"github.com/shifty11/cosmos-gov/authz"
	"github.com/shifty11/cosmos-gov/database"
	"github.com/shifty11/cosmos-gov/ent"
	"github.com/shifty11/cosmos-gov/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

//goland:noinspection GoNameStartsWithPackageName
type VotePermissionServer struct {
	pb.UnimplementedVotePermissionServiceServer
	authzClient   *authz.AuthzClient
	chainManager  *database.ChainManager
	walletManager *database.WalletManager
}

func NewVotePermissionsServer(authzClient *authz.AuthzClient, chainManager *database.ChainManager, walletManager *database.WalletManager) pb.VotePermissionServiceServer {
	return &VotePermissionServer{authzClient: authzClient, chainManager: chainManager, walletManager: walletManager}
}

func getGranteeAddress(entChain *ent.Chain) (string, error) {
	const granteeCosmos = "cosmos1wtcvjqx8097gtkjdemle9c0lm8gczm2aac5p4n"
	granteeBytes, err := sdk.GetFromBech32(granteeCosmos, "cosmos")
	if err != nil {
		log.Sugar.Errorf("Could not convert address %v to Cosmos address", granteeCosmos)
		return "", err
	}

	grantee, err := sdk.Bech32ifyAddressBytes(entChain.AccountPrefix, granteeBytes)
	if err != nil {
		log.Sugar.Errorf("Could not convert address %v to %v address", granteeCosmos, entChain.DisplayName)
		return "", err
	}
	return grantee, nil
}

func chainToProtobuf(c *ent.Chain) (*pb.Chain, error) {
	grantee, err := getGranteeAddress(c)
	if err != nil {
		return nil, err
	}
	return &pb.Chain{
		ChainId:     c.ChainID,
		Name:        c.Name,
		DisplayName: c.DisplayName,
		//RpcAddress:     c.Edges.RPCEndpoints[0].Endpoint,
		RpcAddress:     fmt.Sprintf("https://rpc.cosmos.directory/%v", c.Name),
		Grantee:        grantee,
		Denom:          "u" + c.AccountPrefix,
		AccountPrefix:  c.AccountPrefix,
		IsFeegrantUsed: c.IsFeegrantUsed,
	}, nil
}

func walletToProtobuf(wallet *ent.Wallet) (*pb.Wallet, error) {
	pbChain, err := chainToProtobuf(wallet.Edges.Chain)
	if err != nil {
		return nil, err
	}
	var vperms []*pb.VotePermission
	for _, g := range wallet.Edges.Grants {
		vperms = append(vperms, &pb.VotePermission{
			ExpiresAt: &timestamppb.Timestamp{
				Seconds: g.ExpiresAt.Unix(),
			},
		})
	}
	return &pb.Wallet{
		Chain:           pbChain,
		Address:         wallet.Address,
		VotePermissions: vperms,
	}, nil
}

func (server *VotePermissionServer) GetSupportedChains(context.Context, *emptypb.Empty) (*pb.GetSupportedChainsResponse, error) {
	var chains []*pb.Chain
	for _, c := range server.chainManager.CanVote() {
		if len(c.Edges.RPCEndpoints) == 0 {
			log.Sugar.Errorf("Chain %v has no RPC endpoint", c.DisplayName)
			continue
		}

		pbChain, err := chainToProtobuf(c)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "unknown error")
		}
		chains = append(chains, pbChain)
	}
	return &pb.GetSupportedChainsResponse{Chains: chains}, nil
}

func (server *VotePermissionServer) RegisterWallet(ctx context.Context, req *pb.RegisterWalletRequest) (*emptypb.Empty, error) {
	entUser, ok := ctx.Value("user").(*ent.User)
	if !ok {
		log.Sugar.Error("invalid user")
		return nil, status.Errorf(codes.Internal, "invalid user")
	}

	if req.ChainName == "" || req.WalletAddress == "" {
		log.Sugar.Errorf("RegisterWallet request doesn't have all the necessary arguments: %v", req.String())
		return nil, status.Errorf(codes.InvalidArgument, "arguments are missing")
	}

	//TODO: make this with streaming -> client opens request -> gets a random token -> sends token as memo -> sends response that it's done -> server checks and creates wallet if successful
	entChain, err := server.chainManager.ByName(req.ChainName)
	if err != nil {
		log.Sugar.Errorf("Error while getting chain %v for user %v (%v): %v", req.ChainName, entUser.Name, entUser.ID, err)
		return nil, status.Errorf(codes.Internal, "unknown error")
	}

	if server.walletManager.Exists(entUser, entChain, req.WalletAddress) {
		return &emptypb.Empty{}, nil
	}

	_, err = server.walletManager.Create(entUser, entChain, req.WalletAddress)
	if err != nil {
		log.Sugar.Errorf("Error while creating wallet for user %v (%v): %v", entUser.Name, entUser.ID, err)
		return nil, status.Errorf(codes.Internal, "unknown error")
	}

	grantee, err := getGranteeAddress(entChain)
	if err != nil {
		log.Sugar.Errorf("Error while getting grantee address for user %v (%v): %v", entUser.Name, entUser.ID, err)
		return nil, status.Errorf(codes.Internal, "bad request")
	}

	grant, err := server.authzClient.GetGrant(req.ChainName, req.WalletAddress, grantee)
	if err != nil {
		log.Sugar.Errorf("Error while getting grants for user %v (%v): %v", entUser.Name, entUser.ID, err)
		//return nil, status.Errorf(codes.Internal, "bad request")
		//TODO: check if this is proper
		return &emptypb.Empty{}, nil
	}

	if grant == nil {
		return &emptypb.Empty{}, nil
	}
	_, err = server.walletManager.SaveGrant(entUser, req.ChainName, grant)
	if err != nil {
		log.Sugar.Errorf("Error while saving grant for user %v (%v): %v", entUser.Name, entUser.ID, err)
		return nil, status.Errorf(codes.Internal, "could not save grant")
	}

	return &emptypb.Empty{}, nil
}

func (server *VotePermissionServer) RemoveWallet(ctx context.Context, req *pb.RemoveWalletRequest) (*emptypb.Empty, error) {
	entUser, ok := ctx.Value("user").(*ent.User)
	if !ok {
		log.Sugar.Error("invalid user")
		return nil, status.Errorf(codes.Internal, "invalid user")
	}

	if req.WalletAddress == "" {
		log.Sugar.Errorf("RemoveWallet request doesn't have all the necessary arguments: %v", req.String())
		return nil, status.Errorf(codes.InvalidArgument, "arguments are missing")
	}

	_, err := server.walletManager.Delete(entUser, req.WalletAddress)
	if err != nil {
		log.Sugar.Errorf("Error while deleting wallet of user %v (%v): %v", entUser.Name, entUser.ID, err)
		return nil, status.Errorf(codes.InvalidArgument, "arguments are missing")
	}

	return &emptypb.Empty{}, nil
}

func (server *VotePermissionServer) GetWallets(ctx context.Context, _ *emptypb.Empty) (*pb.GetWalletsResponse, error) {
	entUser, ok := ctx.Value("user").(*ent.User)
	if !ok {
		log.Sugar.Error("invalid user")
		return nil, status.Errorf(codes.Internal, "invalid user")
	}

	wallets, err := server.walletManager.ByUser(entUser)
	if err != nil {
		log.Sugar.Errorf("Error while querying grants of user %v (%v): %v", entUser.Name, entUser.ID, err)
		return nil, status.Errorf(codes.Internal, "unknown error")
	}

	var pbWallets []*pb.Wallet
	for _, w := range wallets {
		pbWallet, err := walletToProtobuf(w)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "unknown error")
		}
		pbWallets = append(pbWallets, pbWallet)
	}
	var res = &pb.GetWalletsResponse{Wallets: pbWallets}
	return res, nil
}

func (server *VotePermissionServer) RefreshVotePermission(ctx context.Context, req *pb.RefreshVotePermissionRequest) (*pb.RefreshVotePermissionResponse, error) {
	entUser, ok := ctx.Value("user").(*ent.User)
	if !ok {
		log.Sugar.Error("invalid user")
		return nil, status.Errorf(codes.Internal, "invalid user")
	}

	wallet, err := server.walletManager.ByUserAndAddress(entUser, req.Granter)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "unknown error")
	}

	grant, err := server.authzClient.GetGrant(req.ChainName, req.Granter, req.Grantee)
	if grant == nil && err != nil {
		_, err = server.walletManager.DeleteGrant(req.ChainName, req.Granter, req.Grantee)
		if err != nil {
			log.Sugar.Errorf("Error while deleting grants for user %v (%v): %v", entUser.Name, entUser.ID, err)
			return nil, status.Errorf(codes.Internal, "bad request")
		}
		return &pb.RefreshVotePermissionResponse{}, nil
	} else if grant != nil && err == nil {
		_, err := server.walletManager.SaveGrant(entUser, req.ChainName, grant)
		if err != nil {
			return nil, err
		}
		if err != nil {
			log.Sugar.Errorf("Error while creating grants for user %v (%v): %v", entUser.Name, entUser.ID, err)
			return nil, status.Errorf(codes.Internal, "bad request")
		}
	}

	pbWallet, err := walletToProtobuf(wallet)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "unknown error")
	}

	return &pb.RefreshVotePermissionResponse{Wallet: pbWallet}, nil
}
