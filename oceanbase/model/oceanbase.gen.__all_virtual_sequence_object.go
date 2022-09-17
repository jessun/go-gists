package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _AllVirtualSequenceObjectMgr struct {
	*_BaseMgr
}

// AllVirtualSequenceObjectMgr open func
func AllVirtualSequenceObjectMgr(db *gorm.DB) *_AllVirtualSequenceObjectMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualSequenceObjectMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualSequenceObjectMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_sequence_object"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualSequenceObjectMgr) GetTableName() string {
	return "__all_virtual_sequence_object"
}

// Reset 重置gorm会话
func (obj *_AllVirtualSequenceObjectMgr) Reset() *_AllVirtualSequenceObjectMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualSequenceObjectMgr) Get() (result AllVirtualSequenceObject, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSequenceObject{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualSequenceObjectMgr) Gets() (results []*AllVirtualSequenceObject, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSequenceObject{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualSequenceObjectMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualSequenceObject{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTenantID tenant_id获取
func (obj *_AllVirtualSequenceObjectMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithSequenceID sequence_id获取
func (obj *_AllVirtualSequenceObjectMgr) WithSequenceID(sequenceID int64) Option {
	return optionFunc(func(o *options) { o.query["sequence_id"] = sequenceID })
}

// WithGmtCreate gmt_create获取
func (obj *_AllVirtualSequenceObjectMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllVirtualSequenceObjectMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithSchemaVersion schema_version获取
func (obj *_AllVirtualSequenceObjectMgr) WithSchemaVersion(schemaVersion int64) Option {
	return optionFunc(func(o *options) { o.query["schema_version"] = schemaVersion })
}

// WithDatabaseID database_id获取
func (obj *_AllVirtualSequenceObjectMgr) WithDatabaseID(databaseID int64) Option {
	return optionFunc(func(o *options) { o.query["database_id"] = databaseID })
}

// WithSequenceName sequence_name获取
func (obj *_AllVirtualSequenceObjectMgr) WithSequenceName(sequenceName string) Option {
	return optionFunc(func(o *options) { o.query["sequence_name"] = sequenceName })
}

// WithMinValue min_value获取
func (obj *_AllVirtualSequenceObjectMgr) WithMinValue(minValue float64) Option {
	return optionFunc(func(o *options) { o.query["min_value"] = minValue })
}

// WithMaxValue max_value获取
func (obj *_AllVirtualSequenceObjectMgr) WithMaxValue(maxValue float64) Option {
	return optionFunc(func(o *options) { o.query["max_value"] = maxValue })
}

// WithIncrementBy increment_by获取
func (obj *_AllVirtualSequenceObjectMgr) WithIncrementBy(incrementBy float64) Option {
	return optionFunc(func(o *options) { o.query["increment_by"] = incrementBy })
}

// WithStartWith start_with获取
func (obj *_AllVirtualSequenceObjectMgr) WithStartWith(startWith float64) Option {
	return optionFunc(func(o *options) { o.query["start_with"] = startWith })
}

// WithCacheSize cache_size获取
func (obj *_AllVirtualSequenceObjectMgr) WithCacheSize(cacheSize float64) Option {
	return optionFunc(func(o *options) { o.query["cache_size"] = cacheSize })
}

// WithOrderFlag order_flag获取
func (obj *_AllVirtualSequenceObjectMgr) WithOrderFlag(orderFlag int8) Option {
	return optionFunc(func(o *options) { o.query["order_flag"] = orderFlag })
}

// WithCycleFlag cycle_flag获取
func (obj *_AllVirtualSequenceObjectMgr) WithCycleFlag(cycleFlag int8) Option {
	return optionFunc(func(o *options) { o.query["cycle_flag"] = cycleFlag })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualSequenceObjectMgr) GetByOption(opts ...Option) (result AllVirtualSequenceObject, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSequenceObject{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualSequenceObjectMgr) GetByOptions(opts ...Option) (results []*AllVirtualSequenceObject, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSequenceObject{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualSequenceObjectMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualSequenceObject, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualSequenceObject{}).Where(options.query)
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
func (obj *_AllVirtualSequenceObjectMgr) GetFromTenantID(tenantID int64) (results []*AllVirtualSequenceObject, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSequenceObject{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllVirtualSequenceObjectMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllVirtualSequenceObject, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSequenceObject{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromSequenceID 通过sequence_id获取内容
func (obj *_AllVirtualSequenceObjectMgr) GetFromSequenceID(sequenceID int64) (results []*AllVirtualSequenceObject, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSequenceObject{}).Where("`sequence_id` = ?", sequenceID).Find(&results).Error

	return
}

// GetBatchFromSequenceID 批量查找
func (obj *_AllVirtualSequenceObjectMgr) GetBatchFromSequenceID(sequenceIDs []int64) (results []*AllVirtualSequenceObject, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSequenceObject{}).Where("`sequence_id` IN (?)", sequenceIDs).Find(&results).Error

	return
}

// GetFromGmtCreate 通过gmt_create获取内容
func (obj *_AllVirtualSequenceObjectMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllVirtualSequenceObject, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSequenceObject{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllVirtualSequenceObjectMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllVirtualSequenceObject, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSequenceObject{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllVirtualSequenceObjectMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllVirtualSequenceObject, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSequenceObject{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllVirtualSequenceObjectMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllVirtualSequenceObject, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSequenceObject{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromSchemaVersion 通过schema_version获取内容
func (obj *_AllVirtualSequenceObjectMgr) GetFromSchemaVersion(schemaVersion int64) (results []*AllVirtualSequenceObject, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSequenceObject{}).Where("`schema_version` = ?", schemaVersion).Find(&results).Error

	return
}

// GetBatchFromSchemaVersion 批量查找
func (obj *_AllVirtualSequenceObjectMgr) GetBatchFromSchemaVersion(schemaVersions []int64) (results []*AllVirtualSequenceObject, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSequenceObject{}).Where("`schema_version` IN (?)", schemaVersions).Find(&results).Error

	return
}

// GetFromDatabaseID 通过database_id获取内容
func (obj *_AllVirtualSequenceObjectMgr) GetFromDatabaseID(databaseID int64) (results []*AllVirtualSequenceObject, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSequenceObject{}).Where("`database_id` = ?", databaseID).Find(&results).Error

	return
}

// GetBatchFromDatabaseID 批量查找
func (obj *_AllVirtualSequenceObjectMgr) GetBatchFromDatabaseID(databaseIDs []int64) (results []*AllVirtualSequenceObject, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSequenceObject{}).Where("`database_id` IN (?)", databaseIDs).Find(&results).Error

	return
}

// GetFromSequenceName 通过sequence_name获取内容
func (obj *_AllVirtualSequenceObjectMgr) GetFromSequenceName(sequenceName string) (results []*AllVirtualSequenceObject, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSequenceObject{}).Where("`sequence_name` = ?", sequenceName).Find(&results).Error

	return
}

// GetBatchFromSequenceName 批量查找
func (obj *_AllVirtualSequenceObjectMgr) GetBatchFromSequenceName(sequenceNames []string) (results []*AllVirtualSequenceObject, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSequenceObject{}).Where("`sequence_name` IN (?)", sequenceNames).Find(&results).Error

	return
}

// GetFromMinValue 通过min_value获取内容
func (obj *_AllVirtualSequenceObjectMgr) GetFromMinValue(minValue float64) (results []*AllVirtualSequenceObject, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSequenceObject{}).Where("`min_value` = ?", minValue).Find(&results).Error

	return
}

// GetBatchFromMinValue 批量查找
func (obj *_AllVirtualSequenceObjectMgr) GetBatchFromMinValue(minValues []float64) (results []*AllVirtualSequenceObject, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSequenceObject{}).Where("`min_value` IN (?)", minValues).Find(&results).Error

	return
}

// GetFromMaxValue 通过max_value获取内容
func (obj *_AllVirtualSequenceObjectMgr) GetFromMaxValue(maxValue float64) (results []*AllVirtualSequenceObject, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSequenceObject{}).Where("`max_value` = ?", maxValue).Find(&results).Error

	return
}

// GetBatchFromMaxValue 批量查找
func (obj *_AllVirtualSequenceObjectMgr) GetBatchFromMaxValue(maxValues []float64) (results []*AllVirtualSequenceObject, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSequenceObject{}).Where("`max_value` IN (?)", maxValues).Find(&results).Error

	return
}

// GetFromIncrementBy 通过increment_by获取内容
func (obj *_AllVirtualSequenceObjectMgr) GetFromIncrementBy(incrementBy float64) (results []*AllVirtualSequenceObject, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSequenceObject{}).Where("`increment_by` = ?", incrementBy).Find(&results).Error

	return
}

// GetBatchFromIncrementBy 批量查找
func (obj *_AllVirtualSequenceObjectMgr) GetBatchFromIncrementBy(incrementBys []float64) (results []*AllVirtualSequenceObject, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSequenceObject{}).Where("`increment_by` IN (?)", incrementBys).Find(&results).Error

	return
}

// GetFromStartWith 通过start_with获取内容
func (obj *_AllVirtualSequenceObjectMgr) GetFromStartWith(startWith float64) (results []*AllVirtualSequenceObject, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSequenceObject{}).Where("`start_with` = ?", startWith).Find(&results).Error

	return
}

// GetBatchFromStartWith 批量查找
func (obj *_AllVirtualSequenceObjectMgr) GetBatchFromStartWith(startWiths []float64) (results []*AllVirtualSequenceObject, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSequenceObject{}).Where("`start_with` IN (?)", startWiths).Find(&results).Error

	return
}

// GetFromCacheSize 通过cache_size获取内容
func (obj *_AllVirtualSequenceObjectMgr) GetFromCacheSize(cacheSize float64) (results []*AllVirtualSequenceObject, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSequenceObject{}).Where("`cache_size` = ?", cacheSize).Find(&results).Error

	return
}

// GetBatchFromCacheSize 批量查找
func (obj *_AllVirtualSequenceObjectMgr) GetBatchFromCacheSize(cacheSizes []float64) (results []*AllVirtualSequenceObject, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSequenceObject{}).Where("`cache_size` IN (?)", cacheSizes).Find(&results).Error

	return
}

// GetFromOrderFlag 通过order_flag获取内容
func (obj *_AllVirtualSequenceObjectMgr) GetFromOrderFlag(orderFlag int8) (results []*AllVirtualSequenceObject, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSequenceObject{}).Where("`order_flag` = ?", orderFlag).Find(&results).Error

	return
}

// GetBatchFromOrderFlag 批量查找
func (obj *_AllVirtualSequenceObjectMgr) GetBatchFromOrderFlag(orderFlags []int8) (results []*AllVirtualSequenceObject, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSequenceObject{}).Where("`order_flag` IN (?)", orderFlags).Find(&results).Error

	return
}

// GetFromCycleFlag 通过cycle_flag获取内容
func (obj *_AllVirtualSequenceObjectMgr) GetFromCycleFlag(cycleFlag int8) (results []*AllVirtualSequenceObject, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSequenceObject{}).Where("`cycle_flag` = ?", cycleFlag).Find(&results).Error

	return
}

// GetBatchFromCycleFlag 批量查找
func (obj *_AllVirtualSequenceObjectMgr) GetBatchFromCycleFlag(cycleFlags []int8) (results []*AllVirtualSequenceObject, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSequenceObject{}).Where("`cycle_flag` IN (?)", cycleFlags).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllVirtualSequenceObjectMgr) FetchByPrimaryKey(tenantID int64, sequenceID int64) (result AllVirtualSequenceObject, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSequenceObject{}).Where("`tenant_id` = ? AND `sequence_id` = ?", tenantID, sequenceID).First(&result).Error

	return
}
