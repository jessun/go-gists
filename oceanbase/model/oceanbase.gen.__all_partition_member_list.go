package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _AllPartitionMemberListMgr struct {
	*_BaseMgr
}

// AllPartitionMemberListMgr open func
func AllPartitionMemberListMgr(db *gorm.DB) *_AllPartitionMemberListMgr {
	if db == nil {
		panic(fmt.Errorf("AllPartitionMemberListMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllPartitionMemberListMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_partition_member_list"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllPartitionMemberListMgr) GetTableName() string {
	return "__all_partition_member_list"
}

// Reset 重置gorm会话
func (obj *_AllPartitionMemberListMgr) Reset() *_AllPartitionMemberListMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllPartitionMemberListMgr) Get() (result AllPartitionMemberList, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPartitionMemberList{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllPartitionMemberListMgr) Gets() (results []*AllPartitionMemberList, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPartitionMemberList{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllPartitionMemberListMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllPartitionMemberList{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithGmtCreate gmt_create获取
func (obj *_AllPartitionMemberListMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllPartitionMemberListMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithTenantID tenant_id获取
func (obj *_AllPartitionMemberListMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithTableID table_id获取
func (obj *_AllPartitionMemberListMgr) WithTableID(tableID int64) Option {
	return optionFunc(func(o *options) { o.query["table_id"] = tableID })
}

// WithPartitionID partition_id获取
func (obj *_AllPartitionMemberListMgr) WithPartitionID(partitionID int64) Option {
	return optionFunc(func(o *options) { o.query["partition_id"] = partitionID })
}

// WithMemberList member_list获取
func (obj *_AllPartitionMemberListMgr) WithMemberList(memberList string) Option {
	return optionFunc(func(o *options) { o.query["member_list"] = memberList })
}

// WithSchemaVersion schema_version获取
func (obj *_AllPartitionMemberListMgr) WithSchemaVersion(schemaVersion int64) Option {
	return optionFunc(func(o *options) { o.query["schema_version"] = schemaVersion })
}

// WithPartitionStatus partition_status获取
func (obj *_AllPartitionMemberListMgr) WithPartitionStatus(partitionStatus int64) Option {
	return optionFunc(func(o *options) { o.query["partition_status"] = partitionStatus })
}

// GetByOption 功能选项模式获取
func (obj *_AllPartitionMemberListMgr) GetByOption(opts ...Option) (result AllPartitionMemberList, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllPartitionMemberList{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllPartitionMemberListMgr) GetByOptions(opts ...Option) (results []*AllPartitionMemberList, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllPartitionMemberList{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllPartitionMemberListMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllPartitionMemberList, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllPartitionMemberList{}).Where(options.query)
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
func (obj *_AllPartitionMemberListMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllPartitionMemberList, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPartitionMemberList{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllPartitionMemberListMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllPartitionMemberList, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPartitionMemberList{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllPartitionMemberListMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllPartitionMemberList, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPartitionMemberList{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllPartitionMemberListMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllPartitionMemberList, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPartitionMemberList{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromTenantID 通过tenant_id获取内容
func (obj *_AllPartitionMemberListMgr) GetFromTenantID(tenantID int64) (results []*AllPartitionMemberList, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPartitionMemberList{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllPartitionMemberListMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllPartitionMemberList, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPartitionMemberList{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromTableID 通过table_id获取内容
func (obj *_AllPartitionMemberListMgr) GetFromTableID(tableID int64) (results []*AllPartitionMemberList, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPartitionMemberList{}).Where("`table_id` = ?", tableID).Find(&results).Error

	return
}

// GetBatchFromTableID 批量查找
func (obj *_AllPartitionMemberListMgr) GetBatchFromTableID(tableIDs []int64) (results []*AllPartitionMemberList, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPartitionMemberList{}).Where("`table_id` IN (?)", tableIDs).Find(&results).Error

	return
}

// GetFromPartitionID 通过partition_id获取内容
func (obj *_AllPartitionMemberListMgr) GetFromPartitionID(partitionID int64) (results []*AllPartitionMemberList, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPartitionMemberList{}).Where("`partition_id` = ?", partitionID).Find(&results).Error

	return
}

// GetBatchFromPartitionID 批量查找
func (obj *_AllPartitionMemberListMgr) GetBatchFromPartitionID(partitionIDs []int64) (results []*AllPartitionMemberList, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPartitionMemberList{}).Where("`partition_id` IN (?)", partitionIDs).Find(&results).Error

	return
}

// GetFromMemberList 通过member_list获取内容
func (obj *_AllPartitionMemberListMgr) GetFromMemberList(memberList string) (results []*AllPartitionMemberList, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPartitionMemberList{}).Where("`member_list` = ?", memberList).Find(&results).Error

	return
}

// GetBatchFromMemberList 批量查找
func (obj *_AllPartitionMemberListMgr) GetBatchFromMemberList(memberLists []string) (results []*AllPartitionMemberList, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPartitionMemberList{}).Where("`member_list` IN (?)", memberLists).Find(&results).Error

	return
}

// GetFromSchemaVersion 通过schema_version获取内容
func (obj *_AllPartitionMemberListMgr) GetFromSchemaVersion(schemaVersion int64) (results []*AllPartitionMemberList, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPartitionMemberList{}).Where("`schema_version` = ?", schemaVersion).Find(&results).Error

	return
}

// GetBatchFromSchemaVersion 批量查找
func (obj *_AllPartitionMemberListMgr) GetBatchFromSchemaVersion(schemaVersions []int64) (results []*AllPartitionMemberList, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPartitionMemberList{}).Where("`schema_version` IN (?)", schemaVersions).Find(&results).Error

	return
}

// GetFromPartitionStatus 通过partition_status获取内容
func (obj *_AllPartitionMemberListMgr) GetFromPartitionStatus(partitionStatus int64) (results []*AllPartitionMemberList, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPartitionMemberList{}).Where("`partition_status` = ?", partitionStatus).Find(&results).Error

	return
}

// GetBatchFromPartitionStatus 批量查找
func (obj *_AllPartitionMemberListMgr) GetBatchFromPartitionStatus(partitionStatuss []int64) (results []*AllPartitionMemberList, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPartitionMemberList{}).Where("`partition_status` IN (?)", partitionStatuss).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllPartitionMemberListMgr) FetchByPrimaryKey(tenantID int64, tableID int64, partitionID int64) (result AllPartitionMemberList, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllPartitionMemberList{}).Where("`tenant_id` = ? AND `table_id` = ? AND `partition_id` = ?", tenantID, tableID, partitionID).First(&result).Error

	return
}
