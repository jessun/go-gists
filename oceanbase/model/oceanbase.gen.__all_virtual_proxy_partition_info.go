package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _AllVirtualProxyPartitionInfoMgr struct {
	*_BaseMgr
}

// AllVirtualProxyPartitionInfoMgr open func
func AllVirtualProxyPartitionInfoMgr(db *gorm.DB) *_AllVirtualProxyPartitionInfoMgr {
	if db == nil {
		panic(fmt.Errorf("AllVirtualProxyPartitionInfoMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AllVirtualProxyPartitionInfoMgr{_BaseMgr: &_BaseMgr{DB: db.Table("__all_virtual_proxy_partition_info"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AllVirtualProxyPartitionInfoMgr) GetTableName() string {
	return "__all_virtual_proxy_partition_info"
}

// Reset 重置gorm会话
func (obj *_AllVirtualProxyPartitionInfoMgr) Reset() *_AllVirtualProxyPartitionInfoMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AllVirtualProxyPartitionInfoMgr) Get() (result AllVirtualProxyPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxyPartitionInfo{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AllVirtualProxyPartitionInfoMgr) Gets() (results []*AllVirtualProxyPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxyPartitionInfo{}).Find(&results).Error

	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AllVirtualProxyPartitionInfoMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(AllVirtualProxyPartitionInfo{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTableID table_id获取
func (obj *_AllVirtualProxyPartitionInfoMgr) WithTableID(tableID int64) Option {
	return optionFunc(func(o *options) { o.query["table_id"] = tableID })
}

// WithTenantID tenant_id获取
func (obj *_AllVirtualProxyPartitionInfoMgr) WithTenantID(tenantID int64) Option {
	return optionFunc(func(o *options) { o.query["tenant_id"] = tenantID })
}

// WithPartLevel part_level获取
func (obj *_AllVirtualProxyPartitionInfoMgr) WithPartLevel(partLevel int64) Option {
	return optionFunc(func(o *options) { o.query["part_level"] = partLevel })
}

// WithAllPartNum all_part_num获取
func (obj *_AllVirtualProxyPartitionInfoMgr) WithAllPartNum(allPartNum int64) Option {
	return optionFunc(func(o *options) { o.query["all_part_num"] = allPartNum })
}

// WithTemplateNum template_num获取
func (obj *_AllVirtualProxyPartitionInfoMgr) WithTemplateNum(templateNum int64) Option {
	return optionFunc(func(o *options) { o.query["template_num"] = templateNum })
}

// WithPartIDRuleVer part_id_rule_ver获取
func (obj *_AllVirtualProxyPartitionInfoMgr) WithPartIDRuleVer(partIDRuleVer int64) Option {
	return optionFunc(func(o *options) { o.query["part_id_rule_ver"] = partIDRuleVer })
}

// WithPartType part_type获取
func (obj *_AllVirtualProxyPartitionInfoMgr) WithPartType(partType int64) Option {
	return optionFunc(func(o *options) { o.query["part_type"] = partType })
}

// WithPartNum part_num获取
func (obj *_AllVirtualProxyPartitionInfoMgr) WithPartNum(partNum int64) Option {
	return optionFunc(func(o *options) { o.query["part_num"] = partNum })
}

// WithIsColumnType is_column_type获取
func (obj *_AllVirtualProxyPartitionInfoMgr) WithIsColumnType(isColumnType int8) Option {
	return optionFunc(func(o *options) { o.query["is_column_type"] = isColumnType })
}

// WithPartSpace part_space获取
func (obj *_AllVirtualProxyPartitionInfoMgr) WithPartSpace(partSpace int64) Option {
	return optionFunc(func(o *options) { o.query["part_space"] = partSpace })
}

// WithPartExpr part_expr获取
func (obj *_AllVirtualProxyPartitionInfoMgr) WithPartExpr(partExpr string) Option {
	return optionFunc(func(o *options) { o.query["part_expr"] = partExpr })
}

// WithPartExprBin part_expr_bin获取
func (obj *_AllVirtualProxyPartitionInfoMgr) WithPartExprBin(partExprBin string) Option {
	return optionFunc(func(o *options) { o.query["part_expr_bin"] = partExprBin })
}

// WithPartRangeType part_range_type获取
func (obj *_AllVirtualProxyPartitionInfoMgr) WithPartRangeType(partRangeType string) Option {
	return optionFunc(func(o *options) { o.query["part_range_type"] = partRangeType })
}

// WithPartInterval part_interval获取
func (obj *_AllVirtualProxyPartitionInfoMgr) WithPartInterval(partInterval string) Option {
	return optionFunc(func(o *options) { o.query["part_interval"] = partInterval })
}

// WithPartIntervalBin part_interval_bin获取
func (obj *_AllVirtualProxyPartitionInfoMgr) WithPartIntervalBin(partIntervalBin string) Option {
	return optionFunc(func(o *options) { o.query["part_interval_bin"] = partIntervalBin })
}

// WithIntervalStart interval_start获取
func (obj *_AllVirtualProxyPartitionInfoMgr) WithIntervalStart(intervalStart string) Option {
	return optionFunc(func(o *options) { o.query["interval_start"] = intervalStart })
}

// WithIntervalStartBin interval_start_bin获取
func (obj *_AllVirtualProxyPartitionInfoMgr) WithIntervalStartBin(intervalStartBin string) Option {
	return optionFunc(func(o *options) { o.query["interval_start_bin"] = intervalStartBin })
}

// WithSubPartType sub_part_type获取
func (obj *_AllVirtualProxyPartitionInfoMgr) WithSubPartType(subPartType int64) Option {
	return optionFunc(func(o *options) { o.query["sub_part_type"] = subPartType })
}

// WithSubPartNum sub_part_num获取
func (obj *_AllVirtualProxyPartitionInfoMgr) WithSubPartNum(subPartNum int64) Option {
	return optionFunc(func(o *options) { o.query["sub_part_num"] = subPartNum })
}

// WithIsSubColumnType is_sub_column_type获取
func (obj *_AllVirtualProxyPartitionInfoMgr) WithIsSubColumnType(isSubColumnType int8) Option {
	return optionFunc(func(o *options) { o.query["is_sub_column_type"] = isSubColumnType })
}

// WithSubPartSpace sub_part_space获取
func (obj *_AllVirtualProxyPartitionInfoMgr) WithSubPartSpace(subPartSpace int64) Option {
	return optionFunc(func(o *options) { o.query["sub_part_space"] = subPartSpace })
}

// WithSubPartExpr sub_part_expr获取
func (obj *_AllVirtualProxyPartitionInfoMgr) WithSubPartExpr(subPartExpr string) Option {
	return optionFunc(func(o *options) { o.query["sub_part_expr"] = subPartExpr })
}

// WithSubPartExprBin sub_part_expr_bin获取
func (obj *_AllVirtualProxyPartitionInfoMgr) WithSubPartExprBin(subPartExprBin string) Option {
	return optionFunc(func(o *options) { o.query["sub_part_expr_bin"] = subPartExprBin })
}

// WithSubPartRangeType sub_part_range_type获取
func (obj *_AllVirtualProxyPartitionInfoMgr) WithSubPartRangeType(subPartRangeType string) Option {
	return optionFunc(func(o *options) { o.query["sub_part_range_type"] = subPartRangeType })
}

// WithDefSubPartInterval def_sub_part_interval获取
func (obj *_AllVirtualProxyPartitionInfoMgr) WithDefSubPartInterval(defSubPartInterval string) Option {
	return optionFunc(func(o *options) { o.query["def_sub_part_interval"] = defSubPartInterval })
}

// WithDefSubPartIntervalBin def_sub_part_interval_bin获取
func (obj *_AllVirtualProxyPartitionInfoMgr) WithDefSubPartIntervalBin(defSubPartIntervalBin string) Option {
	return optionFunc(func(o *options) { o.query["def_sub_part_interval_bin"] = defSubPartIntervalBin })
}

// WithDefSubIntervalStart def_sub_interval_start获取
func (obj *_AllVirtualProxyPartitionInfoMgr) WithDefSubIntervalStart(defSubIntervalStart string) Option {
	return optionFunc(func(o *options) { o.query["def_sub_interval_start"] = defSubIntervalStart })
}

// WithDefSubIntervalStartBin def_sub_interval_start_bin获取
func (obj *_AllVirtualProxyPartitionInfoMgr) WithDefSubIntervalStartBin(defSubIntervalStartBin string) Option {
	return optionFunc(func(o *options) { o.query["def_sub_interval_start_bin"] = defSubIntervalStartBin })
}

// WithPartKeyNum part_key_num获取
func (obj *_AllVirtualProxyPartitionInfoMgr) WithPartKeyNum(partKeyNum int64) Option {
	return optionFunc(func(o *options) { o.query["part_key_num"] = partKeyNum })
}

// WithPartKeyName part_key_name获取
func (obj *_AllVirtualProxyPartitionInfoMgr) WithPartKeyName(partKeyName string) Option {
	return optionFunc(func(o *options) { o.query["part_key_name"] = partKeyName })
}

// WithPartKeyType part_key_type获取
func (obj *_AllVirtualProxyPartitionInfoMgr) WithPartKeyType(partKeyType int64) Option {
	return optionFunc(func(o *options) { o.query["part_key_type"] = partKeyType })
}

// WithPartKeyIDx part_key_idx获取
func (obj *_AllVirtualProxyPartitionInfoMgr) WithPartKeyIDx(partKeyIDx int64) Option {
	return optionFunc(func(o *options) { o.query["part_key_idx"] = partKeyIDx })
}

// WithPartKeyLevel part_key_level获取
func (obj *_AllVirtualProxyPartitionInfoMgr) WithPartKeyLevel(partKeyLevel int64) Option {
	return optionFunc(func(o *options) { o.query["part_key_level"] = partKeyLevel })
}

// WithPartKeyExtra part_key_extra获取
func (obj *_AllVirtualProxyPartitionInfoMgr) WithPartKeyExtra(partKeyExtra string) Option {
	return optionFunc(func(o *options) { o.query["part_key_extra"] = partKeyExtra })
}

// WithSpare1 spare1获取
func (obj *_AllVirtualProxyPartitionInfoMgr) WithSpare1(spare1 int64) Option {
	return optionFunc(func(o *options) { o.query["spare1"] = spare1 })
}

// WithSpare2 spare2获取
func (obj *_AllVirtualProxyPartitionInfoMgr) WithSpare2(spare2 int64) Option {
	return optionFunc(func(o *options) { o.query["spare2"] = spare2 })
}

// WithSpare3 spare3获取
func (obj *_AllVirtualProxyPartitionInfoMgr) WithSpare3(spare3 int64) Option {
	return optionFunc(func(o *options) { o.query["spare3"] = spare3 })
}

// WithSpare4 spare4获取
func (obj *_AllVirtualProxyPartitionInfoMgr) WithSpare4(spare4 string) Option {
	return optionFunc(func(o *options) { o.query["spare4"] = spare4 })
}

// WithSpare5 spare5获取
func (obj *_AllVirtualProxyPartitionInfoMgr) WithSpare5(spare5 string) Option {
	return optionFunc(func(o *options) { o.query["spare5"] = spare5 })
}

// WithSpare6 spare6获取
func (obj *_AllVirtualProxyPartitionInfoMgr) WithSpare6(spare6 string) Option {
	return optionFunc(func(o *options) { o.query["spare6"] = spare6 })
}

// GetByOption 功能选项模式获取
func (obj *_AllVirtualProxyPartitionInfoMgr) GetByOption(opts ...Option) (result AllVirtualProxyPartitionInfo, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxyPartitionInfo{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AllVirtualProxyPartitionInfoMgr) GetByOptions(opts ...Option) (results []*AllVirtualProxyPartitionInfo, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxyPartitionInfo{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AllVirtualProxyPartitionInfoMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]AllVirtualProxyPartitionInfo, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(AllVirtualProxyPartitionInfo{}).Where(options.query)
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

// GetFromTableID 通过table_id获取内容
func (obj *_AllVirtualProxyPartitionInfoMgr) GetFromTableID(tableID int64) (result AllVirtualProxyPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxyPartitionInfo{}).Where("`table_id` = ?", tableID).First(&result).Error

	return
}

// GetBatchFromTableID 批量查找
func (obj *_AllVirtualProxyPartitionInfoMgr) GetBatchFromTableID(tableIDs []int64) (results []*AllVirtualProxyPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxyPartitionInfo{}).Where("`table_id` IN (?)", tableIDs).Find(&results).Error

	return
}

// GetFromTenantID 通过tenant_id获取内容
func (obj *_AllVirtualProxyPartitionInfoMgr) GetFromTenantID(tenantID int64) (results []*AllVirtualProxyPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxyPartitionInfo{}).Where("`tenant_id` = ?", tenantID).Find(&results).Error

	return
}

// GetBatchFromTenantID 批量查找
func (obj *_AllVirtualProxyPartitionInfoMgr) GetBatchFromTenantID(tenantIDs []int64) (results []*AllVirtualProxyPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxyPartitionInfo{}).Where("`tenant_id` IN (?)", tenantIDs).Find(&results).Error

	return
}

// GetFromPartLevel 通过part_level获取内容
func (obj *_AllVirtualProxyPartitionInfoMgr) GetFromPartLevel(partLevel int64) (results []*AllVirtualProxyPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxyPartitionInfo{}).Where("`part_level` = ?", partLevel).Find(&results).Error

	return
}

// GetBatchFromPartLevel 批量查找
func (obj *_AllVirtualProxyPartitionInfoMgr) GetBatchFromPartLevel(partLevels []int64) (results []*AllVirtualProxyPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxyPartitionInfo{}).Where("`part_level` IN (?)", partLevels).Find(&results).Error

	return
}

// GetFromAllPartNum 通过all_part_num获取内容
func (obj *_AllVirtualProxyPartitionInfoMgr) GetFromAllPartNum(allPartNum int64) (results []*AllVirtualProxyPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxyPartitionInfo{}).Where("`all_part_num` = ?", allPartNum).Find(&results).Error

	return
}

// GetBatchFromAllPartNum 批量查找
func (obj *_AllVirtualProxyPartitionInfoMgr) GetBatchFromAllPartNum(allPartNums []int64) (results []*AllVirtualProxyPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxyPartitionInfo{}).Where("`all_part_num` IN (?)", allPartNums).Find(&results).Error

	return
}

// GetFromTemplateNum 通过template_num获取内容
func (obj *_AllVirtualProxyPartitionInfoMgr) GetFromTemplateNum(templateNum int64) (results []*AllVirtualProxyPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxyPartitionInfo{}).Where("`template_num` = ?", templateNum).Find(&results).Error

	return
}

// GetBatchFromTemplateNum 批量查找
func (obj *_AllVirtualProxyPartitionInfoMgr) GetBatchFromTemplateNum(templateNums []int64) (results []*AllVirtualProxyPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxyPartitionInfo{}).Where("`template_num` IN (?)", templateNums).Find(&results).Error

	return
}

// GetFromPartIDRuleVer 通过part_id_rule_ver获取内容
func (obj *_AllVirtualProxyPartitionInfoMgr) GetFromPartIDRuleVer(partIDRuleVer int64) (results []*AllVirtualProxyPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxyPartitionInfo{}).Where("`part_id_rule_ver` = ?", partIDRuleVer).Find(&results).Error

	return
}

// GetBatchFromPartIDRuleVer 批量查找
func (obj *_AllVirtualProxyPartitionInfoMgr) GetBatchFromPartIDRuleVer(partIDRuleVers []int64) (results []*AllVirtualProxyPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxyPartitionInfo{}).Where("`part_id_rule_ver` IN (?)", partIDRuleVers).Find(&results).Error

	return
}

// GetFromPartType 通过part_type获取内容
func (obj *_AllVirtualProxyPartitionInfoMgr) GetFromPartType(partType int64) (results []*AllVirtualProxyPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxyPartitionInfo{}).Where("`part_type` = ?", partType).Find(&results).Error

	return
}

// GetBatchFromPartType 批量查找
func (obj *_AllVirtualProxyPartitionInfoMgr) GetBatchFromPartType(partTypes []int64) (results []*AllVirtualProxyPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxyPartitionInfo{}).Where("`part_type` IN (?)", partTypes).Find(&results).Error

	return
}

// GetFromPartNum 通过part_num获取内容
func (obj *_AllVirtualProxyPartitionInfoMgr) GetFromPartNum(partNum int64) (results []*AllVirtualProxyPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxyPartitionInfo{}).Where("`part_num` = ?", partNum).Find(&results).Error

	return
}

// GetBatchFromPartNum 批量查找
func (obj *_AllVirtualProxyPartitionInfoMgr) GetBatchFromPartNum(partNums []int64) (results []*AllVirtualProxyPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxyPartitionInfo{}).Where("`part_num` IN (?)", partNums).Find(&results).Error

	return
}

// GetFromIsColumnType 通过is_column_type获取内容
func (obj *_AllVirtualProxyPartitionInfoMgr) GetFromIsColumnType(isColumnType int8) (results []*AllVirtualProxyPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxyPartitionInfo{}).Where("`is_column_type` = ?", isColumnType).Find(&results).Error

	return
}

// GetBatchFromIsColumnType 批量查找
func (obj *_AllVirtualProxyPartitionInfoMgr) GetBatchFromIsColumnType(isColumnTypes []int8) (results []*AllVirtualProxyPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxyPartitionInfo{}).Where("`is_column_type` IN (?)", isColumnTypes).Find(&results).Error

	return
}

// GetFromPartSpace 通过part_space获取内容
func (obj *_AllVirtualProxyPartitionInfoMgr) GetFromPartSpace(partSpace int64) (results []*AllVirtualProxyPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxyPartitionInfo{}).Where("`part_space` = ?", partSpace).Find(&results).Error

	return
}

// GetBatchFromPartSpace 批量查找
func (obj *_AllVirtualProxyPartitionInfoMgr) GetBatchFromPartSpace(partSpaces []int64) (results []*AllVirtualProxyPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxyPartitionInfo{}).Where("`part_space` IN (?)", partSpaces).Find(&results).Error

	return
}

// GetFromPartExpr 通过part_expr获取内容
func (obj *_AllVirtualProxyPartitionInfoMgr) GetFromPartExpr(partExpr string) (results []*AllVirtualProxyPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxyPartitionInfo{}).Where("`part_expr` = ?", partExpr).Find(&results).Error

	return
}

// GetBatchFromPartExpr 批量查找
func (obj *_AllVirtualProxyPartitionInfoMgr) GetBatchFromPartExpr(partExprs []string) (results []*AllVirtualProxyPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxyPartitionInfo{}).Where("`part_expr` IN (?)", partExprs).Find(&results).Error

	return
}

// GetFromPartExprBin 通过part_expr_bin获取内容
func (obj *_AllVirtualProxyPartitionInfoMgr) GetFromPartExprBin(partExprBin string) (results []*AllVirtualProxyPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxyPartitionInfo{}).Where("`part_expr_bin` = ?", partExprBin).Find(&results).Error

	return
}

// GetBatchFromPartExprBin 批量查找
func (obj *_AllVirtualProxyPartitionInfoMgr) GetBatchFromPartExprBin(partExprBins []string) (results []*AllVirtualProxyPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxyPartitionInfo{}).Where("`part_expr_bin` IN (?)", partExprBins).Find(&results).Error

	return
}

// GetFromPartRangeType 通过part_range_type获取内容
func (obj *_AllVirtualProxyPartitionInfoMgr) GetFromPartRangeType(partRangeType string) (results []*AllVirtualProxyPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxyPartitionInfo{}).Where("`part_range_type` = ?", partRangeType).Find(&results).Error

	return
}

// GetBatchFromPartRangeType 批量查找
func (obj *_AllVirtualProxyPartitionInfoMgr) GetBatchFromPartRangeType(partRangeTypes []string) (results []*AllVirtualProxyPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxyPartitionInfo{}).Where("`part_range_type` IN (?)", partRangeTypes).Find(&results).Error

	return
}

// GetFromPartInterval 通过part_interval获取内容
func (obj *_AllVirtualProxyPartitionInfoMgr) GetFromPartInterval(partInterval string) (results []*AllVirtualProxyPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxyPartitionInfo{}).Where("`part_interval` = ?", partInterval).Find(&results).Error

	return
}

// GetBatchFromPartInterval 批量查找
func (obj *_AllVirtualProxyPartitionInfoMgr) GetBatchFromPartInterval(partIntervals []string) (results []*AllVirtualProxyPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxyPartitionInfo{}).Where("`part_interval` IN (?)", partIntervals).Find(&results).Error

	return
}

// GetFromPartIntervalBin 通过part_interval_bin获取内容
func (obj *_AllVirtualProxyPartitionInfoMgr) GetFromPartIntervalBin(partIntervalBin string) (results []*AllVirtualProxyPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxyPartitionInfo{}).Where("`part_interval_bin` = ?", partIntervalBin).Find(&results).Error

	return
}

// GetBatchFromPartIntervalBin 批量查找
func (obj *_AllVirtualProxyPartitionInfoMgr) GetBatchFromPartIntervalBin(partIntervalBins []string) (results []*AllVirtualProxyPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxyPartitionInfo{}).Where("`part_interval_bin` IN (?)", partIntervalBins).Find(&results).Error

	return
}

// GetFromIntervalStart 通过interval_start获取内容
func (obj *_AllVirtualProxyPartitionInfoMgr) GetFromIntervalStart(intervalStart string) (results []*AllVirtualProxyPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxyPartitionInfo{}).Where("`interval_start` = ?", intervalStart).Find(&results).Error

	return
}

// GetBatchFromIntervalStart 批量查找
func (obj *_AllVirtualProxyPartitionInfoMgr) GetBatchFromIntervalStart(intervalStarts []string) (results []*AllVirtualProxyPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxyPartitionInfo{}).Where("`interval_start` IN (?)", intervalStarts).Find(&results).Error

	return
}

// GetFromIntervalStartBin 通过interval_start_bin获取内容
func (obj *_AllVirtualProxyPartitionInfoMgr) GetFromIntervalStartBin(intervalStartBin string) (results []*AllVirtualProxyPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxyPartitionInfo{}).Where("`interval_start_bin` = ?", intervalStartBin).Find(&results).Error

	return
}

// GetBatchFromIntervalStartBin 批量查找
func (obj *_AllVirtualProxyPartitionInfoMgr) GetBatchFromIntervalStartBin(intervalStartBins []string) (results []*AllVirtualProxyPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxyPartitionInfo{}).Where("`interval_start_bin` IN (?)", intervalStartBins).Find(&results).Error

	return
}

// GetFromSubPartType 通过sub_part_type获取内容
func (obj *_AllVirtualProxyPartitionInfoMgr) GetFromSubPartType(subPartType int64) (results []*AllVirtualProxyPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxyPartitionInfo{}).Where("`sub_part_type` = ?", subPartType).Find(&results).Error

	return
}

// GetBatchFromSubPartType 批量查找
func (obj *_AllVirtualProxyPartitionInfoMgr) GetBatchFromSubPartType(subPartTypes []int64) (results []*AllVirtualProxyPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxyPartitionInfo{}).Where("`sub_part_type` IN (?)", subPartTypes).Find(&results).Error

	return
}

// GetFromSubPartNum 通过sub_part_num获取内容
func (obj *_AllVirtualProxyPartitionInfoMgr) GetFromSubPartNum(subPartNum int64) (results []*AllVirtualProxyPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxyPartitionInfo{}).Where("`sub_part_num` = ?", subPartNum).Find(&results).Error

	return
}

// GetBatchFromSubPartNum 批量查找
func (obj *_AllVirtualProxyPartitionInfoMgr) GetBatchFromSubPartNum(subPartNums []int64) (results []*AllVirtualProxyPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxyPartitionInfo{}).Where("`sub_part_num` IN (?)", subPartNums).Find(&results).Error

	return
}

// GetFromIsSubColumnType 通过is_sub_column_type获取内容
func (obj *_AllVirtualProxyPartitionInfoMgr) GetFromIsSubColumnType(isSubColumnType int8) (results []*AllVirtualProxyPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxyPartitionInfo{}).Where("`is_sub_column_type` = ?", isSubColumnType).Find(&results).Error

	return
}

// GetBatchFromIsSubColumnType 批量查找
func (obj *_AllVirtualProxyPartitionInfoMgr) GetBatchFromIsSubColumnType(isSubColumnTypes []int8) (results []*AllVirtualProxyPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxyPartitionInfo{}).Where("`is_sub_column_type` IN (?)", isSubColumnTypes).Find(&results).Error

	return
}

// GetFromSubPartSpace 通过sub_part_space获取内容
func (obj *_AllVirtualProxyPartitionInfoMgr) GetFromSubPartSpace(subPartSpace int64) (results []*AllVirtualProxyPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxyPartitionInfo{}).Where("`sub_part_space` = ?", subPartSpace).Find(&results).Error

	return
}

// GetBatchFromSubPartSpace 批量查找
func (obj *_AllVirtualProxyPartitionInfoMgr) GetBatchFromSubPartSpace(subPartSpaces []int64) (results []*AllVirtualProxyPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxyPartitionInfo{}).Where("`sub_part_space` IN (?)", subPartSpaces).Find(&results).Error

	return
}

// GetFromSubPartExpr 通过sub_part_expr获取内容
func (obj *_AllVirtualProxyPartitionInfoMgr) GetFromSubPartExpr(subPartExpr string) (results []*AllVirtualProxyPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxyPartitionInfo{}).Where("`sub_part_expr` = ?", subPartExpr).Find(&results).Error

	return
}

// GetBatchFromSubPartExpr 批量查找
func (obj *_AllVirtualProxyPartitionInfoMgr) GetBatchFromSubPartExpr(subPartExprs []string) (results []*AllVirtualProxyPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxyPartitionInfo{}).Where("`sub_part_expr` IN (?)", subPartExprs).Find(&results).Error

	return
}

// GetFromSubPartExprBin 通过sub_part_expr_bin获取内容
func (obj *_AllVirtualProxyPartitionInfoMgr) GetFromSubPartExprBin(subPartExprBin string) (results []*AllVirtualProxyPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxyPartitionInfo{}).Where("`sub_part_expr_bin` = ?", subPartExprBin).Find(&results).Error

	return
}

// GetBatchFromSubPartExprBin 批量查找
func (obj *_AllVirtualProxyPartitionInfoMgr) GetBatchFromSubPartExprBin(subPartExprBins []string) (results []*AllVirtualProxyPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxyPartitionInfo{}).Where("`sub_part_expr_bin` IN (?)", subPartExprBins).Find(&results).Error

	return
}

// GetFromSubPartRangeType 通过sub_part_range_type获取内容
func (obj *_AllVirtualProxyPartitionInfoMgr) GetFromSubPartRangeType(subPartRangeType string) (results []*AllVirtualProxyPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxyPartitionInfo{}).Where("`sub_part_range_type` = ?", subPartRangeType).Find(&results).Error

	return
}

// GetBatchFromSubPartRangeType 批量查找
func (obj *_AllVirtualProxyPartitionInfoMgr) GetBatchFromSubPartRangeType(subPartRangeTypes []string) (results []*AllVirtualProxyPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxyPartitionInfo{}).Where("`sub_part_range_type` IN (?)", subPartRangeTypes).Find(&results).Error

	return
}

// GetFromDefSubPartInterval 通过def_sub_part_interval获取内容
func (obj *_AllVirtualProxyPartitionInfoMgr) GetFromDefSubPartInterval(defSubPartInterval string) (results []*AllVirtualProxyPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxyPartitionInfo{}).Where("`def_sub_part_interval` = ?", defSubPartInterval).Find(&results).Error

	return
}

// GetBatchFromDefSubPartInterval 批量查找
func (obj *_AllVirtualProxyPartitionInfoMgr) GetBatchFromDefSubPartInterval(defSubPartIntervals []string) (results []*AllVirtualProxyPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxyPartitionInfo{}).Where("`def_sub_part_interval` IN (?)", defSubPartIntervals).Find(&results).Error

	return
}

// GetFromDefSubPartIntervalBin 通过def_sub_part_interval_bin获取内容
func (obj *_AllVirtualProxyPartitionInfoMgr) GetFromDefSubPartIntervalBin(defSubPartIntervalBin string) (results []*AllVirtualProxyPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxyPartitionInfo{}).Where("`def_sub_part_interval_bin` = ?", defSubPartIntervalBin).Find(&results).Error

	return
}

// GetBatchFromDefSubPartIntervalBin 批量查找
func (obj *_AllVirtualProxyPartitionInfoMgr) GetBatchFromDefSubPartIntervalBin(defSubPartIntervalBins []string) (results []*AllVirtualProxyPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxyPartitionInfo{}).Where("`def_sub_part_interval_bin` IN (?)", defSubPartIntervalBins).Find(&results).Error

	return
}

// GetFromDefSubIntervalStart 通过def_sub_interval_start获取内容
func (obj *_AllVirtualProxyPartitionInfoMgr) GetFromDefSubIntervalStart(defSubIntervalStart string) (results []*AllVirtualProxyPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxyPartitionInfo{}).Where("`def_sub_interval_start` = ?", defSubIntervalStart).Find(&results).Error

	return
}

// GetBatchFromDefSubIntervalStart 批量查找
func (obj *_AllVirtualProxyPartitionInfoMgr) GetBatchFromDefSubIntervalStart(defSubIntervalStarts []string) (results []*AllVirtualProxyPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxyPartitionInfo{}).Where("`def_sub_interval_start` IN (?)", defSubIntervalStarts).Find(&results).Error

	return
}

// GetFromDefSubIntervalStartBin 通过def_sub_interval_start_bin获取内容
func (obj *_AllVirtualProxyPartitionInfoMgr) GetFromDefSubIntervalStartBin(defSubIntervalStartBin string) (results []*AllVirtualProxyPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxyPartitionInfo{}).Where("`def_sub_interval_start_bin` = ?", defSubIntervalStartBin).Find(&results).Error

	return
}

// GetBatchFromDefSubIntervalStartBin 批量查找
func (obj *_AllVirtualProxyPartitionInfoMgr) GetBatchFromDefSubIntervalStartBin(defSubIntervalStartBins []string) (results []*AllVirtualProxyPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxyPartitionInfo{}).Where("`def_sub_interval_start_bin` IN (?)", defSubIntervalStartBins).Find(&results).Error

	return
}

// GetFromPartKeyNum 通过part_key_num获取内容
func (obj *_AllVirtualProxyPartitionInfoMgr) GetFromPartKeyNum(partKeyNum int64) (results []*AllVirtualProxyPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxyPartitionInfo{}).Where("`part_key_num` = ?", partKeyNum).Find(&results).Error

	return
}

// GetBatchFromPartKeyNum 批量查找
func (obj *_AllVirtualProxyPartitionInfoMgr) GetBatchFromPartKeyNum(partKeyNums []int64) (results []*AllVirtualProxyPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxyPartitionInfo{}).Where("`part_key_num` IN (?)", partKeyNums).Find(&results).Error

	return
}

// GetFromPartKeyName 通过part_key_name获取内容
func (obj *_AllVirtualProxyPartitionInfoMgr) GetFromPartKeyName(partKeyName string) (results []*AllVirtualProxyPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxyPartitionInfo{}).Where("`part_key_name` = ?", partKeyName).Find(&results).Error

	return
}

// GetBatchFromPartKeyName 批量查找
func (obj *_AllVirtualProxyPartitionInfoMgr) GetBatchFromPartKeyName(partKeyNames []string) (results []*AllVirtualProxyPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxyPartitionInfo{}).Where("`part_key_name` IN (?)", partKeyNames).Find(&results).Error

	return
}

// GetFromPartKeyType 通过part_key_type获取内容
func (obj *_AllVirtualProxyPartitionInfoMgr) GetFromPartKeyType(partKeyType int64) (results []*AllVirtualProxyPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxyPartitionInfo{}).Where("`part_key_type` = ?", partKeyType).Find(&results).Error

	return
}

// GetBatchFromPartKeyType 批量查找
func (obj *_AllVirtualProxyPartitionInfoMgr) GetBatchFromPartKeyType(partKeyTypes []int64) (results []*AllVirtualProxyPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxyPartitionInfo{}).Where("`part_key_type` IN (?)", partKeyTypes).Find(&results).Error

	return
}

// GetFromPartKeyIDx 通过part_key_idx获取内容
func (obj *_AllVirtualProxyPartitionInfoMgr) GetFromPartKeyIDx(partKeyIDx int64) (results []*AllVirtualProxyPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxyPartitionInfo{}).Where("`part_key_idx` = ?", partKeyIDx).Find(&results).Error

	return
}

// GetBatchFromPartKeyIDx 批量查找
func (obj *_AllVirtualProxyPartitionInfoMgr) GetBatchFromPartKeyIDx(partKeyIDxs []int64) (results []*AllVirtualProxyPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxyPartitionInfo{}).Where("`part_key_idx` IN (?)", partKeyIDxs).Find(&results).Error

	return
}

// GetFromPartKeyLevel 通过part_key_level获取内容
func (obj *_AllVirtualProxyPartitionInfoMgr) GetFromPartKeyLevel(partKeyLevel int64) (results []*AllVirtualProxyPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxyPartitionInfo{}).Where("`part_key_level` = ?", partKeyLevel).Find(&results).Error

	return
}

// GetBatchFromPartKeyLevel 批量查找
func (obj *_AllVirtualProxyPartitionInfoMgr) GetBatchFromPartKeyLevel(partKeyLevels []int64) (results []*AllVirtualProxyPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxyPartitionInfo{}).Where("`part_key_level` IN (?)", partKeyLevels).Find(&results).Error

	return
}

// GetFromPartKeyExtra 通过part_key_extra获取内容
func (obj *_AllVirtualProxyPartitionInfoMgr) GetFromPartKeyExtra(partKeyExtra string) (results []*AllVirtualProxyPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxyPartitionInfo{}).Where("`part_key_extra` = ?", partKeyExtra).Find(&results).Error

	return
}

// GetBatchFromPartKeyExtra 批量查找
func (obj *_AllVirtualProxyPartitionInfoMgr) GetBatchFromPartKeyExtra(partKeyExtras []string) (results []*AllVirtualProxyPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxyPartitionInfo{}).Where("`part_key_extra` IN (?)", partKeyExtras).Find(&results).Error

	return
}

// GetFromSpare1 通过spare1获取内容
func (obj *_AllVirtualProxyPartitionInfoMgr) GetFromSpare1(spare1 int64) (results []*AllVirtualProxyPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxyPartitionInfo{}).Where("`spare1` = ?", spare1).Find(&results).Error

	return
}

// GetBatchFromSpare1 批量查找
func (obj *_AllVirtualProxyPartitionInfoMgr) GetBatchFromSpare1(spare1s []int64) (results []*AllVirtualProxyPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxyPartitionInfo{}).Where("`spare1` IN (?)", spare1s).Find(&results).Error

	return
}

// GetFromSpare2 通过spare2获取内容
func (obj *_AllVirtualProxyPartitionInfoMgr) GetFromSpare2(spare2 int64) (results []*AllVirtualProxyPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxyPartitionInfo{}).Where("`spare2` = ?", spare2).Find(&results).Error

	return
}

// GetBatchFromSpare2 批量查找
func (obj *_AllVirtualProxyPartitionInfoMgr) GetBatchFromSpare2(spare2s []int64) (results []*AllVirtualProxyPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxyPartitionInfo{}).Where("`spare2` IN (?)", spare2s).Find(&results).Error

	return
}

// GetFromSpare3 通过spare3获取内容
func (obj *_AllVirtualProxyPartitionInfoMgr) GetFromSpare3(spare3 int64) (results []*AllVirtualProxyPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxyPartitionInfo{}).Where("`spare3` = ?", spare3).Find(&results).Error

	return
}

// GetBatchFromSpare3 批量查找
func (obj *_AllVirtualProxyPartitionInfoMgr) GetBatchFromSpare3(spare3s []int64) (results []*AllVirtualProxyPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxyPartitionInfo{}).Where("`spare3` IN (?)", spare3s).Find(&results).Error

	return
}

// GetFromSpare4 通过spare4获取内容
func (obj *_AllVirtualProxyPartitionInfoMgr) GetFromSpare4(spare4 string) (results []*AllVirtualProxyPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxyPartitionInfo{}).Where("`spare4` = ?", spare4).Find(&results).Error

	return
}

// GetBatchFromSpare4 批量查找
func (obj *_AllVirtualProxyPartitionInfoMgr) GetBatchFromSpare4(spare4s []string) (results []*AllVirtualProxyPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxyPartitionInfo{}).Where("`spare4` IN (?)", spare4s).Find(&results).Error

	return
}

// GetFromSpare5 通过spare5获取内容
func (obj *_AllVirtualProxyPartitionInfoMgr) GetFromSpare5(spare5 string) (results []*AllVirtualProxyPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxyPartitionInfo{}).Where("`spare5` = ?", spare5).Find(&results).Error

	return
}

// GetBatchFromSpare5 批量查找
func (obj *_AllVirtualProxyPartitionInfoMgr) GetBatchFromSpare5(spare5s []string) (results []*AllVirtualProxyPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxyPartitionInfo{}).Where("`spare5` IN (?)", spare5s).Find(&results).Error

	return
}

// GetFromSpare6 通过spare6获取内容
func (obj *_AllVirtualProxyPartitionInfoMgr) GetFromSpare6(spare6 string) (results []*AllVirtualProxyPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxyPartitionInfo{}).Where("`spare6` = ?", spare6).Find(&results).Error

	return
}

// GetBatchFromSpare6 批量查找
func (obj *_AllVirtualProxyPartitionInfoMgr) GetBatchFromSpare6(spare6s []string) (results []*AllVirtualProxyPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxyPartitionInfo{}).Where("`spare6` IN (?)", spare6s).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AllVirtualProxyPartitionInfoMgr) FetchByPrimaryKey(tableID int64) (result AllVirtualProxyPartitionInfo, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(AllVirtualProxyPartitionInfo{}).Where("`table_id` = ?", tableID).First(&result).Error

	return
}
