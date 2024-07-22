package biz

import (
	"context"
	"errors"
	"social-todo-list/common"
	"social-todo-list/modules/item/model"
)

type UpdateItemStorage interface {
	GetItem(ctx context.Context, cond map[string]interface{}) (*model.TodoItem, error)
	UpdateItem(ctx context.Context, cond map[string]interface{}, dataUpdate *model.TodoItemUpdate) error
}

type UpdateItemBiz struct {
	storage UpdateItemStorage
}

func NewUpdateItemBiz(storage UpdateItemStorage) *UpdateItemBiz {
	return &UpdateItemBiz{storage: storage}
}

func (biz *UpdateItemBiz) UpdateItemById(ctx context.Context, id int, dataUpdate *model.TodoItemUpdate) error {
	data, err := biz.storage.GetItem(ctx, map[string]interface{}{"id": id})

	if err != nil {
		if errors.Is(err, common.RecordNotFoundError) {
			return common.ErrCannotGetEntity(model.EntityName, err)
		}
		return common.ErrCannotUpdateEntity(model.EntityName, err)
	}

	if *data.Status == model.ItemStatusDeleted {
		return model.ErrItemDeleted
	}

	if err := biz.storage.UpdateItem(ctx, map[string]interface{}{"id": id}, dataUpdate); err != nil {
		return common.ErrCannotUpdateEntity(model.EntityName, err)
	}
	return nil
}
