package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _AllVirtualProxySubPartitionMgr struct {
	*_BaseMgr
}

// AllVirtualProxySubPartitionMgr open func
func AllVirtualProxySubPartitionMgr(db *gorm.DB) *_AllVirtualProxySubPartitionMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualProxySubPartitionMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualProxySubPartitionMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_proxy_sub_partition"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualProxySubPartitionMgr) GetTableName() string {
	return "__all_virtual_proxy_sub_partition"
}

// Reset 重置gorm会话
func (obj *_AllVirtualProxySubPartitionMgr) Reset() *_AllVirtualProxySubPartitionMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualProxySubPartitionMgr) Get() (result AllVirtualProxySubPartition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxySubPartition{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualProxySubPartitionMgr) Gets() (results []*AllVirtualProxySubPartition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxySubPartition{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualProxySubPartitionMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualProxySubPartition{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTableID table_id获取
func (obj *_AllVirtualProxySubPartitionMgr) WithTableID(tableID int64) Option {
	return optionFunc(func(o *options) { o.query["table_id"] = tableID })
}

// WithPartID part_id获取
func (obj *_AllVirtualProxySubPartitionMgr) WithPartID(partID int64) Option {
	return optionFunc(func(o *options) { o.query["part_id"] = partID })
}

// WithSubPartID sub_part_id获取
func (obj *_AllVirtualProxySubPartitionMgr) WithSubPartID(subPartID int64) Option {
	return optionFunc(func(o *options) { o.query["sub_part_id"] = subPartID })
}

// WithTenantID tenant_id获取
func (obj *_AllVirtualProxySubPartitionMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithPartName part_name获取
func (obj *_AllVirtualProxySubPartitionMgr) WithPartName(partName string) Option {
	return optionFunc(func(o *options) { o.query["part_name"] = partName })
}

// WithStatus status获取
func (obj *_AllVirtualProxySubPartitionMgr) WithStatus(status int64) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithLowBoundVal low_bound_val获取
func (obj *_AllVirtualProxySubPartitionMgr) WithLowBoundVal(lowBoundVal string) Option {
	return optionFunc(func(o *options) { o.query["low_bound_val"] = lowBoundVal })
}

// WithLowBoundValBin low_bound_val_bin获取
func (obj *_AllVirtualProxySubPartitionMgr) WithLowBoundValBin(lowBoundValBin string) Option {
	return optionFunc(func(o *options) { o.query["low_bound_val_bin"] = lowBoundValBin })
}

// WithHighBoundVal high_bound_val获取
func (obj *_AllVirtualProxySubPartitionMgr) WithHighBoundVal(highBoundVal string) Option {
	return optionFunc(func(o *options) { o.query["high_bound_val"] = highBoundVal })
}

// WithHighBoundValBin high_bound_val_bin获取
func (obj *_AllVirtualProxySubPartitionMgr) WithHighBoundValBin(highBoundValBin string) Option {
	return optionFunc(func(o *options) { o.query["high_bound_val_bin"] = highBoundValBin })
}

// WithSpare1 spare1获取
func (obj *_AllVirtualProxySubPartitionMgr) WithSpare1(spare1 int64) Option {
	return optionFunc(func(o *options) { o.query["spare1"] = spare1 })
}

// WithSpare2 spare2获取
func (obj *_AllVirtualProxySubPartitionMgr) WithSpare2(spare2 int64) Option {
	return optionFunc(func(o *options) { o.query["spare2"] = spare2 })
}

// WithSpare3 spare3获取
func (obj *_AllVirtualProxySubPartitionMgr) WithSpare3(spare3 int64) Option {
	return optionFunc(func(o *options) { o.query["spare3"] = spare3 })
}

// WithSpare4 spare4获取
func (obj *_AllVirtualProxySubPartitionMgr) WithSpare4(spare4 string) Option {
	return optionFunc(func(o *options) { o.query["spare4"] = spare4 })
}

// WithSpare5 spare5获取
func (obj *_AllVirtualProxySubPartitionMgr) WithSpare5(spare5 string) Option {
	return optionFunc(func(o *options) { o.query["spare5"] = spare5 })
}

// WithSpare6 spare6获取
func (obj *_AllVirtualProxySubPartitionMgr) WithSpare6(spare6 string) Option {
	return optionFunc(func(o *options) { o.query["spare6"] = spare6 })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualProxySubPartitionMgr) GetByOption(opts ...Option) (result AllVirtualProxySubPartition, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxySubPartition{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualProxySubPartitionMgr) GetByOptions(opts ...Option) (results []*AllVirtualProxySubPartition, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxySubPartition{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualProxySubPartitionMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualProxySubPartition, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualProxySubPartition{}).Where(options.query)
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

// GetFromTableID 通过table_id获取内容
func (obj *_AllVirtualProxySubPartitionMgr) GetFromTableID(tableID int64) (results []*AllVirtualProxySubPartition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxySubPartition{}).Where("`table_id` = ?", tableID).Find(&results).Error

	return
}

// GetBatchFromTableID 批量查找
func (obj *_AllVirtualProxySubPartitionMgr) GetBatchFromTableID(tableIDs []int64) (results []*AllVirtualProxySubPartition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxySubPartition{}).Where("`table_id` IN (?)", tableIDs).Find(&results).Error

	return
}

// GetFromPartID 通过part_id获取内容
func (obj *_AllVirtualProxySubPartitionMgr) GetFromPartID(partID int64) (results []*AllVirtualProxySubPartition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxySubPartition{}).Where("`part_id` = ?", partID).Find(&results).Error

	return
}

// GetBatchFromPartID 批量查找
func (obj *_AllVirtualProxySubPartitionMgr) GetBatchFromPartID(partIDs []int64) (results []*AllVirtualProxySubPartition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxySubPartition{}).Where("`part_id` IN (?)", partIDs).Find(&results).Error

	return
}

// GetFromSubPartID 通过sub_part_id获取内容
func (obj *_AllVirtualProxySubPartitionMgr) GetFromSubPartID(subPartID int64) (results []*AllVirtualProxySubPartition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxySubPartition{}).Where("`sub_part_id` = ?", subPartID).Find(&results).Error

	return
}

// GetBatchFromSubPartID 批量查找
func (obj *_AllVirtualProxySubPartitionMgr) GetBatchFromSubPartID(subPartIDs []int64) (results []*AllVirtualProxySubPartition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxySubPartition{}).Where("`sub_part_id` IN (?)", subPartIDs).Find(&results).Error

	return
}

// GetFromTenantID 通过tenant_id获取内容
func (obj *_AllVirtualProxySubPartitionMgr) GetFromTenantID(tenantID int64) (results []*AllVirtualProxySubPartition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxySubPartition{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllVirtualProxySubPartitionMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllVirtualProxySubPartition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxySubPartition{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromPartName 通过part_name获取内容
func (obj *_AllVirtualProxySubPartitionMgr) GetFromPartName(partName string) (results []*AllVirtualProxySubPartition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxySubPartition{}).Where("`part_name` = ?", partName).Find(&results).Error

	return
}

// GetBatchFromPartName 批量查找
func (obj *_AllVirtualProxySubPartitionMgr) GetBatchFromPartName(partNames []string) (results []*AllVirtualProxySubPartition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxySubPartition{}).Where("`part_name` IN (?)", partNames).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容
func (obj *_AllVirtualProxySubPartitionMgr) GetFromStatus(status int64) (results []*AllVirtualProxySubPartition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxySubPartition{}).Where("`status` = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量查找
func (obj *_AllVirtualProxySubPartitionMgr) GetBatchFromStatus(statuss []int64) (results []*AllVirtualProxySubPartition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxySubPartition{}).Where("`status` IN (?)", statuss).Find(&results).Error

	return
}

// GetFromLowBoundVal 通过low_bound_val获取内容
func (obj *_AllVirtualProxySubPartitionMgr) GetFromLowBoundVal(lowBoundVal string) (results []*AllVirtualProxySubPartition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxySubPartition{}).Where("`low_bound_val` = ?", lowBoundVal).Find(&results).Error

	return
}

// GetBatchFromLowBoundVal 批量查找
func (obj *_AllVirtualProxySubPartitionMgr) GetBatchFromLowBoundVal(lowBoundVals []string) (results []*AllVirtualProxySubPartition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxySubPartition{}).Where("`low_bound_val` IN (?)", lowBoundVals).Find(&results).Error

	return
}

// GetFromLowBoundValBin 通过low_bound_val_bin获取内容
func (obj *_AllVirtualProxySubPartitionMgr) GetFromLowBoundValBin(lowBoundValBin string) (results []*AllVirtualProxySubPartition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxySubPartition{}).Where("`low_bound_val_bin` = ?", lowBoundValBin).Find(&results).Error

	return
}

// GetBatchFromLowBoundValBin 批量查找
func (obj *_AllVirtualProxySubPartitionMgr) GetBatchFromLowBoundValBin(lowBoundValBins []string) (results []*AllVirtualProxySubPartition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxySubPartition{}).Where("`low_bound_val_bin` IN (?)", lowBoundValBins).Find(&results).Error

	return
}

// GetFromHighBoundVal 通过high_bound_val获取内容
func (obj *_AllVirtualProxySubPartitionMgr) GetFromHighBoundVal(highBoundVal string) (results []*AllVirtualProxySubPartition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxySubPartition{}).Where("`high_bound_val` = ?", highBoundVal).Find(&results).Error

	return
}

// GetBatchFromHighBoundVal 批量查找
func (obj *_AllVirtualProxySubPartitionMgr) GetBatchFromHighBoundVal(highBoundVals []string) (results []*AllVirtualProxySubPartition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxySubPartition{}).Where("`high_bound_val` IN (?)", highBoundVals).Find(&results).Error

	return
}

// GetFromHighBoundValBin 通过high_bound_val_bin获取内容
func (obj *_AllVirtualProxySubPartitionMgr) GetFromHighBoundValBin(highBoundValBin string) (results []*AllVirtualProxySubPartition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxySubPartition{}).Where("`high_bound_val_bin` = ?", highBoundValBin).Find(&results).Error

	return
}

// GetBatchFromHighBoundValBin 批量查找
func (obj *_AllVirtualProxySubPartitionMgr) GetBatchFromHighBoundValBin(highBoundValBins []string) (results []*AllVirtualProxySubPartition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxySubPartition{}).Where("`high_bound_val_bin` IN (?)", highBoundValBins).Find(&results).Error

	return
}

// GetFromSpare1 通过spare1获取内容
func (obj *_AllVirtualProxySubPartitionMgr) GetFromSpare1(spare1 int64) (results []*AllVirtualProxySubPartition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxySubPartition{}).Where("`spare1` = ?", spare1).Find(&results).Error

	return
}

// GetBatchFromSpare1 批量查找
func (obj *_AllVirtualProxySubPartitionMgr) GetBatchFromSpare1(spare1s []int64) (results []*AllVirtualProxySubPartition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxySubPartition{}).Where("`spare1` IN (?)", spare1s).Find(&results).Error

	return
}

// GetFromSpare2 通过spare2获取内容
func (obj *_AllVirtualProxySubPartitionMgr) GetFromSpare2(spare2 int64) (results []*AllVirtualProxySubPartition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxySubPartition{}).Where("`spare2` = ?", spare2).Find(&results).Error

	return
}

// GetBatchFromSpare2 批量查找
func (obj *_AllVirtualProxySubPartitionMgr) GetBatchFromSpare2(spare2s []int64) (results []*AllVirtualProxySubPartition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxySubPartition{}).Where("`spare2` IN (?)", spare2s).Find(&results).Error

	return
}

// GetFromSpare3 通过spare3获取内容
func (obj *_AllVirtualProxySubPartitionMgr) GetFromSpare3(spare3 int64) (results []*AllVirtualProxySubPartition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxySubPartition{}).Where("`spare3` = ?", spare3).Find(&results).Error

	return
}

// GetBatchFromSpare3 批量查找
func (obj *_AllVirtualProxySubPartitionMgr) GetBatchFromSpare3(spare3s []int64) (results []*AllVirtualProxySubPartition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxySubPartition{}).Where("`spare3` IN (?)", spare3s).Find(&results).Error

	return
}

// GetFromSpare4 通过spare4获取内容
func (obj *_AllVirtualProxySubPartitionMgr) GetFromSpare4(spare4 string) (results []*AllVirtualProxySubPartition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxySubPartition{}).Where("`spare4` = ?", spare4).Find(&results).Error

	return
}

// GetBatchFromSpare4 批量查找
func (obj *_AllVirtualProxySubPartitionMgr) GetBatchFromSpare4(spare4s []string) (results []*AllVirtualProxySubPartition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxySubPartition{}).Where("`spare4` IN (?)", spare4s).Find(&results).Error

	return
}

// GetFromSpare5 通过spare5获取内容
func (obj *_AllVirtualProxySubPartitionMgr) GetFromSpare5(spare5 string) (results []*AllVirtualProxySubPartition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxySubPartition{}).Where("`spare5` = ?", spare5).Find(&results).Error

	return
}

// GetBatchFromSpare5 批量查找
func (obj *_AllVirtualProxySubPartitionMgr) GetBatchFromSpare5(spare5s []string) (results []*AllVirtualProxySubPartition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxySubPartition{}).Where("`spare5` IN (?)", spare5s).Find(&results).Error

	return
}

// GetFromSpare6 通过spare6获取内容
func (obj *_AllVirtualProxySubPartitionMgr) GetFromSpare6(spare6 string) (results []*AllVirtualProxySubPartition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxySubPartition{}).Where("`spare6` = ?", spare6).Find(&results).Error

	return
}

// GetBatchFromSpare6 批量查找
func (obj *_AllVirtualProxySubPartitionMgr) GetBatchFromSpare6(spare6s []string) (results []*AllVirtualProxySubPartition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxySubPartition{}).Where("`spare6` IN (?)", spare6s).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllVirtualProxySubPartitionMgr) FetchByPrimaryKey(tableID int64, partID int64, subPartID int64) (result AllVirtualProxySubPartition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxySubPartition{}).Where("`table_id` = ? AND `part_id` = ? AND `sub_part_id` = ?", tableID, partID, subPartID).First(&result).Error

	return
}
