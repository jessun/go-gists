package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _AllSubPartMgr struct {
	*_BaseMgr
}

// AllSubPartMgr open func
func AllSubPartMgr(db *gorm.DB) *_AllSubPartMgr {
	if db == nil {
		panic(fmt.Errorf("AllSubPartMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllSubPartMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_sub_part"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllSubPartMgr) GetTableName() string {
	return "__all_sub_part"
}

// Reset 重置gorm会话
func (obj *_AllSubPartMgr) Reset() *_AllSubPartMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllSubPartMgr) Get() (result AllSubPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSubPart{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllSubPartMgr) Gets() (results []*AllSubPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSubPart{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllSubPartMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllSubPart{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithGmtCreate gmt_create获取
func (obj *_AllSubPartMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllSubPartMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithTenantID tenant_id获取
func (obj *_AllSubPartMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithTableID table_id获取
func (obj *_AllSubPartMgr) WithTableID(tableID int64) Option {
	return optionFunc(func(o *options) { o.query["table_id"] = tableID })
}

// WithPartID part_id获取
func (obj *_AllSubPartMgr) WithPartID(partID int64) Option {
	return optionFunc(func(o *options) { o.query["part_id"] = partID })
}

// WithSubPartID sub_part_id获取
func (obj *_AllSubPartMgr) WithSubPartID(subPartID int64) Option {
	return optionFunc(func(o *options) { o.query["sub_part_id"] = subPartID })
}

// WithSubPartName sub_part_name获取
func (obj *_AllSubPartMgr) WithSubPartName(subPartName string) Option {
	return optionFunc(func(o *options) { o.query["sub_part_name"] = subPartName })
}

// WithSchemaVersion schema_version获取
func (obj *_AllSubPartMgr) WithSchemaVersion(schemaVersion int64) Option {
	return optionFunc(func(o *options) { o.query["schema_version"] = schemaVersion })
}

// WithHighBoundVal high_bound_val获取
func (obj *_AllSubPartMgr) WithHighBoundVal(highBoundVal string) Option {
	return optionFunc(func(o *options) { o.query["high_bound_val"] = highBoundVal })
}

// WithBHighBoundVal b_high_bound_val获取
func (obj *_AllSubPartMgr) WithBHighBoundVal(bHighBoundVal string) Option {
	return optionFunc(func(o *options) { o.query["b_high_bound_val"] = bHighBoundVal })
}

// WithBlockSize block_size获取
func (obj *_AllSubPartMgr) WithBlockSize(blockSize int64) Option {
	return optionFunc(func(o *options) { o.query["block_size"] = blockSize })
}

// WithReplicaNum replica_num获取
func (obj *_AllSubPartMgr) WithReplicaNum(replicaNum int64) Option {
	return optionFunc(func(o *options) { o.query["replica_num"] = replicaNum })
}

// WithCompressFuncName compress_func_name获取
func (obj *_AllSubPartMgr) WithCompressFuncName(compressFuncName string) Option {
	return optionFunc(func(o *options) { o.query["compress_func_name"] = compressFuncName })
}

// WithStatus status获取
func (obj *_AllSubPartMgr) WithStatus(status int64) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithSpare1 spare1获取
func (obj *_AllSubPartMgr) WithSpare1(spare1 int64) Option {
	return optionFunc(func(o *options) { o.query["spare1"] = spare1 })
}

// WithSpare2 spare2获取
func (obj *_AllSubPartMgr) WithSpare2(spare2 int64) Option {
	return optionFunc(func(o *options) { o.query["spare2"] = spare2 })
}

// WithSpare3 spare3获取
func (obj *_AllSubPartMgr) WithSpare3(spare3 string) Option {
	return optionFunc(func(o *options) { o.query["spare3"] = spare3 })
}

// WithComment comment获取
func (obj *_AllSubPartMgr) WithComment(comment string) Option {
	return optionFunc(func(o *options) { o.query["comment"] = comment })
}

// WithListVal list_val获取
func (obj *_AllSubPartMgr) WithListVal(listVal string) Option {
	return optionFunc(func(o *options) { o.query["list_val"] = listVal })
}

// WithBListVal b_list_val获取
func (obj *_AllSubPartMgr) WithBListVal(bListVal string) Option {
	return optionFunc(func(o *options) { o.query["b_list_val"] = bListVal })
}

// WithTablespaceID tablespace_id获取
func (obj *_AllSubPartMgr) WithTablespaceID(tablespaceID int64) Option {
	return optionFunc(func(o *options) { o.query["tablespace_id"] = tablespaceID })
}

// WithSubPartIDx sub_part_idx获取
func (obj *_AllSubPartMgr) WithSubPartIDx(subPartIDx int64) Option {
	return optionFunc(func(o *options) { o.query["sub_part_idx"] = subPartIDx })
}

// WithSourcePartitionID source_partition_id获取
func (obj *_AllSubPartMgr) WithSourcePartitionID(sourcePartitionID string) Option {
	return optionFunc(func(o *options) { o.query["source_partition_id"] = sourcePartitionID })
}

// WithMappingPgSubPartID mapping_pg_sub_part_id获取
func (obj *_AllSubPartMgr) WithMappingPgSubPartID(mappingPgSubPartID int64) Option {
	return optionFunc(func(o *options) { o.query["mapping_pg_sub_part_id"] = mappingPgSubPartID })
}

// WithDropSchemaVersion drop_schema_version获取
func (obj *_AllSubPartMgr) WithDropSchemaVersion(dropSchemaVersion int64) Option {
	return optionFunc(func(o *options) { o.query["drop_schema_version"] = dropSchemaVersion })
}

// GetByOption 功能选项模式获取
func (obj *_AllSubPartMgr) GetByOption(opts ...Option) (result AllSubPart, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllSubPart{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllSubPartMgr) GetByOptions(opts ...Option) (results []*AllSubPart, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllSubPart{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllSubPartMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllSubPart, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllSubPart{}).Where(options.query)
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
func (obj *_AllSubPartMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllSubPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSubPart{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllSubPartMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllSubPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSubPart{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllSubPartMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllSubPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSubPart{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllSubPartMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllSubPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSubPart{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromTenantID 通过tenant_id获取内容
func (obj *_AllSubPartMgr) GetFromTenantID(tenantID int64) (results []*AllSubPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSubPart{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllSubPartMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllSubPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSubPart{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromTableID 通过table_id获取内容
func (obj *_AllSubPartMgr) GetFromTableID(tableID int64) (results []*AllSubPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSubPart{}).Where("`table_id` = ?", tableID).Find(&results).Error

	return
}

// GetBatchFromTableID 批量查找
func (obj *_AllSubPartMgr) GetBatchFromTableID(tableIDs []int64) (results []*AllSubPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSubPart{}).Where("`table_id` IN (?)", tableIDs).Find(&results).Error

	return
}

// GetFromPartID 通过part_id获取内容
func (obj *_AllSubPartMgr) GetFromPartID(partID int64) (results []*AllSubPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSubPart{}).Where("`part_id` = ?", partID).Find(&results).Error

	return
}

// GetBatchFromPartID 批量查找
func (obj *_AllSubPartMgr) GetBatchFromPartID(partIDs []int64) (results []*AllSubPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSubPart{}).Where("`part_id` IN (?)", partIDs).Find(&results).Error

	return
}

// GetFromSubPartID 通过sub_part_id获取内容
func (obj *_AllSubPartMgr) GetFromSubPartID(subPartID int64) (results []*AllSubPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSubPart{}).Where("`sub_part_id` = ?", subPartID).Find(&results).Error

	return
}

// GetBatchFromSubPartID 批量查找
func (obj *_AllSubPartMgr) GetBatchFromSubPartID(subPartIDs []int64) (results []*AllSubPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSubPart{}).Where("`sub_part_id` IN (?)", subPartIDs).Find(&results).Error

	return
}

// GetFromSubPartName 通过sub_part_name获取内容
func (obj *_AllSubPartMgr) GetFromSubPartName(subPartName string) (results []*AllSubPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSubPart{}).Where("`sub_part_name` = ?", subPartName).Find(&results).Error

	return
}

// GetBatchFromSubPartName 批量查找
func (obj *_AllSubPartMgr) GetBatchFromSubPartName(subPartNames []string) (results []*AllSubPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSubPart{}).Where("`sub_part_name` IN (?)", subPartNames).Find(&results).Error

	return
}

// GetFromSchemaVersion 通过schema_version获取内容
func (obj *_AllSubPartMgr) GetFromSchemaVersion(schemaVersion int64) (results []*AllSubPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSubPart{}).Where("`schema_version` = ?", schemaVersion).Find(&results).Error

	return
}

// GetBatchFromSchemaVersion 批量查找
func (obj *_AllSubPartMgr) GetBatchFromSchemaVersion(schemaVersions []int64) (results []*AllSubPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSubPart{}).Where("`schema_version` IN (?)", schemaVersions).Find(&results).Error

	return
}

// GetFromHighBoundVal 通过high_bound_val获取内容
func (obj *_AllSubPartMgr) GetFromHighBoundVal(highBoundVal string) (results []*AllSubPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSubPart{}).Where("`high_bound_val` = ?", highBoundVal).Find(&results).Error

	return
}

// GetBatchFromHighBoundVal 批量查找
func (obj *_AllSubPartMgr) GetBatchFromHighBoundVal(highBoundVals []string) (results []*AllSubPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSubPart{}).Where("`high_bound_val` IN (?)", highBoundVals).Find(&results).Error

	return
}

// GetFromBHighBoundVal 通过b_high_bound_val获取内容
func (obj *_AllSubPartMgr) GetFromBHighBoundVal(bHighBoundVal string) (results []*AllSubPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSubPart{}).Where("`b_high_bound_val` = ?", bHighBoundVal).Find(&results).Error

	return
}

// GetBatchFromBHighBoundVal 批量查找
func (obj *_AllSubPartMgr) GetBatchFromBHighBoundVal(bHighBoundVals []string) (results []*AllSubPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSubPart{}).Where("`b_high_bound_val` IN (?)", bHighBoundVals).Find(&results).Error

	return
}

// GetFromBlockSize 通过block_size获取内容
func (obj *_AllSubPartMgr) GetFromBlockSize(blockSize int64) (results []*AllSubPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSubPart{}).Where("`block_size` = ?", blockSize).Find(&results).Error

	return
}

// GetBatchFromBlockSize 批量查找
func (obj *_AllSubPartMgr) GetBatchFromBlockSize(blockSizes []int64) (results []*AllSubPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSubPart{}).Where("`block_size` IN (?)", blockSizes).Find(&results).Error

	return
}

// GetFromReplicaNum 通过replica_num获取内容
func (obj *_AllSubPartMgr) GetFromReplicaNum(replicaNum int64) (results []*AllSubPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSubPart{}).Where("`replica_num` = ?", replicaNum).Find(&results).Error

	return
}

// GetBatchFromReplicaNum 批量查找
func (obj *_AllSubPartMgr) GetBatchFromReplicaNum(replicaNums []int64) (results []*AllSubPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSubPart{}).Where("`replica_num` IN (?)", replicaNums).Find(&results).Error

	return
}

// GetFromCompressFuncName 通过compress_func_name获取内容
func (obj *_AllSubPartMgr) GetFromCompressFuncName(compressFuncName string) (results []*AllSubPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSubPart{}).Where("`compress_func_name` = ?", compressFuncName).Find(&results).Error

	return
}

// GetBatchFromCompressFuncName 批量查找
func (obj *_AllSubPartMgr) GetBatchFromCompressFuncName(compressFuncNames []string) (results []*AllSubPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSubPart{}).Where("`compress_func_name` IN (?)", compressFuncNames).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容
func (obj *_AllSubPartMgr) GetFromStatus(status int64) (results []*AllSubPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSubPart{}).Where("`status` = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量查找
func (obj *_AllSubPartMgr) GetBatchFromStatus(statuss []int64) (results []*AllSubPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSubPart{}).Where("`status` IN (?)", statuss).Find(&results).Error

	return
}

// GetFromSpare1 通过spare1获取内容
func (obj *_AllSubPartMgr) GetFromSpare1(spare1 int64) (results []*AllSubPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSubPart{}).Where("`spare1` = ?", spare1).Find(&results).Error

	return
}

// GetBatchFromSpare1 批量查找
func (obj *_AllSubPartMgr) GetBatchFromSpare1(spare1s []int64) (results []*AllSubPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSubPart{}).Where("`spare1` IN (?)", spare1s).Find(&results).Error

	return
}

// GetFromSpare2 通过spare2获取内容
func (obj *_AllSubPartMgr) GetFromSpare2(spare2 int64) (results []*AllSubPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSubPart{}).Where("`spare2` = ?", spare2).Find(&results).Error

	return
}

// GetBatchFromSpare2 批量查找
func (obj *_AllSubPartMgr) GetBatchFromSpare2(spare2s []int64) (results []*AllSubPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSubPart{}).Where("`spare2` IN (?)", spare2s).Find(&results).Error

	return
}

// GetFromSpare3 通过spare3获取内容
func (obj *_AllSubPartMgr) GetFromSpare3(spare3 string) (results []*AllSubPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSubPart{}).Where("`spare3` = ?", spare3).Find(&results).Error

	return
}

// GetBatchFromSpare3 批量查找
func (obj *_AllSubPartMgr) GetBatchFromSpare3(spare3s []string) (results []*AllSubPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSubPart{}).Where("`spare3` IN (?)", spare3s).Find(&results).Error

	return
}

// GetFromComment 通过comment获取内容
func (obj *_AllSubPartMgr) GetFromComment(comment string) (results []*AllSubPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSubPart{}).Where("`comment` = ?", comment).Find(&results).Error

	return
}

// GetBatchFromComment 批量查找
func (obj *_AllSubPartMgr) GetBatchFromComment(comments []string) (results []*AllSubPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSubPart{}).Where("`comment` IN (?)", comments).Find(&results).Error

	return
}

// GetFromListVal 通过list_val获取内容
func (obj *_AllSubPartMgr) GetFromListVal(listVal string) (results []*AllSubPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSubPart{}).Where("`list_val` = ?", listVal).Find(&results).Error

	return
}

// GetBatchFromListVal 批量查找
func (obj *_AllSubPartMgr) GetBatchFromListVal(listVals []string) (results []*AllSubPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSubPart{}).Where("`list_val` IN (?)", listVals).Find(&results).Error

	return
}

// GetFromBListVal 通过b_list_val获取内容
func (obj *_AllSubPartMgr) GetFromBListVal(bListVal string) (results []*AllSubPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSubPart{}).Where("`b_list_val` = ?", bListVal).Find(&results).Error

	return
}

// GetBatchFromBListVal 批量查找
func (obj *_AllSubPartMgr) GetBatchFromBListVal(bListVals []string) (results []*AllSubPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSubPart{}).Where("`b_list_val` IN (?)", bListVals).Find(&results).Error

	return
}

// GetFromTablespaceID 通过tablespace_id获取内容
func (obj *_AllSubPartMgr) GetFromTablespaceID(tablespaceID int64) (results []*AllSubPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSubPart{}).Where("`tablespace_id` = ?", tablespaceID).Find(&results).Error

	return
}

// GetBatchFromTablespaceID 批量查找
func (obj *_AllSubPartMgr) GetBatchFromTablespaceID(tablespaceIDs []int64) (results []*AllSubPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSubPart{}).Where("`tablespace_id` IN (?)", tablespaceIDs).Find(&results).Error

	return
}

// GetFromSubPartIDx 通过sub_part_idx获取内容
func (obj *_AllSubPartMgr) GetFromSubPartIDx(subPartIDx int64) (results []*AllSubPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSubPart{}).Where("`sub_part_idx` = ?", subPartIDx).Find(&results).Error

	return
}

// GetBatchFromSubPartIDx 批量查找
func (obj *_AllSubPartMgr) GetBatchFromSubPartIDx(subPartIDxs []int64) (results []*AllSubPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSubPart{}).Where("`sub_part_idx` IN (?)", subPartIDxs).Find(&results).Error

	return
}

// GetFromSourcePartitionID 通过source_partition_id获取内容
func (obj *_AllSubPartMgr) GetFromSourcePartitionID(sourcePartitionID string) (results []*AllSubPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSubPart{}).Where("`source_partition_id` = ?", sourcePartitionID).Find(&results).Error

	return
}

// GetBatchFromSourcePartitionID 批量查找
func (obj *_AllSubPartMgr) GetBatchFromSourcePartitionID(sourcePartitionIDs []string) (results []*AllSubPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSubPart{}).Where("`source_partition_id` IN (?)", sourcePartitionIDs).Find(&results).Error

	return
}

// GetFromMappingPgSubPartID 通过mapping_pg_sub_part_id获取内容
func (obj *_AllSubPartMgr) GetFromMappingPgSubPartID(mappingPgSubPartID int64) (results []*AllSubPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSubPart{}).Where("`mapping_pg_sub_part_id` = ?", mappingPgSubPartID).Find(&results).Error

	return
}

// GetBatchFromMappingPgSubPartID 批量查找
func (obj *_AllSubPartMgr) GetBatchFromMappingPgSubPartID(mappingPgSubPartIDs []int64) (results []*AllSubPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSubPart{}).Where("`mapping_pg_sub_part_id` IN (?)", mappingPgSubPartIDs).Find(&results).Error

	return
}

// GetFromDropSchemaVersion 通过drop_schema_version获取内容
func (obj *_AllSubPartMgr) GetFromDropSchemaVersion(dropSchemaVersion int64) (results []*AllSubPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSubPart{}).Where("`drop_schema_version` = ?", dropSchemaVersion).Find(&results).Error

	return
}

// GetBatchFromDropSchemaVersion 批量查找
func (obj *_AllSubPartMgr) GetBatchFromDropSchemaVersion(dropSchemaVersions []int64) (results []*AllSubPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSubPart{}).Where("`drop_schema_version` IN (?)", dropSchemaVersions).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllSubPartMgr) FetchByPrimaryKey(tenantID int64, tableID int64, partID int64, subPartID int64) (result AllSubPart, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSubPart{}).Where("`tenant_id` = ? AND `table_id` = ? AND `part_id` = ? AND `sub_part_id` = ?", tenantID, tableID, partID, subPartID).First(&result).Error

	return
}
