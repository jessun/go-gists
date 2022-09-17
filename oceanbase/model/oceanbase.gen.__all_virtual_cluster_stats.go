package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _AllVirtualClusterStatsMgr struct {
	*_BaseMgr
}

// AllVirtualClusterStatsMgr open func
func AllVirtualClusterStatsMgr(db *gorm.DB) *_AllVirtualClusterStatsMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualClusterStatsMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualClusterStatsMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_cluster_stats"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualClusterStatsMgr) GetTableName() string {
	return "__all_virtual_cluster_stats"
}

// Reset 重置gorm会话
func (obj *_AllVirtualClusterStatsMgr) Reset() *_AllVirtualClusterStatsMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualClusterStatsMgr) Get() (result AllVirtualClusterStats, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualClusterStats{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualClusterStatsMgr) Gets() (results []*AllVirtualClusterStats, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualClusterStats{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualClusterStatsMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualClusterStats{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTenantID tenant_id获取
func (obj *_AllVirtualClusterStatsMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithRefreshedSchemaVersion refreshed_schema_version获取
func (obj *_AllVirtualClusterStatsMgr) WithRefreshedSchemaVersion(refreshedSchemaVersion int64) Option {
	return optionFunc(func(o *options) { o.query["refreshed_schema_version"] = refreshedSchemaVersion })
}

// WithDdlLag ddl_lag获取
func (obj *_AllVirtualClusterStatsMgr) WithDdlLag(ddlLag int64) Option {
	return optionFunc(func(o *options) { o.query["ddl_lag"] = ddlLag })
}

// WithMinSysTableScn min_sys_table_scn获取
func (obj *_AllVirtualClusterStatsMgr) WithMinSysTableScn(minSysTableScn int64) Option {
	return optionFunc(func(o *options) { o.query["min_sys_table_scn"] = minSysTableScn })
}

// WithMinUserTableScn min_user_table_scn获取
func (obj *_AllVirtualClusterStatsMgr) WithMinUserTableScn(minUserTableScn int64) Option {
	return optionFunc(func(o *options) { o.query["min_user_table_scn"] = minUserTableScn })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualClusterStatsMgr) GetByOption(opts ...Option) (result AllVirtualClusterStats, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualClusterStats{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualClusterStatsMgr) GetByOptions(opts ...Option) (results []*AllVirtualClusterStats, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualClusterStats{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualClusterStatsMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualClusterStats, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualClusterStats{}).Where(options.query)
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
func (obj *_AllVirtualClusterStatsMgr) GetFromTenantID(tenantID int64) (result AllVirtualClusterStats, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualClusterStats{}).Where("`tenant_id` = ?", tenantID).First(&result).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllVirtualClusterStatsMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllVirtualClusterStats, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualClusterStats{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromRefreshedSchemaVersion 通过refreshed_schema_version获取内容
func (obj *_AllVirtualClusterStatsMgr) GetFromRefreshedSchemaVersion(refreshedSchemaVersion int64) (results []*AllVirtualClusterStats, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualClusterStats{}).Where("`refreshed_schema_version` = ?", refreshedSchemaVersion).Find(&results).Error

	return
}

// GetBatchFromRefreshedSchemaVersion 批量查找
func (obj *_AllVirtualClusterStatsMgr) GetBatchFromRefreshedSchemaVersion(refreshedSchemaVersions []int64) (results []*AllVirtualClusterStats, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualClusterStats{}).Where("`refreshed_schema_version` IN (?)", refreshedSchemaVersions).Find(&results).Error

	return
}

// GetFromDdlLag 通过ddl_lag获取内容
func (obj *_AllVirtualClusterStatsMgr) GetFromDdlLag(ddlLag int64) (results []*AllVirtualClusterStats, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualClusterStats{}).Where("`ddl_lag` = ?", ddlLag).Find(&results).Error

	return
}

// GetBatchFromDdlLag 批量查找
func (obj *_AllVirtualClusterStatsMgr) GetBatchFromDdlLag(ddlLags []int64) (results []*AllVirtualClusterStats, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualClusterStats{}).Where("`ddl_lag` IN (?)", ddlLags).Find(&results).Error

	return
}

// GetFromMinSysTableScn 通过min_sys_table_scn获取内容
func (obj *_AllVirtualClusterStatsMgr) GetFromMinSysTableScn(minSysTableScn int64) (results []*AllVirtualClusterStats, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualClusterStats{}).Where("`min_sys_table_scn` = ?", minSysTableScn).Find(&results).Error

	return
}

// GetBatchFromMinSysTableScn 批量查找
func (obj *_AllVirtualClusterStatsMgr) GetBatchFromMinSysTableScn(minSysTableScns []int64) (results []*AllVirtualClusterStats, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualClusterStats{}).Where("`min_sys_table_scn` IN (?)", minSysTableScns).Find(&results).Error

	return
}

// GetFromMinUserTableScn 通过min_user_table_scn获取内容
func (obj *_AllVirtualClusterStatsMgr) GetFromMinUserTableScn(minUserTableScn int64) (results []*AllVirtualClusterStats, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualClusterStats{}).Where("`min_user_table_scn` = ?", minUserTableScn).Find(&results).Error

	return
}

// GetBatchFromMinUserTableScn 批量查找
func (obj *_AllVirtualClusterStatsMgr) GetBatchFromMinUserTableScn(minUserTableScns []int64) (results []*AllVirtualClusterStats, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualClusterStats{}).Where("`min_user_table_scn` IN (?)", minUserTableScns).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllVirtualClusterStatsMgr) FetchByPrimaryKey(tenantID int64) (result AllVirtualClusterStats, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualClusterStats{}).Where("`tenant_id` = ?", tenantID).First(&result).Error

	return
}
