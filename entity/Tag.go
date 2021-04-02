package entity

import (
	"errors"
	"database/sql"

	"CURD/database"
)

type Tag struct {
	TagId, Title, TagTypeId string
}

func (tag *Tag) InitTagId() {
	comp := tag.makeQueryTagIdComponent()

	if rows, err := database.Query(comp); (nil == err) && rows.Next() {
		defer rows.Close()
		err = rows.Scan(&tag.TagId)

		if nil != err {
			panic (err)
		}	
	}
}

func (tag *Tag) makeQueryTagIdComponent() database.QueryComponent {
	return database.QueryComponent {
		Tables: []string {"TAG"},
		Columns: []string {"TAGID"},
		Selection: "TITLE = ?",
		SelectionArgs: []string {tag.Title},
	}
}

func FetchAllTags() []Tag {
	comp := makeFetchAllComponent()
	return FetchTags(comp)
}

func FetchTagsOfServer(ServerId string) []Tag {
	comp := makeFetchTagsServerComponent(ServerId)
	return FetchTags(comp)
}

func FetchTags(comp qcomp) []Tag {
	rows, err := database.Query(comp)
	defer rows.Close()

	if nil != err {
		panic (err)
	}

	tags := make([]Tag, 0)
	tag := &Tag{}

	for rows.Next() && nil != tag {
		tag = ParseTag(rows)
		tags = append(tags, *tag)
	}

	return tags
}

func FetchTagsByTagType(TagType string) []Tag {
	comp := makeFetchTagByTagTypeComp(TagType)
	return fetchTagByComp(comp)
}

func fetchTagByComp(comp qcomp) []Tag {
	tags := make([]Tag, 0)
	return tags
}

func makeFetchTagByTagTypeComp(TagType string) qcomp {
	return makeFetchTagComp (
		"TagType = ? AND Tag.TagTypeId = TagType.Id",
				[]string {
					TagType,
				})
}

func makeFetchTagComp(selection string, selectionArgs []string) qcomp {
	return qcomp {}
}

func fetchTagsByTypeComp(TagType string) qcomp {
	return qcomp {
		Tables: []string {"TAG", "TAGTYPE"},
		Columns: []string {"ID", },
	}
}

func makeFetchAllComponent() qcomp {
	return qcomp {
		Tables: []string {"TAG"},
		Columns: []string {"TagId", "Title"},
	}
}

func makeFetchTagsServerComponent(ServerId string) qcomp {
	return qcomp {
		Tables: []string {"TAG"},
		Columns: []string {"TagId", "Title"},
		Selection: "ServerId = ?",
		SelectionArgs: []string {ServerId},
	}
}

func ParseTag(row *sql.Rows) *Tag {
	var tag Tag

	err := row.Scan(&tag.TagId, &tag.Title)
	if nil != err {
		return nil
	}

	return &tag
}

func DeleteServerTags(ServerId string) {
	comp := makeDeleteServerTagsComponent(ServerId)

	err := database.Delete(comp)
	if nil != err {
		panic (err)
	}
}

func makeDeleteServerTagsComponent(ServerId string) database.DeleteComponent {
	return database.DeleteComponent {
		Table: "SERVER_TAG",
		Selection: "SERVERID = ?",
		SelectionArgs: []string {ServerId},
	}
}

func InsertServerTags(ServerId string, tags []Tag) error {
	if 0 != len(tags) {
		comp := makeInsertServerTagsComponent(ServerId, tags)
		return database.Insert(comp)
	}

	return errors.New("Tag array is null")
}

func makeInsertServerTagsComponent(ServerId string, tags []Tag) database.InsertComponent {
	return database.InsertComponent {
		Table: "SERVER_TAG",
		Columns: []string {"SERVERID", "TAGID"},
		Values: makeValuesInsertStatement(ServerId, tags),
	}
}

func makeValuesInsertStatement(ServerId string, tags []Tag) [][]string {
	values := make([][]string, len(tags))

	for i := range values {
		values[i] = []string {ServerId, tags[i].TagId}
	}

	return values
}
