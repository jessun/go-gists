package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _AllVirtualTransMgrStatMgr struct {
	*_BaseMgr
}

// AllVirtualTransMgrStatMgr open func
func AllVirtualTransMgrStatMgr(db *gorm.DB) *_AllVirtualTransMgrStatMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualTransMgrStatMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualTransMgrStatMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_trans_mgr_stat"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualTransMgrStatMgr) GetTableName() string {
	return "__all_virtual_trans_mgr_stat"
}

// Reset 重置gorm会话
func (obj *_AllVirtualTransMgrStatMgr) Reset() *_AllVirtualTransMgrStatMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualTransMgrStatMgr) Get() (result AllVirtualTransMgrStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransMgrStat{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualTransMgrStatMgr) Gets() (results []*AllVirtualTransMgrStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransMgrStat{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualTransMgrStatMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualTransMgrStat{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithSvrIP svr_ip获取
func (obj *_AllVirtualTransMgrStatMgr) WithSvrIP(svrIP string) Option {
	return optionFunc(func(o *options) { o.query["svr_ip"] = svrIP })
}

// WithSvrPort svr_port获取
func (obj *_AllVirtualTransMgrStatMgr) WithSvrPort(svrPort int64) Option {
	return optionFunc(func(o *options) { o.query["svr_port"] = svrPort })
}

// WithTableID table_id获取
func (obj *_AllVirtualTransMgrStatMgr) WithTableID(tableID int64) Option {
	return optionFunc(func(o *options) { o.query["table_id"] = tableID })
}

// WithPartitionIDx partition_idx获取
func (obj *_AllVirtualTransMgrStatMgr) WithPartitionIDx(partitionIDx int64) Option {
	return optionFunc(func(o *options) { o.query["partition_idx"] = partitionIDx })
}

// WithCtxType ctx_type获取
func (obj *_AllVirtualTransMgrStatMgr) WithCtxType(ctxType int64) Option {
	return optionFunc(func(o *options) { o.query["ctx_type"] = ctxType })
}

// WithIsMaster is_master获取
func (obj *_AllVirtualTransMgrStatMgr) WithIsMaster(isMaster int64) Option {
	return optionFunc(func(o *options) { o.query["is_master"] = isMaster })
}

// WithIsFrozen is_frozen获取
func (obj *_AllVirtualTransMgrStatMgr) WithIsFrozen(isFrozen int64) Option {
	return optionFunc(func(o *options) { o.query["is_frozen"] = isFrozen })
}

// WithIsStopped is_stopped获取
func (obj *_AllVirtualTransMgrStatMgr) WithIsStopped(isStopped int64) Option {
	return optionFunc(func(o *options) { o.query["is_stopped"] = isStopped })
}

// WithReadOnlyCount read_only_count获取
func (obj *_AllVirtualTransMgrStatMgr) WithReadOnlyCount(readOnlyCount int64) Option {
	return optionFunc(func(o *options) { o.query["read_only_count"] = readOnlyCount })
}

// WithActiveReadWriteCount active_read_write_count获取
func (obj *_AllVirtualTransMgrStatMgr) WithActiveReadWriteCount(activeReadWriteCount int64) Option {
	return optionFunc(func(o *options) { o.query["active_read_write_count"] = activeReadWriteCount })
}

// WithActiveMemstoreVersion active_memstore_version获取
func (obj *_AllVirtualTransMgrStatMgr) WithActiveMemstoreVersion(activeMemstoreVersion string) Option {
	return optionFunc(func(o *options) { o.query["active_memstore_version"] = activeMemstoreVersion })
}

// WithTotalCtxCount total_ctx_count获取
func (obj *_AllVirtualTransMgrStatMgr) WithTotalCtxCount(totalCtxCount int64) Option {
	return optionFunc(func(o *options) { o.query["total_ctx_count"] = totalCtxCount })
}

// WithWithDepTransCount with_dep_trans_count获取
func (obj *_AllVirtualTransMgrStatMgr) WithWithDepTransCount(withDepTransCount int64) Option {
	return optionFunc(func(o *options) { o.query["with_dep_trans_count"] = withDepTransCount })
}

// WithWithoutDepTransCount without_dep_trans_count获取
func (obj *_AllVirtualTransMgrStatMgr) WithWithoutDepTransCount(withoutDepTransCount int64) Option {
	return optionFunc(func(o *options) { o.query["without_dep_trans_count"] = withoutDepTransCount })
}

// WithEndtransByPrevCount endtrans_by_prev_count获取
func (obj *_AllVirtualTransMgrStatMgr) WithEndtransByPrevCount(endtransByPrevCount int64) Option {
	return optionFunc(func(o *options) { o.query["endtrans_by_prev_count"] = endtransByPrevCount })
}

// WithEndtransByCheckpointCount endtrans_by_checkpoint_count获取
func (obj *_AllVirtualTransMgrStatMgr) WithEndtransByCheckpointCount(endtransByCheckpointCount int64) Option {
	return optionFunc(func(o *options) { o.query["endtrans_by_checkpoint_count"] = endtransByCheckpointCount })
}

// WithEndtransBySelfCount endtrans_by_self_count获取
func (obj *_AllVirtualTransMgrStatMgr) WithEndtransBySelfCount(endtransBySelfCount int64) Option {
	return optionFunc(func(o *options) { o.query["endtrans_by_self_count"] = endtransBySelfCount })
}

// WithMgrAddr mgr_addr获取
func (obj *_AllVirtualTransMgrStatMgr) WithMgrAddr(mgrAddr int64) Option {
	return optionFunc(func(o *options) { o.query["mgr_addr"] = mgrAddr })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualTransMgrStatMgr) GetByOption(opts ...Option) (result AllVirtualTransMgrStat, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransMgrStat{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualTransMgrStatMgr) GetByOptions(opts ...Option) (results []*AllVirtualTransMgrStat, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransMgrStat{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualTransMgrStatMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualTransMgrStat, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualTransMgrStat{}).Where(options.query)
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
func (obj *_AllVirtualTransMgrStatMgr) GetFromSvrIP(svrIP string) (results []*AllVirtualTransMgrStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransMgrStat{}).Where("`svr_ip` = ?", svrIP).Find(&results).Error

	return
}

// GetBatchFromSvrIP 批量查找
func (obj *_AllVirtualTransMgrStatMgr) GetBatchFromSvrIP(svrIPs []string) (results []*AllVirtualTransMgrStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransMgrStat{}).Where("`svr_ip` IN (?)", svrIPs).Find(&results).Error

	return
}

// GetFromSvrPort 通过svr_port获取内容
func (obj *_AllVirtualTransMgrStatMgr) GetFromSvrPort(svrPort int64) (results []*AllVirtualTransMgrStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransMgrStat{}).Where("`svr_port` = ?", svrPort).Find(&results).Error

	return
}

// GetBatchFromSvrPort 批量查找
func (obj *_AllVirtualTransMgrStatMgr) GetBatchFromSvrPort(svrPorts []int64) (results []*AllVirtualTransMgrStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransMgrStat{}).Where("`svr_port` IN (?)", svrPorts).Find(&results).Error

	return
}

// GetFromTableID 通过table_id获取内容
func (obj *_AllVirtualTransMgrStatMgr) GetFromTableID(tableID int64) (results []*AllVirtualTransMgrStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransMgrStat{}).Where("`table_id` = ?", tableID).Find(&results).Error

	return
}

// GetBatchFromTableID 批量查找
func (obj *_AllVirtualTransMgrStatMgr) GetBatchFromTableID(tableIDs []int64) (results []*AllVirtualTransMgrStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransMgrStat{}).Where("`table_id` IN (?)", tableIDs).Find(&results).Error

	return
}

// GetFromPartitionIDx 通过partition_idx获取内容
func (obj *_AllVirtualTransMgrStatMgr) GetFromPartitionIDx(partitionIDx int64) (results []*AllVirtualTransMgrStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransMgrStat{}).Where("`partition_idx` = ?", partitionIDx).Find(&results).Error

	return
}

// GetBatchFromPartitionIDx 批量查找
func (obj *_AllVirtualTransMgrStatMgr) GetBatchFromPartitionIDx(partitionIDxs []int64) (results []*AllVirtualTransMgrStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransMgrStat{}).Where("`partition_idx` IN (?)", partitionIDxs).Find(&results).Error

	return
}

// GetFromCtxType 通过ctx_type获取内容
func (obj *_AllVirtualTransMgrStatMgr) GetFromCtxType(ctxType int64) (results []*AllVirtualTransMgrStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransMgrStat{}).Where("`ctx_type` = ?", ctxType).Find(&results).Error

	return
}

// GetBatchFromCtxType 批量查找
func (obj *_AllVirtualTransMgrStatMgr) GetBatchFromCtxType(ctxTypes []int64) (results []*AllVirtualTransMgrStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransMgrStat{}).Where("`ctx_type` IN (?)", ctxTypes).Find(&results).Error

	return
}

// GetFromIsMaster 通过is_master获取内容
func (obj *_AllVirtualTransMgrStatMgr) GetFromIsMaster(isMaster int64) (results []*AllVirtualTransMgrStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransMgrStat{}).Where("`is_master` = ?", isMaster).Find(&results).Error

	return
}

// GetBatchFromIsMaster 批量查找
func (obj *_AllVirtualTransMgrStatMgr) GetBatchFromIsMaster(isMasters []int64) (results []*AllVirtualTransMgrStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransMgrStat{}).Where("`is_master` IN (?)", isMasters).Find(&results).Error

	return
}

// GetFromIsFrozen 通过is_frozen获取内容
func (obj *_AllVirtualTransMgrStatMgr) GetFromIsFrozen(isFrozen int64) (results []*AllVirtualTransMgrStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransMgrStat{}).Where("`is_frozen` = ?", isFrozen).Find(&results).Error

	return
}

// GetBatchFromIsFrozen 批量查找
func (obj *_AllVirtualTransMgrStatMgr) GetBatchFromIsFrozen(isFrozens []int64) (results []*AllVirtualTransMgrStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransMgrStat{}).Where("`is_frozen` IN (?)", isFrozens).Find(&results).Error

	return
}

// GetFromIsStopped 通过is_stopped获取内容
func (obj *_AllVirtualTransMgrStatMgr) GetFromIsStopped(isStopped int64) (results []*AllVirtualTransMgrStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransMgrStat{}).Where("`is_stopped` = ?", isStopped).Find(&results).Error

	return
}

// GetBatchFromIsStopped 批量查找
func (obj *_AllVirtualTransMgrStatMgr) GetBatchFromIsStopped(isStoppeds []int64) (results []*AllVirtualTransMgrStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransMgrStat{}).Where("`is_stopped` IN (?)", isStoppeds).Find(&results).Error

	return
}

// GetFromReadOnlyCount 通过read_only_count获取内容
func (obj *_AllVirtualTransMgrStatMgr) GetFromReadOnlyCount(readOnlyCount int64) (results []*AllVirtualTransMgrStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransMgrStat{}).Where("`read_only_count` = ?", readOnlyCount).Find(&results).Error

	return
}

// GetBatchFromReadOnlyCount 批量查找
func (obj *_AllVirtualTransMgrStatMgr) GetBatchFromReadOnlyCount(readOnlyCounts []int64) (results []*AllVirtualTransMgrStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransMgrStat{}).Where("`read_only_count` IN (?)", readOnlyCounts).Find(&results).Error

	return
}

// GetFromActiveReadWriteCount 通过active_read_write_count获取内容
func (obj *_AllVirtualTransMgrStatMgr) GetFromActiveReadWriteCount(activeReadWriteCount int64) (results []*AllVirtualTransMgrStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransMgrStat{}).Where("`active_read_write_count` = ?", activeReadWriteCount).Find(&results).Error

	return
}

// GetBatchFromActiveReadWriteCount 批量查找
func (obj *_AllVirtualTransMgrStatMgr) GetBatchFromActiveReadWriteCount(activeReadWriteCounts []int64) (results []*AllVirtualTransMgrStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransMgrStat{}).Where("`active_read_write_count` IN (?)", activeReadWriteCounts).Find(&results).Error

	return
}

// GetFromActiveMemstoreVersion 通过active_memstore_version获取内容
func (obj *_AllVirtualTransMgrStatMgr) GetFromActiveMemstoreVersion(activeMemstoreVersion string) (results []*AllVirtualTransMgrStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransMgrStat{}).Where("`active_memstore_version` = ?", activeMemstoreVersion).Find(&results).Error

	return
}

// GetBatchFromActiveMemstoreVersion 批量查找
func (obj *_AllVirtualTransMgrStatMgr) GetBatchFromActiveMemstoreVersion(activeMemstoreVersions []string) (results []*AllVirtualTransMgrStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransMgrStat{}).Where("`active_memstore_version` IN (?)", activeMemstoreVersions).Find(&results).Error

	return
}

// GetFromTotalCtxCount 通过total_ctx_count获取内容
func (obj *_AllVirtualTransMgrStatMgr) GetFromTotalCtxCount(totalCtxCount int64) (results []*AllVirtualTransMgrStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransMgrStat{}).Where("`total_ctx_count` = ?", totalCtxCount).Find(&results).Error

	return
}

// GetBatchFromTotalCtxCount 批量查找
func (obj *_AllVirtualTransMgrStatMgr) GetBatchFromTotalCtxCount(totalCtxCounts []int64) (results []*AllVirtualTransMgrStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransMgrStat{}).Where("`total_ctx_count` IN (?)", totalCtxCounts).Find(&results).Error

	return
}

// GetFromWithDepTransCount 通过with_dep_trans_count获取内容
func (obj *_AllVirtualTransMgrStatMgr) GetFromWithDepTransCount(withDepTransCount int64) (results []*AllVirtualTransMgrStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransMgrStat{}).Where("`with_dep_trans_count` = ?", withDepTransCount).Find(&results).Error

	return
}

// GetBatchFromWithDepTransCount 批量查找
func (obj *_AllVirtualTransMgrStatMgr) GetBatchFromWithDepTransCount(withDepTransCounts []int64) (results []*AllVirtualTransMgrStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransMgrStat{}).Where("`with_dep_trans_count` IN (?)", withDepTransCounts).Find(&results).Error

	return
}

// GetFromWithoutDepTransCount 通过without_dep_trans_count获取内容
func (obj *_AllVirtualTransMgrStatMgr) GetFromWithoutDepTransCount(withoutDepTransCount int64) (results []*AllVirtualTransMgrStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransMgrStat{}).Where("`without_dep_trans_count` = ?", withoutDepTransCount).Find(&results).Error

	return
}

// GetBatchFromWithoutDepTransCount 批量查找
func (obj *_AllVirtualTransMgrStatMgr) GetBatchFromWithoutDepTransCount(withoutDepTransCounts []int64) (results []*AllVirtualTransMgrStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransMgrStat{}).Where("`without_dep_trans_count` IN (?)", withoutDepTransCounts).Find(&results).Error

	return
}

// GetFromEndtransByPrevCount 通过endtrans_by_prev_count获取内容
func (obj *_AllVirtualTransMgrStatMgr) GetFromEndtransByPrevCount(endtransByPrevCount int64) (results []*AllVirtualTransMgrStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransMgrStat{}).Where("`endtrans_by_prev_count` = ?", endtransByPrevCount).Find(&results).Error

	return
}

// GetBatchFromEndtransByPrevCount 批量查找
func (obj *_AllVirtualTransMgrStatMgr) GetBatchFromEndtransByPrevCount(endtransByPrevCounts []int64) (results []*AllVirtualTransMgrStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransMgrStat{}).Where("`endtrans_by_prev_count` IN (?)", endtransByPrevCounts).Find(&results).Error

	return
}

// GetFromEndtransByCheckpointCount 通过endtrans_by_checkpoint_count获取内容
func (obj *_AllVirtualTransMgrStatMgr) GetFromEndtransByCheckpointCount(endtransByCheckpointCount int64) (results []*AllVirtualTransMgrStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransMgrStat{}).Where("`endtrans_by_checkpoint_count` = ?", endtransByCheckpointCount).Find(&results).Error

	return
}

// GetBatchFromEndtransByCheckpointCount 批量查找
func (obj *_AllVirtualTransMgrStatMgr) GetBatchFromEndtransByCheckpointCount(endtransByCheckpointCounts []int64) (results []*AllVirtualTransMgrStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransMgrStat{}).Where("`endtrans_by_checkpoint_count` IN (?)", endtransByCheckpointCounts).Find(&results).Error

	return
}

// GetFromEndtransBySelfCount 通过endtrans_by_self_count获取内容
func (obj *_AllVirtualTransMgrStatMgr) GetFromEndtransBySelfCount(endtransBySelfCount int64) (results []*AllVirtualTransMgrStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransMgrStat{}).Where("`endtrans_by_self_count` = ?", endtransBySelfCount).Find(&results).Error

	return
}

// GetBatchFromEndtransBySelfCount 批量查找
func (obj *_AllVirtualTransMgrStatMgr) GetBatchFromEndtransBySelfCount(endtransBySelfCounts []int64) (results []*AllVirtualTransMgrStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransMgrStat{}).Where("`endtrans_by_self_count` IN (?)", endtransBySelfCounts).Find(&results).Error

	return
}

// GetFromMgrAddr 通过mgr_addr获取内容
func (obj *_AllVirtualTransMgrStatMgr) GetFromMgrAddr(mgrAddr int64) (results []*AllVirtualTransMgrStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransMgrStat{}).Where("`mgr_addr` = ?", mgrAddr).Find(&results).Error

	return
}

// GetBatchFromMgrAddr 批量查找
func (obj *_AllVirtualTransMgrStatMgr) GetBatchFromMgrAddr(mgrAddrs []int64) (results []*AllVirtualTransMgrStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransMgrStat{}).Where("`mgr_addr` IN (?)", mgrAddrs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////
