// Copyright 2019 shimingyah.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// ee the License for the specific language governing permissions and
// limitations under the License.

// +build linux

package ioengine

import (
	"os"
	"syscall"
)

const (
	// AlignSize size to align the buffer
	AlignSize = 512
	// BlockSize direct IO minimum number of bytes to write
	BlockSize = 4096
)

// OpenFileWithDIO open files with O_DIRECT flag
func OpenFileWithDIO(name string, flag int, perm os.FileMode) (*os.File, error) {
	return os.OpenFile(name, syscall.O_DIRECT|flag, perm)
}

// WriteAtv like linux pwritev, write to the specifies offset and dose not change the file offset.
func (dio *DirectIO) WriteAtv(bs [][]byte, off int64) (int, error) {
	return linuxWriteAtv(dio, bs, off)
}

// Append write data to the end of file.
func (dio *DirectIO) Append(bs [][]byte) (int, error) {
	return genericAppend(dio, bs)
}
