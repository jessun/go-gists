package model

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type _AllTenantDependencyMgr struct {
	*_BaseMgr
}

// AllTenantDependencyMgr open func
func AllTenantDependencyMgr(db *gorm.DB) *_AllTenantDependencyMgr {
	if db == nil {
		panic(fmt.Errorf("AllTenantDependencyMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllTenantDependencyMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_tenant_dependency"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllTenantDependencyMgr) GetTableName() string {
	return "__all_tenant_dependency"
}

// Reset 重置gorm会话
func (obj *_AllTenantDependencyMgr) Reset() *_AllTenantDependencyMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllTenantDependencyMgr) Get() (result AllTenantDependency, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantDependency{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllTenantDependencyMgr) Gets() (results []*AllTenantDependency, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantDependency{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllTenantDependencyMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllTenantDependency{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithGmtCreate gmt_create获取
func (obj *_AllTenantDependencyMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllTenantDependencyMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithTenantID tenant_id获取
func (obj *_AllTenantDependencyMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithDepObjType dep_obj_type获取
func (obj *_AllTenantDependencyMgr) WithDepObjType(depObjType int64) Option {
	return optionFunc(func(o *options) { o.query["dep_obj_type"] = depObjType })
}

// WithDepObjID dep_obj_id获取
func (obj *_AllTenantDependencyMgr) WithDepObjID(depObjID int64) Option {
	return optionFunc(func(o *options) { o.query["dep_obj_id"] = depObjID })
}

// WithDepOrder dep_order获取
func (obj *_AllTenantDependencyMgr) WithDepOrder(depOrder int64) Option {
	return optionFunc(func(o *options) { o.query["dep_order"] = depOrder })
}

// WithSchemaVersion schema_version获取
func (obj *_AllTenantDependencyMgr) WithSchemaVersion(schemaVersion int64) Option {
	return optionFunc(func(o *options) { o.query["schema_version"] = schemaVersion })
}

// WithDepTimestamp dep_timestamp获取
func (obj *_AllTenantDependencyMgr) WithDepTimestamp(depTimestamp int64) Option {
	return optionFunc(func(o *options) { o.query["dep_timestamp"] = depTimestamp })
}

// WithRefObjType ref_obj_type获取
func (obj *_AllTenantDependencyMgr) WithRefObjType(refObjType int64) Option {
	return optionFunc(func(o *options) { o.query["ref_obj_type"] = refObjType })
}

// WithRefObjID ref_obj_id获取
func (obj *_AllTenantDependencyMgr) WithRefObjID(refObjID int64) Option {
	return optionFunc(func(o *options) { o.query["ref_obj_id"] = refObjID })
}

// WithRefTimestamp ref_timestamp获取
func (obj *_AllTenantDependencyMgr) WithRefTimestamp(refTimestamp int64) Option {
	return optionFunc(func(o *options) { o.query["ref_timestamp"] = refTimestamp })
}

// WithDepObjOwnerID dep_obj_owner_id获取
func (obj *_AllTenantDependencyMgr) WithDepObjOwnerID(depObjOwnerID int64) Option {
	return optionFunc(func(o *options) { o.query["dep_obj_owner_id"] = depObjOwnerID })
}

// WithProperty property获取
func (obj *_AllTenantDependencyMgr) WithProperty(property int64) Option {
	return optionFunc(func(o *options) { o.query["property"] = property })
}

// WithDepAttrs dep_attrs获取
func (obj *_AllTenantDependencyMgr) WithDepAttrs(depAttrs []byte) Option {
	return optionFunc(func(o *options) { o.query["dep_attrs"] = depAttrs })
}

// WithDepReason dep_reason获取
func (obj *_AllTenantDependencyMgr) WithDepReason(depReason []byte) Option {
	return optionFunc(func(o *options) { o.query["dep_reason"] = depReason })
}

// WithRefObjName ref_obj_name获取
func (obj *_AllTenantDependencyMgr) WithRefObjName(refObjName string) Option {
	return optionFunc(func(o *options) { o.query["ref_obj_name"] = refObjName })
}

// GetByOption 功能选项模式获取
func (obj *_AllTenantDependencyMgr) GetByOption(opts ...Option) (result AllTenantDependency, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllTenantDependency{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllTenantDependencyMgr) GetByOptions(opts ...Option) (results []*AllTenantDependency, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllTenantDependency{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllTenantDependencyMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllTenantDependency, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllTenantDependency{}).Where(options.query)
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
func (obj *_AllTenantDependencyMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllTenantDependency, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantDependency{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllTenantDependencyMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllTenantDependency, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantDependency{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllTenantDependencyMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllTenantDependency, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantDependency{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllTenantDependencyMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllTenantDependency, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantDependency{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromTenantID 通过tenant_id获取内容
func (obj *_AllTenantDependencyMgr) GetFromTenantID(tenantID int64) (results []*AllTenantDependency, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantDependency{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllTenantDependencyMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllTenantDependency, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantDependency{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromDepObjType 通过dep_obj_type获取内容
func (obj *_AllTenantDependencyMgr) GetFromDepObjType(depObjType int64) (results []*AllTenantDependency, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantDependency{}).Where("`dep_obj_type` = ?", depObjType).Find(&results).Error

	return
}

// GetBatchFromDepObjType 批量查找
func (obj *_AllTenantDependencyMgr) GetBatchFromDepObjType(depObjTypes []int64) (results []*AllTenantDependency, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantDependency{}).Where("`dep_obj_type` IN (?)", depObjTypes).Find(&results).Error

	return
}

// GetFromDepObjID 通过dep_obj_id获取内容
func (obj *_AllTenantDependencyMgr) GetFromDepObjID(depObjID int64) (results []*AllTenantDependency, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantDependency{}).Where("`dep_obj_id` = ?", depObjID).Find(&results).Error

	return
}

// GetBatchFromDepObjID 批量查找
func (obj *_AllTenantDependencyMgr) GetBatchFromDepObjID(depObjIDs []int64) (results []*AllTenantDependency, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantDependency{}).Where("`dep_obj_id` IN (?)", depObjIDs).Find(&results).Error

	return
}

// GetFromDepOrder 通过dep_order获取内容
func (obj *_AllTenantDependencyMgr) GetFromDepOrder(depOrder int64) (results []*AllTenantDependency, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantDependency{}).Where("`dep_order` = ?", depOrder).Find(&results).Error

	return
}

// GetBatchFromDepOrder 批量查找
func (obj *_AllTenantDependencyMgr) GetBatchFromDepOrder(depOrders []int64) (results []*AllTenantDependency, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantDependency{}).Where("`dep_order` IN (?)", depOrders).Find(&results).Error

	return
}

// GetFromSchemaVersion 通过schema_version获取内容
func (obj *_AllTenantDependencyMgr) GetFromSchemaVersion(schemaVersion int64) (results []*AllTenantDependency, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantDependency{}).Where("`schema_version` = ?", schemaVersion).Find(&results).Error

	return
}

// GetBatchFromSchemaVersion 批量查找
func (obj *_AllTenantDependencyMgr) GetBatchFromSchemaVersion(schemaVersions []int64) (results []*AllTenantDependency, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantDependency{}).Where("`schema_version` IN (?)", schemaVersions).Find(&results).Error

	return
}

// GetFromDepTimestamp 通过dep_timestamp获取内容
func (obj *_AllTenantDependencyMgr) GetFromDepTimestamp(depTimestamp int64) (results []*AllTenantDependency, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantDependency{}).Where("`dep_timestamp` = ?", depTimestamp).Find(&results).Error

	return
}

// GetBatchFromDepTimestamp 批量查找
func (obj *_AllTenantDependencyMgr) GetBatchFromDepTimestamp(depTimestamps []int64) (results []*AllTenantDependency, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantDependency{}).Where("`dep_timestamp` IN (?)", depTimestamps).Find(&results).Error

	return
}

// GetFromRefObjType 通过ref_obj_type获取内容
func (obj *_AllTenantDependencyMgr) GetFromRefObjType(refObjType int64) (results []*AllTenantDependency, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantDependency{}).Where("`ref_obj_type` = ?", refObjType).Find(&results).Error

	return
}

// GetBatchFromRefObjType 批量查找
func (obj *_AllTenantDependencyMgr) GetBatchFromRefObjType(refObjTypes []int64) (results []*AllTenantDependency, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantDependency{}).Where("`ref_obj_type` IN (?)", refObjTypes).Find(&results).Error

	return
}

// GetFromRefObjID 通过ref_obj_id获取内容
func (obj *_AllTenantDependencyMgr) GetFromRefObjID(refObjID int64) (results []*AllTenantDependency, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantDependency{}).Where("`ref_obj_id` = ?", refObjID).Find(&results).Error

	return
}

// GetBatchFromRefObjID 批量查找
func (obj *_AllTenantDependencyMgr) GetBatchFromRefObjID(refObjIDs []int64) (results []*AllTenantDependency, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantDependency{}).Where("`ref_obj_id` IN (?)", refObjIDs).Find(&results).Error

	return
}

// GetFromRefTimestamp 通过ref_timestamp获取内容
func (obj *_AllTenantDependencyMgr) GetFromRefTimestamp(refTimestamp int64) (results []*AllTenantDependency, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantDependency{}).Where("`ref_timestamp` = ?", refTimestamp).Find(&results).Error

	return
}

// GetBatchFromRefTimestamp 批量查找
func (obj *_AllTenantDependencyMgr) GetBatchFromRefTimestamp(refTimestamps []int64) (results []*AllTenantDependency, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantDependency{}).Where("`ref_timestamp` IN (?)", refTimestamps).Find(&results).Error

	return
}

// GetFromDepObjOwnerID 通过dep_obj_owner_id获取内容
func (obj *_AllTenantDependencyMgr) GetFromDepObjOwnerID(depObjOwnerID int64) (results []*AllTenantDependency, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantDependency{}).Where("`dep_obj_owner_id` = ?", depObjOwnerID).Find(&results).Error

	return
}

// GetBatchFromDepObjOwnerID 批量查找
func (obj *_AllTenantDependencyMgr) GetBatchFromDepObjOwnerID(depObjOwnerIDs []int64) (results []*AllTenantDependency, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantDependency{}).Where("`dep_obj_owner_id` IN (?)", depObjOwnerIDs).Find(&results).Error

	return
}

// GetFromProperty 通过property获取内容
func (obj *_AllTenantDependencyMgr) GetFromProperty(property int64) (results []*AllTenantDependency, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantDependency{}).Where("`property` = ?", property).Find(&results).Error

	return
}

// GetBatchFromProperty 批量查找
func (obj *_AllTenantDependencyMgr) GetBatchFromProperty(propertys []int64) (results []*AllTenantDependency, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantDependency{}).Where("`property` IN (?)", propertys).Find(&results).Error

	return
}

// GetFromDepAttrs 通过dep_attrs获取内容
func (obj *_AllTenantDependencyMgr) GetFromDepAttrs(depAttrs []byte) (results []*AllTenantDependency, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantDependency{}).Where("`dep_attrs` = ?", depAttrs).Find(&results).Error

	return
}

// GetBatchFromDepAttrs 批量查找
func (obj *_AllTenantDependencyMgr) GetBatchFromDepAttrs(depAttrss [][]byte) (results []*AllTenantDependency, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantDependency{}).Where("`dep_attrs` IN (?)", depAttrss).Find(&results).Error

	return
}

// GetFromDepReason 通过dep_reason获取内容
func (obj *_AllTenantDependencyMgr) GetFromDepReason(depReason []byte) (results []*AllTenantDependency, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantDependency{}).Where("`dep_reason` = ?", depReason).Find(&results).Error

	return
}

// GetBatchFromDepReason 批量查找
func (obj *_AllTenantDependencyMgr) GetBatchFromDepReason(depReasons [][]byte) (results []*AllTenantDependency, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantDependency{}).Where("`dep_reason` IN (?)", depReasons).Find(&results).Error

	return
}

// GetFromRefObjName 通过ref_obj_name获取内容
func (obj *_AllTenantDependencyMgr) GetFromRefObjName(refObjName string) (results []*AllTenantDependency, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantDependency{}).Where("`ref_obj_name` = ?", refObjName).Find(&results).Error

	return
}

// GetBatchFromRefObjName 批量查找
func (obj *_AllTenantDependencyMgr) GetBatchFromRefObjName(refObjNames []string) (results []*AllTenantDependency, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantDependency{}).Where("`ref_obj_name` IN (?)", refObjNames).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllTenantDependencyMgr) FetchByPrimaryKey(tenantID int64, depObjType int64, depObjID int64, depOrder int64) (result AllTenantDependency, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllTenantDependency{}).Where("`tenant_id` = ? AND `dep_obj_type` = ? AND `dep_obj_id` = ? AND `dep_order` = ?", tenantID, depObjType, depObjID, depOrder).First(&result).Error

	return
}
