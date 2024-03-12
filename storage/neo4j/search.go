package neo4j

import (
	"fmt"
	"noname_team_project/model"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func (n *Neo4j) Init() error {
	_, err := neo4j.ExecuteQuery(n.context, n.conn,
		`CREATE (:Lesson {id_lesson: 1, id_lecture: 1})
		CREATE (:Lesson {id_lesson: 2, id_lecture: 2})
		CREATE (:Lesson {id_lesson: 3, id_lecture: 3})
		CREATE (:Lesson {id_lesson: 4, id_lecture: 4})
		CREATE (:Lesson {id_lesson: 5, id_lecture: 5})
		CREATE (:Lesson {id_lesson: 6, id_lecture: 6})
		CREATE (:Lesson {id_lesson: 7, id_lecture: 7})
		CREATE (:Lesson {id_lesson: 8, id_lecture: 8})
		
		CREATE (:Schedule {id_schedule: 1, id_lesson: 1, id_group: 1})
		CREATE (:Schedule {id_schedule: 2, id_lesson: 2, id_group: 1})
		CREATE (:Schedule {id_schedule: 3, id_lesson: 3, id_group: 1})
		CREATE (:Schedule {id_schedule: 4, id_lesson: 4, id_group: 1})
		CREATE (:Schedule {id_schedule: 5, id_lesson: 5, id_group: 1})
		CREATE (:Schedule {id_schedule: 6, id_lesson: 6, id_group: 1})
		CREATE (:Schedule {id_schedule: 7, id_lesson: 7, id_group: 1})
		CREATE (:Schedule {id_schedule: 8, id_lesson: 8, id_group: 1})
		
		CREATE (:Schedule {id_schedule: 9, id_lesson: 1, id_group: 2})
		CREATE (:Schedule {id_schedule: 10, id_lesson: 2, id_group: 2})
		CREATE (:Schedule {id_schedule: 11, id_lesson: 3, id_group: 2})
		CREATE (:Schedule {id_schedule: 12, id_lesson: 4, id_group: 2})
		CREATE (:Schedule {id_schedule: 13, id_lesson: 5, id_group: 2})
		CREATE (:Schedule {id_schedule: 14, id_lesson: 6, id_group: 2})
		CREATE (:Schedule {id_schedule: 15, id_lesson: 7, id_group: 2})
		CREATE (:Schedule {id_schedule: 16, id_lesson: 8, id_group: 2})
		
		CREATE (:Schedule {id_schedule: 17, id_lesson: 1, id_group: 3})
		CREATE (:Schedule {id_schedule: 18, id_lesson: 2, id_group: 3})
		CREATE (:Schedule {id_schedule: 19, id_lesson: 3, id_group: 3})
		CREATE (:Schedule {id_schedule: 20, id_lesson: 4, id_group: 3})
		CREATE (:Schedule {id_schedule: 21, id_lesson: 5, id_group: 3})
		CREATE (:Schedule {id_schedule: 22, id_lesson: 6, id_group: 3})
		CREATE (:Schedule {id_schedule: 23, id_lesson: 7, id_group: 3})
		CREATE (:Schedule {id_schedule: 24, id_lesson: 8, id_group: 3})
		
		CREATE (:Group {id_group: 1})
		CREATE (:Group {id_group: 2})
		CREATE (:Group {id_group: 3})
		
		CREATE (:Student {id_student: 1, group_id: 1})
		CREATE (:Student {id_student: 2, group_id: 1})
		CREATE (:Student {id_student: 3, group_id: 1})
		CREATE (:Student {id_student: 4, group_id: 1})
		CREATE (:Student {id_student: 5, group_id: 1})
		CREATE (:Student {id_student: 6, group_id: 1})
		
		CREATE (:Student {id_student: 7, group_id: 2})
		CREATE (:Student {id_student: 8, group_id: 2})
		CREATE (:Student {id_student: 9, group_id: 2})
		CREATE (:Student {id_student: 10, group_id: 2})
		CREATE (:Student {id_student: 11, group_id: 2})
		CREATE (:Student {id_student: 12, group_id: 2})
		
		CREATE (:Student {id_student: 13, group_id: 3})
		CREATE (:Student {id_student: 14, group_id: 3})
		CREATE (:Student {id_student: 15, group_id: 3})
		CREATE (:Student {id_student: 16, group_id: 3})
		CREATE (:Student {id_student: 17, group_id: 3})
		CREATE (:Student {id_student: 18, group_id: 3})`,
	nil, neo4j.EagerResultTransformer,
	neo4j.ExecuteQueryWithDatabase("neo4j"))
	if err != nil {
		return err
	}

	_, err = neo4j.ExecuteQuery(n.context, n.conn,
		`MATCH (l:Lesson), (s:Schedule)
		WHERE l.id_lesson = s.id_lesson
		CREATE (s)-[:A]->(l)`,
	nil, neo4j.EagerResultTransformer,
	neo4j.ExecuteQueryWithDatabase("neo4j"))
	if err != nil {
		return err
	}

	_, err = neo4j.ExecuteQuery(n.context, n.conn,
		`MATCH (s:Schedule), (g:Group)
		WHERE s.id_group = g.id_group
		CREATE (s)-[:B]->(g)`,
	nil, neo4j.EagerResultTransformer,
	neo4j.ExecuteQueryWithDatabase("neo4j"))
	if err != nil {
		return err
	}

	_, err = neo4j.ExecuteQuery(n.context, n.conn,
		`MATCH (s:Student), (g:Group)
		WHERE s.id_student <= 6
		AND g.id_group = 1
		CREATE (g)-[:C]->(s)`,
	nil, neo4j.EagerResultTransformer,
	neo4j.ExecuteQueryWithDatabase("neo4j"))
	if err != nil {
		return err
	}

	_, err = neo4j.ExecuteQuery(n.context, n.conn,
		`MATCH (s:Student), (g:Group)
		WHERE s.id_student > 6 AND s.id_student <= 12
		AND g.id_group = 2
		CREATE (g)-[:C]->(s)`,
	nil, neo4j.EagerResultTransformer,
	neo4j.ExecuteQueryWithDatabase("neo4j"))
	if err != nil {
		return err
	}

	_, err = neo4j.ExecuteQuery(n.context, n.conn,
		`MATCH (s:Student), (g:Group)
		WHERE s.id_student > 12
		AND g.id_group = 3
		CREATE (g)-[:C]->(s)`,
	nil, neo4j.EagerResultTransformer,
	neo4j.ExecuteQueryWithDatabase("neo4j"))
	if err != nil {
		return err
	}
		
	fmt.Println("init pass")

	return nil
}

func (n *Neo4j) GetVisited(lectures []int) ([]int, []int, error) {
	var lessons, students []int

	var itemsInterface []interface{}
    for _, item := range lectures {
        itemsInterface = append(itemsInterface, item)
    }

	fmt.Println("interface pass")

	result, err := neo4j.ExecuteQuery(n.context, n.conn,
		`MATCH (l:Lesson) 
		WHERE l.id_lecture IN $lectures 
		WITH l.id_lesson AS lessonId 
		MATCH (s:Schedule) 
		WHERE lessonId = s.id_lesson 
		WITH lessonId, s.id_group AS groupsId
		MATCH (g:Group) 
		WHERE groupsId = g.id_group 
		WITH lessonId, g.id_group AS newGroupsId
		MATCH (s:Student)
		WHERE s.group_id = newGroupsId
		RETURN lessonId, s.id_student AS studentsID`,
		map[string]interface{}{
            "lectures": itemsInterface,
        }, neo4j.EagerResultTransformer,
		neo4j.ExecuteQueryWithDatabase("neo4j"))
	if err != nil {
		return nil, nil, err
	}
		
	fmt.Println("search pass")

	for _, record := range result.Records {
		l, _ := record.Get("lessonId")
		lessons = append(lessons, int(l.(int64)))

		s, _ := record.Get("studentsID")
		students = append(students, int(s.(int64)))
	}

	return lessons, students, nil
}
 
func (n *Neo4j) StudentVisit(studentId int) (*model.StudVisit, error) {
	stVisit := model.StudVisit{
		StudId: studentId,
	}

	res, err := neo4j.ExecuteQuery(n.context, n.conn,
		`MATCH (s: Student) WHERE s.id_student = $stud_id WITH s.group_id AS group MATCH (sch: Schedule) WHERE sch.id_group = group WITH sch.id_lesson AS lessons MATCH (l:Lesson) WHERE l.id_lesson = lessons RETURN l.id_lecture AS lectures`,
		map[string]any{
            "stud_id": studentId,
        }, neo4j.EagerResultTransformer,
		neo4j.ExecuteQueryWithDatabase("neo4j"))
	if err != nil {
		return nil, err
	}

	for _, record := range res.Records {
		l, _ := record.Get("lectures")
		stVisit.Lessons = append(stVisit.Lessons, int(l.(int64)))
	}

	return &stVisit, nil
}