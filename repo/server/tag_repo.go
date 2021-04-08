package server

import (
	"CURD/entity"
	"database/sql"
	"strings"
)

type TagRepo struct {
	SqliteRepo
}

func (tag TagRepo) Fetch(comp qcomp,
	scan func(obj interface{}, rows *sql.Rows) (interface{}, error)) ([]entity.Tag, error) {
	makeTag := func() interface{} { return entity.Tag{} }
	entities, err := tag.SqliteRepo.Query(comp, makeTag, scan)
	if nil != err {
		return nil, err
	}

	tags := make([]entity.Tag, len(entities))

	for i := range tags {
		tags[i] = entities[i].(entity.Tag)
	}

	return tags, nil
}

func (repo TagRepo) FetchAll() ([]entity.Tag, error) {
	comp := qcomp{
		Tables: []string {"TAG"},
		Columns: []string {"TAGID", "TITLE", "TAGTYPEID"},
	}

	scan := func (obj interface{}, row *sql.Rows) (interface{}, error) {
		tag := obj.(entity.Tag)
		err := row.Scan(&tag.TagId, &tag.Title, &tag.TagTypeId)
		return tag, err
	}

	return repo.Fetch(comp, scan)
}

func (repo TagRepo) IdOf(title string) (string, error) {
	comp := qcomp{
		Tables: []string {"TAG"},
		Columns: []string {"TAGID"},
		Selection: "TITLE = ?",
		SelectionArgs: []string {title},
	}

	makeId := func() interface{} { return "" }

	scan := func(obj interface{}, row *sql.Rows) (interface{}, error) {
		id := obj.(string)
		err := row.Scan(&id)
		return id, err
	}

	values, err := repo.SqliteRepo.Query(comp, makeId, scan)
	if len(values) > 0 {
		return values[0].(string), err
	}

	return "", err
}

func (repo TagRepo) FetchTaggedServer(ServerId string) ([]entity.Tag, error) {
	comp := qcomp{
		Tables: []string {"SERVER_TAG", "TAG"},
		Columns: []string {"TAG.TAGID", "TAG.TITLE"},
		Selection: strings.Join([]string {"SERVERID = ?",
				"SERVER_TAG.TAGID = TAG.TAGID"}, " AND "),
		SelectionArgs: []string {ServerId},
	}
	return repo.Fetch(comp, scanTag)
}

func (repo TagRepo) FetchUntaggedServer(ServerId string) ([]entity.Tag, error) {
	selection := "TAGID NOT IN (SELECT S.TAGID " +
		"FROM SERVER_TAG AS S " +
		"WHERE S.SERVERID = ?) AND TAGTYPE.DESCRIPTION = ? AND TAG.TAGTYPEID = TAGTYPE.ID"
	comp := qcomp{
		Tables: []string {"TAG", "TAGTYPE"},
		Columns: []string {"TAG.TAGID", "TAG.TITLE"},
		Selection: selection,
		SelectionArgs: []string {ServerId, "SERVER"},
	}
	return repo.Fetch(comp, scanTag)
}

func scanTag(obj interface{}, row *sql.Rows) (interface{}, error) {
	tag := obj.(entity.Tag)
	err := row.Scan(&tag.TagId, &tag.Title)
	return tag, err
}