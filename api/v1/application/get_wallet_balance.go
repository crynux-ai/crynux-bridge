package application

import (
	"context"
	"crynux_bridge/api/v1/response"
	"crynux_bridge/blockchain"
	"crynux_bridge/config"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
)

type WalletBalance struct {
	Address string   `json:"address"`
	ETH     *big.Int `json:"eth"`
	CNX     *big.Int `json:"cnx"`
}

type GetWalletBalanceResponse struct {
	response.Response
	Data *WalletBalance `json:"data"`
}

func GetWalletBalance(_ *gin.Context) (*GetWalletBalanceResponse, error) {
	appConfig := config.GetConfig()
	applicationWalletAddress := common.HexToAddress(appConfig.Blockchain.Account.Address)

	client, err := blockchain.GetRpcClient()
	if err != nil {
		return nil, response.NewExceptionResponse(err)
	}

	ethBalance, err := client.BalanceAt(
		context.Background(),
		applicationWalletAddress,
		nil,
	)

	if err != nil {
		return nil, response.NewExceptionResponse(err)
	}

	cnxTokenInstance, err := blockchain.GetCrynuxTokenContractInstance()
	if err != nil {
		return nil, response.NewExceptionResponse(err)
	}

	cnxBalance, err := cnxTokenInstance.BalanceOf(
		&bind.CallOpts{
			Pending: false,
			Context: context.Background(),
		},
		applicationWalletAddress,
	)
	if err != nil {
		return nil, response.NewExceptionResponse(err)
	}

	return &GetWalletBalanceResponse{
		Data: &WalletBalance{
			Address: appConfig.Blockchain.Account.Address,
			CNX:     cnxBalance,
			ETH:     ethBalance,
		},
	}, nil
}
