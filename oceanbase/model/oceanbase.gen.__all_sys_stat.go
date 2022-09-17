package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _AllSysStatMgr struct {
	*_BaseMgr
}

// AllSysStatMgr open func
func AllSysStatMgr(db *gorm.DB) *_AllSysStatMgr {
	if db == nil {
		panic(fmt.Errorf("AllSysStatMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllSysStatMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_sys_stat"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllSysStatMgr) GetTableName() string {
	return "__all_sys_stat"
}

// Reset 重置gorm会话
func (obj *_AllSysStatMgr) Reset() *_AllSysStatMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllSysStatMgr) Get() (result AllSysStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSysStat{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllSysStatMgr) Gets() (results []*AllSysStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSysStat{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllSysStatMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllSysStat{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithGmtCreate gmt_create获取
func (obj *_AllSysStatMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllSysStatMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithTenantID tenant_id获取
func (obj *_AllSysStatMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithZone zone获取
func (obj *_AllSysStatMgr) WithZone(zone string) Option {
	return optionFunc(func(o *options) { o.query["zone"] = zone })
}

// WithName name获取
func (obj *_AllSysStatMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithDataType data_type获取
func (obj *_AllSysStatMgr) WithDataType(dataType int64) Option {
	return optionFunc(func(o *options) { o.query["data_type"] = dataType })
}

// WithValue value获取
func (obj *_AllSysStatMgr) WithValue(value string) Option {
	return optionFunc(func(o *options) { o.query["value"] = value })
}

// WithInfo info获取
func (obj *_AllSysStatMgr) WithInfo(info string) Option {
	return optionFunc(func(o *options) { o.query["info"] = info })
}

// GetByOption 功能选项模式获取
func (obj *_AllSysStatMgr) GetByOption(opts ...Option) (result AllSysStat, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllSysStat{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllSysStatMgr) GetByOptions(opts ...Option) (results []*AllSysStat, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllSysStat{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllSysStatMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllSysStat, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllSysStat{}).Where(options.query)
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
func (obj *_AllSysStatMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllSysStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSysStat{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllSysStatMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllSysStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSysStat{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllSysStatMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllSysStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSysStat{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllSysStatMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllSysStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSysStat{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromTenantID 通过tenant_id获取内容
func (obj *_AllSysStatMgr) GetFromTenantID(tenantID int64) (results []*AllSysStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSysStat{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllSysStatMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllSysStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSysStat{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromZone 通过zone获取内容
func (obj *_AllSysStatMgr) GetFromZone(zone string) (results []*AllSysStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSysStat{}).Where("`zone` = ?", zone).Find(&results).Error

	return
}

// GetBatchFromZone 批量查找
func (obj *_AllSysStatMgr) GetBatchFromZone(zones []string) (results []*AllSysStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSysStat{}).Where("`zone` IN (?)", zones).Find(&results).Error

	return
}

// GetFromName 通过name获取内容
func (obj *_AllSysStatMgr) GetFromName(name string) (results []*AllSysStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSysStat{}).Where("`name` = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量查找
func (obj *_AllSysStatMgr) GetBatchFromName(names []string) (results []*AllSysStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSysStat{}).Where("`name` IN (?)", names).Find(&results).Error

	return
}

// GetFromDataType 通过data_type获取内容
func (obj *_AllSysStatMgr) GetFromDataType(dataType int64) (results []*AllSysStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSysStat{}).Where("`data_type` = ?", dataType).Find(&results).Error

	return
}

// GetBatchFromDataType 批量查找
func (obj *_AllSysStatMgr) GetBatchFromDataType(dataTypes []int64) (results []*AllSysStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSysStat{}).Where("`data_type` IN (?)", dataTypes).Find(&results).Error

	return
}

// GetFromValue 通过value获取内容
func (obj *_AllSysStatMgr) GetFromValue(value string) (results []*AllSysStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSysStat{}).Where("`value` = ?", value).Find(&results).Error

	return
}

// GetBatchFromValue 批量查找
func (obj *_AllSysStatMgr) GetBatchFromValue(values []string) (results []*AllSysStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSysStat{}).Where("`value` IN (?)", values).Find(&results).Error

	return
}

// GetFromInfo 通过info获取内容
func (obj *_AllSysStatMgr) GetFromInfo(info string) (results []*AllSysStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSysStat{}).Where("`info` = ?", info).Find(&results).Error

	return
}

// GetBatchFromInfo 批量查找
func (obj *_AllSysStatMgr) GetBatchFromInfo(infos []string) (results []*AllSysStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSysStat{}).Where("`info` IN (?)", infos).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllSysStatMgr) FetchByPrimaryKey(tenantID int64, zone string, name string) (result AllSysStat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSysStat{}).Where("`tenant_id` = ? AND `zone` = ? AND `name` = ?", tenantID, zone, name).First(&result).Error

	return
}
