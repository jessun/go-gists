package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _AllServerRecoveryStatusMgr struct {
	*_BaseMgr
}

// AllServerRecoveryStatusMgr open func
func AllServerRecoveryStatusMgr(db *gorm.DB) *_AllServerRecoveryStatusMgr {
	if db == nil {
		panic(fmt.Errorf("AllServerRecoveryStatusMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllServerRecoveryStatusMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_server_recovery_status"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllServerRecoveryStatusMgr) GetTableName() string {
	return "__all_server_recovery_status"
}

// Reset 重置gorm会话
func (obj *_AllServerRecoveryStatusMgr) Reset() *_AllServerRecoveryStatusMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllServerRecoveryStatusMgr) Get() (result AllServerRecoveryStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllServerRecoveryStatus{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllServerRecoveryStatusMgr) Gets() (results []*AllServerRecoveryStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllServerRecoveryStatus{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllServerRecoveryStatusMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllServerRecoveryStatus{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithGmtCreate gmt_create获取
func (obj *_AllServerRecoveryStatusMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllServerRecoveryStatusMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithSvrIP svr_ip获取
func (obj *_AllServerRecoveryStatusMgr) WithSvrIP(svrIP string) Option {
	return optionFunc(func(o *options) { o.query["svr_ip"] = svrIP })
}

// WithSvrPort svr_port获取
func (obj *_AllServerRecoveryStatusMgr) WithSvrPort(svrPort int64) Option {
	return optionFunc(func(o *options) { o.query["svr_port"] = svrPort })
}

// WithRescueSvrIP rescue_svr_ip获取
func (obj *_AllServerRecoveryStatusMgr) WithRescueSvrIP(rescueSvrIP string) Option {
	return optionFunc(func(o *options) { o.query["rescue_svr_ip"] = rescueSvrIP })
}

// WithRescueSvrPort rescue_svr_port获取
func (obj *_AllServerRecoveryStatusMgr) WithRescueSvrPort(rescueSvrPort int64) Option {
	return optionFunc(func(o *options) { o.query["rescue_svr_port"] = rescueSvrPort })
}

// WithRescueProgress rescue_progress获取
func (obj *_AllServerRecoveryStatusMgr) WithRescueProgress(rescueProgress int64) Option {
	return optionFunc(func(o *options) { o.query["rescue_progress"] = rescueProgress })
}

// GetByOption 功能选项模式获取
func (obj *_AllServerRecoveryStatusMgr) GetByOption(opts ...Option) (result AllServerRecoveryStatus, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllServerRecoveryStatus{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllServerRecoveryStatusMgr) GetByOptions(opts ...Option) (results []*AllServerRecoveryStatus, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllServerRecoveryStatus{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllServerRecoveryStatusMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllServerRecoveryStatus, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllServerRecoveryStatus{}).Where(options.query)
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
func (obj *_AllServerRecoveryStatusMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllServerRecoveryStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllServerRecoveryStatus{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllServerRecoveryStatusMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllServerRecoveryStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllServerRecoveryStatus{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllServerRecoveryStatusMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllServerRecoveryStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllServerRecoveryStatus{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllServerRecoveryStatusMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllServerRecoveryStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllServerRecoveryStatus{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromSvrIP 通过svr_ip获取内容
func (obj *_AllServerRecoveryStatusMgr) GetFromSvrIP(svrIP string) (results []*AllServerRecoveryStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllServerRecoveryStatus{}).Where("`svr_ip` = ?", svrIP).Find(&results).Error

	return
}

// GetBatchFromSvrIP 批量查找
func (obj *_AllServerRecoveryStatusMgr) GetBatchFromSvrIP(svrIPs []string) (results []*AllServerRecoveryStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllServerRecoveryStatus{}).Where("`svr_ip` IN (?)", svrIPs).Find(&results).Error

	return
}

// GetFromSvrPort 通过svr_port获取内容
func (obj *_AllServerRecoveryStatusMgr) GetFromSvrPort(svrPort int64) (results []*AllServerRecoveryStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllServerRecoveryStatus{}).Where("`svr_port` = ?", svrPort).Find(&results).Error

	return
}

// GetBatchFromSvrPort 批量查找
func (obj *_AllServerRecoveryStatusMgr) GetBatchFromSvrPort(svrPorts []int64) (results []*AllServerRecoveryStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllServerRecoveryStatus{}).Where("`svr_port` IN (?)", svrPorts).Find(&results).Error

	return
}

// GetFromRescueSvrIP 通过rescue_svr_ip获取内容
func (obj *_AllServerRecoveryStatusMgr) GetFromRescueSvrIP(rescueSvrIP string) (results []*AllServerRecoveryStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllServerRecoveryStatus{}).Where("`rescue_svr_ip` = ?", rescueSvrIP).Find(&results).Error

	return
}

// GetBatchFromRescueSvrIP 批量查找
func (obj *_AllServerRecoveryStatusMgr) GetBatchFromRescueSvrIP(rescueSvrIPs []string) (results []*AllServerRecoveryStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllServerRecoveryStatus{}).Where("`rescue_svr_ip` IN (?)", rescueSvrIPs).Find(&results).Error

	return
}

// GetFromRescueSvrPort 通过rescue_svr_port获取内容
func (obj *_AllServerRecoveryStatusMgr) GetFromRescueSvrPort(rescueSvrPort int64) (results []*AllServerRecoveryStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllServerRecoveryStatus{}).Where("`rescue_svr_port` = ?", rescueSvrPort).Find(&results).Error

	return
}

// GetBatchFromRescueSvrPort 批量查找
func (obj *_AllServerRecoveryStatusMgr) GetBatchFromRescueSvrPort(rescueSvrPorts []int64) (results []*AllServerRecoveryStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllServerRecoveryStatus{}).Where("`rescue_svr_port` IN (?)", rescueSvrPorts).Find(&results).Error

	return
}

// GetFromRescueProgress 通过rescue_progress获取内容
func (obj *_AllServerRecoveryStatusMgr) GetFromRescueProgress(rescueProgress int64) (results []*AllServerRecoveryStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllServerRecoveryStatus{}).Where("`rescue_progress` = ?", rescueProgress).Find(&results).Error

	return
}

// GetBatchFromRescueProgress 批量查找
func (obj *_AllServerRecoveryStatusMgr) GetBatchFromRescueProgress(rescueProgresss []int64) (results []*AllServerRecoveryStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllServerRecoveryStatus{}).Where("`rescue_progress` IN (?)", rescueProgresss).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllServerRecoveryStatusMgr) FetchByPrimaryKey(svrIP string, svrPort int64) (result AllServerRecoveryStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllServerRecoveryStatus{}).Where("`svr_ip` = ? AND `svr_port` = ?", svrIP, svrPort).First(&result).Error

	return
}
