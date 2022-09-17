package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _AllSysParameterMgr struct {
	*_BaseMgr
}

// AllSysParameterMgr open func
func AllSysParameterMgr(db *gorm.DB) *_AllSysParameterMgr {
	if db == nil {
		panic(fmt.Errorf("AllSysParameterMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllSysParameterMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_sys_parameter"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllSysParameterMgr) GetTableName() string {
	return "__all_sys_parameter"
}

// Reset 重置gorm会话
func (obj *_AllSysParameterMgr) Reset() *_AllSysParameterMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllSysParameterMgr) Get() (result AllSysParameter, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSysParameter{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllSysParameterMgr) Gets() (results []*AllSysParameter, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSysParameter{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllSysParameterMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllSysParameter{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithGmtCreate gmt_create获取
func (obj *_AllSysParameterMgr) WithGmtCreate(gmtCreate time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_create"] = gmtCreate })
}

// WithGmtModified gmt_modified获取
func (obj *_AllSysParameterMgr) WithGmtModified(gmtModified time.Time) Option {
	return optionFunc(func(o *options) { o.query["gmt_modified"] = gmtModified })
}

// WithZone zone获取
func (obj *_AllSysParameterMgr) WithZone(zone string) Option {
	return optionFunc(func(o *options) { o.query["zone"] = zone })
}

// WithSvrType svr_type获取
func (obj *_AllSysParameterMgr) WithSvrType(svrType string) Option {
	return optionFunc(func(o *options) { o.query["svr_type"] = svrType })
}

// WithSvrIP svr_ip获取
func (obj *_AllSysParameterMgr) WithSvrIP(svrIP string) Option {
	return optionFunc(func(o *options) { o.query["svr_ip"] = svrIP })
}

// WithSvrPort svr_port获取
func (obj *_AllSysParameterMgr) WithSvrPort(svrPort int64) Option {
	return optionFunc(func(o *options) { o.query["svr_port"] = svrPort })
}

// WithName name获取
func (obj *_AllSysParameterMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithDataType data_type获取
func (obj *_AllSysParameterMgr) WithDataType(dataType string) Option {
	return optionFunc(func(o *options) { o.query["data_type"] = dataType })
}

// WithValue value获取
func (obj *_AllSysParameterMgr) WithValue(value string) Option {
	return optionFunc(func(o *options) { o.query["value"] = value })
}

// WithValueStrict value_strict获取
func (obj *_AllSysParameterMgr) WithValueStrict(valueStrict string) Option {
	return optionFunc(func(o *options) { o.query["value_strict"] = valueStrict })
}

// WithInfo info获取
func (obj *_AllSysParameterMgr) WithInfo(info string) Option {
	return optionFunc(func(o *options) { o.query["info"] = info })
}

// WithNeedReboot need_reboot获取
func (obj *_AllSysParameterMgr) WithNeedReboot(needReboot int64) Option {
	return optionFunc(func(o *options) { o.query["need_reboot"] = needReboot })
}

// WithSection section获取
func (obj *_AllSysParameterMgr) WithSection(section string) Option {
	return optionFunc(func(o *options) { o.query["section"] = section })
}

// WithVisibleLevel visible_level获取
func (obj *_AllSysParameterMgr) WithVisibleLevel(visibleLevel string) Option {
	return optionFunc(func(o *options) { o.query["visible_level"] = visibleLevel })
}

// WithConfigVersion config_version获取
func (obj *_AllSysParameterMgr) WithConfigVersion(configVersion int64) Option {
	return optionFunc(func(o *options) { o.query["config_version"] = configVersion })
}

// WithScope scope获取
func (obj *_AllSysParameterMgr) WithScope(scope string) Option {
	return optionFunc(func(o *options) { o.query["scope"] = scope })
}

// WithSource source获取
func (obj *_AllSysParameterMgr) WithSource(source string) Option {
	return optionFunc(func(o *options) { o.query["source"] = source })
}

// WithEditLevel edit_level获取
func (obj *_AllSysParameterMgr) WithEditLevel(editLevel string) Option {
	return optionFunc(func(o *options) { o.query["edit_level"] = editLevel })
}

// GetByOption 功能选项模式获取
func (obj *_AllSysParameterMgr) GetByOption(opts ...Option) (result AllSysParameter, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllSysParameter{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllSysParameterMgr) GetByOptions(opts ...Option) (results []*AllSysParameter, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllSysParameter{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllSysParameterMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllSysParameter, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllSysParameter{}).Where(options.query)
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
func (obj *_AllSysParameterMgr) GetFromGmtCreate(gmtCreate time.Time) (results []*AllSysParameter, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSysParameter{}).Where("`gmt_create` = ?", gmtCreate).Find(&results).Error

	return
}

// GetBatchFromGmtCreate 批量查找
func (obj *_AllSysParameterMgr) GetBatchFromGmtCreate(gmtCreates []time.Time) (results []*AllSysParameter, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSysParameter{}).Where("`gmt_create` IN (?)", gmtCreates).Find(&results).Error

	return
}

// GetFromGmtModified 通过gmt_modified获取内容
func (obj *_AllSysParameterMgr) GetFromGmtModified(gmtModified time.Time) (results []*AllSysParameter, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSysParameter{}).Where("`gmt_modified` = ?", gmtModified).Find(&results).Error

	return
}

// GetBatchFromGmtModified 批量查找
func (obj *_AllSysParameterMgr) GetBatchFromGmtModified(gmtModifieds []time.Time) (results []*AllSysParameter, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSysParameter{}).Where("`gmt_modified` IN (?)", gmtModifieds).Find(&results).Error

	return
}

// GetFromZone 通过zone获取内容
func (obj *_AllSysParameterMgr) GetFromZone(zone string) (results []*AllSysParameter, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSysParameter{}).Where("`zone` = ?", zone).Find(&results).Error

	return
}

// GetBatchFromZone 批量查找
func (obj *_AllSysParameterMgr) GetBatchFromZone(zones []string) (results []*AllSysParameter, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSysParameter{}).Where("`zone` IN (?)", zones).Find(&results).Error

	return
}

// GetFromSvrType 通过svr_type获取内容
func (obj *_AllSysParameterMgr) GetFromSvrType(svrType string) (results []*AllSysParameter, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSysParameter{}).Where("`svr_type` = ?", svrType).Find(&results).Error

	return
}

// GetBatchFromSvrType 批量查找
func (obj *_AllSysParameterMgr) GetBatchFromSvrType(svrTypes []string) (results []*AllSysParameter, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSysParameter{}).Where("`svr_type` IN (?)", svrTypes).Find(&results).Error

	return
}

// GetFromSvrIP 通过svr_ip获取内容
func (obj *_AllSysParameterMgr) GetFromSvrIP(svrIP string) (results []*AllSysParameter, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSysParameter{}).Where("`svr_ip` = ?", svrIP).Find(&results).Error

	return
}

// GetBatchFromSvrIP 批量查找
func (obj *_AllSysParameterMgr) GetBatchFromSvrIP(svrIPs []string) (results []*AllSysParameter, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSysParameter{}).Where("`svr_ip` IN (?)", svrIPs).Find(&results).Error

	return
}

// GetFromSvrPort 通过svr_port获取内容
func (obj *_AllSysParameterMgr) GetFromSvrPort(svrPort int64) (results []*AllSysParameter, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSysParameter{}).Where("`svr_port` = ?", svrPort).Find(&results).Error

	return
}

// GetBatchFromSvrPort 批量查找
func (obj *_AllSysParameterMgr) GetBatchFromSvrPort(svrPorts []int64) (results []*AllSysParameter, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSysParameter{}).Where("`svr_port` IN (?)", svrPorts).Find(&results).Error

	return
}

// GetFromName 通过name获取内容
func (obj *_AllSysParameterMgr) GetFromName(name string) (results []*AllSysParameter, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSysParameter{}).Where("`name` = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量查找
func (obj *_AllSysParameterMgr) GetBatchFromName(names []string) (results []*AllSysParameter, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSysParameter{}).Where("`name` IN (?)", names).Find(&results).Error

	return
}

// GetFromDataType 通过data_type获取内容
func (obj *_AllSysParameterMgr) GetFromDataType(dataType string) (results []*AllSysParameter, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSysParameter{}).Where("`data_type` = ?", dataType).Find(&results).Error

	return
}

// GetBatchFromDataType 批量查找
func (obj *_AllSysParameterMgr) GetBatchFromDataType(dataTypes []string) (results []*AllSysParameter, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSysParameter{}).Where("`data_type` IN (?)", dataTypes).Find(&results).Error

	return
}

// GetFromValue 通过value获取内容
func (obj *_AllSysParameterMgr) GetFromValue(value string) (results []*AllSysParameter, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSysParameter{}).Where("`value` = ?", value).Find(&results).Error

	return
}

// GetBatchFromValue 批量查找
func (obj *_AllSysParameterMgr) GetBatchFromValue(values []string) (results []*AllSysParameter, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSysParameter{}).Where("`value` IN (?)", values).Find(&results).Error

	return
}

// GetFromValueStrict 通过value_strict获取内容
func (obj *_AllSysParameterMgr) GetFromValueStrict(valueStrict string) (results []*AllSysParameter, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSysParameter{}).Where("`value_strict` = ?", valueStrict).Find(&results).Error

	return
}

// GetBatchFromValueStrict 批量查找
func (obj *_AllSysParameterMgr) GetBatchFromValueStrict(valueStricts []string) (results []*AllSysParameter, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSysParameter{}).Where("`value_strict` IN (?)", valueStricts).Find(&results).Error

	return
}

// GetFromInfo 通过info获取内容
func (obj *_AllSysParameterMgr) GetFromInfo(info string) (results []*AllSysParameter, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSysParameter{}).Where("`info` = ?", info).Find(&results).Error

	return
}

// GetBatchFromInfo 批量查找
func (obj *_AllSysParameterMgr) GetBatchFromInfo(infos []string) (results []*AllSysParameter, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSysParameter{}).Where("`info` IN (?)", infos).Find(&results).Error

	return
}

// GetFromNeedReboot 通过need_reboot获取内容
func (obj *_AllSysParameterMgr) GetFromNeedReboot(needReboot int64) (results []*AllSysParameter, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSysParameter{}).Where("`need_reboot` = ?", needReboot).Find(&results).Error

	return
}

// GetBatchFromNeedReboot 批量查找
func (obj *_AllSysParameterMgr) GetBatchFromNeedReboot(needReboots []int64) (results []*AllSysParameter, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSysParameter{}).Where("`need_reboot` IN (?)", needReboots).Find(&results).Error

	return
}

// GetFromSection 通过section获取内容
func (obj *_AllSysParameterMgr) GetFromSection(section string) (results []*AllSysParameter, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSysParameter{}).Where("`section` = ?", section).Find(&results).Error

	return
}

// GetBatchFromSection 批量查找
func (obj *_AllSysParameterMgr) GetBatchFromSection(sections []string) (results []*AllSysParameter, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSysParameter{}).Where("`section` IN (?)", sections).Find(&results).Error

	return
}

// GetFromVisibleLevel 通过visible_level获取内容
func (obj *_AllSysParameterMgr) GetFromVisibleLevel(visibleLevel string) (results []*AllSysParameter, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSysParameter{}).Where("`visible_level` = ?", visibleLevel).Find(&results).Error

	return
}

// GetBatchFromVisibleLevel 批量查找
func (obj *_AllSysParameterMgr) GetBatchFromVisibleLevel(visibleLevels []string) (results []*AllSysParameter, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSysParameter{}).Where("`visible_level` IN (?)", visibleLevels).Find(&results).Error

	return
}

// GetFromConfigVersion 通过config_version获取内容
func (obj *_AllSysParameterMgr) GetFromConfigVersion(configVersion int64) (results []*AllSysParameter, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSysParameter{}).Where("`config_version` = ?", configVersion).Find(&results).Error

	return
}

// GetBatchFromConfigVersion 批量查找
func (obj *_AllSysParameterMgr) GetBatchFromConfigVersion(configVersions []int64) (results []*AllSysParameter, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSysParameter{}).Where("`config_version` IN (?)", configVersions).Find(&results).Error

	return
}

// GetFromScope 通过scope获取内容
func (obj *_AllSysParameterMgr) GetFromScope(scope string) (results []*AllSysParameter, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSysParameter{}).Where("`scope` = ?", scope).Find(&results).Error

	return
}

// GetBatchFromScope 批量查找
func (obj *_AllSysParameterMgr) GetBatchFromScope(scopes []string) (results []*AllSysParameter, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSysParameter{}).Where("`scope` IN (?)", scopes).Find(&results).Error

	return
}

// GetFromSource 通过source获取内容
func (obj *_AllSysParameterMgr) GetFromSource(source string) (results []*AllSysParameter, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSysParameter{}).Where("`source` = ?", source).Find(&results).Error

	return
}

// GetBatchFromSource 批量查找
func (obj *_AllSysParameterMgr) GetBatchFromSource(sources []string) (results []*AllSysParameter, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSysParameter{}).Where("`source` IN (?)", sources).Find(&results).Error

	return
}

// GetFromEditLevel 通过edit_level获取内容
func (obj *_AllSysParameterMgr) GetFromEditLevel(editLevel string) (results []*AllSysParameter, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSysParameter{}).Where("`edit_level` = ?", editLevel).Find(&results).Error

	return
}

// GetBatchFromEditLevel 批量查找
func (obj *_AllSysParameterMgr) GetBatchFromEditLevel(editLevels []string) (results []*AllSysParameter, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSysParameter{}).Where("`edit_level` IN (?)", editLevels).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllSysParameterMgr) FetchByPrimaryKey(zone string, svrType string, svrIP string, svrPort int64, name string) (result AllSysParameter, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllSysParameter{}).Where("`zone` = ? AND `svr_type` = ? AND `svr_ip` = ? AND `svr_port` = ? AND `name` = ?", zone, svrType, svrIP, svrPort, name).First(&result).Error

	return
}
