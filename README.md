# HN.HSBLHSN.ME

A [HackerNews](https://news.ycombinator.com/) reader written in Go and React. It focuses on content and readability.

## Demo

You can check the final result here at [https://hn.hsblhsn.me/](https://hn.hsblhsn.me/).


## Why?

I like Hacker News and I find it cool. But when it comes to readability and content I struggle a lot. I have to click
the link to read the content from another website, then I will be able to understand the comments. That does not go well
with me. And I found no Hacker News reader that preloads the content of the link. So I had to build it.

## Why load the third party content?

This app is completely web based. Initially I wanted to make a native browser-like app that preloads and prerenders the
content of the link. So that when I click the link, I get the article instantly. Unfortunately, I don't know anything
about developing a native app. I did not know frontend (react) either. But I knew backend. That's why most of the task
of this app is done in Go. So, I scraped the content of the link and displayed it here in the browser.

## What about the crappy content layouts?

Since the content is loaded from the third party site, I can't control the layout. So I had to make the best effort to
make it readable. I am constantly trying to prevent badly presented content from being displayed. Please open an issue
to report any broken content.

## Links

- [Development](docs/development.md).
- [Contributors](docs/contributors.md).
- [Security](docs/security.md).
- [License](docs/license.md).




