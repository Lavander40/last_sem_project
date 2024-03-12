package model

type Rate struct {
	Student Student `json:"student"`
	Score float64 `jsom:"score"`
}

type Student struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	GroupId string `json:"group"`
}

type Lection struct {
	Title string `json:"title"`
}

type StudVisit struct {
	StudId int `json:"stud_id"`
	Lessons []int `json:"lessons_ids"`
}
