package	model	
import (	
"gorm.io/gorm"	
"time"	
"fmt"	
"context"	
)	

type _V$obrpcIncomingMgr struct {
	*_BaseMgr
}

// V$obrpcIncomingMgr open func
func V$obrpcIncomingMgr(db *gorm.DB) *_V$obrpcIncomingMgr {
	if db == nil {
		panic(fmt.Errorf("V$obrpcIncomingMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_V$obrpcIncomingMgr{_BaseMgr: &_BaseMgr{DB: db.Table("v$obrpc_incoming"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_V$obrpcIncomingMgr) GetTableName() string {
	return "v$obrpc_incoming"
}

// Reset 重置gorm会话
func (obj *_V$obrpcIncomingMgr) Reset() *_V$obrpcIncomingMgr {
	obj.New()
	return obj
}

// Get 获取 
func (obj *_V$obrpcIncomingMgr) Get() (result V$obrpcIncoming, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$obrpcIncoming{}).First(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_V$obrpcIncomingMgr) Gets() (results []*V$obrpcIncoming, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$obrpcIncoming{}).Find(&results).Error
	
	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_V$obrpcIncomingMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(V$obrpcIncoming{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTenantID TENANT_ID获取 
func (obj *_V$obrpcIncomingMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["TENANT_ID"] = tenantID })
}

// WithIP IP获取 
func (obj *_V$obrpcIncomingMgr) WithIP(ip string) Option {
	return optionFunc(func(o *options) { o.query["IP"] = ip })
}

// WithPort PORT获取 
func (obj *_V$obrpcIncomingMgr) WithPort(port int64) Option {
	return optionFunc(func(o *options) { o.query["PORT"] = port })
}

// WithPcode PCODE获取 
func (obj *_V$obrpcIncomingMgr) WithPcode(pcode int64) Option {
	return optionFunc(func(o *options) { o.query["PCODE"] = pcode })
}

// WithPcodeName PCODE_NAME获取 
func (obj *_V$obrpcIncomingMgr) WithPcodeName(pcodeName string) Option {
	return optionFunc(func(o *options) { o.query["PCODE_NAME"] = pcodeName })
}

// WithCount COUNT获取 
func (obj *_V$obrpcIncomingMgr) WithCount(count int64) Option {
	return optionFunc(func(o *options) { o.query["COUNT"] = count })
}

// WithTotalSize TOTAL_SIZE获取 
func (obj *_V$obrpcIncomingMgr) WithTotalSize(totalSize int64) Option {
	return optionFunc(func(o *options) { o.query["TOTAL_SIZE"] = totalSize })
}

// WithNetTime NET_TIME获取 
func (obj *_V$obrpcIncomingMgr) WithNetTime(netTime int64) Option {
	return optionFunc(func(o *options) { o.query["NET_TIME"] = netTime })
}

// WithWaitTime WAIT_TIME获取 
func (obj *_V$obrpcIncomingMgr) WithWaitTime(waitTime int64) Option {
	return optionFunc(func(o *options) { o.query["WAIT_TIME"] = waitTime })
}

// WithQueueTime QUEUE_TIME获取 
func (obj *_V$obrpcIncomingMgr) WithQueueTime(queueTime int64) Option {
	return optionFunc(func(o *options) { o.query["QUEUE_TIME"] = queueTime })
}

// WithProcessTime PROCESS_TIME获取 
func (obj *_V$obrpcIncomingMgr) WithProcessTime(processTime int64) Option {
	return optionFunc(func(o *options) { o.query["PROCESS_TIME"] = processTime })
}

// WithLastTimestamp LAST_TIMESTAMP获取 
func (obj *_V$obrpcIncomingMgr) WithLastTimestamp(lastTimestamp time.Time) Option {
	return optionFunc(func(o *options) { o.query["LAST_TIMESTAMP"] = lastTimestamp })
}


// GetByOption 功能选项模式获取
func (obj *_V$obrpcIncomingMgr) GetByOption(opts ...Option) (result V$obrpcIncoming, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(V$obrpcIncoming{}).Where(options.query).First(&result).Error
	
	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_V$obrpcIncomingMgr) GetByOptions(opts ...Option) (results []*V$obrpcIncoming, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(V$obrpcIncoming{}).Where(options.query).Find(&results).Error
	
	return
}


// SelectPage 分页查询
func (obj *_V$obrpcIncomingMgr) SelectPage(page IPage,opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]V$obrpcIncoming,0)
	var count int64 // 统计总的记录数
	query :=  obj.DB.WithContext(obj.ctx).Model(V$obrpcIncoming{}).Where(options.query)
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
func (obj *_V$obrpcIncomingMgr) GetFromTenantID(tenantID int64) (results []*V$obrpcIncoming, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$obrpcIncoming{}).Where("`TENANT_ID` = ?", tenantID).Find(&results).Error
	
	return
}

// GetBatchFromTenantID 批量查找 
func (obj *_V$obrpcIncomingMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*V$obrpcIncoming, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$obrpcIncoming{}).Where("`TENANT_ID` IN (?)", tenantIDs).Find(&results).Error
	
	return
}
 
// GetFromIP 通过IP获取内容  
func (obj *_V$obrpcIncomingMgr) GetFromIP(ip string) (results []*V$obrpcIncoming, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$obrpcIncoming{}).Where("`IP` = ?", ip).Find(&results).Error
	
	return
}

// GetBatchFromIP 批量查找 
func (obj *_V$obrpcIncomingMgr) GetBatchFromIP(ips []string) (results []*V$obrpcIncoming, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$obrpcIncoming{}).Where("`IP` IN (?)", ips).Find(&results).Error
	
	return
}
 
// GetFromPort 通过PORT获取内容  
func (obj *_V$obrpcIncomingMgr) GetFromPort(port int64) (results []*V$obrpcIncoming, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$obrpcIncoming{}).Where("`PORT` = ?", port).Find(&results).Error
	
	return
}

// GetBatchFromPort 批量查找 
func (obj *_V$obrpcIncomingMgr) GetBatchFromPort(ports []int64) (results []*V$obrpcIncoming, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$obrpcIncoming{}).Where("`PORT` IN (?)", ports).Find(&results).Error
	
	return
}
 
// GetFromPcode 通过PCODE获取内容  
func (obj *_V$obrpcIncomingMgr) GetFromPcode(pcode int64) (results []*V$obrpcIncoming, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$obrpcIncoming{}).Where("`PCODE` = ?", pcode).Find(&results).Error
	
	return
}

// GetBatchFromPcode 批量查找 
func (obj *_V$obrpcIncomingMgr) GetBatchFromPcode(pcodes []int64) (results []*V$obrpcIncoming, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$obrpcIncoming{}).Where("`PCODE` IN (?)", pcodes).Find(&results).Error
	
	return
}
 
// GetFromPcodeName 通过PCODE_NAME获取内容  
func (obj *_V$obrpcIncomingMgr) GetFromPcodeName(pcodeName string) (results []*V$obrpcIncoming, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$obrpcIncoming{}).Where("`PCODE_NAME` = ?", pcodeName).Find(&results).Error
	
	return
}

// GetBatchFromPcodeName 批量查找 
func (obj *_V$obrpcIncomingMgr) GetBatchFromPcodeName(pcodeNames []string) (results []*V$obrpcIncoming, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$obrpcIncoming{}).Where("`PCODE_NAME` IN (?)", pcodeNames).Find(&results).Error
	
	return
}
 
// GetFromCount 通过COUNT获取内容  
func (obj *_V$obrpcIncomingMgr) GetFromCount(count int64) (results []*V$obrpcIncoming, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$obrpcIncoming{}).Where("`COUNT` = ?", count).Find(&results).Error
	
	return
}

// GetBatchFromCount 批量查找 
func (obj *_V$obrpcIncomingMgr) GetBatchFromCount(counts []int64) (results []*V$obrpcIncoming, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$obrpcIncoming{}).Where("`COUNT` IN (?)", counts).Find(&results).Error
	
	return
}
 
// GetFromTotalSize 通过TOTAL_SIZE获取内容  
func (obj *_V$obrpcIncomingMgr) GetFromTotalSize(totalSize int64) (results []*V$obrpcIncoming, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$obrpcIncoming{}).Where("`TOTAL_SIZE` = ?", totalSize).Find(&results).Error
	
	return
}

// GetBatchFromTotalSize 批量查找 
func (obj *_V$obrpcIncomingMgr) GetBatchFromTotalSize(totalSizes []int64) (results []*V$obrpcIncoming, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$obrpcIncoming{}).Where("`TOTAL_SIZE` IN (?)", totalSizes).Find(&results).Error
	
	return
}
 
// GetFromNetTime 通过NET_TIME获取内容  
func (obj *_V$obrpcIncomingMgr) GetFromNetTime(netTime int64) (results []*V$obrpcIncoming, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$obrpcIncoming{}).Where("`NET_TIME` = ?", netTime).Find(&results).Error
	
	return
}

// GetBatchFromNetTime 批量查找 
func (obj *_V$obrpcIncomingMgr) GetBatchFromNetTime(netTimes []int64) (results []*V$obrpcIncoming, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$obrpcIncoming{}).Where("`NET_TIME` IN (?)", netTimes).Find(&results).Error
	
	return
}
 
// GetFromWaitTime 通过WAIT_TIME获取内容  
func (obj *_V$obrpcIncomingMgr) GetFromWaitTime(waitTime int64) (results []*V$obrpcIncoming, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$obrpcIncoming{}).Where("`WAIT_TIME` = ?", waitTime).Find(&results).Error
	
	return
}

// GetBatchFromWaitTime 批量查找 
func (obj *_V$obrpcIncomingMgr) GetBatchFromWaitTime(waitTimes []int64) (results []*V$obrpcIncoming, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$obrpcIncoming{}).Where("`WAIT_TIME` IN (?)", waitTimes).Find(&results).Error
	
	return
}
 
// GetFromQueueTime 通过QUEUE_TIME获取内容  
func (obj *_V$obrpcIncomingMgr) GetFromQueueTime(queueTime int64) (results []*V$obrpcIncoming, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$obrpcIncoming{}).Where("`QUEUE_TIME` = ?", queueTime).Find(&results).Error
	
	return
}

// GetBatchFromQueueTime 批量查找 
func (obj *_V$obrpcIncomingMgr) GetBatchFromQueueTime(queueTimes []int64) (results []*V$obrpcIncoming, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$obrpcIncoming{}).Where("`QUEUE_TIME` IN (?)", queueTimes).Find(&results).Error
	
	return
}
 
// GetFromProcessTime 通过PROCESS_TIME获取内容  
func (obj *_V$obrpcIncomingMgr) GetFromProcessTime(processTime int64) (results []*V$obrpcIncoming, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$obrpcIncoming{}).Where("`PROCESS_TIME` = ?", processTime).Find(&results).Error
	
	return
}

// GetBatchFromProcessTime 批量查找 
func (obj *_V$obrpcIncomingMgr) GetBatchFromProcessTime(processTimes []int64) (results []*V$obrpcIncoming, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$obrpcIncoming{}).Where("`PROCESS_TIME` IN (?)", processTimes).Find(&results).Error
	
	return
}
 
// GetFromLastTimestamp 通过LAST_TIMESTAMP获取内容  
func (obj *_V$obrpcIncomingMgr) GetFromLastTimestamp(lastTimestamp time.Time) (results []*V$obrpcIncoming, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$obrpcIncoming{}).Where("`LAST_TIMESTAMP` = ?", lastTimestamp).Find(&results).Error
	
	return
}

// GetBatchFromLastTimestamp 批量查找 
func (obj *_V$obrpcIncomingMgr) GetBatchFromLastTimestamp(lastTimestamps []time.Time) (results []*V$obrpcIncoming, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$obrpcIncoming{}).Where("`LAST_TIMESTAMP` IN (?)", lastTimestamps).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 

 

	

