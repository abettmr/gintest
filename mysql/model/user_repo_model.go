package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ UserRepoModel = (*customUserRepoModel)(nil)

type (
	// UserRepoModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserRepoModel.
	UserRepoModel interface {
		userRepoModel
		withSession(session sqlx.Session) UserRepoModel
	}

	customUserRepoModel struct {
		*defaultUserRepoModel
	}
)

// NewUserRepoModel returns a model for the database table.
func NewUserRepoModel(conn sqlx.SqlConn) UserRepoModel {
	return &customUserRepoModel{
		defaultUserRepoModel: newUserRepoModel(conn),
	}
}

func (m *customUserRepoModel) withSession(session sqlx.Session) UserRepoModel {
	return NewUserRepoModel(sqlx.NewSqlConnFromSession(session))
}
