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

While using ethernet cable connected to the router:

```bash
------------------------------------------------------------
Client connecting to (**hidden**), TCP port 5200
TCP window size: 16.0 KByte (default)
------------------------------------------------------------
[  1] local *.*.*.79 port 55392 connected with *.*.*.85 port 5200
[ ID] Interval       Transfer     Bandwidth
[  1] 0.0000-10.0420 sec  1.03 GBytes   882 Mbits/sec
```

The above output shows that the network bandwidth between my PC and the home server inside the LAN is 882 Mbits/sec.

To run the test on the internet, we need to run the following command on the client:

```bash
iperf -c ping.online.net -p 5200 --hide-ips
```

```bash
------------------------------------------------------------
Client connecting to (**hidden**), TCP port 5200
TCP window size: 16.0 KByte (default)
------------------------------------------------------------
[  1] local *.*.*.14 port 50456 connected with *.*.*.21 port 5200
[ ID] Interval       Transfer     Bandwidth
[  1] 0.0000-9.6136 sec  35.6 MBytes  31.1 Mbits/sec
```

The above output shows that the network bandwidth between my PC and the server on the internet is 19.6 Mbits/sec.

## 2. CPU

We can use the `lscpu` tool to get the CPU model.

```bash
lscpu
```

```bash
Architecture:             x86_64
  CPU op-mode(s):         32-bit, 64-bit
  Address sizes:          48 bits physical, 48 bits virtual
  Byte Order:             Little Endian
CPU(s):                   16
  On-line CPU(s) list:    0-15
Vendor ID:                AuthenticAMD
  BIOS Vendor ID:         Advanced Micro Devices, Inc.
  Model name:             AMD Ryzen 7 4800H with Radeon Graphics
    BIOS Model name:      AMD Ryzen 7 4800H with Radeon Graphics          Unknown CPU @ 2.9GHz
    BIOS CPU family:      107
    CPU family:           23
    Model:                96
    Thread(s) per core:   2
    Core(s) per socket:   8
    Socket(s):            1
    Stepping:             1
    Frequency boost:      enabled
    CPU(s) scaling MHz:   108%
    CPU max MHz:          2900.0000
    CPU min MHz:          1400.0000
    BogoMIPS:             5789.09
    Flags:                fpu vme de pse tsc msr pae mce cx8 apic sep mtrr pge mca cmov pat pse36 clflush mmx fxsr sse sse2 ht syscall nx mmxext fxsr_opt pdpe1gb rd
                          tscp lm constant_tsc rep_good nopl nonstop_tsc cpuid extd_apicid aperfmperf rapl pni pclmulqdq monitor ssse3 fma cx16 sse4_1 sse4_2 movbe
                          popcnt aes xsave avx f16c rdrand lahf_lm cmp_legacy svm extapic cr8_legacy abm sse4a misalignsse 3dnowprefetch osvw ibs skinit wdt tce top
                          oext perfctr_core perfctr_nb bpext perfctr_llc mwaitx cpb cat_l3 cdp_l3 hw_pstate ssbd mba ibrs ibpb stibp vmmcall fsgsbase bmi1 avx2 smep
                           bmi2 cqm rdt_a rdseed adx smap clflushopt clwb sha_ni xsaveopt xsavec xgetbv1 cqm_llc cqm_occup_llc cqm_mbm_total cqm_mbm_local clzero ir
                          perf xsaveerptr rdpru wbnoinvd cppc arat npt lbrv svm_lock nrip_save tsc_scale vmcb_clean flushbyasid decodeassists pausefilter pfthreshol
                          d avic v_vmsave_vmload vgif v_spec_ctrl umip rdpid overflow_recov succor smca
Virtualization features:
  Virtualization:         AMD-V
Caches (sum of all):
  L1d:                    256 KiB (8 instances)
  L1i:                    256 KiB (8 instances)
  L2:                     4 MiB (8 instances)
  L3:                     8 MiB (2 instances)
NUMA:
  NUMA node(s):           1
  NUMA node0 CPU(s):      0-15
```

## 3. Memory

We can use the `lsmem` tool to get the memory model.

```bash
lsmem
```

```bash
RANGE                                 SIZE  STATE REMOVABLE BLOCK
0x0000000000000000-0x000000107fffffff  66G online       yes  0-32

Memory block size:         2G
Total online memory:      66G
Total offline memory:      0B
```

## 4. Disk

NVMe SSD is used to install the operating system, proxmox and hold backups for the VMS locally. HDD is used to hold the virtual machines.
Currently, I have not configured any RAID level for the HDD.

To test the disk performance, we can use the `dd` tool to write and read data from the disk.

```bash
dd if=/dev/zero of=/tmp/test bs=64k count=64k conv=fdatasync
```

Results from bohr:

```bash
65536+0 records in
65536+0 records out
4294967296 bytes (4.3 GB, 4.0 GiB) copied, 4.44474 s, 966 MB/s
```

### Benchmark utilities

#### bench.sh for my PC

```bash
wget -qO- bench.sh | bash
```

```bash
-------------------- A Bench.sh Script By Teddysun -------------------
 Version            : v2024-11-11
 Usage              : wget -qO- bench.sh | bash
----------------------------------------------------------------------
 CPU Model          : AMD Ryzen 7 7735HS with Radeon Graphics
 CPU Cores          : 16 @ 4341.376 MHz
 CPU Cache          : 512 KB
 AES-NI             : ✓ Enabled
 VM-x/AMD-V         : ✓ Enabled
 Total Disk         : 4.6 TB (3.0 TB Used)
 Total Mem          : 27.1 GB (11.0 GB Used)
 Total Swap         : 54.3 GB (14.5 MB Used)
 System uptime      : 0 days, 7 hour 28 min
 Load average       : 4.40, 2.48, 2.08
 OS                 : Arch Linux
 Arch               : x86_64 (64 Bit)
 Kernel             : 6.13.8-arch1-1
 TCP CC             : cubic
 Virtualization     : Dedicated
 IPv4/IPv6          : ✓ Online / ✗ Offline
Prepended http:// to 'ipinfo.io/org'
Prepended http:// to 'ipinfo.io/city'
Prepended http:// to 'ipinfo.io/country'
Prepended http:// to 'ipinfo.io/region
 Organization       : AS33771 Safaricom Limited
 Location           : Nairobi / KE
 Region             : Nairobi Area
----------------------------------------------------------------------
I/O Speed(1st run) : 965 MB/s
I/O Speed(2nd run) : 890 MB/s
I/O Speed(3rd run) : 820 MB/s
I/O Speed(average) : 891.7 MB/s
----------------------------------------------------------------------
Node Name        Upload Speed      Download Speed      Latency
Speedtest.net    28.81 Mbps        27.86 Mbps          11.37 ms
Paris, FR        29.52 Mbps        32.82 Mbps          169.86 ms
Amsterdam, NL    29.94 Mbps        29.65 Mbps          175.93 ms
Shanghai, CN     1.91 Mbps         26.48 Mbps          566.82 ms
Hong Kong, CN    27.50 Mbps        31.24 Mbps          188.28 ms
Singapore, SG    30.39 Mbps        34.24 Mbps          247.31 ms
Tokyo, JP        30.50 Mbps        29.29 Mbps          292.47 ms
----------------------------------------------------------------------
Finished in        : 4 min 15 sec
Timestamp          : 2025-03-29 12:55:35 EAT
----------------------------------------------------------------------
```

## nench for my PC

```bash
(wget -qO- wget.racing/nench.sh | bash; wget -qO- wget.racing/nench.sh | bash) 2>&1 | tee nench.log
```

```bash
Prepended http:// to 'wget.racing/nench.sh'
-------------------------------------------------
 nench.sh v2019.07.20 -- https://git.io/nench.sh
 benchmark timestamp:    2025-03-29 10:02:20 UTC
-------------------------------------------------

Processor:    AMD Ryzen 7 7735HS with Radeon Graphics
CPU cores:    16
Frequency:    4586.167 MHz
RAM:          27Gi
Swap:         54Gi
Kernel:       Linux 6.13.8-arch1-1 x86_64

Disks:
nvme0n1  931.5G  SSD
zram0   54.3G  SSD

CPU: SHA256-hashing 500 MB
    0.312 seconds
CPU: bzip2-compressing 500 MB
    2.422 seconds
CPU: AES-encrypting 500 MB
    0.574 seconds

ioping: seek rate
    min/avg/max/mdev = 63.3 us / 77.9 us / 7.78 ms / 55.2 us
ioping: sequential read speed
    generated 21.9 k requests in 5.00 s, 5.36 GiB, 4.39 k iops, 1.07 GiB/s

dd: sequential write speed
    1st run:    1049.04 MiB/s
    2nd run:    1049.04 MiB/s
    3rd run:    1049.04 MiB/s
    average:    1049.04 MiB/s

IPv4 speedtests
    your IPv4:    105.163.158.xxxx

    Cachefly CDN:         3.57 MiB/s
    Leaseweb (NL):        0.02 MiB/s
    Softlayer DAL (US):   0.00 MiB/s
    Online.net (FR):      3.46 MiB/s
    OVH BHS (CA):         3.00 MiB/s

No IPv6 connectivity detected
-------------------------------------------------
```

## bench.sh for bohr

```bash
wget -qO- bench.sh | bash
```

```bash
-------------------- A Bench.sh Script By Teddysun -------------------
 Version            : v2024-11-11
 Usage              : wget -qO- bench.sh | bash
----------------------------------------------------------------------
 CPU Model          : QEMU Virtual CPU version 2.5+
 CPU Cores          : 8 @ 2894.560 MHz
 CPU Cache          : 512 KB
 AES-NI             : ✓ Enabled
 VM-x/AMD-V         : ✗ Disabled
 Total Disk         : 492.1 GB (221.8 GB Used)
 Total Mem          : 19.5 GB (4.3 GB Used)
 System uptime      : 0 days, 1 hour 1 min
 Load average       : 0.46, 0.78, 1.09
 OS                 : Debian GNU/Linux 12
 Arch               : x86_64 (64 Bit)
 Kernel             : 6.1.0-32-amd64
 TCP CC             :
 Virtualization     : KVM
 IPv4/IPv6          : ✓ Online / ✗ Offline
 Organization       : AS33771 Safaricom Limited
 Location           : Nairobi / KE
 Region             : Nairobi Area
----------------------------------------------------------------------
I/O Speed(1st run) : 792 MB/s
I/O Speed(2nd run) : 917 MB/s
I/O Speed(3rd run) : 765 MB/s
I/O Speed(average) : 824.7 MB/s
----------------------------------------------------------------------
Node Name        Upload Speed      Download Speed      Latency
Speedtest.net    28.86 Mbps        28.90 Mbps          11.98 ms
Paris, FR        29.67 Mbps        32.36 Mbps          176.92 ms
Amsterdam, NL    29.83 Mbps        34.01 Mbps          176.65 ms
Shanghai, CN     2.80 Mbps         31.05 Mbps          928.30 ms
Hong Kong, CN    30.31 Mbps        33.44 Mbps          186.17 ms
Singapore, SG    29.66 Mbps        32.97 Mbps          245.92 ms
Tokyo, JP        32.60 Mbps        36.11 Mbps          294.09 ms
----------------------------------------------------------------------
Finished in        : 4 min 9 sec
Timestamp          : 2025-03-29 10:08:54 UTC
----------------------------------------------------------------------
```

## nench for bohr

```bash
(wget -qO- wget.racing/nench.sh | bash; wget -qO- wget.racing/nench.sh | bash) 2>&1 | tee nench.log
```

```bash
-------------------------------------------------
 nench.sh v2019.07.20 -- https://git.io/nench.sh
 benchmark timestamp:    2025-03-29 10:09:26 UTC
-------------------------------------------------

Processor:    QEMU Virtual CPU version 2.5+
CPU cores:    8
Frequency:    2894.560 MHz
RAM:          19Gi
bash: line 156: swapon: command not found
Swap:         -
Kernel:       Linux 6.1.0-32-amd64 x86_64

Disks:
sda    500G  HDD

CPU: SHA256-hashing 500 MB
    1.856 seconds
CPU: bzip2-compressing 500 MB
    3.691 seconds
CPU: AES-encrypting 500 MB
    0.633 seconds

ioping: seek rate
    min/avg/max/mdev = 71.5 us / 151.2 us / 10.2 ms / 151.9 us
ioping: sequential read speed
    generated 19.8 k requests in 5.00 s, 4.83 GiB, 3.96 k iops, 989.8 MiB/s

dd: sequential write speed
    1st run:    635.15 MiB/s
    2nd run:    1049.04 MiB/s
    3rd run:    1049.04 MiB/s
    average:    911.08 MiB/s

IPv4 speedtests
    your IPv4:    105.163.158.xxxx

    Cachefly CDN:         0.00 MiB/s
    Leaseweb (NL):        0.01 MiB/s
    Softlayer DAL (US):   0.00 MiB/s
    Online.net (FR):      3.34 MiB/s
    OVH BHS (CA):         2.99 MiB/s

No IPv6 connectivity detected
-------------------------------------------------
```

## bench.sh for odin

```bash
wget -qO- bench.sh | bash
```

```bash
-------------------- A Bench.sh Script By Teddysun -------------------
 Version            : v2024-11-11
 Usage              : wget -qO- bench.sh | bash
----------------------------------------------------------------------
 CPU Model          : AMD Ryzen 7 4800H with Radeon Graphics
 CPU Cores          : 16 @ 3028.863 MHz
 CPU Cache          : 512 KB
 AES-NI             : ✓ Enabled
 VM-x/AMD-V         : ✓ Enabled
 Total Disk         : 2.7 TB (812.7 GB Used)
 Total Mem          : 62.2 GB (28.6 GB Used)
 System uptime      : 0 days, 1 hour 13 min
 Load average       : 1.43, 2.68, 2.45
 OS                 : Debian GNU/Linux 12
 Arch               : x86_64 (64 Bit)
 Kernel             : 6.8.12-8-pve
 TCP CC             : cubic
 Virtualization     : Dedicated
 IPv4/IPv6          : ✓ Online / ✗ Offline
 Organization       : AS33771 Safaricom Limited
 Location           : Nairobi / KE
 Region             : Nairobi Area
----------------------------------------------------------------------
I/O Speed(1st run) : 863 MB/s
I/O Speed(2nd run) : 857 MB/s
I/O Speed(3rd run) : 858 MB/s
I/O Speed(average) : 859.3 MB/s
----------------------------------------------------------------------
Node Name        Upload Speed      Download Speed      Latency
Speedtest.net    28.88 Mbps        23.81 Mbps          12.06 ms
Paris, FR        29.41 Mbps        30.77 Mbps          175.75 ms
Amsterdam, NL    29.99 Mbps        20.28 Mbps          182.95 ms
Shanghai, CN     0.75 Mbps         19.41 Mbps          569.65 ms
Hong Kong, CN    30.65 Mbps        28.06 Mbps          187.84 ms
Singapore, SG    29.69 Mbps        28.38 Mbps          256.93 ms
Tokyo, JP        30.95 Mbps        28.11 Mbps          294.33 ms
----------------------------------------------------------------------
Finished in        : 4 min 4 sec
Timestamp          : 2025-03-29 13:19:57 EAT
----------------------------------------------------------------------
```

## nench for odin

```bash
(wget -qO- wget.racing/nench.sh | bash; wget -qO- wget.racing/nench.sh | bash) 2>&1 | tee nench.log
```

```bash
-------------------------------------------------
 nench.sh v2019.07.20 -- https://git.io/nench.sh
 benchmark timestamp:    2025-03-29 10:13:07 UTC
-------------------------------------------------

Processor:    AMD Ryzen 7 4800H with Radeon Graphics
CPU cores:    16
Frequency:    4240.619 MHz
RAM:          62Gi
Swap:         -
Kernel:       Linux 6.8.12-8-pve x86_64

Disks:
nvme0n1  953.9G  SSD
sda    1.8T  HDD

CPU: SHA256-hashing 500 MB
    1.787 seconds
CPU: bzip2-compressing 500 MB
    3.584 seconds
CPU: AES-encrypting 500 MB
    0.667 seconds

ioping: seek rate
    min/avg/max/mdev = 39.3 us / 68.5 us / 76.4 ms / 286.4 us
ioping: sequential read speed
    generated 22.2 k requests in 5.00 s, 5.42 GiB, 4.44 k iops, 1.08 GiB/s

dd: sequential write speed
    1st run:    872.61 MiB/s
    2nd run:    899.31 MiB/s
    3rd run:    872.61 MiB/s
    average:    881.51 MiB/s

IPv4 speedtests
    your IPv4:    105.163.158.xxxx

    Cachefly CDN:         3.26 MiB/s
    Leaseweb (NL):        0.01 MiB/s
    Softlayer DAL (US):   0.00 MiB/s
    Online.net (FR):      2.59 MiB/s
    OVH BHS (CA):         0.05 MiB/s

No IPv6 connectivity detected
-------------------------------------------------
```

## bench.sh for heimdall

```bash
wget -qO- bench.sh | bash
```

```bash
-------------------- A Bench.sh Script By Teddysun -------------------
 Version            : v2024-11-11
 Usage              : wget -qO- bench.sh | bash
----------------------------------------------------------------------
 CPU Model          : Intel(R) N100
 CPU Cores          : 4 @ 3120.872 MHz
 CPU Cache          : 6144 KB
 AES-NI             : ✓ Enabled
 VM-x/AMD-V         : ✓ Enabled
 Total Disk         : 0 (0 Used)
 Total Mem          : 15.4 GB (8.2 GB Used)
 System uptime      : 1 days, 13 hour 26 min
 Load average       : 0.30, 0.13, 0.05
 OS                 : Debian GNU/Linux 12
 Arch               : x86_64 (64 Bit)
 Kernel             : 6.8.12-8-pve
 TCP CC             : cubic
 Virtualization     : Dedicated
 IPv4/IPv6          : ✓ Online / ✗ Offline
 Organization       : AS33771 Safaricom Limited
 Location           : Nairobi / KE
 Region             : Nairobi Area
 ----------------------------------------------------------------------
  I/O Speed(1st run) : 3.8 GB/s
  I/O Speed(2nd run) : 3.8 GB/s
  I/O Speed(3rd run) : 3.8 GB/s
  I/O Speed(average) : 3891.2 MB/s
 ----------------------------------------------------------------------
  Node Name        Upload Speed      Download Speed      Latency
  Speedtest.net    29.34 Mbps        28.20 Mbps          13.04 ms
  Paris, FR        28.77 Mbps        32.70 Mbps          204.76 ms
  Amsterdam, NL    32.06 Mbps        29.57 Mbps          197.91 ms
  Hong Kong, CN    31.28 Mbps        33.47 Mbps          190.52 ms
  Singapore, SG    29.12 Mbps        31.90 Mbps          259.39 ms
  Tokyo, JP        33.39 Mbps        33.98 Mbps          306.92 ms
 ----------------------------------------------------------------------
  Finished in        : 3 min 51 sec
  Timestamp          : 2025-03-29 14:06:11 EAT
 ----------------------------------------------------------------------
```

## nench for heimdall

```bash
(wget -qO- wget.racing/nench.sh | bash; wget -qO- wget.racing/nench.sh | bash) 2>&1 | tee nench.log
```

```bash
-------------------------------------------------
 nench.sh v2019.07.20 -- https://git.io/nench.sh
 benchmark timestamp:    2025-03-29 11:00:29 UTC
-------------------------------------------------

Processor:    Intel(R) N100
CPU cores:    4
Frequency:    2900.792 MHz
RAM:          15Gi
Swap:         -
Kernel:       Linux 6.8.12-8-pve x86_64

Disks:
nvme0n1  476.9G  SSD
sda    3.6T  HDD

CPU: SHA256-hashing 500 MB
    2.134 seconds
CPU: bzip2-compressing 500 MB
    3.813 seconds
CPU: AES-encrypting 500 MB
    0.630 seconds

ioping: seek rate
    min/avg/max/mdev = 1.63 us / 2.17 us / 106.3 us / 756 ns
ioping: sequential read speed
    generated 225.2 k requests in 5.00 s, 55.0 GiB, 45.0 k iops, 11.0 GiB/s

dd: sequential write speed
    1st run:    3051.76 MiB/s
    2nd run:    2956.39 MiB/s
    3rd run:    3051.76 MiB/s
    average:    3019.97 MiB/s

IPv4 speedtests
    your IPv4:    105.163.158.xxxx

    Cachefly CDN:         3.48 MiB/s
    Leaseweb (NL):        0.02 MiB/s
    Softlayer DAL (US):   0.00 MiB/s
    Online.net (FR):      3.49 MiB/s
    OVH BHS (CA):         0.83 MiB/s

No IPv6 connectivity detected
-------------------------------------------------
```
