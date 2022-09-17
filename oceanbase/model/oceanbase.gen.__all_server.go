package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _AllServerMgr struct {
	*_BaseMgr
}

// AllServerMgr open func
func AllServerMgr(db *gorm.DB) *_AllServerMgr {
	if db == nil {
		panic(fmt.Errorf("AllServerMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllServerMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_server"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllServerMgr) GetTableName() string {
	return "__all_server"
}

// Reset 重置gorm会话
func (obj *_AllServerMgr) Reset() *_AllServerMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllServerMgr) Get() (result AllServer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllServer{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllServerMgr) Gets() (results []*AllServer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllServer{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllServerMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllServer{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithGmtCreate gmt_create获取
func (obj *_AllServerMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllServerMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithSvrIP svr_ip获取
func (obj *_AllServerMgr) WithSvrIP(svrIP string) Option {
	return optionFunc(func(o *options) { o.query["svr_ip"] = svrIP })
}

// WithSvrPort svr_port获取
func (obj *_AllServerMgr) WithSvrPort(svrPort int64) Option {
	return optionFunc(func(o *options) { o.query["svr_port"] = svrPort })
}

// WithID id获取
func (obj *_AllServerMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithZone zone获取
func (obj *_AllServerMgr) WithZone(zone string) Option {
	return optionFunc(func(o *options) { o.query["zone"] = zone })
}

// WithInnerPort inner_port获取
func (obj *_AllServerMgr) WithInnerPort(innerPort int64) Option {
	return optionFunc(func(o *options) { o.query["inner_port"] = innerPort })
}

// WithWithRootserver with_rootserver获取
func (obj *_AllServerMgr) WithWithRootserver(withRootserver int64) Option {
	return optionFunc(func(o *options) { o.query["with_rootserver"] = withRootserver })
}

// WithStatus status获取
func (obj *_AllServerMgr) WithStatus(status string) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithBlockMigrateInTime block_migrate_in_time获取
func (obj *_AllServerMgr) WithBlockMigrateInTime(blockMigrateInTime int64) Option {
	return optionFunc(func(o *options) { o.query["block_migrate_in_time"] = blockMigrateInTime })
}

// WithBuildVersion build_version获取
func (obj *_AllServerMgr) WithBuildVersion(buildVersion string) Option {
	return optionFunc(func(o *options) { o.query["build_version"] = buildVersion })
}

// WithStopTime stop_time获取
func (obj *_AllServerMgr) WithStopTime(stopTime int64) Option {
	return optionFunc(func(o *options) { o.query["stop_time"] = stopTime })
}

// WithStartServiceTime start_service_time获取
func (obj *_AllServerMgr) WithStartServiceTime(startServiceTime int64) Option {
	return optionFunc(func(o *options) { o.query["start_service_time"] = startServiceTime })
}

// WithFirstSessid first_sessid获取
func (obj *_AllServerMgr) WithFirstSessid(firstSessid int64) Option {
	return optionFunc(func(o *options) { o.query["first_sessid"] = firstSessid })
}

// WithWithPartition with_partition获取
func (obj *_AllServerMgr) WithWithPartition(withPartition int64) Option {
	return optionFunc(func(o *options) { o.query["with_partition"] = withPartition })
}

// WithLastOfflineTime last_offline_time获取
func (obj *_AllServerMgr) WithLastOfflineTime(lastOfflineTime int64) Option {
	return optionFunc(func(o *options) { o.query["last_offline_time"] = lastOfflineTime })
}

// GetByOption 功能选项模式获取
func (obj *_AllServerMgr) GetByOption(opts ...Option) (result AllServer, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllServer{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllServerMgr) GetByOptions(opts ...Option) (results []*AllServer, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllServer{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllServerMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllServer, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllServer{}).Where(options.query)
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

// GetFromGmtCreate 通过gmt_create获取内容
func (obj *_AllServerMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllServer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllServer{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllServerMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllServer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllServer{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllServerMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllServer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllServer{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllServerMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllServer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllServer{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromSvrIP 通过svr_ip获取内容
func (obj *_AllServerMgr) GetFromSvrIP(svrIP string) (results []*AllServer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllServer{}).Where("`svr_ip` = ?", svrIP).Find(&results).Error

	return
}

// GetBatchFromSvrIP 批量查找
func (obj *_AllServerMgr) GetBatchFromSvrIP(svrIPs []string) (results []*AllServer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllServer{}).Where("`svr_ip` IN (?)", svrIPs).Find(&results).Error

	return
}

// GetFromSvrPort 通过svr_port获取内容
func (obj *_AllServerMgr) GetFromSvrPort(svrPort int64) (results []*AllServer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllServer{}).Where("`svr_port` = ?", svrPort).Find(&results).Error

	return
}

// GetBatchFromSvrPort 批量查找
func (obj *_AllServerMgr) GetBatchFromSvrPort(svrPorts []int64) (results []*AllServer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllServer{}).Where("`svr_port` IN (?)", svrPorts).Find(&results).Error

	return
}

// GetFromID 通过id获取内容
func (obj *_AllServerMgr) GetFromID(id int64) (results []*AllServer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllServer{}).Where("`id` = ?", id).Find(&results).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_AllServerMgr) GetBatchFromID(ids []int64) (results []*AllServer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllServer{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromZone 通过zone获取内容
func (obj *_AllServerMgr) GetFromZone(zone string) (results []*AllServer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllServer{}).Where("`zone` = ?", zone).Find(&results).Error

	return
}

// GetBatchFromZone 批量查找
func (obj *_AllServerMgr) GetBatchFromZone(zones []string) (results []*AllServer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllServer{}).Where("`zone` IN (?)", zones).Find(&results).Error

	return
}

// GetFromInnerPort 通过inner_port获取内容
func (obj *_AllServerMgr) GetFromInnerPort(innerPort int64) (results []*AllServer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllServer{}).Where("`inner_port` = ?", innerPort).Find(&results).Error

	return
}

// GetBatchFromInnerPort 批量查找
func (obj *_AllServerMgr) GetBatchFromInnerPort(innerPorts []int64) (results []*AllServer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllServer{}).Where("`inner_port` IN (?)", innerPorts).Find(&results).Error

	return
}

// GetFromWithRootserver 通过with_rootserver获取内容
func (obj *_AllServerMgr) GetFromWithRootserver(withRootserver int64) (results []*AllServer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllServer{}).Where("`with_rootserver` = ?", withRootserver).Find(&results).Error

	return
}

// GetBatchFromWithRootserver 批量查找
func (obj *_AllServerMgr) GetBatchFromWithRootserver(withRootservers []int64) (results []*AllServer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllServer{}).Where("`with_rootserver` IN (?)", withRootservers).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容
func (obj *_AllServerMgr) GetFromStatus(status string) (results []*AllServer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllServer{}).Where("`status` = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量查找
func (obj *_AllServerMgr) GetBatchFromStatus(statuss []string) (results []*AllServer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllServer{}).Where("`status` IN (?)", statuss).Find(&results).Error

	return
}

// GetFromBlockMigrateInTime 通过block_migrate_in_time获取内容
func (obj *_AllServerMgr) GetFromBlockMigrateInTime(blockMigrateInTime int64) (results []*AllServer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllServer{}).Where("`block_migrate_in_time` = ?", blockMigrateInTime).Find(&results).Error

	return
}

// GetBatchFromBlockMigrateInTime 批量查找
func (obj *_AllServerMgr) GetBatchFromBlockMigrateInTime(blockMigrateInTimes []int64) (results []*AllServer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllServer{}).Where("`block_migrate_in_time` IN (?)", blockMigrateInTimes).Find(&results).Error

	return
}

// GetFromBuildVersion 通过build_version获取内容
func (obj *_AllServerMgr) GetFromBuildVersion(buildVersion string) (results []*AllServer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllServer{}).Where("`build_version` = ?", buildVersion).Find(&results).Error

	return
}

// GetBatchFromBuildVersion 批量查找
func (obj *_AllServerMgr) GetBatchFromBuildVersion(buildVersions []string) (results []*AllServer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllServer{}).Where("`build_version` IN (?)", buildVersions).Find(&results).Error

	return
}

// GetFromStopTime 通过stop_time获取内容
func (obj *_AllServerMgr) GetFromStopTime(stopTime int64) (results []*AllServer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllServer{}).Where("`stop_time` = ?", stopTime).Find(&results).Error

	return
}

// GetBatchFromStopTime 批量查找
func (obj *_AllServerMgr) GetBatchFromStopTime(stopTimes []int64) (results []*AllServer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllServer{}).Where("`stop_time` IN (?)", stopTimes).Find(&results).Error

	return
}

// GetFromStartServiceTime 通过start_service_time获取内容
func (obj *_AllServerMgr) GetFromStartServiceTime(startServiceTime int64) (results []*AllServer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllServer{}).Where("`start_service_time` = ?", startServiceTime).Find(&results).Error

	return
}

// GetBatchFromStartServiceTime 批量查找
func (obj *_AllServerMgr) GetBatchFromStartServiceTime(startServiceTimes []int64) (results []*AllServer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllServer{}).Where("`start_service_time` IN (?)", startServiceTimes).Find(&results).Error

	return
}

// GetFromFirstSessid 通过first_sessid获取内容
func (obj *_AllServerMgr) GetFromFirstSessid(firstSessid int64) (results []*AllServer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllServer{}).Where("`first_sessid` = ?", firstSessid).Find(&results).Error

	return
}

// GetBatchFromFirstSessid 批量查找
func (obj *_AllServerMgr) GetBatchFromFirstSessid(firstSessids []int64) (results []*AllServer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllServer{}).Where("`first_sessid` IN (?)", firstSessids).Find(&results).Error

	return
}

// GetFromWithPartition 通过with_partition获取内容
func (obj *_AllServerMgr) GetFromWithPartition(withPartition int64) (results []*AllServer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllServer{}).Where("`with_partition` = ?", withPartition).Find(&results).Error

	return
}

// GetBatchFromWithPartition 批量查找
func (obj *_AllServerMgr) GetBatchFromWithPartition(withPartitions []int64) (results []*AllServer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllServer{}).Where("`with_partition` IN (?)", withPartitions).Find(&results).Error

	return
}

// GetFromLastOfflineTime 通过last_offline_time获取内容
func (obj *_AllServerMgr) GetFromLastOfflineTime(lastOfflineTime int64) (results []*AllServer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllServer{}).Where("`last_offline_time` = ?", lastOfflineTime).Find(&results).Error

	return
}

// GetBatchFromLastOfflineTime 批量查找
func (obj *_AllServerMgr) GetBatchFromLastOfflineTime(lastOfflineTimes []int64) (results []*AllServer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllServer{}).Where("`last_offline_time` IN (?)", lastOfflineTimes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllServerMgr) FetchByPrimaryKey(svrIP string, svrPort int64) (result AllServer, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllServer{}).Where("`svr_ip` = ? AND `svr_port` = ?", svrIP, svrPort).First(&result).Error

	return
}
