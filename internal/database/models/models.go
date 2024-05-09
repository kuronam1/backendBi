package models

import "time"

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
	Level        int
	Date         time.Time
	Comment      string
}

type Journal struct {
	Grades map[string][]*Grade
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
}

type Group struct {
	Id         int
	Name       string
	Speciality string
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
