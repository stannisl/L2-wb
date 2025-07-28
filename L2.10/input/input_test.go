package input

import (
	"io"
	"os"
	"strings"
	"testing"
)

// Вспомогательная функция для создания временного файла с содержимым
func createTempFile(t *testing.T, content string) *os.File {
	t.Helper()

	tmpFile, err := os.CreateTemp("", "testfile")
	if err != nil {
		t.Fatalf("failed to create temp file: %v", err)
	}

	_, err = tmpFile.WriteString(content)
	if err != nil {
		t.Fatalf("failed to write to temp file: %v", err)
	}

	// Не забудь закрыть, чтобы потом открыть его снова
	if err := tmpFile.Close(); err != nil {
		t.Fatalf("failed to close temp file: %v", err)
	}

	return tmpFile
}

func TestReadFile(t *testing.T) {
	expected := []string{"line1", "line2", "line3"}
	content := strings.Join(expected, "\n")

	tmpFile := createTempFile(t, content)
	defer os.Remove(tmpFile.Name())

	lines, err := ReadFile(tmpFile.Name())
	if err != nil {
		t.Fatalf("ReadFile() error = %v", err)
	}

	if len(lines) != len(expected) {
		t.Errorf("Expected %d lines, got %d", len(expected), len(lines))
	}

	for i := range expected {
		if lines[i] != expected[i] {
			t.Errorf("Expected line %d = %q, got %q", i, expected[i], lines[i])
		}
	}
}
func TestReadStdin(t *testing.T) {
	expected := []string{"hello", "world"}
	input := strings.Join(expected, "\n")

	// Создаём pipe
	reader, writer, err := os.Pipe()
	if err != nil {
		t.Fatalf("failed to create pipe: %v", err)
	}

	// Сохраняем оригинальный stdin
	originalStdin := os.Stdin
	defer func() {
		os.Stdin = originalStdin
		reader.Close()
		writer.Close()
	}()

	// Пишем в pipe (это будет "вход" в stdin)
	go func() {
		defer writer.Close()
		io.WriteString(writer, input)
	}()

	// Подменяем stdin
	os.Stdin = reader

	lines, err := ReadStdin()
	if err != nil {
		t.Fatalf("ReadStdin() error = %v", err)
	}

	if len(lines) != len(expected) {
		t.Errorf("Expected %d lines, got %d", len(expected), len(lines))
	}

	for i := range expected {
		if lines[i] != expected[i] {
			t.Errorf("Expected line %d = %q, got %q", i, expected[i], lines[i])
		}
	}
}
