package controller

import (
	"errors"
	"testing"

	"github.com/anxmukul/todo/mocks"
	"github.com/anxmukul/todo/model"
	"github.com/anxmukul/todo/view"
	"github.com/stretchr/testify/assert"
)

func TestTodoController_Create(t *testing.T) {
	t.Run("shoudl return the inserted todo when no error in model and view", func(t *testing.T) {
		mockedModel := &mocks.TodoModel{}
		mockedView := &mocks.TodoDisplayer{}
		mockedTodo := &model.ToDo{
			Id:      100,
			Title:   "foo",
			Content: "bar",
		}
		mockedModel.On("CreateTodo", "foo", "bar").Return(mockedTodo, nil)
		mockedView.On("ShowTodo", int64(100), "foo", "bar").Return(nil)

		testController := TodoController{
			model: mockedModel,
			view:  mockedView,
		}

		receivedTodo, receivedError := testController.Create("foo", "bar")
		assert.NotNil(t, receivedTodo)
		assert.Nil(t, receivedError)
		assert.Equal(t, "foo", receivedTodo.Title)
		assert.Equal(t, "bar", receivedTodo.Content)
		assert.Equal(t, int64(100), receivedTodo.Id)

	})

	t.Run("should return nil todo and error if model returns error", func(t *testing.T) {
		mockedModel := &mocks.TodoModel{}
		// mockedView := &mocks.TodoDisplayer{}
		// mockedTodo := &nil
		err := errors.New("todo creation failed")

		mockedModel.On("CreateTodo", "foo", "bar").Return(nil, err)
		// mockedView.On("ShowTodo", int64(100), "foo", "bar").Return(nil)

		testController := TodoController{
			model: mockedModel,
			// view:  mockedView,
		}

		receivedTodo, receivedError := testController.Create("foo", "bar")
		assert.Nil(t, receivedTodo)
		assert.NotNil(t, receivedError)
	})

	t.Run("should return todo and error if modelworks fine but view return error", func(t *testing.T) {
		mockedModel := &mocks.TodoModel{}
		mockedView := &mocks.TodoDisplayer{}
		mockedTodo := &model.ToDo{
			Id:      10,
			Title:   "foo",
			Content: "bar",
		}
		mockedModel.On("CreateTodo", "foo", "bar").Return(mockedTodo, nil)
		err := errors.New("Can't display todo")
		mockedView.On("ShowTodo", int64(10), "foo", "bar").Return(err)

		testController := TodoController{
			model: mockedModel,
			view:  mockedView,
		}

		receivedTodo, receivedError := testController.Create("foo", "bar")
		assert.NotNil(t, receivedTodo)
		assert.NotNil(t, receivedError)
		assert.Equal(t, "foo", receivedTodo.Title)
		assert.Equal(t, "bar", receivedTodo.Content)
		assert.Equal(t, int64(10), receivedTodo.Id)

	})

}

func TestTodoController_SearchById(t *testing.T) {
	t.Run("should return the todo with given Id when no error in model and view", func(t *testing.T) {
		mockedModel := &mocks.TodoModel{}
		mockedView := &mocks.TodoDisplayer{}
		mockedTodo := &model.ToDo{
			Id:      10,
			Title:   "foo",
			Content: "bar",
		}
		mockedModel.On("GetTodoById", int64(10)).Return(mockedTodo, nil)
		mockedView.On("ShowTodo", int64(10), "foo", "bar").Return(nil)
		testController := TodoController{
			model: mockedModel,
			view:  mockedView,
		}

		receivedTodo, receivedError := testController.SearchById(int64(10))
		assert.NotNil(t, receivedTodo)
		assert.Nil(t, receivedError)
		assert.Equal(t, int64(10), receivedTodo.Id)
		assert.Equal(t, "foo", receivedTodo.Title)
		assert.Equal(t, "bar", receivedTodo.Content)

	})
	t.Run("should return the nil todo when model return error", func(t *testing.T) {
		mockedModel := &mocks.TodoModel{}
		err := errors.New("Todo with given id not present")
		mockedModel.On("GetTodoById", int64(10)).Return(nil, err)
		testController := TodoController{
			model: mockedModel,
		}
		receivedTodo, receivedError := testController.SearchById(int64(10))
		assert.Nil(t, receivedTodo)
		assert.NotNil(t, receivedError)
	})

	t.Run("should return todo and error if GetTodoById of model works fine but view return error", func(t *testing.T) {
		mockedModel := &mocks.TodoModel{}
		mockedView := &mocks.TodoDisplayer{}
		mockedTodo := &model.ToDo{
			Id:      10,
			Title:   "foo",
			Content: "bar",
		}
		mockedModel.On("GetTodoById", int64(10)).Return(mockedTodo, nil)
		err := errors.New("Can't display todo")
		mockedView.On("ShowTodo", int64(10), "foo", "bar").Return(err)

		testController := TodoController{
			model: mockedModel,
			view:  mockedView,
		}

		receivedTodo, receivedError := testController.SearchById(int64(10))
		assert.NotNil(t, receivedTodo)
		assert.NotNil(t, receivedError)
		assert.Equal(t, int64(10), receivedTodo.Id)
		assert.Equal(t, "foo", receivedTodo.Title)
		assert.Equal(t, "bar", receivedTodo.Content)
	})

}

func TestTodoController_DeleteByTitle(t *testing.T) {
	t.Run("should return the todo when no error in model and view", func(t *testing.T) {
		mockedModel := &mocks.TodoModel{}
		mockedView := &mocks.TodoDisplayer{}
		mockedTodo := &model.ToDo{
			Id:      10,
			Title:   "",
			Content: "",
		}
		mockedModel.On("DeleteByTitle", "foo").Return(mockedTodo, nil)
		mockedView.On("ShowTodo", int64(10), "", "").Return(nil)
		testController := TodoController{
			model: mockedModel,
			view:  mockedView,
		}

		receivedTodo, receivedError := testController.DeleteByTitle("foo")
		assert.NotNil(t, receivedTodo)
		assert.Nil(t, receivedError)
		assert.Equal(t, int64(10), receivedTodo.Id)
		assert.Equal(t, "", receivedTodo.Title)
		assert.Equal(t, "", receivedTodo.Content)
	})
	t.Run("should return the nil todo when model return error", func(t *testing.T) {
		mockedModel := &mocks.TodoModel{}
		err := errors.New("Todo with given id not present")
		mockedModel.On("DeleteByTitle", "foo").Return(nil, err)
		testController := TodoController{
			model: mockedModel,
		}
		receivedTodo, receivedError := testController.DeleteByTitle("foo")
		assert.Nil(t, receivedTodo)
		assert.NotNil(t, receivedError)
	})

	t.Run("should return todo and error if DeleteByTitle of model works fine but view return error", func(t *testing.T) {
		mockedModel := &mocks.TodoModel{}
		mockedView := &mocks.TodoDisplayer{}
		mockedTodo := &model.ToDo{
			Id:      10,
			Title:   "",
			Content: "",
		}
		mockedModel.On("DeleteByTitle", "foo").Return(mockedTodo, nil)
		err := errors.New("Can't display todo")
		mockedView.On("ShowTodo", int64(10), "", "").Return(err)

		testController := TodoController{
			model: mockedModel,
			view:  mockedView,
		}

		receivedTodo, receivedError := testController.DeleteByTitle("foo")
		assert.NotNil(t, receivedTodo)
		assert.NotNil(t, receivedError)
		assert.Equal(t, int64(10), receivedTodo.Id)
		assert.Equal(t, "", receivedTodo.Title)
		assert.Equal(t, "", receivedTodo.Content)
	})

}

func TestTodoController_SearchByTitle(t *testing.T) {
	t.Run("should return the todo array when no error in model and view", func(t *testing.T) {
		mockedModel := &mocks.TodoModel{}
		mockedView := &mocks.TodoDisplayer{}
		mockedTodo := &[]model.ToDo{{
			Id:      10,
			Title:   "foo",
			Content: "bar"},
			{
				Id:      11,
				Title:   "foo",
				Content: "baar",
			},
		}
		mockedviewTodo := &[]view.Todo{
			{
				Id:      10,
				Title:   "foo",
				Content: "bar"},
			{
				Id:      11,
				Title:   "foo",
				Content: "baar",
			},
		}
		mockedModel.On("GetTodoByTitle", "foo").Return(mockedTodo, nil)
		mockedView.On("ShowManyTodo", mockedviewTodo).Return(nil)
		testController := TodoController{
			model: mockedModel,
			view:  mockedView,
		}

		receivedTodo, receivedError := testController.SearchByTitle("foo")
		assert.NotNil(t, receivedTodo)
		assert.Nil(t, receivedError)
		// assert.Equal(t, int64(10), len(receivedTodo))
		// assert.Equal(t, "", receivedTodo.Title)
		// assert.Equal(t, "", receivedTodo.Content)
	})
	t.Run("should return the nil todo array and non nil error when GetTodoByTitile gives error", func(t *testing.T) {
		mockedModel := &mocks.TodoModel{}
		err := errors.New("No todo exits with given title")
		mockedModel.On("GetTodoByTitle", "foo").Return(nil, err)
		testController := TodoController{
			model: mockedModel,
		}
		receivedTodo, receivedError := testController.SearchByTitle("foo")
		assert.Nil(t, receivedTodo)
		assert.NotNil(t, receivedError)
	})

	t.Run("should return the todo struct array and on nil errot when no error in model but view return error", func(t *testing.T) {
		mockedModel := &mocks.TodoModel{}
		mockedView := &mocks.TodoDisplayer{}
		mockedTodo := &[]model.ToDo{{
			Id:      10,
			Title:   "foo",
			Content: "bar"},
			{
				Id:      11,
				Title:   "foo",
				Content: "baar",
			},
		}
		mockedviewTodo := &[]view.Todo{
			{
				Id:      10,
				Title:   "foo",
				Content: "bar"},
			{
				Id:      11,
				Title:   "foo",
				Content: "baar",
			},
		}
		mockedModel.On("GetTodoByTitle", "foo").Return(mockedTodo, nil)
		err := errors.New("Can't display todo")
		mockedView.On("ShowManyTodo", mockedviewTodo).Return(err)
		testController := TodoController{
			model: mockedModel,
			view:  mockedView,
		}

		receivedTodo, receivedError := testController.SearchByTitle("foo")
		assert.NotNil(t, receivedTodo)
		assert.NotNil(t, receivedError)
		// assert.Equal(t, int64(10), len(receivedTodo))
		// assert.Equal(t, "", receivedTodo.Title)
		// assert.Equal(t, "", receivedTodo.Content)
	})

}
