package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ RepoPoolModel = (*customRepoPoolModel)(nil)

type (
	// RepoPoolModel is an interface to be customized, add more methods here,
	// and implement the added methods in customRepoPoolModel.
	RepoPoolModel interface {
		repoPoolModel
		withSession(session sqlx.Session) RepoPoolModel
	}

	customRepoPoolModel struct {
		*defaultRepoPoolModel
	}
)

// NewRepoPoolModel returns a model for the database table.
func NewRepoPoolModel(conn sqlx.SqlConn) RepoPoolModel {
	return &customRepoPoolModel{
		defaultRepoPoolModel: newRepoPoolModel(conn),
	}
}

func (m *customRepoPoolModel) withSession(session sqlx.Session) RepoPoolModel {
	return NewRepoPoolModel(sqlx.NewSqlConnFromSession(session))
}
