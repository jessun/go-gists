package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
)	

type _V$latchMgr struct {
	*_BaseMgr
}

// V$latchMgr open func
func V$latchMgr(db *gorm.DB) *_V$latchMgr {
	if db == nil {
		panic(fmt.Errorf("V$latchMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_V$latchMgr{_BaseMgr: &_BaseMgr{DB: db.Table("v$latch"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_V$latchMgr) GetTableName() string {
	return "v$latch"
}

// Reset 重置gorm会话
func (obj *_V$latchMgr) Reset() *_V$latchMgr {
	obj.New()
	return obj
}

// Get 获取 
func (obj *_V$latchMgr) Get() (result V$latch, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$latch{}).First(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_V$latchMgr) Gets() (results []*V$latch, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$latch{}).Find(&results).Error
	
	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_V$latchMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(V$latch{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithConID CON_ID获取 
func (obj *_V$latchMgr) WithConID(conID int64) Option {
	return optionFunc(func(o *options) { o.query["CON_ID"] = conID })
}

// WithSvrIP SVR_IP获取 
func (obj *_V$latchMgr) WithSvrIP(svrIP string) Option {
	return optionFunc(func(o *options) { o.query["SVR_IP"] = svrIP })
}

// WithSvrPort SVR_PORT获取 
func (obj *_V$latchMgr) WithSvrPort(svrPort int64) Option {
	return optionFunc(func(o *options) { o.query["SVR_PORT"] = svrPort })
}

// WithAddr ADDR获取 
func (obj *_V$latchMgr) WithAddr(addr string) Option {
	return optionFunc(func(o *options) { o.query["ADDR"] = addr })
}

// WithLatch# LATCH#获取 
func (obj *_V$latchMgr) WithLatch#(latch# int64) Option {
	return optionFunc(func(o *options) { o.query["LATCH#"] = latch# })
}

// WithLevel# LEVEL#获取 
func (obj *_V$latchMgr) WithLevel#(level# int64) Option {
	return optionFunc(func(o *options) { o.query["LEVEL#"] = level# })
}

// WithName NAME获取 
func (obj *_V$latchMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["NAME"] = name })
}

// WithHash HASH获取 
func (obj *_V$latchMgr) WithHash(hash int64) Option {
	return optionFunc(func(o *options) { o.query["HASH"] = hash })
}

// WithGets GETS获取 
func (obj *_V$latchMgr) WithGets(gets int64) Option {
	return optionFunc(func(o *options) { o.query["GETS"] = gets })
}

// WithMisses MISSES获取 
func (obj *_V$latchMgr) WithMisses(misses int64) Option {
	return optionFunc(func(o *options) { o.query["MISSES"] = misses })
}

// WithSleeps SLEEPS获取 
func (obj *_V$latchMgr) WithSleeps(sleeps int64) Option {
	return optionFunc(func(o *options) { o.query["SLEEPS"] = sleeps })
}

// WithImmediateGets IMMEDIATE_GETS获取 
func (obj *_V$latchMgr) WithImmediateGets(immediateGets int64) Option {
	return optionFunc(func(o *options) { o.query["IMMEDIATE_GETS"] = immediateGets })
}

// WithImmediateMisses IMMEDIATE_MISSES获取 
func (obj *_V$latchMgr) WithImmediateMisses(immediateMisses int64) Option {
	return optionFunc(func(o *options) { o.query["IMMEDIATE_MISSES"] = immediateMisses })
}

// WithSpinGets SPIN_GETS获取 
func (obj *_V$latchMgr) WithSpinGets(spinGets int64) Option {
	return optionFunc(func(o *options) { o.query["SPIN_GETS"] = spinGets })
}

// WithWaitTime WAIT_TIME获取 
func (obj *_V$latchMgr) WithWaitTime(waitTime int64) Option {
	return optionFunc(func(o *options) { o.query["WAIT_TIME"] = waitTime })
}


// GetByOption 功能选项模式获取
func (obj *_V$latchMgr) GetByOption(opts ...Option) (result V$latch, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(V$latch{}).Where(options.query).First(&result).Error
	
	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_V$latchMgr) GetByOptions(opts ...Option) (results []*V$latch, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(V$latch{}).Where(options.query).Find(&results).Error
	
	return
}


// SelectPage 分页查询
func (obj *_V$latchMgr) SelectPage(page IPage,opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]V$latch,0)
	var count int64 // 统计总的记录数
	query :=  obj.DB.WithContext(obj.ctx).Model(V$latch{}).Where(options.query)
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


// GetFromConID 通过CON_ID获取内容  
func (obj *_V$latchMgr) GetFromConID(conID int64) (results []*V$latch, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$latch{}).Where("`CON_ID` = ?", conID).Find(&results).Error
	
	return
}

// GetBatchFromConID 批量查找 
func (obj *_V$latchMgr) GetBatchFromConID(conIDs []int64) (results []*V$latch, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$latch{}).Where("`CON_ID` IN (?)", conIDs).Find(&results).Error
	
	return
}
 
// GetFromSvrIP 通过SVR_IP获取内容  
func (obj *_V$latchMgr) GetFromSvrIP(svrIP string) (results []*V$latch, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$latch{}).Where("`SVR_IP` = ?", svrIP).Find(&results).Error
	
	return
}

// GetBatchFromSvrIP 批量查找 
func (obj *_V$latchMgr) GetBatchFromSvrIP(svrIPs []string) (results []*V$latch, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$latch{}).Where("`SVR_IP` IN (?)", svrIPs).Find(&results).Error
	
	return
}
 
// GetFromSvrPort 通过SVR_PORT获取内容  
func (obj *_V$latchMgr) GetFromSvrPort(svrPort int64) (results []*V$latch, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$latch{}).Where("`SVR_PORT` = ?", svrPort).Find(&results).Error
	
	return
}

// GetBatchFromSvrPort 批量查找 
func (obj *_V$latchMgr) GetBatchFromSvrPort(svrPorts []int64) (results []*V$latch, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$latch{}).Where("`SVR_PORT` IN (?)", svrPorts).Find(&results).Error
	
	return
}
 
// GetFromAddr 通过ADDR获取内容  
func (obj *_V$latchMgr) GetFromAddr(addr string) (results []*V$latch, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$latch{}).Where("`ADDR` = ?", addr).Find(&results).Error
	
	return
}

// GetBatchFromAddr 批量查找 
func (obj *_V$latchMgr) GetBatchFromAddr(addrs []string) (results []*V$latch, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$latch{}).Where("`ADDR` IN (?)", addrs).Find(&results).Error
	
	return
}
 
// GetFromLatch# 通过LATCH#获取内容  
func (obj *_V$latchMgr) GetFromLatch#(latch# int64) (results []*V$latch, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$latch{}).Where("`LATCH#` = ?", latch#).Find(&results).Error
	
	return
}

// GetBatchFromLatch# 批量查找 
func (obj *_V$latchMgr) GetBatchFromLatch#(latch#s []int64) (results []*V$latch, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$latch{}).Where("`LATCH#` IN (?)", latch#s).Find(&results).Error
	
	return
}
 
// GetFromLevel# 通过LEVEL#获取内容  
func (obj *_V$latchMgr) GetFromLevel#(level# int64) (results []*V$latch, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$latch{}).Where("`LEVEL#` = ?", level#).Find(&results).Error
	
	return
}

// GetBatchFromLevel# 批量查找 
func (obj *_V$latchMgr) GetBatchFromLevel#(level#s []int64) (results []*V$latch, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$latch{}).Where("`LEVEL#` IN (?)", level#s).Find(&results).Error
	
	return
}
 
// GetFromName 通过NAME获取内容  
func (obj *_V$latchMgr) GetFromName(name string) (results []*V$latch, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$latch{}).Where("`NAME` = ?", name).Find(&results).Error
	
	return
}

// GetBatchFromName 批量查找 
func (obj *_V$latchMgr) GetBatchFromName(names []string) (results []*V$latch, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$latch{}).Where("`NAME` IN (?)", names).Find(&results).Error
	
	return
}
 
// GetFromHash 通过HASH获取内容  
func (obj *_V$latchMgr) GetFromHash(hash int64) (results []*V$latch, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$latch{}).Where("`HASH` = ?", hash).Find(&results).Error
	
	return
}

// GetBatchFromHash 批量查找 
func (obj *_V$latchMgr) GetBatchFromHash(hashs []int64) (results []*V$latch, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$latch{}).Where("`HASH` IN (?)", hashs).Find(&results).Error
	
	return
}
 
// GetFromGets 通过GETS获取内容  
func (obj *_V$latchMgr) GetFromGets(gets int64) (results []*V$latch, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$latch{}).Where("`GETS` = ?", gets).Find(&results).Error
	
	return
}

// GetBatchFromGets 批量查找 
func (obj *_V$latchMgr) GetBatchFromGets(getss []int64) (results []*V$latch, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$latch{}).Where("`GETS` IN (?)", getss).Find(&results).Error
	
	return
}
 
// GetFromMisses 通过MISSES获取内容  
func (obj *_V$latchMgr) GetFromMisses(misses int64) (results []*V$latch, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$latch{}).Where("`MISSES` = ?", misses).Find(&results).Error
	
	return
}

// GetBatchFromMisses 批量查找 
func (obj *_V$latchMgr) GetBatchFromMisses(missess []int64) (results []*V$latch, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$latch{}).Where("`MISSES` IN (?)", missess).Find(&results).Error
	
	return
}
 
// GetFromSleeps 通过SLEEPS获取内容  
func (obj *_V$latchMgr) GetFromSleeps(sleeps int64) (results []*V$latch, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$latch{}).Where("`SLEEPS` = ?", sleeps).Find(&results).Error
	
	return
}

// GetBatchFromSleeps 批量查找 
func (obj *_V$latchMgr) GetBatchFromSleeps(sleepss []int64) (results []*V$latch, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$latch{}).Where("`SLEEPS` IN (?)", sleepss).Find(&results).Error
	
	return
}
 
// GetFromImmediateGets 通过IMMEDIATE_GETS获取内容  
func (obj *_V$latchMgr) GetFromImmediateGets(immediateGets int64) (results []*V$latch, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$latch{}).Where("`IMMEDIATE_GETS` = ?", immediateGets).Find(&results).Error
	
	return
}

// GetBatchFromImmediateGets 批量查找 
func (obj *_V$latchMgr) GetBatchFromImmediateGets(immediateGetss []int64) (results []*V$latch, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$latch{}).Where("`IMMEDIATE_GETS` IN (?)", immediateGetss).Find(&results).Error
	
	return
}
 
// GetFromImmediateMisses 通过IMMEDIATE_MISSES获取内容  
func (obj *_V$latchMgr) GetFromImmediateMisses(immediateMisses int64) (results []*V$latch, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$latch{}).Where("`IMMEDIATE_MISSES` = ?", immediateMisses).Find(&results).Error
	
	return
}

// GetBatchFromImmediateMisses 批量查找 
func (obj *_V$latchMgr) GetBatchFromImmediateMisses(immediateMissess []int64) (results []*V$latch, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$latch{}).Where("`IMMEDIATE_MISSES` IN (?)", immediateMissess).Find(&results).Error
	
	return
}
 
// GetFromSpinGets 通过SPIN_GETS获取内容  
func (obj *_V$latchMgr) GetFromSpinGets(spinGets int64) (results []*V$latch, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$latch{}).Where("`SPIN_GETS` = ?", spinGets).Find(&results).Error
	
	return
}

// GetBatchFromSpinGets 批量查找 
func (obj *_V$latchMgr) GetBatchFromSpinGets(spinGetss []int64) (results []*V$latch, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$latch{}).Where("`SPIN_GETS` IN (?)", spinGetss).Find(&results).Error
	
	return
}
 
// GetFromWaitTime 通过WAIT_TIME获取内容  
func (obj *_V$latchMgr) GetFromWaitTime(waitTime int64) (results []*V$latch, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$latch{}).Where("`WAIT_TIME` = ?", waitTime).Find(&results).Error
	
	return
}

// GetBatchFromWaitTime 批量查找 
func (obj *_V$latchMgr) GetBatchFromWaitTime(waitTimes []int64) (results []*V$latch, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$latch{}).Where("`WAIT_TIME` IN (?)", waitTimes).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 

 

	

