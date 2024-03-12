package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"noname_team_project/model"
	"noname_team_project/util"
	"sort"
	"time"
)

type EsRequest struct {
	Phrase string `json:"phrase"`
	Date   time.Time `json:"date"`
}

func (s *Server) handleLab1() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var esReq EsRequest

		err := json.NewDecoder(r.Body).Decode(&esReq)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		fmt.Println("json pass")

		lectureList, err := s.storage.Elastic.GetByPhrase(esReq.Phrase)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		fmt.Println("elastic pass")

		_, studentArray, err := s.storage.Neo4j.GetVisited(lectureList)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		studentArray = util.Unique(studentArray)
		fmt.Println("neo pass")
		fmt.Println(studentArray)
		var student_visit []model.StudVisit

		for _, st := range studentArray {
			stv, err := s.storage.Neo4j.StudentVisit(st)
			if err != nil {
				return
			}
			student_visit = append(student_visit, *stv)
			fmt.Printf("%+v", *stv)
		}

		fmt.Println("stv pass")
		var result []model.Rate

		for _, v := range student_visit {
			visitRate, err := s.storage.Postgre.GetVisitRate(v, esReq.Date)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			result = append(result, *visitRate)
			fmt.Println(visitRate)
		}

		sort.Slice(result, func(i, j int) bool {
			return result[i].Score < result[j].Score
		})
		if len(result) > 10 {
			result = result[:10]
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(result)
	}
}
