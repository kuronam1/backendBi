package models

import (
	"time"
)

type User struct {
	UserID   int
	Login    string
	Password string
	FullName string
	Role     string
}

type Grade struct {
	GradeID      int
	StudentID    int
	DisciplineID int
	Level        string
	Date         time.Time
	Comment      string
}

type Journal struct {
	Grades  map[string][]Grade
	Headers []string
}

type Schedule struct {
	Lessons map[time.Weekday][]Lesson
}

type Lesson struct {
	LessonId       int
	GroupName      string
	Time           time.Time
	DisciplineName string
	Audience       string
	Description    string
	TeacherName    string
	LessonOrder    int
}

type Group struct {
	Id         int
	Speciality string
	Name       string
	Number     int
	Course     int
}

type Discipline struct {
	DisciplineID   int
	TeacherID      int
	DisciplineName string
	Speciality     string
	Course         int
}

type ParseLessons struct {
	LessonID     int
	GroupID      int
	Time         time.Time
	DisciplineID int
	TeacherID    int
	Audience     string
	Description  string
	LessonOrder  int
}
