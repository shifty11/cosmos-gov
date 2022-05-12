package admin

import (
	"context"
	pb "github.com/shifty11/cosmos-gov/api/grpc/protobuf/go/admin_service"
	"github.com/shifty11/cosmos-gov/database"
	"github.com/shifty11/cosmos-gov/ent"
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

func chainToProtobuf(chain *ent.Chain) *pb.ChainSettings {
	return &pb.ChainSettings{
		ChainId:         chain.ChainID,
		Name:            chain.Name,
		DisplayName:     chain.DisplayName,
		IsEnabled:       chain.IsEnabled,
		IsVotingEnabled: chain.IsVotingEnabled,
		IsFeegrantUsed:  chain.IsFeegrantUsed,
	}
}

func (server *ChainServer) GetChains(_ context.Context, _ *emptypb.Empty) (*pb.GetChainsResponse, error) {
	chains := server.chainManager.All()
	var chainSettings []*pb.ChainSettings
	for _, chain := range chains {
		chainSettings = append(chainSettings, chainToProtobuf(chain))
	}

	var res = &pb.GetChainsResponse{Chains: chainSettings}
	return res, nil
}

func (server *ChainServer) UpdateChain(_ context.Context, req *pb.UpdateChainRequest) (*pb.UpdateChainResponse, error) {
	chain, err := server.chainManager.Update(req.ChainName, req.IsEnabled, req.IsVotingEnabled, req.IsFeegrantUsed)
	if err != nil {
		log.Sugar.Errorf("error while updating chain: %v", err)
		return nil, status.Errorf(codes.Internal, "Unknown error occurred")
	}
	var res = &pb.UpdateChainResponse{Chain: chainToProtobuf(chain)}
	return res, nil
}
