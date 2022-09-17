package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _AllVirtualWeakReadStatMgr struct {
	*_BaseMgr
}

// AllVirtualWeakReadStatMgr open func
func AllVirtualWeakReadStatMgr(db *gorm.DB) *_AllVirtualWeakReadStatMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualWeakReadStatMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualWeakReadStatMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_weak_read_stat"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualWeakReadStatMgr) GetTableName() string {
	return "__all_virtual_weak_read_stat"
}

// Reset 重置gorm会话
func (obj *_AllVirtualWeakReadStatMgr) Reset() *_AllVirtualWeakReadStatMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualWeakReadStatMgr) Get() (result AllVirtualWeakReadStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualWeakReadStat{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualWeakReadStatMgr) Gets() (results []*AllVirtualWeakReadStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualWeakReadStat{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualWeakReadStatMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualWeakReadStat{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTenantID tenant_id获取
func (obj *_AllVirtualWeakReadStatMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithSvrIP svr_ip获取
func (obj *_AllVirtualWeakReadStatMgr) WithSvrIP(svrIP string) Option {
	return optionFunc(func(o *options) { o.query["svr_ip"] = svrIP })
}

// WithSvrPort svr_port获取
func (obj *_AllVirtualWeakReadStatMgr) WithSvrPort(svrPort int64) Option {
	return optionFunc(func(o *options) { o.query["svr_port"] = svrPort })
}

// WithServerVersion server_version获取
func (obj *_AllVirtualWeakReadStatMgr) WithServerVersion(serverVersion int64) Option {
	return optionFunc(func(o *options) { o.query["server_version"] = serverVersion })
}

// WithServerVersionDelta server_version_delta获取
func (obj *_AllVirtualWeakReadStatMgr) WithServerVersionDelta(serverVersionDelta int64) Option {
	return optionFunc(func(o *options) { o.query["server_version_delta"] = serverVersionDelta })
}

// WithLocalClusterVersion local_cluster_version获取
func (obj *_AllVirtualWeakReadStatMgr) WithLocalClusterVersion(localClusterVersion int64) Option {
	return optionFunc(func(o *options) { o.query["local_cluster_version"] = localClusterVersion })
}

// WithLocalClusterVersionDelta local_cluster_version_delta获取
func (obj *_AllVirtualWeakReadStatMgr) WithLocalClusterVersionDelta(localClusterVersionDelta int64) Option {
	return optionFunc(func(o *options) { o.query["local_cluster_version_delta"] = localClusterVersionDelta })
}

// WithTotalPartCount total_part_count获取
func (obj *_AllVirtualWeakReadStatMgr) WithTotalPartCount(totalPartCount int64) Option {
	return optionFunc(func(o *options) { o.query["total_part_count"] = totalPartCount })
}

// WithValidInnerPartCount valid_inner_part_count获取
func (obj *_AllVirtualWeakReadStatMgr) WithValidInnerPartCount(validInnerPartCount int64) Option {
	return optionFunc(func(o *options) { o.query["valid_inner_part_count"] = validInnerPartCount })
}

// WithValidUserPartCount valid_user_part_count获取
func (obj *_AllVirtualWeakReadStatMgr) WithValidUserPartCount(validUserPartCount int64) Option {
	return optionFunc(func(o *options) { o.query["valid_user_part_count"] = validUserPartCount })
}

// WithClusterMasterIP cluster_master_ip获取
func (obj *_AllVirtualWeakReadStatMgr) WithClusterMasterIP(clusterMasterIP string) Option {
	return optionFunc(func(o *options) { o.query["cluster_master_ip"] = clusterMasterIP })
}

// WithClusterMasterPort cluster_master_port获取
func (obj *_AllVirtualWeakReadStatMgr) WithClusterMasterPort(clusterMasterPort int64) Option {
	return optionFunc(func(o *options) { o.query["cluster_master_port"] = clusterMasterPort })
}

// WithClusterHeartbeatPostTstamp cluster_heartbeat_post_tstamp获取
func (obj *_AllVirtualWeakReadStatMgr) WithClusterHeartbeatPostTstamp(clusterHeartbeatPostTstamp int64) Option {
	return optionFunc(func(o *options) { o.query["cluster_heartbeat_post_tstamp"] = clusterHeartbeatPostTstamp })
}

// WithClusterHeartbeatPostCount cluster_heartbeat_post_count获取
func (obj *_AllVirtualWeakReadStatMgr) WithClusterHeartbeatPostCount(clusterHeartbeatPostCount int64) Option {
	return optionFunc(func(o *options) { o.query["cluster_heartbeat_post_count"] = clusterHeartbeatPostCount })
}

// WithClusterHeartbeatSuccTstamp cluster_heartbeat_succ_tstamp获取
func (obj *_AllVirtualWeakReadStatMgr) WithClusterHeartbeatSuccTstamp(clusterHeartbeatSuccTstamp int64) Option {
	return optionFunc(func(o *options) { o.query["cluster_heartbeat_succ_tstamp"] = clusterHeartbeatSuccTstamp })
}

// WithClusterHeartbeatSuccCount cluster_heartbeat_succ_count获取
func (obj *_AllVirtualWeakReadStatMgr) WithClusterHeartbeatSuccCount(clusterHeartbeatSuccCount int64) Option {
	return optionFunc(func(o *options) { o.query["cluster_heartbeat_succ_count"] = clusterHeartbeatSuccCount })
}

// WithSelfCheckTstamp self_check_tstamp获取
func (obj *_AllVirtualWeakReadStatMgr) WithSelfCheckTstamp(selfCheckTstamp int64) Option {
	return optionFunc(func(o *options) { o.query["self_check_tstamp"] = selfCheckTstamp })
}

// WithLocalCurrentTstamp local_current_tstamp获取
func (obj *_AllVirtualWeakReadStatMgr) WithLocalCurrentTstamp(localCurrentTstamp int64) Option {
	return optionFunc(func(o *options) { o.query["local_current_tstamp"] = localCurrentTstamp })
}

// WithInClusterService in_cluster_service获取
func (obj *_AllVirtualWeakReadStatMgr) WithInClusterService(inClusterService int64) Option {
	return optionFunc(func(o *options) { o.query["in_cluster_service"] = inClusterService })
}

// WithIsClusterServiceMaster is_cluster_service_master获取
func (obj *_AllVirtualWeakReadStatMgr) WithIsClusterServiceMaster(isClusterServiceMaster int64) Option {
	return optionFunc(func(o *options) { o.query["is_cluster_service_master"] = isClusterServiceMaster })
}

// WithClusterServiceEpoch cluster_service_epoch获取
func (obj *_AllVirtualWeakReadStatMgr) WithClusterServiceEpoch(clusterServiceEpoch int64) Option {
	return optionFunc(func(o *options) { o.query["cluster_service_epoch"] = clusterServiceEpoch })
}

// WithClusterTotalServerCount cluster_total_server_count获取
func (obj *_AllVirtualWeakReadStatMgr) WithClusterTotalServerCount(clusterTotalServerCount int64) Option {
	return optionFunc(func(o *options) { o.query["cluster_total_server_count"] = clusterTotalServerCount })
}

// WithClusterSkippedServerCount cluster_skipped_server_count获取
func (obj *_AllVirtualWeakReadStatMgr) WithClusterSkippedServerCount(clusterSkippedServerCount int64) Option {
	return optionFunc(func(o *options) { o.query["cluster_skipped_server_count"] = clusterSkippedServerCount })
}

// WithClusterVersionGenTstamp cluster_version_gen_tstamp获取
func (obj *_AllVirtualWeakReadStatMgr) WithClusterVersionGenTstamp(clusterVersionGenTstamp int64) Option {
	return optionFunc(func(o *options) { o.query["cluster_version_gen_tstamp"] = clusterVersionGenTstamp })
}

// WithClusterVersion cluster_version获取
func (obj *_AllVirtualWeakReadStatMgr) WithClusterVersion(clusterVersion int64) Option {
	return optionFunc(func(o *options) { o.query["cluster_version"] = clusterVersion })
}

// WithClusterVersionDelta cluster_version_delta获取
func (obj *_AllVirtualWeakReadStatMgr) WithClusterVersionDelta(clusterVersionDelta int64) Option {
	return optionFunc(func(o *options) { o.query["cluster_version_delta"] = clusterVersionDelta })
}

// WithMinClusterVersion min_cluster_version获取
func (obj *_AllVirtualWeakReadStatMgr) WithMinClusterVersion(minClusterVersion int64) Option {
	return optionFunc(func(o *options) { o.query["min_cluster_version"] = minClusterVersion })
}

// WithMaxClusterVersion max_cluster_version获取
func (obj *_AllVirtualWeakReadStatMgr) WithMaxClusterVersion(maxClusterVersion int64) Option {
	return optionFunc(func(o *options) { o.query["max_cluster_version"] = maxClusterVersion })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualWeakReadStatMgr) GetByOption(opts ...Option) (result AllVirtualWeakReadStat, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualWeakReadStat{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualWeakReadStatMgr) GetByOptions(opts ...Option) (results []*AllVirtualWeakReadStat, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualWeakReadStat{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualWeakReadStatMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualWeakReadStat, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualWeakReadStat{}).Where(options.query)
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
func (obj *_AllVirtualWeakReadStatMgr) GetFromTenantID(tenantID int64) (results []*AllVirtualWeakReadStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualWeakReadStat{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllVirtualWeakReadStatMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllVirtualWeakReadStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualWeakReadStat{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromSvrIP 通过svr_ip获取内容
func (obj *_AllVirtualWeakReadStatMgr) GetFromSvrIP(svrIP string) (results []*AllVirtualWeakReadStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualWeakReadStat{}).Where("`svr_ip` = ?", svrIP).Find(&results).Error

	return
}

// GetBatchFromSvrIP 批量查找
func (obj *_AllVirtualWeakReadStatMgr) GetBatchFromSvrIP(svrIPs []string) (results []*AllVirtualWeakReadStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualWeakReadStat{}).Where("`svr_ip` IN (?)", svrIPs).Find(&results).Error

	return
}

// GetFromSvrPort 通过svr_port获取内容
func (obj *_AllVirtualWeakReadStatMgr) GetFromSvrPort(svrPort int64) (results []*AllVirtualWeakReadStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualWeakReadStat{}).Where("`svr_port` = ?", svrPort).Find(&results).Error

	return
}

// GetBatchFromSvrPort 批量查找
func (obj *_AllVirtualWeakReadStatMgr) GetBatchFromSvrPort(svrPorts []int64) (results []*AllVirtualWeakReadStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualWeakReadStat{}).Where("`svr_port` IN (?)", svrPorts).Find(&results).Error

	return
}

// GetFromServerVersion 通过server_version获取内容
func (obj *_AllVirtualWeakReadStatMgr) GetFromServerVersion(serverVersion int64) (results []*AllVirtualWeakReadStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualWeakReadStat{}).Where("`server_version` = ?", serverVersion).Find(&results).Error

	return
}

// GetBatchFromServerVersion 批量查找
func (obj *_AllVirtualWeakReadStatMgr) GetBatchFromServerVersion(serverVersions []int64) (results []*AllVirtualWeakReadStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualWeakReadStat{}).Where("`server_version` IN (?)", serverVersions).Find(&results).Error

	return
}

// GetFromServerVersionDelta 通过server_version_delta获取内容
func (obj *_AllVirtualWeakReadStatMgr) GetFromServerVersionDelta(serverVersionDelta int64) (results []*AllVirtualWeakReadStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualWeakReadStat{}).Where("`server_version_delta` = ?", serverVersionDelta).Find(&results).Error

	return
}

// GetBatchFromServerVersionDelta 批量查找
func (obj *_AllVirtualWeakReadStatMgr) GetBatchFromServerVersionDelta(serverVersionDeltas []int64) (results []*AllVirtualWeakReadStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualWeakReadStat{}).Where("`server_version_delta` IN (?)", serverVersionDeltas).Find(&results).Error

	return
}

// GetFromLocalClusterVersion 通过local_cluster_version获取内容
func (obj *_AllVirtualWeakReadStatMgr) GetFromLocalClusterVersion(localClusterVersion int64) (results []*AllVirtualWeakReadStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualWeakReadStat{}).Where("`local_cluster_version` = ?", localClusterVersion).Find(&results).Error

	return
}

// GetBatchFromLocalClusterVersion 批量查找
func (obj *_AllVirtualWeakReadStatMgr) GetBatchFromLocalClusterVersion(localClusterVersions []int64) (results []*AllVirtualWeakReadStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualWeakReadStat{}).Where("`local_cluster_version` IN (?)", localClusterVersions).Find(&results).Error

	return
}

// GetFromLocalClusterVersionDelta 通过local_cluster_version_delta获取内容
func (obj *_AllVirtualWeakReadStatMgr) GetFromLocalClusterVersionDelta(localClusterVersionDelta int64) (results []*AllVirtualWeakReadStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualWeakReadStat{}).Where("`local_cluster_version_delta` = ?", localClusterVersionDelta).Find(&results).Error

	return
}

// GetBatchFromLocalClusterVersionDelta 批量查找
func (obj *_AllVirtualWeakReadStatMgr) GetBatchFromLocalClusterVersionDelta(localClusterVersionDeltas []int64) (results []*AllVirtualWeakReadStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualWeakReadStat{}).Where("`local_cluster_version_delta` IN (?)", localClusterVersionDeltas).Find(&results).Error

	return
}

// GetFromTotalPartCount 通过total_part_count获取内容
func (obj *_AllVirtualWeakReadStatMgr) GetFromTotalPartCount(totalPartCount int64) (results []*AllVirtualWeakReadStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualWeakReadStat{}).Where("`total_part_count` = ?", totalPartCount).Find(&results).Error

	return
}

// GetBatchFromTotalPartCount 批量查找
func (obj *_AllVirtualWeakReadStatMgr) GetBatchFromTotalPartCount(totalPartCounts []int64) (results []*AllVirtualWeakReadStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualWeakReadStat{}).Where("`total_part_count` IN (?)", totalPartCounts).Find(&results).Error

	return
}

// GetFromValidInnerPartCount 通过valid_inner_part_count获取内容
func (obj *_AllVirtualWeakReadStatMgr) GetFromValidInnerPartCount(validInnerPartCount int64) (results []*AllVirtualWeakReadStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualWeakReadStat{}).Where("`valid_inner_part_count` = ?", validInnerPartCount).Find(&results).Error

	return
}

// GetBatchFromValidInnerPartCount 批量查找
func (obj *_AllVirtualWeakReadStatMgr) GetBatchFromValidInnerPartCount(validInnerPartCounts []int64) (results []*AllVirtualWeakReadStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualWeakReadStat{}).Where("`valid_inner_part_count` IN (?)", validInnerPartCounts).Find(&results).Error

	return
}

// GetFromValidUserPartCount 通过valid_user_part_count获取内容
func (obj *_AllVirtualWeakReadStatMgr) GetFromValidUserPartCount(validUserPartCount int64) (results []*AllVirtualWeakReadStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualWeakReadStat{}).Where("`valid_user_part_count` = ?", validUserPartCount).Find(&results).Error

	return
}

// GetBatchFromValidUserPartCount 批量查找
func (obj *_AllVirtualWeakReadStatMgr) GetBatchFromValidUserPartCount(validUserPartCounts []int64) (results []*AllVirtualWeakReadStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualWeakReadStat{}).Where("`valid_user_part_count` IN (?)", validUserPartCounts).Find(&results).Error

	return
}

// GetFromClusterMasterIP 通过cluster_master_ip获取内容
func (obj *_AllVirtualWeakReadStatMgr) GetFromClusterMasterIP(clusterMasterIP string) (results []*AllVirtualWeakReadStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualWeakReadStat{}).Where("`cluster_master_ip` = ?", clusterMasterIP).Find(&results).Error

	return
}

// GetBatchFromClusterMasterIP 批量查找
func (obj *_AllVirtualWeakReadStatMgr) GetBatchFromClusterMasterIP(clusterMasterIPs []string) (results []*AllVirtualWeakReadStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualWeakReadStat{}).Where("`cluster_master_ip` IN (?)", clusterMasterIPs).Find(&results).Error

	return
}

// GetFromClusterMasterPort 通过cluster_master_port获取内容
func (obj *_AllVirtualWeakReadStatMgr) GetFromClusterMasterPort(clusterMasterPort int64) (results []*AllVirtualWeakReadStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualWeakReadStat{}).Where("`cluster_master_port` = ?", clusterMasterPort).Find(&results).Error

	return
}

// GetBatchFromClusterMasterPort 批量查找
func (obj *_AllVirtualWeakReadStatMgr) GetBatchFromClusterMasterPort(clusterMasterPorts []int64) (results []*AllVirtualWeakReadStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualWeakReadStat{}).Where("`cluster_master_port` IN (?)", clusterMasterPorts).Find(&results).Error

	return
}

// GetFromClusterHeartbeatPostTstamp 通过cluster_heartbeat_post_tstamp获取内容
func (obj *_AllVirtualWeakReadStatMgr) GetFromClusterHeartbeatPostTstamp(clusterHeartbeatPostTstamp int64) (results []*AllVirtualWeakReadStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualWeakReadStat{}).Where("`cluster_heartbeat_post_tstamp` = ?", clusterHeartbeatPostTstamp).Find(&results).Error

	return
}

// GetBatchFromClusterHeartbeatPostTstamp 批量查找
func (obj *_AllVirtualWeakReadStatMgr) GetBatchFromClusterHeartbeatPostTstamp(clusterHeartbeatPostTstamps []int64) (results []*AllVirtualWeakReadStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualWeakReadStat{}).Where("`cluster_heartbeat_post_tstamp` IN (?)", clusterHeartbeatPostTstamps).Find(&results).Error

	return
}

// GetFromClusterHeartbeatPostCount 通过cluster_heartbeat_post_count获取内容
func (obj *_AllVirtualWeakReadStatMgr) GetFromClusterHeartbeatPostCount(clusterHeartbeatPostCount int64) (results []*AllVirtualWeakReadStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualWeakReadStat{}).Where("`cluster_heartbeat_post_count` = ?", clusterHeartbeatPostCount).Find(&results).Error

	return
}

// GetBatchFromClusterHeartbeatPostCount 批量查找
func (obj *_AllVirtualWeakReadStatMgr) GetBatchFromClusterHeartbeatPostCount(clusterHeartbeatPostCounts []int64) (results []*AllVirtualWeakReadStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualWeakReadStat{}).Where("`cluster_heartbeat_post_count` IN (?)", clusterHeartbeatPostCounts).Find(&results).Error

	return
}

// GetFromClusterHeartbeatSuccTstamp 通过cluster_heartbeat_succ_tstamp获取内容
func (obj *_AllVirtualWeakReadStatMgr) GetFromClusterHeartbeatSuccTstamp(clusterHeartbeatSuccTstamp int64) (results []*AllVirtualWeakReadStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualWeakReadStat{}).Where("`cluster_heartbeat_succ_tstamp` = ?", clusterHeartbeatSuccTstamp).Find(&results).Error

	return
}

// GetBatchFromClusterHeartbeatSuccTstamp 批量查找
func (obj *_AllVirtualWeakReadStatMgr) GetBatchFromClusterHeartbeatSuccTstamp(clusterHeartbeatSuccTstamps []int64) (results []*AllVirtualWeakReadStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualWeakReadStat{}).Where("`cluster_heartbeat_succ_tstamp` IN (?)", clusterHeartbeatSuccTstamps).Find(&results).Error

	return
}

// GetFromClusterHeartbeatSuccCount 通过cluster_heartbeat_succ_count获取内容
func (obj *_AllVirtualWeakReadStatMgr) GetFromClusterHeartbeatSuccCount(clusterHeartbeatSuccCount int64) (results []*AllVirtualWeakReadStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualWeakReadStat{}).Where("`cluster_heartbeat_succ_count` = ?", clusterHeartbeatSuccCount).Find(&results).Error

	return
}

// GetBatchFromClusterHeartbeatSuccCount 批量查找
func (obj *_AllVirtualWeakReadStatMgr) GetBatchFromClusterHeartbeatSuccCount(clusterHeartbeatSuccCounts []int64) (results []*AllVirtualWeakReadStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualWeakReadStat{}).Where("`cluster_heartbeat_succ_count` IN (?)", clusterHeartbeatSuccCounts).Find(&results).Error

	return
}

// GetFromSelfCheckTstamp 通过self_check_tstamp获取内容
func (obj *_AllVirtualWeakReadStatMgr) GetFromSelfCheckTstamp(selfCheckTstamp int64) (results []*AllVirtualWeakReadStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualWeakReadStat{}).Where("`self_check_tstamp` = ?", selfCheckTstamp).Find(&results).Error

	return
}

// GetBatchFromSelfCheckTstamp 批量查找
func (obj *_AllVirtualWeakReadStatMgr) GetBatchFromSelfCheckTstamp(selfCheckTstamps []int64) (results []*AllVirtualWeakReadStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualWeakReadStat{}).Where("`self_check_tstamp` IN (?)", selfCheckTstamps).Find(&results).Error

	return
}

// GetFromLocalCurrentTstamp 通过local_current_tstamp获取内容
func (obj *_AllVirtualWeakReadStatMgr) GetFromLocalCurrentTstamp(localCurrentTstamp int64) (results []*AllVirtualWeakReadStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualWeakReadStat{}).Where("`local_current_tstamp` = ?", localCurrentTstamp).Find(&results).Error

	return
}

// GetBatchFromLocalCurrentTstamp 批量查找
func (obj *_AllVirtualWeakReadStatMgr) GetBatchFromLocalCurrentTstamp(localCurrentTstamps []int64) (results []*AllVirtualWeakReadStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualWeakReadStat{}).Where("`local_current_tstamp` IN (?)", localCurrentTstamps).Find(&results).Error

	return
}

// GetFromInClusterService 通过in_cluster_service获取内容
func (obj *_AllVirtualWeakReadStatMgr) GetFromInClusterService(inClusterService int64) (results []*AllVirtualWeakReadStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualWeakReadStat{}).Where("`in_cluster_service` = ?", inClusterService).Find(&results).Error

	return
}

// GetBatchFromInClusterService 批量查找
func (obj *_AllVirtualWeakReadStatMgr) GetBatchFromInClusterService(inClusterServices []int64) (results []*AllVirtualWeakReadStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualWeakReadStat{}).Where("`in_cluster_service` IN (?)", inClusterServices).Find(&results).Error

	return
}

// GetFromIsClusterServiceMaster 通过is_cluster_service_master获取内容
func (obj *_AllVirtualWeakReadStatMgr) GetFromIsClusterServiceMaster(isClusterServiceMaster int64) (results []*AllVirtualWeakReadStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualWeakReadStat{}).Where("`is_cluster_service_master` = ?", isClusterServiceMaster).Find(&results).Error

	return
}

// GetBatchFromIsClusterServiceMaster 批量查找
func (obj *_AllVirtualWeakReadStatMgr) GetBatchFromIsClusterServiceMaster(isClusterServiceMasters []int64) (results []*AllVirtualWeakReadStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualWeakReadStat{}).Where("`is_cluster_service_master` IN (?)", isClusterServiceMasters).Find(&results).Error

	return
}

// GetFromClusterServiceEpoch 通过cluster_service_epoch获取内容
func (obj *_AllVirtualWeakReadStatMgr) GetFromClusterServiceEpoch(clusterServiceEpoch int64) (results []*AllVirtualWeakReadStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualWeakReadStat{}).Where("`cluster_service_epoch` = ?", clusterServiceEpoch).Find(&results).Error

	return
}

// GetBatchFromClusterServiceEpoch 批量查找
func (obj *_AllVirtualWeakReadStatMgr) GetBatchFromClusterServiceEpoch(clusterServiceEpochs []int64) (results []*AllVirtualWeakReadStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualWeakReadStat{}).Where("`cluster_service_epoch` IN (?)", clusterServiceEpochs).Find(&results).Error

	return
}

// GetFromClusterTotalServerCount 通过cluster_total_server_count获取内容
func (obj *_AllVirtualWeakReadStatMgr) GetFromClusterTotalServerCount(clusterTotalServerCount int64) (results []*AllVirtualWeakReadStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualWeakReadStat{}).Where("`cluster_total_server_count` = ?", clusterTotalServerCount).Find(&results).Error

	return
}

// GetBatchFromClusterTotalServerCount 批量查找
func (obj *_AllVirtualWeakReadStatMgr) GetBatchFromClusterTotalServerCount(clusterTotalServerCounts []int64) (results []*AllVirtualWeakReadStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualWeakReadStat{}).Where("`cluster_total_server_count` IN (?)", clusterTotalServerCounts).Find(&results).Error

	return
}

// GetFromClusterSkippedServerCount 通过cluster_skipped_server_count获取内容
func (obj *_AllVirtualWeakReadStatMgr) GetFromClusterSkippedServerCount(clusterSkippedServerCount int64) (results []*AllVirtualWeakReadStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualWeakReadStat{}).Where("`cluster_skipped_server_count` = ?", clusterSkippedServerCount).Find(&results).Error

	return
}

// GetBatchFromClusterSkippedServerCount 批量查找
func (obj *_AllVirtualWeakReadStatMgr) GetBatchFromClusterSkippedServerCount(clusterSkippedServerCounts []int64) (results []*AllVirtualWeakReadStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualWeakReadStat{}).Where("`cluster_skipped_server_count` IN (?)", clusterSkippedServerCounts).Find(&results).Error

	return
}

// GetFromClusterVersionGenTstamp 通过cluster_version_gen_tstamp获取内容
func (obj *_AllVirtualWeakReadStatMgr) GetFromClusterVersionGenTstamp(clusterVersionGenTstamp int64) (results []*AllVirtualWeakReadStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualWeakReadStat{}).Where("`cluster_version_gen_tstamp` = ?", clusterVersionGenTstamp).Find(&results).Error

	return
}

// GetBatchFromClusterVersionGenTstamp 批量查找
func (obj *_AllVirtualWeakReadStatMgr) GetBatchFromClusterVersionGenTstamp(clusterVersionGenTstamps []int64) (results []*AllVirtualWeakReadStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualWeakReadStat{}).Where("`cluster_version_gen_tstamp` IN (?)", clusterVersionGenTstamps).Find(&results).Error

	return
}

// GetFromClusterVersion 通过cluster_version获取内容
func (obj *_AllVirtualWeakReadStatMgr) GetFromClusterVersion(clusterVersion int64) (results []*AllVirtualWeakReadStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualWeakReadStat{}).Where("`cluster_version` = ?", clusterVersion).Find(&results).Error

	return
}

// GetBatchFromClusterVersion 批量查找
func (obj *_AllVirtualWeakReadStatMgr) GetBatchFromClusterVersion(clusterVersions []int64) (results []*AllVirtualWeakReadStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualWeakReadStat{}).Where("`cluster_version` IN (?)", clusterVersions).Find(&results).Error

	return
}

// GetFromClusterVersionDelta 通过cluster_version_delta获取内容
func (obj *_AllVirtualWeakReadStatMgr) GetFromClusterVersionDelta(clusterVersionDelta int64) (results []*AllVirtualWeakReadStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualWeakReadStat{}).Where("`cluster_version_delta` = ?", clusterVersionDelta).Find(&results).Error

	return
}

// GetBatchFromClusterVersionDelta 批量查找
func (obj *_AllVirtualWeakReadStatMgr) GetBatchFromClusterVersionDelta(clusterVersionDeltas []int64) (results []*AllVirtualWeakReadStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualWeakReadStat{}).Where("`cluster_version_delta` IN (?)", clusterVersionDeltas).Find(&results).Error

	return
}

// GetFromMinClusterVersion 通过min_cluster_version获取内容
func (obj *_AllVirtualWeakReadStatMgr) GetFromMinClusterVersion(minClusterVersion int64) (results []*AllVirtualWeakReadStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualWeakReadStat{}).Where("`min_cluster_version` = ?", minClusterVersion).Find(&results).Error

	return
}

// GetBatchFromMinClusterVersion 批量查找
func (obj *_AllVirtualWeakReadStatMgr) GetBatchFromMinClusterVersion(minClusterVersions []int64) (results []*AllVirtualWeakReadStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualWeakReadStat{}).Where("`min_cluster_version` IN (?)", minClusterVersions).Find(&results).Error

	return
}

// GetFromMaxClusterVersion 通过max_cluster_version获取内容
func (obj *_AllVirtualWeakReadStatMgr) GetFromMaxClusterVersion(maxClusterVersion int64) (results []*AllVirtualWeakReadStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualWeakReadStat{}).Where("`max_cluster_version` = ?", maxClusterVersion).Find(&results).Error

	return
}

// GetBatchFromMaxClusterVersion 批量查找
func (obj *_AllVirtualWeakReadStatMgr) GetBatchFromMaxClusterVersion(maxClusterVersions []int64) (results []*AllVirtualWeakReadStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualWeakReadStat{}).Where("`max_cluster_version` IN (?)", maxClusterVersions).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////
