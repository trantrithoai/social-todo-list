package biz

import (
	"context"
	"social-todo-list/common"
	"social-todo-list/modules/item/model"
	"strings"
)

type CreateItemStorage interface {
	CreateItem(ctx context.Context, data *model.TodoItemCreation) error
}

type CreateItemBiz struct {
	storage CreateItemStorage
}

func NewCreateItemBiz(storage CreateItemStorage) *CreateItemBiz {
	return &CreateItemBiz{storage: storage}
}

func (biz *CreateItemBiz) CreateNewItem(ctx context.Context, data *model.TodoItemCreation) error {
	title := strings.TrimSpace(data.Title)

	if title == "" {
		return model.ErrTitleIsEmpty
	}

	if err := biz.storage.CreateItem(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(model.EntityName, err)
	}
	return nil
}
