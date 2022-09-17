package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _AllPartHistoryMgr struct {
	*_BaseMgr
}

// AllPartHistoryMgr open func
func AllPartHistoryMgr(db *gorm.DB) *_AllPartHistoryMgr {
	if db == nil {
		panic(fmt.Errorf("AllPartHistoryMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllPartHistoryMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_part_history"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllPartHistoryMgr) GetTableName() string {
	return "__all_part_history"
}

// Reset 重置gorm会话
func (obj *_AllPartHistoryMgr) Reset() *_AllPartHistoryMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllPartHistoryMgr) Get() (result AllPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPartHistory{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllPartHistoryMgr) Gets() (results []*AllPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPartHistory{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllPartHistoryMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllPartHistory{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithGmtCreate gmt_create获取
func (obj *_AllPartHistoryMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllPartHistoryMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithTenantID tenant_id获取
func (obj *_AllPartHistoryMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithTableID table_id获取
func (obj *_AllPartHistoryMgr) WithTableID(tableID int64) Option {
	return optionFunc(func(o *options) { o.query["table_id"] = tableID })
}

// WithPartID part_id获取
func (obj *_AllPartHistoryMgr) WithPartID(partID int64) Option {
	return optionFunc(func(o *options) { o.query["part_id"] = partID })
}

// WithSchemaVersion schema_version获取
func (obj *_AllPartHistoryMgr) WithSchemaVersion(schemaVersion int64) Option {
	return optionFunc(func(o *options) { o.query["schema_version"] = schemaVersion })
}

// WithIsDeleted is_deleted获取
func (obj *_AllPartHistoryMgr) WithIsDeleted(isDeleted int64) Option {
	return optionFunc(func(o *options) { o.query["is_deleted"] = isDeleted })
}

// WithPartName part_name获取
func (obj *_AllPartHistoryMgr) WithPartName(partName string) Option {
	return optionFunc(func(o *options) { o.query["part_name"] = partName })
}

// WithHighBoundVal high_bound_val获取
func (obj *_AllPartHistoryMgr) WithHighBoundVal(highBoundVal string) Option {
	return optionFunc(func(o *options) { o.query["high_bound_val"] = highBoundVal })
}

// WithBHighBoundVal b_high_bound_val获取
func (obj *_AllPartHistoryMgr) WithBHighBoundVal(bHighBoundVal string) Option {
	return optionFunc(func(o *options) { o.query["b_high_bound_val"] = bHighBoundVal })
}

// WithSubPartNum sub_part_num获取
func (obj *_AllPartHistoryMgr) WithSubPartNum(subPartNum int64) Option {
	return optionFunc(func(o *options) { o.query["sub_part_num"] = subPartNum })
}

// WithSubPartSpace sub_part_space获取
func (obj *_AllPartHistoryMgr) WithSubPartSpace(subPartSpace int64) Option {
	return optionFunc(func(o *options) { o.query["sub_part_space"] = subPartSpace })
}

// WithNewSubPartNum new_sub_part_num获取
func (obj *_AllPartHistoryMgr) WithNewSubPartNum(newSubPartNum int64) Option {
	return optionFunc(func(o *options) { o.query["new_sub_part_num"] = newSubPartNum })
}

// WithNewSubPartSpace new_sub_part_space获取
func (obj *_AllPartHistoryMgr) WithNewSubPartSpace(newSubPartSpace int64) Option {
	return optionFunc(func(o *options) { o.query["new_sub_part_space"] = newSubPartSpace })
}

// WithSubPartInterval sub_part_interval获取
func (obj *_AllPartHistoryMgr) WithSubPartInterval(subPartInterval string) Option {
	return optionFunc(func(o *options) { o.query["sub_part_interval"] = subPartInterval })
}

// WithSubIntervalStart sub_interval_start获取
func (obj *_AllPartHistoryMgr) WithSubIntervalStart(subIntervalStart string) Option {
	return optionFunc(func(o *options) { o.query["sub_interval_start"] = subIntervalStart })
}

// WithNewSubPartInterval new_sub_part_interval获取
func (obj *_AllPartHistoryMgr) WithNewSubPartInterval(newSubPartInterval string) Option {
	return optionFunc(func(o *options) { o.query["new_sub_part_interval"] = newSubPartInterval })
}

// WithNewSubIntervalStart new_sub_interval_start获取
func (obj *_AllPartHistoryMgr) WithNewSubIntervalStart(newSubIntervalStart string) Option {
	return optionFunc(func(o *options) { o.query["new_sub_interval_start"] = newSubIntervalStart })
}

// WithBlockSize block_size获取
func (obj *_AllPartHistoryMgr) WithBlockSize(blockSize int64) Option {
	return optionFunc(func(o *options) { o.query["block_size"] = blockSize })
}

// WithReplicaNum replica_num获取
func (obj *_AllPartHistoryMgr) WithReplicaNum(replicaNum int64) Option {
	return optionFunc(func(o *options) { o.query["replica_num"] = replicaNum })
}

// WithCompressFuncName compress_func_name获取
func (obj *_AllPartHistoryMgr) WithCompressFuncName(compressFuncName string) Option {
	return optionFunc(func(o *options) { o.query["compress_func_name"] = compressFuncName })
}

// WithStatus status获取
func (obj *_AllPartHistoryMgr) WithStatus(status int64) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithSpare1 spare1获取
func (obj *_AllPartHistoryMgr) WithSpare1(spare1 int64) Option {
	return optionFunc(func(o *options) { o.query["spare1"] = spare1 })
}

// WithSpare2 spare2获取
func (obj *_AllPartHistoryMgr) WithSpare2(spare2 int64) Option {
	return optionFunc(func(o *options) { o.query["spare2"] = spare2 })
}

// WithSpare3 spare3获取
func (obj *_AllPartHistoryMgr) WithSpare3(spare3 string) Option {
	return optionFunc(func(o *options) { o.query["spare3"] = spare3 })
}

// WithComment comment获取
func (obj *_AllPartHistoryMgr) WithComment(comment string) Option {
	return optionFunc(func(o *options) { o.query["comment"] = comment })
}

// WithListVal list_val获取
func (obj *_AllPartHistoryMgr) WithListVal(listVal string) Option {
	return optionFunc(func(o *options) { o.query["list_val"] = listVal })
}

// WithBListVal b_list_val获取
func (obj *_AllPartHistoryMgr) WithBListVal(bListVal string) Option {
	return optionFunc(func(o *options) { o.query["b_list_val"] = bListVal })
}

// WithPartIDx part_idx获取
func (obj *_AllPartHistoryMgr) WithPartIDx(partIDx int64) Option {
	return optionFunc(func(o *options) { o.query["part_idx"] = partIDx })
}

// WithSourcePartitionID source_partition_id获取
func (obj *_AllPartHistoryMgr) WithSourcePartitionID(sourcePartitionID string) Option {
	return optionFunc(func(o *options) { o.query["source_partition_id"] = sourcePartitionID })
}

// WithMappingPgPartID mapping_pg_part_id获取
func (obj *_AllPartHistoryMgr) WithMappingPgPartID(mappingPgPartID int64) Option {
	return optionFunc(func(o *options) { o.query["mapping_pg_part_id"] = mappingPgPartID })
}

// WithTablespaceID tablespace_id获取
func (obj *_AllPartHistoryMgr) WithTablespaceID(tablespaceID int64) Option {
	return optionFunc(func(o *options) { o.query["tablespace_id"] = tablespaceID })
}

// WithDropSchemaVersion drop_schema_version获取
func (obj *_AllPartHistoryMgr) WithDropSchemaVersion(dropSchemaVersion int64) Option {
	return optionFunc(func(o *options) { o.query["drop_schema_version"] = dropSchemaVersion })
}

// WithMaxUsedSubPartID max_used_sub_part_id获取
func (obj *_AllPartHistoryMgr) WithMaxUsedSubPartID(maxUsedSubPartID int64) Option {
	return optionFunc(func(o *options) { o.query["max_used_sub_part_id"] = maxUsedSubPartID })
}

// GetByOption 功能选项模式获取
func (obj *_AllPartHistoryMgr) GetByOption(opts ...Option) (result AllPartHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllPartHistory{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllPartHistoryMgr) GetByOptions(opts ...Option) (results []*AllPartHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllPartHistory{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllPartHistoryMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllPartHistory, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllPartHistory{}).Where(options.query)
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
func (obj *_AllPartHistoryMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPartHistory{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllPartHistoryMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPartHistory{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllPartHistoryMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPartHistory{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllPartHistoryMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPartHistory{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromTenantID 通过tenant_id获取内容
func (obj *_AllPartHistoryMgr) GetFromTenantID(tenantID int64) (results []*AllPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPartHistory{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllPartHistoryMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPartHistory{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromTableID 通过table_id获取内容
func (obj *_AllPartHistoryMgr) GetFromTableID(tableID int64) (results []*AllPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPartHistory{}).Where("`table_id` = ?", tableID).Find(&results).Error

	return
}

// GetBatchFromTableID 批量查找
func (obj *_AllPartHistoryMgr) GetBatchFromTableID(tableIDs []int64) (results []*AllPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPartHistory{}).Where("`table_id` IN (?)", tableIDs).Find(&results).Error

	return
}

// GetFromPartID 通过part_id获取内容
func (obj *_AllPartHistoryMgr) GetFromPartID(partID int64) (results []*AllPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPartHistory{}).Where("`part_id` = ?", partID).Find(&results).Error

	return
}

// GetBatchFromPartID 批量查找
func (obj *_AllPartHistoryMgr) GetBatchFromPartID(partIDs []int64) (results []*AllPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPartHistory{}).Where("`part_id` IN (?)", partIDs).Find(&results).Error

	return
}

// GetFromSchemaVersion 通过schema_version获取内容
func (obj *_AllPartHistoryMgr) GetFromSchemaVersion(schemaVersion int64) (results []*AllPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPartHistory{}).Where("`schema_version` = ?", schemaVersion).Find(&results).Error

	return
}

// GetBatchFromSchemaVersion 批量查找
func (obj *_AllPartHistoryMgr) GetBatchFromSchemaVersion(schemaVersions []int64) (results []*AllPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPartHistory{}).Where("`schema_version` IN (?)", schemaVersions).Find(&results).Error

	return
}

// GetFromIsDeleted 通过is_deleted获取内容
func (obj *_AllPartHistoryMgr) GetFromIsDeleted(isDeleted int64) (results []*AllPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPartHistory{}).Where("`is_deleted` = ?", isDeleted).Find(&results).Error

	return
}

// GetBatchFromIsDeleted 批量查找
func (obj *_AllPartHistoryMgr) GetBatchFromIsDeleted(isDeleteds []int64) (results []*AllPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPartHistory{}).Where("`is_deleted` IN (?)", isDeleteds).Find(&results).Error

	return
}

// GetFromPartName 通过part_name获取内容
func (obj *_AllPartHistoryMgr) GetFromPartName(partName string) (results []*AllPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPartHistory{}).Where("`part_name` = ?", partName).Find(&results).Error

	return
}

// GetBatchFromPartName 批量查找
func (obj *_AllPartHistoryMgr) GetBatchFromPartName(partNames []string) (results []*AllPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPartHistory{}).Where("`part_name` IN (?)", partNames).Find(&results).Error

	return
}

// GetFromHighBoundVal 通过high_bound_val获取内容
func (obj *_AllPartHistoryMgr) GetFromHighBoundVal(highBoundVal string) (results []*AllPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPartHistory{}).Where("`high_bound_val` = ?", highBoundVal).Find(&results).Error

	return
}

// GetBatchFromHighBoundVal 批量查找
func (obj *_AllPartHistoryMgr) GetBatchFromHighBoundVal(highBoundVals []string) (results []*AllPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPartHistory{}).Where("`high_bound_val` IN (?)", highBoundVals).Find(&results).Error

	return
}

// GetFromBHighBoundVal 通过b_high_bound_val获取内容
func (obj *_AllPartHistoryMgr) GetFromBHighBoundVal(bHighBoundVal string) (results []*AllPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPartHistory{}).Where("`b_high_bound_val` = ?", bHighBoundVal).Find(&results).Error

	return
}

// GetBatchFromBHighBoundVal 批量查找
func (obj *_AllPartHistoryMgr) GetBatchFromBHighBoundVal(bHighBoundVals []string) (results []*AllPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPartHistory{}).Where("`b_high_bound_val` IN (?)", bHighBoundVals).Find(&results).Error

	return
}

// GetFromSubPartNum 通过sub_part_num获取内容
func (obj *_AllPartHistoryMgr) GetFromSubPartNum(subPartNum int64) (results []*AllPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPartHistory{}).Where("`sub_part_num` = ?", subPartNum).Find(&results).Error

	return
}

// GetBatchFromSubPartNum 批量查找
func (obj *_AllPartHistoryMgr) GetBatchFromSubPartNum(subPartNums []int64) (results []*AllPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPartHistory{}).Where("`sub_part_num` IN (?)", subPartNums).Find(&results).Error

	return
}

// GetFromSubPartSpace 通过sub_part_space获取内容
func (obj *_AllPartHistoryMgr) GetFromSubPartSpace(subPartSpace int64) (results []*AllPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPartHistory{}).Where("`sub_part_space` = ?", subPartSpace).Find(&results).Error

	return
}

// GetBatchFromSubPartSpace 批量查找
func (obj *_AllPartHistoryMgr) GetBatchFromSubPartSpace(subPartSpaces []int64) (results []*AllPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPartHistory{}).Where("`sub_part_space` IN (?)", subPartSpaces).Find(&results).Error

	return
}

// GetFromNewSubPartNum 通过new_sub_part_num获取内容
func (obj *_AllPartHistoryMgr) GetFromNewSubPartNum(newSubPartNum int64) (results []*AllPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPartHistory{}).Where("`new_sub_part_num` = ?", newSubPartNum).Find(&results).Error

	return
}

// GetBatchFromNewSubPartNum 批量查找
func (obj *_AllPartHistoryMgr) GetBatchFromNewSubPartNum(newSubPartNums []int64) (results []*AllPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPartHistory{}).Where("`new_sub_part_num` IN (?)", newSubPartNums).Find(&results).Error

	return
}

// GetFromNewSubPartSpace 通过new_sub_part_space获取内容
func (obj *_AllPartHistoryMgr) GetFromNewSubPartSpace(newSubPartSpace int64) (results []*AllPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPartHistory{}).Where("`new_sub_part_space` = ?", newSubPartSpace).Find(&results).Error

	return
}

// GetBatchFromNewSubPartSpace 批量查找
func (obj *_AllPartHistoryMgr) GetBatchFromNewSubPartSpace(newSubPartSpaces []int64) (results []*AllPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPartHistory{}).Where("`new_sub_part_space` IN (?)", newSubPartSpaces).Find(&results).Error

	return
}

// GetFromSubPartInterval 通过sub_part_interval获取内容
func (obj *_AllPartHistoryMgr) GetFromSubPartInterval(subPartInterval string) (results []*AllPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPartHistory{}).Where("`sub_part_interval` = ?", subPartInterval).Find(&results).Error

	return
}

// GetBatchFromSubPartInterval 批量查找
func (obj *_AllPartHistoryMgr) GetBatchFromSubPartInterval(subPartIntervals []string) (results []*AllPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPartHistory{}).Where("`sub_part_interval` IN (?)", subPartIntervals).Find(&results).Error

	return
}

// GetFromSubIntervalStart 通过sub_interval_start获取内容
func (obj *_AllPartHistoryMgr) GetFromSubIntervalStart(subIntervalStart string) (results []*AllPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPartHistory{}).Where("`sub_interval_start` = ?", subIntervalStart).Find(&results).Error

	return
}

// GetBatchFromSubIntervalStart 批量查找
func (obj *_AllPartHistoryMgr) GetBatchFromSubIntervalStart(subIntervalStarts []string) (results []*AllPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPartHistory{}).Where("`sub_interval_start` IN (?)", subIntervalStarts).Find(&results).Error

	return
}

// GetFromNewSubPartInterval 通过new_sub_part_interval获取内容
func (obj *_AllPartHistoryMgr) GetFromNewSubPartInterval(newSubPartInterval string) (results []*AllPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPartHistory{}).Where("`new_sub_part_interval` = ?", newSubPartInterval).Find(&results).Error

	return
}

// GetBatchFromNewSubPartInterval 批量查找
func (obj *_AllPartHistoryMgr) GetBatchFromNewSubPartInterval(newSubPartIntervals []string) (results []*AllPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPartHistory{}).Where("`new_sub_part_interval` IN (?)", newSubPartIntervals).Find(&results).Error

	return
}

// GetFromNewSubIntervalStart 通过new_sub_interval_start获取内容
func (obj *_AllPartHistoryMgr) GetFromNewSubIntervalStart(newSubIntervalStart string) (results []*AllPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPartHistory{}).Where("`new_sub_interval_start` = ?", newSubIntervalStart).Find(&results).Error

	return
}

// GetBatchFromNewSubIntervalStart 批量查找
func (obj *_AllPartHistoryMgr) GetBatchFromNewSubIntervalStart(newSubIntervalStarts []string) (results []*AllPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPartHistory{}).Where("`new_sub_interval_start` IN (?)", newSubIntervalStarts).Find(&results).Error

	return
}

// GetFromBlockSize 通过block_size获取内容
func (obj *_AllPartHistoryMgr) GetFromBlockSize(blockSize int64) (results []*AllPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPartHistory{}).Where("`block_size` = ?", blockSize).Find(&results).Error

	return
}

// GetBatchFromBlockSize 批量查找
func (obj *_AllPartHistoryMgr) GetBatchFromBlockSize(blockSizes []int64) (results []*AllPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPartHistory{}).Where("`block_size` IN (?)", blockSizes).Find(&results).Error

	return
}

// GetFromReplicaNum 通过replica_num获取内容
func (obj *_AllPartHistoryMgr) GetFromReplicaNum(replicaNum int64) (results []*AllPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPartHistory{}).Where("`replica_num` = ?", replicaNum).Find(&results).Error

	return
}

// GetBatchFromReplicaNum 批量查找
func (obj *_AllPartHistoryMgr) GetBatchFromReplicaNum(replicaNums []int64) (results []*AllPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPartHistory{}).Where("`replica_num` IN (?)", replicaNums).Find(&results).Error

	return
}

// GetFromCompressFuncName 通过compress_func_name获取内容
func (obj *_AllPartHistoryMgr) GetFromCompressFuncName(compressFuncName string) (results []*AllPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPartHistory{}).Where("`compress_func_name` = ?", compressFuncName).Find(&results).Error

	return
}

// GetBatchFromCompressFuncName 批量查找
func (obj *_AllPartHistoryMgr) GetBatchFromCompressFuncName(compressFuncNames []string) (results []*AllPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPartHistory{}).Where("`compress_func_name` IN (?)", compressFuncNames).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容
func (obj *_AllPartHistoryMgr) GetFromStatus(status int64) (results []*AllPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPartHistory{}).Where("`status` = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量查找
func (obj *_AllPartHistoryMgr) GetBatchFromStatus(statuss []int64) (results []*AllPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPartHistory{}).Where("`status` IN (?)", statuss).Find(&results).Error

	return
}

// GetFromSpare1 通过spare1获取内容
func (obj *_AllPartHistoryMgr) GetFromSpare1(spare1 int64) (results []*AllPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPartHistory{}).Where("`spare1` = ?", spare1).Find(&results).Error

	return
}

// GetBatchFromSpare1 批量查找
func (obj *_AllPartHistoryMgr) GetBatchFromSpare1(spare1s []int64) (results []*AllPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPartHistory{}).Where("`spare1` IN (?)", spare1s).Find(&results).Error

	return
}

// GetFromSpare2 通过spare2获取内容
func (obj *_AllPartHistoryMgr) GetFromSpare2(spare2 int64) (results []*AllPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPartHistory{}).Where("`spare2` = ?", spare2).Find(&results).Error

	return
}

// GetBatchFromSpare2 批量查找
func (obj *_AllPartHistoryMgr) GetBatchFromSpare2(spare2s []int64) (results []*AllPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPartHistory{}).Where("`spare2` IN (?)", spare2s).Find(&results).Error

	return
}

// GetFromSpare3 通过spare3获取内容
func (obj *_AllPartHistoryMgr) GetFromSpare3(spare3 string) (results []*AllPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPartHistory{}).Where("`spare3` = ?", spare3).Find(&results).Error

	return
}

// GetBatchFromSpare3 批量查找
func (obj *_AllPartHistoryMgr) GetBatchFromSpare3(spare3s []string) (results []*AllPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPartHistory{}).Where("`spare3` IN (?)", spare3s).Find(&results).Error

	return
}

// GetFromComment 通过comment获取内容
func (obj *_AllPartHistoryMgr) GetFromComment(comment string) (results []*AllPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPartHistory{}).Where("`comment` = ?", comment).Find(&results).Error

	return
}

// GetBatchFromComment 批量查找
func (obj *_AllPartHistoryMgr) GetBatchFromComment(comments []string) (results []*AllPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPartHistory{}).Where("`comment` IN (?)", comments).Find(&results).Error

	return
}

// GetFromListVal 通过list_val获取内容
func (obj *_AllPartHistoryMgr) GetFromListVal(listVal string) (results []*AllPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPartHistory{}).Where("`list_val` = ?", listVal).Find(&results).Error

	return
}

// GetBatchFromListVal 批量查找
func (obj *_AllPartHistoryMgr) GetBatchFromListVal(listVals []string) (results []*AllPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPartHistory{}).Where("`list_val` IN (?)", listVals).Find(&results).Error

	return
}

// GetFromBListVal 通过b_list_val获取内容
func (obj *_AllPartHistoryMgr) GetFromBListVal(bListVal string) (results []*AllPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPartHistory{}).Where("`b_list_val` = ?", bListVal).Find(&results).Error

	return
}

// GetBatchFromBListVal 批量查找
func (obj *_AllPartHistoryMgr) GetBatchFromBListVal(bListVals []string) (results []*AllPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPartHistory{}).Where("`b_list_val` IN (?)", bListVals).Find(&results).Error

	return
}

// GetFromPartIDx 通过part_idx获取内容
func (obj *_AllPartHistoryMgr) GetFromPartIDx(partIDx int64) (results []*AllPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPartHistory{}).Where("`part_idx` = ?", partIDx).Find(&results).Error

	return
}

// GetBatchFromPartIDx 批量查找
func (obj *_AllPartHistoryMgr) GetBatchFromPartIDx(partIDxs []int64) (results []*AllPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPartHistory{}).Where("`part_idx` IN (?)", partIDxs).Find(&results).Error

	return
}

// GetFromSourcePartitionID 通过source_partition_id获取内容
func (obj *_AllPartHistoryMgr) GetFromSourcePartitionID(sourcePartitionID string) (results []*AllPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPartHistory{}).Where("`source_partition_id` = ?", sourcePartitionID).Find(&results).Error

	return
}

// GetBatchFromSourcePartitionID 批量查找
func (obj *_AllPartHistoryMgr) GetBatchFromSourcePartitionID(sourcePartitionIDs []string) (results []*AllPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPartHistory{}).Where("`source_partition_id` IN (?)", sourcePartitionIDs).Find(&results).Error

	return
}

// GetFromMappingPgPartID 通过mapping_pg_part_id获取内容
func (obj *_AllPartHistoryMgr) GetFromMappingPgPartID(mappingPgPartID int64) (results []*AllPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPartHistory{}).Where("`mapping_pg_part_id` = ?", mappingPgPartID).Find(&results).Error

	return
}

// GetBatchFromMappingPgPartID 批量查找
func (obj *_AllPartHistoryMgr) GetBatchFromMappingPgPartID(mappingPgPartIDs []int64) (results []*AllPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPartHistory{}).Where("`mapping_pg_part_id` IN (?)", mappingPgPartIDs).Find(&results).Error

	return
}

// GetFromTablespaceID 通过tablespace_id获取内容
func (obj *_AllPartHistoryMgr) GetFromTablespaceID(tablespaceID int64) (results []*AllPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPartHistory{}).Where("`tablespace_id` = ?", tablespaceID).Find(&results).Error

	return
}

// GetBatchFromTablespaceID 批量查找
func (obj *_AllPartHistoryMgr) GetBatchFromTablespaceID(tablespaceIDs []int64) (results []*AllPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPartHistory{}).Where("`tablespace_id` IN (?)", tablespaceIDs).Find(&results).Error

	return
}

// GetFromDropSchemaVersion 通过drop_schema_version获取内容
func (obj *_AllPartHistoryMgr) GetFromDropSchemaVersion(dropSchemaVersion int64) (results []*AllPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPartHistory{}).Where("`drop_schema_version` = ?", dropSchemaVersion).Find(&results).Error

	return
}

// GetBatchFromDropSchemaVersion 批量查找
func (obj *_AllPartHistoryMgr) GetBatchFromDropSchemaVersion(dropSchemaVersions []int64) (results []*AllPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPartHistory{}).Where("`drop_schema_version` IN (?)", dropSchemaVersions).Find(&results).Error

	return
}

// GetFromMaxUsedSubPartID 通过max_used_sub_part_id获取内容
func (obj *_AllPartHistoryMgr) GetFromMaxUsedSubPartID(maxUsedSubPartID int64) (results []*AllPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPartHistory{}).Where("`max_used_sub_part_id` = ?", maxUsedSubPartID).Find(&results).Error

	return
}

// GetBatchFromMaxUsedSubPartID 批量查找
func (obj *_AllPartHistoryMgr) GetBatchFromMaxUsedSubPartID(maxUsedSubPartIDs []int64) (results []*AllPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPartHistory{}).Where("`max_used_sub_part_id` IN (?)", maxUsedSubPartIDs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllPartHistoryMgr) FetchByPrimaryKey(tenantID int64, tableID int64, partID int64, schemaVersion int64) (result AllPartHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPartHistory{}).Where("`tenant_id` = ? AND `table_id` = ? AND `part_id` = ? AND `schema_version` = ?", tenantID, tableID, partID, schemaVersion).First(&result).Error

	return
}
