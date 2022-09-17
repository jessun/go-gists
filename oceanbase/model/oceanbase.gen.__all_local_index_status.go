package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _AllLocalIndexStatusMgr struct {
	*_BaseMgr
}

// AllLocalIndexStatusMgr open func
func AllLocalIndexStatusMgr(db *gorm.DB) *_AllLocalIndexStatusMgr {
	if db == nil {
		panic(fmt.Errorf("AllLocalIndexStatusMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllLocalIndexStatusMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_local_index_status"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllLocalIndexStatusMgr) GetTableName() string {
	return "__all_local_index_status"
}

// Reset 重置gorm会话
func (obj *_AllLocalIndexStatusMgr) Reset() *_AllLocalIndexStatusMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllLocalIndexStatusMgr) Get() (result AllLocalIndexStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllLocalIndexStatus{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllLocalIndexStatusMgr) Gets() (results []*AllLocalIndexStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllLocalIndexStatus{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllLocalIndexStatusMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllLocalIndexStatus{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithGmtCreate gmt_create获取
func (obj *_AllLocalIndexStatusMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllLocalIndexStatusMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithTenantID tenant_id获取
func (obj *_AllLocalIndexStatusMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithIndexTableID index_table_id获取
func (obj *_AllLocalIndexStatusMgr) WithIndexTableID(indexTableID int64) Option {
	return optionFunc(func(o *options) { o.query["index_table_id"] = indexTableID })
}

// WithPartitionID partition_id获取
func (obj *_AllLocalIndexStatusMgr) WithPartitionID(partitionID int64) Option {
	return optionFunc(func(o *options) { o.query["partition_id"] = partitionID })
}

// WithSvrIP svr_ip获取
func (obj *_AllLocalIndexStatusMgr) WithSvrIP(svrIP string) Option {
	return optionFunc(func(o *options) { o.query["svr_ip"] = svrIP })
}

// WithSvrPort svr_port获取
func (obj *_AllLocalIndexStatusMgr) WithSvrPort(svrPort int64) Option {
	return optionFunc(func(o *options) { o.query["svr_port"] = svrPort })
}

// WithIndexStatus index_status获取
func (obj *_AllLocalIndexStatusMgr) WithIndexStatus(indexStatus int64) Option {
	return optionFunc(func(o *options) { o.query["index_status"] = indexStatus })
}

// WithRetCode ret_code获取
func (obj *_AllLocalIndexStatusMgr) WithRetCode(retCode int64) Option {
	return optionFunc(func(o *options) { o.query["ret_code"] = retCode })
}

// WithRole role获取
func (obj *_AllLocalIndexStatusMgr) WithRole(role int64) Option {
	return optionFunc(func(o *options) { o.query["role"] = role })
}

// GetByOption 功能选项模式获取
func (obj *_AllLocalIndexStatusMgr) GetByOption(opts ...Option) (result AllLocalIndexStatus, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllLocalIndexStatus{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllLocalIndexStatusMgr) GetByOptions(opts ...Option) (results []*AllLocalIndexStatus, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllLocalIndexStatus{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllLocalIndexStatusMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllLocalIndexStatus, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllLocalIndexStatus{}).Where(options.query)
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
func (obj *_AllLocalIndexStatusMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllLocalIndexStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllLocalIndexStatus{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllLocalIndexStatusMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllLocalIndexStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllLocalIndexStatus{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllLocalIndexStatusMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllLocalIndexStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllLocalIndexStatus{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllLocalIndexStatusMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllLocalIndexStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllLocalIndexStatus{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromTenantID 通过tenant_id获取内容
func (obj *_AllLocalIndexStatusMgr) GetFromTenantID(tenantID int64) (results []*AllLocalIndexStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllLocalIndexStatus{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllLocalIndexStatusMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllLocalIndexStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllLocalIndexStatus{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromIndexTableID 通过index_table_id获取内容
func (obj *_AllLocalIndexStatusMgr) GetFromIndexTableID(indexTableID int64) (results []*AllLocalIndexStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllLocalIndexStatus{}).Where("`index_table_id` = ?", indexTableID).Find(&results).Error

	return
}

// GetBatchFromIndexTableID 批量查找
func (obj *_AllLocalIndexStatusMgr) GetBatchFromIndexTableID(indexTableIDs []int64) (results []*AllLocalIndexStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllLocalIndexStatus{}).Where("`index_table_id` IN (?)", indexTableIDs).Find(&results).Error

	return
}

// GetFromPartitionID 通过partition_id获取内容
func (obj *_AllLocalIndexStatusMgr) GetFromPartitionID(partitionID int64) (results []*AllLocalIndexStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllLocalIndexStatus{}).Where("`partition_id` = ?", partitionID).Find(&results).Error

	return
}

// GetBatchFromPartitionID 批量查找
func (obj *_AllLocalIndexStatusMgr) GetBatchFromPartitionID(partitionIDs []int64) (results []*AllLocalIndexStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllLocalIndexStatus{}).Where("`partition_id` IN (?)", partitionIDs).Find(&results).Error

	return
}

// GetFromSvrIP 通过svr_ip获取内容
func (obj *_AllLocalIndexStatusMgr) GetFromSvrIP(svrIP string) (results []*AllLocalIndexStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllLocalIndexStatus{}).Where("`svr_ip` = ?", svrIP).Find(&results).Error

	return
}

// GetBatchFromSvrIP 批量查找
func (obj *_AllLocalIndexStatusMgr) GetBatchFromSvrIP(svrIPs []string) (results []*AllLocalIndexStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllLocalIndexStatus{}).Where("`svr_ip` IN (?)", svrIPs).Find(&results).Error

	return
}

// GetFromSvrPort 通过svr_port获取内容
func (obj *_AllLocalIndexStatusMgr) GetFromSvrPort(svrPort int64) (results []*AllLocalIndexStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllLocalIndexStatus{}).Where("`svr_port` = ?", svrPort).Find(&results).Error

	return
}

// GetBatchFromSvrPort 批量查找
func (obj *_AllLocalIndexStatusMgr) GetBatchFromSvrPort(svrPorts []int64) (results []*AllLocalIndexStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllLocalIndexStatus{}).Where("`svr_port` IN (?)", svrPorts).Find(&results).Error

	return
}

// GetFromIndexStatus 通过index_status获取内容
func (obj *_AllLocalIndexStatusMgr) GetFromIndexStatus(indexStatus int64) (results []*AllLocalIndexStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllLocalIndexStatus{}).Where("`index_status` = ?", indexStatus).Find(&results).Error

	return
}

// GetBatchFromIndexStatus 批量查找
func (obj *_AllLocalIndexStatusMgr) GetBatchFromIndexStatus(indexStatuss []int64) (results []*AllLocalIndexStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllLocalIndexStatus{}).Where("`index_status` IN (?)", indexStatuss).Find(&results).Error

	return
}

// GetFromRetCode 通过ret_code获取内容
func (obj *_AllLocalIndexStatusMgr) GetFromRetCode(retCode int64) (results []*AllLocalIndexStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllLocalIndexStatus{}).Where("`ret_code` = ?", retCode).Find(&results).Error

	return
}

// GetBatchFromRetCode 批量查找
func (obj *_AllLocalIndexStatusMgr) GetBatchFromRetCode(retCodes []int64) (results []*AllLocalIndexStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllLocalIndexStatus{}).Where("`ret_code` IN (?)", retCodes).Find(&results).Error

	return
}

// GetFromRole 通过role获取内容
func (obj *_AllLocalIndexStatusMgr) GetFromRole(role int64) (results []*AllLocalIndexStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllLocalIndexStatus{}).Where("`role` = ?", role).Find(&results).Error

	return
}

// GetBatchFromRole 批量查找
func (obj *_AllLocalIndexStatusMgr) GetBatchFromRole(roles []int64) (results []*AllLocalIndexStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllLocalIndexStatus{}).Where("`role` IN (?)", roles).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllLocalIndexStatusMgr) FetchByPrimaryKey(tenantID int64, indexTableID int64, partitionID int64, svrIP string, svrPort int64) (result AllLocalIndexStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllLocalIndexStatus{}).Where("`tenant_id` = ? AND `index_table_id` = ? AND `partition_id` = ? AND `svr_ip` = ? AND `svr_port` = ?", tenantID, indexTableID, partitionID, svrIP, svrPort).First(&result).Error

	return
}
