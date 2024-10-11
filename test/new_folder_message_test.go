package services

import (
	"context"
	"encoding/json"
	"errors"
	"saved_messages_classifier/constants"
	"saved_messages_classifier/services"
	"testing"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockRedisClient struct {
	mock.Mock
}

func (m *MockRedisClient) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.StatusCmd {
	args := m.Called(ctx, key, value, expiration)
	return redis.NewStatusResult(args.String(0), args.Error(1))
}

func TestSetState(t *testing.T) {
	mockRedisClient := new(MockRedisClient)
	stateService := services.StateService{RedisClient: mockRedisClient}

	state := services.State{
		ChatId:    12345,
		StateName: constants.AddFolder,
	}

	stateJSON, err := json.Marshal(state)
	assert.NoError(t, err)

	mockRedisClient.On("Set", mock.Anything, "state", stateJSON, time.Duration(0)).Return("", nil)

	err = stateService.SetState(state)
	assert.NoError(t, err)

	mockRedisClient.AssertExpectations(t)
}

func TestSetState_Error(t *testing.T) {
	mockRedisClient := new(MockRedisClient)
	stateService := services.StateService{RedisClient: mockRedisClient}

	state := services.State{
		ChatId:    12345,
		StateName: "AddFolder",
	}

	stateJSON, err := json.Marshal(state)
	assert.NoError(t, err)

	// Set up expectation for Set method to return an error
	mockRedisClient.On("Set", mock.Anything, "state", stateJSON, time.Duration(0)).Return("", errors.New("some error"))

	err = stateService.SetState(state)
	assert.Error(t, err)

	mockRedisClient.AssertExpectations(t)
}
