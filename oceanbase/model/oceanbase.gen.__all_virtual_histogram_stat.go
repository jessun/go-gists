package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _AllVirtualHistogramStatMgr struct {
	*_BaseMgr
}

// AllVirtualHistogramStatMgr open func
func AllVirtualHistogramStatMgr(db *gorm.DB) *_AllVirtualHistogramStatMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualHistogramStatMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualHistogramStatMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_histogram_stat"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualHistogramStatMgr) GetTableName() string {
	return "__all_virtual_histogram_stat"
}

// Reset 重置gorm会话
func (obj *_AllVirtualHistogramStatMgr) Reset() *_AllVirtualHistogramStatMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualHistogramStatMgr) Get() (result AllVirtualHistogramStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualHistogramStat{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualHistogramStatMgr) Gets() (results []*AllVirtualHistogramStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualHistogramStat{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualHistogramStatMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualHistogramStat{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTenantID tenant_id获取
func (obj *_AllVirtualHistogramStatMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithTableID table_id获取
func (obj *_AllVirtualHistogramStatMgr) WithTableID(tableID int64) Option {
	return optionFunc(func(o *options) { o.query["table_id"] = tableID })
}

// WithPartitionID partition_id获取
func (obj *_AllVirtualHistogramStatMgr) WithPartitionID(partitionID int64) Option {
	return optionFunc(func(o *options) { o.query["partition_id"] = partitionID })
}

// WithColumnID column_id获取
func (obj *_AllVirtualHistogramStatMgr) WithColumnID(columnID int64) Option {
	return optionFunc(func(o *options) { o.query["column_id"] = columnID })
}

// WithEndpointNum endpoint_num获取
func (obj *_AllVirtualHistogramStatMgr) WithEndpointNum(endpointNum int64) Option {
	return optionFunc(func(o *options) { o.query["endpoint_num"] = endpointNum })
}

// WithGmtCreate gmt_create获取
func (obj *_AllVirtualHistogramStatMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllVirtualHistogramStatMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithObjectType object_type获取
func (obj *_AllVirtualHistogramStatMgr) WithObjectType(objectType int64) Option {
	return optionFunc(func(o *options) { o.query["object_type"] = objectType })
}

// WithEndpointValue endpoint_value获取
func (obj *_AllVirtualHistogramStatMgr) WithEndpointValue(endpointValue string) Option {
	return optionFunc(func(o *options) { o.query["endpoint_value"] = endpointValue })
}

// WithBEndpointValue b_endpoint_value获取
func (obj *_AllVirtualHistogramStatMgr) WithBEndpointValue(bEndpointValue string) Option {
	return optionFunc(func(o *options) { o.query["b_endpoint_value"] = bEndpointValue })
}

// WithEndpointRepeatCnt endpoint_repeat_cnt获取
func (obj *_AllVirtualHistogramStatMgr) WithEndpointRepeatCnt(endpointRepeatCnt int64) Option {
	return optionFunc(func(o *options) { o.query["endpoint_repeat_cnt"] = endpointRepeatCnt })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualHistogramStatMgr) GetByOption(opts ...Option) (result AllVirtualHistogramStat, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualHistogramStat{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualHistogramStatMgr) GetByOptions(opts ...Option) (results []*AllVirtualHistogramStat, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualHistogramStat{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualHistogramStatMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualHistogramStat, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualHistogramStat{}).Where(options.query)
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
func (obj *_AllVirtualHistogramStatMgr) GetFromTenantID(tenantID int64) (results []*AllVirtualHistogramStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualHistogramStat{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllVirtualHistogramStatMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllVirtualHistogramStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualHistogramStat{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromTableID 通过table_id获取内容
func (obj *_AllVirtualHistogramStatMgr) GetFromTableID(tableID int64) (results []*AllVirtualHistogramStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualHistogramStat{}).Where("`table_id` = ?", tableID).Find(&results).Error

	return
}

// GetBatchFromTableID 批量查找
func (obj *_AllVirtualHistogramStatMgr) GetBatchFromTableID(tableIDs []int64) (results []*AllVirtualHistogramStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualHistogramStat{}).Where("`table_id` IN (?)", tableIDs).Find(&results).Error

	return
}

// GetFromPartitionID 通过partition_id获取内容
func (obj *_AllVirtualHistogramStatMgr) GetFromPartitionID(partitionID int64) (results []*AllVirtualHistogramStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualHistogramStat{}).Where("`partition_id` = ?", partitionID).Find(&results).Error

	return
}

// GetBatchFromPartitionID 批量查找
func (obj *_AllVirtualHistogramStatMgr) GetBatchFromPartitionID(partitionIDs []int64) (results []*AllVirtualHistogramStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualHistogramStat{}).Where("`partition_id` IN (?)", partitionIDs).Find(&results).Error

	return
}

// GetFromColumnID 通过column_id获取内容
func (obj *_AllVirtualHistogramStatMgr) GetFromColumnID(columnID int64) (results []*AllVirtualHistogramStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualHistogramStat{}).Where("`column_id` = ?", columnID).Find(&results).Error

	return
}

// GetBatchFromColumnID 批量查找
func (obj *_AllVirtualHistogramStatMgr) GetBatchFromColumnID(columnIDs []int64) (results []*AllVirtualHistogramStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualHistogramStat{}).Where("`column_id` IN (?)", columnIDs).Find(&results).Error

	return
}

// GetFromEndpointNum 通过endpoint_num获取内容
func (obj *_AllVirtualHistogramStatMgr) GetFromEndpointNum(endpointNum int64) (results []*AllVirtualHistogramStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualHistogramStat{}).Where("`endpoint_num` = ?", endpointNum).Find(&results).Error

	return
}

// GetBatchFromEndpointNum 批量查找
func (obj *_AllVirtualHistogramStatMgr) GetBatchFromEndpointNum(endpointNums []int64) (results []*AllVirtualHistogramStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualHistogramStat{}).Where("`endpoint_num` IN (?)", endpointNums).Find(&results).Error

	return
}

// GetFromGmtCreate 通过gmt_create获取内容
func (obj *_AllVirtualHistogramStatMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllVirtualHistogramStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualHistogramStat{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllVirtualHistogramStatMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllVirtualHistogramStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualHistogramStat{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllVirtualHistogramStatMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllVirtualHistogramStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualHistogramStat{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllVirtualHistogramStatMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllVirtualHistogramStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualHistogramStat{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromObjectType 通过object_type获取内容
func (obj *_AllVirtualHistogramStatMgr) GetFromObjectType(objectType int64) (results []*AllVirtualHistogramStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualHistogramStat{}).Where("`object_type` = ?", objectType).Find(&results).Error

	return
}

// GetBatchFromObjectType 批量查找
func (obj *_AllVirtualHistogramStatMgr) GetBatchFromObjectType(objectTypes []int64) (results []*AllVirtualHistogramStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualHistogramStat{}).Where("`object_type` IN (?)", objectTypes).Find(&results).Error

	return
}

// GetFromEndpointValue 通过endpoint_value获取内容
func (obj *_AllVirtualHistogramStatMgr) GetFromEndpointValue(endpointValue string) (results []*AllVirtualHistogramStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualHistogramStat{}).Where("`endpoint_value` = ?", endpointValue).Find(&results).Error

	return
}

// GetBatchFromEndpointValue 批量查找
func (obj *_AllVirtualHistogramStatMgr) GetBatchFromEndpointValue(endpointValues []string) (results []*AllVirtualHistogramStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualHistogramStat{}).Where("`endpoint_value` IN (?)", endpointValues).Find(&results).Error

	return
}

// GetFromBEndpointValue 通过b_endpoint_value获取内容
func (obj *_AllVirtualHistogramStatMgr) GetFromBEndpointValue(bEndpointValue string) (results []*AllVirtualHistogramStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualHistogramStat{}).Where("`b_endpoint_value` = ?", bEndpointValue).Find(&results).Error

	return
}

// GetBatchFromBEndpointValue 批量查找
func (obj *_AllVirtualHistogramStatMgr) GetBatchFromBEndpointValue(bEndpointValues []string) (results []*AllVirtualHistogramStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualHistogramStat{}).Where("`b_endpoint_value` IN (?)", bEndpointValues).Find(&results).Error

	return
}

// GetFromEndpointRepeatCnt 通过endpoint_repeat_cnt获取内容
func (obj *_AllVirtualHistogramStatMgr) GetFromEndpointRepeatCnt(endpointRepeatCnt int64) (results []*AllVirtualHistogramStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualHistogramStat{}).Where("`endpoint_repeat_cnt` = ?", endpointRepeatCnt).Find(&results).Error

	return
}

// GetBatchFromEndpointRepeatCnt 批量查找
func (obj *_AllVirtualHistogramStatMgr) GetBatchFromEndpointRepeatCnt(endpointRepeatCnts []int64) (results []*AllVirtualHistogramStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualHistogramStat{}).Where("`endpoint_repeat_cnt` IN (?)", endpointRepeatCnts).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllVirtualHistogramStatMgr) FetchByPrimaryKey(tenantID int64, tableID int64, partitionID int64, columnID int64, endpointNum int64) (result AllVirtualHistogramStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualHistogramStat{}).Where("`tenant_id` = ? AND `table_id` = ? AND `partition_id` = ? AND `column_id` = ? AND `endpoint_num` = ?", tenantID, tableID, partitionID, columnID, endpointNum).First(&result).Error

	return
}
