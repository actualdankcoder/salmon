<p align="center"><a href="https://github.com/actualdankcoder/salmon" target="_blank" rel="noopener noreferrer"><img width="1000" src="https://i.imgur.com/otKLIGl.png" alt="Salmon logo"></a></p>

<p align="center">
  <a href="https://github.com/actualdankcoder/salmon" title="Go to GitHub repo"><img src="https://img.shields.io/static/v1?label=actualdankcoder&message=salmon&color=blue&logo=github" alt="actualdankcoder - salmon"></a>
  <a href="https://github.com/actualdankcoder/salmon/releases/"><img src="https://img.shields.io/github/release/actualdankcoder/salmon?include_prereleases=&sort=semver&color=blue" alt="GitHub release"></a>
  <a href="#license"><img src="https://img.shields.io/badge/License-MIT-blue" alt="License"></a>
  <a href="https://github.com/actualdankcoder/salmon/issues"><img src="https://img.shields.io/github/issues/actualdankcoder/salmon" alt="issues - salmon"></a>
  <a href="https://discord.gg/2nRUM6auZM"><img src="https://img.shields.io/discord/932256432760426556?logo=discord"alt="chat on Discord"></a>
  <br/>
  Salmon is a music tracking server that connects with different sources to display a presence
</p>

## Dependencies:
+`git`	- To clone repo (optional).
+`go`	- Required the compile the binary (Required, Make Dependency).
+`make`	- Build system (Optional, Make Dependency).

## Installation

### Windows:
There are two ways to get salmon, you can either head to [releases](https://github.com/actualdankcoder/salmon/releases) and download the latest binary 

**or**

Compile the binary with the following commands:

```bash
git clone https://github.com/actualdankcoder/salmon
cd salmon
go mod tidy
go build
```
Once the binary is compiled, you can uninstall all dependencies and add salmon.exe to your PATH variable.

## Linux & MacOS:
Install Go and Make with your package manager and then run the following in the project directory:

+ `git clone https://github.com/actualdankcoder/salmon`
+ `cd salmon`
	+ `make` for a dynamically linked binary.
	+ `make static` for a statically linked binary.
+ `sudo make install`

Note: If you don't know the difference between dynamic & static linking then just use dynamic linker (`make`)

## Usage

*Ensure that salmon server is running before usage*

- Install the salmon plugin to your favourite music player
- Head to [http://localhost:8406](http://localhost:8406) to verify if music data is being fetched

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
[MIT](https://choosealicense.com/licenses/mit/)
