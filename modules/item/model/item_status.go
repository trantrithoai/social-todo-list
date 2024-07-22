package model

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"strings"
)

type ItemStatus int

const (
	ItemStatusDoing ItemStatus = iota
	ItemStatusDone
	ItemStatusDeleted
)

var allItemStatus = [3]string{"Doing", "Done", "Deleted"}

func (item *ItemStatus) String() string {
	return allItemStatus[*item]
}

func parseStringToItemStatus(s string) (ItemStatus, error) {
	for i := range allItemStatus {
		if allItemStatus[i] == s {
			return ItemStatus(i), nil
		}
	}
	return ItemStatusDoing, errors.New(s + " is not a valid item status")
}

func (item *ItemStatus) Scan(value interface{}) error {
	bytes, ok := value.([]byte)

	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}

	v, err := parseStringToItemStatus(string(bytes))

	if err != nil {
		return err
	}

	*item = v

	return nil
}

func (item *ItemStatus) Value() (driver.Value, error) {
	if item == nil {
		return nil, nil
	}
	return item.String(), nil
}

func (item *ItemStatus) MarshalJSON() ([]byte, error) {
	if item == nil {
		return nil, nil
	}

	return []byte(fmt.Sprintf("\"%s\"", item.String())), nil
}

func (item *ItemStatus) UnmarshalJSON(b []byte) error {
	str := strings.ReplaceAll(string(b), "\"", "")

	itemStatus, err := parseStringToItemStatus(str)

	if err != nil {
		return err
	}

	*item = itemStatus
	return nil
}
