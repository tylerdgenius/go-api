package utils

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

type migFile struct {
	version int
	name    string
}

func RunMigrations(dsn string) error {
	abs, _ := filepath.Abs("./migrations")

	// 1) Collect & sort *.up.sql by numeric prefix
	var all []migFile
	entries, err := os.ReadDir(abs)
	if err != nil {
		return err
	}
	for _, e := range entries {
		if e.IsDir() {
			continue
		}
		n := e.Name()
		if !strings.HasSuffix(n, ".up.sql") {
			continue
		}
		prefix := strings.SplitN(n, "_", 2)[0]
		v, err := strconv.Atoi(prefix)
		if err != nil {
			continue
		}
		all = append(all, migFile{version: v, name: n})
	}
	sort.Slice(all, func(i, j int) bool { return all[i].version < all[j].version })

	// 2) Open migrate
	m, err := migrate.New("file://"+abs, dsn)
	if err != nil {
		return err
	}
	defer func() {
		_, _ = m.Close()
	}()

	// 3) Get current DB version
	curVersion := 0
	if v, dirty, err := m.Version(); err == nil {
		if dirty {
			return fmt.Errorf("database is in a dirty migration state at version %d", v)
		}
		curVersion = int(v)
	} else if err != migrate.ErrNilVersion {
		return err
	}
	log.Printf("Current DB version: %04d", curVersion)

	// 4) Filter to pending files (version > current)
	var pending []migFile
	for _, f := range all {
		if f.version > curVersion {
			pending = append(pending, f)
		}
	}
	if len(pending) == 0 {
		log.Println("No pending migrations.")
		return nil
	}

	log.Printf("Pending migrations (%d): %s", len(pending), joinNames(pending))

	// 5) Apply pending one-by-one and log each file
	for range pending {
		if err := m.Steps(1); err != nil {
			// On error, surface where we stopped
			if v, dirty, e2 := m.Version(); e2 == nil {
				return fmt.Errorf("failed applying migrations at version %04d (dirty=%v): %w", v, dirty, err)
			}
			return err
		}
		// After each step, the DB version advanced by 1 file
		if v, _, err := m.Version(); err == nil {
			// find name by version
			for _, f := range all {
				if f.version == int(v) {
					log.Printf("Applied migration: %s", f.name)
					break
				}
			}
		}
	}

	log.Println("Database migrations applied successfully.")
	return nil
}

func joinNames(files []migFile) string {
	names := make([]string, len(files))
	for i, f := range files {
		names[i] = f.name
	}
	return strings.Join(names, ", ")
}
