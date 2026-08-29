package main

import (
	"context"
	"fmt"
	lesson "grpc/proto_"
	"net"

	http "github.com/go-kratos/kratos/v2/transport/http"

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

type LessonServiceHTTPServerImpl struct{}

func (l *LessonServiceHTTPServerImpl) GetLesson(ctx context.Context, id *lesson.LessonId) (*lesson.Lesson, error) {
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
	go func() {
		err := server.Serve(listener)
		if err != nil {
			fmt.Println("启动失败: " + err.Error())
			return
		}
	}()

	// 添加 http 相关代码
	httpListener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	httpServer := http.NewServer()
	lesson.RegisterLessonServiceHTTPServer(httpServer, &LessonServiceHTTPServerImpl{})
	httpServer.Serve(httpListener)

}
