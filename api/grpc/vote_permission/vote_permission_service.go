package vote_permission

import (
	"context"
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

func (server *VotePermissionServer) GetSupportedChains(context.Context, *emptypb.Empty) (*pb.GetSupportedChainsResponse, error) {
	var chains []*pb.Chain
	options := &database.ChainQueryOptions{WithRpcAddresses: true}
	for _, c := range server.chainManager.Enabled(options) {
		if len(c.Edges.RPCEndpoints) == 0 {
			log.Sugar.Errorf("Chain %v has no RPC endpoint", c.DisplayName)
			continue
		}

		grantee, err := getGranteeAddress(c)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "unknown error")
		}

		chains = append(chains, &pb.Chain{
			ChainId:     c.ChainID,
			Name:        c.Name,
			DisplayName: c.DisplayName,
			RpcAddress:  c.Edges.RPCEndpoints[0].Endpoint,
			Grantee:     grantee,
			Denom:       "u" + c.AccountPrefix,
		})
	}
	return &pb.GetSupportedChainsResponse{Chains: chains}, nil
}

func (server *VotePermissionServer) CreateVotePermission(ctx context.Context, req *pb.CreateVotePermissionRequest) (*pb.CreateVotePermissionResponse, error) {
	entUser, ok := ctx.Value("user").(*ent.User)
	if !ok {
		log.Sugar.Error("invalid user")
		return nil, status.Errorf(codes.Internal, "invalid user")
	}

	if req.VotePermission == nil ||
		req.VotePermission.Chain.Name == "" ||
		req.VotePermission.Chain.DisplayName == "" ||
		req.VotePermission.Chain.ChainId == "" ||
		req.VotePermission.Chain.Grantee == "" ||
		req.VotePermission.Granter == "" {
		log.Sugar.Errorf("CreateVotePermission request doesn't have all the necessary arguments: %v", req.VotePermission.String())
		return nil, status.Errorf(codes.InvalidArgument, "arguments are missing")
	}

	grant, err := server.authzClient.GetGrant(req.VotePermission.Chain.Name, req.VotePermission.Granter, req.VotePermission.Chain.Grantee)
	if err != nil {
		log.Sugar.Errorf("Error while getting grants for user %v (%v): %v", entUser.Name, entUser.ID, err)
		return nil, status.Errorf(codes.Internal, "bad request")
	}
	if grant == nil {
		log.Sugar.Errorf("Grant for user %v (%v) was not found", entUser.Name, entUser.ID)
		return nil, status.Errorf(codes.NotFound, "grant was not found")
	}

	entGrant, err := server.walletManager.SaveGrant(entUser, req.VotePermission.Chain.Name, grant)
	if err != nil {
		log.Sugar.Errorf("Error while saving grants for user %v (%v): %v", entUser.Name, entUser.ID, err)
		return nil, status.Errorf(codes.Internal, "could not save grant")
	}

	var votePerm = &pb.VotePermission{
		Chain:   req.VotePermission.Chain,
		Granter: req.VotePermission.Granter,
		ExpiresAt: &timestamppb.Timestamp{
			Seconds: entGrant.ExpiresAt.Unix(),
		},
	}
	var res = &pb.CreateVotePermissionResponse{VotePermission: votePerm}
	return res, nil
}

func (server *VotePermissionServer) GetVotePermissions(ctx context.Context, _ *emptypb.Empty) (*pb.GetVotePermissionsResponse, error) {
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

	var vperms []*pb.VotePermission
	for _, w := range wallets {
		for _, g := range w.Edges.Grants {
			grantee, err := getGranteeAddress(w.Edges.Chain)
			if err != nil {
				return nil, status.Errorf(codes.Internal, "unknown error")
			}
			vperms = append(vperms, &pb.VotePermission{
				Chain: &pb.Chain{
					ChainId:     w.Edges.Chain.ChainID,
					Name:        w.Edges.Chain.Name,
					DisplayName: w.Edges.Chain.DisplayName,
					RpcAddress:  server.chainManager.GetFirstRpc(w.Edges.Chain).Endpoint,
					Grantee:     grantee,
					Denom:       "u" + w.Edges.Chain.AccountPrefix,
				},
				Granter: w.Address,
				ExpiresAt: &timestamppb.Timestamp{
					Seconds: g.ExpiresAt.Unix(),
				},
			})
		}
	}
	var res = &pb.GetVotePermissionsResponse{VotePermissions: vperms}
	return res, nil
}

func (server *VotePermissionServer) RefreshVotePermission(ctx context.Context, req *pb.RefreshVotePermissionRequest) (*pb.RefreshVotePermissionResponse, error) {
	entUser, ok := ctx.Value("user").(*ent.User)
	if !ok {
		log.Sugar.Error("invalid user")
		return nil, status.Errorf(codes.Internal, "invalid user")
	}

	grant, err := server.authzClient.GetGrant(req.VotePermission.Chain.Name, req.VotePermission.Granter, req.VotePermission.Chain.Grantee)
	if grant == nil && err != nil {
		_, err = server.walletManager.DeleteGrant(req.VotePermission.Chain.Name, req.VotePermission.Granter, req.VotePermission.Chain.Grantee)
		if err != nil {
			log.Sugar.Errorf("Error while deleting grants for user %v (%v): %v", entUser.Name, entUser.ID, err)
			return nil, status.Errorf(codes.Internal, "bad request")
		}
		return &pb.RefreshVotePermissionResponse{}, nil
	}
	return &pb.RefreshVotePermissionResponse{VotePermission: &pb.VotePermission{
		Chain:     req.VotePermission.Chain,
		Granter:   grant.Granter,
		ExpiresAt: &timestamppb.Timestamp{Seconds: grant.ExpiresAt.Unix()},
	}}, nil
}
