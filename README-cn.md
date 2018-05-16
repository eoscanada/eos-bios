基于EOS.IO的区块链网络启动工具
--------------------------------------------------------

#### v1.0 ZhaoYu from EosLaoMao
#### v1.1 Ian(Yahuang Wu) from MEET.ONE

[Click me switch to English version](./README.md)

`eos-bios` 是一个命令行工具，给那些想要启动基于EOS.IO的区块链网络的人而准备.

它提供以下功能:
* 主网分阶段启动
* 启动本地开发环境Booting local development environments
* 启动测试环境Booting testnets
* 启动联合网络或者私有网络

首先你需要知道的是**discovery protocol**. 点击这个链接可以看到介绍:

[![The Disco Dance](https://i.ytimg.com/vi/8aNZ_ZnKS-A/hqdefault.jpg?sqp=-oaymwEZCNACELwBSFXyq4qpAwsIARUAAIhCGAFwAQ==&rs=AOn4CLCZeGSGv9Qix8mHX77R4-d0rzDkgA)](https://youtu.be/8aNZ_ZnKS-A)

https://youtu.be/8aNZ_ZnKS-A

[点击查看示例配置](./sample_config)

下列 **视频** 解释了`eos-bios`的其他概念:

| 备注 | 链接 |
| ----------- | ----:|
| Details of the Discovery file  | https://youtu.be/uE5J7LCqcUc |
| Network meshing algorithm  | https://youtu.be/I-PrnlmLNnQ |
| The boot sequence | https://youtu.be/SbVzINnqWAE |
| Code review of the core flow  | https://youtu.be/p1B6jvx25O0 |
| The hooks to integrate with your infrastructure  | https://youtu.be/oZNV6fUoyqM |

`eos-bios`的一些问答视频:

| 备注 | 链接 |
| ----------- | ----:|
| Are accounts carried over from stage to stage?  | https://youtu.be/amyMm5gVpLg |
| How block producers agree on the content of the chain? | https://youtu.be/WIR7nab40qk |



### 本地开发环境
-----------------------------

[点击下载`eos-bios`](https://github.com/eoscanada/eos-bios/releases),
克隆当前仓库并且复制 `sample_config` 文件夹到新的目录

修改 `my_discovery_file.yaml` 文件,把target_http_address指向本地地址:

```
target_http_address: http://localhost:8888
```

然后运行:

    ./eos-bios boot --single

这个命令会在你本地运行一个完整的开发环境，包含所有的系统合约，跟你在主网启动时的环境很类似.

这个示例配置启动了一个单节点，同时它不会指向其他的出块候选节点.


### 社区协作启动网络
------------------------------

启动社区网络的时候，*每个人*都需要执行:

    ./eos-bios orchestrate

根据特定算法和各个节点提供的 discovery 文件，各个节点将会被自动赋予相应的角色:

1. _BIOS Boot node_, 该节点将单独执行 `eos-bios boot` 命令，完成系统合约的部署，快照的导入，EOS 代币的分配等任务.
2. _Appointed Block Producer_, 这类节点将会执行 `eos-bios join --verify` 命令，加入 BIOS Boot 的网络，并负责验证 BIOS Boot 的行为.
3. _other participant_, 将会执行`eos-bios join`加入网络.

以上执行完成以后，所有节点都需要等待`seed_network_launch_block`达成共识.


### 练习加入网络


让[种子网络](#seed-networks)里面的任何一个人邀请你，他们会执行:

    ./eos-bios invite [youraccount] [your pubkey]

这条命令会在种子网络中创建一个账号给你，修改你的`my_discovery_file.yaml`:

* `seed_network_account_name`, 需要跟你被邀请的账号一致.
* `seed_network_http_address`, 你想加入的种子网络的地址.

把你的私钥添加到`privkey.keys`, 因为需要运行发布命令:

    ./eos-bios publish

同时你也需要修改`my_discovery_file.yaml`中的这些值:

* `target_http_address` 你当前正在启动的节点HTTP地址.
* `target_p2p_address` 你当前正在启动的节点P2P地址.
* `target_appointed_block_producer_signing_key` and `target_initial_authority`: 这两个是跟你账户相关的授权信息.
* `seed_network_peers` [点击查看细节](#network-peers).

其他比较重要的字段:
* `seed_network_launch_block` 种子网络启动时的区块高度.
* `target_contents` 这里面配置的是我们已经达成共识可以写入链中的内容，比如系统合约，ERC-20快照等等. 



### 练习启动网络

检查你的`my_discovery_file.yaml`,确认没问题运行:

    ./eos-bios boot

这个会测试你的所有`hook_boot_*.sh`.


查看种子网络的数据
-----------------------

如果你的配置文件已经指向了一个种子网络，你可以运行:

    ./eos-bios discover

将会打印出节点的信息, 权重等.


Network peers
-------------

The `seed_network_peers` section of your discovery file looks like this:

```
seed_network_peers:
- account: eosexample
  comment: "They are good"
  weight: 10  # Weights are between 0 and 100 (INT value)
- account: eosmore
  comment: "They are better"
  weight: 20
```

This means you are comfortable launching the network with both
`eosexample` (at 10% vote weight), and `eosmore` (at 20%). `eos-bios`
will compute a graph of the network based on that peering information.

These are all account names on the seed network used to boot a new
network.


Seed networks
-------------

We keep an updated list of the different stages launched with `eos-bios` here:

https://stages.eoscanada.com



Example flow and interventions in the orchestrated launch
---------------------------------------------------------

1. Everyone runs `eos-bios orchestrate`.
1. `eos-bios` downloads the network topology pointed to by your `my_discovery_file.yaml`, as does everyone.
1. The network topology is sorted by weight according to how people voted in their `peers` section.
1. The `launch_ethereum_block` is taken from the top 20% in the topology: if they all agree, with continue with that number. Otherwise, we wait until they do (and periodically retraverse the network graph)




Install / Download
------------------

You can download the latest release here:
https://github.com/eoscanada/eos-bios/releases .. it is a single
binary that you can download on all major platforms. Simply make
executable and run. It has zero dependencies.

Alternatively, you can build from source with:

    go get -v github.com/eoscanada/eos-bios/eos-bios

This will install the binary in `~/go/bin` provided you have the Go
tool installed (quick install at https://golang.org/dl)



Join the discussion
-------------------

On Telegram through this invite link:
https://t.me/joinchat/GSUv1UaI5QIuifHZs8k_eA (EOSIO BIOS Boot channel)



Previous proposition
--------------------

See the previous proposition in this repo in README.v0.md



Readiness checklist
-------------------

* Did I update my `target_p2p_address` to reflect the IP of the NEW network we're booting ?
* Did I update my `target_http_address` to point to my node, reachable from `eos-bios`'s machine ?
*


Troubleshooting
---------------

* Do your `PRIVKEY` and `PUBKEY` in `hook_join_network.sh` match what
  you published in your discovery file under
  `target_initial_authority` and `target_initial_block_signing_key` ?

* Forked ? Did someone point to an old p2p address ? If so, removed
  them from the network.



TODO
----

* In Orchestrate, compute the LaunchData by the most votes, weighted by the highest Weight

  * Find out what we do for the chain_id.. do we vote for it too ?
    Top 20% must agree on the chain_id ?
    Top 20% must agree on the constitution ?

* boot_connect_mesh: Make sure we don't mesh with the first BIOS boot..
  it's most probably not running..

* Do connectivity checks when doing `discovery`.. and get a report upon orchestration
  that the peers are up ?

* Implement `eos-bios boot --reset` or something.. through eosio.disco::delgenesis

* create the "RAM" currency, issue an initial base ? is it with `setram` ?
* call `setram`, agree on it.. start with 32GB ?

* delegatebw, from/to eosio, do the transfer with it ?

* undelegatebw never removes my "voters" entry.. sunk forever ?



Role       Seed Account  Target Acct   Weight  Contents          Launch block (local time)
----       ------------  -----------   ------  ----------------  ------------
BIOS NODE  eosmama       eoscanadacom  10      1112211           500 (Fri Nov 7th 23:36, local time)
ABP 01     eosmarc       eosmarc       5       1111112           572 (Fri Nov 8th 00:25, local time)
ABP 02     eosrita       eosrita       2       1111111           572 (Fri Nov 8th 00:25, local time)
ABP 03     eosguy        eosguy        1       1111111           572 (Fri Nov 8th 00:25, local time)
ABP 04     eosbob        eosbob        1

Contents disagreements:
* About column 4: `boot_sequence.yaml`
  * eosmarc, eoscanadacom, eospouet says 1: /ipfs/Qmakjdsflakjdslfkjaldsfk
  * eosmama says 2: /ipfs/Qmhellkajdlakjdsflkj
