package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _AllDdlOperationMgr struct {
	*_BaseMgr
}

// AllDdlOperationMgr open func
func AllDdlOperationMgr(db *gorm.DB) *_AllDdlOperationMgr {
	if db == nil {
		panic(fmt.Errorf("AllDdlOperationMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllDdlOperationMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_ddl_operation"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllDdlOperationMgr) GetTableName() string {
	return "__all_ddl_operation"
}

// Reset 重置gorm会话
func (obj *_AllDdlOperationMgr) Reset() *_AllDdlOperationMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllDdlOperationMgr) Get() (result AllDdlOperation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDdlOperation{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllDdlOperationMgr) Gets() (results []*AllDdlOperation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDdlOperation{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllDdlOperationMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllDdlOperation{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithGmtCreate gmt_create获取
func (obj *_AllDdlOperationMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllDdlOperationMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithSchemaVersion schema_version获取
func (obj *_AllDdlOperationMgr) WithSchemaVersion(schemaVersion int64) Option {
	return optionFunc(func(o *options) { o.query["schema_version"] = schemaVersion })
}

// WithTenantID tenant_id获取
func (obj *_AllDdlOperationMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithUserID user_id获取
func (obj *_AllDdlOperationMgr) WithUserID(userID int64) Option {
	return optionFunc(func(o *options) { o.query["user_id"] = userID })
}

// WithDatabaseID database_id获取
func (obj *_AllDdlOperationMgr) WithDatabaseID(databaseID int64) Option {
	return optionFunc(func(o *options) { o.query["database_id"] = databaseID })
}

// WithDatabaseName database_name获取
func (obj *_AllDdlOperationMgr) WithDatabaseName(databaseName string) Option {
	return optionFunc(func(o *options) { o.query["database_name"] = databaseName })
}

// WithTablegroupID tablegroup_id获取
func (obj *_AllDdlOperationMgr) WithTablegroupID(tablegroupID int64) Option {
	return optionFunc(func(o *options) { o.query["tablegroup_id"] = tablegroupID })
}

// WithTableID table_id获取
func (obj *_AllDdlOperationMgr) WithTableID(tableID int64) Option {
	return optionFunc(func(o *options) { o.query["table_id"] = tableID })
}

// WithTableName table_name获取
func (obj *_AllDdlOperationMgr) WithTableName(tableName string) Option {
	return optionFunc(func(o *options) { o.query["table_name"] = tableName })
}

// WithOperationType operation_type获取
func (obj *_AllDdlOperationMgr) WithOperationType(operationType int64) Option {
	return optionFunc(func(o *options) { o.query["operation_type"] = operationType })
}

// WithDdlStmtStr ddl_stmt_str获取
func (obj *_AllDdlOperationMgr) WithDdlStmtStr(ddlStmtStr string) Option {
	return optionFunc(func(o *options) { o.query["ddl_stmt_str"] = ddlStmtStr })
}

// WithExecTenantID exec_tenant_id获取
func (obj *_AllDdlOperationMgr) WithExecTenantID(execTenantID int64) Option {
	return optionFunc(func(o *options) { o.query["exec_tenant_id"] = execTenantID })
}

// GetByOption 功能选项模式获取
func (obj *_AllDdlOperationMgr) GetByOption(opts ...Option) (result AllDdlOperation, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllDdlOperation{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllDdlOperationMgr) GetByOptions(opts ...Option) (results []*AllDdlOperation, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllDdlOperation{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllDdlOperationMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllDdlOperation, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllDdlOperation{}).Where(options.query)
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

// GetFromGmtCreate 通过gmt_create获取内容
func (obj *_AllDdlOperationMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllDdlOperation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDdlOperation{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllDdlOperationMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllDdlOperation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDdlOperation{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllDdlOperationMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllDdlOperation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDdlOperation{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllDdlOperationMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllDdlOperation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDdlOperation{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromSchemaVersion 通过schema_version获取内容
func (obj *_AllDdlOperationMgr) GetFromSchemaVersion(schemaVersion int64) (result AllDdlOperation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDdlOperation{}).Where("`schema_version` = ?", schemaVersion).First(&result).Error

	return
}

// GetBatchFromSchemaVersion 批量查找
func (obj *_AllDdlOperationMgr) GetBatchFromSchemaVersion(schemaVersions []int64) (results []*AllDdlOperation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDdlOperation{}).Where("`schema_version` IN (?)", schemaVersions).Find(&results).Error

	return
}

// GetFromTenantID 通过tenant_id获取内容
func (obj *_AllDdlOperationMgr) GetFromTenantID(tenantID int64) (results []*AllDdlOperation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDdlOperation{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllDdlOperationMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllDdlOperation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDdlOperation{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromUserID 通过user_id获取内容
func (obj *_AllDdlOperationMgr) GetFromUserID(userID int64) (results []*AllDdlOperation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDdlOperation{}).Where("`user_id` = ?", userID).Find(&results).Error

	return
}

// GetBatchFromUserID 批量查找
func (obj *_AllDdlOperationMgr) GetBatchFromUserID(userIDs []int64) (results []*AllDdlOperation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDdlOperation{}).Where("`user_id` IN (?)", userIDs).Find(&results).Error

	return
}

// GetFromDatabaseID 通过database_id获取内容
func (obj *_AllDdlOperationMgr) GetFromDatabaseID(databaseID int64) (results []*AllDdlOperation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDdlOperation{}).Where("`database_id` = ?", databaseID).Find(&results).Error

	return
}

// GetBatchFromDatabaseID 批量查找
func (obj *_AllDdlOperationMgr) GetBatchFromDatabaseID(databaseIDs []int64) (results []*AllDdlOperation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDdlOperation{}).Where("`database_id` IN (?)", databaseIDs).Find(&results).Error

	return
}

// GetFromDatabaseName 通过database_name获取内容
func (obj *_AllDdlOperationMgr) GetFromDatabaseName(databaseName string) (results []*AllDdlOperation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDdlOperation{}).Where("`database_name` = ?", databaseName).Find(&results).Error

	return
}

// GetBatchFromDatabaseName 批量查找
func (obj *_AllDdlOperationMgr) GetBatchFromDatabaseName(databaseNames []string) (results []*AllDdlOperation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDdlOperation{}).Where("`database_name` IN (?)", databaseNames).Find(&results).Error

	return
}

// GetFromTablegroupID 通过tablegroup_id获取内容
func (obj *_AllDdlOperationMgr) GetFromTablegroupID(tablegroupID int64) (results []*AllDdlOperation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDdlOperation{}).Where("`tablegroup_id` = ?", tablegroupID).Find(&results).Error

	return
}

// GetBatchFromTablegroupID 批量查找
func (obj *_AllDdlOperationMgr) GetBatchFromTablegroupID(tablegroupIDs []int64) (results []*AllDdlOperation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDdlOperation{}).Where("`tablegroup_id` IN (?)", tablegroupIDs).Find(&results).Error

	return
}

// GetFromTableID 通过table_id获取内容
func (obj *_AllDdlOperationMgr) GetFromTableID(tableID int64) (results []*AllDdlOperation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDdlOperation{}).Where("`table_id` = ?", tableID).Find(&results).Error

	return
}

// GetBatchFromTableID 批量查找
func (obj *_AllDdlOperationMgr) GetBatchFromTableID(tableIDs []int64) (results []*AllDdlOperation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDdlOperation{}).Where("`table_id` IN (?)", tableIDs).Find(&results).Error

	return
}

// GetFromTableName 通过table_name获取内容
func (obj *_AllDdlOperationMgr) GetFromTableName(tableName string) (results []*AllDdlOperation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDdlOperation{}).Where("`table_name` = ?", tableName).Find(&results).Error

	return
}

// GetBatchFromTableName 批量查找
func (obj *_AllDdlOperationMgr) GetBatchFromTableName(tableNames []string) (results []*AllDdlOperation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDdlOperation{}).Where("`table_name` IN (?)", tableNames).Find(&results).Error

	return
}

// GetFromOperationType 通过operation_type获取内容
func (obj *_AllDdlOperationMgr) GetFromOperationType(operationType int64) (results []*AllDdlOperation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDdlOperation{}).Where("`operation_type` = ?", operationType).Find(&results).Error

	return
}

// GetBatchFromOperationType 批量查找
func (obj *_AllDdlOperationMgr) GetBatchFromOperationType(operationTypes []int64) (results []*AllDdlOperation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDdlOperation{}).Where("`operation_type` IN (?)", operationTypes).Find(&results).Error

	return
}

// GetFromDdlStmtStr 通过ddl_stmt_str获取内容
func (obj *_AllDdlOperationMgr) GetFromDdlStmtStr(ddlStmtStr string) (results []*AllDdlOperation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDdlOperation{}).Where("`ddl_stmt_str` = ?", ddlStmtStr).Find(&results).Error

	return
}

// GetBatchFromDdlStmtStr 批量查找
func (obj *_AllDdlOperationMgr) GetBatchFromDdlStmtStr(ddlStmtStrs []string) (results []*AllDdlOperation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDdlOperation{}).Where("`ddl_stmt_str` IN (?)", ddlStmtStrs).Find(&results).Error

	return
}

// GetFromExecTenantID 通过exec_tenant_id获取内容
func (obj *_AllDdlOperationMgr) GetFromExecTenantID(execTenantID int64) (results []*AllDdlOperation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDdlOperation{}).Where("`exec_tenant_id` = ?", execTenantID).Find(&results).Error

	return
}

// GetBatchFromExecTenantID 批量查找
func (obj *_AllDdlOperationMgr) GetBatchFromExecTenantID(execTenantIDs []int64) (results []*AllDdlOperation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDdlOperation{}).Where("`exec_tenant_id` IN (?)", execTenantIDs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllDdlOperationMgr) FetchByPrimaryKey(schemaVersion int64) (result AllDdlOperation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDdlOperation{}).Where("`schema_version` = ?", schemaVersion).First(&result).Error

	return
}
