package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
)

type Relation struct {
	Id    int64 // Tag the rowId
	Id1   int64 // Tag the userId
	Id2   int64 // Tag the other userId
	State int8  // Tag state enum [like, dislike, match] state by 0 1 2
}

type RetRelationships struct {
	Id    int64  `json:"user_id"`
	State string `json:"state"`
	Type  string `json:"type"`
}

const (
	LIKE    int8 = 0
	DISLIKE int8 = 1
	MATCH   int8 = 2
)

var StateMaps = map[int8]string{
	LIKE:    "like",
	DISLIKE: "dislike",
	MATCH:   "match",
}

var StateRevMaps = map[string]int8{
	"like":    LIKE,
	"dislike": DISLIKE,
	"match":   MATCH,
}

func init() {
	// 需要在init中注册定义的model
	orm.RegisterModel(new(Relation))
}

func Relationships() orm.QuerySeter {
	return orm.NewOrm().QueryTable(new(Relation))
}

// judge userId exits in id1 fields
func CheckUserId(userId int64) bool {
	exist := Relationships().Filter("id1", userId).Exist()
	return exist
}

func GetRelationshipsById(id int64) (ml []RetRelationships, err error) {

	fmt.Println(id)
	if !CheckUserId(id) {
		return nil, err
	}

	var userRelations []RetRelationships
	var relations []Relation

	if _, err := orm.NewOrm().QueryTable("relation").Filter("id1", id).RelatedSel().All(&relations); err == nil {
		for _, v := range relations {
			userRelations = append(userRelations, RetRelationships{v.Id2, StateMaps[v.State], "relationship"})
		}
		return userRelations, err
	}
	return nil, err
}

func SetRelationships(usrID int64, relation RetRelationships) (*Relation, error) {

	o := Relationships()
	stateRevMap := StateRevMaps[relation.Type]
	rel := Relation{
		Id1:   usrID,
		Id2:   relation.Id,
		State: stateRevMap,
	}
	fmt.Println(rel)
	var err error
	if stateRevMap == LIKE {
		r := o.Filter("id1", relation.Id).Filter("id2", usrID)
		if r.Exist() {
			var usr Relation
			r.RelatedSel().One(&usr)
			if usr.State == LIKE {
				//update id2 ==> match
				r.Update(orm.Params{"state": MATCH})

				// update or insert id1 ==> id2 match
				rr := o.Filter("id1", usrID).Filter("id2", relation.Id)
				if exist := rr.Exist(); exist {
					_, err := rr.Update(orm.Params{"state": MATCH})
					if err == nil {
						return &rel, err
					}
				} else {
					rel.State = MATCH
					_, err := orm.NewOrm().Insert(&rel)
					if err == nil {
						return &rel, err
					}
				}

			}

		}

	} else {
		rr := o.Filter("id1", usrID).Filter("id2", relation.Id)
		if exist := rr.Exist(); exist {
			_, err := rr.Update(orm.Params{"state": stateRevMap})
			if err == nil {
				return &rel, err
			}
		} else {
			_, err := orm.NewOrm().Insert(&rel)
			if err == nil {
				return &rel, err
			}
		}
	}

	return nil, err

}
