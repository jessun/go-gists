package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _AllVirtualProxyPartitionMgr struct {
	*_BaseMgr
}

// AllVirtualProxyPartitionMgr open func
func AllVirtualProxyPartitionMgr(db *gorm.DB) *_AllVirtualProxyPartitionMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualProxyPartitionMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualProxyPartitionMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_proxy_partition"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualProxyPartitionMgr) GetTableName() string {
	return "__all_virtual_proxy_partition"
}

// Reset 重置gorm会话
func (obj *_AllVirtualProxyPartitionMgr) Reset() *_AllVirtualProxyPartitionMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualProxyPartitionMgr) Get() (result AllVirtualProxyPartition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxyPartition{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualProxyPartitionMgr) Gets() (results []*AllVirtualProxyPartition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxyPartition{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualProxyPartitionMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualProxyPartition{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTableID table_id获取
func (obj *_AllVirtualProxyPartitionMgr) WithTableID(tableID int64) Option {
	return optionFunc(func(o *options) { o.query["table_id"] = tableID })
}

// WithPartID part_id获取
func (obj *_AllVirtualProxyPartitionMgr) WithPartID(partID int64) Option {
	return optionFunc(func(o *options) { o.query["part_id"] = partID })
}

// WithTenantID tenant_id获取
func (obj *_AllVirtualProxyPartitionMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithPartName part_name获取
func (obj *_AllVirtualProxyPartitionMgr) WithPartName(partName string) Option {
	return optionFunc(func(o *options) { o.query["part_name"] = partName })
}

// WithStatus status获取
func (obj *_AllVirtualProxyPartitionMgr) WithStatus(status int64) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithLowBoundVal low_bound_val获取
func (obj *_AllVirtualProxyPartitionMgr) WithLowBoundVal(lowBoundVal string) Option {
	return optionFunc(func(o *options) { o.query["low_bound_val"] = lowBoundVal })
}

// WithLowBoundValBin low_bound_val_bin获取
func (obj *_AllVirtualProxyPartitionMgr) WithLowBoundValBin(lowBoundValBin string) Option {
	return optionFunc(func(o *options) { o.query["low_bound_val_bin"] = lowBoundValBin })
}

// WithHighBoundVal high_bound_val获取
func (obj *_AllVirtualProxyPartitionMgr) WithHighBoundVal(highBoundVal string) Option {
	return optionFunc(func(o *options) { o.query["high_bound_val"] = highBoundVal })
}

// WithHighBoundValBin high_bound_val_bin获取
func (obj *_AllVirtualProxyPartitionMgr) WithHighBoundValBin(highBoundValBin string) Option {
	return optionFunc(func(o *options) { o.query["high_bound_val_bin"] = highBoundValBin })
}

// WithPartIDx part_idx获取
func (obj *_AllVirtualProxyPartitionMgr) WithPartIDx(partIDx int64) Option {
	return optionFunc(func(o *options) { o.query["part_idx"] = partIDx })
}

// WithSubPartNum sub_part_num获取
func (obj *_AllVirtualProxyPartitionMgr) WithSubPartNum(subPartNum int64) Option {
	return optionFunc(func(o *options) { o.query["sub_part_num"] = subPartNum })
}

// WithSubPartSpace sub_part_space获取
func (obj *_AllVirtualProxyPartitionMgr) WithSubPartSpace(subPartSpace int64) Option {
	return optionFunc(func(o *options) { o.query["sub_part_space"] = subPartSpace })
}

// WithSubPartInterval sub_part_interval获取
func (obj *_AllVirtualProxyPartitionMgr) WithSubPartInterval(subPartInterval string) Option {
	return optionFunc(func(o *options) { o.query["sub_part_interval"] = subPartInterval })
}

// WithSubPartIntervalBin sub_part_interval_bin获取
func (obj *_AllVirtualProxyPartitionMgr) WithSubPartIntervalBin(subPartIntervalBin string) Option {
	return optionFunc(func(o *options) { o.query["sub_part_interval_bin"] = subPartIntervalBin })
}

// WithSubIntervalStart sub_interval_start获取
func (obj *_AllVirtualProxyPartitionMgr) WithSubIntervalStart(subIntervalStart string) Option {
	return optionFunc(func(o *options) { o.query["sub_interval_start"] = subIntervalStart })
}

// WithSubIntervalStartBin sub_interval_start_bin获取
func (obj *_AllVirtualProxyPartitionMgr) WithSubIntervalStartBin(subIntervalStartBin string) Option {
	return optionFunc(func(o *options) { o.query["sub_interval_start_bin"] = subIntervalStartBin })
}

// WithSpare1 spare1获取
func (obj *_AllVirtualProxyPartitionMgr) WithSpare1(spare1 int64) Option {
	return optionFunc(func(o *options) { o.query["spare1"] = spare1 })
}

// WithSpare2 spare2获取
func (obj *_AllVirtualProxyPartitionMgr) WithSpare2(spare2 int64) Option {
	return optionFunc(func(o *options) { o.query["spare2"] = spare2 })
}

// WithSpare3 spare3获取
func (obj *_AllVirtualProxyPartitionMgr) WithSpare3(spare3 int64) Option {
	return optionFunc(func(o *options) { o.query["spare3"] = spare3 })
}

// WithSpare4 spare4获取
func (obj *_AllVirtualProxyPartitionMgr) WithSpare4(spare4 string) Option {
	return optionFunc(func(o *options) { o.query["spare4"] = spare4 })
}

// WithSpare5 spare5获取
func (obj *_AllVirtualProxyPartitionMgr) WithSpare5(spare5 string) Option {
	return optionFunc(func(o *options) { o.query["spare5"] = spare5 })
}

// WithSpare6 spare6获取
func (obj *_AllVirtualProxyPartitionMgr) WithSpare6(spare6 string) Option {
	return optionFunc(func(o *options) { o.query["spare6"] = spare6 })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualProxyPartitionMgr) GetByOption(opts ...Option) (result AllVirtualProxyPartition, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxyPartition{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualProxyPartitionMgr) GetByOptions(opts ...Option) (results []*AllVirtualProxyPartition, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxyPartition{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualProxyPartitionMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualProxyPartition, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualProxyPartition{}).Where(options.query)
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
func (obj *_AllVirtualProxyPartitionMgr) GetFromTableID(tableID int64) (results []*AllVirtualProxyPartition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxyPartition{}).Where("`table_id` = ?", tableID).Find(&results).Error

	return
}

// GetBatchFromTableID 批量查找
func (obj *_AllVirtualProxyPartitionMgr) GetBatchFromTableID(tableIDs []int64) (results []*AllVirtualProxyPartition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxyPartition{}).Where("`table_id` IN (?)", tableIDs).Find(&results).Error

	return
}

// GetFromPartID 通过part_id获取内容
func (obj *_AllVirtualProxyPartitionMgr) GetFromPartID(partID int64) (results []*AllVirtualProxyPartition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxyPartition{}).Where("`part_id` = ?", partID).Find(&results).Error

	return
}

// GetBatchFromPartID 批量查找
func (obj *_AllVirtualProxyPartitionMgr) GetBatchFromPartID(partIDs []int64) (results []*AllVirtualProxyPartition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxyPartition{}).Where("`part_id` IN (?)", partIDs).Find(&results).Error

	return
}

// GetFromTenantID 通过tenant_id获取内容
func (obj *_AllVirtualProxyPartitionMgr) GetFromTenantID(tenantID int64) (results []*AllVirtualProxyPartition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxyPartition{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllVirtualProxyPartitionMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllVirtualProxyPartition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxyPartition{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromPartName 通过part_name获取内容
func (obj *_AllVirtualProxyPartitionMgr) GetFromPartName(partName string) (results []*AllVirtualProxyPartition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxyPartition{}).Where("`part_name` = ?", partName).Find(&results).Error

	return
}

// GetBatchFromPartName 批量查找
func (obj *_AllVirtualProxyPartitionMgr) GetBatchFromPartName(partNames []string) (results []*AllVirtualProxyPartition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxyPartition{}).Where("`part_name` IN (?)", partNames).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容
func (obj *_AllVirtualProxyPartitionMgr) GetFromStatus(status int64) (results []*AllVirtualProxyPartition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxyPartition{}).Where("`status` = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量查找
func (obj *_AllVirtualProxyPartitionMgr) GetBatchFromStatus(statuss []int64) (results []*AllVirtualProxyPartition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxyPartition{}).Where("`status` IN (?)", statuss).Find(&results).Error

	return
}

// GetFromLowBoundVal 通过low_bound_val获取内容
func (obj *_AllVirtualProxyPartitionMgr) GetFromLowBoundVal(lowBoundVal string) (results []*AllVirtualProxyPartition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxyPartition{}).Where("`low_bound_val` = ?", lowBoundVal).Find(&results).Error

	return
}

// GetBatchFromLowBoundVal 批量查找
func (obj *_AllVirtualProxyPartitionMgr) GetBatchFromLowBoundVal(lowBoundVals []string) (results []*AllVirtualProxyPartition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxyPartition{}).Where("`low_bound_val` IN (?)", lowBoundVals).Find(&results).Error

	return
}

// GetFromLowBoundValBin 通过low_bound_val_bin获取内容
func (obj *_AllVirtualProxyPartitionMgr) GetFromLowBoundValBin(lowBoundValBin string) (results []*AllVirtualProxyPartition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxyPartition{}).Where("`low_bound_val_bin` = ?", lowBoundValBin).Find(&results).Error

	return
}

// GetBatchFromLowBoundValBin 批量查找
func (obj *_AllVirtualProxyPartitionMgr) GetBatchFromLowBoundValBin(lowBoundValBins []string) (results []*AllVirtualProxyPartition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxyPartition{}).Where("`low_bound_val_bin` IN (?)", lowBoundValBins).Find(&results).Error

	return
}

// GetFromHighBoundVal 通过high_bound_val获取内容
func (obj *_AllVirtualProxyPartitionMgr) GetFromHighBoundVal(highBoundVal string) (results []*AllVirtualProxyPartition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxyPartition{}).Where("`high_bound_val` = ?", highBoundVal).Find(&results).Error

	return
}

// GetBatchFromHighBoundVal 批量查找
func (obj *_AllVirtualProxyPartitionMgr) GetBatchFromHighBoundVal(highBoundVals []string) (results []*AllVirtualProxyPartition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxyPartition{}).Where("`high_bound_val` IN (?)", highBoundVals).Find(&results).Error

	return
}

// GetFromHighBoundValBin 通过high_bound_val_bin获取内容
func (obj *_AllVirtualProxyPartitionMgr) GetFromHighBoundValBin(highBoundValBin string) (results []*AllVirtualProxyPartition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxyPartition{}).Where("`high_bound_val_bin` = ?", highBoundValBin).Find(&results).Error

	return
}

// GetBatchFromHighBoundValBin 批量查找
func (obj *_AllVirtualProxyPartitionMgr) GetBatchFromHighBoundValBin(highBoundValBins []string) (results []*AllVirtualProxyPartition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxyPartition{}).Where("`high_bound_val_bin` IN (?)", highBoundValBins).Find(&results).Error

	return
}

// GetFromPartIDx 通过part_idx获取内容
func (obj *_AllVirtualProxyPartitionMgr) GetFromPartIDx(partIDx int64) (results []*AllVirtualProxyPartition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxyPartition{}).Where("`part_idx` = ?", partIDx).Find(&results).Error

	return
}

// GetBatchFromPartIDx 批量查找
func (obj *_AllVirtualProxyPartitionMgr) GetBatchFromPartIDx(partIDxs []int64) (results []*AllVirtualProxyPartition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxyPartition{}).Where("`part_idx` IN (?)", partIDxs).Find(&results).Error

	return
}

// GetFromSubPartNum 通过sub_part_num获取内容
func (obj *_AllVirtualProxyPartitionMgr) GetFromSubPartNum(subPartNum int64) (results []*AllVirtualProxyPartition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxyPartition{}).Where("`sub_part_num` = ?", subPartNum).Find(&results).Error

	return
}

// GetBatchFromSubPartNum 批量查找
func (obj *_AllVirtualProxyPartitionMgr) GetBatchFromSubPartNum(subPartNums []int64) (results []*AllVirtualProxyPartition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxyPartition{}).Where("`sub_part_num` IN (?)", subPartNums).Find(&results).Error

	return
}

// GetFromSubPartSpace 通过sub_part_space获取内容
func (obj *_AllVirtualProxyPartitionMgr) GetFromSubPartSpace(subPartSpace int64) (results []*AllVirtualProxyPartition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxyPartition{}).Where("`sub_part_space` = ?", subPartSpace).Find(&results).Error

	return
}

// GetBatchFromSubPartSpace 批量查找
func (obj *_AllVirtualProxyPartitionMgr) GetBatchFromSubPartSpace(subPartSpaces []int64) (results []*AllVirtualProxyPartition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxyPartition{}).Where("`sub_part_space` IN (?)", subPartSpaces).Find(&results).Error

	return
}

// GetFromSubPartInterval 通过sub_part_interval获取内容
func (obj *_AllVirtualProxyPartitionMgr) GetFromSubPartInterval(subPartInterval string) (results []*AllVirtualProxyPartition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxyPartition{}).Where("`sub_part_interval` = ?", subPartInterval).Find(&results).Error

	return
}

// GetBatchFromSubPartInterval 批量查找
func (obj *_AllVirtualProxyPartitionMgr) GetBatchFromSubPartInterval(subPartIntervals []string) (results []*AllVirtualProxyPartition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxyPartition{}).Where("`sub_part_interval` IN (?)", subPartIntervals).Find(&results).Error

	return
}

// GetFromSubPartIntervalBin 通过sub_part_interval_bin获取内容
func (obj *_AllVirtualProxyPartitionMgr) GetFromSubPartIntervalBin(subPartIntervalBin string) (results []*AllVirtualProxyPartition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxyPartition{}).Where("`sub_part_interval_bin` = ?", subPartIntervalBin).Find(&results).Error

	return
}

// GetBatchFromSubPartIntervalBin 批量查找
func (obj *_AllVirtualProxyPartitionMgr) GetBatchFromSubPartIntervalBin(subPartIntervalBins []string) (results []*AllVirtualProxyPartition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxyPartition{}).Where("`sub_part_interval_bin` IN (?)", subPartIntervalBins).Find(&results).Error

	return
}

// GetFromSubIntervalStart 通过sub_interval_start获取内容
func (obj *_AllVirtualProxyPartitionMgr) GetFromSubIntervalStart(subIntervalStart string) (results []*AllVirtualProxyPartition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxyPartition{}).Where("`sub_interval_start` = ?", subIntervalStart).Find(&results).Error

	return
}

// GetBatchFromSubIntervalStart 批量查找
func (obj *_AllVirtualProxyPartitionMgr) GetBatchFromSubIntervalStart(subIntervalStarts []string) (results []*AllVirtualProxyPartition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxyPartition{}).Where("`sub_interval_start` IN (?)", subIntervalStarts).Find(&results).Error

	return
}

// GetFromSubIntervalStartBin 通过sub_interval_start_bin获取内容
func (obj *_AllVirtualProxyPartitionMgr) GetFromSubIntervalStartBin(subIntervalStartBin string) (results []*AllVirtualProxyPartition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxyPartition{}).Where("`sub_interval_start_bin` = ?", subIntervalStartBin).Find(&results).Error

	return
}

// GetBatchFromSubIntervalStartBin 批量查找
func (obj *_AllVirtualProxyPartitionMgr) GetBatchFromSubIntervalStartBin(subIntervalStartBins []string) (results []*AllVirtualProxyPartition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxyPartition{}).Where("`sub_interval_start_bin` IN (?)", subIntervalStartBins).Find(&results).Error

	return
}

// GetFromSpare1 通过spare1获取内容
func (obj *_AllVirtualProxyPartitionMgr) GetFromSpare1(spare1 int64) (results []*AllVirtualProxyPartition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxyPartition{}).Where("`spare1` = ?", spare1).Find(&results).Error

	return
}

// GetBatchFromSpare1 批量查找
func (obj *_AllVirtualProxyPartitionMgr) GetBatchFromSpare1(spare1s []int64) (results []*AllVirtualProxyPartition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxyPartition{}).Where("`spare1` IN (?)", spare1s).Find(&results).Error

	return
}

// GetFromSpare2 通过spare2获取内容
func (obj *_AllVirtualProxyPartitionMgr) GetFromSpare2(spare2 int64) (results []*AllVirtualProxyPartition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxyPartition{}).Where("`spare2` = ?", spare2).Find(&results).Error

	return
}

// GetBatchFromSpare2 批量查找
func (obj *_AllVirtualProxyPartitionMgr) GetBatchFromSpare2(spare2s []int64) (results []*AllVirtualProxyPartition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxyPartition{}).Where("`spare2` IN (?)", spare2s).Find(&results).Error

	return
}

// GetFromSpare3 通过spare3获取内容
func (obj *_AllVirtualProxyPartitionMgr) GetFromSpare3(spare3 int64) (results []*AllVirtualProxyPartition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxyPartition{}).Where("`spare3` = ?", spare3).Find(&results).Error

	return
}

// GetBatchFromSpare3 批量查找
func (obj *_AllVirtualProxyPartitionMgr) GetBatchFromSpare3(spare3s []int64) (results []*AllVirtualProxyPartition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxyPartition{}).Where("`spare3` IN (?)", spare3s).Find(&results).Error

	return
}

// GetFromSpare4 通过spare4获取内容
func (obj *_AllVirtualProxyPartitionMgr) GetFromSpare4(spare4 string) (results []*AllVirtualProxyPartition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxyPartition{}).Where("`spare4` = ?", spare4).Find(&results).Error

	return
}

// GetBatchFromSpare4 批量查找
func (obj *_AllVirtualProxyPartitionMgr) GetBatchFromSpare4(spare4s []string) (results []*AllVirtualProxyPartition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxyPartition{}).Where("`spare4` IN (?)", spare4s).Find(&results).Error

	return
}

// GetFromSpare5 通过spare5获取内容
func (obj *_AllVirtualProxyPartitionMgr) GetFromSpare5(spare5 string) (results []*AllVirtualProxyPartition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxyPartition{}).Where("`spare5` = ?", spare5).Find(&results).Error

	return
}

// GetBatchFromSpare5 批量查找
func (obj *_AllVirtualProxyPartitionMgr) GetBatchFromSpare5(spare5s []string) (results []*AllVirtualProxyPartition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxyPartition{}).Where("`spare5` IN (?)", spare5s).Find(&results).Error

	return
}

// GetFromSpare6 通过spare6获取内容
func (obj *_AllVirtualProxyPartitionMgr) GetFromSpare6(spare6 string) (results []*AllVirtualProxyPartition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxyPartition{}).Where("`spare6` = ?", spare6).Find(&results).Error

	return
}

// GetBatchFromSpare6 批量查找
func (obj *_AllVirtualProxyPartitionMgr) GetBatchFromSpare6(spare6s []string) (results []*AllVirtualProxyPartition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxyPartition{}).Where("`spare6` IN (?)", spare6s).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllVirtualProxyPartitionMgr) FetchByPrimaryKey(tableID int64, partID int64) (result AllVirtualProxyPartition, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxyPartition{}).Where("`table_id` = ? AND `part_id` = ?", tableID, partID).First(&result).Error

	return
}
