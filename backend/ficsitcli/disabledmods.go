package ficsitcli

import (
	"fmt"
	"log/slog"
	"os"
	"path/filepath"

	"github.com/satisfactorymodding/ficsit-cli/cli"
)

func modsDir(install *cli.Installation) string {
	return filepath.Join(install.BasePath(), "FactoryGame", "Mods")
}

func disabledModsDir(install *cli.Installation) string {
	return filepath.Join(install.BasePath(), "FactoryGame", "disabledMods")
}

func moveModsToDisabled(install *cli.Installation) error {
	return moveModsWithMarker(modsDir(install), disabledModsDir(install))
}

func restoreModsFromDisabled(install *cli.Installation) error {
	return moveModsWithMarker(disabledModsDir(install), modsDir(install))
}

func moveModsWithMarker(srcDir, dstDir string) error {
	entries, err := os.ReadDir(srcDir)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return fmt.Errorf("failed to read directory %s: %w", srcDir, err)
	}

	if err := os.MkdirAll(dstDir, 0755); err != nil {
		return fmt.Errorf("failed to create directory %s: %w", dstDir, err)
	}

	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}
		modDir := filepath.Join(srcDir, entry.Name())
		smmFile := filepath.Join(modDir, ".smm")
		if _, err := os.Stat(smmFile); err != nil {
			continue
		}
		dstPath := filepath.Join(dstDir, entry.Name())
		if err := os.Rename(modDir, dstPath); err != nil {
			slog.Warn("failed to move mod directory", slog.String("src", modDir), slog.String("dst", dstPath), slog.Any("error", err))
			continue
		}
		slog.Info("moved mod", slog.String("mod", entry.Name()), slog.String("from", srcDir), slog.String("to", dstDir))
	}

	return nil
}

func cleanDisabledMods(install *cli.Installation, activeMods map[string]cli.ProfileMod) error {
	disabledDir := disabledModsDir(install)
	entries, err := os.ReadDir(disabledDir)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return fmt.Errorf("failed to read disabled mods directory %s: %w", disabledDir, err)
	}

	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}
		modDir := filepath.Join(disabledDir, entry.Name())
		smmFile := filepath.Join(modDir, ".smm")
		if _, err := os.Stat(smmFile); err != nil {
			continue
		}
		if _, stillActive := activeMods[entry.Name()]; !stillActive {
			slog.Info("removing mod from disabledMods (no longer in profile)", slog.String("mod", entry.Name()))
			if err := os.RemoveAll(modDir); err != nil {
				return fmt.Errorf("failed to remove disabled mod %s: %w", entry.Name(), err)
			}
		}
	}

	return nil
}
