# linux_coreutils 项目规格说明书

## 1. 项目概述

**项目名称**：linux_coreutils

**项目目标**：用 Go 语言实现一组 Linux/Unix 常用命令的可移植版本，覆盖日常开发、运维和脚本场景中的高频工具。项目在 Linux 上提供与系统命令相近的行为，并通过 Go 的原生交叉编译能力保证在 Windows 和 macOS 上可编译运行。

**项目范围**：
- 实现 44 个命令（见第 3 节命令清单）。
- 不追求 100% 复刻 GNU Coreutils 的全部选项，优先覆盖 POSIX 子集和日常使用路径。
- 不实现需要完整解释器、协议栈、文件系统底层或 Linux 特有内核接口的命令。

## 2. 目标与范围

### 2.1 核心目标
- 提供一组自包含、单二进制、零外部依赖或低依赖的命令行工具。
- 源代码通过 Go 交叉编译在 Linux、Windows 和 macOS 上均可编译。
- 每个命令提供 `-h` / `--help` 帮助信息和基本 manual 页。
- 提供自动化测试套件，验证命令行为与系统命令的一致性。

### 2.2 不在范围内
- 完整实现 `sed`、`awk`、`diff`、`patch` 等需要解释器或复杂算法的工具。
- 实现依赖 `/proc`、`/sys`、`utmp`、`wtmp`、挂载、信号等 Linux 特有机制的系统命令（如 `ps`、`top`、`kill`、`df`、`mount`）。
- 实现 SSH、SCP 等需要完整协议栈的网络工具。
- 实现磁盘分区、文件系统检查和格式化工具。

## 3. 命令清单（44 个）

### 3.1 文件与目录操作（13 个）

| 命令 | 复杂度 | 说明 |
|---|---|---|
| `ls` | 中 | 列出目录内容，支持 `-l`、`-a`、`-h`、`-R` 等常用选项 |
| `cat` | 低 | 连接并输出文件内容，支持 `-n`、`-b` |
| `cp` | 低 | 复制文件/目录，支持 `-r`、`-v`、`-i` |
| `mv` | 低 | 移动/重命名文件，支持 `-i`、`-v` |
| `rm` | 低 | 删除文件/目录，支持 `-r`、`-f`、`-i` |
| `mkdir` | 低 | 创建目录，支持 `-p`、`-m` |
| `rmdir` | 低 | 删除空目录，支持 `-p` |
| `touch` | 低 | 创建空文件或更新时间戳 |
| `find` | 中 | 在目录树中查找文件，支持按名称、类型、大小过滤 |
| `which` | 低 | 在 PATH 中查找可执行文件 |
| `pwd` | 低 | 输出当前工作目录 |
| `realpath` | 低 | 输出规范化的绝对路径 |
| `basename` / `dirname` | 低 | 提取路径的基名或目录名 |

### 3.2 文本处理（13 个）

| 命令 | 复杂度 | 说明 |
|---|---|---|
| `grep` | 中 | 按模式搜索文本，支持基础正则、`-i`、`-v`、`-n`、`-r` |
| `cut` | 低 | 按字节/字符/字段切分行 |
| `sort` | 中 | 排序文本行，支持字典序、数值序、去重 |
| `uniq` | 低 | 过滤相邻重复行，支持计数 |
| `wc` | 低 | 统计行数、单词数、字节数 |
| `head` | 低 | 输出文件开头若干行/字节 |
| `tail` | 低 | 输出文件末尾若干行/字节（`-f` 可选） |
| `tr` | 低 | 字符转换/删除 |
| `tee` | 低 | 同时输出到标准输出和文件 |
| `xargs` | 中 | 从标准输入构建并执行命令 |
| `echo` | 低 | 输出字符串，支持 `-n`、转义序列 |
| `printf` | 低 | 格式化输出 |
| `paste` | 低 | 合并文件行 |
| `join` | 中 | 按公共字段连接两个文件 |

### 3.3 系统信息（3 个）

| 命令 | 复杂度 | 说明 |
|---|---|---|
| `uname` | 低 | 输出系统名称、版本、架构 |
| `hostname` | 低 | 输出主机名 |
| `whoami` | 低 | 输出当前用户名 |

### 3.4 网络（3 个）

| 命令 | 复杂度 | 说明 |
|---|---|---|
| `wget` | 中 | HTTP/HTTPS 文件下载（基础版） |
| `telnet` | 中 | TCP 连接测试（简化交互客户端） |
| `nc` | 中 | netcat：TCP/UDP 端口扫描、监听、转发 |

### 3.5 压缩 / 归档（6 个）

| 命令 | 复杂度 | 说明 |
|---|---|---|
| `tar` | 中 | 归档创建、解压、列表（基于 `archive/tar`，支持 ustar 格式） |
| `gzip` / `gunzip` | 中 | gzip 压缩/解压（基于 `compress/gzip`） |
| `zip` / `unzip` | 中 | zip 压缩/解压（基于 `archive/zip`） |
| `bzip2` | 中 | bzip2 压缩/解压（可引入纯 Go 实现，如 `github.com/dsnet/compress/bzip2`） |
| `xz` | 中 | xz 压缩/解压（可引入纯 Go 实现，如 `github.com/ulikunitz/xz`） |

### 3.6 其他实用工具（6 个）

| 命令 | 复杂度 | 说明 |
|---|---|---|
| `date` | 低 | 输出/格式化日期时间 |
| `cal` | 低 | 输出月历 |
| `sleep` | 低 | 暂停指定时间 |
| `timeout` | 中 | 在指定时间内运行命令 |
| `watch` | 中 | 定时重复执行命令并刷新输出 |
| `clear` / `reset` | 低 | 清屏/重置终端 |

## 4. 技术架构

### 4.1 编程语言
- **Go 1.24+**，使用标准库完成大部分功能。
- 复杂命令（如 `grep` 正则、`tar` 归档、`gzip` 解压）优先使用 Go 标准库或成熟纯 Go 库。
- 利用 Go 的交叉编译能力，通过 `GOOS` / `GOARCH` 生成 Linux、Windows、macOS 可执行文件。

### 4.2 构建系统
- 顶层使用 **Go Modules**（`go.mod` / `go.mod.sum`）管理依赖和模块版本。
- 使用 `go build ./cmd/<command>` 编译单个命令，使用 `go build ./cmd/...` 编译全部命令。
- 提供 `Makefile` 或 `Taskfile.yml` 包装常用操作：
  - `make build` 或 `go build ./cmd/...`：编译所有命令
  - `make install`：安装到 `PREFIX`（默认 `/usr/local/bin` 或 Windows 下指定目录）
  - `make test` 或 `go test ./...`：运行测试套件
  - `make clean`：清理构建产物
- 每个命令位于 `cmd/<command>/main.go`，生成独立可执行文件，便于单独测试和替换系统命令。

### 4.3 依赖策略
- **零依赖优先**：简单命令仅使用 Go 标准库。
- **允许可选依赖**：压缩/归档命令可依赖纯 Go 包，如 `github.com/klauspost/compress/gzip`、`github.com/klauspost/compress/zstd`（如需要）、`archive/zip`、`archive/tar`，构建时通过 Go Modules 版本控制。
- **公共包**：提取 `pkg/common` 用于错误处理、路径处理、跨平台抽象；提取 `pkg/cliutil` 用于统一命令行入口和帮助信息。

### 4.4 跨平台策略
| 平台 | 策略 |
|---|---|
| Linux | 直接使用 Go 标准库和 `golang.org/x/sys/unix` 等扩展包 |
| Windows | 使用 Go 标准库的 `os` / `path/filepath` / `os/exec` 封装，必要时使用 `golang.org/x/sys/windows` |
| macOS | 与 Linux 共享大部分代码，使用 `runtime.GOOS` 做少量差异处理 |

## 5. 目录结构

```text
linux_coreutils/
├── go.mod
├── go.sum
├── Makefile
├── README.md
├── SPEC.md
├── pkg/
│   ├── common/             # 公共工具函数
│   │   ├── errors.go       # 错误处理封装
│   │   ├── path.go         # 路径处理（跨平台）
│   │   └── io.go           # I/O 辅助函数
│   ├── cliutil/            # 命令行通用入口
│   │   ├── help.go         # -h / --help 统一输出
│   │   └── flags.go        # 通用标志解析辅助
│   ├── regex/              # grep 等命令的正则引擎封装
│   │   └── regex.go
│   └── tarfmt/             # tar 格式解析与生成
│       └── tarfmt.go
├── cmd/
│   ├── ls/
│   │   └── main.go
│   ├── cat/
│   │   └── main.go
│   ├── cp/
│   │   └── main.go
│   ├── mv/
│   │   └── main.go
│   ├── rm/
│   │   └── main.go
│   ├── mkdir/
│   │   └── main.go
│   ├── rmdir/
│   │   └── main.go
│   ├── touch/
│   │   └── main.go
│   ├── find/
│   │   └── main.go
│   ├── which/
│   │   └── main.go
│   ├── pwd/
│   │   └── main.go
│   ├── realpath/
│   │   └── main.go
│   ├── basename_dirname/
│   │   └── main.go
│   ├── grep/
│   │   └── main.go
│   ├── cut/
│   │   └── main.go
│   ├── sort/
│   │   └── main.go
│   ├── uniq/
│   │   └── main.go
│   ├── wc/
│   │   └── main.go
│   ├── head/
│   │   └── main.go
│   ├── tail/
│   │   └── main.go
│   ├── tr/
│   │   └── main.go
│   ├── tee/
│   │   └── main.go
│   ├── xargs/
│   │   └── main.go
│   ├── echo/
│   │   └── main.go
│   ├── printf/
│   │   └── main.go
│   ├── paste/
│   │   └── main.go
│   ├── join/
│   │   └── main.go
│   ├── uname/
│   │   └── main.go
│   ├── hostname/
│   │   └── main.go
│   ├── whoami/
│   │   └── main.go
│   ├── wget/
│   │   └── main.go
│   ├── telnet/
│   │   └── main.go
│   ├── nc/
│   │   └── main.go
│   ├── tar/
│   │   └── main.go
│   ├── gzip_gunzip/
│   │   └── main.go
│   ├── zip_unzip/
│   │   └── main.go
│   ├── bzip2/
│   │   └── main.go
│   ├── xz/
│   │   └── main.go
│   ├── date/
│   │   └── main.go
│   ├── cal/
│   │   └── main.go
│   ├── sleep/
│   │   └── main.go
│   ├── timeout/
│   │   └── main.go
│   ├── watch/
│   │   └── main.go
│   └── clear_reset/
│       └── main.go
├── internal/
│   └── testutil/           # 测试辅助包
│       ├── diff.go
│       └── tempdir.go
├── tests/
│   ├── common.sh           # 测试公共函数
│   ├── run_tests.sh        # 测试入口
│   └── cases/              # 每个命令的集成测试用例目录
├── docs/
│   └── man/                # manual 页源文件
└── scripts/
    └── build.sh            # CI 构建脚本（交叉编译矩阵）
```

## 6. 开发阶段

### 阶段一：基础设施（第 1-2 周）
- 初始化 Go Modules 工程：`go mod init github.com/user/linux_coreutils`。
- 实现公共包：
  - `pkg/common`：错误处理、路径处理、I/O 辅助函数
  - `pkg/cliutil`：统一 `-h` / `--help` 输出、退出码规范
- 实现测试框架：Go 单元测试 + shell 脚本集成测试。
- 确定命令行选项解析风格：每个命令使用标准库 `flag` 包，支持 POSIX 短选项和 GNU 长选项。

### 阶段二：文件与目录命令（第 3-5 周）
- 按优先级实现：`pwd`、`echo`、`printf`、`cat`、`mkdir`、`rmdir`、`touch`、`rm`、`cp`、`mv`、`ls`、`which`、`realpath`、`basename`/`dirname`、`find`。
- 每个命令完成后立即添加 `_test.go` 单元测试和 shell 集成测试。
- 重点验证跨平台路径行为（Windows 盘符、路径分隔符）。

### 阶段三：文本处理命令（第 6-9 周）
- 实现：`wc`、`head`、`tail`、`cut`、`tr`、`tee`、`uniq`、`sort`、`paste`、`join`、`xargs`、`grep`。
- `grep` 使用 Go 标准库 `regexp` 引擎，支持 `-i`、`-v`、`-n`、`-r`。
- `sort` 处理大文件时考虑流式读取或内存限制策略。

### 阶段四：系统信息与网络命令（第 10-12 周）
- 系统信息：`uname`、`hostname`、`whoami`。
- 网络：`wget`（HTTP/1.1 基础下载，基于 `net/http`）、`telnet`（TCP 连接，基于 `net`）、`nc`（监听/连接/端口扫描）。
- Windows 下网络命令使用 Go 标准库网络包，无需额外 WinSock2 适配。

### 阶段五：压缩与归档命令（第 13-16 周）
- 实现：`tar`（基于 `archive/tar`）、`gzip`/`gunzip`（基于 `compress/gzip`）、`zip`/`unzip`（基于 `archive/zip`）、`bzip2`、`xz`。
- `tar` 优先支持 ustar 格式，后续可扩展 GNU/pax 扩展。
- 压缩命令优先使用 Go 标准库或成熟纯 Go 库，保证正确性和跨平台性。

### 阶段六：其他实用工具（第 17-18 周）
- 实现：`date`、`cal`、`sleep`、`timeout`、`watch`、`clear`/`reset`。
- 终端相关命令在 Windows 下使用标准库和 `golang.org/x/sys/windows` 做控制台适配。

### 阶段七：集成测试与发布（第 19-20 周）
- 全量回归测试：`go test ./...` + shell 集成测试，修复与系统命令的差异。
- 完善 manual 页和 README。
- 提供预编译安装包（Linux、Windows、macOS 的 `.tar.gz` / `.zip`），通过 `GOOS` / `GOARCH` 交叉编译生成。
- 发布 v1.0。

## 7. 测试策略

### 7.1 单元测试
- 对公共包函数编写 Go 单元测试（`_test.go`）。
- 对复杂算法（如 `grep` 正则模式、`sort` 排序逻辑）编写独立测试。

### 7.2 集成测试
- 每个命令编写 shell 测试用例，与系统命令输出做 diff 对比。
- 覆盖常见选项组合和边界条件（空文件、大文件、特殊字符、目录递归）。

### 7.3 跨平台测试
- Linux：Ubuntu / Debian / Fedora 容器或虚拟机，运行 `go test ./...`。
- Windows：本地或 GitHub Actions Windows runner 直接运行二进制。
- macOS：验证交叉编译产物可运行。
- CI：GitHub Actions 配置多平台构建矩阵（`ubuntu-latest`、`windows-latest`、`macos-latest`）。

### 7.4 测试目录约定
```text
tests/cases/<command>/
├── test_01.sh
├── input.txt
├── expected.out
└── README.md
```

## 8. 兼容性要求

### 8.1 Linux
- 内核 3.10+
- Go 1.24+
- glibc 2.17+ 或 musl libc（静态链接可选）

### 8.2 Windows
- Windows 10+（或 Windows Server 2016+）
- 直接运行编译后的 `.exe`，无需 MinGW/MSYS2/Cygwin

### 8.3 macOS
- macOS 12+（ Monterey ）
- 同时支持 Intel 和 Apple Silicon 架构

### 8.4 行为兼容性
- 输出格式尽量贴近 GNU Coreutils，但不保证完全一致。
- 命令返回值遵循 POSIX 约定：成功返回 0，失败返回非 0。
- 错误信息使用英文，输出到标准错误流。

## 9. 发布标准

v1.0 发布前必须满足：
- [ ] 44 个命令全部实现并通过集成测试。
- [ ] Linux、Windows、macOS 平台均能编译通过。
- [ ] 所有命令提供 `-h` / `--help`。
- [ ] README 包含构建、安装、使用说明。
- [ ] CI 持续集成通过。
- [ ] 通过 `go test -race ./...` 检测数据竞争，无内存安全问题（依托 Go 垃圾回收器和内存安全特性）。

## 10. 时间线

| 阶段 | 内容 | 周期 |
|---|---|---|
| 一 | 基础设施 | 第 1-2 周 |
| 二 | 文件与目录命令 | 第 3-5 周 |
| 三 | 文本处理命令 | 第 6-9 周 |
| 四 | 系统信息与网络命令 | 第 10-12 周 |
| 五 | 压缩与归档命令 | 第 13-16 周 |
| 六 | 其他实用工具 | 第 17-18 周 |
| 七 | 集成测试与发布 | 第 19-20 周 |
| **合计** | | **20 周（约 5 个月）** |

## 11. 风险与依赖

| 风险 | 影响 | 缓解措施 |
|---|---|---|
| 跨平台路径/权限差异 | 中 | 统一使用 `path/filepath` 和 `os` 包处理路径和文件属性 |
| 压缩/归档格式兼容性 | 中 | 优先使用 Go 标准库，复杂格式引入成熟纯 Go 依赖 |
| Windows 控制台行为差异 | 中 | `watch`、`clear` 使用 `golang.org/x/sys/windows` 控制台 API |
| 测试环境配置复杂 | 低 | 提供 Docker/容器化测试环境 |
| 个别命令（如 `grep` 正则）复杂度超预期 | 中 | 先实现标准库 `regexp` 支持的子集，后续扩展 |
| Go 二进制体积较大 | 低 | 使用 `go build -ldflags="-s -w"` 裁剪符号，必要时使用 `upx` 压缩 |

---

**文档版本**：v1.0
**最后更新**：2026-06-23
