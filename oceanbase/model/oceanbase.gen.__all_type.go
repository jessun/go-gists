package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _AllTypeMgr struct {
	*_BaseMgr
}

// AllTypeMgr open func
func AllTypeMgr(db *gorm.DB) *_AllTypeMgr {
	if db == nil {
		panic(fmt.Errorf("AllTypeMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllTypeMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_type"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllTypeMgr) GetTableName() string {
	return "__all_type"
}

// Reset 重置gorm会话
func (obj *_AllTypeMgr) Reset() *_AllTypeMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllTypeMgr) Get() (result AllType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllType{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllTypeMgr) Gets() (results []*AllType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllType{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllTypeMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllType{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithGmtCreate gmt_create获取
func (obj *_AllTypeMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllTypeMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithTenantID tenant_id获取
func (obj *_AllTypeMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithTypeID type_id获取
func (obj *_AllTypeMgr) WithTypeID(typeID int64) Option {
	return optionFunc(func(o *options) { o.query["type_id"] = typeID })
}

// WithDatabaseID database_id获取
func (obj *_AllTypeMgr) WithDatabaseID(databaseID int64) Option {
	return optionFunc(func(o *options) { o.query["database_id"] = databaseID })
}

// WithSchemaVersion schema_version获取
func (obj *_AllTypeMgr) WithSchemaVersion(schemaVersion int64) Option {
	return optionFunc(func(o *options) { o.query["schema_version"] = schemaVersion })
}

// WithTypecode typecode获取
func (obj *_AllTypeMgr) WithTypecode(typecode int64) Option {
	return optionFunc(func(o *options) { o.query["typecode"] = typecode })
}

// WithProperties properties获取
func (obj *_AllTypeMgr) WithProperties(properties int64) Option {
	return optionFunc(func(o *options) { o.query["properties"] = properties })
}

// WithAttributes attributes获取
func (obj *_AllTypeMgr) WithAttributes(attributes int64) Option {
	return optionFunc(func(o *options) { o.query["attributes"] = attributes })
}

// WithMethods methods获取
func (obj *_AllTypeMgr) WithMethods(methods int64) Option {
	return optionFunc(func(o *options) { o.query["methods"] = methods })
}

// WithHiddenmethods hiddenmethods获取
func (obj *_AllTypeMgr) WithHiddenmethods(hiddenmethods int64) Option {
	return optionFunc(func(o *options) { o.query["hiddenmethods"] = hiddenmethods })
}

// WithSupertypes supertypes获取
func (obj *_AllTypeMgr) WithSupertypes(supertypes int64) Option {
	return optionFunc(func(o *options) { o.query["supertypes"] = supertypes })
}

// WithSubtypes subtypes获取
func (obj *_AllTypeMgr) WithSubtypes(subtypes int64) Option {
	return optionFunc(func(o *options) { o.query["subtypes"] = subtypes })
}

// WithExterntype externtype获取
func (obj *_AllTypeMgr) WithExterntype(externtype int64) Option {
	return optionFunc(func(o *options) { o.query["externtype"] = externtype })
}

// WithExternname externname获取
func (obj *_AllTypeMgr) WithExternname(externname string) Option {
	return optionFunc(func(o *options) { o.query["externname"] = externname })
}

// WithHelperclassname helperclassname获取
func (obj *_AllTypeMgr) WithHelperclassname(helperclassname string) Option {
	return optionFunc(func(o *options) { o.query["helperclassname"] = helperclassname })
}

// WithLocalAttrs local_attrs获取
func (obj *_AllTypeMgr) WithLocalAttrs(localAttrs int64) Option {
	return optionFunc(func(o *options) { o.query["local_attrs"] = localAttrs })
}

// WithLocalMethods local_methods获取
func (obj *_AllTypeMgr) WithLocalMethods(localMethods int64) Option {
	return optionFunc(func(o *options) { o.query["local_methods"] = localMethods })
}

// WithSupertypeid supertypeid获取
func (obj *_AllTypeMgr) WithSupertypeid(supertypeid int64) Option {
	return optionFunc(func(o *options) { o.query["supertypeid"] = supertypeid })
}

// WithTypeName type_name获取
func (obj *_AllTypeMgr) WithTypeName(typeName string) Option {
	return optionFunc(func(o *options) { o.query["type_name"] = typeName })
}

// WithPackageID package_id获取
func (obj *_AllTypeMgr) WithPackageID(packageID int64) Option {
	return optionFunc(func(o *options) { o.query["package_id"] = packageID })
}

// GetByOption 功能选项模式获取
func (obj *_AllTypeMgr) GetByOption(opts ...Option) (result AllType, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllType{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllTypeMgr) GetByOptions(opts ...Option) (results []*AllType, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllType{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllTypeMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllType, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllType{}).Where(options.query)
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
func (obj *_AllTypeMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllType{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllTypeMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllType{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllTypeMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllType{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllTypeMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllType{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromTenantID 通过tenant_id获取内容
func (obj *_AllTypeMgr) GetFromTenantID(tenantID int64) (results []*AllType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllType{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllTypeMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllType{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromTypeID 通过type_id获取内容
func (obj *_AllTypeMgr) GetFromTypeID(typeID int64) (results []*AllType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllType{}).Where("`type_id` = ?", typeID).Find(&results).Error

	return
}

// GetBatchFromTypeID 批量查找
func (obj *_AllTypeMgr) GetBatchFromTypeID(typeIDs []int64) (results []*AllType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllType{}).Where("`type_id` IN (?)", typeIDs).Find(&results).Error

	return
}

// GetFromDatabaseID 通过database_id获取内容
func (obj *_AllTypeMgr) GetFromDatabaseID(databaseID int64) (results []*AllType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllType{}).Where("`database_id` = ?", databaseID).Find(&results).Error

	return
}

// GetBatchFromDatabaseID 批量查找
func (obj *_AllTypeMgr) GetBatchFromDatabaseID(databaseIDs []int64) (results []*AllType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllType{}).Where("`database_id` IN (?)", databaseIDs).Find(&results).Error

	return
}

// GetFromSchemaVersion 通过schema_version获取内容
func (obj *_AllTypeMgr) GetFromSchemaVersion(schemaVersion int64) (results []*AllType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllType{}).Where("`schema_version` = ?", schemaVersion).Find(&results).Error

	return
}

// GetBatchFromSchemaVersion 批量查找
func (obj *_AllTypeMgr) GetBatchFromSchemaVersion(schemaVersions []int64) (results []*AllType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllType{}).Where("`schema_version` IN (?)", schemaVersions).Find(&results).Error

	return
}

// GetFromTypecode 通过typecode获取内容
func (obj *_AllTypeMgr) GetFromTypecode(typecode int64) (results []*AllType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllType{}).Where("`typecode` = ?", typecode).Find(&results).Error

	return
}

// GetBatchFromTypecode 批量查找
func (obj *_AllTypeMgr) GetBatchFromTypecode(typecodes []int64) (results []*AllType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllType{}).Where("`typecode` IN (?)", typecodes).Find(&results).Error

	return
}

// GetFromProperties 通过properties获取内容
func (obj *_AllTypeMgr) GetFromProperties(properties int64) (results []*AllType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllType{}).Where("`properties` = ?", properties).Find(&results).Error

	return
}

// GetBatchFromProperties 批量查找
func (obj *_AllTypeMgr) GetBatchFromProperties(propertiess []int64) (results []*AllType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllType{}).Where("`properties` IN (?)", propertiess).Find(&results).Error

	return
}

// GetFromAttributes 通过attributes获取内容
func (obj *_AllTypeMgr) GetFromAttributes(attributes int64) (results []*AllType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllType{}).Where("`attributes` = ?", attributes).Find(&results).Error

	return
}

// GetBatchFromAttributes 批量查找
func (obj *_AllTypeMgr) GetBatchFromAttributes(attributess []int64) (results []*AllType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllType{}).Where("`attributes` IN (?)", attributess).Find(&results).Error

	return
}

// GetFromMethods 通过methods获取内容
func (obj *_AllTypeMgr) GetFromMethods(methods int64) (results []*AllType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllType{}).Where("`methods` = ?", methods).Find(&results).Error

	return
}

// GetBatchFromMethods 批量查找
func (obj *_AllTypeMgr) GetBatchFromMethods(methodss []int64) (results []*AllType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllType{}).Where("`methods` IN (?)", methodss).Find(&results).Error

	return
}

// GetFromHiddenmethods 通过hiddenmethods获取内容
func (obj *_AllTypeMgr) GetFromHiddenmethods(hiddenmethods int64) (results []*AllType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllType{}).Where("`hiddenmethods` = ?", hiddenmethods).Find(&results).Error

	return
}

// GetBatchFromHiddenmethods 批量查找
func (obj *_AllTypeMgr) GetBatchFromHiddenmethods(hiddenmethodss []int64) (results []*AllType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllType{}).Where("`hiddenmethods` IN (?)", hiddenmethodss).Find(&results).Error

	return
}

// GetFromSupertypes 通过supertypes获取内容
func (obj *_AllTypeMgr) GetFromSupertypes(supertypes int64) (results []*AllType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllType{}).Where("`supertypes` = ?", supertypes).Find(&results).Error

	return
}

// GetBatchFromSupertypes 批量查找
func (obj *_AllTypeMgr) GetBatchFromSupertypes(supertypess []int64) (results []*AllType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllType{}).Where("`supertypes` IN (?)", supertypess).Find(&results).Error

	return
}

// GetFromSubtypes 通过subtypes获取内容
func (obj *_AllTypeMgr) GetFromSubtypes(subtypes int64) (results []*AllType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllType{}).Where("`subtypes` = ?", subtypes).Find(&results).Error

	return
}

// GetBatchFromSubtypes 批量查找
func (obj *_AllTypeMgr) GetBatchFromSubtypes(subtypess []int64) (results []*AllType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllType{}).Where("`subtypes` IN (?)", subtypess).Find(&results).Error

	return
}

// GetFromExterntype 通过externtype获取内容
func (obj *_AllTypeMgr) GetFromExterntype(externtype int64) (results []*AllType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllType{}).Where("`externtype` = ?", externtype).Find(&results).Error

	return
}

// GetBatchFromExterntype 批量查找
func (obj *_AllTypeMgr) GetBatchFromExterntype(externtypes []int64) (results []*AllType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllType{}).Where("`externtype` IN (?)", externtypes).Find(&results).Error

	return
}

// GetFromExternname 通过externname获取内容
func (obj *_AllTypeMgr) GetFromExternname(externname string) (results []*AllType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllType{}).Where("`externname` = ?", externname).Find(&results).Error

	return
}

// GetBatchFromExternname 批量查找
func (obj *_AllTypeMgr) GetBatchFromExternname(externnames []string) (results []*AllType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllType{}).Where("`externname` IN (?)", externnames).Find(&results).Error

	return
}

// GetFromHelperclassname 通过helperclassname获取内容
func (obj *_AllTypeMgr) GetFromHelperclassname(helperclassname string) (results []*AllType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllType{}).Where("`helperclassname` = ?", helperclassname).Find(&results).Error

	return
}

// GetBatchFromHelperclassname 批量查找
func (obj *_AllTypeMgr) GetBatchFromHelperclassname(helperclassnames []string) (results []*AllType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllType{}).Where("`helperclassname` IN (?)", helperclassnames).Find(&results).Error

	return
}

// GetFromLocalAttrs 通过local_attrs获取内容
func (obj *_AllTypeMgr) GetFromLocalAttrs(localAttrs int64) (results []*AllType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllType{}).Where("`local_attrs` = ?", localAttrs).Find(&results).Error

	return
}

// GetBatchFromLocalAttrs 批量查找
func (obj *_AllTypeMgr) GetBatchFromLocalAttrs(localAttrss []int64) (results []*AllType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllType{}).Where("`local_attrs` IN (?)", localAttrss).Find(&results).Error

	return
}

// GetFromLocalMethods 通过local_methods获取内容
func (obj *_AllTypeMgr) GetFromLocalMethods(localMethods int64) (results []*AllType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllType{}).Where("`local_methods` = ?", localMethods).Find(&results).Error

	return
}

// GetBatchFromLocalMethods 批量查找
func (obj *_AllTypeMgr) GetBatchFromLocalMethods(localMethodss []int64) (results []*AllType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllType{}).Where("`local_methods` IN (?)", localMethodss).Find(&results).Error

	return
}

// GetFromSupertypeid 通过supertypeid获取内容
func (obj *_AllTypeMgr) GetFromSupertypeid(supertypeid int64) (results []*AllType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllType{}).Where("`supertypeid` = ?", supertypeid).Find(&results).Error

	return
}

// GetBatchFromSupertypeid 批量查找
func (obj *_AllTypeMgr) GetBatchFromSupertypeid(supertypeids []int64) (results []*AllType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllType{}).Where("`supertypeid` IN (?)", supertypeids).Find(&results).Error

	return
}

// GetFromTypeName 通过type_name获取内容
func (obj *_AllTypeMgr) GetFromTypeName(typeName string) (results []*AllType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllType{}).Where("`type_name` = ?", typeName).Find(&results).Error

	return
}

// GetBatchFromTypeName 批量查找
func (obj *_AllTypeMgr) GetBatchFromTypeName(typeNames []string) (results []*AllType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllType{}).Where("`type_name` IN (?)", typeNames).Find(&results).Error

	return
}

// GetFromPackageID 通过package_id获取内容
func (obj *_AllTypeMgr) GetFromPackageID(packageID int64) (results []*AllType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllType{}).Where("`package_id` = ?", packageID).Find(&results).Error

	return
}

// GetBatchFromPackageID 批量查找
func (obj *_AllTypeMgr) GetBatchFromPackageID(packageIDs []int64) (results []*AllType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllType{}).Where("`package_id` IN (?)", packageIDs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllTypeMgr) FetchByPrimaryKey(tenantID int64, typeID int64) (result AllType, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllType{}).Where("`tenant_id` = ? AND `type_id` = ?", tenantID, typeID).First(&result).Error

	return
}
