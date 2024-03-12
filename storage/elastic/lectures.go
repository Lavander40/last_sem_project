package elastic

import (
	"bytes"
	"encoding/json"
	"io"
	"noname_team_project/model"
	"strconv"
	"strings"
)

type EsResponse struct{
	Hits struct {
		Hits     []struct {
			ID     string  `json:"_id"`
		} `json:"hits"`
	} `json:"hits"`
}

func (e *Elastic) Init() error {
	jlec, _ := json.Marshal(model.Lection{Title: "Основы квантовой механики"})
	_, err := e.conn.Index("lectures", bytes.NewBuffer(jlec), e.conn.Index.WithDocumentID("1"))
	if err != nil {
		return err
	}
	
	jlec, _ = json.Marshal(model.Lection{Title: "Основы проектирования распределённых систем"})
	_, err = e.conn.Index("lectures", bytes.NewBuffer(jlec), e.conn.Index.WithDocumentID("2"))
	if err != nil {
		return err
	}
	
	jlec, _ = json.Marshal(model.Lection{Title: "Эксплуатация распределённых систем"})
	_, err = e.conn.Index("lectures", bytes.NewBuffer(jlec), e.conn.Index.WithDocumentID("3"))
	if err != nil {
		return err
	}

	return nil
}

func (e *Elastic) GetByPhrase(phrase string) ([]int, error) {
	var lectures EsResponse
	var lectureIds []int

	res, err := e.conn.Search(
		e.conn.Search.WithBody(strings.NewReader(`{
		  "query": {
			"query_string": {
				"query": "title:\"` + phrase + `\""
			}
		  }
		}`)),
		e.conn.Search.WithPretty(),
	)
	if err != nil {
		return nil, err
	}

	temp, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(temp, &lectures); err != nil {
		return nil, err
	}

	for _, v := range lectures.Hits.Hits {
		id, err := strconv.Atoi(v.ID)
		if err != nil{
			return nil, err
		}

		lectureIds = append(lectureIds, id)
	}

	return lectureIds, nil
}
