package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _AllVirtualReservedTableMgrMgr struct {
	*_BaseMgr
}

// AllVirtualReservedTableMgrMgr open func
func AllVirtualReservedTableMgrMgr(db *gorm.DB) *_AllVirtualReservedTableMgrMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualReservedTableMgrMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualReservedTableMgrMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_reserved_table_mgr"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualReservedTableMgrMgr) GetTableName() string {
	return "__all_virtual_reserved_table_mgr"
}

// Reset 重置gorm会话
func (obj *_AllVirtualReservedTableMgrMgr) Reset() *_AllVirtualReservedTableMgrMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualReservedTableMgrMgr) Get() (result AllVirtualReservedTableMgr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualReservedTableMgr{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualReservedTableMgrMgr) Gets() (results []*AllVirtualReservedTableMgr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualReservedTableMgr{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualReservedTableMgrMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualReservedTableMgr{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithSvrIP svr_ip获取
func (obj *_AllVirtualReservedTableMgrMgr) WithSvrIP(svrIP string) Option {
	return optionFunc(func(o *options) { o.query["svr_ip"] = svrIP })
}

// WithSvrPort svr_port获取
func (obj *_AllVirtualReservedTableMgrMgr) WithSvrPort(svrPort int64) Option {
	return optionFunc(func(o *options) { o.query["svr_port"] = svrPort })
}

// WithTenantID tenant_id获取
func (obj *_AllVirtualReservedTableMgrMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithTableID table_id获取
func (obj *_AllVirtualReservedTableMgrMgr) WithTableID(tableID int64) Option {
	return optionFunc(func(o *options) { o.query["table_id"] = tableID })
}

// WithTableType table_type获取
func (obj *_AllVirtualReservedTableMgrMgr) WithTableType(tableType int64) Option {
	return optionFunc(func(o *options) { o.query["table_type"] = tableType })
}

// WithPartitionID partition_id获取
func (obj *_AllVirtualReservedTableMgrMgr) WithPartitionID(partitionID int64) Option {
	return optionFunc(func(o *options) { o.query["partition_id"] = partitionID })
}

// WithIndexID index_id获取
func (obj *_AllVirtualReservedTableMgrMgr) WithIndexID(indexID int64) Option {
	return optionFunc(func(o *options) { o.query["index_id"] = indexID })
}

// WithBaseVersion base_version获取
func (obj *_AllVirtualReservedTableMgrMgr) WithBaseVersion(baseVersion int64) Option {
	return optionFunc(func(o *options) { o.query["base_version"] = baseVersion })
}

// WithMultiVersionStart multi_version_start获取
func (obj *_AllVirtualReservedTableMgrMgr) WithMultiVersionStart(multiVersionStart int64) Option {
	return optionFunc(func(o *options) { o.query["multi_version_start"] = multiVersionStart })
}

// WithSnapshotVersion snapshot_version获取
func (obj *_AllVirtualReservedTableMgrMgr) WithSnapshotVersion(snapshotVersion int64) Option {
	return optionFunc(func(o *options) { o.query["snapshot_version"] = snapshotVersion })
}

// WithVersion version获取
func (obj *_AllVirtualReservedTableMgrMgr) WithVersion(version int64) Option {
	return optionFunc(func(o *options) { o.query["version"] = version })
}

// WithSize size获取
func (obj *_AllVirtualReservedTableMgrMgr) WithSize(size int64) Option {
	return optionFunc(func(o *options) { o.query["size"] = size })
}

// WithRef ref获取
func (obj *_AllVirtualReservedTableMgrMgr) WithRef(ref int64) Option {
	return optionFunc(func(o *options) { o.query["ref"] = ref })
}

// WithReserveType reserve_type获取
func (obj *_AllVirtualReservedTableMgrMgr) WithReserveType(reserveType string) Option {
	return optionFunc(func(o *options) { o.query["reserve_type"] = reserveType })
}

// WithReservePointVersion reserve_point_version获取
func (obj *_AllVirtualReservedTableMgrMgr) WithReservePointVersion(reservePointVersion int64) Option {
	return optionFunc(func(o *options) { o.query["reserve_point_version"] = reservePointVersion })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualReservedTableMgrMgr) GetByOption(opts ...Option) (result AllVirtualReservedTableMgr, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualReservedTableMgr{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualReservedTableMgrMgr) GetByOptions(opts ...Option) (results []*AllVirtualReservedTableMgr, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualReservedTableMgr{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualReservedTableMgrMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualReservedTableMgr, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualReservedTableMgr{}).Where(options.query)
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
func (obj *_AllVirtualReservedTableMgrMgr) GetFromSvrIP(svrIP string) (results []*AllVirtualReservedTableMgr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualReservedTableMgr{}).Where("`svr_ip` = ?", svrIP).Find(&results).Error

	return
}

// GetBatchFromSvrIP 批量查找
func (obj *_AllVirtualReservedTableMgrMgr) GetBatchFromSvrIP(svrIPs []string) (results []*AllVirtualReservedTableMgr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualReservedTableMgr{}).Where("`svr_ip` IN (?)", svrIPs).Find(&results).Error

	return
}

// GetFromSvrPort 通过svr_port获取内容
func (obj *_AllVirtualReservedTableMgrMgr) GetFromSvrPort(svrPort int64) (results []*AllVirtualReservedTableMgr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualReservedTableMgr{}).Where("`svr_port` = ?", svrPort).Find(&results).Error

	return
}

// GetBatchFromSvrPort 批量查找
func (obj *_AllVirtualReservedTableMgrMgr) GetBatchFromSvrPort(svrPorts []int64) (results []*AllVirtualReservedTableMgr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualReservedTableMgr{}).Where("`svr_port` IN (?)", svrPorts).Find(&results).Error

	return
}

// GetFromTenantID 通过tenant_id获取内容
func (obj *_AllVirtualReservedTableMgrMgr) GetFromTenantID(tenantID int64) (results []*AllVirtualReservedTableMgr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualReservedTableMgr{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllVirtualReservedTableMgrMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllVirtualReservedTableMgr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualReservedTableMgr{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromTableID 通过table_id获取内容
func (obj *_AllVirtualReservedTableMgrMgr) GetFromTableID(tableID int64) (results []*AllVirtualReservedTableMgr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualReservedTableMgr{}).Where("`table_id` = ?", tableID).Find(&results).Error

	return
}

// GetBatchFromTableID 批量查找
func (obj *_AllVirtualReservedTableMgrMgr) GetBatchFromTableID(tableIDs []int64) (results []*AllVirtualReservedTableMgr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualReservedTableMgr{}).Where("`table_id` IN (?)", tableIDs).Find(&results).Error

	return
}

// GetFromTableType 通过table_type获取内容
func (obj *_AllVirtualReservedTableMgrMgr) GetFromTableType(tableType int64) (results []*AllVirtualReservedTableMgr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualReservedTableMgr{}).Where("`table_type` = ?", tableType).Find(&results).Error

	return
}

// GetBatchFromTableType 批量查找
func (obj *_AllVirtualReservedTableMgrMgr) GetBatchFromTableType(tableTypes []int64) (results []*AllVirtualReservedTableMgr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualReservedTableMgr{}).Where("`table_type` IN (?)", tableTypes).Find(&results).Error

	return
}

// GetFromPartitionID 通过partition_id获取内容
func (obj *_AllVirtualReservedTableMgrMgr) GetFromPartitionID(partitionID int64) (results []*AllVirtualReservedTableMgr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualReservedTableMgr{}).Where("`partition_id` = ?", partitionID).Find(&results).Error

	return
}

// GetBatchFromPartitionID 批量查找
func (obj *_AllVirtualReservedTableMgrMgr) GetBatchFromPartitionID(partitionIDs []int64) (results []*AllVirtualReservedTableMgr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualReservedTableMgr{}).Where("`partition_id` IN (?)", partitionIDs).Find(&results).Error

	return
}

// GetFromIndexID 通过index_id获取内容
func (obj *_AllVirtualReservedTableMgrMgr) GetFromIndexID(indexID int64) (results []*AllVirtualReservedTableMgr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualReservedTableMgr{}).Where("`index_id` = ?", indexID).Find(&results).Error

	return
}

// GetBatchFromIndexID 批量查找
func (obj *_AllVirtualReservedTableMgrMgr) GetBatchFromIndexID(indexIDs []int64) (results []*AllVirtualReservedTableMgr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualReservedTableMgr{}).Where("`index_id` IN (?)", indexIDs).Find(&results).Error

	return
}

// GetFromBaseVersion 通过base_version获取内容
func (obj *_AllVirtualReservedTableMgrMgr) GetFromBaseVersion(baseVersion int64) (results []*AllVirtualReservedTableMgr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualReservedTableMgr{}).Where("`base_version` = ?", baseVersion).Find(&results).Error

	return
}

// GetBatchFromBaseVersion 批量查找
func (obj *_AllVirtualReservedTableMgrMgr) GetBatchFromBaseVersion(baseVersions []int64) (results []*AllVirtualReservedTableMgr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualReservedTableMgr{}).Where("`base_version` IN (?)", baseVersions).Find(&results).Error

	return
}

// GetFromMultiVersionStart 通过multi_version_start获取内容
func (obj *_AllVirtualReservedTableMgrMgr) GetFromMultiVersionStart(multiVersionStart int64) (results []*AllVirtualReservedTableMgr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualReservedTableMgr{}).Where("`multi_version_start` = ?", multiVersionStart).Find(&results).Error

	return
}

// GetBatchFromMultiVersionStart 批量查找
func (obj *_AllVirtualReservedTableMgrMgr) GetBatchFromMultiVersionStart(multiVersionStarts []int64) (results []*AllVirtualReservedTableMgr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualReservedTableMgr{}).Where("`multi_version_start` IN (?)", multiVersionStarts).Find(&results).Error

	return
}

// GetFromSnapshotVersion 通过snapshot_version获取内容
func (obj *_AllVirtualReservedTableMgrMgr) GetFromSnapshotVersion(snapshotVersion int64) (results []*AllVirtualReservedTableMgr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualReservedTableMgr{}).Where("`snapshot_version` = ?", snapshotVersion).Find(&results).Error

	return
}

// GetBatchFromSnapshotVersion 批量查找
func (obj *_AllVirtualReservedTableMgrMgr) GetBatchFromSnapshotVersion(snapshotVersions []int64) (results []*AllVirtualReservedTableMgr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualReservedTableMgr{}).Where("`snapshot_version` IN (?)", snapshotVersions).Find(&results).Error

	return
}

// GetFromVersion 通过version获取内容
func (obj *_AllVirtualReservedTableMgrMgr) GetFromVersion(version int64) (results []*AllVirtualReservedTableMgr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualReservedTableMgr{}).Where("`version` = ?", version).Find(&results).Error

	return
}

// GetBatchFromVersion 批量查找
func (obj *_AllVirtualReservedTableMgrMgr) GetBatchFromVersion(versions []int64) (results []*AllVirtualReservedTableMgr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualReservedTableMgr{}).Where("`version` IN (?)", versions).Find(&results).Error

	return
}

// GetFromSize 通过size获取内容
func (obj *_AllVirtualReservedTableMgrMgr) GetFromSize(size int64) (results []*AllVirtualReservedTableMgr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualReservedTableMgr{}).Where("`size` = ?", size).Find(&results).Error

	return
}

// GetBatchFromSize 批量查找
func (obj *_AllVirtualReservedTableMgrMgr) GetBatchFromSize(sizes []int64) (results []*AllVirtualReservedTableMgr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualReservedTableMgr{}).Where("`size` IN (?)", sizes).Find(&results).Error

	return
}

// GetFromRef 通过ref获取内容
func (obj *_AllVirtualReservedTableMgrMgr) GetFromRef(ref int64) (results []*AllVirtualReservedTableMgr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualReservedTableMgr{}).Where("`ref` = ?", ref).Find(&results).Error

	return
}

// GetBatchFromRef 批量查找
func (obj *_AllVirtualReservedTableMgrMgr) GetBatchFromRef(refs []int64) (results []*AllVirtualReservedTableMgr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualReservedTableMgr{}).Where("`ref` IN (?)", refs).Find(&results).Error

	return
}

// GetFromReserveType 通过reserve_type获取内容
func (obj *_AllVirtualReservedTableMgrMgr) GetFromReserveType(reserveType string) (results []*AllVirtualReservedTableMgr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualReservedTableMgr{}).Where("`reserve_type` = ?", reserveType).Find(&results).Error

	return
}

// GetBatchFromReserveType 批量查找
func (obj *_AllVirtualReservedTableMgrMgr) GetBatchFromReserveType(reserveTypes []string) (results []*AllVirtualReservedTableMgr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualReservedTableMgr{}).Where("`reserve_type` IN (?)", reserveTypes).Find(&results).Error

	return
}

// GetFromReservePointVersion 通过reserve_point_version获取内容
func (obj *_AllVirtualReservedTableMgrMgr) GetFromReservePointVersion(reservePointVersion int64) (results []*AllVirtualReservedTableMgr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualReservedTableMgr{}).Where("`reserve_point_version` = ?", reservePointVersion).Find(&results).Error

	return
}

// GetBatchFromReservePointVersion 批量查找
func (obj *_AllVirtualReservedTableMgrMgr) GetBatchFromReservePointVersion(reservePointVersions []int64) (results []*AllVirtualReservedTableMgr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualReservedTableMgr{}).Where("`reserve_point_version` IN (?)", reservePointVersions).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllVirtualReservedTableMgrMgr) FetchByPrimaryKey(svrIP string, svrPort int64, tenantID int64, tableID int64, tableType int64, partitionID int64, indexID int64, baseVersion int64, multiVersionStart int64, snapshotVersion int64, version int64) (result AllVirtualReservedTableMgr, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualReservedTableMgr{}).Where("`svr_ip` = ? AND `svr_port` = ? AND `tenant_id` = ? AND `table_id` = ? AND `table_type` = ? AND `partition_id` = ? AND `index_id` = ? AND `base_version` = ? AND `multi_version_start` = ? AND `snapshot_version` = ? AND `version` = ?", svrIP, svrPort, tenantID, tableID, tableType, partitionID, indexID, baseVersion, multiVersionStart, snapshotVersion, version).First(&result).Error

	return
}
