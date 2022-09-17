package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _AllFuncHistoryMgr struct {
	*_BaseMgr
}

// AllFuncHistoryMgr open func
func AllFuncHistoryMgr(db *gorm.DB) *_AllFuncHistoryMgr {
	if db == nil {
		panic(fmt.Errorf("AllFuncHistoryMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllFuncHistoryMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_func_history"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllFuncHistoryMgr) GetTableName() string {
	return "__all_func_history"
}

// Reset 重置gorm会话
func (obj *_AllFuncHistoryMgr) Reset() *_AllFuncHistoryMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllFuncHistoryMgr) Get() (result AllFuncHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllFuncHistory{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllFuncHistoryMgr) Gets() (results []*AllFuncHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllFuncHistory{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllFuncHistoryMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllFuncHistory{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithGmtCreate gmt_create获取
func (obj *_AllFuncHistoryMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllFuncHistoryMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithTenantID tenant_id获取
func (obj *_AllFuncHistoryMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithName name获取
func (obj *_AllFuncHistoryMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithSchemaVersion schema_version获取
func (obj *_AllFuncHistoryMgr) WithSchemaVersion(schemaVersion int64) Option {
	return optionFunc(func(o *options) { o.query["schema_version"] = schemaVersion })
}

// WithIsDeleted is_deleted获取
func (obj *_AllFuncHistoryMgr) WithIsDeleted(isDeleted int64) Option {
	return optionFunc(func(o *options) { o.query["is_deleted"] = isDeleted })
}

// WithRet ret获取
func (obj *_AllFuncHistoryMgr) WithRet(ret int64) Option {
	return optionFunc(func(o *options) { o.query["ret"] = ret })
}

// WithDl dl获取
func (obj *_AllFuncHistoryMgr) WithDl(dl string) Option {
	return optionFunc(func(o *options) { o.query["dl"] = dl })
}

// WithUdfID udf_id获取
func (obj *_AllFuncHistoryMgr) WithUdfID(udfID int64) Option {
	return optionFunc(func(o *options) { o.query["udf_id"] = udfID })
}

// WithType type获取
func (obj *_AllFuncHistoryMgr) WithType(_type int64) Option {
	return optionFunc(func(o *options) { o.query["type"] = _type })
}

// GetByOption 功能选项模式获取
func (obj *_AllFuncHistoryMgr) GetByOption(opts ...Option) (result AllFuncHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllFuncHistory{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllFuncHistoryMgr) GetByOptions(opts ...Option) (results []*AllFuncHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllFuncHistory{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllFuncHistoryMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllFuncHistory, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllFuncHistory{}).Where(options.query)
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
func (obj *_AllFuncHistoryMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllFuncHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllFuncHistory{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllFuncHistoryMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllFuncHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllFuncHistory{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllFuncHistoryMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllFuncHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllFuncHistory{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllFuncHistoryMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllFuncHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllFuncHistory{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromTenantID 通过tenant_id获取内容
func (obj *_AllFuncHistoryMgr) GetFromTenantID(tenantID int64) (results []*AllFuncHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllFuncHistory{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllFuncHistoryMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllFuncHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllFuncHistory{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromName 通过name获取内容
func (obj *_AllFuncHistoryMgr) GetFromName(name string) (results []*AllFuncHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllFuncHistory{}).Where("`name` = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量查找
func (obj *_AllFuncHistoryMgr) GetBatchFromName(names []string) (results []*AllFuncHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllFuncHistory{}).Where("`name` IN (?)", names).Find(&results).Error

	return
}

// GetFromSchemaVersion 通过schema_version获取内容
func (obj *_AllFuncHistoryMgr) GetFromSchemaVersion(schemaVersion int64) (results []*AllFuncHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllFuncHistory{}).Where("`schema_version` = ?", schemaVersion).Find(&results).Error

	return
}

// GetBatchFromSchemaVersion 批量查找
func (obj *_AllFuncHistoryMgr) GetBatchFromSchemaVersion(schemaVersions []int64) (results []*AllFuncHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllFuncHistory{}).Where("`schema_version` IN (?)", schemaVersions).Find(&results).Error

	return
}

// GetFromIsDeleted 通过is_deleted获取内容
func (obj *_AllFuncHistoryMgr) GetFromIsDeleted(isDeleted int64) (results []*AllFuncHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllFuncHistory{}).Where("`is_deleted` = ?", isDeleted).Find(&results).Error

	return
}

// GetBatchFromIsDeleted 批量查找
func (obj *_AllFuncHistoryMgr) GetBatchFromIsDeleted(isDeleteds []int64) (results []*AllFuncHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllFuncHistory{}).Where("`is_deleted` IN (?)", isDeleteds).Find(&results).Error

	return
}

// GetFromRet 通过ret获取内容
func (obj *_AllFuncHistoryMgr) GetFromRet(ret int64) (results []*AllFuncHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllFuncHistory{}).Where("`ret` = ?", ret).Find(&results).Error

	return
}

// GetBatchFromRet 批量查找
func (obj *_AllFuncHistoryMgr) GetBatchFromRet(rets []int64) (results []*AllFuncHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllFuncHistory{}).Where("`ret` IN (?)", rets).Find(&results).Error

	return
}

// GetFromDl 通过dl获取内容
func (obj *_AllFuncHistoryMgr) GetFromDl(dl string) (results []*AllFuncHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllFuncHistory{}).Where("`dl` = ?", dl).Find(&results).Error

	return
}

// GetBatchFromDl 批量查找
func (obj *_AllFuncHistoryMgr) GetBatchFromDl(dls []string) (results []*AllFuncHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllFuncHistory{}).Where("`dl` IN (?)", dls).Find(&results).Error

	return
}

// GetFromUdfID 通过udf_id获取内容
func (obj *_AllFuncHistoryMgr) GetFromUdfID(udfID int64) (results []*AllFuncHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllFuncHistory{}).Where("`udf_id` = ?", udfID).Find(&results).Error

	return
}

// GetBatchFromUdfID 批量查找
func (obj *_AllFuncHistoryMgr) GetBatchFromUdfID(udfIDs []int64) (results []*AllFuncHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllFuncHistory{}).Where("`udf_id` IN (?)", udfIDs).Find(&results).Error

	return
}

// GetFromType 通过type获取内容
func (obj *_AllFuncHistoryMgr) GetFromType(_type int64) (results []*AllFuncHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllFuncHistory{}).Where("`type` = ?", _type).Find(&results).Error

	return
}

// GetBatchFromType 批量查找
func (obj *_AllFuncHistoryMgr) GetBatchFromType(_types []int64) (results []*AllFuncHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllFuncHistory{}).Where("`type` IN (?)", _types).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllFuncHistoryMgr) FetchByPrimaryKey(tenantID int64, name string, schemaVersion int64) (result AllFuncHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllFuncHistory{}).Where("`tenant_id` = ? AND `name` = ? AND `schema_version` = ?", tenantID, name, schemaVersion).First(&result).Error

	return
}
