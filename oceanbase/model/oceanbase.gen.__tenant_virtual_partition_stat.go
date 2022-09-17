package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _TenantVirtualPartitionStatMgr struct {
	*_BaseMgr
}

// TenantVirtualPartitionStatMgr open func
func TenantVirtualPartitionStatMgr(db *gorm.DB) *_TenantVirtualPartitionStatMgr {
	if db == nil {
		panic(fmt.Errorf("TenantVirtualPartitionStatMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_TenantVirtualPartitionStatMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__tenant_virtual_partition_stat"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_TenantVirtualPartitionStatMgr) GetTableName() string {
	return "__tenant_virtual_partition_stat"
}

// Reset 重置gorm会话
func (obj *_TenantVirtualPartitionStatMgr) Reset() *_TenantVirtualPartitionStatMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_TenantVirtualPartitionStatMgr) Get() (result TenantVirtualPartitionStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualPartitionStat{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_TenantVirtualPartitionStatMgr) Gets() (results []*TenantVirtualPartitionStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualPartitionStat{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_TenantVirtualPartitionStatMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(TenantVirtualPartitionStat{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTableID table_id获取
func (obj *_TenantVirtualPartitionStatMgr) WithTableID(tableID int64) Option {
	return optionFunc(func(o *options) { o.query["table_id"] = tableID })
}

// WithPartitionID partition_id获取
func (obj *_TenantVirtualPartitionStatMgr) WithPartitionID(partitionID int64) Option {
	return optionFunc(func(o *options) { o.query["partition_id"] = partitionID })
}

// WithPartitionCnt partition_cnt获取
func (obj *_TenantVirtualPartitionStatMgr) WithPartitionCnt(partitionCnt int64) Option {
	return optionFunc(func(o *options) { o.query["partition_cnt"] = partitionCnt })
}

// WithRowCount row_count获取
func (obj *_TenantVirtualPartitionStatMgr) WithRowCount(rowCount int64) Option {
	return optionFunc(func(o *options) { o.query["row_count"] = rowCount })
}

// WithDiffPercentage diff_percentage获取
func (obj *_TenantVirtualPartitionStatMgr) WithDiffPercentage(diffPercentage int64) Option {
	return optionFunc(func(o *options) { o.query["diff_percentage"] = diffPercentage })
}

// GetByOption 功能选项模式获取
func (obj *_TenantVirtualPartitionStatMgr) GetByOption(opts ...Option) (result TenantVirtualPartitionStat, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualPartitionStat{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_TenantVirtualPartitionStatMgr) GetByOptions(opts ...Option) (results []*TenantVirtualPartitionStat, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualPartitionStat{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_TenantVirtualPartitionStatMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]TenantVirtualPartitionStat, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(TenantVirtualPartitionStat{}).Where(options.query)
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
func (obj *_TenantVirtualPartitionStatMgr) GetFromTableID(tableID int64) (results []*TenantVirtualPartitionStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualPartitionStat{}).Where("`table_id` = ?", tableID).Find(&results).Error

	return
}

// GetBatchFromTableID 批量查找
func (obj *_TenantVirtualPartitionStatMgr) GetBatchFromTableID(tableIDs []int64) (results []*TenantVirtualPartitionStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualPartitionStat{}).Where("`table_id` IN (?)", tableIDs).Find(&results).Error

	return
}

// GetFromPartitionID 通过partition_id获取内容
func (obj *_TenantVirtualPartitionStatMgr) GetFromPartitionID(partitionID int64) (results []*TenantVirtualPartitionStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualPartitionStat{}).Where("`partition_id` = ?", partitionID).Find(&results).Error

	return
}

// GetBatchFromPartitionID 批量查找
func (obj *_TenantVirtualPartitionStatMgr) GetBatchFromPartitionID(partitionIDs []int64) (results []*TenantVirtualPartitionStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualPartitionStat{}).Where("`partition_id` IN (?)", partitionIDs).Find(&results).Error

	return
}

// GetFromPartitionCnt 通过partition_cnt获取内容
func (obj *_TenantVirtualPartitionStatMgr) GetFromPartitionCnt(partitionCnt int64) (results []*TenantVirtualPartitionStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualPartitionStat{}).Where("`partition_cnt` = ?", partitionCnt).Find(&results).Error

	return
}

// GetBatchFromPartitionCnt 批量查找
func (obj *_TenantVirtualPartitionStatMgr) GetBatchFromPartitionCnt(partitionCnts []int64) (results []*TenantVirtualPartitionStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualPartitionStat{}).Where("`partition_cnt` IN (?)", partitionCnts).Find(&results).Error

	return
}

// GetFromRowCount 通过row_count获取内容
func (obj *_TenantVirtualPartitionStatMgr) GetFromRowCount(rowCount int64) (results []*TenantVirtualPartitionStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualPartitionStat{}).Where("`row_count` = ?", rowCount).Find(&results).Error

	return
}

// GetBatchFromRowCount 批量查找
func (obj *_TenantVirtualPartitionStatMgr) GetBatchFromRowCount(rowCounts []int64) (results []*TenantVirtualPartitionStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualPartitionStat{}).Where("`row_count` IN (?)", rowCounts).Find(&results).Error

	return
}

// GetFromDiffPercentage 通过diff_percentage获取内容
func (obj *_TenantVirtualPartitionStatMgr) GetFromDiffPercentage(diffPercentage int64) (results []*TenantVirtualPartitionStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualPartitionStat{}).Where("`diff_percentage` = ?", diffPercentage).Find(&results).Error

	return
}

// GetBatchFromDiffPercentage 批量查找
func (obj *_TenantVirtualPartitionStatMgr) GetBatchFromDiffPercentage(diffPercentages []int64) (results []*TenantVirtualPartitionStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualPartitionStat{}).Where("`diff_percentage` IN (?)", diffPercentages).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////
