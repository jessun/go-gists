package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _AllVirtualMacroBlockMarkerStatusMgr struct {
	*_BaseMgr
}

// AllVirtualMacroBlockMarkerStatusMgr open func
func AllVirtualMacroBlockMarkerStatusMgr(db *gorm.DB) *_AllVirtualMacroBlockMarkerStatusMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualMacroBlockMarkerStatusMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualMacroBlockMarkerStatusMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_macro_block_marker_status"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualMacroBlockMarkerStatusMgr) GetTableName() string {
	return "__all_virtual_macro_block_marker_status"
}

// Reset 重置gorm会话
func (obj *_AllVirtualMacroBlockMarkerStatusMgr) Reset() *_AllVirtualMacroBlockMarkerStatusMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualMacroBlockMarkerStatusMgr) Get() (result AllVirtualMacroBlockMarkerStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMacroBlockMarkerStatus{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualMacroBlockMarkerStatusMgr) Gets() (results []*AllVirtualMacroBlockMarkerStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMacroBlockMarkerStatus{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualMacroBlockMarkerStatusMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualMacroBlockMarkerStatus{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithSvrIP svr_ip获取
func (obj *_AllVirtualMacroBlockMarkerStatusMgr) WithSvrIP(svrIP string) Option {
	return optionFunc(func(o *options) { o.query["svr_ip"] = svrIP })
}

// WithSvrPort svr_port获取
func (obj *_AllVirtualMacroBlockMarkerStatusMgr) WithSvrPort(svrPort int64) Option {
	return optionFunc(func(o *options) { o.query["svr_port"] = svrPort })
}

// WithTotalCount total_count获取
func (obj *_AllVirtualMacroBlockMarkerStatusMgr) WithTotalCount(totalCount int64) Option {
	return optionFunc(func(o *options) { o.query["total_count"] = totalCount })
}

// WithReservedCount reserved_count获取
func (obj *_AllVirtualMacroBlockMarkerStatusMgr) WithReservedCount(reservedCount int64) Option {
	return optionFunc(func(o *options) { o.query["reserved_count"] = reservedCount })
}

// WithMacroMetaCount macro_meta_count获取
func (obj *_AllVirtualMacroBlockMarkerStatusMgr) WithMacroMetaCount(macroMetaCount int64) Option {
	return optionFunc(func(o *options) { o.query["macro_meta_count"] = macroMetaCount })
}

// WithPartitionMetaCount partition_meta_count获取
func (obj *_AllVirtualMacroBlockMarkerStatusMgr) WithPartitionMetaCount(partitionMetaCount int64) Option {
	return optionFunc(func(o *options) { o.query["partition_meta_count"] = partitionMetaCount })
}

// WithDataCount data_count获取
func (obj *_AllVirtualMacroBlockMarkerStatusMgr) WithDataCount(dataCount int64) Option {
	return optionFunc(func(o *options) { o.query["data_count"] = dataCount })
}

// WithSecondIndexCount second_index_count获取
func (obj *_AllVirtualMacroBlockMarkerStatusMgr) WithSecondIndexCount(secondIndexCount int64) Option {
	return optionFunc(func(o *options) { o.query["second_index_count"] = secondIndexCount })
}

// WithLobDataCount lob_data_count获取
func (obj *_AllVirtualMacroBlockMarkerStatusMgr) WithLobDataCount(lobDataCount int64) Option {
	return optionFunc(func(o *options) { o.query["lob_data_count"] = lobDataCount })
}

// WithLobSecondIndexCount lob_second_index_count获取
func (obj *_AllVirtualMacroBlockMarkerStatusMgr) WithLobSecondIndexCount(lobSecondIndexCount int64) Option {
	return optionFunc(func(o *options) { o.query["lob_second_index_count"] = lobSecondIndexCount })
}

// WithBloomfilterCount bloomfilter_count获取
func (obj *_AllVirtualMacroBlockMarkerStatusMgr) WithBloomfilterCount(bloomfilterCount int64) Option {
	return optionFunc(func(o *options) { o.query["bloomfilter_count"] = bloomfilterCount })
}

// WithHoldCount hold_count获取
func (obj *_AllVirtualMacroBlockMarkerStatusMgr) WithHoldCount(holdCount int64) Option {
	return optionFunc(func(o *options) { o.query["hold_count"] = holdCount })
}

// WithPendingFreeCount pending_free_count获取
func (obj *_AllVirtualMacroBlockMarkerStatusMgr) WithPendingFreeCount(pendingFreeCount int64) Option {
	return optionFunc(func(o *options) { o.query["pending_free_count"] = pendingFreeCount })
}

// WithFreeCount free_count获取
func (obj *_AllVirtualMacroBlockMarkerStatusMgr) WithFreeCount(freeCount int64) Option {
	return optionFunc(func(o *options) { o.query["free_count"] = freeCount })
}

// WithMarkCostTime mark_cost_time获取
func (obj *_AllVirtualMacroBlockMarkerStatusMgr) WithMarkCostTime(markCostTime int64) Option {
	return optionFunc(func(o *options) { o.query["mark_cost_time"] = markCostTime })
}

// WithSweepCostTime sweep_cost_time获取
func (obj *_AllVirtualMacroBlockMarkerStatusMgr) WithSweepCostTime(sweepCostTime int64) Option {
	return optionFunc(func(o *options) { o.query["sweep_cost_time"] = sweepCostTime })
}

// WithComment comment获取
func (obj *_AllVirtualMacroBlockMarkerStatusMgr) WithComment(comment string) Option {
	return optionFunc(func(o *options) { o.query["comment"] = comment })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualMacroBlockMarkerStatusMgr) GetByOption(opts ...Option) (result AllVirtualMacroBlockMarkerStatus, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMacroBlockMarkerStatus{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualMacroBlockMarkerStatusMgr) GetByOptions(opts ...Option) (results []*AllVirtualMacroBlockMarkerStatus, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMacroBlockMarkerStatus{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualMacroBlockMarkerStatusMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualMacroBlockMarkerStatus, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualMacroBlockMarkerStatus{}).Where(options.query)
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
func (obj *_AllVirtualMacroBlockMarkerStatusMgr) GetFromSvrIP(svrIP string) (results []*AllVirtualMacroBlockMarkerStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMacroBlockMarkerStatus{}).Where("`svr_ip` = ?", svrIP).Find(&results).Error

	return
}

// GetBatchFromSvrIP 批量查找
func (obj *_AllVirtualMacroBlockMarkerStatusMgr) GetBatchFromSvrIP(svrIPs []string) (results []*AllVirtualMacroBlockMarkerStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMacroBlockMarkerStatus{}).Where("`svr_ip` IN (?)", svrIPs).Find(&results).Error

	return
}

// GetFromSvrPort 通过svr_port获取内容
func (obj *_AllVirtualMacroBlockMarkerStatusMgr) GetFromSvrPort(svrPort int64) (results []*AllVirtualMacroBlockMarkerStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMacroBlockMarkerStatus{}).Where("`svr_port` = ?", svrPort).Find(&results).Error

	return
}

// GetBatchFromSvrPort 批量查找
func (obj *_AllVirtualMacroBlockMarkerStatusMgr) GetBatchFromSvrPort(svrPorts []int64) (results []*AllVirtualMacroBlockMarkerStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMacroBlockMarkerStatus{}).Where("`svr_port` IN (?)", svrPorts).Find(&results).Error

	return
}

// GetFromTotalCount 通过total_count获取内容
func (obj *_AllVirtualMacroBlockMarkerStatusMgr) GetFromTotalCount(totalCount int64) (results []*AllVirtualMacroBlockMarkerStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMacroBlockMarkerStatus{}).Where("`total_count` = ?", totalCount).Find(&results).Error

	return
}

// GetBatchFromTotalCount 批量查找
func (obj *_AllVirtualMacroBlockMarkerStatusMgr) GetBatchFromTotalCount(totalCounts []int64) (results []*AllVirtualMacroBlockMarkerStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMacroBlockMarkerStatus{}).Where("`total_count` IN (?)", totalCounts).Find(&results).Error

	return
}

// GetFromReservedCount 通过reserved_count获取内容
func (obj *_AllVirtualMacroBlockMarkerStatusMgr) GetFromReservedCount(reservedCount int64) (results []*AllVirtualMacroBlockMarkerStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMacroBlockMarkerStatus{}).Where("`reserved_count` = ?", reservedCount).Find(&results).Error

	return
}

// GetBatchFromReservedCount 批量查找
func (obj *_AllVirtualMacroBlockMarkerStatusMgr) GetBatchFromReservedCount(reservedCounts []int64) (results []*AllVirtualMacroBlockMarkerStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMacroBlockMarkerStatus{}).Where("`reserved_count` IN (?)", reservedCounts).Find(&results).Error

	return
}

// GetFromMacroMetaCount 通过macro_meta_count获取内容
func (obj *_AllVirtualMacroBlockMarkerStatusMgr) GetFromMacroMetaCount(macroMetaCount int64) (results []*AllVirtualMacroBlockMarkerStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMacroBlockMarkerStatus{}).Where("`macro_meta_count` = ?", macroMetaCount).Find(&results).Error

	return
}

// GetBatchFromMacroMetaCount 批量查找
func (obj *_AllVirtualMacroBlockMarkerStatusMgr) GetBatchFromMacroMetaCount(macroMetaCounts []int64) (results []*AllVirtualMacroBlockMarkerStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMacroBlockMarkerStatus{}).Where("`macro_meta_count` IN (?)", macroMetaCounts).Find(&results).Error

	return
}

// GetFromPartitionMetaCount 通过partition_meta_count获取内容
func (obj *_AllVirtualMacroBlockMarkerStatusMgr) GetFromPartitionMetaCount(partitionMetaCount int64) (results []*AllVirtualMacroBlockMarkerStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMacroBlockMarkerStatus{}).Where("`partition_meta_count` = ?", partitionMetaCount).Find(&results).Error

	return
}

// GetBatchFromPartitionMetaCount 批量查找
func (obj *_AllVirtualMacroBlockMarkerStatusMgr) GetBatchFromPartitionMetaCount(partitionMetaCounts []int64) (results []*AllVirtualMacroBlockMarkerStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMacroBlockMarkerStatus{}).Where("`partition_meta_count` IN (?)", partitionMetaCounts).Find(&results).Error

	return
}

// GetFromDataCount 通过data_count获取内容
func (obj *_AllVirtualMacroBlockMarkerStatusMgr) GetFromDataCount(dataCount int64) (results []*AllVirtualMacroBlockMarkerStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMacroBlockMarkerStatus{}).Where("`data_count` = ?", dataCount).Find(&results).Error

	return
}

// GetBatchFromDataCount 批量查找
func (obj *_AllVirtualMacroBlockMarkerStatusMgr) GetBatchFromDataCount(dataCounts []int64) (results []*AllVirtualMacroBlockMarkerStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMacroBlockMarkerStatus{}).Where("`data_count` IN (?)", dataCounts).Find(&results).Error

	return
}

// GetFromSecondIndexCount 通过second_index_count获取内容
func (obj *_AllVirtualMacroBlockMarkerStatusMgr) GetFromSecondIndexCount(secondIndexCount int64) (results []*AllVirtualMacroBlockMarkerStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMacroBlockMarkerStatus{}).Where("`second_index_count` = ?", secondIndexCount).Find(&results).Error

	return
}

// GetBatchFromSecondIndexCount 批量查找
func (obj *_AllVirtualMacroBlockMarkerStatusMgr) GetBatchFromSecondIndexCount(secondIndexCounts []int64) (results []*AllVirtualMacroBlockMarkerStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMacroBlockMarkerStatus{}).Where("`second_index_count` IN (?)", secondIndexCounts).Find(&results).Error

	return
}

// GetFromLobDataCount 通过lob_data_count获取内容
func (obj *_AllVirtualMacroBlockMarkerStatusMgr) GetFromLobDataCount(lobDataCount int64) (results []*AllVirtualMacroBlockMarkerStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMacroBlockMarkerStatus{}).Where("`lob_data_count` = ?", lobDataCount).Find(&results).Error

	return
}

// GetBatchFromLobDataCount 批量查找
func (obj *_AllVirtualMacroBlockMarkerStatusMgr) GetBatchFromLobDataCount(lobDataCounts []int64) (results []*AllVirtualMacroBlockMarkerStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMacroBlockMarkerStatus{}).Where("`lob_data_count` IN (?)", lobDataCounts).Find(&results).Error

	return
}

// GetFromLobSecondIndexCount 通过lob_second_index_count获取内容
func (obj *_AllVirtualMacroBlockMarkerStatusMgr) GetFromLobSecondIndexCount(lobSecondIndexCount int64) (results []*AllVirtualMacroBlockMarkerStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMacroBlockMarkerStatus{}).Where("`lob_second_index_count` = ?", lobSecondIndexCount).Find(&results).Error

	return
}

// GetBatchFromLobSecondIndexCount 批量查找
func (obj *_AllVirtualMacroBlockMarkerStatusMgr) GetBatchFromLobSecondIndexCount(lobSecondIndexCounts []int64) (results []*AllVirtualMacroBlockMarkerStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMacroBlockMarkerStatus{}).Where("`lob_second_index_count` IN (?)", lobSecondIndexCounts).Find(&results).Error

	return
}

// GetFromBloomfilterCount 通过bloomfilter_count获取内容
func (obj *_AllVirtualMacroBlockMarkerStatusMgr) GetFromBloomfilterCount(bloomfilterCount int64) (results []*AllVirtualMacroBlockMarkerStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMacroBlockMarkerStatus{}).Where("`bloomfilter_count` = ?", bloomfilterCount).Find(&results).Error

	return
}

// GetBatchFromBloomfilterCount 批量查找
func (obj *_AllVirtualMacroBlockMarkerStatusMgr) GetBatchFromBloomfilterCount(bloomfilterCounts []int64) (results []*AllVirtualMacroBlockMarkerStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMacroBlockMarkerStatus{}).Where("`bloomfilter_count` IN (?)", bloomfilterCounts).Find(&results).Error

	return
}

// GetFromHoldCount 通过hold_count获取内容
func (obj *_AllVirtualMacroBlockMarkerStatusMgr) GetFromHoldCount(holdCount int64) (results []*AllVirtualMacroBlockMarkerStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMacroBlockMarkerStatus{}).Where("`hold_count` = ?", holdCount).Find(&results).Error

	return
}

// GetBatchFromHoldCount 批量查找
func (obj *_AllVirtualMacroBlockMarkerStatusMgr) GetBatchFromHoldCount(holdCounts []int64) (results []*AllVirtualMacroBlockMarkerStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMacroBlockMarkerStatus{}).Where("`hold_count` IN (?)", holdCounts).Find(&results).Error

	return
}

// GetFromPendingFreeCount 通过pending_free_count获取内容
func (obj *_AllVirtualMacroBlockMarkerStatusMgr) GetFromPendingFreeCount(pendingFreeCount int64) (results []*AllVirtualMacroBlockMarkerStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMacroBlockMarkerStatus{}).Where("`pending_free_count` = ?", pendingFreeCount).Find(&results).Error

	return
}

// GetBatchFromPendingFreeCount 批量查找
func (obj *_AllVirtualMacroBlockMarkerStatusMgr) GetBatchFromPendingFreeCount(pendingFreeCounts []int64) (results []*AllVirtualMacroBlockMarkerStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMacroBlockMarkerStatus{}).Where("`pending_free_count` IN (?)", pendingFreeCounts).Find(&results).Error

	return
}

// GetFromFreeCount 通过free_count获取内容
func (obj *_AllVirtualMacroBlockMarkerStatusMgr) GetFromFreeCount(freeCount int64) (results []*AllVirtualMacroBlockMarkerStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMacroBlockMarkerStatus{}).Where("`free_count` = ?", freeCount).Find(&results).Error

	return
}

// GetBatchFromFreeCount 批量查找
func (obj *_AllVirtualMacroBlockMarkerStatusMgr) GetBatchFromFreeCount(freeCounts []int64) (results []*AllVirtualMacroBlockMarkerStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMacroBlockMarkerStatus{}).Where("`free_count` IN (?)", freeCounts).Find(&results).Error

	return
}

// GetFromMarkCostTime 通过mark_cost_time获取内容
func (obj *_AllVirtualMacroBlockMarkerStatusMgr) GetFromMarkCostTime(markCostTime int64) (results []*AllVirtualMacroBlockMarkerStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMacroBlockMarkerStatus{}).Where("`mark_cost_time` = ?", markCostTime).Find(&results).Error

	return
}

// GetBatchFromMarkCostTime 批量查找
func (obj *_AllVirtualMacroBlockMarkerStatusMgr) GetBatchFromMarkCostTime(markCostTimes []int64) (results []*AllVirtualMacroBlockMarkerStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMacroBlockMarkerStatus{}).Where("`mark_cost_time` IN (?)", markCostTimes).Find(&results).Error

	return
}

// GetFromSweepCostTime 通过sweep_cost_time获取内容
func (obj *_AllVirtualMacroBlockMarkerStatusMgr) GetFromSweepCostTime(sweepCostTime int64) (results []*AllVirtualMacroBlockMarkerStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMacroBlockMarkerStatus{}).Where("`sweep_cost_time` = ?", sweepCostTime).Find(&results).Error

	return
}

// GetBatchFromSweepCostTime 批量查找
func (obj *_AllVirtualMacroBlockMarkerStatusMgr) GetBatchFromSweepCostTime(sweepCostTimes []int64) (results []*AllVirtualMacroBlockMarkerStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMacroBlockMarkerStatus{}).Where("`sweep_cost_time` IN (?)", sweepCostTimes).Find(&results).Error

	return
}

// GetFromComment 通过comment获取内容
func (obj *_AllVirtualMacroBlockMarkerStatusMgr) GetFromComment(comment string) (results []*AllVirtualMacroBlockMarkerStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMacroBlockMarkerStatus{}).Where("`comment` = ?", comment).Find(&results).Error

	return
}

// GetBatchFromComment 批量查找
func (obj *_AllVirtualMacroBlockMarkerStatusMgr) GetBatchFromComment(comments []string) (results []*AllVirtualMacroBlockMarkerStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMacroBlockMarkerStatus{}).Where("`comment` IN (?)", comments).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllVirtualMacroBlockMarkerStatusMgr) FetchByPrimaryKey(svrIP string, svrPort int64) (result AllVirtualMacroBlockMarkerStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMacroBlockMarkerStatus{}).Where("`svr_ip` = ? AND `svr_port` = ?", svrIP, svrPort).First(&result).Error

	return
}
