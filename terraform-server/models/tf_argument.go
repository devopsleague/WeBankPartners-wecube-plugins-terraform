package models

type TfArgumentTable struct {
	Id                       string `json:"id" xorm:"id"`
	Name                     string `json:"name" xorm:"name"`
	Source                   string `json:"source" xorm:"source"`
	Parameter                string `json:"parameter" xorm:"parameter"`
	DefaultValue             string `json:"defaultValue" xorm:"default_value"`
	IsNull                   string `json:"isNull" xorm:"is_null"`
	Type                     string `json:"type" xorm:"type"`
	ObjectName               string `json:"objectName" xorm:"object_name"`
	IsMulti                  string `json:"isMulti" xorm:"is_multi"`
	ConvertWay               string `json:"convertWay" xorm:"convert_way"`
	RelativeSource           string `json:"relativeSource" xorm:"relative_source"`
	RelativeTfstateAttribute string `json:"relativeTfstateAttribute" xorm:"relative_tfstate_attribute"`
	RelativeParameter        string `json:"relativeParameter" xorm:"relative_parameter"`
	RelativeParameterValue   string `json:"relativeParameterValue" xorm:"relative_parameter_value"`
	FunctionDefine           string `json:"functionDefine" xorm:"function_define"`
	CreateTime               string `json:"createTime" xorm:"create_time"`
	CreateUser               string `json:"createUser" xorm:"create_user"`
	UpdateTime               string `json:"updateTime" xorm:"update_time"`
	UpdateUser               string `json:"updateUser" xorm:"update_user"`
	KeyArgument              string `json:"keyArgument" xorm:"key_argument"`
}

type TfArgumentQuery struct {
	Id                            string `json:"id" xorm:"id"`
	Name                          string `json:"name" xorm:"name"`
	Source                        string `json:"source" xorm:"source"`
	SourceTitle                   string `json:"sourceTitle" xorm:"source_title"`
	Parameter                     string `json:"parameter" xorm:"parameter"`
	ParameterTitle                string `json:"parameterTitle" xorm:"parameter_title"`
	DefaultValue                  string `json:"defaultValue" xorm:"default_value"`
	IsNull                        string `json:"isNull" xorm:"is_null"`
	Type                          string `json:"type" xorm:"type"`
	ObjectName                    string `json:"objectName" xorm:"object_name"`
	ObjectNameTitle               string `json:"objectNameTitle" xorm:"object_name_title"`
	IsMulti                       string `json:"isMulti" xorm:"is_multi"`
	ConvertWay                    string `json:"convertWay" xorm:"convert_way"`
	RelativeSource                string `json:"relativeSource" xorm:"relative_source"`
	RelativeSourceTitle           string `json:"relativeSourceTitle" xorm:"relative_source_title"`
	RelativeTfstateAttribute      string `json:"relativeTfstateAttribute" xorm:"relative_tfstate_attribute"`
	RelativeTfstateAttributeTitle string `json:"relativeTfstateAttributeTitle" xorm:"relative_tfstate_attribute_title"`
	RelativeParameter             string `json:"relativeParameter" xorm:"relative_parameter"`
	RelativeParameterTitle        string `json:"relativeParameterTitle" xorm:"relative_parameter_title"`
	RelativeParameterValue        string `json:"relativeParameterValue" xorm:"relative_parameter_value"`
	FunctionDefine                string `json:"functionDefine" xorm:"function_define"`
	CreateTime                    string `json:"createTime" xorm:"create_time"`
	CreateUser                    string `json:"createUser" xorm:"create_user"`
	UpdateTime                    string `json:"updateTime" xorm:"update_time"`
	UpdateUser                    string `json:"updateUser" xorm:"update_user"`
	KeyArgument                   string `json:"keyArgument" xorm:"key_argument"`
}
