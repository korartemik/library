package service

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"library/api/model"
	mocks_actions "library/api/service/actions/mocks"
	mocks_storage "library/api/storage/mocks"
	modelapi "library/genproto/model"
	"testing"
)

func TestLibraryServiceMock(t *testing.T) {
	t.Run("get by tittle", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		tester := NewLibraryServiceTester(ctrl)

		tester.con.EXPECT().GetConnection().Return(&sql.DB{}, nil)

		tester.st.EXPECT().GetByTitle(gomock.Any(), "some_title", gomock.Any()).Return([]*model.Book{{Title: "some_title", Author: "some"}}, nil)

		service := NewLibraryService(tester.con, tester.st)

		resp, err := service.GetBooksByTitle(context.Background(), &modelapi.GetByTitleRequest{Title: "some_title"})
		require.NoError(t, err)
		require.NotNil(t, resp)
		require.Equal(t, len(resp.Books), 1)
		assert.Equal(t, resp.Books[0].Author, "some")
	})

	t.Run("not found", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		tester := NewLibraryServiceTester(ctrl)

		tester.con.EXPECT().GetConnection().Return(&sql.DB{}, nil)

		tester.st.EXPECT().GetByTitle(gomock.Any(), "some_title", gomock.Any()).Return(nil, model.NotFound(fmt.Errorf("some error")))

		service := NewLibraryService(tester.con, tester.st)

		resp, err := service.GetBooksByTitle(context.Background(), &modelapi.GetByTitleRequest{Title: "some_title"})
		require.Error(t, err)
		require.Nil(t, resp)
		assert.Equal(t, status.Code(err), codes.NotFound)
	})
}

type libaryServiceTester struct {
	con *mocks_actions.MockConnector
	st  *mocks_storage.MockLibraryStorage
}

func NewLibraryServiceTester(ctrl *gomock.Controller) *libaryServiceTester {
	return &libaryServiceTester{
		con: mocks_actions.NewMockConnector(ctrl),
		st:  mocks_storage.NewMockLibraryStorage(ctrl),
	}
}
