package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _AllVirtualDeadlockStatMgr struct {
	*_BaseMgr
}

// AllVirtualDeadlockStatMgr open func
func AllVirtualDeadlockStatMgr(db *gorm.DB) *_AllVirtualDeadlockStatMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualDeadlockStatMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualDeadlockStatMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_deadlock_stat"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualDeadlockStatMgr) GetTableName() string {
	return "__all_virtual_deadlock_stat"
}

// Reset 重置gorm会话
func (obj *_AllVirtualDeadlockStatMgr) Reset() *_AllVirtualDeadlockStatMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualDeadlockStatMgr) Get() (result AllVirtualDeadlockStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDeadlockStat{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualDeadlockStatMgr) Gets() (results []*AllVirtualDeadlockStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDeadlockStat{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualDeadlockStatMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualDeadlockStat{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithSvrIP svr_ip获取
func (obj *_AllVirtualDeadlockStatMgr) WithSvrIP(svrIP string) Option {
	return optionFunc(func(o *options) { o.query["svr_ip"] = svrIP })
}

// WithSvrPort svr_port获取
func (obj *_AllVirtualDeadlockStatMgr) WithSvrPort(svrPort int64) Option {
	return optionFunc(func(o *options) { o.query["svr_port"] = svrPort })
}

// WithCycleID cycle_id获取
func (obj *_AllVirtualDeadlockStatMgr) WithCycleID(cycleID uint64) Option {
	return optionFunc(func(o *options) { o.query["cycle_id"] = cycleID })
}

// WithCycleSeq cycle_seq获取
func (obj *_AllVirtualDeadlockStatMgr) WithCycleSeq(cycleSeq int64) Option {
	return optionFunc(func(o *options) { o.query["cycle_seq"] = cycleSeq })
}

// WithSessionID session_id获取
func (obj *_AllVirtualDeadlockStatMgr) WithSessionID(sessionID int64) Option {
	return optionFunc(func(o *options) { o.query["session_id"] = sessionID })
}

// WithTableID table_id获取
func (obj *_AllVirtualDeadlockStatMgr) WithTableID(tableID int64) Option {
	return optionFunc(func(o *options) { o.query["table_id"] = tableID })
}

// WithRowKey row_key获取
func (obj *_AllVirtualDeadlockStatMgr) WithRowKey(rowKey string) Option {
	return optionFunc(func(o *options) { o.query["row_key"] = rowKey })
}

// WithWaiterTransID waiter_trans_id获取
func (obj *_AllVirtualDeadlockStatMgr) WithWaiterTransID(waiterTransID string) Option {
	return optionFunc(func(o *options) { o.query["waiter_trans_id"] = waiterTransID })
}

// WithHolderTransID holder_trans_id获取
func (obj *_AllVirtualDeadlockStatMgr) WithHolderTransID(holderTransID string) Option {
	return optionFunc(func(o *options) { o.query["holder_trans_id"] = holderTransID })
}

// WithDeadlockRollbacked deadlock_rollbacked获取
func (obj *_AllVirtualDeadlockStatMgr) WithDeadlockRollbacked(deadlockRollbacked int8) Option {
	return optionFunc(func(o *options) { o.query["deadlock_rollbacked"] = deadlockRollbacked })
}

// WithCycleDetectTs cycle_detect_ts获取
func (obj *_AllVirtualDeadlockStatMgr) WithCycleDetectTs(cycleDetectTs int64) Option {
	return optionFunc(func(o *options) { o.query["cycle_detect_ts"] = cycleDetectTs })
}

// WithLockWaitTs lock_wait_ts获取
func (obj *_AllVirtualDeadlockStatMgr) WithLockWaitTs(lockWaitTs int64) Option {
	return optionFunc(func(o *options) { o.query["lock_wait_ts"] = lockWaitTs })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualDeadlockStatMgr) GetByOption(opts ...Option) (result AllVirtualDeadlockStat, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDeadlockStat{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualDeadlockStatMgr) GetByOptions(opts ...Option) (results []*AllVirtualDeadlockStat, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDeadlockStat{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualDeadlockStatMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualDeadlockStat, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualDeadlockStat{}).Where(options.query)
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

// GetFromSvrIP 通过svr_ip获取内容
func (obj *_AllVirtualDeadlockStatMgr) GetFromSvrIP(svrIP string) (results []*AllVirtualDeadlockStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDeadlockStat{}).Where("`svr_ip` = ?", svrIP).Find(&results).Error

	return
}

// GetBatchFromSvrIP 批量查找
func (obj *_AllVirtualDeadlockStatMgr) GetBatchFromSvrIP(svrIPs []string) (results []*AllVirtualDeadlockStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDeadlockStat{}).Where("`svr_ip` IN (?)", svrIPs).Find(&results).Error

	return
}

// GetFromSvrPort 通过svr_port获取内容
func (obj *_AllVirtualDeadlockStatMgr) GetFromSvrPort(svrPort int64) (results []*AllVirtualDeadlockStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDeadlockStat{}).Where("`svr_port` = ?", svrPort).Find(&results).Error

	return
}

// GetBatchFromSvrPort 批量查找
func (obj *_AllVirtualDeadlockStatMgr) GetBatchFromSvrPort(svrPorts []int64) (results []*AllVirtualDeadlockStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDeadlockStat{}).Where("`svr_port` IN (?)", svrPorts).Find(&results).Error

	return
}

// GetFromCycleID 通过cycle_id获取内容
func (obj *_AllVirtualDeadlockStatMgr) GetFromCycleID(cycleID uint64) (results []*AllVirtualDeadlockStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDeadlockStat{}).Where("`cycle_id` = ?", cycleID).Find(&results).Error

	return
}

// GetBatchFromCycleID 批量查找
func (obj *_AllVirtualDeadlockStatMgr) GetBatchFromCycleID(cycleIDs []uint64) (results []*AllVirtualDeadlockStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDeadlockStat{}).Where("`cycle_id` IN (?)", cycleIDs).Find(&results).Error

	return
}

// GetFromCycleSeq 通过cycle_seq获取内容
func (obj *_AllVirtualDeadlockStatMgr) GetFromCycleSeq(cycleSeq int64) (results []*AllVirtualDeadlockStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDeadlockStat{}).Where("`cycle_seq` = ?", cycleSeq).Find(&results).Error

	return
}

// GetBatchFromCycleSeq 批量查找
func (obj *_AllVirtualDeadlockStatMgr) GetBatchFromCycleSeq(cycleSeqs []int64) (results []*AllVirtualDeadlockStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDeadlockStat{}).Where("`cycle_seq` IN (?)", cycleSeqs).Find(&results).Error

	return
}

// GetFromSessionID 通过session_id获取内容
func (obj *_AllVirtualDeadlockStatMgr) GetFromSessionID(sessionID int64) (results []*AllVirtualDeadlockStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDeadlockStat{}).Where("`session_id` = ?", sessionID).Find(&results).Error

	return
}

// GetBatchFromSessionID 批量查找
func (obj *_AllVirtualDeadlockStatMgr) GetBatchFromSessionID(sessionIDs []int64) (results []*AllVirtualDeadlockStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDeadlockStat{}).Where("`session_id` IN (?)", sessionIDs).Find(&results).Error

	return
}

// GetFromTableID 通过table_id获取内容
func (obj *_AllVirtualDeadlockStatMgr) GetFromTableID(tableID int64) (results []*AllVirtualDeadlockStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDeadlockStat{}).Where("`table_id` = ?", tableID).Find(&results).Error

	return
}

// GetBatchFromTableID 批量查找
func (obj *_AllVirtualDeadlockStatMgr) GetBatchFromTableID(tableIDs []int64) (results []*AllVirtualDeadlockStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDeadlockStat{}).Where("`table_id` IN (?)", tableIDs).Find(&results).Error

	return
}

// GetFromRowKey 通过row_key获取内容
func (obj *_AllVirtualDeadlockStatMgr) GetFromRowKey(rowKey string) (results []*AllVirtualDeadlockStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDeadlockStat{}).Where("`row_key` = ?", rowKey).Find(&results).Error

	return
}

// GetBatchFromRowKey 批量查找
func (obj *_AllVirtualDeadlockStatMgr) GetBatchFromRowKey(rowKeys []string) (results []*AllVirtualDeadlockStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDeadlockStat{}).Where("`row_key` IN (?)", rowKeys).Find(&results).Error

	return
}

// GetFromWaiterTransID 通过waiter_trans_id获取内容
func (obj *_AllVirtualDeadlockStatMgr) GetFromWaiterTransID(waiterTransID string) (results []*AllVirtualDeadlockStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDeadlockStat{}).Where("`waiter_trans_id` = ?", waiterTransID).Find(&results).Error

	return
}

// GetBatchFromWaiterTransID 批量查找
func (obj *_AllVirtualDeadlockStatMgr) GetBatchFromWaiterTransID(waiterTransIDs []string) (results []*AllVirtualDeadlockStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDeadlockStat{}).Where("`waiter_trans_id` IN (?)", waiterTransIDs).Find(&results).Error

	return
}

// GetFromHolderTransID 通过holder_trans_id获取内容
func (obj *_AllVirtualDeadlockStatMgr) GetFromHolderTransID(holderTransID string) (results []*AllVirtualDeadlockStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDeadlockStat{}).Where("`holder_trans_id` = ?", holderTransID).Find(&results).Error

	return
}

// GetBatchFromHolderTransID 批量查找
func (obj *_AllVirtualDeadlockStatMgr) GetBatchFromHolderTransID(holderTransIDs []string) (results []*AllVirtualDeadlockStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDeadlockStat{}).Where("`holder_trans_id` IN (?)", holderTransIDs).Find(&results).Error

	return
}

// GetFromDeadlockRollbacked 通过deadlock_rollbacked获取内容
func (obj *_AllVirtualDeadlockStatMgr) GetFromDeadlockRollbacked(deadlockRollbacked int8) (results []*AllVirtualDeadlockStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDeadlockStat{}).Where("`deadlock_rollbacked` = ?", deadlockRollbacked).Find(&results).Error

	return
}

// GetBatchFromDeadlockRollbacked 批量查找
func (obj *_AllVirtualDeadlockStatMgr) GetBatchFromDeadlockRollbacked(deadlockRollbackeds []int8) (results []*AllVirtualDeadlockStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDeadlockStat{}).Where("`deadlock_rollbacked` IN (?)", deadlockRollbackeds).Find(&results).Error

	return
}

// GetFromCycleDetectTs 通过cycle_detect_ts获取内容
func (obj *_AllVirtualDeadlockStatMgr) GetFromCycleDetectTs(cycleDetectTs int64) (results []*AllVirtualDeadlockStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDeadlockStat{}).Where("`cycle_detect_ts` = ?", cycleDetectTs).Find(&results).Error

	return
}

// GetBatchFromCycleDetectTs 批量查找
func (obj *_AllVirtualDeadlockStatMgr) GetBatchFromCycleDetectTs(cycleDetectTss []int64) (results []*AllVirtualDeadlockStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDeadlockStat{}).Where("`cycle_detect_ts` IN (?)", cycleDetectTss).Find(&results).Error

	return
}

// GetFromLockWaitTs 通过lock_wait_ts获取内容
func (obj *_AllVirtualDeadlockStatMgr) GetFromLockWaitTs(lockWaitTs int64) (results []*AllVirtualDeadlockStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDeadlockStat{}).Where("`lock_wait_ts` = ?", lockWaitTs).Find(&results).Error

	return
}

// GetBatchFromLockWaitTs 批量查找
func (obj *_AllVirtualDeadlockStatMgr) GetBatchFromLockWaitTs(lockWaitTss []int64) (results []*AllVirtualDeadlockStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDeadlockStat{}).Where("`lock_wait_ts` IN (?)", lockWaitTss).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////
