package db

import (
	"fmt"
	"strings"
	"time"

	"github.com/WeBankPartners/wecube-plugins-terraform/terraform-server/common-lib/guid"
	"github.com/WeBankPartners/wecube-plugins-terraform/terraform-server/common/log"
	"github.com/WeBankPartners/wecube-plugins-terraform/terraform-server/models"
)

func InterfaceList(paramsMap map[string]interface{}) (rowData []*models.InterfaceTable, err error) {
	sqlCmd := "SELECT * FROM interface WHERE 1=1"
	paramArgs := []interface{}{}
	for k, v := range paramsMap {
		sqlCmd += " AND " + k + "=?"
		paramArgs = append(paramArgs, v)
	}
	sqlCmd += " ORDER BY id DESC"
	err = x.SQL(sqlCmd, paramArgs...).Find(&rowData)
	if err != nil {
		log.Logger.Error("Get interface list error", log.Error(err))
	}
	return
}

func InterfaceBatchCreate(user string, param []*models.InterfaceTable) (rowData []*models.InterfaceTable, err error) {
	actions := []*execAction{}
	tableName := "interface"
	createTime := time.Now().Format(models.DateTimeFormat)

	for i := range param {
		id := guid.CreateGuid()
		data := &models.InterfaceTable{Id: id, Name: param[i].Name, Plugin: param[i].Plugin, Description: param[i].Description,
			CreateUser: user, CreateTime: createTime, UpdateUser: user, UpdateTime: createTime}
		rowData = append(rowData, data)
	}

	for i := range rowData {
		action, tmpErr := GetInsertTableExecAction(tableName, *rowData[i], nil)
		if tmpErr != nil {
			err = fmt.Errorf("Try to create interface fail,%s ", tmpErr.Error())
			return
		}
		actions = append(actions, action)

		// Auto insert system parameters
		systemParams := make(map[string]map[string][]string)
		systemParams["apply"] = make(map[string][]string)
		systemParams["query"] = make(map[string][]string)
		systemParams["destroy"] = make(map[string][]string)

		systemParams["apply"]["input"] = []string{"id", "asset_id", "region_id", "provider_info"}
		systemParams["query"]["input"] = []string{"id", "asset_id", "region_id", "provider_info"}
		systemParams["destroy"]["input"] = []string{"id", "region_id", "provider_info"}

		systemParams["apply"]["output"] = []string{"id", "asset_id"}
		systemParams["query"]["output"] = []string{"id", "asset_id"}
		systemParams["destroy"]["output"] = []string{"id"}

		// 当 transNullStr 的 key 表示的字段为空时，表示需要将其插入 null
		transNullStr := make(map[string]string)
		transNullStr["template"] = "true"
		transNullStr["object_name"] = "true"
		paramTableName := "parameter"
		for paramType, _ := range systemParams[rowData[i].Name] {
			for _, paramName := range systemParams[rowData[i].Name][paramType] {
				paramId := guid.CreateGuid()
				paramData := &models.ParameterTable{Id: paramId, Name: paramName, Type: paramType, Multiple: "N", Interface: rowData[i].Id, DataType: "string", Source: "system", CreateUser: user, CreateTime: createTime, UpdateTime: createTime, UpdateUser: user, Sensitive: "N", Nullable: "N"}
				action, tmpErr := GetInsertTableExecAction(paramTableName, *paramData, transNullStr)
				if tmpErr != nil {
					err = fmt.Errorf("Try to create parameter fail,%s ", tmpErr.Error())
					return
				}
				actions = append(actions, action)
			}
		}
	}


	err = transaction(actions)
	if err != nil {
		err = fmt.Errorf("Try to create interface fail,%s ", err.Error())
	}
	return
}

func InterfaceBatchDelete(ids []string) (err error) {
	actions := []*execAction{}

	// get the parameter by interface id
	interfaceidsStr := strings.Join(ids, "','")
	sqlCmd := "SELECT * FROM parameter WHERE interface IN ('" + interfaceidsStr + "')" + "ORDER BY object_name DESC"
	var parameterList []*models.ParameterTable
	err = x.SQL(sqlCmd).Find(&parameterList)
	if err != nil {
		log.Logger.Error("Get parameter list error", log.Error(err))
	}
	tableName := "parameter"
	for i := range parameterList {
		action, tmpErr := GetDeleteTableExecAction(tableName, "id", parameterList[i].Id)
		if tmpErr != nil {
			err = fmt.Errorf("Try to delete parameter fail,%s ", tmpErr.Error())
			return
		}
		actions = append(actions, action)
	}

	tableName = "interface"
	for i := range ids {
		action, tmpErr := GetDeleteTableExecAction(tableName, "id", ids[i])
		if tmpErr != nil {
			err = fmt.Errorf("Try to delete interface fail,%s ", tmpErr.Error())
			return
		}
		actions = append(actions, action)
	}
	err = transaction(actions)
	if err != nil {
		err = fmt.Errorf("Try to delete interface fail,%s ", err.Error())
	}
	return
}

func InterfaceBatchUpdate(user string, param []*models.InterfaceTable) (err error) {
	actions := []*execAction{}
	tableName := "interface"
	updateTime := time.Now().Format(models.DateTimeFormat)
	for i := range param {
		param[i].UpdateTime = updateTime
		param[i].UpdateUser = user
		action, tmpErr := GetUpdateTableExecAction(tableName, "id", param[i].Id, *param[i], nil)
		if tmpErr != nil {
			err = fmt.Errorf("Try to update interface fail,%s ", tmpErr.Error())
			return
		}
		actions = append(actions, action)
	}

	err = transaction(actions)
	if err != nil {
		err = fmt.Errorf("Try to update interface fail,%s ", err.Error())
	}
	return
}
