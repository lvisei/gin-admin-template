package bson

import (
	"context"
	"encoding/json"
	"time"

	"gin-admin-template/pkg/logger"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Config 配置参数
type Config struct {
	URI        string
	Database   string
	Collection string
	Timeout    time.Duration
}

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}

// New 创建基于bson的钩子实例(需要指定表名)
func New(cfg *Config) *Hook {
	var (
		ctx    = context.Background()
		cancel context.CancelFunc
	)

	if t := cfg.Timeout; t > 0 {
		ctx, cancel = context.WithTimeout(ctx, t)
		defer cancel()
	}

	cli, err := mongo.Connect(ctx, options.Client().ApplyURI(cfg.URI))
	handleError(err)
	c := cli.Database(cfg.Database).Collection(cfg.Collection)

	return &Hook{
		Client:     cli,
		Collection: c,
	}
}

// Hook bson日志钩子
type Hook struct {
	Client     *mongo.Client
	Collection *mongo.Collection
}

// Exec 执行日志写入
func (h *Hook) Exec(entry *logrus.Entry) error {
	item := &LogItem{
		ID:        primitive.NewObjectID().Hex(),
		Level:     entry.Level.String(),
		Message:   entry.Message,
		CreatedAt: entry.Time,
	}

	data := entry.Data
	if v, ok := data[logger.TraceIDKey]; ok {
		item.TraceID, _ = v.(string)
		delete(data, logger.TraceIDKey)
	}
	if v, ok := data[logger.UserIDKey]; ok {
		item.UserID, _ = v.(string)
		delete(data, logger.UserIDKey)
	}
	if v, ok := data[logger.TagKey]; ok {
		item.Tag, _ = v.(string)
		delete(data, logger.TagKey)
	}
	if v, ok := data[logger.StackKey]; ok {
		item.ErrorStack, _ = v.(string)
		delete(data, logger.StackKey)
	}
	if v, ok := data[logger.VersionKey]; ok {
		item.Version, _ = v.(string)
		delete(data, logger.VersionKey)
	}

	if len(data) > 0 {
		b, _ := json.Marshal(data)
		item.Data = string(b)
	}

	_, err := h.Collection.InsertOne(context.Background(), item)
	return err
}

// Close 关闭钩子
func (h *Hook) Close() error {
	return h.Client.Disconnect(context.Background())
}

// LogItem 存储日志项
type LogItem struct {
	ID         string    `bson:"_id"`         // id
	Level      string    `bson:"level"`       // 日志级别
	TraceID    string    `bson:"trace_id"`    // 跟踪ID
	UserID     string    `bson:"user_id"`     // 用户ID
	Tag        string    `bson:"tag"`         // Tag
	Version    string    `bson:"version"`     // 版本号
	Message    string    `bson:"message"`     // 消息
	Data       string    `bson:"data"`        // 日志数据(json)
	ErrorStack string    `bson:"error_stack"` // Error Stack
	CreatedAt  time.Time `bson:"created_at"`  // 创建时间
}
