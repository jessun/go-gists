package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _AllSequenceObjectHistoryMgr struct {
	*_BaseMgr
}

// AllSequenceObjectHistoryMgr open func
func AllSequenceObjectHistoryMgr(db *gorm.DB) *_AllSequenceObjectHistoryMgr {
	if db == nil {
		panic(fmt.Errorf("AllSequenceObjectHistoryMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllSequenceObjectHistoryMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_sequence_object_history"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllSequenceObjectHistoryMgr) GetTableName() string {
	return "__all_sequence_object_history"
}

// Reset 重置gorm会话
func (obj *_AllSequenceObjectHistoryMgr) Reset() *_AllSequenceObjectHistoryMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllSequenceObjectHistoryMgr) Get() (result AllSequenceObjectHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSequenceObjectHistory{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllSequenceObjectHistoryMgr) Gets() (results []*AllSequenceObjectHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSequenceObjectHistory{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllSequenceObjectHistoryMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllSequenceObjectHistory{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithGmtCreate gmt_create获取
func (obj *_AllSequenceObjectHistoryMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllSequenceObjectHistoryMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithTenantID tenant_id获取
func (obj *_AllSequenceObjectHistoryMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithSequenceID sequence_id获取
func (obj *_AllSequenceObjectHistoryMgr) WithSequenceID(sequenceID int64) Option {
	return optionFunc(func(o *options) { o.query["sequence_id"] = sequenceID })
}

// WithSchemaVersion schema_version获取
func (obj *_AllSequenceObjectHistoryMgr) WithSchemaVersion(schemaVersion int64) Option {
	return optionFunc(func(o *options) { o.query["schema_version"] = schemaVersion })
}

// WithIsDeleted is_deleted获取
func (obj *_AllSequenceObjectHistoryMgr) WithIsDeleted(isDeleted int64) Option {
	return optionFunc(func(o *options) { o.query["is_deleted"] = isDeleted })
}

// WithDatabaseID database_id获取
func (obj *_AllSequenceObjectHistoryMgr) WithDatabaseID(databaseID int64) Option {
	return optionFunc(func(o *options) { o.query["database_id"] = databaseID })
}

// WithSequenceName sequence_name获取
func (obj *_AllSequenceObjectHistoryMgr) WithSequenceName(sequenceName string) Option {
	return optionFunc(func(o *options) { o.query["sequence_name"] = sequenceName })
}

// WithMinValue min_value获取
func (obj *_AllSequenceObjectHistoryMgr) WithMinValue(minValue float64) Option {
	return optionFunc(func(o *options) { o.query["min_value"] = minValue })
}

// WithMaxValue max_value获取
func (obj *_AllSequenceObjectHistoryMgr) WithMaxValue(maxValue float64) Option {
	return optionFunc(func(o *options) { o.query["max_value"] = maxValue })
}

// WithIncrementBy increment_by获取
func (obj *_AllSequenceObjectHistoryMgr) WithIncrementBy(incrementBy float64) Option {
	return optionFunc(func(o *options) { o.query["increment_by"] = incrementBy })
}

// WithStartWith start_with获取
func (obj *_AllSequenceObjectHistoryMgr) WithStartWith(startWith float64) Option {
	return optionFunc(func(o *options) { o.query["start_with"] = startWith })
}

// WithCacheSize cache_size获取
func (obj *_AllSequenceObjectHistoryMgr) WithCacheSize(cacheSize float64) Option {
	return optionFunc(func(o *options) { o.query["cache_size"] = cacheSize })
}

// WithOrderFlag order_flag获取
func (obj *_AllSequenceObjectHistoryMgr) WithOrderFlag(orderFlag int8) Option {
	return optionFunc(func(o *options) { o.query["order_flag"] = orderFlag })
}

// WithCycleFlag cycle_flag获取
func (obj *_AllSequenceObjectHistoryMgr) WithCycleFlag(cycleFlag int8) Option {
	return optionFunc(func(o *options) { o.query["cycle_flag"] = cycleFlag })
}

// GetByOption 功能选项模式获取
func (obj *_AllSequenceObjectHistoryMgr) GetByOption(opts ...Option) (result AllSequenceObjectHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllSequenceObjectHistory{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllSequenceObjectHistoryMgr) GetByOptions(opts ...Option) (results []*AllSequenceObjectHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllSequenceObjectHistory{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllSequenceObjectHistoryMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllSequenceObjectHistory, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllSequenceObjectHistory{}).Where(options.query)
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
func (obj *_AllSequenceObjectHistoryMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllSequenceObjectHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSequenceObjectHistory{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllSequenceObjectHistoryMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllSequenceObjectHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSequenceObjectHistory{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllSequenceObjectHistoryMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllSequenceObjectHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSequenceObjectHistory{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllSequenceObjectHistoryMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllSequenceObjectHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSequenceObjectHistory{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromTenantID 通过tenant_id获取内容
func (obj *_AllSequenceObjectHistoryMgr) GetFromTenantID(tenantID int64) (results []*AllSequenceObjectHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSequenceObjectHistory{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllSequenceObjectHistoryMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllSequenceObjectHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSequenceObjectHistory{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromSequenceID 通过sequence_id获取内容
func (obj *_AllSequenceObjectHistoryMgr) GetFromSequenceID(sequenceID int64) (results []*AllSequenceObjectHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSequenceObjectHistory{}).Where("`sequence_id` = ?", sequenceID).Find(&results).Error

	return
}

// GetBatchFromSequenceID 批量查找
func (obj *_AllSequenceObjectHistoryMgr) GetBatchFromSequenceID(sequenceIDs []int64) (results []*AllSequenceObjectHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSequenceObjectHistory{}).Where("`sequence_id` IN (?)", sequenceIDs).Find(&results).Error

	return
}

// GetFromSchemaVersion 通过schema_version获取内容
func (obj *_AllSequenceObjectHistoryMgr) GetFromSchemaVersion(schemaVersion int64) (results []*AllSequenceObjectHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSequenceObjectHistory{}).Where("`schema_version` = ?", schemaVersion).Find(&results).Error

	return
}

// GetBatchFromSchemaVersion 批量查找
func (obj *_AllSequenceObjectHistoryMgr) GetBatchFromSchemaVersion(schemaVersions []int64) (results []*AllSequenceObjectHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSequenceObjectHistory{}).Where("`schema_version` IN (?)", schemaVersions).Find(&results).Error

	return
}

// GetFromIsDeleted 通过is_deleted获取内容
func (obj *_AllSequenceObjectHistoryMgr) GetFromIsDeleted(isDeleted int64) (results []*AllSequenceObjectHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSequenceObjectHistory{}).Where("`is_deleted` = ?", isDeleted).Find(&results).Error

	return
}

// GetBatchFromIsDeleted 批量查找
func (obj *_AllSequenceObjectHistoryMgr) GetBatchFromIsDeleted(isDeleteds []int64) (results []*AllSequenceObjectHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSequenceObjectHistory{}).Where("`is_deleted` IN (?)", isDeleteds).Find(&results).Error

	return
}

// GetFromDatabaseID 通过database_id获取内容
func (obj *_AllSequenceObjectHistoryMgr) GetFromDatabaseID(databaseID int64) (results []*AllSequenceObjectHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSequenceObjectHistory{}).Where("`database_id` = ?", databaseID).Find(&results).Error

	return
}

// GetBatchFromDatabaseID 批量查找
func (obj *_AllSequenceObjectHistoryMgr) GetBatchFromDatabaseID(databaseIDs []int64) (results []*AllSequenceObjectHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSequenceObjectHistory{}).Where("`database_id` IN (?)", databaseIDs).Find(&results).Error

	return
}

// GetFromSequenceName 通过sequence_name获取内容
func (obj *_AllSequenceObjectHistoryMgr) GetFromSequenceName(sequenceName string) (results []*AllSequenceObjectHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSequenceObjectHistory{}).Where("`sequence_name` = ?", sequenceName).Find(&results).Error

	return
}

// GetBatchFromSequenceName 批量查找
func (obj *_AllSequenceObjectHistoryMgr) GetBatchFromSequenceName(sequenceNames []string) (results []*AllSequenceObjectHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSequenceObjectHistory{}).Where("`sequence_name` IN (?)", sequenceNames).Find(&results).Error

	return
}

// GetFromMinValue 通过min_value获取内容
func (obj *_AllSequenceObjectHistoryMgr) GetFromMinValue(minValue float64) (results []*AllSequenceObjectHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSequenceObjectHistory{}).Where("`min_value` = ?", minValue).Find(&results).Error

	return
}

// GetBatchFromMinValue 批量查找
func (obj *_AllSequenceObjectHistoryMgr) GetBatchFromMinValue(minValues []float64) (results []*AllSequenceObjectHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSequenceObjectHistory{}).Where("`min_value` IN (?)", minValues).Find(&results).Error

	return
}

// GetFromMaxValue 通过max_value获取内容
func (obj *_AllSequenceObjectHistoryMgr) GetFromMaxValue(maxValue float64) (results []*AllSequenceObjectHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSequenceObjectHistory{}).Where("`max_value` = ?", maxValue).Find(&results).Error

	return
}

// GetBatchFromMaxValue 批量查找
func (obj *_AllSequenceObjectHistoryMgr) GetBatchFromMaxValue(maxValues []float64) (results []*AllSequenceObjectHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSequenceObjectHistory{}).Where("`max_value` IN (?)", maxValues).Find(&results).Error

	return
}

// GetFromIncrementBy 通过increment_by获取内容
func (obj *_AllSequenceObjectHistoryMgr) GetFromIncrementBy(incrementBy float64) (results []*AllSequenceObjectHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSequenceObjectHistory{}).Where("`increment_by` = ?", incrementBy).Find(&results).Error

	return
}

// GetBatchFromIncrementBy 批量查找
func (obj *_AllSequenceObjectHistoryMgr) GetBatchFromIncrementBy(incrementBys []float64) (results []*AllSequenceObjectHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSequenceObjectHistory{}).Where("`increment_by` IN (?)", incrementBys).Find(&results).Error

	return
}

// GetFromStartWith 通过start_with获取内容
func (obj *_AllSequenceObjectHistoryMgr) GetFromStartWith(startWith float64) (results []*AllSequenceObjectHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSequenceObjectHistory{}).Where("`start_with` = ?", startWith).Find(&results).Error

	return
}

// GetBatchFromStartWith 批量查找
func (obj *_AllSequenceObjectHistoryMgr) GetBatchFromStartWith(startWiths []float64) (results []*AllSequenceObjectHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSequenceObjectHistory{}).Where("`start_with` IN (?)", startWiths).Find(&results).Error

	return
}

// GetFromCacheSize 通过cache_size获取内容
func (obj *_AllSequenceObjectHistoryMgr) GetFromCacheSize(cacheSize float64) (results []*AllSequenceObjectHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSequenceObjectHistory{}).Where("`cache_size` = ?", cacheSize).Find(&results).Error

	return
}

// GetBatchFromCacheSize 批量查找
func (obj *_AllSequenceObjectHistoryMgr) GetBatchFromCacheSize(cacheSizes []float64) (results []*AllSequenceObjectHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSequenceObjectHistory{}).Where("`cache_size` IN (?)", cacheSizes).Find(&results).Error

	return
}

// GetFromOrderFlag 通过order_flag获取内容
func (obj *_AllSequenceObjectHistoryMgr) GetFromOrderFlag(orderFlag int8) (results []*AllSequenceObjectHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSequenceObjectHistory{}).Where("`order_flag` = ?", orderFlag).Find(&results).Error

	return
}

// GetBatchFromOrderFlag 批量查找
func (obj *_AllSequenceObjectHistoryMgr) GetBatchFromOrderFlag(orderFlags []int8) (results []*AllSequenceObjectHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSequenceObjectHistory{}).Where("`order_flag` IN (?)", orderFlags).Find(&results).Error

	return
}

// GetFromCycleFlag 通过cycle_flag获取内容
func (obj *_AllSequenceObjectHistoryMgr) GetFromCycleFlag(cycleFlag int8) (results []*AllSequenceObjectHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSequenceObjectHistory{}).Where("`cycle_flag` = ?", cycleFlag).Find(&results).Error

	return
}

// GetBatchFromCycleFlag 批量查找
func (obj *_AllSequenceObjectHistoryMgr) GetBatchFromCycleFlag(cycleFlags []int8) (results []*AllSequenceObjectHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSequenceObjectHistory{}).Where("`cycle_flag` IN (?)", cycleFlags).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllSequenceObjectHistoryMgr) FetchByPrimaryKey(tenantID int64, sequenceID int64, schemaVersion int64) (result AllSequenceObjectHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSequenceObjectHistory{}).Where("`tenant_id` = ? AND `sequence_id` = ? AND `schema_version` = ?", tenantID, sequenceID, schemaVersion).First(&result).Error

	return
}
