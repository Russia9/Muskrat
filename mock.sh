#!/bin/bash

# Mock Repositories
mockgen -source pkg/domain/player.go -destination internal/player/repository/mock/player_mock.go -package mock -exclude_interfaces PlayerUsecase
