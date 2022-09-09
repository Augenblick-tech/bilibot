package user

import (
	"github.com/Augenblick-tech/bilibot/pkg/dao"
	"github.com/Augenblick-tech/bilibot/pkg/e"
	"github.com/Augenblick-tech/bilibot/pkg/model"
	"github.com/Augenblick-tech/bilibot/pkg/services/author"
	"github.com/Augenblick-tech/bilibot/pkg/services/bot"
)

func Add(user model.User) error {
	return dao.Create(&user)
}

func Get(username string) (model.User, error) {
	return dao.GetUserByName(username)
}

func GetByID(id string) (model.User, error) {
	return dao.GetUserByID(id)
}

func CheckRecordWithID(id uint, buID ...string) error {
	if len(buID) <= 0 {
		return e.ErrInvalidParam
	}
	Bot, err := bot.Get(buID[0])
	if err != nil {
		return err
	}

	if Bot.UserID == id {
		if len(buID) == 2 {
			Author, err := author.Get(buID[1])
			if err != nil {
				return err
			}

			if Author.BotID == buID[0] {
				return nil
			} else {
				return e.ErrNotFound
			}
		}
		return nil
	} else {
		return e.ErrNotFound
	}
}

// 待改进，无法区分 Bot 与 Author
// func CheckRecordWithID(id uint, checkID string) error {
// 	Author,err := author.Get(checkID)
// 	if err != nil {
// 		Bot, err1 := bot.Get(checkID)
// 		if err1 != nil {
// 			return err1
// 		}
// 		if Bot.UserID != id {
// 			return e.RespCode_NO_RECORD
// 		} else {
// 			return nil
// 		}
// 	} else {
// 		Bot, err1 := bot.Get(Author.BotID)
// 		if err1 != nil {
// 			return err1
// 		}
// 		if Bot.UserID != id {
// 			return e.RespCode_NO_RECORD
// 		} else {
// 			return nil
// 		}
// 	}
// }
