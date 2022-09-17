package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _AllIndexWaitTransactionStatusMgr struct {
	*_BaseMgr
}

// AllIndexWaitTransactionStatusMgr open func
func AllIndexWaitTransactionStatusMgr(db *gorm.DB) *_AllIndexWaitTransactionStatusMgr {
	if db == nil {
		panic(fmt.Errorf("AllIndexWaitTransactionStatusMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllIndexWaitTransactionStatusMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_index_wait_transaction_status"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllIndexWaitTransactionStatusMgr) GetTableName() string {
	return "__all_index_wait_transaction_status"
}

// Reset 重置gorm会话
func (obj *_AllIndexWaitTransactionStatusMgr) Reset() *_AllIndexWaitTransactionStatusMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllIndexWaitTransactionStatusMgr) Get() (result AllIndexWaitTransactionStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllIndexWaitTransactionStatus{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllIndexWaitTransactionStatusMgr) Gets() (results []*AllIndexWaitTransactionStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllIndexWaitTransactionStatus{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllIndexWaitTransactionStatusMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllIndexWaitTransactionStatus{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithGmtCreate gmt_create获取
func (obj *_AllIndexWaitTransactionStatusMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllIndexWaitTransactionStatusMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithTenantID tenant_id获取
func (obj *_AllIndexWaitTransactionStatusMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithIndexTableID index_table_id获取
func (obj *_AllIndexWaitTransactionStatusMgr) WithIndexTableID(indexTableID int64) Option {
	return optionFunc(func(o *options) { o.query["index_table_id"] = indexTableID })
}

// WithSvrType svr_type获取
func (obj *_AllIndexWaitTransactionStatusMgr) WithSvrType(svrType int64) Option {
	return optionFunc(func(o *options) { o.query["svr_type"] = svrType })
}

// WithPartitionID partition_id获取
func (obj *_AllIndexWaitTransactionStatusMgr) WithPartitionID(partitionID int64) Option {
	return optionFunc(func(o *options) { o.query["partition_id"] = partitionID })
}

// WithSvrIP svr_ip获取
func (obj *_AllIndexWaitTransactionStatusMgr) WithSvrIP(svrIP string) Option {
	return optionFunc(func(o *options) { o.query["svr_ip"] = svrIP })
}

// WithSvrPort svr_port获取
func (obj *_AllIndexWaitTransactionStatusMgr) WithSvrPort(svrPort int64) Option {
	return optionFunc(func(o *options) { o.query["svr_port"] = svrPort })
}

// WithTransStatus trans_status获取
func (obj *_AllIndexWaitTransactionStatusMgr) WithTransStatus(transStatus int64) Option {
	return optionFunc(func(o *options) { o.query["trans_status"] = transStatus })
}

// WithSnapshotVersion snapshot_version获取
func (obj *_AllIndexWaitTransactionStatusMgr) WithSnapshotVersion(snapshotVersion int64) Option {
	return optionFunc(func(o *options) { o.query["snapshot_version"] = snapshotVersion })
}

// WithFrozenVersion frozen_version获取
func (obj *_AllIndexWaitTransactionStatusMgr) WithFrozenVersion(frozenVersion int64) Option {
	return optionFunc(func(o *options) { o.query["frozen_version"] = frozenVersion })
}

// WithSchemaVersion schema_version获取
func (obj *_AllIndexWaitTransactionStatusMgr) WithSchemaVersion(schemaVersion int64) Option {
	return optionFunc(func(o *options) { o.query["schema_version"] = schemaVersion })
}

// GetByOption 功能选项模式获取
func (obj *_AllIndexWaitTransactionStatusMgr) GetByOption(opts ...Option) (result AllIndexWaitTransactionStatus, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllIndexWaitTransactionStatus{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllIndexWaitTransactionStatusMgr) GetByOptions(opts ...Option) (results []*AllIndexWaitTransactionStatus, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllIndexWaitTransactionStatus{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllIndexWaitTransactionStatusMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllIndexWaitTransactionStatus, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllIndexWaitTransactionStatus{}).Where(options.query)
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
func (obj *_AllIndexWaitTransactionStatusMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllIndexWaitTransactionStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllIndexWaitTransactionStatus{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllIndexWaitTransactionStatusMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllIndexWaitTransactionStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllIndexWaitTransactionStatus{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllIndexWaitTransactionStatusMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllIndexWaitTransactionStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllIndexWaitTransactionStatus{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllIndexWaitTransactionStatusMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllIndexWaitTransactionStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllIndexWaitTransactionStatus{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromTenantID 通过tenant_id获取内容
func (obj *_AllIndexWaitTransactionStatusMgr) GetFromTenantID(tenantID int64) (results []*AllIndexWaitTransactionStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllIndexWaitTransactionStatus{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllIndexWaitTransactionStatusMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllIndexWaitTransactionStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllIndexWaitTransactionStatus{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromIndexTableID 通过index_table_id获取内容
func (obj *_AllIndexWaitTransactionStatusMgr) GetFromIndexTableID(indexTableID int64) (results []*AllIndexWaitTransactionStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllIndexWaitTransactionStatus{}).Where("`index_table_id` = ?", indexTableID).Find(&results).Error

	return
}

// GetBatchFromIndexTableID 批量查找
func (obj *_AllIndexWaitTransactionStatusMgr) GetBatchFromIndexTableID(indexTableIDs []int64) (results []*AllIndexWaitTransactionStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllIndexWaitTransactionStatus{}).Where("`index_table_id` IN (?)", indexTableIDs).Find(&results).Error

	return
}

// GetFromSvrType 通过svr_type获取内容
func (obj *_AllIndexWaitTransactionStatusMgr) GetFromSvrType(svrType int64) (results []*AllIndexWaitTransactionStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllIndexWaitTransactionStatus{}).Where("`svr_type` = ?", svrType).Find(&results).Error

	return
}

// GetBatchFromSvrType 批量查找
func (obj *_AllIndexWaitTransactionStatusMgr) GetBatchFromSvrType(svrTypes []int64) (results []*AllIndexWaitTransactionStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllIndexWaitTransactionStatus{}).Where("`svr_type` IN (?)", svrTypes).Find(&results).Error

	return
}

// GetFromPartitionID 通过partition_id获取内容
func (obj *_AllIndexWaitTransactionStatusMgr) GetFromPartitionID(partitionID int64) (results []*AllIndexWaitTransactionStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllIndexWaitTransactionStatus{}).Where("`partition_id` = ?", partitionID).Find(&results).Error

	return
}

// GetBatchFromPartitionID 批量查找
func (obj *_AllIndexWaitTransactionStatusMgr) GetBatchFromPartitionID(partitionIDs []int64) (results []*AllIndexWaitTransactionStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllIndexWaitTransactionStatus{}).Where("`partition_id` IN (?)", partitionIDs).Find(&results).Error

	return
}

// GetFromSvrIP 通过svr_ip获取内容
func (obj *_AllIndexWaitTransactionStatusMgr) GetFromSvrIP(svrIP string) (results []*AllIndexWaitTransactionStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllIndexWaitTransactionStatus{}).Where("`svr_ip` = ?", svrIP).Find(&results).Error

	return
}

// GetBatchFromSvrIP 批量查找
func (obj *_AllIndexWaitTransactionStatusMgr) GetBatchFromSvrIP(svrIPs []string) (results []*AllIndexWaitTransactionStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllIndexWaitTransactionStatus{}).Where("`svr_ip` IN (?)", svrIPs).Find(&results).Error

	return
}

// GetFromSvrPort 通过svr_port获取内容
func (obj *_AllIndexWaitTransactionStatusMgr) GetFromSvrPort(svrPort int64) (results []*AllIndexWaitTransactionStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllIndexWaitTransactionStatus{}).Where("`svr_port` = ?", svrPort).Find(&results).Error

	return
}

// GetBatchFromSvrPort 批量查找
func (obj *_AllIndexWaitTransactionStatusMgr) GetBatchFromSvrPort(svrPorts []int64) (results []*AllIndexWaitTransactionStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllIndexWaitTransactionStatus{}).Where("`svr_port` IN (?)", svrPorts).Find(&results).Error

	return
}

// GetFromTransStatus 通过trans_status获取内容
func (obj *_AllIndexWaitTransactionStatusMgr) GetFromTransStatus(transStatus int64) (results []*AllIndexWaitTransactionStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllIndexWaitTransactionStatus{}).Where("`trans_status` = ?", transStatus).Find(&results).Error

	return
}

// GetBatchFromTransStatus 批量查找
func (obj *_AllIndexWaitTransactionStatusMgr) GetBatchFromTransStatus(transStatuss []int64) (results []*AllIndexWaitTransactionStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllIndexWaitTransactionStatus{}).Where("`trans_status` IN (?)", transStatuss).Find(&results).Error

	return
}

// GetFromSnapshotVersion 通过snapshot_version获取内容
func (obj *_AllIndexWaitTransactionStatusMgr) GetFromSnapshotVersion(snapshotVersion int64) (results []*AllIndexWaitTransactionStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllIndexWaitTransactionStatus{}).Where("`snapshot_version` = ?", snapshotVersion).Find(&results).Error

	return
}

// GetBatchFromSnapshotVersion 批量查找
func (obj *_AllIndexWaitTransactionStatusMgr) GetBatchFromSnapshotVersion(snapshotVersions []int64) (results []*AllIndexWaitTransactionStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllIndexWaitTransactionStatus{}).Where("`snapshot_version` IN (?)", snapshotVersions).Find(&results).Error

	return
}

// GetFromFrozenVersion 通过frozen_version获取内容
func (obj *_AllIndexWaitTransactionStatusMgr) GetFromFrozenVersion(frozenVersion int64) (results []*AllIndexWaitTransactionStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllIndexWaitTransactionStatus{}).Where("`frozen_version` = ?", frozenVersion).Find(&results).Error

	return
}

// GetBatchFromFrozenVersion 批量查找
func (obj *_AllIndexWaitTransactionStatusMgr) GetBatchFromFrozenVersion(frozenVersions []int64) (results []*AllIndexWaitTransactionStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllIndexWaitTransactionStatus{}).Where("`frozen_version` IN (?)", frozenVersions).Find(&results).Error

	return
}

// GetFromSchemaVersion 通过schema_version获取内容
func (obj *_AllIndexWaitTransactionStatusMgr) GetFromSchemaVersion(schemaVersion int64) (results []*AllIndexWaitTransactionStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllIndexWaitTransactionStatus{}).Where("`schema_version` = ?", schemaVersion).Find(&results).Error

	return
}

// GetBatchFromSchemaVersion 批量查找
func (obj *_AllIndexWaitTransactionStatusMgr) GetBatchFromSchemaVersion(schemaVersions []int64) (results []*AllIndexWaitTransactionStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllIndexWaitTransactionStatus{}).Where("`schema_version` IN (?)", schemaVersions).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllIndexWaitTransactionStatusMgr) FetchByPrimaryKey(tenantID int64, indexTableID int64, svrType int64, partitionID int64) (result AllIndexWaitTransactionStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllIndexWaitTransactionStatus{}).Where("`tenant_id` = ? AND `index_table_id` = ? AND `svr_type` = ? AND `partition_id` = ?", tenantID, indexTableID, svrType, partitionID).First(&result).Error

	return
}
