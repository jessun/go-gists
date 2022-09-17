package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _AllVirtualDumpTenantInfoMgr struct {
	*_BaseMgr
}

// AllVirtualDumpTenantInfoMgr open func
func AllVirtualDumpTenantInfoMgr(db *gorm.DB) *_AllVirtualDumpTenantInfoMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualDumpTenantInfoMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualDumpTenantInfoMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_dump_tenant_info"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualDumpTenantInfoMgr) GetTableName() string {
	return "__all_virtual_dump_tenant_info"
}

// Reset 重置gorm会话
func (obj *_AllVirtualDumpTenantInfoMgr) Reset() *_AllVirtualDumpTenantInfoMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualDumpTenantInfoMgr) Get() (result AllVirtualDumpTenantInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDumpTenantInfo{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualDumpTenantInfoMgr) Gets() (results []*AllVirtualDumpTenantInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDumpTenantInfo{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualDumpTenantInfoMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualDumpTenantInfo{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithSvrIP svr_ip获取
func (obj *_AllVirtualDumpTenantInfoMgr) WithSvrIP(svrIP string) Option {
	return optionFunc(func(o *options) { o.query["svr_ip"] = svrIP })
}

// WithSvrPort svr_port获取
func (obj *_AllVirtualDumpTenantInfoMgr) WithSvrPort(svrPort int64) Option {
	return optionFunc(func(o *options) { o.query["svr_port"] = svrPort })
}

// WithTenantID tenant_id获取
func (obj *_AllVirtualDumpTenantInfoMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithCompatMode compat_mode获取
func (obj *_AllVirtualDumpTenantInfoMgr) WithCompatMode(compatMode int64) Option {
	return optionFunc(func(o *options) { o.query["compat_mode"] = compatMode })
}

// WithUnitMinCPU unit_min_cpu获取
func (obj *_AllVirtualDumpTenantInfoMgr) WithUnitMinCPU(unitMinCPU float64) Option {
	return optionFunc(func(o *options) { o.query["unit_min_cpu"] = unitMinCPU })
}

// WithUnitMaxCPU unit_max_cpu获取
func (obj *_AllVirtualDumpTenantInfoMgr) WithUnitMaxCPU(unitMaxCPU float64) Option {
	return optionFunc(func(o *options) { o.query["unit_max_cpu"] = unitMaxCPU })
}

// WithSlice slice获取
func (obj *_AllVirtualDumpTenantInfoMgr) WithSlice(slice float64) Option {
	return optionFunc(func(o *options) { o.query["slice"] = slice })
}

// WithRemainSlice remain_slice获取
func (obj *_AllVirtualDumpTenantInfoMgr) WithRemainSlice(remainSlice float64) Option {
	return optionFunc(func(o *options) { o.query["remain_slice"] = remainSlice })
}

// WithTokenCnt token_cnt获取
func (obj *_AllVirtualDumpTenantInfoMgr) WithTokenCnt(tokenCnt int64) Option {
	return optionFunc(func(o *options) { o.query["token_cnt"] = tokenCnt })
}

// WithAssTokenCnt ass_token_cnt获取
func (obj *_AllVirtualDumpTenantInfoMgr) WithAssTokenCnt(assTokenCnt int64) Option {
	return optionFunc(func(o *options) { o.query["ass_token_cnt"] = assTokenCnt })
}

// WithLqTokens lq_tokens获取
func (obj *_AllVirtualDumpTenantInfoMgr) WithLqTokens(lqTokens int64) Option {
	return optionFunc(func(o *options) { o.query["lq_tokens"] = lqTokens })
}

// WithUsedLqTokens used_lq_tokens获取
func (obj *_AllVirtualDumpTenantInfoMgr) WithUsedLqTokens(usedLqTokens int64) Option {
	return optionFunc(func(o *options) { o.query["used_lq_tokens"] = usedLqTokens })
}

// WithStopped stopped获取
func (obj *_AllVirtualDumpTenantInfoMgr) WithStopped(stopped int64) Option {
	return optionFunc(func(o *options) { o.query["stopped"] = stopped })
}

// WithIDleUs idle_us获取
func (obj *_AllVirtualDumpTenantInfoMgr) WithIDleUs(idleUs int64) Option {
	return optionFunc(func(o *options) { o.query["idle_us"] = idleUs })
}

// WithRecvHpRPCCnt recv_hp_rpc_cnt获取
func (obj *_AllVirtualDumpTenantInfoMgr) WithRecvHpRPCCnt(recvHpRPCCnt int64) Option {
	return optionFunc(func(o *options) { o.query["recv_hp_rpc_cnt"] = recvHpRPCCnt })
}

// WithRecvNpRPCCnt recv_np_rpc_cnt获取
func (obj *_AllVirtualDumpTenantInfoMgr) WithRecvNpRPCCnt(recvNpRPCCnt int64) Option {
	return optionFunc(func(o *options) { o.query["recv_np_rpc_cnt"] = recvNpRPCCnt })
}

// WithRecvLpRPCCnt recv_lp_rpc_cnt获取
func (obj *_AllVirtualDumpTenantInfoMgr) WithRecvLpRPCCnt(recvLpRPCCnt int64) Option {
	return optionFunc(func(o *options) { o.query["recv_lp_rpc_cnt"] = recvLpRPCCnt })
}

// WithRecvMysqlCnt recv_mysql_cnt获取
func (obj *_AllVirtualDumpTenantInfoMgr) WithRecvMysqlCnt(recvMysqlCnt int64) Option {
	return optionFunc(func(o *options) { o.query["recv_mysql_cnt"] = recvMysqlCnt })
}

// WithRecvTaskCnt recv_task_cnt获取
func (obj *_AllVirtualDumpTenantInfoMgr) WithRecvTaskCnt(recvTaskCnt int64) Option {
	return optionFunc(func(o *options) { o.query["recv_task_cnt"] = recvTaskCnt })
}

// WithRecvLargeReqCnt recv_large_req_cnt获取
func (obj *_AllVirtualDumpTenantInfoMgr) WithRecvLargeReqCnt(recvLargeReqCnt int64) Option {
	return optionFunc(func(o *options) { o.query["recv_large_req_cnt"] = recvLargeReqCnt })
}

// WithRecvLargeQueries recv_large_queries获取
func (obj *_AllVirtualDumpTenantInfoMgr) WithRecvLargeQueries(recvLargeQueries int64) Option {
	return optionFunc(func(o *options) { o.query["recv_large_queries"] = recvLargeQueries })
}

// WithActives actives获取
func (obj *_AllVirtualDumpTenantInfoMgr) WithActives(actives int64) Option {
	return optionFunc(func(o *options) { o.query["actives"] = actives })
}

// WithWorkers workers获取
func (obj *_AllVirtualDumpTenantInfoMgr) WithWorkers(workers int64) Option {
	return optionFunc(func(o *options) { o.query["workers"] = workers })
}

// WithLqWaitingWorkers lq_waiting_workers获取
func (obj *_AllVirtualDumpTenantInfoMgr) WithLqWaitingWorkers(lqWaitingWorkers int64) Option {
	return optionFunc(func(o *options) { o.query["lq_waiting_workers"] = lqWaitingWorkers })
}

// WithReqQueueTotalSize req_queue_total_size获取
func (obj *_AllVirtualDumpTenantInfoMgr) WithReqQueueTotalSize(reqQueueTotalSize int64) Option {
	return optionFunc(func(o *options) { o.query["req_queue_total_size"] = reqQueueTotalSize })
}

// WithQueue0 queue_0获取
func (obj *_AllVirtualDumpTenantInfoMgr) WithQueue0(queue0 int64) Option {
	return optionFunc(func(o *options) { o.query["queue_0"] = queue0 })
}

// WithQueue1 queue_1获取
func (obj *_AllVirtualDumpTenantInfoMgr) WithQueue1(queue1 int64) Option {
	return optionFunc(func(o *options) { o.query["queue_1"] = queue1 })
}

// WithQueue2 queue_2获取
func (obj *_AllVirtualDumpTenantInfoMgr) WithQueue2(queue2 int64) Option {
	return optionFunc(func(o *options) { o.query["queue_2"] = queue2 })
}

// WithQueue3 queue_3获取
func (obj *_AllVirtualDumpTenantInfoMgr) WithQueue3(queue3 int64) Option {
	return optionFunc(func(o *options) { o.query["queue_3"] = queue3 })
}

// WithQueue4 queue_4获取
func (obj *_AllVirtualDumpTenantInfoMgr) WithQueue4(queue4 int64) Option {
	return optionFunc(func(o *options) { o.query["queue_4"] = queue4 })
}

// WithQueue5 queue_5获取
func (obj *_AllVirtualDumpTenantInfoMgr) WithQueue5(queue5 int64) Option {
	return optionFunc(func(o *options) { o.query["queue_5"] = queue5 })
}

// WithLargeQueued large_queued获取
func (obj *_AllVirtualDumpTenantInfoMgr) WithLargeQueued(largeQueued int64) Option {
	return optionFunc(func(o *options) { o.query["large_queued"] = largeQueued })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualDumpTenantInfoMgr) GetByOption(opts ...Option) (result AllVirtualDumpTenantInfo, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDumpTenantInfo{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualDumpTenantInfoMgr) GetByOptions(opts ...Option) (results []*AllVirtualDumpTenantInfo, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDumpTenantInfo{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualDumpTenantInfoMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualDumpTenantInfo, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualDumpTenantInfo{}).Where(options.query)
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
func (obj *_AllVirtualDumpTenantInfoMgr) GetFromSvrIP(svrIP string) (results []*AllVirtualDumpTenantInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDumpTenantInfo{}).Where("`svr_ip` = ?", svrIP).Find(&results).Error

	return
}

// GetBatchFromSvrIP 批量查找
func (obj *_AllVirtualDumpTenantInfoMgr) GetBatchFromSvrIP(svrIPs []string) (results []*AllVirtualDumpTenantInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDumpTenantInfo{}).Where("`svr_ip` IN (?)", svrIPs).Find(&results).Error

	return
}

// GetFromSvrPort 通过svr_port获取内容
func (obj *_AllVirtualDumpTenantInfoMgr) GetFromSvrPort(svrPort int64) (results []*AllVirtualDumpTenantInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDumpTenantInfo{}).Where("`svr_port` = ?", svrPort).Find(&results).Error

	return
}

// GetBatchFromSvrPort 批量查找
func (obj *_AllVirtualDumpTenantInfoMgr) GetBatchFromSvrPort(svrPorts []int64) (results []*AllVirtualDumpTenantInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDumpTenantInfo{}).Where("`svr_port` IN (?)", svrPorts).Find(&results).Error

	return
}

// GetFromTenantID 通过tenant_id获取内容
func (obj *_AllVirtualDumpTenantInfoMgr) GetFromTenantID(tenantID int64) (results []*AllVirtualDumpTenantInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDumpTenantInfo{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllVirtualDumpTenantInfoMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllVirtualDumpTenantInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDumpTenantInfo{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromCompatMode 通过compat_mode获取内容
func (obj *_AllVirtualDumpTenantInfoMgr) GetFromCompatMode(compatMode int64) (results []*AllVirtualDumpTenantInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDumpTenantInfo{}).Where("`compat_mode` = ?", compatMode).Find(&results).Error

	return
}

// GetBatchFromCompatMode 批量查找
func (obj *_AllVirtualDumpTenantInfoMgr) GetBatchFromCompatMode(compatModes []int64) (results []*AllVirtualDumpTenantInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDumpTenantInfo{}).Where("`compat_mode` IN (?)", compatModes).Find(&results).Error

	return
}

// GetFromUnitMinCPU 通过unit_min_cpu获取内容
func (obj *_AllVirtualDumpTenantInfoMgr) GetFromUnitMinCPU(unitMinCPU float64) (results []*AllVirtualDumpTenantInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDumpTenantInfo{}).Where("`unit_min_cpu` = ?", unitMinCPU).Find(&results).Error

	return
}

// GetBatchFromUnitMinCPU 批量查找
func (obj *_AllVirtualDumpTenantInfoMgr) GetBatchFromUnitMinCPU(unitMinCPUs []float64) (results []*AllVirtualDumpTenantInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDumpTenantInfo{}).Where("`unit_min_cpu` IN (?)", unitMinCPUs).Find(&results).Error

	return
}

// GetFromUnitMaxCPU 通过unit_max_cpu获取内容
func (obj *_AllVirtualDumpTenantInfoMgr) GetFromUnitMaxCPU(unitMaxCPU float64) (results []*AllVirtualDumpTenantInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDumpTenantInfo{}).Where("`unit_max_cpu` = ?", unitMaxCPU).Find(&results).Error

	return
}

// GetBatchFromUnitMaxCPU 批量查找
func (obj *_AllVirtualDumpTenantInfoMgr) GetBatchFromUnitMaxCPU(unitMaxCPUs []float64) (results []*AllVirtualDumpTenantInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDumpTenantInfo{}).Where("`unit_max_cpu` IN (?)", unitMaxCPUs).Find(&results).Error

	return
}

// GetFromSlice 通过slice获取内容
func (obj *_AllVirtualDumpTenantInfoMgr) GetFromSlice(slice float64) (results []*AllVirtualDumpTenantInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDumpTenantInfo{}).Where("`slice` = ?", slice).Find(&results).Error

	return
}

// GetBatchFromSlice 批量查找
func (obj *_AllVirtualDumpTenantInfoMgr) GetBatchFromSlice(slices []float64) (results []*AllVirtualDumpTenantInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDumpTenantInfo{}).Where("`slice` IN (?)", slices).Find(&results).Error

	return
}

// GetFromRemainSlice 通过remain_slice获取内容
func (obj *_AllVirtualDumpTenantInfoMgr) GetFromRemainSlice(remainSlice float64) (results []*AllVirtualDumpTenantInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDumpTenantInfo{}).Where("`remain_slice` = ?", remainSlice).Find(&results).Error

	return
}

// GetBatchFromRemainSlice 批量查找
func (obj *_AllVirtualDumpTenantInfoMgr) GetBatchFromRemainSlice(remainSlices []float64) (results []*AllVirtualDumpTenantInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDumpTenantInfo{}).Where("`remain_slice` IN (?)", remainSlices).Find(&results).Error

	return
}

// GetFromTokenCnt 通过token_cnt获取内容
func (obj *_AllVirtualDumpTenantInfoMgr) GetFromTokenCnt(tokenCnt int64) (results []*AllVirtualDumpTenantInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDumpTenantInfo{}).Where("`token_cnt` = ?", tokenCnt).Find(&results).Error

	return
}

// GetBatchFromTokenCnt 批量查找
func (obj *_AllVirtualDumpTenantInfoMgr) GetBatchFromTokenCnt(tokenCnts []int64) (results []*AllVirtualDumpTenantInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDumpTenantInfo{}).Where("`token_cnt` IN (?)", tokenCnts).Find(&results).Error

	return
}

// GetFromAssTokenCnt 通过ass_token_cnt获取内容
func (obj *_AllVirtualDumpTenantInfoMgr) GetFromAssTokenCnt(assTokenCnt int64) (results []*AllVirtualDumpTenantInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDumpTenantInfo{}).Where("`ass_token_cnt` = ?", assTokenCnt).Find(&results).Error

	return
}

// GetBatchFromAssTokenCnt 批量查找
func (obj *_AllVirtualDumpTenantInfoMgr) GetBatchFromAssTokenCnt(assTokenCnts []int64) (results []*AllVirtualDumpTenantInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDumpTenantInfo{}).Where("`ass_token_cnt` IN (?)", assTokenCnts).Find(&results).Error

	return
}

// GetFromLqTokens 通过lq_tokens获取内容
func (obj *_AllVirtualDumpTenantInfoMgr) GetFromLqTokens(lqTokens int64) (results []*AllVirtualDumpTenantInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDumpTenantInfo{}).Where("`lq_tokens` = ?", lqTokens).Find(&results).Error

	return
}

// GetBatchFromLqTokens 批量查找
func (obj *_AllVirtualDumpTenantInfoMgr) GetBatchFromLqTokens(lqTokenss []int64) (results []*AllVirtualDumpTenantInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDumpTenantInfo{}).Where("`lq_tokens` IN (?)", lqTokenss).Find(&results).Error

	return
}

// GetFromUsedLqTokens 通过used_lq_tokens获取内容
func (obj *_AllVirtualDumpTenantInfoMgr) GetFromUsedLqTokens(usedLqTokens int64) (results []*AllVirtualDumpTenantInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDumpTenantInfo{}).Where("`used_lq_tokens` = ?", usedLqTokens).Find(&results).Error

	return
}

// GetBatchFromUsedLqTokens 批量查找
func (obj *_AllVirtualDumpTenantInfoMgr) GetBatchFromUsedLqTokens(usedLqTokenss []int64) (results []*AllVirtualDumpTenantInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDumpTenantInfo{}).Where("`used_lq_tokens` IN (?)", usedLqTokenss).Find(&results).Error

	return
}

// GetFromStopped 通过stopped获取内容
func (obj *_AllVirtualDumpTenantInfoMgr) GetFromStopped(stopped int64) (results []*AllVirtualDumpTenantInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDumpTenantInfo{}).Where("`stopped` = ?", stopped).Find(&results).Error

	return
}

// GetBatchFromStopped 批量查找
func (obj *_AllVirtualDumpTenantInfoMgr) GetBatchFromStopped(stoppeds []int64) (results []*AllVirtualDumpTenantInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDumpTenantInfo{}).Where("`stopped` IN (?)", stoppeds).Find(&results).Error

	return
}

// GetFromIDleUs 通过idle_us获取内容
func (obj *_AllVirtualDumpTenantInfoMgr) GetFromIDleUs(idleUs int64) (results []*AllVirtualDumpTenantInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDumpTenantInfo{}).Where("`idle_us` = ?", idleUs).Find(&results).Error

	return
}

// GetBatchFromIDleUs 批量查找
func (obj *_AllVirtualDumpTenantInfoMgr) GetBatchFromIDleUs(idleUss []int64) (results []*AllVirtualDumpTenantInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDumpTenantInfo{}).Where("`idle_us` IN (?)", idleUss).Find(&results).Error

	return
}

// GetFromRecvHpRPCCnt 通过recv_hp_rpc_cnt获取内容
func (obj *_AllVirtualDumpTenantInfoMgr) GetFromRecvHpRPCCnt(recvHpRPCCnt int64) (results []*AllVirtualDumpTenantInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDumpTenantInfo{}).Where("`recv_hp_rpc_cnt` = ?", recvHpRPCCnt).Find(&results).Error

	return
}

// GetBatchFromRecvHpRPCCnt 批量查找
func (obj *_AllVirtualDumpTenantInfoMgr) GetBatchFromRecvHpRPCCnt(recvHpRPCCnts []int64) (results []*AllVirtualDumpTenantInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDumpTenantInfo{}).Where("`recv_hp_rpc_cnt` IN (?)", recvHpRPCCnts).Find(&results).Error

	return
}

// GetFromRecvNpRPCCnt 通过recv_np_rpc_cnt获取内容
func (obj *_AllVirtualDumpTenantInfoMgr) GetFromRecvNpRPCCnt(recvNpRPCCnt int64) (results []*AllVirtualDumpTenantInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDumpTenantInfo{}).Where("`recv_np_rpc_cnt` = ?", recvNpRPCCnt).Find(&results).Error

	return
}

// GetBatchFromRecvNpRPCCnt 批量查找
func (obj *_AllVirtualDumpTenantInfoMgr) GetBatchFromRecvNpRPCCnt(recvNpRPCCnts []int64) (results []*AllVirtualDumpTenantInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDumpTenantInfo{}).Where("`recv_np_rpc_cnt` IN (?)", recvNpRPCCnts).Find(&results).Error

	return
}

// GetFromRecvLpRPCCnt 通过recv_lp_rpc_cnt获取内容
func (obj *_AllVirtualDumpTenantInfoMgr) GetFromRecvLpRPCCnt(recvLpRPCCnt int64) (results []*AllVirtualDumpTenantInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDumpTenantInfo{}).Where("`recv_lp_rpc_cnt` = ?", recvLpRPCCnt).Find(&results).Error

	return
}

// GetBatchFromRecvLpRPCCnt 批量查找
func (obj *_AllVirtualDumpTenantInfoMgr) GetBatchFromRecvLpRPCCnt(recvLpRPCCnts []int64) (results []*AllVirtualDumpTenantInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDumpTenantInfo{}).Where("`recv_lp_rpc_cnt` IN (?)", recvLpRPCCnts).Find(&results).Error

	return
}

// GetFromRecvMysqlCnt 通过recv_mysql_cnt获取内容
func (obj *_AllVirtualDumpTenantInfoMgr) GetFromRecvMysqlCnt(recvMysqlCnt int64) (results []*AllVirtualDumpTenantInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDumpTenantInfo{}).Where("`recv_mysql_cnt` = ?", recvMysqlCnt).Find(&results).Error

	return
}

// GetBatchFromRecvMysqlCnt 批量查找
func (obj *_AllVirtualDumpTenantInfoMgr) GetBatchFromRecvMysqlCnt(recvMysqlCnts []int64) (results []*AllVirtualDumpTenantInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDumpTenantInfo{}).Where("`recv_mysql_cnt` IN (?)", recvMysqlCnts).Find(&results).Error

	return
}

// GetFromRecvTaskCnt 通过recv_task_cnt获取内容
func (obj *_AllVirtualDumpTenantInfoMgr) GetFromRecvTaskCnt(recvTaskCnt int64) (results []*AllVirtualDumpTenantInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDumpTenantInfo{}).Where("`recv_task_cnt` = ?", recvTaskCnt).Find(&results).Error

	return
}

// GetBatchFromRecvTaskCnt 批量查找
func (obj *_AllVirtualDumpTenantInfoMgr) GetBatchFromRecvTaskCnt(recvTaskCnts []int64) (results []*AllVirtualDumpTenantInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDumpTenantInfo{}).Where("`recv_task_cnt` IN (?)", recvTaskCnts).Find(&results).Error

	return
}

// GetFromRecvLargeReqCnt 通过recv_large_req_cnt获取内容
func (obj *_AllVirtualDumpTenantInfoMgr) GetFromRecvLargeReqCnt(recvLargeReqCnt int64) (results []*AllVirtualDumpTenantInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDumpTenantInfo{}).Where("`recv_large_req_cnt` = ?", recvLargeReqCnt).Find(&results).Error

	return
}

// GetBatchFromRecvLargeReqCnt 批量查找
func (obj *_AllVirtualDumpTenantInfoMgr) GetBatchFromRecvLargeReqCnt(recvLargeReqCnts []int64) (results []*AllVirtualDumpTenantInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDumpTenantInfo{}).Where("`recv_large_req_cnt` IN (?)", recvLargeReqCnts).Find(&results).Error

	return
}

// GetFromRecvLargeQueries 通过recv_large_queries获取内容
func (obj *_AllVirtualDumpTenantInfoMgr) GetFromRecvLargeQueries(recvLargeQueries int64) (results []*AllVirtualDumpTenantInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDumpTenantInfo{}).Where("`recv_large_queries` = ?", recvLargeQueries).Find(&results).Error

	return
}

// GetBatchFromRecvLargeQueries 批量查找
func (obj *_AllVirtualDumpTenantInfoMgr) GetBatchFromRecvLargeQueries(recvLargeQueriess []int64) (results []*AllVirtualDumpTenantInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDumpTenantInfo{}).Where("`recv_large_queries` IN (?)", recvLargeQueriess).Find(&results).Error

	return
}

// GetFromActives 通过actives获取内容
func (obj *_AllVirtualDumpTenantInfoMgr) GetFromActives(actives int64) (results []*AllVirtualDumpTenantInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDumpTenantInfo{}).Where("`actives` = ?", actives).Find(&results).Error

	return
}

// GetBatchFromActives 批量查找
func (obj *_AllVirtualDumpTenantInfoMgr) GetBatchFromActives(activess []int64) (results []*AllVirtualDumpTenantInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDumpTenantInfo{}).Where("`actives` IN (?)", activess).Find(&results).Error

	return
}

// GetFromWorkers 通过workers获取内容
func (obj *_AllVirtualDumpTenantInfoMgr) GetFromWorkers(workers int64) (results []*AllVirtualDumpTenantInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDumpTenantInfo{}).Where("`workers` = ?", workers).Find(&results).Error

	return
}

// GetBatchFromWorkers 批量查找
func (obj *_AllVirtualDumpTenantInfoMgr) GetBatchFromWorkers(workerss []int64) (results []*AllVirtualDumpTenantInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDumpTenantInfo{}).Where("`workers` IN (?)", workerss).Find(&results).Error

	return
}

// GetFromLqWaitingWorkers 通过lq_waiting_workers获取内容
func (obj *_AllVirtualDumpTenantInfoMgr) GetFromLqWaitingWorkers(lqWaitingWorkers int64) (results []*AllVirtualDumpTenantInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDumpTenantInfo{}).Where("`lq_waiting_workers` = ?", lqWaitingWorkers).Find(&results).Error

	return
}

// GetBatchFromLqWaitingWorkers 批量查找
func (obj *_AllVirtualDumpTenantInfoMgr) GetBatchFromLqWaitingWorkers(lqWaitingWorkerss []int64) (results []*AllVirtualDumpTenantInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDumpTenantInfo{}).Where("`lq_waiting_workers` IN (?)", lqWaitingWorkerss).Find(&results).Error

	return
}

// GetFromReqQueueTotalSize 通过req_queue_total_size获取内容
func (obj *_AllVirtualDumpTenantInfoMgr) GetFromReqQueueTotalSize(reqQueueTotalSize int64) (results []*AllVirtualDumpTenantInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDumpTenantInfo{}).Where("`req_queue_total_size` = ?", reqQueueTotalSize).Find(&results).Error

	return
}

// GetBatchFromReqQueueTotalSize 批量查找
func (obj *_AllVirtualDumpTenantInfoMgr) GetBatchFromReqQueueTotalSize(reqQueueTotalSizes []int64) (results []*AllVirtualDumpTenantInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDumpTenantInfo{}).Where("`req_queue_total_size` IN (?)", reqQueueTotalSizes).Find(&results).Error

	return
}

// GetFromQueue0 通过queue_0获取内容
func (obj *_AllVirtualDumpTenantInfoMgr) GetFromQueue0(queue0 int64) (results []*AllVirtualDumpTenantInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDumpTenantInfo{}).Where("`queue_0` = ?", queue0).Find(&results).Error

	return
}

// GetBatchFromQueue0 批量查找
func (obj *_AllVirtualDumpTenantInfoMgr) GetBatchFromQueue0(queue0s []int64) (results []*AllVirtualDumpTenantInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDumpTenantInfo{}).Where("`queue_0` IN (?)", queue0s).Find(&results).Error

	return
}

// GetFromQueue1 通过queue_1获取内容
func (obj *_AllVirtualDumpTenantInfoMgr) GetFromQueue1(queue1 int64) (results []*AllVirtualDumpTenantInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDumpTenantInfo{}).Where("`queue_1` = ?", queue1).Find(&results).Error

	return
}

// GetBatchFromQueue1 批量查找
func (obj *_AllVirtualDumpTenantInfoMgr) GetBatchFromQueue1(queue1s []int64) (results []*AllVirtualDumpTenantInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDumpTenantInfo{}).Where("`queue_1` IN (?)", queue1s).Find(&results).Error

	return
}

// GetFromQueue2 通过queue_2获取内容
func (obj *_AllVirtualDumpTenantInfoMgr) GetFromQueue2(queue2 int64) (results []*AllVirtualDumpTenantInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDumpTenantInfo{}).Where("`queue_2` = ?", queue2).Find(&results).Error

	return
}

// GetBatchFromQueue2 批量查找
func (obj *_AllVirtualDumpTenantInfoMgr) GetBatchFromQueue2(queue2s []int64) (results []*AllVirtualDumpTenantInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDumpTenantInfo{}).Where("`queue_2` IN (?)", queue2s).Find(&results).Error

	return
}

// GetFromQueue3 通过queue_3获取内容
func (obj *_AllVirtualDumpTenantInfoMgr) GetFromQueue3(queue3 int64) (results []*AllVirtualDumpTenantInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDumpTenantInfo{}).Where("`queue_3` = ?", queue3).Find(&results).Error

	return
}

// GetBatchFromQueue3 批量查找
func (obj *_AllVirtualDumpTenantInfoMgr) GetBatchFromQueue3(queue3s []int64) (results []*AllVirtualDumpTenantInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDumpTenantInfo{}).Where("`queue_3` IN (?)", queue3s).Find(&results).Error

	return
}

// GetFromQueue4 通过queue_4获取内容
func (obj *_AllVirtualDumpTenantInfoMgr) GetFromQueue4(queue4 int64) (results []*AllVirtualDumpTenantInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDumpTenantInfo{}).Where("`queue_4` = ?", queue4).Find(&results).Error

	return
}

// GetBatchFromQueue4 批量查找
func (obj *_AllVirtualDumpTenantInfoMgr) GetBatchFromQueue4(queue4s []int64) (results []*AllVirtualDumpTenantInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDumpTenantInfo{}).Where("`queue_4` IN (?)", queue4s).Find(&results).Error

	return
}

// GetFromQueue5 通过queue_5获取内容
func (obj *_AllVirtualDumpTenantInfoMgr) GetFromQueue5(queue5 int64) (results []*AllVirtualDumpTenantInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDumpTenantInfo{}).Where("`queue_5` = ?", queue5).Find(&results).Error

	return
}

// GetBatchFromQueue5 批量查找
func (obj *_AllVirtualDumpTenantInfoMgr) GetBatchFromQueue5(queue5s []int64) (results []*AllVirtualDumpTenantInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDumpTenantInfo{}).Where("`queue_5` IN (?)", queue5s).Find(&results).Error

	return
}

// GetFromLargeQueued 通过large_queued获取内容
func (obj *_AllVirtualDumpTenantInfoMgr) GetFromLargeQueued(largeQueued int64) (results []*AllVirtualDumpTenantInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDumpTenantInfo{}).Where("`large_queued` = ?", largeQueued).Find(&results).Error

	return
}

// GetBatchFromLargeQueued 批量查找
func (obj *_AllVirtualDumpTenantInfoMgr) GetBatchFromLargeQueued(largeQueueds []int64) (results []*AllVirtualDumpTenantInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDumpTenantInfo{}).Where("`large_queued` IN (?)", largeQueueds).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllVirtualDumpTenantInfoMgr) FetchByPrimaryKey(svrIP string, svrPort int64, tenantID int64) (result AllVirtualDumpTenantInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDumpTenantInfo{}).Where("`svr_ip` = ? AND `svr_port` = ? AND `tenant_id` = ?", svrIP, svrPort, tenantID).First(&result).Error

	return
}
