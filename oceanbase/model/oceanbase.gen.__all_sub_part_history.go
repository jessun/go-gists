package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _AllSubPartHistoryMgr struct {
	*_BaseMgr
}

// AllSubPartHistoryMgr open func
func AllSubPartHistoryMgr(db *gorm.DB) *_AllSubPartHistoryMgr {
	if db == nil {
		panic(fmt.Errorf("AllSubPartHistoryMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllSubPartHistoryMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_sub_part_history"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllSubPartHistoryMgr) GetTableName() string {
	return "__all_sub_part_history"
}

// Reset 重置gorm会话
func (obj *_AllSubPartHistoryMgr) Reset() *_AllSubPartHistoryMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllSubPartHistoryMgr) Get() (result AllSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSubPartHistory{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllSubPartHistoryMgr) Gets() (results []*AllSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSubPartHistory{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllSubPartHistoryMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllSubPartHistory{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithGmtCreate gmt_create获取
func (obj *_AllSubPartHistoryMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllSubPartHistoryMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithTenantID tenant_id获取
func (obj *_AllSubPartHistoryMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithTableID table_id获取
func (obj *_AllSubPartHistoryMgr) WithTableID(tableID int64) Option {
	return optionFunc(func(o *options) { o.query["table_id"] = tableID })
}

// WithPartID part_id获取
func (obj *_AllSubPartHistoryMgr) WithPartID(partID int64) Option {
	return optionFunc(func(o *options) { o.query["part_id"] = partID })
}

// WithSubPartID sub_part_id获取
func (obj *_AllSubPartHistoryMgr) WithSubPartID(subPartID int64) Option {
	return optionFunc(func(o *options) { o.query["sub_part_id"] = subPartID })
}

// WithSchemaVersion schema_version获取
func (obj *_AllSubPartHistoryMgr) WithSchemaVersion(schemaVersion int64) Option {
	return optionFunc(func(o *options) { o.query["schema_version"] = schemaVersion })
}

// WithIsDeleted is_deleted获取
func (obj *_AllSubPartHistoryMgr) WithIsDeleted(isDeleted int64) Option {
	return optionFunc(func(o *options) { o.query["is_deleted"] = isDeleted })
}

// WithSubPartName sub_part_name获取
func (obj *_AllSubPartHistoryMgr) WithSubPartName(subPartName string) Option {
	return optionFunc(func(o *options) { o.query["sub_part_name"] = subPartName })
}

// WithHighBoundVal high_bound_val获取
func (obj *_AllSubPartHistoryMgr) WithHighBoundVal(highBoundVal string) Option {
	return optionFunc(func(o *options) { o.query["high_bound_val"] = highBoundVal })
}

// WithBHighBoundVal b_high_bound_val获取
func (obj *_AllSubPartHistoryMgr) WithBHighBoundVal(bHighBoundVal string) Option {
	return optionFunc(func(o *options) { o.query["b_high_bound_val"] = bHighBoundVal })
}

// WithBlockSize block_size获取
func (obj *_AllSubPartHistoryMgr) WithBlockSize(blockSize int64) Option {
	return optionFunc(func(o *options) { o.query["block_size"] = blockSize })
}

// WithReplicaNum replica_num获取
func (obj *_AllSubPartHistoryMgr) WithReplicaNum(replicaNum int64) Option {
	return optionFunc(func(o *options) { o.query["replica_num"] = replicaNum })
}

// WithCompressFuncName compress_func_name获取
func (obj *_AllSubPartHistoryMgr) WithCompressFuncName(compressFuncName string) Option {
	return optionFunc(func(o *options) { o.query["compress_func_name"] = compressFuncName })
}

// WithStatus status获取
func (obj *_AllSubPartHistoryMgr) WithStatus(status int64) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithSpare1 spare1获取
func (obj *_AllSubPartHistoryMgr) WithSpare1(spare1 int64) Option {
	return optionFunc(func(o *options) { o.query["spare1"] = spare1 })
}

// WithSpare2 spare2获取
func (obj *_AllSubPartHistoryMgr) WithSpare2(spare2 int64) Option {
	return optionFunc(func(o *options) { o.query["spare2"] = spare2 })
}

// WithSpare3 spare3获取
func (obj *_AllSubPartHistoryMgr) WithSpare3(spare3 string) Option {
	return optionFunc(func(o *options) { o.query["spare3"] = spare3 })
}

// WithComment comment获取
func (obj *_AllSubPartHistoryMgr) WithComment(comment string) Option {
	return optionFunc(func(o *options) { o.query["comment"] = comment })
}

// WithListVal list_val获取
func (obj *_AllSubPartHistoryMgr) WithListVal(listVal string) Option {
	return optionFunc(func(o *options) { o.query["list_val"] = listVal })
}

// WithBListVal b_list_val获取
func (obj *_AllSubPartHistoryMgr) WithBListVal(bListVal string) Option {
	return optionFunc(func(o *options) { o.query["b_list_val"] = bListVal })
}

// WithTablespaceID tablespace_id获取
func (obj *_AllSubPartHistoryMgr) WithTablespaceID(tablespaceID int64) Option {
	return optionFunc(func(o *options) { o.query["tablespace_id"] = tablespaceID })
}

// WithSubPartIDx sub_part_idx获取
func (obj *_AllSubPartHistoryMgr) WithSubPartIDx(subPartIDx int64) Option {
	return optionFunc(func(o *options) { o.query["sub_part_idx"] = subPartIDx })
}

// WithSourcePartitionID source_partition_id获取
func (obj *_AllSubPartHistoryMgr) WithSourcePartitionID(sourcePartitionID string) Option {
	return optionFunc(func(o *options) { o.query["source_partition_id"] = sourcePartitionID })
}

// WithMappingPgSubPartID mapping_pg_sub_part_id获取
func (obj *_AllSubPartHistoryMgr) WithMappingPgSubPartID(mappingPgSubPartID int64) Option {
	return optionFunc(func(o *options) { o.query["mapping_pg_sub_part_id"] = mappingPgSubPartID })
}

// WithDropSchemaVersion drop_schema_version获取
func (obj *_AllSubPartHistoryMgr) WithDropSchemaVersion(dropSchemaVersion int64) Option {
	return optionFunc(func(o *options) { o.query["drop_schema_version"] = dropSchemaVersion })
}

// GetByOption 功能选项模式获取
func (obj *_AllSubPartHistoryMgr) GetByOption(opts ...Option) (result AllSubPartHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllSubPartHistory{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllSubPartHistoryMgr) GetByOptions(opts ...Option) (results []*AllSubPartHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllSubPartHistory{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllSubPartHistoryMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllSubPartHistory, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllSubPartHistory{}).Where(options.query)
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
func (obj *_AllSubPartHistoryMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSubPartHistory{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllSubPartHistoryMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSubPartHistory{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllSubPartHistoryMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSubPartHistory{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllSubPartHistoryMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSubPartHistory{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromTenantID 通过tenant_id获取内容
func (obj *_AllSubPartHistoryMgr) GetFromTenantID(tenantID int64) (results []*AllSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSubPartHistory{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllSubPartHistoryMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSubPartHistory{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromTableID 通过table_id获取内容
func (obj *_AllSubPartHistoryMgr) GetFromTableID(tableID int64) (results []*AllSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSubPartHistory{}).Where("`table_id` = ?", tableID).Find(&results).Error

	return
}

// GetBatchFromTableID 批量查找
func (obj *_AllSubPartHistoryMgr) GetBatchFromTableID(tableIDs []int64) (results []*AllSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSubPartHistory{}).Where("`table_id` IN (?)", tableIDs).Find(&results).Error

	return
}

// GetFromPartID 通过part_id获取内容
func (obj *_AllSubPartHistoryMgr) GetFromPartID(partID int64) (results []*AllSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSubPartHistory{}).Where("`part_id` = ?", partID).Find(&results).Error

	return
}

// GetBatchFromPartID 批量查找
func (obj *_AllSubPartHistoryMgr) GetBatchFromPartID(partIDs []int64) (results []*AllSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSubPartHistory{}).Where("`part_id` IN (?)", partIDs).Find(&results).Error

	return
}

// GetFromSubPartID 通过sub_part_id获取内容
func (obj *_AllSubPartHistoryMgr) GetFromSubPartID(subPartID int64) (results []*AllSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSubPartHistory{}).Where("`sub_part_id` = ?", subPartID).Find(&results).Error

	return
}

// GetBatchFromSubPartID 批量查找
func (obj *_AllSubPartHistoryMgr) GetBatchFromSubPartID(subPartIDs []int64) (results []*AllSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSubPartHistory{}).Where("`sub_part_id` IN (?)", subPartIDs).Find(&results).Error

	return
}

// GetFromSchemaVersion 通过schema_version获取内容
func (obj *_AllSubPartHistoryMgr) GetFromSchemaVersion(schemaVersion int64) (results []*AllSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSubPartHistory{}).Where("`schema_version` = ?", schemaVersion).Find(&results).Error

	return
}

// GetBatchFromSchemaVersion 批量查找
func (obj *_AllSubPartHistoryMgr) GetBatchFromSchemaVersion(schemaVersions []int64) (results []*AllSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSubPartHistory{}).Where("`schema_version` IN (?)", schemaVersions).Find(&results).Error

	return
}

// GetFromIsDeleted 通过is_deleted获取内容
func (obj *_AllSubPartHistoryMgr) GetFromIsDeleted(isDeleted int64) (results []*AllSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSubPartHistory{}).Where("`is_deleted` = ?", isDeleted).Find(&results).Error

	return
}

// GetBatchFromIsDeleted 批量查找
func (obj *_AllSubPartHistoryMgr) GetBatchFromIsDeleted(isDeleteds []int64) (results []*AllSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSubPartHistory{}).Where("`is_deleted` IN (?)", isDeleteds).Find(&results).Error

	return
}

// GetFromSubPartName 通过sub_part_name获取内容
func (obj *_AllSubPartHistoryMgr) GetFromSubPartName(subPartName string) (results []*AllSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSubPartHistory{}).Where("`sub_part_name` = ?", subPartName).Find(&results).Error

	return
}

// GetBatchFromSubPartName 批量查找
func (obj *_AllSubPartHistoryMgr) GetBatchFromSubPartName(subPartNames []string) (results []*AllSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSubPartHistory{}).Where("`sub_part_name` IN (?)", subPartNames).Find(&results).Error

	return
}

// GetFromHighBoundVal 通过high_bound_val获取内容
func (obj *_AllSubPartHistoryMgr) GetFromHighBoundVal(highBoundVal string) (results []*AllSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSubPartHistory{}).Where("`high_bound_val` = ?", highBoundVal).Find(&results).Error

	return
}

// GetBatchFromHighBoundVal 批量查找
func (obj *_AllSubPartHistoryMgr) GetBatchFromHighBoundVal(highBoundVals []string) (results []*AllSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSubPartHistory{}).Where("`high_bound_val` IN (?)", highBoundVals).Find(&results).Error

	return
}

// GetFromBHighBoundVal 通过b_high_bound_val获取内容
func (obj *_AllSubPartHistoryMgr) GetFromBHighBoundVal(bHighBoundVal string) (results []*AllSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSubPartHistory{}).Where("`b_high_bound_val` = ?", bHighBoundVal).Find(&results).Error

	return
}

// GetBatchFromBHighBoundVal 批量查找
func (obj *_AllSubPartHistoryMgr) GetBatchFromBHighBoundVal(bHighBoundVals []string) (results []*AllSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSubPartHistory{}).Where("`b_high_bound_val` IN (?)", bHighBoundVals).Find(&results).Error

	return
}

// GetFromBlockSize 通过block_size获取内容
func (obj *_AllSubPartHistoryMgr) GetFromBlockSize(blockSize int64) (results []*AllSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSubPartHistory{}).Where("`block_size` = ?", blockSize).Find(&results).Error

	return
}

// GetBatchFromBlockSize 批量查找
func (obj *_AllSubPartHistoryMgr) GetBatchFromBlockSize(blockSizes []int64) (results []*AllSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSubPartHistory{}).Where("`block_size` IN (?)", blockSizes).Find(&results).Error

	return
}

// GetFromReplicaNum 通过replica_num获取内容
func (obj *_AllSubPartHistoryMgr) GetFromReplicaNum(replicaNum int64) (results []*AllSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSubPartHistory{}).Where("`replica_num` = ?", replicaNum).Find(&results).Error

	return
}

// GetBatchFromReplicaNum 批量查找
func (obj *_AllSubPartHistoryMgr) GetBatchFromReplicaNum(replicaNums []int64) (results []*AllSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSubPartHistory{}).Where("`replica_num` IN (?)", replicaNums).Find(&results).Error

	return
}

// GetFromCompressFuncName 通过compress_func_name获取内容
func (obj *_AllSubPartHistoryMgr) GetFromCompressFuncName(compressFuncName string) (results []*AllSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSubPartHistory{}).Where("`compress_func_name` = ?", compressFuncName).Find(&results).Error

	return
}

// GetBatchFromCompressFuncName 批量查找
func (obj *_AllSubPartHistoryMgr) GetBatchFromCompressFuncName(compressFuncNames []string) (results []*AllSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSubPartHistory{}).Where("`compress_func_name` IN (?)", compressFuncNames).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容
func (obj *_AllSubPartHistoryMgr) GetFromStatus(status int64) (results []*AllSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSubPartHistory{}).Where("`status` = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量查找
func (obj *_AllSubPartHistoryMgr) GetBatchFromStatus(statuss []int64) (results []*AllSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSubPartHistory{}).Where("`status` IN (?)", statuss).Find(&results).Error

	return
}

// GetFromSpare1 通过spare1获取内容
func (obj *_AllSubPartHistoryMgr) GetFromSpare1(spare1 int64) (results []*AllSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSubPartHistory{}).Where("`spare1` = ?", spare1).Find(&results).Error

	return
}

// GetBatchFromSpare1 批量查找
func (obj *_AllSubPartHistoryMgr) GetBatchFromSpare1(spare1s []int64) (results []*AllSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSubPartHistory{}).Where("`spare1` IN (?)", spare1s).Find(&results).Error

	return
}

// GetFromSpare2 通过spare2获取内容
func (obj *_AllSubPartHistoryMgr) GetFromSpare2(spare2 int64) (results []*AllSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSubPartHistory{}).Where("`spare2` = ?", spare2).Find(&results).Error

	return
}

// GetBatchFromSpare2 批量查找
func (obj *_AllSubPartHistoryMgr) GetBatchFromSpare2(spare2s []int64) (results []*AllSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSubPartHistory{}).Where("`spare2` IN (?)", spare2s).Find(&results).Error

	return
}

// GetFromSpare3 通过spare3获取内容
func (obj *_AllSubPartHistoryMgr) GetFromSpare3(spare3 string) (results []*AllSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSubPartHistory{}).Where("`spare3` = ?", spare3).Find(&results).Error

	return
}

// GetBatchFromSpare3 批量查找
func (obj *_AllSubPartHistoryMgr) GetBatchFromSpare3(spare3s []string) (results []*AllSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSubPartHistory{}).Where("`spare3` IN (?)", spare3s).Find(&results).Error

	return
}

// GetFromComment 通过comment获取内容
func (obj *_AllSubPartHistoryMgr) GetFromComment(comment string) (results []*AllSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSubPartHistory{}).Where("`comment` = ?", comment).Find(&results).Error

	return
}

// GetBatchFromComment 批量查找
func (obj *_AllSubPartHistoryMgr) GetBatchFromComment(comments []string) (results []*AllSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSubPartHistory{}).Where("`comment` IN (?)", comments).Find(&results).Error

	return
}

// GetFromListVal 通过list_val获取内容
func (obj *_AllSubPartHistoryMgr) GetFromListVal(listVal string) (results []*AllSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSubPartHistory{}).Where("`list_val` = ?", listVal).Find(&results).Error

	return
}

// GetBatchFromListVal 批量查找
func (obj *_AllSubPartHistoryMgr) GetBatchFromListVal(listVals []string) (results []*AllSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSubPartHistory{}).Where("`list_val` IN (?)", listVals).Find(&results).Error

	return
}

// GetFromBListVal 通过b_list_val获取内容
func (obj *_AllSubPartHistoryMgr) GetFromBListVal(bListVal string) (results []*AllSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSubPartHistory{}).Where("`b_list_val` = ?", bListVal).Find(&results).Error

	return
}

// GetBatchFromBListVal 批量查找
func (obj *_AllSubPartHistoryMgr) GetBatchFromBListVal(bListVals []string) (results []*AllSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSubPartHistory{}).Where("`b_list_val` IN (?)", bListVals).Find(&results).Error

	return
}

// GetFromTablespaceID 通过tablespace_id获取内容
func (obj *_AllSubPartHistoryMgr) GetFromTablespaceID(tablespaceID int64) (results []*AllSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSubPartHistory{}).Where("`tablespace_id` = ?", tablespaceID).Find(&results).Error

	return
}

// GetBatchFromTablespaceID 批量查找
func (obj *_AllSubPartHistoryMgr) GetBatchFromTablespaceID(tablespaceIDs []int64) (results []*AllSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSubPartHistory{}).Where("`tablespace_id` IN (?)", tablespaceIDs).Find(&results).Error

	return
}

// GetFromSubPartIDx 通过sub_part_idx获取内容
func (obj *_AllSubPartHistoryMgr) GetFromSubPartIDx(subPartIDx int64) (results []*AllSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSubPartHistory{}).Where("`sub_part_idx` = ?", subPartIDx).Find(&results).Error

	return
}

// GetBatchFromSubPartIDx 批量查找
func (obj *_AllSubPartHistoryMgr) GetBatchFromSubPartIDx(subPartIDxs []int64) (results []*AllSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSubPartHistory{}).Where("`sub_part_idx` IN (?)", subPartIDxs).Find(&results).Error

	return
}

// GetFromSourcePartitionID 通过source_partition_id获取内容
func (obj *_AllSubPartHistoryMgr) GetFromSourcePartitionID(sourcePartitionID string) (results []*AllSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSubPartHistory{}).Where("`source_partition_id` = ?", sourcePartitionID).Find(&results).Error

	return
}

// GetBatchFromSourcePartitionID 批量查找
func (obj *_AllSubPartHistoryMgr) GetBatchFromSourcePartitionID(sourcePartitionIDs []string) (results []*AllSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSubPartHistory{}).Where("`source_partition_id` IN (?)", sourcePartitionIDs).Find(&results).Error

	return
}

// GetFromMappingPgSubPartID 通过mapping_pg_sub_part_id获取内容
func (obj *_AllSubPartHistoryMgr) GetFromMappingPgSubPartID(mappingPgSubPartID int64) (results []*AllSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSubPartHistory{}).Where("`mapping_pg_sub_part_id` = ?", mappingPgSubPartID).Find(&results).Error

	return
}

// GetBatchFromMappingPgSubPartID 批量查找
func (obj *_AllSubPartHistoryMgr) GetBatchFromMappingPgSubPartID(mappingPgSubPartIDs []int64) (results []*AllSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSubPartHistory{}).Where("`mapping_pg_sub_part_id` IN (?)", mappingPgSubPartIDs).Find(&results).Error

	return
}

// GetFromDropSchemaVersion 通过drop_schema_version获取内容
func (obj *_AllSubPartHistoryMgr) GetFromDropSchemaVersion(dropSchemaVersion int64) (results []*AllSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSubPartHistory{}).Where("`drop_schema_version` = ?", dropSchemaVersion).Find(&results).Error

	return
}

// GetBatchFromDropSchemaVersion 批量查找
func (obj *_AllSubPartHistoryMgr) GetBatchFromDropSchemaVersion(dropSchemaVersions []int64) (results []*AllSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSubPartHistory{}).Where("`drop_schema_version` IN (?)", dropSchemaVersions).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllSubPartHistoryMgr) FetchByPrimaryKey(tenantID int64, tableID int64, partID int64, subPartID int64, schemaVersion int64) (result AllSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSubPartHistory{}).Where("`tenant_id` = ? AND `table_id` = ? AND `part_id` = ? AND `sub_part_id` = ? AND `schema_version` = ?", tenantID, tableID, partID, subPartID, schemaVersion).First(&result).Error

	return
}
