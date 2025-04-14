package tools

import (
	"context"
	"crynux_bridge/models"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func GenerateAPIKey(ctx context.Context, db *gorm.DB, clientID string) (string, int64, error) {
	randKey := make([]byte, 32)
	_, err := rand.Read(randKey)
	if err != nil {
		return "", 0, err
	}
	apiKeyStr := base64.StdEncoding.EncodeToString(randKey)
	keyPrefix := apiKeyStr[:8]
	hashKey, err := bcrypt.GenerateFromPassword(randKey, bcrypt.DefaultCost)
	if err != nil {
		return "", 0, err
	}
	hashKeyStr := base64.StdEncoding.EncodeToString(hashKey)

	now := time.Now()
	expiresAt := now.Add(time.Hour * 24 * 365) // 1 year expiration
	// Generate a new API key
	var apiKey *models.ClientAPIKey
	apiKey, err = models.GetAPIKeyByClientID(ctx, db, clientID)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		apiKey = &models.ClientAPIKey{
			ClientID:   clientID,
			KeyPrefix:  keyPrefix,
			KeyHash:    hashKeyStr,
			LastUsedAt: now,
			ExpiresAt:  expiresAt, // 1 year expiration
		}
		if err := apiKey.Save(ctx, db); err != nil {
			return "", 0, err
		}
	} else if err != nil {
		return "", 0, err
	} else {
		apiKey.KeyPrefix = keyPrefix
		apiKey.KeyHash = hashKeyStr
		apiKey.LastUsedAt = now
		apiKey.ExpiresAt = expiresAt
		if err := apiKey.Save(ctx, db); err != nil {
			return "", 0, err
		}
	}

	return apiKeyStr, expiresAt.Unix(), nil
}

var ErrAPIKeyInvalid = errors.New("API key is invalid")
var ErrAPIKeyExpired = errors.New("API key is expired")

func ValidateAPIKey(ctx context.Context, db *gorm.DB, apiKeyStr string) (*models.ClientAPIKey, error) {
	rawKey, err := base64.StdEncoding.DecodeString(apiKeyStr)
	if err != nil {
		return nil, ErrAPIKeyInvalid
	}
	keyPrefix := apiKeyStr[:8]
	apiKey, err := models.GetAPIKeyByKeyPrefix(ctx, db, keyPrefix)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrAPIKeyInvalid
		}
		return nil, err
	}
	hashKey, _ := base64.StdEncoding.DecodeString(apiKey.KeyHash)
	if apiKey.ExpiresAt.Before(time.Now()) {
		return nil, ErrAPIKeyExpired
	}
	err = bcrypt.CompareHashAndPassword(hashKey, rawKey)
	if err != nil {
		return nil, ErrAPIKeyInvalid
	}
	return apiKey, nil
}

func DeleteAPIKey(ctx context.Context, db *gorm.DB, clientID string) error {
	apiKey, err := models.GetAPIKeyByClientID(ctx, db, clientID)
	if err != nil {
		return err
	}
	return apiKey.Delete(ctx, db)
}

func UpdateAPIKeyUsedAt(ctx context.Context, db *gorm.DB, clientID string) error {
	apiKey, err := models.GetAPIKeyByClientID(ctx, db, clientID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		}
		return err
	}
	return apiKey.Update(ctx, db, &models.ClientAPIKey{
		LastUsedAt: time.Now(),
	})
}
