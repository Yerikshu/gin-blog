package service

import (
	"context"
	"github.com/golang-minibear2333/gin-blog/global"
	"github.com/golang-minibear2333/gin-blog/internal/dao"
)

type Service struct {
	ctx context.Context
	dao *dao.Dao
}
// New 所有Service都走这里的逻辑
func New(ctx context.Context) Service {
	svc := Service{ctx: ctx}
	svc.dao = dao.New(global.DBEngine)
	return svc
}