package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _AllSstableChecksumMgr struct {
	*_BaseMgr
}

// AllSstableChecksumMgr open func
func AllSstableChecksumMgr(db *gorm.DB) *_AllSstableChecksumMgr {
	if db == nil {
		panic(fmt.Errorf("AllSstableChecksumMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllSstableChecksumMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_sstable_checksum"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllSstableChecksumMgr) GetTableName() string {
	return "__all_sstable_checksum"
}

// Reset 重置gorm会话
func (obj *_AllSstableChecksumMgr) Reset() *_AllSstableChecksumMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllSstableChecksumMgr) Get() (result AllSstableChecksum, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSstableChecksum{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllSstableChecksumMgr) Gets() (results []*AllSstableChecksum, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSstableChecksum{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllSstableChecksumMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllSstableChecksum{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithGmtCreate gmt_create获取
func (obj *_AllSstableChecksumMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllSstableChecksumMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithTenantID tenant_id获取
func (obj *_AllSstableChecksumMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithDataTableID data_table_id获取
func (obj *_AllSstableChecksumMgr) WithDataTableID(dataTableID int64) Option {
	return optionFunc(func(o *options) { o.query["data_table_id"] = dataTableID })
}

// WithSstableID sstable_id获取
func (obj *_AllSstableChecksumMgr) WithSstableID(sstableID int64) Option {
	return optionFunc(func(o *options) { o.query["sstable_id"] = sstableID })
}

// WithPartitionID partition_id获取
func (obj *_AllSstableChecksumMgr) WithPartitionID(partitionID int64) Option {
	return optionFunc(func(o *options) { o.query["partition_id"] = partitionID })
}

// WithSstableType sstable_type获取
func (obj *_AllSstableChecksumMgr) WithSstableType(sstableType int64) Option {
	return optionFunc(func(o *options) { o.query["sstable_type"] = sstableType })
}

// WithSvrIP svr_ip获取
func (obj *_AllSstableChecksumMgr) WithSvrIP(svrIP string) Option {
	return optionFunc(func(o *options) { o.query["svr_ip"] = svrIP })
}

// WithSvrPort svr_port获取
func (obj *_AllSstableChecksumMgr) WithSvrPort(svrPort int64) Option {
	return optionFunc(func(o *options) { o.query["svr_port"] = svrPort })
}

// WithRowChecksum row_checksum获取
func (obj *_AllSstableChecksumMgr) WithRowChecksum(rowChecksum int64) Option {
	return optionFunc(func(o *options) { o.query["row_checksum"] = rowChecksum })
}

// WithDataChecksum data_checksum获取
func (obj *_AllSstableChecksumMgr) WithDataChecksum(dataChecksum int64) Option {
	return optionFunc(func(o *options) { o.query["data_checksum"] = dataChecksum })
}

// WithRowCount row_count获取
func (obj *_AllSstableChecksumMgr) WithRowCount(rowCount int64) Option {
	return optionFunc(func(o *options) { o.query["row_count"] = rowCount })
}

// WithSnapshotVersion snapshot_version获取
func (obj *_AllSstableChecksumMgr) WithSnapshotVersion(snapshotVersion int64) Option {
	return optionFunc(func(o *options) { o.query["snapshot_version"] = snapshotVersion })
}

// WithReplicaType replica_type获取
func (obj *_AllSstableChecksumMgr) WithReplicaType(replicaType int64) Option {
	return optionFunc(func(o *options) { o.query["replica_type"] = replicaType })
}

// GetByOption 功能选项模式获取
func (obj *_AllSstableChecksumMgr) GetByOption(opts ...Option) (result AllSstableChecksum, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllSstableChecksum{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllSstableChecksumMgr) GetByOptions(opts ...Option) (results []*AllSstableChecksum, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllSstableChecksum{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllSstableChecksumMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllSstableChecksum, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllSstableChecksum{}).Where(options.query)
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
func (obj *_AllSstableChecksumMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllSstableChecksum, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSstableChecksum{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllSstableChecksumMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllSstableChecksum, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSstableChecksum{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllSstableChecksumMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllSstableChecksum, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSstableChecksum{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllSstableChecksumMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllSstableChecksum, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSstableChecksum{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromTenantID 通过tenant_id获取内容
func (obj *_AllSstableChecksumMgr) GetFromTenantID(tenantID int64) (results []*AllSstableChecksum, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSstableChecksum{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllSstableChecksumMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllSstableChecksum, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSstableChecksum{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromDataTableID 通过data_table_id获取内容
func (obj *_AllSstableChecksumMgr) GetFromDataTableID(dataTableID int64) (results []*AllSstableChecksum, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSstableChecksum{}).Where("`data_table_id` = ?", dataTableID).Find(&results).Error

	return
}

// GetBatchFromDataTableID 批量查找
func (obj *_AllSstableChecksumMgr) GetBatchFromDataTableID(dataTableIDs []int64) (results []*AllSstableChecksum, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSstableChecksum{}).Where("`data_table_id` IN (?)", dataTableIDs).Find(&results).Error

	return
}

// GetFromSstableID 通过sstable_id获取内容
func (obj *_AllSstableChecksumMgr) GetFromSstableID(sstableID int64) (results []*AllSstableChecksum, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSstableChecksum{}).Where("`sstable_id` = ?", sstableID).Find(&results).Error

	return
}

// GetBatchFromSstableID 批量查找
func (obj *_AllSstableChecksumMgr) GetBatchFromSstableID(sstableIDs []int64) (results []*AllSstableChecksum, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSstableChecksum{}).Where("`sstable_id` IN (?)", sstableIDs).Find(&results).Error

	return
}

// GetFromPartitionID 通过partition_id获取内容
func (obj *_AllSstableChecksumMgr) GetFromPartitionID(partitionID int64) (results []*AllSstableChecksum, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSstableChecksum{}).Where("`partition_id` = ?", partitionID).Find(&results).Error

	return
}

// GetBatchFromPartitionID 批量查找
func (obj *_AllSstableChecksumMgr) GetBatchFromPartitionID(partitionIDs []int64) (results []*AllSstableChecksum, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSstableChecksum{}).Where("`partition_id` IN (?)", partitionIDs).Find(&results).Error

	return
}

// GetFromSstableType 通过sstable_type获取内容
func (obj *_AllSstableChecksumMgr) GetFromSstableType(sstableType int64) (results []*AllSstableChecksum, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSstableChecksum{}).Where("`sstable_type` = ?", sstableType).Find(&results).Error

	return
}

// GetBatchFromSstableType 批量查找
func (obj *_AllSstableChecksumMgr) GetBatchFromSstableType(sstableTypes []int64) (results []*AllSstableChecksum, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSstableChecksum{}).Where("`sstable_type` IN (?)", sstableTypes).Find(&results).Error

	return
}

// GetFromSvrIP 通过svr_ip获取内容
func (obj *_AllSstableChecksumMgr) GetFromSvrIP(svrIP string) (results []*AllSstableChecksum, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSstableChecksum{}).Where("`svr_ip` = ?", svrIP).Find(&results).Error

	return
}

// GetBatchFromSvrIP 批量查找
func (obj *_AllSstableChecksumMgr) GetBatchFromSvrIP(svrIPs []string) (results []*AllSstableChecksum, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSstableChecksum{}).Where("`svr_ip` IN (?)", svrIPs).Find(&results).Error

	return
}

// GetFromSvrPort 通过svr_port获取内容
func (obj *_AllSstableChecksumMgr) GetFromSvrPort(svrPort int64) (results []*AllSstableChecksum, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSstableChecksum{}).Where("`svr_port` = ?", svrPort).Find(&results).Error

	return
}

// GetBatchFromSvrPort 批量查找
func (obj *_AllSstableChecksumMgr) GetBatchFromSvrPort(svrPorts []int64) (results []*AllSstableChecksum, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSstableChecksum{}).Where("`svr_port` IN (?)", svrPorts).Find(&results).Error

	return
}

// GetFromRowChecksum 通过row_checksum获取内容
func (obj *_AllSstableChecksumMgr) GetFromRowChecksum(rowChecksum int64) (results []*AllSstableChecksum, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSstableChecksum{}).Where("`row_checksum` = ?", rowChecksum).Find(&results).Error

	return
}

// GetBatchFromRowChecksum 批量查找
func (obj *_AllSstableChecksumMgr) GetBatchFromRowChecksum(rowChecksums []int64) (results []*AllSstableChecksum, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSstableChecksum{}).Where("`row_checksum` IN (?)", rowChecksums).Find(&results).Error

	return
}

// GetFromDataChecksum 通过data_checksum获取内容
func (obj *_AllSstableChecksumMgr) GetFromDataChecksum(dataChecksum int64) (results []*AllSstableChecksum, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSstableChecksum{}).Where("`data_checksum` = ?", dataChecksum).Find(&results).Error

	return
}

// GetBatchFromDataChecksum 批量查找
func (obj *_AllSstableChecksumMgr) GetBatchFromDataChecksum(dataChecksums []int64) (results []*AllSstableChecksum, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSstableChecksum{}).Where("`data_checksum` IN (?)", dataChecksums).Find(&results).Error

	return
}

// GetFromRowCount 通过row_count获取内容
func (obj *_AllSstableChecksumMgr) GetFromRowCount(rowCount int64) (results []*AllSstableChecksum, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSstableChecksum{}).Where("`row_count` = ?", rowCount).Find(&results).Error

	return
}

// GetBatchFromRowCount 批量查找
func (obj *_AllSstableChecksumMgr) GetBatchFromRowCount(rowCounts []int64) (results []*AllSstableChecksum, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSstableChecksum{}).Where("`row_count` IN (?)", rowCounts).Find(&results).Error

	return
}

// GetFromSnapshotVersion 通过snapshot_version获取内容
func (obj *_AllSstableChecksumMgr) GetFromSnapshotVersion(snapshotVersion int64) (results []*AllSstableChecksum, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSstableChecksum{}).Where("`snapshot_version` = ?", snapshotVersion).Find(&results).Error

	return
}

// GetBatchFromSnapshotVersion 批量查找
func (obj *_AllSstableChecksumMgr) GetBatchFromSnapshotVersion(snapshotVersions []int64) (results []*AllSstableChecksum, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSstableChecksum{}).Where("`snapshot_version` IN (?)", snapshotVersions).Find(&results).Error

	return
}

// GetFromReplicaType 通过replica_type获取内容
func (obj *_AllSstableChecksumMgr) GetFromReplicaType(replicaType int64) (results []*AllSstableChecksum, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSstableChecksum{}).Where("`replica_type` = ?", replicaType).Find(&results).Error

	return
}

// GetBatchFromReplicaType 批量查找
func (obj *_AllSstableChecksumMgr) GetBatchFromReplicaType(replicaTypes []int64) (results []*AllSstableChecksum, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSstableChecksum{}).Where("`replica_type` IN (?)", replicaTypes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllSstableChecksumMgr) FetchByPrimaryKey(tenantID int64, dataTableID int64, sstableID int64, partitionID int64, sstableType int64, svrIP string, svrPort int64) (result AllSstableChecksum, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSstableChecksum{}).Where("`tenant_id` = ? AND `data_table_id` = ? AND `sstable_id` = ? AND `partition_id` = ? AND `sstable_type` = ? AND `svr_ip` = ? AND `svr_port` = ?", tenantID, dataTableID, sstableID, partitionID, sstableType, svrIP, svrPort).First(&result).Error

	return
}
