package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _AllVirtualObjauthHistoryMgr struct {
	*_BaseMgr
}

// AllVirtualObjauthHistoryMgr open func
func AllVirtualObjauthHistoryMgr(db *gorm.DB) *_AllVirtualObjauthHistoryMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualObjauthHistoryMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualObjauthHistoryMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_objauth_history"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualObjauthHistoryMgr) GetTableName() string {
	return "__all_virtual_objauth_history"
}

// Reset 重置gorm会话
func (obj *_AllVirtualObjauthHistoryMgr) Reset() *_AllVirtualObjauthHistoryMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualObjauthHistoryMgr) Get() (result AllVirtualObjauthHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualObjauthHistory{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualObjauthHistoryMgr) Gets() (results []*AllVirtualObjauthHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualObjauthHistory{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualObjauthHistoryMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualObjauthHistory{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTenantID tenant_id获取
func (obj *_AllVirtualObjauthHistoryMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithObjID obj_id获取
func (obj *_AllVirtualObjauthHistoryMgr) WithObjID(objID int64) Option {
	return optionFunc(func(o *options) { o.query["obj_id"] = objID })
}

// WithObjtype objtype获取
func (obj *_AllVirtualObjauthHistoryMgr) WithObjtype(objtype int64) Option {
	return optionFunc(func(o *options) { o.query["objtype"] = objtype })
}

// WithColID col_id获取
func (obj *_AllVirtualObjauthHistoryMgr) WithColID(colID int64) Option {
	return optionFunc(func(o *options) { o.query["col_id"] = colID })
}

// WithGrantorID grantor_id获取
func (obj *_AllVirtualObjauthHistoryMgr) WithGrantorID(grantorID int64) Option {
	return optionFunc(func(o *options) { o.query["grantor_id"] = grantorID })
}

// WithGranteeID grantee_id获取
func (obj *_AllVirtualObjauthHistoryMgr) WithGranteeID(granteeID int64) Option {
	return optionFunc(func(o *options) { o.query["grantee_id"] = granteeID })
}

// WithPrivID priv_id获取
func (obj *_AllVirtualObjauthHistoryMgr) WithPrivID(privID int64) Option {
	return optionFunc(func(o *options) { o.query["priv_id"] = privID })
}

// WithSchemaVersion schema_version获取
func (obj *_AllVirtualObjauthHistoryMgr) WithSchemaVersion(schemaVersion int64) Option {
	return optionFunc(func(o *options) { o.query["schema_version"] = schemaVersion })
}

// WithGmtCreate gmt_create获取
func (obj *_AllVirtualObjauthHistoryMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllVirtualObjauthHistoryMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithIsDeleted is_deleted获取
func (obj *_AllVirtualObjauthHistoryMgr) WithIsDeleted(isDeleted int64) Option {
	return optionFunc(func(o *options) { o.query["is_deleted"] = isDeleted })
}

// WithPrivOption priv_option获取
func (obj *_AllVirtualObjauthHistoryMgr) WithPrivOption(privOption int64) Option {
	return optionFunc(func(o *options) { o.query["priv_option"] = privOption })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualObjauthHistoryMgr) GetByOption(opts ...Option) (result AllVirtualObjauthHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualObjauthHistory{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualObjauthHistoryMgr) GetByOptions(opts ...Option) (results []*AllVirtualObjauthHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualObjauthHistory{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualObjauthHistoryMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualObjauthHistory, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualObjauthHistory{}).Where(options.query)
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
func (obj *_AllVirtualObjauthHistoryMgr) GetFromTenantID(tenantID int64) (results []*AllVirtualObjauthHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualObjauthHistory{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllVirtualObjauthHistoryMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllVirtualObjauthHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualObjauthHistory{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromObjID 通过obj_id获取内容
func (obj *_AllVirtualObjauthHistoryMgr) GetFromObjID(objID int64) (results []*AllVirtualObjauthHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualObjauthHistory{}).Where("`obj_id` = ?", objID).Find(&results).Error

	return
}

// GetBatchFromObjID 批量查找
func (obj *_AllVirtualObjauthHistoryMgr) GetBatchFromObjID(objIDs []int64) (results []*AllVirtualObjauthHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualObjauthHistory{}).Where("`obj_id` IN (?)", objIDs).Find(&results).Error

	return
}

// GetFromObjtype 通过objtype获取内容
func (obj *_AllVirtualObjauthHistoryMgr) GetFromObjtype(objtype int64) (results []*AllVirtualObjauthHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualObjauthHistory{}).Where("`objtype` = ?", objtype).Find(&results).Error

	return
}

// GetBatchFromObjtype 批量查找
func (obj *_AllVirtualObjauthHistoryMgr) GetBatchFromObjtype(objtypes []int64) (results []*AllVirtualObjauthHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualObjauthHistory{}).Where("`objtype` IN (?)", objtypes).Find(&results).Error

	return
}

// GetFromColID 通过col_id获取内容
func (obj *_AllVirtualObjauthHistoryMgr) GetFromColID(colID int64) (results []*AllVirtualObjauthHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualObjauthHistory{}).Where("`col_id` = ?", colID).Find(&results).Error

	return
}

// GetBatchFromColID 批量查找
func (obj *_AllVirtualObjauthHistoryMgr) GetBatchFromColID(colIDs []int64) (results []*AllVirtualObjauthHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualObjauthHistory{}).Where("`col_id` IN (?)", colIDs).Find(&results).Error

	return
}

// GetFromGrantorID 通过grantor_id获取内容
func (obj *_AllVirtualObjauthHistoryMgr) GetFromGrantorID(grantorID int64) (results []*AllVirtualObjauthHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualObjauthHistory{}).Where("`grantor_id` = ?", grantorID).Find(&results).Error

	return
}

// GetBatchFromGrantorID 批量查找
func (obj *_AllVirtualObjauthHistoryMgr) GetBatchFromGrantorID(grantorIDs []int64) (results []*AllVirtualObjauthHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualObjauthHistory{}).Where("`grantor_id` IN (?)", grantorIDs).Find(&results).Error

	return
}

// GetFromGranteeID 通过grantee_id获取内容
func (obj *_AllVirtualObjauthHistoryMgr) GetFromGranteeID(granteeID int64) (results []*AllVirtualObjauthHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualObjauthHistory{}).Where("`grantee_id` = ?", granteeID).Find(&results).Error

	return
}

// GetBatchFromGranteeID 批量查找
func (obj *_AllVirtualObjauthHistoryMgr) GetBatchFromGranteeID(granteeIDs []int64) (results []*AllVirtualObjauthHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualObjauthHistory{}).Where("`grantee_id` IN (?)", granteeIDs).Find(&results).Error

	return
}

// GetFromPrivID 通过priv_id获取内容
func (obj *_AllVirtualObjauthHistoryMgr) GetFromPrivID(privID int64) (results []*AllVirtualObjauthHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualObjauthHistory{}).Where("`priv_id` = ?", privID).Find(&results).Error

	return
}

// GetBatchFromPrivID 批量查找
func (obj *_AllVirtualObjauthHistoryMgr) GetBatchFromPrivID(privIDs []int64) (results []*AllVirtualObjauthHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualObjauthHistory{}).Where("`priv_id` IN (?)", privIDs).Find(&results).Error

	return
}

// GetFromSchemaVersion 通过schema_version获取内容
func (obj *_AllVirtualObjauthHistoryMgr) GetFromSchemaVersion(schemaVersion int64) (results []*AllVirtualObjauthHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualObjauthHistory{}).Where("`schema_version` = ?", schemaVersion).Find(&results).Error

	return
}

// GetBatchFromSchemaVersion 批量查找
func (obj *_AllVirtualObjauthHistoryMgr) GetBatchFromSchemaVersion(schemaVersions []int64) (results []*AllVirtualObjauthHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualObjauthHistory{}).Where("`schema_version` IN (?)", schemaVersions).Find(&results).Error

	return
}

// GetFromGmtCreate 通过gmt_create获取内容
func (obj *_AllVirtualObjauthHistoryMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllVirtualObjauthHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualObjauthHistory{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllVirtualObjauthHistoryMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllVirtualObjauthHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualObjauthHistory{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllVirtualObjauthHistoryMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllVirtualObjauthHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualObjauthHistory{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllVirtualObjauthHistoryMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllVirtualObjauthHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualObjauthHistory{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromIsDeleted 通过is_deleted获取内容
func (obj *_AllVirtualObjauthHistoryMgr) GetFromIsDeleted(isDeleted int64) (results []*AllVirtualObjauthHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualObjauthHistory{}).Where("`is_deleted` = ?", isDeleted).Find(&results).Error

	return
}

// GetBatchFromIsDeleted 批量查找
func (obj *_AllVirtualObjauthHistoryMgr) GetBatchFromIsDeleted(isDeleteds []int64) (results []*AllVirtualObjauthHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualObjauthHistory{}).Where("`is_deleted` IN (?)", isDeleteds).Find(&results).Error

	return
}

// GetFromPrivOption 通过priv_option获取内容
func (obj *_AllVirtualObjauthHistoryMgr) GetFromPrivOption(privOption int64) (results []*AllVirtualObjauthHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualObjauthHistory{}).Where("`priv_option` = ?", privOption).Find(&results).Error

	return
}

// GetBatchFromPrivOption 批量查找
func (obj *_AllVirtualObjauthHistoryMgr) GetBatchFromPrivOption(privOptions []int64) (results []*AllVirtualObjauthHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualObjauthHistory{}).Where("`priv_option` IN (?)", privOptions).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllVirtualObjauthHistoryMgr) FetchByPrimaryKey(tenantID int64, objID int64, objtype int64, colID int64, grantorID int64, granteeID int64, privID int64, schemaVersion int64) (result AllVirtualObjauthHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualObjauthHistory{}).Where("`tenant_id` = ? AND `obj_id` = ? AND `objtype` = ? AND `col_id` = ? AND `grantor_id` = ? AND `grantee_id` = ? AND `priv_id` = ? AND `schema_version` = ?", tenantID, objID, objtype, colID, grantorID, granteeID, privID, schemaVersion).First(&result).Error

	return
}
