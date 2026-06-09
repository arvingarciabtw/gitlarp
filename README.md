<a id="readme-top"></a>

<!-- PROJECT LOGO -->

<div align="center">
  <img src="https://img.shields.io/badge/Go-%2300ADD8.svg?logo=go&logoColor=white" alt="Go">
</div>

<div align="center">
  <p align="center">
    A CLI tool to larp your Git contribution history
  </p>
</div>

<!-- TABLE OF CONTENTS -->
<details>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
    <li>
      <a href="#usage">Usage</a>
      <ul>
        <li><a href="#flags">Flags</a></li>
      </ul>
    </li>
    <li><a href="#license">License</a></li>
  </ol>
</details>

<!-- ABOUT THE PROJECT -->

## About The Project

The idea is simple: if you want to easily larp your Git contribution history for a certain date range, you can use this tool. Ideally used for private repos, as commit messages are repeated for each date.

_(I really just wanted to build a simple and fun CLI project to learn Go...)_

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- GETTING STARTED -->

## Getting Started

To get a local copy up and running follow these simple steps.

### Prerequisites

- [Go](https://go.dev/dl/)
- [Git](https://git-scm.com/install/)

### Installation

You can install the tool directly and run it without building the binary:

```sh
go install github.com/arvingarciabtw/gitlarp@latest

# Sample command
gitlarp -s 2026-01-01 -e 2026-01-31 -m "my commit message" -c 3
```

If you want to set it up locally:

```sh
# 1. Clone the repo
git clone https://github.com/arvingarciabtw/gitlarp.git
cd gitlarp

# 2. Build the binary (optional)
go build -o gitlarp

# 3. Run the tool
# For each day in the specified date range, it will
# make three commits with the specified commit message
./gitlarp -s 2026-01-01 -e 2026-01-31 -m "my commit message" -c 3
```

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- USAGE -->

## Usage

Intended usage is by using the tool directly. Refer to the installation section above. Available flags for the tool are below.

### Flags

| Flag       | Short | Description                        | Default      |
| ---------- | ----- | ---------------------------------- | ------------ |
| `-message` | `-m`  | Commit message                     | `larping...` |
| `-count`   | `-c`  | Number of commits per day (max 50) | `1`          |
| `-start`   | `-s`  | Start date (YYYY-MM-DD)            | today        |
| `-end`     | `-e`  | End date (YYYY-MM-DD)              | today        |

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- LICENSE -->

## License

Distributed under the MIT License. See `LICENSE` for more information.

<p align="right">(<a href="#readme-top">back to top</a>)</p>
