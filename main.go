package main

import (
	"log"

	"github.com/dezhishen/onebot-plus/pkg/plugin"
	"github.com/dezhishen/onebot-sdk/pkg/api"
	"github.com/dezhishen/onebot-sdk/pkg/model"
)

func main() {
	plugin.OnebotPluginBuilder().
		Id("demo").
		Name("demo").
		Init(func(cli api.OnebotApiClientInterface) error {
			log.Println("init")
			return nil
		}).
		BeforeExit(func(cli api.OnebotApiClientInterface) error {
			log.Println("before exit")
			return nil
		}).
		HandleMessagePrivate(func(data *model.EventMessagePrivate, onebotApi api.OnebotApiClientInterface) error {
			log.Println("message private")
			// reply hello
			onebotApi.SendPrivateMsg(
				&model.PrivateMsg{
					UserId: data.Sender.UserId,
					Message: []*model.MessageSegment{
						{
							Type: "text",
							Data: &model.MessageElementText{
								Text: "hello" + data.Sender.Nickname,
							},
						},
					},
				},
			)
			return nil
		}).
		HandleMessageGroup(func(data *model.EventMessageGroup, onebotApi api.OnebotApiClientInterface) error {
			// reply hello
			onebotApi.SendGroupMsg(
				&model.GroupMsg{
					GroupId: data.GroupId,
					Message: []*model.MessageSegment{
						{
							Type: "text",
							Data: &model.MessageElementText{
								Text: "hello" + data.Sender.Nickname,
							},
						},
					},
				},
			)
			return nil
		}).
		Build().
		Start()
}
