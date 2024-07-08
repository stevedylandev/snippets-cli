# Snippets.so CLI

![img](https://www.snippets.so/og.png)

✂︎ Clean and simple code sharing ✂︎

Quickly upload local code files and get links in return that you can view and share with others!

## Installation

### Mac OS
```
brew tap stevedylandev/snippets-cli
brew install snippets-cli
```

### Go
```
git clone https://github.com/stevedylandev/snippets-cli && cd snippets-cli && go install .
```

## Usage
Use the command `snip` followed by the name or path of the file you want to upload as a snippet.
```
snip hello.ts

//  https://snippets.so/snip/bafkreia2cv4bsi6xzc6den6b2zveffnmpbxlplwia2otsvp5fowrp7morm
```
After a file is upload you will get a link in return for the snippet on snippet.so.

If you want to customize the name of the snippet, you can use the `--name` flag.
```
snip --name for_steve hello.ts
```

## Questions

If you have any questions or requests feel free to [contact me!](mailto:hello@stevedylan.dev)
