package log

import (
	"bytes"
	"log"
	"os"
	"strings"
	"testing"
)

func captureOutput(logger *log.Logger, f func()) string {
	var buf bytes.Buffer
	logger.SetOutput(&buf)
	f()
	logger.SetOutput(os.Stdout) // reset to default output
	return buf.String()
}

func assertLog(t *testing.T, logger *log.Logger, level string, funcUnderTest func(string, ...interface{}), msg string) {
	output := captureOutput(logger, func() {
		funcUnderTest(msg)
	})
	expected := "[" + level + "] " + msg
	if !strings.Contains(output, expected) {
		t.Errorf("Expected log message: %s, but got: %s", expected, output)
	}
}

func TestHappyLog(t *testing.T) {
	hlog := GetInstance("", "testLogger")
	hlog.SetLevel(TRACE)

	t.Run("testVar", func(t *testing.T) {
		output := captureOutput(hlog.logger, func() {
			hlog.Var("foo", 1)
		})
		expected := "[TRACE] var->foo=1"
		if !strings.Contains(output, expected) {
			t.Errorf("Expected log message: %s, but got: %s", expected, output)
		}
	})

	t.Run("testCritical", func(t *testing.T) {
		assertLog(t, hlog.logger, "CRITICAL", hlog.Critical, "critical info")
	})

	t.Run("testError", func(t *testing.T) {
		assertLog(t, hlog.logger, "ERROR", hlog.Error, "error info")
	})

	t.Run("testWarning", func(t *testing.T) {
		assertLog(t, hlog.logger, "WARNING", hlog.Warning, "warning info")
	})

	t.Run("testInfo", func(t *testing.T) {
		assertLog(t, hlog.logger, "INFO", hlog.Info, "info info")
	})

	t.Run("testDebug", func(t *testing.T) {
		assertLog(t, hlog.logger, "DEBUG", hlog.Debug, "debug info")
	})

	t.Run("testTrace", func(t *testing.T) {
		assertLog(t, hlog.logger, "TRACE", hlog.Trace, "trace info")
	})

	t.Run("testInput", func(t *testing.T) {
		output := captureOutput(hlog.logger, func() {
			hlog.Input("foo", 1)
		})
		expected := "[TRACE] input->foo=1"
		if !strings.Contains(output, expected) {
			t.Errorf("Expected log message: %s, but got: %s", expected, output)
		}
	})

	t.Run("testOutput", func(t *testing.T) {
		output := captureOutput(hlog.logger, func() {
			hlog.Output("foo", 1)
		})
		expected := "[TRACE] output->foo=1"
		if !strings.Contains(output, expected) {
			t.Errorf("Expected log message: %s, but got: %s", expected, output)
		}
	})

	t.Run("testEnterFunc", func(t *testing.T) {
		output := captureOutput(hlog.logger, func() {
			hlog.EnterFunc("testEnterFunc")
		})
		expected := "[TRACE] Enter function: testEnterFunc"
		if !strings.Contains(output, expected) {
			t.Errorf("Expected log message: %s, but got: %s", expected, output)
		}
	})

	t.Run("testExitFunc", func(t *testing.T) {
		output := captureOutput(hlog.logger, func() {
			hlog.ExitFunc("testExitFunc")
		})
		expected := "[TRACE] Exit function: testExitFunc"
		if !strings.Contains(output, expected) {
			t.Errorf("Expected log message: %s, but got: %s", expected, output)
		}
	})
}
