package genproto

import (
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"library/api/model"
	"library/api/service"
	"library/api/service/actions"
	"library/api/storage/library"
	modelapi "library/genproto/model"
	"log"
	"net"
)

func StartServer(options model.ServerOptions) (*grpc.Server, error) {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *options.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	RegisterLibraryServiceServer(s, NewLibraryServer(options.DBOptions))
	log.Printf("server listening at %v", lis.Addr())
	go s.Serve(lis)

	return s, nil
}

// libraryServer is used to implement genproto.LibraryServiceServer.
type libraryServer struct {
	service *service.LibraryService
}

func NewLibraryServer(options model.DBOptions) *libraryServer {
	return &libraryServer{
		service: service.NewLibraryService(actions.NewStorageConnector(options), library.NewLibraryStorage()),
	}
}

// GetBooksByTitle implements genproto.LibraryServiceServer
func (s *libraryServer) GetBooksByTitle(ctx context.Context, req *modelapi.GetByTitleRequest) (*modelapi.GetBooksResponse, error) {
	return s.service.GetBooksByTitle(ctx, req)
}

// GetBooksByAuthor implements genproto.LibraryServiceServer
func (s *libraryServer) GetBooksByAuthor(ctx context.Context, req *modelapi.GetByAuthorRequest) (*modelapi.GetBooksResponse, error) {
	return s.service.GetBooksByAuthor(ctx, req)
}

func (s *libraryServer) mustEmbedUnimplementedLibraryServiceServer() {}
