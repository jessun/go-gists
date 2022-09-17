package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _AllVirtualPartitionStoreInfoMgr struct {
	*_BaseMgr
}

// AllVirtualPartitionStoreInfoMgr open func
func AllVirtualPartitionStoreInfoMgr(db *gorm.DB) *_AllVirtualPartitionStoreInfoMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualPartitionStoreInfoMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualPartitionStoreInfoMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_partition_store_info"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualPartitionStoreInfoMgr) GetTableName() string {
	return "__all_virtual_partition_store_info"
}

// Reset 重置gorm会话
func (obj *_AllVirtualPartitionStoreInfoMgr) Reset() *_AllVirtualPartitionStoreInfoMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualPartitionStoreInfoMgr) Get() (result AllVirtualPartitionStoreInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionStoreInfo{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualPartitionStoreInfoMgr) Gets() (results []*AllVirtualPartitionStoreInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionStoreInfo{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualPartitionStoreInfoMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionStoreInfo{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithSvrIP svr_ip获取
func (obj *_AllVirtualPartitionStoreInfoMgr) WithSvrIP(svrIP string) Option {
	return optionFunc(func(o *options) { o.query["svr_ip"] = svrIP })
}

// WithSvrPort svr_port获取
func (obj *_AllVirtualPartitionStoreInfoMgr) WithSvrPort(svrPort int64) Option {
	return optionFunc(func(o *options) { o.query["svr_port"] = svrPort })
}

// WithTenantID tenant_id获取
func (obj *_AllVirtualPartitionStoreInfoMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithTableID table_id获取
func (obj *_AllVirtualPartitionStoreInfoMgr) WithTableID(tableID int64) Option {
	return optionFunc(func(o *options) { o.query["table_id"] = tableID })
}

// WithPartitionIDx partition_idx获取
func (obj *_AllVirtualPartitionStoreInfoMgr) WithPartitionIDx(partitionIDx int64) Option {
	return optionFunc(func(o *options) { o.query["partition_idx"] = partitionIDx })
}

// WithPartitionCnt partition_cnt获取
func (obj *_AllVirtualPartitionStoreInfoMgr) WithPartitionCnt(partitionCnt int64) Option {
	return optionFunc(func(o *options) { o.query["partition_cnt"] = partitionCnt })
}

// WithIsRestore is_restore获取
func (obj *_AllVirtualPartitionStoreInfoMgr) WithIsRestore(isRestore int64) Option {
	return optionFunc(func(o *options) { o.query["is_restore"] = isRestore })
}

// WithMigrateStatus migrate_status获取
func (obj *_AllVirtualPartitionStoreInfoMgr) WithMigrateStatus(migrateStatus int64) Option {
	return optionFunc(func(o *options) { o.query["migrate_status"] = migrateStatus })
}

// WithMigrateTimestamp migrate_timestamp获取
func (obj *_AllVirtualPartitionStoreInfoMgr) WithMigrateTimestamp(migrateTimestamp int64) Option {
	return optionFunc(func(o *options) { o.query["migrate_timestamp"] = migrateTimestamp })
}

// WithReplicaType replica_type获取
func (obj *_AllVirtualPartitionStoreInfoMgr) WithReplicaType(replicaType int64) Option {
	return optionFunc(func(o *options) { o.query["replica_type"] = replicaType })
}

// WithSplitState split_state获取
func (obj *_AllVirtualPartitionStoreInfoMgr) WithSplitState(splitState int64) Option {
	return optionFunc(func(o *options) { o.query["split_state"] = splitState })
}

// WithMultiVersionStart multi_version_start获取
func (obj *_AllVirtualPartitionStoreInfoMgr) WithMultiVersionStart(multiVersionStart int64) Option {
	return optionFunc(func(o *options) { o.query["multi_version_start"] = multiVersionStart })
}

// WithReportVersion report_version获取
func (obj *_AllVirtualPartitionStoreInfoMgr) WithReportVersion(reportVersion int64) Option {
	return optionFunc(func(o *options) { o.query["report_version"] = reportVersion })
}

// WithReportRowCount report_row_count获取
func (obj *_AllVirtualPartitionStoreInfoMgr) WithReportRowCount(reportRowCount int64) Option {
	return optionFunc(func(o *options) { o.query["report_row_count"] = reportRowCount })
}

// WithReportDataChecksum report_data_checksum获取
func (obj *_AllVirtualPartitionStoreInfoMgr) WithReportDataChecksum(reportDataChecksum int64) Option {
	return optionFunc(func(o *options) { o.query["report_data_checksum"] = reportDataChecksum })
}

// WithReportDataSize report_data_size获取
func (obj *_AllVirtualPartitionStoreInfoMgr) WithReportDataSize(reportDataSize int64) Option {
	return optionFunc(func(o *options) { o.query["report_data_size"] = reportDataSize })
}

// WithReportRequiredSize report_required_size获取
func (obj *_AllVirtualPartitionStoreInfoMgr) WithReportRequiredSize(reportRequiredSize int64) Option {
	return optionFunc(func(o *options) { o.query["report_required_size"] = reportRequiredSize })
}

// WithReadableTs readable_ts获取
func (obj *_AllVirtualPartitionStoreInfoMgr) WithReadableTs(readableTs int64) Option {
	return optionFunc(func(o *options) { o.query["readable_ts"] = readableTs })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualPartitionStoreInfoMgr) GetByOption(opts ...Option) (result AllVirtualPartitionStoreInfo, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionStoreInfo{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualPartitionStoreInfoMgr) GetByOptions(opts ...Option) (results []*AllVirtualPartitionStoreInfo, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionStoreInfo{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualPartitionStoreInfoMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualPartitionStoreInfo, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionStoreInfo{}).Where(options.query)
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
func (obj *_AllVirtualPartitionStoreInfoMgr) GetFromSvrIP(svrIP string) (results []*AllVirtualPartitionStoreInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionStoreInfo{}).Where("`svr_ip` = ?", svrIP).Find(&results).Error

	return
}

// GetBatchFromSvrIP 批量查找
func (obj *_AllVirtualPartitionStoreInfoMgr) GetBatchFromSvrIP(svrIPs []string) (results []*AllVirtualPartitionStoreInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionStoreInfo{}).Where("`svr_ip` IN (?)", svrIPs).Find(&results).Error

	return
}

// GetFromSvrPort 通过svr_port获取内容
func (obj *_AllVirtualPartitionStoreInfoMgr) GetFromSvrPort(svrPort int64) (results []*AllVirtualPartitionStoreInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionStoreInfo{}).Where("`svr_port` = ?", svrPort).Find(&results).Error

	return
}

// GetBatchFromSvrPort 批量查找
func (obj *_AllVirtualPartitionStoreInfoMgr) GetBatchFromSvrPort(svrPorts []int64) (results []*AllVirtualPartitionStoreInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionStoreInfo{}).Where("`svr_port` IN (?)", svrPorts).Find(&results).Error

	return
}

// GetFromTenantID 通过tenant_id获取内容
func (obj *_AllVirtualPartitionStoreInfoMgr) GetFromTenantID(tenantID int64) (results []*AllVirtualPartitionStoreInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionStoreInfo{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllVirtualPartitionStoreInfoMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllVirtualPartitionStoreInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionStoreInfo{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromTableID 通过table_id获取内容
func (obj *_AllVirtualPartitionStoreInfoMgr) GetFromTableID(tableID int64) (results []*AllVirtualPartitionStoreInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionStoreInfo{}).Where("`table_id` = ?", tableID).Find(&results).Error

	return
}

// GetBatchFromTableID 批量查找
func (obj *_AllVirtualPartitionStoreInfoMgr) GetBatchFromTableID(tableIDs []int64) (results []*AllVirtualPartitionStoreInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionStoreInfo{}).Where("`table_id` IN (?)", tableIDs).Find(&results).Error

	return
}

// GetFromPartitionIDx 通过partition_idx获取内容
func (obj *_AllVirtualPartitionStoreInfoMgr) GetFromPartitionIDx(partitionIDx int64) (results []*AllVirtualPartitionStoreInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionStoreInfo{}).Where("`partition_idx` = ?", partitionIDx).Find(&results).Error

	return
}

// GetBatchFromPartitionIDx 批量查找
func (obj *_AllVirtualPartitionStoreInfoMgr) GetBatchFromPartitionIDx(partitionIDxs []int64) (results []*AllVirtualPartitionStoreInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionStoreInfo{}).Where("`partition_idx` IN (?)", partitionIDxs).Find(&results).Error

	return
}

// GetFromPartitionCnt 通过partition_cnt获取内容
func (obj *_AllVirtualPartitionStoreInfoMgr) GetFromPartitionCnt(partitionCnt int64) (results []*AllVirtualPartitionStoreInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionStoreInfo{}).Where("`partition_cnt` = ?", partitionCnt).Find(&results).Error

	return
}

// GetBatchFromPartitionCnt 批量查找
func (obj *_AllVirtualPartitionStoreInfoMgr) GetBatchFromPartitionCnt(partitionCnts []int64) (results []*AllVirtualPartitionStoreInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionStoreInfo{}).Where("`partition_cnt` IN (?)", partitionCnts).Find(&results).Error

	return
}

// GetFromIsRestore 通过is_restore获取内容
func (obj *_AllVirtualPartitionStoreInfoMgr) GetFromIsRestore(isRestore int64) (results []*AllVirtualPartitionStoreInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionStoreInfo{}).Where("`is_restore` = ?", isRestore).Find(&results).Error

	return
}

// GetBatchFromIsRestore 批量查找
func (obj *_AllVirtualPartitionStoreInfoMgr) GetBatchFromIsRestore(isRestores []int64) (results []*AllVirtualPartitionStoreInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionStoreInfo{}).Where("`is_restore` IN (?)", isRestores).Find(&results).Error

	return
}

// GetFromMigrateStatus 通过migrate_status获取内容
func (obj *_AllVirtualPartitionStoreInfoMgr) GetFromMigrateStatus(migrateStatus int64) (results []*AllVirtualPartitionStoreInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionStoreInfo{}).Where("`migrate_status` = ?", migrateStatus).Find(&results).Error

	return
}

// GetBatchFromMigrateStatus 批量查找
func (obj *_AllVirtualPartitionStoreInfoMgr) GetBatchFromMigrateStatus(migrateStatuss []int64) (results []*AllVirtualPartitionStoreInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionStoreInfo{}).Where("`migrate_status` IN (?)", migrateStatuss).Find(&results).Error

	return
}

// GetFromMigrateTimestamp 通过migrate_timestamp获取内容
func (obj *_AllVirtualPartitionStoreInfoMgr) GetFromMigrateTimestamp(migrateTimestamp int64) (results []*AllVirtualPartitionStoreInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionStoreInfo{}).Where("`migrate_timestamp` = ?", migrateTimestamp).Find(&results).Error

	return
}

// GetBatchFromMigrateTimestamp 批量查找
func (obj *_AllVirtualPartitionStoreInfoMgr) GetBatchFromMigrateTimestamp(migrateTimestamps []int64) (results []*AllVirtualPartitionStoreInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionStoreInfo{}).Where("`migrate_timestamp` IN (?)", migrateTimestamps).Find(&results).Error

	return
}

// GetFromReplicaType 通过replica_type获取内容
func (obj *_AllVirtualPartitionStoreInfoMgr) GetFromReplicaType(replicaType int64) (results []*AllVirtualPartitionStoreInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionStoreInfo{}).Where("`replica_type` = ?", replicaType).Find(&results).Error

	return
}

// GetBatchFromReplicaType 批量查找
func (obj *_AllVirtualPartitionStoreInfoMgr) GetBatchFromReplicaType(replicaTypes []int64) (results []*AllVirtualPartitionStoreInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionStoreInfo{}).Where("`replica_type` IN (?)", replicaTypes).Find(&results).Error

	return
}

// GetFromSplitState 通过split_state获取内容
func (obj *_AllVirtualPartitionStoreInfoMgr) GetFromSplitState(splitState int64) (results []*AllVirtualPartitionStoreInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionStoreInfo{}).Where("`split_state` = ?", splitState).Find(&results).Error

	return
}

// GetBatchFromSplitState 批量查找
func (obj *_AllVirtualPartitionStoreInfoMgr) GetBatchFromSplitState(splitStates []int64) (results []*AllVirtualPartitionStoreInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionStoreInfo{}).Where("`split_state` IN (?)", splitStates).Find(&results).Error

	return
}

// GetFromMultiVersionStart 通过multi_version_start获取内容
func (obj *_AllVirtualPartitionStoreInfoMgr) GetFromMultiVersionStart(multiVersionStart int64) (results []*AllVirtualPartitionStoreInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionStoreInfo{}).Where("`multi_version_start` = ?", multiVersionStart).Find(&results).Error

	return
}

// GetBatchFromMultiVersionStart 批量查找
func (obj *_AllVirtualPartitionStoreInfoMgr) GetBatchFromMultiVersionStart(multiVersionStarts []int64) (results []*AllVirtualPartitionStoreInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionStoreInfo{}).Where("`multi_version_start` IN (?)", multiVersionStarts).Find(&results).Error

	return
}

// GetFromReportVersion 通过report_version获取内容
func (obj *_AllVirtualPartitionStoreInfoMgr) GetFromReportVersion(reportVersion int64) (results []*AllVirtualPartitionStoreInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionStoreInfo{}).Where("`report_version` = ?", reportVersion).Find(&results).Error

	return
}

// GetBatchFromReportVersion 批量查找
func (obj *_AllVirtualPartitionStoreInfoMgr) GetBatchFromReportVersion(reportVersions []int64) (results []*AllVirtualPartitionStoreInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionStoreInfo{}).Where("`report_version` IN (?)", reportVersions).Find(&results).Error

	return
}

// GetFromReportRowCount 通过report_row_count获取内容
func (obj *_AllVirtualPartitionStoreInfoMgr) GetFromReportRowCount(reportRowCount int64) (results []*AllVirtualPartitionStoreInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionStoreInfo{}).Where("`report_row_count` = ?", reportRowCount).Find(&results).Error

	return
}

// GetBatchFromReportRowCount 批量查找
func (obj *_AllVirtualPartitionStoreInfoMgr) GetBatchFromReportRowCount(reportRowCounts []int64) (results []*AllVirtualPartitionStoreInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionStoreInfo{}).Where("`report_row_count` IN (?)", reportRowCounts).Find(&results).Error

	return
}

// GetFromReportDataChecksum 通过report_data_checksum获取内容
func (obj *_AllVirtualPartitionStoreInfoMgr) GetFromReportDataChecksum(reportDataChecksum int64) (results []*AllVirtualPartitionStoreInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionStoreInfo{}).Where("`report_data_checksum` = ?", reportDataChecksum).Find(&results).Error

	return
}

// GetBatchFromReportDataChecksum 批量查找
func (obj *_AllVirtualPartitionStoreInfoMgr) GetBatchFromReportDataChecksum(reportDataChecksums []int64) (results []*AllVirtualPartitionStoreInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionStoreInfo{}).Where("`report_data_checksum` IN (?)", reportDataChecksums).Find(&results).Error

	return
}

// GetFromReportDataSize 通过report_data_size获取内容
func (obj *_AllVirtualPartitionStoreInfoMgr) GetFromReportDataSize(reportDataSize int64) (results []*AllVirtualPartitionStoreInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionStoreInfo{}).Where("`report_data_size` = ?", reportDataSize).Find(&results).Error

	return
}

// GetBatchFromReportDataSize 批量查找
func (obj *_AllVirtualPartitionStoreInfoMgr) GetBatchFromReportDataSize(reportDataSizes []int64) (results []*AllVirtualPartitionStoreInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionStoreInfo{}).Where("`report_data_size` IN (?)", reportDataSizes).Find(&results).Error

	return
}

// GetFromReportRequiredSize 通过report_required_size获取内容
func (obj *_AllVirtualPartitionStoreInfoMgr) GetFromReportRequiredSize(reportRequiredSize int64) (results []*AllVirtualPartitionStoreInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionStoreInfo{}).Where("`report_required_size` = ?", reportRequiredSize).Find(&results).Error

	return
}

// GetBatchFromReportRequiredSize 批量查找
func (obj *_AllVirtualPartitionStoreInfoMgr) GetBatchFromReportRequiredSize(reportRequiredSizes []int64) (results []*AllVirtualPartitionStoreInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionStoreInfo{}).Where("`report_required_size` IN (?)", reportRequiredSizes).Find(&results).Error

	return
}

// GetFromReadableTs 通过readable_ts获取内容
func (obj *_AllVirtualPartitionStoreInfoMgr) GetFromReadableTs(readableTs int64) (results []*AllVirtualPartitionStoreInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionStoreInfo{}).Where("`readable_ts` = ?", readableTs).Find(&results).Error

	return
}

// GetBatchFromReadableTs 批量查找
func (obj *_AllVirtualPartitionStoreInfoMgr) GetBatchFromReadableTs(readableTss []int64) (results []*AllVirtualPartitionStoreInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualPartitionStoreInfo{}).Where("`readable_ts` IN (?)", readableTss).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////
