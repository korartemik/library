package service

import (
	"context"
	"google.golang.org/genproto/googleapis/rpc/code"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"library/api/model"
	"library/api/service/actions"
	"library/api/storage"
	modelapi "library/genproto/model"
)

type LibraryService struct {
	st  storage.LibraryStorage
	con actions.Connector
}

func NewLibraryService(con actions.Connector, st storage.LibraryStorage) *LibraryService {
	return &LibraryService{
		st:  st,
		con: con,
	}
}

func (s *LibraryService) GetBooksByTitle(ctx context.Context, req *modelapi.GetByTitleRequest) (*modelapi.GetBooksResponse, error) {
	//TODO: using Reactive planning pattern
	db, err := s.con.GetConnection()
	if err != nil {
		return nil, status.Errorf(codes.Code(code.Code_UNAVAILABLE), err.Error())
	}
	books, err := s.st.GetByTitle(ctx, req.Title, db)

	if err != nil {
		if model.IsNotFound(err) {
			return nil, status.Errorf(codes.Code(code.Code_NOT_FOUND), err.Error())
		}
		return nil, err
	}

	var protoBooks []*modelapi.Book

	for _, book := range books {
		protoBooks = append(protoBooks, modelapi.ConvertToAPI(book))
	}

	return &modelapi.GetBooksResponse{Books: protoBooks}, nil
}

func (s *LibraryService) GetBooksByAuthor(ctx context.Context, req *modelapi.GetByAuthorRequest) (*modelapi.GetBooksResponse, error) {
	db, err := s.con.GetConnection()
	if err != nil {
		return nil, status.Errorf(codes.Code(code.Code_UNAVAILABLE), err.Error())
	}
	books, err := s.st.GetByAuthor(ctx, req.Author, db)

	if err != nil {
		if model.IsNotFound(err) {
			return nil, status.Errorf(codes.Code(code.Code_NOT_FOUND), err.Error())
		}
		return nil, err
	}

	var protoBooks []*modelapi.Book

	for _, book := range books {
		protoBooks = append(protoBooks, modelapi.ConvertToAPI(book))
	}

	return &modelapi.GetBooksResponse{Books: protoBooks}, nil
}
