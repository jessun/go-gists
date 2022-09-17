package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _AllVirtualTypeMgr struct {
	*_BaseMgr
}

// AllVirtualTypeMgr open func
func AllVirtualTypeMgr(db *gorm.DB) *_AllVirtualTypeMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualTypeMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualTypeMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_type"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualTypeMgr) GetTableName() string {
	return "__all_virtual_type"
}

// Reset 重置gorm会话
func (obj *_AllVirtualTypeMgr) Reset() *_AllVirtualTypeMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualTypeMgr) Get() (result AllVirtualType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualType{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualTypeMgr) Gets() (results []*AllVirtualType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualType{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualTypeMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualType{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTenantID tenant_id获取
func (obj *_AllVirtualTypeMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithTypeID type_id获取
func (obj *_AllVirtualTypeMgr) WithTypeID(typeID int64) Option {
	return optionFunc(func(o *options) { o.query["type_id"] = typeID })
}

// WithGmtCreate gmt_create获取
func (obj *_AllVirtualTypeMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllVirtualTypeMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithDatabaseID database_id获取
func (obj *_AllVirtualTypeMgr) WithDatabaseID(databaseID int64) Option {
	return optionFunc(func(o *options) { o.query["database_id"] = databaseID })
}

// WithSchemaVersion schema_version获取
func (obj *_AllVirtualTypeMgr) WithSchemaVersion(schemaVersion int64) Option {
	return optionFunc(func(o *options) { o.query["schema_version"] = schemaVersion })
}

// WithTypecode typecode获取
func (obj *_AllVirtualTypeMgr) WithTypecode(typecode int64) Option {
	return optionFunc(func(o *options) { o.query["typecode"] = typecode })
}

// WithProperties properties获取
func (obj *_AllVirtualTypeMgr) WithProperties(properties int64) Option {
	return optionFunc(func(o *options) { o.query["properties"] = properties })
}

// WithAttributes attributes获取
func (obj *_AllVirtualTypeMgr) WithAttributes(attributes int64) Option {
	return optionFunc(func(o *options) { o.query["attributes"] = attributes })
}

// WithMethods methods获取
func (obj *_AllVirtualTypeMgr) WithMethods(methods int64) Option {
	return optionFunc(func(o *options) { o.query["methods"] = methods })
}

// WithHiddenmethods hiddenmethods获取
func (obj *_AllVirtualTypeMgr) WithHiddenmethods(hiddenmethods int64) Option {
	return optionFunc(func(o *options) { o.query["hiddenmethods"] = hiddenmethods })
}

// WithSupertypes supertypes获取
func (obj *_AllVirtualTypeMgr) WithSupertypes(supertypes int64) Option {
	return optionFunc(func(o *options) { o.query["supertypes"] = supertypes })
}

// WithSubtypes subtypes获取
func (obj *_AllVirtualTypeMgr) WithSubtypes(subtypes int64) Option {
	return optionFunc(func(o *options) { o.query["subtypes"] = subtypes })
}

// WithExterntype externtype获取
func (obj *_AllVirtualTypeMgr) WithExterntype(externtype int64) Option {
	return optionFunc(func(o *options) { o.query["externtype"] = externtype })
}

// WithExternname externname获取
func (obj *_AllVirtualTypeMgr) WithExternname(externname string) Option {
	return optionFunc(func(o *options) { o.query["externname"] = externname })
}

// WithHelperclassname helperclassname获取
func (obj *_AllVirtualTypeMgr) WithHelperclassname(helperclassname string) Option {
	return optionFunc(func(o *options) { o.query["helperclassname"] = helperclassname })
}

// WithLocalAttrs local_attrs获取
func (obj *_AllVirtualTypeMgr) WithLocalAttrs(localAttrs int64) Option {
	return optionFunc(func(o *options) { o.query["local_attrs"] = localAttrs })
}

// WithLocalMethods local_methods获取
func (obj *_AllVirtualTypeMgr) WithLocalMethods(localMethods int64) Option {
	return optionFunc(func(o *options) { o.query["local_methods"] = localMethods })
}

// WithSupertypeid supertypeid获取
func (obj *_AllVirtualTypeMgr) WithSupertypeid(supertypeid int64) Option {
	return optionFunc(func(o *options) { o.query["supertypeid"] = supertypeid })
}

// WithTypeName type_name获取
func (obj *_AllVirtualTypeMgr) WithTypeName(typeName string) Option {
	return optionFunc(func(o *options) { o.query["type_name"] = typeName })
}

// WithPackageID package_id获取
func (obj *_AllVirtualTypeMgr) WithPackageID(packageID int64) Option {
	return optionFunc(func(o *options) { o.query["package_id"] = packageID })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualTypeMgr) GetByOption(opts ...Option) (result AllVirtualType, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualType{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualTypeMgr) GetByOptions(opts ...Option) (results []*AllVirtualType, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualType{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualTypeMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualType, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualType{}).Where(options.query)
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
func (obj *_AllVirtualTypeMgr) GetFromTenantID(tenantID int64) (results []*AllVirtualType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualType{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllVirtualTypeMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllVirtualType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualType{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromTypeID 通过type_id获取内容
func (obj *_AllVirtualTypeMgr) GetFromTypeID(typeID int64) (results []*AllVirtualType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualType{}).Where("`type_id` = ?", typeID).Find(&results).Error

	return
}

// GetBatchFromTypeID 批量查找
func (obj *_AllVirtualTypeMgr) GetBatchFromTypeID(typeIDs []int64) (results []*AllVirtualType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualType{}).Where("`type_id` IN (?)", typeIDs).Find(&results).Error

	return
}

// GetFromGmtCreate 通过gmt_create获取内容
func (obj *_AllVirtualTypeMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllVirtualType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualType{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllVirtualTypeMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllVirtualType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualType{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllVirtualTypeMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllVirtualType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualType{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllVirtualTypeMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllVirtualType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualType{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromDatabaseID 通过database_id获取内容
func (obj *_AllVirtualTypeMgr) GetFromDatabaseID(databaseID int64) (results []*AllVirtualType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualType{}).Where("`database_id` = ?", databaseID).Find(&results).Error

	return
}

// GetBatchFromDatabaseID 批量查找
func (obj *_AllVirtualTypeMgr) GetBatchFromDatabaseID(databaseIDs []int64) (results []*AllVirtualType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualType{}).Where("`database_id` IN (?)", databaseIDs).Find(&results).Error

	return
}

// GetFromSchemaVersion 通过schema_version获取内容
func (obj *_AllVirtualTypeMgr) GetFromSchemaVersion(schemaVersion int64) (results []*AllVirtualType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualType{}).Where("`schema_version` = ?", schemaVersion).Find(&results).Error

	return
}

// GetBatchFromSchemaVersion 批量查找
func (obj *_AllVirtualTypeMgr) GetBatchFromSchemaVersion(schemaVersions []int64) (results []*AllVirtualType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualType{}).Where("`schema_version` IN (?)", schemaVersions).Find(&results).Error

	return
}

// GetFromTypecode 通过typecode获取内容
func (obj *_AllVirtualTypeMgr) GetFromTypecode(typecode int64) (results []*AllVirtualType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualType{}).Where("`typecode` = ?", typecode).Find(&results).Error

	return
}

// GetBatchFromTypecode 批量查找
func (obj *_AllVirtualTypeMgr) GetBatchFromTypecode(typecodes []int64) (results []*AllVirtualType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualType{}).Where("`typecode` IN (?)", typecodes).Find(&results).Error

	return
}

// GetFromProperties 通过properties获取内容
func (obj *_AllVirtualTypeMgr) GetFromProperties(properties int64) (results []*AllVirtualType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualType{}).Where("`properties` = ?", properties).Find(&results).Error

	return
}

// GetBatchFromProperties 批量查找
func (obj *_AllVirtualTypeMgr) GetBatchFromProperties(propertiess []int64) (results []*AllVirtualType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualType{}).Where("`properties` IN (?)", propertiess).Find(&results).Error

	return
}

// GetFromAttributes 通过attributes获取内容
func (obj *_AllVirtualTypeMgr) GetFromAttributes(attributes int64) (results []*AllVirtualType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualType{}).Where("`attributes` = ?", attributes).Find(&results).Error

	return
}

// GetBatchFromAttributes 批量查找
func (obj *_AllVirtualTypeMgr) GetBatchFromAttributes(attributess []int64) (results []*AllVirtualType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualType{}).Where("`attributes` IN (?)", attributess).Find(&results).Error

	return
}

// GetFromMethods 通过methods获取内容
func (obj *_AllVirtualTypeMgr) GetFromMethods(methods int64) (results []*AllVirtualType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualType{}).Where("`methods` = ?", methods).Find(&results).Error

	return
}

// GetBatchFromMethods 批量查找
func (obj *_AllVirtualTypeMgr) GetBatchFromMethods(methodss []int64) (results []*AllVirtualType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualType{}).Where("`methods` IN (?)", methodss).Find(&results).Error

	return
}

// GetFromHiddenmethods 通过hiddenmethods获取内容
func (obj *_AllVirtualTypeMgr) GetFromHiddenmethods(hiddenmethods int64) (results []*AllVirtualType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualType{}).Where("`hiddenmethods` = ?", hiddenmethods).Find(&results).Error

	return
}

// GetBatchFromHiddenmethods 批量查找
func (obj *_AllVirtualTypeMgr) GetBatchFromHiddenmethods(hiddenmethodss []int64) (results []*AllVirtualType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualType{}).Where("`hiddenmethods` IN (?)", hiddenmethodss).Find(&results).Error

	return
}

// GetFromSupertypes 通过supertypes获取内容
func (obj *_AllVirtualTypeMgr) GetFromSupertypes(supertypes int64) (results []*AllVirtualType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualType{}).Where("`supertypes` = ?", supertypes).Find(&results).Error

	return
}

// GetBatchFromSupertypes 批量查找
func (obj *_AllVirtualTypeMgr) GetBatchFromSupertypes(supertypess []int64) (results []*AllVirtualType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualType{}).Where("`supertypes` IN (?)", supertypess).Find(&results).Error

	return
}

// GetFromSubtypes 通过subtypes获取内容
func (obj *_AllVirtualTypeMgr) GetFromSubtypes(subtypes int64) (results []*AllVirtualType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualType{}).Where("`subtypes` = ?", subtypes).Find(&results).Error

	return
}

// GetBatchFromSubtypes 批量查找
func (obj *_AllVirtualTypeMgr) GetBatchFromSubtypes(subtypess []int64) (results []*AllVirtualType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualType{}).Where("`subtypes` IN (?)", subtypess).Find(&results).Error

	return
}

// GetFromExterntype 通过externtype获取内容
func (obj *_AllVirtualTypeMgr) GetFromExterntype(externtype int64) (results []*AllVirtualType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualType{}).Where("`externtype` = ?", externtype).Find(&results).Error

	return
}

// GetBatchFromExterntype 批量查找
func (obj *_AllVirtualTypeMgr) GetBatchFromExterntype(externtypes []int64) (results []*AllVirtualType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualType{}).Where("`externtype` IN (?)", externtypes).Find(&results).Error

	return
}

// GetFromExternname 通过externname获取内容
func (obj *_AllVirtualTypeMgr) GetFromExternname(externname string) (results []*AllVirtualType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualType{}).Where("`externname` = ?", externname).Find(&results).Error

	return
}

// GetBatchFromExternname 批量查找
func (obj *_AllVirtualTypeMgr) GetBatchFromExternname(externnames []string) (results []*AllVirtualType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualType{}).Where("`externname` IN (?)", externnames).Find(&results).Error

	return
}

// GetFromHelperclassname 通过helperclassname获取内容
func (obj *_AllVirtualTypeMgr) GetFromHelperclassname(helperclassname string) (results []*AllVirtualType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualType{}).Where("`helperclassname` = ?", helperclassname).Find(&results).Error

	return
}

// GetBatchFromHelperclassname 批量查找
func (obj *_AllVirtualTypeMgr) GetBatchFromHelperclassname(helperclassnames []string) (results []*AllVirtualType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualType{}).Where("`helperclassname` IN (?)", helperclassnames).Find(&results).Error

	return
}

// GetFromLocalAttrs 通过local_attrs获取内容
func (obj *_AllVirtualTypeMgr) GetFromLocalAttrs(localAttrs int64) (results []*AllVirtualType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualType{}).Where("`local_attrs` = ?", localAttrs).Find(&results).Error

	return
}

// GetBatchFromLocalAttrs 批量查找
func (obj *_AllVirtualTypeMgr) GetBatchFromLocalAttrs(localAttrss []int64) (results []*AllVirtualType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualType{}).Where("`local_attrs` IN (?)", localAttrss).Find(&results).Error

	return
}

// GetFromLocalMethods 通过local_methods获取内容
func (obj *_AllVirtualTypeMgr) GetFromLocalMethods(localMethods int64) (results []*AllVirtualType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualType{}).Where("`local_methods` = ?", localMethods).Find(&results).Error

	return
}

// GetBatchFromLocalMethods 批量查找
func (obj *_AllVirtualTypeMgr) GetBatchFromLocalMethods(localMethodss []int64) (results []*AllVirtualType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualType{}).Where("`local_methods` IN (?)", localMethodss).Find(&results).Error

	return
}

// GetFromSupertypeid 通过supertypeid获取内容
func (obj *_AllVirtualTypeMgr) GetFromSupertypeid(supertypeid int64) (results []*AllVirtualType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualType{}).Where("`supertypeid` = ?", supertypeid).Find(&results).Error

	return
}

// GetBatchFromSupertypeid 批量查找
func (obj *_AllVirtualTypeMgr) GetBatchFromSupertypeid(supertypeids []int64) (results []*AllVirtualType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualType{}).Where("`supertypeid` IN (?)", supertypeids).Find(&results).Error

	return
}

// GetFromTypeName 通过type_name获取内容
func (obj *_AllVirtualTypeMgr) GetFromTypeName(typeName string) (results []*AllVirtualType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualType{}).Where("`type_name` = ?", typeName).Find(&results).Error

	return
}

// GetBatchFromTypeName 批量查找
func (obj *_AllVirtualTypeMgr) GetBatchFromTypeName(typeNames []string) (results []*AllVirtualType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualType{}).Where("`type_name` IN (?)", typeNames).Find(&results).Error

	return
}

// GetFromPackageID 通过package_id获取内容
func (obj *_AllVirtualTypeMgr) GetFromPackageID(packageID int64) (results []*AllVirtualType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualType{}).Where("`package_id` = ?", packageID).Find(&results).Error

	return
}

// GetBatchFromPackageID 批量查找
func (obj *_AllVirtualTypeMgr) GetBatchFromPackageID(packageIDs []int64) (results []*AllVirtualType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualType{}).Where("`package_id` IN (?)", packageIDs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllVirtualTypeMgr) FetchByPrimaryKey(tenantID int64, typeID int64) (result AllVirtualType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualType{}).Where("`tenant_id` = ? AND `type_id` = ?", tenantID, typeID).First(&result).Error

	return
}
