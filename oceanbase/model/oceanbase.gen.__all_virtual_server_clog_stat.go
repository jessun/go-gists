package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _AllVirtualServerClogStatMgr struct {
	*_BaseMgr
}

// AllVirtualServerClogStatMgr open func
func AllVirtualServerClogStatMgr(db *gorm.DB) *_AllVirtualServerClogStatMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualServerClogStatMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualServerClogStatMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_server_clog_stat"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualServerClogStatMgr) GetTableName() string {
	return "__all_virtual_server_clog_stat"
}

// Reset 重置gorm会话
func (obj *_AllVirtualServerClogStatMgr) Reset() *_AllVirtualServerClogStatMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualServerClogStatMgr) Get() (result AllVirtualServerClogStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerClogStat{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualServerClogStatMgr) Gets() (results []*AllVirtualServerClogStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerClogStat{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualServerClogStatMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualServerClogStat{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithSvrIP svr_ip获取
func (obj *_AllVirtualServerClogStatMgr) WithSvrIP(svrIP string) Option {
	return optionFunc(func(o *options) { o.query["svr_ip"] = svrIP })
}

// WithSvrPort svr_port获取
func (obj *_AllVirtualServerClogStatMgr) WithSvrPort(svrPort int64) Option {
	return optionFunc(func(o *options) { o.query["svr_port"] = svrPort })
}

// WithSystemClogMinUsingFileID system_clog_min_using_file_id获取
func (obj *_AllVirtualServerClogStatMgr) WithSystemClogMinUsingFileID(systemClogMinUsingFileID int64) Option {
	return optionFunc(func(o *options) { o.query["system_clog_min_using_file_id"] = systemClogMinUsingFileID })
}

// WithUserClogMinUsingFileID user_clog_min_using_file_id获取
func (obj *_AllVirtualServerClogStatMgr) WithUserClogMinUsingFileID(userClogMinUsingFileID int64) Option {
	return optionFunc(func(o *options) { o.query["user_clog_min_using_file_id"] = userClogMinUsingFileID })
}

// WithSystemIlogMinUsingFileID system_ilog_min_using_file_id获取
func (obj *_AllVirtualServerClogStatMgr) WithSystemIlogMinUsingFileID(systemIlogMinUsingFileID int64) Option {
	return optionFunc(func(o *options) { o.query["system_ilog_min_using_file_id"] = systemIlogMinUsingFileID })
}

// WithUserIlogMinUsingFileID user_ilog_min_using_file_id获取
func (obj *_AllVirtualServerClogStatMgr) WithUserIlogMinUsingFileID(userIlogMinUsingFileID int64) Option {
	return optionFunc(func(o *options) { o.query["user_ilog_min_using_file_id"] = userIlogMinUsingFileID })
}

// WithZone zone获取
func (obj *_AllVirtualServerClogStatMgr) WithZone(zone string) Option {
	return optionFunc(func(o *options) { o.query["zone"] = zone })
}

// WithRegion region获取
func (obj *_AllVirtualServerClogStatMgr) WithRegion(region string) Option {
	return optionFunc(func(o *options) { o.query["region"] = region })
}

// WithIDc idc获取
func (obj *_AllVirtualServerClogStatMgr) WithIDc(idc string) Option {
	return optionFunc(func(o *options) { o.query["idc"] = idc })
}

// WithZoneType zone_type获取
func (obj *_AllVirtualServerClogStatMgr) WithZoneType(zoneType string) Option {
	return optionFunc(func(o *options) { o.query["zone_type"] = zoneType })
}

// WithMergeStatus merge_status获取
func (obj *_AllVirtualServerClogStatMgr) WithMergeStatus(mergeStatus string) Option {
	return optionFunc(func(o *options) { o.query["merge_status"] = mergeStatus })
}

// WithZoneStatus zone_status获取
func (obj *_AllVirtualServerClogStatMgr) WithZoneStatus(zoneStatus string) Option {
	return optionFunc(func(o *options) { o.query["zone_status"] = zoneStatus })
}

// WithSvrMinLogTimestamp svr_min_log_timestamp获取
func (obj *_AllVirtualServerClogStatMgr) WithSvrMinLogTimestamp(svrMinLogTimestamp int64) Option {
	return optionFunc(func(o *options) { o.query["svr_min_log_timestamp"] = svrMinLogTimestamp })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualServerClogStatMgr) GetByOption(opts ...Option) (result AllVirtualServerClogStat, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerClogStat{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualServerClogStatMgr) GetByOptions(opts ...Option) (results []*AllVirtualServerClogStat, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerClogStat{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualServerClogStatMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualServerClogStat, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualServerClogStat{}).Where(options.query)
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
func (obj *_AllVirtualServerClogStatMgr) GetFromSvrIP(svrIP string) (results []*AllVirtualServerClogStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerClogStat{}).Where("`svr_ip` = ?", svrIP).Find(&results).Error

	return
}

// GetBatchFromSvrIP 批量查找
func (obj *_AllVirtualServerClogStatMgr) GetBatchFromSvrIP(svrIPs []string) (results []*AllVirtualServerClogStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerClogStat{}).Where("`svr_ip` IN (?)", svrIPs).Find(&results).Error

	return
}

// GetFromSvrPort 通过svr_port获取内容
func (obj *_AllVirtualServerClogStatMgr) GetFromSvrPort(svrPort int64) (results []*AllVirtualServerClogStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerClogStat{}).Where("`svr_port` = ?", svrPort).Find(&results).Error

	return
}

// GetBatchFromSvrPort 批量查找
func (obj *_AllVirtualServerClogStatMgr) GetBatchFromSvrPort(svrPorts []int64) (results []*AllVirtualServerClogStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerClogStat{}).Where("`svr_port` IN (?)", svrPorts).Find(&results).Error

	return
}

// GetFromSystemClogMinUsingFileID 通过system_clog_min_using_file_id获取内容
func (obj *_AllVirtualServerClogStatMgr) GetFromSystemClogMinUsingFileID(systemClogMinUsingFileID int64) (results []*AllVirtualServerClogStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerClogStat{}).Where("`system_clog_min_using_file_id` = ?", systemClogMinUsingFileID).Find(&results).Error

	return
}

// GetBatchFromSystemClogMinUsingFileID 批量查找
func (obj *_AllVirtualServerClogStatMgr) GetBatchFromSystemClogMinUsingFileID(systemClogMinUsingFileIDs []int64) (results []*AllVirtualServerClogStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerClogStat{}).Where("`system_clog_min_using_file_id` IN (?)", systemClogMinUsingFileIDs).Find(&results).Error

	return
}

// GetFromUserClogMinUsingFileID 通过user_clog_min_using_file_id获取内容
func (obj *_AllVirtualServerClogStatMgr) GetFromUserClogMinUsingFileID(userClogMinUsingFileID int64) (results []*AllVirtualServerClogStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerClogStat{}).Where("`user_clog_min_using_file_id` = ?", userClogMinUsingFileID).Find(&results).Error

	return
}

// GetBatchFromUserClogMinUsingFileID 批量查找
func (obj *_AllVirtualServerClogStatMgr) GetBatchFromUserClogMinUsingFileID(userClogMinUsingFileIDs []int64) (results []*AllVirtualServerClogStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerClogStat{}).Where("`user_clog_min_using_file_id` IN (?)", userClogMinUsingFileIDs).Find(&results).Error

	return
}

// GetFromSystemIlogMinUsingFileID 通过system_ilog_min_using_file_id获取内容
func (obj *_AllVirtualServerClogStatMgr) GetFromSystemIlogMinUsingFileID(systemIlogMinUsingFileID int64) (results []*AllVirtualServerClogStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerClogStat{}).Where("`system_ilog_min_using_file_id` = ?", systemIlogMinUsingFileID).Find(&results).Error

	return
}

// GetBatchFromSystemIlogMinUsingFileID 批量查找
func (obj *_AllVirtualServerClogStatMgr) GetBatchFromSystemIlogMinUsingFileID(systemIlogMinUsingFileIDs []int64) (results []*AllVirtualServerClogStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerClogStat{}).Where("`system_ilog_min_using_file_id` IN (?)", systemIlogMinUsingFileIDs).Find(&results).Error

	return
}

// GetFromUserIlogMinUsingFileID 通过user_ilog_min_using_file_id获取内容
func (obj *_AllVirtualServerClogStatMgr) GetFromUserIlogMinUsingFileID(userIlogMinUsingFileID int64) (results []*AllVirtualServerClogStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerClogStat{}).Where("`user_ilog_min_using_file_id` = ?", userIlogMinUsingFileID).Find(&results).Error

	return
}

// GetBatchFromUserIlogMinUsingFileID 批量查找
func (obj *_AllVirtualServerClogStatMgr) GetBatchFromUserIlogMinUsingFileID(userIlogMinUsingFileIDs []int64) (results []*AllVirtualServerClogStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerClogStat{}).Where("`user_ilog_min_using_file_id` IN (?)", userIlogMinUsingFileIDs).Find(&results).Error

	return
}

// GetFromZone 通过zone获取内容
func (obj *_AllVirtualServerClogStatMgr) GetFromZone(zone string) (results []*AllVirtualServerClogStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerClogStat{}).Where("`zone` = ?", zone).Find(&results).Error

	return
}

// GetBatchFromZone 批量查找
func (obj *_AllVirtualServerClogStatMgr) GetBatchFromZone(zones []string) (results []*AllVirtualServerClogStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerClogStat{}).Where("`zone` IN (?)", zones).Find(&results).Error

	return
}

// GetFromRegion 通过region获取内容
func (obj *_AllVirtualServerClogStatMgr) GetFromRegion(region string) (results []*AllVirtualServerClogStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerClogStat{}).Where("`region` = ?", region).Find(&results).Error

	return
}

// GetBatchFromRegion 批量查找
func (obj *_AllVirtualServerClogStatMgr) GetBatchFromRegion(regions []string) (results []*AllVirtualServerClogStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerClogStat{}).Where("`region` IN (?)", regions).Find(&results).Error

	return
}

// GetFromIDc 通过idc获取内容
func (obj *_AllVirtualServerClogStatMgr) GetFromIDc(idc string) (results []*AllVirtualServerClogStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerClogStat{}).Where("`idc` = ?", idc).Find(&results).Error

	return
}

// GetBatchFromIDc 批量查找
func (obj *_AllVirtualServerClogStatMgr) GetBatchFromIDc(idcs []string) (results []*AllVirtualServerClogStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerClogStat{}).Where("`idc` IN (?)", idcs).Find(&results).Error

	return
}

// GetFromZoneType 通过zone_type获取内容
func (obj *_AllVirtualServerClogStatMgr) GetFromZoneType(zoneType string) (results []*AllVirtualServerClogStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerClogStat{}).Where("`zone_type` = ?", zoneType).Find(&results).Error

	return
}

// GetBatchFromZoneType 批量查找
func (obj *_AllVirtualServerClogStatMgr) GetBatchFromZoneType(zoneTypes []string) (results []*AllVirtualServerClogStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerClogStat{}).Where("`zone_type` IN (?)", zoneTypes).Find(&results).Error

	return
}

// GetFromMergeStatus 通过merge_status获取内容
func (obj *_AllVirtualServerClogStatMgr) GetFromMergeStatus(mergeStatus string) (results []*AllVirtualServerClogStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerClogStat{}).Where("`merge_status` = ?", mergeStatus).Find(&results).Error

	return
}

// GetBatchFromMergeStatus 批量查找
func (obj *_AllVirtualServerClogStatMgr) GetBatchFromMergeStatus(mergeStatuss []string) (results []*AllVirtualServerClogStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerClogStat{}).Where("`merge_status` IN (?)", mergeStatuss).Find(&results).Error

	return
}

// GetFromZoneStatus 通过zone_status获取内容
func (obj *_AllVirtualServerClogStatMgr) GetFromZoneStatus(zoneStatus string) (results []*AllVirtualServerClogStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerClogStat{}).Where("`zone_status` = ?", zoneStatus).Find(&results).Error

	return
}

// GetBatchFromZoneStatus 批量查找
func (obj *_AllVirtualServerClogStatMgr) GetBatchFromZoneStatus(zoneStatuss []string) (results []*AllVirtualServerClogStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerClogStat{}).Where("`zone_status` IN (?)", zoneStatuss).Find(&results).Error

	return
}

// GetFromSvrMinLogTimestamp 通过svr_min_log_timestamp获取内容
func (obj *_AllVirtualServerClogStatMgr) GetFromSvrMinLogTimestamp(svrMinLogTimestamp int64) (results []*AllVirtualServerClogStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerClogStat{}).Where("`svr_min_log_timestamp` = ?", svrMinLogTimestamp).Find(&results).Error

	return
}

// GetBatchFromSvrMinLogTimestamp 批量查找
func (obj *_AllVirtualServerClogStatMgr) GetBatchFromSvrMinLogTimestamp(svrMinLogTimestamps []int64) (results []*AllVirtualServerClogStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerClogStat{}).Where("`svr_min_log_timestamp` IN (?)", svrMinLogTimestamps).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllVirtualServerClogStatMgr) FetchByPrimaryKey(svrIP string, svrPort int64) (result AllVirtualServerClogStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerClogStat{}).Where("`svr_ip` = ? AND `svr_port` = ?", svrIP, svrPort).First(&result).Error

	return
}
