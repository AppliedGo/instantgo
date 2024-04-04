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
description = "Run Go code in the browser with Klipse and Yaegi. No backend required."
author = "Christoph Berger"
email = "chris@appliedgo.net"
date = "2021-09-04"
draft = "false"
categories = ["Go Ecosystem"]
tags = ["repl", "wasm", "browser", "yaegi"]
articletypes = ["Tools and Libraries"]
+++

Edit and run Go code right in the browser. No backend required.

<!--more-->

## A Go playground without a server

Surely you know the Go playground already. It's been around since Go 1.0. The Go playground consists of a Web UI and a server component that runs the code entered in the UI and serves the output. It's a classic Single-Page Application (SPA). The obvious drawback for anybody who wants to run a similar service is the cost (in terms of time, sweat, tears, and money) of setting up and running the backend.

What if there was a playground that runs entirely in the browser?

Well, now there is.


## Meet Klipse

There is a project named [Klipse](https://book.klipse.tech/) that provides browser-side REPLs (**R**ead-**E**val(uate)-**P**rint **L**oops) for many languages since years. Each language needs a different approach though. A few days ago, on August 29th, the Klipse team [announced support for Go](https://blog.klipse.tech/golang/2021/08/29/blog-go.html).

This means nothing less that there is now a server-less Go playground available for online documentation, blogs, tutorials, and so forth, to provide code that the reader can immediately run – and edit.

And yes, this is literally *server-less* in the sense of *"no backend required"*, and not the kind of "serverless" that commonly is an euphemism for *"the backend runs on servers that other people manage for me"*.


## How Klipse runs Go in the browser

Two words: Yægi and WASM.

[Yægi](https://github.com/traefik/yaegi) is a Go interpreter written and maintained by the Træfik team. (Yeah, they obviously have a weakness for names with an "æ" ligature.) Yægi supports the latest two Go versions, which is pretty neat, since typical Go blog article tend to report new things around Go that might not run on older Go versions.

Now Yægi usually runs in the terminal or as a Go library, but thanks to the [Web Assembly](https://en.wikipedia.org/wiki/WebAssembly) (or short: WASM) standard, this is no excuse for not running Yægi in the browser.

And that's pretty much all there is, although I can imagine that it took quite some effort to make all this work smoothly. So kudos to the Klipse team for their universal REPL framework, and to [Miguel Liezun](https://github.com/mliezun) who [provided the Go/Yægi integration](https://github.com/viebel/klipse/pull/393).


## What you can do with Klipse

Each Klipse snippet is a REPL that runs continuously. You can edit the code right in the browser, and as soon as the code compiles without errors, the eval loop runs and shows the result in the output pane. This is really neat!

Try for yourself with the below code which I borrowed from my [regexp article](https://appliedgo.net/regexp). (I stripped the comments for brevity, so feel free to visit the regexp article for more details.) Play with the regular expressions in `exps` and see how the output changes.


<!-- Klipse integration part 1 -->

<link rel="stylesheet" type="text/css" href="https://storage.googleapis.com/app.klipse.tech/css/codemirror.css">

<script>
	window.klipse_settings = {
		selector_golang: '.language-klipse-go',
        selector_eval_js: '.lang-eval-js',
	};
</script>

<!-- See also the script tag at the end of this document -->

<pre><code class="language-klipse-go">

package main

import (
	"fmt"
	"regexp"
	"runtime"
)

func prettyMatches(m []string) string {
	s := "["
	for i, e := range m {
		s += e
		if i < len(m)-1 {
			s += "|"
		}
	}
	s += "]"
	return s
}

func prettySubmatches(m [][]string) string {
	s := "[\n"
	for _, e := range m {
		s += "    " + prettyMatches(e) + "\n"
	}
	s += "]"
	return s
}

var (
	exps = []string{"b.*tter", "b(i|u)tter", `batter (\w+)`}

	text = `Betty Botter bought some butter
But she said the butter’s bitter
If I put it in my batter, it will make my batter bitter
But a bit of better butter will make my batter better
So ‘twas better Betty Botter bought a bit of better butter`
)

// Klipse automatically runs func main() whenever the source changes.

func main() {
	fmt.Println("Go version:", runtime.Version())
	for _, e := range exps {
		re := regexp.MustCompile(e)
		fmt.Println(e + ":")
		fmt.Println("1. FindString: ", re.FindString(text))
		fmt.Println("2. FindStringIndex: ", re.FindStringIndex(text))
		fmt.Println("3. FindStringSubmatch: ", re.FindStringSubmatch(text))
		fmt.Printf("4. FindAllString: %v\n", prettyMatches(re.FindAllString(text, -1)))
		fmt.Printf("5. FindAllStringIndex: %v\n", re.FindAllStringIndex(text, -1))
		fmt.Printf("6. FindAllStringSubmatch: %v\n\n", prettySubmatches(re.FindAllStringSubmatch(text, -1)))
	}
}

</code></pre>


## Code from a gist

Code can also be served from a GitHub gist, if code and blog text shall remain separate. The code window then shows the gist URL in an auto-generated comment.

Below is the code from my article about [futures in Go](https://appliedgo.net/futures/) (again, with comments stripped for brevity) that I extracted into a gist.

<pre><code class="language-klipse-go" data-gist-id="christophberger/2378d127b561c7f08332326cda205db8">
</code></pre>


## How to integrate the Klipse REPL in your Web page

The exact steps for setting up a Klipse Go REPL are spread across the Klipse [blog](https://blog.klipse.tech/golang/2021/08/29/blog-go.html) and [repo pages](https://github.com/viebel/klipse), so let me summarize here what I have done on this page.

### Step 1: import the scripts

Two HTML snippets are required for importing the Klipse/Yaegi REPL.

The first one goes into the header or in the body, but before the Go code snippets.

```html
<link rel="stylesheet" type="text/css" href="https://storage.googleapis.com/app.klipse.tech/css/codemirror.css">

<script>
    window.klipse_settings = {
        selector_golang: '.language-klipse-go',
    };
</script>
```

The second one should be added to the end of the page, before the closing `<body>` tag.

```html
<script src="https://storage.googleapis.com/app.klipse.tech/plugin_prod/js/klipse_plugin.min.js"></script>
```

### Step 2: add Go Code snippets

Each Go snippet that shall run in the browser must be placed inside a `<code>` block with class "language-klipse-go".

A package declaration is not needed. The code always runs inside package `main` and needs a `main` function as entry point. Each Go snippet runs its own Yægi interpreter, hence multiple Go snippets on the same page do not interfere with each other.

```html
<pre><code class="language-klipse-go">
import "fmt"

func main() {
  fmt.Println("Hello World!")
}
</code></pre>
```

For GitHub gists, the HTML snippet looks like this:

```html
<pre><code class="language-klipse-go" data-gist-id="christophberger/2378d127b561c7f08332326cda205db8">
</code></pre>
```

Easy enough, right?

## Limitations

Of course, it is not all sunshine and roses. There are a few limitations to consider.

- Go code snippets are isolated from each other. It is not possible to spread code across snippets and run them in a single `eval` loop, with a single output. That's somehow expected because each snippet runs its own Yægi instance.
- Right now, importing third-party libraries does not work.
- Yægi comes with a few [limitations](https://github.com/traefik/yaegi#Limitations), so don't expect 100.0% compatibility. For example, I tried to run the code from the [article about balanced trees](https://appliedgo.net/balancedtree/) but got an error from the `reflect` package. (No, the binary tree code does not use `reflect`.)
- A few features are still on the wish list, such as saving or exporting modified code.

Besides this, I found that the Klipse mechanisms do not play well with the way my blog generator works. I use a custom preprocessor for dividing code and comments to display them nicely side by side. This tears the code apart and makes it unusable for Klipse. I guess I'll find a solution for this, but for the time being, I'll have to add a single Klipse window to the end of the article to provide a runnable version of the code discussed in the article.


## Conclusion

With Klipse, running Go in the browser is almost effortless. A few snippets of HTML for loading scripts and CSS, and you're done. The available functionality is minimalistic but for the purpose of demonstrating things through small code snippets, it is certainly sufficient.

## Update 2024

Revisiting this article and the Klipse repository two and a half years later, I discovered that the Go version is still at Go 1.16 (yes, that's pre-generics), and there aren't any signs to keep the Go version up to date. (I added a `Println()` to the first code snippet to display the Go version, in case there are any changes.)

I therefore recommend using Klipse for Go only if the Go version does not matter.

If you are looking for a flexible, maintained alternative, take a look at [Codapi.org](https://codapi.org). It's server-based, but responses are swift, and the project is open source.

**Happy coding!**

<!-- Klipse integration part 2 -->

<script src="https://storage.googleapis.com/app.klipse.tech/plugin_prod/js/klipse_plugin.min.js"></script>

*/
