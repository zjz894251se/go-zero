package logx

import (
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/zjz894251se/go-zero/core/fs"
)

func TestDailyRotateRuleMarkRotated(t *testing.T) {
	var rule DailyRotateRule
	rule.MarkRotated()
	assert.Equal(t, getNowDate(), rule.rotatedTime)
}

func TestDailyRotateRuleOutdatedFiles(t *testing.T) {
	var rule DailyRotateRule
	assert.Empty(t, rule.OutdatedFiles())
	rule.days = 1
	assert.Empty(t, rule.OutdatedFiles())
	rule.gzip = true
	assert.Empty(t, rule.OutdatedFiles())
}

func TestDailyRotateRuleShallRotate(t *testing.T) {
	var rule DailyRotateRule
	rule.rotatedTime = time.Now().Add(time.Hour * 24).Format(dateFormat)
	assert.True(t, rule.ShallRotate())
}

func TestRotateLoggerClose(t *testing.T) {
	filename, err := fs.TempFilenameWithText("foo")
	assert.Nil(t, err)
	if len(filename) > 0 {
		defer os.Remove(filename)
	}
	logger, err := NewLogger(filename, new(DailyRotateRule), false)
	assert.Nil(t, err)
	assert.Nil(t, logger.Close())
}

func TestRotateLoggerGetBackupFilename(t *testing.T) {
	filename, err := fs.TempFilenameWithText("foo")
	assert.Nil(t, err)
	if len(filename) > 0 {
		defer os.Remove(filename)
	}
	logger, err := NewLogger(filename, new(DailyRotateRule), false)
	assert.Nil(t, err)
	assert.True(t, len(logger.getBackupFilename()) > 0)
	logger.backup = ""
	assert.True(t, len(logger.getBackupFilename()) > 0)
}

func TestRotateLoggerMayCompressFile(t *testing.T) {
	filename, err := fs.TempFilenameWithText("foo")
	assert.Nil(t, err)
	if len(filename) > 0 {
		defer os.Remove(filename)
	}
	logger, err := NewLogger(filename, new(DailyRotateRule), false)
	assert.Nil(t, err)
	logger.maybeCompressFile(filename)
	_, err = os.Stat(filename)
	assert.Nil(t, err)
}

func TestRotateLoggerMayCompressFileTrue(t *testing.T) {
	filename, err := fs.TempFilenameWithText("foo")
	assert.Nil(t, err)
	logger, err := NewLogger(filename, new(DailyRotateRule), true)
	assert.Nil(t, err)
	if len(filename) > 0 {
		defer func() {
			os.Remove(filename)
			os.Remove(filepath.Base(logger.getBackupFilename()) + ".gz")
		}()
	}
	logger.maybeCompressFile(filename)
	_, err = os.Stat(filename)
	assert.NotNil(t, err)
}

func TestRotateLoggerRotate(t *testing.T) {
	filename, err := fs.TempFilenameWithText("foo")
	assert.Nil(t, err)
	logger, err := NewLogger(filename, new(DailyRotateRule), true)
	assert.Nil(t, err)
	if len(filename) > 0 {
		defer func() {
			os.Remove(filename)
			os.Remove(logger.getBackupFilename())
			os.Remove(filepath.Base(logger.getBackupFilename()) + ".gz")
		}()
	}
	err = logger.rotate()
	assert.Nil(t, err)
}

func TestRotateLoggerWrite(t *testing.T) {
	filename, err := fs.TempFilenameWithText("foo")
	assert.Nil(t, err)
	rule := new(DailyRotateRule)
	logger, err := NewLogger(filename, rule, true)
	assert.Nil(t, err)
	if len(filename) > 0 {
		defer func() {
			os.Remove(filename)
			os.Remove(logger.getBackupFilename())
			os.Remove(filepath.Base(logger.getBackupFilename()) + ".gz")
		}()
	}
	logger.write([]byte(`foo`))
	rule.rotatedTime = time.Now().Add(-time.Hour * 24).Format(dateFormat)
	logger.write([]byte(`bar`))
}
