// Copyright (C) 2019-2020, Xiongfa Li.
// @author xiongfa.li
// @version V1.0
// Description:

package test

import (
	"github.com/xfali/gobatis"
	"testing"
)

func TestSessionTpl(t *testing.T) {
	gobatis.RegisterTemplateFile("./postgresql.tpl")
	var param = TestTable{Id: 1, Username: "user", Password: "pw"}
	t.Run("select", func(t *testing.T) {
		mgr := gobatis.NewSessionManager(connect())
		sess := mgr.NewSession()
		var ret []TestTable
		sess.Select("selectTestTable").Param(param).Result(&ret)
		t.Log(ret)
	})

	t.Run("insert", func(t *testing.T) {
		mgr := gobatis.NewSessionManager(connect())
		sess := mgr.NewSession()
		var ret int
		err := sess.Insert("insertTestTable").Param(param).Result(&ret)
		t.Log(err)
		t.Log(ret)
	})

	t.Run("insertBatch", func(t *testing.T) {
		mgr := gobatis.NewSessionManager(connect())
		sess := mgr.NewSession()
		var ret int
		err := sess.Insert("insertBatchTestTable").Param([]TestTable{
			{Id: 13, Username: "user13", Password: "pw13"},
			{Id: 14, Username: "user14", Password: "pw14"},
		}).Result(&ret)
		t.Log(err)
		t.Log(ret)
	})

	t.Run("update", func(t *testing.T) {
		mgr := gobatis.NewSessionManager(connect())
		sess := mgr.NewSession()
		var ret int
		err := sess.Update("updateTestTable").Param(TestTable{Id: 1, Username: "user2", Password: "pw2"}).Result(&ret)
		t.Log(err)
		t.Log(ret)
	})

	t.Run("delete", func(t *testing.T) {
		mgr := gobatis.NewSessionManager(connect())
		sess := mgr.NewSession()
		var ret int
		sess.Delete("deleteTestTable").Param(TestTable{Id: 1}).Result(&ret)
		t.Log(ret)
	})
}

func TestSessionXml(t *testing.T) {
	gobatis.RegisterMapperFile("./postgresql.xml")
	var param = TestTable{Id: 1, Username: "user", Password: "pw"}
	t.Run("select", func(t *testing.T) {
		mgr := gobatis.NewSessionManager(connect())
		sess := mgr.NewSession()
		var ret []TestTable
		sess.Select("selectTestTable").Param(param).Result(&ret)
		t.Log(ret)
	})

	t.Run("insert", func(t *testing.T) {
		mgr := gobatis.NewSessionManager(connect())
		sess := mgr.NewSession()
		var ret int
		err := sess.Insert("insertTestTable").Param(param).Result(&ret)
		t.Log(err)
		t.Log(ret)
	})

	t.Run("update", func(t *testing.T) {
		mgr := gobatis.NewSessionManager(connect())
		sess := mgr.NewSession()
		var ret int
		err := sess.Update("updateTestTable").Param(TestTable{Id: 1, Username: "user2", Password: "pw2"}).Result(&ret)
		t.Log(err)
		t.Log(ret)
	})

	t.Run("delete", func(t *testing.T) {
		mgr := gobatis.NewSessionManager(connect())
		sess := mgr.NewSession()
		var ret int
		sess.Delete("deleteTestTable").Param(TestTable{Id: 1}).Result(&ret)
		t.Log(ret)
	})
}
