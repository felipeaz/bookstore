package service

import (
	"github.com/stretchr/testify/mock"
)

type MockMySQLService struct {
	mock.Mock
}

func (s *MockMySQLService) FetchAll(domainObj interface{}) (interface{}, error) {
	resp := s.Called(domainObj)
	return resp.Get(0), resp.Get(1).(error)
}

func (s *MockMySQLService) FetchAllWithPreload(domainObj interface{}, preload string) (interface{}, error) {
	resp := s.Called(domainObj, preload)
	return resp.Get(0), resp.Get(1).(error)
}

func (s *MockMySQLService) FetchAllWithQueryAndPreload(domainObj interface{}, query, preload, join, group string) (interface{}, error) {
	resp := s.Called(domainObj, query, preload, join, group)
	return resp.Get(0), resp.Get(1).(error)
}

func (s *MockMySQLService) Fetch(domainObj interface{}, id string) (interface{}, error) {
	resp := s.Called(domainObj, id)
	return resp.Get(0), resp.Get(1).(error)
}

func (s *MockMySQLService) FetchWithPreload(domainObj interface{}, id, preload string) (interface{}, error) {
	resp := s.Called(domainObj, id, preload)
	return resp.Get(0), resp.Get(1).(error)
}

func (s *MockMySQLService) FetchAllWhere(domainObj interface{}, fieldName, fieldValue string) (interface{}, error) {
	resp := s.Called(domainObj, fieldName, fieldValue)
	return resp.Get(0), resp.Get(1).(error)
}

func (s *MockMySQLService) Persist(domainObj interface{}) error {
	resp := s.Called(domainObj)
	return resp.Error(0)
}

func (s *MockMySQLService) Refresh(domainObj interface{}, id string) error {
	resp := s.Called(id)
	return resp.Error(0)
}

func (s *MockMySQLService) Remove(domainObj interface{}, id string) error {
	resp := s.Called(domainObj, id)
	return resp.Error(0)
}

func (s *MockMySQLService) RemoveWhere(domainObj interface{}, query string) error {
	resp := s.Called(domainObj, query)
	return resp.Error(0)
}

func (s *MockMySQLService) GetErrorStatusCode(err error) int {
	resp := s.Called(err)
	return resp.Int(0)
}
