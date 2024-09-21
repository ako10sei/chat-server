package main

import (
	"context"
	"fmt"
	"log"
	"net"

	desc "github.com/chat-server/pkg/chat_server_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"
)

const grpcPort = 50051

type chatServer struct {
	desc.UnimplementedChatServiceServer
}

// Create - создание нового чата
func (s *chatServer) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	// Просто выводим на экран usernames
	log.Printf("Creating chat with users: %v", req.GetUsernames())

	// В качестве ответа возвращаем фиктивный ID
	return &desc.CreateResponse{
		Id: 1, // Фиксированный ID для примера
	}, nil
}

// Delete - удаление чата
func (s *chatServer) Delete(ctx context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {
	// Просто выводим на экран ID чата
	log.Printf("Deleting chat with ID: %d", req.GetId())

	return &emptypb.Empty{}, nil
}

// SendMessage - отправка сообщения
func (s *chatServer) SendMessage(ctx context.Context, req *desc.SendMessageRequest) (*emptypb.Empty, error) {
	// Просто выводим на экран детали сообщения
	log.Printf("Message from %s: %s at %v", req.GetFrom(), req.GetText(), req.GetTimestamp().AsTime())

	return &emptypb.Empty{}, nil
}

func main() {
	// Создаем TCP listener на указанном порту
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Создаем новый gRPC сервер
	s := grpc.NewServer()
	reflection.Register(s)

	// Регистрируем сервис чатов
	desc.RegisterChatServiceServer(s, &chatServer{})

	log.Printf("Chat server listening at %v", lis.Addr())

	// Запускаем сервер
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
