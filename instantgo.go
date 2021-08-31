/*
<!--
Copyright (c) 2019 Christoph Berger. Some rights reserved.

Use of the text in this file is governed by a Creative Commons Attribution Non-Commercial
Share-Alike License that can be found in the LICENSE.txt file.

Use of the code in this file is governed by a BSD 3-clause license that can be found
in the LICENSE.txt file.

The source code contained in this file may import third-party source code
whose licenses are provided in the respective license files.
-->

<!--
NOTE: The comments in this file are NOT godoc compliant. This is not an oversight.

Comments and code in this file are used for describing and explaining a particular topic to the reader. While this file is a syntactically valid Go source file, its main purpose is to get converted into a blog article. The comments were created for learning and not for code documentation.
-->

+++
title = "Instant Go"
description = "Run Go code in the browser with Klipse and Yaegi"
author = "Christoph Berger"
email = "chris@appliedgo.net"
date = "2021-08-31"
draft = "true"
categories = [""]
tags = ["", "", ""]
articletypes = ["Tutorial"]
+++

Blog articles about Go are mostly static. Well, this might change, as authors now can include Go code that runs in the browser right away.

<!--more-->

## A Go playground without a server

Surely you know the Go playground already. It's been around since Go 1.0. The Go playground consists of a Web UI and a server component that runs the code entered in the UI and serves the output. A classic SPA Web App if you want. The obvious drawback for anybody who wants to run a similiar service is the cost (in terms of time and money) of setting up and running the backend.

What if there was a playground that runs entirely in the browser?

Well, now there is.


## Meet Klipse

There is a project named Klipse that provides browser-side REPLs (Read-Eval(uate)-Print Loops) for many languages since years. Each language needs a different approach though. And recently, the Klipse team announced support for Go.

This means there is now a workable server-less (yes, really, server-less, not just "serverless" in the sense of "running on servers that I don't have to care about") Go-playground-like solution for online documentation, blogs, tutorials, and so forth, to provide code that the reader can immediately run – and play with.


## How Klipse runs Go in the browser

Two words: Yaegi and WASM.

Yægi is a Go interpreter written and maintained by the Træfik team. (Yeah, they obviously have a faible for the "æ" ligature.) Yægi


## The code

{{< klipse_go >}}

*/

// ## Imports and globals
package main

import "fmt"

// Klipse automatically runs func main() whenever the source changes.
func main() {
	fmt.Println("Hulloh World")
}

/*
## How to get and run the code

Step 1: `go get` the code. Note the `-d` flag that prevents auto-installing
the binary into `$GOPATH/bin`.

    go get -d github.com/appliedgo/TODO:

Step 2: `cd` to the source code directory.

    cd $GOPATH/src/github.com/appliedgo/TODO:

Step 3. Run the binary.

    go run TODO:.go


## Odds and ends
## Some remarks
## Tips
## Links


**Happy coding!**

*/
