package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _AllVirtualSysStatMgr struct {
	*_BaseMgr
}

// AllVirtualSysStatMgr open func
func AllVirtualSysStatMgr(db *gorm.DB) *_AllVirtualSysStatMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualSysStatMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualSysStatMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_sys_stat"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualSysStatMgr) GetTableName() string {
	return "__all_virtual_sys_stat"
}

// Reset 重置gorm会话
func (obj *_AllVirtualSysStatMgr) Reset() *_AllVirtualSysStatMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualSysStatMgr) Get() (result AllVirtualSysStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSysStat{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualSysStatMgr) Gets() (results []*AllVirtualSysStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSysStat{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualSysStatMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualSysStat{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTenantID tenant_id获取
func (obj *_AllVirtualSysStatMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithZone zone获取
func (obj *_AllVirtualSysStatMgr) WithZone(zone string) Option {
	return optionFunc(func(o *options) { o.query["zone"] = zone })
}

// WithName name获取
func (obj *_AllVirtualSysStatMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithGmtCreate gmt_create获取
func (obj *_AllVirtualSysStatMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllVirtualSysStatMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithDataType data_type获取
func (obj *_AllVirtualSysStatMgr) WithDataType(dataType int64) Option {
	return optionFunc(func(o *options) { o.query["data_type"] = dataType })
}

// WithValue value获取
func (obj *_AllVirtualSysStatMgr) WithValue(value string) Option {
	return optionFunc(func(o *options) { o.query["value"] = value })
}

// WithInfo info获取
func (obj *_AllVirtualSysStatMgr) WithInfo(info string) Option {
	return optionFunc(func(o *options) { o.query["info"] = info })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualSysStatMgr) GetByOption(opts ...Option) (result AllVirtualSysStat, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSysStat{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualSysStatMgr) GetByOptions(opts ...Option) (results []*AllVirtualSysStat, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSysStat{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualSysStatMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualSysStat, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualSysStat{}).Where(options.query)
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
func (obj *_AllVirtualSysStatMgr) GetFromTenantID(tenantID int64) (results []*AllVirtualSysStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSysStat{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllVirtualSysStatMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllVirtualSysStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSysStat{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromZone 通过zone获取内容
func (obj *_AllVirtualSysStatMgr) GetFromZone(zone string) (results []*AllVirtualSysStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSysStat{}).Where("`zone` = ?", zone).Find(&results).Error

	return
}

// GetBatchFromZone 批量查找
func (obj *_AllVirtualSysStatMgr) GetBatchFromZone(zones []string) (results []*AllVirtualSysStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSysStat{}).Where("`zone` IN (?)", zones).Find(&results).Error

	return
}

// GetFromName 通过name获取内容
func (obj *_AllVirtualSysStatMgr) GetFromName(name string) (results []*AllVirtualSysStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSysStat{}).Where("`name` = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量查找
func (obj *_AllVirtualSysStatMgr) GetBatchFromName(names []string) (results []*AllVirtualSysStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSysStat{}).Where("`name` IN (?)", names).Find(&results).Error

	return
}

// GetFromGmtCreate 通过gmt_create获取内容
func (obj *_AllVirtualSysStatMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllVirtualSysStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSysStat{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllVirtualSysStatMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllVirtualSysStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSysStat{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllVirtualSysStatMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllVirtualSysStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSysStat{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllVirtualSysStatMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllVirtualSysStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSysStat{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromDataType 通过data_type获取内容
func (obj *_AllVirtualSysStatMgr) GetFromDataType(dataType int64) (results []*AllVirtualSysStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSysStat{}).Where("`data_type` = ?", dataType).Find(&results).Error

	return
}

// GetBatchFromDataType 批量查找
func (obj *_AllVirtualSysStatMgr) GetBatchFromDataType(dataTypes []int64) (results []*AllVirtualSysStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSysStat{}).Where("`data_type` IN (?)", dataTypes).Find(&results).Error

	return
}

// GetFromValue 通过value获取内容
func (obj *_AllVirtualSysStatMgr) GetFromValue(value string) (results []*AllVirtualSysStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSysStat{}).Where("`value` = ?", value).Find(&results).Error

	return
}

// GetBatchFromValue 批量查找
func (obj *_AllVirtualSysStatMgr) GetBatchFromValue(values []string) (results []*AllVirtualSysStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSysStat{}).Where("`value` IN (?)", values).Find(&results).Error

	return
}

// GetFromInfo 通过info获取内容
func (obj *_AllVirtualSysStatMgr) GetFromInfo(info string) (results []*AllVirtualSysStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSysStat{}).Where("`info` = ?", info).Find(&results).Error

	return
}

// GetBatchFromInfo 批量查找
func (obj *_AllVirtualSysStatMgr) GetBatchFromInfo(infos []string) (results []*AllVirtualSysStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSysStat{}).Where("`info` IN (?)", infos).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllVirtualSysStatMgr) FetchByPrimaryKey(tenantID int64, zone string, name string) (result AllVirtualSysStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualSysStat{}).Where("`tenant_id` = ? AND `zone` = ? AND `name` = ?", tenantID, zone, name).First(&result).Error

	return
}
