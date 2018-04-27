基于 EOS.IO 的启动工具
--------------------------------------------------------


`eos-bios` 是一个方便开发者快速启动 EOS 区块链网络的命令行工具.


该工具可用于以下目的:

* 启动 EOS 主网
* 启动 EOS 测试网络
* 搭建本地开发环境
* 搭建私有测试网络
* 基于网络发现协议，多节点协作启动 EOS 网络.

为了使用该工具，你需要发布一个 `discovery` 配置文件，以便节点之间互相发现。


启动本地单一节点
-----------------------------------------

[下载编译好的 `eos-bios` 可执行文件](https://github.com/eoscanada/eos-bios/releases),
clone 代码, 进入 `sample_config/docker` 文件夹，并执行如下命令：

    git clone https://github.com/eoscanada/eos-bios
    cd eos-bios/sample_config
    wget https://github.com/eoscanada/eos-bios/releases/download/......tar.gz  # 选择最新版本
    tar -zxvf eos-bios*tar.gz
    

然后运行：

    ./eos-bios boot

至此，本地节点启动完成。

这个节点提供了一个部署有全部系统合约的全功能的开发环境，其功能和即将于 6 月份上线的主网基本一致。

示例配置只会启动一个节点，不会和其他的 BP 节点组成网络（通过 `wingmen` 参数可以进行配置）。


启动本地网络
--------------

只需要简单更新配置，就可以搭建一个多节点网络。首先，执行 boot 命令:


    ./eos-bios boot

然后在另一台机器上执行:

    ./eos-bios join --verify

在 `boot` 命令执行完毕之后，你将看到一个 
_Kickstart Data_ 字符串。其他节点在执行 `join` 命令的时候，可以提供这串字符（目前该字符串是一个 base64 编码的 json），以便加入你的网络。

上述功能的实现，需要配合 `config.yaml` 和 `discovery` 文件的正确配置，并且需要将 `discovery` 文件，以链接的形式提供给其他想要加入的节点。


加入网络
------------------------

To join a network, tweak your discovery file to point to the network you're trying to join and publish it. Make sure other participants in the network link to your discovery file as their `wingmen`.

想要加入已有的网络，需要配置 `config.yaml` 和 `discovery` 文件，并将该文件的链接告知其他节点。`discovery` 中需要将其他节点配置在 `wingmen` 参数中。

* [config.yaml 示例配置](sample_config/config.yaml)
* [discovery 示例配置](https://github.com/eoscanada/network-discovery)

执行下面的命令:

    ./eos-bios join [--verify]

`--verify` 参数表示该节点在加入网络之后，需要验证 BIOS 节点的启动顺序。

多节点协作启动网络
------------------------------

当多个节点达成一致想要共同启动网络的时候, *所有节点* 只需执行下面这一命令:

    eos-bios orchestrate

根据特定算法和各个节点提供的 `discovery` 文件，各个节点将会被自动赋予相应的角色。

节点的角色可以分为下列三类：

1.BIOS Boot 节点，该节点将独自执行 `eos-bios boot` 命令，完成系统合约的部署，快照的导入，EOS 代币的分配等任务。
2.指定的区块生产节点（A-BP），这类节点将会执行 `eos-bios join --verify` 命令，加入 BIOS Boot 的网络，并负责验证 BIOS Boot 的行为。
3.其他节点，将执行 `eos-bios join` 命令，加入网络。

该工具提供了一系列 hooks 脚本，在执行 `boot`, `join` and `orchestrate` 命令的时候，这些脚本将被调用，所以请务必确保 hooks 脚本的准确，并且要多加练习。


安装 / 下载
------------------

你可以在这里下载该工具的最新版本:

[https://github.com/eoscanada/eos-bios/releases](https://github.com/eoscanada/eos-bios/releases)

该工具是一个单独的二进制文件，支持主流的操作系统。该工具没有其他依赖，只需被赋予可执行权限，即可运行。

当然，你也可以自行编译（Go 语言）：

    go get -v github.com/eoscanada/eos-bios/eos-bios

编译成功之后，可执行文件将默认安装到 `~/go/bin` 目录下。


加入讨论
-------------------

加入我们的 telegram 讨论组:
https://t.me/joinchat/GSUv1UaI5QIuifHZs8k_eA (EOSIO BIOS Boot channel)




TODO
----

* Shuffling of the top 5 for Boot selection
* Wait on Bitcoin Block
  * Add bitcoin_block_height in LaunchData
* In Orchestrate, compute the LaunchData by the most votes, weighted by the highest Weight
* Do we auto-publish the `my_discovery_file.yaml` ? Make it a hook?
* canonical_url ?
* convention regarding URLs, which pieces we want to see in there (the name of the organization, `testnet-[network_name])
