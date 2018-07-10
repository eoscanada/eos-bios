package bios

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"

	"github.com/abourget/llerrgroup"
)

func (b *BIOS) DownloadReferences() error {
	if err := b.ensureCacheExists(); err != nil {
		return fmt.Errorf("error creating cache path: %s", err)
	}

	eg := llerrgroup.New(10)
	for _, contentRef := range b.BootSequence.Contents {
		if eg.Stop() {
			continue
		}

		contentRef := contentRef
		eg.Go(func() error {
			if err := b.DownloadURL(contentRef.URL, contentRef.Hash); err != nil {
				return fmt.Errorf("content %q: %s", contentRef.Name, err)
			}
			return nil
		})
	}
	if err := eg.Wait(); err != nil {
		return err
	}

	return nil
}

func (b *BIOS) ensureCacheExists() error {
	return os.MkdirAll(b.CachePath, 0777)
}

func (b *BIOS) DownloadURL(ref string, hash string) error {
	if b.isInCache(ref) {
		//net.Log.Printf("ipfs ref: %q in cache\n", ref)
		return nil
	}

	b.Log.Printf("Downloading and caching content from %q\n", ref)
	cnt, err := b.downloadURL(ref)
	if err != nil {
		return err
	}

	h := sha256.New()
	_, _ = h.Write(cnt)
	contentHash := hex.EncodeToString(h.Sum(nil))

	if contentHash != hash {
		return fmt.Errorf("hash in boot sequence [%q] not equal to computed hash on downloaded file [%q]", hash, contentHash)
	}

	if err := b.writeToCache(ref, cnt); err != nil {
		return err
	}

	b.Log.Printf("- %q done\n", ref)

	return nil
}

func (b *BIOS) downloadURL(destURL string) ([]byte, error) {
	req, err := http.NewRequest("GET", destURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, errors.New("download attempts failed")
	}
	defer resp.Body.Close()

	cnt, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode > 299 {
		if len(cnt) > 50 {
			cnt = cnt[:50]
		}
		return nil, fmt.Errorf("couldn't get %s, return code: %d, server error: %q", destURL, resp.StatusCode, cnt)
	}

	return cnt, nil
}

func (b *BIOS) writeToCache(ref string, content []byte) error {
	fileName := replaceAllWeirdities(ref)
	return ioutil.WriteFile(filepath.Join(b.CachePath, fileName), content, 0666)
}

func (b *BIOS) isInCache(ref string) bool {
	fileName := filepath.Join(b.CachePath, replaceAllWeirdities(ref))

	if _, err := os.Stat(fileName); err == nil {
		return true
	}
	return false
}

func (b *BIOS) ReadFromCache(ref string) ([]byte, error) {
	fileName := replaceAllWeirdities(ref)
	return ioutil.ReadFile(filepath.Join(b.CachePath, fileName))
}

func (b *BIOS) ReaderFromCache(ref string) (io.ReadCloser, error) {
	fileName := replaceAllWeirdities(ref)
	return os.Open(filepath.Join(b.CachePath, fileName))
}

func (b *BIOS) FileNameFromCache(ref string) string {
	fileName := replaceAllWeirdities(ref)
	return filepath.Join(b.CachePath, fileName)
}
