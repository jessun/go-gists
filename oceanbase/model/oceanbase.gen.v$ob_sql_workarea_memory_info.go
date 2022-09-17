package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
)	

type _V$obSQLWorkareaMemoryInfoMgr struct {
	*_BaseMgr
}

// V$obSQLWorkareaMemoryInfoMgr open func
func V$obSQLWorkareaMemoryInfoMgr(db *gorm.DB) *_V$obSQLWorkareaMemoryInfoMgr {
	if db == nil {
		panic(fmt.Errorf("V$obSQLWorkareaMemoryInfoMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_V$obSQLWorkareaMemoryInfoMgr{_BaseMgr: &_BaseMgr{DB: db.Table("v$ob_sql_workarea_memory_info"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_V$obSQLWorkareaMemoryInfoMgr) GetTableName() string {
	return "v$ob_sql_workarea_memory_info"
}

// Reset 重置gorm会话
func (obj *_V$obSQLWorkareaMemoryInfoMgr) Reset() *_V$obSQLWorkareaMemoryInfoMgr {
	obj.New()
	return obj
}

// Get 获取 
func (obj *_V$obSQLWorkareaMemoryInfoMgr) Get() (result V$obSQLWorkareaMemoryInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$obSQLWorkareaMemoryInfo{}).First(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_V$obSQLWorkareaMemoryInfoMgr) Gets() (results []*V$obSQLWorkareaMemoryInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$obSQLWorkareaMemoryInfo{}).Find(&results).Error
	
	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_V$obSQLWorkareaMemoryInfoMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(V$obSQLWorkareaMemoryInfo{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithMaxWorkareaSize max_workarea_size获取 
func (obj *_V$obSQLWorkareaMemoryInfoMgr) WithMaxWorkareaSize(maxWorkareaSize int64) Option {
	return optionFunc(func(o *options) { o.query["max_workarea_size"] = maxWorkareaSize })
}

// WithWorkareaHoldSize workarea_hold_size获取 
func (obj *_V$obSQLWorkareaMemoryInfoMgr) WithWorkareaHoldSize(workareaHoldSize int64) Option {
	return optionFunc(func(o *options) { o.query["workarea_hold_size"] = workareaHoldSize })
}

// WithMaxAutoWorkareaSize max_auto_workarea_size获取 
func (obj *_V$obSQLWorkareaMemoryInfoMgr) WithMaxAutoWorkareaSize(maxAutoWorkareaSize int64) Option {
	return optionFunc(func(o *options) { o.query["max_auto_workarea_size"] = maxAutoWorkareaSize })
}

// WithMemTarget mem_target获取 
func (obj *_V$obSQLWorkareaMemoryInfoMgr) WithMemTarget(memTarget int64) Option {
	return optionFunc(func(o *options) { o.query["mem_target"] = memTarget })
}

// WithTotalMemUsed total_mem_used获取 
func (obj *_V$obSQLWorkareaMemoryInfoMgr) WithTotalMemUsed(totalMemUsed int64) Option {
	return optionFunc(func(o *options) { o.query["total_mem_used"] = totalMemUsed })
}

// WithGlobalMemBound global_mem_bound获取 
func (obj *_V$obSQLWorkareaMemoryInfoMgr) WithGlobalMemBound(globalMemBound int64) Option {
	return optionFunc(func(o *options) { o.query["global_mem_bound"] = globalMemBound })
}

// WithDriftSize drift_size获取 
func (obj *_V$obSQLWorkareaMemoryInfoMgr) WithDriftSize(driftSize int64) Option {
	return optionFunc(func(o *options) { o.query["drift_size"] = driftSize })
}

// WithWorkareaCount workarea_count获取 
func (obj *_V$obSQLWorkareaMemoryInfoMgr) WithWorkareaCount(workareaCount int64) Option {
	return optionFunc(func(o *options) { o.query["workarea_count"] = workareaCount })
}

// WithManualCalcCount manual_calc_count获取 
func (obj *_V$obSQLWorkareaMemoryInfoMgr) WithManualCalcCount(manualCalcCount int64) Option {
	return optionFunc(func(o *options) { o.query["manual_calc_count"] = manualCalcCount })
}


// GetByOption 功能选项模式获取
func (obj *_V$obSQLWorkareaMemoryInfoMgr) GetByOption(opts ...Option) (result V$obSQLWorkareaMemoryInfo, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(V$obSQLWorkareaMemoryInfo{}).Where(options.query).First(&result).Error
	
	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_V$obSQLWorkareaMemoryInfoMgr) GetByOptions(opts ...Option) (results []*V$obSQLWorkareaMemoryInfo, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(V$obSQLWorkareaMemoryInfo{}).Where(options.query).Find(&results).Error
	
	return
}


// SelectPage 分页查询
func (obj *_V$obSQLWorkareaMemoryInfoMgr) SelectPage(page IPage,opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]V$obSQLWorkareaMemoryInfo,0)
	var count int64 // 统计总的记录数
	query :=  obj.DB.WithContext(obj.ctx).Model(V$obSQLWorkareaMemoryInfo{}).Where(options.query)
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


// GetFromMaxWorkareaSize 通过max_workarea_size获取内容  
func (obj *_V$obSQLWorkareaMemoryInfoMgr) GetFromMaxWorkareaSize(maxWorkareaSize int64) (results []*V$obSQLWorkareaMemoryInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$obSQLWorkareaMemoryInfo{}).Where("`max_workarea_size` = ?", maxWorkareaSize).Find(&results).Error
	
	return
}

// GetBatchFromMaxWorkareaSize 批量查找 
func (obj *_V$obSQLWorkareaMemoryInfoMgr) GetBatchFromMaxWorkareaSize(maxWorkareaSizes []int64) (results []*V$obSQLWorkareaMemoryInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$obSQLWorkareaMemoryInfo{}).Where("`max_workarea_size` IN (?)", maxWorkareaSizes).Find(&results).Error
	
	return
}
 
// GetFromWorkareaHoldSize 通过workarea_hold_size获取内容  
func (obj *_V$obSQLWorkareaMemoryInfoMgr) GetFromWorkareaHoldSize(workareaHoldSize int64) (results []*V$obSQLWorkareaMemoryInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$obSQLWorkareaMemoryInfo{}).Where("`workarea_hold_size` = ?", workareaHoldSize).Find(&results).Error
	
	return
}

// GetBatchFromWorkareaHoldSize 批量查找 
func (obj *_V$obSQLWorkareaMemoryInfoMgr) GetBatchFromWorkareaHoldSize(workareaHoldSizes []int64) (results []*V$obSQLWorkareaMemoryInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$obSQLWorkareaMemoryInfo{}).Where("`workarea_hold_size` IN (?)", workareaHoldSizes).Find(&results).Error
	
	return
}
 
// GetFromMaxAutoWorkareaSize 通过max_auto_workarea_size获取内容  
func (obj *_V$obSQLWorkareaMemoryInfoMgr) GetFromMaxAutoWorkareaSize(maxAutoWorkareaSize int64) (results []*V$obSQLWorkareaMemoryInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$obSQLWorkareaMemoryInfo{}).Where("`max_auto_workarea_size` = ?", maxAutoWorkareaSize).Find(&results).Error
	
	return
}

// GetBatchFromMaxAutoWorkareaSize 批量查找 
func (obj *_V$obSQLWorkareaMemoryInfoMgr) GetBatchFromMaxAutoWorkareaSize(maxAutoWorkareaSizes []int64) (results []*V$obSQLWorkareaMemoryInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$obSQLWorkareaMemoryInfo{}).Where("`max_auto_workarea_size` IN (?)", maxAutoWorkareaSizes).Find(&results).Error
	
	return
}
 
// GetFromMemTarget 通过mem_target获取内容  
func (obj *_V$obSQLWorkareaMemoryInfoMgr) GetFromMemTarget(memTarget int64) (results []*V$obSQLWorkareaMemoryInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$obSQLWorkareaMemoryInfo{}).Where("`mem_target` = ?", memTarget).Find(&results).Error
	
	return
}

// GetBatchFromMemTarget 批量查找 
func (obj *_V$obSQLWorkareaMemoryInfoMgr) GetBatchFromMemTarget(memTargets []int64) (results []*V$obSQLWorkareaMemoryInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$obSQLWorkareaMemoryInfo{}).Where("`mem_target` IN (?)", memTargets).Find(&results).Error
	
	return
}
 
// GetFromTotalMemUsed 通过total_mem_used获取内容  
func (obj *_V$obSQLWorkareaMemoryInfoMgr) GetFromTotalMemUsed(totalMemUsed int64) (results []*V$obSQLWorkareaMemoryInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$obSQLWorkareaMemoryInfo{}).Where("`total_mem_used` = ?", totalMemUsed).Find(&results).Error
	
	return
}

// GetBatchFromTotalMemUsed 批量查找 
func (obj *_V$obSQLWorkareaMemoryInfoMgr) GetBatchFromTotalMemUsed(totalMemUseds []int64) (results []*V$obSQLWorkareaMemoryInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$obSQLWorkareaMemoryInfo{}).Where("`total_mem_used` IN (?)", totalMemUseds).Find(&results).Error
	
	return
}
 
// GetFromGlobalMemBound 通过global_mem_bound获取内容  
func (obj *_V$obSQLWorkareaMemoryInfoMgr) GetFromGlobalMemBound(globalMemBound int64) (results []*V$obSQLWorkareaMemoryInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$obSQLWorkareaMemoryInfo{}).Where("`global_mem_bound` = ?", globalMemBound).Find(&results).Error
	
	return
}

// GetBatchFromGlobalMemBound 批量查找 
func (obj *_V$obSQLWorkareaMemoryInfoMgr) GetBatchFromGlobalMemBound(globalMemBounds []int64) (results []*V$obSQLWorkareaMemoryInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$obSQLWorkareaMemoryInfo{}).Where("`global_mem_bound` IN (?)", globalMemBounds).Find(&results).Error
	
	return
}
 
// GetFromDriftSize 通过drift_size获取内容  
func (obj *_V$obSQLWorkareaMemoryInfoMgr) GetFromDriftSize(driftSize int64) (results []*V$obSQLWorkareaMemoryInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$obSQLWorkareaMemoryInfo{}).Where("`drift_size` = ?", driftSize).Find(&results).Error
	
	return
}

// GetBatchFromDriftSize 批量查找 
func (obj *_V$obSQLWorkareaMemoryInfoMgr) GetBatchFromDriftSize(driftSizes []int64) (results []*V$obSQLWorkareaMemoryInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$obSQLWorkareaMemoryInfo{}).Where("`drift_size` IN (?)", driftSizes).Find(&results).Error
	
	return
}
 
// GetFromWorkareaCount 通过workarea_count获取内容  
func (obj *_V$obSQLWorkareaMemoryInfoMgr) GetFromWorkareaCount(workareaCount int64) (results []*V$obSQLWorkareaMemoryInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$obSQLWorkareaMemoryInfo{}).Where("`workarea_count` = ?", workareaCount).Find(&results).Error
	
	return
}

// GetBatchFromWorkareaCount 批量查找 
func (obj *_V$obSQLWorkareaMemoryInfoMgr) GetBatchFromWorkareaCount(workareaCounts []int64) (results []*V$obSQLWorkareaMemoryInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$obSQLWorkareaMemoryInfo{}).Where("`workarea_count` IN (?)", workareaCounts).Find(&results).Error
	
	return
}
 
// GetFromManualCalcCount 通过manual_calc_count获取内容  
func (obj *_V$obSQLWorkareaMemoryInfoMgr) GetFromManualCalcCount(manualCalcCount int64) (results []*V$obSQLWorkareaMemoryInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$obSQLWorkareaMemoryInfo{}).Where("`manual_calc_count` = ?", manualCalcCount).Find(&results).Error
	
	return
}

// GetBatchFromManualCalcCount 批量查找 
func (obj *_V$obSQLWorkareaMemoryInfoMgr) GetBatchFromManualCalcCount(manualCalcCounts []int64) (results []*V$obSQLWorkareaMemoryInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$obSQLWorkareaMemoryInfo{}).Where("`manual_calc_count` IN (?)", manualCalcCounts).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 

 

	

