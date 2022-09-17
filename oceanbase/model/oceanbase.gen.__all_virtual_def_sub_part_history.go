package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _AllVirtualDefSubPartHistoryMgr struct {
	*_BaseMgr
}

// AllVirtualDefSubPartHistoryMgr open func
func AllVirtualDefSubPartHistoryMgr(db *gorm.DB) *_AllVirtualDefSubPartHistoryMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualDefSubPartHistoryMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualDefSubPartHistoryMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_def_sub_part_history"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualDefSubPartHistoryMgr) GetTableName() string {
	return "__all_virtual_def_sub_part_history"
}

// Reset 重置gorm会话
func (obj *_AllVirtualDefSubPartHistoryMgr) Reset() *_AllVirtualDefSubPartHistoryMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualDefSubPartHistoryMgr) Get() (result AllVirtualDefSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDefSubPartHistory{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualDefSubPartHistoryMgr) Gets() (results []*AllVirtualDefSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDefSubPartHistory{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualDefSubPartHistoryMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualDefSubPartHistory{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTenantID tenant_id获取
func (obj *_AllVirtualDefSubPartHistoryMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithTableID table_id获取
func (obj *_AllVirtualDefSubPartHistoryMgr) WithTableID(tableID int64) Option {
	return optionFunc(func(o *options) { o.query["table_id"] = tableID })
}

// WithSubPartID sub_part_id获取
func (obj *_AllVirtualDefSubPartHistoryMgr) WithSubPartID(subPartID int64) Option {
	return optionFunc(func(o *options) { o.query["sub_part_id"] = subPartID })
}

// WithSchemaVersion schema_version获取
func (obj *_AllVirtualDefSubPartHistoryMgr) WithSchemaVersion(schemaVersion int64) Option {
	return optionFunc(func(o *options) { o.query["schema_version"] = schemaVersion })
}

// WithGmtCreate gmt_create获取
func (obj *_AllVirtualDefSubPartHistoryMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllVirtualDefSubPartHistoryMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithIsDeleted is_deleted获取
func (obj *_AllVirtualDefSubPartHistoryMgr) WithIsDeleted(isDeleted int64) Option {
	return optionFunc(func(o *options) { o.query["is_deleted"] = isDeleted })
}

// WithSubPartName sub_part_name获取
func (obj *_AllVirtualDefSubPartHistoryMgr) WithSubPartName(subPartName string) Option {
	return optionFunc(func(o *options) { o.query["sub_part_name"] = subPartName })
}

// WithHighBoundVal high_bound_val获取
func (obj *_AllVirtualDefSubPartHistoryMgr) WithHighBoundVal(highBoundVal string) Option {
	return optionFunc(func(o *options) { o.query["high_bound_val"] = highBoundVal })
}

// WithBHighBoundVal b_high_bound_val获取
func (obj *_AllVirtualDefSubPartHistoryMgr) WithBHighBoundVal(bHighBoundVal string) Option {
	return optionFunc(func(o *options) { o.query["b_high_bound_val"] = bHighBoundVal })
}

// WithBlockSize block_size获取
func (obj *_AllVirtualDefSubPartHistoryMgr) WithBlockSize(blockSize int64) Option {
	return optionFunc(func(o *options) { o.query["block_size"] = blockSize })
}

// WithReplicaNum replica_num获取
func (obj *_AllVirtualDefSubPartHistoryMgr) WithReplicaNum(replicaNum int64) Option {
	return optionFunc(func(o *options) { o.query["replica_num"] = replicaNum })
}

// WithCompressFuncName compress_func_name获取
func (obj *_AllVirtualDefSubPartHistoryMgr) WithCompressFuncName(compressFuncName string) Option {
	return optionFunc(func(o *options) { o.query["compress_func_name"] = compressFuncName })
}

// WithSpare1 spare1获取
func (obj *_AllVirtualDefSubPartHistoryMgr) WithSpare1(spare1 int64) Option {
	return optionFunc(func(o *options) { o.query["spare1"] = spare1 })
}

// WithSpare2 spare2获取
func (obj *_AllVirtualDefSubPartHistoryMgr) WithSpare2(spare2 int64) Option {
	return optionFunc(func(o *options) { o.query["spare2"] = spare2 })
}

// WithSpare3 spare3获取
func (obj *_AllVirtualDefSubPartHistoryMgr) WithSpare3(spare3 string) Option {
	return optionFunc(func(o *options) { o.query["spare3"] = spare3 })
}

// WithComment comment获取
func (obj *_AllVirtualDefSubPartHistoryMgr) WithComment(comment string) Option {
	return optionFunc(func(o *options) { o.query["comment"] = comment })
}

// WithListVal list_val获取
func (obj *_AllVirtualDefSubPartHistoryMgr) WithListVal(listVal string) Option {
	return optionFunc(func(o *options) { o.query["list_val"] = listVal })
}

// WithBListVal b_list_val获取
func (obj *_AllVirtualDefSubPartHistoryMgr) WithBListVal(bListVal string) Option {
	return optionFunc(func(o *options) { o.query["b_list_val"] = bListVal })
}

// WithSubPartIDx sub_part_idx获取
func (obj *_AllVirtualDefSubPartHistoryMgr) WithSubPartIDx(subPartIDx int64) Option {
	return optionFunc(func(o *options) { o.query["sub_part_idx"] = subPartIDx })
}

// WithSourcePartitionID source_partition_id获取
func (obj *_AllVirtualDefSubPartHistoryMgr) WithSourcePartitionID(sourcePartitionID string) Option {
	return optionFunc(func(o *options) { o.query["source_partition_id"] = sourcePartitionID })
}

// WithMappingPgSubPartID mapping_pg_sub_part_id获取
func (obj *_AllVirtualDefSubPartHistoryMgr) WithMappingPgSubPartID(mappingPgSubPartID int64) Option {
	return optionFunc(func(o *options) { o.query["mapping_pg_sub_part_id"] = mappingPgSubPartID })
}

// WithTablespaceID tablespace_id获取
func (obj *_AllVirtualDefSubPartHistoryMgr) WithTablespaceID(tablespaceID int64) Option {
	return optionFunc(func(o *options) { o.query["tablespace_id"] = tablespaceID })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualDefSubPartHistoryMgr) GetByOption(opts ...Option) (result AllVirtualDefSubPartHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDefSubPartHistory{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualDefSubPartHistoryMgr) GetByOptions(opts ...Option) (results []*AllVirtualDefSubPartHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDefSubPartHistory{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualDefSubPartHistoryMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualDefSubPartHistory, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualDefSubPartHistory{}).Where(options.query)
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
func (obj *_AllVirtualDefSubPartHistoryMgr) GetFromTenantID(tenantID int64) (results []*AllVirtualDefSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDefSubPartHistory{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllVirtualDefSubPartHistoryMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllVirtualDefSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDefSubPartHistory{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromTableID 通过table_id获取内容
func (obj *_AllVirtualDefSubPartHistoryMgr) GetFromTableID(tableID int64) (results []*AllVirtualDefSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDefSubPartHistory{}).Where("`table_id` = ?", tableID).Find(&results).Error

	return
}

// GetBatchFromTableID 批量查找
func (obj *_AllVirtualDefSubPartHistoryMgr) GetBatchFromTableID(tableIDs []int64) (results []*AllVirtualDefSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDefSubPartHistory{}).Where("`table_id` IN (?)", tableIDs).Find(&results).Error

	return
}

// GetFromSubPartID 通过sub_part_id获取内容
func (obj *_AllVirtualDefSubPartHistoryMgr) GetFromSubPartID(subPartID int64) (results []*AllVirtualDefSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDefSubPartHistory{}).Where("`sub_part_id` = ?", subPartID).Find(&results).Error

	return
}

// GetBatchFromSubPartID 批量查找
func (obj *_AllVirtualDefSubPartHistoryMgr) GetBatchFromSubPartID(subPartIDs []int64) (results []*AllVirtualDefSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDefSubPartHistory{}).Where("`sub_part_id` IN (?)", subPartIDs).Find(&results).Error

	return
}

// GetFromSchemaVersion 通过schema_version获取内容
func (obj *_AllVirtualDefSubPartHistoryMgr) GetFromSchemaVersion(schemaVersion int64) (results []*AllVirtualDefSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDefSubPartHistory{}).Where("`schema_version` = ?", schemaVersion).Find(&results).Error

	return
}

// GetBatchFromSchemaVersion 批量查找
func (obj *_AllVirtualDefSubPartHistoryMgr) GetBatchFromSchemaVersion(schemaVersions []int64) (results []*AllVirtualDefSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDefSubPartHistory{}).Where("`schema_version` IN (?)", schemaVersions).Find(&results).Error

	return
}

// GetFromGmtCreate 通过gmt_create获取内容
func (obj *_AllVirtualDefSubPartHistoryMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllVirtualDefSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDefSubPartHistory{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllVirtualDefSubPartHistoryMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllVirtualDefSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDefSubPartHistory{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllVirtualDefSubPartHistoryMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllVirtualDefSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDefSubPartHistory{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllVirtualDefSubPartHistoryMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllVirtualDefSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDefSubPartHistory{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromIsDeleted 通过is_deleted获取内容
func (obj *_AllVirtualDefSubPartHistoryMgr) GetFromIsDeleted(isDeleted int64) (results []*AllVirtualDefSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDefSubPartHistory{}).Where("`is_deleted` = ?", isDeleted).Find(&results).Error

	return
}

// GetBatchFromIsDeleted 批量查找
func (obj *_AllVirtualDefSubPartHistoryMgr) GetBatchFromIsDeleted(isDeleteds []int64) (results []*AllVirtualDefSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDefSubPartHistory{}).Where("`is_deleted` IN (?)", isDeleteds).Find(&results).Error

	return
}

// GetFromSubPartName 通过sub_part_name获取内容
func (obj *_AllVirtualDefSubPartHistoryMgr) GetFromSubPartName(subPartName string) (results []*AllVirtualDefSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDefSubPartHistory{}).Where("`sub_part_name` = ?", subPartName).Find(&results).Error

	return
}

// GetBatchFromSubPartName 批量查找
func (obj *_AllVirtualDefSubPartHistoryMgr) GetBatchFromSubPartName(subPartNames []string) (results []*AllVirtualDefSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDefSubPartHistory{}).Where("`sub_part_name` IN (?)", subPartNames).Find(&results).Error

	return
}

// GetFromHighBoundVal 通过high_bound_val获取内容
func (obj *_AllVirtualDefSubPartHistoryMgr) GetFromHighBoundVal(highBoundVal string) (results []*AllVirtualDefSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDefSubPartHistory{}).Where("`high_bound_val` = ?", highBoundVal).Find(&results).Error

	return
}

// GetBatchFromHighBoundVal 批量查找
func (obj *_AllVirtualDefSubPartHistoryMgr) GetBatchFromHighBoundVal(highBoundVals []string) (results []*AllVirtualDefSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDefSubPartHistory{}).Where("`high_bound_val` IN (?)", highBoundVals).Find(&results).Error

	return
}

// GetFromBHighBoundVal 通过b_high_bound_val获取内容
func (obj *_AllVirtualDefSubPartHistoryMgr) GetFromBHighBoundVal(bHighBoundVal string) (results []*AllVirtualDefSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDefSubPartHistory{}).Where("`b_high_bound_val` = ?", bHighBoundVal).Find(&results).Error

	return
}

// GetBatchFromBHighBoundVal 批量查找
func (obj *_AllVirtualDefSubPartHistoryMgr) GetBatchFromBHighBoundVal(bHighBoundVals []string) (results []*AllVirtualDefSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDefSubPartHistory{}).Where("`b_high_bound_val` IN (?)", bHighBoundVals).Find(&results).Error

	return
}

// GetFromBlockSize 通过block_size获取内容
func (obj *_AllVirtualDefSubPartHistoryMgr) GetFromBlockSize(blockSize int64) (results []*AllVirtualDefSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDefSubPartHistory{}).Where("`block_size` = ?", blockSize).Find(&results).Error

	return
}

// GetBatchFromBlockSize 批量查找
func (obj *_AllVirtualDefSubPartHistoryMgr) GetBatchFromBlockSize(blockSizes []int64) (results []*AllVirtualDefSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDefSubPartHistory{}).Where("`block_size` IN (?)", blockSizes).Find(&results).Error

	return
}

// GetFromReplicaNum 通过replica_num获取内容
func (obj *_AllVirtualDefSubPartHistoryMgr) GetFromReplicaNum(replicaNum int64) (results []*AllVirtualDefSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDefSubPartHistory{}).Where("`replica_num` = ?", replicaNum).Find(&results).Error

	return
}

// GetBatchFromReplicaNum 批量查找
func (obj *_AllVirtualDefSubPartHistoryMgr) GetBatchFromReplicaNum(replicaNums []int64) (results []*AllVirtualDefSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDefSubPartHistory{}).Where("`replica_num` IN (?)", replicaNums).Find(&results).Error

	return
}

// GetFromCompressFuncName 通过compress_func_name获取内容
func (obj *_AllVirtualDefSubPartHistoryMgr) GetFromCompressFuncName(compressFuncName string) (results []*AllVirtualDefSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDefSubPartHistory{}).Where("`compress_func_name` = ?", compressFuncName).Find(&results).Error

	return
}

// GetBatchFromCompressFuncName 批量查找
func (obj *_AllVirtualDefSubPartHistoryMgr) GetBatchFromCompressFuncName(compressFuncNames []string) (results []*AllVirtualDefSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDefSubPartHistory{}).Where("`compress_func_name` IN (?)", compressFuncNames).Find(&results).Error

	return
}

// GetFromSpare1 通过spare1获取内容
func (obj *_AllVirtualDefSubPartHistoryMgr) GetFromSpare1(spare1 int64) (results []*AllVirtualDefSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDefSubPartHistory{}).Where("`spare1` = ?", spare1).Find(&results).Error

	return
}

// GetBatchFromSpare1 批量查找
func (obj *_AllVirtualDefSubPartHistoryMgr) GetBatchFromSpare1(spare1s []int64) (results []*AllVirtualDefSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDefSubPartHistory{}).Where("`spare1` IN (?)", spare1s).Find(&results).Error

	return
}

// GetFromSpare2 通过spare2获取内容
func (obj *_AllVirtualDefSubPartHistoryMgr) GetFromSpare2(spare2 int64) (results []*AllVirtualDefSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDefSubPartHistory{}).Where("`spare2` = ?", spare2).Find(&results).Error

	return
}

// GetBatchFromSpare2 批量查找
func (obj *_AllVirtualDefSubPartHistoryMgr) GetBatchFromSpare2(spare2s []int64) (results []*AllVirtualDefSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDefSubPartHistory{}).Where("`spare2` IN (?)", spare2s).Find(&results).Error

	return
}

// GetFromSpare3 通过spare3获取内容
func (obj *_AllVirtualDefSubPartHistoryMgr) GetFromSpare3(spare3 string) (results []*AllVirtualDefSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDefSubPartHistory{}).Where("`spare3` = ?", spare3).Find(&results).Error

	return
}

// GetBatchFromSpare3 批量查找
func (obj *_AllVirtualDefSubPartHistoryMgr) GetBatchFromSpare3(spare3s []string) (results []*AllVirtualDefSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDefSubPartHistory{}).Where("`spare3` IN (?)", spare3s).Find(&results).Error

	return
}

// GetFromComment 通过comment获取内容
func (obj *_AllVirtualDefSubPartHistoryMgr) GetFromComment(comment string) (results []*AllVirtualDefSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDefSubPartHistory{}).Where("`comment` = ?", comment).Find(&results).Error

	return
}

// GetBatchFromComment 批量查找
func (obj *_AllVirtualDefSubPartHistoryMgr) GetBatchFromComment(comments []string) (results []*AllVirtualDefSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDefSubPartHistory{}).Where("`comment` IN (?)", comments).Find(&results).Error

	return
}

// GetFromListVal 通过list_val获取内容
func (obj *_AllVirtualDefSubPartHistoryMgr) GetFromListVal(listVal string) (results []*AllVirtualDefSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDefSubPartHistory{}).Where("`list_val` = ?", listVal).Find(&results).Error

	return
}

// GetBatchFromListVal 批量查找
func (obj *_AllVirtualDefSubPartHistoryMgr) GetBatchFromListVal(listVals []string) (results []*AllVirtualDefSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDefSubPartHistory{}).Where("`list_val` IN (?)", listVals).Find(&results).Error

	return
}

// GetFromBListVal 通过b_list_val获取内容
func (obj *_AllVirtualDefSubPartHistoryMgr) GetFromBListVal(bListVal string) (results []*AllVirtualDefSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDefSubPartHistory{}).Where("`b_list_val` = ?", bListVal).Find(&results).Error

	return
}

// GetBatchFromBListVal 批量查找
func (obj *_AllVirtualDefSubPartHistoryMgr) GetBatchFromBListVal(bListVals []string) (results []*AllVirtualDefSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDefSubPartHistory{}).Where("`b_list_val` IN (?)", bListVals).Find(&results).Error

	return
}

// GetFromSubPartIDx 通过sub_part_idx获取内容
func (obj *_AllVirtualDefSubPartHistoryMgr) GetFromSubPartIDx(subPartIDx int64) (results []*AllVirtualDefSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDefSubPartHistory{}).Where("`sub_part_idx` = ?", subPartIDx).Find(&results).Error

	return
}

// GetBatchFromSubPartIDx 批量查找
func (obj *_AllVirtualDefSubPartHistoryMgr) GetBatchFromSubPartIDx(subPartIDxs []int64) (results []*AllVirtualDefSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDefSubPartHistory{}).Where("`sub_part_idx` IN (?)", subPartIDxs).Find(&results).Error

	return
}

// GetFromSourcePartitionID 通过source_partition_id获取内容
func (obj *_AllVirtualDefSubPartHistoryMgr) GetFromSourcePartitionID(sourcePartitionID string) (results []*AllVirtualDefSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDefSubPartHistory{}).Where("`source_partition_id` = ?", sourcePartitionID).Find(&results).Error

	return
}

// GetBatchFromSourcePartitionID 批量查找
func (obj *_AllVirtualDefSubPartHistoryMgr) GetBatchFromSourcePartitionID(sourcePartitionIDs []string) (results []*AllVirtualDefSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDefSubPartHistory{}).Where("`source_partition_id` IN (?)", sourcePartitionIDs).Find(&results).Error

	return
}

// GetFromMappingPgSubPartID 通过mapping_pg_sub_part_id获取内容
func (obj *_AllVirtualDefSubPartHistoryMgr) GetFromMappingPgSubPartID(mappingPgSubPartID int64) (results []*AllVirtualDefSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDefSubPartHistory{}).Where("`mapping_pg_sub_part_id` = ?", mappingPgSubPartID).Find(&results).Error

	return
}

// GetBatchFromMappingPgSubPartID 批量查找
func (obj *_AllVirtualDefSubPartHistoryMgr) GetBatchFromMappingPgSubPartID(mappingPgSubPartIDs []int64) (results []*AllVirtualDefSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDefSubPartHistory{}).Where("`mapping_pg_sub_part_id` IN (?)", mappingPgSubPartIDs).Find(&results).Error

	return
}

// GetFromTablespaceID 通过tablespace_id获取内容
func (obj *_AllVirtualDefSubPartHistoryMgr) GetFromTablespaceID(tablespaceID int64) (results []*AllVirtualDefSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDefSubPartHistory{}).Where("`tablespace_id` = ?", tablespaceID).Find(&results).Error

	return
}

// GetBatchFromTablespaceID 批量查找
func (obj *_AllVirtualDefSubPartHistoryMgr) GetBatchFromTablespaceID(tablespaceIDs []int64) (results []*AllVirtualDefSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDefSubPartHistory{}).Where("`tablespace_id` IN (?)", tablespaceIDs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllVirtualDefSubPartHistoryMgr) FetchByPrimaryKey(tenantID int64, tableID int64, subPartID int64, schemaVersion int64) (result AllVirtualDefSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDefSubPartHistory{}).Where("`tenant_id` = ? AND `table_id` = ? AND `sub_part_id` = ? AND `schema_version` = ?", tenantID, tableID, subPartID, schemaVersion).First(&result).Error

	return
}
