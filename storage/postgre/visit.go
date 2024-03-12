package postgre

import (
	"noname_team_project/model"
	"time"
	"github.com/lib/pq"
)

func (p *Postgre) GetVisitRate(stud model.StudVisit, date time.Time) (*model.Rate, error) {
	var rate model.Rate

	rate.Student, _ = p.GetStudent(stud.StudId)
	stud.Lessons, _ = p.FilterLessonByDate(stud.Lessons, date)
	
	err := p.conn.QueryRow("SELECT ((SELECT count(id) FROM attendances WHERE stud_id = $1 ) / $2::float) AS rate", stud.StudId, len(stud.Lessons)).Scan(&rate.Score)
	if err != nil{
		return nil, err
	}

	return &rate, nil
}

func (p *Postgre) GetStudent(id int) (model.Student, error) {
	var student model.Student
	
	err := p.conn.QueryRow("SELECT * FROM students WHERE id = $1", id).Scan(&student.Id, &student.Name, &student.GroupId)
	if err != nil{
		return student, err
	}

	return student, nil
}

func (p *Postgre) FilterLessonByDate(lessonsArray []int, date time.Time) ([]int, error) {
	var filteredArray []int

	rows, err := p.conn.Query("SELECT id FROM lessons WHERE id = ANY($1) AND date >= $2;", pq.Array(lessonsArray), pq.FormatTimestamp(date))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next(){
		var id int
        err := rows.Scan(&id)
        if err != nil{
            return nil, err
        }
        filteredArray = append(filteredArray, id)
    }

	return filteredArray, nil
}
