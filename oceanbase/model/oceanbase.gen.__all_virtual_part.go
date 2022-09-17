package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _AllVirtualPartMgr struct {
	*_BaseMgr
}

// AllVirtualPartMgr open func
func AllVirtualPartMgr(db *gorm.DB) *_AllVirtualPartMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualPartMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualPartMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_part"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualPartMgr) GetTableName() string {
	return "__all_virtual_part"
}

// Reset 重置gorm会话
func (obj *_AllVirtualPartMgr) Reset() *_AllVirtualPartMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualPartMgr) Get() (result AllVirtualPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPart{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualPartMgr) Gets() (results []*AllVirtualPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPart{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualPartMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualPart{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTenantID tenant_id获取
func (obj *_AllVirtualPartMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithTableID table_id获取
func (obj *_AllVirtualPartMgr) WithTableID(tableID int64) Option {
	return optionFunc(func(o *options) { o.query["table_id"] = tableID })
}

// WithPartID part_id获取
func (obj *_AllVirtualPartMgr) WithPartID(partID int64) Option {
	return optionFunc(func(o *options) { o.query["part_id"] = partID })
}

// WithGmtCreate gmt_create获取
func (obj *_AllVirtualPartMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllVirtualPartMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithPartName part_name获取
func (obj *_AllVirtualPartMgr) WithPartName(partName string) Option {
	return optionFunc(func(o *options) { o.query["part_name"] = partName })
}

// WithSchemaVersion schema_version获取
func (obj *_AllVirtualPartMgr) WithSchemaVersion(schemaVersion int64) Option {
	return optionFunc(func(o *options) { o.query["schema_version"] = schemaVersion })
}

// WithHighBoundVal high_bound_val获取
func (obj *_AllVirtualPartMgr) WithHighBoundVal(highBoundVal string) Option {
	return optionFunc(func(o *options) { o.query["high_bound_val"] = highBoundVal })
}

// WithBHighBoundVal b_high_bound_val获取
func (obj *_AllVirtualPartMgr) WithBHighBoundVal(bHighBoundVal string) Option {
	return optionFunc(func(o *options) { o.query["b_high_bound_val"] = bHighBoundVal })
}

// WithSubPartNum sub_part_num获取
func (obj *_AllVirtualPartMgr) WithSubPartNum(subPartNum int64) Option {
	return optionFunc(func(o *options) { o.query["sub_part_num"] = subPartNum })
}

// WithSubPartSpace sub_part_space获取
func (obj *_AllVirtualPartMgr) WithSubPartSpace(subPartSpace int64) Option {
	return optionFunc(func(o *options) { o.query["sub_part_space"] = subPartSpace })
}

// WithNewSubPartNum new_sub_part_num获取
func (obj *_AllVirtualPartMgr) WithNewSubPartNum(newSubPartNum int64) Option {
	return optionFunc(func(o *options) { o.query["new_sub_part_num"] = newSubPartNum })
}

// WithNewSubPartSpace new_sub_part_space获取
func (obj *_AllVirtualPartMgr) WithNewSubPartSpace(newSubPartSpace int64) Option {
	return optionFunc(func(o *options) { o.query["new_sub_part_space"] = newSubPartSpace })
}

// WithSubPartInterval sub_part_interval获取
func (obj *_AllVirtualPartMgr) WithSubPartInterval(subPartInterval string) Option {
	return optionFunc(func(o *options) { o.query["sub_part_interval"] = subPartInterval })
}

// WithSubIntervalStart sub_interval_start获取
func (obj *_AllVirtualPartMgr) WithSubIntervalStart(subIntervalStart string) Option {
	return optionFunc(func(o *options) { o.query["sub_interval_start"] = subIntervalStart })
}

// WithNewSubPartInterval new_sub_part_interval获取
func (obj *_AllVirtualPartMgr) WithNewSubPartInterval(newSubPartInterval string) Option {
	return optionFunc(func(o *options) { o.query["new_sub_part_interval"] = newSubPartInterval })
}

// WithNewSubIntervalStart new_sub_interval_start获取
func (obj *_AllVirtualPartMgr) WithNewSubIntervalStart(newSubIntervalStart string) Option {
	return optionFunc(func(o *options) { o.query["new_sub_interval_start"] = newSubIntervalStart })
}

// WithBlockSize block_size获取
func (obj *_AllVirtualPartMgr) WithBlockSize(blockSize int64) Option {
	return optionFunc(func(o *options) { o.query["block_size"] = blockSize })
}

// WithReplicaNum replica_num获取
func (obj *_AllVirtualPartMgr) WithReplicaNum(replicaNum int64) Option {
	return optionFunc(func(o *options) { o.query["replica_num"] = replicaNum })
}

// WithCompressFuncName compress_func_name获取
func (obj *_AllVirtualPartMgr) WithCompressFuncName(compressFuncName string) Option {
	return optionFunc(func(o *options) { o.query["compress_func_name"] = compressFuncName })
}

// WithStatus status获取
func (obj *_AllVirtualPartMgr) WithStatus(status int64) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithSpare1 spare1获取
func (obj *_AllVirtualPartMgr) WithSpare1(spare1 int64) Option {
	return optionFunc(func(o *options) { o.query["spare1"] = spare1 })
}

// WithSpare2 spare2获取
func (obj *_AllVirtualPartMgr) WithSpare2(spare2 int64) Option {
	return optionFunc(func(o *options) { o.query["spare2"] = spare2 })
}

// WithSpare3 spare3获取
func (obj *_AllVirtualPartMgr) WithSpare3(spare3 string) Option {
	return optionFunc(func(o *options) { o.query["spare3"] = spare3 })
}

// WithComment comment获取
func (obj *_AllVirtualPartMgr) WithComment(comment string) Option {
	return optionFunc(func(o *options) { o.query["comment"] = comment })
}

// WithListVal list_val获取
func (obj *_AllVirtualPartMgr) WithListVal(listVal string) Option {
	return optionFunc(func(o *options) { o.query["list_val"] = listVal })
}

// WithBListVal b_list_val获取
func (obj *_AllVirtualPartMgr) WithBListVal(bListVal string) Option {
	return optionFunc(func(o *options) { o.query["b_list_val"] = bListVal })
}

// WithPartIDx part_idx获取
func (obj *_AllVirtualPartMgr) WithPartIDx(partIDx int64) Option {
	return optionFunc(func(o *options) { o.query["part_idx"] = partIDx })
}

// WithSourcePartitionID source_partition_id获取
func (obj *_AllVirtualPartMgr) WithSourcePartitionID(sourcePartitionID string) Option {
	return optionFunc(func(o *options) { o.query["source_partition_id"] = sourcePartitionID })
}

// WithMappingPgPartID mapping_pg_part_id获取
func (obj *_AllVirtualPartMgr) WithMappingPgPartID(mappingPgPartID int64) Option {
	return optionFunc(func(o *options) { o.query["mapping_pg_part_id"] = mappingPgPartID })
}

// WithTablespaceID tablespace_id获取
func (obj *_AllVirtualPartMgr) WithTablespaceID(tablespaceID int64) Option {
	return optionFunc(func(o *options) { o.query["tablespace_id"] = tablespaceID })
}

// WithDropSchemaVersion drop_schema_version获取
func (obj *_AllVirtualPartMgr) WithDropSchemaVersion(dropSchemaVersion int64) Option {
	return optionFunc(func(o *options) { o.query["drop_schema_version"] = dropSchemaVersion })
}

// WithMaxUsedSubPartID max_used_sub_part_id获取
func (obj *_AllVirtualPartMgr) WithMaxUsedSubPartID(maxUsedSubPartID int64) Option {
	return optionFunc(func(o *options) { o.query["max_used_sub_part_id"] = maxUsedSubPartID })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualPartMgr) GetByOption(opts ...Option) (result AllVirtualPart, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPart{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualPartMgr) GetByOptions(opts ...Option) (results []*AllVirtualPart, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPart{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualPartMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualPart, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualPart{}).Where(options.query)
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
func (obj *_AllVirtualPartMgr) GetFromTenantID(tenantID int64) (results []*AllVirtualPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPart{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllVirtualPartMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllVirtualPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPart{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromTableID 通过table_id获取内容
func (obj *_AllVirtualPartMgr) GetFromTableID(tableID int64) (results []*AllVirtualPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPart{}).Where("`table_id` = ?", tableID).Find(&results).Error

	return
}

// GetBatchFromTableID 批量查找
func (obj *_AllVirtualPartMgr) GetBatchFromTableID(tableIDs []int64) (results []*AllVirtualPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPart{}).Where("`table_id` IN (?)", tableIDs).Find(&results).Error

	return
}

// GetFromPartID 通过part_id获取内容
func (obj *_AllVirtualPartMgr) GetFromPartID(partID int64) (results []*AllVirtualPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPart{}).Where("`part_id` = ?", partID).Find(&results).Error

	return
}

// GetBatchFromPartID 批量查找
func (obj *_AllVirtualPartMgr) GetBatchFromPartID(partIDs []int64) (results []*AllVirtualPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPart{}).Where("`part_id` IN (?)", partIDs).Find(&results).Error

	return
}

// GetFromGmtCreate 通过gmt_create获取内容
func (obj *_AllVirtualPartMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllVirtualPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPart{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllVirtualPartMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllVirtualPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPart{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllVirtualPartMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllVirtualPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPart{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllVirtualPartMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllVirtualPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPart{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromPartName 通过part_name获取内容
func (obj *_AllVirtualPartMgr) GetFromPartName(partName string) (results []*AllVirtualPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPart{}).Where("`part_name` = ?", partName).Find(&results).Error

	return
}

// GetBatchFromPartName 批量查找
func (obj *_AllVirtualPartMgr) GetBatchFromPartName(partNames []string) (results []*AllVirtualPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPart{}).Where("`part_name` IN (?)", partNames).Find(&results).Error

	return
}

// GetFromSchemaVersion 通过schema_version获取内容
func (obj *_AllVirtualPartMgr) GetFromSchemaVersion(schemaVersion int64) (results []*AllVirtualPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPart{}).Where("`schema_version` = ?", schemaVersion).Find(&results).Error

	return
}

// GetBatchFromSchemaVersion 批量查找
func (obj *_AllVirtualPartMgr) GetBatchFromSchemaVersion(schemaVersions []int64) (results []*AllVirtualPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPart{}).Where("`schema_version` IN (?)", schemaVersions).Find(&results).Error

	return
}

// GetFromHighBoundVal 通过high_bound_val获取内容
func (obj *_AllVirtualPartMgr) GetFromHighBoundVal(highBoundVal string) (results []*AllVirtualPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPart{}).Where("`high_bound_val` = ?", highBoundVal).Find(&results).Error

	return
}

// GetBatchFromHighBoundVal 批量查找
func (obj *_AllVirtualPartMgr) GetBatchFromHighBoundVal(highBoundVals []string) (results []*AllVirtualPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPart{}).Where("`high_bound_val` IN (?)", highBoundVals).Find(&results).Error

	return
}

// GetFromBHighBoundVal 通过b_high_bound_val获取内容
func (obj *_AllVirtualPartMgr) GetFromBHighBoundVal(bHighBoundVal string) (results []*AllVirtualPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPart{}).Where("`b_high_bound_val` = ?", bHighBoundVal).Find(&results).Error

	return
}

// GetBatchFromBHighBoundVal 批量查找
func (obj *_AllVirtualPartMgr) GetBatchFromBHighBoundVal(bHighBoundVals []string) (results []*AllVirtualPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPart{}).Where("`b_high_bound_val` IN (?)", bHighBoundVals).Find(&results).Error

	return
}

// GetFromSubPartNum 通过sub_part_num获取内容
func (obj *_AllVirtualPartMgr) GetFromSubPartNum(subPartNum int64) (results []*AllVirtualPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPart{}).Where("`sub_part_num` = ?", subPartNum).Find(&results).Error

	return
}

// GetBatchFromSubPartNum 批量查找
func (obj *_AllVirtualPartMgr) GetBatchFromSubPartNum(subPartNums []int64) (results []*AllVirtualPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPart{}).Where("`sub_part_num` IN (?)", subPartNums).Find(&results).Error

	return
}

// GetFromSubPartSpace 通过sub_part_space获取内容
func (obj *_AllVirtualPartMgr) GetFromSubPartSpace(subPartSpace int64) (results []*AllVirtualPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPart{}).Where("`sub_part_space` = ?", subPartSpace).Find(&results).Error

	return
}

// GetBatchFromSubPartSpace 批量查找
func (obj *_AllVirtualPartMgr) GetBatchFromSubPartSpace(subPartSpaces []int64) (results []*AllVirtualPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPart{}).Where("`sub_part_space` IN (?)", subPartSpaces).Find(&results).Error

	return
}

// GetFromNewSubPartNum 通过new_sub_part_num获取内容
func (obj *_AllVirtualPartMgr) GetFromNewSubPartNum(newSubPartNum int64) (results []*AllVirtualPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPart{}).Where("`new_sub_part_num` = ?", newSubPartNum).Find(&results).Error

	return
}

// GetBatchFromNewSubPartNum 批量查找
func (obj *_AllVirtualPartMgr) GetBatchFromNewSubPartNum(newSubPartNums []int64) (results []*AllVirtualPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPart{}).Where("`new_sub_part_num` IN (?)", newSubPartNums).Find(&results).Error

	return
}

// GetFromNewSubPartSpace 通过new_sub_part_space获取内容
func (obj *_AllVirtualPartMgr) GetFromNewSubPartSpace(newSubPartSpace int64) (results []*AllVirtualPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPart{}).Where("`new_sub_part_space` = ?", newSubPartSpace).Find(&results).Error

	return
}

// GetBatchFromNewSubPartSpace 批量查找
func (obj *_AllVirtualPartMgr) GetBatchFromNewSubPartSpace(newSubPartSpaces []int64) (results []*AllVirtualPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPart{}).Where("`new_sub_part_space` IN (?)", newSubPartSpaces).Find(&results).Error

	return
}

// GetFromSubPartInterval 通过sub_part_interval获取内容
func (obj *_AllVirtualPartMgr) GetFromSubPartInterval(subPartInterval string) (results []*AllVirtualPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPart{}).Where("`sub_part_interval` = ?", subPartInterval).Find(&results).Error

	return
}

// GetBatchFromSubPartInterval 批量查找
func (obj *_AllVirtualPartMgr) GetBatchFromSubPartInterval(subPartIntervals []string) (results []*AllVirtualPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPart{}).Where("`sub_part_interval` IN (?)", subPartIntervals).Find(&results).Error

	return
}

// GetFromSubIntervalStart 通过sub_interval_start获取内容
func (obj *_AllVirtualPartMgr) GetFromSubIntervalStart(subIntervalStart string) (results []*AllVirtualPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPart{}).Where("`sub_interval_start` = ?", subIntervalStart).Find(&results).Error

	return
}

// GetBatchFromSubIntervalStart 批量查找
func (obj *_AllVirtualPartMgr) GetBatchFromSubIntervalStart(subIntervalStarts []string) (results []*AllVirtualPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPart{}).Where("`sub_interval_start` IN (?)", subIntervalStarts).Find(&results).Error

	return
}

// GetFromNewSubPartInterval 通过new_sub_part_interval获取内容
func (obj *_AllVirtualPartMgr) GetFromNewSubPartInterval(newSubPartInterval string) (results []*AllVirtualPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPart{}).Where("`new_sub_part_interval` = ?", newSubPartInterval).Find(&results).Error

	return
}

// GetBatchFromNewSubPartInterval 批量查找
func (obj *_AllVirtualPartMgr) GetBatchFromNewSubPartInterval(newSubPartIntervals []string) (results []*AllVirtualPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPart{}).Where("`new_sub_part_interval` IN (?)", newSubPartIntervals).Find(&results).Error

	return
}

// GetFromNewSubIntervalStart 通过new_sub_interval_start获取内容
func (obj *_AllVirtualPartMgr) GetFromNewSubIntervalStart(newSubIntervalStart string) (results []*AllVirtualPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPart{}).Where("`new_sub_interval_start` = ?", newSubIntervalStart).Find(&results).Error

	return
}

// GetBatchFromNewSubIntervalStart 批量查找
func (obj *_AllVirtualPartMgr) GetBatchFromNewSubIntervalStart(newSubIntervalStarts []string) (results []*AllVirtualPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPart{}).Where("`new_sub_interval_start` IN (?)", newSubIntervalStarts).Find(&results).Error

	return
}

// GetFromBlockSize 通过block_size获取内容
func (obj *_AllVirtualPartMgr) GetFromBlockSize(blockSize int64) (results []*AllVirtualPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPart{}).Where("`block_size` = ?", blockSize).Find(&results).Error

	return
}

// GetBatchFromBlockSize 批量查找
func (obj *_AllVirtualPartMgr) GetBatchFromBlockSize(blockSizes []int64) (results []*AllVirtualPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPart{}).Where("`block_size` IN (?)", blockSizes).Find(&results).Error

	return
}

// GetFromReplicaNum 通过replica_num获取内容
func (obj *_AllVirtualPartMgr) GetFromReplicaNum(replicaNum int64) (results []*AllVirtualPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPart{}).Where("`replica_num` = ?", replicaNum).Find(&results).Error

	return
}

// GetBatchFromReplicaNum 批量查找
func (obj *_AllVirtualPartMgr) GetBatchFromReplicaNum(replicaNums []int64) (results []*AllVirtualPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPart{}).Where("`replica_num` IN (?)", replicaNums).Find(&results).Error

	return
}

// GetFromCompressFuncName 通过compress_func_name获取内容
func (obj *_AllVirtualPartMgr) GetFromCompressFuncName(compressFuncName string) (results []*AllVirtualPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPart{}).Where("`compress_func_name` = ?", compressFuncName).Find(&results).Error

	return
}

// GetBatchFromCompressFuncName 批量查找
func (obj *_AllVirtualPartMgr) GetBatchFromCompressFuncName(compressFuncNames []string) (results []*AllVirtualPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPart{}).Where("`compress_func_name` IN (?)", compressFuncNames).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容
func (obj *_AllVirtualPartMgr) GetFromStatus(status int64) (results []*AllVirtualPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPart{}).Where("`status` = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量查找
func (obj *_AllVirtualPartMgr) GetBatchFromStatus(statuss []int64) (results []*AllVirtualPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPart{}).Where("`status` IN (?)", statuss).Find(&results).Error

	return
}

// GetFromSpare1 通过spare1获取内容
func (obj *_AllVirtualPartMgr) GetFromSpare1(spare1 int64) (results []*AllVirtualPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPart{}).Where("`spare1` = ?", spare1).Find(&results).Error

	return
}

// GetBatchFromSpare1 批量查找
func (obj *_AllVirtualPartMgr) GetBatchFromSpare1(spare1s []int64) (results []*AllVirtualPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPart{}).Where("`spare1` IN (?)", spare1s).Find(&results).Error

	return
}

// GetFromSpare2 通过spare2获取内容
func (obj *_AllVirtualPartMgr) GetFromSpare2(spare2 int64) (results []*AllVirtualPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPart{}).Where("`spare2` = ?", spare2).Find(&results).Error

	return
}

// GetBatchFromSpare2 批量查找
func (obj *_AllVirtualPartMgr) GetBatchFromSpare2(spare2s []int64) (results []*AllVirtualPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPart{}).Where("`spare2` IN (?)", spare2s).Find(&results).Error

	return
}

// GetFromSpare3 通过spare3获取内容
func (obj *_AllVirtualPartMgr) GetFromSpare3(spare3 string) (results []*AllVirtualPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPart{}).Where("`spare3` = ?", spare3).Find(&results).Error

	return
}

// GetBatchFromSpare3 批量查找
func (obj *_AllVirtualPartMgr) GetBatchFromSpare3(spare3s []string) (results []*AllVirtualPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPart{}).Where("`spare3` IN (?)", spare3s).Find(&results).Error

	return
}

// GetFromComment 通过comment获取内容
func (obj *_AllVirtualPartMgr) GetFromComment(comment string) (results []*AllVirtualPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPart{}).Where("`comment` = ?", comment).Find(&results).Error

	return
}

// GetBatchFromComment 批量查找
func (obj *_AllVirtualPartMgr) GetBatchFromComment(comments []string) (results []*AllVirtualPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPart{}).Where("`comment` IN (?)", comments).Find(&results).Error

	return
}

// GetFromListVal 通过list_val获取内容
func (obj *_AllVirtualPartMgr) GetFromListVal(listVal string) (results []*AllVirtualPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPart{}).Where("`list_val` = ?", listVal).Find(&results).Error

	return
}

// GetBatchFromListVal 批量查找
func (obj *_AllVirtualPartMgr) GetBatchFromListVal(listVals []string) (results []*AllVirtualPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPart{}).Where("`list_val` IN (?)", listVals).Find(&results).Error

	return
}

// GetFromBListVal 通过b_list_val获取内容
func (obj *_AllVirtualPartMgr) GetFromBListVal(bListVal string) (results []*AllVirtualPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPart{}).Where("`b_list_val` = ?", bListVal).Find(&results).Error

	return
}

// GetBatchFromBListVal 批量查找
func (obj *_AllVirtualPartMgr) GetBatchFromBListVal(bListVals []string) (results []*AllVirtualPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPart{}).Where("`b_list_val` IN (?)", bListVals).Find(&results).Error

	return
}

// GetFromPartIDx 通过part_idx获取内容
func (obj *_AllVirtualPartMgr) GetFromPartIDx(partIDx int64) (results []*AllVirtualPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPart{}).Where("`part_idx` = ?", partIDx).Find(&results).Error

	return
}

// GetBatchFromPartIDx 批量查找
func (obj *_AllVirtualPartMgr) GetBatchFromPartIDx(partIDxs []int64) (results []*AllVirtualPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPart{}).Where("`part_idx` IN (?)", partIDxs).Find(&results).Error

	return
}

// GetFromSourcePartitionID 通过source_partition_id获取内容
func (obj *_AllVirtualPartMgr) GetFromSourcePartitionID(sourcePartitionID string) (results []*AllVirtualPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPart{}).Where("`source_partition_id` = ?", sourcePartitionID).Find(&results).Error

	return
}

// GetBatchFromSourcePartitionID 批量查找
func (obj *_AllVirtualPartMgr) GetBatchFromSourcePartitionID(sourcePartitionIDs []string) (results []*AllVirtualPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPart{}).Where("`source_partition_id` IN (?)", sourcePartitionIDs).Find(&results).Error

	return
}

// GetFromMappingPgPartID 通过mapping_pg_part_id获取内容
func (obj *_AllVirtualPartMgr) GetFromMappingPgPartID(mappingPgPartID int64) (results []*AllVirtualPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPart{}).Where("`mapping_pg_part_id` = ?", mappingPgPartID).Find(&results).Error

	return
}

// GetBatchFromMappingPgPartID 批量查找
func (obj *_AllVirtualPartMgr) GetBatchFromMappingPgPartID(mappingPgPartIDs []int64) (results []*AllVirtualPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPart{}).Where("`mapping_pg_part_id` IN (?)", mappingPgPartIDs).Find(&results).Error

	return
}

// GetFromTablespaceID 通过tablespace_id获取内容
func (obj *_AllVirtualPartMgr) GetFromTablespaceID(tablespaceID int64) (results []*AllVirtualPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPart{}).Where("`tablespace_id` = ?", tablespaceID).Find(&results).Error

	return
}

// GetBatchFromTablespaceID 批量查找
func (obj *_AllVirtualPartMgr) GetBatchFromTablespaceID(tablespaceIDs []int64) (results []*AllVirtualPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPart{}).Where("`tablespace_id` IN (?)", tablespaceIDs).Find(&results).Error

	return
}

// GetFromDropSchemaVersion 通过drop_schema_version获取内容
func (obj *_AllVirtualPartMgr) GetFromDropSchemaVersion(dropSchemaVersion int64) (results []*AllVirtualPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPart{}).Where("`drop_schema_version` = ?", dropSchemaVersion).Find(&results).Error

	return
}

// GetBatchFromDropSchemaVersion 批量查找
func (obj *_AllVirtualPartMgr) GetBatchFromDropSchemaVersion(dropSchemaVersions []int64) (results []*AllVirtualPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPart{}).Where("`drop_schema_version` IN (?)", dropSchemaVersions).Find(&results).Error

	return
}

// GetFromMaxUsedSubPartID 通过max_used_sub_part_id获取内容
func (obj *_AllVirtualPartMgr) GetFromMaxUsedSubPartID(maxUsedSubPartID int64) (results []*AllVirtualPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPart{}).Where("`max_used_sub_part_id` = ?", maxUsedSubPartID).Find(&results).Error

	return
}

// GetBatchFromMaxUsedSubPartID 批量查找
func (obj *_AllVirtualPartMgr) GetBatchFromMaxUsedSubPartID(maxUsedSubPartIDs []int64) (results []*AllVirtualPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPart{}).Where("`max_used_sub_part_id` IN (?)", maxUsedSubPartIDs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllVirtualPartMgr) FetchByPrimaryKey(tenantID int64, tableID int64, partID int64) (result AllVirtualPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPart{}).Where("`tenant_id` = ? AND `table_id` = ? AND `part_id` = ?", tenantID, tableID, partID).First(&result).Error

	return
}
