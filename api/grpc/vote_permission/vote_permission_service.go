package vote_permission

import (
	"context"
	pb "github.com/shifty11/cosmos-gov/api/grpc/protobuf/go/vote_permission_service"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

//goland:noinspection GoNameStartsWithPackageName
type VotePermissionServer struct {
	pb.UnimplementedVotePermissionServiceServer
	//votePermissionManager *database.VotePermissionManager
}

//func NewVotePermissionsServer(votePermissionManager *database.VotePermissionManager) pb.VotePermissionServiceServer {
//	return &VotePermissionServer{votePermissionManager: votePermissionManager}
//}

func NewVotePermissionsServer() pb.VotePermissionServiceServer {
	return &VotePermissionServer{}
}

func (server *VotePermissionServer) GetVotePermissions(ctx context.Context, _ *emptypb.Empty) (*pb.GetVotePermissionsResponse, error) {
	//entUser, ok := ctx.Value("user").(*ent.User)
	//if !ok {
	//	log.Sugar.Error("invalid user")
	//	return nil, status.Errorf(codes.Internal, "invalid user")
	//}

	//subsDtos := server.votePermissionManager.GetVotePermissions(entUser.ChatID, entUser.Type)
	//var vperms []*pb.VotePermission
	//for _, sub := range subsDtos {
	//	vperms = append(vperms, &pb.VotePermission{
	//		Name:         sub.Name,
	//		DisplayName:  sub.DisplayName,
	//		IsSubscribed: sub.Notify,
	//	})
	//}

	var vperms []*pb.VotePermission
	var expiresAt = timestamppb.Timestamp{
		Seconds: time.Now().Add(time.Hour * 24 * time.Duration(365)).Unix(),
	}
	vperms = append(vperms, &pb.VotePermission{
		Address:   "cosmos1fhp54fwlfmpwwgrnfwk3v47v53yjtp8ffqng4q",
		ExpiresAt: &expiresAt,
	})

	var res = &pb.GetVotePermissionsResponse{VotePermissions: vperms}
	return res, nil
}

func (server *VotePermissionServer) RefreshVotePermission(ctx context.Context, req *pb.RefreshVotePermissionRequest) (*pb.RefreshVotePermissionResponse, error) {
	//entUser, ok := ctx.Value("user").(*ent.User)
	//if !ok {
	//	log.Sugar.Error("invalid user")
	//	return nil, status.Errorf(codes.Internal, "invalid user")
	//}
	//
	//isSubscribed, err := server.votePermissionManager.RefreshVotePermission(entUser.ChatID, entUser.Type, req.Name)
	//if err != nil {
	//	log.Sugar.Error("error while toggling votePermission: %v", err)
	//	return nil, status.Errorf(codes.Internal, "Unknown error occured")
	//}

	var res = &pb.RefreshVotePermissionResponse{Success: true}
	return res, nil
}
