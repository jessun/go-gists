package	model	
import (	
"gorm.io/gorm"	
"fmt"	
"context"	
)	

type _Gv$sqlWorkareaMgr struct {
	*_BaseMgr
}

// Gv$sqlWorkareaMgr open func
func Gv$sqlWorkareaMgr(db *gorm.DB) *_Gv$sqlWorkareaMgr {
	if db == nil {
		panic(fmt.Errorf("Gv$sqlWorkareaMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_Gv$sqlWorkareaMgr{_BaseMgr: &_BaseMgr{DB: db.Table("gv$sql_workarea"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_Gv$sqlWorkareaMgr) GetTableName() string {
	return "gv$sql_workarea"
}

// Reset 重置gorm会话
func (obj *_Gv$sqlWorkareaMgr) Reset() *_Gv$sqlWorkareaMgr {
	obj.New()
	return obj
}

// Get 获取 
func (obj *_Gv$sqlWorkareaMgr) Get() (result Gv$sqlWorkarea, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlWorkarea{}).First(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_Gv$sqlWorkareaMgr) Gets() (results []*Gv$sqlWorkarea, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlWorkarea{}).Find(&results).Error
	
	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_Gv$sqlWorkareaMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(Gv$sqlWorkarea{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithAddress address获取 
func (obj *_Gv$sqlWorkareaMgr) WithAddress(address []byte) Option {
	return optionFunc(func(o *options) { o.query["address"] = address })
}

// WithHashValue hash_value获取 
func (obj *_Gv$sqlWorkareaMgr) WithHashValue(hashValue int64) Option {
	return optionFunc(func(o *options) { o.query["hash_value"] = hashValue })
}

// WithSQLID sql_id获取 
func (obj *_Gv$sqlWorkareaMgr) WithSQLID(sqlID string) Option {
	return optionFunc(func(o *options) { o.query["sql_id"] = sqlID })
}

// WithChildNumber child_number获取 
func (obj *_Gv$sqlWorkareaMgr) WithChildNumber(childNumber int64) Option {
	return optionFunc(func(o *options) { o.query["child_number"] = childNumber })
}

// WithWorkareaAddress workarea_address获取 
func (obj *_Gv$sqlWorkareaMgr) WithWorkareaAddress(workareaAddress []byte) Option {
	return optionFunc(func(o *options) { o.query["workarea_address"] = workareaAddress })
}

// WithOperationType operation_type获取 
func (obj *_Gv$sqlWorkareaMgr) WithOperationType(operationType string) Option {
	return optionFunc(func(o *options) { o.query["operation_type"] = operationType })
}

// WithOperationID operation_id获取 
func (obj *_Gv$sqlWorkareaMgr) WithOperationID(operationID int64) Option {
	return optionFunc(func(o *options) { o.query["operation_id"] = operationID })
}

// WithPolicy policy获取 
func (obj *_Gv$sqlWorkareaMgr) WithPolicy(policy string) Option {
	return optionFunc(func(o *options) { o.query["policy"] = policy })
}

// WithEstimatedOptimalSize estimated_optimal_size获取 
func (obj *_Gv$sqlWorkareaMgr) WithEstimatedOptimalSize(estimatedOptimalSize int64) Option {
	return optionFunc(func(o *options) { o.query["estimated_optimal_size"] = estimatedOptimalSize })
}

// WithEstimatedOnepassSize estimated_onepass_size获取 
func (obj *_Gv$sqlWorkareaMgr) WithEstimatedOnepassSize(estimatedOnepassSize int64) Option {
	return optionFunc(func(o *options) { o.query["estimated_onepass_size"] = estimatedOnepassSize })
}

// WithLastMemoryUsed last_memory_used获取 
func (obj *_Gv$sqlWorkareaMgr) WithLastMemoryUsed(lastMemoryUsed int64) Option {
	return optionFunc(func(o *options) { o.query["last_memory_used"] = lastMemoryUsed })
}

// WithLastExecution last_execution获取 
func (obj *_Gv$sqlWorkareaMgr) WithLastExecution(lastExecution string) Option {
	return optionFunc(func(o *options) { o.query["last_execution"] = lastExecution })
}

// WithLastDegree last_degree获取 
func (obj *_Gv$sqlWorkareaMgr) WithLastDegree(lastDegree int64) Option {
	return optionFunc(func(o *options) { o.query["last_degree"] = lastDegree })
}

// WithTotalExecutions total_executions获取 
func (obj *_Gv$sqlWorkareaMgr) WithTotalExecutions(totalExecutions int64) Option {
	return optionFunc(func(o *options) { o.query["total_executions"] = totalExecutions })
}

// WithOptimalExecutions optimal_executions获取 
func (obj *_Gv$sqlWorkareaMgr) WithOptimalExecutions(optimalExecutions int64) Option {
	return optionFunc(func(o *options) { o.query["optimal_executions"] = optimalExecutions })
}

// WithOnepassExecutions onepass_executions获取 
func (obj *_Gv$sqlWorkareaMgr) WithOnepassExecutions(onepassExecutions int64) Option {
	return optionFunc(func(o *options) { o.query["onepass_executions"] = onepassExecutions })
}

// WithMultipassesExecutions multipasses_executions获取 
func (obj *_Gv$sqlWorkareaMgr) WithMultipassesExecutions(multipassesExecutions int64) Option {
	return optionFunc(func(o *options) { o.query["multipasses_executions"] = multipassesExecutions })
}

// WithActiveTime active_time获取 
func (obj *_Gv$sqlWorkareaMgr) WithActiveTime(activeTime int64) Option {
	return optionFunc(func(o *options) { o.query["active_time"] = activeTime })
}

// WithMaxTempsegSize max_tempseg_size获取 
func (obj *_Gv$sqlWorkareaMgr) WithMaxTempsegSize(maxTempsegSize int64) Option {
	return optionFunc(func(o *options) { o.query["max_tempseg_size"] = maxTempsegSize })
}

// WithLastTempsegSize last_tempseg_size获取 
func (obj *_Gv$sqlWorkareaMgr) WithLastTempsegSize(lastTempsegSize int64) Option {
	return optionFunc(func(o *options) { o.query["last_tempseg_size"] = lastTempsegSize })
}

// WithConID con_id获取 
func (obj *_Gv$sqlWorkareaMgr) WithConID(conID int64) Option {
	return optionFunc(func(o *options) { o.query["con_id"] = conID })
}


// GetByOption 功能选项模式获取
func (obj *_Gv$sqlWorkareaMgr) GetByOption(opts ...Option) (result Gv$sqlWorkarea, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlWorkarea{}).Where(options.query).First(&result).Error
	
	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_Gv$sqlWorkareaMgr) GetByOptions(opts ...Option) (results []*Gv$sqlWorkarea, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlWorkarea{}).Where(options.query).Find(&results).Error
	
	return
}


// SelectPage 分页查询
func (obj *_Gv$sqlWorkareaMgr) SelectPage(page IPage,opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]Gv$sqlWorkarea,0)
	var count int64 // 统计总的记录数
	query :=  obj.DB.WithContext(obj.ctx).Model(Gv$sqlWorkarea{}).Where(options.query)
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


// GetFromAddress 通过address获取内容  
func (obj *_Gv$sqlWorkareaMgr) GetFromAddress(address []byte) (results []*Gv$sqlWorkarea, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlWorkarea{}).Where("`address` = ?", address).Find(&results).Error
	
	return
}

// GetBatchFromAddress 批量查找 
func (obj *_Gv$sqlWorkareaMgr) GetBatchFromAddress(addresss [][]byte) (results []*Gv$sqlWorkarea, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlWorkarea{}).Where("`address` IN (?)", addresss).Find(&results).Error
	
	return
}
 
// GetFromHashValue 通过hash_value获取内容  
func (obj *_Gv$sqlWorkareaMgr) GetFromHashValue(hashValue int64) (results []*Gv$sqlWorkarea, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlWorkarea{}).Where("`hash_value` = ?", hashValue).Find(&results).Error
	
	return
}

// GetBatchFromHashValue 批量查找 
func (obj *_Gv$sqlWorkareaMgr) GetBatchFromHashValue(hashValues []int64) (results []*Gv$sqlWorkarea, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlWorkarea{}).Where("`hash_value` IN (?)", hashValues).Find(&results).Error
	
	return
}
 
// GetFromSQLID 通过sql_id获取内容  
func (obj *_Gv$sqlWorkareaMgr) GetFromSQLID(sqlID string) (results []*Gv$sqlWorkarea, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlWorkarea{}).Where("`sql_id` = ?", sqlID).Find(&results).Error
	
	return
}

// GetBatchFromSQLID 批量查找 
func (obj *_Gv$sqlWorkareaMgr) GetBatchFromSQLID(sqlIDs []string) (results []*Gv$sqlWorkarea, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlWorkarea{}).Where("`sql_id` IN (?)", sqlIDs).Find(&results).Error
	
	return
}
 
// GetFromChildNumber 通过child_number获取内容  
func (obj *_Gv$sqlWorkareaMgr) GetFromChildNumber(childNumber int64) (results []*Gv$sqlWorkarea, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlWorkarea{}).Where("`child_number` = ?", childNumber).Find(&results).Error
	
	return
}

// GetBatchFromChildNumber 批量查找 
func (obj *_Gv$sqlWorkareaMgr) GetBatchFromChildNumber(childNumbers []int64) (results []*Gv$sqlWorkarea, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlWorkarea{}).Where("`child_number` IN (?)", childNumbers).Find(&results).Error
	
	return
}
 
// GetFromWorkareaAddress 通过workarea_address获取内容  
func (obj *_Gv$sqlWorkareaMgr) GetFromWorkareaAddress(workareaAddress []byte) (results []*Gv$sqlWorkarea, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlWorkarea{}).Where("`workarea_address` = ?", workareaAddress).Find(&results).Error
	
	return
}

// GetBatchFromWorkareaAddress 批量查找 
func (obj *_Gv$sqlWorkareaMgr) GetBatchFromWorkareaAddress(workareaAddresss [][]byte) (results []*Gv$sqlWorkarea, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlWorkarea{}).Where("`workarea_address` IN (?)", workareaAddresss).Find(&results).Error
	
	return
}
 
// GetFromOperationType 通过operation_type获取内容  
func (obj *_Gv$sqlWorkareaMgr) GetFromOperationType(operationType string) (results []*Gv$sqlWorkarea, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlWorkarea{}).Where("`operation_type` = ?", operationType).Find(&results).Error
	
	return
}

// GetBatchFromOperationType 批量查找 
func (obj *_Gv$sqlWorkareaMgr) GetBatchFromOperationType(operationTypes []string) (results []*Gv$sqlWorkarea, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlWorkarea{}).Where("`operation_type` IN (?)", operationTypes).Find(&results).Error
	
	return
}
 
// GetFromOperationID 通过operation_id获取内容  
func (obj *_Gv$sqlWorkareaMgr) GetFromOperationID(operationID int64) (results []*Gv$sqlWorkarea, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlWorkarea{}).Where("`operation_id` = ?", operationID).Find(&results).Error
	
	return
}

// GetBatchFromOperationID 批量查找 
func (obj *_Gv$sqlWorkareaMgr) GetBatchFromOperationID(operationIDs []int64) (results []*Gv$sqlWorkarea, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlWorkarea{}).Where("`operation_id` IN (?)", operationIDs).Find(&results).Error
	
	return
}
 
// GetFromPolicy 通过policy获取内容  
func (obj *_Gv$sqlWorkareaMgr) GetFromPolicy(policy string) (results []*Gv$sqlWorkarea, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlWorkarea{}).Where("`policy` = ?", policy).Find(&results).Error
	
	return
}

// GetBatchFromPolicy 批量查找 
func (obj *_Gv$sqlWorkareaMgr) GetBatchFromPolicy(policys []string) (results []*Gv$sqlWorkarea, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlWorkarea{}).Where("`policy` IN (?)", policys).Find(&results).Error
	
	return
}
 
// GetFromEstimatedOptimalSize 通过estimated_optimal_size获取内容  
func (obj *_Gv$sqlWorkareaMgr) GetFromEstimatedOptimalSize(estimatedOptimalSize int64) (results []*Gv$sqlWorkarea, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlWorkarea{}).Where("`estimated_optimal_size` = ?", estimatedOptimalSize).Find(&results).Error
	
	return
}

// GetBatchFromEstimatedOptimalSize 批量查找 
func (obj *_Gv$sqlWorkareaMgr) GetBatchFromEstimatedOptimalSize(estimatedOptimalSizes []int64) (results []*Gv$sqlWorkarea, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlWorkarea{}).Where("`estimated_optimal_size` IN (?)", estimatedOptimalSizes).Find(&results).Error
	
	return
}
 
// GetFromEstimatedOnepassSize 通过estimated_onepass_size获取内容  
func (obj *_Gv$sqlWorkareaMgr) GetFromEstimatedOnepassSize(estimatedOnepassSize int64) (results []*Gv$sqlWorkarea, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlWorkarea{}).Where("`estimated_onepass_size` = ?", estimatedOnepassSize).Find(&results).Error
	
	return
}

// GetBatchFromEstimatedOnepassSize 批量查找 
func (obj *_Gv$sqlWorkareaMgr) GetBatchFromEstimatedOnepassSize(estimatedOnepassSizes []int64) (results []*Gv$sqlWorkarea, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlWorkarea{}).Where("`estimated_onepass_size` IN (?)", estimatedOnepassSizes).Find(&results).Error
	
	return
}
 
// GetFromLastMemoryUsed 通过last_memory_used获取内容  
func (obj *_Gv$sqlWorkareaMgr) GetFromLastMemoryUsed(lastMemoryUsed int64) (results []*Gv$sqlWorkarea, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlWorkarea{}).Where("`last_memory_used` = ?", lastMemoryUsed).Find(&results).Error
	
	return
}

// GetBatchFromLastMemoryUsed 批量查找 
func (obj *_Gv$sqlWorkareaMgr) GetBatchFromLastMemoryUsed(lastMemoryUseds []int64) (results []*Gv$sqlWorkarea, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlWorkarea{}).Where("`last_memory_used` IN (?)", lastMemoryUseds).Find(&results).Error
	
	return
}
 
// GetFromLastExecution 通过last_execution获取内容  
func (obj *_Gv$sqlWorkareaMgr) GetFromLastExecution(lastExecution string) (results []*Gv$sqlWorkarea, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlWorkarea{}).Where("`last_execution` = ?", lastExecution).Find(&results).Error
	
	return
}

// GetBatchFromLastExecution 批量查找 
func (obj *_Gv$sqlWorkareaMgr) GetBatchFromLastExecution(lastExecutions []string) (results []*Gv$sqlWorkarea, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlWorkarea{}).Where("`last_execution` IN (?)", lastExecutions).Find(&results).Error
	
	return
}
 
// GetFromLastDegree 通过last_degree获取内容  
func (obj *_Gv$sqlWorkareaMgr) GetFromLastDegree(lastDegree int64) (results []*Gv$sqlWorkarea, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlWorkarea{}).Where("`last_degree` = ?", lastDegree).Find(&results).Error
	
	return
}

// GetBatchFromLastDegree 批量查找 
func (obj *_Gv$sqlWorkareaMgr) GetBatchFromLastDegree(lastDegrees []int64) (results []*Gv$sqlWorkarea, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlWorkarea{}).Where("`last_degree` IN (?)", lastDegrees).Find(&results).Error
	
	return
}
 
// GetFromTotalExecutions 通过total_executions获取内容  
func (obj *_Gv$sqlWorkareaMgr) GetFromTotalExecutions(totalExecutions int64) (results []*Gv$sqlWorkarea, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlWorkarea{}).Where("`total_executions` = ?", totalExecutions).Find(&results).Error
	
	return
}

// GetBatchFromTotalExecutions 批量查找 
func (obj *_Gv$sqlWorkareaMgr) GetBatchFromTotalExecutions(totalExecutionss []int64) (results []*Gv$sqlWorkarea, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlWorkarea{}).Where("`total_executions` IN (?)", totalExecutionss).Find(&results).Error
	
	return
}
 
// GetFromOptimalExecutions 通过optimal_executions获取内容  
func (obj *_Gv$sqlWorkareaMgr) GetFromOptimalExecutions(optimalExecutions int64) (results []*Gv$sqlWorkarea, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlWorkarea{}).Where("`optimal_executions` = ?", optimalExecutions).Find(&results).Error
	
	return
}

// GetBatchFromOptimalExecutions 批量查找 
func (obj *_Gv$sqlWorkareaMgr) GetBatchFromOptimalExecutions(optimalExecutionss []int64) (results []*Gv$sqlWorkarea, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlWorkarea{}).Where("`optimal_executions` IN (?)", optimalExecutionss).Find(&results).Error
	
	return
}
 
// GetFromOnepassExecutions 通过onepass_executions获取内容  
func (obj *_Gv$sqlWorkareaMgr) GetFromOnepassExecutions(onepassExecutions int64) (results []*Gv$sqlWorkarea, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlWorkarea{}).Where("`onepass_executions` = ?", onepassExecutions).Find(&results).Error
	
	return
}

// GetBatchFromOnepassExecutions 批量查找 
func (obj *_Gv$sqlWorkareaMgr) GetBatchFromOnepassExecutions(onepassExecutionss []int64) (results []*Gv$sqlWorkarea, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlWorkarea{}).Where("`onepass_executions` IN (?)", onepassExecutionss).Find(&results).Error
	
	return
}
 
// GetFromMultipassesExecutions 通过multipasses_executions获取内容  
func (obj *_Gv$sqlWorkareaMgr) GetFromMultipassesExecutions(multipassesExecutions int64) (results []*Gv$sqlWorkarea, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlWorkarea{}).Where("`multipasses_executions` = ?", multipassesExecutions).Find(&results).Error
	
	return
}

// GetBatchFromMultipassesExecutions 批量查找 
func (obj *_Gv$sqlWorkareaMgr) GetBatchFromMultipassesExecutions(multipassesExecutionss []int64) (results []*Gv$sqlWorkarea, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlWorkarea{}).Where("`multipasses_executions` IN (?)", multipassesExecutionss).Find(&results).Error
	
	return
}
 
// GetFromActiveTime 通过active_time获取内容  
func (obj *_Gv$sqlWorkareaMgr) GetFromActiveTime(activeTime int64) (results []*Gv$sqlWorkarea, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlWorkarea{}).Where("`active_time` = ?", activeTime).Find(&results).Error
	
	return
}

// GetBatchFromActiveTime 批量查找 
func (obj *_Gv$sqlWorkareaMgr) GetBatchFromActiveTime(activeTimes []int64) (results []*Gv$sqlWorkarea, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlWorkarea{}).Where("`active_time` IN (?)", activeTimes).Find(&results).Error
	
	return
}
 
// GetFromMaxTempsegSize 通过max_tempseg_size获取内容  
func (obj *_Gv$sqlWorkareaMgr) GetFromMaxTempsegSize(maxTempsegSize int64) (results []*Gv$sqlWorkarea, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlWorkarea{}).Where("`max_tempseg_size` = ?", maxTempsegSize).Find(&results).Error
	
	return
}

// GetBatchFromMaxTempsegSize 批量查找 
func (obj *_Gv$sqlWorkareaMgr) GetBatchFromMaxTempsegSize(maxTempsegSizes []int64) (results []*Gv$sqlWorkarea, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlWorkarea{}).Where("`max_tempseg_size` IN (?)", maxTempsegSizes).Find(&results).Error
	
	return
}
 
// GetFromLastTempsegSize 通过last_tempseg_size获取内容  
func (obj *_Gv$sqlWorkareaMgr) GetFromLastTempsegSize(lastTempsegSize int64) (results []*Gv$sqlWorkarea, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlWorkarea{}).Where("`last_tempseg_size` = ?", lastTempsegSize).Find(&results).Error
	
	return
}

// GetBatchFromLastTempsegSize 批量查找 
func (obj *_Gv$sqlWorkareaMgr) GetBatchFromLastTempsegSize(lastTempsegSizes []int64) (results []*Gv$sqlWorkarea, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlWorkarea{}).Where("`last_tempseg_size` IN (?)", lastTempsegSizes).Find(&results).Error
	
	return
}
 
// GetFromConID 通过con_id获取内容  
func (obj *_Gv$sqlWorkareaMgr) GetFromConID(conID int64) (results []*Gv$sqlWorkarea, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlWorkarea{}).Where("`con_id` = ?", conID).Find(&results).Error
	
	return
}

// GetBatchFromConID 批量查找 
func (obj *_Gv$sqlWorkareaMgr) GetBatchFromConID(conIDs []int64) (results []*Gv$sqlWorkarea, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlWorkarea{}).Where("`con_id` IN (?)", conIDs).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 

 

	

