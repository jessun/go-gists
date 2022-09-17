package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _AllVirtualPartitionSstableMacroInfoMgr struct {
	*_BaseMgr
}

// AllVirtualPartitionSstableMacroInfoMgr open func
func AllVirtualPartitionSstableMacroInfoMgr(db *gorm.DB) *_AllVirtualPartitionSstableMacroInfoMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualPartitionSstableMacroInfoMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualPartitionSstableMacroInfoMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_partition_sstable_macro_info"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualPartitionSstableMacroInfoMgr) GetTableName() string {
	return "__all_virtual_partition_sstable_macro_info"
}

// Reset 重置gorm会话
func (obj *_AllVirtualPartitionSstableMacroInfoMgr) Reset() *_AllVirtualPartitionSstableMacroInfoMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualPartitionSstableMacroInfoMgr) Get() (result AllVirtualPartitionSstableMacroInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionSstableMacroInfo{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualPartitionSstableMacroInfoMgr) Gets() (results []*AllVirtualPartitionSstableMacroInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionSstableMacroInfo{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualPartitionSstableMacroInfoMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionSstableMacroInfo{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithSvrIP svr_ip获取
func (obj *_AllVirtualPartitionSstableMacroInfoMgr) WithSvrIP(svrIP string) Option {
	return optionFunc(func(o *options) { o.query["svr_ip"] = svrIP })
}

// WithSvrPort svr_port获取
func (obj *_AllVirtualPartitionSstableMacroInfoMgr) WithSvrPort(svrPort int64) Option {
	return optionFunc(func(o *options) { o.query["svr_port"] = svrPort })
}

// WithTenantID tenant_id获取
func (obj *_AllVirtualPartitionSstableMacroInfoMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithTableID table_id获取
func (obj *_AllVirtualPartitionSstableMacroInfoMgr) WithTableID(tableID int64) Option {
	return optionFunc(func(o *options) { o.query["table_id"] = tableID })
}

// WithPartitionID partition_id获取
func (obj *_AllVirtualPartitionSstableMacroInfoMgr) WithPartitionID(partitionID int64) Option {
	return optionFunc(func(o *options) { o.query["partition_id"] = partitionID })
}

// WithDataVersion data_version获取
func (obj *_AllVirtualPartitionSstableMacroInfoMgr) WithDataVersion(dataVersion int64) Option {
	return optionFunc(func(o *options) { o.query["data_version"] = dataVersion })
}

// WithBaseVersion base_version获取
func (obj *_AllVirtualPartitionSstableMacroInfoMgr) WithBaseVersion(baseVersion int64) Option {
	return optionFunc(func(o *options) { o.query["base_version"] = baseVersion })
}

// WithMultiVersionStart multi_version_start获取
func (obj *_AllVirtualPartitionSstableMacroInfoMgr) WithMultiVersionStart(multiVersionStart int64) Option {
	return optionFunc(func(o *options) { o.query["multi_version_start"] = multiVersionStart })
}

// WithSnapshotVersion snapshot_version获取
func (obj *_AllVirtualPartitionSstableMacroInfoMgr) WithSnapshotVersion(snapshotVersion int64) Option {
	return optionFunc(func(o *options) { o.query["snapshot_version"] = snapshotVersion })
}

// WithMacroIDxInSstable macro_idx_in_sstable获取
func (obj *_AllVirtualPartitionSstableMacroInfoMgr) WithMacroIDxInSstable(macroIDxInSstable int64) Option {
	return optionFunc(func(o *options) { o.query["macro_idx_in_sstable"] = macroIDxInSstable })
}

// WithMajorTableID major_table_id获取
func (obj *_AllVirtualPartitionSstableMacroInfoMgr) WithMajorTableID(majorTableID int64) Option {
	return optionFunc(func(o *options) { o.query["major_table_id"] = majorTableID })
}

// WithMacroDataVersion macro_data_version获取
func (obj *_AllVirtualPartitionSstableMacroInfoMgr) WithMacroDataVersion(macroDataVersion int64) Option {
	return optionFunc(func(o *options) { o.query["macro_data_version"] = macroDataVersion })
}

// WithMacroIDxInDataFile macro_idx_in_data_file获取
func (obj *_AllVirtualPartitionSstableMacroInfoMgr) WithMacroIDxInDataFile(macroIDxInDataFile int64) Option {
	return optionFunc(func(o *options) { o.query["macro_idx_in_data_file"] = macroIDxInDataFile })
}

// WithDataSeq data_seq获取
func (obj *_AllVirtualPartitionSstableMacroInfoMgr) WithDataSeq(dataSeq int64) Option {
	return optionFunc(func(o *options) { o.query["data_seq"] = dataSeq })
}

// WithRowCount row_count获取
func (obj *_AllVirtualPartitionSstableMacroInfoMgr) WithRowCount(rowCount int64) Option {
	return optionFunc(func(o *options) { o.query["row_count"] = rowCount })
}

// WithOccupySize occupy_size获取
func (obj *_AllVirtualPartitionSstableMacroInfoMgr) WithOccupySize(occupySize int64) Option {
	return optionFunc(func(o *options) { o.query["occupy_size"] = occupySize })
}

// WithMicroBlockCount micro_block_count获取
func (obj *_AllVirtualPartitionSstableMacroInfoMgr) WithMicroBlockCount(microBlockCount int64) Option {
	return optionFunc(func(o *options) { o.query["micro_block_count"] = microBlockCount })
}

// WithDataChecksum data_checksum获取
func (obj *_AllVirtualPartitionSstableMacroInfoMgr) WithDataChecksum(dataChecksum int64) Option {
	return optionFunc(func(o *options) { o.query["data_checksum"] = dataChecksum })
}

// WithSchemaVersion schema_version获取
func (obj *_AllVirtualPartitionSstableMacroInfoMgr) WithSchemaVersion(schemaVersion int64) Option {
	return optionFunc(func(o *options) { o.query["schema_version"] = schemaVersion })
}

// WithMacroRange macro_range获取
func (obj *_AllVirtualPartitionSstableMacroInfoMgr) WithMacroRange(macroRange string) Option {
	return optionFunc(func(o *options) { o.query["macro_range"] = macroRange })
}

// WithRowCountDelta row_count_delta获取
func (obj *_AllVirtualPartitionSstableMacroInfoMgr) WithRowCountDelta(rowCountDelta int64) Option {
	return optionFunc(func(o *options) { o.query["row_count_delta"] = rowCountDelta })
}

// WithMacroBlockType macro_block_type获取
func (obj *_AllVirtualPartitionSstableMacroInfoMgr) WithMacroBlockType(macroBlockType string) Option {
	return optionFunc(func(o *options) { o.query["macro_block_type"] = macroBlockType })
}

// WithCompressorName compressor_name获取
func (obj *_AllVirtualPartitionSstableMacroInfoMgr) WithCompressorName(compressorName string) Option {
	return optionFunc(func(o *options) { o.query["compressor_name"] = compressorName })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualPartitionSstableMacroInfoMgr) GetByOption(opts ...Option) (result AllVirtualPartitionSstableMacroInfo, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionSstableMacroInfo{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualPartitionSstableMacroInfoMgr) GetByOptions(opts ...Option) (results []*AllVirtualPartitionSstableMacroInfo, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionSstableMacroInfo{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualPartitionSstableMacroInfoMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualPartitionSstableMacroInfo, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionSstableMacroInfo{}).Where(options.query)
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
func (obj *_AllVirtualPartitionSstableMacroInfoMgr) GetFromSvrIP(svrIP string) (results []*AllVirtualPartitionSstableMacroInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionSstableMacroInfo{}).Where("`svr_ip` = ?", svrIP).Find(&results).Error

	return
}

// GetBatchFromSvrIP 批量查找
func (obj *_AllVirtualPartitionSstableMacroInfoMgr) GetBatchFromSvrIP(svrIPs []string) (results []*AllVirtualPartitionSstableMacroInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionSstableMacroInfo{}).Where("`svr_ip` IN (?)", svrIPs).Find(&results).Error

	return
}

// GetFromSvrPort 通过svr_port获取内容
func (obj *_AllVirtualPartitionSstableMacroInfoMgr) GetFromSvrPort(svrPort int64) (results []*AllVirtualPartitionSstableMacroInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionSstableMacroInfo{}).Where("`svr_port` = ?", svrPort).Find(&results).Error

	return
}

// GetBatchFromSvrPort 批量查找
func (obj *_AllVirtualPartitionSstableMacroInfoMgr) GetBatchFromSvrPort(svrPorts []int64) (results []*AllVirtualPartitionSstableMacroInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionSstableMacroInfo{}).Where("`svr_port` IN (?)", svrPorts).Find(&results).Error

	return
}

// GetFromTenantID 通过tenant_id获取内容
func (obj *_AllVirtualPartitionSstableMacroInfoMgr) GetFromTenantID(tenantID int64) (results []*AllVirtualPartitionSstableMacroInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionSstableMacroInfo{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllVirtualPartitionSstableMacroInfoMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllVirtualPartitionSstableMacroInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionSstableMacroInfo{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromTableID 通过table_id获取内容
func (obj *_AllVirtualPartitionSstableMacroInfoMgr) GetFromTableID(tableID int64) (results []*AllVirtualPartitionSstableMacroInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionSstableMacroInfo{}).Where("`table_id` = ?", tableID).Find(&results).Error

	return
}

// GetBatchFromTableID 批量查找
func (obj *_AllVirtualPartitionSstableMacroInfoMgr) GetBatchFromTableID(tableIDs []int64) (results []*AllVirtualPartitionSstableMacroInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionSstableMacroInfo{}).Where("`table_id` IN (?)", tableIDs).Find(&results).Error

	return
}

// GetFromPartitionID 通过partition_id获取内容
func (obj *_AllVirtualPartitionSstableMacroInfoMgr) GetFromPartitionID(partitionID int64) (results []*AllVirtualPartitionSstableMacroInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionSstableMacroInfo{}).Where("`partition_id` = ?", partitionID).Find(&results).Error

	return
}

// GetBatchFromPartitionID 批量查找
func (obj *_AllVirtualPartitionSstableMacroInfoMgr) GetBatchFromPartitionID(partitionIDs []int64) (results []*AllVirtualPartitionSstableMacroInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionSstableMacroInfo{}).Where("`partition_id` IN (?)", partitionIDs).Find(&results).Error

	return
}

// GetFromDataVersion 通过data_version获取内容
func (obj *_AllVirtualPartitionSstableMacroInfoMgr) GetFromDataVersion(dataVersion int64) (results []*AllVirtualPartitionSstableMacroInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionSstableMacroInfo{}).Where("`data_version` = ?", dataVersion).Find(&results).Error

	return
}

// GetBatchFromDataVersion 批量查找
func (obj *_AllVirtualPartitionSstableMacroInfoMgr) GetBatchFromDataVersion(dataVersions []int64) (results []*AllVirtualPartitionSstableMacroInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionSstableMacroInfo{}).Where("`data_version` IN (?)", dataVersions).Find(&results).Error

	return
}

// GetFromBaseVersion 通过base_version获取内容
func (obj *_AllVirtualPartitionSstableMacroInfoMgr) GetFromBaseVersion(baseVersion int64) (results []*AllVirtualPartitionSstableMacroInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionSstableMacroInfo{}).Where("`base_version` = ?", baseVersion).Find(&results).Error

	return
}

// GetBatchFromBaseVersion 批量查找
func (obj *_AllVirtualPartitionSstableMacroInfoMgr) GetBatchFromBaseVersion(baseVersions []int64) (results []*AllVirtualPartitionSstableMacroInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionSstableMacroInfo{}).Where("`base_version` IN (?)", baseVersions).Find(&results).Error

	return
}

// GetFromMultiVersionStart 通过multi_version_start获取内容
func (obj *_AllVirtualPartitionSstableMacroInfoMgr) GetFromMultiVersionStart(multiVersionStart int64) (results []*AllVirtualPartitionSstableMacroInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionSstableMacroInfo{}).Where("`multi_version_start` = ?", multiVersionStart).Find(&results).Error

	return
}

// GetBatchFromMultiVersionStart 批量查找
func (obj *_AllVirtualPartitionSstableMacroInfoMgr) GetBatchFromMultiVersionStart(multiVersionStarts []int64) (results []*AllVirtualPartitionSstableMacroInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionSstableMacroInfo{}).Where("`multi_version_start` IN (?)", multiVersionStarts).Find(&results).Error

	return
}

// GetFromSnapshotVersion 通过snapshot_version获取内容
func (obj *_AllVirtualPartitionSstableMacroInfoMgr) GetFromSnapshotVersion(snapshotVersion int64) (results []*AllVirtualPartitionSstableMacroInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionSstableMacroInfo{}).Where("`snapshot_version` = ?", snapshotVersion).Find(&results).Error

	return
}

// GetBatchFromSnapshotVersion 批量查找
func (obj *_AllVirtualPartitionSstableMacroInfoMgr) GetBatchFromSnapshotVersion(snapshotVersions []int64) (results []*AllVirtualPartitionSstableMacroInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionSstableMacroInfo{}).Where("`snapshot_version` IN (?)", snapshotVersions).Find(&results).Error

	return
}

// GetFromMacroIDxInSstable 通过macro_idx_in_sstable获取内容
func (obj *_AllVirtualPartitionSstableMacroInfoMgr) GetFromMacroIDxInSstable(macroIDxInSstable int64) (results []*AllVirtualPartitionSstableMacroInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionSstableMacroInfo{}).Where("`macro_idx_in_sstable` = ?", macroIDxInSstable).Find(&results).Error

	return
}

// GetBatchFromMacroIDxInSstable 批量查找
func (obj *_AllVirtualPartitionSstableMacroInfoMgr) GetBatchFromMacroIDxInSstable(macroIDxInSstables []int64) (results []*AllVirtualPartitionSstableMacroInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionSstableMacroInfo{}).Where("`macro_idx_in_sstable` IN (?)", macroIDxInSstables).Find(&results).Error

	return
}

// GetFromMajorTableID 通过major_table_id获取内容
func (obj *_AllVirtualPartitionSstableMacroInfoMgr) GetFromMajorTableID(majorTableID int64) (results []*AllVirtualPartitionSstableMacroInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionSstableMacroInfo{}).Where("`major_table_id` = ?", majorTableID).Find(&results).Error

	return
}

// GetBatchFromMajorTableID 批量查找
func (obj *_AllVirtualPartitionSstableMacroInfoMgr) GetBatchFromMajorTableID(majorTableIDs []int64) (results []*AllVirtualPartitionSstableMacroInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionSstableMacroInfo{}).Where("`major_table_id` IN (?)", majorTableIDs).Find(&results).Error

	return
}

// GetFromMacroDataVersion 通过macro_data_version获取内容
func (obj *_AllVirtualPartitionSstableMacroInfoMgr) GetFromMacroDataVersion(macroDataVersion int64) (results []*AllVirtualPartitionSstableMacroInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionSstableMacroInfo{}).Where("`macro_data_version` = ?", macroDataVersion).Find(&results).Error

	return
}

// GetBatchFromMacroDataVersion 批量查找
func (obj *_AllVirtualPartitionSstableMacroInfoMgr) GetBatchFromMacroDataVersion(macroDataVersions []int64) (results []*AllVirtualPartitionSstableMacroInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionSstableMacroInfo{}).Where("`macro_data_version` IN (?)", macroDataVersions).Find(&results).Error

	return
}

// GetFromMacroIDxInDataFile 通过macro_idx_in_data_file获取内容
func (obj *_AllVirtualPartitionSstableMacroInfoMgr) GetFromMacroIDxInDataFile(macroIDxInDataFile int64) (results []*AllVirtualPartitionSstableMacroInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionSstableMacroInfo{}).Where("`macro_idx_in_data_file` = ?", macroIDxInDataFile).Find(&results).Error

	return
}

// GetBatchFromMacroIDxInDataFile 批量查找
func (obj *_AllVirtualPartitionSstableMacroInfoMgr) GetBatchFromMacroIDxInDataFile(macroIDxInDataFiles []int64) (results []*AllVirtualPartitionSstableMacroInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionSstableMacroInfo{}).Where("`macro_idx_in_data_file` IN (?)", macroIDxInDataFiles).Find(&results).Error

	return
}

// GetFromDataSeq 通过data_seq获取内容
func (obj *_AllVirtualPartitionSstableMacroInfoMgr) GetFromDataSeq(dataSeq int64) (results []*AllVirtualPartitionSstableMacroInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionSstableMacroInfo{}).Where("`data_seq` = ?", dataSeq).Find(&results).Error

	return
}

// GetBatchFromDataSeq 批量查找
func (obj *_AllVirtualPartitionSstableMacroInfoMgr) GetBatchFromDataSeq(dataSeqs []int64) (results []*AllVirtualPartitionSstableMacroInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionSstableMacroInfo{}).Where("`data_seq` IN (?)", dataSeqs).Find(&results).Error

	return
}

// GetFromRowCount 通过row_count获取内容
func (obj *_AllVirtualPartitionSstableMacroInfoMgr) GetFromRowCount(rowCount int64) (results []*AllVirtualPartitionSstableMacroInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionSstableMacroInfo{}).Where("`row_count` = ?", rowCount).Find(&results).Error

	return
}

// GetBatchFromRowCount 批量查找
func (obj *_AllVirtualPartitionSstableMacroInfoMgr) GetBatchFromRowCount(rowCounts []int64) (results []*AllVirtualPartitionSstableMacroInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionSstableMacroInfo{}).Where("`row_count` IN (?)", rowCounts).Find(&results).Error

	return
}

// GetFromOccupySize 通过occupy_size获取内容
func (obj *_AllVirtualPartitionSstableMacroInfoMgr) GetFromOccupySize(occupySize int64) (results []*AllVirtualPartitionSstableMacroInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionSstableMacroInfo{}).Where("`occupy_size` = ?", occupySize).Find(&results).Error

	return
}

// GetBatchFromOccupySize 批量查找
func (obj *_AllVirtualPartitionSstableMacroInfoMgr) GetBatchFromOccupySize(occupySizes []int64) (results []*AllVirtualPartitionSstableMacroInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionSstableMacroInfo{}).Where("`occupy_size` IN (?)", occupySizes).Find(&results).Error

	return
}

// GetFromMicroBlockCount 通过micro_block_count获取内容
func (obj *_AllVirtualPartitionSstableMacroInfoMgr) GetFromMicroBlockCount(microBlockCount int64) (results []*AllVirtualPartitionSstableMacroInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionSstableMacroInfo{}).Where("`micro_block_count` = ?", microBlockCount).Find(&results).Error

	return
}

// GetBatchFromMicroBlockCount 批量查找
func (obj *_AllVirtualPartitionSstableMacroInfoMgr) GetBatchFromMicroBlockCount(microBlockCounts []int64) (results []*AllVirtualPartitionSstableMacroInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionSstableMacroInfo{}).Where("`micro_block_count` IN (?)", microBlockCounts).Find(&results).Error

	return
}

// GetFromDataChecksum 通过data_checksum获取内容
func (obj *_AllVirtualPartitionSstableMacroInfoMgr) GetFromDataChecksum(dataChecksum int64) (results []*AllVirtualPartitionSstableMacroInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionSstableMacroInfo{}).Where("`data_checksum` = ?", dataChecksum).Find(&results).Error

	return
}

// GetBatchFromDataChecksum 批量查找
func (obj *_AllVirtualPartitionSstableMacroInfoMgr) GetBatchFromDataChecksum(dataChecksums []int64) (results []*AllVirtualPartitionSstableMacroInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionSstableMacroInfo{}).Where("`data_checksum` IN (?)", dataChecksums).Find(&results).Error

	return
}

// GetFromSchemaVersion 通过schema_version获取内容
func (obj *_AllVirtualPartitionSstableMacroInfoMgr) GetFromSchemaVersion(schemaVersion int64) (results []*AllVirtualPartitionSstableMacroInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionSstableMacroInfo{}).Where("`schema_version` = ?", schemaVersion).Find(&results).Error

	return
}

// GetBatchFromSchemaVersion 批量查找
func (obj *_AllVirtualPartitionSstableMacroInfoMgr) GetBatchFromSchemaVersion(schemaVersions []int64) (results []*AllVirtualPartitionSstableMacroInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionSstableMacroInfo{}).Where("`schema_version` IN (?)", schemaVersions).Find(&results).Error

	return
}

// GetFromMacroRange 通过macro_range获取内容
func (obj *_AllVirtualPartitionSstableMacroInfoMgr) GetFromMacroRange(macroRange string) (results []*AllVirtualPartitionSstableMacroInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionSstableMacroInfo{}).Where("`macro_range` = ?", macroRange).Find(&results).Error

	return
}

// GetBatchFromMacroRange 批量查找
func (obj *_AllVirtualPartitionSstableMacroInfoMgr) GetBatchFromMacroRange(macroRanges []string) (results []*AllVirtualPartitionSstableMacroInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionSstableMacroInfo{}).Where("`macro_range` IN (?)", macroRanges).Find(&results).Error

	return
}

// GetFromRowCountDelta 通过row_count_delta获取内容
func (obj *_AllVirtualPartitionSstableMacroInfoMgr) GetFromRowCountDelta(rowCountDelta int64) (results []*AllVirtualPartitionSstableMacroInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionSstableMacroInfo{}).Where("`row_count_delta` = ?", rowCountDelta).Find(&results).Error

	return
}

// GetBatchFromRowCountDelta 批量查找
func (obj *_AllVirtualPartitionSstableMacroInfoMgr) GetBatchFromRowCountDelta(rowCountDeltas []int64) (results []*AllVirtualPartitionSstableMacroInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionSstableMacroInfo{}).Where("`row_count_delta` IN (?)", rowCountDeltas).Find(&results).Error

	return
}

// GetFromMacroBlockType 通过macro_block_type获取内容
func (obj *_AllVirtualPartitionSstableMacroInfoMgr) GetFromMacroBlockType(macroBlockType string) (results []*AllVirtualPartitionSstableMacroInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionSstableMacroInfo{}).Where("`macro_block_type` = ?", macroBlockType).Find(&results).Error

	return
}

// GetBatchFromMacroBlockType 批量查找
func (obj *_AllVirtualPartitionSstableMacroInfoMgr) GetBatchFromMacroBlockType(macroBlockTypes []string) (results []*AllVirtualPartitionSstableMacroInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionSstableMacroInfo{}).Where("`macro_block_type` IN (?)", macroBlockTypes).Find(&results).Error

	return
}

// GetFromCompressorName 通过compressor_name获取内容
func (obj *_AllVirtualPartitionSstableMacroInfoMgr) GetFromCompressorName(compressorName string) (results []*AllVirtualPartitionSstableMacroInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionSstableMacroInfo{}).Where("`compressor_name` = ?", compressorName).Find(&results).Error

	return
}

// GetBatchFromCompressorName 批量查找
func (obj *_AllVirtualPartitionSstableMacroInfoMgr) GetBatchFromCompressorName(compressorNames []string) (results []*AllVirtualPartitionSstableMacroInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionSstableMacroInfo{}).Where("`compressor_name` IN (?)", compressorNames).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllVirtualPartitionSstableMacroInfoMgr) FetchByPrimaryKey(svrIP string, svrPort int64, tenantID int64, tableID int64, partitionID int64, dataVersion int64, baseVersion int64, multiVersionStart int64, snapshotVersion int64, macroIDxInSstable int64) (result AllVirtualPartitionSstableMacroInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionSstableMacroInfo{}).Where("`svr_ip` = ? AND `svr_port` = ? AND `tenant_id` = ? AND `table_id` = ? AND `partition_id` = ? AND `data_version` = ? AND `base_version` = ? AND `multi_version_start` = ? AND `snapshot_version` = ? AND `macro_idx_in_sstable` = ?", svrIP, svrPort, tenantID, tableID, partitionID, dataVersion, baseVersion, multiVersionStart, snapshotVersion, macroIDxInSstable).First(&result).Error

	return
}
