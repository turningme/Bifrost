package src_test

import (
	"github.com/brokercap/Bifrost/sdk/pluginTestData"
	"testing"
)

func TestAppendSyncAndChekcData(t *testing.T){
	beforeTest()
	initDBTable(true)
	conn := getPluginConn("LogAppend")
	e := pluginTestData.NewEvent()
	insertdata := e.GetTestInsertData()
	conn.Insert(insertdata)
	updateData := e.GetTestUpdateData()
	conn.Update(updateData)
	deleteData := e.GetTestDeleteData()
	conn.Del(deleteData)
	_,err2 := conn.Commit()
	if err2 != nil{
		t.Fatal(err2)
	}

	n , err := getTableCount()
	if err != nil{
		t.Fatal(err)
	}

	if n != 3{
		t.Fatal("append result count != 3")
	}
	t.Log("success")
}

