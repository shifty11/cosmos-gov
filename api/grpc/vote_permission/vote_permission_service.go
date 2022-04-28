package vote_permission

import (
	"context"
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
	//votePermissionManager *database.VotePermissionManager
	authzClient   *authz.AuthzClient
	walletManager *database.WalletManager
}

func NewVotePermissionsServer(authzClient *authz.AuthzClient, walletManager *database.WalletManager) pb.VotePermissionServiceServer {
	return &VotePermissionServer{authzClient: authzClient, walletManager: walletManager}
}

func (server *VotePermissionServer) CreateVotePermission(ctx context.Context, req *pb.CreateVotePermissionRequest) (*pb.CreateVotePermissionResponse, error) {
	entUser, ok := ctx.Value("user").(*ent.User)
	if !ok {
		log.Sugar.Error("invalid user")
		return nil, status.Errorf(codes.Internal, "invalid user")
	}

	if req.VotePermission == nil || req.VotePermission.ChainName == "" || req.VotePermission.Granter == "" || req.VotePermission.Grantee == "" {
		return nil, status.Errorf(codes.InvalidArgument, "arguments are missing")
	}

	grant, err := server.authzClient.GetGrant(req.VotePermission.Granter, req.VotePermission.Grantee)
	if err != nil {
		log.Sugar.Errorf("Error while getting grants for user %v (%v): %v", entUser.Name, entUser.ID, err)
		return nil, status.Errorf(codes.Internal, "bad request")
	}
	if grant == nil {
		log.Sugar.Errorf("Grant for user %v (%v) was not found", entUser.Name, entUser.ID)
		return nil, status.Errorf(codes.NotFound, "grant was not found")
	}

	entGrant, err := server.walletManager.SaveGrant(entUser, req.VotePermission.ChainName, grant)
	if err != nil {
		log.Sugar.Errorf("Error while saving grants for user %v (%v): %v", entUser.Name, entUser.ID, err)
		return nil, status.Errorf(codes.Internal, "could not save grant")
	}

	var votePerm = &pb.VotePermission{
		ChainName: req.VotePermission.ChainName,
		Granter:   req.VotePermission.Granter,
		Grantee:   entGrant.Grantee,
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

	wallets, err := server.walletManager.OfUser(entUser)
	if err != nil {
		log.Sugar.Errorf("Error while querying grants of user %v (%v): %v", entUser.Name, entUser.ID, err)
		return nil, status.Errorf(codes.Internal, "unknown error")
	}

	var vperms []*pb.VotePermission
	for _, w := range wallets {
		for _, g := range w.Edges.Grants {
			vperms = append(vperms, &pb.VotePermission{
				ChainName: w.Edges.Chain.Name,
				Granter:   w.Address,
				Grantee:   g.Grantee,
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

	grant, err := server.authzClient.GetGrant(req.VotePermission.Granter, req.VotePermission.Grantee)
	if err != nil {
		log.Sugar.Errorf("Error while getting grants for user %v (%v): %v", entUser.Name, entUser.ID, err)
		return nil, status.Errorf(codes.Internal, "bad request")
	}

	if grant == nil {
		_, err = server.walletManager.DeleteGrant(req.VotePermission.ChainName, req.VotePermission.Granter, req.VotePermission.Grantee)
		if err != nil {
			log.Sugar.Errorf("Error while deleting grants for user %v (%v): %v", entUser.Name, entUser.ID, err)
			return nil, status.Errorf(codes.Internal, "bad request")
		}
		return &pb.RefreshVotePermissionResponse{}, nil
	}
	return &pb.RefreshVotePermissionResponse{VotePermission: &pb.VotePermission{
		ChainName: req.VotePermission.ChainName,
		Granter:   grant.Granter,
		Grantee:   grant.Grantee,
		ExpiresAt: &timestamppb.Timestamp{Seconds: grant.ExpiresAt.Unix()},
	}}, nil
}
