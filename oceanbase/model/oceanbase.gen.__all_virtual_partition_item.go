package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _AllVirtualPartitionItemMgr struct {
	*_BaseMgr
}

// AllVirtualPartitionItemMgr open func
func AllVirtualPartitionItemMgr(db *gorm.DB) *_AllVirtualPartitionItemMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualPartitionItemMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualPartitionItemMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_partition_item"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualPartitionItemMgr) GetTableName() string {
	return "__all_virtual_partition_item"
}

// Reset 重置gorm会话
func (obj *_AllVirtualPartitionItemMgr) Reset() *_AllVirtualPartitionItemMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualPartitionItemMgr) Get() (result AllVirtualPartitionItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionItem{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualPartitionItemMgr) Gets() (results []*AllVirtualPartitionItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionItem{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualPartitionItemMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionItem{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTenantID tenant_id获取
func (obj *_AllVirtualPartitionItemMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithTableID table_id获取
func (obj *_AllVirtualPartitionItemMgr) WithTableID(tableID int64) Option {
	return optionFunc(func(o *options) { o.query["table_id"] = tableID })
}

// WithPartitionID partition_id获取
func (obj *_AllVirtualPartitionItemMgr) WithPartitionID(partitionID int64) Option {
	return optionFunc(func(o *options) { o.query["partition_id"] = partitionID })
}

// WithTenantName tenant_name获取
func (obj *_AllVirtualPartitionItemMgr) WithTenantName(tenantName string) Option {
	return optionFunc(func(o *options) { o.query["tenant_name"] = tenantName })
}

// WithTableName table_name获取
func (obj *_AllVirtualPartitionItemMgr) WithTableName(tableName string) Option {
	return optionFunc(func(o *options) { o.query["table_name"] = tableName })
}

// WithPartitionLevel partition_level获取
func (obj *_AllVirtualPartitionItemMgr) WithPartitionLevel(partitionLevel int64) Option {
	return optionFunc(func(o *options) { o.query["partition_level"] = partitionLevel })
}

// WithPartitionNum partition_num获取
func (obj *_AllVirtualPartitionItemMgr) WithPartitionNum(partitionNum int64) Option {
	return optionFunc(func(o *options) { o.query["partition_num"] = partitionNum })
}

// WithPartitionIDx partition_idx获取
func (obj *_AllVirtualPartitionItemMgr) WithPartitionIDx(partitionIDx int64) Option {
	return optionFunc(func(o *options) { o.query["partition_idx"] = partitionIDx })
}

// WithPartFuncType part_func_type获取
func (obj *_AllVirtualPartitionItemMgr) WithPartFuncType(partFuncType string) Option {
	return optionFunc(func(o *options) { o.query["part_func_type"] = partFuncType })
}

// WithPartNum part_num获取
func (obj *_AllVirtualPartitionItemMgr) WithPartNum(partNum int64) Option {
	return optionFunc(func(o *options) { o.query["part_num"] = partNum })
}

// WithPartName part_name获取
func (obj *_AllVirtualPartitionItemMgr) WithPartName(partName string) Option {
	return optionFunc(func(o *options) { o.query["part_name"] = partName })
}

// WithPartIDx part_idx获取
func (obj *_AllVirtualPartitionItemMgr) WithPartIDx(partIDx int64) Option {
	return optionFunc(func(o *options) { o.query["part_idx"] = partIDx })
}

// WithPartID part_id获取
func (obj *_AllVirtualPartitionItemMgr) WithPartID(partID int64) Option {
	return optionFunc(func(o *options) { o.query["part_id"] = partID })
}

// WithPartHighBound part_high_bound获取
func (obj *_AllVirtualPartitionItemMgr) WithPartHighBound(partHighBound string) Option {
	return optionFunc(func(o *options) { o.query["part_high_bound"] = partHighBound })
}

// WithSubpartFuncType subpart_func_type获取
func (obj *_AllVirtualPartitionItemMgr) WithSubpartFuncType(subpartFuncType string) Option {
	return optionFunc(func(o *options) { o.query["subpart_func_type"] = subpartFuncType })
}

// WithSubpartNum subpart_num获取
func (obj *_AllVirtualPartitionItemMgr) WithSubpartNum(subpartNum int64) Option {
	return optionFunc(func(o *options) { o.query["subpart_num"] = subpartNum })
}

// WithSubpartName subpart_name获取
func (obj *_AllVirtualPartitionItemMgr) WithSubpartName(subpartName string) Option {
	return optionFunc(func(o *options) { o.query["subpart_name"] = subpartName })
}

// WithSubpartIDx subpart_idx获取
func (obj *_AllVirtualPartitionItemMgr) WithSubpartIDx(subpartIDx int64) Option {
	return optionFunc(func(o *options) { o.query["subpart_idx"] = subpartIDx })
}

// WithSubpartID subpart_id获取
func (obj *_AllVirtualPartitionItemMgr) WithSubpartID(subpartID int64) Option {
	return optionFunc(func(o *options) { o.query["subpart_id"] = subpartID })
}

// WithSubpartHighBound subpart_high_bound获取
func (obj *_AllVirtualPartitionItemMgr) WithSubpartHighBound(subpartHighBound string) Option {
	return optionFunc(func(o *options) { o.query["subpart_high_bound"] = subpartHighBound })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualPartitionItemMgr) GetByOption(opts ...Option) (result AllVirtualPartitionItem, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionItem{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualPartitionItemMgr) GetByOptions(opts ...Option) (results []*AllVirtualPartitionItem, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionItem{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualPartitionItemMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualPartitionItem, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionItem{}).Where(options.query)
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
func (obj *_AllVirtualPartitionItemMgr) GetFromTenantID(tenantID int64) (results []*AllVirtualPartitionItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionItem{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllVirtualPartitionItemMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllVirtualPartitionItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionItem{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromTableID 通过table_id获取内容
func (obj *_AllVirtualPartitionItemMgr) GetFromTableID(tableID int64) (results []*AllVirtualPartitionItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionItem{}).Where("`table_id` = ?", tableID).Find(&results).Error

	return
}

// GetBatchFromTableID 批量查找
func (obj *_AllVirtualPartitionItemMgr) GetBatchFromTableID(tableIDs []int64) (results []*AllVirtualPartitionItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionItem{}).Where("`table_id` IN (?)", tableIDs).Find(&results).Error

	return
}

// GetFromPartitionID 通过partition_id获取内容
func (obj *_AllVirtualPartitionItemMgr) GetFromPartitionID(partitionID int64) (results []*AllVirtualPartitionItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionItem{}).Where("`partition_id` = ?", partitionID).Find(&results).Error

	return
}

// GetBatchFromPartitionID 批量查找
func (obj *_AllVirtualPartitionItemMgr) GetBatchFromPartitionID(partitionIDs []int64) (results []*AllVirtualPartitionItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionItem{}).Where("`partition_id` IN (?)", partitionIDs).Find(&results).Error

	return
}

// GetFromTenantName 通过tenant_name获取内容
func (obj *_AllVirtualPartitionItemMgr) GetFromTenantName(tenantName string) (results []*AllVirtualPartitionItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionItem{}).Where("`tenant_name` = ?", tenantName).Find(&results).Error

	return
}

// GetBatchFromTenantName 批量查找
func (obj *_AllVirtualPartitionItemMgr) GetBatchFromTenantName(tenantNames []string) (results []*AllVirtualPartitionItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionItem{}).Where("`tenant_name` IN (?)", tenantNames).Find(&results).Error

	return
}

// GetFromTableName 通过table_name获取内容
func (obj *_AllVirtualPartitionItemMgr) GetFromTableName(tableName string) (results []*AllVirtualPartitionItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionItem{}).Where("`table_name` = ?", tableName).Find(&results).Error

	return
}

// GetBatchFromTableName 批量查找
func (obj *_AllVirtualPartitionItemMgr) GetBatchFromTableName(tableNames []string) (results []*AllVirtualPartitionItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionItem{}).Where("`table_name` IN (?)", tableNames).Find(&results).Error

	return
}

// GetFromPartitionLevel 通过partition_level获取内容
func (obj *_AllVirtualPartitionItemMgr) GetFromPartitionLevel(partitionLevel int64) (results []*AllVirtualPartitionItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionItem{}).Where("`partition_level` = ?", partitionLevel).Find(&results).Error

	return
}

// GetBatchFromPartitionLevel 批量查找
func (obj *_AllVirtualPartitionItemMgr) GetBatchFromPartitionLevel(partitionLevels []int64) (results []*AllVirtualPartitionItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionItem{}).Where("`partition_level` IN (?)", partitionLevels).Find(&results).Error

	return
}

// GetFromPartitionNum 通过partition_num获取内容
func (obj *_AllVirtualPartitionItemMgr) GetFromPartitionNum(partitionNum int64) (results []*AllVirtualPartitionItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionItem{}).Where("`partition_num` = ?", partitionNum).Find(&results).Error

	return
}

// GetBatchFromPartitionNum 批量查找
func (obj *_AllVirtualPartitionItemMgr) GetBatchFromPartitionNum(partitionNums []int64) (results []*AllVirtualPartitionItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionItem{}).Where("`partition_num` IN (?)", partitionNums).Find(&results).Error

	return
}

// GetFromPartitionIDx 通过partition_idx获取内容
func (obj *_AllVirtualPartitionItemMgr) GetFromPartitionIDx(partitionIDx int64) (results []*AllVirtualPartitionItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionItem{}).Where("`partition_idx` = ?", partitionIDx).Find(&results).Error

	return
}

// GetBatchFromPartitionIDx 批量查找
func (obj *_AllVirtualPartitionItemMgr) GetBatchFromPartitionIDx(partitionIDxs []int64) (results []*AllVirtualPartitionItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionItem{}).Where("`partition_idx` IN (?)", partitionIDxs).Find(&results).Error

	return
}

// GetFromPartFuncType 通过part_func_type获取内容
func (obj *_AllVirtualPartitionItemMgr) GetFromPartFuncType(partFuncType string) (results []*AllVirtualPartitionItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionItem{}).Where("`part_func_type` = ?", partFuncType).Find(&results).Error

	return
}

// GetBatchFromPartFuncType 批量查找
func (obj *_AllVirtualPartitionItemMgr) GetBatchFromPartFuncType(partFuncTypes []string) (results []*AllVirtualPartitionItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionItem{}).Where("`part_func_type` IN (?)", partFuncTypes).Find(&results).Error

	return
}

// GetFromPartNum 通过part_num获取内容
func (obj *_AllVirtualPartitionItemMgr) GetFromPartNum(partNum int64) (results []*AllVirtualPartitionItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionItem{}).Where("`part_num` = ?", partNum).Find(&results).Error

	return
}

// GetBatchFromPartNum 批量查找
func (obj *_AllVirtualPartitionItemMgr) GetBatchFromPartNum(partNums []int64) (results []*AllVirtualPartitionItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionItem{}).Where("`part_num` IN (?)", partNums).Find(&results).Error

	return
}

// GetFromPartName 通过part_name获取内容
func (obj *_AllVirtualPartitionItemMgr) GetFromPartName(partName string) (results []*AllVirtualPartitionItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionItem{}).Where("`part_name` = ?", partName).Find(&results).Error

	return
}

// GetBatchFromPartName 批量查找
func (obj *_AllVirtualPartitionItemMgr) GetBatchFromPartName(partNames []string) (results []*AllVirtualPartitionItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionItem{}).Where("`part_name` IN (?)", partNames).Find(&results).Error

	return
}

// GetFromPartIDx 通过part_idx获取内容
func (obj *_AllVirtualPartitionItemMgr) GetFromPartIDx(partIDx int64) (results []*AllVirtualPartitionItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionItem{}).Where("`part_idx` = ?", partIDx).Find(&results).Error

	return
}

// GetBatchFromPartIDx 批量查找
func (obj *_AllVirtualPartitionItemMgr) GetBatchFromPartIDx(partIDxs []int64) (results []*AllVirtualPartitionItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionItem{}).Where("`part_idx` IN (?)", partIDxs).Find(&results).Error

	return
}

// GetFromPartID 通过part_id获取内容
func (obj *_AllVirtualPartitionItemMgr) GetFromPartID(partID int64) (results []*AllVirtualPartitionItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionItem{}).Where("`part_id` = ?", partID).Find(&results).Error

	return
}

// GetBatchFromPartID 批量查找
func (obj *_AllVirtualPartitionItemMgr) GetBatchFromPartID(partIDs []int64) (results []*AllVirtualPartitionItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionItem{}).Where("`part_id` IN (?)", partIDs).Find(&results).Error

	return
}

// GetFromPartHighBound 通过part_high_bound获取内容
func (obj *_AllVirtualPartitionItemMgr) GetFromPartHighBound(partHighBound string) (results []*AllVirtualPartitionItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionItem{}).Where("`part_high_bound` = ?", partHighBound).Find(&results).Error

	return
}

// GetBatchFromPartHighBound 批量查找
func (obj *_AllVirtualPartitionItemMgr) GetBatchFromPartHighBound(partHighBounds []string) (results []*AllVirtualPartitionItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionItem{}).Where("`part_high_bound` IN (?)", partHighBounds).Find(&results).Error

	return
}

// GetFromSubpartFuncType 通过subpart_func_type获取内容
func (obj *_AllVirtualPartitionItemMgr) GetFromSubpartFuncType(subpartFuncType string) (results []*AllVirtualPartitionItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionItem{}).Where("`subpart_func_type` = ?", subpartFuncType).Find(&results).Error

	return
}

// GetBatchFromSubpartFuncType 批量查找
func (obj *_AllVirtualPartitionItemMgr) GetBatchFromSubpartFuncType(subpartFuncTypes []string) (results []*AllVirtualPartitionItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionItem{}).Where("`subpart_func_type` IN (?)", subpartFuncTypes).Find(&results).Error

	return
}

// GetFromSubpartNum 通过subpart_num获取内容
func (obj *_AllVirtualPartitionItemMgr) GetFromSubpartNum(subpartNum int64) (results []*AllVirtualPartitionItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionItem{}).Where("`subpart_num` = ?", subpartNum).Find(&results).Error

	return
}

// GetBatchFromSubpartNum 批量查找
func (obj *_AllVirtualPartitionItemMgr) GetBatchFromSubpartNum(subpartNums []int64) (results []*AllVirtualPartitionItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionItem{}).Where("`subpart_num` IN (?)", subpartNums).Find(&results).Error

	return
}

// GetFromSubpartName 通过subpart_name获取内容
func (obj *_AllVirtualPartitionItemMgr) GetFromSubpartName(subpartName string) (results []*AllVirtualPartitionItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionItem{}).Where("`subpart_name` = ?", subpartName).Find(&results).Error

	return
}

// GetBatchFromSubpartName 批量查找
func (obj *_AllVirtualPartitionItemMgr) GetBatchFromSubpartName(subpartNames []string) (results []*AllVirtualPartitionItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionItem{}).Where("`subpart_name` IN (?)", subpartNames).Find(&results).Error

	return
}

// GetFromSubpartIDx 通过subpart_idx获取内容
func (obj *_AllVirtualPartitionItemMgr) GetFromSubpartIDx(subpartIDx int64) (results []*AllVirtualPartitionItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionItem{}).Where("`subpart_idx` = ?", subpartIDx).Find(&results).Error

	return
}

// GetBatchFromSubpartIDx 批量查找
func (obj *_AllVirtualPartitionItemMgr) GetBatchFromSubpartIDx(subpartIDxs []int64) (results []*AllVirtualPartitionItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionItem{}).Where("`subpart_idx` IN (?)", subpartIDxs).Find(&results).Error

	return
}

// GetFromSubpartID 通过subpart_id获取内容
func (obj *_AllVirtualPartitionItemMgr) GetFromSubpartID(subpartID int64) (results []*AllVirtualPartitionItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionItem{}).Where("`subpart_id` = ?", subpartID).Find(&results).Error

	return
}

// GetBatchFromSubpartID 批量查找
func (obj *_AllVirtualPartitionItemMgr) GetBatchFromSubpartID(subpartIDs []int64) (results []*AllVirtualPartitionItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionItem{}).Where("`subpart_id` IN (?)", subpartIDs).Find(&results).Error

	return
}

// GetFromSubpartHighBound 通过subpart_high_bound获取内容
func (obj *_AllVirtualPartitionItemMgr) GetFromSubpartHighBound(subpartHighBound string) (results []*AllVirtualPartitionItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionItem{}).Where("`subpart_high_bound` = ?", subpartHighBound).Find(&results).Error

	return
}

// GetBatchFromSubpartHighBound 批量查找
func (obj *_AllVirtualPartitionItemMgr) GetBatchFromSubpartHighBound(subpartHighBounds []string) (results []*AllVirtualPartitionItem, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionItem{}).Where("`subpart_high_bound` IN (?)", subpartHighBounds).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////
