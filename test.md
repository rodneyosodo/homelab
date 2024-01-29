# Tests

## 1. Network IO

We can use the `iperf` tool to measure the network bandwidth.

To measure the network bandwidth, we need to run the following command on the server, bohr:

```bash
iperf -s -p 5200
```

Then, run the following command on the client, desktop:

```bash
iperf -c 192.168.100.32 -p 5200 --hide-ips --enhanced
```

```bash
------------------------------------------------------------
Client connecting to (**hidden**), TCP port 5200
TCP window size: 16.0 KByte (default)
------------------------------------------------------------
[  1] local *.*.*.2 port 42392 connected with *.*.*.32 port 5200 (icwnd/mss/irtt=14/1448/13675)
[ ID] Interval       Transfer     Bandwidth
[  1] 0.0000-10.6232 sec  55.3 MBytes  43.6 Mbits/sec
```

The above output shows that the network bandwidth between my PC and the home server inside the WLAN is 64.4 Mbits/sec.

To run the test on the internet, we need to run the following command on the client:

```bash
iperf -c ping.online.net -p 5200 --hide-ips --enhanced
```

```bash
------------------------------------------------------------
Client connecting to (**hidden**), TCP port 5200
TCP window size: 16.0 KByte (default)
------------------------------------------------------------
[  1] local *.*.*.2 port 54576 connected with *.*.*.21 port 5200 (icwnd/mss/irtt=13/1400/135758)
[ ID] Interval       Transfer     Bandwidth
[  1] 0.0000-13.5145 sec  16.1 MBytes  10.0 Mbits/sec
```

The above output shows that the network bandwidth between my PC and the server on the internet is 10.0 Mbits/sec.

## 2. CPU

We can use the `lscpu` tool to get the CPU model.

```bash
lscpu
```

```bash
Architecture:            x86_64
  CPU op-mode(s):        32-bit, 64-bit
  Address sizes:         48 bits physical, 48 bits virtual
  Byte Order:            Little Endian
CPU(s):                  16
  On-line CPU(s) list:   0-15
Vendor ID:               AuthenticAMD
  BIOS Vendor ID:        Advanced Micro Devices, Inc.
  Model name:            AMD Ryzen 7 4800H with Radeon Graphics
    BIOS Model name:     AMD Ryzen 7 4800H with Radeon Graphics          Unknown CPU @ 2.9GHz
    BIOS CPU family:     107
    CPU family:          23
    Model:               96
    Thread(s) per core:  2
    Core(s) per socket:  8
    Socket(s):           1
    Stepping:            1
    Frequency boost:     enabled
    CPU(s) scaling MHz:  93%
    CPU max MHz:         2900.0000
    CPU min MHz:         1400.0000
    BogoMIPS:            5789.17
    Flags:               fpu vme de pse tsc msr pae mce cx8 apic sep mtrr pge mca cmov pat pse36 clflush mmx fxsr sse sse2 ht syscall nx mmxext fxsr_opt pdpe1gb rdtscp lm constant_tsc r
                         ep_good nopl nonstop_tsc cpuid extd_apicid aperfmperf rapl pni pclmulqdq monitor ssse3 fma cx16 sse4_1 sse4_2 movbe popcnt aes xsave avx f16c rdrand lahf_lm cmp
                         _legacy svm extapic cr8_legacy abm sse4a misalignsse 3dnowprefetch osvw ibs skinit wdt tce topoext perfctr_core perfctr_nb bpext perfctr_llc mwaitx cpb cat_l3 c
                         dp_l3 hw_pstate ssbd mba ibrs ibpb stibp vmmcall fsgsbase bmi1 avx2 smep bmi2 cqm rdt_a rdseed adx smap clflushopt clwb sha_ni xsaveopt xsavec xgetbv1 cqm_llc c
                         qm_occup_llc cqm_mbm_total cqm_mbm_local clzero irperf xsaveerptr rdpru wbnoinvd cppc arat npt lbrv svm_lock nrip_save tsc_scale vmcb_clean flushbyasid decodeas
                         sists pausefilter pfthreshold avic v_vmsave_vmload vgif v_spec_ctrl umip rdpid overflow_recov succor smca
Virtualization features:
  Virtualization:        AMD-V
Caches (sum of all):
  L1d:                   256 KiB (8 instances)
  L1i:                   256 KiB (8 instances)
  L2:                    4 MiB (8 instances)
  L3:                    8 MiB (2 instances)
NUMA:
  NUMA node(s):          1
  NUMA node0 CPU(s):     0-15
```

## 3. Memory

We can use the `lsmem` tool to get the memory model.

```bash
lsmem
```

```bash
RANGE                                  SIZE  STATE REMOVABLE  BLOCK
0x0000000000000000-0x00000000cfffffff  3.3G online       yes   0-25
0x0000000100000000-0x000000080fffffff 28.3G online       yes 32-257

Memory block size:       128M
Total online memory:    31.5G
Total offline memory:      0B
```

## 4. Disk

We can use the `lsblk` tool to get the disk model.

```bash
lsblk
```

Results from odin:

```bash
NAME        MAJ:MIN RM   SIZE RO TYPE MOUNTPOINTS
sda           8:0    0   1.8T  0 disk
├─sda1        8:1    0   1.8T  0 part
└─sda9        8:9    0     8M  0 part
zd0         230:0    0     1M  0 disk
zd16        230:16   0   500G  0 disk
├─zd16p1    230:17   0   512M  0 part
├─zd16p2    230:18   0  27.9G  0 part
├─zd16p3    230:19   0   977M  0 part
└─zd16p4    230:20   0 470.6G  0 part
zd32        230:32   0     4M  0 disk
nvme0n1     259:0    0 953.9G  0 disk
├─nvme0n1p1 259:1    0  1007K  0 part
├─nvme0n1p2 259:2    0     1G  0 part
└─nvme0n1p3 259:3    0 952.9G  0 part
```

NVMe SSD: `nvme0n1` is used to install the operating system, proxmox and hold backups for the VMS locally. HDD: `sda` is used to hold the virtual machines.
Currently, I have not configured any RAID level for the HDD.

To test the disk performance, we can use the `dd` tool to write and read data from the disk.

```bash
dd if=/dev/zero of=/tmp/test bs=64k count=64k conv=fdatasync
```

Results from bohr:

```bash
65536+0 records in
65536+0 records out
4294967296 bytes (4.3 GB, 4.0 GiB) copied, 21.0751 s, 204 MB/s
```

### Benchmark utilities

#### bench.sh for my PC

```bash
wget -qO- bench.sh | bash
```

```bash
-------------------- A Bench.sh Script By Teddysun -------------------
 Version            : v2023-10-15
 Usage              : wget -qO- bench.sh | bash
----------------------------------------------------------------------
 CPU Model          : AMD Ryzen 7 7735HS with Radeon Graphics
 CPU Cores          : 16 @ 2092.532 MHz
 CPU Cache          : 512 KB
 AES-NI             : ✓ Enabled
 VM-x/AMD-V         : ✓ Enabled
 Total Disk         : 4.6 TB (1.4 TB Used)
 Total Mem          : 27.2 GB (14.6 GB Used)
 Total Swap         : 4.0 GB (1.6 GB Used)
 System uptime      : 2 days, 21 hour 32 min
 Load average       : 0.98, 0.69, 0.75
 OS                 : Arch Linux
 Arch               : x86_64 (64 Bit)
 Kernel             : 6.7.0-arch3-1
 TCP CC             : cubic
 Virtualization     : Dedicated
 IPv4/IPv6          : ✓ Online / ✗ Offline
 Organization       : AS33771 Safaricom Limited
 Location           : Nairobi / KE
 Region             : Nairobi Area
----------------------------------------------------------------------
 I/O Speed(1st run) : 1.3 GB/s
 I/O Speed(2nd run) : 1.3 GB/s
 I/O Speed(3rd run) : 1.3 GB/s
 I/O Speed(average) : 1331.2 MB/s
----------------------------------------------------------------------
 Node Name        Upload Speed      Download Speed      Latency
 Speedtest.net    10.24 Mbps        9.61 Mbps           12.93 ms
 Los Angeles, US  9.94 Mbps         11.32 Mbps          330.54 ms
 Dallas, US       4.56 Mbps         11.29 Mbps          322.77 ms
 Montreal, CA     3.44 Mbps         10.86 Mbps          270.39 ms
 Paris, FR        4.35 Mbps         10.83 Mbps          195.20 ms
 Amsterdam, NL    4.35 Mbps         10.88 Mbps          163.91 ms
 Shanghai, CN     6.23 Mbps         10.76 Mbps          430.61 ms
 Hongkong, CN     5.08 Mbps         11.02 Mbps          298.79 ms
 Mumbai, IN       7.94 Mbps         12.26 Mbps          286.12 ms
 Singapore, SG    5.61 Mbps         11.43 Mbps          223.43 ms
 Tokyo, JP        6.84 Mbps         10.84 Mbps          276.56 ms
----------------------------------------------------------------------
 Finished in        : 6 min 38 sec
 Timestamp          : 2024-01-29 20:49:39 EAT
----------------------------------------------------------------------
```

## nench for my PC

```bash
(wget -qO- wget.racing/nench.sh | bash; wget -qO- wget.racing/nench.sh | bash) 2>&1 | tee nench.log
```

```bash
-------------------------------------------------
 nench.sh v2019.07.20 -- https://git.io/nench.sh
 benchmark timestamp:    2024-01-29 17:54:53 UTC
-------------------------------------------------

Processor:    AMD Ryzen 7 7735HS with Radeon Graphics
CPU cores:    16
Frequency:    1993.591 MHz
RAM:          27Gi
Swap:         4.0Gi
Kernel:       Linux 6.7.0-arch3-1 x86_64

Disks:
nvme0n1  931.5G  SSD
zram0      4G  SSD

CPU: SHA256-hashing 500 MB
    0.311 seconds
CPU: bzip2-compressing 500 MB
    2.321 seconds
CPU: AES-encrypting 500 MB
    0.557 seconds

ioping: seek rate
    min/avg/max/mdev = 23.3 us / 71.4 us / 3.28 ms / 26.9 us
ioping: sequential read speed
    generated 21.8 k requests in 5.00 s, 5.33 GiB, 4.36 k iops, 1.07 GiB/s

dd: sequential write speed
    1st run:    1144.41 MiB/s
    2nd run:    1144.41 MiB/s
    3rd run:    1144.41 MiB/s
    average:    1144.41 MiB/s

IPv4 speedtests
    your IPv4:    105.163.2.xxxx

    Cachefly CDN:         1.11 MiB/s
    Leaseweb (NL):        0.02 MiB/s
    Softlayer DAL (US):   0.00 MiB/s
    Online.net (FR):      1.01 MiB/s
    OVH BHS (CA):         0.34 MiB/s

No IPv6 connectivity detected
-------------------------------------------------
```

## bench.sh for bohr

```bash
wget -qO- bench.sh | bash
```

```bash
-------------------- A Bench.sh Script By Teddysun -------------------
 Version            : v2023-10-15
 Usage              : wget -qO- bench.sh | bash
----------------------------------------------------------------------
 CPU Model          : QEMU Virtual CPU version 2.5+
 CPU Cores          : 6 @ 2894.578 MHz
 CPU Cache          : 512 KB
 AES-NI             : ✓ Enabled
 VM-x/AMD-V         : ✗ Disabled
 Total Disk         : 490.9 GB (30.3 GB Used)
 Total Mem          : 15.6 GB (2.5 GB Used)
 Total Swap         : 977.0 MB (0 Used)
 System uptime      : 1 days, 0 hour 40 min
 Load average       : 0.23, 0.14, 0.12
 OS                 : Debian GNU/Linux 12
 Arch               : x86_64 (64 Bit)
 Kernel             : 6.1.0-17-amd64
 TCP CC             :
 Virtualization     : KVM
 IPv4/IPv6          : ✓ Online / ✗ Offline
 Organization       : AS33771 Safaricom Limited
 Location           : Nairobi / KE
 Region             : Nairobi Area
----------------------------------------------------------------------
 I/O Speed(1st run) : 212 MB/s
 I/O Speed(2nd run) : 97.2 MB/s
 I/O Speed(3rd run) : 102 MB/s
 I/O Speed(average) : 137.1 MB/s
----------------------------------------------------------------------
 Node Name        Upload Speed      Download Speed      Latency
 Speedtest.net    9.62 Mbps         9.74 Mbps           12.63 ms
 Los Angeles, US  10.81 Mbps        11.15 Mbps          319.97 ms
 Dallas, US       10.81 Mbps        10.92 Mbps          351.30 ms
 Montreal, CA     10.33 Mbps        10.56 Mbps          235.70 ms
 Paris, FR        9.87 Mbps         9.65 Mbps           191.84 ms
 Amsterdam, NL    10.08 Mbps        9.51 Mbps           155.17 ms
 Hongkong, CN     10.03 Mbps        10.83 Mbps          273.09 ms
 Mumbai, IN       10.27 Mbps        10.90 Mbps          236.78 ms
 Singapore, SG    9.83 Mbps         11.20 Mbps          215.12 ms
 Tokyo, JP        9.92 Mbps         10.83 Mbps          268.81 ms
----------------------------------------------------------------------
 Finished in        : 5 min 45 sec
 Timestamp          : 2024-01-29 21:02:31 EAT
----------------------------------------------------------------------
```

## nench for bohr

```bash
(wget -qO- wget.racing/nench.sh | bash; wget -qO- wget.racing/nench.sh | bash) 2>&1 | tee nench.log
```

```bash
-------------------------------------------------
 nench.sh v2019.07.20 -- https://git.io/nench.sh
 benchmark timestamp:    2024-01-29 18:05:02 UTC
-------------------------------------------------

Processor:    QEMU Virtual CPU version 2.5+
CPU cores:    6
Frequency:    2894.578 MHz
RAM:          15Gi
bash: line 156: swapon: command not found
Swap:         -
Kernel:       Linux 6.1.0-17-amd64 x86_64

Disks:
sda    500G  HDD

CPU: SHA256-hashing 500 MB
    1.753 seconds
CPU: bzip2-compressing 500 MB
    3.515 seconds
CPU: AES-encrypting 500 MB
    0.641 seconds

ioping: seek rate
    min/avg/max/mdev = 36.9 us / 56.7 us / 753.8 us / 14.1 us
ioping: sequential read speed
    generated 30.0 k requests in 5.00 s, 7.32 GiB, 6.00 k iops, 1.46 GiB/s

dd: sequential write speed
    1st run:    161.17 MiB/s
    2nd run:    87.83 MiB/s
    3rd run:    159.26 MiB/s
    average:    136.09 MiB/s

IPv4 speedtests
    your IPv4:    105.163.2.xxxx

    Cachefly CDN:         1.25 MiB/s
    Leaseweb (NL):        0.03 MiB/s
    Softlayer DAL (US):   0.00 MiB/s
    Online.net (FR):      0.49 MiB/s
    OVH BHS (CA):         0.89 MiB/s

No IPv6 connectivity detected
-------------------------------------------------
```

## bench.sh for odin

```bash
wget -qO- bench.sh | bash
```

```bash
-------------------- A Bench.sh Script By Teddysun -------------------
 Version            : v2023-10-15
 Usage              : wget -qO- bench.sh | bash
----------------------------------------------------------------------
 CPU Model          : AMD Ryzen 7 4800H with Radeon Graphics
 CPU Cores          : 16 @ 1896.275 MHz
 CPU Cache          : 512 KB
 AES-NI             : ✓ Enabled
 VM-x/AMD-V         : ✓ Enabled
 Total Disk         : 0 (0 Used)
 Total Mem          : 30.8 GB (13.2 GB Used)
 System uptime      : 1 days, 0 hour 55 min
 Load average       : 0.48, 1.22, 1.00
 OS                 : Debian GNU/Linux 12
 Arch               : x86_64 (64 Bit)
 Kernel             : 6.5.11-7-pve
 TCP CC             : cubic
 Virtualization     : Dedicated
 IPv4/IPv6          : ✓ Online / ✗ Offline
 Organization       : AS33771 Safaricom Limited
 Location           : Nairobi / KE
 Region             : Nairobi Area
----------------------------------------------------------------------
 I/O Speed(1st run) : 2.0 GB/s
 I/O Speed(2nd run) : 2.0 GB/s
 I/O Speed(3rd run) : 2.0 GB/s
 I/O Speed(average) : 2048.0 MB/s
----------------------------------------------------------------------
 Node Name        Upload Speed      Download Speed      Latency
 Speedtest.net    9.59 Mbps         9.47 Mbps           12.37 ms
 Los Angeles, US  10.05 Mbps        11.37 Mbps          314.92 ms
 Dallas, US       10.31 Mbps        10.12 Mbps          299.67 ms
 Montreal, CA     10.51 Mbps        10.68 Mbps          238.69 ms
 Paris, FR        9.88 Mbps         11.47 Mbps          222.00 ms
 Amsterdam, NL    10.13 Mbps        11.52 Mbps          153.43 ms
 Shanghai, CN     10.31 Mbps        10.66 Mbps          379.71 ms
 Hongkong, CN     10.00 Mbps        10.64 Mbps          288.27 ms
 Mumbai, IN       10.13 Mbps        10.06 Mbps          240.09 ms
 Singapore, SG    9.98 Mbps         11.15 Mbps          210.09 ms
 Tokyo, JP        10.51 Mbps        10.91 Mbps          267.09 ms
----------------------------------------------------------------------
 Finished in        : 5 min 54 sec
 Timestamp          : 2024-01-29 21:17:09 EAT
----------------------------------------------------------------------
```

## nench for odin

```bash
(wget -qO- wget.racing/nench.sh | bash; wget -qO- wget.racing/nench.sh | bash) 2>&1 | tee nench.log
```

```bash
-------------------------------------------------
 nench.sh v2019.07.20 -- https://git.io/nench.sh
 benchmark timestamp:    2024-01-29 18:09:27 UTC
-------------------------------------------------

Processor:    AMD Ryzen 7 4800H with Radeon Graphics
CPU cores:    16
Frequency:    2900.000 MHz
RAM:          30Gi
Swap:         -
Kernel:       Linux 6.5.11-7-pve x86_64

Disks:
nvme0n1  953.9G  SSD
sda    1.8T  HDD
zd0    500G  SSD
zd16      4M  SSD
zd32      1M  SSD

CPU: SHA256-hashing 500 MB
    1.735 seconds
CPU: bzip2-compressing 500 MB
    3.431 seconds
CPU: AES-encrypting 500 MB
    0.657 seconds

ioping: seek rate
    min/avg/max/mdev = 3.84 us / 6.75 us / 295.7 us / 2.06 us
ioping: sequential read speed
    generated 157.9 k requests in 5.00 s, 38.5 GiB, 31.6 k iops, 7.71 GiB/s

dd: sequential write speed
    1st run:    1811.98 MiB/s
    2nd run:    1907.35 MiB/s
    3rd run:    1907.35 MiB/s
    average:    1875.56 MiB/s

IPv4 speedtests
    your IPv4:    105.163.2.xxxx

    Cachefly CDN:         1.26 MiB/s
    Leaseweb (NL):        0.03 MiB/s
    Softlayer DAL (US):   0.00 MiB/s
    Online.net (FR):      1.17 MiB/s
    OVH BHS (CA):         1.06 MiB/s

No IPv6 connectivity detected
-------------------------------------------------
```
