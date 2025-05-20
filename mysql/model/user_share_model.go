package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ UserShareModel = (*customUserShareModel)(nil)

type (
	// UserShareModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserShareModel.
	UserShareModel interface {
		userShareModel
		withSession(session sqlx.Session) UserShareModel
	}

	customUserShareModel struct {
		*defaultUserShareModel
	}
)

// NewUserShareModel returns a model for the database table.
func NewUserShareModel(conn sqlx.SqlConn) UserShareModel {
	return &customUserShareModel{
		defaultUserShareModel: newUserShareModel(conn),
	}
}

func (m *customUserShareModel) withSession(session sqlx.Session) UserShareModel {
	return NewUserShareModel(sqlx.NewSqlConnFromSession(session))
}
