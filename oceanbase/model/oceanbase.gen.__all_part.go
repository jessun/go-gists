package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _AllPartMgr struct {
	*_BaseMgr
}

// AllPartMgr open func
func AllPartMgr(db *gorm.DB) *_AllPartMgr {
	if db == nil {
		panic(fmt.Errorf("AllPartMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllPartMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_part"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllPartMgr) GetTableName() string {
	return "__all_part"
}

// Reset 重置gorm会话
func (obj *_AllPartMgr) Reset() *_AllPartMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllPartMgr) Get() (result AllPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPart{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllPartMgr) Gets() (results []*AllPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPart{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllPartMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllPart{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithGmtCreate gmt_create获取
func (obj *_AllPartMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllPartMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithTenantID tenant_id获取
func (obj *_AllPartMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithTableID table_id获取
func (obj *_AllPartMgr) WithTableID(tableID int64) Option {
	return optionFunc(func(o *options) { o.query["table_id"] = tableID })
}

// WithPartID part_id获取
func (obj *_AllPartMgr) WithPartID(partID int64) Option {
	return optionFunc(func(o *options) { o.query["part_id"] = partID })
}

// WithPartName part_name获取
func (obj *_AllPartMgr) WithPartName(partName string) Option {
	return optionFunc(func(o *options) { o.query["part_name"] = partName })
}

// WithSchemaVersion schema_version获取
func (obj *_AllPartMgr) WithSchemaVersion(schemaVersion int64) Option {
	return optionFunc(func(o *options) { o.query["schema_version"] = schemaVersion })
}

// WithHighBoundVal high_bound_val获取
func (obj *_AllPartMgr) WithHighBoundVal(highBoundVal string) Option {
	return optionFunc(func(o *options) { o.query["high_bound_val"] = highBoundVal })
}

// WithBHighBoundVal b_high_bound_val获取
func (obj *_AllPartMgr) WithBHighBoundVal(bHighBoundVal string) Option {
	return optionFunc(func(o *options) { o.query["b_high_bound_val"] = bHighBoundVal })
}

// WithSubPartNum sub_part_num获取
func (obj *_AllPartMgr) WithSubPartNum(subPartNum int64) Option {
	return optionFunc(func(o *options) { o.query["sub_part_num"] = subPartNum })
}

// WithSubPartSpace sub_part_space获取
func (obj *_AllPartMgr) WithSubPartSpace(subPartSpace int64) Option {
	return optionFunc(func(o *options) { o.query["sub_part_space"] = subPartSpace })
}

// WithNewSubPartNum new_sub_part_num获取
func (obj *_AllPartMgr) WithNewSubPartNum(newSubPartNum int64) Option {
	return optionFunc(func(o *options) { o.query["new_sub_part_num"] = newSubPartNum })
}

// WithNewSubPartSpace new_sub_part_space获取
func (obj *_AllPartMgr) WithNewSubPartSpace(newSubPartSpace int64) Option {
	return optionFunc(func(o *options) { o.query["new_sub_part_space"] = newSubPartSpace })
}

// WithSubPartInterval sub_part_interval获取
func (obj *_AllPartMgr) WithSubPartInterval(subPartInterval string) Option {
	return optionFunc(func(o *options) { o.query["sub_part_interval"] = subPartInterval })
}

// WithSubIntervalStart sub_interval_start获取
func (obj *_AllPartMgr) WithSubIntervalStart(subIntervalStart string) Option {
	return optionFunc(func(o *options) { o.query["sub_interval_start"] = subIntervalStart })
}

// WithNewSubPartInterval new_sub_part_interval获取
func (obj *_AllPartMgr) WithNewSubPartInterval(newSubPartInterval string) Option {
	return optionFunc(func(o *options) { o.query["new_sub_part_interval"] = newSubPartInterval })
}

// WithNewSubIntervalStart new_sub_interval_start获取
func (obj *_AllPartMgr) WithNewSubIntervalStart(newSubIntervalStart string) Option {
	return optionFunc(func(o *options) { o.query["new_sub_interval_start"] = newSubIntervalStart })
}

// WithBlockSize block_size获取
func (obj *_AllPartMgr) WithBlockSize(blockSize int64) Option {
	return optionFunc(func(o *options) { o.query["block_size"] = blockSize })
}

// WithReplicaNum replica_num获取
func (obj *_AllPartMgr) WithReplicaNum(replicaNum int64) Option {
	return optionFunc(func(o *options) { o.query["replica_num"] = replicaNum })
}

// WithCompressFuncName compress_func_name获取
func (obj *_AllPartMgr) WithCompressFuncName(compressFuncName string) Option {
	return optionFunc(func(o *options) { o.query["compress_func_name"] = compressFuncName })
}

// WithStatus status获取
func (obj *_AllPartMgr) WithStatus(status int64) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithSpare1 spare1获取
func (obj *_AllPartMgr) WithSpare1(spare1 int64) Option {
	return optionFunc(func(o *options) { o.query["spare1"] = spare1 })
}

// WithSpare2 spare2获取
func (obj *_AllPartMgr) WithSpare2(spare2 int64) Option {
	return optionFunc(func(o *options) { o.query["spare2"] = spare2 })
}

// WithSpare3 spare3获取
func (obj *_AllPartMgr) WithSpare3(spare3 string) Option {
	return optionFunc(func(o *options) { o.query["spare3"] = spare3 })
}

// WithComment comment获取
func (obj *_AllPartMgr) WithComment(comment string) Option {
	return optionFunc(func(o *options) { o.query["comment"] = comment })
}

// WithListVal list_val获取
func (obj *_AllPartMgr) WithListVal(listVal string) Option {
	return optionFunc(func(o *options) { o.query["list_val"] = listVal })
}

// WithBListVal b_list_val获取
func (obj *_AllPartMgr) WithBListVal(bListVal string) Option {
	return optionFunc(func(o *options) { o.query["b_list_val"] = bListVal })
}

// WithPartIDx part_idx获取
func (obj *_AllPartMgr) WithPartIDx(partIDx int64) Option {
	return optionFunc(func(o *options) { o.query["part_idx"] = partIDx })
}

// WithSourcePartitionID source_partition_id获取
func (obj *_AllPartMgr) WithSourcePartitionID(sourcePartitionID string) Option {
	return optionFunc(func(o *options) { o.query["source_partition_id"] = sourcePartitionID })
}

// WithMappingPgPartID mapping_pg_part_id获取
func (obj *_AllPartMgr) WithMappingPgPartID(mappingPgPartID int64) Option {
	return optionFunc(func(o *options) { o.query["mapping_pg_part_id"] = mappingPgPartID })
}

// WithTablespaceID tablespace_id获取
func (obj *_AllPartMgr) WithTablespaceID(tablespaceID int64) Option {
	return optionFunc(func(o *options) { o.query["tablespace_id"] = tablespaceID })
}

// WithDropSchemaVersion drop_schema_version获取
func (obj *_AllPartMgr) WithDropSchemaVersion(dropSchemaVersion int64) Option {
	return optionFunc(func(o *options) { o.query["drop_schema_version"] = dropSchemaVersion })
}

// WithMaxUsedSubPartID max_used_sub_part_id获取
func (obj *_AllPartMgr) WithMaxUsedSubPartID(maxUsedSubPartID int64) Option {
	return optionFunc(func(o *options) { o.query["max_used_sub_part_id"] = maxUsedSubPartID })
}

// GetByOption 功能选项模式获取
func (obj *_AllPartMgr) GetByOption(opts ...Option) (result AllPart, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllPart{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllPartMgr) GetByOptions(opts ...Option) (results []*AllPart, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllPart{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllPartMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllPart, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllPart{}).Where(options.query)
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
func (obj *_AllPartMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPart{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllPartMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPart{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllPartMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPart{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllPartMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPart{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromTenantID 通过tenant_id获取内容
func (obj *_AllPartMgr) GetFromTenantID(tenantID int64) (results []*AllPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPart{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllPartMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPart{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromTableID 通过table_id获取内容
func (obj *_AllPartMgr) GetFromTableID(tableID int64) (results []*AllPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPart{}).Where("`table_id` = ?", tableID).Find(&results).Error

	return
}

// GetBatchFromTableID 批量查找
func (obj *_AllPartMgr) GetBatchFromTableID(tableIDs []int64) (results []*AllPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPart{}).Where("`table_id` IN (?)", tableIDs).Find(&results).Error

	return
}

// GetFromPartID 通过part_id获取内容
func (obj *_AllPartMgr) GetFromPartID(partID int64) (results []*AllPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPart{}).Where("`part_id` = ?", partID).Find(&results).Error

	return
}

// GetBatchFromPartID 批量查找
func (obj *_AllPartMgr) GetBatchFromPartID(partIDs []int64) (results []*AllPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPart{}).Where("`part_id` IN (?)", partIDs).Find(&results).Error

	return
}

// GetFromPartName 通过part_name获取内容
func (obj *_AllPartMgr) GetFromPartName(partName string) (results []*AllPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPart{}).Where("`part_name` = ?", partName).Find(&results).Error

	return
}

// GetBatchFromPartName 批量查找
func (obj *_AllPartMgr) GetBatchFromPartName(partNames []string) (results []*AllPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPart{}).Where("`part_name` IN (?)", partNames).Find(&results).Error

	return
}

// GetFromSchemaVersion 通过schema_version获取内容
func (obj *_AllPartMgr) GetFromSchemaVersion(schemaVersion int64) (results []*AllPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPart{}).Where("`schema_version` = ?", schemaVersion).Find(&results).Error

	return
}

// GetBatchFromSchemaVersion 批量查找
func (obj *_AllPartMgr) GetBatchFromSchemaVersion(schemaVersions []int64) (results []*AllPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPart{}).Where("`schema_version` IN (?)", schemaVersions).Find(&results).Error

	return
}

// GetFromHighBoundVal 通过high_bound_val获取内容
func (obj *_AllPartMgr) GetFromHighBoundVal(highBoundVal string) (results []*AllPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPart{}).Where("`high_bound_val` = ?", highBoundVal).Find(&results).Error

	return
}

// GetBatchFromHighBoundVal 批量查找
func (obj *_AllPartMgr) GetBatchFromHighBoundVal(highBoundVals []string) (results []*AllPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPart{}).Where("`high_bound_val` IN (?)", highBoundVals).Find(&results).Error

	return
}

// GetFromBHighBoundVal 通过b_high_bound_val获取内容
func (obj *_AllPartMgr) GetFromBHighBoundVal(bHighBoundVal string) (results []*AllPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPart{}).Where("`b_high_bound_val` = ?", bHighBoundVal).Find(&results).Error

	return
}

// GetBatchFromBHighBoundVal 批量查找
func (obj *_AllPartMgr) GetBatchFromBHighBoundVal(bHighBoundVals []string) (results []*AllPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPart{}).Where("`b_high_bound_val` IN (?)", bHighBoundVals).Find(&results).Error

	return
}

// GetFromSubPartNum 通过sub_part_num获取内容
func (obj *_AllPartMgr) GetFromSubPartNum(subPartNum int64) (results []*AllPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPart{}).Where("`sub_part_num` = ?", subPartNum).Find(&results).Error

	return
}

// GetBatchFromSubPartNum 批量查找
func (obj *_AllPartMgr) GetBatchFromSubPartNum(subPartNums []int64) (results []*AllPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPart{}).Where("`sub_part_num` IN (?)", subPartNums).Find(&results).Error

	return
}

// GetFromSubPartSpace 通过sub_part_space获取内容
func (obj *_AllPartMgr) GetFromSubPartSpace(subPartSpace int64) (results []*AllPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPart{}).Where("`sub_part_space` = ?", subPartSpace).Find(&results).Error

	return
}

// GetBatchFromSubPartSpace 批量查找
func (obj *_AllPartMgr) GetBatchFromSubPartSpace(subPartSpaces []int64) (results []*AllPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPart{}).Where("`sub_part_space` IN (?)", subPartSpaces).Find(&results).Error

	return
}

// GetFromNewSubPartNum 通过new_sub_part_num获取内容
func (obj *_AllPartMgr) GetFromNewSubPartNum(newSubPartNum int64) (results []*AllPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPart{}).Where("`new_sub_part_num` = ?", newSubPartNum).Find(&results).Error

	return
}

// GetBatchFromNewSubPartNum 批量查找
func (obj *_AllPartMgr) GetBatchFromNewSubPartNum(newSubPartNums []int64) (results []*AllPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPart{}).Where("`new_sub_part_num` IN (?)", newSubPartNums).Find(&results).Error

	return
}

// GetFromNewSubPartSpace 通过new_sub_part_space获取内容
func (obj *_AllPartMgr) GetFromNewSubPartSpace(newSubPartSpace int64) (results []*AllPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPart{}).Where("`new_sub_part_space` = ?", newSubPartSpace).Find(&results).Error

	return
}

// GetBatchFromNewSubPartSpace 批量查找
func (obj *_AllPartMgr) GetBatchFromNewSubPartSpace(newSubPartSpaces []int64) (results []*AllPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPart{}).Where("`new_sub_part_space` IN (?)", newSubPartSpaces).Find(&results).Error

	return
}

// GetFromSubPartInterval 通过sub_part_interval获取内容
func (obj *_AllPartMgr) GetFromSubPartInterval(subPartInterval string) (results []*AllPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPart{}).Where("`sub_part_interval` = ?", subPartInterval).Find(&results).Error

	return
}

// GetBatchFromSubPartInterval 批量查找
func (obj *_AllPartMgr) GetBatchFromSubPartInterval(subPartIntervals []string) (results []*AllPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPart{}).Where("`sub_part_interval` IN (?)", subPartIntervals).Find(&results).Error

	return
}

// GetFromSubIntervalStart 通过sub_interval_start获取内容
func (obj *_AllPartMgr) GetFromSubIntervalStart(subIntervalStart string) (results []*AllPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPart{}).Where("`sub_interval_start` = ?", subIntervalStart).Find(&results).Error

	return
}

// GetBatchFromSubIntervalStart 批量查找
func (obj *_AllPartMgr) GetBatchFromSubIntervalStart(subIntervalStarts []string) (results []*AllPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPart{}).Where("`sub_interval_start` IN (?)", subIntervalStarts).Find(&results).Error

	return
}

// GetFromNewSubPartInterval 通过new_sub_part_interval获取内容
func (obj *_AllPartMgr) GetFromNewSubPartInterval(newSubPartInterval string) (results []*AllPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPart{}).Where("`new_sub_part_interval` = ?", newSubPartInterval).Find(&results).Error

	return
}

// GetBatchFromNewSubPartInterval 批量查找
func (obj *_AllPartMgr) GetBatchFromNewSubPartInterval(newSubPartIntervals []string) (results []*AllPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPart{}).Where("`new_sub_part_interval` IN (?)", newSubPartIntervals).Find(&results).Error

	return
}

// GetFromNewSubIntervalStart 通过new_sub_interval_start获取内容
func (obj *_AllPartMgr) GetFromNewSubIntervalStart(newSubIntervalStart string) (results []*AllPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPart{}).Where("`new_sub_interval_start` = ?", newSubIntervalStart).Find(&results).Error

	return
}

// GetBatchFromNewSubIntervalStart 批量查找
func (obj *_AllPartMgr) GetBatchFromNewSubIntervalStart(newSubIntervalStarts []string) (results []*AllPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPart{}).Where("`new_sub_interval_start` IN (?)", newSubIntervalStarts).Find(&results).Error

	return
}

// GetFromBlockSize 通过block_size获取内容
func (obj *_AllPartMgr) GetFromBlockSize(blockSize int64) (results []*AllPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPart{}).Where("`block_size` = ?", blockSize).Find(&results).Error

	return
}

// GetBatchFromBlockSize 批量查找
func (obj *_AllPartMgr) GetBatchFromBlockSize(blockSizes []int64) (results []*AllPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPart{}).Where("`block_size` IN (?)", blockSizes).Find(&results).Error

	return
}

// GetFromReplicaNum 通过replica_num获取内容
func (obj *_AllPartMgr) GetFromReplicaNum(replicaNum int64) (results []*AllPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPart{}).Where("`replica_num` = ?", replicaNum).Find(&results).Error

	return
}

// GetBatchFromReplicaNum 批量查找
func (obj *_AllPartMgr) GetBatchFromReplicaNum(replicaNums []int64) (results []*AllPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPart{}).Where("`replica_num` IN (?)", replicaNums).Find(&results).Error

	return
}

// GetFromCompressFuncName 通过compress_func_name获取内容
func (obj *_AllPartMgr) GetFromCompressFuncName(compressFuncName string) (results []*AllPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPart{}).Where("`compress_func_name` = ?", compressFuncName).Find(&results).Error

	return
}

// GetBatchFromCompressFuncName 批量查找
func (obj *_AllPartMgr) GetBatchFromCompressFuncName(compressFuncNames []string) (results []*AllPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPart{}).Where("`compress_func_name` IN (?)", compressFuncNames).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容
func (obj *_AllPartMgr) GetFromStatus(status int64) (results []*AllPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPart{}).Where("`status` = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量查找
func (obj *_AllPartMgr) GetBatchFromStatus(statuss []int64) (results []*AllPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPart{}).Where("`status` IN (?)", statuss).Find(&results).Error

	return
}

// GetFromSpare1 通过spare1获取内容
func (obj *_AllPartMgr) GetFromSpare1(spare1 int64) (results []*AllPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPart{}).Where("`spare1` = ?", spare1).Find(&results).Error

	return
}

// GetBatchFromSpare1 批量查找
func (obj *_AllPartMgr) GetBatchFromSpare1(spare1s []int64) (results []*AllPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPart{}).Where("`spare1` IN (?)", spare1s).Find(&results).Error

	return
}

// GetFromSpare2 通过spare2获取内容
func (obj *_AllPartMgr) GetFromSpare2(spare2 int64) (results []*AllPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPart{}).Where("`spare2` = ?", spare2).Find(&results).Error

	return
}

// GetBatchFromSpare2 批量查找
func (obj *_AllPartMgr) GetBatchFromSpare2(spare2s []int64) (results []*AllPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPart{}).Where("`spare2` IN (?)", spare2s).Find(&results).Error

	return
}

// GetFromSpare3 通过spare3获取内容
func (obj *_AllPartMgr) GetFromSpare3(spare3 string) (results []*AllPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPart{}).Where("`spare3` = ?", spare3).Find(&results).Error

	return
}

// GetBatchFromSpare3 批量查找
func (obj *_AllPartMgr) GetBatchFromSpare3(spare3s []string) (results []*AllPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPart{}).Where("`spare3` IN (?)", spare3s).Find(&results).Error

	return
}

// GetFromComment 通过comment获取内容
func (obj *_AllPartMgr) GetFromComment(comment string) (results []*AllPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPart{}).Where("`comment` = ?", comment).Find(&results).Error

	return
}

// GetBatchFromComment 批量查找
func (obj *_AllPartMgr) GetBatchFromComment(comments []string) (results []*AllPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPart{}).Where("`comment` IN (?)", comments).Find(&results).Error

	return
}

// GetFromListVal 通过list_val获取内容
func (obj *_AllPartMgr) GetFromListVal(listVal string) (results []*AllPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPart{}).Where("`list_val` = ?", listVal).Find(&results).Error

	return
}

// GetBatchFromListVal 批量查找
func (obj *_AllPartMgr) GetBatchFromListVal(listVals []string) (results []*AllPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPart{}).Where("`list_val` IN (?)", listVals).Find(&results).Error

	return
}

// GetFromBListVal 通过b_list_val获取内容
func (obj *_AllPartMgr) GetFromBListVal(bListVal string) (results []*AllPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPart{}).Where("`b_list_val` = ?", bListVal).Find(&results).Error

	return
}

// GetBatchFromBListVal 批量查找
func (obj *_AllPartMgr) GetBatchFromBListVal(bListVals []string) (results []*AllPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPart{}).Where("`b_list_val` IN (?)", bListVals).Find(&results).Error

	return
}

// GetFromPartIDx 通过part_idx获取内容
func (obj *_AllPartMgr) GetFromPartIDx(partIDx int64) (results []*AllPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPart{}).Where("`part_idx` = ?", partIDx).Find(&results).Error

	return
}

// GetBatchFromPartIDx 批量查找
func (obj *_AllPartMgr) GetBatchFromPartIDx(partIDxs []int64) (results []*AllPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPart{}).Where("`part_idx` IN (?)", partIDxs).Find(&results).Error

	return
}

// GetFromSourcePartitionID 通过source_partition_id获取内容
func (obj *_AllPartMgr) GetFromSourcePartitionID(sourcePartitionID string) (results []*AllPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPart{}).Where("`source_partition_id` = ?", sourcePartitionID).Find(&results).Error

	return
}

// GetBatchFromSourcePartitionID 批量查找
func (obj *_AllPartMgr) GetBatchFromSourcePartitionID(sourcePartitionIDs []string) (results []*AllPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPart{}).Where("`source_partition_id` IN (?)", sourcePartitionIDs).Find(&results).Error

	return
}

// GetFromMappingPgPartID 通过mapping_pg_part_id获取内容
func (obj *_AllPartMgr) GetFromMappingPgPartID(mappingPgPartID int64) (results []*AllPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPart{}).Where("`mapping_pg_part_id` = ?", mappingPgPartID).Find(&results).Error

	return
}

// GetBatchFromMappingPgPartID 批量查找
func (obj *_AllPartMgr) GetBatchFromMappingPgPartID(mappingPgPartIDs []int64) (results []*AllPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPart{}).Where("`mapping_pg_part_id` IN (?)", mappingPgPartIDs).Find(&results).Error

	return
}

// GetFromTablespaceID 通过tablespace_id获取内容
func (obj *_AllPartMgr) GetFromTablespaceID(tablespaceID int64) (results []*AllPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPart{}).Where("`tablespace_id` = ?", tablespaceID).Find(&results).Error

	return
}

// GetBatchFromTablespaceID 批量查找
func (obj *_AllPartMgr) GetBatchFromTablespaceID(tablespaceIDs []int64) (results []*AllPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPart{}).Where("`tablespace_id` IN (?)", tablespaceIDs).Find(&results).Error

	return
}

// GetFromDropSchemaVersion 通过drop_schema_version获取内容
func (obj *_AllPartMgr) GetFromDropSchemaVersion(dropSchemaVersion int64) (results []*AllPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPart{}).Where("`drop_schema_version` = ?", dropSchemaVersion).Find(&results).Error

	return
}

// GetBatchFromDropSchemaVersion 批量查找
func (obj *_AllPartMgr) GetBatchFromDropSchemaVersion(dropSchemaVersions []int64) (results []*AllPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPart{}).Where("`drop_schema_version` IN (?)", dropSchemaVersions).Find(&results).Error

	return
}

// GetFromMaxUsedSubPartID 通过max_used_sub_part_id获取内容
func (obj *_AllPartMgr) GetFromMaxUsedSubPartID(maxUsedSubPartID int64) (results []*AllPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPart{}).Where("`max_used_sub_part_id` = ?", maxUsedSubPartID).Find(&results).Error

	return
}

// GetBatchFromMaxUsedSubPartID 批量查找
func (obj *_AllPartMgr) GetBatchFromMaxUsedSubPartID(maxUsedSubPartIDs []int64) (results []*AllPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPart{}).Where("`max_used_sub_part_id` IN (?)", maxUsedSubPartIDs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllPartMgr) FetchByPrimaryKey(tenantID int64, tableID int64, partID int64) (result AllPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPart{}).Where("`tenant_id` = ? AND `table_id` = ? AND `part_id` = ?", tenantID, tableID, partID).First(&result).Error

	return
}
