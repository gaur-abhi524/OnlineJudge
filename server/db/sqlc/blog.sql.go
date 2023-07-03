// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: blog.sql

package db

import (
	"context"
	"database/sql"
)

const createBlog = `-- name: CreateBlog :one
INSERT INTO blogs (
  blog_title,
  blog_content,
  created_by,
  ispublish
) VALUES (
  $1, $2, $3, $4
) RETURNING id, blog_title, blog_content, created_by, created_at, ispublish, votes_count
`

type CreateBlogParams struct {
	BlogTitle   string       `json:"blog_title"`
	BlogContent string       `json:"blog_content"`
	CreatedBy   string       `json:"created_by"`
	Ispublish   sql.NullBool `json:"ispublish"`
}

func (q *Queries) CreateBlog(ctx context.Context, arg CreateBlogParams) (Blog, error) {
	row := q.db.QueryRowContext(ctx, createBlog,
		arg.BlogTitle,
		arg.BlogContent,
		arg.CreatedBy,
		arg.Ispublish,
	)
	var i Blog
	err := row.Scan(
		&i.ID,
		&i.BlogTitle,
		&i.BlogContent,
		&i.CreatedBy,
		&i.CreatedAt,
		&i.Ispublish,
		&i.VotesCount,
	)
	return i, err
}

const deleteBlog = `-- name: DeleteBlog :exec
DELETE FROM blogs
WHERE id = $1
`

func (q *Queries) DeleteBlog(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteBlog, id)
	return err
}

const getBlog = `-- name: GetBlog :one
SELECT id, blog_title, blog_content, created_by, created_at, ispublish, votes_count FROM blogs
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetBlog(ctx context.Context, id int64) (Blog, error) {
	row := q.db.QueryRowContext(ctx, getBlog, id)
	var i Blog
	err := row.Scan(
		&i.ID,
		&i.BlogTitle,
		&i.BlogContent,
		&i.CreatedBy,
		&i.CreatedAt,
		&i.Ispublish,
		&i.VotesCount,
	)
	return i, err
}

const listBlogs = `-- name: ListBlogs :many
SELECT id, blog_title, blog_content, created_by, created_at, ispublish, votes_count FROM blogs
ORDER BY id
LIMIT $1
OFFSET $2
`

type ListBlogsParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListBlogs(ctx context.Context, arg ListBlogsParams) ([]Blog, error) {
	rows, err := q.db.QueryContext(ctx, listBlogs, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Blog{}
	for rows.Next() {
		var i Blog
		if err := rows.Scan(
			&i.ID,
			&i.BlogTitle,
			&i.BlogContent,
			&i.CreatedBy,
			&i.CreatedAt,
			&i.Ispublish,
			&i.VotesCount,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateBlog = `-- name: UpdateBlog :one
UPDATE blogs
SET
  blog_title = COALESCE($1, blog_title),
  blog_content = COALESCE($2, blog_content),
  ispublish = COALESCE($3, ispublish)
WHERE id = $4
RETURNING id, blog_title, blog_content, created_by, created_at, ispublish, votes_count
`

type UpdateBlogParams struct {
	BlogTitle   sql.NullString `json:"blog_title"`
	BlogContent sql.NullString `json:"blog_content"`
	Ispublish   sql.NullBool   `json:"ispublish"`
	ID          int64          `json:"id"`
}

func (q *Queries) UpdateBlog(ctx context.Context, arg UpdateBlogParams) (Blog, error) {
	row := q.db.QueryRowContext(ctx, updateBlog,
		arg.BlogTitle,
		arg.BlogContent,
		arg.Ispublish,
		arg.ID,
	)
	var i Blog
	err := row.Scan(
		&i.ID,
		&i.BlogTitle,
		&i.BlogContent,
		&i.CreatedBy,
		&i.CreatedAt,
		&i.Ispublish,
		&i.VotesCount,
	)
	return i, err
}
