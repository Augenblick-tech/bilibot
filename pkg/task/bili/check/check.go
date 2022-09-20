package check

import (
	"net/http"

	bilibot "github.com/Augenblick-tech/bilibot/lib/bili_bot"
	"github.com/Augenblick-tech/bilibot/pkg/e"
	"github.com/Augenblick-tech/bilibot/pkg/email"
	"github.com/Augenblick-tech/bilibot/pkg/services/bot"
	"github.com/Augenblick-tech/bilibot/pkg/task/basetask"
	"github.com/Augenblick-tech/bilibot/pkg/utils"
)

type BotLoginInfo struct {
	basetask.BaseTask
	spec  string
	BotID string
}

func New(spec, BotID string) *BotLoginInfo {
	return &BotLoginInfo{
		spec:  spec,
		BotID: BotID,
	}
}

func NewWithAttr(spec string, attr map[string]interface{}) *BotLoginInfo {
	return New(spec, attr["BotID"].(string))
}

func (b *BotLoginInfo) Run() {
	defer func() {
		if r := recover(); r != nil {
			if b.Status == basetask.Running {
				email.SendEmail(1, "发生错误", r)
				b.Status = basetask.Warning
				panic(r)
			}
		}
	}()

	Bot, err := bot.Get(b.BotID)
	if err != nil {
		panic(err)
	}
	botLoginInfo, err := bilibot.GetBotInfo(&http.Cookie{Name: "SESSDATA", Value: utils.StrToMap(Bot.Cookie)["SESSDATA"]})
	if err != nil {
		panic(err)
	}
	if !botLoginInfo.Data.IsLogin {
		panic(e.ErrNotLogin)
	}
}

func (b *BotLoginInfo) Name() string {
	return "Check " + b.BotID
}

func (b *BotLoginInfo) Attribute() interface{} {
	return struct {
		BotID string
	}{
		BotID: b.BotID,
	}
}

func (b *BotLoginInfo) Data() interface{} {
	return nil
}

func (b *BotLoginInfo) SetStatus(s basetask.Status) {
	b.Status = s
}

func (b *BotLoginInfo) Spec() string {
	return b.spec
}
