package relay

import (
	"context"
	"crynux_bridge/config"
	"encoding/json"
	"io"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

type QueuedTasksCountResponse struct {
	Message string `json:"message"`
	Data    int64  `json:"data"`
}

func GetQueuedTasks(ctx context.Context) (int64, error) {
	appConfig := config.GetConfig()
	reqUrl := appConfig.Relay.BaseURL + "/v1/stats/queue/count"

	timeoutCtx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	req, _ := http.NewRequestWithContext(timeoutCtx, "GET", reqUrl, nil)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	if err := processRelayResponse(resp); err != nil {
		log.Errorf("Relay: GetQueuedTasks error: %v", err)
		return 0, err
	}

	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	res := QueuedTasksCountResponse{}
	if err := json.Unmarshal(respBytes, &res); err != nil {
		return 0, err
	}
	return res.Data, nil
}
