package services

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
)

type StateService struct {
	RedisClient *redis.Client
}

type State struct {
	State string `json:"state"`
}

func (env StateService) SetState(chatId string, state string) error {
	rd := env.RedisClient

	stateJSON, err := json.Marshal(State{
		State: state,
	})
	if err != nil {
		return err
	}

	err = rd.Set(context.Background(), chatId, stateJSON, 0).Err() // TODO: Set expiration time
	if err != nil {
		return err
	}
	return nil
}

func (env StateService) GetState(chatId string) (string, error) {
	rd := env.RedisClient

	stateJSON, err := rd.Get(context.Background(), chatId).Result()
	log.Println("stateJSON:", stateJSON, err)
	if err == redis.Nil {
		return "", nil
	} else if err != nil {
		return "", err
	}
	if stateJSON == "" {
		fmt.Println("State not found")
		return "", nil
	}

	var state State
	err = json.Unmarshal([]byte(stateJSON), &state)
	if err != nil {
		return "", err
	}

	return state.State, nil
}

func (env StateService) DeleteState(chatId string) error {
	rd := env.RedisClient

	err := rd.Del(context.Background(), chatId).Err()
	if err != nil {
		return err
	}
	return nil
}
