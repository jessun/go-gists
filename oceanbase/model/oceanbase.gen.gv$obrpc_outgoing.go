package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
"time"	
)	

type _Gv$obrpcOutgoingMgr struct {
	*_BaseMgr
}

// Gv$obrpcOutgoingMgr open func
func Gv$obrpcOutgoingMgr(db *gorm.DB) *_Gv$obrpcOutgoingMgr {
	if db == nil {
		panic(fmt.Errorf("Gv$obrpcOutgoingMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_Gv$obrpcOutgoingMgr{_BaseMgr: &_BaseMgr{DB: db.Table("gv$obrpc_outgoing"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_Gv$obrpcOutgoingMgr) GetTableName() string {
	return "gv$obrpc_outgoing"
}

// Reset 重置gorm会话
func (obj *_Gv$obrpcOutgoingMgr) Reset() *_Gv$obrpcOutgoingMgr {
	obj.New()
	return obj
}

// Get 获取 
func (obj *_Gv$obrpcOutgoingMgr) Get() (result Gv$obrpcOutgoing, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$obrpcOutgoing{}).First(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_Gv$obrpcOutgoingMgr) Gets() (results []*Gv$obrpcOutgoing, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$obrpcOutgoing{}).Find(&results).Error
	
	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_Gv$obrpcOutgoingMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(Gv$obrpcOutgoing{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTenantID TENANT_ID获取 
func (obj *_Gv$obrpcOutgoingMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["TENANT_ID"] = tenantID })
}

// WithIP IP获取 
func (obj *_Gv$obrpcOutgoingMgr) WithIP(ip string) Option {
	return optionFunc(func(o *options) { o.query["IP"] = ip })
}

// WithPort PORT获取 
func (obj *_Gv$obrpcOutgoingMgr) WithPort(port int64) Option {
	return optionFunc(func(o *options) { o.query["PORT"] = port })
}

// WithPcode PCODE获取 
func (obj *_Gv$obrpcOutgoingMgr) WithPcode(pcode int64) Option {
	return optionFunc(func(o *options) { o.query["PCODE"] = pcode })
}

// WithPcodeName PCODE_NAME获取 
func (obj *_Gv$obrpcOutgoingMgr) WithPcodeName(pcodeName string) Option {
	return optionFunc(func(o *options) { o.query["PCODE_NAME"] = pcodeName })
}

// WithCount COUNT获取 
func (obj *_Gv$obrpcOutgoingMgr) WithCount(count int64) Option {
	return optionFunc(func(o *options) { o.query["COUNT"] = count })
}

// WithTotalTime TOTAL_TIME获取 
func (obj *_Gv$obrpcOutgoingMgr) WithTotalTime(totalTime int64) Option {
	return optionFunc(func(o *options) { o.query["TOTAL_TIME"] = totalTime })
}

// WithTotalSize TOTAL_SIZE获取 
func (obj *_Gv$obrpcOutgoingMgr) WithTotalSize(totalSize int64) Option {
	return optionFunc(func(o *options) { o.query["TOTAL_SIZE"] = totalSize })
}

// WithFailure FAILURE获取 
func (obj *_Gv$obrpcOutgoingMgr) WithFailure(failure int64) Option {
	return optionFunc(func(o *options) { o.query["FAILURE"] = failure })
}

// WithTimeout TIMEOUT获取 
func (obj *_Gv$obrpcOutgoingMgr) WithTimeout(timeout int64) Option {
	return optionFunc(func(o *options) { o.query["TIMEOUT"] = timeout })
}

// WithSync SYNC获取 
func (obj *_Gv$obrpcOutgoingMgr) WithSync(sync int64) Option {
	return optionFunc(func(o *options) { o.query["SYNC"] = sync })
}

// WithAsync ASYNC获取 
func (obj *_Gv$obrpcOutgoingMgr) WithAsync(async int64) Option {
	return optionFunc(func(o *options) { o.query["ASYNC"] = async })
}

// WithLastTimestamp LAST_TIMESTAMP获取 
func (obj *_Gv$obrpcOutgoingMgr) WithLastTimestamp(lastTimestamp time.Time) Option {
	return optionFunc(func(o *options) { o.query["LAST_TIMESTAMP"] = lastTimestamp })
}


// GetByOption 功能选项模式获取
func (obj *_Gv$obrpcOutgoingMgr) GetByOption(opts ...Option) (result Gv$obrpcOutgoing, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Gv$obrpcOutgoing{}).Where(options.query).First(&result).Error
	
	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_Gv$obrpcOutgoingMgr) GetByOptions(opts ...Option) (results []*Gv$obrpcOutgoing, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Gv$obrpcOutgoing{}).Where(options.query).Find(&results).Error
	
	return
}


// SelectPage 分页查询
func (obj *_Gv$obrpcOutgoingMgr) SelectPage(page IPage,opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]Gv$obrpcOutgoing,0)
	var count int64 // 统计总的记录数
	query :=  obj.DB.WithContext(obj.ctx).Model(Gv$obrpcOutgoing{}).Where(options.query)
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


// GetFromTenantID 通过TENANT_ID获取内容  
func (obj *_Gv$obrpcOutgoingMgr) GetFromTenantID(tenantID int64) (results []*Gv$obrpcOutgoing, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$obrpcOutgoing{}).Where("`TENANT_ID` = ?", tenantID).Find(&results).Error
	
	return
}

// GetBatchFromTenantID 批量查找 
func (obj *_Gv$obrpcOutgoingMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*Gv$obrpcOutgoing, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$obrpcOutgoing{}).Where("`TENANT_ID` IN (?)", tenantIDs).Find(&results).Error
	
	return
}
 
// GetFromIP 通过IP获取内容  
func (obj *_Gv$obrpcOutgoingMgr) GetFromIP(ip string) (results []*Gv$obrpcOutgoing, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$obrpcOutgoing{}).Where("`IP` = ?", ip).Find(&results).Error
	
	return
}

// GetBatchFromIP 批量查找 
func (obj *_Gv$obrpcOutgoingMgr) GetBatchFromIP(ips []string) (results []*Gv$obrpcOutgoing, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$obrpcOutgoing{}).Where("`IP` IN (?)", ips).Find(&results).Error
	
	return
}
 
// GetFromPort 通过PORT获取内容  
func (obj *_Gv$obrpcOutgoingMgr) GetFromPort(port int64) (results []*Gv$obrpcOutgoing, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$obrpcOutgoing{}).Where("`PORT` = ?", port).Find(&results).Error
	
	return
}

// GetBatchFromPort 批量查找 
func (obj *_Gv$obrpcOutgoingMgr) GetBatchFromPort(ports []int64) (results []*Gv$obrpcOutgoing, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$obrpcOutgoing{}).Where("`PORT` IN (?)", ports).Find(&results).Error
	
	return
}
 
// GetFromPcode 通过PCODE获取内容  
func (obj *_Gv$obrpcOutgoingMgr) GetFromPcode(pcode int64) (results []*Gv$obrpcOutgoing, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$obrpcOutgoing{}).Where("`PCODE` = ?", pcode).Find(&results).Error
	
	return
}

// GetBatchFromPcode 批量查找 
func (obj *_Gv$obrpcOutgoingMgr) GetBatchFromPcode(pcodes []int64) (results []*Gv$obrpcOutgoing, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$obrpcOutgoing{}).Where("`PCODE` IN (?)", pcodes).Find(&results).Error
	
	return
}
 
// GetFromPcodeName 通过PCODE_NAME获取内容  
func (obj *_Gv$obrpcOutgoingMgr) GetFromPcodeName(pcodeName string) (results []*Gv$obrpcOutgoing, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$obrpcOutgoing{}).Where("`PCODE_NAME` = ?", pcodeName).Find(&results).Error
	
	return
}

// GetBatchFromPcodeName 批量查找 
func (obj *_Gv$obrpcOutgoingMgr) GetBatchFromPcodeName(pcodeNames []string) (results []*Gv$obrpcOutgoing, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$obrpcOutgoing{}).Where("`PCODE_NAME` IN (?)", pcodeNames).Find(&results).Error
	
	return
}
 
// GetFromCount 通过COUNT获取内容  
func (obj *_Gv$obrpcOutgoingMgr) GetFromCount(count int64) (results []*Gv$obrpcOutgoing, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$obrpcOutgoing{}).Where("`COUNT` = ?", count).Find(&results).Error
	
	return
}

// GetBatchFromCount 批量查找 
func (obj *_Gv$obrpcOutgoingMgr) GetBatchFromCount(counts []int64) (results []*Gv$obrpcOutgoing, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$obrpcOutgoing{}).Where("`COUNT` IN (?)", counts).Find(&results).Error
	
	return
}
 
// GetFromTotalTime 通过TOTAL_TIME获取内容  
func (obj *_Gv$obrpcOutgoingMgr) GetFromTotalTime(totalTime int64) (results []*Gv$obrpcOutgoing, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$obrpcOutgoing{}).Where("`TOTAL_TIME` = ?", totalTime).Find(&results).Error
	
	return
}

// GetBatchFromTotalTime 批量查找 
func (obj *_Gv$obrpcOutgoingMgr) GetBatchFromTotalTime(totalTimes []int64) (results []*Gv$obrpcOutgoing, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$obrpcOutgoing{}).Where("`TOTAL_TIME` IN (?)", totalTimes).Find(&results).Error
	
	return
}
 
// GetFromTotalSize 通过TOTAL_SIZE获取内容  
func (obj *_Gv$obrpcOutgoingMgr) GetFromTotalSize(totalSize int64) (results []*Gv$obrpcOutgoing, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$obrpcOutgoing{}).Where("`TOTAL_SIZE` = ?", totalSize).Find(&results).Error
	
	return
}

// GetBatchFromTotalSize 批量查找 
func (obj *_Gv$obrpcOutgoingMgr) GetBatchFromTotalSize(totalSizes []int64) (results []*Gv$obrpcOutgoing, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$obrpcOutgoing{}).Where("`TOTAL_SIZE` IN (?)", totalSizes).Find(&results).Error
	
	return
}
 
// GetFromFailure 通过FAILURE获取内容  
func (obj *_Gv$obrpcOutgoingMgr) GetFromFailure(failure int64) (results []*Gv$obrpcOutgoing, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$obrpcOutgoing{}).Where("`FAILURE` = ?", failure).Find(&results).Error
	
	return
}

// GetBatchFromFailure 批量查找 
func (obj *_Gv$obrpcOutgoingMgr) GetBatchFromFailure(failures []int64) (results []*Gv$obrpcOutgoing, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$obrpcOutgoing{}).Where("`FAILURE` IN (?)", failures).Find(&results).Error
	
	return
}
 
// GetFromTimeout 通过TIMEOUT获取内容  
func (obj *_Gv$obrpcOutgoingMgr) GetFromTimeout(timeout int64) (results []*Gv$obrpcOutgoing, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$obrpcOutgoing{}).Where("`TIMEOUT` = ?", timeout).Find(&results).Error
	
	return
}

// GetBatchFromTimeout 批量查找 
func (obj *_Gv$obrpcOutgoingMgr) GetBatchFromTimeout(timeouts []int64) (results []*Gv$obrpcOutgoing, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$obrpcOutgoing{}).Where("`TIMEOUT` IN (?)", timeouts).Find(&results).Error
	
	return
}
 
// GetFromSync 通过SYNC获取内容  
func (obj *_Gv$obrpcOutgoingMgr) GetFromSync(sync int64) (results []*Gv$obrpcOutgoing, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$obrpcOutgoing{}).Where("`SYNC` = ?", sync).Find(&results).Error
	
	return
}

// GetBatchFromSync 批量查找 
func (obj *_Gv$obrpcOutgoingMgr) GetBatchFromSync(syncs []int64) (results []*Gv$obrpcOutgoing, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$obrpcOutgoing{}).Where("`SYNC` IN (?)", syncs).Find(&results).Error
	
	return
}
 
// GetFromAsync 通过ASYNC获取内容  
func (obj *_Gv$obrpcOutgoingMgr) GetFromAsync(async int64) (results []*Gv$obrpcOutgoing, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$obrpcOutgoing{}).Where("`ASYNC` = ?", async).Find(&results).Error
	
	return
}

// GetBatchFromAsync 批量查找 
func (obj *_Gv$obrpcOutgoingMgr) GetBatchFromAsync(asyncs []int64) (results []*Gv$obrpcOutgoing, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$obrpcOutgoing{}).Where("`ASYNC` IN (?)", asyncs).Find(&results).Error
	
	return
}
 
// GetFromLastTimestamp 通过LAST_TIMESTAMP获取内容  
func (obj *_Gv$obrpcOutgoingMgr) GetFromLastTimestamp(lastTimestamp time.Time) (results []*Gv$obrpcOutgoing, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$obrpcOutgoing{}).Where("`LAST_TIMESTAMP` = ?", lastTimestamp).Find(&results).Error
	
	return
}

// GetBatchFromLastTimestamp 批量查找 
func (obj *_Gv$obrpcOutgoingMgr) GetBatchFromLastTimestamp(lastTimestamps []time.Time) (results []*Gv$obrpcOutgoing, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$obrpcOutgoing{}).Where("`LAST_TIMESTAMP` IN (?)", lastTimestamps).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 

 

	

