package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
)	

type _V$sqlWorkareaHistogramMgr struct {
	*_BaseMgr
}

// V$sqlWorkareaHistogramMgr open func
func V$sqlWorkareaHistogramMgr(db *gorm.DB) *_V$sqlWorkareaHistogramMgr {
	if db == nil {
		panic(fmt.Errorf("V$sqlWorkareaHistogramMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_V$sqlWorkareaHistogramMgr{_BaseMgr: &_BaseMgr{DB: db.Table("v$sql_workarea_histogram"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_V$sqlWorkareaHistogramMgr) GetTableName() string {
	return "v$sql_workarea_histogram"
}

// Reset 重置gorm会话
func (obj *_V$sqlWorkareaHistogramMgr) Reset() *_V$sqlWorkareaHistogramMgr {
	obj.New()
	return obj
}

// Get 获取 
func (obj *_V$sqlWorkareaHistogramMgr) Get() (result V$sqlWorkareaHistogram, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlWorkareaHistogram{}).First(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_V$sqlWorkareaHistogramMgr) Gets() (results []*V$sqlWorkareaHistogram, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlWorkareaHistogram{}).Find(&results).Error
	
	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_V$sqlWorkareaHistogramMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(V$sqlWorkareaHistogram{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithLowOptimalSize low_optimal_size获取 
func (obj *_V$sqlWorkareaHistogramMgr) WithLowOptimalSize(lowOptimalSize int64) Option {
	return optionFunc(func(o *options) { o.query["low_optimal_size"] = lowOptimalSize })
}

// WithHighOptimalSize high_optimal_size获取 
func (obj *_V$sqlWorkareaHistogramMgr) WithHighOptimalSize(highOptimalSize int64) Option {
	return optionFunc(func(o *options) { o.query["high_optimal_size"] = highOptimalSize })
}

// WithOptimalExecutions optimal_executions获取 
func (obj *_V$sqlWorkareaHistogramMgr) WithOptimalExecutions(optimalExecutions int64) Option {
	return optionFunc(func(o *options) { o.query["optimal_executions"] = optimalExecutions })
}

// WithOnepassExecutions onepass_executions获取 
func (obj *_V$sqlWorkareaHistogramMgr) WithOnepassExecutions(onepassExecutions int64) Option {
	return optionFunc(func(o *options) { o.query["onepass_executions"] = onepassExecutions })
}

// WithMultipassesExecutions multipasses_executions获取 
func (obj *_V$sqlWorkareaHistogramMgr) WithMultipassesExecutions(multipassesExecutions int64) Option {
	return optionFunc(func(o *options) { o.query["multipasses_executions"] = multipassesExecutions })
}

// WithTotalExecutions total_executions获取 
func (obj *_V$sqlWorkareaHistogramMgr) WithTotalExecutions(totalExecutions int64) Option {
	return optionFunc(func(o *options) { o.query["total_executions"] = totalExecutions })
}

// WithConID con_id获取 
func (obj *_V$sqlWorkareaHistogramMgr) WithConID(conID int64) Option {
	return optionFunc(func(o *options) { o.query["con_id"] = conID })
}


// GetByOption 功能选项模式获取
func (obj *_V$sqlWorkareaHistogramMgr) GetByOption(opts ...Option) (result V$sqlWorkareaHistogram, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(V$sqlWorkareaHistogram{}).Where(options.query).First(&result).Error
	
	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_V$sqlWorkareaHistogramMgr) GetByOptions(opts ...Option) (results []*V$sqlWorkareaHistogram, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(V$sqlWorkareaHistogram{}).Where(options.query).Find(&results).Error
	
	return
}


// SelectPage 分页查询
func (obj *_V$sqlWorkareaHistogramMgr) SelectPage(page IPage,opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]V$sqlWorkareaHistogram,0)
	var count int64 // 统计总的记录数
	query :=  obj.DB.WithContext(obj.ctx).Model(V$sqlWorkareaHistogram{}).Where(options.query)
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


// GetFromLowOptimalSize 通过low_optimal_size获取内容  
func (obj *_V$sqlWorkareaHistogramMgr) GetFromLowOptimalSize(lowOptimalSize int64) (results []*V$sqlWorkareaHistogram, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlWorkareaHistogram{}).Where("`low_optimal_size` = ?", lowOptimalSize).Find(&results).Error
	
	return
}

// GetBatchFromLowOptimalSize 批量查找 
func (obj *_V$sqlWorkareaHistogramMgr) GetBatchFromLowOptimalSize(lowOptimalSizes []int64) (results []*V$sqlWorkareaHistogram, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlWorkareaHistogram{}).Where("`low_optimal_size` IN (?)", lowOptimalSizes).Find(&results).Error
	
	return
}
 
// GetFromHighOptimalSize 通过high_optimal_size获取内容  
func (obj *_V$sqlWorkareaHistogramMgr) GetFromHighOptimalSize(highOptimalSize int64) (results []*V$sqlWorkareaHistogram, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlWorkareaHistogram{}).Where("`high_optimal_size` = ?", highOptimalSize).Find(&results).Error
	
	return
}

// GetBatchFromHighOptimalSize 批量查找 
func (obj *_V$sqlWorkareaHistogramMgr) GetBatchFromHighOptimalSize(highOptimalSizes []int64) (results []*V$sqlWorkareaHistogram, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlWorkareaHistogram{}).Where("`high_optimal_size` IN (?)", highOptimalSizes).Find(&results).Error
	
	return
}
 
// GetFromOptimalExecutions 通过optimal_executions获取内容  
func (obj *_V$sqlWorkareaHistogramMgr) GetFromOptimalExecutions(optimalExecutions int64) (results []*V$sqlWorkareaHistogram, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlWorkareaHistogram{}).Where("`optimal_executions` = ?", optimalExecutions).Find(&results).Error
	
	return
}

// GetBatchFromOptimalExecutions 批量查找 
func (obj *_V$sqlWorkareaHistogramMgr) GetBatchFromOptimalExecutions(optimalExecutionss []int64) (results []*V$sqlWorkareaHistogram, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlWorkareaHistogram{}).Where("`optimal_executions` IN (?)", optimalExecutionss).Find(&results).Error
	
	return
}
 
// GetFromOnepassExecutions 通过onepass_executions获取内容  
func (obj *_V$sqlWorkareaHistogramMgr) GetFromOnepassExecutions(onepassExecutions int64) (results []*V$sqlWorkareaHistogram, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlWorkareaHistogram{}).Where("`onepass_executions` = ?", onepassExecutions).Find(&results).Error
	
	return
}

// GetBatchFromOnepassExecutions 批量查找 
func (obj *_V$sqlWorkareaHistogramMgr) GetBatchFromOnepassExecutions(onepassExecutionss []int64) (results []*V$sqlWorkareaHistogram, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlWorkareaHistogram{}).Where("`onepass_executions` IN (?)", onepassExecutionss).Find(&results).Error
	
	return
}
 
// GetFromMultipassesExecutions 通过multipasses_executions获取内容  
func (obj *_V$sqlWorkareaHistogramMgr) GetFromMultipassesExecutions(multipassesExecutions int64) (results []*V$sqlWorkareaHistogram, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlWorkareaHistogram{}).Where("`multipasses_executions` = ?", multipassesExecutions).Find(&results).Error
	
	return
}

// GetBatchFromMultipassesExecutions 批量查找 
func (obj *_V$sqlWorkareaHistogramMgr) GetBatchFromMultipassesExecutions(multipassesExecutionss []int64) (results []*V$sqlWorkareaHistogram, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlWorkareaHistogram{}).Where("`multipasses_executions` IN (?)", multipassesExecutionss).Find(&results).Error
	
	return
}
 
// GetFromTotalExecutions 通过total_executions获取内容  
func (obj *_V$sqlWorkareaHistogramMgr) GetFromTotalExecutions(totalExecutions int64) (results []*V$sqlWorkareaHistogram, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlWorkareaHistogram{}).Where("`total_executions` = ?", totalExecutions).Find(&results).Error
	
	return
}

// GetBatchFromTotalExecutions 批量查找 
func (obj *_V$sqlWorkareaHistogramMgr) GetBatchFromTotalExecutions(totalExecutionss []int64) (results []*V$sqlWorkareaHistogram, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlWorkareaHistogram{}).Where("`total_executions` IN (?)", totalExecutionss).Find(&results).Error
	
	return
}
 
// GetFromConID 通过con_id获取内容  
func (obj *_V$sqlWorkareaHistogramMgr) GetFromConID(conID int64) (results []*V$sqlWorkareaHistogram, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlWorkareaHistogram{}).Where("`con_id` = ?", conID).Find(&results).Error
	
	return
}

// GetBatchFromConID 批量查找 
func (obj *_V$sqlWorkareaHistogramMgr) GetBatchFromConID(conIDs []int64) (results []*V$sqlWorkareaHistogram, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$sqlWorkareaHistogram{}).Where("`con_id` IN (?)", conIDs).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 

 

	

