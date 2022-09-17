package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _AllVirtualColumnStatMgr struct {
	*_BaseMgr
}

// AllVirtualColumnStatMgr open func
func AllVirtualColumnStatMgr(db *gorm.DB) *_AllVirtualColumnStatMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualColumnStatMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualColumnStatMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_column_stat"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualColumnStatMgr) GetTableName() string {
	return "__all_virtual_column_stat"
}

// Reset 重置gorm会话
func (obj *_AllVirtualColumnStatMgr) Reset() *_AllVirtualColumnStatMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualColumnStatMgr) Get() (result AllVirtualColumnStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualColumnStat{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualColumnStatMgr) Gets() (results []*AllVirtualColumnStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualColumnStat{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualColumnStatMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualColumnStat{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTenantID tenant_id获取
func (obj *_AllVirtualColumnStatMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithTableID table_id获取
func (obj *_AllVirtualColumnStatMgr) WithTableID(tableID int64) Option {
	return optionFunc(func(o *options) { o.query["table_id"] = tableID })
}

// WithPartitionID partition_id获取
func (obj *_AllVirtualColumnStatMgr) WithPartitionID(partitionID int64) Option {
	return optionFunc(func(o *options) { o.query["partition_id"] = partitionID })
}

// WithColumnID column_id获取
func (obj *_AllVirtualColumnStatMgr) WithColumnID(columnID int64) Option {
	return optionFunc(func(o *options) { o.query["column_id"] = columnID })
}

// WithGmtCreate gmt_create获取
func (obj *_AllVirtualColumnStatMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllVirtualColumnStatMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithObjectType object_type获取
func (obj *_AllVirtualColumnStatMgr) WithObjectType(objectType int64) Option {
	return optionFunc(func(o *options) { o.query["object_type"] = objectType })
}

// WithLastAnalyzed last_analyzed获取
func (obj *_AllVirtualColumnStatMgr) WithLastAnalyzed(lastAnalyzed time.Time) Option {
	return optionFunc(func(o *options) { o.query["last_analyzed"] = lastAnalyzed })
}

// WithDistinctCnt distinct_cnt获取
func (obj *_AllVirtualColumnStatMgr) WithDistinctCnt(distinctCnt int64) Option {
	return optionFunc(func(o *options) { o.query["distinct_cnt"] = distinctCnt })
}

// WithNullCnt null_cnt获取
func (obj *_AllVirtualColumnStatMgr) WithNullCnt(nullCnt int64) Option {
	return optionFunc(func(o *options) { o.query["null_cnt"] = nullCnt })
}

// WithMaxValue max_value获取
func (obj *_AllVirtualColumnStatMgr) WithMaxValue(maxValue string) Option {
	return optionFunc(func(o *options) { o.query["max_value"] = maxValue })
}

// WithBMaxValue b_max_value获取
func (obj *_AllVirtualColumnStatMgr) WithBMaxValue(bMaxValue string) Option {
	return optionFunc(func(o *options) { o.query["b_max_value"] = bMaxValue })
}

// WithMinValue min_value获取
func (obj *_AllVirtualColumnStatMgr) WithMinValue(minValue string) Option {
	return optionFunc(func(o *options) { o.query["min_value"] = minValue })
}

// WithBMinValue b_min_value获取
func (obj *_AllVirtualColumnStatMgr) WithBMinValue(bMinValue string) Option {
	return optionFunc(func(o *options) { o.query["b_min_value"] = bMinValue })
}

// WithAvgLen avg_len获取
func (obj *_AllVirtualColumnStatMgr) WithAvgLen(avgLen int64) Option {
	return optionFunc(func(o *options) { o.query["avg_len"] = avgLen })
}

// WithDistinctCntSynopsis distinct_cnt_synopsis获取
func (obj *_AllVirtualColumnStatMgr) WithDistinctCntSynopsis(distinctCntSynopsis string) Option {
	return optionFunc(func(o *options) { o.query["distinct_cnt_synopsis"] = distinctCntSynopsis })
}

// WithDistinctCntSynopsisSize distinct_cnt_synopsis_size获取
func (obj *_AllVirtualColumnStatMgr) WithDistinctCntSynopsisSize(distinctCntSynopsisSize int64) Option {
	return optionFunc(func(o *options) { o.query["distinct_cnt_synopsis_size"] = distinctCntSynopsisSize })
}

// WithSampleSize sample_size获取
func (obj *_AllVirtualColumnStatMgr) WithSampleSize(sampleSize float64) Option {
	return optionFunc(func(o *options) { o.query["sample_size"] = sampleSize })
}

// WithDensity density获取
func (obj *_AllVirtualColumnStatMgr) WithDensity(density float64) Option {
	return optionFunc(func(o *options) { o.query["density"] = density })
}

// WithBucketCnt bucket_cnt获取
func (obj *_AllVirtualColumnStatMgr) WithBucketCnt(bucketCnt int64) Option {
	return optionFunc(func(o *options) { o.query["bucket_cnt"] = bucketCnt })
}

// WithHistogramType histogram_type获取
func (obj *_AllVirtualColumnStatMgr) WithHistogramType(histogramType int64) Option {
	return optionFunc(func(o *options) { o.query["histogram_type"] = histogramType })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualColumnStatMgr) GetByOption(opts ...Option) (result AllVirtualColumnStat, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualColumnStat{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualColumnStatMgr) GetByOptions(opts ...Option) (results []*AllVirtualColumnStat, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualColumnStat{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualColumnStatMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualColumnStat, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualColumnStat{}).Where(options.query)
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
func (obj *_AllVirtualColumnStatMgr) GetFromTenantID(tenantID int64) (results []*AllVirtualColumnStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualColumnStat{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllVirtualColumnStatMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllVirtualColumnStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualColumnStat{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromTableID 通过table_id获取内容
func (obj *_AllVirtualColumnStatMgr) GetFromTableID(tableID int64) (results []*AllVirtualColumnStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualColumnStat{}).Where("`table_id` = ?", tableID).Find(&results).Error

	return
}

// GetBatchFromTableID 批量查找
func (obj *_AllVirtualColumnStatMgr) GetBatchFromTableID(tableIDs []int64) (results []*AllVirtualColumnStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualColumnStat{}).Where("`table_id` IN (?)", tableIDs).Find(&results).Error

	return
}

// GetFromPartitionID 通过partition_id获取内容
func (obj *_AllVirtualColumnStatMgr) GetFromPartitionID(partitionID int64) (results []*AllVirtualColumnStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualColumnStat{}).Where("`partition_id` = ?", partitionID).Find(&results).Error

	return
}

// GetBatchFromPartitionID 批量查找
func (obj *_AllVirtualColumnStatMgr) GetBatchFromPartitionID(partitionIDs []int64) (results []*AllVirtualColumnStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualColumnStat{}).Where("`partition_id` IN (?)", partitionIDs).Find(&results).Error

	return
}

// GetFromColumnID 通过column_id获取内容
func (obj *_AllVirtualColumnStatMgr) GetFromColumnID(columnID int64) (results []*AllVirtualColumnStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualColumnStat{}).Where("`column_id` = ?", columnID).Find(&results).Error

	return
}

// GetBatchFromColumnID 批量查找
func (obj *_AllVirtualColumnStatMgr) GetBatchFromColumnID(columnIDs []int64) (results []*AllVirtualColumnStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualColumnStat{}).Where("`column_id` IN (?)", columnIDs).Find(&results).Error

	return
}

// GetFromGmtCreate 通过gmt_create获取内容
func (obj *_AllVirtualColumnStatMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllVirtualColumnStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualColumnStat{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllVirtualColumnStatMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllVirtualColumnStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualColumnStat{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllVirtualColumnStatMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllVirtualColumnStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualColumnStat{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllVirtualColumnStatMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllVirtualColumnStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualColumnStat{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromObjectType 通过object_type获取内容
func (obj *_AllVirtualColumnStatMgr) GetFromObjectType(objectType int64) (results []*AllVirtualColumnStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualColumnStat{}).Where("`object_type` = ?", objectType).Find(&results).Error

	return
}

// GetBatchFromObjectType 批量查找
func (obj *_AllVirtualColumnStatMgr) GetBatchFromObjectType(objectTypes []int64) (results []*AllVirtualColumnStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualColumnStat{}).Where("`object_type` IN (?)", objectTypes).Find(&results).Error

	return
}

// GetFromLastAnalyzed 通过last_analyzed获取内容
func (obj *_AllVirtualColumnStatMgr) GetFromLastAnalyzed(lastAnalyzed time.Time) (results []*AllVirtualColumnStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualColumnStat{}).Where("`last_analyzed` = ?", lastAnalyzed).Find(&results).Error

	return
}

// GetBatchFromLastAnalyzed 批量查找
func (obj *_AllVirtualColumnStatMgr) GetBatchFromLastAnalyzed(lastAnalyzeds []time.Time) (results []*AllVirtualColumnStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualColumnStat{}).Where("`last_analyzed` IN (?)", lastAnalyzeds).Find(&results).Error

	return
}

// GetFromDistinctCnt 通过distinct_cnt获取内容
func (obj *_AllVirtualColumnStatMgr) GetFromDistinctCnt(distinctCnt int64) (results []*AllVirtualColumnStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualColumnStat{}).Where("`distinct_cnt` = ?", distinctCnt).Find(&results).Error

	return
}

// GetBatchFromDistinctCnt 批量查找
func (obj *_AllVirtualColumnStatMgr) GetBatchFromDistinctCnt(distinctCnts []int64) (results []*AllVirtualColumnStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualColumnStat{}).Where("`distinct_cnt` IN (?)", distinctCnts).Find(&results).Error

	return
}

// GetFromNullCnt 通过null_cnt获取内容
func (obj *_AllVirtualColumnStatMgr) GetFromNullCnt(nullCnt int64) (results []*AllVirtualColumnStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualColumnStat{}).Where("`null_cnt` = ?", nullCnt).Find(&results).Error

	return
}

// GetBatchFromNullCnt 批量查找
func (obj *_AllVirtualColumnStatMgr) GetBatchFromNullCnt(nullCnts []int64) (results []*AllVirtualColumnStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualColumnStat{}).Where("`null_cnt` IN (?)", nullCnts).Find(&results).Error

	return
}

// GetFromMaxValue 通过max_value获取内容
func (obj *_AllVirtualColumnStatMgr) GetFromMaxValue(maxValue string) (results []*AllVirtualColumnStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualColumnStat{}).Where("`max_value` = ?", maxValue).Find(&results).Error

	return
}

// GetBatchFromMaxValue 批量查找
func (obj *_AllVirtualColumnStatMgr) GetBatchFromMaxValue(maxValues []string) (results []*AllVirtualColumnStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualColumnStat{}).Where("`max_value` IN (?)", maxValues).Find(&results).Error

	return
}

// GetFromBMaxValue 通过b_max_value获取内容
func (obj *_AllVirtualColumnStatMgr) GetFromBMaxValue(bMaxValue string) (results []*AllVirtualColumnStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualColumnStat{}).Where("`b_max_value` = ?", bMaxValue).Find(&results).Error

	return
}

// GetBatchFromBMaxValue 批量查找
func (obj *_AllVirtualColumnStatMgr) GetBatchFromBMaxValue(bMaxValues []string) (results []*AllVirtualColumnStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualColumnStat{}).Where("`b_max_value` IN (?)", bMaxValues).Find(&results).Error

	return
}

// GetFromMinValue 通过min_value获取内容
func (obj *_AllVirtualColumnStatMgr) GetFromMinValue(minValue string) (results []*AllVirtualColumnStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualColumnStat{}).Where("`min_value` = ?", minValue).Find(&results).Error

	return
}

// GetBatchFromMinValue 批量查找
func (obj *_AllVirtualColumnStatMgr) GetBatchFromMinValue(minValues []string) (results []*AllVirtualColumnStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualColumnStat{}).Where("`min_value` IN (?)", minValues).Find(&results).Error

	return
}

// GetFromBMinValue 通过b_min_value获取内容
func (obj *_AllVirtualColumnStatMgr) GetFromBMinValue(bMinValue string) (results []*AllVirtualColumnStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualColumnStat{}).Where("`b_min_value` = ?", bMinValue).Find(&results).Error

	return
}

// GetBatchFromBMinValue 批量查找
func (obj *_AllVirtualColumnStatMgr) GetBatchFromBMinValue(bMinValues []string) (results []*AllVirtualColumnStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualColumnStat{}).Where("`b_min_value` IN (?)", bMinValues).Find(&results).Error

	return
}

// GetFromAvgLen 通过avg_len获取内容
func (obj *_AllVirtualColumnStatMgr) GetFromAvgLen(avgLen int64) (results []*AllVirtualColumnStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualColumnStat{}).Where("`avg_len` = ?", avgLen).Find(&results).Error

	return
}

// GetBatchFromAvgLen 批量查找
func (obj *_AllVirtualColumnStatMgr) GetBatchFromAvgLen(avgLens []int64) (results []*AllVirtualColumnStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualColumnStat{}).Where("`avg_len` IN (?)", avgLens).Find(&results).Error

	return
}

// GetFromDistinctCntSynopsis 通过distinct_cnt_synopsis获取内容
func (obj *_AllVirtualColumnStatMgr) GetFromDistinctCntSynopsis(distinctCntSynopsis string) (results []*AllVirtualColumnStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualColumnStat{}).Where("`distinct_cnt_synopsis` = ?", distinctCntSynopsis).Find(&results).Error

	return
}

// GetBatchFromDistinctCntSynopsis 批量查找
func (obj *_AllVirtualColumnStatMgr) GetBatchFromDistinctCntSynopsis(distinctCntSynopsiss []string) (results []*AllVirtualColumnStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualColumnStat{}).Where("`distinct_cnt_synopsis` IN (?)", distinctCntSynopsiss).Find(&results).Error

	return
}

// GetFromDistinctCntSynopsisSize 通过distinct_cnt_synopsis_size获取内容
func (obj *_AllVirtualColumnStatMgr) GetFromDistinctCntSynopsisSize(distinctCntSynopsisSize int64) (results []*AllVirtualColumnStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualColumnStat{}).Where("`distinct_cnt_synopsis_size` = ?", distinctCntSynopsisSize).Find(&results).Error

	return
}

// GetBatchFromDistinctCntSynopsisSize 批量查找
func (obj *_AllVirtualColumnStatMgr) GetBatchFromDistinctCntSynopsisSize(distinctCntSynopsisSizes []int64) (results []*AllVirtualColumnStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualColumnStat{}).Where("`distinct_cnt_synopsis_size` IN (?)", distinctCntSynopsisSizes).Find(&results).Error

	return
}

// GetFromSampleSize 通过sample_size获取内容
func (obj *_AllVirtualColumnStatMgr) GetFromSampleSize(sampleSize float64) (results []*AllVirtualColumnStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualColumnStat{}).Where("`sample_size` = ?", sampleSize).Find(&results).Error

	return
}

// GetBatchFromSampleSize 批量查找
func (obj *_AllVirtualColumnStatMgr) GetBatchFromSampleSize(sampleSizes []float64) (results []*AllVirtualColumnStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualColumnStat{}).Where("`sample_size` IN (?)", sampleSizes).Find(&results).Error

	return
}

// GetFromDensity 通过density获取内容
func (obj *_AllVirtualColumnStatMgr) GetFromDensity(density float64) (results []*AllVirtualColumnStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualColumnStat{}).Where("`density` = ?", density).Find(&results).Error

	return
}

// GetBatchFromDensity 批量查找
func (obj *_AllVirtualColumnStatMgr) GetBatchFromDensity(densitys []float64) (results []*AllVirtualColumnStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualColumnStat{}).Where("`density` IN (?)", densitys).Find(&results).Error

	return
}

// GetFromBucketCnt 通过bucket_cnt获取内容
func (obj *_AllVirtualColumnStatMgr) GetFromBucketCnt(bucketCnt int64) (results []*AllVirtualColumnStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualColumnStat{}).Where("`bucket_cnt` = ?", bucketCnt).Find(&results).Error

	return
}

// GetBatchFromBucketCnt 批量查找
func (obj *_AllVirtualColumnStatMgr) GetBatchFromBucketCnt(bucketCnts []int64) (results []*AllVirtualColumnStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualColumnStat{}).Where("`bucket_cnt` IN (?)", bucketCnts).Find(&results).Error

	return
}

// GetFromHistogramType 通过histogram_type获取内容
func (obj *_AllVirtualColumnStatMgr) GetFromHistogramType(histogramType int64) (results []*AllVirtualColumnStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualColumnStat{}).Where("`histogram_type` = ?", histogramType).Find(&results).Error

	return
}

// GetBatchFromHistogramType 批量查找
func (obj *_AllVirtualColumnStatMgr) GetBatchFromHistogramType(histogramTypes []int64) (results []*AllVirtualColumnStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualColumnStat{}).Where("`histogram_type` IN (?)", histogramTypes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllVirtualColumnStatMgr) FetchByPrimaryKey(tenantID int64, tableID int64, partitionID int64, columnID int64) (result AllVirtualColumnStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualColumnStat{}).Where("`tenant_id` = ? AND `table_id` = ? AND `partition_id` = ? AND `column_id` = ?", tenantID, tableID, partitionID, columnID).First(&result).Error

	return
}
