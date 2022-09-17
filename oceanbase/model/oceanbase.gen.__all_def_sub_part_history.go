package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _AllDefSubPartHistoryMgr struct {
	*_BaseMgr
}

// AllDefSubPartHistoryMgr open func
func AllDefSubPartHistoryMgr(db *gorm.DB) *_AllDefSubPartHistoryMgr {
	if db == nil {
		panic(fmt.Errorf("AllDefSubPartHistoryMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllDefSubPartHistoryMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_def_sub_part_history"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllDefSubPartHistoryMgr) GetTableName() string {
	return "__all_def_sub_part_history"
}

// Reset 重置gorm会话
func (obj *_AllDefSubPartHistoryMgr) Reset() *_AllDefSubPartHistoryMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllDefSubPartHistoryMgr) Get() (result AllDefSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDefSubPartHistory{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllDefSubPartHistoryMgr) Gets() (results []*AllDefSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDefSubPartHistory{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllDefSubPartHistoryMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllDefSubPartHistory{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithGmtCreate gmt_create获取
func (obj *_AllDefSubPartHistoryMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllDefSubPartHistoryMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithTenantID tenant_id获取
func (obj *_AllDefSubPartHistoryMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithTableID table_id获取
func (obj *_AllDefSubPartHistoryMgr) WithTableID(tableID int64) Option {
	return optionFunc(func(o *options) { o.query["table_id"] = tableID })
}

// WithSubPartID sub_part_id获取
func (obj *_AllDefSubPartHistoryMgr) WithSubPartID(subPartID int64) Option {
	return optionFunc(func(o *options) { o.query["sub_part_id"] = subPartID })
}

// WithSchemaVersion schema_version获取
func (obj *_AllDefSubPartHistoryMgr) WithSchemaVersion(schemaVersion int64) Option {
	return optionFunc(func(o *options) { o.query["schema_version"] = schemaVersion })
}

// WithIsDeleted is_deleted获取
func (obj *_AllDefSubPartHistoryMgr) WithIsDeleted(isDeleted int64) Option {
	return optionFunc(func(o *options) { o.query["is_deleted"] = isDeleted })
}

// WithSubPartName sub_part_name获取
func (obj *_AllDefSubPartHistoryMgr) WithSubPartName(subPartName string) Option {
	return optionFunc(func(o *options) { o.query["sub_part_name"] = subPartName })
}

// WithHighBoundVal high_bound_val获取
func (obj *_AllDefSubPartHistoryMgr) WithHighBoundVal(highBoundVal string) Option {
	return optionFunc(func(o *options) { o.query["high_bound_val"] = highBoundVal })
}

// WithBHighBoundVal b_high_bound_val获取
func (obj *_AllDefSubPartHistoryMgr) WithBHighBoundVal(bHighBoundVal string) Option {
	return optionFunc(func(o *options) { o.query["b_high_bound_val"] = bHighBoundVal })
}

// WithBlockSize block_size获取
func (obj *_AllDefSubPartHistoryMgr) WithBlockSize(blockSize int64) Option {
	return optionFunc(func(o *options) { o.query["block_size"] = blockSize })
}

// WithReplicaNum replica_num获取
func (obj *_AllDefSubPartHistoryMgr) WithReplicaNum(replicaNum int64) Option {
	return optionFunc(func(o *options) { o.query["replica_num"] = replicaNum })
}

// WithCompressFuncName compress_func_name获取
func (obj *_AllDefSubPartHistoryMgr) WithCompressFuncName(compressFuncName string) Option {
	return optionFunc(func(o *options) { o.query["compress_func_name"] = compressFuncName })
}

// WithSpare1 spare1获取
func (obj *_AllDefSubPartHistoryMgr) WithSpare1(spare1 int64) Option {
	return optionFunc(func(o *options) { o.query["spare1"] = spare1 })
}

// WithSpare2 spare2获取
func (obj *_AllDefSubPartHistoryMgr) WithSpare2(spare2 int64) Option {
	return optionFunc(func(o *options) { o.query["spare2"] = spare2 })
}

// WithSpare3 spare3获取
func (obj *_AllDefSubPartHistoryMgr) WithSpare3(spare3 string) Option {
	return optionFunc(func(o *options) { o.query["spare3"] = spare3 })
}

// WithComment comment获取
func (obj *_AllDefSubPartHistoryMgr) WithComment(comment string) Option {
	return optionFunc(func(o *options) { o.query["comment"] = comment })
}

// WithListVal list_val获取
func (obj *_AllDefSubPartHistoryMgr) WithListVal(listVal string) Option {
	return optionFunc(func(o *options) { o.query["list_val"] = listVal })
}

// WithBListVal b_list_val获取
func (obj *_AllDefSubPartHistoryMgr) WithBListVal(bListVal string) Option {
	return optionFunc(func(o *options) { o.query["b_list_val"] = bListVal })
}

// WithSubPartIDx sub_part_idx获取
func (obj *_AllDefSubPartHistoryMgr) WithSubPartIDx(subPartIDx int64) Option {
	return optionFunc(func(o *options) { o.query["sub_part_idx"] = subPartIDx })
}

// WithSourcePartitionID source_partition_id获取
func (obj *_AllDefSubPartHistoryMgr) WithSourcePartitionID(sourcePartitionID string) Option {
	return optionFunc(func(o *options) { o.query["source_partition_id"] = sourcePartitionID })
}

// WithMappingPgSubPartID mapping_pg_sub_part_id获取
func (obj *_AllDefSubPartHistoryMgr) WithMappingPgSubPartID(mappingPgSubPartID int64) Option {
	return optionFunc(func(o *options) { o.query["mapping_pg_sub_part_id"] = mappingPgSubPartID })
}

// WithTablespaceID tablespace_id获取
func (obj *_AllDefSubPartHistoryMgr) WithTablespaceID(tablespaceID int64) Option {
	return optionFunc(func(o *options) { o.query["tablespace_id"] = tablespaceID })
}

// GetByOption 功能选项模式获取
func (obj *_AllDefSubPartHistoryMgr) GetByOption(opts ...Option) (result AllDefSubPartHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllDefSubPartHistory{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllDefSubPartHistoryMgr) GetByOptions(opts ...Option) (results []*AllDefSubPartHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllDefSubPartHistory{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllDefSubPartHistoryMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllDefSubPartHistory, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllDefSubPartHistory{}).Where(options.query)
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
func (obj *_AllDefSubPartHistoryMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllDefSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDefSubPartHistory{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllDefSubPartHistoryMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllDefSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDefSubPartHistory{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllDefSubPartHistoryMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllDefSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDefSubPartHistory{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllDefSubPartHistoryMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllDefSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDefSubPartHistory{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromTenantID 通过tenant_id获取内容
func (obj *_AllDefSubPartHistoryMgr) GetFromTenantID(tenantID int64) (results []*AllDefSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDefSubPartHistory{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllDefSubPartHistoryMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllDefSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDefSubPartHistory{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromTableID 通过table_id获取内容
func (obj *_AllDefSubPartHistoryMgr) GetFromTableID(tableID int64) (results []*AllDefSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDefSubPartHistory{}).Where("`table_id` = ?", tableID).Find(&results).Error

	return
}

// GetBatchFromTableID 批量查找
func (obj *_AllDefSubPartHistoryMgr) GetBatchFromTableID(tableIDs []int64) (results []*AllDefSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDefSubPartHistory{}).Where("`table_id` IN (?)", tableIDs).Find(&results).Error

	return
}

// GetFromSubPartID 通过sub_part_id获取内容
func (obj *_AllDefSubPartHistoryMgr) GetFromSubPartID(subPartID int64) (results []*AllDefSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDefSubPartHistory{}).Where("`sub_part_id` = ?", subPartID).Find(&results).Error

	return
}

// GetBatchFromSubPartID 批量查找
func (obj *_AllDefSubPartHistoryMgr) GetBatchFromSubPartID(subPartIDs []int64) (results []*AllDefSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDefSubPartHistory{}).Where("`sub_part_id` IN (?)", subPartIDs).Find(&results).Error

	return
}

// GetFromSchemaVersion 通过schema_version获取内容
func (obj *_AllDefSubPartHistoryMgr) GetFromSchemaVersion(schemaVersion int64) (results []*AllDefSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDefSubPartHistory{}).Where("`schema_version` = ?", schemaVersion).Find(&results).Error

	return
}

// GetBatchFromSchemaVersion 批量查找
func (obj *_AllDefSubPartHistoryMgr) GetBatchFromSchemaVersion(schemaVersions []int64) (results []*AllDefSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDefSubPartHistory{}).Where("`schema_version` IN (?)", schemaVersions).Find(&results).Error

	return
}

// GetFromIsDeleted 通过is_deleted获取内容
func (obj *_AllDefSubPartHistoryMgr) GetFromIsDeleted(isDeleted int64) (results []*AllDefSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDefSubPartHistory{}).Where("`is_deleted` = ?", isDeleted).Find(&results).Error

	return
}

// GetBatchFromIsDeleted 批量查找
func (obj *_AllDefSubPartHistoryMgr) GetBatchFromIsDeleted(isDeleteds []int64) (results []*AllDefSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDefSubPartHistory{}).Where("`is_deleted` IN (?)", isDeleteds).Find(&results).Error

	return
}

// GetFromSubPartName 通过sub_part_name获取内容
func (obj *_AllDefSubPartHistoryMgr) GetFromSubPartName(subPartName string) (results []*AllDefSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDefSubPartHistory{}).Where("`sub_part_name` = ?", subPartName).Find(&results).Error

	return
}

// GetBatchFromSubPartName 批量查找
func (obj *_AllDefSubPartHistoryMgr) GetBatchFromSubPartName(subPartNames []string) (results []*AllDefSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDefSubPartHistory{}).Where("`sub_part_name` IN (?)", subPartNames).Find(&results).Error

	return
}

// GetFromHighBoundVal 通过high_bound_val获取内容
func (obj *_AllDefSubPartHistoryMgr) GetFromHighBoundVal(highBoundVal string) (results []*AllDefSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDefSubPartHistory{}).Where("`high_bound_val` = ?", highBoundVal).Find(&results).Error

	return
}

// GetBatchFromHighBoundVal 批量查找
func (obj *_AllDefSubPartHistoryMgr) GetBatchFromHighBoundVal(highBoundVals []string) (results []*AllDefSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDefSubPartHistory{}).Where("`high_bound_val` IN (?)", highBoundVals).Find(&results).Error

	return
}

// GetFromBHighBoundVal 通过b_high_bound_val获取内容
func (obj *_AllDefSubPartHistoryMgr) GetFromBHighBoundVal(bHighBoundVal string) (results []*AllDefSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDefSubPartHistory{}).Where("`b_high_bound_val` = ?", bHighBoundVal).Find(&results).Error

	return
}

// GetBatchFromBHighBoundVal 批量查找
func (obj *_AllDefSubPartHistoryMgr) GetBatchFromBHighBoundVal(bHighBoundVals []string) (results []*AllDefSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDefSubPartHistory{}).Where("`b_high_bound_val` IN (?)", bHighBoundVals).Find(&results).Error

	return
}

// GetFromBlockSize 通过block_size获取内容
func (obj *_AllDefSubPartHistoryMgr) GetFromBlockSize(blockSize int64) (results []*AllDefSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDefSubPartHistory{}).Where("`block_size` = ?", blockSize).Find(&results).Error

	return
}

// GetBatchFromBlockSize 批量查找
func (obj *_AllDefSubPartHistoryMgr) GetBatchFromBlockSize(blockSizes []int64) (results []*AllDefSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDefSubPartHistory{}).Where("`block_size` IN (?)", blockSizes).Find(&results).Error

	return
}

// GetFromReplicaNum 通过replica_num获取内容
func (obj *_AllDefSubPartHistoryMgr) GetFromReplicaNum(replicaNum int64) (results []*AllDefSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDefSubPartHistory{}).Where("`replica_num` = ?", replicaNum).Find(&results).Error

	return
}

// GetBatchFromReplicaNum 批量查找
func (obj *_AllDefSubPartHistoryMgr) GetBatchFromReplicaNum(replicaNums []int64) (results []*AllDefSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDefSubPartHistory{}).Where("`replica_num` IN (?)", replicaNums).Find(&results).Error

	return
}

// GetFromCompressFuncName 通过compress_func_name获取内容
func (obj *_AllDefSubPartHistoryMgr) GetFromCompressFuncName(compressFuncName string) (results []*AllDefSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDefSubPartHistory{}).Where("`compress_func_name` = ?", compressFuncName).Find(&results).Error

	return
}

// GetBatchFromCompressFuncName 批量查找
func (obj *_AllDefSubPartHistoryMgr) GetBatchFromCompressFuncName(compressFuncNames []string) (results []*AllDefSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDefSubPartHistory{}).Where("`compress_func_name` IN (?)", compressFuncNames).Find(&results).Error

	return
}

// GetFromSpare1 通过spare1获取内容
func (obj *_AllDefSubPartHistoryMgr) GetFromSpare1(spare1 int64) (results []*AllDefSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDefSubPartHistory{}).Where("`spare1` = ?", spare1).Find(&results).Error

	return
}

// GetBatchFromSpare1 批量查找
func (obj *_AllDefSubPartHistoryMgr) GetBatchFromSpare1(spare1s []int64) (results []*AllDefSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDefSubPartHistory{}).Where("`spare1` IN (?)", spare1s).Find(&results).Error

	return
}

// GetFromSpare2 通过spare2获取内容
func (obj *_AllDefSubPartHistoryMgr) GetFromSpare2(spare2 int64) (results []*AllDefSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDefSubPartHistory{}).Where("`spare2` = ?", spare2).Find(&results).Error

	return
}

// GetBatchFromSpare2 批量查找
func (obj *_AllDefSubPartHistoryMgr) GetBatchFromSpare2(spare2s []int64) (results []*AllDefSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDefSubPartHistory{}).Where("`spare2` IN (?)", spare2s).Find(&results).Error

	return
}

// GetFromSpare3 通过spare3获取内容
func (obj *_AllDefSubPartHistoryMgr) GetFromSpare3(spare3 string) (results []*AllDefSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDefSubPartHistory{}).Where("`spare3` = ?", spare3).Find(&results).Error

	return
}

// GetBatchFromSpare3 批量查找
func (obj *_AllDefSubPartHistoryMgr) GetBatchFromSpare3(spare3s []string) (results []*AllDefSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDefSubPartHistory{}).Where("`spare3` IN (?)", spare3s).Find(&results).Error

	return
}

// GetFromComment 通过comment获取内容
func (obj *_AllDefSubPartHistoryMgr) GetFromComment(comment string) (results []*AllDefSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDefSubPartHistory{}).Where("`comment` = ?", comment).Find(&results).Error

	return
}

// GetBatchFromComment 批量查找
func (obj *_AllDefSubPartHistoryMgr) GetBatchFromComment(comments []string) (results []*AllDefSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDefSubPartHistory{}).Where("`comment` IN (?)", comments).Find(&results).Error

	return
}

// GetFromListVal 通过list_val获取内容
func (obj *_AllDefSubPartHistoryMgr) GetFromListVal(listVal string) (results []*AllDefSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDefSubPartHistory{}).Where("`list_val` = ?", listVal).Find(&results).Error

	return
}

// GetBatchFromListVal 批量查找
func (obj *_AllDefSubPartHistoryMgr) GetBatchFromListVal(listVals []string) (results []*AllDefSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDefSubPartHistory{}).Where("`list_val` IN (?)", listVals).Find(&results).Error

	return
}

// GetFromBListVal 通过b_list_val获取内容
func (obj *_AllDefSubPartHistoryMgr) GetFromBListVal(bListVal string) (results []*AllDefSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDefSubPartHistory{}).Where("`b_list_val` = ?", bListVal).Find(&results).Error

	return
}

// GetBatchFromBListVal 批量查找
func (obj *_AllDefSubPartHistoryMgr) GetBatchFromBListVal(bListVals []string) (results []*AllDefSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDefSubPartHistory{}).Where("`b_list_val` IN (?)", bListVals).Find(&results).Error

	return
}

// GetFromSubPartIDx 通过sub_part_idx获取内容
func (obj *_AllDefSubPartHistoryMgr) GetFromSubPartIDx(subPartIDx int64) (results []*AllDefSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDefSubPartHistory{}).Where("`sub_part_idx` = ?", subPartIDx).Find(&results).Error

	return
}

// GetBatchFromSubPartIDx 批量查找
func (obj *_AllDefSubPartHistoryMgr) GetBatchFromSubPartIDx(subPartIDxs []int64) (results []*AllDefSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDefSubPartHistory{}).Where("`sub_part_idx` IN (?)", subPartIDxs).Find(&results).Error

	return
}

// GetFromSourcePartitionID 通过source_partition_id获取内容
func (obj *_AllDefSubPartHistoryMgr) GetFromSourcePartitionID(sourcePartitionID string) (results []*AllDefSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDefSubPartHistory{}).Where("`source_partition_id` = ?", sourcePartitionID).Find(&results).Error

	return
}

// GetBatchFromSourcePartitionID 批量查找
func (obj *_AllDefSubPartHistoryMgr) GetBatchFromSourcePartitionID(sourcePartitionIDs []string) (results []*AllDefSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDefSubPartHistory{}).Where("`source_partition_id` IN (?)", sourcePartitionIDs).Find(&results).Error

	return
}

// GetFromMappingPgSubPartID 通过mapping_pg_sub_part_id获取内容
func (obj *_AllDefSubPartHistoryMgr) GetFromMappingPgSubPartID(mappingPgSubPartID int64) (results []*AllDefSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDefSubPartHistory{}).Where("`mapping_pg_sub_part_id` = ?", mappingPgSubPartID).Find(&results).Error

	return
}

// GetBatchFromMappingPgSubPartID 批量查找
func (obj *_AllDefSubPartHistoryMgr) GetBatchFromMappingPgSubPartID(mappingPgSubPartIDs []int64) (results []*AllDefSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDefSubPartHistory{}).Where("`mapping_pg_sub_part_id` IN (?)", mappingPgSubPartIDs).Find(&results).Error

	return
}

// GetFromTablespaceID 通过tablespace_id获取内容
func (obj *_AllDefSubPartHistoryMgr) GetFromTablespaceID(tablespaceID int64) (results []*AllDefSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDefSubPartHistory{}).Where("`tablespace_id` = ?", tablespaceID).Find(&results).Error

	return
}

// GetBatchFromTablespaceID 批量查找
func (obj *_AllDefSubPartHistoryMgr) GetBatchFromTablespaceID(tablespaceIDs []int64) (results []*AllDefSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDefSubPartHistory{}).Where("`tablespace_id` IN (?)", tablespaceIDs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllDefSubPartHistoryMgr) FetchByPrimaryKey(tenantID int64, tableID int64, subPartID int64, schemaVersion int64) (result AllDefSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDefSubPartHistory{}).Where("`tenant_id` = ? AND `table_id` = ? AND `sub_part_id` = ? AND `schema_version` = ?", tenantID, tableID, subPartID, schemaVersion).First(&result).Error

	return
}
