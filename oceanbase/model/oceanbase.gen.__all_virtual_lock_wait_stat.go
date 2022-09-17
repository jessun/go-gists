package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _AllVirtualLockWaitStatMgr struct {
	*_BaseMgr
}

// AllVirtualLockWaitStatMgr open func
func AllVirtualLockWaitStatMgr(db *gorm.DB) *_AllVirtualLockWaitStatMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualLockWaitStatMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualLockWaitStatMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_lock_wait_stat"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualLockWaitStatMgr) GetTableName() string {
	return "__all_virtual_lock_wait_stat"
}

// Reset 重置gorm会话
func (obj *_AllVirtualLockWaitStatMgr) Reset() *_AllVirtualLockWaitStatMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualLockWaitStatMgr) Get() (result AllVirtualLockWaitStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualLockWaitStat{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualLockWaitStatMgr) Gets() (results []*AllVirtualLockWaitStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualLockWaitStat{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualLockWaitStatMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualLockWaitStat{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithSvrIP svr_ip获取
func (obj *_AllVirtualLockWaitStatMgr) WithSvrIP(svrIP string) Option {
	return optionFunc(func(o *options) { o.query["svr_ip"] = svrIP })
}

// WithSvrPort svr_port获取
func (obj *_AllVirtualLockWaitStatMgr) WithSvrPort(svrPort int64) Option {
	return optionFunc(func(o *options) { o.query["svr_port"] = svrPort })
}

// WithTableID table_id获取
func (obj *_AllVirtualLockWaitStatMgr) WithTableID(tableID int64) Option {
	return optionFunc(func(o *options) { o.query["table_id"] = tableID })
}

// WithRowkey rowkey获取
func (obj *_AllVirtualLockWaitStatMgr) WithRowkey(rowkey string) Option {
	return optionFunc(func(o *options) { o.query["rowkey"] = rowkey })
}

// WithAddr addr获取
func (obj *_AllVirtualLockWaitStatMgr) WithAddr(addr uint64) Option {
	return optionFunc(func(o *options) { o.query["addr"] = addr })
}

// WithNeedWait need_wait获取
func (obj *_AllVirtualLockWaitStatMgr) WithNeedWait(needWait int8) Option {
	return optionFunc(func(o *options) { o.query["need_wait"] = needWait })
}

// WithRecvTs recv_ts获取
func (obj *_AllVirtualLockWaitStatMgr) WithRecvTs(recvTs int64) Option {
	return optionFunc(func(o *options) { o.query["recv_ts"] = recvTs })
}

// WithLockTs lock_ts获取
func (obj *_AllVirtualLockWaitStatMgr) WithLockTs(lockTs int64) Option {
	return optionFunc(func(o *options) { o.query["lock_ts"] = lockTs })
}

// WithAbsTimeout abs_timeout获取
func (obj *_AllVirtualLockWaitStatMgr) WithAbsTimeout(absTimeout int64) Option {
	return optionFunc(func(o *options) { o.query["abs_timeout"] = absTimeout })
}

// WithTryLockTimes try_lock_times获取
func (obj *_AllVirtualLockWaitStatMgr) WithTryLockTimes(tryLockTimes int64) Option {
	return optionFunc(func(o *options) { o.query["try_lock_times"] = tryLockTimes })
}

// WithTimeAfterRecv time_after_recv获取
func (obj *_AllVirtualLockWaitStatMgr) WithTimeAfterRecv(timeAfterRecv int64) Option {
	return optionFunc(func(o *options) { o.query["time_after_recv"] = timeAfterRecv })
}

// WithSessionID session_id获取
func (obj *_AllVirtualLockWaitStatMgr) WithSessionID(sessionID int64) Option {
	return optionFunc(func(o *options) { o.query["session_id"] = sessionID })
}

// WithBlockSessionID block_session_id获取
func (obj *_AllVirtualLockWaitStatMgr) WithBlockSessionID(blockSessionID int64) Option {
	return optionFunc(func(o *options) { o.query["block_session_id"] = blockSessionID })
}

// WithType type获取
func (obj *_AllVirtualLockWaitStatMgr) WithType(_type int64) Option {
	return optionFunc(func(o *options) { o.query["type"] = _type })
}

// WithLockMode lock_mode获取
func (obj *_AllVirtualLockWaitStatMgr) WithLockMode(lockMode int64) Option {
	return optionFunc(func(o *options) { o.query["lock_mode"] = lockMode })
}

// WithTotalUpdateCnt total_update_cnt获取
func (obj *_AllVirtualLockWaitStatMgr) WithTotalUpdateCnt(totalUpdateCnt int64) Option {
	return optionFunc(func(o *options) { o.query["total_update_cnt"] = totalUpdateCnt })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualLockWaitStatMgr) GetByOption(opts ...Option) (result AllVirtualLockWaitStat, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualLockWaitStat{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualLockWaitStatMgr) GetByOptions(opts ...Option) (results []*AllVirtualLockWaitStat, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualLockWaitStat{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualLockWaitStatMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualLockWaitStat, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualLockWaitStat{}).Where(options.query)
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
func (obj *_AllVirtualLockWaitStatMgr) GetFromSvrIP(svrIP string) (results []*AllVirtualLockWaitStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualLockWaitStat{}).Where("`svr_ip` = ?", svrIP).Find(&results).Error

	return
}

// GetBatchFromSvrIP 批量查找
func (obj *_AllVirtualLockWaitStatMgr) GetBatchFromSvrIP(svrIPs []string) (results []*AllVirtualLockWaitStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualLockWaitStat{}).Where("`svr_ip` IN (?)", svrIPs).Find(&results).Error

	return
}

// GetFromSvrPort 通过svr_port获取内容
func (obj *_AllVirtualLockWaitStatMgr) GetFromSvrPort(svrPort int64) (results []*AllVirtualLockWaitStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualLockWaitStat{}).Where("`svr_port` = ?", svrPort).Find(&results).Error

	return
}

// GetBatchFromSvrPort 批量查找
func (obj *_AllVirtualLockWaitStatMgr) GetBatchFromSvrPort(svrPorts []int64) (results []*AllVirtualLockWaitStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualLockWaitStat{}).Where("`svr_port` IN (?)", svrPorts).Find(&results).Error

	return
}

// GetFromTableID 通过table_id获取内容
func (obj *_AllVirtualLockWaitStatMgr) GetFromTableID(tableID int64) (results []*AllVirtualLockWaitStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualLockWaitStat{}).Where("`table_id` = ?", tableID).Find(&results).Error

	return
}

// GetBatchFromTableID 批量查找
func (obj *_AllVirtualLockWaitStatMgr) GetBatchFromTableID(tableIDs []int64) (results []*AllVirtualLockWaitStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualLockWaitStat{}).Where("`table_id` IN (?)", tableIDs).Find(&results).Error

	return
}

// GetFromRowkey 通过rowkey获取内容
func (obj *_AllVirtualLockWaitStatMgr) GetFromRowkey(rowkey string) (results []*AllVirtualLockWaitStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualLockWaitStat{}).Where("`rowkey` = ?", rowkey).Find(&results).Error

	return
}

// GetBatchFromRowkey 批量查找
func (obj *_AllVirtualLockWaitStatMgr) GetBatchFromRowkey(rowkeys []string) (results []*AllVirtualLockWaitStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualLockWaitStat{}).Where("`rowkey` IN (?)", rowkeys).Find(&results).Error

	return
}

// GetFromAddr 通过addr获取内容
func (obj *_AllVirtualLockWaitStatMgr) GetFromAddr(addr uint64) (results []*AllVirtualLockWaitStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualLockWaitStat{}).Where("`addr` = ?", addr).Find(&results).Error

	return
}

// GetBatchFromAddr 批量查找
func (obj *_AllVirtualLockWaitStatMgr) GetBatchFromAddr(addrs []uint64) (results []*AllVirtualLockWaitStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualLockWaitStat{}).Where("`addr` IN (?)", addrs).Find(&results).Error

	return
}

// GetFromNeedWait 通过need_wait获取内容
func (obj *_AllVirtualLockWaitStatMgr) GetFromNeedWait(needWait int8) (results []*AllVirtualLockWaitStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualLockWaitStat{}).Where("`need_wait` = ?", needWait).Find(&results).Error

	return
}

// GetBatchFromNeedWait 批量查找
func (obj *_AllVirtualLockWaitStatMgr) GetBatchFromNeedWait(needWaits []int8) (results []*AllVirtualLockWaitStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualLockWaitStat{}).Where("`need_wait` IN (?)", needWaits).Find(&results).Error

	return
}

// GetFromRecvTs 通过recv_ts获取内容
func (obj *_AllVirtualLockWaitStatMgr) GetFromRecvTs(recvTs int64) (results []*AllVirtualLockWaitStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualLockWaitStat{}).Where("`recv_ts` = ?", recvTs).Find(&results).Error

	return
}

// GetBatchFromRecvTs 批量查找
func (obj *_AllVirtualLockWaitStatMgr) GetBatchFromRecvTs(recvTss []int64) (results []*AllVirtualLockWaitStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualLockWaitStat{}).Where("`recv_ts` IN (?)", recvTss).Find(&results).Error

	return
}

// GetFromLockTs 通过lock_ts获取内容
func (obj *_AllVirtualLockWaitStatMgr) GetFromLockTs(lockTs int64) (results []*AllVirtualLockWaitStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualLockWaitStat{}).Where("`lock_ts` = ?", lockTs).Find(&results).Error

	return
}

// GetBatchFromLockTs 批量查找
func (obj *_AllVirtualLockWaitStatMgr) GetBatchFromLockTs(lockTss []int64) (results []*AllVirtualLockWaitStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualLockWaitStat{}).Where("`lock_ts` IN (?)", lockTss).Find(&results).Error

	return
}

// GetFromAbsTimeout 通过abs_timeout获取内容
func (obj *_AllVirtualLockWaitStatMgr) GetFromAbsTimeout(absTimeout int64) (results []*AllVirtualLockWaitStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualLockWaitStat{}).Where("`abs_timeout` = ?", absTimeout).Find(&results).Error

	return
}

// GetBatchFromAbsTimeout 批量查找
func (obj *_AllVirtualLockWaitStatMgr) GetBatchFromAbsTimeout(absTimeouts []int64) (results []*AllVirtualLockWaitStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualLockWaitStat{}).Where("`abs_timeout` IN (?)", absTimeouts).Find(&results).Error

	return
}

// GetFromTryLockTimes 通过try_lock_times获取内容
func (obj *_AllVirtualLockWaitStatMgr) GetFromTryLockTimes(tryLockTimes int64) (results []*AllVirtualLockWaitStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualLockWaitStat{}).Where("`try_lock_times` = ?", tryLockTimes).Find(&results).Error

	return
}

// GetBatchFromTryLockTimes 批量查找
func (obj *_AllVirtualLockWaitStatMgr) GetBatchFromTryLockTimes(tryLockTimess []int64) (results []*AllVirtualLockWaitStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualLockWaitStat{}).Where("`try_lock_times` IN (?)", tryLockTimess).Find(&results).Error

	return
}

// GetFromTimeAfterRecv 通过time_after_recv获取内容
func (obj *_AllVirtualLockWaitStatMgr) GetFromTimeAfterRecv(timeAfterRecv int64) (results []*AllVirtualLockWaitStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualLockWaitStat{}).Where("`time_after_recv` = ?", timeAfterRecv).Find(&results).Error

	return
}

// GetBatchFromTimeAfterRecv 批量查找
func (obj *_AllVirtualLockWaitStatMgr) GetBatchFromTimeAfterRecv(timeAfterRecvs []int64) (results []*AllVirtualLockWaitStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualLockWaitStat{}).Where("`time_after_recv` IN (?)", timeAfterRecvs).Find(&results).Error

	return
}

// GetFromSessionID 通过session_id获取内容
func (obj *_AllVirtualLockWaitStatMgr) GetFromSessionID(sessionID int64) (results []*AllVirtualLockWaitStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualLockWaitStat{}).Where("`session_id` = ?", sessionID).Find(&results).Error

	return
}

// GetBatchFromSessionID 批量查找
func (obj *_AllVirtualLockWaitStatMgr) GetBatchFromSessionID(sessionIDs []int64) (results []*AllVirtualLockWaitStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualLockWaitStat{}).Where("`session_id` IN (?)", sessionIDs).Find(&results).Error

	return
}

// GetFromBlockSessionID 通过block_session_id获取内容
func (obj *_AllVirtualLockWaitStatMgr) GetFromBlockSessionID(blockSessionID int64) (results []*AllVirtualLockWaitStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualLockWaitStat{}).Where("`block_session_id` = ?", blockSessionID).Find(&results).Error

	return
}

// GetBatchFromBlockSessionID 批量查找
func (obj *_AllVirtualLockWaitStatMgr) GetBatchFromBlockSessionID(blockSessionIDs []int64) (results []*AllVirtualLockWaitStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualLockWaitStat{}).Where("`block_session_id` IN (?)", blockSessionIDs).Find(&results).Error

	return
}

// GetFromType 通过type获取内容
func (obj *_AllVirtualLockWaitStatMgr) GetFromType(_type int64) (results []*AllVirtualLockWaitStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualLockWaitStat{}).Where("`type` = ?", _type).Find(&results).Error

	return
}

// GetBatchFromType 批量查找
func (obj *_AllVirtualLockWaitStatMgr) GetBatchFromType(_types []int64) (results []*AllVirtualLockWaitStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualLockWaitStat{}).Where("`type` IN (?)", _types).Find(&results).Error

	return
}

// GetFromLockMode 通过lock_mode获取内容
func (obj *_AllVirtualLockWaitStatMgr) GetFromLockMode(lockMode int64) (results []*AllVirtualLockWaitStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualLockWaitStat{}).Where("`lock_mode` = ?", lockMode).Find(&results).Error

	return
}

// GetBatchFromLockMode 批量查找
func (obj *_AllVirtualLockWaitStatMgr) GetBatchFromLockMode(lockModes []int64) (results []*AllVirtualLockWaitStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualLockWaitStat{}).Where("`lock_mode` IN (?)", lockModes).Find(&results).Error

	return
}

// GetFromTotalUpdateCnt 通过total_update_cnt获取内容
func (obj *_AllVirtualLockWaitStatMgr) GetFromTotalUpdateCnt(totalUpdateCnt int64) (results []*AllVirtualLockWaitStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualLockWaitStat{}).Where("`total_update_cnt` = ?", totalUpdateCnt).Find(&results).Error

	return
}

// GetBatchFromTotalUpdateCnt 批量查找
func (obj *_AllVirtualLockWaitStatMgr) GetBatchFromTotalUpdateCnt(totalUpdateCnts []int64) (results []*AllVirtualLockWaitStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualLockWaitStat{}).Where("`total_update_cnt` IN (?)", totalUpdateCnts).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////
