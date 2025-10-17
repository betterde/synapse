package service

import restate "github.com/restatedev/sdk-go"

type MessageService struct{}

func (MessageService) GetMessages(ctx restate.Context) error {
	return nil
}

func (MessageService) GetMessage(ctx restate.Context) error {
	return nil
}

func (MessageService) CreateMessage(ctx restate.Context) error {
	return nil
}

func (MessageService) DeleteMessage(ctx restate.Context) error {
	return nil
}

func (MessageService) SendMessage(ctx restate.Context) error {
	return nil
}
