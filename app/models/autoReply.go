package models

import (
	"encoding/json"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

type AutoReply struct {
	Id      bson.ObjectId `bson:"_id,omitempty"`
	MsgType string        `bson:"msgType"`
	KeyWord string        `bson:"keyWord"`
	Content string        `bson:"content"`
}

func getAutoReplyCollection(s *mgo.Session) *mgo.Collection {
	return s.DB("weixin").C("autoReply")
}

func SaveTextAutoReply(s *mgo.Session, dataString string) error {
	autoReply := &AutoReply{}
	jsonErr := json.Unmarshal([]byte(dataString), autoReply)
	if jsonErr != nil {
		panic(jsonErr)
		return jsonErr
	}
	err := getAutoReplyCollection(s).Insert(autoReply)
	if err != nil {
		panic(err)
		return err
	}
	return nil
}

func GetTextAutoReply(s *mgo.Session, keyWord string) (autoReply *AutoReply, err error) {
	err = getAutoReplyCollection(s).Find(bson.M{"keyWord": keyWord}).One(&autoReply)
	return
}
