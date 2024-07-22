package storage

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"social-todo-list/common"
	"social-todo-list/modules/item/model"
)

func (s *SqlStore) GetItem(ctx context.Context, cond map[string]interface{}) (*model.TodoItem, error) {
	var data model.TodoItem

	if err := s.db.Where(cond).First(&data).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, common.RecordNotFoundError
		}
		return nil, common.ErrDB(err)
	}
	return &data, nil
}
