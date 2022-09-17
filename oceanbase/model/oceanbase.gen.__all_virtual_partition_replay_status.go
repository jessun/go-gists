package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _AllVirtualPartitionReplayStatusMgr struct {
	*_BaseMgr
}

// AllVirtualPartitionReplayStatusMgr open func
func AllVirtualPartitionReplayStatusMgr(db *gorm.DB) *_AllVirtualPartitionReplayStatusMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualPartitionReplayStatusMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualPartitionReplayStatusMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_partition_replay_status"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualPartitionReplayStatusMgr) GetTableName() string {
	return "__all_virtual_partition_replay_status"
}

// Reset 重置gorm会话
func (obj *_AllVirtualPartitionReplayStatusMgr) Reset() *_AllVirtualPartitionReplayStatusMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualPartitionReplayStatusMgr) Get() (result AllVirtualPartitionReplayStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionReplayStatus{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualPartitionReplayStatusMgr) Gets() (results []*AllVirtualPartitionReplayStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionReplayStatus{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualPartitionReplayStatusMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionReplayStatus{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithSvrIP svr_ip获取
func (obj *_AllVirtualPartitionReplayStatusMgr) WithSvrIP(svrIP string) Option {
	return optionFunc(func(o *options) { o.query["svr_ip"] = svrIP })
}

// WithSvrPort svr_port获取
func (obj *_AllVirtualPartitionReplayStatusMgr) WithSvrPort(svrPort int64) Option {
	return optionFunc(func(o *options) { o.query["svr_port"] = svrPort })
}

// WithTenantID tenant_id获取
func (obj *_AllVirtualPartitionReplayStatusMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithTableID table_id获取
func (obj *_AllVirtualPartitionReplayStatusMgr) WithTableID(tableID int64) Option {
	return optionFunc(func(o *options) { o.query["table_id"] = tableID })
}

// WithPartitionIDx partition_idx获取
func (obj *_AllVirtualPartitionReplayStatusMgr) WithPartitionIDx(partitionIDx int64) Option {
	return optionFunc(func(o *options) { o.query["partition_idx"] = partitionIDx })
}

// WithPartitionCnt partition_cnt获取
func (obj *_AllVirtualPartitionReplayStatusMgr) WithPartitionCnt(partitionCnt int64) Option {
	return optionFunc(func(o *options) { o.query["partition_cnt"] = partitionCnt })
}

// WithPendingTaskCount pending_task_count获取
func (obj *_AllVirtualPartitionReplayStatusMgr) WithPendingTaskCount(pendingTaskCount int64) Option {
	return optionFunc(func(o *options) { o.query["pending_task_count"] = pendingTaskCount })
}

// WithRetriedTaskCount retried_task_count获取
func (obj *_AllVirtualPartitionReplayStatusMgr) WithRetriedTaskCount(retriedTaskCount int64) Option {
	return optionFunc(func(o *options) { o.query["retried_task_count"] = retriedTaskCount })
}

// WithPostBarrierStatus post_barrier_status获取
func (obj *_AllVirtualPartitionReplayStatusMgr) WithPostBarrierStatus(postBarrierStatus string) Option {
	return optionFunc(func(o *options) { o.query["post_barrier_status"] = postBarrierStatus })
}

// WithIsEnabled is_enabled获取
func (obj *_AllVirtualPartitionReplayStatusMgr) WithIsEnabled(isEnabled int64) Option {
	return optionFunc(func(o *options) { o.query["is_enabled"] = isEnabled })
}

// WithMaxConfirmedLogID max_confirmed_log_id获取
func (obj *_AllVirtualPartitionReplayStatusMgr) WithMaxConfirmedLogID(maxConfirmedLogID uint64) Option {
	return optionFunc(func(o *options) { o.query["max_confirmed_log_id"] = maxConfirmedLogID })
}

// WithLastReplayLogID last_replay_log_id获取
func (obj *_AllVirtualPartitionReplayStatusMgr) WithLastReplayLogID(lastReplayLogID uint64) Option {
	return optionFunc(func(o *options) { o.query["last_replay_log_id"] = lastReplayLogID })
}

// WithLastReplayLogType last_replay_log_type获取
func (obj *_AllVirtualPartitionReplayStatusMgr) WithLastReplayLogType(lastReplayLogType string) Option {
	return optionFunc(func(o *options) { o.query["last_replay_log_type"] = lastReplayLogType })
}

// WithTotalSubmmitedTaskCount total_submmited_task_count获取
func (obj *_AllVirtualPartitionReplayStatusMgr) WithTotalSubmmitedTaskCount(totalSubmmitedTaskCount int64) Option {
	return optionFunc(func(o *options) { o.query["total_submmited_task_count"] = totalSubmmitedTaskCount })
}

// WithTotalReplayedTaskCount total_replayed_task_count获取
func (obj *_AllVirtualPartitionReplayStatusMgr) WithTotalReplayedTaskCount(totalReplayedTaskCount int64) Option {
	return optionFunc(func(o *options) { o.query["total_replayed_task_count"] = totalReplayedTaskCount })
}

// WithNextSubmitLogID next_submit_log_id获取
func (obj *_AllVirtualPartitionReplayStatusMgr) WithNextSubmitLogID(nextSubmitLogID uint64) Option {
	return optionFunc(func(o *options) { o.query["next_submit_log_id"] = nextSubmitLogID })
}

// WithNextSubmitLogTs next_submit_log_ts获取
func (obj *_AllVirtualPartitionReplayStatusMgr) WithNextSubmitLogTs(nextSubmitLogTs int64) Option {
	return optionFunc(func(o *options) { o.query["next_submit_log_ts"] = nextSubmitLogTs })
}

// WithLastSlideOutLogID last_slide_out_log_id获取
func (obj *_AllVirtualPartitionReplayStatusMgr) WithLastSlideOutLogID(lastSlideOutLogID uint64) Option {
	return optionFunc(func(o *options) { o.query["last_slide_out_log_id"] = lastSlideOutLogID })
}

// WithLastSlideOutLogTs last_slide_out_log_ts获取
func (obj *_AllVirtualPartitionReplayStatusMgr) WithLastSlideOutLogTs(lastSlideOutLogTs int64) Option {
	return optionFunc(func(o *options) { o.query["last_slide_out_log_ts"] = lastSlideOutLogTs })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualPartitionReplayStatusMgr) GetByOption(opts ...Option) (result AllVirtualPartitionReplayStatus, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionReplayStatus{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualPartitionReplayStatusMgr) GetByOptions(opts ...Option) (results []*AllVirtualPartitionReplayStatus, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionReplayStatus{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualPartitionReplayStatusMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualPartitionReplayStatus, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionReplayStatus{}).Where(options.query)
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
func (obj *_AllVirtualPartitionReplayStatusMgr) GetFromSvrIP(svrIP string) (results []*AllVirtualPartitionReplayStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionReplayStatus{}).Where("`svr_ip` = ?", svrIP).Find(&results).Error

	return
}

// GetBatchFromSvrIP 批量查找
func (obj *_AllVirtualPartitionReplayStatusMgr) GetBatchFromSvrIP(svrIPs []string) (results []*AllVirtualPartitionReplayStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionReplayStatus{}).Where("`svr_ip` IN (?)", svrIPs).Find(&results).Error

	return
}

// GetFromSvrPort 通过svr_port获取内容
func (obj *_AllVirtualPartitionReplayStatusMgr) GetFromSvrPort(svrPort int64) (results []*AllVirtualPartitionReplayStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionReplayStatus{}).Where("`svr_port` = ?", svrPort).Find(&results).Error

	return
}

// GetBatchFromSvrPort 批量查找
func (obj *_AllVirtualPartitionReplayStatusMgr) GetBatchFromSvrPort(svrPorts []int64) (results []*AllVirtualPartitionReplayStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionReplayStatus{}).Where("`svr_port` IN (?)", svrPorts).Find(&results).Error

	return
}

// GetFromTenantID 通过tenant_id获取内容
func (obj *_AllVirtualPartitionReplayStatusMgr) GetFromTenantID(tenantID int64) (results []*AllVirtualPartitionReplayStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionReplayStatus{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllVirtualPartitionReplayStatusMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllVirtualPartitionReplayStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionReplayStatus{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromTableID 通过table_id获取内容
func (obj *_AllVirtualPartitionReplayStatusMgr) GetFromTableID(tableID int64) (results []*AllVirtualPartitionReplayStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionReplayStatus{}).Where("`table_id` = ?", tableID).Find(&results).Error

	return
}

// GetBatchFromTableID 批量查找
func (obj *_AllVirtualPartitionReplayStatusMgr) GetBatchFromTableID(tableIDs []int64) (results []*AllVirtualPartitionReplayStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionReplayStatus{}).Where("`table_id` IN (?)", tableIDs).Find(&results).Error

	return
}

// GetFromPartitionIDx 通过partition_idx获取内容
func (obj *_AllVirtualPartitionReplayStatusMgr) GetFromPartitionIDx(partitionIDx int64) (results []*AllVirtualPartitionReplayStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionReplayStatus{}).Where("`partition_idx` = ?", partitionIDx).Find(&results).Error

	return
}

// GetBatchFromPartitionIDx 批量查找
func (obj *_AllVirtualPartitionReplayStatusMgr) GetBatchFromPartitionIDx(partitionIDxs []int64) (results []*AllVirtualPartitionReplayStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionReplayStatus{}).Where("`partition_idx` IN (?)", partitionIDxs).Find(&results).Error

	return
}

// GetFromPartitionCnt 通过partition_cnt获取内容
func (obj *_AllVirtualPartitionReplayStatusMgr) GetFromPartitionCnt(partitionCnt int64) (results []*AllVirtualPartitionReplayStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionReplayStatus{}).Where("`partition_cnt` = ?", partitionCnt).Find(&results).Error

	return
}

// GetBatchFromPartitionCnt 批量查找
func (obj *_AllVirtualPartitionReplayStatusMgr) GetBatchFromPartitionCnt(partitionCnts []int64) (results []*AllVirtualPartitionReplayStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionReplayStatus{}).Where("`partition_cnt` IN (?)", partitionCnts).Find(&results).Error

	return
}

// GetFromPendingTaskCount 通过pending_task_count获取内容
func (obj *_AllVirtualPartitionReplayStatusMgr) GetFromPendingTaskCount(pendingTaskCount int64) (results []*AllVirtualPartitionReplayStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionReplayStatus{}).Where("`pending_task_count` = ?", pendingTaskCount).Find(&results).Error

	return
}

// GetBatchFromPendingTaskCount 批量查找
func (obj *_AllVirtualPartitionReplayStatusMgr) GetBatchFromPendingTaskCount(pendingTaskCounts []int64) (results []*AllVirtualPartitionReplayStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionReplayStatus{}).Where("`pending_task_count` IN (?)", pendingTaskCounts).Find(&results).Error

	return
}

// GetFromRetriedTaskCount 通过retried_task_count获取内容
func (obj *_AllVirtualPartitionReplayStatusMgr) GetFromRetriedTaskCount(retriedTaskCount int64) (results []*AllVirtualPartitionReplayStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionReplayStatus{}).Where("`retried_task_count` = ?", retriedTaskCount).Find(&results).Error

	return
}

// GetBatchFromRetriedTaskCount 批量查找
func (obj *_AllVirtualPartitionReplayStatusMgr) GetBatchFromRetriedTaskCount(retriedTaskCounts []int64) (results []*AllVirtualPartitionReplayStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionReplayStatus{}).Where("`retried_task_count` IN (?)", retriedTaskCounts).Find(&results).Error

	return
}

// GetFromPostBarrierStatus 通过post_barrier_status获取内容
func (obj *_AllVirtualPartitionReplayStatusMgr) GetFromPostBarrierStatus(postBarrierStatus string) (results []*AllVirtualPartitionReplayStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionReplayStatus{}).Where("`post_barrier_status` = ?", postBarrierStatus).Find(&results).Error

	return
}

// GetBatchFromPostBarrierStatus 批量查找
func (obj *_AllVirtualPartitionReplayStatusMgr) GetBatchFromPostBarrierStatus(postBarrierStatuss []string) (results []*AllVirtualPartitionReplayStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionReplayStatus{}).Where("`post_barrier_status` IN (?)", postBarrierStatuss).Find(&results).Error

	return
}

// GetFromIsEnabled 通过is_enabled获取内容
func (obj *_AllVirtualPartitionReplayStatusMgr) GetFromIsEnabled(isEnabled int64) (results []*AllVirtualPartitionReplayStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionReplayStatus{}).Where("`is_enabled` = ?", isEnabled).Find(&results).Error

	return
}

// GetBatchFromIsEnabled 批量查找
func (obj *_AllVirtualPartitionReplayStatusMgr) GetBatchFromIsEnabled(isEnableds []int64) (results []*AllVirtualPartitionReplayStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionReplayStatus{}).Where("`is_enabled` IN (?)", isEnableds).Find(&results).Error

	return
}

// GetFromMaxConfirmedLogID 通过max_confirmed_log_id获取内容
func (obj *_AllVirtualPartitionReplayStatusMgr) GetFromMaxConfirmedLogID(maxConfirmedLogID uint64) (results []*AllVirtualPartitionReplayStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionReplayStatus{}).Where("`max_confirmed_log_id` = ?", maxConfirmedLogID).Find(&results).Error

	return
}

// GetBatchFromMaxConfirmedLogID 批量查找
func (obj *_AllVirtualPartitionReplayStatusMgr) GetBatchFromMaxConfirmedLogID(maxConfirmedLogIDs []uint64) (results []*AllVirtualPartitionReplayStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionReplayStatus{}).Where("`max_confirmed_log_id` IN (?)", maxConfirmedLogIDs).Find(&results).Error

	return
}

// GetFromLastReplayLogID 通过last_replay_log_id获取内容
func (obj *_AllVirtualPartitionReplayStatusMgr) GetFromLastReplayLogID(lastReplayLogID uint64) (results []*AllVirtualPartitionReplayStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionReplayStatus{}).Where("`last_replay_log_id` = ?", lastReplayLogID).Find(&results).Error

	return
}

// GetBatchFromLastReplayLogID 批量查找
func (obj *_AllVirtualPartitionReplayStatusMgr) GetBatchFromLastReplayLogID(lastReplayLogIDs []uint64) (results []*AllVirtualPartitionReplayStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionReplayStatus{}).Where("`last_replay_log_id` IN (?)", lastReplayLogIDs).Find(&results).Error

	return
}

// GetFromLastReplayLogType 通过last_replay_log_type获取内容
func (obj *_AllVirtualPartitionReplayStatusMgr) GetFromLastReplayLogType(lastReplayLogType string) (results []*AllVirtualPartitionReplayStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionReplayStatus{}).Where("`last_replay_log_type` = ?", lastReplayLogType).Find(&results).Error

	return
}

// GetBatchFromLastReplayLogType 批量查找
func (obj *_AllVirtualPartitionReplayStatusMgr) GetBatchFromLastReplayLogType(lastReplayLogTypes []string) (results []*AllVirtualPartitionReplayStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionReplayStatus{}).Where("`last_replay_log_type` IN (?)", lastReplayLogTypes).Find(&results).Error

	return
}

// GetFromTotalSubmmitedTaskCount 通过total_submmited_task_count获取内容
func (obj *_AllVirtualPartitionReplayStatusMgr) GetFromTotalSubmmitedTaskCount(totalSubmmitedTaskCount int64) (results []*AllVirtualPartitionReplayStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionReplayStatus{}).Where("`total_submmited_task_count` = ?", totalSubmmitedTaskCount).Find(&results).Error

	return
}

// GetBatchFromTotalSubmmitedTaskCount 批量查找
func (obj *_AllVirtualPartitionReplayStatusMgr) GetBatchFromTotalSubmmitedTaskCount(totalSubmmitedTaskCounts []int64) (results []*AllVirtualPartitionReplayStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionReplayStatus{}).Where("`total_submmited_task_count` IN (?)", totalSubmmitedTaskCounts).Find(&results).Error

	return
}

// GetFromTotalReplayedTaskCount 通过total_replayed_task_count获取内容
func (obj *_AllVirtualPartitionReplayStatusMgr) GetFromTotalReplayedTaskCount(totalReplayedTaskCount int64) (results []*AllVirtualPartitionReplayStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionReplayStatus{}).Where("`total_replayed_task_count` = ?", totalReplayedTaskCount).Find(&results).Error

	return
}

// GetBatchFromTotalReplayedTaskCount 批量查找
func (obj *_AllVirtualPartitionReplayStatusMgr) GetBatchFromTotalReplayedTaskCount(totalReplayedTaskCounts []int64) (results []*AllVirtualPartitionReplayStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionReplayStatus{}).Where("`total_replayed_task_count` IN (?)", totalReplayedTaskCounts).Find(&results).Error

	return
}

// GetFromNextSubmitLogID 通过next_submit_log_id获取内容
func (obj *_AllVirtualPartitionReplayStatusMgr) GetFromNextSubmitLogID(nextSubmitLogID uint64) (results []*AllVirtualPartitionReplayStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionReplayStatus{}).Where("`next_submit_log_id` = ?", nextSubmitLogID).Find(&results).Error

	return
}

// GetBatchFromNextSubmitLogID 批量查找
func (obj *_AllVirtualPartitionReplayStatusMgr) GetBatchFromNextSubmitLogID(nextSubmitLogIDs []uint64) (results []*AllVirtualPartitionReplayStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionReplayStatus{}).Where("`next_submit_log_id` IN (?)", nextSubmitLogIDs).Find(&results).Error

	return
}

// GetFromNextSubmitLogTs 通过next_submit_log_ts获取内容
func (obj *_AllVirtualPartitionReplayStatusMgr) GetFromNextSubmitLogTs(nextSubmitLogTs int64) (results []*AllVirtualPartitionReplayStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionReplayStatus{}).Where("`next_submit_log_ts` = ?", nextSubmitLogTs).Find(&results).Error

	return
}

// GetBatchFromNextSubmitLogTs 批量查找
func (obj *_AllVirtualPartitionReplayStatusMgr) GetBatchFromNextSubmitLogTs(nextSubmitLogTss []int64) (results []*AllVirtualPartitionReplayStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionReplayStatus{}).Where("`next_submit_log_ts` IN (?)", nextSubmitLogTss).Find(&results).Error

	return
}

// GetFromLastSlideOutLogID 通过last_slide_out_log_id获取内容
func (obj *_AllVirtualPartitionReplayStatusMgr) GetFromLastSlideOutLogID(lastSlideOutLogID uint64) (results []*AllVirtualPartitionReplayStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionReplayStatus{}).Where("`last_slide_out_log_id` = ?", lastSlideOutLogID).Find(&results).Error

	return
}

// GetBatchFromLastSlideOutLogID 批量查找
func (obj *_AllVirtualPartitionReplayStatusMgr) GetBatchFromLastSlideOutLogID(lastSlideOutLogIDs []uint64) (results []*AllVirtualPartitionReplayStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionReplayStatus{}).Where("`last_slide_out_log_id` IN (?)", lastSlideOutLogIDs).Find(&results).Error

	return
}

// GetFromLastSlideOutLogTs 通过last_slide_out_log_ts获取内容
func (obj *_AllVirtualPartitionReplayStatusMgr) GetFromLastSlideOutLogTs(lastSlideOutLogTs int64) (results []*AllVirtualPartitionReplayStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionReplayStatus{}).Where("`last_slide_out_log_ts` = ?", lastSlideOutLogTs).Find(&results).Error

	return
}

// GetBatchFromLastSlideOutLogTs 批量查找
func (obj *_AllVirtualPartitionReplayStatusMgr) GetBatchFromLastSlideOutLogTs(lastSlideOutLogTss []int64) (results []*AllVirtualPartitionReplayStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionReplayStatus{}).Where("`last_slide_out_log_ts` IN (?)", lastSlideOutLogTss).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////
