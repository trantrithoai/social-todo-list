package storage

import (
	"context"
	"social-todo-list/modules/item/model"
)

func (s *SqlStore) DeleteItem(ctx context.Context, cond map[string]interface{}) error {
	deletedStatus := model.ItemStatusDeleted
	err := s.db.Table(model.TodoItem{}.TableName()).Where(cond).Updates(map[string]interface{}{"status": deletedStatus.String()}).Error
	if err != nil {
		return err
	}
	return nil
}
