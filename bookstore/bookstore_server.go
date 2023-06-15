package bookstore

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "path/to/generated/bookstore" // Update the path to the generated code
	// Add any necessary book store dependencies
)

type bookServer struct {
	books []*pb.Book // In-memory book storage for simplicity
}

func (s *bookServer) GetBooks(ctx context.Context, req *pb.GetBooksRequest) (*pb.GetBooksResponse, error) {
	return &pb.GetBooksResponse{
		Books: s.books,
	}, nil
}

func (s *bookServer) GetBook(ctx context.Context, req *pb.GetBookRequest) (*pb.BookResponse, error) {
	bookID := req.GetBookId()

	for _, book := range s.books {
		if book.GetId() == bookID {
			return &pb.BookResponse{
				Book: book,
			}, nil
		}
	}

	return nil, grpc.Errorf(grpc.NotFound, "Book not found")
}

func main() {
	// Create a new gRPC server
	server := grpc.NewServer()

	// Initialize the book server
	bookstore := &bookServer{
		books: []*pb.Book{
			{Id: "1", Title: "Book 1", Author: "Author 1", Price: 10.99, Quantity: 5},
			{Id: "2", Title: "Book 2", Author: "Author 2", Price: 12.99, Quantity: 3},
			{Id: "3", Title: "Book 3", Author: "Author 3", Price: 9.99, Quantity: 8},
		},
	}

	// Register the book server with the gRPC server
	pb.RegisterBookstoreServiceServer(server, bookstore)

	// Start listening on port 50051
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// Serve incoming requests
	if err := server.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
