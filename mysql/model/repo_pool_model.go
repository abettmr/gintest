package model

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ RepoPoolModel = (*customRepoPoolModel)(nil)

type (
	// RepoPoolModel is an interface to be customized, add more methods here,
	// and implement the added methods in customRepoPoolModel.
	RepoPoolModel interface {
		repoPoolModel
		withSession(session sqlx.Session) RepoPoolModel
		FindPath(ctx context.Context, path string) (*[]RepoPool, error)
		FindHash(ctx context.Context, hash string) (*RepoPool, error)
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

func (m *customRepoPoolModel) FindHash(ctx context.Context, hash string) (*RepoPool, error) {
	query := fmt.Sprintf("select %s from %s where `hash` = ?", repoPoolRows, m.table)
	var resp RepoPool
	err := m.conn.QueryRowCtx(ctx, &resp, query, hash)
	switch err {
	case nil:
		return &resp, nil
	case sqlx.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *customRepoPoolModel) FindPath(ctx context.Context, id string) (*[]RepoPool, error) {
	query := fmt.Sprintf("select %s from %s where `path` = ?", repoPoolRows, m.table)
	var resp []RepoPool
	err := m.conn.QueryRowCtx(ctx, &resp, query, id)
	switch err {
	case nil:
		return &resp, nil
	case sqlx.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
