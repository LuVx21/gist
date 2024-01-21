package rocksdb

import (
    "github.com/linxGnu/grocksdb"
    "luvx/api/logs"
    "strconv"
    "testing"
)

func Test_a(t *testing.T) {
    db, _ := OpenDB()

    readOptions := grocksdb.NewDefaultReadOptions()
    readOptions.SetFillCache(true)

    writeOptions := grocksdb.NewDefaultWriteOptions()
    writeOptions.SetSync(true)

    for i := 0; i < 100; i++ {
        keyStr := "key_" + strconv.Itoa(i)
        valueStr := "value_" + strconv.Itoa(i)
        var key = []byte(keyStr)
        var value = []byte(valueStr)

        db.Put(writeOptions, key, value)
        slice, err2 := db.Get(readOptions, key)
        if err2 != nil {
            logs.Log.Println("获取数据异常:", key, err2)
            continue
        }
        logs.Log.Println("获取数据-> size:", slice.Size(), "值:", string(slice.Data()))

        //err2 = db.Delete(writeOptions, key)
    }

    //defer readOptions.Destroy()
    //defer writeOptions.Destroy()
}

func Test_Iterator(t *testing.T) {
    db, _ := OpenDB()
    get("key_0", db)

    all(db)
}
