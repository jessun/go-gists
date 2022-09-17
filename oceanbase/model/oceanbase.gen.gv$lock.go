package	model	
import (	
"gorm.io/gorm"	
"fmt"	
"context"	
)	

type _Gv$lockMgr struct {
	*_BaseMgr
}

// Gv$lockMgr open func
func Gv$lockMgr(db *gorm.DB) *_Gv$lockMgr {
	if db == nil {
		panic(fmt.Errorf("Gv$lockMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_Gv$lockMgr{_BaseMgr: &_BaseMgr{DB: db.Table("gv$lock"), isRelated: globalIsRelated,ctx:ctx,cancel:cancel,timeout:-1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_Gv$lockMgr) GetTableName() string {
	return "gv$lock"
}

// Reset 重置gorm会话
func (obj *_Gv$lockMgr) Reset() *_Gv$lockMgr {
	obj.New()
	return obj
}

// Get 获取 
func (obj *_Gv$lockMgr) Get() (result Gv$lock, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$lock{}).First(&result).Error
	
	return
}

// Gets 获取批量结果
func (obj *_Gv$lockMgr) Gets() (results []*Gv$lock, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$lock{}).Find(&results).Error
	
	return
}

////////////////////////////////// gorm replace /////////////////////////////////
func (obj *_Gv$lockMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(Gv$lock{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithSvrIP SVR_IP获取 
func (obj *_Gv$lockMgr) WithSvrIP(svrIP string) Option {
	return optionFunc(func(o *options) { o.query["SVR_IP"] = svrIP })
}

// WithSvrPort SVR_PORT获取 
func (obj *_Gv$lockMgr) WithSvrPort(svrPort int64) Option {
	return optionFunc(func(o *options) { o.query["SVR_PORT"] = svrPort })
}

// WithTableID TABLE_ID获取 
func (obj *_Gv$lockMgr) WithTableID(tableID int64) Option {
	return optionFunc(func(o *options) { o.query["TABLE_ID"] = tableID })
}

// WithAddr ADDR获取 
func (obj *_Gv$lockMgr) WithAddr(addr string) Option {
	return optionFunc(func(o *options) { o.query["ADDR"] = addr })
}

// WithKaddr KADDR获取 
func (obj *_Gv$lockMgr) WithKaddr(kaddr uint64) Option {
	return optionFunc(func(o *options) { o.query["KADDR"] = kaddr })
}

// WithSid SID获取 
func (obj *_Gv$lockMgr) WithSid(sid int64) Option {
	return optionFunc(func(o *options) { o.query["SID"] = sid })
}

// WithType TYPE获取 
func (obj *_Gv$lockMgr) WithType(_type int64) Option {
	return optionFunc(func(o *options) { o.query["TYPE"] = _type })
}

// WithLmode LMODE获取 
func (obj *_Gv$lockMgr) WithLmode(lmode int64) Option {
	return optionFunc(func(o *options) { o.query["LMODE"] = lmode })
}

// WithRequest REQUEST获取 
func (obj *_Gv$lockMgr) WithRequest(request int64) Option {
	return optionFunc(func(o *options) { o.query["REQUEST"] = request })
}

// WithCtime CTIME获取 
func (obj *_Gv$lockMgr) WithCtime(ctime int64) Option {
	return optionFunc(func(o *options) { o.query["CTIME"] = ctime })
}

// WithBlock BLOCK获取 
func (obj *_Gv$lockMgr) WithBlock(block int64) Option {
	return optionFunc(func(o *options) { o.query["BLOCK"] = block })
}

// WithConID CON_ID获取 
func (obj *_Gv$lockMgr) WithConID(conID uint64) Option {
	return optionFunc(func(o *options) { o.query["CON_ID"] = conID })
}


// GetByOption 功能选项模式获取
func (obj *_Gv$lockMgr) GetByOption(opts ...Option) (result Gv$lock, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Gv$lock{}).Where(options.query).First(&result).Error
	
	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_Gv$lockMgr) GetByOptions(opts ...Option) (results []*Gv$lock, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Gv$lock{}).Where(options.query).Find(&results).Error
	
	return
}


// SelectPage 分页查询
func (obj *_Gv$lockMgr) SelectPage(page IPage,opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]Gv$lock,0)
	var count int64 // 统计总的记录数
	query :=  obj.DB.WithContext(obj.ctx).Model(Gv$lock{}).Where(options.query)
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


// GetFromSvrIP 通过SVR_IP获取内容  
func (obj *_Gv$lockMgr) GetFromSvrIP(svrIP string) (results []*Gv$lock, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$lock{}).Where("`SVR_IP` = ?", svrIP).Find(&results).Error
	
	return
}

// GetBatchFromSvrIP 批量查找 
func (obj *_Gv$lockMgr) GetBatchFromSvrIP(svrIPs []string) (results []*Gv$lock, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$lock{}).Where("`SVR_IP` IN (?)", svrIPs).Find(&results).Error
	
	return
}
 
// GetFromSvrPort 通过SVR_PORT获取内容  
func (obj *_Gv$lockMgr) GetFromSvrPort(svrPort int64) (results []*Gv$lock, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$lock{}).Where("`SVR_PORT` = ?", svrPort).Find(&results).Error
	
	return
}

// GetBatchFromSvrPort 批量查找 
func (obj *_Gv$lockMgr) GetBatchFromSvrPort(svrPorts []int64) (results []*Gv$lock, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$lock{}).Where("`SVR_PORT` IN (?)", svrPorts).Find(&results).Error
	
	return
}
 
// GetFromTableID 通过TABLE_ID获取内容  
func (obj *_Gv$lockMgr) GetFromTableID(tableID int64) (results []*Gv$lock, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$lock{}).Where("`TABLE_ID` = ?", tableID).Find(&results).Error
	
	return
}

// GetBatchFromTableID 批量查找 
func (obj *_Gv$lockMgr) GetBatchFromTableID(tableIDs []int64) (results []*Gv$lock, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$lock{}).Where("`TABLE_ID` IN (?)", tableIDs).Find(&results).Error
	
	return
}
 
// GetFromAddr 通过ADDR获取内容  
func (obj *_Gv$lockMgr) GetFromAddr(addr string) (results []*Gv$lock, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$lock{}).Where("`ADDR` = ?", addr).Find(&results).Error
	
	return
}

// GetBatchFromAddr 批量查找 
func (obj *_Gv$lockMgr) GetBatchFromAddr(addrs []string) (results []*Gv$lock, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$lock{}).Where("`ADDR` IN (?)", addrs).Find(&results).Error
	
	return
}
 
// GetFromKaddr 通过KADDR获取内容  
func (obj *_Gv$lockMgr) GetFromKaddr(kaddr uint64) (results []*Gv$lock, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$lock{}).Where("`KADDR` = ?", kaddr).Find(&results).Error
	
	return
}

// GetBatchFromKaddr 批量查找 
func (obj *_Gv$lockMgr) GetBatchFromKaddr(kaddrs []uint64) (results []*Gv$lock, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$lock{}).Where("`KADDR` IN (?)", kaddrs).Find(&results).Error
	
	return
}
 
// GetFromSid 通过SID获取内容  
func (obj *_Gv$lockMgr) GetFromSid(sid int64) (results []*Gv$lock, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$lock{}).Where("`SID` = ?", sid).Find(&results).Error
	
	return
}

// GetBatchFromSid 批量查找 
func (obj *_Gv$lockMgr) GetBatchFromSid(sids []int64) (results []*Gv$lock, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$lock{}).Where("`SID` IN (?)", sids).Find(&results).Error
	
	return
}
 
// GetFromType 通过TYPE获取内容  
func (obj *_Gv$lockMgr) GetFromType(_type int64) (results []*Gv$lock, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$lock{}).Where("`TYPE` = ?", _type).Find(&results).Error
	
	return
}

// GetBatchFromType 批量查找 
func (obj *_Gv$lockMgr) GetBatchFromType(_types []int64) (results []*Gv$lock, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$lock{}).Where("`TYPE` IN (?)", _types).Find(&results).Error
	
	return
}
 
// GetFromLmode 通过LMODE获取内容  
func (obj *_Gv$lockMgr) GetFromLmode(lmode int64) (results []*Gv$lock, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$lock{}).Where("`LMODE` = ?", lmode).Find(&results).Error
	
	return
}

// GetBatchFromLmode 批量查找 
func (obj *_Gv$lockMgr) GetBatchFromLmode(lmodes []int64) (results []*Gv$lock, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$lock{}).Where("`LMODE` IN (?)", lmodes).Find(&results).Error
	
	return
}
 
// GetFromRequest 通过REQUEST获取内容  
func (obj *_Gv$lockMgr) GetFromRequest(request int64) (results []*Gv$lock, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$lock{}).Where("`REQUEST` = ?", request).Find(&results).Error
	
	return
}

// GetBatchFromRequest 批量查找 
func (obj *_Gv$lockMgr) GetBatchFromRequest(requests []int64) (results []*Gv$lock, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$lock{}).Where("`REQUEST` IN (?)", requests).Find(&results).Error
	
	return
}
 
// GetFromCtime 通过CTIME获取内容  
func (obj *_Gv$lockMgr) GetFromCtime(ctime int64) (results []*Gv$lock, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$lock{}).Where("`CTIME` = ?", ctime).Find(&results).Error
	
	return
}

// GetBatchFromCtime 批量查找 
func (obj *_Gv$lockMgr) GetBatchFromCtime(ctimes []int64) (results []*Gv$lock, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$lock{}).Where("`CTIME` IN (?)", ctimes).Find(&results).Error
	
	return
}
 
// GetFromBlock 通过BLOCK获取内容  
func (obj *_Gv$lockMgr) GetFromBlock(block int64) (results []*Gv$lock, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$lock{}).Where("`BLOCK` = ?", block).Find(&results).Error
	
	return
}

// GetBatchFromBlock 批量查找 
func (obj *_Gv$lockMgr) GetBatchFromBlock(blocks []int64) (results []*Gv$lock, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$lock{}).Where("`BLOCK` IN (?)", blocks).Find(&results).Error
	
	return
}
 
// GetFromConID 通过CON_ID获取内容  
func (obj *_Gv$lockMgr) GetFromConID(conID uint64) (results []*Gv$lock, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$lock{}).Where("`CON_ID` = ?", conID).Find(&results).Error
	
	return
}

// GetBatchFromConID 批量查找 
func (obj *_Gv$lockMgr) GetBatchFromConID(conIDs []uint64) (results []*Gv$lock, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Gv$lock{}).Where("`CON_ID` IN (?)", conIDs).Find(&results).Error
	
	return
}
 
 //////////////////////////primary index case ////////////////////////////////////////////
 

 

	

