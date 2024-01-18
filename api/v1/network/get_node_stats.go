package network

import (
	"context"
	"crynux_bridge/api/v1/response"
	"crynux_bridge/blockchain"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/gin-gonic/gin"
)

type NodeStats struct {
	NumTotalNodes     uint64 `json:"num_total_nodes"`
	NumAvailableNodes uint64 `json:"num_available_nodes"`
}

type GetNodeStatsOutput struct {
	response.Response
	Data NodeStats `json:"data"`
}

func GetNodeStats(*gin.Context) (*GetNodeStatsOutput, error) {

	nodeContractInstance, err := blockchain.GetNodeContractInstance()
	if err != nil {
		return nil, response.NewExceptionResponse(err)
	}

	totalNodes, err := nodeContractInstance.TotalNodes(&bind.CallOpts{
		Pending: false,
		Context: context.Background(),
	})
	if err != nil {
		return nil, response.NewExceptionResponse(err)
	}

	availableNodes, err := nodeContractInstance.AvailableNodes(&bind.CallOpts{
		Pending: false,
		Context: context.Background(),
	})

	return &GetNodeStatsOutput{
		Data: NodeStats{
			NumAvailableNodes: availableNodes.Uint64(),
			NumTotalNodes:     totalNodes.Uint64(),
		},
	}, nil
}
