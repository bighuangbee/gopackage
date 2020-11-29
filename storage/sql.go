package storage

import (
	"gopackage/units"
	"strings"
)

/*
	有唯一键更新插入，没有则插入
	@Param tale 表名
	@Param duplicateKey 唯一键
	@param colums 键值对
*/
func BuildInsertOrUpdateSql(table string, duplicateKey string, columns map[string]interface{})(string){

	var keys string
	var values string
	var valuesChar string
	var updateStr string
	var sql string

	duplicateVal, ok := columns[duplicateKey]
	if duplicateKey != ""{
		keys += "`" + duplicateKey + "`,"
		values += "'" + units.ToStr(duplicateVal) +"' ,"
	}

	for key, val := range columns{
		if duplicateKey != key && strings.ToLower(key) != "id" {
			keys += "`" + key + "`,"

			//包含'的字段值不用’包住
			if strings.Contains(units.ToStr(val), "'"){
				values += units.ToStr(val) + ","
			}else{
				values += "'" + units.ToStr(val) + "',"
			}

			valuesChar += "'" + units.PlaceChar(val) + "',"

			if strings.Contains(units.ToStr(val), "'"){
				updateStr += "`" + key + "`=" + units.ToStr(val) + ","
			}else{
				updateStr += "`" + key + "`='" + units.ToStr(val) + "',"
			}
		}
	}

	sql = `INSERT INTO ` + table +`(`
	sql += strings.TrimRight(keys, ",") + ") VALUES( " + strings.TrimRight(values, ",") + ")"

	if ok{
		sql += " ON DUPLICATE KEY UPDATE " + strings.TrimRight(updateStr, ",")
	}

	return sql
}