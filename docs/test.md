# Tests

## 1. Network IO

We can use the `iperf` tool to measure the network bandwidth.

To measure the network bandwidth, we need to run the following command on the server, bohr:

```bash
iperf -s -p 5200
```

Then, run the following command on the client, desktop:

```bash
iperf -c 192.168.100.32 -p 5200 --hide-ips
```

```bash
------------------------------------------------------------
Client connecting to (**hidden**), TCP port 5200
TCP window size: 16.0 KByte (default)
------------------------------------------------------------
[  1] local *.*.*.2 port 56454 connected with *.*.*.32 port 5200 (icwnd/mss/irtt=14/1448/1550)
[ ID] Interval       Transfer     Bandwidth
[  1] 0.0000-10.3744 sec  87.5 MBytes  70.8 Mbits/sec
```

The above output shows that the network bandwidth between my PC and the home server inside the WLAN is 70.8 Mbits/sec.

To run the test on the internet, we need to run the following command on the client:

```bash
iperf -c ping.online.net -p 5200 --hide-ips
```

```bash
------------------------------------------------------------
Client connecting to (**hidden**), TCP port 5200
TCP window size: 16.0 KByte (default)
------------------------------------------------------------
[  1] local *.*.*.2 port 41746 connected with *.*.*.21 port 5200 (icwnd/mss/irtt=13/1400/128092)
[ ID] Interval       Transfer     Bandwidth
[  1] 0.0000-11.3599 sec  26.5 MBytes  19.6 Mbits/sec
```

The above output shows that the network bandwidth between my PC and the server on the internet is 19.6 Mbits/sec.

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
 CPU Cores          : 16 @ 3241.014 MHz
 CPU Cache          : 512 KB
 AES-NI             : ✓ Enabled
 VM-x/AMD-V         : ✓ Enabled
 Total Disk         : 4.6 TB (1.4 TB Used)
 Total Mem          : 27.2 GB (16.6 GB Used)
 Total Swap         : 4.0 GB (3.0 GB Used)
 System uptime      : 1 days, 4 hour 21 min
 Load average       : 0.96, 0.93, 0.88
 OS                 : Arch Linux
 Arch               : x86_64 (64 Bit)
 Kernel             : 6.7.2-arch1-1
 TCP CC             : cubic
 Virtualization     : Dedicated
 IPv4/IPv6          : ✓ Online / ✗ Offline
 Organization       : AS33771 Safaricom Limited
 Location           : Nairobi / KE
 Region             : Nairobi Area
----------------------------------------------------------------------
 I/O Speed(1st run) : 1.2 GB/s
 I/O Speed(2nd run) : 1.3 GB/s
 I/O Speed(3rd run) : 1.3 GB/s
 I/O Speed(average) : 1297.1 MB/s
----------------------------------------------------------------------
 Node Name        Upload Speed      Download Speed      Latency
 Speedtest.net    19.20 Mbps        19.31 Mbps          12.99 ms
 Los Angeles, US  22.10 Mbps        22.02 Mbps          316.33 ms
 Dallas, US       21.19 Mbps        22.39 Mbps          291.47 ms
 Montreal, CA     20.69 Mbps        20.23 Mbps          224.89 ms
 Paris, FR        21.29 Mbps        22.34 Mbps          254.75 ms
 Amsterdam, NL    18.92 Mbps        22.88 Mbps          157.68 ms
 Shanghai, CN     21.90 Mbps        4.30 Mbps           406.80 ms
 Singapore, SG    19.97 Mbps        21.64 Mbps          213.06 ms
 Tokyo, JP        20.59 Mbps        21.79 Mbps          278.28 ms
----------------------------------------------------------------------
 Finished in        : 7 min 2 sec
 Timestamp          : 2024-02-01 22:38:00 EAT
----------------------------------------------------------------------
```

## nench for my PC

```bash
(wget -qO- wget.racing/nench.sh | bash; wget -qO- wget.racing/nench.sh | bash) 2>&1 | tee nench.log
```

```bash
-------------------------------------------------
 nench.sh v2019.07.20 -- https://git.io/nench.sh
 benchmark timestamp:    2024-02-01 19:38:37 UTC
-------------------------------------------------

Processor:    AMD Ryzen 7 7735HS with Radeon Graphics
CPU cores:    16
Frequency:    4093.903 MHz
RAM:          27Gi
Swap:         4.0Gi
Kernel:       Linux 6.7.2-arch1-1 x86_64

Disks:
nvme0n1  931.5G  SSD
zram0      4G  SSD

CPU: SHA256-hashing 500 MB
    0.315 seconds
CPU: bzip2-compressing 500 MB
    2.365 seconds
CPU: AES-encrypting 500 MB
    0.571 seconds

ioping: seek rate
    min/avg/max/mdev = 58.8 us / 73.4 us / 3.38 ms / 27.6 us
ioping: sequential read speed
    generated 20.4 k requests in 5.00 s, 4.97 GiB, 4.07 k iops, 1017.5 MiB/s

dd: sequential write speed
    1st run:    858.31 MiB/s
    2nd run:    1144.41 MiB/s
    3rd run:    1144.41 MiB/s
    average:    1049.04 MiB/s

IPv4 speedtests
    your IPv4:    105.163.158.xxxx

    Cachefly CDN:         2.35 MiB/s
    Leaseweb (NL):        0.00 MiB/s
    Softlayer DAL (US):   0.00 MiB/s
    Online.net (FR):      2.41 MiB/s
    OVH BHS (CA):         2.04 MiB/s

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
 CPU Cores          : 6 @ 2894.572 MHz
 CPU Cache          : 512 KB
 AES-NI             : ✓ Enabled
 VM-x/AMD-V         : ✗ Disabled
 Total Disk         : 490.9 GB (69.2 GB Used)
 Total Mem          : 15.6 GB (3.0 GB Used)
 Total Swap         : 977.0 MB (0 Used)
 System uptime      : 1 days, 4 hour 28 min
 Load average       : 0.34, 0.27, 0.32
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
 I/O Speed(1st run) : 190 MB/s
 I/O Speed(2nd run) : 200 MB/s
 I/O Speed(3rd run) : 110 MB/s
 I/O Speed(average) : 166.7 MB/s
----------------------------------------------------------------------
 Node Name        Upload Speed      Download Speed      Latency
 Speedtest.net    19.26 Mbps        19.29 Mbps          12.55 ms
 Los Angeles, US  21.78 Mbps        22.64 Mbps          347.51 ms
 Dallas, US       20.58 Mbps        22.16 Mbps          286.76 ms
 Montreal, CA     8.99 Mbps         11.58 Mbps          225.75 ms
 Paris, FR        21.35 Mbps        21.77 Mbps          192.05 ms
 Amsterdam, NL    20.32 Mbps        23.10 Mbps          153.59 ms
 Shanghai, CN     20.06 Mbps        17.74 Mbps          397.71 ms
 Hongkong, CN     19.93 Mbps        22.12 Mbps          275.27 ms
 Mumbai, IN       20.78 Mbps        21.18 Mbps          237.37 ms
 Singapore, SG    20.38 Mbps        21.77 Mbps          210.30 ms
 Tokyo, JP        20.71 Mbps        21.85 Mbps          272.67 ms
----------------------------------------------------------------------
 Finished in        : 7 min 35 sec
 Timestamp          : 2024-02-01 22:49:00 EAT
----------------------------------------------------------------------
```

## nench for bohr

```bash
(wget -qO- wget.racing/nench.sh | bash; wget -qO- wget.racing/nench.sh | bash) 2>&1 | tee nench.log
```

```bash
-------------------------------------------------
 nench.sh v2019.07.20 -- https://git.io/nench.sh
 benchmark timestamp:    2024-02-01 19:50:57 UTC
-------------------------------------------------

Processor:    QEMU Virtual CPU version 2.5+
CPU cores:    6
Frequency:    2894.572 MHz
RAM:          15Gi
bash: line 156: swapon: command not found
Swap:         -
Kernel:       Linux 6.1.0-17-amd64 x86_64

Disks:
sda    500G  HDD

CPU: SHA256-hashing 500 MB
    1.765 seconds
CPU: bzip2-compressing 500 MB
    3.521 seconds
CPU: AES-encrypting 500 MB
    0.653 seconds

ioping: seek rate
    min/avg/max/mdev = 35.7 us / 62.5 us / 851.5 us / 19.5 us
ioping: sequential read speed
    generated 31.9 k requests in 5.00 s, 7.79 GiB, 6.38 k iops, 1.56 GiB/s

dd: sequential write speed
    1st run:    85.64 MiB/s
    2nd run:    96.32 MiB/s
    3rd run:    69.24 MiB/s
    average:    83.73 MiB/s

IPv4 speedtests
    your IPv4:    105.163.158.xxxx

    Cachefly CDN:         2.46 MiB/s
    Leaseweb (NL):        0.04 MiB/s
    Softlayer DAL (US):   0.00 MiB/s
    Online.net (FR):      2.34 MiB/s
    OVH BHS (CA):         2.16 MiB/s

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
 CPU Cores          : 16 @ 2271.872 MHz
 CPU Cache          : 512 KB
 AES-NI             : ✓ Enabled
 VM-x/AMD-V         : ✓ Enabled
 Total Disk         : 0 (0 Used)
 Total Mem          : 30.8 GB (17.8 GB Used)
 System uptime      : 1 days, 4 hour 42 min
 Load average       : 0.43, 0.93, 0.69
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
 I/O Speed(1st run) : 2.1 GB/s
 I/O Speed(2nd run) : 2.1 GB/s
 I/O Speed(3rd run) : 2.1 GB/s
 I/O Speed(average) : 2150.4 MB/s
----------------------------------------------------------------------
 Node Name        Upload Speed      Download Speed      Latency
 Speedtest.net    19.31 Mbps        19.29 Mbps          12.30 ms
 Los Angeles, US  19.87 Mbps        22.55 Mbps          320.06 ms
 Dallas, US       20.40 Mbps        21.51 Mbps          287.91 ms
 Montreal, CA     20.79 Mbps        21.80 Mbps          219.82 ms
 Paris, FR        20.28 Mbps        21.72 Mbps          177.02 ms
 Amsterdam, NL    20.46 Mbps        22.88 Mbps          155.18 ms
 Shanghai, CN     22.10 Mbps        1.53 Mbps           397.51 ms
 Hongkong, CN     19.77 Mbps        22.60 Mbps          282.20 ms
 Mumbai, IN       19.71 Mbps        21.76 Mbps          239.42 ms
 Singapore, SG    20.64 Mbps        22.35 Mbps          208.81 ms
 Tokyo, JP        18.93 Mbps        22.67 Mbps          268.28 ms
----------------------------------------------------------------------
 Finished in        : 6 min 9 sec
 Timestamp          : 2024-02-01 23:00:16 EAT
----------------------------------------------------------------------
```

## nench for odin

```bash
(wget -qO- wget.racing/nench.sh | bash; wget -qO- wget.racing/nench.sh | bash) 2>&1 | tee nench.log
```

```bash
-------------------------------------------------
 nench.sh v2019.07.20 -- https://git.io/nench.sh
 benchmark timestamp:    2024-02-01 20:01:43 UTC
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
    1.734 seconds
CPU: bzip2-compressing 500 MB
    3.561 seconds
CPU: AES-encrypting 500 MB
    0.657 seconds

ioping: seek rate
    min/avg/max/mdev = 3.84 us / 6.63 us / 36.2 us / 539 ns
ioping: sequential read speed
    generated 155.5 k requests in 5.00 s, 38.0 GiB, 31.1 k iops, 7.59 GiB/s

dd: sequential write speed
    1st run:    1811.98 MiB/s
    2nd run:    1811.98 MiB/s
    3rd run:    1811.98 MiB/s
    average:    1811.98 MiB/s

IPv4 speedtests
    your IPv4:    105.163.158.xxxx

    Cachefly CDN:         2.45 MiB/s
    Leaseweb (NL):        0.04 MiB/s
    Softlayer DAL (US):   0.00 MiB/s
    Online.net (FR):      2.27 MiB/s
    OVH BHS (CA):         2.22 MiB/s

No IPv6 connectivity detected
-------------------------------------------------
```
