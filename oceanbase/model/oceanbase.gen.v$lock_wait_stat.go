package	model	
import (	
"fmt"	
"context"	
"gorm.io/gorm"	
)	

type _V$lockWaitStatMgr struct {
	*_BaseMgr
}

// V$lockWaitStatMgr open func
func V$lockWaitStatMgr(db *gorm.DB) *_V$lockWaitStatMgr {
	if db == nil {
		panic(fmt.Errorf("V$lockWaitStatMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_V$lockWaitStatMgr{_BaseMgr: &_BaseMgr{DB: db.Table("v$lock_wait_stat"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_V$lockWaitStatMgr) GetTableName() string {
	return "v$lock_wait_stat"
}

// Reset 重置gorm会话
func (obj *_V$lockWaitStatMgr) Reset() *_V$lockWaitStatMgr {
	obj.New()
	return obj
}

// Get 获取 
func (obj *_V$lockWaitStatMgr) Get() (result V$lockWaitStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$lockWaitStat{}).First(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_V$lockWaitStatMgr) Gets() (results []*V$lockWaitStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$lockWaitStat{}).Find(&results).Error
	
	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_V$lockWaitStatMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(V$lockWaitStat{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTenantID TENANT_ID获取 
func (obj *_V$lockWaitStatMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["TENANT_ID"] = tenantID })
}

// WithTableID TABLE_ID获取 
func (obj *_V$lockWaitStatMgr) WithTableID(tableID int64) Option {
	return optionFunc(func(o *options) { o.query["TABLE_ID"] = tableID })
}

// WithTableName TABLE_NAME获取 
func (obj *_V$lockWaitStatMgr) WithTableName(tableName string) Option {
	return optionFunc(func(o *options) { o.query["TABLE_NAME"] = tableName })
}

// WithRowkey ROWKEY获取 
func (obj *_V$lockWaitStatMgr) WithRowkey(rowkey string) Option {
	return optionFunc(func(o *options) { o.query["ROWKEY"] = rowkey })
}

// WithSvrIP SVR_IP获取 
func (obj *_V$lockWaitStatMgr) WithSvrIP(svrIP string) Option {
	return optionFunc(func(o *options) { o.query["SVR_IP"] = svrIP })
}

// WithSvrPort SVR_PORT获取 
func (obj *_V$lockWaitStatMgr) WithSvrPort(svrPort int64) Option {
	return optionFunc(func(o *options) { o.query["SVR_PORT"] = svrPort })
}

// WithSessionID SESSION_ID获取 
func (obj *_V$lockWaitStatMgr) WithSessionID(sessionID int64) Option {
	return optionFunc(func(o *options) { o.query["SESSION_ID"] = sessionID })
}

// WithNeedWait NEED_WAIT获取 
func (obj *_V$lockWaitStatMgr) WithNeedWait(needWait int8) Option {
	return optionFunc(func(o *options) { o.query["NEED_WAIT"] = needWait })
}

// WithRecvTs RECV_TS获取 
func (obj *_V$lockWaitStatMgr) WithRecvTs(recvTs int64) Option {
	return optionFunc(func(o *options) { o.query["RECV_TS"] = recvTs })
}

// WithLockTs LOCK_TS获取 
func (obj *_V$lockWaitStatMgr) WithLockTs(lockTs int64) Option {
	return optionFunc(func(o *options) { o.query["LOCK_TS"] = lockTs })
}

// WithAbsTimeout ABS_TIMEOUT获取 
func (obj *_V$lockWaitStatMgr) WithAbsTimeout(absTimeout int64) Option {
	return optionFunc(func(o *options) { o.query["ABS_TIMEOUT"] = absTimeout })
}

// WithTryLockTimes TRY_LOCK_TIMES获取 
func (obj *_V$lockWaitStatMgr) WithTryLockTimes(tryLockTimes int64) Option {
	return optionFunc(func(o *options) { o.query["TRY_LOCK_TIMES"] = tryLockTimes })
}

// WithTimeAfterRecv TIME_AFTER_RECV获取 
func (obj *_V$lockWaitStatMgr) WithTimeAfterRecv(timeAfterRecv int64) Option {
	return optionFunc(func(o *options) { o.query["TIME_AFTER_RECV"] = timeAfterRecv })
}


// GetByOption 功能选项模式获取
func (obj *_V$lockWaitStatMgr) GetByOption(opts ...Option) (result V$lockWaitStat, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(V$lockWaitStat{}).Where(options.query).First(&result).Error
	
	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_V$lockWaitStatMgr) GetByOptions(opts ...Option) (results []*V$lockWaitStat, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(V$lockWaitStat{}).Where(options.query).Find(&results).Error
	
	return
}


// SelectPage 分页查询
func (obj *_V$lockWaitStatMgr) SelectPage(page IPage,opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]V$lockWaitStat,0)
	var count int64 // 统计总的记录数
	query :=  obj.DB.WithContext(obj.ctx).Model(V$lockWaitStat{}).Where(options.query)
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
func (obj *_V$lockWaitStatMgr) GetFromTenantID(tenantID int64) (results []*V$lockWaitStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$lockWaitStat{}).Where("`TENANT_ID` = ?", tenantID).Find(&results).Error
	
	return
}

// GetBatchFromTenantID 批量查找 
func (obj *_V$lockWaitStatMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*V$lockWaitStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$lockWaitStat{}).Where("`TENANT_ID` IN (?)", tenantIDs).Find(&results).Error
	
	return
}
 
// GetFromTableID 通过TABLE_ID获取内容  
func (obj *_V$lockWaitStatMgr) GetFromTableID(tableID int64) (results []*V$lockWaitStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$lockWaitStat{}).Where("`TABLE_ID` = ?", tableID).Find(&results).Error
	
	return
}

// GetBatchFromTableID 批量查找 
func (obj *_V$lockWaitStatMgr) GetBatchFromTableID(tableIDs []int64) (results []*V$lockWaitStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$lockWaitStat{}).Where("`TABLE_ID` IN (?)", tableIDs).Find(&results).Error
	
	return
}
 
// GetFromTableName 通过TABLE_NAME获取内容  
func (obj *_V$lockWaitStatMgr) GetFromTableName(tableName string) (results []*V$lockWaitStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$lockWaitStat{}).Where("`TABLE_NAME` = ?", tableName).Find(&results).Error
	
	return
}

// GetBatchFromTableName 批量查找 
func (obj *_V$lockWaitStatMgr) GetBatchFromTableName(tableNames []string) (results []*V$lockWaitStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$lockWaitStat{}).Where("`TABLE_NAME` IN (?)", tableNames).Find(&results).Error
	
	return
}
 
// GetFromRowkey 通过ROWKEY获取内容  
func (obj *_V$lockWaitStatMgr) GetFromRowkey(rowkey string) (results []*V$lockWaitStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$lockWaitStat{}).Where("`ROWKEY` = ?", rowkey).Find(&results).Error
	
	return
}

// GetBatchFromRowkey 批量查找 
func (obj *_V$lockWaitStatMgr) GetBatchFromRowkey(rowkeys []string) (results []*V$lockWaitStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$lockWaitStat{}).Where("`ROWKEY` IN (?)", rowkeys).Find(&results).Error
	
	return
}
 
// GetFromSvrIP 通过SVR_IP获取内容  
func (obj *_V$lockWaitStatMgr) GetFromSvrIP(svrIP string) (results []*V$lockWaitStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$lockWaitStat{}).Where("`SVR_IP` = ?", svrIP).Find(&results).Error
	
	return
}

// GetBatchFromSvrIP 批量查找 
func (obj *_V$lockWaitStatMgr) GetBatchFromSvrIP(svrIPs []string) (results []*V$lockWaitStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$lockWaitStat{}).Where("`SVR_IP` IN (?)", svrIPs).Find(&results).Error
	
	return
}
 
// GetFromSvrPort 通过SVR_PORT获取内容  
func (obj *_V$lockWaitStatMgr) GetFromSvrPort(svrPort int64) (results []*V$lockWaitStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$lockWaitStat{}).Where("`SVR_PORT` = ?", svrPort).Find(&results).Error
	
	return
}

// GetBatchFromSvrPort 批量查找 
func (obj *_V$lockWaitStatMgr) GetBatchFromSvrPort(svrPorts []int64) (results []*V$lockWaitStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$lockWaitStat{}).Where("`SVR_PORT` IN (?)", svrPorts).Find(&results).Error
	
	return
}
 
// GetFromSessionID 通过SESSION_ID获取内容  
func (obj *_V$lockWaitStatMgr) GetFromSessionID(sessionID int64) (results []*V$lockWaitStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$lockWaitStat{}).Where("`SESSION_ID` = ?", sessionID).Find(&results).Error
	
	return
}

// GetBatchFromSessionID 批量查找 
func (obj *_V$lockWaitStatMgr) GetBatchFromSessionID(sessionIDs []int64) (results []*V$lockWaitStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$lockWaitStat{}).Where("`SESSION_ID` IN (?)", sessionIDs).Find(&results).Error
	
	return
}
 
// GetFromNeedWait 通过NEED_WAIT获取内容  
func (obj *_V$lockWaitStatMgr) GetFromNeedWait(needWait int8) (results []*V$lockWaitStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$lockWaitStat{}).Where("`NEED_WAIT` = ?", needWait).Find(&results).Error
	
	return
}

// GetBatchFromNeedWait 批量查找 
func (obj *_V$lockWaitStatMgr) GetBatchFromNeedWait(needWaits []int8) (results []*V$lockWaitStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$lockWaitStat{}).Where("`NEED_WAIT` IN (?)", needWaits).Find(&results).Error
	
	return
}
 
// GetFromRecvTs 通过RECV_TS获取内容  
func (obj *_V$lockWaitStatMgr) GetFromRecvTs(recvTs int64) (results []*V$lockWaitStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$lockWaitStat{}).Where("`RECV_TS` = ?", recvTs).Find(&results).Error
	
	return
}

// GetBatchFromRecvTs 批量查找 
func (obj *_V$lockWaitStatMgr) GetBatchFromRecvTs(recvTss []int64) (results []*V$lockWaitStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$lockWaitStat{}).Where("`RECV_TS` IN (?)", recvTss).Find(&results).Error
	
	return
}
 
// GetFromLockTs 通过LOCK_TS获取内容  
func (obj *_V$lockWaitStatMgr) GetFromLockTs(lockTs int64) (results []*V$lockWaitStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$lockWaitStat{}).Where("`LOCK_TS` = ?", lockTs).Find(&results).Error
	
	return
}

// GetBatchFromLockTs 批量查找 
func (obj *_V$lockWaitStatMgr) GetBatchFromLockTs(lockTss []int64) (results []*V$lockWaitStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$lockWaitStat{}).Where("`LOCK_TS` IN (?)", lockTss).Find(&results).Error
	
	return
}
 
// GetFromAbsTimeout 通过ABS_TIMEOUT获取内容  
func (obj *_V$lockWaitStatMgr) GetFromAbsTimeout(absTimeout int64) (results []*V$lockWaitStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$lockWaitStat{}).Where("`ABS_TIMEOUT` = ?", absTimeout).Find(&results).Error
	
	return
}

// GetBatchFromAbsTimeout 批量查找 
func (obj *_V$lockWaitStatMgr) GetBatchFromAbsTimeout(absTimeouts []int64) (results []*V$lockWaitStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$lockWaitStat{}).Where("`ABS_TIMEOUT` IN (?)", absTimeouts).Find(&results).Error
	
	return
}
 
// GetFromTryLockTimes 通过TRY_LOCK_TIMES获取内容  
func (obj *_V$lockWaitStatMgr) GetFromTryLockTimes(tryLockTimes int64) (results []*V$lockWaitStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$lockWaitStat{}).Where("`TRY_LOCK_TIMES` = ?", tryLockTimes).Find(&results).Error
	
	return
}

// GetBatchFromTryLockTimes 批量查找 
func (obj *_V$lockWaitStatMgr) GetBatchFromTryLockTimes(tryLockTimess []int64) (results []*V$lockWaitStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$lockWaitStat{}).Where("`TRY_LOCK_TIMES` IN (?)", tryLockTimess).Find(&results).Error
	
	return
}
 
// GetFromTimeAfterRecv 通过TIME_AFTER_RECV获取内容  
func (obj *_V$lockWaitStatMgr) GetFromTimeAfterRecv(timeAfterRecv int64) (results []*V$lockWaitStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$lockWaitStat{}).Where("`TIME_AFTER_RECV` = ?", timeAfterRecv).Find(&results).Error
	
	return
}

// GetBatchFromTimeAfterRecv 批量查找 
func (obj *_V$lockWaitStatMgr) GetBatchFromTimeAfterRecv(timeAfterRecvs []int64) (results []*V$lockWaitStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$lockWaitStat{}).Where("`TIME_AFTER_RECV` IN (?)", timeAfterRecvs).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 

 

	

