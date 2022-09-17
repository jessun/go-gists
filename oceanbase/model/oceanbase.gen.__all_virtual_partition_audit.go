package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _AllVirtualPartitionAuditMgr struct {
	*_BaseMgr
}

// AllVirtualPartitionAuditMgr open func
func AllVirtualPartitionAuditMgr(db *gorm.DB) *_AllVirtualPartitionAuditMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualPartitionAuditMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualPartitionAuditMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_partition_audit"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualPartitionAuditMgr) GetTableName() string {
	return "__all_virtual_partition_audit"
}

// Reset 重置gorm会话
func (obj *_AllVirtualPartitionAuditMgr) Reset() *_AllVirtualPartitionAuditMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualPartitionAuditMgr) Get() (result AllVirtualPartitionAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionAudit{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualPartitionAuditMgr) Gets() (results []*AllVirtualPartitionAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionAudit{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualPartitionAuditMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionAudit{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithSvrIP svr_ip获取
func (obj *_AllVirtualPartitionAuditMgr) WithSvrIP(svrIP string) Option {
	return optionFunc(func(o *options) { o.query["svr_ip"] = svrIP })
}

// WithSvrPort svr_port获取
func (obj *_AllVirtualPartitionAuditMgr) WithSvrPort(svrPort int64) Option {
	return optionFunc(func(o *options) { o.query["svr_port"] = svrPort })
}

// WithTenantID tenant_id获取
func (obj *_AllVirtualPartitionAuditMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithTableID table_id获取
func (obj *_AllVirtualPartitionAuditMgr) WithTableID(tableID int64) Option {
	return optionFunc(func(o *options) { o.query["table_id"] = tableID })
}

// WithPartitionID partition_id获取
func (obj *_AllVirtualPartitionAuditMgr) WithPartitionID(partitionID int64) Option {
	return optionFunc(func(o *options) { o.query["partition_id"] = partitionID })
}

// WithPartitionStatus partition_status获取
func (obj *_AllVirtualPartitionAuditMgr) WithPartitionStatus(partitionStatus int64) Option {
	return optionFunc(func(o *options) { o.query["partition_status"] = partitionStatus })
}

// WithBaseRowCount base_row_count获取
func (obj *_AllVirtualPartitionAuditMgr) WithBaseRowCount(baseRowCount int64) Option {
	return optionFunc(func(o *options) { o.query["base_row_count"] = baseRowCount })
}

// WithInsertRowCount insert_row_count获取
func (obj *_AllVirtualPartitionAuditMgr) WithInsertRowCount(insertRowCount int64) Option {
	return optionFunc(func(o *options) { o.query["insert_row_count"] = insertRowCount })
}

// WithDeleteRowCount delete_row_count获取
func (obj *_AllVirtualPartitionAuditMgr) WithDeleteRowCount(deleteRowCount int64) Option {
	return optionFunc(func(o *options) { o.query["delete_row_count"] = deleteRowCount })
}

// WithUpdateRowCount update_row_count获取
func (obj *_AllVirtualPartitionAuditMgr) WithUpdateRowCount(updateRowCount int64) Option {
	return optionFunc(func(o *options) { o.query["update_row_count"] = updateRowCount })
}

// WithQueryRowCount query_row_count获取
func (obj *_AllVirtualPartitionAuditMgr) WithQueryRowCount(queryRowCount int64) Option {
	return optionFunc(func(o *options) { o.query["query_row_count"] = queryRowCount })
}

// WithInsertSQLCount insert_sql_count获取
func (obj *_AllVirtualPartitionAuditMgr) WithInsertSQLCount(insertSQLCount int64) Option {
	return optionFunc(func(o *options) { o.query["insert_sql_count"] = insertSQLCount })
}

// WithDeleteSQLCount delete_sql_count获取
func (obj *_AllVirtualPartitionAuditMgr) WithDeleteSQLCount(deleteSQLCount int64) Option {
	return optionFunc(func(o *options) { o.query["delete_sql_count"] = deleteSQLCount })
}

// WithUpdateSQLCount update_sql_count获取
func (obj *_AllVirtualPartitionAuditMgr) WithUpdateSQLCount(updateSQLCount int64) Option {
	return optionFunc(func(o *options) { o.query["update_sql_count"] = updateSQLCount })
}

// WithQuerySQLCount query_sql_count获取
func (obj *_AllVirtualPartitionAuditMgr) WithQuerySQLCount(querySQLCount int64) Option {
	return optionFunc(func(o *options) { o.query["query_sql_count"] = querySQLCount })
}

// WithTransCount trans_count获取
func (obj *_AllVirtualPartitionAuditMgr) WithTransCount(transCount int64) Option {
	return optionFunc(func(o *options) { o.query["trans_count"] = transCount })
}

// WithSQLCount sql_count获取
func (obj *_AllVirtualPartitionAuditMgr) WithSQLCount(sqlCount int64) Option {
	return optionFunc(func(o *options) { o.query["sql_count"] = sqlCount })
}

// WithRollbackInsertRowCount rollback_insert_row_count获取
func (obj *_AllVirtualPartitionAuditMgr) WithRollbackInsertRowCount(rollbackInsertRowCount int64) Option {
	return optionFunc(func(o *options) { o.query["rollback_insert_row_count"] = rollbackInsertRowCount })
}

// WithRollbackDeleteRowCount rollback_delete_row_count获取
func (obj *_AllVirtualPartitionAuditMgr) WithRollbackDeleteRowCount(rollbackDeleteRowCount int64) Option {
	return optionFunc(func(o *options) { o.query["rollback_delete_row_count"] = rollbackDeleteRowCount })
}

// WithRollbackUpdateRowCount rollback_update_row_count获取
func (obj *_AllVirtualPartitionAuditMgr) WithRollbackUpdateRowCount(rollbackUpdateRowCount int64) Option {
	return optionFunc(func(o *options) { o.query["rollback_update_row_count"] = rollbackUpdateRowCount })
}

// WithRollbackInsertSQLCount rollback_insert_sql_count获取
func (obj *_AllVirtualPartitionAuditMgr) WithRollbackInsertSQLCount(rollbackInsertSQLCount int64) Option {
	return optionFunc(func(o *options) { o.query["rollback_insert_sql_count"] = rollbackInsertSQLCount })
}

// WithRollbackDeleteSQLCount rollback_delete_sql_count获取
func (obj *_AllVirtualPartitionAuditMgr) WithRollbackDeleteSQLCount(rollbackDeleteSQLCount int64) Option {
	return optionFunc(func(o *options) { o.query["rollback_delete_sql_count"] = rollbackDeleteSQLCount })
}

// WithRollbackUpdateSQLCount rollback_update_sql_count获取
func (obj *_AllVirtualPartitionAuditMgr) WithRollbackUpdateSQLCount(rollbackUpdateSQLCount int64) Option {
	return optionFunc(func(o *options) { o.query["rollback_update_sql_count"] = rollbackUpdateSQLCount })
}

// WithRollbackTransCount rollback_trans_count获取
func (obj *_AllVirtualPartitionAuditMgr) WithRollbackTransCount(rollbackTransCount int64) Option {
	return optionFunc(func(o *options) { o.query["rollback_trans_count"] = rollbackTransCount })
}

// WithRollbackSQLCount rollback_sql_count获取
func (obj *_AllVirtualPartitionAuditMgr) WithRollbackSQLCount(rollbackSQLCount int64) Option {
	return optionFunc(func(o *options) { o.query["rollback_sql_count"] = rollbackSQLCount })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualPartitionAuditMgr) GetByOption(opts ...Option) (result AllVirtualPartitionAudit, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionAudit{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualPartitionAuditMgr) GetByOptions(opts ...Option) (results []*AllVirtualPartitionAudit, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionAudit{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualPartitionAuditMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualPartitionAudit, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionAudit{}).Where(options.query)
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

// GetFromSvrIP 通过svr_ip获取内容
func (obj *_AllVirtualPartitionAuditMgr) GetFromSvrIP(svrIP string) (results []*AllVirtualPartitionAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionAudit{}).Where("`svr_ip` = ?", svrIP).Find(&results).Error

	return
}

// GetBatchFromSvrIP 批量查找
func (obj *_AllVirtualPartitionAuditMgr) GetBatchFromSvrIP(svrIPs []string) (results []*AllVirtualPartitionAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionAudit{}).Where("`svr_ip` IN (?)", svrIPs).Find(&results).Error

	return
}

// GetFromSvrPort 通过svr_port获取内容
func (obj *_AllVirtualPartitionAuditMgr) GetFromSvrPort(svrPort int64) (results []*AllVirtualPartitionAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionAudit{}).Where("`svr_port` = ?", svrPort).Find(&results).Error

	return
}

// GetBatchFromSvrPort 批量查找
func (obj *_AllVirtualPartitionAuditMgr) GetBatchFromSvrPort(svrPorts []int64) (results []*AllVirtualPartitionAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionAudit{}).Where("`svr_port` IN (?)", svrPorts).Find(&results).Error

	return
}

// GetFromTenantID 通过tenant_id获取内容
func (obj *_AllVirtualPartitionAuditMgr) GetFromTenantID(tenantID int64) (results []*AllVirtualPartitionAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionAudit{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllVirtualPartitionAuditMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllVirtualPartitionAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionAudit{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromTableID 通过table_id获取内容
func (obj *_AllVirtualPartitionAuditMgr) GetFromTableID(tableID int64) (results []*AllVirtualPartitionAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionAudit{}).Where("`table_id` = ?", tableID).Find(&results).Error

	return
}

// GetBatchFromTableID 批量查找
func (obj *_AllVirtualPartitionAuditMgr) GetBatchFromTableID(tableIDs []int64) (results []*AllVirtualPartitionAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionAudit{}).Where("`table_id` IN (?)", tableIDs).Find(&results).Error

	return
}

// GetFromPartitionID 通过partition_id获取内容
func (obj *_AllVirtualPartitionAuditMgr) GetFromPartitionID(partitionID int64) (results []*AllVirtualPartitionAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionAudit{}).Where("`partition_id` = ?", partitionID).Find(&results).Error

	return
}

// GetBatchFromPartitionID 批量查找
func (obj *_AllVirtualPartitionAuditMgr) GetBatchFromPartitionID(partitionIDs []int64) (results []*AllVirtualPartitionAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionAudit{}).Where("`partition_id` IN (?)", partitionIDs).Find(&results).Error

	return
}

// GetFromPartitionStatus 通过partition_status获取内容
func (obj *_AllVirtualPartitionAuditMgr) GetFromPartitionStatus(partitionStatus int64) (results []*AllVirtualPartitionAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionAudit{}).Where("`partition_status` = ?", partitionStatus).Find(&results).Error

	return
}

// GetBatchFromPartitionStatus 批量查找
func (obj *_AllVirtualPartitionAuditMgr) GetBatchFromPartitionStatus(partitionStatuss []int64) (results []*AllVirtualPartitionAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionAudit{}).Where("`partition_status` IN (?)", partitionStatuss).Find(&results).Error

	return
}

// GetFromBaseRowCount 通过base_row_count获取内容
func (obj *_AllVirtualPartitionAuditMgr) GetFromBaseRowCount(baseRowCount int64) (results []*AllVirtualPartitionAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionAudit{}).Where("`base_row_count` = ?", baseRowCount).Find(&results).Error

	return
}

// GetBatchFromBaseRowCount 批量查找
func (obj *_AllVirtualPartitionAuditMgr) GetBatchFromBaseRowCount(baseRowCounts []int64) (results []*AllVirtualPartitionAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionAudit{}).Where("`base_row_count` IN (?)", baseRowCounts).Find(&results).Error

	return
}

// GetFromInsertRowCount 通过insert_row_count获取内容
func (obj *_AllVirtualPartitionAuditMgr) GetFromInsertRowCount(insertRowCount int64) (results []*AllVirtualPartitionAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionAudit{}).Where("`insert_row_count` = ?", insertRowCount).Find(&results).Error

	return
}

// GetBatchFromInsertRowCount 批量查找
func (obj *_AllVirtualPartitionAuditMgr) GetBatchFromInsertRowCount(insertRowCounts []int64) (results []*AllVirtualPartitionAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionAudit{}).Where("`insert_row_count` IN (?)", insertRowCounts).Find(&results).Error

	return
}

// GetFromDeleteRowCount 通过delete_row_count获取内容
func (obj *_AllVirtualPartitionAuditMgr) GetFromDeleteRowCount(deleteRowCount int64) (results []*AllVirtualPartitionAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionAudit{}).Where("`delete_row_count` = ?", deleteRowCount).Find(&results).Error

	return
}

// GetBatchFromDeleteRowCount 批量查找
func (obj *_AllVirtualPartitionAuditMgr) GetBatchFromDeleteRowCount(deleteRowCounts []int64) (results []*AllVirtualPartitionAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionAudit{}).Where("`delete_row_count` IN (?)", deleteRowCounts).Find(&results).Error

	return
}

// GetFromUpdateRowCount 通过update_row_count获取内容
func (obj *_AllVirtualPartitionAuditMgr) GetFromUpdateRowCount(updateRowCount int64) (results []*AllVirtualPartitionAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionAudit{}).Where("`update_row_count` = ?", updateRowCount).Find(&results).Error

	return
}

// GetBatchFromUpdateRowCount 批量查找
func (obj *_AllVirtualPartitionAuditMgr) GetBatchFromUpdateRowCount(updateRowCounts []int64) (results []*AllVirtualPartitionAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionAudit{}).Where("`update_row_count` IN (?)", updateRowCounts).Find(&results).Error

	return
}

// GetFromQueryRowCount 通过query_row_count获取内容
func (obj *_AllVirtualPartitionAuditMgr) GetFromQueryRowCount(queryRowCount int64) (results []*AllVirtualPartitionAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionAudit{}).Where("`query_row_count` = ?", queryRowCount).Find(&results).Error

	return
}

// GetBatchFromQueryRowCount 批量查找
func (obj *_AllVirtualPartitionAuditMgr) GetBatchFromQueryRowCount(queryRowCounts []int64) (results []*AllVirtualPartitionAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionAudit{}).Where("`query_row_count` IN (?)", queryRowCounts).Find(&results).Error

	return
}

// GetFromInsertSQLCount 通过insert_sql_count获取内容
func (obj *_AllVirtualPartitionAuditMgr) GetFromInsertSQLCount(insertSQLCount int64) (results []*AllVirtualPartitionAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionAudit{}).Where("`insert_sql_count` = ?", insertSQLCount).Find(&results).Error

	return
}

// GetBatchFromInsertSQLCount 批量查找
func (obj *_AllVirtualPartitionAuditMgr) GetBatchFromInsertSQLCount(insertSQLCounts []int64) (results []*AllVirtualPartitionAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionAudit{}).Where("`insert_sql_count` IN (?)", insertSQLCounts).Find(&results).Error

	return
}

// GetFromDeleteSQLCount 通过delete_sql_count获取内容
func (obj *_AllVirtualPartitionAuditMgr) GetFromDeleteSQLCount(deleteSQLCount int64) (results []*AllVirtualPartitionAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionAudit{}).Where("`delete_sql_count` = ?", deleteSQLCount).Find(&results).Error

	return
}

// GetBatchFromDeleteSQLCount 批量查找
func (obj *_AllVirtualPartitionAuditMgr) GetBatchFromDeleteSQLCount(deleteSQLCounts []int64) (results []*AllVirtualPartitionAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionAudit{}).Where("`delete_sql_count` IN (?)", deleteSQLCounts).Find(&results).Error

	return
}

// GetFromUpdateSQLCount 通过update_sql_count获取内容
func (obj *_AllVirtualPartitionAuditMgr) GetFromUpdateSQLCount(updateSQLCount int64) (results []*AllVirtualPartitionAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionAudit{}).Where("`update_sql_count` = ?", updateSQLCount).Find(&results).Error

	return
}

// GetBatchFromUpdateSQLCount 批量查找
func (obj *_AllVirtualPartitionAuditMgr) GetBatchFromUpdateSQLCount(updateSQLCounts []int64) (results []*AllVirtualPartitionAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionAudit{}).Where("`update_sql_count` IN (?)", updateSQLCounts).Find(&results).Error

	return
}

// GetFromQuerySQLCount 通过query_sql_count获取内容
func (obj *_AllVirtualPartitionAuditMgr) GetFromQuerySQLCount(querySQLCount int64) (results []*AllVirtualPartitionAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionAudit{}).Where("`query_sql_count` = ?", querySQLCount).Find(&results).Error

	return
}

// GetBatchFromQuerySQLCount 批量查找
func (obj *_AllVirtualPartitionAuditMgr) GetBatchFromQuerySQLCount(querySQLCounts []int64) (results []*AllVirtualPartitionAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionAudit{}).Where("`query_sql_count` IN (?)", querySQLCounts).Find(&results).Error

	return
}

// GetFromTransCount 通过trans_count获取内容
func (obj *_AllVirtualPartitionAuditMgr) GetFromTransCount(transCount int64) (results []*AllVirtualPartitionAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionAudit{}).Where("`trans_count` = ?", transCount).Find(&results).Error

	return
}

// GetBatchFromTransCount 批量查找
func (obj *_AllVirtualPartitionAuditMgr) GetBatchFromTransCount(transCounts []int64) (results []*AllVirtualPartitionAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionAudit{}).Where("`trans_count` IN (?)", transCounts).Find(&results).Error

	return
}

// GetFromSQLCount 通过sql_count获取内容
func (obj *_AllVirtualPartitionAuditMgr) GetFromSQLCount(sqlCount int64) (results []*AllVirtualPartitionAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionAudit{}).Where("`sql_count` = ?", sqlCount).Find(&results).Error

	return
}

// GetBatchFromSQLCount 批量查找
func (obj *_AllVirtualPartitionAuditMgr) GetBatchFromSQLCount(sqlCounts []int64) (results []*AllVirtualPartitionAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionAudit{}).Where("`sql_count` IN (?)", sqlCounts).Find(&results).Error

	return
}

// GetFromRollbackInsertRowCount 通过rollback_insert_row_count获取内容
func (obj *_AllVirtualPartitionAuditMgr) GetFromRollbackInsertRowCount(rollbackInsertRowCount int64) (results []*AllVirtualPartitionAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionAudit{}).Where("`rollback_insert_row_count` = ?", rollbackInsertRowCount).Find(&results).Error

	return
}

// GetBatchFromRollbackInsertRowCount 批量查找
func (obj *_AllVirtualPartitionAuditMgr) GetBatchFromRollbackInsertRowCount(rollbackInsertRowCounts []int64) (results []*AllVirtualPartitionAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionAudit{}).Where("`rollback_insert_row_count` IN (?)", rollbackInsertRowCounts).Find(&results).Error

	return
}

// GetFromRollbackDeleteRowCount 通过rollback_delete_row_count获取内容
func (obj *_AllVirtualPartitionAuditMgr) GetFromRollbackDeleteRowCount(rollbackDeleteRowCount int64) (results []*AllVirtualPartitionAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionAudit{}).Where("`rollback_delete_row_count` = ?", rollbackDeleteRowCount).Find(&results).Error

	return
}

// GetBatchFromRollbackDeleteRowCount 批量查找
func (obj *_AllVirtualPartitionAuditMgr) GetBatchFromRollbackDeleteRowCount(rollbackDeleteRowCounts []int64) (results []*AllVirtualPartitionAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionAudit{}).Where("`rollback_delete_row_count` IN (?)", rollbackDeleteRowCounts).Find(&results).Error

	return
}

// GetFromRollbackUpdateRowCount 通过rollback_update_row_count获取内容
func (obj *_AllVirtualPartitionAuditMgr) GetFromRollbackUpdateRowCount(rollbackUpdateRowCount int64) (results []*AllVirtualPartitionAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionAudit{}).Where("`rollback_update_row_count` = ?", rollbackUpdateRowCount).Find(&results).Error

	return
}

// GetBatchFromRollbackUpdateRowCount 批量查找
func (obj *_AllVirtualPartitionAuditMgr) GetBatchFromRollbackUpdateRowCount(rollbackUpdateRowCounts []int64) (results []*AllVirtualPartitionAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionAudit{}).Where("`rollback_update_row_count` IN (?)", rollbackUpdateRowCounts).Find(&results).Error

	return
}

// GetFromRollbackInsertSQLCount 通过rollback_insert_sql_count获取内容
func (obj *_AllVirtualPartitionAuditMgr) GetFromRollbackInsertSQLCount(rollbackInsertSQLCount int64) (results []*AllVirtualPartitionAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionAudit{}).Where("`rollback_insert_sql_count` = ?", rollbackInsertSQLCount).Find(&results).Error

	return
}

// GetBatchFromRollbackInsertSQLCount 批量查找
func (obj *_AllVirtualPartitionAuditMgr) GetBatchFromRollbackInsertSQLCount(rollbackInsertSQLCounts []int64) (results []*AllVirtualPartitionAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionAudit{}).Where("`rollback_insert_sql_count` IN (?)", rollbackInsertSQLCounts).Find(&results).Error

	return
}

// GetFromRollbackDeleteSQLCount 通过rollback_delete_sql_count获取内容
func (obj *_AllVirtualPartitionAuditMgr) GetFromRollbackDeleteSQLCount(rollbackDeleteSQLCount int64) (results []*AllVirtualPartitionAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionAudit{}).Where("`rollback_delete_sql_count` = ?", rollbackDeleteSQLCount).Find(&results).Error

	return
}

// GetBatchFromRollbackDeleteSQLCount 批量查找
func (obj *_AllVirtualPartitionAuditMgr) GetBatchFromRollbackDeleteSQLCount(rollbackDeleteSQLCounts []int64) (results []*AllVirtualPartitionAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionAudit{}).Where("`rollback_delete_sql_count` IN (?)", rollbackDeleteSQLCounts).Find(&results).Error

	return
}

// GetFromRollbackUpdateSQLCount 通过rollback_update_sql_count获取内容
func (obj *_AllVirtualPartitionAuditMgr) GetFromRollbackUpdateSQLCount(rollbackUpdateSQLCount int64) (results []*AllVirtualPartitionAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionAudit{}).Where("`rollback_update_sql_count` = ?", rollbackUpdateSQLCount).Find(&results).Error

	return
}

// GetBatchFromRollbackUpdateSQLCount 批量查找
func (obj *_AllVirtualPartitionAuditMgr) GetBatchFromRollbackUpdateSQLCount(rollbackUpdateSQLCounts []int64) (results []*AllVirtualPartitionAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionAudit{}).Where("`rollback_update_sql_count` IN (?)", rollbackUpdateSQLCounts).Find(&results).Error

	return
}

// GetFromRollbackTransCount 通过rollback_trans_count获取内容
func (obj *_AllVirtualPartitionAuditMgr) GetFromRollbackTransCount(rollbackTransCount int64) (results []*AllVirtualPartitionAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionAudit{}).Where("`rollback_trans_count` = ?", rollbackTransCount).Find(&results).Error

	return
}

// GetBatchFromRollbackTransCount 批量查找
func (obj *_AllVirtualPartitionAuditMgr) GetBatchFromRollbackTransCount(rollbackTransCounts []int64) (results []*AllVirtualPartitionAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionAudit{}).Where("`rollback_trans_count` IN (?)", rollbackTransCounts).Find(&results).Error

	return
}

// GetFromRollbackSQLCount 通过rollback_sql_count获取内容
func (obj *_AllVirtualPartitionAuditMgr) GetFromRollbackSQLCount(rollbackSQLCount int64) (results []*AllVirtualPartitionAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionAudit{}).Where("`rollback_sql_count` = ?", rollbackSQLCount).Find(&results).Error

	return
}

// GetBatchFromRollbackSQLCount 批量查找
func (obj *_AllVirtualPartitionAuditMgr) GetBatchFromRollbackSQLCount(rollbackSQLCounts []int64) (results []*AllVirtualPartitionAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionAudit{}).Where("`rollback_sql_count` IN (?)", rollbackSQLCounts).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////
