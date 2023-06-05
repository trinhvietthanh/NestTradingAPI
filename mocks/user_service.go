package mocks

import (
	"backend/model"
	"context"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

type MockUserService struct {
	mock.Mock
}

func (m *MockUserService) Get(ctx context.Context, uid uuid.UUID) {
	ret := m.Called(ctx, uid)

	var r0 *model.User
	if ret.Get(0) != nil {
		r0 = ret.Get(0)
	}
}