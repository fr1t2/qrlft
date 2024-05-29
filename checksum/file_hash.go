// Package checksum computes checksums for large files
package checksum

import (
	"bufio"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"fmt"
	"hash"
	"hash/crc32"
	"io"
	"os"

	"golang.org/x/crypto/blake2s"
)

const bufferSize = 65536

// MD5sumReader returns MD5 checksum of content in reader
func MD5sumReader(reader io.Reader) (string, error) {
	return sumReader(md5.New(), reader)
}

// SHA256sumReader returns SHA256 checksum of content in reader
func SHA256sumReader(reader io.Reader) (string, error) {
	return sumReader(sha256.New(), reader)
}

// SHA1sumReader returns SHA1 checksum of content in reader
func SHA1sumReader(reader io.Reader) (string, error) {
	return sumReader(sha1.New(), reader)
}

// Blake2s256Reader returns SHA1 checksum of content in reader
func Blake2s256Reader(reader io.Reader) (string, error) {
	hash, _ := blake2s.New256([]byte{})
	return sumReader(hash, reader)
}

// CRCReader returns CRC-32-IEEE checksum of content in reader
func CRCReader(reader io.Reader) (string, error) {
	table := crc32.MakeTable(crc32.IEEE)
	checksum := crc32.Checksum([]byte(""), table)
	buf := make([]byte, bufferSize)
	for {
		switch n, err := reader.Read(buf); err {
		case nil:
			checksum = crc32.Update(checksum, table, buf[:n])
		case io.EOF:
			return fmt.Sprintf("%08x", checksum), nil
		default:
			return "", err
		}
	}
}

// sumReader calculates the hash based on a provided hash provider
func sumReader(hashAlgorithm hash.Hash, reader io.Reader) (string, error) {
	buf := make([]byte, bufferSize)
	for {
		switch n, err := reader.Read(buf); err {
		case nil:
			hashAlgorithm.Write(buf[:n])
		case io.EOF:
			return fmt.Sprintf("%x", hashAlgorithm.Sum(nil)), nil
		default:
			return "", err
		}
	}
}

// MD5sum returns MD5 checksum of filename
func MD5sum(filename string) (string, error) {
	return sum(md5.New(), filename)
}

// SHA256sum returns SHA256 checksum of filename
func SHA256sum(filename string) (string, error) {
	return sum(sha256.New(), filename)
}

// SHA1sum returns SHA1 checksum of filename
func SHA1sum(filename string) (string, error) {
	return sum(sha1.New(), filename)
}

// Blake2s256 returns BLAKE2s-256 checksum of filename
func Blake2s256(filename string) (string, error) {
	hash, _ := blake2s.New256([]byte{})
	return sum(hash, filename)
}

// CRC32 returns CRC-32-IEEE checksum of filename
func CRC32(filename string) (string, error) {
	if info, err := os.Stat(filename); err != nil || info.IsDir() {
		return "", err
	}

	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer func() { _ = file.Close() }()

	return CRCReader(bufio.NewReader(file))
}

// sum calculates the hash based on a provided hash provider
func sum(hashAlgorithm hash.Hash, filename string) (string, error) {
	if info, err := os.Stat(filename); err != nil || info.IsDir() {
		return "", err
	}

	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer func() { _ = file.Close() }()

	return sumReader(hashAlgorithm, bufio.NewReader(file))
}
