package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _TenantVirtualObjectDefinitionMgr struct {
	*_BaseMgr
}

// TenantVirtualObjectDefinitionMgr open func
func TenantVirtualObjectDefinitionMgr(db *gorm.DB) *_TenantVirtualObjectDefinitionMgr {
	if db == nil {
		panic(fmt.Errorf("TenantVirtualObjectDefinitionMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_TenantVirtualObjectDefinitionMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__tenant_virtual_object_definition"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_TenantVirtualObjectDefinitionMgr) GetTableName() string {
	return "__tenant_virtual_object_definition"
}

// Reset 重置gorm会话
func (obj *_TenantVirtualObjectDefinitionMgr) Reset() *_TenantVirtualObjectDefinitionMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_TenantVirtualObjectDefinitionMgr) Get() (result TenantVirtualObjectDefinition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualObjectDefinition{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_TenantVirtualObjectDefinitionMgr) Gets() (results []*TenantVirtualObjectDefinition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualObjectDefinition{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_TenantVirtualObjectDefinitionMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(TenantVirtualObjectDefinition{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithObjectType object_type获取
func (obj *_TenantVirtualObjectDefinitionMgr) WithObjectType(objectType int64) Option {
	return optionFunc(func(o *options) { o.query["object_type"] = objectType })
}

// WithObjectName object_name获取
func (obj *_TenantVirtualObjectDefinitionMgr) WithObjectName(objectName string) Option {
	return optionFunc(func(o *options) { o.query["object_name"] = objectName })
}

// WithSchema schema获取
func (obj *_TenantVirtualObjectDefinitionMgr) WithSchema(schema string) Option {
	return optionFunc(func(o *options) { o.query["schema"] = schema })
}

// WithVersion version获取
func (obj *_TenantVirtualObjectDefinitionMgr) WithVersion(version string) Option {
	return optionFunc(func(o *options) { o.query["version"] = version })
}

// WithModel model获取
func (obj *_TenantVirtualObjectDefinitionMgr) WithModel(model string) Option {
	return optionFunc(func(o *options) { o.query["model"] = model })
}

// WithTransform transform获取
func (obj *_TenantVirtualObjectDefinitionMgr) WithTransform(transform string) Option {
	return optionFunc(func(o *options) { o.query["transform"] = transform })
}

// WithDefinition definition获取
func (obj *_TenantVirtualObjectDefinitionMgr) WithDefinition(definition string) Option {
	return optionFunc(func(o *options) { o.query["definition"] = definition })
}

// WithCreateDatabaseWithIfNotExists create_database_with_if_not_exists获取
func (obj *_TenantVirtualObjectDefinitionMgr) WithCreateDatabaseWithIfNotExists(createDatabaseWithIfNotExists string) Option {
	return optionFunc(func(o *options) { o.query["create_database_with_if_not_exists"] = createDatabaseWithIfNotExists })
}

// WithCharacterSetClient character_set_client获取
func (obj *_TenantVirtualObjectDefinitionMgr) WithCharacterSetClient(characterSetClient string) Option {
	return optionFunc(func(o *options) { o.query["character_set_client"] = characterSetClient })
}

// WithCollationConnection collation_connection获取
func (obj *_TenantVirtualObjectDefinitionMgr) WithCollationConnection(collationConnection string) Option {
	return optionFunc(func(o *options) { o.query["collation_connection"] = collationConnection })
}

// WithProcType proc_type获取
func (obj *_TenantVirtualObjectDefinitionMgr) WithProcType(procType int64) Option {
	return optionFunc(func(o *options) { o.query["proc_type"] = procType })
}

// WithCollationDatabase collation_database获取
func (obj *_TenantVirtualObjectDefinitionMgr) WithCollationDatabase(collationDatabase string) Option {
	return optionFunc(func(o *options) { o.query["collation_database"] = collationDatabase })
}

// WithSQLMode sql_mode获取
func (obj *_TenantVirtualObjectDefinitionMgr) WithSQLMode(sqlMode string) Option {
	return optionFunc(func(o *options) { o.query["sql_mode"] = sqlMode })
}

// GetByOption 功能选项模式获取
func (obj *_TenantVirtualObjectDefinitionMgr) GetByOption(opts ...Option) (result TenantVirtualObjectDefinition, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualObjectDefinition{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_TenantVirtualObjectDefinitionMgr) GetByOptions(opts ...Option) (results []*TenantVirtualObjectDefinition, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualObjectDefinition{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_TenantVirtualObjectDefinitionMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]TenantVirtualObjectDefinition, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(TenantVirtualObjectDefinition{}).Where(options.query)
	query.Count(&count)
	resultPage.SetTotal(count)
	if len(page.GetOrederItemsString()) > 0 {
		query = query.Order(page.GetOrederItemsString())
	}
	err = query.Limit(int(page.GetSize())).Offset(int(page.Offset())).Find(&results).Error

	resultPage.SetRecords(results)
	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromObjectType 通过object_type获取内容
func (obj *_TenantVirtualObjectDefinitionMgr) GetFromObjectType(objectType int64) (results []*TenantVirtualObjectDefinition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualObjectDefinition{}).Where("`object_type` = ?", objectType).Find(&results).Error

	return
}

// GetBatchFromObjectType 批量查找
func (obj *_TenantVirtualObjectDefinitionMgr) GetBatchFromObjectType(objectTypes []int64) (results []*TenantVirtualObjectDefinition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualObjectDefinition{}).Where("`object_type` IN (?)", objectTypes).Find(&results).Error

	return
}

// GetFromObjectName 通过object_name获取内容
func (obj *_TenantVirtualObjectDefinitionMgr) GetFromObjectName(objectName string) (results []*TenantVirtualObjectDefinition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualObjectDefinition{}).Where("`object_name` = ?", objectName).Find(&results).Error

	return
}

// GetBatchFromObjectName 批量查找
func (obj *_TenantVirtualObjectDefinitionMgr) GetBatchFromObjectName(objectNames []string) (results []*TenantVirtualObjectDefinition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualObjectDefinition{}).Where("`object_name` IN (?)", objectNames).Find(&results).Error

	return
}

// GetFromSchema 通过schema获取内容
func (obj *_TenantVirtualObjectDefinitionMgr) GetFromSchema(schema string) (results []*TenantVirtualObjectDefinition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualObjectDefinition{}).Where("`schema` = ?", schema).Find(&results).Error

	return
}

// GetBatchFromSchema 批量查找
func (obj *_TenantVirtualObjectDefinitionMgr) GetBatchFromSchema(schemas []string) (results []*TenantVirtualObjectDefinition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualObjectDefinition{}).Where("`schema` IN (?)", schemas).Find(&results).Error

	return
}

// GetFromVersion 通过version获取内容
func (obj *_TenantVirtualObjectDefinitionMgr) GetFromVersion(version string) (results []*TenantVirtualObjectDefinition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualObjectDefinition{}).Where("`version` = ?", version).Find(&results).Error

	return
}

// GetBatchFromVersion 批量查找
func (obj *_TenantVirtualObjectDefinitionMgr) GetBatchFromVersion(versions []string) (results []*TenantVirtualObjectDefinition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualObjectDefinition{}).Where("`version` IN (?)", versions).Find(&results).Error

	return
}

// GetFromModel 通过model获取内容
func (obj *_TenantVirtualObjectDefinitionMgr) GetFromModel(model string) (results []*TenantVirtualObjectDefinition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualObjectDefinition{}).Where("`model` = ?", model).Find(&results).Error

	return
}

// GetBatchFromModel 批量查找
func (obj *_TenantVirtualObjectDefinitionMgr) GetBatchFromModel(models []string) (results []*TenantVirtualObjectDefinition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualObjectDefinition{}).Where("`model` IN (?)", models).Find(&results).Error

	return
}

// GetFromTransform 通过transform获取内容
func (obj *_TenantVirtualObjectDefinitionMgr) GetFromTransform(transform string) (results []*TenantVirtualObjectDefinition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualObjectDefinition{}).Where("`transform` = ?", transform).Find(&results).Error

	return
}

// GetBatchFromTransform 批量查找
func (obj *_TenantVirtualObjectDefinitionMgr) GetBatchFromTransform(transforms []string) (results []*TenantVirtualObjectDefinition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualObjectDefinition{}).Where("`transform` IN (?)", transforms).Find(&results).Error

	return
}

// GetFromDefinition 通过definition获取内容
func (obj *_TenantVirtualObjectDefinitionMgr) GetFromDefinition(definition string) (results []*TenantVirtualObjectDefinition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualObjectDefinition{}).Where("`definition` = ?", definition).Find(&results).Error

	return
}

// GetBatchFromDefinition 批量查找
func (obj *_TenantVirtualObjectDefinitionMgr) GetBatchFromDefinition(definitions []string) (results []*TenantVirtualObjectDefinition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualObjectDefinition{}).Where("`definition` IN (?)", definitions).Find(&results).Error

	return
}

// GetFromCreateDatabaseWithIfNotExists 通过create_database_with_if_not_exists获取内容
func (obj *_TenantVirtualObjectDefinitionMgr) GetFromCreateDatabaseWithIfNotExists(createDatabaseWithIfNotExists string) (results []*TenantVirtualObjectDefinition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualObjectDefinition{}).Where("`create_database_with_if_not_exists` = ?", createDatabaseWithIfNotExists).Find(&results).Error

	return
}

// GetBatchFromCreateDatabaseWithIfNotExists 批量查找
func (obj *_TenantVirtualObjectDefinitionMgr) GetBatchFromCreateDatabaseWithIfNotExists(createDatabaseWithIfNotExistss []string) (results []*TenantVirtualObjectDefinition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualObjectDefinition{}).Where("`create_database_with_if_not_exists` IN (?)", createDatabaseWithIfNotExistss).Find(&results).Error

	return
}

// GetFromCharacterSetClient 通过character_set_client获取内容
func (obj *_TenantVirtualObjectDefinitionMgr) GetFromCharacterSetClient(characterSetClient string) (results []*TenantVirtualObjectDefinition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualObjectDefinition{}).Where("`character_set_client` = ?", characterSetClient).Find(&results).Error

	return
}

// GetBatchFromCharacterSetClient 批量查找
func (obj *_TenantVirtualObjectDefinitionMgr) GetBatchFromCharacterSetClient(characterSetClients []string) (results []*TenantVirtualObjectDefinition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualObjectDefinition{}).Where("`character_set_client` IN (?)", characterSetClients).Find(&results).Error

	return
}

// GetFromCollationConnection 通过collation_connection获取内容
func (obj *_TenantVirtualObjectDefinitionMgr) GetFromCollationConnection(collationConnection string) (results []*TenantVirtualObjectDefinition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualObjectDefinition{}).Where("`collation_connection` = ?", collationConnection).Find(&results).Error

	return
}

// GetBatchFromCollationConnection 批量查找
func (obj *_TenantVirtualObjectDefinitionMgr) GetBatchFromCollationConnection(collationConnections []string) (results []*TenantVirtualObjectDefinition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualObjectDefinition{}).Where("`collation_connection` IN (?)", collationConnections).Find(&results).Error

	return
}

// GetFromProcType 通过proc_type获取内容
func (obj *_TenantVirtualObjectDefinitionMgr) GetFromProcType(procType int64) (results []*TenantVirtualObjectDefinition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualObjectDefinition{}).Where("`proc_type` = ?", procType).Find(&results).Error

	return
}

// GetBatchFromProcType 批量查找
func (obj *_TenantVirtualObjectDefinitionMgr) GetBatchFromProcType(procTypes []int64) (results []*TenantVirtualObjectDefinition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualObjectDefinition{}).Where("`proc_type` IN (?)", procTypes).Find(&results).Error

	return
}

// GetFromCollationDatabase 通过collation_database获取内容
func (obj *_TenantVirtualObjectDefinitionMgr) GetFromCollationDatabase(collationDatabase string) (results []*TenantVirtualObjectDefinition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualObjectDefinition{}).Where("`collation_database` = ?", collationDatabase).Find(&results).Error

	return
}

// GetBatchFromCollationDatabase 批量查找
func (obj *_TenantVirtualObjectDefinitionMgr) GetBatchFromCollationDatabase(collationDatabases []string) (results []*TenantVirtualObjectDefinition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualObjectDefinition{}).Where("`collation_database` IN (?)", collationDatabases).Find(&results).Error

	return
}

// GetFromSQLMode 通过sql_mode获取内容
func (obj *_TenantVirtualObjectDefinitionMgr) GetFromSQLMode(sqlMode string) (results []*TenantVirtualObjectDefinition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualObjectDefinition{}).Where("`sql_mode` = ?", sqlMode).Find(&results).Error

	return
}

// GetBatchFromSQLMode 批量查找
func (obj *_TenantVirtualObjectDefinitionMgr) GetBatchFromSQLMode(sqlModes []string) (results []*TenantVirtualObjectDefinition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualObjectDefinition{}).Where("`sql_mode` IN (?)", sqlModes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_TenantVirtualObjectDefinitionMgr) FetchByPrimaryKey(objectType int64, objectName string, schema string, version string, model string, transform string) (result TenantVirtualObjectDefinition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualObjectDefinition{}).Where("`object_type` = ? AND `object_name` = ? AND `schema` = ? AND `version` = ? AND `model` = ? AND `transform` = ?", objectType, objectName, schema, version, model, transform).First(&result).Error

	return
}
