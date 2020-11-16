package storage

import (
	"fmt"
	"testing"
)

func TestBuildInsertOrUpdateSql(t *testing.T) {
	columns := make(map[string]interface{})
	columns["device_uuid"] = "rrrrssssaaa"
	columns["user_id"] = 1111222
	columns["name"] = "22333344444"
	columns["location"] = "ST_GeomFromText('POINT(1 22)')"

	sql := BuildInsertOrUpdateSql("device", "device_uuid", columns)

	fmt.Println(sql)
}
