package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _AllVirtualTenantPartitionMetaTableMgr struct {
	*_BaseMgr
}

// AllVirtualTenantPartitionMetaTableMgr open func
func AllVirtualTenantPartitionMetaTableMgr(db *gorm.DB) *_AllVirtualTenantPartitionMetaTableMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualTenantPartitionMetaTableMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualTenantPartitionMetaTableMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_tenant_partition_meta_table"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualTenantPartitionMetaTableMgr) GetTableName() string {
	return "__all_virtual_tenant_partition_meta_table"
}

// Reset 重置gorm会话
func (obj *_AllVirtualTenantPartitionMetaTableMgr) Reset() *_AllVirtualTenantPartitionMetaTableMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualTenantPartitionMetaTableMgr) Get() (result AllVirtualTenantPartitionMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantPartitionMetaTable{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualTenantPartitionMetaTableMgr) Gets() (results []*AllVirtualTenantPartitionMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantPartitionMetaTable{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualTenantPartitionMetaTableMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantPartitionMetaTable{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTenantID tenant_id获取
func (obj *_AllVirtualTenantPartitionMetaTableMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithTableID table_id获取
func (obj *_AllVirtualTenantPartitionMetaTableMgr) WithTableID(tableID int64) Option {
	return optionFunc(func(o *options) { o.query["table_id"] = tableID })
}

// WithPartitionID partition_id获取
func (obj *_AllVirtualTenantPartitionMetaTableMgr) WithPartitionID(partitionID int64) Option {
	return optionFunc(func(o *options) { o.query["partition_id"] = partitionID })
}

// WithSvrIP svr_ip获取
func (obj *_AllVirtualTenantPartitionMetaTableMgr) WithSvrIP(svrIP string) Option {
	return optionFunc(func(o *options) { o.query["svr_ip"] = svrIP })
}

// WithSvrPort svr_port获取
func (obj *_AllVirtualTenantPartitionMetaTableMgr) WithSvrPort(svrPort int64) Option {
	return optionFunc(func(o *options) { o.query["svr_port"] = svrPort })
}

// WithGmtCreate gmt_create获取
func (obj *_AllVirtualTenantPartitionMetaTableMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllVirtualTenantPartitionMetaTableMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithRole role获取
func (obj *_AllVirtualTenantPartitionMetaTableMgr) WithRole(role int64) Option {
	return optionFunc(func(o *options) { o.query["role"] = role })
}

// WithRowCount row_count获取
func (obj *_AllVirtualTenantPartitionMetaTableMgr) WithRowCount(rowCount int64) Option {
	return optionFunc(func(o *options) { o.query["row_count"] = rowCount })
}

// WithDataSize data_size获取
func (obj *_AllVirtualTenantPartitionMetaTableMgr) WithDataSize(dataSize int64) Option {
	return optionFunc(func(o *options) { o.query["data_size"] = dataSize })
}

// WithDataVersion data_version获取
func (obj *_AllVirtualTenantPartitionMetaTableMgr) WithDataVersion(dataVersion int64) Option {
	return optionFunc(func(o *options) { o.query["data_version"] = dataVersion })
}

// WithRequiredSize required_size获取
func (obj *_AllVirtualTenantPartitionMetaTableMgr) WithRequiredSize(requiredSize int64) Option {
	return optionFunc(func(o *options) { o.query["required_size"] = requiredSize })
}

// WithReplicaType replica_type获取
func (obj *_AllVirtualTenantPartitionMetaTableMgr) WithReplicaType(replicaType int64) Option {
	return optionFunc(func(o *options) { o.query["replica_type"] = replicaType })
}

// WithStatus status获取
func (obj *_AllVirtualTenantPartitionMetaTableMgr) WithStatus(status string) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithDataChecksum data_checksum获取
func (obj *_AllVirtualTenantPartitionMetaTableMgr) WithDataChecksum(dataChecksum int64) Option {
	return optionFunc(func(o *options) { o.query["data_checksum"] = dataChecksum })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualTenantPartitionMetaTableMgr) GetByOption(opts ...Option) (result AllVirtualTenantPartitionMetaTable, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantPartitionMetaTable{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualTenantPartitionMetaTableMgr) GetByOptions(opts ...Option) (results []*AllVirtualTenantPartitionMetaTable, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantPartitionMetaTable{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualTenantPartitionMetaTableMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualTenantPartitionMetaTable, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantPartitionMetaTable{}).Where(options.query)
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
func (obj *_AllVirtualTenantPartitionMetaTableMgr) GetFromTenantID(tenantID int64) (results []*AllVirtualTenantPartitionMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantPartitionMetaTable{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllVirtualTenantPartitionMetaTableMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllVirtualTenantPartitionMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantPartitionMetaTable{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromTableID 通过table_id获取内容
func (obj *_AllVirtualTenantPartitionMetaTableMgr) GetFromTableID(tableID int64) (results []*AllVirtualTenantPartitionMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantPartitionMetaTable{}).Where("`table_id` = ?", tableID).Find(&results).Error

	return
}

// GetBatchFromTableID 批量查找
func (obj *_AllVirtualTenantPartitionMetaTableMgr) GetBatchFromTableID(tableIDs []int64) (results []*AllVirtualTenantPartitionMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantPartitionMetaTable{}).Where("`table_id` IN (?)", tableIDs).Find(&results).Error

	return
}

// GetFromPartitionID 通过partition_id获取内容
func (obj *_AllVirtualTenantPartitionMetaTableMgr) GetFromPartitionID(partitionID int64) (results []*AllVirtualTenantPartitionMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantPartitionMetaTable{}).Where("`partition_id` = ?", partitionID).Find(&results).Error

	return
}

// GetBatchFromPartitionID 批量查找
func (obj *_AllVirtualTenantPartitionMetaTableMgr) GetBatchFromPartitionID(partitionIDs []int64) (results []*AllVirtualTenantPartitionMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantPartitionMetaTable{}).Where("`partition_id` IN (?)", partitionIDs).Find(&results).Error

	return
}

// GetFromSvrIP 通过svr_ip获取内容
func (obj *_AllVirtualTenantPartitionMetaTableMgr) GetFromSvrIP(svrIP string) (results []*AllVirtualTenantPartitionMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantPartitionMetaTable{}).Where("`svr_ip` = ?", svrIP).Find(&results).Error

	return
}

// GetBatchFromSvrIP 批量查找
func (obj *_AllVirtualTenantPartitionMetaTableMgr) GetBatchFromSvrIP(svrIPs []string) (results []*AllVirtualTenantPartitionMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantPartitionMetaTable{}).Where("`svr_ip` IN (?)", svrIPs).Find(&results).Error

	return
}

// GetFromSvrPort 通过svr_port获取内容
func (obj *_AllVirtualTenantPartitionMetaTableMgr) GetFromSvrPort(svrPort int64) (results []*AllVirtualTenantPartitionMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantPartitionMetaTable{}).Where("`svr_port` = ?", svrPort).Find(&results).Error

	return
}

// GetBatchFromSvrPort 批量查找
func (obj *_AllVirtualTenantPartitionMetaTableMgr) GetBatchFromSvrPort(svrPorts []int64) (results []*AllVirtualTenantPartitionMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantPartitionMetaTable{}).Where("`svr_port` IN (?)", svrPorts).Find(&results).Error

	return
}

// GetFromGmtCreate 通过gmt_create获取内容
func (obj *_AllVirtualTenantPartitionMetaTableMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllVirtualTenantPartitionMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantPartitionMetaTable{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllVirtualTenantPartitionMetaTableMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllVirtualTenantPartitionMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantPartitionMetaTable{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllVirtualTenantPartitionMetaTableMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllVirtualTenantPartitionMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantPartitionMetaTable{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllVirtualTenantPartitionMetaTableMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllVirtualTenantPartitionMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantPartitionMetaTable{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromRole 通过role获取内容
func (obj *_AllVirtualTenantPartitionMetaTableMgr) GetFromRole(role int64) (results []*AllVirtualTenantPartitionMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantPartitionMetaTable{}).Where("`role` = ?", role).Find(&results).Error

	return
}

// GetBatchFromRole 批量查找
func (obj *_AllVirtualTenantPartitionMetaTableMgr) GetBatchFromRole(roles []int64) (results []*AllVirtualTenantPartitionMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantPartitionMetaTable{}).Where("`role` IN (?)", roles).Find(&results).Error

	return
}

// GetFromRowCount 通过row_count获取内容
func (obj *_AllVirtualTenantPartitionMetaTableMgr) GetFromRowCount(rowCount int64) (results []*AllVirtualTenantPartitionMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantPartitionMetaTable{}).Where("`row_count` = ?", rowCount).Find(&results).Error

	return
}

// GetBatchFromRowCount 批量查找
func (obj *_AllVirtualTenantPartitionMetaTableMgr) GetBatchFromRowCount(rowCounts []int64) (results []*AllVirtualTenantPartitionMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantPartitionMetaTable{}).Where("`row_count` IN (?)", rowCounts).Find(&results).Error

	return
}

// GetFromDataSize 通过data_size获取内容
func (obj *_AllVirtualTenantPartitionMetaTableMgr) GetFromDataSize(dataSize int64) (results []*AllVirtualTenantPartitionMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantPartitionMetaTable{}).Where("`data_size` = ?", dataSize).Find(&results).Error

	return
}

// GetBatchFromDataSize 批量查找
func (obj *_AllVirtualTenantPartitionMetaTableMgr) GetBatchFromDataSize(dataSizes []int64) (results []*AllVirtualTenantPartitionMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantPartitionMetaTable{}).Where("`data_size` IN (?)", dataSizes).Find(&results).Error

	return
}

// GetFromDataVersion 通过data_version获取内容
func (obj *_AllVirtualTenantPartitionMetaTableMgr) GetFromDataVersion(dataVersion int64) (results []*AllVirtualTenantPartitionMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantPartitionMetaTable{}).Where("`data_version` = ?", dataVersion).Find(&results).Error

	return
}

// GetBatchFromDataVersion 批量查找
func (obj *_AllVirtualTenantPartitionMetaTableMgr) GetBatchFromDataVersion(dataVersions []int64) (results []*AllVirtualTenantPartitionMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantPartitionMetaTable{}).Where("`data_version` IN (?)", dataVersions).Find(&results).Error

	return
}

// GetFromRequiredSize 通过required_size获取内容
func (obj *_AllVirtualTenantPartitionMetaTableMgr) GetFromRequiredSize(requiredSize int64) (results []*AllVirtualTenantPartitionMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantPartitionMetaTable{}).Where("`required_size` = ?", requiredSize).Find(&results).Error

	return
}

// GetBatchFromRequiredSize 批量查找
func (obj *_AllVirtualTenantPartitionMetaTableMgr) GetBatchFromRequiredSize(requiredSizes []int64) (results []*AllVirtualTenantPartitionMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantPartitionMetaTable{}).Where("`required_size` IN (?)", requiredSizes).Find(&results).Error

	return
}

// GetFromReplicaType 通过replica_type获取内容
func (obj *_AllVirtualTenantPartitionMetaTableMgr) GetFromReplicaType(replicaType int64) (results []*AllVirtualTenantPartitionMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantPartitionMetaTable{}).Where("`replica_type` = ?", replicaType).Find(&results).Error

	return
}

// GetBatchFromReplicaType 批量查找
func (obj *_AllVirtualTenantPartitionMetaTableMgr) GetBatchFromReplicaType(replicaTypes []int64) (results []*AllVirtualTenantPartitionMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantPartitionMetaTable{}).Where("`replica_type` IN (?)", replicaTypes).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容
func (obj *_AllVirtualTenantPartitionMetaTableMgr) GetFromStatus(status string) (results []*AllVirtualTenantPartitionMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantPartitionMetaTable{}).Where("`status` = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量查找
func (obj *_AllVirtualTenantPartitionMetaTableMgr) GetBatchFromStatus(statuss []string) (results []*AllVirtualTenantPartitionMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantPartitionMetaTable{}).Where("`status` IN (?)", statuss).Find(&results).Error

	return
}

// GetFromDataChecksum 通过data_checksum获取内容
func (obj *_AllVirtualTenantPartitionMetaTableMgr) GetFromDataChecksum(dataChecksum int64) (results []*AllVirtualTenantPartitionMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantPartitionMetaTable{}).Where("`data_checksum` = ?", dataChecksum).Find(&results).Error

	return
}

// GetBatchFromDataChecksum 批量查找
func (obj *_AllVirtualTenantPartitionMetaTableMgr) GetBatchFromDataChecksum(dataChecksums []int64) (results []*AllVirtualTenantPartitionMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantPartitionMetaTable{}).Where("`data_checksum` IN (?)", dataChecksums).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllVirtualTenantPartitionMetaTableMgr) FetchByPrimaryKey(tenantID int64, tableID int64, partitionID int64, svrIP string, svrPort int64) (result AllVirtualTenantPartitionMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualTenantPartitionMetaTable{}).Where("`tenant_id` = ? AND `table_id` = ? AND `partition_id` = ? AND `svr_ip` = ? AND `svr_port` = ?", tenantID, tableID, partitionID, svrIP, svrPort).First(&result).Error

	return
}
