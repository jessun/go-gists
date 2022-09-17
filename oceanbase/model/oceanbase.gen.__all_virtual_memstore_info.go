package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _AllVirtualMemstoreInfoMgr struct {
	*_BaseMgr
}

// AllVirtualMemstoreInfoMgr open func
func AllVirtualMemstoreInfoMgr(db *gorm.DB) *_AllVirtualMemstoreInfoMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualMemstoreInfoMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualMemstoreInfoMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_memstore_info"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualMemstoreInfoMgr) GetTableName() string {
	return "__all_virtual_memstore_info"
}

// Reset 重置gorm会话
func (obj *_AllVirtualMemstoreInfoMgr) Reset() *_AllVirtualMemstoreInfoMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualMemstoreInfoMgr) Get() (result AllVirtualMemstoreInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMemstoreInfo{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualMemstoreInfoMgr) Gets() (results []*AllVirtualMemstoreInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMemstoreInfo{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualMemstoreInfoMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualMemstoreInfo{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTenantID tenant_id获取
func (obj *_AllVirtualMemstoreInfoMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithSvrIP svr_ip获取
func (obj *_AllVirtualMemstoreInfoMgr) WithSvrIP(svrIP string) Option {
	return optionFunc(func(o *options) { o.query["svr_ip"] = svrIP })
}

// WithSvrPort svr_port获取
func (obj *_AllVirtualMemstoreInfoMgr) WithSvrPort(svrPort int64) Option {
	return optionFunc(func(o *options) { o.query["svr_port"] = svrPort })
}

// WithTableID table_id获取
func (obj *_AllVirtualMemstoreInfoMgr) WithTableID(tableID int64) Option {
	return optionFunc(func(o *options) { o.query["table_id"] = tableID })
}

// WithPartitionIDx partition_idx获取
func (obj *_AllVirtualMemstoreInfoMgr) WithPartitionIDx(partitionIDx int64) Option {
	return optionFunc(func(o *options) { o.query["partition_idx"] = partitionIDx })
}

// WithPartitionCnt partition_cnt获取
func (obj *_AllVirtualMemstoreInfoMgr) WithPartitionCnt(partitionCnt int64) Option {
	return optionFunc(func(o *options) { o.query["partition_cnt"] = partitionCnt })
}

// WithVersion version获取
func (obj *_AllVirtualMemstoreInfoMgr) WithVersion(version string) Option {
	return optionFunc(func(o *options) { o.query["version"] = version })
}

// WithBaseVersion base_version获取
func (obj *_AllVirtualMemstoreInfoMgr) WithBaseVersion(baseVersion int64) Option {
	return optionFunc(func(o *options) { o.query["base_version"] = baseVersion })
}

// WithMultiVersionStart multi_version_start获取
func (obj *_AllVirtualMemstoreInfoMgr) WithMultiVersionStart(multiVersionStart int64) Option {
	return optionFunc(func(o *options) { o.query["multi_version_start"] = multiVersionStart })
}

// WithSnapshotVersion snapshot_version获取
func (obj *_AllVirtualMemstoreInfoMgr) WithSnapshotVersion(snapshotVersion int64) Option {
	return optionFunc(func(o *options) { o.query["snapshot_version"] = snapshotVersion })
}

// WithIsActive is_active获取
func (obj *_AllVirtualMemstoreInfoMgr) WithIsActive(isActive int64) Option {
	return optionFunc(func(o *options) { o.query["is_active"] = isActive })
}

// WithMemUsed mem_used获取
func (obj *_AllVirtualMemstoreInfoMgr) WithMemUsed(memUsed int64) Option {
	return optionFunc(func(o *options) { o.query["mem_used"] = memUsed })
}

// WithHashItemCount hash_item_count获取
func (obj *_AllVirtualMemstoreInfoMgr) WithHashItemCount(hashItemCount int64) Option {
	return optionFunc(func(o *options) { o.query["hash_item_count"] = hashItemCount })
}

// WithHashMemUsed hash_mem_used获取
func (obj *_AllVirtualMemstoreInfoMgr) WithHashMemUsed(hashMemUsed int64) Option {
	return optionFunc(func(o *options) { o.query["hash_mem_used"] = hashMemUsed })
}

// WithBtreeItemCount btree_item_count获取
func (obj *_AllVirtualMemstoreInfoMgr) WithBtreeItemCount(btreeItemCount int64) Option {
	return optionFunc(func(o *options) { o.query["btree_item_count"] = btreeItemCount })
}

// WithBtreeMemUsed btree_mem_used获取
func (obj *_AllVirtualMemstoreInfoMgr) WithBtreeMemUsed(btreeMemUsed int64) Option {
	return optionFunc(func(o *options) { o.query["btree_mem_used"] = btreeMemUsed })
}

// WithInsertRowCount insert_row_count获取
func (obj *_AllVirtualMemstoreInfoMgr) WithInsertRowCount(insertRowCount int64) Option {
	return optionFunc(func(o *options) { o.query["insert_row_count"] = insertRowCount })
}

// WithUpdateRowCount update_row_count获取
func (obj *_AllVirtualMemstoreInfoMgr) WithUpdateRowCount(updateRowCount int64) Option {
	return optionFunc(func(o *options) { o.query["update_row_count"] = updateRowCount })
}

// WithDeleteRowCount delete_row_count获取
func (obj *_AllVirtualMemstoreInfoMgr) WithDeleteRowCount(deleteRowCount int64) Option {
	return optionFunc(func(o *options) { o.query["delete_row_count"] = deleteRowCount })
}

// WithPurgeRowCount purge_row_count获取
func (obj *_AllVirtualMemstoreInfoMgr) WithPurgeRowCount(purgeRowCount int64) Option {
	return optionFunc(func(o *options) { o.query["purge_row_count"] = purgeRowCount })
}

// WithPurgeQueueCount purge_queue_count获取
func (obj *_AllVirtualMemstoreInfoMgr) WithPurgeQueueCount(purgeQueueCount int64) Option {
	return optionFunc(func(o *options) { o.query["purge_queue_count"] = purgeQueueCount })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualMemstoreInfoMgr) GetByOption(opts ...Option) (result AllVirtualMemstoreInfo, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMemstoreInfo{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualMemstoreInfoMgr) GetByOptions(opts ...Option) (results []*AllVirtualMemstoreInfo, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMemstoreInfo{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualMemstoreInfoMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualMemstoreInfo, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualMemstoreInfo{}).Where(options.query)
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

// GetFromTenantID 通过tenant_id获取内容
func (obj *_AllVirtualMemstoreInfoMgr) GetFromTenantID(tenantID int64) (results []*AllVirtualMemstoreInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMemstoreInfo{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllVirtualMemstoreInfoMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllVirtualMemstoreInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMemstoreInfo{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromSvrIP 通过svr_ip获取内容
func (obj *_AllVirtualMemstoreInfoMgr) GetFromSvrIP(svrIP string) (results []*AllVirtualMemstoreInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMemstoreInfo{}).Where("`svr_ip` = ?", svrIP).Find(&results).Error

	return
}

// GetBatchFromSvrIP 批量查找
func (obj *_AllVirtualMemstoreInfoMgr) GetBatchFromSvrIP(svrIPs []string) (results []*AllVirtualMemstoreInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMemstoreInfo{}).Where("`svr_ip` IN (?)", svrIPs).Find(&results).Error

	return
}

// GetFromSvrPort 通过svr_port获取内容
func (obj *_AllVirtualMemstoreInfoMgr) GetFromSvrPort(svrPort int64) (results []*AllVirtualMemstoreInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMemstoreInfo{}).Where("`svr_port` = ?", svrPort).Find(&results).Error

	return
}

// GetBatchFromSvrPort 批量查找
func (obj *_AllVirtualMemstoreInfoMgr) GetBatchFromSvrPort(svrPorts []int64) (results []*AllVirtualMemstoreInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMemstoreInfo{}).Where("`svr_port` IN (?)", svrPorts).Find(&results).Error

	return
}

// GetFromTableID 通过table_id获取内容
func (obj *_AllVirtualMemstoreInfoMgr) GetFromTableID(tableID int64) (results []*AllVirtualMemstoreInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMemstoreInfo{}).Where("`table_id` = ?", tableID).Find(&results).Error

	return
}

// GetBatchFromTableID 批量查找
func (obj *_AllVirtualMemstoreInfoMgr) GetBatchFromTableID(tableIDs []int64) (results []*AllVirtualMemstoreInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMemstoreInfo{}).Where("`table_id` IN (?)", tableIDs).Find(&results).Error

	return
}

// GetFromPartitionIDx 通过partition_idx获取内容
func (obj *_AllVirtualMemstoreInfoMgr) GetFromPartitionIDx(partitionIDx int64) (results []*AllVirtualMemstoreInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMemstoreInfo{}).Where("`partition_idx` = ?", partitionIDx).Find(&results).Error

	return
}

// GetBatchFromPartitionIDx 批量查找
func (obj *_AllVirtualMemstoreInfoMgr) GetBatchFromPartitionIDx(partitionIDxs []int64) (results []*AllVirtualMemstoreInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMemstoreInfo{}).Where("`partition_idx` IN (?)", partitionIDxs).Find(&results).Error

	return
}

// GetFromPartitionCnt 通过partition_cnt获取内容
func (obj *_AllVirtualMemstoreInfoMgr) GetFromPartitionCnt(partitionCnt int64) (results []*AllVirtualMemstoreInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMemstoreInfo{}).Where("`partition_cnt` = ?", partitionCnt).Find(&results).Error

	return
}

// GetBatchFromPartitionCnt 批量查找
func (obj *_AllVirtualMemstoreInfoMgr) GetBatchFromPartitionCnt(partitionCnts []int64) (results []*AllVirtualMemstoreInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMemstoreInfo{}).Where("`partition_cnt` IN (?)", partitionCnts).Find(&results).Error

	return
}

// GetFromVersion 通过version获取内容
func (obj *_AllVirtualMemstoreInfoMgr) GetFromVersion(version string) (results []*AllVirtualMemstoreInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMemstoreInfo{}).Where("`version` = ?", version).Find(&results).Error

	return
}

// GetBatchFromVersion 批量查找
func (obj *_AllVirtualMemstoreInfoMgr) GetBatchFromVersion(versions []string) (results []*AllVirtualMemstoreInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMemstoreInfo{}).Where("`version` IN (?)", versions).Find(&results).Error

	return
}

// GetFromBaseVersion 通过base_version获取内容
func (obj *_AllVirtualMemstoreInfoMgr) GetFromBaseVersion(baseVersion int64) (results []*AllVirtualMemstoreInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMemstoreInfo{}).Where("`base_version` = ?", baseVersion).Find(&results).Error

	return
}

// GetBatchFromBaseVersion 批量查找
func (obj *_AllVirtualMemstoreInfoMgr) GetBatchFromBaseVersion(baseVersions []int64) (results []*AllVirtualMemstoreInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMemstoreInfo{}).Where("`base_version` IN (?)", baseVersions).Find(&results).Error

	return
}

// GetFromMultiVersionStart 通过multi_version_start获取内容
func (obj *_AllVirtualMemstoreInfoMgr) GetFromMultiVersionStart(multiVersionStart int64) (results []*AllVirtualMemstoreInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMemstoreInfo{}).Where("`multi_version_start` = ?", multiVersionStart).Find(&results).Error

	return
}

// GetBatchFromMultiVersionStart 批量查找
func (obj *_AllVirtualMemstoreInfoMgr) GetBatchFromMultiVersionStart(multiVersionStarts []int64) (results []*AllVirtualMemstoreInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMemstoreInfo{}).Where("`multi_version_start` IN (?)", multiVersionStarts).Find(&results).Error

	return
}

// GetFromSnapshotVersion 通过snapshot_version获取内容
func (obj *_AllVirtualMemstoreInfoMgr) GetFromSnapshotVersion(snapshotVersion int64) (results []*AllVirtualMemstoreInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMemstoreInfo{}).Where("`snapshot_version` = ?", snapshotVersion).Find(&results).Error

	return
}

// GetBatchFromSnapshotVersion 批量查找
func (obj *_AllVirtualMemstoreInfoMgr) GetBatchFromSnapshotVersion(snapshotVersions []int64) (results []*AllVirtualMemstoreInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMemstoreInfo{}).Where("`snapshot_version` IN (?)", snapshotVersions).Find(&results).Error

	return
}

// GetFromIsActive 通过is_active获取内容
func (obj *_AllVirtualMemstoreInfoMgr) GetFromIsActive(isActive int64) (results []*AllVirtualMemstoreInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMemstoreInfo{}).Where("`is_active` = ?", isActive).Find(&results).Error

	return
}

// GetBatchFromIsActive 批量查找
func (obj *_AllVirtualMemstoreInfoMgr) GetBatchFromIsActive(isActives []int64) (results []*AllVirtualMemstoreInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMemstoreInfo{}).Where("`is_active` IN (?)", isActives).Find(&results).Error

	return
}

// GetFromMemUsed 通过mem_used获取内容
func (obj *_AllVirtualMemstoreInfoMgr) GetFromMemUsed(memUsed int64) (results []*AllVirtualMemstoreInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMemstoreInfo{}).Where("`mem_used` = ?", memUsed).Find(&results).Error

	return
}

// GetBatchFromMemUsed 批量查找
func (obj *_AllVirtualMemstoreInfoMgr) GetBatchFromMemUsed(memUseds []int64) (results []*AllVirtualMemstoreInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMemstoreInfo{}).Where("`mem_used` IN (?)", memUseds).Find(&results).Error

	return
}

// GetFromHashItemCount 通过hash_item_count获取内容
func (obj *_AllVirtualMemstoreInfoMgr) GetFromHashItemCount(hashItemCount int64) (results []*AllVirtualMemstoreInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMemstoreInfo{}).Where("`hash_item_count` = ?", hashItemCount).Find(&results).Error

	return
}

// GetBatchFromHashItemCount 批量查找
func (obj *_AllVirtualMemstoreInfoMgr) GetBatchFromHashItemCount(hashItemCounts []int64) (results []*AllVirtualMemstoreInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMemstoreInfo{}).Where("`hash_item_count` IN (?)", hashItemCounts).Find(&results).Error

	return
}

// GetFromHashMemUsed 通过hash_mem_used获取内容
func (obj *_AllVirtualMemstoreInfoMgr) GetFromHashMemUsed(hashMemUsed int64) (results []*AllVirtualMemstoreInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMemstoreInfo{}).Where("`hash_mem_used` = ?", hashMemUsed).Find(&results).Error

	return
}

// GetBatchFromHashMemUsed 批量查找
func (obj *_AllVirtualMemstoreInfoMgr) GetBatchFromHashMemUsed(hashMemUseds []int64) (results []*AllVirtualMemstoreInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMemstoreInfo{}).Where("`hash_mem_used` IN (?)", hashMemUseds).Find(&results).Error

	return
}

// GetFromBtreeItemCount 通过btree_item_count获取内容
func (obj *_AllVirtualMemstoreInfoMgr) GetFromBtreeItemCount(btreeItemCount int64) (results []*AllVirtualMemstoreInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMemstoreInfo{}).Where("`btree_item_count` = ?", btreeItemCount).Find(&results).Error

	return
}

// GetBatchFromBtreeItemCount 批量查找
func (obj *_AllVirtualMemstoreInfoMgr) GetBatchFromBtreeItemCount(btreeItemCounts []int64) (results []*AllVirtualMemstoreInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMemstoreInfo{}).Where("`btree_item_count` IN (?)", btreeItemCounts).Find(&results).Error

	return
}

// GetFromBtreeMemUsed 通过btree_mem_used获取内容
func (obj *_AllVirtualMemstoreInfoMgr) GetFromBtreeMemUsed(btreeMemUsed int64) (results []*AllVirtualMemstoreInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMemstoreInfo{}).Where("`btree_mem_used` = ?", btreeMemUsed).Find(&results).Error

	return
}

// GetBatchFromBtreeMemUsed 批量查找
func (obj *_AllVirtualMemstoreInfoMgr) GetBatchFromBtreeMemUsed(btreeMemUseds []int64) (results []*AllVirtualMemstoreInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMemstoreInfo{}).Where("`btree_mem_used` IN (?)", btreeMemUseds).Find(&results).Error

	return
}

// GetFromInsertRowCount 通过insert_row_count获取内容
func (obj *_AllVirtualMemstoreInfoMgr) GetFromInsertRowCount(insertRowCount int64) (results []*AllVirtualMemstoreInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMemstoreInfo{}).Where("`insert_row_count` = ?", insertRowCount).Find(&results).Error

	return
}

// GetBatchFromInsertRowCount 批量查找
func (obj *_AllVirtualMemstoreInfoMgr) GetBatchFromInsertRowCount(insertRowCounts []int64) (results []*AllVirtualMemstoreInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMemstoreInfo{}).Where("`insert_row_count` IN (?)", insertRowCounts).Find(&results).Error

	return
}

// GetFromUpdateRowCount 通过update_row_count获取内容
func (obj *_AllVirtualMemstoreInfoMgr) GetFromUpdateRowCount(updateRowCount int64) (results []*AllVirtualMemstoreInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMemstoreInfo{}).Where("`update_row_count` = ?", updateRowCount).Find(&results).Error

	return
}

// GetBatchFromUpdateRowCount 批量查找
func (obj *_AllVirtualMemstoreInfoMgr) GetBatchFromUpdateRowCount(updateRowCounts []int64) (results []*AllVirtualMemstoreInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMemstoreInfo{}).Where("`update_row_count` IN (?)", updateRowCounts).Find(&results).Error

	return
}

// GetFromDeleteRowCount 通过delete_row_count获取内容
func (obj *_AllVirtualMemstoreInfoMgr) GetFromDeleteRowCount(deleteRowCount int64) (results []*AllVirtualMemstoreInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMemstoreInfo{}).Where("`delete_row_count` = ?", deleteRowCount).Find(&results).Error

	return
}

// GetBatchFromDeleteRowCount 批量查找
func (obj *_AllVirtualMemstoreInfoMgr) GetBatchFromDeleteRowCount(deleteRowCounts []int64) (results []*AllVirtualMemstoreInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMemstoreInfo{}).Where("`delete_row_count` IN (?)", deleteRowCounts).Find(&results).Error

	return
}

// GetFromPurgeRowCount 通过purge_row_count获取内容
func (obj *_AllVirtualMemstoreInfoMgr) GetFromPurgeRowCount(purgeRowCount int64) (results []*AllVirtualMemstoreInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMemstoreInfo{}).Where("`purge_row_count` = ?", purgeRowCount).Find(&results).Error

	return
}

// GetBatchFromPurgeRowCount 批量查找
func (obj *_AllVirtualMemstoreInfoMgr) GetBatchFromPurgeRowCount(purgeRowCounts []int64) (results []*AllVirtualMemstoreInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMemstoreInfo{}).Where("`purge_row_count` IN (?)", purgeRowCounts).Find(&results).Error

	return
}

// GetFromPurgeQueueCount 通过purge_queue_count获取内容
func (obj *_AllVirtualMemstoreInfoMgr) GetFromPurgeQueueCount(purgeQueueCount int64) (results []*AllVirtualMemstoreInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMemstoreInfo{}).Where("`purge_queue_count` = ?", purgeQueueCount).Find(&results).Error

	return
}

// GetBatchFromPurgeQueueCount 批量查找
func (obj *_AllVirtualMemstoreInfoMgr) GetBatchFromPurgeQueueCount(purgeQueueCounts []int64) (results []*AllVirtualMemstoreInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualMemstoreInfo{}).Where("`purge_queue_count` IN (?)", purgeQueueCounts).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////
