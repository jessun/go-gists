package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _AllVirtualProxyServerStatMgr struct {
	*_BaseMgr
}

// AllVirtualProxyServerStatMgr open func
func AllVirtualProxyServerStatMgr(db *gorm.DB) *_AllVirtualProxyServerStatMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualProxyServerStatMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualProxyServerStatMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_proxy_server_stat"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualProxyServerStatMgr) GetTableName() string {
	return "__all_virtual_proxy_server_stat"
}

// Reset 重置gorm会话
func (obj *_AllVirtualProxyServerStatMgr) Reset() *_AllVirtualProxyServerStatMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualProxyServerStatMgr) Get() (result AllVirtualProxyServerStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxyServerStat{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualProxyServerStatMgr) Gets() (results []*AllVirtualProxyServerStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxyServerStat{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualProxyServerStatMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualProxyServerStat{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithSvrIP svr_ip获取
func (obj *_AllVirtualProxyServerStatMgr) WithSvrIP(svrIP string) Option {
	return optionFunc(func(o *options) { o.query["svr_ip"] = svrIP })
}

// WithSvrPort svr_port获取
func (obj *_AllVirtualProxyServerStatMgr) WithSvrPort(svrPort int64) Option {
	return optionFunc(func(o *options) { o.query["svr_port"] = svrPort })
}

// WithZone zone获取
func (obj *_AllVirtualProxyServerStatMgr) WithZone(zone string) Option {
	return optionFunc(func(o *options) { o.query["zone"] = zone })
}

// WithStartServiceTime start_service_time获取
func (obj *_AllVirtualProxyServerStatMgr) WithStartServiceTime(startServiceTime int64) Option {
	return optionFunc(func(o *options) { o.query["start_service_time"] = startServiceTime })
}

// WithStopTime stop_time获取
func (obj *_AllVirtualProxyServerStatMgr) WithStopTime(stopTime int64) Option {
	return optionFunc(func(o *options) { o.query["stop_time"] = stopTime })
}

// WithStatus status获取
func (obj *_AllVirtualProxyServerStatMgr) WithStatus(status string) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualProxyServerStatMgr) GetByOption(opts ...Option) (result AllVirtualProxyServerStat, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxyServerStat{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualProxyServerStatMgr) GetByOptions(opts ...Option) (results []*AllVirtualProxyServerStat, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxyServerStat{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualProxyServerStatMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualProxyServerStat, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualProxyServerStat{}).Where(options.query)
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
func (obj *_AllVirtualProxyServerStatMgr) GetFromSvrIP(svrIP string) (results []*AllVirtualProxyServerStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxyServerStat{}).Where("`svr_ip` = ?", svrIP).Find(&results).Error

	return
}

// GetBatchFromSvrIP 批量查找
func (obj *_AllVirtualProxyServerStatMgr) GetBatchFromSvrIP(svrIPs []string) (results []*AllVirtualProxyServerStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxyServerStat{}).Where("`svr_ip` IN (?)", svrIPs).Find(&results).Error

	return
}

// GetFromSvrPort 通过svr_port获取内容
func (obj *_AllVirtualProxyServerStatMgr) GetFromSvrPort(svrPort int64) (results []*AllVirtualProxyServerStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxyServerStat{}).Where("`svr_port` = ?", svrPort).Find(&results).Error

	return
}

// GetBatchFromSvrPort 批量查找
func (obj *_AllVirtualProxyServerStatMgr) GetBatchFromSvrPort(svrPorts []int64) (results []*AllVirtualProxyServerStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxyServerStat{}).Where("`svr_port` IN (?)", svrPorts).Find(&results).Error

	return
}

// GetFromZone 通过zone获取内容
func (obj *_AllVirtualProxyServerStatMgr) GetFromZone(zone string) (results []*AllVirtualProxyServerStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxyServerStat{}).Where("`zone` = ?", zone).Find(&results).Error

	return
}

// GetBatchFromZone 批量查找
func (obj *_AllVirtualProxyServerStatMgr) GetBatchFromZone(zones []string) (results []*AllVirtualProxyServerStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxyServerStat{}).Where("`zone` IN (?)", zones).Find(&results).Error

	return
}

// GetFromStartServiceTime 通过start_service_time获取内容
func (obj *_AllVirtualProxyServerStatMgr) GetFromStartServiceTime(startServiceTime int64) (results []*AllVirtualProxyServerStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxyServerStat{}).Where("`start_service_time` = ?", startServiceTime).Find(&results).Error

	return
}

// GetBatchFromStartServiceTime 批量查找
func (obj *_AllVirtualProxyServerStatMgr) GetBatchFromStartServiceTime(startServiceTimes []int64) (results []*AllVirtualProxyServerStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxyServerStat{}).Where("`start_service_time` IN (?)", startServiceTimes).Find(&results).Error

	return
}

// GetFromStopTime 通过stop_time获取内容
func (obj *_AllVirtualProxyServerStatMgr) GetFromStopTime(stopTime int64) (results []*AllVirtualProxyServerStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxyServerStat{}).Where("`stop_time` = ?", stopTime).Find(&results).Error

	return
}

// GetBatchFromStopTime 批量查找
func (obj *_AllVirtualProxyServerStatMgr) GetBatchFromStopTime(stopTimes []int64) (results []*AllVirtualProxyServerStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxyServerStat{}).Where("`stop_time` IN (?)", stopTimes).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容
func (obj *_AllVirtualProxyServerStatMgr) GetFromStatus(status string) (results []*AllVirtualProxyServerStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxyServerStat{}).Where("`status` = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量查找
func (obj *_AllVirtualProxyServerStatMgr) GetBatchFromStatus(statuss []string) (results []*AllVirtualProxyServerStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxyServerStat{}).Where("`status` IN (?)", statuss).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////
