
package storage

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "os"
    "path/filepath"
    "time"

    "github.com/gofrs/flock"
)

func ensureDir(path string) error {
    dir := filepath.Dir(path)
    return os.MkdirAll(dir, 0o755)
}

func EnsureDir(path string) error {
    dir := filepath.Dir(path)
    return os.MkdirAll(dir,0o755)
}

func atomicWriteFile(path string, data []byte) error {
    if err := ensureDir(path); err != nil {
        return err
    }
    tmp := path + ".tmp"
    if err := ioutil.WriteFile(tmp, data, 0o644); err != nil {
        return err
    }
    return os.Rename(tmp, path)
}

func SaveJSON(path string, v interface{}) error {
    b, err := json.MarshalIndent(v, "", "  ")
    if err != nil {
        return err
    }
    return atomicWriteFile(path, b)
}


func LoadJSON(path string, v interface{}) error {
    _, err := os.Stat(path)
    if os.IsNotExist(err) {
        // caller may handle empty file case
        return fmt.Errorf("file not found")
    }
    b, err := ioutil.ReadFile(path)
    if err != nil {
        return err
    }
    return json.Unmarshal(b, v)
}



// Simple file lock helper. Returns a locked flock and nil error on success.
func LockFile(path string, timeout  time.Duration) (*flock.Flock, error) {
    lockPath := path + ".lock"
    f := flock.New(lockPath)
    locked, err := f.TryLock()
    if err != nil {
        return nil, err
    }
    if !locked {
        return f, nil
    }

    deadline := time.Now().Add(timeout)
    for time.Now().Before(deadline){
    	time.Sleep(100 * time.Millisecond)
    	locked, err = f.TryLock()
    	if err != nil {
    		return nil, err
    	}
    	if locked {
    		return f, nil
    	}
    }

    return nil, fmt.Errorf("Could not acquire lock for %s within %s", path,timeout)
}