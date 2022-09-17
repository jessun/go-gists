package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _AllTenantObjauthMgr struct {
	*_BaseMgr
}

// AllTenantObjauthMgr open func
func AllTenantObjauthMgr(db *gorm.DB) *_AllTenantObjauthMgr {
	if db == nil {
		panic(fmt.Errorf("AllTenantObjauthMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllTenantObjauthMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_tenant_objauth"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllTenantObjauthMgr) GetTableName() string {
	return "__all_tenant_objauth"
}

// Reset 重置gorm会话
func (obj *_AllTenantObjauthMgr) Reset() *_AllTenantObjauthMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllTenantObjauthMgr) Get() (result AllTenantObjauth, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantObjauth{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllTenantObjauthMgr) Gets() (results []*AllTenantObjauth, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantObjauth{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllTenantObjauthMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllTenantObjauth{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithGmtCreate gmt_create获取
func (obj *_AllTenantObjauthMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllTenantObjauthMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithTenantID tenant_id获取
func (obj *_AllTenantObjauthMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithObjID obj_id获取
func (obj *_AllTenantObjauthMgr) WithObjID(objID int64) Option {
	return optionFunc(func(o *options) { o.query["obj_id"] = objID })
}

// WithObjtype objtype获取
func (obj *_AllTenantObjauthMgr) WithObjtype(objtype int64) Option {
	return optionFunc(func(o *options) { o.query["objtype"] = objtype })
}

// WithColID col_id获取
func (obj *_AllTenantObjauthMgr) WithColID(colID int64) Option {
	return optionFunc(func(o *options) { o.query["col_id"] = colID })
}

// WithGrantorID grantor_id获取
func (obj *_AllTenantObjauthMgr) WithGrantorID(grantorID int64) Option {
	return optionFunc(func(o *options) { o.query["grantor_id"] = grantorID })
}

// WithGranteeID grantee_id获取
func (obj *_AllTenantObjauthMgr) WithGranteeID(granteeID int64) Option {
	return optionFunc(func(o *options) { o.query["grantee_id"] = granteeID })
}

// WithPrivID priv_id获取
func (obj *_AllTenantObjauthMgr) WithPrivID(privID int64) Option {
	return optionFunc(func(o *options) { o.query["priv_id"] = privID })
}

// WithPrivOption priv_option获取
func (obj *_AllTenantObjauthMgr) WithPrivOption(privOption int64) Option {
	return optionFunc(func(o *options) { o.query["priv_option"] = privOption })
}

// GetByOption 功能选项模式获取
func (obj *_AllTenantObjauthMgr) GetByOption(opts ...Option) (result AllTenantObjauth, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllTenantObjauth{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllTenantObjauthMgr) GetByOptions(opts ...Option) (results []*AllTenantObjauth, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllTenantObjauth{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllTenantObjauthMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllTenantObjauth, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllTenantObjauth{}).Where(options.query)
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
func (obj *_AllTenantObjauthMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllTenantObjauth, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantObjauth{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllTenantObjauthMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllTenantObjauth, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantObjauth{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllTenantObjauthMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllTenantObjauth, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantObjauth{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllTenantObjauthMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllTenantObjauth, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantObjauth{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromTenantID 通过tenant_id获取内容
func (obj *_AllTenantObjauthMgr) GetFromTenantID(tenantID int64) (results []*AllTenantObjauth, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantObjauth{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllTenantObjauthMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllTenantObjauth, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantObjauth{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromObjID 通过obj_id获取内容
func (obj *_AllTenantObjauthMgr) GetFromObjID(objID int64) (results []*AllTenantObjauth, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantObjauth{}).Where("`obj_id` = ?", objID).Find(&results).Error

	return
}

// GetBatchFromObjID 批量查找
func (obj *_AllTenantObjauthMgr) GetBatchFromObjID(objIDs []int64) (results []*AllTenantObjauth, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantObjauth{}).Where("`obj_id` IN (?)", objIDs).Find(&results).Error

	return
}

// GetFromObjtype 通过objtype获取内容
func (obj *_AllTenantObjauthMgr) GetFromObjtype(objtype int64) (results []*AllTenantObjauth, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantObjauth{}).Where("`objtype` = ?", objtype).Find(&results).Error

	return
}

// GetBatchFromObjtype 批量查找
func (obj *_AllTenantObjauthMgr) GetBatchFromObjtype(objtypes []int64) (results []*AllTenantObjauth, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantObjauth{}).Where("`objtype` IN (?)", objtypes).Find(&results).Error

	return
}

// GetFromColID 通过col_id获取内容
func (obj *_AllTenantObjauthMgr) GetFromColID(colID int64) (results []*AllTenantObjauth, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantObjauth{}).Where("`col_id` = ?", colID).Find(&results).Error

	return
}

// GetBatchFromColID 批量查找
func (obj *_AllTenantObjauthMgr) GetBatchFromColID(colIDs []int64) (results []*AllTenantObjauth, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantObjauth{}).Where("`col_id` IN (?)", colIDs).Find(&results).Error

	return
}

// GetFromGrantorID 通过grantor_id获取内容
func (obj *_AllTenantObjauthMgr) GetFromGrantorID(grantorID int64) (results []*AllTenantObjauth, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantObjauth{}).Where("`grantor_id` = ?", grantorID).Find(&results).Error

	return
}

// GetBatchFromGrantorID 批量查找
func (obj *_AllTenantObjauthMgr) GetBatchFromGrantorID(grantorIDs []int64) (results []*AllTenantObjauth, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantObjauth{}).Where("`grantor_id` IN (?)", grantorIDs).Find(&results).Error

	return
}

// GetFromGranteeID 通过grantee_id获取内容
func (obj *_AllTenantObjauthMgr) GetFromGranteeID(granteeID int64) (results []*AllTenantObjauth, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantObjauth{}).Where("`grantee_id` = ?", granteeID).Find(&results).Error

	return
}

// GetBatchFromGranteeID 批量查找
func (obj *_AllTenantObjauthMgr) GetBatchFromGranteeID(granteeIDs []int64) (results []*AllTenantObjauth, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantObjauth{}).Where("`grantee_id` IN (?)", granteeIDs).Find(&results).Error

	return
}

// GetFromPrivID 通过priv_id获取内容
func (obj *_AllTenantObjauthMgr) GetFromPrivID(privID int64) (results []*AllTenantObjauth, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantObjauth{}).Where("`priv_id` = ?", privID).Find(&results).Error

	return
}

// GetBatchFromPrivID 批量查找
func (obj *_AllTenantObjauthMgr) GetBatchFromPrivID(privIDs []int64) (results []*AllTenantObjauth, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantObjauth{}).Where("`priv_id` IN (?)", privIDs).Find(&results).Error

	return
}

// GetFromPrivOption 通过priv_option获取内容
func (obj *_AllTenantObjauthMgr) GetFromPrivOption(privOption int64) (results []*AllTenantObjauth, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantObjauth{}).Where("`priv_option` = ?", privOption).Find(&results).Error

	return
}

// GetBatchFromPrivOption 批量查找
func (obj *_AllTenantObjauthMgr) GetBatchFromPrivOption(privOptions []int64) (results []*AllTenantObjauth, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantObjauth{}).Where("`priv_option` IN (?)", privOptions).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllTenantObjauthMgr) FetchByPrimaryKey(tenantID int64, objID int64, objtype int64, colID int64, grantorID int64, granteeID int64, privID int64) (result AllTenantObjauth, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantObjauth{}).Where("`tenant_id` = ? AND `obj_id` = ? AND `objtype` = ? AND `col_id` = ? AND `grantor_id` = ? AND `grantee_id` = ? AND `priv_id` = ?", tenantID, objID, objtype, colID, grantorID, granteeID, privID).First(&result).Error

	return
}
