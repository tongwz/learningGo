package main

import (
	"errors"
	"fmt"
)

func main() {
	myEx(0)
}

func myEx(userId int) {
	talkingMemberInfo, err := findMember(userId)
	if err != nil {
		fmt.Printf("查询数据库异常：%s \n", err.Error())
		talkingMemberInfo = findMember2(userId)

		// 其他逻辑 使用talkingMemberInfo
		fmt.Printf("err内部的talkingMmber中的messageId = %d \n", talkingMemberInfo.MessageId)
		fmt.Printf("我们的talkingMemberInfo的指针地址：%p \n", &talkingMemberInfo)
	}
	// 其他逻辑
	fmt.Printf("我们的talkingMmber中的messageId = %d \n", talkingMemberInfo.MessageId)
	fmt.Printf("我们的talkingMemberInfo的指针地址：%p \n", &talkingMemberInfo)
}

type TalkingMember struct {
	UserId           int
	MessageId        int
	MessageDirection string
	MessageNickname  string
}

func findMember(userId int) (TalkingMember, error) {
	if userId > 0 {
		return TalkingMember{
			UserId:           userId,
			MessageId:        100,
			MessageDirection: "test",
			MessageNickname:  "test",
		}, nil
	}
	return TalkingMember{}, errors.New("错误J")
}

func findMember2(userId int) TalkingMember {
	return TalkingMember{
		UserId:           userId,
		MessageId:        200,
		MessageDirection: "test2",
		MessageNickname:  "test2",
	}
}
