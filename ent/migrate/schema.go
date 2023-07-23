// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// AttachmentsColumns holds the columns for the "attachments" table.
	AttachmentsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "path", Type: field.TypeString, Default: ""},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "name", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "user_attachments", Type: field.TypeInt, Nullable: true},
	}
	// AttachmentsTable holds the schema information for the "attachments" table.
	AttachmentsTable = &schema.Table{
		Name:       "attachments",
		Columns:    AttachmentsColumns,
		PrimaryKey: []*schema.Column{AttachmentsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "attachments_users_attachments",
				Columns:    []*schema.Column{AttachmentsColumns[5]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// CommentsColumns holds the columns for the "comments" table.
	CommentsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "content", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "parent_comment_id", Type: field.TypeInt, Nullable: true, Default: 0},
		{Name: "user_comments", Type: field.TypeInt, Nullable: true},
	}
	// CommentsTable holds the schema information for the "comments" table.
	CommentsTable = &schema.Table{
		Name:       "comments",
		Columns:    CommentsColumns,
		PrimaryKey: []*schema.Column{CommentsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "comments_users_comments",
				Columns:    []*schema.Column{CommentsColumns[5]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// PostsColumns holds the columns for the "posts" table.
	PostsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "content", Type: field.TypeString, Nullable: true, Default: ""},
	}
	// PostsTable holds the schema information for the "posts" table.
	PostsTable = &schema.Table{
		Name:       "posts",
		Columns:    PostsColumns,
		PrimaryKey: []*schema.Column{PostsColumns[0]},
	}
	// ReactsColumns holds the columns for the "reacts" table.
	ReactsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "react_type", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "post_type", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "user_reacted_user", Type: field.TypeInt, Nullable: true},
	}
	// ReactsTable holds the schema information for the "reacts" table.
	ReactsTable = &schema.Table{
		Name:       "reacts",
		Columns:    ReactsColumns,
		PrimaryKey: []*schema.Column{ReactsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "reacts_users_reacted_user",
				Columns:    []*schema.Column{ReactsColumns[5]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString, Default: ""},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// CommentAttachmentsColumns holds the columns for the "comment_attachments" table.
	CommentAttachmentsColumns = []*schema.Column{
		{Name: "comment_id", Type: field.TypeInt},
		{Name: "attachment_id", Type: field.TypeInt},
	}
	// CommentAttachmentsTable holds the schema information for the "comment_attachments" table.
	CommentAttachmentsTable = &schema.Table{
		Name:       "comment_attachments",
		Columns:    CommentAttachmentsColumns,
		PrimaryKey: []*schema.Column{CommentAttachmentsColumns[0], CommentAttachmentsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "comment_attachments_comment_id",
				Columns:    []*schema.Column{CommentAttachmentsColumns[0]},
				RefColumns: []*schema.Column{CommentsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "comment_attachments_attachment_id",
				Columns:    []*schema.Column{CommentAttachmentsColumns[1]},
				RefColumns: []*schema.Column{AttachmentsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// CommentReactsColumns holds the columns for the "comment_reacts" table.
	CommentReactsColumns = []*schema.Column{
		{Name: "comment_id", Type: field.TypeInt},
		{Name: "react_id", Type: field.TypeInt},
	}
	// CommentReactsTable holds the schema information for the "comment_reacts" table.
	CommentReactsTable = &schema.Table{
		Name:       "comment_reacts",
		Columns:    CommentReactsColumns,
		PrimaryKey: []*schema.Column{CommentReactsColumns[0], CommentReactsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "comment_reacts_comment_id",
				Columns:    []*schema.Column{CommentReactsColumns[0]},
				RefColumns: []*schema.Column{CommentsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "comment_reacts_react_id",
				Columns:    []*schema.Column{CommentReactsColumns[1]},
				RefColumns: []*schema.Column{ReactsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// PostAttachmentsColumns holds the columns for the "post_attachments" table.
	PostAttachmentsColumns = []*schema.Column{
		{Name: "post_id", Type: field.TypeInt},
		{Name: "attachment_id", Type: field.TypeInt},
	}
	// PostAttachmentsTable holds the schema information for the "post_attachments" table.
	PostAttachmentsTable = &schema.Table{
		Name:       "post_attachments",
		Columns:    PostAttachmentsColumns,
		PrimaryKey: []*schema.Column{PostAttachmentsColumns[0], PostAttachmentsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "post_attachments_post_id",
				Columns:    []*schema.Column{PostAttachmentsColumns[0]},
				RefColumns: []*schema.Column{PostsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "post_attachments_attachment_id",
				Columns:    []*schema.Column{PostAttachmentsColumns[1]},
				RefColumns: []*schema.Column{AttachmentsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// PostCommentsColumns holds the columns for the "post_comments" table.
	PostCommentsColumns = []*schema.Column{
		{Name: "post_id", Type: field.TypeInt},
		{Name: "comment_id", Type: field.TypeInt},
	}
	// PostCommentsTable holds the schema information for the "post_comments" table.
	PostCommentsTable = &schema.Table{
		Name:       "post_comments",
		Columns:    PostCommentsColumns,
		PrimaryKey: []*schema.Column{PostCommentsColumns[0], PostCommentsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "post_comments_post_id",
				Columns:    []*schema.Column{PostCommentsColumns[0]},
				RefColumns: []*schema.Column{PostsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "post_comments_comment_id",
				Columns:    []*schema.Column{PostCommentsColumns[1]},
				RefColumns: []*schema.Column{CommentsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// PostReactsColumns holds the columns for the "post_reacts" table.
	PostReactsColumns = []*schema.Column{
		{Name: "post_id", Type: field.TypeInt},
		{Name: "react_id", Type: field.TypeInt},
	}
	// PostReactsTable holds the schema information for the "post_reacts" table.
	PostReactsTable = &schema.Table{
		Name:       "post_reacts",
		Columns:    PostReactsColumns,
		PrimaryKey: []*schema.Column{PostReactsColumns[0], PostReactsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "post_reacts_post_id",
				Columns:    []*schema.Column{PostReactsColumns[0]},
				RefColumns: []*schema.Column{PostsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "post_reacts_react_id",
				Columns:    []*schema.Column{PostReactsColumns[1]},
				RefColumns: []*schema.Column{ReactsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// UserPostsColumns holds the columns for the "user_posts" table.
	UserPostsColumns = []*schema.Column{
		{Name: "user_id", Type: field.TypeInt},
		{Name: "post_id", Type: field.TypeInt},
	}
	// UserPostsTable holds the schema information for the "user_posts" table.
	UserPostsTable = &schema.Table{
		Name:       "user_posts",
		Columns:    UserPostsColumns,
		PrimaryKey: []*schema.Column{UserPostsColumns[0], UserPostsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "user_posts_user_id",
				Columns:    []*schema.Column{UserPostsColumns[0]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "user_posts_post_id",
				Columns:    []*schema.Column{UserPostsColumns[1]},
				RefColumns: []*schema.Column{PostsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AttachmentsTable,
		CommentsTable,
		PostsTable,
		ReactsTable,
		UsersTable,
		CommentAttachmentsTable,
		CommentReactsTable,
		PostAttachmentsTable,
		PostCommentsTable,
		PostReactsTable,
		UserPostsTable,
	}
)

func init() {
	AttachmentsTable.ForeignKeys[0].RefTable = UsersTable
	CommentsTable.ForeignKeys[0].RefTable = UsersTable
	ReactsTable.ForeignKeys[0].RefTable = UsersTable
	CommentAttachmentsTable.ForeignKeys[0].RefTable = CommentsTable
	CommentAttachmentsTable.ForeignKeys[1].RefTable = AttachmentsTable
	CommentReactsTable.ForeignKeys[0].RefTable = CommentsTable
	CommentReactsTable.ForeignKeys[1].RefTable = ReactsTable
	PostAttachmentsTable.ForeignKeys[0].RefTable = PostsTable
	PostAttachmentsTable.ForeignKeys[1].RefTable = AttachmentsTable
	PostCommentsTable.ForeignKeys[0].RefTable = PostsTable
	PostCommentsTable.ForeignKeys[1].RefTable = CommentsTable
	PostReactsTable.ForeignKeys[0].RefTable = PostsTable
	PostReactsTable.ForeignKeys[1].RefTable = ReactsTable
	UserPostsTable.ForeignKeys[0].RefTable = UsersTable
	UserPostsTable.ForeignKeys[1].RefTable = PostsTable
}
