package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _AllGlobalIndexDataSrcMgr struct {
	*_BaseMgr
}

// AllGlobalIndexDataSrcMgr open func
func AllGlobalIndexDataSrcMgr(db *gorm.DB) *_AllGlobalIndexDataSrcMgr {
	if db == nil {
		panic(fmt.Errorf("AllGlobalIndexDataSrcMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllGlobalIndexDataSrcMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_global_index_data_src"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllGlobalIndexDataSrcMgr) GetTableName() string {
	return "__all_global_index_data_src"
}

// Reset 重置gorm会话
func (obj *_AllGlobalIndexDataSrcMgr) Reset() *_AllGlobalIndexDataSrcMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllGlobalIndexDataSrcMgr) Get() (result AllGlobalIndexDataSrc, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllGlobalIndexDataSrc{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllGlobalIndexDataSrcMgr) Gets() (results []*AllGlobalIndexDataSrc, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllGlobalIndexDataSrc{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllGlobalIndexDataSrcMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllGlobalIndexDataSrc{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithGmtCreate gmt_create获取
func (obj *_AllGlobalIndexDataSrcMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllGlobalIndexDataSrcMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithTenantID tenant_id获取
func (obj *_AllGlobalIndexDataSrcMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithIndexTableID index_table_id获取
func (obj *_AllGlobalIndexDataSrcMgr) WithIndexTableID(indexTableID int64) Option {
	return optionFunc(func(o *options) { o.query["index_table_id"] = indexTableID })
}

// WithDataTableID data_table_id获取
func (obj *_AllGlobalIndexDataSrcMgr) WithDataTableID(dataTableID int64) Option {
	return optionFunc(func(o *options) { o.query["data_table_id"] = dataTableID })
}

// WithDataPartitionID data_partition_id获取
func (obj *_AllGlobalIndexDataSrcMgr) WithDataPartitionID(dataPartitionID int64) Option {
	return optionFunc(func(o *options) { o.query["data_partition_id"] = dataPartitionID })
}

// WithSvrIP svr_ip获取
func (obj *_AllGlobalIndexDataSrcMgr) WithSvrIP(svrIP string) Option {
	return optionFunc(func(o *options) { o.query["svr_ip"] = svrIP })
}

// WithSvrPort svr_port获取
func (obj *_AllGlobalIndexDataSrcMgr) WithSvrPort(svrPort int64) Option {
	return optionFunc(func(o *options) { o.query["svr_port"] = svrPort })
}

// GetByOption 功能选项模式获取
func (obj *_AllGlobalIndexDataSrcMgr) GetByOption(opts ...Option) (result AllGlobalIndexDataSrc, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllGlobalIndexDataSrc{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllGlobalIndexDataSrcMgr) GetByOptions(opts ...Option) (results []*AllGlobalIndexDataSrc, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllGlobalIndexDataSrc{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllGlobalIndexDataSrcMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllGlobalIndexDataSrc, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllGlobalIndexDataSrc{}).Where(options.query)
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
func (obj *_AllGlobalIndexDataSrcMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllGlobalIndexDataSrc, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllGlobalIndexDataSrc{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllGlobalIndexDataSrcMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllGlobalIndexDataSrc, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllGlobalIndexDataSrc{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllGlobalIndexDataSrcMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllGlobalIndexDataSrc, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllGlobalIndexDataSrc{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllGlobalIndexDataSrcMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllGlobalIndexDataSrc, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllGlobalIndexDataSrc{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromTenantID 通过tenant_id获取内容
func (obj *_AllGlobalIndexDataSrcMgr) GetFromTenantID(tenantID int64) (results []*AllGlobalIndexDataSrc, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllGlobalIndexDataSrc{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllGlobalIndexDataSrcMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllGlobalIndexDataSrc, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllGlobalIndexDataSrc{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromIndexTableID 通过index_table_id获取内容
func (obj *_AllGlobalIndexDataSrcMgr) GetFromIndexTableID(indexTableID int64) (results []*AllGlobalIndexDataSrc, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllGlobalIndexDataSrc{}).Where("`index_table_id` = ?", indexTableID).Find(&results).Error

	return
}

// GetBatchFromIndexTableID 批量查找
func (obj *_AllGlobalIndexDataSrcMgr) GetBatchFromIndexTableID(indexTableIDs []int64) (results []*AllGlobalIndexDataSrc, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllGlobalIndexDataSrc{}).Where("`index_table_id` IN (?)", indexTableIDs).Find(&results).Error

	return
}

// GetFromDataTableID 通过data_table_id获取内容
func (obj *_AllGlobalIndexDataSrcMgr) GetFromDataTableID(dataTableID int64) (results []*AllGlobalIndexDataSrc, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllGlobalIndexDataSrc{}).Where("`data_table_id` = ?", dataTableID).Find(&results).Error

	return
}

// GetBatchFromDataTableID 批量查找
func (obj *_AllGlobalIndexDataSrcMgr) GetBatchFromDataTableID(dataTableIDs []int64) (results []*AllGlobalIndexDataSrc, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllGlobalIndexDataSrc{}).Where("`data_table_id` IN (?)", dataTableIDs).Find(&results).Error

	return
}

// GetFromDataPartitionID 通过data_partition_id获取内容
func (obj *_AllGlobalIndexDataSrcMgr) GetFromDataPartitionID(dataPartitionID int64) (results []*AllGlobalIndexDataSrc, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllGlobalIndexDataSrc{}).Where("`data_partition_id` = ?", dataPartitionID).Find(&results).Error

	return
}

// GetBatchFromDataPartitionID 批量查找
func (obj *_AllGlobalIndexDataSrcMgr) GetBatchFromDataPartitionID(dataPartitionIDs []int64) (results []*AllGlobalIndexDataSrc, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllGlobalIndexDataSrc{}).Where("`data_partition_id` IN (?)", dataPartitionIDs).Find(&results).Error

	return
}

// GetFromSvrIP 通过svr_ip获取内容
func (obj *_AllGlobalIndexDataSrcMgr) GetFromSvrIP(svrIP string) (results []*AllGlobalIndexDataSrc, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllGlobalIndexDataSrc{}).Where("`svr_ip` = ?", svrIP).Find(&results).Error

	return
}

// GetBatchFromSvrIP 批量查找
func (obj *_AllGlobalIndexDataSrcMgr) GetBatchFromSvrIP(svrIPs []string) (results []*AllGlobalIndexDataSrc, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllGlobalIndexDataSrc{}).Where("`svr_ip` IN (?)", svrIPs).Find(&results).Error

	return
}

// GetFromSvrPort 通过svr_port获取内容
func (obj *_AllGlobalIndexDataSrcMgr) GetFromSvrPort(svrPort int64) (results []*AllGlobalIndexDataSrc, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllGlobalIndexDataSrc{}).Where("`svr_port` = ?", svrPort).Find(&results).Error

	return
}

// GetBatchFromSvrPort 批量查找
func (obj *_AllGlobalIndexDataSrcMgr) GetBatchFromSvrPort(svrPorts []int64) (results []*AllGlobalIndexDataSrc, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllGlobalIndexDataSrc{}).Where("`svr_port` IN (?)", svrPorts).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllGlobalIndexDataSrcMgr) FetchByPrimaryKey(tenantID int64, indexTableID int64, dataTableID int64, dataPartitionID int64) (result AllGlobalIndexDataSrc, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllGlobalIndexDataSrc{}).Where("`tenant_id` = ? AND `index_table_id` = ? AND `data_table_id` = ? AND `data_partition_id` = ?", tenantID, indexTableID, dataTableID, dataPartitionID).First(&result).Error

	return
}
