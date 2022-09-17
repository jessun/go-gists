package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _AllColumnStatMgr struct {
	*_BaseMgr
}

// AllColumnStatMgr open func
func AllColumnStatMgr(db *gorm.DB) *_AllColumnStatMgr {
	if db == nil {
		panic(fmt.Errorf("AllColumnStatMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllColumnStatMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_column_stat"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllColumnStatMgr) GetTableName() string {
	return "__all_column_stat"
}

// Reset 重置gorm会话
func (obj *_AllColumnStatMgr) Reset() *_AllColumnStatMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllColumnStatMgr) Get() (result AllColumnStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllColumnStat{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllColumnStatMgr) Gets() (results []*AllColumnStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllColumnStat{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllColumnStatMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllColumnStat{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithGmtCreate gmt_create获取
func (obj *_AllColumnStatMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllColumnStatMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithTenantID tenant_id获取
func (obj *_AllColumnStatMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithTableID table_id获取
func (obj *_AllColumnStatMgr) WithTableID(tableID int64) Option {
	return optionFunc(func(o *options) { o.query["table_id"] = tableID })
}

// WithPartitionID partition_id获取
func (obj *_AllColumnStatMgr) WithPartitionID(partitionID int64) Option {
	return optionFunc(func(o *options) { o.query["partition_id"] = partitionID })
}

// WithColumnID column_id获取
func (obj *_AllColumnStatMgr) WithColumnID(columnID int64) Option {
	return optionFunc(func(o *options) { o.query["column_id"] = columnID })
}

// WithObjectType object_type获取
func (obj *_AllColumnStatMgr) WithObjectType(objectType int64) Option {
	return optionFunc(func(o *options) { o.query["object_type"] = objectType })
}

// WithLastAnalyzed last_analyzed获取
func (obj *_AllColumnStatMgr) WithLastAnalyzed(lastAnalyzed time.Time) Option {
	return optionFunc(func(o *options) { o.query["last_analyzed"] = lastAnalyzed })
}

// WithDistinctCnt distinct_cnt获取
func (obj *_AllColumnStatMgr) WithDistinctCnt(distinctCnt int64) Option {
	return optionFunc(func(o *options) { o.query["distinct_cnt"] = distinctCnt })
}

// WithNullCnt null_cnt获取
func (obj *_AllColumnStatMgr) WithNullCnt(nullCnt int64) Option {
	return optionFunc(func(o *options) { o.query["null_cnt"] = nullCnt })
}

// WithMaxValue max_value获取
func (obj *_AllColumnStatMgr) WithMaxValue(maxValue string) Option {
	return optionFunc(func(o *options) { o.query["max_value"] = maxValue })
}

// WithBMaxValue b_max_value获取
func (obj *_AllColumnStatMgr) WithBMaxValue(bMaxValue string) Option {
	return optionFunc(func(o *options) { o.query["b_max_value"] = bMaxValue })
}

// WithMinValue min_value获取
func (obj *_AllColumnStatMgr) WithMinValue(minValue string) Option {
	return optionFunc(func(o *options) { o.query["min_value"] = minValue })
}

// WithBMinValue b_min_value获取
func (obj *_AllColumnStatMgr) WithBMinValue(bMinValue string) Option {
	return optionFunc(func(o *options) { o.query["b_min_value"] = bMinValue })
}

// WithAvgLen avg_len获取
func (obj *_AllColumnStatMgr) WithAvgLen(avgLen int64) Option {
	return optionFunc(func(o *options) { o.query["avg_len"] = avgLen })
}

// WithDistinctCntSynopsis distinct_cnt_synopsis获取
func (obj *_AllColumnStatMgr) WithDistinctCntSynopsis(distinctCntSynopsis string) Option {
	return optionFunc(func(o *options) { o.query["distinct_cnt_synopsis"] = distinctCntSynopsis })
}

// WithDistinctCntSynopsisSize distinct_cnt_synopsis_size获取
func (obj *_AllColumnStatMgr) WithDistinctCntSynopsisSize(distinctCntSynopsisSize int64) Option {
	return optionFunc(func(o *options) { o.query["distinct_cnt_synopsis_size"] = distinctCntSynopsisSize })
}

// WithSampleSize sample_size获取
func (obj *_AllColumnStatMgr) WithSampleSize(sampleSize float64) Option {
	return optionFunc(func(o *options) { o.query["sample_size"] = sampleSize })
}

// WithDensity density获取
func (obj *_AllColumnStatMgr) WithDensity(density float64) Option {
	return optionFunc(func(o *options) { o.query["density"] = density })
}

// WithBucketCnt bucket_cnt获取
func (obj *_AllColumnStatMgr) WithBucketCnt(bucketCnt int64) Option {
	return optionFunc(func(o *options) { o.query["bucket_cnt"] = bucketCnt })
}

// WithHistogramType histogram_type获取
func (obj *_AllColumnStatMgr) WithHistogramType(histogramType int64) Option {
	return optionFunc(func(o *options) { o.query["histogram_type"] = histogramType })
}

// GetByOption 功能选项模式获取
func (obj *_AllColumnStatMgr) GetByOption(opts ...Option) (result AllColumnStat, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllColumnStat{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllColumnStatMgr) GetByOptions(opts ...Option) (results []*AllColumnStat, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllColumnStat{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllColumnStatMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllColumnStat, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllColumnStat{}).Where(options.query)
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
func (obj *_AllColumnStatMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllColumnStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllColumnStat{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllColumnStatMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllColumnStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllColumnStat{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllColumnStatMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllColumnStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllColumnStat{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllColumnStatMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllColumnStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllColumnStat{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromTenantID 通过tenant_id获取内容
func (obj *_AllColumnStatMgr) GetFromTenantID(tenantID int64) (results []*AllColumnStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllColumnStat{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllColumnStatMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllColumnStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllColumnStat{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromTableID 通过table_id获取内容
func (obj *_AllColumnStatMgr) GetFromTableID(tableID int64) (results []*AllColumnStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllColumnStat{}).Where("`table_id` = ?", tableID).Find(&results).Error

	return
}

// GetBatchFromTableID 批量查找
func (obj *_AllColumnStatMgr) GetBatchFromTableID(tableIDs []int64) (results []*AllColumnStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllColumnStat{}).Where("`table_id` IN (?)", tableIDs).Find(&results).Error

	return
}

// GetFromPartitionID 通过partition_id获取内容
func (obj *_AllColumnStatMgr) GetFromPartitionID(partitionID int64) (results []*AllColumnStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllColumnStat{}).Where("`partition_id` = ?", partitionID).Find(&results).Error

	return
}

// GetBatchFromPartitionID 批量查找
func (obj *_AllColumnStatMgr) GetBatchFromPartitionID(partitionIDs []int64) (results []*AllColumnStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllColumnStat{}).Where("`partition_id` IN (?)", partitionIDs).Find(&results).Error

	return
}

// GetFromColumnID 通过column_id获取内容
func (obj *_AllColumnStatMgr) GetFromColumnID(columnID int64) (results []*AllColumnStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllColumnStat{}).Where("`column_id` = ?", columnID).Find(&results).Error

	return
}

// GetBatchFromColumnID 批量查找
func (obj *_AllColumnStatMgr) GetBatchFromColumnID(columnIDs []int64) (results []*AllColumnStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllColumnStat{}).Where("`column_id` IN (?)", columnIDs).Find(&results).Error

	return
}

// GetFromObjectType 通过object_type获取内容
func (obj *_AllColumnStatMgr) GetFromObjectType(objectType int64) (results []*AllColumnStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllColumnStat{}).Where("`object_type` = ?", objectType).Find(&results).Error

	return
}

// GetBatchFromObjectType 批量查找
func (obj *_AllColumnStatMgr) GetBatchFromObjectType(objectTypes []int64) (results []*AllColumnStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllColumnStat{}).Where("`object_type` IN (?)", objectTypes).Find(&results).Error

	return
}

// GetFromLastAnalyzed 通过last_analyzed获取内容
func (obj *_AllColumnStatMgr) GetFromLastAnalyzed(lastAnalyzed time.Time) (results []*AllColumnStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllColumnStat{}).Where("`last_analyzed` = ?", lastAnalyzed).Find(&results).Error

	return
}

// GetBatchFromLastAnalyzed 批量查找
func (obj *_AllColumnStatMgr) GetBatchFromLastAnalyzed(lastAnalyzeds []time.Time) (results []*AllColumnStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllColumnStat{}).Where("`last_analyzed` IN (?)", lastAnalyzeds).Find(&results).Error

	return
}

// GetFromDistinctCnt 通过distinct_cnt获取内容
func (obj *_AllColumnStatMgr) GetFromDistinctCnt(distinctCnt int64) (results []*AllColumnStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllColumnStat{}).Where("`distinct_cnt` = ?", distinctCnt).Find(&results).Error

	return
}

// GetBatchFromDistinctCnt 批量查找
func (obj *_AllColumnStatMgr) GetBatchFromDistinctCnt(distinctCnts []int64) (results []*AllColumnStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllColumnStat{}).Where("`distinct_cnt` IN (?)", distinctCnts).Find(&results).Error

	return
}

// GetFromNullCnt 通过null_cnt获取内容
func (obj *_AllColumnStatMgr) GetFromNullCnt(nullCnt int64) (results []*AllColumnStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllColumnStat{}).Where("`null_cnt` = ?", nullCnt).Find(&results).Error

	return
}

// GetBatchFromNullCnt 批量查找
func (obj *_AllColumnStatMgr) GetBatchFromNullCnt(nullCnts []int64) (results []*AllColumnStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllColumnStat{}).Where("`null_cnt` IN (?)", nullCnts).Find(&results).Error

	return
}

// GetFromMaxValue 通过max_value获取内容
func (obj *_AllColumnStatMgr) GetFromMaxValue(maxValue string) (results []*AllColumnStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllColumnStat{}).Where("`max_value` = ?", maxValue).Find(&results).Error

	return
}

// GetBatchFromMaxValue 批量查找
func (obj *_AllColumnStatMgr) GetBatchFromMaxValue(maxValues []string) (results []*AllColumnStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllColumnStat{}).Where("`max_value` IN (?)", maxValues).Find(&results).Error

	return
}

// GetFromBMaxValue 通过b_max_value获取内容
func (obj *_AllColumnStatMgr) GetFromBMaxValue(bMaxValue string) (results []*AllColumnStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllColumnStat{}).Where("`b_max_value` = ?", bMaxValue).Find(&results).Error

	return
}

// GetBatchFromBMaxValue 批量查找
func (obj *_AllColumnStatMgr) GetBatchFromBMaxValue(bMaxValues []string) (results []*AllColumnStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllColumnStat{}).Where("`b_max_value` IN (?)", bMaxValues).Find(&results).Error

	return
}

// GetFromMinValue 通过min_value获取内容
func (obj *_AllColumnStatMgr) GetFromMinValue(minValue string) (results []*AllColumnStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllColumnStat{}).Where("`min_value` = ?", minValue).Find(&results).Error

	return
}

// GetBatchFromMinValue 批量查找
func (obj *_AllColumnStatMgr) GetBatchFromMinValue(minValues []string) (results []*AllColumnStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllColumnStat{}).Where("`min_value` IN (?)", minValues).Find(&results).Error

	return
}

// GetFromBMinValue 通过b_min_value获取内容
func (obj *_AllColumnStatMgr) GetFromBMinValue(bMinValue string) (results []*AllColumnStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllColumnStat{}).Where("`b_min_value` = ?", bMinValue).Find(&results).Error

	return
}

// GetBatchFromBMinValue 批量查找
func (obj *_AllColumnStatMgr) GetBatchFromBMinValue(bMinValues []string) (results []*AllColumnStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllColumnStat{}).Where("`b_min_value` IN (?)", bMinValues).Find(&results).Error

	return
}

// GetFromAvgLen 通过avg_len获取内容
func (obj *_AllColumnStatMgr) GetFromAvgLen(avgLen int64) (results []*AllColumnStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllColumnStat{}).Where("`avg_len` = ?", avgLen).Find(&results).Error

	return
}

// GetBatchFromAvgLen 批量查找
func (obj *_AllColumnStatMgr) GetBatchFromAvgLen(avgLens []int64) (results []*AllColumnStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllColumnStat{}).Where("`avg_len` IN (?)", avgLens).Find(&results).Error

	return
}

// GetFromDistinctCntSynopsis 通过distinct_cnt_synopsis获取内容
func (obj *_AllColumnStatMgr) GetFromDistinctCntSynopsis(distinctCntSynopsis string) (results []*AllColumnStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllColumnStat{}).Where("`distinct_cnt_synopsis` = ?", distinctCntSynopsis).Find(&results).Error

	return
}

// GetBatchFromDistinctCntSynopsis 批量查找
func (obj *_AllColumnStatMgr) GetBatchFromDistinctCntSynopsis(distinctCntSynopsiss []string) (results []*AllColumnStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllColumnStat{}).Where("`distinct_cnt_synopsis` IN (?)", distinctCntSynopsiss).Find(&results).Error

	return
}

// GetFromDistinctCntSynopsisSize 通过distinct_cnt_synopsis_size获取内容
func (obj *_AllColumnStatMgr) GetFromDistinctCntSynopsisSize(distinctCntSynopsisSize int64) (results []*AllColumnStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllColumnStat{}).Where("`distinct_cnt_synopsis_size` = ?", distinctCntSynopsisSize).Find(&results).Error

	return
}

// GetBatchFromDistinctCntSynopsisSize 批量查找
func (obj *_AllColumnStatMgr) GetBatchFromDistinctCntSynopsisSize(distinctCntSynopsisSizes []int64) (results []*AllColumnStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllColumnStat{}).Where("`distinct_cnt_synopsis_size` IN (?)", distinctCntSynopsisSizes).Find(&results).Error

	return
}

// GetFromSampleSize 通过sample_size获取内容
func (obj *_AllColumnStatMgr) GetFromSampleSize(sampleSize float64) (results []*AllColumnStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllColumnStat{}).Where("`sample_size` = ?", sampleSize).Find(&results).Error

	return
}

// GetBatchFromSampleSize 批量查找
func (obj *_AllColumnStatMgr) GetBatchFromSampleSize(sampleSizes []float64) (results []*AllColumnStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllColumnStat{}).Where("`sample_size` IN (?)", sampleSizes).Find(&results).Error

	return
}

// GetFromDensity 通过density获取内容
func (obj *_AllColumnStatMgr) GetFromDensity(density float64) (results []*AllColumnStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllColumnStat{}).Where("`density` = ?", density).Find(&results).Error

	return
}

// GetBatchFromDensity 批量查找
func (obj *_AllColumnStatMgr) GetBatchFromDensity(densitys []float64) (results []*AllColumnStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllColumnStat{}).Where("`density` IN (?)", densitys).Find(&results).Error

	return
}

// GetFromBucketCnt 通过bucket_cnt获取内容
func (obj *_AllColumnStatMgr) GetFromBucketCnt(bucketCnt int64) (results []*AllColumnStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllColumnStat{}).Where("`bucket_cnt` = ?", bucketCnt).Find(&results).Error

	return
}

// GetBatchFromBucketCnt 批量查找
func (obj *_AllColumnStatMgr) GetBatchFromBucketCnt(bucketCnts []int64) (results []*AllColumnStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllColumnStat{}).Where("`bucket_cnt` IN (?)", bucketCnts).Find(&results).Error

	return
}

// GetFromHistogramType 通过histogram_type获取内容
func (obj *_AllColumnStatMgr) GetFromHistogramType(histogramType int64) (results []*AllColumnStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllColumnStat{}).Where("`histogram_type` = ?", histogramType).Find(&results).Error

	return
}

// GetBatchFromHistogramType 批量查找
func (obj *_AllColumnStatMgr) GetBatchFromHistogramType(histogramTypes []int64) (results []*AllColumnStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllColumnStat{}).Where("`histogram_type` IN (?)", histogramTypes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllColumnStatMgr) FetchByPrimaryKey(tenantID int64, tableID int64, partitionID int64, columnID int64) (result AllColumnStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllColumnStat{}).Where("`tenant_id` = ? AND `table_id` = ? AND `partition_id` = ? AND `column_id` = ?", tenantID, tableID, partitionID, columnID).First(&result).Error

	return
}
