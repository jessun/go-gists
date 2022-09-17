package	model	
import (	
"gorm.io/gorm"	
"fmt"	
"context"	
)	

type _V$lockMgr struct {
	*_BaseMgr
}

// V$lockMgr open func
func V$lockMgr(db *gorm.DB) *_V$lockMgr {
	if db == nil {
		panic(fmt.Errorf("V$lockMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_V$lockMgr{_BaseMgr: &_BaseMgr{DB: db.Table("v$lock"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_V$lockMgr) GetTableName() string {
	return "v$lock"
}

// Reset 重置gorm会话
func (obj *_V$lockMgr) Reset() *_V$lockMgr {
	obj.New()
	return obj
}

// Get 获取 
func (obj *_V$lockMgr) Get() (result V$lock, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$lock{}).First(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_V$lockMgr) Gets() (results []*V$lock, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$lock{}).Find(&results).Error
	
	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_V$lockMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(V$lock{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTableID TABLE_ID获取 
func (obj *_V$lockMgr) WithTableID(tableID int64) Option {
	return optionFunc(func(o *options) { o.query["TABLE_ID"] = tableID })
}

// WithAddr ADDR获取 
func (obj *_V$lockMgr) WithAddr(addr string) Option {
	return optionFunc(func(o *options) { o.query["ADDR"] = addr })
}

// WithKaddr KADDR获取 
func (obj *_V$lockMgr) WithKaddr(kaddr uint64) Option {
	return optionFunc(func(o *options) { o.query["KADDR"] = kaddr })
}

// WithSid SID获取 
func (obj *_V$lockMgr) WithSid(sid int64) Option {
	return optionFunc(func(o *options) { o.query["SID"] = sid })
}

// WithType TYPE获取 
func (obj *_V$lockMgr) WithType(_type int64) Option {
	return optionFunc(func(o *options) { o.query["TYPE"] = _type })
}

// WithLmode LMODE获取 
func (obj *_V$lockMgr) WithLmode(lmode int64) Option {
	return optionFunc(func(o *options) { o.query["LMODE"] = lmode })
}

// WithRequest REQUEST获取 
func (obj *_V$lockMgr) WithRequest(request int64) Option {
	return optionFunc(func(o *options) { o.query["REQUEST"] = request })
}

// WithCtime CTIME获取 
func (obj *_V$lockMgr) WithCtime(ctime int64) Option {
	return optionFunc(func(o *options) { o.query["CTIME"] = ctime })
}

// WithBlock BLOCK获取 
func (obj *_V$lockMgr) WithBlock(block int64) Option {
	return optionFunc(func(o *options) { o.query["BLOCK"] = block })
}

// WithConID CON_ID获取 
func (obj *_V$lockMgr) WithConID(conID uint64) Option {
	return optionFunc(func(o *options) { o.query["CON_ID"] = conID })
}


// GetByOption 功能选项模式获取
func (obj *_V$lockMgr) GetByOption(opts ...Option) (result V$lock, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(V$lock{}).Where(options.query).First(&result).Error
	
	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_V$lockMgr) GetByOptions(opts ...Option) (results []*V$lock, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(V$lock{}).Where(options.query).Find(&results).Error
	
	return
}


// SelectPage 分页查询
func (obj *_V$lockMgr) SelectPage(page IPage,opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]V$lock,0)
	var count int64 // 统计总的记录数
	query :=  obj.DB.WithContext(obj.ctx).Model(V$lock{}).Where(options.query)
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


// GetFromTableID 通过TABLE_ID获取内容  
func (obj *_V$lockMgr) GetFromTableID(tableID int64) (results []*V$lock, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$lock{}).Where("`TABLE_ID` = ?", tableID).Find(&results).Error
	
	return
}

// GetBatchFromTableID 批量查找 
func (obj *_V$lockMgr) GetBatchFromTableID(tableIDs []int64) (results []*V$lock, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$lock{}).Where("`TABLE_ID` IN (?)", tableIDs).Find(&results).Error
	
	return
}
 
// GetFromAddr 通过ADDR获取内容  
func (obj *_V$lockMgr) GetFromAddr(addr string) (results []*V$lock, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$lock{}).Where("`ADDR` = ?", addr).Find(&results).Error
	
	return
}

// GetBatchFromAddr 批量查找 
func (obj *_V$lockMgr) GetBatchFromAddr(addrs []string) (results []*V$lock, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$lock{}).Where("`ADDR` IN (?)", addrs).Find(&results).Error
	
	return
}
 
// GetFromKaddr 通过KADDR获取内容  
func (obj *_V$lockMgr) GetFromKaddr(kaddr uint64) (results []*V$lock, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$lock{}).Where("`KADDR` = ?", kaddr).Find(&results).Error
	
	return
}

// GetBatchFromKaddr 批量查找 
func (obj *_V$lockMgr) GetBatchFromKaddr(kaddrs []uint64) (results []*V$lock, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$lock{}).Where("`KADDR` IN (?)", kaddrs).Find(&results).Error
	
	return
}
 
// GetFromSid 通过SID获取内容  
func (obj *_V$lockMgr) GetFromSid(sid int64) (results []*V$lock, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$lock{}).Where("`SID` = ?", sid).Find(&results).Error
	
	return
}

// GetBatchFromSid 批量查找 
func (obj *_V$lockMgr) GetBatchFromSid(sids []int64) (results []*V$lock, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$lock{}).Where("`SID` IN (?)", sids).Find(&results).Error
	
	return
}
 
// GetFromType 通过TYPE获取内容  
func (obj *_V$lockMgr) GetFromType(_type int64) (results []*V$lock, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$lock{}).Where("`TYPE` = ?", _type).Find(&results).Error
	
	return
}

// GetBatchFromType 批量查找 
func (obj *_V$lockMgr) GetBatchFromType(_types []int64) (results []*V$lock, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$lock{}).Where("`TYPE` IN (?)", _types).Find(&results).Error
	
	return
}
 
// GetFromLmode 通过LMODE获取内容  
func (obj *_V$lockMgr) GetFromLmode(lmode int64) (results []*V$lock, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$lock{}).Where("`LMODE` = ?", lmode).Find(&results).Error
	
	return
}

// GetBatchFromLmode 批量查找 
func (obj *_V$lockMgr) GetBatchFromLmode(lmodes []int64) (results []*V$lock, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$lock{}).Where("`LMODE` IN (?)", lmodes).Find(&results).Error
	
	return
}
 
// GetFromRequest 通过REQUEST获取内容  
func (obj *_V$lockMgr) GetFromRequest(request int64) (results []*V$lock, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$lock{}).Where("`REQUEST` = ?", request).Find(&results).Error
	
	return
}

// GetBatchFromRequest 批量查找 
func (obj *_V$lockMgr) GetBatchFromRequest(requests []int64) (results []*V$lock, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$lock{}).Where("`REQUEST` IN (?)", requests).Find(&results).Error
	
	return
}
 
// GetFromCtime 通过CTIME获取内容  
func (obj *_V$lockMgr) GetFromCtime(ctime int64) (results []*V$lock, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$lock{}).Where("`CTIME` = ?", ctime).Find(&results).Error
	
	return
}

// GetBatchFromCtime 批量查找 
func (obj *_V$lockMgr) GetBatchFromCtime(ctimes []int64) (results []*V$lock, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$lock{}).Where("`CTIME` IN (?)", ctimes).Find(&results).Error
	
	return
}
 
// GetFromBlock 通过BLOCK获取内容  
func (obj *_V$lockMgr) GetFromBlock(block int64) (results []*V$lock, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$lock{}).Where("`BLOCK` = ?", block).Find(&results).Error
	
	return
}

// GetBatchFromBlock 批量查找 
func (obj *_V$lockMgr) GetBatchFromBlock(blocks []int64) (results []*V$lock, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$lock{}).Where("`BLOCK` IN (?)", blocks).Find(&results).Error
	
	return
}
 
// GetFromConID 通过CON_ID获取内容  
func (obj *_V$lockMgr) GetFromConID(conID uint64) (results []*V$lock, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$lock{}).Where("`CON_ID` = ?", conID).Find(&results).Error
	
	return
}

// GetBatchFromConID 批量查找 
func (obj *_V$lockMgr) GetBatchFromConID(conIDs []uint64) (results []*V$lock, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(V$lock{}).Where("`CON_ID` IN (?)", conIDs).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 

 

	

