package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _AllVirtualSequenceObjectHistoryMgr struct {
	*_BaseMgr
}

// AllVirtualSequenceObjectHistoryMgr open func
func AllVirtualSequenceObjectHistoryMgr(db *gorm.DB) *_AllVirtualSequenceObjectHistoryMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualSequenceObjectHistoryMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualSequenceObjectHistoryMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_sequence_object_history"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualSequenceObjectHistoryMgr) GetTableName() string {
	return "__all_virtual_sequence_object_history"
}

// Reset 重置gorm会话
func (obj *_AllVirtualSequenceObjectHistoryMgr) Reset() *_AllVirtualSequenceObjectHistoryMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualSequenceObjectHistoryMgr) Get() (result AllVirtualSequenceObjectHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSequenceObjectHistory{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualSequenceObjectHistoryMgr) Gets() (results []*AllVirtualSequenceObjectHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSequenceObjectHistory{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualSequenceObjectHistoryMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualSequenceObjectHistory{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTenantID tenant_id获取
func (obj *_AllVirtualSequenceObjectHistoryMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithSequenceID sequence_id获取
func (obj *_AllVirtualSequenceObjectHistoryMgr) WithSequenceID(sequenceID int64) Option {
	return optionFunc(func(o *options) { o.query["sequence_id"] = sequenceID })
}

// WithSchemaVersion schema_version获取
func (obj *_AllVirtualSequenceObjectHistoryMgr) WithSchemaVersion(schemaVersion int64) Option {
	return optionFunc(func(o *options) { o.query["schema_version"] = schemaVersion })
}

// WithGmtCreate gmt_create获取
func (obj *_AllVirtualSequenceObjectHistoryMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllVirtualSequenceObjectHistoryMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithIsDeleted is_deleted获取
func (obj *_AllVirtualSequenceObjectHistoryMgr) WithIsDeleted(isDeleted int64) Option {
	return optionFunc(func(o *options) { o.query["is_deleted"] = isDeleted })
}

// WithDatabaseID database_id获取
func (obj *_AllVirtualSequenceObjectHistoryMgr) WithDatabaseID(databaseID int64) Option {
	return optionFunc(func(o *options) { o.query["database_id"] = databaseID })
}

// WithSequenceName sequence_name获取
func (obj *_AllVirtualSequenceObjectHistoryMgr) WithSequenceName(sequenceName string) Option {
	return optionFunc(func(o *options) { o.query["sequence_name"] = sequenceName })
}

// WithMinValue min_value获取
func (obj *_AllVirtualSequenceObjectHistoryMgr) WithMinValue(minValue float64) Option {
	return optionFunc(func(o *options) { o.query["min_value"] = minValue })
}

// WithMaxValue max_value获取
func (obj *_AllVirtualSequenceObjectHistoryMgr) WithMaxValue(maxValue float64) Option {
	return optionFunc(func(o *options) { o.query["max_value"] = maxValue })
}

// WithIncrementBy increment_by获取
func (obj *_AllVirtualSequenceObjectHistoryMgr) WithIncrementBy(incrementBy float64) Option {
	return optionFunc(func(o *options) { o.query["increment_by"] = incrementBy })
}

// WithStartWith start_with获取
func (obj *_AllVirtualSequenceObjectHistoryMgr) WithStartWith(startWith float64) Option {
	return optionFunc(func(o *options) { o.query["start_with"] = startWith })
}

// WithCacheSize cache_size获取
func (obj *_AllVirtualSequenceObjectHistoryMgr) WithCacheSize(cacheSize float64) Option {
	return optionFunc(func(o *options) { o.query["cache_size"] = cacheSize })
}

// WithOrderFlag order_flag获取
func (obj *_AllVirtualSequenceObjectHistoryMgr) WithOrderFlag(orderFlag int8) Option {
	return optionFunc(func(o *options) { o.query["order_flag"] = orderFlag })
}

// WithCycleFlag cycle_flag获取
func (obj *_AllVirtualSequenceObjectHistoryMgr) WithCycleFlag(cycleFlag int8) Option {
	return optionFunc(func(o *options) { o.query["cycle_flag"] = cycleFlag })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualSequenceObjectHistoryMgr) GetByOption(opts ...Option) (result AllVirtualSequenceObjectHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSequenceObjectHistory{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualSequenceObjectHistoryMgr) GetByOptions(opts ...Option) (results []*AllVirtualSequenceObjectHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSequenceObjectHistory{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualSequenceObjectHistoryMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualSequenceObjectHistory, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualSequenceObjectHistory{}).Where(options.query)
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
func (obj *_AllVirtualSequenceObjectHistoryMgr) GetFromTenantID(tenantID int64) (results []*AllVirtualSequenceObjectHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSequenceObjectHistory{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllVirtualSequenceObjectHistoryMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllVirtualSequenceObjectHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSequenceObjectHistory{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromSequenceID 通过sequence_id获取内容
func (obj *_AllVirtualSequenceObjectHistoryMgr) GetFromSequenceID(sequenceID int64) (results []*AllVirtualSequenceObjectHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSequenceObjectHistory{}).Where("`sequence_id` = ?", sequenceID).Find(&results).Error

	return
}

// GetBatchFromSequenceID 批量查找
func (obj *_AllVirtualSequenceObjectHistoryMgr) GetBatchFromSequenceID(sequenceIDs []int64) (results []*AllVirtualSequenceObjectHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSequenceObjectHistory{}).Where("`sequence_id` IN (?)", sequenceIDs).Find(&results).Error

	return
}

// GetFromSchemaVersion 通过schema_version获取内容
func (obj *_AllVirtualSequenceObjectHistoryMgr) GetFromSchemaVersion(schemaVersion int64) (results []*AllVirtualSequenceObjectHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSequenceObjectHistory{}).Where("`schema_version` = ?", schemaVersion).Find(&results).Error

	return
}

// GetBatchFromSchemaVersion 批量查找
func (obj *_AllVirtualSequenceObjectHistoryMgr) GetBatchFromSchemaVersion(schemaVersions []int64) (results []*AllVirtualSequenceObjectHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSequenceObjectHistory{}).Where("`schema_version` IN (?)", schemaVersions).Find(&results).Error

	return
}

// GetFromGmtCreate 通过gmt_create获取内容
func (obj *_AllVirtualSequenceObjectHistoryMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllVirtualSequenceObjectHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSequenceObjectHistory{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllVirtualSequenceObjectHistoryMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllVirtualSequenceObjectHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSequenceObjectHistory{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllVirtualSequenceObjectHistoryMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllVirtualSequenceObjectHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSequenceObjectHistory{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllVirtualSequenceObjectHistoryMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllVirtualSequenceObjectHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSequenceObjectHistory{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromIsDeleted 通过is_deleted获取内容
func (obj *_AllVirtualSequenceObjectHistoryMgr) GetFromIsDeleted(isDeleted int64) (results []*AllVirtualSequenceObjectHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSequenceObjectHistory{}).Where("`is_deleted` = ?", isDeleted).Find(&results).Error

	return
}

// GetBatchFromIsDeleted 批量查找
func (obj *_AllVirtualSequenceObjectHistoryMgr) GetBatchFromIsDeleted(isDeleteds []int64) (results []*AllVirtualSequenceObjectHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSequenceObjectHistory{}).Where("`is_deleted` IN (?)", isDeleteds).Find(&results).Error

	return
}

// GetFromDatabaseID 通过database_id获取内容
func (obj *_AllVirtualSequenceObjectHistoryMgr) GetFromDatabaseID(databaseID int64) (results []*AllVirtualSequenceObjectHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSequenceObjectHistory{}).Where("`database_id` = ?", databaseID).Find(&results).Error

	return
}

// GetBatchFromDatabaseID 批量查找
func (obj *_AllVirtualSequenceObjectHistoryMgr) GetBatchFromDatabaseID(databaseIDs []int64) (results []*AllVirtualSequenceObjectHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSequenceObjectHistory{}).Where("`database_id` IN (?)", databaseIDs).Find(&results).Error

	return
}

// GetFromSequenceName 通过sequence_name获取内容
func (obj *_AllVirtualSequenceObjectHistoryMgr) GetFromSequenceName(sequenceName string) (results []*AllVirtualSequenceObjectHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSequenceObjectHistory{}).Where("`sequence_name` = ?", sequenceName).Find(&results).Error

	return
}

// GetBatchFromSequenceName 批量查找
func (obj *_AllVirtualSequenceObjectHistoryMgr) GetBatchFromSequenceName(sequenceNames []string) (results []*AllVirtualSequenceObjectHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSequenceObjectHistory{}).Where("`sequence_name` IN (?)", sequenceNames).Find(&results).Error

	return
}

// GetFromMinValue 通过min_value获取内容
func (obj *_AllVirtualSequenceObjectHistoryMgr) GetFromMinValue(minValue float64) (results []*AllVirtualSequenceObjectHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSequenceObjectHistory{}).Where("`min_value` = ?", minValue).Find(&results).Error

	return
}

// GetBatchFromMinValue 批量查找
func (obj *_AllVirtualSequenceObjectHistoryMgr) GetBatchFromMinValue(minValues []float64) (results []*AllVirtualSequenceObjectHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSequenceObjectHistory{}).Where("`min_value` IN (?)", minValues).Find(&results).Error

	return
}

// GetFromMaxValue 通过max_value获取内容
func (obj *_AllVirtualSequenceObjectHistoryMgr) GetFromMaxValue(maxValue float64) (results []*AllVirtualSequenceObjectHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSequenceObjectHistory{}).Where("`max_value` = ?", maxValue).Find(&results).Error

	return
}

// GetBatchFromMaxValue 批量查找
func (obj *_AllVirtualSequenceObjectHistoryMgr) GetBatchFromMaxValue(maxValues []float64) (results []*AllVirtualSequenceObjectHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSequenceObjectHistory{}).Where("`max_value` IN (?)", maxValues).Find(&results).Error

	return
}

// GetFromIncrementBy 通过increment_by获取内容
func (obj *_AllVirtualSequenceObjectHistoryMgr) GetFromIncrementBy(incrementBy float64) (results []*AllVirtualSequenceObjectHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSequenceObjectHistory{}).Where("`increment_by` = ?", incrementBy).Find(&results).Error

	return
}

// GetBatchFromIncrementBy 批量查找
func (obj *_AllVirtualSequenceObjectHistoryMgr) GetBatchFromIncrementBy(incrementBys []float64) (results []*AllVirtualSequenceObjectHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSequenceObjectHistory{}).Where("`increment_by` IN (?)", incrementBys).Find(&results).Error

	return
}

// GetFromStartWith 通过start_with获取内容
func (obj *_AllVirtualSequenceObjectHistoryMgr) GetFromStartWith(startWith float64) (results []*AllVirtualSequenceObjectHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSequenceObjectHistory{}).Where("`start_with` = ?", startWith).Find(&results).Error

	return
}

// GetBatchFromStartWith 批量查找
func (obj *_AllVirtualSequenceObjectHistoryMgr) GetBatchFromStartWith(startWiths []float64) (results []*AllVirtualSequenceObjectHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSequenceObjectHistory{}).Where("`start_with` IN (?)", startWiths).Find(&results).Error

	return
}

// GetFromCacheSize 通过cache_size获取内容
func (obj *_AllVirtualSequenceObjectHistoryMgr) GetFromCacheSize(cacheSize float64) (results []*AllVirtualSequenceObjectHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSequenceObjectHistory{}).Where("`cache_size` = ?", cacheSize).Find(&results).Error

	return
}

// GetBatchFromCacheSize 批量查找
func (obj *_AllVirtualSequenceObjectHistoryMgr) GetBatchFromCacheSize(cacheSizes []float64) (results []*AllVirtualSequenceObjectHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSequenceObjectHistory{}).Where("`cache_size` IN (?)", cacheSizes).Find(&results).Error

	return
}

// GetFromOrderFlag 通过order_flag获取内容
func (obj *_AllVirtualSequenceObjectHistoryMgr) GetFromOrderFlag(orderFlag int8) (results []*AllVirtualSequenceObjectHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSequenceObjectHistory{}).Where("`order_flag` = ?", orderFlag).Find(&results).Error

	return
}

// GetBatchFromOrderFlag 批量查找
func (obj *_AllVirtualSequenceObjectHistoryMgr) GetBatchFromOrderFlag(orderFlags []int8) (results []*AllVirtualSequenceObjectHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSequenceObjectHistory{}).Where("`order_flag` IN (?)", orderFlags).Find(&results).Error

	return
}

// GetFromCycleFlag 通过cycle_flag获取内容
func (obj *_AllVirtualSequenceObjectHistoryMgr) GetFromCycleFlag(cycleFlag int8) (results []*AllVirtualSequenceObjectHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSequenceObjectHistory{}).Where("`cycle_flag` = ?", cycleFlag).Find(&results).Error

	return
}

// GetBatchFromCycleFlag 批量查找
func (obj *_AllVirtualSequenceObjectHistoryMgr) GetBatchFromCycleFlag(cycleFlags []int8) (results []*AllVirtualSequenceObjectHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSequenceObjectHistory{}).Where("`cycle_flag` IN (?)", cycleFlags).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllVirtualSequenceObjectHistoryMgr) FetchByPrimaryKey(tenantID int64, sequenceID int64, schemaVersion int64) (result AllVirtualSequenceObjectHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSequenceObjectHistory{}).Where("`tenant_id` = ? AND `sequence_id` = ? AND `schema_version` = ?", tenantID, sequenceID, schemaVersion).First(&result).Error

	return
}
