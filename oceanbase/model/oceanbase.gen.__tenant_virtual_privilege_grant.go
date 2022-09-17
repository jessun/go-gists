package model

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _TenantVirtualPrivilegeGrantMgr struct {
	*_BaseMgr
}

// TenantVirtualPrivilegeGrantMgr open func
func TenantVirtualPrivilegeGrantMgr(db *gorm.DB) *_TenantVirtualPrivilegeGrantMgr {
	if db == nil {
		panic(fmt.Errorf("TenantVirtualPrivilegeGrantMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_TenantVirtualPrivilegeGrantMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__tenant_virtual_privilege_grant"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_TenantVirtualPrivilegeGrantMgr) GetTableName() string {
	return "__tenant_virtual_privilege_grant"
}

// Reset 重置gorm会话
func (obj *_TenantVirtualPrivilegeGrantMgr) Reset() *_TenantVirtualPrivilegeGrantMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_TenantVirtualPrivilegeGrantMgr) Get() (result TenantVirtualPrivilegeGrant, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualPrivilegeGrant{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_TenantVirtualPrivilegeGrantMgr) Gets() (results []*TenantVirtualPrivilegeGrant, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualPrivilegeGrant{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_TenantVirtualPrivilegeGrantMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(TenantVirtualPrivilegeGrant{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithUserID user_id获取
func (obj *_TenantVirtualPrivilegeGrantMgr) WithUserID(userID int64) Option {
	return optionFunc(func(o *options) { o.query["user_id"] = userID })
}

// WithGrants grants获取
func (obj *_TenantVirtualPrivilegeGrantMgr) WithGrants(grants string) Option {
	return optionFunc(func(o *options) { o.query["grants"] = grants })
}

// GetByOption 功能选项模式获取
func (obj *_TenantVirtualPrivilegeGrantMgr) GetByOption(opts ...Option) (result TenantVirtualPrivilegeGrant, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualPrivilegeGrant{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_TenantVirtualPrivilegeGrantMgr) GetByOptions(opts ...Option) (results []*TenantVirtualPrivilegeGrant, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualPrivilegeGrant{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_TenantVirtualPrivilegeGrantMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]TenantVirtualPrivilegeGrant, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(TenantVirtualPrivilegeGrant{}).Where(options.query)
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

// GetFromUserID 通过user_id获取内容
func (obj *_TenantVirtualPrivilegeGrantMgr) GetFromUserID(userID int64) (result TenantVirtualPrivilegeGrant, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualPrivilegeGrant{}).Where("`user_id` = ?", userID).First(&result).Error

	return
}

// GetBatchFromUserID 批量查找
func (obj *_TenantVirtualPrivilegeGrantMgr) GetBatchFromUserID(userIDs []int64) (results []*TenantVirtualPrivilegeGrant, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualPrivilegeGrant{}).Where("`user_id` IN (?)", userIDs).Find(&results).Error

	return
}

// GetFromGrants 通过grants获取内容
func (obj *_TenantVirtualPrivilegeGrantMgr) GetFromGrants(grants string) (results []*TenantVirtualPrivilegeGrant, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualPrivilegeGrant{}).Where("`grants` = ?", grants).Find(&results).Error

	return
}

// GetBatchFromGrants 批量查找
func (obj *_TenantVirtualPrivilegeGrantMgr) GetBatchFromGrants(grantss []string) (results []*TenantVirtualPrivilegeGrant, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualPrivilegeGrant{}).Where("`grants` IN (?)", grantss).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_TenantVirtualPrivilegeGrantMgr) FetchByPrimaryKey(userID int64) (result TenantVirtualPrivilegeGrant, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(TenantVirtualPrivilegeGrant{}).Where("`user_id` = ?", userID).First(&result).Error

	return
}
