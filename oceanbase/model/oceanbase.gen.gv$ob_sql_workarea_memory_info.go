package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
)	

type _Gv$obSQLWorkareaMemoryInfoMgr struct {
	*_BaseMgr
}

// Gv$obSQLWorkareaMemoryInfoMgr open func
func Gv$obSQLWorkareaMemoryInfoMgr(db *gorm.DB) *_Gv$obSQLWorkareaMemoryInfoMgr {
	if db == nil {
		panic(fmt.Errorf("Gv$obSQLWorkareaMemoryInfoMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_Gv$obSQLWorkareaMemoryInfoMgr{_BaseMgr: &_BaseMgr{DB: db.Table("gv$ob_sql_workarea_memory_info"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_Gv$obSQLWorkareaMemoryInfoMgr) GetTableName() string {
	return "gv$ob_sql_workarea_memory_info"
}

// Reset 重置gorm会话
func (obj *_Gv$obSQLWorkareaMemoryInfoMgr) Reset() *_Gv$obSQLWorkareaMemoryInfoMgr {
	obj.New()
	return obj
}

// Get 获取 
func (obj *_Gv$obSQLWorkareaMemoryInfoMgr) Get() (result Gv$obSQLWorkareaMemoryInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$obSQLWorkareaMemoryInfo{}).First(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_Gv$obSQLWorkareaMemoryInfoMgr) Gets() (results []*Gv$obSQLWorkareaMemoryInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$obSQLWorkareaMemoryInfo{}).Find(&results).Error
	
	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_Gv$obSQLWorkareaMemoryInfoMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(Gv$obSQLWorkareaMemoryInfo{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithMaxWorkareaSize max_workarea_size获取 
func (obj *_Gv$obSQLWorkareaMemoryInfoMgr) WithMaxWorkareaSize(maxWorkareaSize int64) Option {
	return optionFunc(func(o *options) { o.query["max_workarea_size"] = maxWorkareaSize })
}

// WithWorkareaHoldSize workarea_hold_size获取 
func (obj *_Gv$obSQLWorkareaMemoryInfoMgr) WithWorkareaHoldSize(workareaHoldSize int64) Option {
	return optionFunc(func(o *options) { o.query["workarea_hold_size"] = workareaHoldSize })
}

// WithMaxAutoWorkareaSize max_auto_workarea_size获取 
func (obj *_Gv$obSQLWorkareaMemoryInfoMgr) WithMaxAutoWorkareaSize(maxAutoWorkareaSize int64) Option {
	return optionFunc(func(o *options) { o.query["max_auto_workarea_size"] = maxAutoWorkareaSize })
}

// WithMemTarget mem_target获取 
func (obj *_Gv$obSQLWorkareaMemoryInfoMgr) WithMemTarget(memTarget int64) Option {
	return optionFunc(func(o *options) { o.query["mem_target"] = memTarget })
}

// WithTotalMemUsed total_mem_used获取 
func (obj *_Gv$obSQLWorkareaMemoryInfoMgr) WithTotalMemUsed(totalMemUsed int64) Option {
	return optionFunc(func(o *options) { o.query["total_mem_used"] = totalMemUsed })
}

// WithGlobalMemBound global_mem_bound获取 
func (obj *_Gv$obSQLWorkareaMemoryInfoMgr) WithGlobalMemBound(globalMemBound int64) Option {
	return optionFunc(func(o *options) { o.query["global_mem_bound"] = globalMemBound })
}

// WithDriftSize drift_size获取 
func (obj *_Gv$obSQLWorkareaMemoryInfoMgr) WithDriftSize(driftSize int64) Option {
	return optionFunc(func(o *options) { o.query["drift_size"] = driftSize })
}

// WithWorkareaCount workarea_count获取 
func (obj *_Gv$obSQLWorkareaMemoryInfoMgr) WithWorkareaCount(workareaCount int64) Option {
	return optionFunc(func(o *options) { o.query["workarea_count"] = workareaCount })
}

// WithManualCalcCount manual_calc_count获取 
func (obj *_Gv$obSQLWorkareaMemoryInfoMgr) WithManualCalcCount(manualCalcCount int64) Option {
	return optionFunc(func(o *options) { o.query["manual_calc_count"] = manualCalcCount })
}


// GetByOption 功能选项模式获取
func (obj *_Gv$obSQLWorkareaMemoryInfoMgr) GetByOption(opts ...Option) (result Gv$obSQLWorkareaMemoryInfo, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Gv$obSQLWorkareaMemoryInfo{}).Where(options.query).First(&result).Error
	
	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_Gv$obSQLWorkareaMemoryInfoMgr) GetByOptions(opts ...Option) (results []*Gv$obSQLWorkareaMemoryInfo, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Gv$obSQLWorkareaMemoryInfo{}).Where(options.query).Find(&results).Error
	
	return
}


// SelectPage 分页查询
func (obj *_Gv$obSQLWorkareaMemoryInfoMgr) SelectPage(page IPage,opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]Gv$obSQLWorkareaMemoryInfo,0)
	var count int64 // 统计总的记录数
	query :=  obj.DB.WithContext(obj.ctx).Model(Gv$obSQLWorkareaMemoryInfo{}).Where(options.query)
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
func (obj *_Gv$obSQLWorkareaMemoryInfoMgr) GetFromMaxWorkareaSize(maxWorkareaSize int64) (results []*Gv$obSQLWorkareaMemoryInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$obSQLWorkareaMemoryInfo{}).Where("`max_workarea_size` = ?", maxWorkareaSize).Find(&results).Error
	
	return
}

// GetBatchFromMaxWorkareaSize 批量查找 
func (obj *_Gv$obSQLWorkareaMemoryInfoMgr) GetBatchFromMaxWorkareaSize(maxWorkareaSizes []int64) (results []*Gv$obSQLWorkareaMemoryInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$obSQLWorkareaMemoryInfo{}).Where("`max_workarea_size` IN (?)", maxWorkareaSizes).Find(&results).Error
	
	return
}
 
// GetFromWorkareaHoldSize 通过workarea_hold_size获取内容  
func (obj *_Gv$obSQLWorkareaMemoryInfoMgr) GetFromWorkareaHoldSize(workareaHoldSize int64) (results []*Gv$obSQLWorkareaMemoryInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$obSQLWorkareaMemoryInfo{}).Where("`workarea_hold_size` = ?", workareaHoldSize).Find(&results).Error
	
	return
}

// GetBatchFromWorkareaHoldSize 批量查找 
func (obj *_Gv$obSQLWorkareaMemoryInfoMgr) GetBatchFromWorkareaHoldSize(workareaHoldSizes []int64) (results []*Gv$obSQLWorkareaMemoryInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$obSQLWorkareaMemoryInfo{}).Where("`workarea_hold_size` IN (?)", workareaHoldSizes).Find(&results).Error
	
	return
}
 
// GetFromMaxAutoWorkareaSize 通过max_auto_workarea_size获取内容  
func (obj *_Gv$obSQLWorkareaMemoryInfoMgr) GetFromMaxAutoWorkareaSize(maxAutoWorkareaSize int64) (results []*Gv$obSQLWorkareaMemoryInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$obSQLWorkareaMemoryInfo{}).Where("`max_auto_workarea_size` = ?", maxAutoWorkareaSize).Find(&results).Error
	
	return
}

// GetBatchFromMaxAutoWorkareaSize 批量查找 
func (obj *_Gv$obSQLWorkareaMemoryInfoMgr) GetBatchFromMaxAutoWorkareaSize(maxAutoWorkareaSizes []int64) (results []*Gv$obSQLWorkareaMemoryInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$obSQLWorkareaMemoryInfo{}).Where("`max_auto_workarea_size` IN (?)", maxAutoWorkareaSizes).Find(&results).Error
	
	return
}
 
// GetFromMemTarget 通过mem_target获取内容  
func (obj *_Gv$obSQLWorkareaMemoryInfoMgr) GetFromMemTarget(memTarget int64) (results []*Gv$obSQLWorkareaMemoryInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$obSQLWorkareaMemoryInfo{}).Where("`mem_target` = ?", memTarget).Find(&results).Error
	
	return
}

// GetBatchFromMemTarget 批量查找 
func (obj *_Gv$obSQLWorkareaMemoryInfoMgr) GetBatchFromMemTarget(memTargets []int64) (results []*Gv$obSQLWorkareaMemoryInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$obSQLWorkareaMemoryInfo{}).Where("`mem_target` IN (?)", memTargets).Find(&results).Error
	
	return
}
 
// GetFromTotalMemUsed 通过total_mem_used获取内容  
func (obj *_Gv$obSQLWorkareaMemoryInfoMgr) GetFromTotalMemUsed(totalMemUsed int64) (results []*Gv$obSQLWorkareaMemoryInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$obSQLWorkareaMemoryInfo{}).Where("`total_mem_used` = ?", totalMemUsed).Find(&results).Error
	
	return
}

// GetBatchFromTotalMemUsed 批量查找 
func (obj *_Gv$obSQLWorkareaMemoryInfoMgr) GetBatchFromTotalMemUsed(totalMemUseds []int64) (results []*Gv$obSQLWorkareaMemoryInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$obSQLWorkareaMemoryInfo{}).Where("`total_mem_used` IN (?)", totalMemUseds).Find(&results).Error
	
	return
}
 
// GetFromGlobalMemBound 通过global_mem_bound获取内容  
func (obj *_Gv$obSQLWorkareaMemoryInfoMgr) GetFromGlobalMemBound(globalMemBound int64) (results []*Gv$obSQLWorkareaMemoryInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$obSQLWorkareaMemoryInfo{}).Where("`global_mem_bound` = ?", globalMemBound).Find(&results).Error
	
	return
}

// GetBatchFromGlobalMemBound 批量查找 
func (obj *_Gv$obSQLWorkareaMemoryInfoMgr) GetBatchFromGlobalMemBound(globalMemBounds []int64) (results []*Gv$obSQLWorkareaMemoryInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$obSQLWorkareaMemoryInfo{}).Where("`global_mem_bound` IN (?)", globalMemBounds).Find(&results).Error
	
	return
}
 
// GetFromDriftSize 通过drift_size获取内容  
func (obj *_Gv$obSQLWorkareaMemoryInfoMgr) GetFromDriftSize(driftSize int64) (results []*Gv$obSQLWorkareaMemoryInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$obSQLWorkareaMemoryInfo{}).Where("`drift_size` = ?", driftSize).Find(&results).Error
	
	return
}

// GetBatchFromDriftSize 批量查找 
func (obj *_Gv$obSQLWorkareaMemoryInfoMgr) GetBatchFromDriftSize(driftSizes []int64) (results []*Gv$obSQLWorkareaMemoryInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$obSQLWorkareaMemoryInfo{}).Where("`drift_size` IN (?)", driftSizes).Find(&results).Error
	
	return
}
 
// GetFromWorkareaCount 通过workarea_count获取内容  
func (obj *_Gv$obSQLWorkareaMemoryInfoMgr) GetFromWorkareaCount(workareaCount int64) (results []*Gv$obSQLWorkareaMemoryInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$obSQLWorkareaMemoryInfo{}).Where("`workarea_count` = ?", workareaCount).Find(&results).Error
	
	return
}

// GetBatchFromWorkareaCount 批量查找 
func (obj *_Gv$obSQLWorkareaMemoryInfoMgr) GetBatchFromWorkareaCount(workareaCounts []int64) (results []*Gv$obSQLWorkareaMemoryInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$obSQLWorkareaMemoryInfo{}).Where("`workarea_count` IN (?)", workareaCounts).Find(&results).Error
	
	return
}
 
// GetFromManualCalcCount 通过manual_calc_count获取内容  
func (obj *_Gv$obSQLWorkareaMemoryInfoMgr) GetFromManualCalcCount(manualCalcCount int64) (results []*Gv$obSQLWorkareaMemoryInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$obSQLWorkareaMemoryInfo{}).Where("`manual_calc_count` = ?", manualCalcCount).Find(&results).Error
	
	return
}

// GetBatchFromManualCalcCount 批量查找 
func (obj *_Gv$obSQLWorkareaMemoryInfoMgr) GetBatchFromManualCalcCount(manualCalcCounts []int64) (results []*Gv$obSQLWorkareaMemoryInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$obSQLWorkareaMemoryInfo{}).Where("`manual_calc_count` IN (?)", manualCalcCounts).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 

 

	

