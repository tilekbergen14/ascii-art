package functions

import (
	"os"
)

func FileCheck(filename string) bool {
	fileHashes := make(map[string]string)
	fileHashes["standard"] = "a51f800619146db0c42d26db3114c99f"
	fileHashes["shadow"] = "d44671e556d138171774efbababfc135"
	fileHashes["thinkertoy"] = "0021f26ad06f2f73a0cfa7b7d38d1434"
	data, err := os.ReadFile(filename + ".txt")
	if err != nil {
		return false
	}
	if !(fileHashes[filename] == GetMD5Hash(string(data))) {
		return false
	}
	return true
}
