package models

import (
	"api/database"
)

type Timetable struct {
	Id   int     `db:"id" json:"id"`
	Name string  `db:"name" json:"name"`
	Memo *string `db:"memo" json:"memo"`
}

type TimetableData struct {
	Id        int     `db:"id" json:"id"`
	LectureId int     `db:"lecture_id" json:"lecture_id"`
	Name      string  `db:"name" json:"name"`
	Teacher   string  `db:"teacher" json:"teacher"`
	Classroom *string `db:"classroom" json:"classroom"`
	DayOfWeek string  `db:"day_of_week" json:"day_of_week"`
	Period    int     `db:"period" json:"period"`
}

func FetchTimetables() ([]Timetable, error) {
	pool := database.Pool()

	timetables := []Timetable{}
	err := pool.Select(&timetables, "SELECT id, name, memo FROM timetables")
	return timetables, err
}

func FetchTimetable(id int) (Timetable, error) {
	pool := database.Pool()

	timetable := Timetable{}
	err := pool.Get(&timetable, "SELECT id, name, memo FROM timetables WHERE id = ?", id)
	return timetable, err
}

func FetchTimetableData(timetable_id int) ([]TimetableData, error) {
	pool := database.Pool()

	data := []TimetableData{}
	query := `
    SELECT
      tl.id,
      tl.lecture_id,
      l.name,
      l.teacher,
      tl.classroom,
      tl.day_of_week,
      tl.period
    FROM
      timetable_lecture tl
    INNER JOIN lectures l
      ON tl.lecture_id = l.id
    WHERE
      timetable_id = ?
    ORDER BY
      tl.id
    `
	err := pool.Select(&data, query, timetable_id)
	return data, err
}
