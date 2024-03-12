package storage

import (
	"noname_team_project/config"
	"noname_team_project/storage/neo4j"
	"noname_team_project/storage/elastic"
	"noname_team_project/storage/postgre"
)

type Storage struct {
	Postgre *postgre.Postgre
	Elastic *elastic.Elastic
	Neo4j   *neo4j.Neo4j
}

func New(config *config.Config) *Storage {
	return &Storage{
		Elastic: elastic.New(config),
		Postgre: postgre.New(config),
		Neo4j: neo4j.New(config),
	}
}

func (s *Storage) Open() error {
	if err := s.Postgre.Open(); err != nil {
		return err
	}
	if err := s.Elastic.Open(); err != nil {
		return err
	}
	if err := s.Neo4j.Open(); err != nil {
		return err
	}

	return nil
}
