package openrouter

import (
	"context"
	"crynux_bridge/api/v1/response"
	"crynux_bridge/api/v1/tools"
	"crynux_bridge/models"
	"errors"
	"slices"
	"strings"

	"gorm.io/gorm"
)

// validate api key
func ValidateRequestApiKey(ctx context.Context, db *gorm.DB, authorization string) (*models.ClientAPIKey, error) {
	if !strings.HasPrefix(authorization, "Bearer ") {
		return nil, response.NewValidationErrorResponse("Authorization", "Authorization header must start with 'Bearer '")
	}
	apiKeyStr := authorization[7:]
	apiKey, err := tools.ValidateAPIKey(ctx, db, apiKeyStr)
	if err != nil {
		if errors.Is(err, tools.ErrAPIKeyExpired) {
			return nil, response.NewValidationErrorResponse("Authorization", "expired")
		}
		if errors.Is(err, tools.ErrAPIKeyInvalid) {
			return nil, response.NewValidationErrorResponse("Authorization", "unauthorized")
		}
		return nil, response.NewExceptionResponse(err)
	}
	if !slices.Contains(apiKey.Roles, models.RoleAdmin) && !slices.Contains(apiKey.Roles, models.RoleChat) {
		return nil, response.NewValidationErrorResponse("Authorization", "unauthorized")
	}
	if apiKey.UseLimit > 0 && apiKey.UsedCount >= apiKey.UseLimit {
		return nil, response.NewValidationErrorResponse("Authorization", "use limit exceeded")
	}

	return apiKey, nil
}
