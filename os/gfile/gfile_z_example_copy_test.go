// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package gfile_test

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/test/gtest"
)

func ExampleCopy() {
	// init
	var (
		srcFileName = "gflie_example.txt"
		srcTempDir  = gfile.Temp("gfile_example_copy_src")
		srcTempFile = gfile.Join(srcTempDir, srcFileName)

		// copy file
		dstFileName = "gflie_example_copy.txt"
		dstTempFile = gfile.Join(srcTempDir, dstFileName)

		// copy dir
		dstTempDir = gfile.Temp("gfile_example_copy_dst")
	)

	// write contents
	gfile.PutContents(srcTempFile, "goframe example copy")

	// copy file
	gfile.Copy(srcTempFile, dstTempFile)

	// read contents after copy file
	fmt.Println(gfile.GetContents(dstTempFile))

	// copy dir
	gfile.Copy(srcTempDir, dstTempDir)

	// list copy dir file
	fList, _ := gfile.ScanDir(dstTempDir, "*", false)
	for _, v := range fList {
		fmt.Println(gfile.Basename(v))
	}

	// copy with context cancel
	ctx, copyCancel := context.WithCancel(context.Background())
	srcFile, err := gfile.Open(gtest.DataPath("dir1\\file1"))
	if err != nil {
		log.Fatalln(err)
	}
	dstFile, err := gfile.Create(gtest.DataPath("dir1\\file2"))
	if err != nil {
		log.Fatalln(err)
	}
	go func() {
		time.Sleep(1 * time.Millisecond)
		copyCancel()
	}()
	gfile.CtxCopy(ctx, dstFile, srcFile)

	// Output:
	// goframe example copy
	// gflie_example.txt
	// gflie_example_copy.txt
}
