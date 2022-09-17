package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _AllVirtualObjauthMgr struct {
	*_BaseMgr
}

// AllVirtualObjauthMgr open func
func AllVirtualObjauthMgr(db *gorm.DB) *_AllVirtualObjauthMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualObjauthMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualObjauthMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_objauth"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualObjauthMgr) GetTableName() string {
	return "__all_virtual_objauth"
}

// Reset 重置gorm会话
func (obj *_AllVirtualObjauthMgr) Reset() *_AllVirtualObjauthMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualObjauthMgr) Get() (result AllVirtualObjauth, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualObjauth{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualObjauthMgr) Gets() (results []*AllVirtualObjauth, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualObjauth{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualObjauthMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualObjauth{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTenantID tenant_id获取
func (obj *_AllVirtualObjauthMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithObjID obj_id获取
func (obj *_AllVirtualObjauthMgr) WithObjID(objID int64) Option {
	return optionFunc(func(o *options) { o.query["obj_id"] = objID })
}

// WithObjtype objtype获取
func (obj *_AllVirtualObjauthMgr) WithObjtype(objtype int64) Option {
	return optionFunc(func(o *options) { o.query["objtype"] = objtype })
}

// WithColID col_id获取
func (obj *_AllVirtualObjauthMgr) WithColID(colID int64) Option {
	return optionFunc(func(o *options) { o.query["col_id"] = colID })
}

// WithGrantorID grantor_id获取
func (obj *_AllVirtualObjauthMgr) WithGrantorID(grantorID int64) Option {
	return optionFunc(func(o *options) { o.query["grantor_id"] = grantorID })
}

// WithGranteeID grantee_id获取
func (obj *_AllVirtualObjauthMgr) WithGranteeID(granteeID int64) Option {
	return optionFunc(func(o *options) { o.query["grantee_id"] = granteeID })
}

// WithPrivID priv_id获取
func (obj *_AllVirtualObjauthMgr) WithPrivID(privID int64) Option {
	return optionFunc(func(o *options) { o.query["priv_id"] = privID })
}

// WithGmtCreate gmt_create获取
func (obj *_AllVirtualObjauthMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllVirtualObjauthMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithPrivOption priv_option获取
func (obj *_AllVirtualObjauthMgr) WithPrivOption(privOption int64) Option {
	return optionFunc(func(o *options) { o.query["priv_option"] = privOption })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualObjauthMgr) GetByOption(opts ...Option) (result AllVirtualObjauth, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualObjauth{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualObjauthMgr) GetByOptions(opts ...Option) (results []*AllVirtualObjauth, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualObjauth{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualObjauthMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualObjauth, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualObjauth{}).Where(options.query)
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
func (obj *_AllVirtualObjauthMgr) GetFromTenantID(tenantID int64) (results []*AllVirtualObjauth, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualObjauth{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllVirtualObjauthMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllVirtualObjauth, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualObjauth{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromObjID 通过obj_id获取内容
func (obj *_AllVirtualObjauthMgr) GetFromObjID(objID int64) (results []*AllVirtualObjauth, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualObjauth{}).Where("`obj_id` = ?", objID).Find(&results).Error

	return
}

// GetBatchFromObjID 批量查找
func (obj *_AllVirtualObjauthMgr) GetBatchFromObjID(objIDs []int64) (results []*AllVirtualObjauth, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualObjauth{}).Where("`obj_id` IN (?)", objIDs).Find(&results).Error

	return
}

// GetFromObjtype 通过objtype获取内容
func (obj *_AllVirtualObjauthMgr) GetFromObjtype(objtype int64) (results []*AllVirtualObjauth, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualObjauth{}).Where("`objtype` = ?", objtype).Find(&results).Error

	return
}

// GetBatchFromObjtype 批量查找
func (obj *_AllVirtualObjauthMgr) GetBatchFromObjtype(objtypes []int64) (results []*AllVirtualObjauth, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualObjauth{}).Where("`objtype` IN (?)", objtypes).Find(&results).Error

	return
}

// GetFromColID 通过col_id获取内容
func (obj *_AllVirtualObjauthMgr) GetFromColID(colID int64) (results []*AllVirtualObjauth, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualObjauth{}).Where("`col_id` = ?", colID).Find(&results).Error

	return
}

// GetBatchFromColID 批量查找
func (obj *_AllVirtualObjauthMgr) GetBatchFromColID(colIDs []int64) (results []*AllVirtualObjauth, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualObjauth{}).Where("`col_id` IN (?)", colIDs).Find(&results).Error

	return
}

// GetFromGrantorID 通过grantor_id获取内容
func (obj *_AllVirtualObjauthMgr) GetFromGrantorID(grantorID int64) (results []*AllVirtualObjauth, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualObjauth{}).Where("`grantor_id` = ?", grantorID).Find(&results).Error

	return
}

// GetBatchFromGrantorID 批量查找
func (obj *_AllVirtualObjauthMgr) GetBatchFromGrantorID(grantorIDs []int64) (results []*AllVirtualObjauth, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualObjauth{}).Where("`grantor_id` IN (?)", grantorIDs).Find(&results).Error

	return
}

// GetFromGranteeID 通过grantee_id获取内容
func (obj *_AllVirtualObjauthMgr) GetFromGranteeID(granteeID int64) (results []*AllVirtualObjauth, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualObjauth{}).Where("`grantee_id` = ?", granteeID).Find(&results).Error

	return
}

// GetBatchFromGranteeID 批量查找
func (obj *_AllVirtualObjauthMgr) GetBatchFromGranteeID(granteeIDs []int64) (results []*AllVirtualObjauth, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualObjauth{}).Where("`grantee_id` IN (?)", granteeIDs).Find(&results).Error

	return
}

// GetFromPrivID 通过priv_id获取内容
func (obj *_AllVirtualObjauthMgr) GetFromPrivID(privID int64) (results []*AllVirtualObjauth, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualObjauth{}).Where("`priv_id` = ?", privID).Find(&results).Error

	return
}

// GetBatchFromPrivID 批量查找
func (obj *_AllVirtualObjauthMgr) GetBatchFromPrivID(privIDs []int64) (results []*AllVirtualObjauth, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualObjauth{}).Where("`priv_id` IN (?)", privIDs).Find(&results).Error

	return
}

// GetFromGmtCreate 通过gmt_create获取内容
func (obj *_AllVirtualObjauthMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllVirtualObjauth, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualObjauth{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllVirtualObjauthMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllVirtualObjauth, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualObjauth{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllVirtualObjauthMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllVirtualObjauth, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualObjauth{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllVirtualObjauthMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllVirtualObjauth, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualObjauth{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromPrivOption 通过priv_option获取内容
func (obj *_AllVirtualObjauthMgr) GetFromPrivOption(privOption int64) (results []*AllVirtualObjauth, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualObjauth{}).Where("`priv_option` = ?", privOption).Find(&results).Error

	return
}

// GetBatchFromPrivOption 批量查找
func (obj *_AllVirtualObjauthMgr) GetBatchFromPrivOption(privOptions []int64) (results []*AllVirtualObjauth, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualObjauth{}).Where("`priv_option` IN (?)", privOptions).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllVirtualObjauthMgr) FetchByPrimaryKey(tenantID int64, objID int64, objtype int64, colID int64, grantorID int64, granteeID int64, privID int64) (result AllVirtualObjauth, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualObjauth{}).Where("`tenant_id` = ? AND `obj_id` = ? AND `objtype` = ? AND `col_id` = ? AND `grantor_id` = ? AND `grantee_id` = ? AND `priv_id` = ?", tenantID, objID, objtype, colID, grantorID, granteeID, privID).First(&result).Error

	return
}
