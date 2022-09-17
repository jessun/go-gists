package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _AllVirtualSubPartHistoryMgr struct {
	*_BaseMgr
}

// AllVirtualSubPartHistoryMgr open func
func AllVirtualSubPartHistoryMgr(db *gorm.DB) *_AllVirtualSubPartHistoryMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualSubPartHistoryMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualSubPartHistoryMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_sub_part_history"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualSubPartHistoryMgr) GetTableName() string {
	return "__all_virtual_sub_part_history"
}

// Reset 重置gorm会话
func (obj *_AllVirtualSubPartHistoryMgr) Reset() *_AllVirtualSubPartHistoryMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualSubPartHistoryMgr) Get() (result AllVirtualSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSubPartHistory{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualSubPartHistoryMgr) Gets() (results []*AllVirtualSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSubPartHistory{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualSubPartHistoryMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualSubPartHistory{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTenantID tenant_id获取
func (obj *_AllVirtualSubPartHistoryMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithTableID table_id获取
func (obj *_AllVirtualSubPartHistoryMgr) WithTableID(tableID int64) Option {
	return optionFunc(func(o *options) { o.query["table_id"] = tableID })
}

// WithPartID part_id获取
func (obj *_AllVirtualSubPartHistoryMgr) WithPartID(partID int64) Option {
	return optionFunc(func(o *options) { o.query["part_id"] = partID })
}

// WithSubPartID sub_part_id获取
func (obj *_AllVirtualSubPartHistoryMgr) WithSubPartID(subPartID int64) Option {
	return optionFunc(func(o *options) { o.query["sub_part_id"] = subPartID })
}

// WithSchemaVersion schema_version获取
func (obj *_AllVirtualSubPartHistoryMgr) WithSchemaVersion(schemaVersion int64) Option {
	return optionFunc(func(o *options) { o.query["schema_version"] = schemaVersion })
}

// WithGmtCreate gmt_create获取
func (obj *_AllVirtualSubPartHistoryMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllVirtualSubPartHistoryMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithIsDeleted is_deleted获取
func (obj *_AllVirtualSubPartHistoryMgr) WithIsDeleted(isDeleted int64) Option {
	return optionFunc(func(o *options) { o.query["is_deleted"] = isDeleted })
}

// WithSubPartName sub_part_name获取
func (obj *_AllVirtualSubPartHistoryMgr) WithSubPartName(subPartName string) Option {
	return optionFunc(func(o *options) { o.query["sub_part_name"] = subPartName })
}

// WithHighBoundVal high_bound_val获取
func (obj *_AllVirtualSubPartHistoryMgr) WithHighBoundVal(highBoundVal string) Option {
	return optionFunc(func(o *options) { o.query["high_bound_val"] = highBoundVal })
}

// WithBHighBoundVal b_high_bound_val获取
func (obj *_AllVirtualSubPartHistoryMgr) WithBHighBoundVal(bHighBoundVal string) Option {
	return optionFunc(func(o *options) { o.query["b_high_bound_val"] = bHighBoundVal })
}

// WithBlockSize block_size获取
func (obj *_AllVirtualSubPartHistoryMgr) WithBlockSize(blockSize int64) Option {
	return optionFunc(func(o *options) { o.query["block_size"] = blockSize })
}

// WithReplicaNum replica_num获取
func (obj *_AllVirtualSubPartHistoryMgr) WithReplicaNum(replicaNum int64) Option {
	return optionFunc(func(o *options) { o.query["replica_num"] = replicaNum })
}

// WithCompressFuncName compress_func_name获取
func (obj *_AllVirtualSubPartHistoryMgr) WithCompressFuncName(compressFuncName string) Option {
	return optionFunc(func(o *options) { o.query["compress_func_name"] = compressFuncName })
}

// WithStatus status获取
func (obj *_AllVirtualSubPartHistoryMgr) WithStatus(status int64) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithSpare1 spare1获取
func (obj *_AllVirtualSubPartHistoryMgr) WithSpare1(spare1 int64) Option {
	return optionFunc(func(o *options) { o.query["spare1"] = spare1 })
}

// WithSpare2 spare2获取
func (obj *_AllVirtualSubPartHistoryMgr) WithSpare2(spare2 int64) Option {
	return optionFunc(func(o *options) { o.query["spare2"] = spare2 })
}

// WithSpare3 spare3获取
func (obj *_AllVirtualSubPartHistoryMgr) WithSpare3(spare3 string) Option {
	return optionFunc(func(o *options) { o.query["spare3"] = spare3 })
}

// WithComment comment获取
func (obj *_AllVirtualSubPartHistoryMgr) WithComment(comment string) Option {
	return optionFunc(func(o *options) { o.query["comment"] = comment })
}

// WithListVal list_val获取
func (obj *_AllVirtualSubPartHistoryMgr) WithListVal(listVal string) Option {
	return optionFunc(func(o *options) { o.query["list_val"] = listVal })
}

// WithBListVal b_list_val获取
func (obj *_AllVirtualSubPartHistoryMgr) WithBListVal(bListVal string) Option {
	return optionFunc(func(o *options) { o.query["b_list_val"] = bListVal })
}

// WithTablespaceID tablespace_id获取
func (obj *_AllVirtualSubPartHistoryMgr) WithTablespaceID(tablespaceID int64) Option {
	return optionFunc(func(o *options) { o.query["tablespace_id"] = tablespaceID })
}

// WithSubPartIDx sub_part_idx获取
func (obj *_AllVirtualSubPartHistoryMgr) WithSubPartIDx(subPartIDx int64) Option {
	return optionFunc(func(o *options) { o.query["sub_part_idx"] = subPartIDx })
}

// WithSourcePartitionID source_partition_id获取
func (obj *_AllVirtualSubPartHistoryMgr) WithSourcePartitionID(sourcePartitionID string) Option {
	return optionFunc(func(o *options) { o.query["source_partition_id"] = sourcePartitionID })
}

// WithMappingPgSubPartID mapping_pg_sub_part_id获取
func (obj *_AllVirtualSubPartHistoryMgr) WithMappingPgSubPartID(mappingPgSubPartID int64) Option {
	return optionFunc(func(o *options) { o.query["mapping_pg_sub_part_id"] = mappingPgSubPartID })
}

// WithDropSchemaVersion drop_schema_version获取
func (obj *_AllVirtualSubPartHistoryMgr) WithDropSchemaVersion(dropSchemaVersion int64) Option {
	return optionFunc(func(o *options) { o.query["drop_schema_version"] = dropSchemaVersion })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualSubPartHistoryMgr) GetByOption(opts ...Option) (result AllVirtualSubPartHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSubPartHistory{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualSubPartHistoryMgr) GetByOptions(opts ...Option) (results []*AllVirtualSubPartHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSubPartHistory{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualSubPartHistoryMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualSubPartHistory, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualSubPartHistory{}).Where(options.query)
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
func (obj *_AllVirtualSubPartHistoryMgr) GetFromTenantID(tenantID int64) (results []*AllVirtualSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSubPartHistory{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllVirtualSubPartHistoryMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllVirtualSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSubPartHistory{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromTableID 通过table_id获取内容
func (obj *_AllVirtualSubPartHistoryMgr) GetFromTableID(tableID int64) (results []*AllVirtualSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSubPartHistory{}).Where("`table_id` = ?", tableID).Find(&results).Error

	return
}

// GetBatchFromTableID 批量查找
func (obj *_AllVirtualSubPartHistoryMgr) GetBatchFromTableID(tableIDs []int64) (results []*AllVirtualSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSubPartHistory{}).Where("`table_id` IN (?)", tableIDs).Find(&results).Error

	return
}

// GetFromPartID 通过part_id获取内容
func (obj *_AllVirtualSubPartHistoryMgr) GetFromPartID(partID int64) (results []*AllVirtualSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSubPartHistory{}).Where("`part_id` = ?", partID).Find(&results).Error

	return
}

// GetBatchFromPartID 批量查找
func (obj *_AllVirtualSubPartHistoryMgr) GetBatchFromPartID(partIDs []int64) (results []*AllVirtualSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSubPartHistory{}).Where("`part_id` IN (?)", partIDs).Find(&results).Error

	return
}

// GetFromSubPartID 通过sub_part_id获取内容
func (obj *_AllVirtualSubPartHistoryMgr) GetFromSubPartID(subPartID int64) (results []*AllVirtualSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSubPartHistory{}).Where("`sub_part_id` = ?", subPartID).Find(&results).Error

	return
}

// GetBatchFromSubPartID 批量查找
func (obj *_AllVirtualSubPartHistoryMgr) GetBatchFromSubPartID(subPartIDs []int64) (results []*AllVirtualSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSubPartHistory{}).Where("`sub_part_id` IN (?)", subPartIDs).Find(&results).Error

	return
}

// GetFromSchemaVersion 通过schema_version获取内容
func (obj *_AllVirtualSubPartHistoryMgr) GetFromSchemaVersion(schemaVersion int64) (results []*AllVirtualSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSubPartHistory{}).Where("`schema_version` = ?", schemaVersion).Find(&results).Error

	return
}

// GetBatchFromSchemaVersion 批量查找
func (obj *_AllVirtualSubPartHistoryMgr) GetBatchFromSchemaVersion(schemaVersions []int64) (results []*AllVirtualSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSubPartHistory{}).Where("`schema_version` IN (?)", schemaVersions).Find(&results).Error

	return
}

// GetFromGmtCreate 通过gmt_create获取内容
func (obj *_AllVirtualSubPartHistoryMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllVirtualSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSubPartHistory{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllVirtualSubPartHistoryMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllVirtualSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSubPartHistory{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllVirtualSubPartHistoryMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllVirtualSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSubPartHistory{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllVirtualSubPartHistoryMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllVirtualSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSubPartHistory{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromIsDeleted 通过is_deleted获取内容
func (obj *_AllVirtualSubPartHistoryMgr) GetFromIsDeleted(isDeleted int64) (results []*AllVirtualSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSubPartHistory{}).Where("`is_deleted` = ?", isDeleted).Find(&results).Error

	return
}

// GetBatchFromIsDeleted 批量查找
func (obj *_AllVirtualSubPartHistoryMgr) GetBatchFromIsDeleted(isDeleteds []int64) (results []*AllVirtualSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSubPartHistory{}).Where("`is_deleted` IN (?)", isDeleteds).Find(&results).Error

	return
}

// GetFromSubPartName 通过sub_part_name获取内容
func (obj *_AllVirtualSubPartHistoryMgr) GetFromSubPartName(subPartName string) (results []*AllVirtualSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSubPartHistory{}).Where("`sub_part_name` = ?", subPartName).Find(&results).Error

	return
}

// GetBatchFromSubPartName 批量查找
func (obj *_AllVirtualSubPartHistoryMgr) GetBatchFromSubPartName(subPartNames []string) (results []*AllVirtualSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSubPartHistory{}).Where("`sub_part_name` IN (?)", subPartNames).Find(&results).Error

	return
}

// GetFromHighBoundVal 通过high_bound_val获取内容
func (obj *_AllVirtualSubPartHistoryMgr) GetFromHighBoundVal(highBoundVal string) (results []*AllVirtualSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSubPartHistory{}).Where("`high_bound_val` = ?", highBoundVal).Find(&results).Error

	return
}

// GetBatchFromHighBoundVal 批量查找
func (obj *_AllVirtualSubPartHistoryMgr) GetBatchFromHighBoundVal(highBoundVals []string) (results []*AllVirtualSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSubPartHistory{}).Where("`high_bound_val` IN (?)", highBoundVals).Find(&results).Error

	return
}

// GetFromBHighBoundVal 通过b_high_bound_val获取内容
func (obj *_AllVirtualSubPartHistoryMgr) GetFromBHighBoundVal(bHighBoundVal string) (results []*AllVirtualSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSubPartHistory{}).Where("`b_high_bound_val` = ?", bHighBoundVal).Find(&results).Error

	return
}

// GetBatchFromBHighBoundVal 批量查找
func (obj *_AllVirtualSubPartHistoryMgr) GetBatchFromBHighBoundVal(bHighBoundVals []string) (results []*AllVirtualSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSubPartHistory{}).Where("`b_high_bound_val` IN (?)", bHighBoundVals).Find(&results).Error

	return
}

// GetFromBlockSize 通过block_size获取内容
func (obj *_AllVirtualSubPartHistoryMgr) GetFromBlockSize(blockSize int64) (results []*AllVirtualSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSubPartHistory{}).Where("`block_size` = ?", blockSize).Find(&results).Error

	return
}

// GetBatchFromBlockSize 批量查找
func (obj *_AllVirtualSubPartHistoryMgr) GetBatchFromBlockSize(blockSizes []int64) (results []*AllVirtualSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSubPartHistory{}).Where("`block_size` IN (?)", blockSizes).Find(&results).Error

	return
}

// GetFromReplicaNum 通过replica_num获取内容
func (obj *_AllVirtualSubPartHistoryMgr) GetFromReplicaNum(replicaNum int64) (results []*AllVirtualSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSubPartHistory{}).Where("`replica_num` = ?", replicaNum).Find(&results).Error

	return
}

// GetBatchFromReplicaNum 批量查找
func (obj *_AllVirtualSubPartHistoryMgr) GetBatchFromReplicaNum(replicaNums []int64) (results []*AllVirtualSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSubPartHistory{}).Where("`replica_num` IN (?)", replicaNums).Find(&results).Error

	return
}

// GetFromCompressFuncName 通过compress_func_name获取内容
func (obj *_AllVirtualSubPartHistoryMgr) GetFromCompressFuncName(compressFuncName string) (results []*AllVirtualSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSubPartHistory{}).Where("`compress_func_name` = ?", compressFuncName).Find(&results).Error

	return
}

// GetBatchFromCompressFuncName 批量查找
func (obj *_AllVirtualSubPartHistoryMgr) GetBatchFromCompressFuncName(compressFuncNames []string) (results []*AllVirtualSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSubPartHistory{}).Where("`compress_func_name` IN (?)", compressFuncNames).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容
func (obj *_AllVirtualSubPartHistoryMgr) GetFromStatus(status int64) (results []*AllVirtualSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSubPartHistory{}).Where("`status` = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量查找
func (obj *_AllVirtualSubPartHistoryMgr) GetBatchFromStatus(statuss []int64) (results []*AllVirtualSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSubPartHistory{}).Where("`status` IN (?)", statuss).Find(&results).Error

	return
}

// GetFromSpare1 通过spare1获取内容
func (obj *_AllVirtualSubPartHistoryMgr) GetFromSpare1(spare1 int64) (results []*AllVirtualSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSubPartHistory{}).Where("`spare1` = ?", spare1).Find(&results).Error

	return
}

// GetBatchFromSpare1 批量查找
func (obj *_AllVirtualSubPartHistoryMgr) GetBatchFromSpare1(spare1s []int64) (results []*AllVirtualSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSubPartHistory{}).Where("`spare1` IN (?)", spare1s).Find(&results).Error

	return
}

// GetFromSpare2 通过spare2获取内容
func (obj *_AllVirtualSubPartHistoryMgr) GetFromSpare2(spare2 int64) (results []*AllVirtualSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSubPartHistory{}).Where("`spare2` = ?", spare2).Find(&results).Error

	return
}

// GetBatchFromSpare2 批量查找
func (obj *_AllVirtualSubPartHistoryMgr) GetBatchFromSpare2(spare2s []int64) (results []*AllVirtualSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSubPartHistory{}).Where("`spare2` IN (?)", spare2s).Find(&results).Error

	return
}

// GetFromSpare3 通过spare3获取内容
func (obj *_AllVirtualSubPartHistoryMgr) GetFromSpare3(spare3 string) (results []*AllVirtualSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSubPartHistory{}).Where("`spare3` = ?", spare3).Find(&results).Error

	return
}

// GetBatchFromSpare3 批量查找
func (obj *_AllVirtualSubPartHistoryMgr) GetBatchFromSpare3(spare3s []string) (results []*AllVirtualSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSubPartHistory{}).Where("`spare3` IN (?)", spare3s).Find(&results).Error

	return
}

// GetFromComment 通过comment获取内容
func (obj *_AllVirtualSubPartHistoryMgr) GetFromComment(comment string) (results []*AllVirtualSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSubPartHistory{}).Where("`comment` = ?", comment).Find(&results).Error

	return
}

// GetBatchFromComment 批量查找
func (obj *_AllVirtualSubPartHistoryMgr) GetBatchFromComment(comments []string) (results []*AllVirtualSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSubPartHistory{}).Where("`comment` IN (?)", comments).Find(&results).Error

	return
}

// GetFromListVal 通过list_val获取内容
func (obj *_AllVirtualSubPartHistoryMgr) GetFromListVal(listVal string) (results []*AllVirtualSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSubPartHistory{}).Where("`list_val` = ?", listVal).Find(&results).Error

	return
}

// GetBatchFromListVal 批量查找
func (obj *_AllVirtualSubPartHistoryMgr) GetBatchFromListVal(listVals []string) (results []*AllVirtualSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSubPartHistory{}).Where("`list_val` IN (?)", listVals).Find(&results).Error

	return
}

// GetFromBListVal 通过b_list_val获取内容
func (obj *_AllVirtualSubPartHistoryMgr) GetFromBListVal(bListVal string) (results []*AllVirtualSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSubPartHistory{}).Where("`b_list_val` = ?", bListVal).Find(&results).Error

	return
}

// GetBatchFromBListVal 批量查找
func (obj *_AllVirtualSubPartHistoryMgr) GetBatchFromBListVal(bListVals []string) (results []*AllVirtualSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSubPartHistory{}).Where("`b_list_val` IN (?)", bListVals).Find(&results).Error

	return
}

// GetFromTablespaceID 通过tablespace_id获取内容
func (obj *_AllVirtualSubPartHistoryMgr) GetFromTablespaceID(tablespaceID int64) (results []*AllVirtualSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSubPartHistory{}).Where("`tablespace_id` = ?", tablespaceID).Find(&results).Error

	return
}

// GetBatchFromTablespaceID 批量查找
func (obj *_AllVirtualSubPartHistoryMgr) GetBatchFromTablespaceID(tablespaceIDs []int64) (results []*AllVirtualSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSubPartHistory{}).Where("`tablespace_id` IN (?)", tablespaceIDs).Find(&results).Error

	return
}

// GetFromSubPartIDx 通过sub_part_idx获取内容
func (obj *_AllVirtualSubPartHistoryMgr) GetFromSubPartIDx(subPartIDx int64) (results []*AllVirtualSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSubPartHistory{}).Where("`sub_part_idx` = ?", subPartIDx).Find(&results).Error

	return
}

// GetBatchFromSubPartIDx 批量查找
func (obj *_AllVirtualSubPartHistoryMgr) GetBatchFromSubPartIDx(subPartIDxs []int64) (results []*AllVirtualSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSubPartHistory{}).Where("`sub_part_idx` IN (?)", subPartIDxs).Find(&results).Error

	return
}

// GetFromSourcePartitionID 通过source_partition_id获取内容
func (obj *_AllVirtualSubPartHistoryMgr) GetFromSourcePartitionID(sourcePartitionID string) (results []*AllVirtualSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSubPartHistory{}).Where("`source_partition_id` = ?", sourcePartitionID).Find(&results).Error

	return
}

// GetBatchFromSourcePartitionID 批量查找
func (obj *_AllVirtualSubPartHistoryMgr) GetBatchFromSourcePartitionID(sourcePartitionIDs []string) (results []*AllVirtualSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSubPartHistory{}).Where("`source_partition_id` IN (?)", sourcePartitionIDs).Find(&results).Error

	return
}

// GetFromMappingPgSubPartID 通过mapping_pg_sub_part_id获取内容
func (obj *_AllVirtualSubPartHistoryMgr) GetFromMappingPgSubPartID(mappingPgSubPartID int64) (results []*AllVirtualSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSubPartHistory{}).Where("`mapping_pg_sub_part_id` = ?", mappingPgSubPartID).Find(&results).Error

	return
}

// GetBatchFromMappingPgSubPartID 批量查找
func (obj *_AllVirtualSubPartHistoryMgr) GetBatchFromMappingPgSubPartID(mappingPgSubPartIDs []int64) (results []*AllVirtualSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSubPartHistory{}).Where("`mapping_pg_sub_part_id` IN (?)", mappingPgSubPartIDs).Find(&results).Error

	return
}

// GetFromDropSchemaVersion 通过drop_schema_version获取内容
func (obj *_AllVirtualSubPartHistoryMgr) GetFromDropSchemaVersion(dropSchemaVersion int64) (results []*AllVirtualSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSubPartHistory{}).Where("`drop_schema_version` = ?", dropSchemaVersion).Find(&results).Error

	return
}

// GetBatchFromDropSchemaVersion 批量查找
func (obj *_AllVirtualSubPartHistoryMgr) GetBatchFromDropSchemaVersion(dropSchemaVersions []int64) (results []*AllVirtualSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSubPartHistory{}).Where("`drop_schema_version` IN (?)", dropSchemaVersions).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllVirtualSubPartHistoryMgr) FetchByPrimaryKey(tenantID int64, tableID int64, partID int64, subPartID int64, schemaVersion int64) (result AllVirtualSubPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSubPartHistory{}).Where("`tenant_id` = ? AND `table_id` = ? AND `part_id` = ? AND `sub_part_id` = ? AND `schema_version` = ?", tenantID, tableID, partID, subPartID, schemaVersion).First(&result).Error

	return
}
