package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _AllVirtualDefSubPartMgr struct {
	*_BaseMgr
}

// AllVirtualDefSubPartMgr open func
func AllVirtualDefSubPartMgr(db *gorm.DB) *_AllVirtualDefSubPartMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualDefSubPartMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualDefSubPartMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_def_sub_part"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualDefSubPartMgr) GetTableName() string {
	return "__all_virtual_def_sub_part"
}

// Reset 重置gorm会话
func (obj *_AllVirtualDefSubPartMgr) Reset() *_AllVirtualDefSubPartMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualDefSubPartMgr) Get() (result AllVirtualDefSubPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDefSubPart{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualDefSubPartMgr) Gets() (results []*AllVirtualDefSubPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDefSubPart{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualDefSubPartMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualDefSubPart{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTenantID tenant_id获取
func (obj *_AllVirtualDefSubPartMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithTableID table_id获取
func (obj *_AllVirtualDefSubPartMgr) WithTableID(tableID int64) Option {
	return optionFunc(func(o *options) { o.query["table_id"] = tableID })
}

// WithSubPartID sub_part_id获取
func (obj *_AllVirtualDefSubPartMgr) WithSubPartID(subPartID int64) Option {
	return optionFunc(func(o *options) { o.query["sub_part_id"] = subPartID })
}

// WithGmtCreate gmt_create获取
func (obj *_AllVirtualDefSubPartMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllVirtualDefSubPartMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithSubPartName sub_part_name获取
func (obj *_AllVirtualDefSubPartMgr) WithSubPartName(subPartName string) Option {
	return optionFunc(func(o *options) { o.query["sub_part_name"] = subPartName })
}

// WithSchemaVersion schema_version获取
func (obj *_AllVirtualDefSubPartMgr) WithSchemaVersion(schemaVersion int64) Option {
	return optionFunc(func(o *options) { o.query["schema_version"] = schemaVersion })
}

// WithHighBoundVal high_bound_val获取
func (obj *_AllVirtualDefSubPartMgr) WithHighBoundVal(highBoundVal string) Option {
	return optionFunc(func(o *options) { o.query["high_bound_val"] = highBoundVal })
}

// WithBHighBoundVal b_high_bound_val获取
func (obj *_AllVirtualDefSubPartMgr) WithBHighBoundVal(bHighBoundVal string) Option {
	return optionFunc(func(o *options) { o.query["b_high_bound_val"] = bHighBoundVal })
}

// WithBlockSize block_size获取
func (obj *_AllVirtualDefSubPartMgr) WithBlockSize(blockSize int64) Option {
	return optionFunc(func(o *options) { o.query["block_size"] = blockSize })
}

// WithReplicaNum replica_num获取
func (obj *_AllVirtualDefSubPartMgr) WithReplicaNum(replicaNum int64) Option {
	return optionFunc(func(o *options) { o.query["replica_num"] = replicaNum })
}

// WithCompressFuncName compress_func_name获取
func (obj *_AllVirtualDefSubPartMgr) WithCompressFuncName(compressFuncName string) Option {
	return optionFunc(func(o *options) { o.query["compress_func_name"] = compressFuncName })
}

// WithSpare1 spare1获取
func (obj *_AllVirtualDefSubPartMgr) WithSpare1(spare1 int64) Option {
	return optionFunc(func(o *options) { o.query["spare1"] = spare1 })
}

// WithSpare2 spare2获取
func (obj *_AllVirtualDefSubPartMgr) WithSpare2(spare2 int64) Option {
	return optionFunc(func(o *options) { o.query["spare2"] = spare2 })
}

// WithSpare3 spare3获取
func (obj *_AllVirtualDefSubPartMgr) WithSpare3(spare3 string) Option {
	return optionFunc(func(o *options) { o.query["spare3"] = spare3 })
}

// WithComment comment获取
func (obj *_AllVirtualDefSubPartMgr) WithComment(comment string) Option {
	return optionFunc(func(o *options) { o.query["comment"] = comment })
}

// WithListVal list_val获取
func (obj *_AllVirtualDefSubPartMgr) WithListVal(listVal string) Option {
	return optionFunc(func(o *options) { o.query["list_val"] = listVal })
}

// WithBListVal b_list_val获取
func (obj *_AllVirtualDefSubPartMgr) WithBListVal(bListVal string) Option {
	return optionFunc(func(o *options) { o.query["b_list_val"] = bListVal })
}

// WithSubPartIDx sub_part_idx获取
func (obj *_AllVirtualDefSubPartMgr) WithSubPartIDx(subPartIDx int64) Option {
	return optionFunc(func(o *options) { o.query["sub_part_idx"] = subPartIDx })
}

// WithSourcePartitionID source_partition_id获取
func (obj *_AllVirtualDefSubPartMgr) WithSourcePartitionID(sourcePartitionID string) Option {
	return optionFunc(func(o *options) { o.query["source_partition_id"] = sourcePartitionID })
}

// WithMappingPgSubPartID mapping_pg_sub_part_id获取
func (obj *_AllVirtualDefSubPartMgr) WithMappingPgSubPartID(mappingPgSubPartID int64) Option {
	return optionFunc(func(o *options) { o.query["mapping_pg_sub_part_id"] = mappingPgSubPartID })
}

// WithTablespaceID tablespace_id获取
func (obj *_AllVirtualDefSubPartMgr) WithTablespaceID(tablespaceID int64) Option {
	return optionFunc(func(o *options) { o.query["tablespace_id"] = tablespaceID })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualDefSubPartMgr) GetByOption(opts ...Option) (result AllVirtualDefSubPart, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDefSubPart{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualDefSubPartMgr) GetByOptions(opts ...Option) (results []*AllVirtualDefSubPart, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDefSubPart{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualDefSubPartMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualDefSubPart, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualDefSubPart{}).Where(options.query)
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
func (obj *_AllVirtualDefSubPartMgr) GetFromTenantID(tenantID int64) (results []*AllVirtualDefSubPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDefSubPart{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllVirtualDefSubPartMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllVirtualDefSubPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDefSubPart{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromTableID 通过table_id获取内容
func (obj *_AllVirtualDefSubPartMgr) GetFromTableID(tableID int64) (results []*AllVirtualDefSubPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDefSubPart{}).Where("`table_id` = ?", tableID).Find(&results).Error

	return
}

// GetBatchFromTableID 批量查找
func (obj *_AllVirtualDefSubPartMgr) GetBatchFromTableID(tableIDs []int64) (results []*AllVirtualDefSubPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDefSubPart{}).Where("`table_id` IN (?)", tableIDs).Find(&results).Error

	return
}

// GetFromSubPartID 通过sub_part_id获取内容
func (obj *_AllVirtualDefSubPartMgr) GetFromSubPartID(subPartID int64) (results []*AllVirtualDefSubPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDefSubPart{}).Where("`sub_part_id` = ?", subPartID).Find(&results).Error

	return
}

// GetBatchFromSubPartID 批量查找
func (obj *_AllVirtualDefSubPartMgr) GetBatchFromSubPartID(subPartIDs []int64) (results []*AllVirtualDefSubPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDefSubPart{}).Where("`sub_part_id` IN (?)", subPartIDs).Find(&results).Error

	return
}

// GetFromGmtCreate 通过gmt_create获取内容
func (obj *_AllVirtualDefSubPartMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllVirtualDefSubPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDefSubPart{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllVirtualDefSubPartMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllVirtualDefSubPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDefSubPart{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllVirtualDefSubPartMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllVirtualDefSubPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDefSubPart{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllVirtualDefSubPartMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllVirtualDefSubPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDefSubPart{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromSubPartName 通过sub_part_name获取内容
func (obj *_AllVirtualDefSubPartMgr) GetFromSubPartName(subPartName string) (results []*AllVirtualDefSubPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDefSubPart{}).Where("`sub_part_name` = ?", subPartName).Find(&results).Error

	return
}

// GetBatchFromSubPartName 批量查找
func (obj *_AllVirtualDefSubPartMgr) GetBatchFromSubPartName(subPartNames []string) (results []*AllVirtualDefSubPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDefSubPart{}).Where("`sub_part_name` IN (?)", subPartNames).Find(&results).Error

	return
}

// GetFromSchemaVersion 通过schema_version获取内容
func (obj *_AllVirtualDefSubPartMgr) GetFromSchemaVersion(schemaVersion int64) (results []*AllVirtualDefSubPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDefSubPart{}).Where("`schema_version` = ?", schemaVersion).Find(&results).Error

	return
}

// GetBatchFromSchemaVersion 批量查找
func (obj *_AllVirtualDefSubPartMgr) GetBatchFromSchemaVersion(schemaVersions []int64) (results []*AllVirtualDefSubPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDefSubPart{}).Where("`schema_version` IN (?)", schemaVersions).Find(&results).Error

	return
}

// GetFromHighBoundVal 通过high_bound_val获取内容
func (obj *_AllVirtualDefSubPartMgr) GetFromHighBoundVal(highBoundVal string) (results []*AllVirtualDefSubPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDefSubPart{}).Where("`high_bound_val` = ?", highBoundVal).Find(&results).Error

	return
}

// GetBatchFromHighBoundVal 批量查找
func (obj *_AllVirtualDefSubPartMgr) GetBatchFromHighBoundVal(highBoundVals []string) (results []*AllVirtualDefSubPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDefSubPart{}).Where("`high_bound_val` IN (?)", highBoundVals).Find(&results).Error

	return
}

// GetFromBHighBoundVal 通过b_high_bound_val获取内容
func (obj *_AllVirtualDefSubPartMgr) GetFromBHighBoundVal(bHighBoundVal string) (results []*AllVirtualDefSubPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDefSubPart{}).Where("`b_high_bound_val` = ?", bHighBoundVal).Find(&results).Error

	return
}

// GetBatchFromBHighBoundVal 批量查找
func (obj *_AllVirtualDefSubPartMgr) GetBatchFromBHighBoundVal(bHighBoundVals []string) (results []*AllVirtualDefSubPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDefSubPart{}).Where("`b_high_bound_val` IN (?)", bHighBoundVals).Find(&results).Error

	return
}

// GetFromBlockSize 通过block_size获取内容
func (obj *_AllVirtualDefSubPartMgr) GetFromBlockSize(blockSize int64) (results []*AllVirtualDefSubPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDefSubPart{}).Where("`block_size` = ?", blockSize).Find(&results).Error

	return
}

// GetBatchFromBlockSize 批量查找
func (obj *_AllVirtualDefSubPartMgr) GetBatchFromBlockSize(blockSizes []int64) (results []*AllVirtualDefSubPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDefSubPart{}).Where("`block_size` IN (?)", blockSizes).Find(&results).Error

	return
}

// GetFromReplicaNum 通过replica_num获取内容
func (obj *_AllVirtualDefSubPartMgr) GetFromReplicaNum(replicaNum int64) (results []*AllVirtualDefSubPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDefSubPart{}).Where("`replica_num` = ?", replicaNum).Find(&results).Error

	return
}

// GetBatchFromReplicaNum 批量查找
func (obj *_AllVirtualDefSubPartMgr) GetBatchFromReplicaNum(replicaNums []int64) (results []*AllVirtualDefSubPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDefSubPart{}).Where("`replica_num` IN (?)", replicaNums).Find(&results).Error

	return
}

// GetFromCompressFuncName 通过compress_func_name获取内容
func (obj *_AllVirtualDefSubPartMgr) GetFromCompressFuncName(compressFuncName string) (results []*AllVirtualDefSubPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDefSubPart{}).Where("`compress_func_name` = ?", compressFuncName).Find(&results).Error

	return
}

// GetBatchFromCompressFuncName 批量查找
func (obj *_AllVirtualDefSubPartMgr) GetBatchFromCompressFuncName(compressFuncNames []string) (results []*AllVirtualDefSubPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDefSubPart{}).Where("`compress_func_name` IN (?)", compressFuncNames).Find(&results).Error

	return
}

// GetFromSpare1 通过spare1获取内容
func (obj *_AllVirtualDefSubPartMgr) GetFromSpare1(spare1 int64) (results []*AllVirtualDefSubPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDefSubPart{}).Where("`spare1` = ?", spare1).Find(&results).Error

	return
}

// GetBatchFromSpare1 批量查找
func (obj *_AllVirtualDefSubPartMgr) GetBatchFromSpare1(spare1s []int64) (results []*AllVirtualDefSubPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDefSubPart{}).Where("`spare1` IN (?)", spare1s).Find(&results).Error

	return
}

// GetFromSpare2 通过spare2获取内容
func (obj *_AllVirtualDefSubPartMgr) GetFromSpare2(spare2 int64) (results []*AllVirtualDefSubPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDefSubPart{}).Where("`spare2` = ?", spare2).Find(&results).Error

	return
}

// GetBatchFromSpare2 批量查找
func (obj *_AllVirtualDefSubPartMgr) GetBatchFromSpare2(spare2s []int64) (results []*AllVirtualDefSubPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDefSubPart{}).Where("`spare2` IN (?)", spare2s).Find(&results).Error

	return
}

// GetFromSpare3 通过spare3获取内容
func (obj *_AllVirtualDefSubPartMgr) GetFromSpare3(spare3 string) (results []*AllVirtualDefSubPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDefSubPart{}).Where("`spare3` = ?", spare3).Find(&results).Error

	return
}

// GetBatchFromSpare3 批量查找
func (obj *_AllVirtualDefSubPartMgr) GetBatchFromSpare3(spare3s []string) (results []*AllVirtualDefSubPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDefSubPart{}).Where("`spare3` IN (?)", spare3s).Find(&results).Error

	return
}

// GetFromComment 通过comment获取内容
func (obj *_AllVirtualDefSubPartMgr) GetFromComment(comment string) (results []*AllVirtualDefSubPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDefSubPart{}).Where("`comment` = ?", comment).Find(&results).Error

	return
}

// GetBatchFromComment 批量查找
func (obj *_AllVirtualDefSubPartMgr) GetBatchFromComment(comments []string) (results []*AllVirtualDefSubPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDefSubPart{}).Where("`comment` IN (?)", comments).Find(&results).Error

	return
}

// GetFromListVal 通过list_val获取内容
func (obj *_AllVirtualDefSubPartMgr) GetFromListVal(listVal string) (results []*AllVirtualDefSubPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDefSubPart{}).Where("`list_val` = ?", listVal).Find(&results).Error

	return
}

// GetBatchFromListVal 批量查找
func (obj *_AllVirtualDefSubPartMgr) GetBatchFromListVal(listVals []string) (results []*AllVirtualDefSubPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDefSubPart{}).Where("`list_val` IN (?)", listVals).Find(&results).Error

	return
}

// GetFromBListVal 通过b_list_val获取内容
func (obj *_AllVirtualDefSubPartMgr) GetFromBListVal(bListVal string) (results []*AllVirtualDefSubPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDefSubPart{}).Where("`b_list_val` = ?", bListVal).Find(&results).Error

	return
}

// GetBatchFromBListVal 批量查找
func (obj *_AllVirtualDefSubPartMgr) GetBatchFromBListVal(bListVals []string) (results []*AllVirtualDefSubPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDefSubPart{}).Where("`b_list_val` IN (?)", bListVals).Find(&results).Error

	return
}

// GetFromSubPartIDx 通过sub_part_idx获取内容
func (obj *_AllVirtualDefSubPartMgr) GetFromSubPartIDx(subPartIDx int64) (results []*AllVirtualDefSubPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDefSubPart{}).Where("`sub_part_idx` = ?", subPartIDx).Find(&results).Error

	return
}

// GetBatchFromSubPartIDx 批量查找
func (obj *_AllVirtualDefSubPartMgr) GetBatchFromSubPartIDx(subPartIDxs []int64) (results []*AllVirtualDefSubPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDefSubPart{}).Where("`sub_part_idx` IN (?)", subPartIDxs).Find(&results).Error

	return
}

// GetFromSourcePartitionID 通过source_partition_id获取内容
func (obj *_AllVirtualDefSubPartMgr) GetFromSourcePartitionID(sourcePartitionID string) (results []*AllVirtualDefSubPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDefSubPart{}).Where("`source_partition_id` = ?", sourcePartitionID).Find(&results).Error

	return
}

// GetBatchFromSourcePartitionID 批量查找
func (obj *_AllVirtualDefSubPartMgr) GetBatchFromSourcePartitionID(sourcePartitionIDs []string) (results []*AllVirtualDefSubPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDefSubPart{}).Where("`source_partition_id` IN (?)", sourcePartitionIDs).Find(&results).Error

	return
}

// GetFromMappingPgSubPartID 通过mapping_pg_sub_part_id获取内容
func (obj *_AllVirtualDefSubPartMgr) GetFromMappingPgSubPartID(mappingPgSubPartID int64) (results []*AllVirtualDefSubPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDefSubPart{}).Where("`mapping_pg_sub_part_id` = ?", mappingPgSubPartID).Find(&results).Error

	return
}

// GetBatchFromMappingPgSubPartID 批量查找
func (obj *_AllVirtualDefSubPartMgr) GetBatchFromMappingPgSubPartID(mappingPgSubPartIDs []int64) (results []*AllVirtualDefSubPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDefSubPart{}).Where("`mapping_pg_sub_part_id` IN (?)", mappingPgSubPartIDs).Find(&results).Error

	return
}

// GetFromTablespaceID 通过tablespace_id获取内容
func (obj *_AllVirtualDefSubPartMgr) GetFromTablespaceID(tablespaceID int64) (results []*AllVirtualDefSubPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDefSubPart{}).Where("`tablespace_id` = ?", tablespaceID).Find(&results).Error

	return
}

// GetBatchFromTablespaceID 批量查找
func (obj *_AllVirtualDefSubPartMgr) GetBatchFromTablespaceID(tablespaceIDs []int64) (results []*AllVirtualDefSubPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDefSubPart{}).Where("`tablespace_id` IN (?)", tablespaceIDs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllVirtualDefSubPartMgr) FetchByPrimaryKey(tenantID int64, tableID int64, subPartID int64) (result AllVirtualDefSubPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDefSubPart{}).Where("`tenant_id` = ? AND `table_id` = ? AND `sub_part_id` = ?", tenantID, tableID, subPartID).First(&result).Error

	return
}
