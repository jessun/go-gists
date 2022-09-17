package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _AllVirtualFuncHistoryMgr struct {
	*_BaseMgr
}

// AllVirtualFuncHistoryMgr open func
func AllVirtualFuncHistoryMgr(db *gorm.DB) *_AllVirtualFuncHistoryMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualFuncHistoryMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualFuncHistoryMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_func_history"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualFuncHistoryMgr) GetTableName() string {
	return "__all_virtual_func_history"
}

// Reset 重置gorm会话
func (obj *_AllVirtualFuncHistoryMgr) Reset() *_AllVirtualFuncHistoryMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualFuncHistoryMgr) Get() (result AllVirtualFuncHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualFuncHistory{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualFuncHistoryMgr) Gets() (results []*AllVirtualFuncHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualFuncHistory{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualFuncHistoryMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualFuncHistory{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTenantID tenant_id获取
func (obj *_AllVirtualFuncHistoryMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithName name获取
func (obj *_AllVirtualFuncHistoryMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithSchemaVersion schema_version获取
func (obj *_AllVirtualFuncHistoryMgr) WithSchemaVersion(schemaVersion int64) Option {
	return optionFunc(func(o *options) { o.query["schema_version"] = schemaVersion })
}

// WithGmtCreate gmt_create获取
func (obj *_AllVirtualFuncHistoryMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllVirtualFuncHistoryMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithIsDeleted is_deleted获取
func (obj *_AllVirtualFuncHistoryMgr) WithIsDeleted(isDeleted int64) Option {
	return optionFunc(func(o *options) { o.query["is_deleted"] = isDeleted })
}

// WithRet ret获取
func (obj *_AllVirtualFuncHistoryMgr) WithRet(ret int64) Option {
	return optionFunc(func(o *options) { o.query["ret"] = ret })
}

// WithDl dl获取
func (obj *_AllVirtualFuncHistoryMgr) WithDl(dl string) Option {
	return optionFunc(func(o *options) { o.query["dl"] = dl })
}

// WithUdfID udf_id获取
func (obj *_AllVirtualFuncHistoryMgr) WithUdfID(udfID int64) Option {
	return optionFunc(func(o *options) { o.query["udf_id"] = udfID })
}

// WithType type获取
func (obj *_AllVirtualFuncHistoryMgr) WithType(_type int64) Option {
	return optionFunc(func(o *options) { o.query["type"] = _type })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualFuncHistoryMgr) GetByOption(opts ...Option) (result AllVirtualFuncHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualFuncHistory{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualFuncHistoryMgr) GetByOptions(opts ...Option) (results []*AllVirtualFuncHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualFuncHistory{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualFuncHistoryMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualFuncHistory, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualFuncHistory{}).Where(options.query)
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
func (obj *_AllVirtualFuncHistoryMgr) GetFromTenantID(tenantID int64) (results []*AllVirtualFuncHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualFuncHistory{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllVirtualFuncHistoryMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllVirtualFuncHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualFuncHistory{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromName 通过name获取内容
func (obj *_AllVirtualFuncHistoryMgr) GetFromName(name string) (results []*AllVirtualFuncHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualFuncHistory{}).Where("`name` = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量查找
func (obj *_AllVirtualFuncHistoryMgr) GetBatchFromName(names []string) (results []*AllVirtualFuncHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualFuncHistory{}).Where("`name` IN (?)", names).Find(&results).Error

	return
}

// GetFromSchemaVersion 通过schema_version获取内容
func (obj *_AllVirtualFuncHistoryMgr) GetFromSchemaVersion(schemaVersion int64) (results []*AllVirtualFuncHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualFuncHistory{}).Where("`schema_version` = ?", schemaVersion).Find(&results).Error

	return
}

// GetBatchFromSchemaVersion 批量查找
func (obj *_AllVirtualFuncHistoryMgr) GetBatchFromSchemaVersion(schemaVersions []int64) (results []*AllVirtualFuncHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualFuncHistory{}).Where("`schema_version` IN (?)", schemaVersions).Find(&results).Error

	return
}

// GetFromGmtCreate 通过gmt_create获取内容
func (obj *_AllVirtualFuncHistoryMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllVirtualFuncHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualFuncHistory{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllVirtualFuncHistoryMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllVirtualFuncHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualFuncHistory{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllVirtualFuncHistoryMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllVirtualFuncHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualFuncHistory{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllVirtualFuncHistoryMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllVirtualFuncHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualFuncHistory{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromIsDeleted 通过is_deleted获取内容
func (obj *_AllVirtualFuncHistoryMgr) GetFromIsDeleted(isDeleted int64) (results []*AllVirtualFuncHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualFuncHistory{}).Where("`is_deleted` = ?", isDeleted).Find(&results).Error

	return
}

// GetBatchFromIsDeleted 批量查找
func (obj *_AllVirtualFuncHistoryMgr) GetBatchFromIsDeleted(isDeleteds []int64) (results []*AllVirtualFuncHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualFuncHistory{}).Where("`is_deleted` IN (?)", isDeleteds).Find(&results).Error

	return
}

// GetFromRet 通过ret获取内容
func (obj *_AllVirtualFuncHistoryMgr) GetFromRet(ret int64) (results []*AllVirtualFuncHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualFuncHistory{}).Where("`ret` = ?", ret).Find(&results).Error

	return
}

// GetBatchFromRet 批量查找
func (obj *_AllVirtualFuncHistoryMgr) GetBatchFromRet(rets []int64) (results []*AllVirtualFuncHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualFuncHistory{}).Where("`ret` IN (?)", rets).Find(&results).Error

	return
}

// GetFromDl 通过dl获取内容
func (obj *_AllVirtualFuncHistoryMgr) GetFromDl(dl string) (results []*AllVirtualFuncHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualFuncHistory{}).Where("`dl` = ?", dl).Find(&results).Error

	return
}

// GetBatchFromDl 批量查找
func (obj *_AllVirtualFuncHistoryMgr) GetBatchFromDl(dls []string) (results []*AllVirtualFuncHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualFuncHistory{}).Where("`dl` IN (?)", dls).Find(&results).Error

	return
}

// GetFromUdfID 通过udf_id获取内容
func (obj *_AllVirtualFuncHistoryMgr) GetFromUdfID(udfID int64) (results []*AllVirtualFuncHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualFuncHistory{}).Where("`udf_id` = ?", udfID).Find(&results).Error

	return
}

// GetBatchFromUdfID 批量查找
func (obj *_AllVirtualFuncHistoryMgr) GetBatchFromUdfID(udfIDs []int64) (results []*AllVirtualFuncHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualFuncHistory{}).Where("`udf_id` IN (?)", udfIDs).Find(&results).Error

	return
}

// GetFromType 通过type获取内容
func (obj *_AllVirtualFuncHistoryMgr) GetFromType(_type int64) (results []*AllVirtualFuncHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualFuncHistory{}).Where("`type` = ?", _type).Find(&results).Error

	return
}

// GetBatchFromType 批量查找
func (obj *_AllVirtualFuncHistoryMgr) GetBatchFromType(_types []int64) (results []*AllVirtualFuncHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualFuncHistory{}).Where("`type` IN (?)", _types).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllVirtualFuncHistoryMgr) FetchByPrimaryKey(tenantID int64, name string, schemaVersion int64) (result AllVirtualFuncHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualFuncHistory{}).Where("`tenant_id` = ? AND `name` = ? AND `schema_version` = ?", tenantID, name, schemaVersion).First(&result).Error

	return
}
