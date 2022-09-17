package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _AllVirtualDdlOperationMgr struct {
	*_BaseMgr
}

// AllVirtualDdlOperationMgr open func
func AllVirtualDdlOperationMgr(db *gorm.DB) *_AllVirtualDdlOperationMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualDdlOperationMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualDdlOperationMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_ddl_operation"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualDdlOperationMgr) GetTableName() string {
	return "__all_virtual_ddl_operation"
}

// Reset 重置gorm会话
func (obj *_AllVirtualDdlOperationMgr) Reset() *_AllVirtualDdlOperationMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualDdlOperationMgr) Get() (result AllVirtualDdlOperation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDdlOperation{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualDdlOperationMgr) Gets() (results []*AllVirtualDdlOperation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDdlOperation{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualDdlOperationMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualDdlOperation{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTenantID tenant_id获取
func (obj *_AllVirtualDdlOperationMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithSchemaVersion schema_version获取
func (obj *_AllVirtualDdlOperationMgr) WithSchemaVersion(schemaVersion int64) Option {
	return optionFunc(func(o *options) { o.query["schema_version"] = schemaVersion })
}

// WithGmtCreate gmt_create获取
func (obj *_AllVirtualDdlOperationMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllVirtualDdlOperationMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithUserID user_id获取
func (obj *_AllVirtualDdlOperationMgr) WithUserID(userID int64) Option {
	return optionFunc(func(o *options) { o.query["user_id"] = userID })
}

// WithDatabaseID database_id获取
func (obj *_AllVirtualDdlOperationMgr) WithDatabaseID(databaseID int64) Option {
	return optionFunc(func(o *options) { o.query["database_id"] = databaseID })
}

// WithDatabaseName database_name获取
func (obj *_AllVirtualDdlOperationMgr) WithDatabaseName(databaseName string) Option {
	return optionFunc(func(o *options) { o.query["database_name"] = databaseName })
}

// WithTablegroupID tablegroup_id获取
func (obj *_AllVirtualDdlOperationMgr) WithTablegroupID(tablegroupID int64) Option {
	return optionFunc(func(o *options) { o.query["tablegroup_id"] = tablegroupID })
}

// WithTableID table_id获取
func (obj *_AllVirtualDdlOperationMgr) WithTableID(tableID int64) Option {
	return optionFunc(func(o *options) { o.query["table_id"] = tableID })
}

// WithTableName table_name获取
func (obj *_AllVirtualDdlOperationMgr) WithTableName(tableName string) Option {
	return optionFunc(func(o *options) { o.query["table_name"] = tableName })
}

// WithOperationType operation_type获取
func (obj *_AllVirtualDdlOperationMgr) WithOperationType(operationType int64) Option {
	return optionFunc(func(o *options) { o.query["operation_type"] = operationType })
}

// WithDdlStmtStr ddl_stmt_str获取
func (obj *_AllVirtualDdlOperationMgr) WithDdlStmtStr(ddlStmtStr string) Option {
	return optionFunc(func(o *options) { o.query["ddl_stmt_str"] = ddlStmtStr })
}

// WithExecTenantID exec_tenant_id获取
func (obj *_AllVirtualDdlOperationMgr) WithExecTenantID(execTenantID int64) Option {
	return optionFunc(func(o *options) { o.query["exec_tenant_id"] = execTenantID })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualDdlOperationMgr) GetByOption(opts ...Option) (result AllVirtualDdlOperation, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDdlOperation{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualDdlOperationMgr) GetByOptions(opts ...Option) (results []*AllVirtualDdlOperation, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDdlOperation{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualDdlOperationMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualDdlOperation, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualDdlOperation{}).Where(options.query)
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

// GetFromTenantID 通过tenant_id获取内容
func (obj *_AllVirtualDdlOperationMgr) GetFromTenantID(tenantID int64) (results []*AllVirtualDdlOperation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDdlOperation{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllVirtualDdlOperationMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllVirtualDdlOperation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDdlOperation{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromSchemaVersion 通过schema_version获取内容
func (obj *_AllVirtualDdlOperationMgr) GetFromSchemaVersion(schemaVersion int64) (results []*AllVirtualDdlOperation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDdlOperation{}).Where("`schema_version` = ?", schemaVersion).Find(&results).Error

	return
}

// GetBatchFromSchemaVersion 批量查找
func (obj *_AllVirtualDdlOperationMgr) GetBatchFromSchemaVersion(schemaVersions []int64) (results []*AllVirtualDdlOperation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDdlOperation{}).Where("`schema_version` IN (?)", schemaVersions).Find(&results).Error

	return
}

// GetFromGmtCreate 通过gmt_create获取内容
func (obj *_AllVirtualDdlOperationMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllVirtualDdlOperation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDdlOperation{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllVirtualDdlOperationMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllVirtualDdlOperation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDdlOperation{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllVirtualDdlOperationMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllVirtualDdlOperation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDdlOperation{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllVirtualDdlOperationMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllVirtualDdlOperation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDdlOperation{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromUserID 通过user_id获取内容
func (obj *_AllVirtualDdlOperationMgr) GetFromUserID(userID int64) (results []*AllVirtualDdlOperation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDdlOperation{}).Where("`user_id` = ?", userID).Find(&results).Error

	return
}

// GetBatchFromUserID 批量查找
func (obj *_AllVirtualDdlOperationMgr) GetBatchFromUserID(userIDs []int64) (results []*AllVirtualDdlOperation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDdlOperation{}).Where("`user_id` IN (?)", userIDs).Find(&results).Error

	return
}

// GetFromDatabaseID 通过database_id获取内容
func (obj *_AllVirtualDdlOperationMgr) GetFromDatabaseID(databaseID int64) (results []*AllVirtualDdlOperation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDdlOperation{}).Where("`database_id` = ?", databaseID).Find(&results).Error

	return
}

// GetBatchFromDatabaseID 批量查找
func (obj *_AllVirtualDdlOperationMgr) GetBatchFromDatabaseID(databaseIDs []int64) (results []*AllVirtualDdlOperation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDdlOperation{}).Where("`database_id` IN (?)", databaseIDs).Find(&results).Error

	return
}

// GetFromDatabaseName 通过database_name获取内容
func (obj *_AllVirtualDdlOperationMgr) GetFromDatabaseName(databaseName string) (results []*AllVirtualDdlOperation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDdlOperation{}).Where("`database_name` = ?", databaseName).Find(&results).Error

	return
}

// GetBatchFromDatabaseName 批量查找
func (obj *_AllVirtualDdlOperationMgr) GetBatchFromDatabaseName(databaseNames []string) (results []*AllVirtualDdlOperation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDdlOperation{}).Where("`database_name` IN (?)", databaseNames).Find(&results).Error

	return
}

// GetFromTablegroupID 通过tablegroup_id获取内容
func (obj *_AllVirtualDdlOperationMgr) GetFromTablegroupID(tablegroupID int64) (results []*AllVirtualDdlOperation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDdlOperation{}).Where("`tablegroup_id` = ?", tablegroupID).Find(&results).Error

	return
}

// GetBatchFromTablegroupID 批量查找
func (obj *_AllVirtualDdlOperationMgr) GetBatchFromTablegroupID(tablegroupIDs []int64) (results []*AllVirtualDdlOperation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDdlOperation{}).Where("`tablegroup_id` IN (?)", tablegroupIDs).Find(&results).Error

	return
}

// GetFromTableID 通过table_id获取内容
func (obj *_AllVirtualDdlOperationMgr) GetFromTableID(tableID int64) (results []*AllVirtualDdlOperation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDdlOperation{}).Where("`table_id` = ?", tableID).Find(&results).Error

	return
}

// GetBatchFromTableID 批量查找
func (obj *_AllVirtualDdlOperationMgr) GetBatchFromTableID(tableIDs []int64) (results []*AllVirtualDdlOperation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDdlOperation{}).Where("`table_id` IN (?)", tableIDs).Find(&results).Error

	return
}

// GetFromTableName 通过table_name获取内容
func (obj *_AllVirtualDdlOperationMgr) GetFromTableName(tableName string) (results []*AllVirtualDdlOperation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDdlOperation{}).Where("`table_name` = ?", tableName).Find(&results).Error

	return
}

// GetBatchFromTableName 批量查找
func (obj *_AllVirtualDdlOperationMgr) GetBatchFromTableName(tableNames []string) (results []*AllVirtualDdlOperation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDdlOperation{}).Where("`table_name` IN (?)", tableNames).Find(&results).Error

	return
}

// GetFromOperationType 通过operation_type获取内容
func (obj *_AllVirtualDdlOperationMgr) GetFromOperationType(operationType int64) (results []*AllVirtualDdlOperation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDdlOperation{}).Where("`operation_type` = ?", operationType).Find(&results).Error

	return
}

// GetBatchFromOperationType 批量查找
func (obj *_AllVirtualDdlOperationMgr) GetBatchFromOperationType(operationTypes []int64) (results []*AllVirtualDdlOperation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDdlOperation{}).Where("`operation_type` IN (?)", operationTypes).Find(&results).Error

	return
}

// GetFromDdlStmtStr 通过ddl_stmt_str获取内容
func (obj *_AllVirtualDdlOperationMgr) GetFromDdlStmtStr(ddlStmtStr string) (results []*AllVirtualDdlOperation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDdlOperation{}).Where("`ddl_stmt_str` = ?", ddlStmtStr).Find(&results).Error

	return
}

// GetBatchFromDdlStmtStr 批量查找
func (obj *_AllVirtualDdlOperationMgr) GetBatchFromDdlStmtStr(ddlStmtStrs []string) (results []*AllVirtualDdlOperation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDdlOperation{}).Where("`ddl_stmt_str` IN (?)", ddlStmtStrs).Find(&results).Error

	return
}

// GetFromExecTenantID 通过exec_tenant_id获取内容
func (obj *_AllVirtualDdlOperationMgr) GetFromExecTenantID(execTenantID int64) (results []*AllVirtualDdlOperation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDdlOperation{}).Where("`exec_tenant_id` = ?", execTenantID).Find(&results).Error

	return
}

// GetBatchFromExecTenantID 批量查找
func (obj *_AllVirtualDdlOperationMgr) GetBatchFromExecTenantID(execTenantIDs []int64) (results []*AllVirtualDdlOperation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDdlOperation{}).Where("`exec_tenant_id` IN (?)", execTenantIDs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllVirtualDdlOperationMgr) FetchByPrimaryKey(tenantID int64, schemaVersion int64) (result AllVirtualDdlOperation, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDdlOperation{}).Where("`tenant_id` = ? AND `schema_version` = ?", tenantID, schemaVersion).First(&result).Error

	return
}
