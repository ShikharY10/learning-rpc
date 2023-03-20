package models

import (
	"errors"
	"fmt"
	"strconv"
)

type Student struct {
	Id        int
	FirstName string
	LastName  string
}

func (s *Student) Fullname() string {
	return s.FirstName + " " + s.LastName
}

type College struct {
	database map[int]Student
}

func (c *College) Add(payload Student, reply *Student) error {
	if _, ok := c.database[payload.Id]; ok {
		return errors.New("student with id " + strconv.Itoa(payload.Id) + " already exists")
	}

	c.database[payload.Id] = payload
	*reply = payload
	return nil
}

func (c *College) Get(paylaod int, reply *Student) error {
	result, ok := c.database[paylaod]
	if !ok {
		return fmt.Errorf("student with id '%d' does not exist", paylaod)
	}

	*reply = result
	return nil
}

func NewCollege() *College {
	return &College{
		database: make(map[int]Student),
	}
}
