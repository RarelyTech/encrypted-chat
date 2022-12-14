package ent

import (
    "context"
    "database/sql"
    "entgo.io/ent/dialect"
    entsql "entgo.io/ent/dialect/sql"
    "entgo.io/ent/entc/integration/ent/migrate"
    "fmt"
    "github.com/chatpuppy/puppychat/pkg/snag"
    log "github.com/sirupsen/logrus"
    "strings"

    _ "github.com/chatpuppy/puppychat/internal/ent/runtime"
    _ "github.com/jackc/pgx/v4/stdlib"
)

var Database *Client

func OpenDatabase(dsn string, debug bool) {

    // change timezone to UTC
    arr := strings.Split(dsn, " ")
    ht := false
    for i, s := range arr {
        if strings.HasPrefix(strings.ToLower(s), "timezone") {
            ht = true
            arr[i] = "TimeZone=UTC"
        }
    }
    if !ht {
        arr = append(arr, "TimeZone=UTC")
    }

    dsn = strings.Join(arr, " ")

    pgx, err := sql.Open("pgx", dsn)
    if err != nil {
        log.Fatalf("postgresql open failed: %v", err)
    }

    drv := entsql.OpenDB(dialect.Postgres, pgx)
    c := NewClient(Driver(drv))
    if debug {
        c = c.Debug()
    }

    autoMigrate(c)

    Database = c
}

func autoMigrate(c *Client) {
    ctx := context.Background()
    if err := c.Schema.Create(
        ctx,
        migrate.WithDropIndex(true),
        migrate.WithDropColumn(true),
        migrate.WithForeignKeys(false),
        // schema.WithDiffHook(func(next schema.Differ) schema.Differ {
        //     return schema.DiffFunc(func(current, desired *atlas.Schema) ([]atlas.Change, error) {
        //         changes, err := next.Diff(current, desired)
        //         if err != nil {
        //             return nil, err
        //         }
        //         changes = append([]atlas.Change{
        //             &atlas.AddSchema{
        //                 S: atlas.New("mytestschema").
        //                     AddAttrs(desired.Attrs...).
        //                     AddTables(desired.Tables...).
        //                     SetCharset("utf8"),
        //                 Extra: []atlas.Clause{&atlas.IfNotExists{}},
        //             },
        //         }, changes...)
        //         return changes, nil
        //     })
        // }),
    ); err != nil {
        log.Fatalf("database migrate failed: %v", err)
    }
}

func WithTx(ctx context.Context, fn func(tx *Tx) error) error {
    tx, err := Database.Tx(ctx)
    if err != nil {
        return err
    }
    defer func() {
        if v := recover(); v != nil {
            _ = tx.Rollback()
            panic(v)
        }
    }()
    if err = fn(tx); err != nil {
        if txErr := tx.Rollback(); txErr != nil {
            err = fmt.Errorf("rolling back transaction: %w", txErr)
        }
        return err
    }
    if err = tx.Commit(); err != nil {
        return fmt.Errorf("committing transaction: %w", err)
    }
    return nil
}

func WithTxPanic(ctx context.Context, fn func(tx *Tx) (err error)) {
    err := WithTx(ctx, fn)
    if err != nil {
        snag.Panic(err)
    }
}
