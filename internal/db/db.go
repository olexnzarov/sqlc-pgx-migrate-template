package db

import (
	"context"
	"embed"
	"errors"
	"fmt"
	"regexp"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/jackc/pgx/v5/stdlib"
)

//go:embed migrations/*.sql
var migrations embed.FS

type Database struct {
	*pgxpool.Pool
}

// Setup applies all migrations and creates a pgxpool.
// If the context is canceled, it will gracefully stop the migrations.
func Setup(ctx context.Context, url string) (*Database, error) {
	if err := applyMigrations(ctx, url); err != nil {
		return nil, err
	}
	return New(ctx, url)
}

func New(ctx context.Context, url string) (*Database, error) {
	pool, err := pgxpool.New(ctx, url)
	if err != nil {
		return nil, err
	}

	return &Database{
		Pool: pool,
	}, nil
}

func applyMigrations(ctx context.Context, url string) error {
	source, err := iofs.New(migrations, "migrations")
	if err != nil {
		return err
	}
	config, err := pgx.ParseConfig(url)
	if err != nil {
		return err
	}

	// pgx connection url might have runtime parameters that postgres won't accept,
	// like pool_max_conns or application_name
	for param := range config.RuntimeParams {
		pattern, err := regexp.Compile(fmt.Sprintf("%s=[a-zA-Z0-9]+&?", param))
		if err != nil {
			return err
		}
		url = pattern.ReplaceAllString(url, "")
	}

	m, err := migrate.NewWithSourceInstance("iofs", source, url)
	if err != nil {
		return err
	}

	doneChan := make(chan error)
	defer close(doneChan)
	go func() {
		doneChan <- m.Up()
	}()

	select {
	case err := <-doneChan:
		if err != nil {
			if errors.Is(err, migrate.ErrNoChange) {
				return nil
			}
			return err
		}
		return errors.Join(m.Close())
	case <-ctx.Done():
		m.GracefulStop <- true
		return errors.Join(
			fmt.Errorf("migrations were gracefully stopped: %w", ctx.Err()),
			errors.Join(m.Close()),
		)
	}
}
