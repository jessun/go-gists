package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _AllVirtualColumnStatisticMgr struct {
	*_BaseMgr
}

// AllVirtualColumnStatisticMgr open func
func AllVirtualColumnStatisticMgr(db *gorm.DB) *_AllVirtualColumnStatisticMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualColumnStatisticMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualColumnStatisticMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_column_statistic"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualColumnStatisticMgr) GetTableName() string {
	return "__all_virtual_column_statistic"
}

// Reset 重置gorm会话
func (obj *_AllVirtualColumnStatisticMgr) Reset() *_AllVirtualColumnStatisticMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualColumnStatisticMgr) Get() (result AllVirtualColumnStatistic, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualColumnStatistic{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualColumnStatisticMgr) Gets() (results []*AllVirtualColumnStatistic, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualColumnStatistic{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualColumnStatisticMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualColumnStatistic{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTenantID tenant_id获取
func (obj *_AllVirtualColumnStatisticMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithTableID table_id获取
func (obj *_AllVirtualColumnStatisticMgr) WithTableID(tableID int64) Option {
	return optionFunc(func(o *options) { o.query["table_id"] = tableID })
}

// WithPartitionID partition_id获取
func (obj *_AllVirtualColumnStatisticMgr) WithPartitionID(partitionID int64) Option {
	return optionFunc(func(o *options) { o.query["partition_id"] = partitionID })
}

// WithColumnID column_id获取
func (obj *_AllVirtualColumnStatisticMgr) WithColumnID(columnID int64) Option {
	return optionFunc(func(o *options) { o.query["column_id"] = columnID })
}

// WithGmtCreate gmt_create获取
func (obj *_AllVirtualColumnStatisticMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllVirtualColumnStatisticMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithNumDistinct num_distinct获取
func (obj *_AllVirtualColumnStatisticMgr) WithNumDistinct(numDistinct int64) Option {
	return optionFunc(func(o *options) { o.query["num_distinct"] = numDistinct })
}

// WithNumNull num_null获取
func (obj *_AllVirtualColumnStatisticMgr) WithNumNull(numNull int64) Option {
	return optionFunc(func(o *options) { o.query["num_null"] = numNull })
}

// WithMinValue min_value获取
func (obj *_AllVirtualColumnStatisticMgr) WithMinValue(minValue string) Option {
	return optionFunc(func(o *options) { o.query["min_value"] = minValue })
}

// WithMaxValue max_value获取
func (obj *_AllVirtualColumnStatisticMgr) WithMaxValue(maxValue string) Option {
	return optionFunc(func(o *options) { o.query["max_value"] = maxValue })
}

// WithLlcBitmap llc_bitmap获取
func (obj *_AllVirtualColumnStatisticMgr) WithLlcBitmap(llcBitmap string) Option {
	return optionFunc(func(o *options) { o.query["llc_bitmap"] = llcBitmap })
}

// WithLlcBitmapSize llc_bitmap_size获取
func (obj *_AllVirtualColumnStatisticMgr) WithLlcBitmapSize(llcBitmapSize int64) Option {
	return optionFunc(func(o *options) { o.query["llc_bitmap_size"] = llcBitmapSize })
}

// WithVersion version获取
func (obj *_AllVirtualColumnStatisticMgr) WithVersion(version int64) Option {
	return optionFunc(func(o *options) { o.query["version"] = version })
}

// WithLastRebuildVersion last_rebuild_version获取
func (obj *_AllVirtualColumnStatisticMgr) WithLastRebuildVersion(lastRebuildVersion int64) Option {
	return optionFunc(func(o *options) { o.query["last_rebuild_version"] = lastRebuildVersion })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualColumnStatisticMgr) GetByOption(opts ...Option) (result AllVirtualColumnStatistic, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualColumnStatistic{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualColumnStatisticMgr) GetByOptions(opts ...Option) (results []*AllVirtualColumnStatistic, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualColumnStatistic{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualColumnStatisticMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualColumnStatistic, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualColumnStatistic{}).Where(options.query)
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
func (obj *_AllVirtualColumnStatisticMgr) GetFromTenantID(tenantID int64) (results []*AllVirtualColumnStatistic, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualColumnStatistic{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllVirtualColumnStatisticMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllVirtualColumnStatistic, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualColumnStatistic{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromTableID 通过table_id获取内容
func (obj *_AllVirtualColumnStatisticMgr) GetFromTableID(tableID int64) (results []*AllVirtualColumnStatistic, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualColumnStatistic{}).Where("`table_id` = ?", tableID).Find(&results).Error

	return
}

// GetBatchFromTableID 批量查找
func (obj *_AllVirtualColumnStatisticMgr) GetBatchFromTableID(tableIDs []int64) (results []*AllVirtualColumnStatistic, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualColumnStatistic{}).Where("`table_id` IN (?)", tableIDs).Find(&results).Error

	return
}

// GetFromPartitionID 通过partition_id获取内容
func (obj *_AllVirtualColumnStatisticMgr) GetFromPartitionID(partitionID int64) (results []*AllVirtualColumnStatistic, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualColumnStatistic{}).Where("`partition_id` = ?", partitionID).Find(&results).Error

	return
}

// GetBatchFromPartitionID 批量查找
func (obj *_AllVirtualColumnStatisticMgr) GetBatchFromPartitionID(partitionIDs []int64) (results []*AllVirtualColumnStatistic, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualColumnStatistic{}).Where("`partition_id` IN (?)", partitionIDs).Find(&results).Error

	return
}

// GetFromColumnID 通过column_id获取内容
func (obj *_AllVirtualColumnStatisticMgr) GetFromColumnID(columnID int64) (results []*AllVirtualColumnStatistic, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualColumnStatistic{}).Where("`column_id` = ?", columnID).Find(&results).Error

	return
}

// GetBatchFromColumnID 批量查找
func (obj *_AllVirtualColumnStatisticMgr) GetBatchFromColumnID(columnIDs []int64) (results []*AllVirtualColumnStatistic, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualColumnStatistic{}).Where("`column_id` IN (?)", columnIDs).Find(&results).Error

	return
}

// GetFromGmtCreate 通过gmt_create获取内容
func (obj *_AllVirtualColumnStatisticMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllVirtualColumnStatistic, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualColumnStatistic{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllVirtualColumnStatisticMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllVirtualColumnStatistic, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualColumnStatistic{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllVirtualColumnStatisticMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllVirtualColumnStatistic, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualColumnStatistic{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllVirtualColumnStatisticMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllVirtualColumnStatistic, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualColumnStatistic{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromNumDistinct 通过num_distinct获取内容
func (obj *_AllVirtualColumnStatisticMgr) GetFromNumDistinct(numDistinct int64) (results []*AllVirtualColumnStatistic, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualColumnStatistic{}).Where("`num_distinct` = ?", numDistinct).Find(&results).Error

	return
}

// GetBatchFromNumDistinct 批量查找
func (obj *_AllVirtualColumnStatisticMgr) GetBatchFromNumDistinct(numDistincts []int64) (results []*AllVirtualColumnStatistic, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualColumnStatistic{}).Where("`num_distinct` IN (?)", numDistincts).Find(&results).Error

	return
}

// GetFromNumNull 通过num_null获取内容
func (obj *_AllVirtualColumnStatisticMgr) GetFromNumNull(numNull int64) (results []*AllVirtualColumnStatistic, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualColumnStatistic{}).Where("`num_null` = ?", numNull).Find(&results).Error

	return
}

// GetBatchFromNumNull 批量查找
func (obj *_AllVirtualColumnStatisticMgr) GetBatchFromNumNull(numNulls []int64) (results []*AllVirtualColumnStatistic, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualColumnStatistic{}).Where("`num_null` IN (?)", numNulls).Find(&results).Error

	return
}

// GetFromMinValue 通过min_value获取内容
func (obj *_AllVirtualColumnStatisticMgr) GetFromMinValue(minValue string) (results []*AllVirtualColumnStatistic, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualColumnStatistic{}).Where("`min_value` = ?", minValue).Find(&results).Error

	return
}

// GetBatchFromMinValue 批量查找
func (obj *_AllVirtualColumnStatisticMgr) GetBatchFromMinValue(minValues []string) (results []*AllVirtualColumnStatistic, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualColumnStatistic{}).Where("`min_value` IN (?)", minValues).Find(&results).Error

	return
}

// GetFromMaxValue 通过max_value获取内容
func (obj *_AllVirtualColumnStatisticMgr) GetFromMaxValue(maxValue string) (results []*AllVirtualColumnStatistic, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualColumnStatistic{}).Where("`max_value` = ?", maxValue).Find(&results).Error

	return
}

// GetBatchFromMaxValue 批量查找
func (obj *_AllVirtualColumnStatisticMgr) GetBatchFromMaxValue(maxValues []string) (results []*AllVirtualColumnStatistic, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualColumnStatistic{}).Where("`max_value` IN (?)", maxValues).Find(&results).Error

	return
}

// GetFromLlcBitmap 通过llc_bitmap获取内容
func (obj *_AllVirtualColumnStatisticMgr) GetFromLlcBitmap(llcBitmap string) (results []*AllVirtualColumnStatistic, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualColumnStatistic{}).Where("`llc_bitmap` = ?", llcBitmap).Find(&results).Error

	return
}

// GetBatchFromLlcBitmap 批量查找
func (obj *_AllVirtualColumnStatisticMgr) GetBatchFromLlcBitmap(llcBitmaps []string) (results []*AllVirtualColumnStatistic, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualColumnStatistic{}).Where("`llc_bitmap` IN (?)", llcBitmaps).Find(&results).Error

	return
}

// GetFromLlcBitmapSize 通过llc_bitmap_size获取内容
func (obj *_AllVirtualColumnStatisticMgr) GetFromLlcBitmapSize(llcBitmapSize int64) (results []*AllVirtualColumnStatistic, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualColumnStatistic{}).Where("`llc_bitmap_size` = ?", llcBitmapSize).Find(&results).Error

	return
}

// GetBatchFromLlcBitmapSize 批量查找
func (obj *_AllVirtualColumnStatisticMgr) GetBatchFromLlcBitmapSize(llcBitmapSizes []int64) (results []*AllVirtualColumnStatistic, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualColumnStatistic{}).Where("`llc_bitmap_size` IN (?)", llcBitmapSizes).Find(&results).Error

	return
}

// GetFromVersion 通过version获取内容
func (obj *_AllVirtualColumnStatisticMgr) GetFromVersion(version int64) (results []*AllVirtualColumnStatistic, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualColumnStatistic{}).Where("`version` = ?", version).Find(&results).Error

	return
}

// GetBatchFromVersion 批量查找
func (obj *_AllVirtualColumnStatisticMgr) GetBatchFromVersion(versions []int64) (results []*AllVirtualColumnStatistic, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualColumnStatistic{}).Where("`version` IN (?)", versions).Find(&results).Error

	return
}

// GetFromLastRebuildVersion 通过last_rebuild_version获取内容
func (obj *_AllVirtualColumnStatisticMgr) GetFromLastRebuildVersion(lastRebuildVersion int64) (results []*AllVirtualColumnStatistic, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualColumnStatistic{}).Where("`last_rebuild_version` = ?", lastRebuildVersion).Find(&results).Error

	return
}

// GetBatchFromLastRebuildVersion 批量查找
func (obj *_AllVirtualColumnStatisticMgr) GetBatchFromLastRebuildVersion(lastRebuildVersions []int64) (results []*AllVirtualColumnStatistic, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualColumnStatistic{}).Where("`last_rebuild_version` IN (?)", lastRebuildVersions).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllVirtualColumnStatisticMgr) FetchByPrimaryKey(tenantID int64, tableID int64, partitionID int64, columnID int64) (result AllVirtualColumnStatistic, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualColumnStatistic{}).Where("`tenant_id` = ? AND `table_id` = ? AND `partition_id` = ? AND `column_id` = ?", tenantID, tableID, partitionID, columnID).First(&result).Error

	return
}
