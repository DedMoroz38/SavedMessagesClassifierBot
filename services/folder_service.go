package services

import "github.com/redis/go-redis/v9"

type FolderService struct {
	RedisClient *redis.Client
}

func (env FolderService) CreateFolder(name string) error {
	
	
	return nil
}