package web

import (
	"strconv"

	"github.com/Augenblick-tech/bilibot/lib/engine"
	"github.com/Augenblick-tech/bilibot/pkg/services/author"
	"github.com/Augenblick-tech/bilibot/pkg/services/bot"
	"github.com/Augenblick-tech/bilibot/pkg/services/dynamic"
)

func GetBotList(c *engine.Context) (interface{}, error) {
	UserID, err := strconv.Atoi(c.Query("user_id")) // 暂时通过 Query 获取（测试用）
	if err != nil {
		return nil, err
	}

	return bot.GetList(uint(UserID))
}

func GetAuthorList(c *engine.Context) (interface{}, error) {
	BotID, err := strconv.Atoi(c.Query("bot_id")) // 暂时通过 Query 获取（测试用）
	if err != nil {
		return nil, err
	}

	return author.GetList(uint(BotID))
}

func GetDynamicList(c *engine.Context) (interface{}, error) {
	AuthorID := c.Query("author_id") // 暂时通过 Query 获取（测试用）
	return dynamic.GetList(AuthorID)
}
