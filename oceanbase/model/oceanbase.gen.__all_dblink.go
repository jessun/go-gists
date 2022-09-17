package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _AllDblinkMgr struct {
	*_BaseMgr
}

// AllDblinkMgr open func
func AllDblinkMgr(db *gorm.DB) *_AllDblinkMgr {
	if db == nil {
		panic(fmt.Errorf("AllDblinkMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllDblinkMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_dblink"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllDblinkMgr) GetTableName() string {
	return "__all_dblink"
}

// Reset 重置gorm会话
func (obj *_AllDblinkMgr) Reset() *_AllDblinkMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllDblinkMgr) Get() (result AllDblink, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDblink{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllDblinkMgr) Gets() (results []*AllDblink, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDblink{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllDblinkMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllDblink{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithGmtCreate gmt_create获取
func (obj *_AllDblinkMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllDblinkMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithTenantID tenant_id获取
func (obj *_AllDblinkMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithDblinkID dblink_id获取
func (obj *_AllDblinkMgr) WithDblinkID(dblinkID int64) Option {
	return optionFunc(func(o *options) { o.query["dblink_id"] = dblinkID })
}

// WithDblinkName dblink_name获取
func (obj *_AllDblinkMgr) WithDblinkName(dblinkName string) Option {
	return optionFunc(func(o *options) { o.query["dblink_name"] = dblinkName })
}

// WithOwnerID owner_id获取
func (obj *_AllDblinkMgr) WithOwnerID(ownerID int64) Option {
	return optionFunc(func(o *options) { o.query["owner_id"] = ownerID })
}

// WithHostIP host_ip获取
func (obj *_AllDblinkMgr) WithHostIP(hostIP string) Option {
	return optionFunc(func(o *options) { o.query["host_ip"] = hostIP })
}

// WithHostPort host_port获取
func (obj *_AllDblinkMgr) WithHostPort(hostPort int64) Option {
	return optionFunc(func(o *options) { o.query["host_port"] = hostPort })
}

// WithClusterName cluster_name获取
func (obj *_AllDblinkMgr) WithClusterName(clusterName string) Option {
	return optionFunc(func(o *options) { o.query["cluster_name"] = clusterName })
}

// WithTenantName tenant_name获取
func (obj *_AllDblinkMgr) WithTenantName(tenantName string) Option {
	return optionFunc(func(o *options) { o.query["tenant_name"] = tenantName })
}

// WithUserName user_name获取
func (obj *_AllDblinkMgr) WithUserName(userName string) Option {
	return optionFunc(func(o *options) { o.query["user_name"] = userName })
}

// WithPassword password获取
func (obj *_AllDblinkMgr) WithPassword(password string) Option {
	return optionFunc(func(o *options) { o.query["password"] = password })
}

// GetByOption 功能选项模式获取
func (obj *_AllDblinkMgr) GetByOption(opts ...Option) (result AllDblink, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllDblink{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllDblinkMgr) GetByOptions(opts ...Option) (results []*AllDblink, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllDblink{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllDblinkMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllDblink, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllDblink{}).Where(options.query)
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
func (obj *_AllDblinkMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllDblink, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDblink{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllDblinkMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllDblink, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDblink{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllDblinkMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllDblink, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDblink{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllDblinkMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllDblink, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDblink{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromTenantID 通过tenant_id获取内容
func (obj *_AllDblinkMgr) GetFromTenantID(tenantID int64) (results []*AllDblink, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDblink{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllDblinkMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllDblink, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDblink{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromDblinkID 通过dblink_id获取内容
func (obj *_AllDblinkMgr) GetFromDblinkID(dblinkID int64) (results []*AllDblink, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDblink{}).Where("`dblink_id` = ?", dblinkID).Find(&results).Error

	return
}

// GetBatchFromDblinkID 批量查找
func (obj *_AllDblinkMgr) GetBatchFromDblinkID(dblinkIDs []int64) (results []*AllDblink, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDblink{}).Where("`dblink_id` IN (?)", dblinkIDs).Find(&results).Error

	return
}

// GetFromDblinkName 通过dblink_name获取内容
func (obj *_AllDblinkMgr) GetFromDblinkName(dblinkName string) (results []*AllDblink, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDblink{}).Where("`dblink_name` = ?", dblinkName).Find(&results).Error

	return
}

// GetBatchFromDblinkName 批量查找
func (obj *_AllDblinkMgr) GetBatchFromDblinkName(dblinkNames []string) (results []*AllDblink, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDblink{}).Where("`dblink_name` IN (?)", dblinkNames).Find(&results).Error

	return
}

// GetFromOwnerID 通过owner_id获取内容
func (obj *_AllDblinkMgr) GetFromOwnerID(ownerID int64) (results []*AllDblink, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDblink{}).Where("`owner_id` = ?", ownerID).Find(&results).Error

	return
}

// GetBatchFromOwnerID 批量查找
func (obj *_AllDblinkMgr) GetBatchFromOwnerID(ownerIDs []int64) (results []*AllDblink, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDblink{}).Where("`owner_id` IN (?)", ownerIDs).Find(&results).Error

	return
}

// GetFromHostIP 通过host_ip获取内容
func (obj *_AllDblinkMgr) GetFromHostIP(hostIP string) (results []*AllDblink, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDblink{}).Where("`host_ip` = ?", hostIP).Find(&results).Error

	return
}

// GetBatchFromHostIP 批量查找
func (obj *_AllDblinkMgr) GetBatchFromHostIP(hostIPs []string) (results []*AllDblink, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDblink{}).Where("`host_ip` IN (?)", hostIPs).Find(&results).Error

	return
}

// GetFromHostPort 通过host_port获取内容
func (obj *_AllDblinkMgr) GetFromHostPort(hostPort int64) (results []*AllDblink, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDblink{}).Where("`host_port` = ?", hostPort).Find(&results).Error

	return
}

// GetBatchFromHostPort 批量查找
func (obj *_AllDblinkMgr) GetBatchFromHostPort(hostPorts []int64) (results []*AllDblink, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDblink{}).Where("`host_port` IN (?)", hostPorts).Find(&results).Error

	return
}

// GetFromClusterName 通过cluster_name获取内容
func (obj *_AllDblinkMgr) GetFromClusterName(clusterName string) (results []*AllDblink, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDblink{}).Where("`cluster_name` = ?", clusterName).Find(&results).Error

	return
}

// GetBatchFromClusterName 批量查找
func (obj *_AllDblinkMgr) GetBatchFromClusterName(clusterNames []string) (results []*AllDblink, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDblink{}).Where("`cluster_name` IN (?)", clusterNames).Find(&results).Error

	return
}

// GetFromTenantName 通过tenant_name获取内容
func (obj *_AllDblinkMgr) GetFromTenantName(tenantName string) (results []*AllDblink, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDblink{}).Where("`tenant_name` = ?", tenantName).Find(&results).Error

	return
}

// GetBatchFromTenantName 批量查找
func (obj *_AllDblinkMgr) GetBatchFromTenantName(tenantNames []string) (results []*AllDblink, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDblink{}).Where("`tenant_name` IN (?)", tenantNames).Find(&results).Error

	return
}

// GetFromUserName 通过user_name获取内容
func (obj *_AllDblinkMgr) GetFromUserName(userName string) (results []*AllDblink, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDblink{}).Where("`user_name` = ?", userName).Find(&results).Error

	return
}

// GetBatchFromUserName 批量查找
func (obj *_AllDblinkMgr) GetBatchFromUserName(userNames []string) (results []*AllDblink, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDblink{}).Where("`user_name` IN (?)", userNames).Find(&results).Error

	return
}

// GetFromPassword 通过password获取内容
func (obj *_AllDblinkMgr) GetFromPassword(password string) (results []*AllDblink, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDblink{}).Where("`password` = ?", password).Find(&results).Error

	return
}

// GetBatchFromPassword 批量查找
func (obj *_AllDblinkMgr) GetBatchFromPassword(passwords []string) (results []*AllDblink, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDblink{}).Where("`password` IN (?)", passwords).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllDblinkMgr) FetchByPrimaryKey(tenantID int64, dblinkID int64) (result AllDblink, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllDblink{}).Where("`tenant_id` = ? AND `dblink_id` = ?", tenantID, dblinkID).First(&result).Error

	return
}
