package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _AllVirtualSstableColumnChecksumMgr struct {
	*_BaseMgr
}

// AllVirtualSstableColumnChecksumMgr open func
func AllVirtualSstableColumnChecksumMgr(db *gorm.DB) *_AllVirtualSstableColumnChecksumMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualSstableColumnChecksumMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualSstableColumnChecksumMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_sstable_column_checksum"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualSstableColumnChecksumMgr) GetTableName() string {
	return "__all_virtual_sstable_column_checksum"
}

// Reset 重置gorm会话
func (obj *_AllVirtualSstableColumnChecksumMgr) Reset() *_AllVirtualSstableColumnChecksumMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualSstableColumnChecksumMgr) Get() (result AllVirtualSstableColumnChecksum, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSstableColumnChecksum{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualSstableColumnChecksumMgr) Gets() (results []*AllVirtualSstableColumnChecksum, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSstableColumnChecksum{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualSstableColumnChecksumMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualSstableColumnChecksum{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTenantID tenant_id获取
func (obj *_AllVirtualSstableColumnChecksumMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithDataTableID data_table_id获取
func (obj *_AllVirtualSstableColumnChecksumMgr) WithDataTableID(dataTableID int64) Option {
	return optionFunc(func(o *options) { o.query["data_table_id"] = dataTableID })
}

// WithIndexID index_id获取
func (obj *_AllVirtualSstableColumnChecksumMgr) WithIndexID(indexID int64) Option {
	return optionFunc(func(o *options) { o.query["index_id"] = indexID })
}

// WithPartitionID partition_id获取
func (obj *_AllVirtualSstableColumnChecksumMgr) WithPartitionID(partitionID int64) Option {
	return optionFunc(func(o *options) { o.query["partition_id"] = partitionID })
}

// WithSstableType sstable_type获取
func (obj *_AllVirtualSstableColumnChecksumMgr) WithSstableType(sstableType int64) Option {
	return optionFunc(func(o *options) { o.query["sstable_type"] = sstableType })
}

// WithColumnID column_id获取
func (obj *_AllVirtualSstableColumnChecksumMgr) WithColumnID(columnID int64) Option {
	return optionFunc(func(o *options) { o.query["column_id"] = columnID })
}

// WithSvrIP svr_ip获取
func (obj *_AllVirtualSstableColumnChecksumMgr) WithSvrIP(svrIP string) Option {
	return optionFunc(func(o *options) { o.query["svr_ip"] = svrIP })
}

// WithSvrPort svr_port获取
func (obj *_AllVirtualSstableColumnChecksumMgr) WithSvrPort(svrPort int64) Option {
	return optionFunc(func(o *options) { o.query["svr_port"] = svrPort })
}

// WithGmtCreate gmt_create获取
func (obj *_AllVirtualSstableColumnChecksumMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllVirtualSstableColumnChecksumMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithColumnChecksum column_checksum获取
func (obj *_AllVirtualSstableColumnChecksumMgr) WithColumnChecksum(columnChecksum int64) Option {
	return optionFunc(func(o *options) { o.query["column_checksum"] = columnChecksum })
}

// WithChecksumMethod checksum_method获取
func (obj *_AllVirtualSstableColumnChecksumMgr) WithChecksumMethod(checksumMethod int64) Option {
	return optionFunc(func(o *options) { o.query["checksum_method"] = checksumMethod })
}

// WithSnapshotVersion snapshot_version获取
func (obj *_AllVirtualSstableColumnChecksumMgr) WithSnapshotVersion(snapshotVersion int64) Option {
	return optionFunc(func(o *options) { o.query["snapshot_version"] = snapshotVersion })
}

// WithReplicaType replica_type获取
func (obj *_AllVirtualSstableColumnChecksumMgr) WithReplicaType(replicaType int64) Option {
	return optionFunc(func(o *options) { o.query["replica_type"] = replicaType })
}

// WithMajorVersion major_version获取
func (obj *_AllVirtualSstableColumnChecksumMgr) WithMajorVersion(majorVersion int64) Option {
	return optionFunc(func(o *options) { o.query["major_version"] = majorVersion })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualSstableColumnChecksumMgr) GetByOption(opts ...Option) (result AllVirtualSstableColumnChecksum, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSstableColumnChecksum{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualSstableColumnChecksumMgr) GetByOptions(opts ...Option) (results []*AllVirtualSstableColumnChecksum, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSstableColumnChecksum{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualSstableColumnChecksumMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualSstableColumnChecksum, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualSstableColumnChecksum{}).Where(options.query)
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
func (obj *_AllVirtualSstableColumnChecksumMgr) GetFromTenantID(tenantID int64) (results []*AllVirtualSstableColumnChecksum, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSstableColumnChecksum{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllVirtualSstableColumnChecksumMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllVirtualSstableColumnChecksum, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSstableColumnChecksum{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromDataTableID 通过data_table_id获取内容
func (obj *_AllVirtualSstableColumnChecksumMgr) GetFromDataTableID(dataTableID int64) (results []*AllVirtualSstableColumnChecksum, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSstableColumnChecksum{}).Where("`data_table_id` = ?", dataTableID).Find(&results).Error

	return
}

// GetBatchFromDataTableID 批量查找
func (obj *_AllVirtualSstableColumnChecksumMgr) GetBatchFromDataTableID(dataTableIDs []int64) (results []*AllVirtualSstableColumnChecksum, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSstableColumnChecksum{}).Where("`data_table_id` IN (?)", dataTableIDs).Find(&results).Error

	return
}

// GetFromIndexID 通过index_id获取内容
func (obj *_AllVirtualSstableColumnChecksumMgr) GetFromIndexID(indexID int64) (results []*AllVirtualSstableColumnChecksum, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSstableColumnChecksum{}).Where("`index_id` = ?", indexID).Find(&results).Error

	return
}

// GetBatchFromIndexID 批量查找
func (obj *_AllVirtualSstableColumnChecksumMgr) GetBatchFromIndexID(indexIDs []int64) (results []*AllVirtualSstableColumnChecksum, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSstableColumnChecksum{}).Where("`index_id` IN (?)", indexIDs).Find(&results).Error

	return
}

// GetFromPartitionID 通过partition_id获取内容
func (obj *_AllVirtualSstableColumnChecksumMgr) GetFromPartitionID(partitionID int64) (results []*AllVirtualSstableColumnChecksum, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSstableColumnChecksum{}).Where("`partition_id` = ?", partitionID).Find(&results).Error

	return
}

// GetBatchFromPartitionID 批量查找
func (obj *_AllVirtualSstableColumnChecksumMgr) GetBatchFromPartitionID(partitionIDs []int64) (results []*AllVirtualSstableColumnChecksum, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSstableColumnChecksum{}).Where("`partition_id` IN (?)", partitionIDs).Find(&results).Error

	return
}

// GetFromSstableType 通过sstable_type获取内容
func (obj *_AllVirtualSstableColumnChecksumMgr) GetFromSstableType(sstableType int64) (results []*AllVirtualSstableColumnChecksum, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSstableColumnChecksum{}).Where("`sstable_type` = ?", sstableType).Find(&results).Error

	return
}

// GetBatchFromSstableType 批量查找
func (obj *_AllVirtualSstableColumnChecksumMgr) GetBatchFromSstableType(sstableTypes []int64) (results []*AllVirtualSstableColumnChecksum, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSstableColumnChecksum{}).Where("`sstable_type` IN (?)", sstableTypes).Find(&results).Error

	return
}

// GetFromColumnID 通过column_id获取内容
func (obj *_AllVirtualSstableColumnChecksumMgr) GetFromColumnID(columnID int64) (results []*AllVirtualSstableColumnChecksum, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSstableColumnChecksum{}).Where("`column_id` = ?", columnID).Find(&results).Error

	return
}

// GetBatchFromColumnID 批量查找
func (obj *_AllVirtualSstableColumnChecksumMgr) GetBatchFromColumnID(columnIDs []int64) (results []*AllVirtualSstableColumnChecksum, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSstableColumnChecksum{}).Where("`column_id` IN (?)", columnIDs).Find(&results).Error

	return
}

// GetFromSvrIP 通过svr_ip获取内容
func (obj *_AllVirtualSstableColumnChecksumMgr) GetFromSvrIP(svrIP string) (results []*AllVirtualSstableColumnChecksum, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSstableColumnChecksum{}).Where("`svr_ip` = ?", svrIP).Find(&results).Error

	return
}

// GetBatchFromSvrIP 批量查找
func (obj *_AllVirtualSstableColumnChecksumMgr) GetBatchFromSvrIP(svrIPs []string) (results []*AllVirtualSstableColumnChecksum, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSstableColumnChecksum{}).Where("`svr_ip` IN (?)", svrIPs).Find(&results).Error

	return
}

// GetFromSvrPort 通过svr_port获取内容
func (obj *_AllVirtualSstableColumnChecksumMgr) GetFromSvrPort(svrPort int64) (results []*AllVirtualSstableColumnChecksum, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSstableColumnChecksum{}).Where("`svr_port` = ?", svrPort).Find(&results).Error

	return
}

// GetBatchFromSvrPort 批量查找
func (obj *_AllVirtualSstableColumnChecksumMgr) GetBatchFromSvrPort(svrPorts []int64) (results []*AllVirtualSstableColumnChecksum, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSstableColumnChecksum{}).Where("`svr_port` IN (?)", svrPorts).Find(&results).Error

	return
}

// GetFromGmtCreate 通过gmt_create获取内容
func (obj *_AllVirtualSstableColumnChecksumMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllVirtualSstableColumnChecksum, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSstableColumnChecksum{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllVirtualSstableColumnChecksumMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllVirtualSstableColumnChecksum, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSstableColumnChecksum{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllVirtualSstableColumnChecksumMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllVirtualSstableColumnChecksum, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSstableColumnChecksum{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllVirtualSstableColumnChecksumMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllVirtualSstableColumnChecksum, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSstableColumnChecksum{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromColumnChecksum 通过column_checksum获取内容
func (obj *_AllVirtualSstableColumnChecksumMgr) GetFromColumnChecksum(columnChecksum int64) (results []*AllVirtualSstableColumnChecksum, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSstableColumnChecksum{}).Where("`column_checksum` = ?", columnChecksum).Find(&results).Error

	return
}

// GetBatchFromColumnChecksum 批量查找
func (obj *_AllVirtualSstableColumnChecksumMgr) GetBatchFromColumnChecksum(columnChecksums []int64) (results []*AllVirtualSstableColumnChecksum, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSstableColumnChecksum{}).Where("`column_checksum` IN (?)", columnChecksums).Find(&results).Error

	return
}

// GetFromChecksumMethod 通过checksum_method获取内容
func (obj *_AllVirtualSstableColumnChecksumMgr) GetFromChecksumMethod(checksumMethod int64) (results []*AllVirtualSstableColumnChecksum, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSstableColumnChecksum{}).Where("`checksum_method` = ?", checksumMethod).Find(&results).Error

	return
}

// GetBatchFromChecksumMethod 批量查找
func (obj *_AllVirtualSstableColumnChecksumMgr) GetBatchFromChecksumMethod(checksumMethods []int64) (results []*AllVirtualSstableColumnChecksum, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSstableColumnChecksum{}).Where("`checksum_method` IN (?)", checksumMethods).Find(&results).Error

	return
}

// GetFromSnapshotVersion 通过snapshot_version获取内容
func (obj *_AllVirtualSstableColumnChecksumMgr) GetFromSnapshotVersion(snapshotVersion int64) (results []*AllVirtualSstableColumnChecksum, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSstableColumnChecksum{}).Where("`snapshot_version` = ?", snapshotVersion).Find(&results).Error

	return
}

// GetBatchFromSnapshotVersion 批量查找
func (obj *_AllVirtualSstableColumnChecksumMgr) GetBatchFromSnapshotVersion(snapshotVersions []int64) (results []*AllVirtualSstableColumnChecksum, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSstableColumnChecksum{}).Where("`snapshot_version` IN (?)", snapshotVersions).Find(&results).Error

	return
}

// GetFromReplicaType 通过replica_type获取内容
func (obj *_AllVirtualSstableColumnChecksumMgr) GetFromReplicaType(replicaType int64) (results []*AllVirtualSstableColumnChecksum, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSstableColumnChecksum{}).Where("`replica_type` = ?", replicaType).Find(&results).Error

	return
}

// GetBatchFromReplicaType 批量查找
func (obj *_AllVirtualSstableColumnChecksumMgr) GetBatchFromReplicaType(replicaTypes []int64) (results []*AllVirtualSstableColumnChecksum, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSstableColumnChecksum{}).Where("`replica_type` IN (?)", replicaTypes).Find(&results).Error

	return
}

// GetFromMajorVersion 通过major_version获取内容
func (obj *_AllVirtualSstableColumnChecksumMgr) GetFromMajorVersion(majorVersion int64) (results []*AllVirtualSstableColumnChecksum, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSstableColumnChecksum{}).Where("`major_version` = ?", majorVersion).Find(&results).Error

	return
}

// GetBatchFromMajorVersion 批量查找
func (obj *_AllVirtualSstableColumnChecksumMgr) GetBatchFromMajorVersion(majorVersions []int64) (results []*AllVirtualSstableColumnChecksum, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSstableColumnChecksum{}).Where("`major_version` IN (?)", majorVersions).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllVirtualSstableColumnChecksumMgr) FetchByPrimaryKey(tenantID int64, dataTableID int64, indexID int64, partitionID int64, sstableType int64, columnID int64, svrIP string, svrPort int64) (result AllVirtualSstableColumnChecksum, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSstableColumnChecksum{}).Where("`tenant_id` = ? AND `data_table_id` = ? AND `index_id` = ? AND `partition_id` = ? AND `sstable_type` = ? AND `column_id` = ? AND `svr_ip` = ? AND `svr_port` = ?", tenantID, dataTableID, indexID, partitionID, sstableType, columnID, svrIP, svrPort).First(&result).Error

	return
}
