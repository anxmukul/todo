package controller

import (
	"github.com/anxmukul/todo/model"
)
type Controller interface {
	Create(string,string) (*model.ToDo, error)
	SearchById(int64) (*model.ToDo, error)
	SearchByTitle(string) (*[]model.ToDo, error)
	DeleteByTitle(string) (*model.ToDo)
}
