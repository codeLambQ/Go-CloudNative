package main

import (
	"context"
	"fmt"
	"grpc/proto_"
	"net"

	"google.golang.org/grpc"
)

type LessonServiceImpl struct {
	lesson.UnimplementedLessonServiceServer
}

func (l *LessonServiceImpl) GetLesson(ctx context.Context, id *lesson.LessonId) (*lesson.Lesson, error) {
	less := &lesson.Lesson{
		Id:     123,
		Name:   "Go",
		Rating: 10.0,
	}
	return less, nil
}

func main() {
	listener, err := net.Listen("tcp", ":12345")
	if err != nil {
		fmt.Println(err)
		return
	}

	server := grpc.NewServer()
	lesson.RegisterLessonServiceServer(server, &LessonServiceImpl{})
	err = server.Serve(listener)
	if err != nil {
		fmt.Println("启动失败: " + err.Error())
		return
	}
}
