package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _AllTenantSstableColumnChecksumMgr struct {
	*_BaseMgr
}

// AllTenantSstableColumnChecksumMgr open func
func AllTenantSstableColumnChecksumMgr(db *gorm.DB) *_AllTenantSstableColumnChecksumMgr {
	if db == nil {
		panic(fmt.Errorf("AllTenantSstableColumnChecksumMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllTenantSstableColumnChecksumMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_tenant_sstable_column_checksum"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllTenantSstableColumnChecksumMgr) GetTableName() string {
	return "__all_tenant_sstable_column_checksum"
}

// Reset 重置gorm会话
func (obj *_AllTenantSstableColumnChecksumMgr) Reset() *_AllTenantSstableColumnChecksumMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllTenantSstableColumnChecksumMgr) Get() (result AllTenantSstableColumnChecksum, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantSstableColumnChecksum{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllTenantSstableColumnChecksumMgr) Gets() (results []*AllTenantSstableColumnChecksum, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantSstableColumnChecksum{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllTenantSstableColumnChecksumMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllTenantSstableColumnChecksum{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithGmtCreate gmt_create获取
func (obj *_AllTenantSstableColumnChecksumMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllTenantSstableColumnChecksumMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithTenantID tenant_id获取
func (obj *_AllTenantSstableColumnChecksumMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithDataTableID data_table_id获取
func (obj *_AllTenantSstableColumnChecksumMgr) WithDataTableID(dataTableID int64) Option {
	return optionFunc(func(o *options) { o.query["data_table_id"] = dataTableID })
}

// WithIndexID index_id获取
func (obj *_AllTenantSstableColumnChecksumMgr) WithIndexID(indexID int64) Option {
	return optionFunc(func(o *options) { o.query["index_id"] = indexID })
}

// WithPartitionID partition_id获取
func (obj *_AllTenantSstableColumnChecksumMgr) WithPartitionID(partitionID int64) Option {
	return optionFunc(func(o *options) { o.query["partition_id"] = partitionID })
}

// WithSstableType sstable_type获取
func (obj *_AllTenantSstableColumnChecksumMgr) WithSstableType(sstableType int64) Option {
	return optionFunc(func(o *options) { o.query["sstable_type"] = sstableType })
}

// WithColumnID column_id获取
func (obj *_AllTenantSstableColumnChecksumMgr) WithColumnID(columnID int64) Option {
	return optionFunc(func(o *options) { o.query["column_id"] = columnID })
}

// WithSvrIP svr_ip获取
func (obj *_AllTenantSstableColumnChecksumMgr) WithSvrIP(svrIP string) Option {
	return optionFunc(func(o *options) { o.query["svr_ip"] = svrIP })
}

// WithSvrPort svr_port获取
func (obj *_AllTenantSstableColumnChecksumMgr) WithSvrPort(svrPort int64) Option {
	return optionFunc(func(o *options) { o.query["svr_port"] = svrPort })
}

// WithColumnChecksum column_checksum获取
func (obj *_AllTenantSstableColumnChecksumMgr) WithColumnChecksum(columnChecksum int64) Option {
	return optionFunc(func(o *options) { o.query["column_checksum"] = columnChecksum })
}

// WithChecksumMethod checksum_method获取
func (obj *_AllTenantSstableColumnChecksumMgr) WithChecksumMethod(checksumMethod int64) Option {
	return optionFunc(func(o *options) { o.query["checksum_method"] = checksumMethod })
}

// WithSnapshotVersion snapshot_version获取
func (obj *_AllTenantSstableColumnChecksumMgr) WithSnapshotVersion(snapshotVersion int64) Option {
	return optionFunc(func(o *options) { o.query["snapshot_version"] = snapshotVersion })
}

// WithReplicaType replica_type获取
func (obj *_AllTenantSstableColumnChecksumMgr) WithReplicaType(replicaType int64) Option {
	return optionFunc(func(o *options) { o.query["replica_type"] = replicaType })
}

// WithMajorVersion major_version获取
func (obj *_AllTenantSstableColumnChecksumMgr) WithMajorVersion(majorVersion int64) Option {
	return optionFunc(func(o *options) { o.query["major_version"] = majorVersion })
}

// GetByOption 功能选项模式获取
func (obj *_AllTenantSstableColumnChecksumMgr) GetByOption(opts ...Option) (result AllTenantSstableColumnChecksum, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllTenantSstableColumnChecksum{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllTenantSstableColumnChecksumMgr) GetByOptions(opts ...Option) (results []*AllTenantSstableColumnChecksum, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllTenantSstableColumnChecksum{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllTenantSstableColumnChecksumMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllTenantSstableColumnChecksum, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllTenantSstableColumnChecksum{}).Where(options.query)
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
func (obj *_AllTenantSstableColumnChecksumMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllTenantSstableColumnChecksum, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantSstableColumnChecksum{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllTenantSstableColumnChecksumMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllTenantSstableColumnChecksum, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantSstableColumnChecksum{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllTenantSstableColumnChecksumMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllTenantSstableColumnChecksum, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantSstableColumnChecksum{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllTenantSstableColumnChecksumMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllTenantSstableColumnChecksum, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantSstableColumnChecksum{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromTenantID 通过tenant_id获取内容
func (obj *_AllTenantSstableColumnChecksumMgr) GetFromTenantID(tenantID int64) (results []*AllTenantSstableColumnChecksum, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantSstableColumnChecksum{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllTenantSstableColumnChecksumMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllTenantSstableColumnChecksum, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantSstableColumnChecksum{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromDataTableID 通过data_table_id获取内容
func (obj *_AllTenantSstableColumnChecksumMgr) GetFromDataTableID(dataTableID int64) (results []*AllTenantSstableColumnChecksum, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantSstableColumnChecksum{}).Where("`data_table_id` = ?", dataTableID).Find(&results).Error

	return
}

// GetBatchFromDataTableID 批量查找
func (obj *_AllTenantSstableColumnChecksumMgr) GetBatchFromDataTableID(dataTableIDs []int64) (results []*AllTenantSstableColumnChecksum, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantSstableColumnChecksum{}).Where("`data_table_id` IN (?)", dataTableIDs).Find(&results).Error

	return
}

// GetFromIndexID 通过index_id获取内容
func (obj *_AllTenantSstableColumnChecksumMgr) GetFromIndexID(indexID int64) (results []*AllTenantSstableColumnChecksum, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantSstableColumnChecksum{}).Where("`index_id` = ?", indexID).Find(&results).Error

	return
}

// GetBatchFromIndexID 批量查找
func (obj *_AllTenantSstableColumnChecksumMgr) GetBatchFromIndexID(indexIDs []int64) (results []*AllTenantSstableColumnChecksum, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantSstableColumnChecksum{}).Where("`index_id` IN (?)", indexIDs).Find(&results).Error

	return
}

// GetFromPartitionID 通过partition_id获取内容
func (obj *_AllTenantSstableColumnChecksumMgr) GetFromPartitionID(partitionID int64) (results []*AllTenantSstableColumnChecksum, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantSstableColumnChecksum{}).Where("`partition_id` = ?", partitionID).Find(&results).Error

	return
}

// GetBatchFromPartitionID 批量查找
func (obj *_AllTenantSstableColumnChecksumMgr) GetBatchFromPartitionID(partitionIDs []int64) (results []*AllTenantSstableColumnChecksum, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantSstableColumnChecksum{}).Where("`partition_id` IN (?)", partitionIDs).Find(&results).Error

	return
}

// GetFromSstableType 通过sstable_type获取内容
func (obj *_AllTenantSstableColumnChecksumMgr) GetFromSstableType(sstableType int64) (results []*AllTenantSstableColumnChecksum, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantSstableColumnChecksum{}).Where("`sstable_type` = ?", sstableType).Find(&results).Error

	return
}

// GetBatchFromSstableType 批量查找
func (obj *_AllTenantSstableColumnChecksumMgr) GetBatchFromSstableType(sstableTypes []int64) (results []*AllTenantSstableColumnChecksum, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantSstableColumnChecksum{}).Where("`sstable_type` IN (?)", sstableTypes).Find(&results).Error

	return
}

// GetFromColumnID 通过column_id获取内容
func (obj *_AllTenantSstableColumnChecksumMgr) GetFromColumnID(columnID int64) (results []*AllTenantSstableColumnChecksum, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantSstableColumnChecksum{}).Where("`column_id` = ?", columnID).Find(&results).Error

	return
}

// GetBatchFromColumnID 批量查找
func (obj *_AllTenantSstableColumnChecksumMgr) GetBatchFromColumnID(columnIDs []int64) (results []*AllTenantSstableColumnChecksum, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantSstableColumnChecksum{}).Where("`column_id` IN (?)", columnIDs).Find(&results).Error

	return
}

// GetFromSvrIP 通过svr_ip获取内容
func (obj *_AllTenantSstableColumnChecksumMgr) GetFromSvrIP(svrIP string) (results []*AllTenantSstableColumnChecksum, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantSstableColumnChecksum{}).Where("`svr_ip` = ?", svrIP).Find(&results).Error

	return
}

// GetBatchFromSvrIP 批量查找
func (obj *_AllTenantSstableColumnChecksumMgr) GetBatchFromSvrIP(svrIPs []string) (results []*AllTenantSstableColumnChecksum, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantSstableColumnChecksum{}).Where("`svr_ip` IN (?)", svrIPs).Find(&results).Error

	return
}

// GetFromSvrPort 通过svr_port获取内容
func (obj *_AllTenantSstableColumnChecksumMgr) GetFromSvrPort(svrPort int64) (results []*AllTenantSstableColumnChecksum, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantSstableColumnChecksum{}).Where("`svr_port` = ?", svrPort).Find(&results).Error

	return
}

// GetBatchFromSvrPort 批量查找
func (obj *_AllTenantSstableColumnChecksumMgr) GetBatchFromSvrPort(svrPorts []int64) (results []*AllTenantSstableColumnChecksum, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantSstableColumnChecksum{}).Where("`svr_port` IN (?)", svrPorts).Find(&results).Error

	return
}

// GetFromColumnChecksum 通过column_checksum获取内容
func (obj *_AllTenantSstableColumnChecksumMgr) GetFromColumnChecksum(columnChecksum int64) (results []*AllTenantSstableColumnChecksum, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantSstableColumnChecksum{}).Where("`column_checksum` = ?", columnChecksum).Find(&results).Error

	return
}

// GetBatchFromColumnChecksum 批量查找
func (obj *_AllTenantSstableColumnChecksumMgr) GetBatchFromColumnChecksum(columnChecksums []int64) (results []*AllTenantSstableColumnChecksum, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantSstableColumnChecksum{}).Where("`column_checksum` IN (?)", columnChecksums).Find(&results).Error

	return
}

// GetFromChecksumMethod 通过checksum_method获取内容
func (obj *_AllTenantSstableColumnChecksumMgr) GetFromChecksumMethod(checksumMethod int64) (results []*AllTenantSstableColumnChecksum, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantSstableColumnChecksum{}).Where("`checksum_method` = ?", checksumMethod).Find(&results).Error

	return
}

// GetBatchFromChecksumMethod 批量查找
func (obj *_AllTenantSstableColumnChecksumMgr) GetBatchFromChecksumMethod(checksumMethods []int64) (results []*AllTenantSstableColumnChecksum, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantSstableColumnChecksum{}).Where("`checksum_method` IN (?)", checksumMethods).Find(&results).Error

	return
}

// GetFromSnapshotVersion 通过snapshot_version获取内容
func (obj *_AllTenantSstableColumnChecksumMgr) GetFromSnapshotVersion(snapshotVersion int64) (results []*AllTenantSstableColumnChecksum, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantSstableColumnChecksum{}).Where("`snapshot_version` = ?", snapshotVersion).Find(&results).Error

	return
}

// GetBatchFromSnapshotVersion 批量查找
func (obj *_AllTenantSstableColumnChecksumMgr) GetBatchFromSnapshotVersion(snapshotVersions []int64) (results []*AllTenantSstableColumnChecksum, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantSstableColumnChecksum{}).Where("`snapshot_version` IN (?)", snapshotVersions).Find(&results).Error

	return
}

// GetFromReplicaType 通过replica_type获取内容
func (obj *_AllTenantSstableColumnChecksumMgr) GetFromReplicaType(replicaType int64) (results []*AllTenantSstableColumnChecksum, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantSstableColumnChecksum{}).Where("`replica_type` = ?", replicaType).Find(&results).Error

	return
}

// GetBatchFromReplicaType 批量查找
func (obj *_AllTenantSstableColumnChecksumMgr) GetBatchFromReplicaType(replicaTypes []int64) (results []*AllTenantSstableColumnChecksum, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantSstableColumnChecksum{}).Where("`replica_type` IN (?)", replicaTypes).Find(&results).Error

	return
}

// GetFromMajorVersion 通过major_version获取内容
func (obj *_AllTenantSstableColumnChecksumMgr) GetFromMajorVersion(majorVersion int64) (results []*AllTenantSstableColumnChecksum, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantSstableColumnChecksum{}).Where("`major_version` = ?", majorVersion).Find(&results).Error

	return
}

// GetBatchFromMajorVersion 批量查找
func (obj *_AllTenantSstableColumnChecksumMgr) GetBatchFromMajorVersion(majorVersions []int64) (results []*AllTenantSstableColumnChecksum, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantSstableColumnChecksum{}).Where("`major_version` IN (?)", majorVersions).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllTenantSstableColumnChecksumMgr) FetchByPrimaryKey(tenantID int64, dataTableID int64, indexID int64, partitionID int64, sstableType int64, columnID int64, svrIP string, svrPort int64) (result AllTenantSstableColumnChecksum, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantSstableColumnChecksum{}).Where("`tenant_id` = ? AND `data_table_id` = ? AND `index_id` = ? AND `partition_id` = ? AND `sstable_type` = ? AND `column_id` = ? AND `svr_ip` = ? AND `svr_port` = ?", tenantID, dataTableID, indexID, partitionID, sstableType, columnID, svrIP, svrPort).First(&result).Error

	return
}
