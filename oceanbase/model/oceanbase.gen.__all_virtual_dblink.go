package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _AllVirtualDblinkMgr struct {
	*_BaseMgr
}

// AllVirtualDblinkMgr open func
func AllVirtualDblinkMgr(db *gorm.DB) *_AllVirtualDblinkMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualDblinkMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualDblinkMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_dblink"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualDblinkMgr) GetTableName() string {
	return "__all_virtual_dblink"
}

// Reset 重置gorm会话
func (obj *_AllVirtualDblinkMgr) Reset() *_AllVirtualDblinkMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualDblinkMgr) Get() (result AllVirtualDblink, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDblink{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualDblinkMgr) Gets() (results []*AllVirtualDblink, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDblink{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualDblinkMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualDblink{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTenantID tenant_id获取
func (obj *_AllVirtualDblinkMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithDblinkID dblink_id获取
func (obj *_AllVirtualDblinkMgr) WithDblinkID(dblinkID int64) Option {
	return optionFunc(func(o *options) { o.query["dblink_id"] = dblinkID })
}

// WithGmtCreate gmt_create获取
func (obj *_AllVirtualDblinkMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllVirtualDblinkMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithDblinkName dblink_name获取
func (obj *_AllVirtualDblinkMgr) WithDblinkName(dblinkName string) Option {
	return optionFunc(func(o *options) { o.query["dblink_name"] = dblinkName })
}

// WithOwnerID owner_id获取
func (obj *_AllVirtualDblinkMgr) WithOwnerID(ownerID int64) Option {
	return optionFunc(func(o *options) { o.query["owner_id"] = ownerID })
}

// WithHostIP host_ip获取
func (obj *_AllVirtualDblinkMgr) WithHostIP(hostIP string) Option {
	return optionFunc(func(o *options) { o.query["host_ip"] = hostIP })
}

// WithHostPort host_port获取
func (obj *_AllVirtualDblinkMgr) WithHostPort(hostPort int64) Option {
	return optionFunc(func(o *options) { o.query["host_port"] = hostPort })
}

// WithClusterName cluster_name获取
func (obj *_AllVirtualDblinkMgr) WithClusterName(clusterName string) Option {
	return optionFunc(func(o *options) { o.query["cluster_name"] = clusterName })
}

// WithTenantName tenant_name获取
func (obj *_AllVirtualDblinkMgr) WithTenantName(tenantName string) Option {
	return optionFunc(func(o *options) { o.query["tenant_name"] = tenantName })
}

// WithUserName user_name获取
func (obj *_AllVirtualDblinkMgr) WithUserName(userName string) Option {
	return optionFunc(func(o *options) { o.query["user_name"] = userName })
}

// WithPassword password获取
func (obj *_AllVirtualDblinkMgr) WithPassword(password string) Option {
	return optionFunc(func(o *options) { o.query["password"] = password })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualDblinkMgr) GetByOption(opts ...Option) (result AllVirtualDblink, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDblink{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualDblinkMgr) GetByOptions(opts ...Option) (results []*AllVirtualDblink, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDblink{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualDblinkMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualDblink, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualDblink{}).Where(options.query)
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
func (obj *_AllVirtualDblinkMgr) GetFromTenantID(tenantID int64) (results []*AllVirtualDblink, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDblink{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllVirtualDblinkMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllVirtualDblink, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDblink{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromDblinkID 通过dblink_id获取内容
func (obj *_AllVirtualDblinkMgr) GetFromDblinkID(dblinkID int64) (results []*AllVirtualDblink, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDblink{}).Where("`dblink_id` = ?", dblinkID).Find(&results).Error

	return
}

// GetBatchFromDblinkID 批量查找
func (obj *_AllVirtualDblinkMgr) GetBatchFromDblinkID(dblinkIDs []int64) (results []*AllVirtualDblink, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDblink{}).Where("`dblink_id` IN (?)", dblinkIDs).Find(&results).Error

	return
}

// GetFromGmtCreate 通过gmt_create获取内容
func (obj *_AllVirtualDblinkMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllVirtualDblink, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDblink{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllVirtualDblinkMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllVirtualDblink, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDblink{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllVirtualDblinkMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllVirtualDblink, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDblink{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllVirtualDblinkMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllVirtualDblink, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDblink{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromDblinkName 通过dblink_name获取内容
func (obj *_AllVirtualDblinkMgr) GetFromDblinkName(dblinkName string) (results []*AllVirtualDblink, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDblink{}).Where("`dblink_name` = ?", dblinkName).Find(&results).Error

	return
}

// GetBatchFromDblinkName 批量查找
func (obj *_AllVirtualDblinkMgr) GetBatchFromDblinkName(dblinkNames []string) (results []*AllVirtualDblink, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDblink{}).Where("`dblink_name` IN (?)", dblinkNames).Find(&results).Error

	return
}

// GetFromOwnerID 通过owner_id获取内容
func (obj *_AllVirtualDblinkMgr) GetFromOwnerID(ownerID int64) (results []*AllVirtualDblink, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDblink{}).Where("`owner_id` = ?", ownerID).Find(&results).Error

	return
}

// GetBatchFromOwnerID 批量查找
func (obj *_AllVirtualDblinkMgr) GetBatchFromOwnerID(ownerIDs []int64) (results []*AllVirtualDblink, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDblink{}).Where("`owner_id` IN (?)", ownerIDs).Find(&results).Error

	return
}

// GetFromHostIP 通过host_ip获取内容
func (obj *_AllVirtualDblinkMgr) GetFromHostIP(hostIP string) (results []*AllVirtualDblink, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDblink{}).Where("`host_ip` = ?", hostIP).Find(&results).Error

	return
}

// GetBatchFromHostIP 批量查找
func (obj *_AllVirtualDblinkMgr) GetBatchFromHostIP(hostIPs []string) (results []*AllVirtualDblink, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDblink{}).Where("`host_ip` IN (?)", hostIPs).Find(&results).Error

	return
}

// GetFromHostPort 通过host_port获取内容
func (obj *_AllVirtualDblinkMgr) GetFromHostPort(hostPort int64) (results []*AllVirtualDblink, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDblink{}).Where("`host_port` = ?", hostPort).Find(&results).Error

	return
}

// GetBatchFromHostPort 批量查找
func (obj *_AllVirtualDblinkMgr) GetBatchFromHostPort(hostPorts []int64) (results []*AllVirtualDblink, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDblink{}).Where("`host_port` IN (?)", hostPorts).Find(&results).Error

	return
}

// GetFromClusterName 通过cluster_name获取内容
func (obj *_AllVirtualDblinkMgr) GetFromClusterName(clusterName string) (results []*AllVirtualDblink, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDblink{}).Where("`cluster_name` = ?", clusterName).Find(&results).Error

	return
}

// GetBatchFromClusterName 批量查找
func (obj *_AllVirtualDblinkMgr) GetBatchFromClusterName(clusterNames []string) (results []*AllVirtualDblink, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDblink{}).Where("`cluster_name` IN (?)", clusterNames).Find(&results).Error

	return
}

// GetFromTenantName 通过tenant_name获取内容
func (obj *_AllVirtualDblinkMgr) GetFromTenantName(tenantName string) (results []*AllVirtualDblink, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDblink{}).Where("`tenant_name` = ?", tenantName).Find(&results).Error

	return
}

// GetBatchFromTenantName 批量查找
func (obj *_AllVirtualDblinkMgr) GetBatchFromTenantName(tenantNames []string) (results []*AllVirtualDblink, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDblink{}).Where("`tenant_name` IN (?)", tenantNames).Find(&results).Error

	return
}

// GetFromUserName 通过user_name获取内容
func (obj *_AllVirtualDblinkMgr) GetFromUserName(userName string) (results []*AllVirtualDblink, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDblink{}).Where("`user_name` = ?", userName).Find(&results).Error

	return
}

// GetBatchFromUserName 批量查找
func (obj *_AllVirtualDblinkMgr) GetBatchFromUserName(userNames []string) (results []*AllVirtualDblink, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDblink{}).Where("`user_name` IN (?)", userNames).Find(&results).Error

	return
}

// GetFromPassword 通过password获取内容
func (obj *_AllVirtualDblinkMgr) GetFromPassword(password string) (results []*AllVirtualDblink, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDblink{}).Where("`password` = ?", password).Find(&results).Error

	return
}

// GetBatchFromPassword 批量查找
func (obj *_AllVirtualDblinkMgr) GetBatchFromPassword(passwords []string) (results []*AllVirtualDblink, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDblink{}).Where("`password` IN (?)", passwords).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllVirtualDblinkMgr) FetchByPrimaryKey(tenantID int64, dblinkID int64) (result AllVirtualDblink, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDblink{}).Where("`tenant_id` = ? AND `dblink_id` = ?", tenantID, dblinkID).First(&result).Error

	return
}
