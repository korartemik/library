package client

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"library/genproto"
	"library/genproto/model"
)

type LibraryClient struct {
	client genproto.LibraryServiceClient
	conn   *grpc.ClientConn
}

func NewLibraryClient(addr *string) (*LibraryClient, error) {
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	c := genproto.NewLibraryServiceClient(conn)
	return &LibraryClient{
		client: c,
		conn:   conn,
	}, nil
}

func (c *LibraryClient) Close() {
	c.conn.Close()
}

func (c *LibraryClient) GetBooksByAuthor(ctx context.Context, author string) (*model.GetBooksResponse, error) {
	return c.client.GetBooksByAuthor(ctx, &model.GetByAuthorRequest{Author: author})
}

func (c *LibraryClient) GetBooksByTitle(ctx context.Context, title string) (*model.GetBooksResponse, error) {
	return c.client.GetBooksByTitle(ctx, &model.GetByTitleRequest{Title: title})
}
