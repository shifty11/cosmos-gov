package admin

import (
	"context"
	pb "github.com/shifty11/cosmos-gov/api/grpc/protobuf/go/admin_service"
	"github.com/shifty11/cosmos-gov/database"
	"github.com/shifty11/cosmos-gov/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

//goland:noinspection GoNameStartsWithPackageName
type ChainServer struct {
	pb.UnimplementedAdminServiceServer
	chainManager *database.ChainManager
}

func NewAdminServer(chainManager *database.ChainManager) pb.AdminServiceServer {
	return &ChainServer{chainManager: chainManager}
}

func (server *ChainServer) GetChains(_ context.Context, _ *emptypb.Empty) (*pb.GetChainsResponse, error) {
	chains := server.chainManager.All()
	var chainSettings []*pb.ChainSettings
	for _, c := range chains {
		chainSettings = append(chainSettings, &pb.ChainSettings{
			ChainId:         c.ChainID,
			Name:            c.Name,
			DisplayName:     c.DisplayName,
			IsEnabled:       c.IsEnabled,
			IsVotingEnabled: false,
			IsFeegrantUsed:  false,
		})
	}

	var res = &pb.GetChainsResponse{Chains: chainSettings}
	return res, nil
}

func (server *ChainServer) SetEnabled(_ context.Context, req *pb.SetChainEnabledRequest) (*pb.SetChainEnabledResponse, error) {
	chain, err := server.chainManager.EnableOrDisableChain(req.ChainId)
	if err != nil {
		log.Sugar.Errorf("error while toggling chain: %v", err)
		return nil, status.Errorf(codes.Internal, "Unknown error occurred")
	}
	var res = &pb.SetChainEnabledResponse{IsEnabled: chain.IsEnabled}
	return res, nil
}
