package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _AllVirtualDtlMemoryMgr struct {
	*_BaseMgr
}

// AllVirtualDtlMemoryMgr open func
func AllVirtualDtlMemoryMgr(db *gorm.DB) *_AllVirtualDtlMemoryMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualDtlMemoryMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualDtlMemoryMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_dtl_memory"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualDtlMemoryMgr) GetTableName() string {
	return "__all_virtual_dtl_memory"
}

// Reset 重置gorm会话
func (obj *_AllVirtualDtlMemoryMgr) Reset() *_AllVirtualDtlMemoryMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualDtlMemoryMgr) Get() (result AllVirtualDtlMemory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDtlMemory{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualDtlMemoryMgr) Gets() (results []*AllVirtualDtlMemory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDtlMemory{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualDtlMemoryMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualDtlMemory{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithSvrIP svr_ip获取
func (obj *_AllVirtualDtlMemoryMgr) WithSvrIP(svrIP string) Option {
	return optionFunc(func(o *options) { o.query["svr_ip"] = svrIP })
}

// WithSvrPort svr_port获取
func (obj *_AllVirtualDtlMemoryMgr) WithSvrPort(svrPort int64) Option {
	return optionFunc(func(o *options) { o.query["svr_port"] = svrPort })
}

// WithTenantID tenant_id获取
func (obj *_AllVirtualDtlMemoryMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithChannelTotalCnt channel_total_cnt获取
func (obj *_AllVirtualDtlMemoryMgr) WithChannelTotalCnt(channelTotalCnt int64) Option {
	return optionFunc(func(o *options) { o.query["channel_total_cnt"] = channelTotalCnt })
}

// WithChannelBlockCnt channel_block_cnt获取
func (obj *_AllVirtualDtlMemoryMgr) WithChannelBlockCnt(channelBlockCnt int64) Option {
	return optionFunc(func(o *options) { o.query["channel_block_cnt"] = channelBlockCnt })
}

// WithMaxParallelCnt max_parallel_cnt获取
func (obj *_AllVirtualDtlMemoryMgr) WithMaxParallelCnt(maxParallelCnt int64) Option {
	return optionFunc(func(o *options) { o.query["max_parallel_cnt"] = maxParallelCnt })
}

// WithMaxBlockedBufferSize max_blocked_buffer_size获取
func (obj *_AllVirtualDtlMemoryMgr) WithMaxBlockedBufferSize(maxBlockedBufferSize int64) Option {
	return optionFunc(func(o *options) { o.query["max_blocked_buffer_size"] = maxBlockedBufferSize })
}

// WithAccumulatedBlockedCnt accumulated_blocked_cnt获取
func (obj *_AllVirtualDtlMemoryMgr) WithAccumulatedBlockedCnt(accumulatedBlockedCnt int64) Option {
	return optionFunc(func(o *options) { o.query["accumulated_blocked_cnt"] = accumulatedBlockedCnt })
}

// WithCurrentBufferUsed current_buffer_used获取
func (obj *_AllVirtualDtlMemoryMgr) WithCurrentBufferUsed(currentBufferUsed int64) Option {
	return optionFunc(func(o *options) { o.query["current_buffer_used"] = currentBufferUsed })
}

// WithSeqno seqno获取
func (obj *_AllVirtualDtlMemoryMgr) WithSeqno(seqno int64) Option {
	return optionFunc(func(o *options) { o.query["seqno"] = seqno })
}

// WithAllocCnt alloc_cnt获取
func (obj *_AllVirtualDtlMemoryMgr) WithAllocCnt(allocCnt int64) Option {
	return optionFunc(func(o *options) { o.query["alloc_cnt"] = allocCnt })
}

// WithFreeCnt free_cnt获取
func (obj *_AllVirtualDtlMemoryMgr) WithFreeCnt(freeCnt int64) Option {
	return optionFunc(func(o *options) { o.query["free_cnt"] = freeCnt })
}

// WithFreeQueueLen free_queue_len获取
func (obj *_AllVirtualDtlMemoryMgr) WithFreeQueueLen(freeQueueLen int64) Option {
	return optionFunc(func(o *options) { o.query["free_queue_len"] = freeQueueLen })
}

// WithTotalMemorySize total_memory_size获取
func (obj *_AllVirtualDtlMemoryMgr) WithTotalMemorySize(totalMemorySize int64) Option {
	return optionFunc(func(o *options) { o.query["total_memory_size"] = totalMemorySize })
}

// WithRealAllocCnt real_alloc_cnt获取
func (obj *_AllVirtualDtlMemoryMgr) WithRealAllocCnt(realAllocCnt int64) Option {
	return optionFunc(func(o *options) { o.query["real_alloc_cnt"] = realAllocCnt })
}

// WithRealFreeCnt real_free_cnt获取
func (obj *_AllVirtualDtlMemoryMgr) WithRealFreeCnt(realFreeCnt int64) Option {
	return optionFunc(func(o *options) { o.query["real_free_cnt"] = realFreeCnt })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualDtlMemoryMgr) GetByOption(opts ...Option) (result AllVirtualDtlMemory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDtlMemory{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualDtlMemoryMgr) GetByOptions(opts ...Option) (results []*AllVirtualDtlMemory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDtlMemory{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualDtlMemoryMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualDtlMemory, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualDtlMemory{}).Where(options.query)
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
func (obj *_AllVirtualDtlMemoryMgr) GetFromSvrIP(svrIP string) (results []*AllVirtualDtlMemory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDtlMemory{}).Where("`svr_ip` = ?", svrIP).Find(&results).Error

	return
}

// GetBatchFromSvrIP 批量查找
func (obj *_AllVirtualDtlMemoryMgr) GetBatchFromSvrIP(svrIPs []string) (results []*AllVirtualDtlMemory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDtlMemory{}).Where("`svr_ip` IN (?)", svrIPs).Find(&results).Error

	return
}

// GetFromSvrPort 通过svr_port获取内容
func (obj *_AllVirtualDtlMemoryMgr) GetFromSvrPort(svrPort int64) (results []*AllVirtualDtlMemory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDtlMemory{}).Where("`svr_port` = ?", svrPort).Find(&results).Error

	return
}

// GetBatchFromSvrPort 批量查找
func (obj *_AllVirtualDtlMemoryMgr) GetBatchFromSvrPort(svrPorts []int64) (results []*AllVirtualDtlMemory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDtlMemory{}).Where("`svr_port` IN (?)", svrPorts).Find(&results).Error

	return
}

// GetFromTenantID 通过tenant_id获取内容
func (obj *_AllVirtualDtlMemoryMgr) GetFromTenantID(tenantID int64) (results []*AllVirtualDtlMemory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDtlMemory{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllVirtualDtlMemoryMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllVirtualDtlMemory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDtlMemory{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromChannelTotalCnt 通过channel_total_cnt获取内容
func (obj *_AllVirtualDtlMemoryMgr) GetFromChannelTotalCnt(channelTotalCnt int64) (results []*AllVirtualDtlMemory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDtlMemory{}).Where("`channel_total_cnt` = ?", channelTotalCnt).Find(&results).Error

	return
}

// GetBatchFromChannelTotalCnt 批量查找
func (obj *_AllVirtualDtlMemoryMgr) GetBatchFromChannelTotalCnt(channelTotalCnts []int64) (results []*AllVirtualDtlMemory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDtlMemory{}).Where("`channel_total_cnt` IN (?)", channelTotalCnts).Find(&results).Error

	return
}

// GetFromChannelBlockCnt 通过channel_block_cnt获取内容
func (obj *_AllVirtualDtlMemoryMgr) GetFromChannelBlockCnt(channelBlockCnt int64) (results []*AllVirtualDtlMemory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDtlMemory{}).Where("`channel_block_cnt` = ?", channelBlockCnt).Find(&results).Error

	return
}

// GetBatchFromChannelBlockCnt 批量查找
func (obj *_AllVirtualDtlMemoryMgr) GetBatchFromChannelBlockCnt(channelBlockCnts []int64) (results []*AllVirtualDtlMemory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDtlMemory{}).Where("`channel_block_cnt` IN (?)", channelBlockCnts).Find(&results).Error

	return
}

// GetFromMaxParallelCnt 通过max_parallel_cnt获取内容
func (obj *_AllVirtualDtlMemoryMgr) GetFromMaxParallelCnt(maxParallelCnt int64) (results []*AllVirtualDtlMemory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDtlMemory{}).Where("`max_parallel_cnt` = ?", maxParallelCnt).Find(&results).Error

	return
}

// GetBatchFromMaxParallelCnt 批量查找
func (obj *_AllVirtualDtlMemoryMgr) GetBatchFromMaxParallelCnt(maxParallelCnts []int64) (results []*AllVirtualDtlMemory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDtlMemory{}).Where("`max_parallel_cnt` IN (?)", maxParallelCnts).Find(&results).Error

	return
}

// GetFromMaxBlockedBufferSize 通过max_blocked_buffer_size获取内容
func (obj *_AllVirtualDtlMemoryMgr) GetFromMaxBlockedBufferSize(maxBlockedBufferSize int64) (results []*AllVirtualDtlMemory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDtlMemory{}).Where("`max_blocked_buffer_size` = ?", maxBlockedBufferSize).Find(&results).Error

	return
}

// GetBatchFromMaxBlockedBufferSize 批量查找
func (obj *_AllVirtualDtlMemoryMgr) GetBatchFromMaxBlockedBufferSize(maxBlockedBufferSizes []int64) (results []*AllVirtualDtlMemory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDtlMemory{}).Where("`max_blocked_buffer_size` IN (?)", maxBlockedBufferSizes).Find(&results).Error

	return
}

// GetFromAccumulatedBlockedCnt 通过accumulated_blocked_cnt获取内容
func (obj *_AllVirtualDtlMemoryMgr) GetFromAccumulatedBlockedCnt(accumulatedBlockedCnt int64) (results []*AllVirtualDtlMemory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDtlMemory{}).Where("`accumulated_blocked_cnt` = ?", accumulatedBlockedCnt).Find(&results).Error

	return
}

// GetBatchFromAccumulatedBlockedCnt 批量查找
func (obj *_AllVirtualDtlMemoryMgr) GetBatchFromAccumulatedBlockedCnt(accumulatedBlockedCnts []int64) (results []*AllVirtualDtlMemory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDtlMemory{}).Where("`accumulated_blocked_cnt` IN (?)", accumulatedBlockedCnts).Find(&results).Error

	return
}

// GetFromCurrentBufferUsed 通过current_buffer_used获取内容
func (obj *_AllVirtualDtlMemoryMgr) GetFromCurrentBufferUsed(currentBufferUsed int64) (results []*AllVirtualDtlMemory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDtlMemory{}).Where("`current_buffer_used` = ?", currentBufferUsed).Find(&results).Error

	return
}

// GetBatchFromCurrentBufferUsed 批量查找
func (obj *_AllVirtualDtlMemoryMgr) GetBatchFromCurrentBufferUsed(currentBufferUseds []int64) (results []*AllVirtualDtlMemory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDtlMemory{}).Where("`current_buffer_used` IN (?)", currentBufferUseds).Find(&results).Error

	return
}

// GetFromSeqno 通过seqno获取内容
func (obj *_AllVirtualDtlMemoryMgr) GetFromSeqno(seqno int64) (results []*AllVirtualDtlMemory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDtlMemory{}).Where("`seqno` = ?", seqno).Find(&results).Error

	return
}

// GetBatchFromSeqno 批量查找
func (obj *_AllVirtualDtlMemoryMgr) GetBatchFromSeqno(seqnos []int64) (results []*AllVirtualDtlMemory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDtlMemory{}).Where("`seqno` IN (?)", seqnos).Find(&results).Error

	return
}

// GetFromAllocCnt 通过alloc_cnt获取内容
func (obj *_AllVirtualDtlMemoryMgr) GetFromAllocCnt(allocCnt int64) (results []*AllVirtualDtlMemory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDtlMemory{}).Where("`alloc_cnt` = ?", allocCnt).Find(&results).Error

	return
}

// GetBatchFromAllocCnt 批量查找
func (obj *_AllVirtualDtlMemoryMgr) GetBatchFromAllocCnt(allocCnts []int64) (results []*AllVirtualDtlMemory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDtlMemory{}).Where("`alloc_cnt` IN (?)", allocCnts).Find(&results).Error

	return
}

// GetFromFreeCnt 通过free_cnt获取内容
func (obj *_AllVirtualDtlMemoryMgr) GetFromFreeCnt(freeCnt int64) (results []*AllVirtualDtlMemory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDtlMemory{}).Where("`free_cnt` = ?", freeCnt).Find(&results).Error

	return
}

// GetBatchFromFreeCnt 批量查找
func (obj *_AllVirtualDtlMemoryMgr) GetBatchFromFreeCnt(freeCnts []int64) (results []*AllVirtualDtlMemory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDtlMemory{}).Where("`free_cnt` IN (?)", freeCnts).Find(&results).Error

	return
}

// GetFromFreeQueueLen 通过free_queue_len获取内容
func (obj *_AllVirtualDtlMemoryMgr) GetFromFreeQueueLen(freeQueueLen int64) (results []*AllVirtualDtlMemory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDtlMemory{}).Where("`free_queue_len` = ?", freeQueueLen).Find(&results).Error

	return
}

// GetBatchFromFreeQueueLen 批量查找
func (obj *_AllVirtualDtlMemoryMgr) GetBatchFromFreeQueueLen(freeQueueLens []int64) (results []*AllVirtualDtlMemory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDtlMemory{}).Where("`free_queue_len` IN (?)", freeQueueLens).Find(&results).Error

	return
}

// GetFromTotalMemorySize 通过total_memory_size获取内容
func (obj *_AllVirtualDtlMemoryMgr) GetFromTotalMemorySize(totalMemorySize int64) (results []*AllVirtualDtlMemory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDtlMemory{}).Where("`total_memory_size` = ?", totalMemorySize).Find(&results).Error

	return
}

// GetBatchFromTotalMemorySize 批量查找
func (obj *_AllVirtualDtlMemoryMgr) GetBatchFromTotalMemorySize(totalMemorySizes []int64) (results []*AllVirtualDtlMemory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDtlMemory{}).Where("`total_memory_size` IN (?)", totalMemorySizes).Find(&results).Error

	return
}

// GetFromRealAllocCnt 通过real_alloc_cnt获取内容
func (obj *_AllVirtualDtlMemoryMgr) GetFromRealAllocCnt(realAllocCnt int64) (results []*AllVirtualDtlMemory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDtlMemory{}).Where("`real_alloc_cnt` = ?", realAllocCnt).Find(&results).Error

	return
}

// GetBatchFromRealAllocCnt 批量查找
func (obj *_AllVirtualDtlMemoryMgr) GetBatchFromRealAllocCnt(realAllocCnts []int64) (results []*AllVirtualDtlMemory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDtlMemory{}).Where("`real_alloc_cnt` IN (?)", realAllocCnts).Find(&results).Error

	return
}

// GetFromRealFreeCnt 通过real_free_cnt获取内容
func (obj *_AllVirtualDtlMemoryMgr) GetFromRealFreeCnt(realFreeCnt int64) (results []*AllVirtualDtlMemory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDtlMemory{}).Where("`real_free_cnt` = ?", realFreeCnt).Find(&results).Error

	return
}

// GetBatchFromRealFreeCnt 批量查找
func (obj *_AllVirtualDtlMemoryMgr) GetBatchFromRealFreeCnt(realFreeCnts []int64) (results []*AllVirtualDtlMemory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDtlMemory{}).Where("`real_free_cnt` IN (?)", realFreeCnts).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllVirtualDtlMemoryMgr) FetchByPrimaryKey(svrIP string, svrPort int64) (result AllVirtualDtlMemory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDtlMemory{}).Where("`svr_ip` = ? AND `svr_port` = ?", svrIP, svrPort).First(&result).Error

	return
}
