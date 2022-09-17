package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _AllVirtualServerSchemaInfoMgr struct {
	*_BaseMgr
}

// AllVirtualServerSchemaInfoMgr open func
func AllVirtualServerSchemaInfoMgr(db *gorm.DB) *_AllVirtualServerSchemaInfoMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualServerSchemaInfoMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualServerSchemaInfoMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_server_schema_info"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualServerSchemaInfoMgr) GetTableName() string {
	return "__all_virtual_server_schema_info"
}

// Reset 重置gorm会话
func (obj *_AllVirtualServerSchemaInfoMgr) Reset() *_AllVirtualServerSchemaInfoMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualServerSchemaInfoMgr) Get() (result AllVirtualServerSchemaInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerSchemaInfo{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualServerSchemaInfoMgr) Gets() (results []*AllVirtualServerSchemaInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerSchemaInfo{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualServerSchemaInfoMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualServerSchemaInfo{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithSvrIP svr_ip获取
func (obj *_AllVirtualServerSchemaInfoMgr) WithSvrIP(svrIP string) Option {
	return optionFunc(func(o *options) { o.query["svr_ip"] = svrIP })
}

// WithSvrPort svr_port获取
func (obj *_AllVirtualServerSchemaInfoMgr) WithSvrPort(svrPort int64) Option {
	return optionFunc(func(o *options) { o.query["svr_port"] = svrPort })
}

// WithTenantID tenant_id获取
func (obj *_AllVirtualServerSchemaInfoMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithRefreshedSchemaVersion refreshed_schema_version获取
func (obj *_AllVirtualServerSchemaInfoMgr) WithRefreshedSchemaVersion(refreshedSchemaVersion int64) Option {
	return optionFunc(func(o *options) { o.query["refreshed_schema_version"] = refreshedSchemaVersion })
}

// WithReceivedSchemaVersion received_schema_version获取
func (obj *_AllVirtualServerSchemaInfoMgr) WithReceivedSchemaVersion(receivedSchemaVersion int64) Option {
	return optionFunc(func(o *options) { o.query["received_schema_version"] = receivedSchemaVersion })
}

// WithSchemaCount schema_count获取
func (obj *_AllVirtualServerSchemaInfoMgr) WithSchemaCount(schemaCount int64) Option {
	return optionFunc(func(o *options) { o.query["schema_count"] = schemaCount })
}

// WithSchemaSize schema_size获取
func (obj *_AllVirtualServerSchemaInfoMgr) WithSchemaSize(schemaSize int64) Option {
	return optionFunc(func(o *options) { o.query["schema_size"] = schemaSize })
}

// WithMinSstableSchemaVersion min_sstable_schema_version获取
func (obj *_AllVirtualServerSchemaInfoMgr) WithMinSstableSchemaVersion(minSstableSchemaVersion int64) Option {
	return optionFunc(func(o *options) { o.query["min_sstable_schema_version"] = minSstableSchemaVersion })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualServerSchemaInfoMgr) GetByOption(opts ...Option) (result AllVirtualServerSchemaInfo, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerSchemaInfo{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualServerSchemaInfoMgr) GetByOptions(opts ...Option) (results []*AllVirtualServerSchemaInfo, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerSchemaInfo{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualServerSchemaInfoMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualServerSchemaInfo, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualServerSchemaInfo{}).Where(options.query)
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
func (obj *_AllVirtualServerSchemaInfoMgr) GetFromSvrIP(svrIP string) (results []*AllVirtualServerSchemaInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerSchemaInfo{}).Where("`svr_ip` = ?", svrIP).Find(&results).Error

	return
}

// GetBatchFromSvrIP 批量查找
func (obj *_AllVirtualServerSchemaInfoMgr) GetBatchFromSvrIP(svrIPs []string) (results []*AllVirtualServerSchemaInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerSchemaInfo{}).Where("`svr_ip` IN (?)", svrIPs).Find(&results).Error

	return
}

// GetFromSvrPort 通过svr_port获取内容
func (obj *_AllVirtualServerSchemaInfoMgr) GetFromSvrPort(svrPort int64) (results []*AllVirtualServerSchemaInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerSchemaInfo{}).Where("`svr_port` = ?", svrPort).Find(&results).Error

	return
}

// GetBatchFromSvrPort 批量查找
func (obj *_AllVirtualServerSchemaInfoMgr) GetBatchFromSvrPort(svrPorts []int64) (results []*AllVirtualServerSchemaInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerSchemaInfo{}).Where("`svr_port` IN (?)", svrPorts).Find(&results).Error

	return
}

// GetFromTenantID 通过tenant_id获取内容
func (obj *_AllVirtualServerSchemaInfoMgr) GetFromTenantID(tenantID int64) (results []*AllVirtualServerSchemaInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerSchemaInfo{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllVirtualServerSchemaInfoMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllVirtualServerSchemaInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerSchemaInfo{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromRefreshedSchemaVersion 通过refreshed_schema_version获取内容
func (obj *_AllVirtualServerSchemaInfoMgr) GetFromRefreshedSchemaVersion(refreshedSchemaVersion int64) (results []*AllVirtualServerSchemaInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerSchemaInfo{}).Where("`refreshed_schema_version` = ?", refreshedSchemaVersion).Find(&results).Error

	return
}

// GetBatchFromRefreshedSchemaVersion 批量查找
func (obj *_AllVirtualServerSchemaInfoMgr) GetBatchFromRefreshedSchemaVersion(refreshedSchemaVersions []int64) (results []*AllVirtualServerSchemaInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerSchemaInfo{}).Where("`refreshed_schema_version` IN (?)", refreshedSchemaVersions).Find(&results).Error

	return
}

// GetFromReceivedSchemaVersion 通过received_schema_version获取内容
func (obj *_AllVirtualServerSchemaInfoMgr) GetFromReceivedSchemaVersion(receivedSchemaVersion int64) (results []*AllVirtualServerSchemaInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerSchemaInfo{}).Where("`received_schema_version` = ?", receivedSchemaVersion).Find(&results).Error

	return
}

// GetBatchFromReceivedSchemaVersion 批量查找
func (obj *_AllVirtualServerSchemaInfoMgr) GetBatchFromReceivedSchemaVersion(receivedSchemaVersions []int64) (results []*AllVirtualServerSchemaInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerSchemaInfo{}).Where("`received_schema_version` IN (?)", receivedSchemaVersions).Find(&results).Error

	return
}

// GetFromSchemaCount 通过schema_count获取内容
func (obj *_AllVirtualServerSchemaInfoMgr) GetFromSchemaCount(schemaCount int64) (results []*AllVirtualServerSchemaInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerSchemaInfo{}).Where("`schema_count` = ?", schemaCount).Find(&results).Error

	return
}

// GetBatchFromSchemaCount 批量查找
func (obj *_AllVirtualServerSchemaInfoMgr) GetBatchFromSchemaCount(schemaCounts []int64) (results []*AllVirtualServerSchemaInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerSchemaInfo{}).Where("`schema_count` IN (?)", schemaCounts).Find(&results).Error

	return
}

// GetFromSchemaSize 通过schema_size获取内容
func (obj *_AllVirtualServerSchemaInfoMgr) GetFromSchemaSize(schemaSize int64) (results []*AllVirtualServerSchemaInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerSchemaInfo{}).Where("`schema_size` = ?", schemaSize).Find(&results).Error

	return
}

// GetBatchFromSchemaSize 批量查找
func (obj *_AllVirtualServerSchemaInfoMgr) GetBatchFromSchemaSize(schemaSizes []int64) (results []*AllVirtualServerSchemaInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerSchemaInfo{}).Where("`schema_size` IN (?)", schemaSizes).Find(&results).Error

	return
}

// GetFromMinSstableSchemaVersion 通过min_sstable_schema_version获取内容
func (obj *_AllVirtualServerSchemaInfoMgr) GetFromMinSstableSchemaVersion(minSstableSchemaVersion int64) (results []*AllVirtualServerSchemaInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerSchemaInfo{}).Where("`min_sstable_schema_version` = ?", minSstableSchemaVersion).Find(&results).Error

	return
}

// GetBatchFromMinSstableSchemaVersion 批量查找
func (obj *_AllVirtualServerSchemaInfoMgr) GetBatchFromMinSstableSchemaVersion(minSstableSchemaVersions []int64) (results []*AllVirtualServerSchemaInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualServerSchemaInfo{}).Where("`min_sstable_schema_version` IN (?)", minSstableSchemaVersions).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////
