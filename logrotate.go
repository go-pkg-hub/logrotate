package logrotate

import (
    "os"
    "fmt"
    "sync"
)

type Logrotate struct {
    sync.Mutex

    file *os.File
    size int64

    maxSize int64
    maxFiles int
}

type Option func(*Logrotate)

func WithMaxSize(value int64) Option {
    return func(l *Logrotate) {
        if value > 0 {
            l.maxSize = value
        }
    }
}

func WithMaxFiles(value int) Option {
    return func(l *Logrotate) {
        if value > 0 {
            l.maxFiles = value
        }
    }
}

func New(logfile string, opts ...Option) (*Logrotate, error) {
    f, err := os.OpenFile(logfile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
    if err != nil {
        return nil, err
    }

    l := &Logrotate{file: f}

    // defaults
    l.maxSize = 0
    l.maxFiles = 1

    for _, opt := range opts {
        opt(l)
    }

    // rotate if needed
    if i, err := l.file.Stat(); err == nil {
        if l.maxSize > 0 && i.Size() > l.maxSize {
            if err := l.rotate(); err != nil {
                return nil, err
            }
        }
    }
    return l, nil
}

// Write implements io.Writer
func (l *Logrotate) Write(log []byte) (n int, err error) {
    l.Lock()
    defer l.Unlock()

    writeLen := int64(len(log))

    // rotate based on size
    if l.maxSize > 0 && l.size + writeLen > l.maxSize {
        if err := l.rotate(); err != nil {
            return 0, err
        }
    }

    n, err = l.file.Write(log)
    l.size += int64(n)

    return n, err
}

// Close implements io.Closer, and closes the current logfile
func (l *Logrotate) Close() error {
    l.Lock()
    defer l.Unlock()
    return l.close()
}

// close closes the file if it is open
func (l *Logrotate) close() error {
    if l.file == nil {
        return nil
    }
    err := l.file.Close()
    l.file = nil
    return err
}

// Rotate helper function for rotate
func (l *Logrotate) Rotate() error {
    l.Lock()
    defer l.Unlock()
    return l.rotate()
}

// rotate close existing log file and create a new one
func (l *Logrotate) rotate() error {
    if l.maxFiles > 1 {
        name := l.file.Name()
        l.close()
        maxFiles := l.maxFiles - 2
        // rotate logs
        for i := maxFiles; i >= 0; i-- {
            logfile := fmt.Sprintf("%s.%d", name, i)
            if _, err := os.Stat(logfile); err == nil {
                // delete old file
                if i == maxFiles {
                    os.Remove(logfile)
                } else if err := os.Rename(logfile, fmt.Sprintf("%s.%d", name, i + 1)); err != nil {
                    return err
                }
            }
        }
        // create logfile.log.0
        if err := os.Rename(name, fmt.Sprintf("%s.0", name)); err != nil {
            return err
        }
        // create new log file
        f, err := os.OpenFile(name, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
        if err != nil {
            return err
        }
        l.file = f
    } else {
        l.file.Truncate(0)
        l.file.Seek(0,0)
    }

    l.size = 0

    return nil
}
