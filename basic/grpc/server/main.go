package main

import (
	"context"
	"grpc/proto_"
)

type LessonServiceImpl struct {
	lesson.UnimplementedLessonServiceServer
}

func (l *LessonServiceImpl) GetLesson(ctx context.Context, id *lesson.LessonId) (*lesson.Lesson, error) {
	return nil, nil
}
