package rocket

import (
	"context"
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestRocketService(t *testing.T){
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	t.Run("Test get Rocket Id", func(t *testing.T) {
		rocketStoreMock := NewMockStore(mockCtrl)
		id := "UUID-1"
		rocketStoreMock.
			EXPECT().
			GetRocketById(id).
			Return(Rocket {
				ID:id,
			}, nil)
		
		rocketService := New(rocketStoreMock)
		rocket, err := rocketService.GetRocketById(
			context.Background(),
			id,
		)

		assert.NoError(t, err)
		assert.Equal(t, "UUID-1", rocket.ID)
	})

	t.Run("Test insert a Rocket", func(t *testing.T) {
		rocketStoreMock := NewMockStore(mockCtrl)
		id := "UUID-1"
		rocketStoreMock.
			EXPECT().
			InsertRocket(Rocket{
				ID: id,
			}).
			Return(Rocket {
				ID:id,
			}, nil)
		
		rocketService := New(rocketStoreMock)
		rocket, err := rocketService.InsertRocket(
			context.Background(),
			Rocket {
				ID:      id,
			},
		) 

		assert.NoError(t, err)
		assert.Equal(t, "UUID-1", rocket.ID)
	})

	t.Run("Test Deleting a Rocket", func(t *testing.T) {
		rocketStoreMock := NewMockStore(mockCtrl)
		id := "UUID-1"
		rocketStoreMock.
			EXPECT().
			DeleteRocket(id). 
			Return(nil)
		 
		rocketService := New(rocketStoreMock)
		err := rocketService.DeleteRocket(id)
		assert.NoError(t, err)
		
	})
}