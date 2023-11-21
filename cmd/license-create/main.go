// Copyright 2018 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/drone/go-license/license"
	"github.com/drone/go-license/license/licenseutil"
)

var month = time.Hour * 24 * 31

var (
	in  = flag.String("in", "id_ed25519", "")
	out = flag.String("out", "license.txt", "")
	iss = flag.String("iss", "", "")
	cus = flag.String("cus", "", "")
	sub = flag.String("sub", "", "")
	typ = flag.String("typ", "", "")
	lim = flag.Int("lim", 0, "")
	exp = flag.Duration("exp", month, "")
	dat = flag.String("dat", "", "")
)

func main() {
	flag.Parse()

	privateKey, err := licenseutil.ReadPrivateKey(*in)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	now := "2023-09-11 00:00:00"
	iat, _ := time.Parse("2006-01-02 15:04:05", now)
	last := "2023-12-31 23:59:59"
	exp, _ := time.Parse("2006-01-02 15:04:05", last)
	license := &license.License{
		Iss: *iss,
		Cus: *cus,
		Sub: *sub,
		Typ: *typ,
		Lim: *lim,
		Exp: exp.UTC(),
		Iat: iat.UTC(),
		Dat: []byte(*dat),
	}

	encoded, err := license.Encode(privateKey)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = ioutil.WriteFile(*out, encoded, 0644)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
