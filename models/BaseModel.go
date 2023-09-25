package models

type BaseModel interface {
	TbName() string
	Model() interface{}
	ModelCreate(payload map[string]interface{}) map[string]interface{}
	ModelUpdate(payload map[string]interface{}) map[string]interface{}
}
