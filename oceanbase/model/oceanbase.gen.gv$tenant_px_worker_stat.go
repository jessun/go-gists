package	model	
import (	
"gorm.io/gorm"	
"time"	
"fmt"	
"context"	
)	

type _Gv$tenantPxWorkerStatMgr struct {
	*_BaseMgr
}

// Gv$tenantPxWorkerStatMgr open func
func Gv$tenantPxWorkerStatMgr(db *gorm.DB) *_Gv$tenantPxWorkerStatMgr {
	if db == nil {
		panic(fmt.Errorf("Gv$tenantPxWorkerStatMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_Gv$tenantPxWorkerStatMgr{_BaseMgr: &_BaseMgr{DB: db.Table("gv$tenant_px_worker_stat"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_Gv$tenantPxWorkerStatMgr) GetTableName() string {
	return "gv$tenant_px_worker_stat"
}

// Reset 重置gorm会话
func (obj *_Gv$tenantPxWorkerStatMgr) Reset() *_Gv$tenantPxWorkerStatMgr {
	obj.New()
	return obj
}

// Get 获取 
func (obj *_Gv$tenantPxWorkerStatMgr) Get() (result Gv$tenantPxWorkerStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$tenantPxWorkerStat{}).First(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_Gv$tenantPxWorkerStatMgr) Gets() (results []*Gv$tenantPxWorkerStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$tenantPxWorkerStat{}).Find(&results).Error
	
	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_Gv$tenantPxWorkerStatMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(Gv$tenantPxWorkerStat{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithSessionID session_id获取 
func (obj *_Gv$tenantPxWorkerStatMgr) WithSessionID(sessionID int64) Option {
	return optionFunc(func(o *options) { o.query["session_id"] = sessionID })
}

// WithTenantID tenant_id获取 
func (obj *_Gv$tenantPxWorkerStatMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithSvrIP svr_ip获取 
func (obj *_Gv$tenantPxWorkerStatMgr) WithSvrIP(svrIP string) Option {
	return optionFunc(func(o *options) { o.query["svr_ip"] = svrIP })
}

// WithSvrPort svr_port获取 
func (obj *_Gv$tenantPxWorkerStatMgr) WithSvrPort(svrPort int64) Option {
	return optionFunc(func(o *options) { o.query["svr_port"] = svrPort })
}

// WithTraceID trace_id获取 
func (obj *_Gv$tenantPxWorkerStatMgr) WithTraceID(traceID string) Option {
	return optionFunc(func(o *options) { o.query["trace_id"] = traceID })
}

// WithQcID qc_id获取 
func (obj *_Gv$tenantPxWorkerStatMgr) WithQcID(qcID int64) Option {
	return optionFunc(func(o *options) { o.query["qc_id"] = qcID })
}

// WithSqcID sqc_id获取 
func (obj *_Gv$tenantPxWorkerStatMgr) WithSqcID(sqcID int64) Option {
	return optionFunc(func(o *options) { o.query["sqc_id"] = sqcID })
}

// WithWorkerID worker_id获取 
func (obj *_Gv$tenantPxWorkerStatMgr) WithWorkerID(workerID int64) Option {
	return optionFunc(func(o *options) { o.query["worker_id"] = workerID })
}

// WithDfoID dfo_id获取 
func (obj *_Gv$tenantPxWorkerStatMgr) WithDfoID(dfoID int64) Option {
	return optionFunc(func(o *options) { o.query["dfo_id"] = dfoID })
}

// WithStartTime start_time获取 
func (obj *_Gv$tenantPxWorkerStatMgr) WithStartTime(startTime time.Time) Option {
	return optionFunc(func(o *options) { o.query["start_time"] = startTime })
}


// GetByOption 功能选项模式获取
func (obj *_Gv$tenantPxWorkerStatMgr) GetByOption(opts ...Option) (result Gv$tenantPxWorkerStat, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Gv$tenantPxWorkerStat{}).Where(options.query).First(&result).Error
	
	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_Gv$tenantPxWorkerStatMgr) GetByOptions(opts ...Option) (results []*Gv$tenantPxWorkerStat, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Gv$tenantPxWorkerStat{}).Where(options.query).Find(&results).Error
	
	return
}


// SelectPage 分页查询
func (obj *_Gv$tenantPxWorkerStatMgr) SelectPage(page IPage,opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]Gv$tenantPxWorkerStat,0)
	var count int64 // 统计总的记录数
	query :=  obj.DB.WithContext(obj.ctx).Model(Gv$tenantPxWorkerStat{}).Where(options.query)
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


// GetFromSessionID 通过session_id获取内容  
func (obj *_Gv$tenantPxWorkerStatMgr) GetFromSessionID(sessionID int64) (results []*Gv$tenantPxWorkerStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$tenantPxWorkerStat{}).Where("`session_id` = ?", sessionID).Find(&results).Error
	
	return
}

// GetBatchFromSessionID 批量查找 
func (obj *_Gv$tenantPxWorkerStatMgr) GetBatchFromSessionID(sessionIDs []int64) (results []*Gv$tenantPxWorkerStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$tenantPxWorkerStat{}).Where("`session_id` IN (?)", sessionIDs).Find(&results).Error
	
	return
}
 
// GetFromTenantID 通过tenant_id获取内容  
func (obj *_Gv$tenantPxWorkerStatMgr) GetFromTenantID(tenantID int64) (results []*Gv$tenantPxWorkerStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$tenantPxWorkerStat{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error
	
	return
}

// GetBatchFromTenantID 批量查找 
func (obj *_Gv$tenantPxWorkerStatMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*Gv$tenantPxWorkerStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$tenantPxWorkerStat{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error
	
	return
}
 
// GetFromSvrIP 通过svr_ip获取内容  
func (obj *_Gv$tenantPxWorkerStatMgr) GetFromSvrIP(svrIP string) (results []*Gv$tenantPxWorkerStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$tenantPxWorkerStat{}).Where("`svr_ip` = ?", svrIP).Find(&results).Error
	
	return
}

// GetBatchFromSvrIP 批量查找 
func (obj *_Gv$tenantPxWorkerStatMgr) GetBatchFromSvrIP(svrIPs []string) (results []*Gv$tenantPxWorkerStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$tenantPxWorkerStat{}).Where("`svr_ip` IN (?)", svrIPs).Find(&results).Error
	
	return
}
 
// GetFromSvrPort 通过svr_port获取内容  
func (obj *_Gv$tenantPxWorkerStatMgr) GetFromSvrPort(svrPort int64) (results []*Gv$tenantPxWorkerStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$tenantPxWorkerStat{}).Where("`svr_port` = ?", svrPort).Find(&results).Error
	
	return
}

// GetBatchFromSvrPort 批量查找 
func (obj *_Gv$tenantPxWorkerStatMgr) GetBatchFromSvrPort(svrPorts []int64) (results []*Gv$tenantPxWorkerStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$tenantPxWorkerStat{}).Where("`svr_port` IN (?)", svrPorts).Find(&results).Error
	
	return
}
 
// GetFromTraceID 通过trace_id获取内容  
func (obj *_Gv$tenantPxWorkerStatMgr) GetFromTraceID(traceID string) (results []*Gv$tenantPxWorkerStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$tenantPxWorkerStat{}).Where("`trace_id` = ?", traceID).Find(&results).Error
	
	return
}

// GetBatchFromTraceID 批量查找 
func (obj *_Gv$tenantPxWorkerStatMgr) GetBatchFromTraceID(traceIDs []string) (results []*Gv$tenantPxWorkerStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$tenantPxWorkerStat{}).Where("`trace_id` IN (?)", traceIDs).Find(&results).Error
	
	return
}
 
// GetFromQcID 通过qc_id获取内容  
func (obj *_Gv$tenantPxWorkerStatMgr) GetFromQcID(qcID int64) (results []*Gv$tenantPxWorkerStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$tenantPxWorkerStat{}).Where("`qc_id` = ?", qcID).Find(&results).Error
	
	return
}

// GetBatchFromQcID 批量查找 
func (obj *_Gv$tenantPxWorkerStatMgr) GetBatchFromQcID(qcIDs []int64) (results []*Gv$tenantPxWorkerStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$tenantPxWorkerStat{}).Where("`qc_id` IN (?)", qcIDs).Find(&results).Error
	
	return
}
 
// GetFromSqcID 通过sqc_id获取内容  
func (obj *_Gv$tenantPxWorkerStatMgr) GetFromSqcID(sqcID int64) (results []*Gv$tenantPxWorkerStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$tenantPxWorkerStat{}).Where("`sqc_id` = ?", sqcID).Find(&results).Error
	
	return
}

// GetBatchFromSqcID 批量查找 
func (obj *_Gv$tenantPxWorkerStatMgr) GetBatchFromSqcID(sqcIDs []int64) (results []*Gv$tenantPxWorkerStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$tenantPxWorkerStat{}).Where("`sqc_id` IN (?)", sqcIDs).Find(&results).Error
	
	return
}
 
// GetFromWorkerID 通过worker_id获取内容  
func (obj *_Gv$tenantPxWorkerStatMgr) GetFromWorkerID(workerID int64) (results []*Gv$tenantPxWorkerStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$tenantPxWorkerStat{}).Where("`worker_id` = ?", workerID).Find(&results).Error
	
	return
}

// GetBatchFromWorkerID 批量查找 
func (obj *_Gv$tenantPxWorkerStatMgr) GetBatchFromWorkerID(workerIDs []int64) (results []*Gv$tenantPxWorkerStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$tenantPxWorkerStat{}).Where("`worker_id` IN (?)", workerIDs).Find(&results).Error
	
	return
}
 
// GetFromDfoID 通过dfo_id获取内容  
func (obj *_Gv$tenantPxWorkerStatMgr) GetFromDfoID(dfoID int64) (results []*Gv$tenantPxWorkerStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$tenantPxWorkerStat{}).Where("`dfo_id` = ?", dfoID).Find(&results).Error
	
	return
}

// GetBatchFromDfoID 批量查找 
func (obj *_Gv$tenantPxWorkerStatMgr) GetBatchFromDfoID(dfoIDs []int64) (results []*Gv$tenantPxWorkerStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$tenantPxWorkerStat{}).Where("`dfo_id` IN (?)", dfoIDs).Find(&results).Error
	
	return
}
 
// GetFromStartTime 通过start_time获取内容  
func (obj *_Gv$tenantPxWorkerStatMgr) GetFromStartTime(startTime time.Time) (results []*Gv$tenantPxWorkerStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$tenantPxWorkerStat{}).Where("`start_time` = ?", startTime).Find(&results).Error
	
	return
}

// GetBatchFromStartTime 批量查找 
func (obj *_Gv$tenantPxWorkerStatMgr) GetBatchFromStartTime(startTimes []time.Time) (results []*Gv$tenantPxWorkerStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$tenantPxWorkerStat{}).Where("`start_time` IN (?)", startTimes).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 

 

	

