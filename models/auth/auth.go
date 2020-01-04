package auth

import (
	"github.com/astaxie/beego/orm"
)

func init() {
	orm.RegisterModel(new(JccRecord))
}

func AuthRequestAuth(token string) ([]string, error) {
	db := orm.NewOrm()

	var paths []string

	sql := `SELECT
			  jur.key as path
			FROM
			 jcc_token jt 
			INNER JOIN 
			 jcc_user_role_link jurl on jt.uid = jurl.uid
			INNER JOIN
			 jcc_user_role jur on jurl.role_id = jur.id
			WHERE jt.token = ?
			AND jurl.is_del = 0
			AND jur.is_del = 0`
	_, err := db.Raw(sql, token).QueryRows(&paths)

	return paths, err
}

func QueryAgentInfo(token string) (string, int, int, string, error) {
	db := orm.NewOrm()

	var cid, uid int
	var boss, openId string

	sql := `SELECT
				jt.uid AS uid,
				jt.openid AS openid,
				jb.aff_company AS cid,
                jb.boss as boss
			FROM
				(
					SELECT
						*
					FROM
						jcc_token
					WHERE
						token = ?
				) jt
			INNER JOIN jcc_boss jb ON jt.uid = jb.id`

	err := db.Raw(sql, token).QueryRow(&uid, &openId, &cid, &boss)

	return openId, cid, uid, boss, err
}

func QueryOpenIdBytoken(token string) (string, error) {
	db := orm.NewOrm()

	sql := `select openid from jcc_token where token = ?`

	var openId string

	err := db.Raw(sql, token).QueryRow(&openId)

	return openId, err
}

type JccRecord struct {
	Id        int    `orm:"column(id);auto"`
	Boss      string `orm:"column(boss);size(255)" description:"用户名"`
	CreatedAt int    `orm:"column(created_at)" description:"创建时间"`
	Param     string `orm:"column(param)" description:"请求参数"`
	Api       string `orm:"column(api);size(255);null" description:"后台api"`
	Path      string `orm:"column(path);size(255);null" description:"前端路由"`
	Ip        string `orm:"column(ip);size(255)" description:"用户ip"`
}

func (m *JccRecord) AddJccRecord() (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}
