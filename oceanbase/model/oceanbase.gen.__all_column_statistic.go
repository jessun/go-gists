package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _AllColumnStatisticMgr struct {
	*_BaseMgr
}

// AllColumnStatisticMgr open func
func AllColumnStatisticMgr(db *gorm.DB) *_AllColumnStatisticMgr {
	if db == nil {
		panic(fmt.Errorf("AllColumnStatisticMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllColumnStatisticMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_column_statistic"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllColumnStatisticMgr) GetTableName() string {
	return "__all_column_statistic"
}

// Reset 重置gorm会话
func (obj *_AllColumnStatisticMgr) Reset() *_AllColumnStatisticMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllColumnStatisticMgr) Get() (result AllColumnStatistic, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllColumnStatistic{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllColumnStatisticMgr) Gets() (results []*AllColumnStatistic, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllColumnStatistic{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllColumnStatisticMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllColumnStatistic{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithGmtCreate gmt_create获取
func (obj *_AllColumnStatisticMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllColumnStatisticMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithTenantID tenant_id获取
func (obj *_AllColumnStatisticMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithTableID table_id获取
func (obj *_AllColumnStatisticMgr) WithTableID(tableID int64) Option {
	return optionFunc(func(o *options) { o.query["table_id"] = tableID })
}

// WithPartitionID partition_id获取
func (obj *_AllColumnStatisticMgr) WithPartitionID(partitionID int64) Option {
	return optionFunc(func(o *options) { o.query["partition_id"] = partitionID })
}

// WithColumnID column_id获取
func (obj *_AllColumnStatisticMgr) WithColumnID(columnID int64) Option {
	return optionFunc(func(o *options) { o.query["column_id"] = columnID })
}

// WithNumDistinct num_distinct获取
func (obj *_AllColumnStatisticMgr) WithNumDistinct(numDistinct int64) Option {
	return optionFunc(func(o *options) { o.query["num_distinct"] = numDistinct })
}

// WithNumNull num_null获取
func (obj *_AllColumnStatisticMgr) WithNumNull(numNull int64) Option {
	return optionFunc(func(o *options) { o.query["num_null"] = numNull })
}

// WithMinValue min_value获取
func (obj *_AllColumnStatisticMgr) WithMinValue(minValue string) Option {
	return optionFunc(func(o *options) { o.query["min_value"] = minValue })
}

// WithMaxValue max_value获取
func (obj *_AllColumnStatisticMgr) WithMaxValue(maxValue string) Option {
	return optionFunc(func(o *options) { o.query["max_value"] = maxValue })
}

// WithLlcBitmap llc_bitmap获取
func (obj *_AllColumnStatisticMgr) WithLlcBitmap(llcBitmap string) Option {
	return optionFunc(func(o *options) { o.query["llc_bitmap"] = llcBitmap })
}

// WithLlcBitmapSize llc_bitmap_size获取
func (obj *_AllColumnStatisticMgr) WithLlcBitmapSize(llcBitmapSize int64) Option {
	return optionFunc(func(o *options) { o.query["llc_bitmap_size"] = llcBitmapSize })
}

// WithVersion version获取
func (obj *_AllColumnStatisticMgr) WithVersion(version int64) Option {
	return optionFunc(func(o *options) { o.query["version"] = version })
}

// WithLastRebuildVersion last_rebuild_version获取
func (obj *_AllColumnStatisticMgr) WithLastRebuildVersion(lastRebuildVersion int64) Option {
	return optionFunc(func(o *options) { o.query["last_rebuild_version"] = lastRebuildVersion })
}

// GetByOption 功能选项模式获取
func (obj *_AllColumnStatisticMgr) GetByOption(opts ...Option) (result AllColumnStatistic, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllColumnStatistic{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllColumnStatisticMgr) GetByOptions(opts ...Option) (results []*AllColumnStatistic, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllColumnStatistic{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllColumnStatisticMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllColumnStatistic, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllColumnStatistic{}).Where(options.query)
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
func (obj *_AllColumnStatisticMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllColumnStatistic, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllColumnStatistic{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllColumnStatisticMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllColumnStatistic, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllColumnStatistic{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllColumnStatisticMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllColumnStatistic, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllColumnStatistic{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllColumnStatisticMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllColumnStatistic, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllColumnStatistic{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromTenantID 通过tenant_id获取内容
func (obj *_AllColumnStatisticMgr) GetFromTenantID(tenantID int64) (results []*AllColumnStatistic, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllColumnStatistic{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllColumnStatisticMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllColumnStatistic, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllColumnStatistic{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromTableID 通过table_id获取内容
func (obj *_AllColumnStatisticMgr) GetFromTableID(tableID int64) (results []*AllColumnStatistic, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllColumnStatistic{}).Where("`table_id` = ?", tableID).Find(&results).Error

	return
}

// GetBatchFromTableID 批量查找
func (obj *_AllColumnStatisticMgr) GetBatchFromTableID(tableIDs []int64) (results []*AllColumnStatistic, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllColumnStatistic{}).Where("`table_id` IN (?)", tableIDs).Find(&results).Error

	return
}

// GetFromPartitionID 通过partition_id获取内容
func (obj *_AllColumnStatisticMgr) GetFromPartitionID(partitionID int64) (results []*AllColumnStatistic, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllColumnStatistic{}).Where("`partition_id` = ?", partitionID).Find(&results).Error

	return
}

// GetBatchFromPartitionID 批量查找
func (obj *_AllColumnStatisticMgr) GetBatchFromPartitionID(partitionIDs []int64) (results []*AllColumnStatistic, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllColumnStatistic{}).Where("`partition_id` IN (?)", partitionIDs).Find(&results).Error

	return
}

// GetFromColumnID 通过column_id获取内容
func (obj *_AllColumnStatisticMgr) GetFromColumnID(columnID int64) (results []*AllColumnStatistic, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllColumnStatistic{}).Where("`column_id` = ?", columnID).Find(&results).Error

	return
}

// GetBatchFromColumnID 批量查找
func (obj *_AllColumnStatisticMgr) GetBatchFromColumnID(columnIDs []int64) (results []*AllColumnStatistic, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllColumnStatistic{}).Where("`column_id` IN (?)", columnIDs).Find(&results).Error

	return
}

// GetFromNumDistinct 通过num_distinct获取内容
func (obj *_AllColumnStatisticMgr) GetFromNumDistinct(numDistinct int64) (results []*AllColumnStatistic, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllColumnStatistic{}).Where("`num_distinct` = ?", numDistinct).Find(&results).Error

	return
}

// GetBatchFromNumDistinct 批量查找
func (obj *_AllColumnStatisticMgr) GetBatchFromNumDistinct(numDistincts []int64) (results []*AllColumnStatistic, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllColumnStatistic{}).Where("`num_distinct` IN (?)", numDistincts).Find(&results).Error

	return
}

// GetFromNumNull 通过num_null获取内容
func (obj *_AllColumnStatisticMgr) GetFromNumNull(numNull int64) (results []*AllColumnStatistic, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllColumnStatistic{}).Where("`num_null` = ?", numNull).Find(&results).Error

	return
}

// GetBatchFromNumNull 批量查找
func (obj *_AllColumnStatisticMgr) GetBatchFromNumNull(numNulls []int64) (results []*AllColumnStatistic, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllColumnStatistic{}).Where("`num_null` IN (?)", numNulls).Find(&results).Error

	return
}

// GetFromMinValue 通过min_value获取内容
func (obj *_AllColumnStatisticMgr) GetFromMinValue(minValue string) (results []*AllColumnStatistic, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllColumnStatistic{}).Where("`min_value` = ?", minValue).Find(&results).Error

	return
}

// GetBatchFromMinValue 批量查找
func (obj *_AllColumnStatisticMgr) GetBatchFromMinValue(minValues []string) (results []*AllColumnStatistic, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllColumnStatistic{}).Where("`min_value` IN (?)", minValues).Find(&results).Error

	return
}

// GetFromMaxValue 通过max_value获取内容
func (obj *_AllColumnStatisticMgr) GetFromMaxValue(maxValue string) (results []*AllColumnStatistic, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllColumnStatistic{}).Where("`max_value` = ?", maxValue).Find(&results).Error

	return
}

// GetBatchFromMaxValue 批量查找
func (obj *_AllColumnStatisticMgr) GetBatchFromMaxValue(maxValues []string) (results []*AllColumnStatistic, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllColumnStatistic{}).Where("`max_value` IN (?)", maxValues).Find(&results).Error

	return
}

// GetFromLlcBitmap 通过llc_bitmap获取内容
func (obj *_AllColumnStatisticMgr) GetFromLlcBitmap(llcBitmap string) (results []*AllColumnStatistic, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllColumnStatistic{}).Where("`llc_bitmap` = ?", llcBitmap).Find(&results).Error

	return
}

// GetBatchFromLlcBitmap 批量查找
func (obj *_AllColumnStatisticMgr) GetBatchFromLlcBitmap(llcBitmaps []string) (results []*AllColumnStatistic, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllColumnStatistic{}).Where("`llc_bitmap` IN (?)", llcBitmaps).Find(&results).Error

	return
}

// GetFromLlcBitmapSize 通过llc_bitmap_size获取内容
func (obj *_AllColumnStatisticMgr) GetFromLlcBitmapSize(llcBitmapSize int64) (results []*AllColumnStatistic, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllColumnStatistic{}).Where("`llc_bitmap_size` = ?", llcBitmapSize).Find(&results).Error

	return
}

// GetBatchFromLlcBitmapSize 批量查找
func (obj *_AllColumnStatisticMgr) GetBatchFromLlcBitmapSize(llcBitmapSizes []int64) (results []*AllColumnStatistic, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllColumnStatistic{}).Where("`llc_bitmap_size` IN (?)", llcBitmapSizes).Find(&results).Error

	return
}

// GetFromVersion 通过version获取内容
func (obj *_AllColumnStatisticMgr) GetFromVersion(version int64) (results []*AllColumnStatistic, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllColumnStatistic{}).Where("`version` = ?", version).Find(&results).Error

	return
}

// GetBatchFromVersion 批量查找
func (obj *_AllColumnStatisticMgr) GetBatchFromVersion(versions []int64) (results []*AllColumnStatistic, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllColumnStatistic{}).Where("`version` IN (?)", versions).Find(&results).Error

	return
}

// GetFromLastRebuildVersion 通过last_rebuild_version获取内容
func (obj *_AllColumnStatisticMgr) GetFromLastRebuildVersion(lastRebuildVersion int64) (results []*AllColumnStatistic, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllColumnStatistic{}).Where("`last_rebuild_version` = ?", lastRebuildVersion).Find(&results).Error

	return
}

// GetBatchFromLastRebuildVersion 批量查找
func (obj *_AllColumnStatisticMgr) GetBatchFromLastRebuildVersion(lastRebuildVersions []int64) (results []*AllColumnStatistic, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllColumnStatistic{}).Where("`last_rebuild_version` IN (?)", lastRebuildVersions).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllColumnStatisticMgr) FetchByPrimaryKey(tenantID int64, tableID int64, partitionID int64, columnID int64) (result AllColumnStatistic, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllColumnStatistic{}).Where("`tenant_id` = ? AND `table_id` = ? AND `partition_id` = ? AND `column_id` = ?", tenantID, tableID, partitionID, columnID).First(&result).Error

	return
}
