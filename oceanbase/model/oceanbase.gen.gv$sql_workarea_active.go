package	model	
import (	
"context"	
"gorm.io/gorm"	
"gorm.io/datatypes"	
"fmt"	
)	

type _Gv$sqlWorkareaActiveMgr struct {
	*_BaseMgr
}

// Gv$sqlWorkareaActiveMgr open func
func Gv$sqlWorkareaActiveMgr(db *gorm.DB) *_Gv$sqlWorkareaActiveMgr {
	if db == nil {
		panic(fmt.Errorf("Gv$sqlWorkareaActiveMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_Gv$sqlWorkareaActiveMgr{_BaseMgr: &_BaseMgr{DB: db.Table("gv$sql_workarea_active"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_Gv$sqlWorkareaActiveMgr) GetTableName() string {
	return "gv$sql_workarea_active"
}

// Reset 重置gorm会话
func (obj *_Gv$sqlWorkareaActiveMgr) Reset() *_Gv$sqlWorkareaActiveMgr {
	obj.New()
	return obj
}

// Get 获取 
func (obj *_Gv$sqlWorkareaActiveMgr) Get() (result Gv$sqlWorkareaActive, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlWorkareaActive{}).First(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_Gv$sqlWorkareaActiveMgr) Gets() (results []*Gv$sqlWorkareaActive, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlWorkareaActive{}).Find(&results).Error
	
	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_Gv$sqlWorkareaActiveMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(Gv$sqlWorkareaActive{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithSQLHashValue sql_hash_value获取 
func (obj *_Gv$sqlWorkareaActiveMgr) WithSQLHashValue(sqlHashValue int64) Option {
	return optionFunc(func(o *options) { o.query["sql_hash_value"] = sqlHashValue })
}

// WithSQLID sql_id获取 
func (obj *_Gv$sqlWorkareaActiveMgr) WithSQLID(sqlID string) Option {
	return optionFunc(func(o *options) { o.query["sql_id"] = sqlID })
}

// WithSQLExecStart sql_exec_start获取 
func (obj *_Gv$sqlWorkareaActiveMgr) WithSQLExecStart(sqlExecStart datatypes.Date) Option {
	return optionFunc(func(o *options) { o.query["sql_exec_start"] = sqlExecStart })
}

// WithSQLExecID sql_exec_id获取 
func (obj *_Gv$sqlWorkareaActiveMgr) WithSQLExecID(sqlExecID int64) Option {
	return optionFunc(func(o *options) { o.query["sql_exec_id"] = sqlExecID })
}

// WithWorkareaAddress workarea_address获取 
func (obj *_Gv$sqlWorkareaActiveMgr) WithWorkareaAddress(workareaAddress []byte) Option {
	return optionFunc(func(o *options) { o.query["workarea_address"] = workareaAddress })
}

// WithOperationType operation_type获取 
func (obj *_Gv$sqlWorkareaActiveMgr) WithOperationType(operationType string) Option {
	return optionFunc(func(o *options) { o.query["operation_type"] = operationType })
}

// WithOperationID operation_id获取 
func (obj *_Gv$sqlWorkareaActiveMgr) WithOperationID(operationID int64) Option {
	return optionFunc(func(o *options) { o.query["operation_id"] = operationID })
}

// WithPolicy policy获取 
func (obj *_Gv$sqlWorkareaActiveMgr) WithPolicy(policy string) Option {
	return optionFunc(func(o *options) { o.query["policy"] = policy })
}

// WithSid sid获取 
func (obj *_Gv$sqlWorkareaActiveMgr) WithSid(sid int64) Option {
	return optionFunc(func(o *options) { o.query["sid"] = sid })
}

// WithQcinstID qcinst_id获取 
func (obj *_Gv$sqlWorkareaActiveMgr) WithQcinstID(qcinstID int64) Option {
	return optionFunc(func(o *options) { o.query["qcinst_id"] = qcinstID })
}

// WithQcsid qcsid获取 
func (obj *_Gv$sqlWorkareaActiveMgr) WithQcsid(qcsid int64) Option {
	return optionFunc(func(o *options) { o.query["qcsid"] = qcsid })
}

// WithActiveTime active_time获取 
func (obj *_Gv$sqlWorkareaActiveMgr) WithActiveTime(activeTime int64) Option {
	return optionFunc(func(o *options) { o.query["active_time"] = activeTime })
}

// WithWorkAreaSize work_area_size获取 
func (obj *_Gv$sqlWorkareaActiveMgr) WithWorkAreaSize(workAreaSize int64) Option {
	return optionFunc(func(o *options) { o.query["work_area_size"] = workAreaSize })
}

// WithExpectSize expect_size获取 
func (obj *_Gv$sqlWorkareaActiveMgr) WithExpectSize(expectSize int64) Option {
	return optionFunc(func(o *options) { o.query["expect_size"] = expectSize })
}

// WithActualMemUsed actual_mem_used获取 
func (obj *_Gv$sqlWorkareaActiveMgr) WithActualMemUsed(actualMemUsed int64) Option {
	return optionFunc(func(o *options) { o.query["actual_mem_used"] = actualMemUsed })
}

// WithMaxMemUsed max_mem_used获取 
func (obj *_Gv$sqlWorkareaActiveMgr) WithMaxMemUsed(maxMemUsed int64) Option {
	return optionFunc(func(o *options) { o.query["max_mem_used"] = maxMemUsed })
}

// WithNumberPasses number_passes获取 
func (obj *_Gv$sqlWorkareaActiveMgr) WithNumberPasses(numberPasses int64) Option {
	return optionFunc(func(o *options) { o.query["number_passes"] = numberPasses })
}

// WithTempsegSize tempseg_size获取 
func (obj *_Gv$sqlWorkareaActiveMgr) WithTempsegSize(tempsegSize int64) Option {
	return optionFunc(func(o *options) { o.query["tempseg_size"] = tempsegSize })
}

// WithTablespace tablespace获取 
func (obj *_Gv$sqlWorkareaActiveMgr) WithTablespace(tablespace string) Option {
	return optionFunc(func(o *options) { o.query["tablespace"] = tablespace })
}

// WithSegrfno# segrfno#获取 
func (obj *_Gv$sqlWorkareaActiveMgr) WithSegrfno#(segrfno# int64) Option {
	return optionFunc(func(o *options) { o.query["segrfno#"] = segrfno# })
}

// WithSegblk# segblk#获取 
func (obj *_Gv$sqlWorkareaActiveMgr) WithSegblk#(segblk# int64) Option {
	return optionFunc(func(o *options) { o.query["segblk#"] = segblk# })
}

// WithConID con_id获取 
func (obj *_Gv$sqlWorkareaActiveMgr) WithConID(conID int64) Option {
	return optionFunc(func(o *options) { o.query["con_id"] = conID })
}


// GetByOption 功能选项模式获取
func (obj *_Gv$sqlWorkareaActiveMgr) GetByOption(opts ...Option) (result Gv$sqlWorkareaActive, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlWorkareaActive{}).Where(options.query).First(&result).Error
	
	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_Gv$sqlWorkareaActiveMgr) GetByOptions(opts ...Option) (results []*Gv$sqlWorkareaActive, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlWorkareaActive{}).Where(options.query).Find(&results).Error
	
	return
}


// SelectPage 分页查询
func (obj *_Gv$sqlWorkareaActiveMgr) SelectPage(page IPage,opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]Gv$sqlWorkareaActive,0)
	var count int64 // 统计总的记录数
	query :=  obj.DB.WithContext(obj.ctx).Model(Gv$sqlWorkareaActive{}).Where(options.query)
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


// GetFromSQLHashValue 通过sql_hash_value获取内容  
func (obj *_Gv$sqlWorkareaActiveMgr) GetFromSQLHashValue(sqlHashValue int64) (results []*Gv$sqlWorkareaActive, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlWorkareaActive{}).Where("`sql_hash_value` = ?", sqlHashValue).Find(&results).Error
	
	return
}

// GetBatchFromSQLHashValue 批量查找 
func (obj *_Gv$sqlWorkareaActiveMgr) GetBatchFromSQLHashValue(sqlHashValues []int64) (results []*Gv$sqlWorkareaActive, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlWorkareaActive{}).Where("`sql_hash_value` IN (?)", sqlHashValues).Find(&results).Error
	
	return
}
 
// GetFromSQLID 通过sql_id获取内容  
func (obj *_Gv$sqlWorkareaActiveMgr) GetFromSQLID(sqlID string) (results []*Gv$sqlWorkareaActive, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlWorkareaActive{}).Where("`sql_id` = ?", sqlID).Find(&results).Error
	
	return
}

// GetBatchFromSQLID 批量查找 
func (obj *_Gv$sqlWorkareaActiveMgr) GetBatchFromSQLID(sqlIDs []string) (results []*Gv$sqlWorkareaActive, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlWorkareaActive{}).Where("`sql_id` IN (?)", sqlIDs).Find(&results).Error
	
	return
}
 
// GetFromSQLExecStart 通过sql_exec_start获取内容  
func (obj *_Gv$sqlWorkareaActiveMgr) GetFromSQLExecStart(sqlExecStart datatypes.Date) (results []*Gv$sqlWorkareaActive, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlWorkareaActive{}).Where("`sql_exec_start` = ?", sqlExecStart).Find(&results).Error
	
	return
}

// GetBatchFromSQLExecStart 批量查找 
func (obj *_Gv$sqlWorkareaActiveMgr) GetBatchFromSQLExecStart(sqlExecStarts []datatypes.Date) (results []*Gv$sqlWorkareaActive, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlWorkareaActive{}).Where("`sql_exec_start` IN (?)", sqlExecStarts).Find(&results).Error
	
	return
}
 
// GetFromSQLExecID 通过sql_exec_id获取内容  
func (obj *_Gv$sqlWorkareaActiveMgr) GetFromSQLExecID(sqlExecID int64) (results []*Gv$sqlWorkareaActive, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlWorkareaActive{}).Where("`sql_exec_id` = ?", sqlExecID).Find(&results).Error
	
	return
}

// GetBatchFromSQLExecID 批量查找 
func (obj *_Gv$sqlWorkareaActiveMgr) GetBatchFromSQLExecID(sqlExecIDs []int64) (results []*Gv$sqlWorkareaActive, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlWorkareaActive{}).Where("`sql_exec_id` IN (?)", sqlExecIDs).Find(&results).Error
	
	return
}
 
// GetFromWorkareaAddress 通过workarea_address获取内容  
func (obj *_Gv$sqlWorkareaActiveMgr) GetFromWorkareaAddress(workareaAddress []byte) (results []*Gv$sqlWorkareaActive, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlWorkareaActive{}).Where("`workarea_address` = ?", workareaAddress).Find(&results).Error
	
	return
}

// GetBatchFromWorkareaAddress 批量查找 
func (obj *_Gv$sqlWorkareaActiveMgr) GetBatchFromWorkareaAddress(workareaAddresss [][]byte) (results []*Gv$sqlWorkareaActive, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlWorkareaActive{}).Where("`workarea_address` IN (?)", workareaAddresss).Find(&results).Error
	
	return
}
 
// GetFromOperationType 通过operation_type获取内容  
func (obj *_Gv$sqlWorkareaActiveMgr) GetFromOperationType(operationType string) (results []*Gv$sqlWorkareaActive, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlWorkareaActive{}).Where("`operation_type` = ?", operationType).Find(&results).Error
	
	return
}

// GetBatchFromOperationType 批量查找 
func (obj *_Gv$sqlWorkareaActiveMgr) GetBatchFromOperationType(operationTypes []string) (results []*Gv$sqlWorkareaActive, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlWorkareaActive{}).Where("`operation_type` IN (?)", operationTypes).Find(&results).Error
	
	return
}
 
// GetFromOperationID 通过operation_id获取内容  
func (obj *_Gv$sqlWorkareaActiveMgr) GetFromOperationID(operationID int64) (results []*Gv$sqlWorkareaActive, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlWorkareaActive{}).Where("`operation_id` = ?", operationID).Find(&results).Error
	
	return
}

// GetBatchFromOperationID 批量查找 
func (obj *_Gv$sqlWorkareaActiveMgr) GetBatchFromOperationID(operationIDs []int64) (results []*Gv$sqlWorkareaActive, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlWorkareaActive{}).Where("`operation_id` IN (?)", operationIDs).Find(&results).Error
	
	return
}
 
// GetFromPolicy 通过policy获取内容  
func (obj *_Gv$sqlWorkareaActiveMgr) GetFromPolicy(policy string) (results []*Gv$sqlWorkareaActive, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlWorkareaActive{}).Where("`policy` = ?", policy).Find(&results).Error
	
	return
}

// GetBatchFromPolicy 批量查找 
func (obj *_Gv$sqlWorkareaActiveMgr) GetBatchFromPolicy(policys []string) (results []*Gv$sqlWorkareaActive, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlWorkareaActive{}).Where("`policy` IN (?)", policys).Find(&results).Error
	
	return
}
 
// GetFromSid 通过sid获取内容  
func (obj *_Gv$sqlWorkareaActiveMgr) GetFromSid(sid int64) (results []*Gv$sqlWorkareaActive, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlWorkareaActive{}).Where("`sid` = ?", sid).Find(&results).Error
	
	return
}

// GetBatchFromSid 批量查找 
func (obj *_Gv$sqlWorkareaActiveMgr) GetBatchFromSid(sids []int64) (results []*Gv$sqlWorkareaActive, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlWorkareaActive{}).Where("`sid` IN (?)", sids).Find(&results).Error
	
	return
}
 
// GetFromQcinstID 通过qcinst_id获取内容  
func (obj *_Gv$sqlWorkareaActiveMgr) GetFromQcinstID(qcinstID int64) (results []*Gv$sqlWorkareaActive, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlWorkareaActive{}).Where("`qcinst_id` = ?", qcinstID).Find(&results).Error
	
	return
}

// GetBatchFromQcinstID 批量查找 
func (obj *_Gv$sqlWorkareaActiveMgr) GetBatchFromQcinstID(qcinstIDs []int64) (results []*Gv$sqlWorkareaActive, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlWorkareaActive{}).Where("`qcinst_id` IN (?)", qcinstIDs).Find(&results).Error
	
	return
}
 
// GetFromQcsid 通过qcsid获取内容  
func (obj *_Gv$sqlWorkareaActiveMgr) GetFromQcsid(qcsid int64) (results []*Gv$sqlWorkareaActive, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlWorkareaActive{}).Where("`qcsid` = ?", qcsid).Find(&results).Error
	
	return
}

// GetBatchFromQcsid 批量查找 
func (obj *_Gv$sqlWorkareaActiveMgr) GetBatchFromQcsid(qcsids []int64) (results []*Gv$sqlWorkareaActive, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlWorkareaActive{}).Where("`qcsid` IN (?)", qcsids).Find(&results).Error
	
	return
}
 
// GetFromActiveTime 通过active_time获取内容  
func (obj *_Gv$sqlWorkareaActiveMgr) GetFromActiveTime(activeTime int64) (results []*Gv$sqlWorkareaActive, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlWorkareaActive{}).Where("`active_time` = ?", activeTime).Find(&results).Error
	
	return
}

// GetBatchFromActiveTime 批量查找 
func (obj *_Gv$sqlWorkareaActiveMgr) GetBatchFromActiveTime(activeTimes []int64) (results []*Gv$sqlWorkareaActive, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlWorkareaActive{}).Where("`active_time` IN (?)", activeTimes).Find(&results).Error
	
	return
}
 
// GetFromWorkAreaSize 通过work_area_size获取内容  
func (obj *_Gv$sqlWorkareaActiveMgr) GetFromWorkAreaSize(workAreaSize int64) (results []*Gv$sqlWorkareaActive, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlWorkareaActive{}).Where("`work_area_size` = ?", workAreaSize).Find(&results).Error
	
	return
}

// GetBatchFromWorkAreaSize 批量查找 
func (obj *_Gv$sqlWorkareaActiveMgr) GetBatchFromWorkAreaSize(workAreaSizes []int64) (results []*Gv$sqlWorkareaActive, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlWorkareaActive{}).Where("`work_area_size` IN (?)", workAreaSizes).Find(&results).Error
	
	return
}
 
// GetFromExpectSize 通过expect_size获取内容  
func (obj *_Gv$sqlWorkareaActiveMgr) GetFromExpectSize(expectSize int64) (results []*Gv$sqlWorkareaActive, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlWorkareaActive{}).Where("`expect_size` = ?", expectSize).Find(&results).Error
	
	return
}

// GetBatchFromExpectSize 批量查找 
func (obj *_Gv$sqlWorkareaActiveMgr) GetBatchFromExpectSize(expectSizes []int64) (results []*Gv$sqlWorkareaActive, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlWorkareaActive{}).Where("`expect_size` IN (?)", expectSizes).Find(&results).Error
	
	return
}
 
// GetFromActualMemUsed 通过actual_mem_used获取内容  
func (obj *_Gv$sqlWorkareaActiveMgr) GetFromActualMemUsed(actualMemUsed int64) (results []*Gv$sqlWorkareaActive, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlWorkareaActive{}).Where("`actual_mem_used` = ?", actualMemUsed).Find(&results).Error
	
	return
}

// GetBatchFromActualMemUsed 批量查找 
func (obj *_Gv$sqlWorkareaActiveMgr) GetBatchFromActualMemUsed(actualMemUseds []int64) (results []*Gv$sqlWorkareaActive, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlWorkareaActive{}).Where("`actual_mem_used` IN (?)", actualMemUseds).Find(&results).Error
	
	return
}
 
// GetFromMaxMemUsed 通过max_mem_used获取内容  
func (obj *_Gv$sqlWorkareaActiveMgr) GetFromMaxMemUsed(maxMemUsed int64) (results []*Gv$sqlWorkareaActive, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlWorkareaActive{}).Where("`max_mem_used` = ?", maxMemUsed).Find(&results).Error
	
	return
}

// GetBatchFromMaxMemUsed 批量查找 
func (obj *_Gv$sqlWorkareaActiveMgr) GetBatchFromMaxMemUsed(maxMemUseds []int64) (results []*Gv$sqlWorkareaActive, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlWorkareaActive{}).Where("`max_mem_used` IN (?)", maxMemUseds).Find(&results).Error
	
	return
}
 
// GetFromNumberPasses 通过number_passes获取内容  
func (obj *_Gv$sqlWorkareaActiveMgr) GetFromNumberPasses(numberPasses int64) (results []*Gv$sqlWorkareaActive, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlWorkareaActive{}).Where("`number_passes` = ?", numberPasses).Find(&results).Error
	
	return
}

// GetBatchFromNumberPasses 批量查找 
func (obj *_Gv$sqlWorkareaActiveMgr) GetBatchFromNumberPasses(numberPassess []int64) (results []*Gv$sqlWorkareaActive, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlWorkareaActive{}).Where("`number_passes` IN (?)", numberPassess).Find(&results).Error
	
	return
}
 
// GetFromTempsegSize 通过tempseg_size获取内容  
func (obj *_Gv$sqlWorkareaActiveMgr) GetFromTempsegSize(tempsegSize int64) (results []*Gv$sqlWorkareaActive, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlWorkareaActive{}).Where("`tempseg_size` = ?", tempsegSize).Find(&results).Error
	
	return
}

// GetBatchFromTempsegSize 批量查找 
func (obj *_Gv$sqlWorkareaActiveMgr) GetBatchFromTempsegSize(tempsegSizes []int64) (results []*Gv$sqlWorkareaActive, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlWorkareaActive{}).Where("`tempseg_size` IN (?)", tempsegSizes).Find(&results).Error
	
	return
}
 
// GetFromTablespace 通过tablespace获取内容  
func (obj *_Gv$sqlWorkareaActiveMgr) GetFromTablespace(tablespace string) (results []*Gv$sqlWorkareaActive, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlWorkareaActive{}).Where("`tablespace` = ?", tablespace).Find(&results).Error
	
	return
}

// GetBatchFromTablespace 批量查找 
func (obj *_Gv$sqlWorkareaActiveMgr) GetBatchFromTablespace(tablespaces []string) (results []*Gv$sqlWorkareaActive, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlWorkareaActive{}).Where("`tablespace` IN (?)", tablespaces).Find(&results).Error
	
	return
}
 
// GetFromSegrfno# 通过segrfno#获取内容  
func (obj *_Gv$sqlWorkareaActiveMgr) GetFromSegrfno#(segrfno# int64) (results []*Gv$sqlWorkareaActive, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlWorkareaActive{}).Where("`segrfno#` = ?", segrfno#).Find(&results).Error
	
	return
}

// GetBatchFromSegrfno# 批量查找 
func (obj *_Gv$sqlWorkareaActiveMgr) GetBatchFromSegrfno#(segrfno#s []int64) (results []*Gv$sqlWorkareaActive, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlWorkareaActive{}).Where("`segrfno#` IN (?)", segrfno#s).Find(&results).Error
	
	return
}
 
// GetFromSegblk# 通过segblk#获取内容  
func (obj *_Gv$sqlWorkareaActiveMgr) GetFromSegblk#(segblk# int64) (results []*Gv$sqlWorkareaActive, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlWorkareaActive{}).Where("`segblk#` = ?", segblk#).Find(&results).Error
	
	return
}

// GetBatchFromSegblk# 批量查找 
func (obj *_Gv$sqlWorkareaActiveMgr) GetBatchFromSegblk#(segblk#s []int64) (results []*Gv$sqlWorkareaActive, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlWorkareaActive{}).Where("`segblk#` IN (?)", segblk#s).Find(&results).Error
	
	return
}
 
// GetFromConID 通过con_id获取内容  
func (obj *_Gv$sqlWorkareaActiveMgr) GetFromConID(conID int64) (results []*Gv$sqlWorkareaActive, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlWorkareaActive{}).Where("`con_id` = ?", conID).Find(&results).Error
	
	return
}

// GetBatchFromConID 批量查找 
func (obj *_Gv$sqlWorkareaActiveMgr) GetBatchFromConID(conIDs []int64) (results []*Gv$sqlWorkareaActive, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$sqlWorkareaActive{}).Where("`con_id` IN (?)", conIDs).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 

 

	

