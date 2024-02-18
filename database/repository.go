package database

import (
	"context"
	"errors"

	"github.com/psbernardo/graphql/graph/model"
	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func NewRepository(DB *gorm.DB) Repository {
	return Repository{
		DB: DB,
	}
}

func (r Repository) SaveSearchResult(ctx context.Context, key string, searchResult *model.QueryResult) error {
	tx := r.DB.WithContext(ctx)
	existsData, err := r.GetSaveResult(ctx, key)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	if existsData != nil {
		return nil
	}

	return tx.
		Model(&KeyValue{}).
		Create(&KeyValue{
			Key:  key,
			Data: NewQueryResultJSONB(*searchResult),
		}).Error

}

func (r Repository) GetSaveResult(ctx context.Context, key string) (*model.QueryResult, error) {
	var result KeyValue

	if err := r.DB.WithContext(ctx).
		Model(&KeyValue{}).
		Where("key=?", key).
		First(&result).Error; err != nil {
		return nil, err
	}
	return &result.Data.value, nil
}

func (r Repository) GetSearchHistoryKey(ctx context.Context) ([]string, error) {
	var keyList []string
	if err := r.DB.WithContext(ctx).
		Model(&KeyValue{}).
		Pluck("key", &keyList).Error; err != nil {
		return keyList, err
	}

	return keyList, nil
}
