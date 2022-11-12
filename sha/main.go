package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"hash"
)

func main() {
	funcNameList := []string{"MD5", "SHA1", "SHA224", "SHA256", "SHA384", "SHA512"}
	funcMap := map[string]func(msg []byte) hash.Hash{
		"MD5":    func(msg []byte) hash.Hash { var h hash.Hash = md5.New(); h.Write(msg); return h },
		"SHA1":   func(msg []byte) hash.Hash { var h hash.Hash = sha1.New(); h.Write(msg); return h },
		"SHA224": func(msg []byte) hash.Hash { var h hash.Hash = sha256.New224(); h.Write(msg); return h },
		"SHA256": func(msg []byte) hash.Hash { var h hash.Hash = sha256.New(); h.Write(msg); return h },
		"SHA384": func(msg []byte) hash.Hash { var h hash.Hash = sha512.New384(); h.Write(msg); return h },
		"SHA512": func(msg []byte) hash.Hash { var h hash.Hash = sha512.New(); h.Write(msg); return h },
	}
	msg := []byte("Hello, world!")
	for _, funcName := range funcNameList {
		f := funcMap[funcName]
		s := f(msg).Sum(nil)

		fmt.Printf("%s(%v): %x\n", funcName, len(s), s)
	}
}
