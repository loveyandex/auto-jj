package service

import (
	"auto-jj/db"
	"auto-jj/model"
	"auto-jj/mongoCtx"

	"go.mongodb.org/mongo-driver/bson"
)

type JVPostSrv struct {
	CmnSrv *CmnSrv[model.JobPosts]
}

func NewJVPostSrv() *JVPostSrv {
	return &JVPostSrv{CmnSrv: &CmnSrv[model.JobPosts]{Xcol: db.Collection("jv-jobposts")}}
}

func (receiver *JVPostSrv) GetByjobIf(jid int) (*model.JobPosts, error) {
	var j model.JobPosts
	query:=bson.M{"ID":jid}

	err := mongoCtx.GetByQuery(query, receiver.CmnSrv.Xcol, &j)

	return &j,err

}
