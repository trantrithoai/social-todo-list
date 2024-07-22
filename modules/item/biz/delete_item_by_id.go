package biz

import (
	"context"
	"social-todo-list/modules/item/model"
)

type DeleteItemStorage interface {
	GetItem(ctx context.Context, cond map[string]interface{}) (*model.TodoItem, error)
	DeleteItem(ctx context.Context, cond map[string]interface{}) error
}

type DeleteItemBiz struct {
	storage DeleteItemStorage
}

func NewDeleteItemBiz(storage DeleteItemStorage) *DeleteItemBiz {
	return &DeleteItemBiz{storage: storage}
}

func (biz *DeleteItemBiz) DeleteItemById(ctx context.Context, id int) error {
	data, err := biz.storage.GetItem(ctx, map[string]interface{}{"id": id})

	if err != nil {
		return err
	}

	if *data.Status == model.ItemStatusDeleted {
		return model.ErrItemDeleted
	}

	if err := biz.storage.DeleteItem(ctx, map[string]interface{}{"id": id}); err != nil {
		return err
	}
	return nil
}
