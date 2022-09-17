package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
)	

type _V$partitionAuditMgr struct {
	*_BaseMgr
}

// V$partitionAuditMgr open func
func V$partitionAuditMgr(db *gorm.DB) *_V$partitionAuditMgr {
	if db == nil {
		panic(fmt.Errorf("V$partitionAuditMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_V$partitionAuditMgr{_BaseMgr: &_BaseMgr{DB: db.Table("v$partition_audit"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_V$partitionAuditMgr) GetTableName() string {
	return "v$partition_audit"
}

// Reset 重置gorm会话
func (obj *_V$partitionAuditMgr) Reset() *_V$partitionAuditMgr {
	obj.New()
	return obj
}

// Get 获取 
func (obj *_V$partitionAuditMgr) Get() (result V$partitionAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$partitionAudit{}).First(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_V$partitionAuditMgr) Gets() (results []*V$partitionAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$partitionAudit{}).Find(&results).Error
	
	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_V$partitionAuditMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(V$partitionAudit{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithSvrIP svr_ip获取 
func (obj *_V$partitionAuditMgr) WithSvrIP(svrIP string) Option {
	return optionFunc(func(o *options) { o.query["svr_ip"] = svrIP })
}

// WithSvrPort svr_port获取 
func (obj *_V$partitionAuditMgr) WithSvrPort(svrPort int64) Option {
	return optionFunc(func(o *options) { o.query["svr_port"] = svrPort })
}

// WithTenantID tenant_id获取 
func (obj *_V$partitionAuditMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithTableID table_id获取 
func (obj *_V$partitionAuditMgr) WithTableID(tableID int64) Option {
	return optionFunc(func(o *options) { o.query["table_id"] = tableID })
}

// WithPartitionID partition_id获取 
func (obj *_V$partitionAuditMgr) WithPartitionID(partitionID int64) Option {
	return optionFunc(func(o *options) { o.query["partition_id"] = partitionID })
}

// WithPartitionStatus partition_status获取 
func (obj *_V$partitionAuditMgr) WithPartitionStatus(partitionStatus int64) Option {
	return optionFunc(func(o *options) { o.query["partition_status"] = partitionStatus })
}

// WithBaseRowCount base_row_count获取 
func (obj *_V$partitionAuditMgr) WithBaseRowCount(baseRowCount int64) Option {
	return optionFunc(func(o *options) { o.query["base_row_count"] = baseRowCount })
}

// WithInsertRowCount insert_row_count获取 
func (obj *_V$partitionAuditMgr) WithInsertRowCount(insertRowCount int64) Option {
	return optionFunc(func(o *options) { o.query["insert_row_count"] = insertRowCount })
}

// WithDeleteRowCount delete_row_count获取 
func (obj *_V$partitionAuditMgr) WithDeleteRowCount(deleteRowCount int64) Option {
	return optionFunc(func(o *options) { o.query["delete_row_count"] = deleteRowCount })
}

// WithUpdateRowCount update_row_count获取 
func (obj *_V$partitionAuditMgr) WithUpdateRowCount(updateRowCount int64) Option {
	return optionFunc(func(o *options) { o.query["update_row_count"] = updateRowCount })
}

// WithQueryRowCount query_row_count获取 
func (obj *_V$partitionAuditMgr) WithQueryRowCount(queryRowCount int64) Option {
	return optionFunc(func(o *options) { o.query["query_row_count"] = queryRowCount })
}

// WithInsertSQLCount insert_sql_count获取 
func (obj *_V$partitionAuditMgr) WithInsertSQLCount(insertSQLCount int64) Option {
	return optionFunc(func(o *options) { o.query["insert_sql_count"] = insertSQLCount })
}

// WithDeleteSQLCount delete_sql_count获取 
func (obj *_V$partitionAuditMgr) WithDeleteSQLCount(deleteSQLCount int64) Option {
	return optionFunc(func(o *options) { o.query["delete_sql_count"] = deleteSQLCount })
}

// WithUpdateSQLCount update_sql_count获取 
func (obj *_V$partitionAuditMgr) WithUpdateSQLCount(updateSQLCount int64) Option {
	return optionFunc(func(o *options) { o.query["update_sql_count"] = updateSQLCount })
}

// WithQuerySQLCount query_sql_count获取 
func (obj *_V$partitionAuditMgr) WithQuerySQLCount(querySQLCount int64) Option {
	return optionFunc(func(o *options) { o.query["query_sql_count"] = querySQLCount })
}

// WithTransCount trans_count获取 
func (obj *_V$partitionAuditMgr) WithTransCount(transCount int64) Option {
	return optionFunc(func(o *options) { o.query["trans_count"] = transCount })
}

// WithSQLCount sql_count获取 
func (obj *_V$partitionAuditMgr) WithSQLCount(sqlCount int64) Option {
	return optionFunc(func(o *options) { o.query["sql_count"] = sqlCount })
}

// WithRollbackInsertRowCount rollback_insert_row_count获取 
func (obj *_V$partitionAuditMgr) WithRollbackInsertRowCount(rollbackInsertRowCount int64) Option {
	return optionFunc(func(o *options) { o.query["rollback_insert_row_count"] = rollbackInsertRowCount })
}

// WithRollbackDeleteRowCount rollback_delete_row_count获取 
func (obj *_V$partitionAuditMgr) WithRollbackDeleteRowCount(rollbackDeleteRowCount int64) Option {
	return optionFunc(func(o *options) { o.query["rollback_delete_row_count"] = rollbackDeleteRowCount })
}

// WithRollbackUpdateRowCount rollback_update_row_count获取 
func (obj *_V$partitionAuditMgr) WithRollbackUpdateRowCount(rollbackUpdateRowCount int64) Option {
	return optionFunc(func(o *options) { o.query["rollback_update_row_count"] = rollbackUpdateRowCount })
}

// WithRollbackInsertSQLCount rollback_insert_sql_count获取 
func (obj *_V$partitionAuditMgr) WithRollbackInsertSQLCount(rollbackInsertSQLCount int64) Option {
	return optionFunc(func(o *options) { o.query["rollback_insert_sql_count"] = rollbackInsertSQLCount })
}

// WithRollbackDeleteSQLCount rollback_delete_sql_count获取 
func (obj *_V$partitionAuditMgr) WithRollbackDeleteSQLCount(rollbackDeleteSQLCount int64) Option {
	return optionFunc(func(o *options) { o.query["rollback_delete_sql_count"] = rollbackDeleteSQLCount })
}

// WithRollbackUpdateSQLCount rollback_update_sql_count获取 
func (obj *_V$partitionAuditMgr) WithRollbackUpdateSQLCount(rollbackUpdateSQLCount int64) Option {
	return optionFunc(func(o *options) { o.query["rollback_update_sql_count"] = rollbackUpdateSQLCount })
}

// WithRollbackTransCount rollback_trans_count获取 
func (obj *_V$partitionAuditMgr) WithRollbackTransCount(rollbackTransCount int64) Option {
	return optionFunc(func(o *options) { o.query["rollback_trans_count"] = rollbackTransCount })
}

// WithRollbackSQLCount rollback_sql_count获取 
func (obj *_V$partitionAuditMgr) WithRollbackSQLCount(rollbackSQLCount int64) Option {
	return optionFunc(func(o *options) { o.query["rollback_sql_count"] = rollbackSQLCount })
}


// GetByOption 功能选项模式获取
func (obj *_V$partitionAuditMgr) GetByOption(opts ...Option) (result V$partitionAudit, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(V$partitionAudit{}).Where(options.query).First(&result).Error
	
	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_V$partitionAuditMgr) GetByOptions(opts ...Option) (results []*V$partitionAudit, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(V$partitionAudit{}).Where(options.query).Find(&results).Error
	
	return
}


// SelectPage 分页查询
func (obj *_V$partitionAuditMgr) SelectPage(page IPage,opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]V$partitionAudit,0)
	var count int64 // 统计总的记录数
	query :=  obj.DB.WithContext(obj.ctx).Model(V$partitionAudit{}).Where(options.query)
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
func (obj *_V$partitionAuditMgr) GetFromSvrIP(svrIP string) (results []*V$partitionAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$partitionAudit{}).Where("`svr_ip` = ?", svrIP).Find(&results).Error
	
	return
}

// GetBatchFromSvrIP 批量查找 
func (obj *_V$partitionAuditMgr) GetBatchFromSvrIP(svrIPs []string) (results []*V$partitionAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$partitionAudit{}).Where("`svr_ip` IN (?)", svrIPs).Find(&results).Error
	
	return
}
 
// GetFromSvrPort 通过svr_port获取内容  
func (obj *_V$partitionAuditMgr) GetFromSvrPort(svrPort int64) (results []*V$partitionAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$partitionAudit{}).Where("`svr_port` = ?", svrPort).Find(&results).Error
	
	return
}

// GetBatchFromSvrPort 批量查找 
func (obj *_V$partitionAuditMgr) GetBatchFromSvrPort(svrPorts []int64) (results []*V$partitionAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$partitionAudit{}).Where("`svr_port` IN (?)", svrPorts).Find(&results).Error
	
	return
}
 
// GetFromTenantID 通过tenant_id获取内容  
func (obj *_V$partitionAuditMgr) GetFromTenantID(tenantID int64) (results []*V$partitionAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$partitionAudit{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error
	
	return
}

// GetBatchFromTenantID 批量查找 
func (obj *_V$partitionAuditMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*V$partitionAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$partitionAudit{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error
	
	return
}
 
// GetFromTableID 通过table_id获取内容  
func (obj *_V$partitionAuditMgr) GetFromTableID(tableID int64) (results []*V$partitionAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$partitionAudit{}).Where("`table_id` = ?", tableID).Find(&results).Error
	
	return
}

// GetBatchFromTableID 批量查找 
func (obj *_V$partitionAuditMgr) GetBatchFromTableID(tableIDs []int64) (results []*V$partitionAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$partitionAudit{}).Where("`table_id` IN (?)", tableIDs).Find(&results).Error
	
	return
}
 
// GetFromPartitionID 通过partition_id获取内容  
func (obj *_V$partitionAuditMgr) GetFromPartitionID(partitionID int64) (results []*V$partitionAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$partitionAudit{}).Where("`partition_id` = ?", partitionID).Find(&results).Error
	
	return
}

// GetBatchFromPartitionID 批量查找 
func (obj *_V$partitionAuditMgr) GetBatchFromPartitionID(partitionIDs []int64) (results []*V$partitionAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$partitionAudit{}).Where("`partition_id` IN (?)", partitionIDs).Find(&results).Error
	
	return
}
 
// GetFromPartitionStatus 通过partition_status获取内容  
func (obj *_V$partitionAuditMgr) GetFromPartitionStatus(partitionStatus int64) (results []*V$partitionAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$partitionAudit{}).Where("`partition_status` = ?", partitionStatus).Find(&results).Error
	
	return
}

// GetBatchFromPartitionStatus 批量查找 
func (obj *_V$partitionAuditMgr) GetBatchFromPartitionStatus(partitionStatuss []int64) (results []*V$partitionAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$partitionAudit{}).Where("`partition_status` IN (?)", partitionStatuss).Find(&results).Error
	
	return
}
 
// GetFromBaseRowCount 通过base_row_count获取内容  
func (obj *_V$partitionAuditMgr) GetFromBaseRowCount(baseRowCount int64) (results []*V$partitionAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$partitionAudit{}).Where("`base_row_count` = ?", baseRowCount).Find(&results).Error
	
	return
}

// GetBatchFromBaseRowCount 批量查找 
func (obj *_V$partitionAuditMgr) GetBatchFromBaseRowCount(baseRowCounts []int64) (results []*V$partitionAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$partitionAudit{}).Where("`base_row_count` IN (?)", baseRowCounts).Find(&results).Error
	
	return
}
 
// GetFromInsertRowCount 通过insert_row_count获取内容  
func (obj *_V$partitionAuditMgr) GetFromInsertRowCount(insertRowCount int64) (results []*V$partitionAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$partitionAudit{}).Where("`insert_row_count` = ?", insertRowCount).Find(&results).Error
	
	return
}

// GetBatchFromInsertRowCount 批量查找 
func (obj *_V$partitionAuditMgr) GetBatchFromInsertRowCount(insertRowCounts []int64) (results []*V$partitionAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$partitionAudit{}).Where("`insert_row_count` IN (?)", insertRowCounts).Find(&results).Error
	
	return
}
 
// GetFromDeleteRowCount 通过delete_row_count获取内容  
func (obj *_V$partitionAuditMgr) GetFromDeleteRowCount(deleteRowCount int64) (results []*V$partitionAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$partitionAudit{}).Where("`delete_row_count` = ?", deleteRowCount).Find(&results).Error
	
	return
}

// GetBatchFromDeleteRowCount 批量查找 
func (obj *_V$partitionAuditMgr) GetBatchFromDeleteRowCount(deleteRowCounts []int64) (results []*V$partitionAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$partitionAudit{}).Where("`delete_row_count` IN (?)", deleteRowCounts).Find(&results).Error
	
	return
}
 
// GetFromUpdateRowCount 通过update_row_count获取内容  
func (obj *_V$partitionAuditMgr) GetFromUpdateRowCount(updateRowCount int64) (results []*V$partitionAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$partitionAudit{}).Where("`update_row_count` = ?", updateRowCount).Find(&results).Error
	
	return
}

// GetBatchFromUpdateRowCount 批量查找 
func (obj *_V$partitionAuditMgr) GetBatchFromUpdateRowCount(updateRowCounts []int64) (results []*V$partitionAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$partitionAudit{}).Where("`update_row_count` IN (?)", updateRowCounts).Find(&results).Error
	
	return
}
 
// GetFromQueryRowCount 通过query_row_count获取内容  
func (obj *_V$partitionAuditMgr) GetFromQueryRowCount(queryRowCount int64) (results []*V$partitionAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$partitionAudit{}).Where("`query_row_count` = ?", queryRowCount).Find(&results).Error
	
	return
}

// GetBatchFromQueryRowCount 批量查找 
func (obj *_V$partitionAuditMgr) GetBatchFromQueryRowCount(queryRowCounts []int64) (results []*V$partitionAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$partitionAudit{}).Where("`query_row_count` IN (?)", queryRowCounts).Find(&results).Error
	
	return
}
 
// GetFromInsertSQLCount 通过insert_sql_count获取内容  
func (obj *_V$partitionAuditMgr) GetFromInsertSQLCount(insertSQLCount int64) (results []*V$partitionAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$partitionAudit{}).Where("`insert_sql_count` = ?", insertSQLCount).Find(&results).Error
	
	return
}

// GetBatchFromInsertSQLCount 批量查找 
func (obj *_V$partitionAuditMgr) GetBatchFromInsertSQLCount(insertSQLCounts []int64) (results []*V$partitionAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$partitionAudit{}).Where("`insert_sql_count` IN (?)", insertSQLCounts).Find(&results).Error
	
	return
}
 
// GetFromDeleteSQLCount 通过delete_sql_count获取内容  
func (obj *_V$partitionAuditMgr) GetFromDeleteSQLCount(deleteSQLCount int64) (results []*V$partitionAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$partitionAudit{}).Where("`delete_sql_count` = ?", deleteSQLCount).Find(&results).Error
	
	return
}

// GetBatchFromDeleteSQLCount 批量查找 
func (obj *_V$partitionAuditMgr) GetBatchFromDeleteSQLCount(deleteSQLCounts []int64) (results []*V$partitionAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$partitionAudit{}).Where("`delete_sql_count` IN (?)", deleteSQLCounts).Find(&results).Error
	
	return
}
 
// GetFromUpdateSQLCount 通过update_sql_count获取内容  
func (obj *_V$partitionAuditMgr) GetFromUpdateSQLCount(updateSQLCount int64) (results []*V$partitionAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$partitionAudit{}).Where("`update_sql_count` = ?", updateSQLCount).Find(&results).Error
	
	return
}

// GetBatchFromUpdateSQLCount 批量查找 
func (obj *_V$partitionAuditMgr) GetBatchFromUpdateSQLCount(updateSQLCounts []int64) (results []*V$partitionAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$partitionAudit{}).Where("`update_sql_count` IN (?)", updateSQLCounts).Find(&results).Error
	
	return
}
 
// GetFromQuerySQLCount 通过query_sql_count获取内容  
func (obj *_V$partitionAuditMgr) GetFromQuerySQLCount(querySQLCount int64) (results []*V$partitionAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$partitionAudit{}).Where("`query_sql_count` = ?", querySQLCount).Find(&results).Error
	
	return
}

// GetBatchFromQuerySQLCount 批量查找 
func (obj *_V$partitionAuditMgr) GetBatchFromQuerySQLCount(querySQLCounts []int64) (results []*V$partitionAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$partitionAudit{}).Where("`query_sql_count` IN (?)", querySQLCounts).Find(&results).Error
	
	return
}
 
// GetFromTransCount 通过trans_count获取内容  
func (obj *_V$partitionAuditMgr) GetFromTransCount(transCount int64) (results []*V$partitionAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$partitionAudit{}).Where("`trans_count` = ?", transCount).Find(&results).Error
	
	return
}

// GetBatchFromTransCount 批量查找 
func (obj *_V$partitionAuditMgr) GetBatchFromTransCount(transCounts []int64) (results []*V$partitionAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$partitionAudit{}).Where("`trans_count` IN (?)", transCounts).Find(&results).Error
	
	return
}
 
// GetFromSQLCount 通过sql_count获取内容  
func (obj *_V$partitionAuditMgr) GetFromSQLCount(sqlCount int64) (results []*V$partitionAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$partitionAudit{}).Where("`sql_count` = ?", sqlCount).Find(&results).Error
	
	return
}

// GetBatchFromSQLCount 批量查找 
func (obj *_V$partitionAuditMgr) GetBatchFromSQLCount(sqlCounts []int64) (results []*V$partitionAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$partitionAudit{}).Where("`sql_count` IN (?)", sqlCounts).Find(&results).Error
	
	return
}
 
// GetFromRollbackInsertRowCount 通过rollback_insert_row_count获取内容  
func (obj *_V$partitionAuditMgr) GetFromRollbackInsertRowCount(rollbackInsertRowCount int64) (results []*V$partitionAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$partitionAudit{}).Where("`rollback_insert_row_count` = ?", rollbackInsertRowCount).Find(&results).Error
	
	return
}

// GetBatchFromRollbackInsertRowCount 批量查找 
func (obj *_V$partitionAuditMgr) GetBatchFromRollbackInsertRowCount(rollbackInsertRowCounts []int64) (results []*V$partitionAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$partitionAudit{}).Where("`rollback_insert_row_count` IN (?)", rollbackInsertRowCounts).Find(&results).Error
	
	return
}
 
// GetFromRollbackDeleteRowCount 通过rollback_delete_row_count获取内容  
func (obj *_V$partitionAuditMgr) GetFromRollbackDeleteRowCount(rollbackDeleteRowCount int64) (results []*V$partitionAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$partitionAudit{}).Where("`rollback_delete_row_count` = ?", rollbackDeleteRowCount).Find(&results).Error
	
	return
}

// GetBatchFromRollbackDeleteRowCount 批量查找 
func (obj *_V$partitionAuditMgr) GetBatchFromRollbackDeleteRowCount(rollbackDeleteRowCounts []int64) (results []*V$partitionAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$partitionAudit{}).Where("`rollback_delete_row_count` IN (?)", rollbackDeleteRowCounts).Find(&results).Error
	
	return
}
 
// GetFromRollbackUpdateRowCount 通过rollback_update_row_count获取内容  
func (obj *_V$partitionAuditMgr) GetFromRollbackUpdateRowCount(rollbackUpdateRowCount int64) (results []*V$partitionAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$partitionAudit{}).Where("`rollback_update_row_count` = ?", rollbackUpdateRowCount).Find(&results).Error
	
	return
}

// GetBatchFromRollbackUpdateRowCount 批量查找 
func (obj *_V$partitionAuditMgr) GetBatchFromRollbackUpdateRowCount(rollbackUpdateRowCounts []int64) (results []*V$partitionAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$partitionAudit{}).Where("`rollback_update_row_count` IN (?)", rollbackUpdateRowCounts).Find(&results).Error
	
	return
}
 
// GetFromRollbackInsertSQLCount 通过rollback_insert_sql_count获取内容  
func (obj *_V$partitionAuditMgr) GetFromRollbackInsertSQLCount(rollbackInsertSQLCount int64) (results []*V$partitionAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$partitionAudit{}).Where("`rollback_insert_sql_count` = ?", rollbackInsertSQLCount).Find(&results).Error
	
	return
}

// GetBatchFromRollbackInsertSQLCount 批量查找 
func (obj *_V$partitionAuditMgr) GetBatchFromRollbackInsertSQLCount(rollbackInsertSQLCounts []int64) (results []*V$partitionAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$partitionAudit{}).Where("`rollback_insert_sql_count` IN (?)", rollbackInsertSQLCounts).Find(&results).Error
	
	return
}
 
// GetFromRollbackDeleteSQLCount 通过rollback_delete_sql_count获取内容  
func (obj *_V$partitionAuditMgr) GetFromRollbackDeleteSQLCount(rollbackDeleteSQLCount int64) (results []*V$partitionAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$partitionAudit{}).Where("`rollback_delete_sql_count` = ?", rollbackDeleteSQLCount).Find(&results).Error
	
	return
}

// GetBatchFromRollbackDeleteSQLCount 批量查找 
func (obj *_V$partitionAuditMgr) GetBatchFromRollbackDeleteSQLCount(rollbackDeleteSQLCounts []int64) (results []*V$partitionAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$partitionAudit{}).Where("`rollback_delete_sql_count` IN (?)", rollbackDeleteSQLCounts).Find(&results).Error
	
	return
}
 
// GetFromRollbackUpdateSQLCount 通过rollback_update_sql_count获取内容  
func (obj *_V$partitionAuditMgr) GetFromRollbackUpdateSQLCount(rollbackUpdateSQLCount int64) (results []*V$partitionAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$partitionAudit{}).Where("`rollback_update_sql_count` = ?", rollbackUpdateSQLCount).Find(&results).Error
	
	return
}

// GetBatchFromRollbackUpdateSQLCount 批量查找 
func (obj *_V$partitionAuditMgr) GetBatchFromRollbackUpdateSQLCount(rollbackUpdateSQLCounts []int64) (results []*V$partitionAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$partitionAudit{}).Where("`rollback_update_sql_count` IN (?)", rollbackUpdateSQLCounts).Find(&results).Error
	
	return
}
 
// GetFromRollbackTransCount 通过rollback_trans_count获取内容  
func (obj *_V$partitionAuditMgr) GetFromRollbackTransCount(rollbackTransCount int64) (results []*V$partitionAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$partitionAudit{}).Where("`rollback_trans_count` = ?", rollbackTransCount).Find(&results).Error
	
	return
}

// GetBatchFromRollbackTransCount 批量查找 
func (obj *_V$partitionAuditMgr) GetBatchFromRollbackTransCount(rollbackTransCounts []int64) (results []*V$partitionAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$partitionAudit{}).Where("`rollback_trans_count` IN (?)", rollbackTransCounts).Find(&results).Error
	
	return
}
 
// GetFromRollbackSQLCount 通过rollback_sql_count获取内容  
func (obj *_V$partitionAuditMgr) GetFromRollbackSQLCount(rollbackSQLCount int64) (results []*V$partitionAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$partitionAudit{}).Where("`rollback_sql_count` = ?", rollbackSQLCount).Find(&results).Error
	
	return
}

// GetBatchFromRollbackSQLCount 批量查找 
func (obj *_V$partitionAuditMgr) GetBatchFromRollbackSQLCount(rollbackSQLCounts []int64) (results []*V$partitionAudit, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$partitionAudit{}).Where("`rollback_sql_count` IN (?)", rollbackSQLCounts).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 

 

	

