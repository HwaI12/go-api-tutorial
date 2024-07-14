package model

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	errors "github.com/HwaI12/go-api-tutorial/internal/error"
	transaction "github.com/HwaI12/go-api-tutorial/internal/transaction"
	"github.com/stretchr/testify/require"
)

func TestBook_Validate(t *testing.T) {
	type fields struct {
		ID        string
		Name      string
		Price     int
		CreatedAt time.Time
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
		err     *errors.UserDefinedError
	}{
		{
			name: "valid book",
			fields: fields{
				Name:  "Valid Book",
				Price: 1000,
			},
			args:    args{ctx: context.Background()},
			wantErr: false,
		},
		{
			name: "empty book name",
			fields: fields{
				Name:  "",
				Price: 1000,
			},
			args:    args{ctx: context.Background()},
			wantErr: true,
			err:     errors.BookNameEmptyError(),
		},
		{
			name: "book name too long",
			fields: fields{
				Name:  "This is a very long book name that exceeds the maximum allowed length of fifty characters",
				Price: 1000,
			},
			args:    args{ctx: context.Background()},
			wantErr: true,
			err:     errors.BookNameTooLongError(),
		},
		{
			name: "negative book price",
			fields: fields{
				Name:  "Negative Price Book",
				Price: -100,
			},
			args:    args{ctx: context.Background()},
			wantErr: true,
			err:     errors.BookPriceNegativeError(),
		},
		{
			name: "zero book price",
			fields: fields{
				Name:  "Zero Price Book",
				Price: 0,
			},
			args:    args{ctx: context.Background()},
			wantErr: true,
			err:     errors.BookPriceEmptyError(),
		},
		{
			name: "book price too high",
			fields: fields{
				Name:  "Expensive Book",
				Price: 30000,
			},
			args:    args{ctx: context.Background()},
			wantErr: true,
			err:     errors.BookPriceTooHighError(),
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			ctx := transaction.InitializeTransaction(tt.args.ctx)
			b := &Book{
				ID:        tt.fields.ID,
				Name:      tt.fields.Name,
				Price:     tt.fields.Price,
				CreatedAt: tt.fields.CreatedAt,
			}
			err := b.Validate(ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("Book.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.wantErr {
				if ue, ok := err.(*errors.UserDefinedError); ok {
					if ue.ErrorCode != tt.err.ErrorCode || ue.ErrorMessage != tt.err.ErrorMessage {
						t.Errorf("Book.Validate() error = %v, expected error = %v", err, tt.err)
					}
				} else {
					t.Errorf("Book.Validate() error = %v, expected error = %v", err, tt.err)
				}
			}
		})
	}
}

func TestGetBooks(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	// モックDBとモックエクスペクテーションのセットアップ
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	// モックのタイムスタンプ
	mockTime := time.Now().Truncate(time.Millisecond)

	tests := []struct {
		name       string
		args       args
		mock       func()
		want       []Book
		wantErr    bool
		errMessage string
	}{
		{
			name: "正常に書籍を取得できる場合",
			args: args{ctx: context.Background()},
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "name", "price", "created_at"}).
					AddRow("1", "Book 1", 1000, mockTime).
					AddRow("2", "Book 2", 2000, mockTime)

				mock.ExpectQuery("SELECT id, name, price, created_at FROM books").
					WillReturnRows(rows)
			},
			want: []Book{
				{ID: "1", Name: "Book 1", Price: 1000, CreatedAt: mockTime},
				{ID: "2", Name: "Book 2", Price: 2000, CreatedAt: mockTime},
			},
			wantErr: false,
		},
		{
			name: "データベースからの取得に失敗する場合",
			args: args{ctx: context.Background()},
			mock: func() {
				mock.ExpectQuery("SELECT id, name, price, created_at FROM books").
					WillReturnError(sqlmock.ErrCancelled)
			},
			want:       nil,
			wantErr:    true,
			errMessage: "[500] [DB-ERR-500-01] データベースクエリの実行に失敗しました",
		},
		{
			name: "取得するデータがない場合",
			args: args{ctx: context.Background()},
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "name", "price", "created_at"})
				mock.ExpectQuery("SELECT id, name, price, created_at FROM books").
					WillReturnRows(rows)
			},
			want:    []Book{},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			got, err := GetBooks(tt.args.ctx, db)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetBooks() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr && err.Error() != tt.errMessage {
				t.Errorf("GetBooks() error = %v, wantErrMessage %v", err.Error(), tt.errMessage)
				return
			}
			if !tt.wantErr && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetBooks() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBook_CreateBook(t *testing.T) {
	type fields struct {
		ID        string
		Name      string
		Price     int
		CreatedAt time.Time
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		setupMock  func(sqlmock.Sqlmock)
		wantErr    bool
		errMessage string
	}{
		{
			name: "正常に書籍を作成できる場合",
			fields: fields{
				Name:  "Test Book",
				Price: 1000,
			},
			args: args{
				ctx: context.Background(),
			},
			setupMock: func(mock sqlmock.Sqlmock) {
				mock.ExpectPrepare("INSERT INTO books").
					ExpectExec().
					WithArgs("Test Book", 1000).
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
			wantErr: false,
		},
		{
			name: "SQLステートメントの準備に失敗する場合",
			fields: fields{
				Name:  "Test Book",
				Price: 1000,
			},
			args: args{
				ctx: context.Background(),
			},
			setupMock: func(mock sqlmock.Sqlmock) {
				mock.ExpectPrepare("INSERT INTO books").
					WillReturnError(sqlmock.ErrCancelled)
			},
			wantErr:    true,
			errMessage: "[500] [DB-ERR-500-04] SQLステートメントの準備に失敗しました",
		},
		{
			name: "データベースへの挿入に失敗する場合",
			fields: fields{
				Name:  "Test Book",
				Price: 1000,
			},
			args: args{
				ctx: context.Background(),
			},
			setupMock: func(mock sqlmock.Sqlmock) {
				mock.ExpectPrepare("INSERT INTO books").
					ExpectExec().
					WithArgs("Test Book", 1000).
					WillReturnError(sqlmock.ErrCancelled)
			},
			wantErr:    true,
			errMessage: "[500] [DB-ERR-500-05] データベースへの挿入に失敗しました",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			require.NoError(t, err)
			defer db.Close()

			tt.setupMock(mock)

			b := &Book{
				ID:        tt.fields.ID,
				Name:      tt.fields.Name,
				Price:     tt.fields.Price,
				CreatedAt: tt.fields.CreatedAt,
			}
			err = b.CreateBook(tt.args.ctx, db)
			if (err != nil) != tt.wantErr {
				t.Errorf("Book.CreateBook() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.wantErr && err.Error() != tt.errMessage {
				t.Errorf("Book.CreateBook() error = %v, wantErrMessage %v", err.Error(), tt.errMessage)
			}
		})
	}
}
