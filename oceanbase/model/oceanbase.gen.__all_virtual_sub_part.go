package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _AllVirtualSubPartMgr struct {
	*_BaseMgr
}

// AllVirtualSubPartMgr open func
func AllVirtualSubPartMgr(db *gorm.DB) *_AllVirtualSubPartMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualSubPartMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualSubPartMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_sub_part"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualSubPartMgr) GetTableName() string {
	return "__all_virtual_sub_part"
}

// Reset 重置gorm会话
func (obj *_AllVirtualSubPartMgr) Reset() *_AllVirtualSubPartMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualSubPartMgr) Get() (result AllVirtualSubPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSubPart{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualSubPartMgr) Gets() (results []*AllVirtualSubPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSubPart{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualSubPartMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualSubPart{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTenantID tenant_id获取
func (obj *_AllVirtualSubPartMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithTableID table_id获取
func (obj *_AllVirtualSubPartMgr) WithTableID(tableID int64) Option {
	return optionFunc(func(o *options) { o.query["table_id"] = tableID })
}

// WithPartID part_id获取
func (obj *_AllVirtualSubPartMgr) WithPartID(partID int64) Option {
	return optionFunc(func(o *options) { o.query["part_id"] = partID })
}

// WithSubPartID sub_part_id获取
func (obj *_AllVirtualSubPartMgr) WithSubPartID(subPartID int64) Option {
	return optionFunc(func(o *options) { o.query["sub_part_id"] = subPartID })
}

// WithGmtCreate gmt_create获取
func (obj *_AllVirtualSubPartMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllVirtualSubPartMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithSubPartName sub_part_name获取
func (obj *_AllVirtualSubPartMgr) WithSubPartName(subPartName string) Option {
	return optionFunc(func(o *options) { o.query["sub_part_name"] = subPartName })
}

// WithSchemaVersion schema_version获取
func (obj *_AllVirtualSubPartMgr) WithSchemaVersion(schemaVersion int64) Option {
	return optionFunc(func(o *options) { o.query["schema_version"] = schemaVersion })
}

// WithHighBoundVal high_bound_val获取
func (obj *_AllVirtualSubPartMgr) WithHighBoundVal(highBoundVal string) Option {
	return optionFunc(func(o *options) { o.query["high_bound_val"] = highBoundVal })
}

// WithBHighBoundVal b_high_bound_val获取
func (obj *_AllVirtualSubPartMgr) WithBHighBoundVal(bHighBoundVal string) Option {
	return optionFunc(func(o *options) { o.query["b_high_bound_val"] = bHighBoundVal })
}

// WithBlockSize block_size获取
func (obj *_AllVirtualSubPartMgr) WithBlockSize(blockSize int64) Option {
	return optionFunc(func(o *options) { o.query["block_size"] = blockSize })
}

// WithReplicaNum replica_num获取
func (obj *_AllVirtualSubPartMgr) WithReplicaNum(replicaNum int64) Option {
	return optionFunc(func(o *options) { o.query["replica_num"] = replicaNum })
}

// WithCompressFuncName compress_func_name获取
func (obj *_AllVirtualSubPartMgr) WithCompressFuncName(compressFuncName string) Option {
	return optionFunc(func(o *options) { o.query["compress_func_name"] = compressFuncName })
}

// WithStatus status获取
func (obj *_AllVirtualSubPartMgr) WithStatus(status int64) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithSpare1 spare1获取
func (obj *_AllVirtualSubPartMgr) WithSpare1(spare1 int64) Option {
	return optionFunc(func(o *options) { o.query["spare1"] = spare1 })
}

// WithSpare2 spare2获取
func (obj *_AllVirtualSubPartMgr) WithSpare2(spare2 int64) Option {
	return optionFunc(func(o *options) { o.query["spare2"] = spare2 })
}

// WithSpare3 spare3获取
func (obj *_AllVirtualSubPartMgr) WithSpare3(spare3 string) Option {
	return optionFunc(func(o *options) { o.query["spare3"] = spare3 })
}

// WithComment comment获取
func (obj *_AllVirtualSubPartMgr) WithComment(comment string) Option {
	return optionFunc(func(o *options) { o.query["comment"] = comment })
}

// WithListVal list_val获取
func (obj *_AllVirtualSubPartMgr) WithListVal(listVal string) Option {
	return optionFunc(func(o *options) { o.query["list_val"] = listVal })
}

// WithBListVal b_list_val获取
func (obj *_AllVirtualSubPartMgr) WithBListVal(bListVal string) Option {
	return optionFunc(func(o *options) { o.query["b_list_val"] = bListVal })
}

// WithTablespaceID tablespace_id获取
func (obj *_AllVirtualSubPartMgr) WithTablespaceID(tablespaceID int64) Option {
	return optionFunc(func(o *options) { o.query["tablespace_id"] = tablespaceID })
}

// WithSubPartIDx sub_part_idx获取
func (obj *_AllVirtualSubPartMgr) WithSubPartIDx(subPartIDx int64) Option {
	return optionFunc(func(o *options) { o.query["sub_part_idx"] = subPartIDx })
}

// WithSourcePartitionID source_partition_id获取
func (obj *_AllVirtualSubPartMgr) WithSourcePartitionID(sourcePartitionID string) Option {
	return optionFunc(func(o *options) { o.query["source_partition_id"] = sourcePartitionID })
}

// WithMappingPgSubPartID mapping_pg_sub_part_id获取
func (obj *_AllVirtualSubPartMgr) WithMappingPgSubPartID(mappingPgSubPartID int64) Option {
	return optionFunc(func(o *options) { o.query["mapping_pg_sub_part_id"] = mappingPgSubPartID })
}

// WithDropSchemaVersion drop_schema_version获取
func (obj *_AllVirtualSubPartMgr) WithDropSchemaVersion(dropSchemaVersion int64) Option {
	return optionFunc(func(o *options) { o.query["drop_schema_version"] = dropSchemaVersion })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualSubPartMgr) GetByOption(opts ...Option) (result AllVirtualSubPart, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSubPart{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualSubPartMgr) GetByOptions(opts ...Option) (results []*AllVirtualSubPart, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSubPart{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualSubPartMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualSubPart, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualSubPart{}).Where(options.query)
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
func (obj *_AllVirtualSubPartMgr) GetFromTenantID(tenantID int64) (results []*AllVirtualSubPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSubPart{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllVirtualSubPartMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllVirtualSubPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSubPart{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromTableID 通过table_id获取内容
func (obj *_AllVirtualSubPartMgr) GetFromTableID(tableID int64) (results []*AllVirtualSubPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSubPart{}).Where("`table_id` = ?", tableID).Find(&results).Error

	return
}

// GetBatchFromTableID 批量查找
func (obj *_AllVirtualSubPartMgr) GetBatchFromTableID(tableIDs []int64) (results []*AllVirtualSubPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSubPart{}).Where("`table_id` IN (?)", tableIDs).Find(&results).Error

	return
}

// GetFromPartID 通过part_id获取内容
func (obj *_AllVirtualSubPartMgr) GetFromPartID(partID int64) (results []*AllVirtualSubPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSubPart{}).Where("`part_id` = ?", partID).Find(&results).Error

	return
}

// GetBatchFromPartID 批量查找
func (obj *_AllVirtualSubPartMgr) GetBatchFromPartID(partIDs []int64) (results []*AllVirtualSubPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSubPart{}).Where("`part_id` IN (?)", partIDs).Find(&results).Error

	return
}

// GetFromSubPartID 通过sub_part_id获取内容
func (obj *_AllVirtualSubPartMgr) GetFromSubPartID(subPartID int64) (results []*AllVirtualSubPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSubPart{}).Where("`sub_part_id` = ?", subPartID).Find(&results).Error

	return
}

// GetBatchFromSubPartID 批量查找
func (obj *_AllVirtualSubPartMgr) GetBatchFromSubPartID(subPartIDs []int64) (results []*AllVirtualSubPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSubPart{}).Where("`sub_part_id` IN (?)", subPartIDs).Find(&results).Error

	return
}

// GetFromGmtCreate 通过gmt_create获取内容
func (obj *_AllVirtualSubPartMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllVirtualSubPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSubPart{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllVirtualSubPartMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllVirtualSubPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSubPart{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllVirtualSubPartMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllVirtualSubPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSubPart{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllVirtualSubPartMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllVirtualSubPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSubPart{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromSubPartName 通过sub_part_name获取内容
func (obj *_AllVirtualSubPartMgr) GetFromSubPartName(subPartName string) (results []*AllVirtualSubPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSubPart{}).Where("`sub_part_name` = ?", subPartName).Find(&results).Error

	return
}

// GetBatchFromSubPartName 批量查找
func (obj *_AllVirtualSubPartMgr) GetBatchFromSubPartName(subPartNames []string) (results []*AllVirtualSubPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSubPart{}).Where("`sub_part_name` IN (?)", subPartNames).Find(&results).Error

	return
}

// GetFromSchemaVersion 通过schema_version获取内容
func (obj *_AllVirtualSubPartMgr) GetFromSchemaVersion(schemaVersion int64) (results []*AllVirtualSubPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSubPart{}).Where("`schema_version` = ?", schemaVersion).Find(&results).Error

	return
}

// GetBatchFromSchemaVersion 批量查找
func (obj *_AllVirtualSubPartMgr) GetBatchFromSchemaVersion(schemaVersions []int64) (results []*AllVirtualSubPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSubPart{}).Where("`schema_version` IN (?)", schemaVersions).Find(&results).Error

	return
}

// GetFromHighBoundVal 通过high_bound_val获取内容
func (obj *_AllVirtualSubPartMgr) GetFromHighBoundVal(highBoundVal string) (results []*AllVirtualSubPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSubPart{}).Where("`high_bound_val` = ?", highBoundVal).Find(&results).Error

	return
}

// GetBatchFromHighBoundVal 批量查找
func (obj *_AllVirtualSubPartMgr) GetBatchFromHighBoundVal(highBoundVals []string) (results []*AllVirtualSubPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSubPart{}).Where("`high_bound_val` IN (?)", highBoundVals).Find(&results).Error

	return
}

// GetFromBHighBoundVal 通过b_high_bound_val获取内容
func (obj *_AllVirtualSubPartMgr) GetFromBHighBoundVal(bHighBoundVal string) (results []*AllVirtualSubPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSubPart{}).Where("`b_high_bound_val` = ?", bHighBoundVal).Find(&results).Error

	return
}

// GetBatchFromBHighBoundVal 批量查找
func (obj *_AllVirtualSubPartMgr) GetBatchFromBHighBoundVal(bHighBoundVals []string) (results []*AllVirtualSubPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSubPart{}).Where("`b_high_bound_val` IN (?)", bHighBoundVals).Find(&results).Error

	return
}

// GetFromBlockSize 通过block_size获取内容
func (obj *_AllVirtualSubPartMgr) GetFromBlockSize(blockSize int64) (results []*AllVirtualSubPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSubPart{}).Where("`block_size` = ?", blockSize).Find(&results).Error

	return
}

// GetBatchFromBlockSize 批量查找
func (obj *_AllVirtualSubPartMgr) GetBatchFromBlockSize(blockSizes []int64) (results []*AllVirtualSubPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSubPart{}).Where("`block_size` IN (?)", blockSizes).Find(&results).Error

	return
}

// GetFromReplicaNum 通过replica_num获取内容
func (obj *_AllVirtualSubPartMgr) GetFromReplicaNum(replicaNum int64) (results []*AllVirtualSubPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSubPart{}).Where("`replica_num` = ?", replicaNum).Find(&results).Error

	return
}

// GetBatchFromReplicaNum 批量查找
func (obj *_AllVirtualSubPartMgr) GetBatchFromReplicaNum(replicaNums []int64) (results []*AllVirtualSubPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSubPart{}).Where("`replica_num` IN (?)", replicaNums).Find(&results).Error

	return
}

// GetFromCompressFuncName 通过compress_func_name获取内容
func (obj *_AllVirtualSubPartMgr) GetFromCompressFuncName(compressFuncName string) (results []*AllVirtualSubPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSubPart{}).Where("`compress_func_name` = ?", compressFuncName).Find(&results).Error

	return
}

// GetBatchFromCompressFuncName 批量查找
func (obj *_AllVirtualSubPartMgr) GetBatchFromCompressFuncName(compressFuncNames []string) (results []*AllVirtualSubPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSubPart{}).Where("`compress_func_name` IN (?)", compressFuncNames).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容
func (obj *_AllVirtualSubPartMgr) GetFromStatus(status int64) (results []*AllVirtualSubPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSubPart{}).Where("`status` = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量查找
func (obj *_AllVirtualSubPartMgr) GetBatchFromStatus(statuss []int64) (results []*AllVirtualSubPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSubPart{}).Where("`status` IN (?)", statuss).Find(&results).Error

	return
}

// GetFromSpare1 通过spare1获取内容
func (obj *_AllVirtualSubPartMgr) GetFromSpare1(spare1 int64) (results []*AllVirtualSubPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSubPart{}).Where("`spare1` = ?", spare1).Find(&results).Error

	return
}

// GetBatchFromSpare1 批量查找
func (obj *_AllVirtualSubPartMgr) GetBatchFromSpare1(spare1s []int64) (results []*AllVirtualSubPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSubPart{}).Where("`spare1` IN (?)", spare1s).Find(&results).Error

	return
}

// GetFromSpare2 通过spare2获取内容
func (obj *_AllVirtualSubPartMgr) GetFromSpare2(spare2 int64) (results []*AllVirtualSubPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSubPart{}).Where("`spare2` = ?", spare2).Find(&results).Error

	return
}

// GetBatchFromSpare2 批量查找
func (obj *_AllVirtualSubPartMgr) GetBatchFromSpare2(spare2s []int64) (results []*AllVirtualSubPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSubPart{}).Where("`spare2` IN (?)", spare2s).Find(&results).Error

	return
}

// GetFromSpare3 通过spare3获取内容
func (obj *_AllVirtualSubPartMgr) GetFromSpare3(spare3 string) (results []*AllVirtualSubPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSubPart{}).Where("`spare3` = ?", spare3).Find(&results).Error

	return
}

// GetBatchFromSpare3 批量查找
func (obj *_AllVirtualSubPartMgr) GetBatchFromSpare3(spare3s []string) (results []*AllVirtualSubPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSubPart{}).Where("`spare3` IN (?)", spare3s).Find(&results).Error

	return
}

// GetFromComment 通过comment获取内容
func (obj *_AllVirtualSubPartMgr) GetFromComment(comment string) (results []*AllVirtualSubPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSubPart{}).Where("`comment` = ?", comment).Find(&results).Error

	return
}

// GetBatchFromComment 批量查找
func (obj *_AllVirtualSubPartMgr) GetBatchFromComment(comments []string) (results []*AllVirtualSubPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSubPart{}).Where("`comment` IN (?)", comments).Find(&results).Error

	return
}

// GetFromListVal 通过list_val获取内容
func (obj *_AllVirtualSubPartMgr) GetFromListVal(listVal string) (results []*AllVirtualSubPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSubPart{}).Where("`list_val` = ?", listVal).Find(&results).Error

	return
}

// GetBatchFromListVal 批量查找
func (obj *_AllVirtualSubPartMgr) GetBatchFromListVal(listVals []string) (results []*AllVirtualSubPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSubPart{}).Where("`list_val` IN (?)", listVals).Find(&results).Error

	return
}

// GetFromBListVal 通过b_list_val获取内容
func (obj *_AllVirtualSubPartMgr) GetFromBListVal(bListVal string) (results []*AllVirtualSubPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSubPart{}).Where("`b_list_val` = ?", bListVal).Find(&results).Error

	return
}

// GetBatchFromBListVal 批量查找
func (obj *_AllVirtualSubPartMgr) GetBatchFromBListVal(bListVals []string) (results []*AllVirtualSubPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSubPart{}).Where("`b_list_val` IN (?)", bListVals).Find(&results).Error

	return
}

// GetFromTablespaceID 通过tablespace_id获取内容
func (obj *_AllVirtualSubPartMgr) GetFromTablespaceID(tablespaceID int64) (results []*AllVirtualSubPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSubPart{}).Where("`tablespace_id` = ?", tablespaceID).Find(&results).Error

	return
}

// GetBatchFromTablespaceID 批量查找
func (obj *_AllVirtualSubPartMgr) GetBatchFromTablespaceID(tablespaceIDs []int64) (results []*AllVirtualSubPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSubPart{}).Where("`tablespace_id` IN (?)", tablespaceIDs).Find(&results).Error

	return
}

// GetFromSubPartIDx 通过sub_part_idx获取内容
func (obj *_AllVirtualSubPartMgr) GetFromSubPartIDx(subPartIDx int64) (results []*AllVirtualSubPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSubPart{}).Where("`sub_part_idx` = ?", subPartIDx).Find(&results).Error

	return
}

// GetBatchFromSubPartIDx 批量查找
func (obj *_AllVirtualSubPartMgr) GetBatchFromSubPartIDx(subPartIDxs []int64) (results []*AllVirtualSubPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSubPart{}).Where("`sub_part_idx` IN (?)", subPartIDxs).Find(&results).Error

	return
}

// GetFromSourcePartitionID 通过source_partition_id获取内容
func (obj *_AllVirtualSubPartMgr) GetFromSourcePartitionID(sourcePartitionID string) (results []*AllVirtualSubPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSubPart{}).Where("`source_partition_id` = ?", sourcePartitionID).Find(&results).Error

	return
}

// GetBatchFromSourcePartitionID 批量查找
func (obj *_AllVirtualSubPartMgr) GetBatchFromSourcePartitionID(sourcePartitionIDs []string) (results []*AllVirtualSubPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSubPart{}).Where("`source_partition_id` IN (?)", sourcePartitionIDs).Find(&results).Error

	return
}

// GetFromMappingPgSubPartID 通过mapping_pg_sub_part_id获取内容
func (obj *_AllVirtualSubPartMgr) GetFromMappingPgSubPartID(mappingPgSubPartID int64) (results []*AllVirtualSubPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSubPart{}).Where("`mapping_pg_sub_part_id` = ?", mappingPgSubPartID).Find(&results).Error

	return
}

// GetBatchFromMappingPgSubPartID 批量查找
func (obj *_AllVirtualSubPartMgr) GetBatchFromMappingPgSubPartID(mappingPgSubPartIDs []int64) (results []*AllVirtualSubPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSubPart{}).Where("`mapping_pg_sub_part_id` IN (?)", mappingPgSubPartIDs).Find(&results).Error

	return
}

// GetFromDropSchemaVersion 通过drop_schema_version获取内容
func (obj *_AllVirtualSubPartMgr) GetFromDropSchemaVersion(dropSchemaVersion int64) (results []*AllVirtualSubPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSubPart{}).Where("`drop_schema_version` = ?", dropSchemaVersion).Find(&results).Error

	return
}

// GetBatchFromDropSchemaVersion 批量查找
func (obj *_AllVirtualSubPartMgr) GetBatchFromDropSchemaVersion(dropSchemaVersions []int64) (results []*AllVirtualSubPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSubPart{}).Where("`drop_schema_version` IN (?)", dropSchemaVersions).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllVirtualSubPartMgr) FetchByPrimaryKey(tenantID int64, tableID int64, partID int64, subPartID int64) (result AllVirtualSubPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSubPart{}).Where("`tenant_id` = ? AND `table_id` = ? AND `part_id` = ? AND `sub_part_id` = ?", tenantID, tableID, partID, subPartID).First(&result).Error

	return
}
