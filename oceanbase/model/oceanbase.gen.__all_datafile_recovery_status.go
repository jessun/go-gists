package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _AllDatafileRecoveryStatusMgr struct {
	*_BaseMgr
}

// AllDatafileRecoveryStatusMgr open func
func AllDatafileRecoveryStatusMgr(db *gorm.DB) *_AllDatafileRecoveryStatusMgr {
	if db == nil {
		panic(fmt.Errorf("AllDatafileRecoveryStatusMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllDatafileRecoveryStatusMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_datafile_recovery_status"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllDatafileRecoveryStatusMgr) GetTableName() string {
	return "__all_datafile_recovery_status"
}

// Reset 重置gorm会话
func (obj *_AllDatafileRecoveryStatusMgr) Reset() *_AllDatafileRecoveryStatusMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllDatafileRecoveryStatusMgr) Get() (result AllDatafileRecoveryStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDatafileRecoveryStatus{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllDatafileRecoveryStatusMgr) Gets() (results []*AllDatafileRecoveryStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDatafileRecoveryStatus{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllDatafileRecoveryStatusMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllDatafileRecoveryStatus{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithGmtCreate gmt_create获取
func (obj *_AllDatafileRecoveryStatusMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllDatafileRecoveryStatusMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithSvrIP svr_ip获取
func (obj *_AllDatafileRecoveryStatusMgr) WithSvrIP(svrIP string) Option {
	return optionFunc(func(o *options) { o.query["svr_ip"] = svrIP })
}

// WithSvrPort svr_port获取
func (obj *_AllDatafileRecoveryStatusMgr) WithSvrPort(svrPort int64) Option {
	return optionFunc(func(o *options) { o.query["svr_port"] = svrPort })
}

// WithTenantID tenant_id获取
func (obj *_AllDatafileRecoveryStatusMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithFileID file_id获取
func (obj *_AllDatafileRecoveryStatusMgr) WithFileID(fileID int64) Option {
	return optionFunc(func(o *options) { o.query["file_id"] = fileID })
}

// WithDestSvrIP dest_svr_ip获取
func (obj *_AllDatafileRecoveryStatusMgr) WithDestSvrIP(destSvrIP string) Option {
	return optionFunc(func(o *options) { o.query["dest_svr_ip"] = destSvrIP })
}

// WithDestSvrPort dest_svr_port获取
func (obj *_AllDatafileRecoveryStatusMgr) WithDestSvrPort(destSvrPort int64) Option {
	return optionFunc(func(o *options) { o.query["dest_svr_port"] = destSvrPort })
}

// WithDestUnitID dest_unit_id获取
func (obj *_AllDatafileRecoveryStatusMgr) WithDestUnitID(destUnitID int64) Option {
	return optionFunc(func(o *options) { o.query["dest_unit_id"] = destUnitID })
}

// WithStatus status获取
func (obj *_AllDatafileRecoveryStatusMgr) WithStatus(status int64) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// GetByOption 功能选项模式获取
func (obj *_AllDatafileRecoveryStatusMgr) GetByOption(opts ...Option) (result AllDatafileRecoveryStatus, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllDatafileRecoveryStatus{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllDatafileRecoveryStatusMgr) GetByOptions(opts ...Option) (results []*AllDatafileRecoveryStatus, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllDatafileRecoveryStatus{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllDatafileRecoveryStatusMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllDatafileRecoveryStatus, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllDatafileRecoveryStatus{}).Where(options.query)
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
func (obj *_AllDatafileRecoveryStatusMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllDatafileRecoveryStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDatafileRecoveryStatus{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllDatafileRecoveryStatusMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllDatafileRecoveryStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDatafileRecoveryStatus{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllDatafileRecoveryStatusMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllDatafileRecoveryStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDatafileRecoveryStatus{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllDatafileRecoveryStatusMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllDatafileRecoveryStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDatafileRecoveryStatus{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromSvrIP 通过svr_ip获取内容
func (obj *_AllDatafileRecoveryStatusMgr) GetFromSvrIP(svrIP string) (results []*AllDatafileRecoveryStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDatafileRecoveryStatus{}).Where("`svr_ip` = ?", svrIP).Find(&results).Error

	return
}

// GetBatchFromSvrIP 批量查找
func (obj *_AllDatafileRecoveryStatusMgr) GetBatchFromSvrIP(svrIPs []string) (results []*AllDatafileRecoveryStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDatafileRecoveryStatus{}).Where("`svr_ip` IN (?)", svrIPs).Find(&results).Error

	return
}

// GetFromSvrPort 通过svr_port获取内容
func (obj *_AllDatafileRecoveryStatusMgr) GetFromSvrPort(svrPort int64) (results []*AllDatafileRecoveryStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDatafileRecoveryStatus{}).Where("`svr_port` = ?", svrPort).Find(&results).Error

	return
}

// GetBatchFromSvrPort 批量查找
func (obj *_AllDatafileRecoveryStatusMgr) GetBatchFromSvrPort(svrPorts []int64) (results []*AllDatafileRecoveryStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDatafileRecoveryStatus{}).Where("`svr_port` IN (?)", svrPorts).Find(&results).Error

	return
}

// GetFromTenantID 通过tenant_id获取内容
func (obj *_AllDatafileRecoveryStatusMgr) GetFromTenantID(tenantID int64) (results []*AllDatafileRecoveryStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDatafileRecoveryStatus{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllDatafileRecoveryStatusMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllDatafileRecoveryStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDatafileRecoveryStatus{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromFileID 通过file_id获取内容
func (obj *_AllDatafileRecoveryStatusMgr) GetFromFileID(fileID int64) (results []*AllDatafileRecoveryStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDatafileRecoveryStatus{}).Where("`file_id` = ?", fileID).Find(&results).Error

	return
}

// GetBatchFromFileID 批量查找
func (obj *_AllDatafileRecoveryStatusMgr) GetBatchFromFileID(fileIDs []int64) (results []*AllDatafileRecoveryStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDatafileRecoveryStatus{}).Where("`file_id` IN (?)", fileIDs).Find(&results).Error

	return
}

// GetFromDestSvrIP 通过dest_svr_ip获取内容
func (obj *_AllDatafileRecoveryStatusMgr) GetFromDestSvrIP(destSvrIP string) (results []*AllDatafileRecoveryStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDatafileRecoveryStatus{}).Where("`dest_svr_ip` = ?", destSvrIP).Find(&results).Error

	return
}

// GetBatchFromDestSvrIP 批量查找
func (obj *_AllDatafileRecoveryStatusMgr) GetBatchFromDestSvrIP(destSvrIPs []string) (results []*AllDatafileRecoveryStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDatafileRecoveryStatus{}).Where("`dest_svr_ip` IN (?)", destSvrIPs).Find(&results).Error

	return
}

// GetFromDestSvrPort 通过dest_svr_port获取内容
func (obj *_AllDatafileRecoveryStatusMgr) GetFromDestSvrPort(destSvrPort int64) (results []*AllDatafileRecoveryStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDatafileRecoveryStatus{}).Where("`dest_svr_port` = ?", destSvrPort).Find(&results).Error

	return
}

// GetBatchFromDestSvrPort 批量查找
func (obj *_AllDatafileRecoveryStatusMgr) GetBatchFromDestSvrPort(destSvrPorts []int64) (results []*AllDatafileRecoveryStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDatafileRecoveryStatus{}).Where("`dest_svr_port` IN (?)", destSvrPorts).Find(&results).Error

	return
}

// GetFromDestUnitID 通过dest_unit_id获取内容
func (obj *_AllDatafileRecoveryStatusMgr) GetFromDestUnitID(destUnitID int64) (results []*AllDatafileRecoveryStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDatafileRecoveryStatus{}).Where("`dest_unit_id` = ?", destUnitID).Find(&results).Error

	return
}

// GetBatchFromDestUnitID 批量查找
func (obj *_AllDatafileRecoveryStatusMgr) GetBatchFromDestUnitID(destUnitIDs []int64) (results []*AllDatafileRecoveryStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDatafileRecoveryStatus{}).Where("`dest_unit_id` IN (?)", destUnitIDs).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容
func (obj *_AllDatafileRecoveryStatusMgr) GetFromStatus(status int64) (results []*AllDatafileRecoveryStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDatafileRecoveryStatus{}).Where("`status` = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量查找
func (obj *_AllDatafileRecoveryStatusMgr) GetBatchFromStatus(statuss []int64) (results []*AllDatafileRecoveryStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDatafileRecoveryStatus{}).Where("`status` IN (?)", statuss).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllDatafileRecoveryStatusMgr) FetchByPrimaryKey(svrIP string, svrPort int64, tenantID int64, fileID int64) (result AllDatafileRecoveryStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDatafileRecoveryStatus{}).Where("`svr_ip` = ? AND `svr_port` = ? AND `tenant_id` = ? AND `file_id` = ?", svrIP, svrPort, tenantID, fileID).First(&result).Error

	return
}
