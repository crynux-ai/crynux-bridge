package relay

import (
	"context"
	"crynux_bridge/config"
	"crynux_bridge/models"
	"net/http"
	"strconv"
	"time"

	log "github.com/sirupsen/logrus"
)

type GetNodeByAddressInput struct {
	Address string `json:"address"`
}

func GetNodeByAddress(ctx context.Context, address string) (*models.RelayNode, error) {
	appConfig := config.GetConfig()

	params := &GetNodeByAddressInput{
		Address: address,
	}

	timestamp, signature, err := SignData(params, appConfig.Blockchain.Account.PrivateKey)
	if err != nil {
		return nil, err
	}

	reqUrl := appConfig.Relay.BaseURL + "/v1/node/" + address
	timeoutCtx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	req, _ := http.NewRequestWithContext(timeoutCtx, "GET", reqUrl, nil)
	query := req.URL.Query()
	query.Add("timestamp", strconv.FormatInt(timestamp, 10))
	query.Add("signature", signature)
	req.URL.RawQuery = query.Encode()
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if err := processRelayResponse(resp); err != nil {
		log.Errorf("Relay: get node by address %s error: %v", address, err)
		return nil, err
	}

	node := new(models.RelayNode)
	err = parseRelayResponseData(resp, node)
	if err != nil {
		log.Errorf("Relay: get node by address %s error: %v", address, err)
		return nil, err
	}

	log.Debugf("Relay: get node %s", address)
	return node, nil
}
