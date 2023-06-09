package models

import (
	"errors"
	"reflect"

	"github.com/aabdullahgungor/mybookcase/database"
	"github.com/aabdullahgungor/mybookcase/entities"
)

var (
	ErrPublisherIDIsNotValid   = errors.New("id is not valid")
	ErrPublisherNameIsNotEmpty = errors.New("publisher name cannot be empty")
	ErrPublisherNotFound       = errors.New("the publisher cannot be found")
)

type PublisherModel struct {
}

func (publisherModel PublisherModel) GetAll() ([]entities.Publisher, error) {
	db, err := database.GetDB()
	if err != nil {
		return nil, err
	} else {
		var publishers []entities.Publisher
		db.Find(&publishers)
		return publishers, nil
	}
}

func (publisherModel PublisherModel) GetById(id int) (entities.Publisher, error) {
	if id <= 0 || reflect.TypeOf(id).Kind() != reflect.Int {
		return entities.Publisher{}, ErrPublisherIDIsNotValid
	}
	db, err := database.GetDB()
	if err != nil {
		return entities.Publisher{}, err
	} else {
		var publisher entities.Publisher
		db.Where("id = ?", id).First(&publisher)
		return publisher, nil
	}
}

func (publisherModel PublisherModel) Create(publisher *entities.Publisher) error {

	if publisher.PublisherName == "" {
		return ErrPublisherNameIsNotEmpty
	} else {
		db, err := database.GetDB()
		if err != nil {
			return err
		} else {
			db.Create(&publisher)
			return nil
		}
	}

}

func (publisherModel PublisherModel) Edit(publisher *entities.Publisher) error {

	if publisher.PublisherName == "" {
		return ErrPublisherNameIsNotEmpty
	} else {
		db, err := database.GetDB()
		if err != nil {
			return err
		} else {
			db.Save(&publisher)
			return nil
		}

	}
}

func (publisherModel PublisherModel) Delete(id int) error {
	publisher, err := publisherModel.GetById(id)
	if err != nil {
		return err
	}
	db, err := database.GetDB()
	if err != nil {
		return err
	}
	db.Delete(publisher)
	return nil
}
