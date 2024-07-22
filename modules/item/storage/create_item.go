package storage

import (
	"golang.org/x/net/context"
	"social-todo-list/common"
	"social-todo-list/modules/item/model"
)

func (s *SqlStore) CreateItem(ctx context.Context, data *model.TodoItemCreation) error {
	if err := s.db.Create(&data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
