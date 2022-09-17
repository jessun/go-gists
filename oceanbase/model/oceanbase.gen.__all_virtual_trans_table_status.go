package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _AllVirtualTransTableStatusMgr struct {
	*_BaseMgr
}

// AllVirtualTransTableStatusMgr open func
func AllVirtualTransTableStatusMgr(db *gorm.DB) *_AllVirtualTransTableStatusMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualTransTableStatusMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualTransTableStatusMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_trans_table_status"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualTransTableStatusMgr) GetTableName() string {
	return "__all_virtual_trans_table_status"
}

// Reset 重置gorm会话
func (obj *_AllVirtualTransTableStatusMgr) Reset() *_AllVirtualTransTableStatusMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualTransTableStatusMgr) Get() (result AllVirtualTransTableStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransTableStatus{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualTransTableStatusMgr) Gets() (results []*AllVirtualTransTableStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransTableStatus{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualTransTableStatusMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualTransTableStatus{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTenantID tenant_id获取
func (obj *_AllVirtualTransTableStatusMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithSvrIP svr_ip获取
func (obj *_AllVirtualTransTableStatusMgr) WithSvrIP(svrIP string) Option {
	return optionFunc(func(o *options) { o.query["svr_ip"] = svrIP })
}

// WithSvrPort svr_port获取
func (obj *_AllVirtualTransTableStatusMgr) WithSvrPort(svrPort int64) Option {
	return optionFunc(func(o *options) { o.query["svr_port"] = svrPort })
}

// WithZone zone获取
func (obj *_AllVirtualTransTableStatusMgr) WithZone(zone string) Option {
	return optionFunc(func(o *options) { o.query["zone"] = zone })
}

// WithTableID table_id获取
func (obj *_AllVirtualTransTableStatusMgr) WithTableID(tableID int64) Option {
	return optionFunc(func(o *options) { o.query["table_id"] = tableID })
}

// WithPartitionID partition_id获取
func (obj *_AllVirtualTransTableStatusMgr) WithPartitionID(partitionID int64) Option {
	return optionFunc(func(o *options) { o.query["partition_id"] = partitionID })
}

// WithEndLogID end_log_id获取
func (obj *_AllVirtualTransTableStatusMgr) WithEndLogID(endLogID int64) Option {
	return optionFunc(func(o *options) { o.query["end_log_id"] = endLogID })
}

// WithTransCnt trans_cnt获取
func (obj *_AllVirtualTransTableStatusMgr) WithTransCnt(transCnt int64) Option {
	return optionFunc(func(o *options) { o.query["trans_cnt"] = transCnt })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualTransTableStatusMgr) GetByOption(opts ...Option) (result AllVirtualTransTableStatus, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransTableStatus{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualTransTableStatusMgr) GetByOptions(opts ...Option) (results []*AllVirtualTransTableStatus, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransTableStatus{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualTransTableStatusMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualTransTableStatus, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualTransTableStatus{}).Where(options.query)
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
func (obj *_AllVirtualTransTableStatusMgr) GetFromTenantID(tenantID int64) (results []*AllVirtualTransTableStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransTableStatus{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllVirtualTransTableStatusMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllVirtualTransTableStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransTableStatus{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromSvrIP 通过svr_ip获取内容
func (obj *_AllVirtualTransTableStatusMgr) GetFromSvrIP(svrIP string) (results []*AllVirtualTransTableStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransTableStatus{}).Where("`svr_ip` = ?", svrIP).Find(&results).Error

	return
}

// GetBatchFromSvrIP 批量查找
func (obj *_AllVirtualTransTableStatusMgr) GetBatchFromSvrIP(svrIPs []string) (results []*AllVirtualTransTableStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransTableStatus{}).Where("`svr_ip` IN (?)", svrIPs).Find(&results).Error

	return
}

// GetFromSvrPort 通过svr_port获取内容
func (obj *_AllVirtualTransTableStatusMgr) GetFromSvrPort(svrPort int64) (results []*AllVirtualTransTableStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransTableStatus{}).Where("`svr_port` = ?", svrPort).Find(&results).Error

	return
}

// GetBatchFromSvrPort 批量查找
func (obj *_AllVirtualTransTableStatusMgr) GetBatchFromSvrPort(svrPorts []int64) (results []*AllVirtualTransTableStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransTableStatus{}).Where("`svr_port` IN (?)", svrPorts).Find(&results).Error

	return
}

// GetFromZone 通过zone获取内容
func (obj *_AllVirtualTransTableStatusMgr) GetFromZone(zone string) (results []*AllVirtualTransTableStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransTableStatus{}).Where("`zone` = ?", zone).Find(&results).Error

	return
}

// GetBatchFromZone 批量查找
func (obj *_AllVirtualTransTableStatusMgr) GetBatchFromZone(zones []string) (results []*AllVirtualTransTableStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransTableStatus{}).Where("`zone` IN (?)", zones).Find(&results).Error

	return
}

// GetFromTableID 通过table_id获取内容
func (obj *_AllVirtualTransTableStatusMgr) GetFromTableID(tableID int64) (results []*AllVirtualTransTableStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransTableStatus{}).Where("`table_id` = ?", tableID).Find(&results).Error

	return
}

// GetBatchFromTableID 批量查找
func (obj *_AllVirtualTransTableStatusMgr) GetBatchFromTableID(tableIDs []int64) (results []*AllVirtualTransTableStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransTableStatus{}).Where("`table_id` IN (?)", tableIDs).Find(&results).Error

	return
}

// GetFromPartitionID 通过partition_id获取内容
func (obj *_AllVirtualTransTableStatusMgr) GetFromPartitionID(partitionID int64) (results []*AllVirtualTransTableStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransTableStatus{}).Where("`partition_id` = ?", partitionID).Find(&results).Error

	return
}

// GetBatchFromPartitionID 批量查找
func (obj *_AllVirtualTransTableStatusMgr) GetBatchFromPartitionID(partitionIDs []int64) (results []*AllVirtualTransTableStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransTableStatus{}).Where("`partition_id` IN (?)", partitionIDs).Find(&results).Error

	return
}

// GetFromEndLogID 通过end_log_id获取内容
func (obj *_AllVirtualTransTableStatusMgr) GetFromEndLogID(endLogID int64) (results []*AllVirtualTransTableStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransTableStatus{}).Where("`end_log_id` = ?", endLogID).Find(&results).Error

	return
}

// GetBatchFromEndLogID 批量查找
func (obj *_AllVirtualTransTableStatusMgr) GetBatchFromEndLogID(endLogIDs []int64) (results []*AllVirtualTransTableStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransTableStatus{}).Where("`end_log_id` IN (?)", endLogIDs).Find(&results).Error

	return
}

// GetFromTransCnt 通过trans_cnt获取内容
func (obj *_AllVirtualTransTableStatusMgr) GetFromTransCnt(transCnt int64) (results []*AllVirtualTransTableStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransTableStatus{}).Where("`trans_cnt` = ?", transCnt).Find(&results).Error

	return
}

// GetBatchFromTransCnt 批量查找
func (obj *_AllVirtualTransTableStatusMgr) GetBatchFromTransCnt(transCnts []int64) (results []*AllVirtualTransTableStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTransTableStatus{}).Where("`trans_cnt` IN (?)", transCnts).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////
