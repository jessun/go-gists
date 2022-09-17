package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _TenantVirtualShowCreateProcedureMgr struct {
	*_BaseMgr
}

// TenantVirtualShowCreateProcedureMgr open func
func TenantVirtualShowCreateProcedureMgr(db *gorm.DB) *_TenantVirtualShowCreateProcedureMgr {
	if db == nil {
		panic(fmt.Errorf("TenantVirtualShowCreateProcedureMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_TenantVirtualShowCreateProcedureMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__tenant_virtual_show_create_procedure"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_TenantVirtualShowCreateProcedureMgr) GetTableName() string {
	return "__tenant_virtual_show_create_procedure"
}

// Reset 重置gorm会话
func (obj *_TenantVirtualShowCreateProcedureMgr) Reset() *_TenantVirtualShowCreateProcedureMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_TenantVirtualShowCreateProcedureMgr) Get() (result TenantVirtualShowCreateProcedure, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualShowCreateProcedure{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_TenantVirtualShowCreateProcedureMgr) Gets() (results []*TenantVirtualShowCreateProcedure, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualShowCreateProcedure{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_TenantVirtualShowCreateProcedureMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(TenantVirtualShowCreateProcedure{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithRoutineID routine_id获取
func (obj *_TenantVirtualShowCreateProcedureMgr) WithRoutineID(routineID int64) Option {
	return optionFunc(func(o *options) { o.query["routine_id"] = routineID })
}

// WithRoutineName routine_name获取
func (obj *_TenantVirtualShowCreateProcedureMgr) WithRoutineName(routineName string) Option {
	return optionFunc(func(o *options) { o.query["routine_name"] = routineName })
}

// WithCreateRoutine create_routine获取
func (obj *_TenantVirtualShowCreateProcedureMgr) WithCreateRoutine(createRoutine string) Option {
	return optionFunc(func(o *options) { o.query["create_routine"] = createRoutine })
}

// WithProcType proc_type获取
func (obj *_TenantVirtualShowCreateProcedureMgr) WithProcType(procType int64) Option {
	return optionFunc(func(o *options) { o.query["proc_type"] = procType })
}

// WithCharacterSetClient character_set_client获取
func (obj *_TenantVirtualShowCreateProcedureMgr) WithCharacterSetClient(characterSetClient string) Option {
	return optionFunc(func(o *options) { o.query["character_set_client"] = characterSetClient })
}

// WithCollationConnection collation_connection获取
func (obj *_TenantVirtualShowCreateProcedureMgr) WithCollationConnection(collationConnection string) Option {
	return optionFunc(func(o *options) { o.query["collation_connection"] = collationConnection })
}

// WithCollationDatabase collation_database获取
func (obj *_TenantVirtualShowCreateProcedureMgr) WithCollationDatabase(collationDatabase string) Option {
	return optionFunc(func(o *options) { o.query["collation_database"] = collationDatabase })
}

// WithSQLMode sql_mode获取
func (obj *_TenantVirtualShowCreateProcedureMgr) WithSQLMode(sqlMode string) Option {
	return optionFunc(func(o *options) { o.query["sql_mode"] = sqlMode })
}

// GetByOption 功能选项模式获取
func (obj *_TenantVirtualShowCreateProcedureMgr) GetByOption(opts ...Option) (result TenantVirtualShowCreateProcedure, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualShowCreateProcedure{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_TenantVirtualShowCreateProcedureMgr) GetByOptions(opts ...Option) (results []*TenantVirtualShowCreateProcedure, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualShowCreateProcedure{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_TenantVirtualShowCreateProcedureMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]TenantVirtualShowCreateProcedure, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(TenantVirtualShowCreateProcedure{}).Where(options.query)
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

// GetFromRoutineID 通过routine_id获取内容
func (obj *_TenantVirtualShowCreateProcedureMgr) GetFromRoutineID(routineID int64) (result TenantVirtualShowCreateProcedure, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualShowCreateProcedure{}).Where("`routine_id` = ?", routineID).First(&result).Error

	return
}

// GetBatchFromRoutineID 批量查找
func (obj *_TenantVirtualShowCreateProcedureMgr) GetBatchFromRoutineID(routineIDs []int64) (results []*TenantVirtualShowCreateProcedure, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualShowCreateProcedure{}).Where("`routine_id` IN (?)", routineIDs).Find(&results).Error

	return
}

// GetFromRoutineName 通过routine_name获取内容
func (obj *_TenantVirtualShowCreateProcedureMgr) GetFromRoutineName(routineName string) (results []*TenantVirtualShowCreateProcedure, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualShowCreateProcedure{}).Where("`routine_name` = ?", routineName).Find(&results).Error

	return
}

// GetBatchFromRoutineName 批量查找
func (obj *_TenantVirtualShowCreateProcedureMgr) GetBatchFromRoutineName(routineNames []string) (results []*TenantVirtualShowCreateProcedure, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualShowCreateProcedure{}).Where("`routine_name` IN (?)", routineNames).Find(&results).Error

	return
}

// GetFromCreateRoutine 通过create_routine获取内容
func (obj *_TenantVirtualShowCreateProcedureMgr) GetFromCreateRoutine(createRoutine string) (results []*TenantVirtualShowCreateProcedure, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualShowCreateProcedure{}).Where("`create_routine` = ?", createRoutine).Find(&results).Error

	return
}

// GetBatchFromCreateRoutine 批量查找
func (obj *_TenantVirtualShowCreateProcedureMgr) GetBatchFromCreateRoutine(createRoutines []string) (results []*TenantVirtualShowCreateProcedure, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualShowCreateProcedure{}).Where("`create_routine` IN (?)", createRoutines).Find(&results).Error

	return
}

// GetFromProcType 通过proc_type获取内容
func (obj *_TenantVirtualShowCreateProcedureMgr) GetFromProcType(procType int64) (results []*TenantVirtualShowCreateProcedure, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualShowCreateProcedure{}).Where("`proc_type` = ?", procType).Find(&results).Error

	return
}

// GetBatchFromProcType 批量查找
func (obj *_TenantVirtualShowCreateProcedureMgr) GetBatchFromProcType(procTypes []int64) (results []*TenantVirtualShowCreateProcedure, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualShowCreateProcedure{}).Where("`proc_type` IN (?)", procTypes).Find(&results).Error

	return
}

// GetFromCharacterSetClient 通过character_set_client获取内容
func (obj *_TenantVirtualShowCreateProcedureMgr) GetFromCharacterSetClient(characterSetClient string) (results []*TenantVirtualShowCreateProcedure, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualShowCreateProcedure{}).Where("`character_set_client` = ?", characterSetClient).Find(&results).Error

	return
}

// GetBatchFromCharacterSetClient 批量查找
func (obj *_TenantVirtualShowCreateProcedureMgr) GetBatchFromCharacterSetClient(characterSetClients []string) (results []*TenantVirtualShowCreateProcedure, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualShowCreateProcedure{}).Where("`character_set_client` IN (?)", characterSetClients).Find(&results).Error

	return
}

// GetFromCollationConnection 通过collation_connection获取内容
func (obj *_TenantVirtualShowCreateProcedureMgr) GetFromCollationConnection(collationConnection string) (results []*TenantVirtualShowCreateProcedure, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualShowCreateProcedure{}).Where("`collation_connection` = ?", collationConnection).Find(&results).Error

	return
}

// GetBatchFromCollationConnection 批量查找
func (obj *_TenantVirtualShowCreateProcedureMgr) GetBatchFromCollationConnection(collationConnections []string) (results []*TenantVirtualShowCreateProcedure, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualShowCreateProcedure{}).Where("`collation_connection` IN (?)", collationConnections).Find(&results).Error

	return
}

// GetFromCollationDatabase 通过collation_database获取内容
func (obj *_TenantVirtualShowCreateProcedureMgr) GetFromCollationDatabase(collationDatabase string) (results []*TenantVirtualShowCreateProcedure, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualShowCreateProcedure{}).Where("`collation_database` = ?", collationDatabase).Find(&results).Error

	return
}

// GetBatchFromCollationDatabase 批量查找
func (obj *_TenantVirtualShowCreateProcedureMgr) GetBatchFromCollationDatabase(collationDatabases []string) (results []*TenantVirtualShowCreateProcedure, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualShowCreateProcedure{}).Where("`collation_database` IN (?)", collationDatabases).Find(&results).Error

	return
}

// GetFromSQLMode 通过sql_mode获取内容
func (obj *_TenantVirtualShowCreateProcedureMgr) GetFromSQLMode(sqlMode string) (results []*TenantVirtualShowCreateProcedure, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualShowCreateProcedure{}).Where("`sql_mode` = ?", sqlMode).Find(&results).Error

	return
}

// GetBatchFromSQLMode 批量查找
func (obj *_TenantVirtualShowCreateProcedureMgr) GetBatchFromSQLMode(sqlModes []string) (results []*TenantVirtualShowCreateProcedure, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualShowCreateProcedure{}).Where("`sql_mode` IN (?)", sqlModes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_TenantVirtualShowCreateProcedureMgr) FetchByPrimaryKey(routineID int64) (result TenantVirtualShowCreateProcedure, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualShowCreateProcedure{}).Where("`routine_id` = ?", routineID).First(&result).Error

	return
}
