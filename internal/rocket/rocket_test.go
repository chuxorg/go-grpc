package rocket

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
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
}