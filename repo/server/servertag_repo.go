package server

import (
	"CURD/entity"
	"database/sql"
	"strings"
)

type ServerTagRepo struct {
	SqliteRepo
}

func (s ServerTagRepo) Insert(ServerId string, tags ...entity.Tag) error {
	comp := s.MakeInsComp(ServerId, tags...)
	return s.SqliteRepo.Insert(comp)
}

func (s ServerTagRepo) MakeInsComp(ServerId string, tags ...entity.Tag) icomp {
	return icomp{
		Table: "SERVER_TAG",
		Columns: []string {"SERVERID", "TAGID"},
		Values: s.makeValues(ServerId ,tags...),
	}
}

func (repo ServerTagRepo)  makeValues(ServerId string, tags ...entity.Tag) [][]string {
	values := make([][]string, len(tags))
	for i := range tags {
		values = append(values,
			[]string{ServerId,tags[i].TagId})
	}
	return values
}

func (repo ServerTagRepo) Delte(comp dcomp) error {
	return repo.SqliteRepo.Delete(comp)
}

func (repo ServerTagRepo) MakeDelteComp(ServerId string) dcomp {
	return dcomp{
		Table: "SERVER_TAG",
		Selection: "SERVERID = ?",
		SelectionArgs: []string {ServerId},
	}
}

func (s ServerTagRepo) FetchTags(serverId string) ([]entity.Tag, error) {
	comp := qcomp{
		Tables: []string {"TAG", "SERVER_TAG", "TAGTYPE"},
		Columns: []string {"TAG.ID", "TAG.TITLE", "TAGTYPE.ID"},
		Selection: strings.Join(
			[]string {"SERVER_TAG.SERVERID = ?",
				"SERVER_TAG.TAGID = ?",
				"TAG.TAGTYPEID = ?"},
			" AND "),

		SelectionArgs: []string {serverId, "TAG.TAGID", "TAGTYPE.ID"},
	}

	scanTag := func (obj interface{}, rows *sql.Rows) (interface{}, error) {
		tag := obj.(entity.Tag)
		err := rows.Scan(&tag.TagId,
			&tag.Title,
			&tag.TagTypeId)
		return tag, err
	}

	return TagRepo{}.Fetch(comp, scanTag)
}
