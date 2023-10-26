package codetest

import (
	"github.com/xinliangnote/go-gin-api/configs"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	"github.com/xinliangnote/go-gin-api/internal/repository/mysql"
	"github.com/xinliangnote/go-gin-api/internal/repository/redis"
	"github.com/xinliangnote/go-gin-api/internal/services/codetest"
	"github.com/xinliangnote/go-gin-api/pkg/hash"

	"go.uber.org/zap"
)

var _ Handler = (*handler)(nil)

type Handler interface {
	i()

	// Create 新增调用方
	// @Tags API.codeTest
	// @Router /api/testCreate [post]
	Create() core.HandlerFunc

	// Delete 删除调用方
	// @Tags API.codeTest
	// @Router /api/testDelete/{id} [delete]
	Delete() core.HandlerFunc

	Update() core.HandlerFunc
}

type handler struct {
	logger           *zap.Logger
	cache            redis.Repo
	contestedService codetest.Service
	hashids          hash.Hash
}

func New(logger *zap.Logger, db mysql.Repo, cache redis.Repo) Handler {
	return &handler{
		logger:           logger,
		cache:            cache,
		contestedService: codetest.New(db, cache),
		hashids:          hash.New(configs.Get().HashIds.Secret, configs.Get().HashIds.Length),
	}
}

func (h *handler) i() {}
