package model

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/tal-tech/go-zero/core/stores/cache"
	"github.com/tal-tech/go-zero/core/stores/sqlc"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"github.com/tal-tech/go-zero/core/stringx"
	"github.com/tal-tech/go-zero/tools/goctl/model/sql/builderx"
)

var (
	xqdAlbumAnnex0FieldNames          = builderx.RawFieldNames(&XqdAlbumAnnex0{})
	xqdAlbumAnnex0Rows                = strings.Join(xqdAlbumAnnex0FieldNames, ",")
	xqdAlbumAnnex0RowsExpectAutoSet   = strings.Join(stringx.Remove(xqdAlbumAnnex0FieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	xqdAlbumAnnex0RowsWithPlaceHolder = strings.Join(stringx.Remove(xqdAlbumAnnex0FieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheXqdAlbumAnnex0IdPrefix = "cache::xqdAlbumAnnex0:id:"
)

type (
	XqdAlbumAnnex0Model interface {
		Insert(data XqdAlbumAnnex0) (sql.Result, error)
		FindOne(id int64) (*XqdAlbumAnnex0, error)
		Update(data XqdAlbumAnnex0) error
		Delete(id int64) error
		FindAll(sql string, param []interface{}) ([]*XqdAlbumAnnex0, error)
		Collation(param []interface{}, rep map[string]interface{}) (string, []interface{})
	}

	defaultXqdAlbumAnnex0Model struct {
		sqlc.CachedConn
		table string
	}

	XqdAlbumAnnex0 struct {
		Id          int64          `db:"id"`           // 主键id
		Uid         int64          `db:"uid"`          // 用户ID
		GroupId     int64          `db:"group_id"`     // 群组id
		Drive       string         `db:"drive"`        // 驱动:oss
		UploadType  string         `db:"upload_type"`  // 上传类型
		MimeType    string         `db:"mime_type"`    // 类别
		BaseUrl     string         `db:"base_url"`     // url
		Path        string         `db:"path"`         // 本地路径
		Md5         string         `db:"md5"`          // md5校验码
		Name        string         `db:"name"`         // 文件原始名
		Extension   string         `db:"extension"`    // 扩展名
		Size        int64          `db:"size"`         // 长度
		Year        int64          `db:"year"`         // 年份
		Month       int64          `db:"month"`        // 月份
		Day         int64          `db:"day"`          // 日
		Width       int64          `db:"width"`        // 宽度
		Height      int64          `db:"height"`       // 高度
		UploadIp    string         `db:"upload_ip"`    // 上传者ip
		Duration    string         `db:"duration"`     // 音频时长（秒）
		CreatedTime int64          `db:"created_time"` // 生成时间/拍摄时间
		Log         string         `db:"log"`          // 经度
		Lat         string         `db:"lat"`          // 维度
		Origin      int64          `db:"origin"`       // 来源 1android 2ios 3微信 4社区 5MCN 6社区后台 7相册 8老照片 9影集
		CopyId      int64          `db:"copy_id"`      // 复制原id
		Status      int64          `db:"status"`       // 状态[1正常(审核成功);2已禁用;3已删除;4待审核;5人工审核;6审核失败]
		FailReason  string         `db:"fail_reason"`  // 审核失败原因
		CreatedAt   sql.NullInt64  `db:"created_at"`   // 创建时间
		UpdatedAt   sql.NullInt64  `db:"updated_at"`   // 修改时间
		LocalPath   sql.NullString `db:"local_path"`   // APP端路径
		Author      string         `db:"author"`       // 作者
		Type        int64          `db:"type"`         // 附件类型 1 :视频 2图片 3音频  4,5空着  6 音乐
		CheckType   int64          `db:"check_type"`   // 审核结果: [1:正常, 2:嫌疑, 3:违禁]
		CheckInfo   string         `db:"check_info"`   // 审核结果信息
		Position    int64          `db:"position"`     // 使用位置 [0:未知, 1:头像, 2:社区, 3:照片墙, 4:老照片, 5:四点定位]
	}
)

func NewXqdAlbumAnnex0Model(conn sqlx.SqlConn, c cache.CacheConf) XqdAlbumAnnex0Model {
	return &defaultXqdAlbumAnnex0Model{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`xqd_album_annex_0`",
	}
}

func (m *defaultXqdAlbumAnnex0Model) Insert(data XqdAlbumAnnex0) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, xqdAlbumAnnex0RowsExpectAutoSet)
	ret, err := m.ExecNoCache(query, data.Uid, data.GroupId, data.Drive, data.UploadType, data.MimeType, data.BaseUrl, data.Path, data.Md5, data.Name, data.Extension, data.Size, data.Year, data.Month, data.Day, data.Width, data.Height, data.UploadIp, data.Duration, data.CreatedTime, data.Log, data.Lat, data.Origin, data.CopyId, data.Status, data.FailReason, data.CreatedAt, data.UpdatedAt, data.LocalPath, data.Author, data.Type, data.CheckType, data.CheckInfo, data.Position)

	return ret, err
}

func (m *defaultXqdAlbumAnnex0Model) FindOne(id int64) (*XqdAlbumAnnex0, error) {
	xqdAlbumAnnex0IdKey := fmt.Sprintf("%s%v", cacheXqdAlbumAnnex0IdPrefix, id)
	var resp XqdAlbumAnnex0
	err := m.QueryRow(&resp, xqdAlbumAnnex0IdKey, func(conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", xqdAlbumAnnex0Rows, m.table)
		return conn.QueryRow(v, query, id)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultXqdAlbumAnnex0Model) FindAll(sql string, param []interface{}) ([]*XqdAlbumAnnex0, error) {
	annex := make([]*XqdAlbumAnnex0, 0)
	query := fmt.Sprintf("select %s from %s where "+sql, xqdAlbumAnnex0Rows, m.table)
	err := m.QueryRowsNoCache(&annex, query, param...)
	switch err {
	case nil:
		return annex, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultXqdAlbumAnnex0Model) Collation(param []interface{}, rep map[string]interface{}) (string, []interface{}) {
	var querySql string
	var arr []interface{}
	for _, v := range param {
		s := v.([]interface{})
		querySql = querySql + " `" + s[0].(string) + "` " + s[1].(string) + " ? and"
		arr = append(arr, s[2])
	}
	querySql = strings.TrimRight(querySql, "and")
	orderBySql := "id desc"
	if _, ok := rep["orderBy"]; !ok {
		orderBySql = rep["orderBy"].(string)
	}
	limitSql := ""
	var limitOk bool
	var offsetOk bool
	if _, ok := rep["limit"]; ok {
		limitOk = true
	}
	if _, ok := rep["offset"]; ok {
		offsetOk = true
	}
	if limitOk && offsetOk {
		limitSql = "limit " + rep["limit"].(string) + " offset " + rep["offset"].(string)
	}
	querySql = querySql + " order by " + orderBySql + limitSql
	return querySql, arr
}

func (m *defaultXqdAlbumAnnex0Model) Update(data XqdAlbumAnnex0) error {
	xqdAlbumAnnex0IdKey := fmt.Sprintf("%s%v", cacheXqdAlbumAnnex0IdPrefix, data.Id)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, xqdAlbumAnnex0RowsWithPlaceHolder)
		return conn.Exec(query, data.Uid, data.GroupId, data.Drive, data.UploadType, data.MimeType, data.BaseUrl, data.Path, data.Md5, data.Name, data.Extension, data.Size, data.Year, data.Month, data.Day, data.Width, data.Height, data.UploadIp, data.Duration, data.CreatedTime, data.Log, data.Lat, data.Origin, data.CopyId, data.Status, data.FailReason, data.CreatedAt, data.UpdatedAt, data.LocalPath, data.Author, data.Type, data.CheckType, data.CheckInfo, data.Position, data.Id)
	}, xqdAlbumAnnex0IdKey)
	return err
}

func (m *defaultXqdAlbumAnnex0Model) Delete(id int64) error {

	xqdAlbumAnnex0IdKey := fmt.Sprintf("%s%v", cacheXqdAlbumAnnex0IdPrefix, id)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.Exec(query, id)
	}, xqdAlbumAnnex0IdKey)
	return err
}

func (m *defaultXqdAlbumAnnex0Model) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheXqdAlbumAnnex0IdPrefix, primary)
}

func (m *defaultXqdAlbumAnnex0Model) queryPrimary(conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", xqdAlbumAnnex0Rows, m.table)
	return conn.QueryRow(v, query, primary)
}
