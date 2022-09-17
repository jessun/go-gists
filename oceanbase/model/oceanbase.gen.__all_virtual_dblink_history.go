package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _AllVirtualDblinkHistoryMgr struct {
	*_BaseMgr
}

// AllVirtualDblinkHistoryMgr open func
func AllVirtualDblinkHistoryMgr(db *gorm.DB) *_AllVirtualDblinkHistoryMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualDblinkHistoryMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualDblinkHistoryMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_dblink_history"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualDblinkHistoryMgr) GetTableName() string {
	return "__all_virtual_dblink_history"
}

// Reset 重置gorm会话
func (obj *_AllVirtualDblinkHistoryMgr) Reset() *_AllVirtualDblinkHistoryMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualDblinkHistoryMgr) Get() (result AllVirtualDblinkHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDblinkHistory{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualDblinkHistoryMgr) Gets() (results []*AllVirtualDblinkHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDblinkHistory{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualDblinkHistoryMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualDblinkHistory{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTenantID tenant_id获取
func (obj *_AllVirtualDblinkHistoryMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithDblinkID dblink_id获取
func (obj *_AllVirtualDblinkHistoryMgr) WithDblinkID(dblinkID int64) Option {
	return optionFunc(func(o *options) { o.query["dblink_id"] = dblinkID })
}

// WithSchemaVersion schema_version获取
func (obj *_AllVirtualDblinkHistoryMgr) WithSchemaVersion(schemaVersion int64) Option {
	return optionFunc(func(o *options) { o.query["schema_version"] = schemaVersion })
}

// WithGmtCreate gmt_create获取
func (obj *_AllVirtualDblinkHistoryMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllVirtualDblinkHistoryMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithIsDeleted is_deleted获取
func (obj *_AllVirtualDblinkHistoryMgr) WithIsDeleted(isDeleted int64) Option {
	return optionFunc(func(o *options) { o.query["is_deleted"] = isDeleted })
}

// WithDblinkName dblink_name获取
func (obj *_AllVirtualDblinkHistoryMgr) WithDblinkName(dblinkName string) Option {
	return optionFunc(func(o *options) { o.query["dblink_name"] = dblinkName })
}

// WithOwnerID owner_id获取
func (obj *_AllVirtualDblinkHistoryMgr) WithOwnerID(ownerID int64) Option {
	return optionFunc(func(o *options) { o.query["owner_id"] = ownerID })
}

// WithHostIP host_ip获取
func (obj *_AllVirtualDblinkHistoryMgr) WithHostIP(hostIP string) Option {
	return optionFunc(func(o *options) { o.query["host_ip"] = hostIP })
}

// WithHostPort host_port获取
func (obj *_AllVirtualDblinkHistoryMgr) WithHostPort(hostPort int64) Option {
	return optionFunc(func(o *options) { o.query["host_port"] = hostPort })
}

// WithClusterName cluster_name获取
func (obj *_AllVirtualDblinkHistoryMgr) WithClusterName(clusterName string) Option {
	return optionFunc(func(o *options) { o.query["cluster_name"] = clusterName })
}

// WithTenantName tenant_name获取
func (obj *_AllVirtualDblinkHistoryMgr) WithTenantName(tenantName string) Option {
	return optionFunc(func(o *options) { o.query["tenant_name"] = tenantName })
}

// WithUserName user_name获取
func (obj *_AllVirtualDblinkHistoryMgr) WithUserName(userName string) Option {
	return optionFunc(func(o *options) { o.query["user_name"] = userName })
}

// WithPassword password获取
func (obj *_AllVirtualDblinkHistoryMgr) WithPassword(password string) Option {
	return optionFunc(func(o *options) { o.query["password"] = password })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualDblinkHistoryMgr) GetByOption(opts ...Option) (result AllVirtualDblinkHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDblinkHistory{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualDblinkHistoryMgr) GetByOptions(opts ...Option) (results []*AllVirtualDblinkHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDblinkHistory{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualDblinkHistoryMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualDblinkHistory, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualDblinkHistory{}).Where(options.query)
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
func (obj *_AllVirtualDblinkHistoryMgr) GetFromTenantID(tenantID int64) (results []*AllVirtualDblinkHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDblinkHistory{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllVirtualDblinkHistoryMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllVirtualDblinkHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDblinkHistory{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromDblinkID 通过dblink_id获取内容
func (obj *_AllVirtualDblinkHistoryMgr) GetFromDblinkID(dblinkID int64) (results []*AllVirtualDblinkHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDblinkHistory{}).Where("`dblink_id` = ?", dblinkID).Find(&results).Error

	return
}

// GetBatchFromDblinkID 批量查找
func (obj *_AllVirtualDblinkHistoryMgr) GetBatchFromDblinkID(dblinkIDs []int64) (results []*AllVirtualDblinkHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDblinkHistory{}).Where("`dblink_id` IN (?)", dblinkIDs).Find(&results).Error

	return
}

// GetFromSchemaVersion 通过schema_version获取内容
func (obj *_AllVirtualDblinkHistoryMgr) GetFromSchemaVersion(schemaVersion int64) (results []*AllVirtualDblinkHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDblinkHistory{}).Where("`schema_version` = ?", schemaVersion).Find(&results).Error

	return
}

// GetBatchFromSchemaVersion 批量查找
func (obj *_AllVirtualDblinkHistoryMgr) GetBatchFromSchemaVersion(schemaVersions []int64) (results []*AllVirtualDblinkHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDblinkHistory{}).Where("`schema_version` IN (?)", schemaVersions).Find(&results).Error

	return
}

// GetFromGmtCreate 通过gmt_create获取内容
func (obj *_AllVirtualDblinkHistoryMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllVirtualDblinkHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDblinkHistory{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllVirtualDblinkHistoryMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllVirtualDblinkHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDblinkHistory{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllVirtualDblinkHistoryMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllVirtualDblinkHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDblinkHistory{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllVirtualDblinkHistoryMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllVirtualDblinkHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDblinkHistory{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromIsDeleted 通过is_deleted获取内容
func (obj *_AllVirtualDblinkHistoryMgr) GetFromIsDeleted(isDeleted int64) (results []*AllVirtualDblinkHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDblinkHistory{}).Where("`is_deleted` = ?", isDeleted).Find(&results).Error

	return
}

// GetBatchFromIsDeleted 批量查找
func (obj *_AllVirtualDblinkHistoryMgr) GetBatchFromIsDeleted(isDeleteds []int64) (results []*AllVirtualDblinkHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDblinkHistory{}).Where("`is_deleted` IN (?)", isDeleteds).Find(&results).Error

	return
}

// GetFromDblinkName 通过dblink_name获取内容
func (obj *_AllVirtualDblinkHistoryMgr) GetFromDblinkName(dblinkName string) (results []*AllVirtualDblinkHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDblinkHistory{}).Where("`dblink_name` = ?", dblinkName).Find(&results).Error

	return
}

// GetBatchFromDblinkName 批量查找
func (obj *_AllVirtualDblinkHistoryMgr) GetBatchFromDblinkName(dblinkNames []string) (results []*AllVirtualDblinkHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDblinkHistory{}).Where("`dblink_name` IN (?)", dblinkNames).Find(&results).Error

	return
}

// GetFromOwnerID 通过owner_id获取内容
func (obj *_AllVirtualDblinkHistoryMgr) GetFromOwnerID(ownerID int64) (results []*AllVirtualDblinkHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDblinkHistory{}).Where("`owner_id` = ?", ownerID).Find(&results).Error

	return
}

// GetBatchFromOwnerID 批量查找
func (obj *_AllVirtualDblinkHistoryMgr) GetBatchFromOwnerID(ownerIDs []int64) (results []*AllVirtualDblinkHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDblinkHistory{}).Where("`owner_id` IN (?)", ownerIDs).Find(&results).Error

	return
}

// GetFromHostIP 通过host_ip获取内容
func (obj *_AllVirtualDblinkHistoryMgr) GetFromHostIP(hostIP string) (results []*AllVirtualDblinkHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDblinkHistory{}).Where("`host_ip` = ?", hostIP).Find(&results).Error

	return
}

// GetBatchFromHostIP 批量查找
func (obj *_AllVirtualDblinkHistoryMgr) GetBatchFromHostIP(hostIPs []string) (results []*AllVirtualDblinkHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDblinkHistory{}).Where("`host_ip` IN (?)", hostIPs).Find(&results).Error

	return
}

// GetFromHostPort 通过host_port获取内容
func (obj *_AllVirtualDblinkHistoryMgr) GetFromHostPort(hostPort int64) (results []*AllVirtualDblinkHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDblinkHistory{}).Where("`host_port` = ?", hostPort).Find(&results).Error

	return
}

// GetBatchFromHostPort 批量查找
func (obj *_AllVirtualDblinkHistoryMgr) GetBatchFromHostPort(hostPorts []int64) (results []*AllVirtualDblinkHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDblinkHistory{}).Where("`host_port` IN (?)", hostPorts).Find(&results).Error

	return
}

// GetFromClusterName 通过cluster_name获取内容
func (obj *_AllVirtualDblinkHistoryMgr) GetFromClusterName(clusterName string) (results []*AllVirtualDblinkHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDblinkHistory{}).Where("`cluster_name` = ?", clusterName).Find(&results).Error

	return
}

// GetBatchFromClusterName 批量查找
func (obj *_AllVirtualDblinkHistoryMgr) GetBatchFromClusterName(clusterNames []string) (results []*AllVirtualDblinkHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDblinkHistory{}).Where("`cluster_name` IN (?)", clusterNames).Find(&results).Error

	return
}

// GetFromTenantName 通过tenant_name获取内容
func (obj *_AllVirtualDblinkHistoryMgr) GetFromTenantName(tenantName string) (results []*AllVirtualDblinkHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDblinkHistory{}).Where("`tenant_name` = ?", tenantName).Find(&results).Error

	return
}

// GetBatchFromTenantName 批量查找
func (obj *_AllVirtualDblinkHistoryMgr) GetBatchFromTenantName(tenantNames []string) (results []*AllVirtualDblinkHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDblinkHistory{}).Where("`tenant_name` IN (?)", tenantNames).Find(&results).Error

	return
}

// GetFromUserName 通过user_name获取内容
func (obj *_AllVirtualDblinkHistoryMgr) GetFromUserName(userName string) (results []*AllVirtualDblinkHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDblinkHistory{}).Where("`user_name` = ?", userName).Find(&results).Error

	return
}

// GetBatchFromUserName 批量查找
func (obj *_AllVirtualDblinkHistoryMgr) GetBatchFromUserName(userNames []string) (results []*AllVirtualDblinkHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDblinkHistory{}).Where("`user_name` IN (?)", userNames).Find(&results).Error

	return
}

// GetFromPassword 通过password获取内容
func (obj *_AllVirtualDblinkHistoryMgr) GetFromPassword(password string) (results []*AllVirtualDblinkHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDblinkHistory{}).Where("`password` = ?", password).Find(&results).Error

	return
}

// GetBatchFromPassword 批量查找
func (obj *_AllVirtualDblinkHistoryMgr) GetBatchFromPassword(passwords []string) (results []*AllVirtualDblinkHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDblinkHistory{}).Where("`password` IN (?)", passwords).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllVirtualDblinkHistoryMgr) FetchByPrimaryKey(tenantID int64, dblinkID int64, schemaVersion int64) (result AllVirtualDblinkHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualDblinkHistory{}).Where("`tenant_id` = ? AND `dblink_id` = ? AND `schema_version` = ?", tenantID, dblinkID, schemaVersion).First(&result).Error

	return
}
