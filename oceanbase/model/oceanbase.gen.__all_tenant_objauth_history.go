package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _AllTenantObjauthHistoryMgr struct {
	*_BaseMgr
}

// AllTenantObjauthHistoryMgr open func
func AllTenantObjauthHistoryMgr(db *gorm.DB) *_AllTenantObjauthHistoryMgr {
	if db == nil {
		panic(fmt.Errorf("AllTenantObjauthHistoryMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllTenantObjauthHistoryMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_tenant_objauth_history"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllTenantObjauthHistoryMgr) GetTableName() string {
	return "__all_tenant_objauth_history"
}

// Reset 重置gorm会话
func (obj *_AllTenantObjauthHistoryMgr) Reset() *_AllTenantObjauthHistoryMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllTenantObjauthHistoryMgr) Get() (result AllTenantObjauthHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantObjauthHistory{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllTenantObjauthHistoryMgr) Gets() (results []*AllTenantObjauthHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantObjauthHistory{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllTenantObjauthHistoryMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllTenantObjauthHistory{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithGmtCreate gmt_create获取
func (obj *_AllTenantObjauthHistoryMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllTenantObjauthHistoryMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithTenantID tenant_id获取
func (obj *_AllTenantObjauthHistoryMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithObjID obj_id获取
func (obj *_AllTenantObjauthHistoryMgr) WithObjID(objID int64) Option {
	return optionFunc(func(o *options) { o.query["obj_id"] = objID })
}

// WithObjtype objtype获取
func (obj *_AllTenantObjauthHistoryMgr) WithObjtype(objtype int64) Option {
	return optionFunc(func(o *options) { o.query["objtype"] = objtype })
}

// WithColID col_id获取
func (obj *_AllTenantObjauthHistoryMgr) WithColID(colID int64) Option {
	return optionFunc(func(o *options) { o.query["col_id"] = colID })
}

// WithGrantorID grantor_id获取
func (obj *_AllTenantObjauthHistoryMgr) WithGrantorID(grantorID int64) Option {
	return optionFunc(func(o *options) { o.query["grantor_id"] = grantorID })
}

// WithGranteeID grantee_id获取
func (obj *_AllTenantObjauthHistoryMgr) WithGranteeID(granteeID int64) Option {
	return optionFunc(func(o *options) { o.query["grantee_id"] = granteeID })
}

// WithPrivID priv_id获取
func (obj *_AllTenantObjauthHistoryMgr) WithPrivID(privID int64) Option {
	return optionFunc(func(o *options) { o.query["priv_id"] = privID })
}

// WithSchemaVersion schema_version获取
func (obj *_AllTenantObjauthHistoryMgr) WithSchemaVersion(schemaVersion int64) Option {
	return optionFunc(func(o *options) { o.query["schema_version"] = schemaVersion })
}

// WithIsDeleted is_deleted获取
func (obj *_AllTenantObjauthHistoryMgr) WithIsDeleted(isDeleted int64) Option {
	return optionFunc(func(o *options) { o.query["is_deleted"] = isDeleted })
}

// WithPrivOption priv_option获取
func (obj *_AllTenantObjauthHistoryMgr) WithPrivOption(privOption int64) Option {
	return optionFunc(func(o *options) { o.query["priv_option"] = privOption })
}

// GetByOption 功能选项模式获取
func (obj *_AllTenantObjauthHistoryMgr) GetByOption(opts ...Option) (result AllTenantObjauthHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllTenantObjauthHistory{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllTenantObjauthHistoryMgr) GetByOptions(opts ...Option) (results []*AllTenantObjauthHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllTenantObjauthHistory{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllTenantObjauthHistoryMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllTenantObjauthHistory, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllTenantObjauthHistory{}).Where(options.query)
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
func (obj *_AllTenantObjauthHistoryMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllTenantObjauthHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantObjauthHistory{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllTenantObjauthHistoryMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllTenantObjauthHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantObjauthHistory{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllTenantObjauthHistoryMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllTenantObjauthHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantObjauthHistory{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllTenantObjauthHistoryMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllTenantObjauthHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantObjauthHistory{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromTenantID 通过tenant_id获取内容
func (obj *_AllTenantObjauthHistoryMgr) GetFromTenantID(tenantID int64) (results []*AllTenantObjauthHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantObjauthHistory{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllTenantObjauthHistoryMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllTenantObjauthHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantObjauthHistory{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromObjID 通过obj_id获取内容
func (obj *_AllTenantObjauthHistoryMgr) GetFromObjID(objID int64) (results []*AllTenantObjauthHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantObjauthHistory{}).Where("`obj_id` = ?", objID).Find(&results).Error

	return
}

// GetBatchFromObjID 批量查找
func (obj *_AllTenantObjauthHistoryMgr) GetBatchFromObjID(objIDs []int64) (results []*AllTenantObjauthHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantObjauthHistory{}).Where("`obj_id` IN (?)", objIDs).Find(&results).Error

	return
}

// GetFromObjtype 通过objtype获取内容
func (obj *_AllTenantObjauthHistoryMgr) GetFromObjtype(objtype int64) (results []*AllTenantObjauthHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantObjauthHistory{}).Where("`objtype` = ?", objtype).Find(&results).Error

	return
}

// GetBatchFromObjtype 批量查找
func (obj *_AllTenantObjauthHistoryMgr) GetBatchFromObjtype(objtypes []int64) (results []*AllTenantObjauthHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantObjauthHistory{}).Where("`objtype` IN (?)", objtypes).Find(&results).Error

	return
}

// GetFromColID 通过col_id获取内容
func (obj *_AllTenantObjauthHistoryMgr) GetFromColID(colID int64) (results []*AllTenantObjauthHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantObjauthHistory{}).Where("`col_id` = ?", colID).Find(&results).Error

	return
}

// GetBatchFromColID 批量查找
func (obj *_AllTenantObjauthHistoryMgr) GetBatchFromColID(colIDs []int64) (results []*AllTenantObjauthHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantObjauthHistory{}).Where("`col_id` IN (?)", colIDs).Find(&results).Error

	return
}

// GetFromGrantorID 通过grantor_id获取内容
func (obj *_AllTenantObjauthHistoryMgr) GetFromGrantorID(grantorID int64) (results []*AllTenantObjauthHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantObjauthHistory{}).Where("`grantor_id` = ?", grantorID).Find(&results).Error

	return
}

// GetBatchFromGrantorID 批量查找
func (obj *_AllTenantObjauthHistoryMgr) GetBatchFromGrantorID(grantorIDs []int64) (results []*AllTenantObjauthHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantObjauthHistory{}).Where("`grantor_id` IN (?)", grantorIDs).Find(&results).Error

	return
}

// GetFromGranteeID 通过grantee_id获取内容
func (obj *_AllTenantObjauthHistoryMgr) GetFromGranteeID(granteeID int64) (results []*AllTenantObjauthHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantObjauthHistory{}).Where("`grantee_id` = ?", granteeID).Find(&results).Error

	return
}

// GetBatchFromGranteeID 批量查找
func (obj *_AllTenantObjauthHistoryMgr) GetBatchFromGranteeID(granteeIDs []int64) (results []*AllTenantObjauthHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantObjauthHistory{}).Where("`grantee_id` IN (?)", granteeIDs).Find(&results).Error

	return
}

// GetFromPrivID 通过priv_id获取内容
func (obj *_AllTenantObjauthHistoryMgr) GetFromPrivID(privID int64) (results []*AllTenantObjauthHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantObjauthHistory{}).Where("`priv_id` = ?", privID).Find(&results).Error

	return
}

// GetBatchFromPrivID 批量查找
func (obj *_AllTenantObjauthHistoryMgr) GetBatchFromPrivID(privIDs []int64) (results []*AllTenantObjauthHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantObjauthHistory{}).Where("`priv_id` IN (?)", privIDs).Find(&results).Error

	return
}

// GetFromSchemaVersion 通过schema_version获取内容
func (obj *_AllTenantObjauthHistoryMgr) GetFromSchemaVersion(schemaVersion int64) (results []*AllTenantObjauthHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantObjauthHistory{}).Where("`schema_version` = ?", schemaVersion).Find(&results).Error

	return
}

// GetBatchFromSchemaVersion 批量查找
func (obj *_AllTenantObjauthHistoryMgr) GetBatchFromSchemaVersion(schemaVersions []int64) (results []*AllTenantObjauthHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantObjauthHistory{}).Where("`schema_version` IN (?)", schemaVersions).Find(&results).Error

	return
}

// GetFromIsDeleted 通过is_deleted获取内容
func (obj *_AllTenantObjauthHistoryMgr) GetFromIsDeleted(isDeleted int64) (results []*AllTenantObjauthHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantObjauthHistory{}).Where("`is_deleted` = ?", isDeleted).Find(&results).Error

	return
}

// GetBatchFromIsDeleted 批量查找
func (obj *_AllTenantObjauthHistoryMgr) GetBatchFromIsDeleted(isDeleteds []int64) (results []*AllTenantObjauthHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantObjauthHistory{}).Where("`is_deleted` IN (?)", isDeleteds).Find(&results).Error

	return
}

// GetFromPrivOption 通过priv_option获取内容
func (obj *_AllTenantObjauthHistoryMgr) GetFromPrivOption(privOption int64) (results []*AllTenantObjauthHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantObjauthHistory{}).Where("`priv_option` = ?", privOption).Find(&results).Error

	return
}

// GetBatchFromPrivOption 批量查找
func (obj *_AllTenantObjauthHistoryMgr) GetBatchFromPrivOption(privOptions []int64) (results []*AllTenantObjauthHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantObjauthHistory{}).Where("`priv_option` IN (?)", privOptions).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllTenantObjauthHistoryMgr) FetchByPrimaryKey(tenantID int64, objID int64, objtype int64, colID int64, grantorID int64, granteeID int64, privID int64, schemaVersion int64) (result AllTenantObjauthHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantObjauthHistory{}).Where("`tenant_id` = ? AND `obj_id` = ? AND `objtype` = ? AND `col_id` = ? AND `grantor_id` = ? AND `grantee_id` = ? AND `priv_id` = ? AND `schema_version` = ?", tenantID, objID, objtype, colID, grantorID, granteeID, privID, schemaVersion).First(&result).Error

	return
}
