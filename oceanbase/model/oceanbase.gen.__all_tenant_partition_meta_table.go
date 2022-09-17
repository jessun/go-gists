package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _AllTenantPartitionMetaTableMgr struct {
	*_BaseMgr
}

// AllTenantPartitionMetaTableMgr open func
func AllTenantPartitionMetaTableMgr(db *gorm.DB) *_AllTenantPartitionMetaTableMgr {
	if db == nil {
		panic(fmt.Errorf("AllTenantPartitionMetaTableMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllTenantPartitionMetaTableMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_tenant_partition_meta_table"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllTenantPartitionMetaTableMgr) GetTableName() string {
	return "__all_tenant_partition_meta_table"
}

// Reset 重置gorm会话
func (obj *_AllTenantPartitionMetaTableMgr) Reset() *_AllTenantPartitionMetaTableMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllTenantPartitionMetaTableMgr) Get() (result AllTenantPartitionMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantPartitionMetaTable{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllTenantPartitionMetaTableMgr) Gets() (results []*AllTenantPartitionMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantPartitionMetaTable{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllTenantPartitionMetaTableMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllTenantPartitionMetaTable{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithGmtCreate gmt_create获取
func (obj *_AllTenantPartitionMetaTableMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllTenantPartitionMetaTableMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithTenantID tenant_id获取
func (obj *_AllTenantPartitionMetaTableMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithTableID table_id获取
func (obj *_AllTenantPartitionMetaTableMgr) WithTableID(tableID int64) Option {
	return optionFunc(func(o *options) { o.query["table_id"] = tableID })
}

// WithPartitionID partition_id获取
func (obj *_AllTenantPartitionMetaTableMgr) WithPartitionID(partitionID int64) Option {
	return optionFunc(func(o *options) { o.query["partition_id"] = partitionID })
}

// WithSvrIP svr_ip获取
func (obj *_AllTenantPartitionMetaTableMgr) WithSvrIP(svrIP string) Option {
	return optionFunc(func(o *options) { o.query["svr_ip"] = svrIP })
}

// WithSvrPort svr_port获取
func (obj *_AllTenantPartitionMetaTableMgr) WithSvrPort(svrPort int64) Option {
	return optionFunc(func(o *options) { o.query["svr_port"] = svrPort })
}

// WithRole role获取
func (obj *_AllTenantPartitionMetaTableMgr) WithRole(role int64) Option {
	return optionFunc(func(o *options) { o.query["role"] = role })
}

// WithRowCount row_count获取
func (obj *_AllTenantPartitionMetaTableMgr) WithRowCount(rowCount int64) Option {
	return optionFunc(func(o *options) { o.query["row_count"] = rowCount })
}

// WithDataSize data_size获取
func (obj *_AllTenantPartitionMetaTableMgr) WithDataSize(dataSize int64) Option {
	return optionFunc(func(o *options) { o.query["data_size"] = dataSize })
}

// WithDataVersion data_version获取
func (obj *_AllTenantPartitionMetaTableMgr) WithDataVersion(dataVersion int64) Option {
	return optionFunc(func(o *options) { o.query["data_version"] = dataVersion })
}

// WithRequiredSize required_size获取
func (obj *_AllTenantPartitionMetaTableMgr) WithRequiredSize(requiredSize int64) Option {
	return optionFunc(func(o *options) { o.query["required_size"] = requiredSize })
}

// WithReplicaType replica_type获取
func (obj *_AllTenantPartitionMetaTableMgr) WithReplicaType(replicaType int64) Option {
	return optionFunc(func(o *options) { o.query["replica_type"] = replicaType })
}

// WithStatus status获取
func (obj *_AllTenantPartitionMetaTableMgr) WithStatus(status string) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithDataChecksum data_checksum获取
func (obj *_AllTenantPartitionMetaTableMgr) WithDataChecksum(dataChecksum int64) Option {
	return optionFunc(func(o *options) { o.query["data_checksum"] = dataChecksum })
}

// GetByOption 功能选项模式获取
func (obj *_AllTenantPartitionMetaTableMgr) GetByOption(opts ...Option) (result AllTenantPartitionMetaTable, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllTenantPartitionMetaTable{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllTenantPartitionMetaTableMgr) GetByOptions(opts ...Option) (results []*AllTenantPartitionMetaTable, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllTenantPartitionMetaTable{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllTenantPartitionMetaTableMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllTenantPartitionMetaTable, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllTenantPartitionMetaTable{}).Where(options.query)
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
func (obj *_AllTenantPartitionMetaTableMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllTenantPartitionMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantPartitionMetaTable{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllTenantPartitionMetaTableMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllTenantPartitionMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantPartitionMetaTable{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllTenantPartitionMetaTableMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllTenantPartitionMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantPartitionMetaTable{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllTenantPartitionMetaTableMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllTenantPartitionMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantPartitionMetaTable{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromTenantID 通过tenant_id获取内容
func (obj *_AllTenantPartitionMetaTableMgr) GetFromTenantID(tenantID int64) (results []*AllTenantPartitionMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantPartitionMetaTable{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllTenantPartitionMetaTableMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllTenantPartitionMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantPartitionMetaTable{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromTableID 通过table_id获取内容
func (obj *_AllTenantPartitionMetaTableMgr) GetFromTableID(tableID int64) (results []*AllTenantPartitionMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantPartitionMetaTable{}).Where("`table_id` = ?", tableID).Find(&results).Error

	return
}

// GetBatchFromTableID 批量查找
func (obj *_AllTenantPartitionMetaTableMgr) GetBatchFromTableID(tableIDs []int64) (results []*AllTenantPartitionMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantPartitionMetaTable{}).Where("`table_id` IN (?)", tableIDs).Find(&results).Error

	return
}

// GetFromPartitionID 通过partition_id获取内容
func (obj *_AllTenantPartitionMetaTableMgr) GetFromPartitionID(partitionID int64) (results []*AllTenantPartitionMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantPartitionMetaTable{}).Where("`partition_id` = ?", partitionID).Find(&results).Error

	return
}

// GetBatchFromPartitionID 批量查找
func (obj *_AllTenantPartitionMetaTableMgr) GetBatchFromPartitionID(partitionIDs []int64) (results []*AllTenantPartitionMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantPartitionMetaTable{}).Where("`partition_id` IN (?)", partitionIDs).Find(&results).Error

	return
}

// GetFromSvrIP 通过svr_ip获取内容
func (obj *_AllTenantPartitionMetaTableMgr) GetFromSvrIP(svrIP string) (results []*AllTenantPartitionMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantPartitionMetaTable{}).Where("`svr_ip` = ?", svrIP).Find(&results).Error

	return
}

// GetBatchFromSvrIP 批量查找
func (obj *_AllTenantPartitionMetaTableMgr) GetBatchFromSvrIP(svrIPs []string) (results []*AllTenantPartitionMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantPartitionMetaTable{}).Where("`svr_ip` IN (?)", svrIPs).Find(&results).Error

	return
}

// GetFromSvrPort 通过svr_port获取内容
func (obj *_AllTenantPartitionMetaTableMgr) GetFromSvrPort(svrPort int64) (results []*AllTenantPartitionMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantPartitionMetaTable{}).Where("`svr_port` = ?", svrPort).Find(&results).Error

	return
}

// GetBatchFromSvrPort 批量查找
func (obj *_AllTenantPartitionMetaTableMgr) GetBatchFromSvrPort(svrPorts []int64) (results []*AllTenantPartitionMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantPartitionMetaTable{}).Where("`svr_port` IN (?)", svrPorts).Find(&results).Error

	return
}

// GetFromRole 通过role获取内容
func (obj *_AllTenantPartitionMetaTableMgr) GetFromRole(role int64) (results []*AllTenantPartitionMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantPartitionMetaTable{}).Where("`role` = ?", role).Find(&results).Error

	return
}

// GetBatchFromRole 批量查找
func (obj *_AllTenantPartitionMetaTableMgr) GetBatchFromRole(roles []int64) (results []*AllTenantPartitionMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantPartitionMetaTable{}).Where("`role` IN (?)", roles).Find(&results).Error

	return
}

// GetFromRowCount 通过row_count获取内容
func (obj *_AllTenantPartitionMetaTableMgr) GetFromRowCount(rowCount int64) (results []*AllTenantPartitionMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantPartitionMetaTable{}).Where("`row_count` = ?", rowCount).Find(&results).Error

	return
}

// GetBatchFromRowCount 批量查找
func (obj *_AllTenantPartitionMetaTableMgr) GetBatchFromRowCount(rowCounts []int64) (results []*AllTenantPartitionMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantPartitionMetaTable{}).Where("`row_count` IN (?)", rowCounts).Find(&results).Error

	return
}

// GetFromDataSize 通过data_size获取内容
func (obj *_AllTenantPartitionMetaTableMgr) GetFromDataSize(dataSize int64) (results []*AllTenantPartitionMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantPartitionMetaTable{}).Where("`data_size` = ?", dataSize).Find(&results).Error

	return
}

// GetBatchFromDataSize 批量查找
func (obj *_AllTenantPartitionMetaTableMgr) GetBatchFromDataSize(dataSizes []int64) (results []*AllTenantPartitionMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantPartitionMetaTable{}).Where("`data_size` IN (?)", dataSizes).Find(&results).Error

	return
}

// GetFromDataVersion 通过data_version获取内容
func (obj *_AllTenantPartitionMetaTableMgr) GetFromDataVersion(dataVersion int64) (results []*AllTenantPartitionMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantPartitionMetaTable{}).Where("`data_version` = ?", dataVersion).Find(&results).Error

	return
}

// GetBatchFromDataVersion 批量查找
func (obj *_AllTenantPartitionMetaTableMgr) GetBatchFromDataVersion(dataVersions []int64) (results []*AllTenantPartitionMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantPartitionMetaTable{}).Where("`data_version` IN (?)", dataVersions).Find(&results).Error

	return
}

// GetFromRequiredSize 通过required_size获取内容
func (obj *_AllTenantPartitionMetaTableMgr) GetFromRequiredSize(requiredSize int64) (results []*AllTenantPartitionMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantPartitionMetaTable{}).Where("`required_size` = ?", requiredSize).Find(&results).Error

	return
}

// GetBatchFromRequiredSize 批量查找
func (obj *_AllTenantPartitionMetaTableMgr) GetBatchFromRequiredSize(requiredSizes []int64) (results []*AllTenantPartitionMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantPartitionMetaTable{}).Where("`required_size` IN (?)", requiredSizes).Find(&results).Error

	return
}

// GetFromReplicaType 通过replica_type获取内容
func (obj *_AllTenantPartitionMetaTableMgr) GetFromReplicaType(replicaType int64) (results []*AllTenantPartitionMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantPartitionMetaTable{}).Where("`replica_type` = ?", replicaType).Find(&results).Error

	return
}

// GetBatchFromReplicaType 批量查找
func (obj *_AllTenantPartitionMetaTableMgr) GetBatchFromReplicaType(replicaTypes []int64) (results []*AllTenantPartitionMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantPartitionMetaTable{}).Where("`replica_type` IN (?)", replicaTypes).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容
func (obj *_AllTenantPartitionMetaTableMgr) GetFromStatus(status string) (results []*AllTenantPartitionMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantPartitionMetaTable{}).Where("`status` = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量查找
func (obj *_AllTenantPartitionMetaTableMgr) GetBatchFromStatus(statuss []string) (results []*AllTenantPartitionMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantPartitionMetaTable{}).Where("`status` IN (?)", statuss).Find(&results).Error

	return
}

// GetFromDataChecksum 通过data_checksum获取内容
func (obj *_AllTenantPartitionMetaTableMgr) GetFromDataChecksum(dataChecksum int64) (results []*AllTenantPartitionMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantPartitionMetaTable{}).Where("`data_checksum` = ?", dataChecksum).Find(&results).Error

	return
}

// GetBatchFromDataChecksum 批量查找
func (obj *_AllTenantPartitionMetaTableMgr) GetBatchFromDataChecksum(dataChecksums []int64) (results []*AllTenantPartitionMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantPartitionMetaTable{}).Where("`data_checksum` IN (?)", dataChecksums).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllTenantPartitionMetaTableMgr) FetchByPrimaryKey(tenantID int64, tableID int64, partitionID int64, svrIP string, svrPort int64) (result AllTenantPartitionMetaTable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantPartitionMetaTable{}).Where("`tenant_id` = ? AND `table_id` = ? AND `partition_id` = ? AND `svr_ip` = ? AND `svr_port` = ?", tenantID, tableID, partitionID, svrIP, svrPort).First(&result).Error

	return
}
