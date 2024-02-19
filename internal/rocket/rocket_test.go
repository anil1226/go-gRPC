//go:generate mockgen -destination=rocket_mocks_test.go -package=rocket github.com/anil1226/go-gRPC/internal/rocket Store

package rocket

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestRocketService(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	t.Run("test rocket by id", func(t *testing.T) {
		rckStoreMock := NewMockStore(mockCtrl)
		ID := "UUID-1"
		rckStoreMock.EXPECT().GetRocketByID(ID).Return(Rocket{
			ID: ID,
		}, nil)

		rockService := New(rckStoreMock)

		rkt, err := rockService.GetRocketByID(ID)
		assert.NoError(t, err)
		assert.Equal(t, ID, rkt.ID)
	})

	t.Run("test insert rocket", func(t *testing.T) {
		rckStoreMock := NewMockStore(mockCtrl)
		ID := "UUID-1"
		rckStoreMock.EXPECT().InsertRocket(Rocket{ID: ID}).Return(nil)

		rockService := New(rckStoreMock)

		err := rockService.InsertRocket(Rocket{ID: ID})
		assert.NoError(t, err)

	})

	t.Run("test delete rocket", func(t *testing.T) {
		rckStoreMock := NewMockStore(mockCtrl)
		ID := "UUID-1"
		rckStoreMock.EXPECT().DeleteRocket(ID).Return(nil)

		rockService := New(rckStoreMock)

		err := rockService.DeleteRocket(ID)
		assert.NoError(t, err)

	})
}
