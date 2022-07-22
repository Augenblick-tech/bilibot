package main

import (
	"encoding/json"
	"errors"
	"log"
	"io/ioutil"

	"github.com/lonzzi/BiliUpDynamicBot/e"
)

func MakeReply(dynamic BriefDynamic, message string) error {
	filePath := "./dynamic_history.json"

	// 读出历史数据
	fileContent, err := ioutil.ReadFile(filePath)
	if err != nil {
		return errors.New(e.ReadFileError + ":" + err.Error())
	}
	var oldDynamics []BriefDynamic
	if err := json.Unmarshal(fileContent, &oldDynamics); err != nil && len(fileContent) != 0 {
		return errors.New(e.UnmarshalFileError + ":" + err.Error())
	}
	log.Println("Old Dynamics: ", oldDynamics)

	oldDynamics, err = AddNewDynamic(oldDynamics, dynamic)
	if err != nil {
		return errors.New(e.AddNewDynamicError + ":" + err.Error())
	} else {
		_, err := DynamicReply(dynamic, message)
		if err != nil {
			return errors.New(e.DynamicReplyError + ":" + err.Error())
		}
	}

	// 写入新数据
	fileContent, err = json.Marshal(oldDynamics)
	if err != nil {
		return errors.New(e.MarshalFileError + ":" + err.Error())
	}
	if err := ioutil.WriteFile(filePath, fileContent, 0644); err != nil {
		return errors.New(e.WriteFileError + ":" + err.Error())
	}

	return nil
}