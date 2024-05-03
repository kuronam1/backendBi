package models

import "time"

type User struct {
	UserId   int
	Login    string
	Password string
	FullName string
	Role     string
	GroupID  int
}

type Grade struct {
	GradeID      int
	StudentID    int
	DisciplineID int
	Level        int
	Date         time.Time
	Comment      string
}

type Journal struct {
	Grades map[string][]Grade
}

type Schedule struct {
	Lessons []Lesson
}

type Lesson struct {
	LessonId     int
	GroupId      int
	time         time.Time
	disciplineId int
	audience     string
}
