package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _AllVirtualStorageStatMgr struct {
	*_BaseMgr
}

// AllVirtualStorageStatMgr open func
func AllVirtualStorageStatMgr(db *gorm.DB) *_AllVirtualStorageStatMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualStorageStatMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualStorageStatMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_storage_stat"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualStorageStatMgr) GetTableName() string {
	return "__all_virtual_storage_stat"
}

// Reset 重置gorm会话
func (obj *_AllVirtualStorageStatMgr) Reset() *_AllVirtualStorageStatMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualStorageStatMgr) Get() (result AllVirtualStorageStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualStorageStat{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualStorageStatMgr) Gets() (results []*AllVirtualStorageStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualStorageStat{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualStorageStatMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualStorageStat{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTenantID tenant_id获取
func (obj *_AllVirtualStorageStatMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithSvrIP svr_ip获取
func (obj *_AllVirtualStorageStatMgr) WithSvrIP(svrIP string) Option {
	return optionFunc(func(o *options) { o.query["svr_ip"] = svrIP })
}

// WithSvrPort svr_port获取
func (obj *_AllVirtualStorageStatMgr) WithSvrPort(svrPort int64) Option {
	return optionFunc(func(o *options) { o.query["svr_port"] = svrPort })
}

// WithTableID table_id获取
func (obj *_AllVirtualStorageStatMgr) WithTableID(tableID int64) Option {
	return optionFunc(func(o *options) { o.query["table_id"] = tableID })
}

// WithPartitionCnt partition_cnt获取
func (obj *_AllVirtualStorageStatMgr) WithPartitionCnt(partitionCnt int64) Option {
	return optionFunc(func(o *options) { o.query["partition_cnt"] = partitionCnt })
}

// WithPartitionID partition_id获取
func (obj *_AllVirtualStorageStatMgr) WithPartitionID(partitionID int64) Option {
	return optionFunc(func(o *options) { o.query["partition_id"] = partitionID })
}

// WithMajorVersion major_version获取
func (obj *_AllVirtualStorageStatMgr) WithMajorVersion(majorVersion int64) Option {
	return optionFunc(func(o *options) { o.query["major_version"] = majorVersion })
}

// WithMinorVersion minor_version获取
func (obj *_AllVirtualStorageStatMgr) WithMinorVersion(minorVersion int64) Option {
	return optionFunc(func(o *options) { o.query["minor_version"] = minorVersion })
}

// WithSstableID sstable_id获取
func (obj *_AllVirtualStorageStatMgr) WithSstableID(sstableID int64) Option {
	return optionFunc(func(o *options) { o.query["sstable_id"] = sstableID })
}

// WithRole role获取
func (obj *_AllVirtualStorageStatMgr) WithRole(role int64) Option {
	return optionFunc(func(o *options) { o.query["role"] = role })
}

// WithDataChecksum data_checksum获取
func (obj *_AllVirtualStorageStatMgr) WithDataChecksum(dataChecksum int64) Option {
	return optionFunc(func(o *options) { o.query["data_checksum"] = dataChecksum })
}

// WithColumnChecksum column_checksum获取
func (obj *_AllVirtualStorageStatMgr) WithColumnChecksum(columnChecksum string) Option {
	return optionFunc(func(o *options) { o.query["column_checksum"] = columnChecksum })
}

// WithMacroBlocks macro_blocks获取
func (obj *_AllVirtualStorageStatMgr) WithMacroBlocks(macroBlocks string) Option {
	return optionFunc(func(o *options) { o.query["macro_blocks"] = macroBlocks })
}

// WithOccupySize occupy_size获取
func (obj *_AllVirtualStorageStatMgr) WithOccupySize(occupySize int64) Option {
	return optionFunc(func(o *options) { o.query["occupy_size"] = occupySize })
}

// WithUsedSize used_size获取
func (obj *_AllVirtualStorageStatMgr) WithUsedSize(usedSize int64) Option {
	return optionFunc(func(o *options) { o.query["used_size"] = usedSize })
}

// WithRowCount row_count获取
func (obj *_AllVirtualStorageStatMgr) WithRowCount(rowCount int64) Option {
	return optionFunc(func(o *options) { o.query["row_count"] = rowCount })
}

// WithStoreType store_type获取
func (obj *_AllVirtualStorageStatMgr) WithStoreType(storeType int64) Option {
	return optionFunc(func(o *options) { o.query["store_type"] = storeType })
}

// WithProgressiveMergeStartVersion progressive_merge_start_version获取
func (obj *_AllVirtualStorageStatMgr) WithProgressiveMergeStartVersion(progressiveMergeStartVersion int64) Option {
	return optionFunc(func(o *options) { o.query["progressive_merge_start_version"] = progressiveMergeStartVersion })
}

// WithProgressiveMergeEndVersion progressive_merge_end_version获取
func (obj *_AllVirtualStorageStatMgr) WithProgressiveMergeEndVersion(progressiveMergeEndVersion int64) Option {
	return optionFunc(func(o *options) { o.query["progressive_merge_end_version"] = progressiveMergeEndVersion })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualStorageStatMgr) GetByOption(opts ...Option) (result AllVirtualStorageStat, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualStorageStat{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualStorageStatMgr) GetByOptions(opts ...Option) (results []*AllVirtualStorageStat, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualStorageStat{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualStorageStatMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualStorageStat, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualStorageStat{}).Where(options.query)
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
func (obj *_AllVirtualStorageStatMgr) GetFromTenantID(tenantID int64) (results []*AllVirtualStorageStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualStorageStat{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllVirtualStorageStatMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllVirtualStorageStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualStorageStat{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromSvrIP 通过svr_ip获取内容
func (obj *_AllVirtualStorageStatMgr) GetFromSvrIP(svrIP string) (results []*AllVirtualStorageStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualStorageStat{}).Where("`svr_ip` = ?", svrIP).Find(&results).Error

	return
}

// GetBatchFromSvrIP 批量查找
func (obj *_AllVirtualStorageStatMgr) GetBatchFromSvrIP(svrIPs []string) (results []*AllVirtualStorageStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualStorageStat{}).Where("`svr_ip` IN (?)", svrIPs).Find(&results).Error

	return
}

// GetFromSvrPort 通过svr_port获取内容
func (obj *_AllVirtualStorageStatMgr) GetFromSvrPort(svrPort int64) (results []*AllVirtualStorageStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualStorageStat{}).Where("`svr_port` = ?", svrPort).Find(&results).Error

	return
}

// GetBatchFromSvrPort 批量查找
func (obj *_AllVirtualStorageStatMgr) GetBatchFromSvrPort(svrPorts []int64) (results []*AllVirtualStorageStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualStorageStat{}).Where("`svr_port` IN (?)", svrPorts).Find(&results).Error

	return
}

// GetFromTableID 通过table_id获取内容
func (obj *_AllVirtualStorageStatMgr) GetFromTableID(tableID int64) (results []*AllVirtualStorageStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualStorageStat{}).Where("`table_id` = ?", tableID).Find(&results).Error

	return
}

// GetBatchFromTableID 批量查找
func (obj *_AllVirtualStorageStatMgr) GetBatchFromTableID(tableIDs []int64) (results []*AllVirtualStorageStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualStorageStat{}).Where("`table_id` IN (?)", tableIDs).Find(&results).Error

	return
}

// GetFromPartitionCnt 通过partition_cnt获取内容
func (obj *_AllVirtualStorageStatMgr) GetFromPartitionCnt(partitionCnt int64) (results []*AllVirtualStorageStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualStorageStat{}).Where("`partition_cnt` = ?", partitionCnt).Find(&results).Error

	return
}

// GetBatchFromPartitionCnt 批量查找
func (obj *_AllVirtualStorageStatMgr) GetBatchFromPartitionCnt(partitionCnts []int64) (results []*AllVirtualStorageStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualStorageStat{}).Where("`partition_cnt` IN (?)", partitionCnts).Find(&results).Error

	return
}

// GetFromPartitionID 通过partition_id获取内容
func (obj *_AllVirtualStorageStatMgr) GetFromPartitionID(partitionID int64) (results []*AllVirtualStorageStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualStorageStat{}).Where("`partition_id` = ?", partitionID).Find(&results).Error

	return
}

// GetBatchFromPartitionID 批量查找
func (obj *_AllVirtualStorageStatMgr) GetBatchFromPartitionID(partitionIDs []int64) (results []*AllVirtualStorageStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualStorageStat{}).Where("`partition_id` IN (?)", partitionIDs).Find(&results).Error

	return
}

// GetFromMajorVersion 通过major_version获取内容
func (obj *_AllVirtualStorageStatMgr) GetFromMajorVersion(majorVersion int64) (results []*AllVirtualStorageStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualStorageStat{}).Where("`major_version` = ?", majorVersion).Find(&results).Error

	return
}

// GetBatchFromMajorVersion 批量查找
func (obj *_AllVirtualStorageStatMgr) GetBatchFromMajorVersion(majorVersions []int64) (results []*AllVirtualStorageStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualStorageStat{}).Where("`major_version` IN (?)", majorVersions).Find(&results).Error

	return
}

// GetFromMinorVersion 通过minor_version获取内容
func (obj *_AllVirtualStorageStatMgr) GetFromMinorVersion(minorVersion int64) (results []*AllVirtualStorageStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualStorageStat{}).Where("`minor_version` = ?", minorVersion).Find(&results).Error

	return
}

// GetBatchFromMinorVersion 批量查找
func (obj *_AllVirtualStorageStatMgr) GetBatchFromMinorVersion(minorVersions []int64) (results []*AllVirtualStorageStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualStorageStat{}).Where("`minor_version` IN (?)", minorVersions).Find(&results).Error

	return
}

// GetFromSstableID 通过sstable_id获取内容
func (obj *_AllVirtualStorageStatMgr) GetFromSstableID(sstableID int64) (results []*AllVirtualStorageStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualStorageStat{}).Where("`sstable_id` = ?", sstableID).Find(&results).Error

	return
}

// GetBatchFromSstableID 批量查找
func (obj *_AllVirtualStorageStatMgr) GetBatchFromSstableID(sstableIDs []int64) (results []*AllVirtualStorageStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualStorageStat{}).Where("`sstable_id` IN (?)", sstableIDs).Find(&results).Error

	return
}

// GetFromRole 通过role获取内容
func (obj *_AllVirtualStorageStatMgr) GetFromRole(role int64) (results []*AllVirtualStorageStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualStorageStat{}).Where("`role` = ?", role).Find(&results).Error

	return
}

// GetBatchFromRole 批量查找
func (obj *_AllVirtualStorageStatMgr) GetBatchFromRole(roles []int64) (results []*AllVirtualStorageStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualStorageStat{}).Where("`role` IN (?)", roles).Find(&results).Error

	return
}

// GetFromDataChecksum 通过data_checksum获取内容
func (obj *_AllVirtualStorageStatMgr) GetFromDataChecksum(dataChecksum int64) (results []*AllVirtualStorageStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualStorageStat{}).Where("`data_checksum` = ?", dataChecksum).Find(&results).Error

	return
}

// GetBatchFromDataChecksum 批量查找
func (obj *_AllVirtualStorageStatMgr) GetBatchFromDataChecksum(dataChecksums []int64) (results []*AllVirtualStorageStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualStorageStat{}).Where("`data_checksum` IN (?)", dataChecksums).Find(&results).Error

	return
}

// GetFromColumnChecksum 通过column_checksum获取内容
func (obj *_AllVirtualStorageStatMgr) GetFromColumnChecksum(columnChecksum string) (results []*AllVirtualStorageStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualStorageStat{}).Where("`column_checksum` = ?", columnChecksum).Find(&results).Error

	return
}

// GetBatchFromColumnChecksum 批量查找
func (obj *_AllVirtualStorageStatMgr) GetBatchFromColumnChecksum(columnChecksums []string) (results []*AllVirtualStorageStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualStorageStat{}).Where("`column_checksum` IN (?)", columnChecksums).Find(&results).Error

	return
}

// GetFromMacroBlocks 通过macro_blocks获取内容
func (obj *_AllVirtualStorageStatMgr) GetFromMacroBlocks(macroBlocks string) (results []*AllVirtualStorageStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualStorageStat{}).Where("`macro_blocks` = ?", macroBlocks).Find(&results).Error

	return
}

// GetBatchFromMacroBlocks 批量查找
func (obj *_AllVirtualStorageStatMgr) GetBatchFromMacroBlocks(macroBlockss []string) (results []*AllVirtualStorageStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualStorageStat{}).Where("`macro_blocks` IN (?)", macroBlockss).Find(&results).Error

	return
}

// GetFromOccupySize 通过occupy_size获取内容
func (obj *_AllVirtualStorageStatMgr) GetFromOccupySize(occupySize int64) (results []*AllVirtualStorageStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualStorageStat{}).Where("`occupy_size` = ?", occupySize).Find(&results).Error

	return
}

// GetBatchFromOccupySize 批量查找
func (obj *_AllVirtualStorageStatMgr) GetBatchFromOccupySize(occupySizes []int64) (results []*AllVirtualStorageStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualStorageStat{}).Where("`occupy_size` IN (?)", occupySizes).Find(&results).Error

	return
}

// GetFromUsedSize 通过used_size获取内容
func (obj *_AllVirtualStorageStatMgr) GetFromUsedSize(usedSize int64) (results []*AllVirtualStorageStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualStorageStat{}).Where("`used_size` = ?", usedSize).Find(&results).Error

	return
}

// GetBatchFromUsedSize 批量查找
func (obj *_AllVirtualStorageStatMgr) GetBatchFromUsedSize(usedSizes []int64) (results []*AllVirtualStorageStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualStorageStat{}).Where("`used_size` IN (?)", usedSizes).Find(&results).Error

	return
}

// GetFromRowCount 通过row_count获取内容
func (obj *_AllVirtualStorageStatMgr) GetFromRowCount(rowCount int64) (results []*AllVirtualStorageStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualStorageStat{}).Where("`row_count` = ?", rowCount).Find(&results).Error

	return
}

// GetBatchFromRowCount 批量查找
func (obj *_AllVirtualStorageStatMgr) GetBatchFromRowCount(rowCounts []int64) (results []*AllVirtualStorageStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualStorageStat{}).Where("`row_count` IN (?)", rowCounts).Find(&results).Error

	return
}

// GetFromStoreType 通过store_type获取内容
func (obj *_AllVirtualStorageStatMgr) GetFromStoreType(storeType int64) (results []*AllVirtualStorageStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualStorageStat{}).Where("`store_type` = ?", storeType).Find(&results).Error

	return
}

// GetBatchFromStoreType 批量查找
func (obj *_AllVirtualStorageStatMgr) GetBatchFromStoreType(storeTypes []int64) (results []*AllVirtualStorageStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualStorageStat{}).Where("`store_type` IN (?)", storeTypes).Find(&results).Error

	return
}

// GetFromProgressiveMergeStartVersion 通过progressive_merge_start_version获取内容
func (obj *_AllVirtualStorageStatMgr) GetFromProgressiveMergeStartVersion(progressiveMergeStartVersion int64) (results []*AllVirtualStorageStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualStorageStat{}).Where("`progressive_merge_start_version` = ?", progressiveMergeStartVersion).Find(&results).Error

	return
}

// GetBatchFromProgressiveMergeStartVersion 批量查找
func (obj *_AllVirtualStorageStatMgr) GetBatchFromProgressiveMergeStartVersion(progressiveMergeStartVersions []int64) (results []*AllVirtualStorageStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualStorageStat{}).Where("`progressive_merge_start_version` IN (?)", progressiveMergeStartVersions).Find(&results).Error

	return
}

// GetFromProgressiveMergeEndVersion 通过progressive_merge_end_version获取内容
func (obj *_AllVirtualStorageStatMgr) GetFromProgressiveMergeEndVersion(progressiveMergeEndVersion int64) (results []*AllVirtualStorageStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualStorageStat{}).Where("`progressive_merge_end_version` = ?", progressiveMergeEndVersion).Find(&results).Error

	return
}

// GetBatchFromProgressiveMergeEndVersion 批量查找
func (obj *_AllVirtualStorageStatMgr) GetBatchFromProgressiveMergeEndVersion(progressiveMergeEndVersions []int64) (results []*AllVirtualStorageStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualStorageStat{}).Where("`progressive_merge_end_version` IN (?)", progressiveMergeEndVersions).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////
