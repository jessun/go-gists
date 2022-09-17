package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
"time"	
)	

type _Gv$tenantSequenceObjectMgr struct {
	*_BaseMgr
}

// Gv$tenantSequenceObjectMgr open func
func Gv$tenantSequenceObjectMgr(db *gorm.DB) *_Gv$tenantSequenceObjectMgr {
	if db == nil {
		panic(fmt.Errorf("Gv$tenantSequenceObjectMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_Gv$tenantSequenceObjectMgr{_BaseMgr: &_BaseMgr{DB: db.Table("gv$tenant_sequence_object"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_Gv$tenantSequenceObjectMgr) GetTableName() string {
	return "gv$tenant_sequence_object"
}

// Reset 重置gorm会话
func (obj *_Gv$tenantSequenceObjectMgr) Reset() *_Gv$tenantSequenceObjectMgr {
	obj.New()
	return obj
}

// Get 获取 
func (obj *_Gv$tenantSequenceObjectMgr) Get() (result Gv$tenantSequenceObject, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$tenantSequenceObject{}).First(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_Gv$tenantSequenceObjectMgr) Gets() (results []*Gv$tenantSequenceObject, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$tenantSequenceObject{}).Find(&results).Error
	
	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_Gv$tenantSequenceObjectMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(Gv$tenantSequenceObject{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTenantID tenant_id获取 
func (obj *_Gv$tenantSequenceObjectMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithSequenceID sequence_id获取 
func (obj *_Gv$tenantSequenceObjectMgr) WithSequenceID(sequenceID int64) Option {
	return optionFunc(func(o *options) { o.query["sequence_id"] = sequenceID })
}

// WithGmtCreate gmt_create获取 
func (obj *_Gv$tenantSequenceObjectMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取 
func (obj *_Gv$tenantSequenceObjectMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithSchemaVersion schema_version获取 
func (obj *_Gv$tenantSequenceObjectMgr) WithSchemaVersion(schemaVersion int64) Option {
	return optionFunc(func(o *options) { o.query["schema_version"] = schemaVersion })
}

// WithDatabaseID database_id获取 
func (obj *_Gv$tenantSequenceObjectMgr) WithDatabaseID(databaseID int64) Option {
	return optionFunc(func(o *options) { o.query["database_id"] = databaseID })
}

// WithSequenceName sequence_name获取 
func (obj *_Gv$tenantSequenceObjectMgr) WithSequenceName(sequenceName string) Option {
	return optionFunc(func(o *options) { o.query["sequence_name"] = sequenceName })
}

// WithMinValue min_value获取 
func (obj *_Gv$tenantSequenceObjectMgr) WithMinValue(minValue float64) Option {
	return optionFunc(func(o *options) { o.query["min_value"] = minValue })
}

// WithMaxValue max_value获取 
func (obj *_Gv$tenantSequenceObjectMgr) WithMaxValue(maxValue float64) Option {
	return optionFunc(func(o *options) { o.query["max_value"] = maxValue })
}

// WithIncrementBy increment_by获取 
func (obj *_Gv$tenantSequenceObjectMgr) WithIncrementBy(incrementBy float64) Option {
	return optionFunc(func(o *options) { o.query["increment_by"] = incrementBy })
}

// WithStartWith start_with获取 
func (obj *_Gv$tenantSequenceObjectMgr) WithStartWith(startWith float64) Option {
	return optionFunc(func(o *options) { o.query["start_with"] = startWith })
}

// WithCacheSize cache_size获取 
func (obj *_Gv$tenantSequenceObjectMgr) WithCacheSize(cacheSize float64) Option {
	return optionFunc(func(o *options) { o.query["cache_size"] = cacheSize })
}

// WithOrderFlag order_flag获取 
func (obj *_Gv$tenantSequenceObjectMgr) WithOrderFlag(orderFlag int8) Option {
	return optionFunc(func(o *options) { o.query["order_flag"] = orderFlag })
}

// WithCycleFlag cycle_flag获取 
func (obj *_Gv$tenantSequenceObjectMgr) WithCycleFlag(cycleFlag int8) Option {
	return optionFunc(func(o *options) { o.query["cycle_flag"] = cycleFlag })
}


// GetByOption 功能选项模式获取
func (obj *_Gv$tenantSequenceObjectMgr) GetByOption(opts ...Option) (result Gv$tenantSequenceObject, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Gv$tenantSequenceObject{}).Where(options.query).First(&result).Error
	
	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_Gv$tenantSequenceObjectMgr) GetByOptions(opts ...Option) (results []*Gv$tenantSequenceObject, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Gv$tenantSequenceObject{}).Where(options.query).Find(&results).Error
	
	return
}


// SelectPage 分页查询
func (obj *_Gv$tenantSequenceObjectMgr) SelectPage(page IPage,opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]Gv$tenantSequenceObject,0)
	var count int64 // 统计总的记录数
	query :=  obj.DB.WithContext(obj.ctx).Model(Gv$tenantSequenceObject{}).Where(options.query)
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
func (obj *_Gv$tenantSequenceObjectMgr) GetFromTenantID(tenantID int64) (results []*Gv$tenantSequenceObject, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$tenantSequenceObject{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error
	
	return
}

// GetBatchFromTenantID 批量查找 
func (obj *_Gv$tenantSequenceObjectMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*Gv$tenantSequenceObject, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$tenantSequenceObject{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error
	
	return
}
 
// GetFromSequenceID 通过sequence_id获取内容  
func (obj *_Gv$tenantSequenceObjectMgr) GetFromSequenceID(sequenceID int64) (results []*Gv$tenantSequenceObject, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$tenantSequenceObject{}).Where("`sequence_id` = ?", sequenceID).Find(&results).Error
	
	return
}

// GetBatchFromSequenceID 批量查找 
func (obj *_Gv$tenantSequenceObjectMgr) GetBatchFromSequenceID(sequenceIDs []int64) (results []*Gv$tenantSequenceObject, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$tenantSequenceObject{}).Where("`sequence_id` IN (?)", sequenceIDs).Find(&results).Error
	
	return
}
 
// GetFromGmtCreate 通过gmt_create获取内容  
func (obj *_Gv$tenantSequenceObjectMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*Gv$tenantSequenceObject, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$tenantSequenceObject{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error
	
	return
}

// GetBatchFromGmtCreate 批量查找 
func (obj *_Gv$tenantSequenceObjectMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*Gv$tenantSequenceObject, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$tenantSequenceObject{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error
	
	return
}
 
// GetFromGmtModified 通过gmt_modified获取内容  
func (obj *_Gv$tenantSequenceObjectMgr) GetFromGmtModified(gmtModified time.Time) (results []*Gv$tenantSequenceObject, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$tenantSequenceObject{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error
	
	return
}

// GetBatchFromGmtModified 批量查找 
func (obj *_Gv$tenantSequenceObjectMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*Gv$tenantSequenceObject, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$tenantSequenceObject{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error
	
	return
}
 
// GetFromSchemaVersion 通过schema_version获取内容  
func (obj *_Gv$tenantSequenceObjectMgr) GetFromSchemaVersion(schemaVersion int64) (results []*Gv$tenantSequenceObject, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$tenantSequenceObject{}).Where("`schema_version` = ?", schemaVersion).Find(&results).Error
	
	return
}

// GetBatchFromSchemaVersion 批量查找 
func (obj *_Gv$tenantSequenceObjectMgr) GetBatchFromSchemaVersion(schemaVersions []int64) (results []*Gv$tenantSequenceObject, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$tenantSequenceObject{}).Where("`schema_version` IN (?)", schemaVersions).Find(&results).Error
	
	return
}
 
// GetFromDatabaseID 通过database_id获取内容  
func (obj *_Gv$tenantSequenceObjectMgr) GetFromDatabaseID(databaseID int64) (results []*Gv$tenantSequenceObject, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$tenantSequenceObject{}).Where("`database_id` = ?", databaseID).Find(&results).Error
	
	return
}

// GetBatchFromDatabaseID 批量查找 
func (obj *_Gv$tenantSequenceObjectMgr) GetBatchFromDatabaseID(databaseIDs []int64) (results []*Gv$tenantSequenceObject, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$tenantSequenceObject{}).Where("`database_id` IN (?)", databaseIDs).Find(&results).Error
	
	return
}
 
// GetFromSequenceName 通过sequence_name获取内容  
func (obj *_Gv$tenantSequenceObjectMgr) GetFromSequenceName(sequenceName string) (results []*Gv$tenantSequenceObject, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$tenantSequenceObject{}).Where("`sequence_name` = ?", sequenceName).Find(&results).Error
	
	return
}

// GetBatchFromSequenceName 批量查找 
func (obj *_Gv$tenantSequenceObjectMgr) GetBatchFromSequenceName(sequenceNames []string) (results []*Gv$tenantSequenceObject, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$tenantSequenceObject{}).Where("`sequence_name` IN (?)", sequenceNames).Find(&results).Error
	
	return
}
 
// GetFromMinValue 通过min_value获取内容  
func (obj *_Gv$tenantSequenceObjectMgr) GetFromMinValue(minValue float64) (results []*Gv$tenantSequenceObject, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$tenantSequenceObject{}).Where("`min_value` = ?", minValue).Find(&results).Error
	
	return
}

// GetBatchFromMinValue 批量查找 
func (obj *_Gv$tenantSequenceObjectMgr) GetBatchFromMinValue(minValues []float64) (results []*Gv$tenantSequenceObject, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$tenantSequenceObject{}).Where("`min_value` IN (?)", minValues).Find(&results).Error
	
	return
}
 
// GetFromMaxValue 通过max_value获取内容  
func (obj *_Gv$tenantSequenceObjectMgr) GetFromMaxValue(maxValue float64) (results []*Gv$tenantSequenceObject, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$tenantSequenceObject{}).Where("`max_value` = ?", maxValue).Find(&results).Error
	
	return
}

// GetBatchFromMaxValue 批量查找 
func (obj *_Gv$tenantSequenceObjectMgr) GetBatchFromMaxValue(maxValues []float64) (results []*Gv$tenantSequenceObject, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$tenantSequenceObject{}).Where("`max_value` IN (?)", maxValues).Find(&results).Error
	
	return
}
 
// GetFromIncrementBy 通过increment_by获取内容  
func (obj *_Gv$tenantSequenceObjectMgr) GetFromIncrementBy(incrementBy float64) (results []*Gv$tenantSequenceObject, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$tenantSequenceObject{}).Where("`increment_by` = ?", incrementBy).Find(&results).Error
	
	return
}

// GetBatchFromIncrementBy 批量查找 
func (obj *_Gv$tenantSequenceObjectMgr) GetBatchFromIncrementBy(incrementBys []float64) (results []*Gv$tenantSequenceObject, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$tenantSequenceObject{}).Where("`increment_by` IN (?)", incrementBys).Find(&results).Error
	
	return
}
 
// GetFromStartWith 通过start_with获取内容  
func (obj *_Gv$tenantSequenceObjectMgr) GetFromStartWith(startWith float64) (results []*Gv$tenantSequenceObject, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$tenantSequenceObject{}).Where("`start_with` = ?", startWith).Find(&results).Error
	
	return
}

// GetBatchFromStartWith 批量查找 
func (obj *_Gv$tenantSequenceObjectMgr) GetBatchFromStartWith(startWiths []float64) (results []*Gv$tenantSequenceObject, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$tenantSequenceObject{}).Where("`start_with` IN (?)", startWiths).Find(&results).Error
	
	return
}
 
// GetFromCacheSize 通过cache_size获取内容  
func (obj *_Gv$tenantSequenceObjectMgr) GetFromCacheSize(cacheSize float64) (results []*Gv$tenantSequenceObject, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$tenantSequenceObject{}).Where("`cache_size` = ?", cacheSize).Find(&results).Error
	
	return
}

// GetBatchFromCacheSize 批量查找 
func (obj *_Gv$tenantSequenceObjectMgr) GetBatchFromCacheSize(cacheSizes []float64) (results []*Gv$tenantSequenceObject, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$tenantSequenceObject{}).Where("`cache_size` IN (?)", cacheSizes).Find(&results).Error
	
	return
}
 
// GetFromOrderFlag 通过order_flag获取内容  
func (obj *_Gv$tenantSequenceObjectMgr) GetFromOrderFlag(orderFlag int8) (results []*Gv$tenantSequenceObject, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$tenantSequenceObject{}).Where("`order_flag` = ?", orderFlag).Find(&results).Error
	
	return
}

// GetBatchFromOrderFlag 批量查找 
func (obj *_Gv$tenantSequenceObjectMgr) GetBatchFromOrderFlag(orderFlags []int8) (results []*Gv$tenantSequenceObject, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$tenantSequenceObject{}).Where("`order_flag` IN (?)", orderFlags).Find(&results).Error
	
	return
}
 
// GetFromCycleFlag 通过cycle_flag获取内容  
func (obj *_Gv$tenantSequenceObjectMgr) GetFromCycleFlag(cycleFlag int8) (results []*Gv$tenantSequenceObject, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$tenantSequenceObject{}).Where("`cycle_flag` = ?", cycleFlag).Find(&results).Error
	
	return
}

// GetBatchFromCycleFlag 批量查找 
func (obj *_Gv$tenantSequenceObjectMgr) GetBatchFromCycleFlag(cycleFlags []int8) (results []*Gv$tenantSequenceObject, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$tenantSequenceObject{}).Where("`cycle_flag` IN (?)", cycleFlags).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 

 

	

