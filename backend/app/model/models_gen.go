// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
	"time"
)

type Connection interface {
	IsConnection()
}

type Edge interface {
	IsEdge()
}

type Node interface {
	IsNode()
}

type CreateTaskInput struct {
	Title     string     `json:"title"`
	Notes     *string    `json:"notes"`
	Completed *bool      `json:"completed"`
	Due       *time.Time `json:"due"`
}

type PageInfo struct {
	EndCursor   string `json:"endCursor"`
	HasNextPage bool   `json:"hasNextPage"`
}

type PaginationInput struct {
	First *int    `json:"first"`
	After *string `json:"after"`
}

type TaskConnection struct {
	PageInfo *PageInfo   `json:"pageInfo"`
	Edges    []*TaskEdge `json:"edges"`
}

func (TaskConnection) IsConnection() {}

type TaskEdge struct {
	Cursor string `json:"cursor"`
	Node   *Task  `json:"node"`
}

func (TaskEdge) IsEdge() {}

type TasksInput struct {
	Completed *bool `json:"completed"`
}

type UpdateTaskInput struct {
	TaskID    string     `json:"taskID"`
	Title     *string    `json:"title"`
	Notes     *string    `json:"notes"`
	Completed *bool      `json:"completed"`
	Due       *time.Time `json:"due"`
}

type TaskOrderFields string

const (
	TaskOrderFieldsLatest TaskOrderFields = "LATEST"
	TaskOrderFieldsDue    TaskOrderFields = "DUE"
)

var AllTaskOrderFields = []TaskOrderFields{
	TaskOrderFieldsLatest,
	TaskOrderFieldsDue,
}

func (e TaskOrderFields) IsValid() bool {
	switch e {
	case TaskOrderFieldsLatest, TaskOrderFieldsDue:
		return true
	}
	return false
}

func (e TaskOrderFields) String() string {
	return string(e)
}

func (e *TaskOrderFields) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = TaskOrderFields(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid TaskOrderFields", str)
	}
	return nil
}

func (e TaskOrderFields) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
