package main

import (
	"context"
	"fmt"
	lesson "grpc/proto_"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("127.0.0.1:12345", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("创建连接失败: " + err.Error())
		return
	}
	client := lesson.NewLessonServiceClient(conn)
	Lesson, err := client.GetLesson(context.Background(), &lesson.LessonId{Id: 1})
	if err != nil {
		fmt.Println("调用方法失败: " + err.Error())
		return
	}
	fmt.Println(Lesson)

}
