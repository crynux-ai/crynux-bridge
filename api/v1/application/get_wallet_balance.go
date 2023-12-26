package application

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"ig_server/api/v1/response"
	"ig_server/blockchain"
	"ig_server/config"
	"math/big"
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
