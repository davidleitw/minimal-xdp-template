# minimalXdpTemplate

`minimalXdpTemplate` is a project for learning xdp, which hopes to provide the simplest xdp program and the main.go that loads the xdp elf file into the kernel, so that users can concentrate on learning xdp as much as possible without dealing with tedious environment settings. This project is suitable for users who want to quickly understand the basics of xdp.


In addition, I have also written a Chinese blog introducing the concept of the whole project, see [基於 goebpf 搭建最小化的 xdp 實驗環境](https://davidleitw.github.io/posts/xdp_example_01/)。

### Install prerequisites

```
# Install clang/llvm to be able to compile C files into bpf arch
$ apt-get install clang llvm make
```

### Compile xdp program

```bash
$ cd bpf/
$ make
```

Instead of using the Makefile, you can choose to manually compile the xdp program using the following command:

```
$ cd bpf/
$ clang -I../includes -O2 -target bpf -c xdp_main.c -o xdp_main.o
```

This command uses the clang compiler to compile the `xdp_main.c` file, and outputs the compiled file to `xdp_main.o`. The `-I` option specifies the include path, `-O2` specifies the optimization level, and `-target` bpf specifies that the target is a BPF (Berkeley Packet Filter) program.

### Run it

After compiling the xdp program using the previous command, you can use the following command to load the xdp_main.o file into the kernel:

```
$ sudo go run main.go
```

This command uses the main.go program to load the `xdp_main.o` file into the kernel. In our example, the xdp_main.o file is loaded into the `lo` network interface, which allows for subsequent testing.

In order to execute the command to load the xdp_main.o file into the kernel, you must use the sudo command. This is because loading a program into the kernel requires elevated privileges, and the sudo command allows the user to temporarily escalate their privileges in order to perform the action. Without using sudo, the command will not be able to access the necessary resources and will fail to execute.

### Test it

After loading the xdp program into the kernel, you can use the `ping 127.0.0.1` command to test it. In our example, we will discard ICMP packets with even sequence numbers. As shown in the following figure, the test indeed achieves the expected effect, and all even packets are discarded among the 10 packets.

```
# Test the xdp program that was just loaded into the lo interface
$ ping 127.0.0.1
```

![](https://i.imgur.com/Up4bJRc.png)

