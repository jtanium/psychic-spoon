package handlers

import (
	"database/sql"
	"examples/go-echo-vue/models"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type H map[string]interface{}

type TaskHandler struct {
	Db *sql.DB
}

func (th *TaskHandler) GetTasks(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, models.GetTasks(th.Db))
}

func (th *TaskHandler) PutTasks(ctx echo.Context) error {
	var task models.Task
	err := ctx.Bind(&task)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	err = models.CreateTask(th.Db, &task)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, &task)
}

func (th *TaskHandler) DeleteTasks(ctx echo.Context) error {
	id, _ := strconv.Atoi(ctx.Param("id"))
	n, err := models.DeleteTask(th.Db, id)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, map[string]string{"numRowsDeleted": fmt.Sprintf("%d", n)})
}
