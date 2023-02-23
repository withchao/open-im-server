package fcm

import (
	"OpenIM/internal/push"
	"OpenIM/pkg/common/config"
	"OpenIM/pkg/common/constant"
	"OpenIM/pkg/common/db/cache"
	"context"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"github.com/go-redis/redis/v8"
	"google.golang.org/api/option"
	"path/filepath"
)

const SinglePushCountLimit = 400

var Terminal = []int{constant.IOSPlatformID, constant.AndroidPlatformID, constant.WebPlatformID}

type Fcm struct {
	fcmMsgCli *messaging.Client
	cache     cache.MsgCache
}

func NewClient(cache cache.MsgCache) *Fcm {
	opt := option.WithCredentialsFile(filepath.Join(config.Root, "config", config.Config.Push.Fcm.ServiceAccount))
	fcmApp, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return nil
	}
	// auth
	// fcmClient, err := fcmApp.Auth(context.Background())
	// if err != nil {
	// 	return
	// }
	ctx := context.Background()
	fcmMsgClient, err := fcmApp.Messaging(ctx)
	if err != nil {
		panic(err.Error())
		return nil
	}
	return &Fcm{fcmMsgCli: fcmMsgClient}
}

func (f *Fcm) Push(ctx context.Context, userIDs []string, title, content string, opts *push.Opts) error {
	// accounts->registrationToken
	allTokens := make(map[string][]string, 0)
	for _, account := range userIDs {
		var personTokens []string
		for _, v := range Terminal {
			Token, err := f.cache.GetFcmToken(ctx, account, v)
			if err == nil {
				personTokens = append(personTokens, Token)
			}
		}
		allTokens[account] = personTokens
	}
	Success := 0
	Fail := 0
	notification := &messaging.Notification{}
	notification.Body = content
	notification.Title = title
	var messages []*messaging.Message
	for userID, personTokens := range allTokens {
		apns := &messaging.APNSConfig{Payload: &messaging.APNSPayload{Aps: &messaging.Aps{Sound: opts.IOSPushSound}}}
		messageCount := len(messages)
		if messageCount >= SinglePushCountLimit {
			response, err := f.fcmMsgCli.SendAll(ctx, messages)
			if err != nil {
				Fail = Fail + messageCount
			} else {
				Success = Success + response.SuccessCount
				Fail = Fail + response.FailureCount
			}
			messages = messages[0:0]
		}
		if opts.IOSBadgeCount {
			unreadCountSum, err := f.cache.IncrUserBadgeUnreadCountSum(ctx, userID)
			if err == nil {
				apns.Payload.Aps.Badge = &unreadCountSum
			} else {
				//log.Error(operationID, "IncrUserBadgeUnreadCountSum redis err", err.Error(), uid)
				Fail++
				continue
			}
		} else {
			unreadCountSum, err := f.cache.GetUserBadgeUnreadCountSum(ctx, userID)
			if err == nil && unreadCountSum != 0 {
				apns.Payload.Aps.Badge = &unreadCountSum
			} else if err == redis.Nil || unreadCountSum == 0 {
				zero := 1
				apns.Payload.Aps.Badge = &zero
			} else {
				//log.Error(operationID, "GetUserBadgeUnreadCountSum redis err", err.Error(), uid)
				Fail++
				continue
			}
		}
		for _, token := range personTokens {
			temp := &messaging.Message{
				Data:         map[string]string{"ex": opts.Ex},
				Token:        token,
				Notification: notification,
				APNS:         apns,
			}
			messages = append(messages, temp)
		}
	}
	messageCount := len(messages)
	if messageCount > 0 {
		response, err := f.fcmMsgCli.SendAll(ctx, messages)
		if err != nil {
			Fail = Fail + messageCount
			//log.Info(operationID, "some token push err", err.Error(), messageCount)
		} else {
			Success = Success + response.SuccessCount
			Fail = Fail + response.FailureCount
		}
	}
	return nil
}
