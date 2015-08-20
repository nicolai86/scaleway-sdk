// Copyright (C) 2015 Scaleway. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE.md file.

package commands

import (
	"bytes"
	"strings"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func ExampleRunImages() {
	ctx := ExampleCommandContext()
	args := ImagesArgs{}
	RunImages(ctx, args)
}

func ExampleRunImages_complex() {
	ctx := ExampleCommandContext()
	args := ImagesArgs{
		All:     false,
		NoTrunc: false,
		Quiet:   false,
	}
	RunImages(ctx, args)
}

func ExampleRunImages_quiet() {
	ctx := ExampleCommandContext()
	args := ImagesArgs{
		All:     false,
		NoTrunc: false,
		Quiet:   true,
	}
	RunImages(ctx, args)
}

func ExampleRunImages_all() {
	ctx := ExampleCommandContext()
	args := ImagesArgs{
		All:     true,
		NoTrunc: false,
		Quiet:   false,
	}
	RunImages(ctx, args)
}

func ExampleRunImages_notrunc() {
	ctx := ExampleCommandContext()
	args := ImagesArgs{
		All:     false,
		NoTrunc: true,
		Quiet:   false,
	}
	RunImages(ctx, args)
}

func TestRunImages_realAPI(t *testing.T) {
	ctx := RealAPIContext()
	if ctx == nil {
		t.Skip()
	}
	Convey("Testing RunImages() on real API", t, func() {
		Convey("no options", func() {
			args := ImagesArgs{
				All:     false,
				NoTrunc: false,
				Quiet:   false,
			}
			err := RunImages(*ctx, args)
			So(err, ShouldBeNil)

			stderr := ctx.Stderr.(*bytes.Buffer).String()
			So(stderr, ShouldBeEmpty)

			stdout := ctx.Stdout.(*bytes.Buffer).String()
			lines := strings.Split(stdout, "\n")
			So(len(lines), ShouldBeGreaterThan, 0)

			firstLine := lines[0]
			colNames := strings.Fields(firstLine)
			So(colNames, ShouldResemble, []string{"REPOSITORY", "TAG", "IMAGE", "ID", "CREATED", "VIRTUAL", "SIZE"})

			// FIXME: test public images
		})
		Convey("--all", func() {
			args := ImagesArgs{
				All:     true,
				NoTrunc: false,
				Quiet:   false,
			}
			err := RunImages(*ctx, args)
			So(err, ShouldBeNil)

			stderr := ctx.Stderr.(*bytes.Buffer).String()
			So(stderr, ShouldBeEmpty)

			stdout := ctx.Stdout.(*bytes.Buffer).String()
			lines := strings.Split(stdout, "\n")
			So(len(lines), ShouldBeGreaterThan, 0)

			firstLine := lines[0]
			colNames := strings.Fields(firstLine)
			So(colNames, ShouldResemble, []string{"REPOSITORY", "TAG", "IMAGE", "ID", "CREATED", "VIRTUAL", "SIZE"})

			// FIXME: test public images
			// FIXME: test bootscripts
			// FIXME: test snapshots
		})
		Convey("--quiet", func() {
			args := ImagesArgs{
				All:     false,
				NoTrunc: false,
				Quiet:   true,
			}
			err := RunImages(*ctx, args)
			So(err, ShouldBeNil)

			stderr := ctx.Stderr.(*bytes.Buffer).String()
			So(stderr, ShouldBeEmpty)

			stdout := ctx.Stdout.(*bytes.Buffer).String()
			lines := strings.Split(stdout, "\n")
			// So(len(lines), ShouldBeGreaterThan, 0)

			if len(lines) > 0 {
				firstLine := lines[0]
				colNames := strings.Fields(firstLine)
				So(colNames, ShouldNotResemble, []string{"REPOSITORY", "TAG", "IMAGE", "ID", "CREATED", "VIRTUAL", "SIZE"})

				// FIXME: test public images
				// FIXME: test bootscripts
				// FIXME: test snapshots
			}
		})
		Reset(func() {
			ctx.Stdout.(*bytes.Buffer).Reset()
			ctx.Stderr.(*bytes.Buffer).Reset()
		})
	})
}
