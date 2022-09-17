package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
)	

type _Gv$latchMgr struct {
	*_BaseMgr
}

// Gv$latchMgr open func
func Gv$latchMgr(db *gorm.DB) *_Gv$latchMgr {
	if db == nil {
		panic(fmt.Errorf("Gv$latchMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_Gv$latchMgr{_BaseMgr: &_BaseMgr{DB: db.Table("gv$latch"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_Gv$latchMgr) GetTableName() string {
	return "gv$latch"
}

// Reset 重置gorm会话
func (obj *_Gv$latchMgr) Reset() *_Gv$latchMgr {
	obj.New()
	return obj
}

// Get 获取 
func (obj *_Gv$latchMgr) Get() (result Gv$latch, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$latch{}).First(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_Gv$latchMgr) Gets() (results []*Gv$latch, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$latch{}).Find(&results).Error
	
	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_Gv$latchMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(Gv$latch{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithConID CON_ID获取 
func (obj *_Gv$latchMgr) WithConID(conID int64) Option {
	return optionFunc(func(o *options) { o.query["CON_ID"] = conID })
}

// WithSvrIP SVR_IP获取 
func (obj *_Gv$latchMgr) WithSvrIP(svrIP string) Option {
	return optionFunc(func(o *options) { o.query["SVR_IP"] = svrIP })
}

// WithSvrPort SVR_PORT获取 
func (obj *_Gv$latchMgr) WithSvrPort(svrPort int64) Option {
	return optionFunc(func(o *options) { o.query["SVR_PORT"] = svrPort })
}

// WithAddr ADDR获取 
func (obj *_Gv$latchMgr) WithAddr(addr string) Option {
	return optionFunc(func(o *options) { o.query["ADDR"] = addr })
}

// WithLatch# LATCH#获取 
func (obj *_Gv$latchMgr) WithLatch#(latch# int64) Option {
	return optionFunc(func(o *options) { o.query["LATCH#"] = latch# })
}

// WithLevel# LEVEL#获取 
func (obj *_Gv$latchMgr) WithLevel#(level# int64) Option {
	return optionFunc(func(o *options) { o.query["LEVEL#"] = level# })
}

// WithName NAME获取 
func (obj *_Gv$latchMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["NAME"] = name })
}

// WithHash HASH获取 
func (obj *_Gv$latchMgr) WithHash(hash int64) Option {
	return optionFunc(func(o *options) { o.query["HASH"] = hash })
}

// WithGets GETS获取 
func (obj *_Gv$latchMgr) WithGets(gets int64) Option {
	return optionFunc(func(o *options) { o.query["GETS"] = gets })
}

// WithMisses MISSES获取 
func (obj *_Gv$latchMgr) WithMisses(misses int64) Option {
	return optionFunc(func(o *options) { o.query["MISSES"] = misses })
}

// WithSleeps SLEEPS获取 
func (obj *_Gv$latchMgr) WithSleeps(sleeps int64) Option {
	return optionFunc(func(o *options) { o.query["SLEEPS"] = sleeps })
}

// WithImmediateGets IMMEDIATE_GETS获取 
func (obj *_Gv$latchMgr) WithImmediateGets(immediateGets int64) Option {
	return optionFunc(func(o *options) { o.query["IMMEDIATE_GETS"] = immediateGets })
}

// WithImmediateMisses IMMEDIATE_MISSES获取 
func (obj *_Gv$latchMgr) WithImmediateMisses(immediateMisses int64) Option {
	return optionFunc(func(o *options) { o.query["IMMEDIATE_MISSES"] = immediateMisses })
}

// WithSpinGets SPIN_GETS获取 
func (obj *_Gv$latchMgr) WithSpinGets(spinGets int64) Option {
	return optionFunc(func(o *options) { o.query["SPIN_GETS"] = spinGets })
}

// WithWaitTime WAIT_TIME获取 
func (obj *_Gv$latchMgr) WithWaitTime(waitTime int64) Option {
	return optionFunc(func(o *options) { o.query["WAIT_TIME"] = waitTime })
}


// GetByOption 功能选项模式获取
func (obj *_Gv$latchMgr) GetByOption(opts ...Option) (result Gv$latch, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Gv$latch{}).Where(options.query).First(&result).Error
	
	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_Gv$latchMgr) GetByOptions(opts ...Option) (results []*Gv$latch, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Gv$latch{}).Where(options.query).Find(&results).Error
	
	return
}


// SelectPage 分页查询
func (obj *_Gv$latchMgr) SelectPage(page IPage,opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]Gv$latch,0)
	var count int64 // 统计总的记录数
	query :=  obj.DB.WithContext(obj.ctx).Model(Gv$latch{}).Where(options.query)
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
func (obj *_Gv$latchMgr) GetFromConID(conID int64) (results []*Gv$latch, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$latch{}).Where("`CON_ID` = ?", conID).Find(&results).Error
	
	return
}

// GetBatchFromConID 批量查找 
func (obj *_Gv$latchMgr) GetBatchFromConID(conIDs []int64) (results []*Gv$latch, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$latch{}).Where("`CON_ID` IN (?)", conIDs).Find(&results).Error
	
	return
}
 
// GetFromSvrIP 通过SVR_IP获取内容  
func (obj *_Gv$latchMgr) GetFromSvrIP(svrIP string) (results []*Gv$latch, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$latch{}).Where("`SVR_IP` = ?", svrIP).Find(&results).Error
	
	return
}

// GetBatchFromSvrIP 批量查找 
func (obj *_Gv$latchMgr) GetBatchFromSvrIP(svrIPs []string) (results []*Gv$latch, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$latch{}).Where("`SVR_IP` IN (?)", svrIPs).Find(&results).Error
	
	return
}
 
// GetFromSvrPort 通过SVR_PORT获取内容  
func (obj *_Gv$latchMgr) GetFromSvrPort(svrPort int64) (results []*Gv$latch, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$latch{}).Where("`SVR_PORT` = ?", svrPort).Find(&results).Error
	
	return
}

// GetBatchFromSvrPort 批量查找 
func (obj *_Gv$latchMgr) GetBatchFromSvrPort(svrPorts []int64) (results []*Gv$latch, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$latch{}).Where("`SVR_PORT` IN (?)", svrPorts).Find(&results).Error
	
	return
}
 
// GetFromAddr 通过ADDR获取内容  
func (obj *_Gv$latchMgr) GetFromAddr(addr string) (results []*Gv$latch, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$latch{}).Where("`ADDR` = ?", addr).Find(&results).Error
	
	return
}

// GetBatchFromAddr 批量查找 
func (obj *_Gv$latchMgr) GetBatchFromAddr(addrs []string) (results []*Gv$latch, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$latch{}).Where("`ADDR` IN (?)", addrs).Find(&results).Error
	
	return
}
 
// GetFromLatch# 通过LATCH#获取内容  
func (obj *_Gv$latchMgr) GetFromLatch#(latch# int64) (results []*Gv$latch, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$latch{}).Where("`LATCH#` = ?", latch#).Find(&results).Error
	
	return
}

// GetBatchFromLatch# 批量查找 
func (obj *_Gv$latchMgr) GetBatchFromLatch#(latch#s []int64) (results []*Gv$latch, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$latch{}).Where("`LATCH#` IN (?)", latch#s).Find(&results).Error
	
	return
}
 
// GetFromLevel# 通过LEVEL#获取内容  
func (obj *_Gv$latchMgr) GetFromLevel#(level# int64) (results []*Gv$latch, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$latch{}).Where("`LEVEL#` = ?", level#).Find(&results).Error
	
	return
}

// GetBatchFromLevel# 批量查找 
func (obj *_Gv$latchMgr) GetBatchFromLevel#(level#s []int64) (results []*Gv$latch, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$latch{}).Where("`LEVEL#` IN (?)", level#s).Find(&results).Error
	
	return
}
 
// GetFromName 通过NAME获取内容  
func (obj *_Gv$latchMgr) GetFromName(name string) (results []*Gv$latch, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$latch{}).Where("`NAME` = ?", name).Find(&results).Error
	
	return
}

// GetBatchFromName 批量查找 
func (obj *_Gv$latchMgr) GetBatchFromName(names []string) (results []*Gv$latch, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$latch{}).Where("`NAME` IN (?)", names).Find(&results).Error
	
	return
}
 
// GetFromHash 通过HASH获取内容  
func (obj *_Gv$latchMgr) GetFromHash(hash int64) (results []*Gv$latch, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$latch{}).Where("`HASH` = ?", hash).Find(&results).Error
	
	return
}

// GetBatchFromHash 批量查找 
func (obj *_Gv$latchMgr) GetBatchFromHash(hashs []int64) (results []*Gv$latch, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$latch{}).Where("`HASH` IN (?)", hashs).Find(&results).Error
	
	return
}
 
// GetFromGets 通过GETS获取内容  
func (obj *_Gv$latchMgr) GetFromGets(gets int64) (results []*Gv$latch, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$latch{}).Where("`GETS` = ?", gets).Find(&results).Error
	
	return
}

// GetBatchFromGets 批量查找 
func (obj *_Gv$latchMgr) GetBatchFromGets(getss []int64) (results []*Gv$latch, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$latch{}).Where("`GETS` IN (?)", getss).Find(&results).Error
	
	return
}
 
// GetFromMisses 通过MISSES获取内容  
func (obj *_Gv$latchMgr) GetFromMisses(misses int64) (results []*Gv$latch, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$latch{}).Where("`MISSES` = ?", misses).Find(&results).Error
	
	return
}

// GetBatchFromMisses 批量查找 
func (obj *_Gv$latchMgr) GetBatchFromMisses(missess []int64) (results []*Gv$latch, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$latch{}).Where("`MISSES` IN (?)", missess).Find(&results).Error
	
	return
}
 
// GetFromSleeps 通过SLEEPS获取内容  
func (obj *_Gv$latchMgr) GetFromSleeps(sleeps int64) (results []*Gv$latch, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$latch{}).Where("`SLEEPS` = ?", sleeps).Find(&results).Error
	
	return
}

// GetBatchFromSleeps 批量查找 
func (obj *_Gv$latchMgr) GetBatchFromSleeps(sleepss []int64) (results []*Gv$latch, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$latch{}).Where("`SLEEPS` IN (?)", sleepss).Find(&results).Error
	
	return
}
 
// GetFromImmediateGets 通过IMMEDIATE_GETS获取内容  
func (obj *_Gv$latchMgr) GetFromImmediateGets(immediateGets int64) (results []*Gv$latch, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$latch{}).Where("`IMMEDIATE_GETS` = ?", immediateGets).Find(&results).Error
	
	return
}

// GetBatchFromImmediateGets 批量查找 
func (obj *_Gv$latchMgr) GetBatchFromImmediateGets(immediateGetss []int64) (results []*Gv$latch, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$latch{}).Where("`IMMEDIATE_GETS` IN (?)", immediateGetss).Find(&results).Error
	
	return
}
 
// GetFromImmediateMisses 通过IMMEDIATE_MISSES获取内容  
func (obj *_Gv$latchMgr) GetFromImmediateMisses(immediateMisses int64) (results []*Gv$latch, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$latch{}).Where("`IMMEDIATE_MISSES` = ?", immediateMisses).Find(&results).Error
	
	return
}

// GetBatchFromImmediateMisses 批量查找 
func (obj *_Gv$latchMgr) GetBatchFromImmediateMisses(immediateMissess []int64) (results []*Gv$latch, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$latch{}).Where("`IMMEDIATE_MISSES` IN (?)", immediateMissess).Find(&results).Error
	
	return
}
 
// GetFromSpinGets 通过SPIN_GETS获取内容  
func (obj *_Gv$latchMgr) GetFromSpinGets(spinGets int64) (results []*Gv$latch, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$latch{}).Where("`SPIN_GETS` = ?", spinGets).Find(&results).Error
	
	return
}

// GetBatchFromSpinGets 批量查找 
func (obj *_Gv$latchMgr) GetBatchFromSpinGets(spinGetss []int64) (results []*Gv$latch, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$latch{}).Where("`SPIN_GETS` IN (?)", spinGetss).Find(&results).Error
	
	return
}
 
// GetFromWaitTime 通过WAIT_TIME获取内容  
func (obj *_Gv$latchMgr) GetFromWaitTime(waitTime int64) (results []*Gv$latch, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$latch{}).Where("`WAIT_TIME` = ?", waitTime).Find(&results).Error
	
	return
}

// GetBatchFromWaitTime 批量查找 
func (obj *_Gv$latchMgr) GetBatchFromWaitTime(waitTimes []int64) (results []*Gv$latch, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$latch{}).Where("`WAIT_TIME` IN (?)", waitTimes).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 

 

	

