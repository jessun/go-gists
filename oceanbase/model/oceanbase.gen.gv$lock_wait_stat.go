package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
)	

type _Gv$lockWaitStatMgr struct {
	*_BaseMgr
}

// Gv$lockWaitStatMgr open func
func Gv$lockWaitStatMgr(db *gorm.DB) *_Gv$lockWaitStatMgr {
	if db == nil {
		panic(fmt.Errorf("Gv$lockWaitStatMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_Gv$lockWaitStatMgr{_BaseMgr: &_BaseMgr{DB: db.Table("gv$lock_wait_stat"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_Gv$lockWaitStatMgr) GetTableName() string {
	return "gv$lock_wait_stat"
}

// Reset 重置gorm会话
func (obj *_Gv$lockWaitStatMgr) Reset() *_Gv$lockWaitStatMgr {
	obj.New()
	return obj
}

// Get 获取 
func (obj *_Gv$lockWaitStatMgr) Get() (result Gv$lockWaitStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$lockWaitStat{}).First(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_Gv$lockWaitStatMgr) Gets() (results []*Gv$lockWaitStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$lockWaitStat{}).Find(&results).Error
	
	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_Gv$lockWaitStatMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(Gv$lockWaitStat{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTenantID TENANT_ID获取 
func (obj *_Gv$lockWaitStatMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["TENANT_ID"] = tenantID })
}

// WithTableID TABLE_ID获取 
func (obj *_Gv$lockWaitStatMgr) WithTableID(tableID int64) Option {
	return optionFunc(func(o *options) { o.query["TABLE_ID"] = tableID })
}

// WithTableName TABLE_NAME获取 
func (obj *_Gv$lockWaitStatMgr) WithTableName(tableName string) Option {
	return optionFunc(func(o *options) { o.query["TABLE_NAME"] = tableName })
}

// WithRowkey ROWKEY获取 
func (obj *_Gv$lockWaitStatMgr) WithRowkey(rowkey string) Option {
	return optionFunc(func(o *options) { o.query["ROWKEY"] = rowkey })
}

// WithSvrIP SVR_IP获取 
func (obj *_Gv$lockWaitStatMgr) WithSvrIP(svrIP string) Option {
	return optionFunc(func(o *options) { o.query["SVR_IP"] = svrIP })
}

// WithSvrPort SVR_PORT获取 
func (obj *_Gv$lockWaitStatMgr) WithSvrPort(svrPort int64) Option {
	return optionFunc(func(o *options) { o.query["SVR_PORT"] = svrPort })
}

// WithSessionID SESSION_ID获取 
func (obj *_Gv$lockWaitStatMgr) WithSessionID(sessionID int64) Option {
	return optionFunc(func(o *options) { o.query["SESSION_ID"] = sessionID })
}

// WithNeedWait NEED_WAIT获取 
func (obj *_Gv$lockWaitStatMgr) WithNeedWait(needWait int8) Option {
	return optionFunc(func(o *options) { o.query["NEED_WAIT"] = needWait })
}

// WithRecvTs RECV_TS获取 
func (obj *_Gv$lockWaitStatMgr) WithRecvTs(recvTs int64) Option {
	return optionFunc(func(o *options) { o.query["RECV_TS"] = recvTs })
}

// WithLockTs LOCK_TS获取 
func (obj *_Gv$lockWaitStatMgr) WithLockTs(lockTs int64) Option {
	return optionFunc(func(o *options) { o.query["LOCK_TS"] = lockTs })
}

// WithAbsTimeout ABS_TIMEOUT获取 
func (obj *_Gv$lockWaitStatMgr) WithAbsTimeout(absTimeout int64) Option {
	return optionFunc(func(o *options) { o.query["ABS_TIMEOUT"] = absTimeout })
}

// WithTryLockTimes TRY_LOCK_TIMES获取 
func (obj *_Gv$lockWaitStatMgr) WithTryLockTimes(tryLockTimes int64) Option {
	return optionFunc(func(o *options) { o.query["TRY_LOCK_TIMES"] = tryLockTimes })
}

// WithTimeAfterRecv TIME_AFTER_RECV获取 
func (obj *_Gv$lockWaitStatMgr) WithTimeAfterRecv(timeAfterRecv int64) Option {
	return optionFunc(func(o *options) { o.query["TIME_AFTER_RECV"] = timeAfterRecv })
}


// GetByOption 功能选项模式获取
func (obj *_Gv$lockWaitStatMgr) GetByOption(opts ...Option) (result Gv$lockWaitStat, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Gv$lockWaitStat{}).Where(options.query).First(&result).Error
	
	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_Gv$lockWaitStatMgr) GetByOptions(opts ...Option) (results []*Gv$lockWaitStat, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Gv$lockWaitStat{}).Where(options.query).Find(&results).Error
	
	return
}


// SelectPage 分页查询
func (obj *_Gv$lockWaitStatMgr) SelectPage(page IPage,opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]Gv$lockWaitStat,0)
	var count int64 // 统计总的记录数
	query :=  obj.DB.WithContext(obj.ctx).Model(Gv$lockWaitStat{}).Where(options.query)
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
func (obj *_Gv$lockWaitStatMgr) GetFromTenantID(tenantID int64) (results []*Gv$lockWaitStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$lockWaitStat{}).Where("`TENANT_ID` = ?", tenantID).Find(&results).Error
	
	return
}

// GetBatchFromTenantID 批量查找 
func (obj *_Gv$lockWaitStatMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*Gv$lockWaitStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$lockWaitStat{}).Where("`TENANT_ID` IN (?)", tenantIDs).Find(&results).Error
	
	return
}
 
// GetFromTableID 通过TABLE_ID获取内容  
func (obj *_Gv$lockWaitStatMgr) GetFromTableID(tableID int64) (results []*Gv$lockWaitStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$lockWaitStat{}).Where("`TABLE_ID` = ?", tableID).Find(&results).Error
	
	return
}

// GetBatchFromTableID 批量查找 
func (obj *_Gv$lockWaitStatMgr) GetBatchFromTableID(tableIDs []int64) (results []*Gv$lockWaitStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$lockWaitStat{}).Where("`TABLE_ID` IN (?)", tableIDs).Find(&results).Error
	
	return
}
 
// GetFromTableName 通过TABLE_NAME获取内容  
func (obj *_Gv$lockWaitStatMgr) GetFromTableName(tableName string) (results []*Gv$lockWaitStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$lockWaitStat{}).Where("`TABLE_NAME` = ?", tableName).Find(&results).Error
	
	return
}

// GetBatchFromTableName 批量查找 
func (obj *_Gv$lockWaitStatMgr) GetBatchFromTableName(tableNames []string) (results []*Gv$lockWaitStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$lockWaitStat{}).Where("`TABLE_NAME` IN (?)", tableNames).Find(&results).Error
	
	return
}
 
// GetFromRowkey 通过ROWKEY获取内容  
func (obj *_Gv$lockWaitStatMgr) GetFromRowkey(rowkey string) (results []*Gv$lockWaitStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$lockWaitStat{}).Where("`ROWKEY` = ?", rowkey).Find(&results).Error
	
	return
}

// GetBatchFromRowkey 批量查找 
func (obj *_Gv$lockWaitStatMgr) GetBatchFromRowkey(rowkeys []string) (results []*Gv$lockWaitStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$lockWaitStat{}).Where("`ROWKEY` IN (?)", rowkeys).Find(&results).Error
	
	return
}
 
// GetFromSvrIP 通过SVR_IP获取内容  
func (obj *_Gv$lockWaitStatMgr) GetFromSvrIP(svrIP string) (results []*Gv$lockWaitStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$lockWaitStat{}).Where("`SVR_IP` = ?", svrIP).Find(&results).Error
	
	return
}

// GetBatchFromSvrIP 批量查找 
func (obj *_Gv$lockWaitStatMgr) GetBatchFromSvrIP(svrIPs []string) (results []*Gv$lockWaitStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$lockWaitStat{}).Where("`SVR_IP` IN (?)", svrIPs).Find(&results).Error
	
	return
}
 
// GetFromSvrPort 通过SVR_PORT获取内容  
func (obj *_Gv$lockWaitStatMgr) GetFromSvrPort(svrPort int64) (results []*Gv$lockWaitStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$lockWaitStat{}).Where("`SVR_PORT` = ?", svrPort).Find(&results).Error
	
	return
}

// GetBatchFromSvrPort 批量查找 
func (obj *_Gv$lockWaitStatMgr) GetBatchFromSvrPort(svrPorts []int64) (results []*Gv$lockWaitStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$lockWaitStat{}).Where("`SVR_PORT` IN (?)", svrPorts).Find(&results).Error
	
	return
}
 
// GetFromSessionID 通过SESSION_ID获取内容  
func (obj *_Gv$lockWaitStatMgr) GetFromSessionID(sessionID int64) (results []*Gv$lockWaitStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$lockWaitStat{}).Where("`SESSION_ID` = ?", sessionID).Find(&results).Error
	
	return
}

// GetBatchFromSessionID 批量查找 
func (obj *_Gv$lockWaitStatMgr) GetBatchFromSessionID(sessionIDs []int64) (results []*Gv$lockWaitStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$lockWaitStat{}).Where("`SESSION_ID` IN (?)", sessionIDs).Find(&results).Error
	
	return
}
 
// GetFromNeedWait 通过NEED_WAIT获取内容  
func (obj *_Gv$lockWaitStatMgr) GetFromNeedWait(needWait int8) (results []*Gv$lockWaitStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$lockWaitStat{}).Where("`NEED_WAIT` = ?", needWait).Find(&results).Error
	
	return
}

// GetBatchFromNeedWait 批量查找 
func (obj *_Gv$lockWaitStatMgr) GetBatchFromNeedWait(needWaits []int8) (results []*Gv$lockWaitStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$lockWaitStat{}).Where("`NEED_WAIT` IN (?)", needWaits).Find(&results).Error
	
	return
}
 
// GetFromRecvTs 通过RECV_TS获取内容  
func (obj *_Gv$lockWaitStatMgr) GetFromRecvTs(recvTs int64) (results []*Gv$lockWaitStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$lockWaitStat{}).Where("`RECV_TS` = ?", recvTs).Find(&results).Error
	
	return
}

// GetBatchFromRecvTs 批量查找 
func (obj *_Gv$lockWaitStatMgr) GetBatchFromRecvTs(recvTss []int64) (results []*Gv$lockWaitStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$lockWaitStat{}).Where("`RECV_TS` IN (?)", recvTss).Find(&results).Error
	
	return
}
 
// GetFromLockTs 通过LOCK_TS获取内容  
func (obj *_Gv$lockWaitStatMgr) GetFromLockTs(lockTs int64) (results []*Gv$lockWaitStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$lockWaitStat{}).Where("`LOCK_TS` = ?", lockTs).Find(&results).Error
	
	return
}

// GetBatchFromLockTs 批量查找 
func (obj *_Gv$lockWaitStatMgr) GetBatchFromLockTs(lockTss []int64) (results []*Gv$lockWaitStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$lockWaitStat{}).Where("`LOCK_TS` IN (?)", lockTss).Find(&results).Error
	
	return
}
 
// GetFromAbsTimeout 通过ABS_TIMEOUT获取内容  
func (obj *_Gv$lockWaitStatMgr) GetFromAbsTimeout(absTimeout int64) (results []*Gv$lockWaitStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$lockWaitStat{}).Where("`ABS_TIMEOUT` = ?", absTimeout).Find(&results).Error
	
	return
}

// GetBatchFromAbsTimeout 批量查找 
func (obj *_Gv$lockWaitStatMgr) GetBatchFromAbsTimeout(absTimeouts []int64) (results []*Gv$lockWaitStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$lockWaitStat{}).Where("`ABS_TIMEOUT` IN (?)", absTimeouts).Find(&results).Error
	
	return
}
 
// GetFromTryLockTimes 通过TRY_LOCK_TIMES获取内容  
func (obj *_Gv$lockWaitStatMgr) GetFromTryLockTimes(tryLockTimes int64) (results []*Gv$lockWaitStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$lockWaitStat{}).Where("`TRY_LOCK_TIMES` = ?", tryLockTimes).Find(&results).Error
	
	return
}

// GetBatchFromTryLockTimes 批量查找 
func (obj *_Gv$lockWaitStatMgr) GetBatchFromTryLockTimes(tryLockTimess []int64) (results []*Gv$lockWaitStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$lockWaitStat{}).Where("`TRY_LOCK_TIMES` IN (?)", tryLockTimess).Find(&results).Error
	
	return
}
 
// GetFromTimeAfterRecv 通过TIME_AFTER_RECV获取内容  
func (obj *_Gv$lockWaitStatMgr) GetFromTimeAfterRecv(timeAfterRecv int64) (results []*Gv$lockWaitStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$lockWaitStat{}).Where("`TIME_AFTER_RECV` = ?", timeAfterRecv).Find(&results).Error
	
	return
}

// GetBatchFromTimeAfterRecv 批量查找 
func (obj *_Gv$lockWaitStatMgr) GetBatchFromTimeAfterRecv(timeAfterRecvs []int64) (results []*Gv$lockWaitStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$lockWaitStat{}).Where("`TIME_AFTER_RECV` IN (?)", timeAfterRecvs).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 

 

	

