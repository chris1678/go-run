package mycasbin

import (
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gormAdapter "github.com/casbin/gorm-adapter/v3"
	"github.com/chris1678/go-run/logger"
	"gorm.io/gorm"
)

// Initialize the model from a string.
var text = `
[request_definition]
r = sub, obj, act

[policy_definition]
p = sub, obj, act

[policy_effect]
e = some(where (p.eft == allow))

[matchers]
m = r.sub == p.sub && (keyMatch2(r.obj, p.obj) || keyMatch(r.obj, p.obj)) && (r.act == p.act || p.act == "*")
`

func Setup(db *gorm.DB, tablename string) *casbin.SyncedEnforcer {
	Apter, err := gormAdapter.NewAdapterByDBUseTableName(db, "", tablename)
	if err != nil {
		logger.LogHelper.Panic(err)
		return nil
	}
	m, err2 := model.NewModelFromString(text)
	if err2 != nil {
		logger.LogHelper.Panic(err2)
		return nil
	}
	e, err3 := casbin.NewSyncedEnforcer(m, Apter)
	if err3 != nil {
		logger.LogHelper.Panic(err3)
		return nil
	}

	l := &Logger{}
	l.EnableLog(true)
	e.SetLogger(l)

	err4 := e.LoadPolicy()
	if err4 != nil {
		logger.LogHelper.Panic(err4)
		return nil
	}

	return e
}
