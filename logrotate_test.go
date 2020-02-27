package logrotate

import (
    "os"
    "fmt"
    "log"
    "testing"
    "io/ioutil"
    "path/filepath"
)

func TestNew(t *testing.T) {
    tmpfile, err := ioutil.TempFile("", "TestNew")
    if err != nil {
        t.Error(err)
    }
    defer os.Remove(tmpfile.Name()) // clean up

    type Expected struct {
        maxSize int64
        maxFiles int
    }

    var testArgs = []struct {
        maxSize int64
        maxFiles int
        expected Expected
    }{
        {0, 0, Expected{0, 1}},
        {0, 1, Expected{0, 1}},
        {1, 0, Expected{1, 1}},
        {1, 1, Expected{1, 1}},
        {1, 2, Expected{1, 2}},
        {2, 0, Expected{2, 1}},
        {3, 1, Expected{3, 1}},
    }

    for _, a := range testArgs {
        l, err := New(tmpfile.Name(), WithMaxSize(a.maxSize), WithMaxFiles(a.maxFiles))
        if err != nil {
            t.Error(err)
        }
        if l.maxFiles != a.expected.maxFiles {
            t.Errorf("Expecting max-file %v, got: %v", a.expected.maxFiles, l.maxFiles)
        }
        if l.maxSize != a.expected.maxSize {
            t.Errorf("Expecting max-size %v, got: %v", a.expected.maxSize, l.maxSize)
        }
    }
}

func TestRotate(t *testing.T) {
    var testRotate = []struct {
        maxSize int64
        maxFiles int
        expected int
    }{
        {0, 0, 1},
        {0, 1, 1},
        {1, 0, 1},
        {0, 2, 1},
        {1, 2, 2},
        {1, 3, 3},
        {1024, 3, 1},
   }

    for _, a := range testRotate {
        dir, err := ioutil.TempDir("", "TestRotate")
        if err != nil {
            t.Error(err)
        }
        tmplog := filepath.Join(dir, "test.log")
        l, err := New(tmplog, WithMaxSize(a.maxSize), WithMaxFiles(a.maxFiles))
        if err != nil {
            t.Error(err)
        }
        log.SetOutput(l)
        for i := 0; i <= 5; i++ {
            log.Println(i)
        }
        files, err := ioutil.ReadDir(dir)
        if err != nil {
            t.Fatal(err)
        }
        if len(files) != a.expected {
            os.RemoveAll(dir)
            t.Fatalf("Expecting %v got %v", a.expected, len(files))
        }
        os.RemoveAll(dir)
    }
}

func TestRotateRotate(t *testing.T) {
    dir, err := ioutil.TempDir("", "TestRotateRotate")
    if err != nil {
        t.Fatal(err)
    }
    defer os.RemoveAll(dir)

    for i := 0; i <= 2; i++ {
        filesLen := i
        if filesLen < 1 {
            filesLen = 1
        }

        os.RemoveAll(dir)
        os.MkdirAll(dir, os.ModePerm)

        tmplog := filepath.Join(dir, fmt.Sprintf("test-%d.log", i))
        d1 := []byte("not\nempty\n")
        err = ioutil.WriteFile(tmplog, d1, 0644)
        if err != nil {
            t.Error(err)
        }
        l, err := New(tmplog, WithMaxSize(int64(i)), WithMaxFiles(i))
        if err != nil {
            t.Error(err)
        }

        l.Rotate()
        log.SetOutput(l)
        for i := 0; i <= 100; i++ {
            log.Println(i)
        }
        files, err := ioutil.ReadDir(dir)
        if err != nil {
            t.Fatal(err)
        }

        if len(files) != filesLen {
            t.Errorf("Expecting %v files got: %v", filesLen, len(files))
        }

        l.Rotate()
        files, err = ioutil.ReadDir(dir)
        if err != nil {
            t.Fatal(err)
        }
        if len(files) != filesLen {
            t.Errorf("Expecting %v files got: %v", filesLen, len(files))
        }
    }
}

func TestNewRotateSize(t *testing.T) {
    dir, err := ioutil.TempDir("", "TestRotateSize")
    if err != nil {
        t.Fatal(err)
    }
    defer os.RemoveAll(dir)
    tmplog := filepath.Join(dir, "test.log")
    d1 := []byte("not\nempty\n")
    err = ioutil.WriteFile(tmplog, d1, 0644)
    if err != nil {
        t.Error(err)
    }
    err = os.Truncate(tmplog, 1048577)
    if err != nil {
        fmt.Println(err)
    }
    _, err = New(tmplog, WithMaxSize(1048576), WithMaxFiles(2))
    if err != nil {
        t.Error(err)
    }
    files, err := ioutil.ReadDir(dir)
    if err != nil {
        t.Fatal(err)
    }
    if len(files) != 2 {
        t.Errorf("Expecting 2 files got: %v", len(files))
    }
}
