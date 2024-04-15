# Soak Test Result

The following is the soak test dump, you can run it yourself.

The test used an infinimap for 30 minutes, and during that period it performed 700 million random operations. An operation can be: insert, get, update, delete.

The operations were executed in both: an infinimap and in a traditional golang map used as reference.
Each operation was applied to both maps and its behaviour was asserted to be similar. i.e. if not found in the reference map, it should not be in the infinimap,
if an update, the old value should be the same in both maps.

At the end both of the process, both maps were fully compared and found to be exact.
Then, the infinimap was compacted, and the resulting was verified against the referenc.

```shell
$  make && demo/soak 30
infinimap: soak-test
2024/04/15 19:12:07 demo/soak: soak test
2024/04/15 19:12:07 running for 30 minute(s).
2024/04/15 19:12:07 (If you want other value, just call this with a numeric parameter.)
2024/04/15 19:12:37 [1%] 30.69M ops, 5.35M entries: 10.60M inserts, 5.25M updates, 5.25M deletes, 9.59M gets, 0.0% clog
2024/04/15 19:12:38   ... disk space: 2.0G allocated, 1.4G in use, 0.6G reclaimable, 15623.0G available
2024/04/15 19:13:07 [3%] 55.06M ops, 9.38M entries: 18.71M inserts, 9.33M updates, 9.33M deletes, 17.70M gets, 0.0% clog
2024/04/15 19:13:08   ... disk space: 2.7G allocated, 1.7G in use, 1.1G reclaimable, 15622.3G available
2024/04/15 19:13:37 [5%] 79.84M ops, 10.00M entries: 26.07M inserts, 12.63M updates, 16.07M deletes, 25.08M gets, 0.0% clog
2024/04/15 19:13:38   ... disk space: 3.3G allocated, 1.7G in use, 1.7G reclaimable, 15621.7G available
2024/04/15 19:14:07 [6%] 104.22M ops, 10.00M entries: 33.37M inserts, 15.10M updates, 23.37M deletes, 32.38M gets, 0.0% clog
2024/04/15 19:14:08   ... disk space: 3.9G allocated, 1.7G in use, 2.2G reclaimable, 15621.1G available
2024/04/15 19:14:37 [8%] 126.33M ops, 10.00M entries: 40.12M inserts, 16.94M updates, 30.12M deletes, 39.14M gets, 0.0% clog
2024/04/15 19:14:38   ... disk space: 4.4G allocated, 1.7G in use, 2.7G reclaimable, 15620.6G available
2024/04/15 19:15:07 [10%] 146.99M ops, 10.00M entries: 46.52M inserts, 18.42M updates, 36.52M deletes, 45.54M gets, 0.0% clog
2024/04/15 19:15:08   ... disk space: 4.8G allocated, 1.7G in use, 3.1G reclaimable, 15620.2G available
2024/04/15 19:15:37 [11%] 165.54M ops, 10.00M entries: 52.31M inserts, 19.59M updates, 42.31M deletes, 51.33M gets, 0.0% clog
2024/04/15 19:15:38   ... disk space: 5.1G allocated, 1.6G in use, 3.5G reclaimable, 15619.9G available
2024/04/15 19:16:07 [13%] 184.44M ops, 10.00M entries: 58.25M inserts, 20.67M updates, 48.25M deletes, 57.26M gets, 0.0% clog
2024/04/15 19:16:08   ... disk space: 5.5G allocated, 1.6G in use, 3.9G reclaimable, 15619.5G available
2024/04/15 19:16:37 [15%] 202.80M ops, 10.00M entries: 64.06M inserts, 21.62M updates, 54.06M deletes, 63.07M gets, 0.0% clog
2024/04/15 19:16:38   ... disk space: 5.9G allocated, 1.6G in use, 4.2G reclaimable, 15619.1G available
2024/04/15 19:17:07 [16%] 220.15M ops, 10.00M entries: 69.57M inserts, 22.44M updates, 59.57M deletes, 68.58M gets, 0.0% clog
2024/04/15 19:17:08   ... disk space: 6.2G allocated, 1.6G in use, 4.5G reclaimable, 15618.8G available
2024/04/15 19:17:37 [18%] 237.25M ops, 10.00M entries: 75.02M inserts, 23.20M updates, 65.02M deletes, 74.03M gets, 0.0% clog
2024/04/15 19:17:38   ... disk space: 6.5G allocated, 1.6G in use, 4.9G reclaimable, 15618.5G available
2024/04/15 19:18:07 [20%] 253.70M ops, 10.00M entries: 80.27M inserts, 23.87M updates, 70.27M deletes, 79.28M gets, 0.0% clog
2024/04/15 19:18:08   ... disk space: 6.8G allocated, 1.6G in use, 5.2G reclaimable, 15618.2G available
2024/04/15 19:18:37 [21%] 269.48M ops, 10.00M entries: 85.33M inserts, 24.48M updates, 75.33M deletes, 84.34M gets, 0.0% clog
2024/04/15 19:18:38   ... disk space: 7.1G allocated, 1.6G in use, 5.5G reclaimable, 15617.9G available
2024/04/15 19:19:07 [23%] 284.76M ops, 10.00M entries: 90.23M inserts, 25.04M updates, 80.23M deletes, 89.25M gets, 0.0% clog
2024/04/15 19:19:08   ... disk space: 7.3G allocated, 1.6G in use, 5.7G reclaimable, 15617.7G available
2024/04/15 19:19:37 [25%] 298.89M ops, 10.00M entries: 94.78M inserts, 25.53M updates, 84.78M deletes, 93.80M gets, 0.0% clog
2024/04/15 19:19:38   ... disk space: 7.6G allocated, 1.6G in use, 6.0G reclaimable, 15617.4G available
2024/04/15 19:20:07 [26%] 312.71M ops, 10.00M entries: 99.23M inserts, 25.99M updates, 89.23M deletes, 98.25M gets, 0.0% clog
2024/04/15 19:20:09   ... disk space: 7.8G allocated, 1.6G in use, 6.2G reclaimable, 15617.2G available
2024/04/15 19:20:37 [28%] 326.65M ops, 10.00M entries: 103.73M inserts, 26.44M updates, 93.73M deletes, 102.74M gets, 0.0% clog
2024/04/15 19:20:39   ... disk space: 8.1G allocated, 1.6G in use, 6.5G reclaimable, 15616.9G available
2024/04/15 19:21:07 [30%] 340.42M ops, 10.00M entries: 108.19M inserts, 26.86M updates, 98.19M deletes, 107.19M gets, 0.0% clog
2024/04/15 19:21:09   ... disk space: 8.3G allocated, 1.6G in use, 6.7G reclaimable, 15616.7G available
2024/04/15 19:21:37 [31%] 352.90M ops, 10.00M entries: 112.22M inserts, 27.22M updates, 102.22M deletes, 111.23M gets, 0.0% clog
2024/04/15 19:21:39   ... disk space: 8.6G allocated, 1.6G in use, 6.9G reclaimable, 15616.4G available
2024/04/15 19:22:07 [33%] 365.17M ops, 10.00M entries: 116.20M inserts, 27.57M updates, 106.20M deletes, 115.21M gets, 0.0% clog
2024/04/15 19:22:09   ... disk space: 8.8G allocated, 1.6G in use, 7.2G reclaimable, 15616.2G available
2024/04/15 19:22:37 [35%] 377.74M ops, 10.00M entries: 120.27M inserts, 27.92M updates, 110.27M deletes, 119.28M gets, 0.0% clog
2024/04/15 19:22:39   ... disk space: 9.0G allocated, 1.6G in use, 7.4G reclaimable, 15616.0G available
2024/04/15 19:23:07 [36%] 389.92M ops, 10.00M entries: 124.22M inserts, 28.24M updates, 114.22M deletes, 123.24M gets, 0.0% clog
2024/04/15 19:23:09   ... disk space: 9.2G allocated, 1.6G in use, 7.6G reclaimable, 15615.8G available
2024/04/15 19:23:37 [38%] 401.38M ops, 10.00M entries: 127.95M inserts, 28.53M updates, 117.95M deletes, 126.96M gets, 0.0% clog
2024/04/15 19:23:39   ... disk space: 9.4G allocated, 1.6G in use, 7.8G reclaimable, 15615.6G available
2024/04/15 19:24:07 [40%] 413.03M ops, 10.00M entries: 131.73M inserts, 28.83M updates, 121.73M deletes, 130.75M gets, 0.0% clog
2024/04/15 19:24:09   ... disk space: 9.6G allocated, 1.6G in use, 8.0G reclaimable, 15615.4G available
2024/04/15 19:24:37 [41%] 424.50M ops, 10.00M entries: 135.46M inserts, 29.10M updates, 125.46M deletes, 134.48M gets, 0.0% clog
2024/04/15 19:24:39   ... disk space: 9.8G allocated, 1.6G in use, 8.2G reclaimable, 15615.2G available
2024/04/15 19:25:07 [43%] 434.87M ops, 10.00M entries: 138.84M inserts, 29.35M updates, 128.84M deletes, 137.85M gets, 0.0% clog
2024/04/15 19:25:09   ... disk space: 10.0G allocated, 1.6G in use, 8.4G reclaimable, 15615.0G available
2024/04/15 19:25:37 [45%] 445.78M ops, 10.00M entries: 142.38M inserts, 29.60M updates, 132.38M deletes, 141.40M gets, 0.0% clog
2024/04/15 19:25:39   ... disk space: 10.2G allocated, 1.6G in use, 8.6G reclaimable, 15614.8G available
2024/04/15 19:26:07 [46%] 456.30M ops, 10.00M entries: 145.81M inserts, 29.84M updates, 135.81M deletes, 144.83M gets, 0.0% clog
2024/04/15 19:26:09   ... disk space: 10.4G allocated, 1.6G in use, 8.8G reclaimable, 15614.6G available
2024/04/15 19:26:37 [48%] 466.29M ops, 10.00M entries: 149.07M inserts, 30.06M updates, 139.07M deletes, 148.09M gets, 0.0% clog
2024/04/15 19:26:39   ... disk space: 10.5G allocated, 1.6G in use, 8.9G reclaimable, 15614.5G available
2024/04/15 19:27:07 [50%] 476.88M ops, 10.00M entries: 152.52M inserts, 30.29M updates, 142.52M deletes, 151.55M gets, 0.0% clog
2024/04/15 19:27:09   ... disk space: 10.7G allocated, 1.6G in use, 9.1G reclaimable, 15614.3G available
2024/04/15 19:27:37 [51%] 486.07M ops, 10.00M entries: 155.52M inserts, 30.49M updates, 145.52M deletes, 154.55M gets, 0.0% clog
2024/04/15 19:27:39   ... disk space: 10.9G allocated, 1.6G in use, 9.3G reclaimable, 15614.1G available
2024/04/15 19:28:07 [53%] 495.51M ops, 10.00M entries: 158.60M inserts, 30.68M updates, 148.60M deletes, 157.63M gets, 0.0% clog
2024/04/15 19:28:09   ... disk space: 11.1G allocated, 1.6G in use, 9.4G reclaimable, 15613.9G available
2024/04/15 19:28:37 [55%] 504.99M ops, 10.00M entries: 161.70M inserts, 30.87M updates, 151.70M deletes, 160.72M gets, 0.0% clog
2024/04/15 19:28:39   ... disk space: 11.2G allocated, 1.6G in use, 9.6G reclaimable, 15613.8G available
2024/04/15 19:29:07 [56%] 513.40M ops, 10.00M entries: 164.44M inserts, 31.04M updates, 154.44M deletes, 163.47M gets, 0.0% clog
2024/04/15 19:29:09   ... disk space: 11.4G allocated, 1.6G in use, 9.7G reclaimable, 15613.6G available
2024/04/15 19:29:37 [58%] 522.25M ops, 10.00M entries: 167.34M inserts, 31.22M updates, 157.34M deletes, 166.36M gets, 0.0% clog
2024/04/15 19:29:39   ... disk space: 11.5G allocated, 1.6G in use, 9.9G reclaimable, 15613.5G available
2024/04/15 19:30:07 [60%] 530.96M ops, 10.00M entries: 170.18M inserts, 31.39M updates, 160.18M deletes, 169.21M gets, 0.0% clog
2024/04/15 19:30:09   ... disk space: 11.7G allocated, 1.6G in use, 10.0G reclaimable, 15613.3G available
2024/04/15 19:30:37 [61%] 539.43M ops, 10.00M entries: 172.95M inserts, 31.55M updates, 162.95M deletes, 171.98M gets, 0.0% clog
2024/04/15 19:30:39   ... disk space: 11.8G allocated, 1.6G in use, 10.2G reclaimable, 15613.2G available
2024/04/15 19:31:07 [63%] 548.18M ops, 10.00M entries: 175.82M inserts, 31.71M updates, 165.82M deletes, 174.84M gets, 0.0% clog
2024/04/15 19:31:09   ... disk space: 12.0G allocated, 1.6G in use, 10.3G reclaimable, 15613.0G available
2024/04/15 19:31:37 [65%] 556.28M ops, 10.00M entries: 178.47M inserts, 31.86M updates, 168.47M deletes, 177.48M gets, 0.0% clog
2024/04/15 19:31:39   ... disk space: 12.1G allocated, 1.6G in use, 10.5G reclaimable, 15612.9G available
2024/04/15 19:32:07 [66%] 564.28M ops, 10.00M entries: 181.08M inserts, 32.01M updates, 171.08M deletes, 180.10M gets, 0.0% clog
2024/04/15 19:32:10   ... disk space: 12.2G allocated, 1.6G in use, 10.6G reclaimable, 15612.8G available
2024/04/15 19:32:37 [68%] 571.70M ops, 10.00M entries: 183.51M inserts, 32.14M updates, 173.51M deletes, 182.53M gets, 0.0% clog
2024/04/15 19:32:39   ... disk space: 12.4G allocated, 1.6G in use, 10.7G reclaimable, 15612.6G available
2024/04/15 19:33:07 [70%] 579.34M ops, 10.00M entries: 186.02M inserts, 32.28M updates, 176.02M deletes, 185.04M gets, 0.0% clog
2024/04/15 19:33:09   ... disk space: 12.5G allocated, 1.6G in use, 10.9G reclaimable, 15612.5G available
2024/04/15 19:33:37 [71%] 586.81M ops, 10.00M entries: 188.46M inserts, 32.41M updates, 178.46M deletes, 187.48M gets, 0.0% clog
2024/04/15 19:33:39   ... disk space: 12.6G allocated, 1.6G in use, 11.0G reclaimable, 15612.4G available
2024/04/15 19:34:07 [73%] 594.64M ops, 10.00M entries: 191.03M inserts, 32.54M updates, 181.03M deletes, 190.04M gets, 0.0% clog
2024/04/15 19:34:09   ... disk space: 12.7G allocated, 1.6G in use, 11.1G reclaimable, 15612.3G available
2024/04/15 19:34:37 [75%] 602.37M ops, 10.00M entries: 193.56M inserts, 32.67M updates, 183.56M deletes, 192.57M gets, 0.0% clog
2024/04/15 19:34:39   ... disk space: 12.9G allocated, 1.6G in use, 11.3G reclaimable, 15612.1G available
2024/04/15 19:35:07 [76%] 609.74M ops, 10.00M entries: 195.97M inserts, 32.80M updates, 185.97M deletes, 194.99M gets, 0.0% clog
2024/04/15 19:35:09   ... disk space: 13.0G allocated, 1.6G in use, 11.4G reclaimable, 15612.0G available
2024/04/15 19:35:37 [78%] 617.36M ops, 10.00M entries: 198.47M inserts, 32.92M updates, 188.47M deletes, 197.49M gets, 0.0% clog
2024/04/15 19:35:40   ... disk space: 13.1G allocated, 1.6G in use, 11.5G reclaimable, 15611.9G available
2024/04/15 19:36:07 [80%] 624.63M ops, 10.00M entries: 200.86M inserts, 33.04M updates, 190.86M deletes, 199.87M gets, 0.0% clog
2024/04/15 19:36:09   ... disk space: 13.3G allocated, 1.6G in use, 11.6G reclaimable, 15611.7G available
2024/04/15 19:36:37 [81%] 631.51M ops, 10.00M entries: 203.11M inserts, 33.16M updates, 193.11M deletes, 202.13M gets, 0.0% clog
2024/04/15 19:36:39   ... disk space: 13.4G allocated, 1.6G in use, 11.8G reclaimable, 15611.6G available
2024/04/15 19:37:07 [83%] 638.77M ops, 10.00M entries: 205.49M inserts, 33.27M updates, 195.49M deletes, 204.51M gets, 0.0% clog
2024/04/15 19:37:10   ... disk space: 13.5G allocated, 1.6G in use, 11.9G reclaimable, 15611.5G available
2024/04/15 19:37:37 [85%] 645.33M ops, 10.00M entries: 207.65M inserts, 33.38M updates, 197.65M deletes, 206.66M gets, 0.0% clog
2024/04/15 19:37:39   ... disk space: 13.6G allocated, 1.6G in use, 12.0G reclaimable, 15611.4G available
2024/04/15 19:38:07 [86%] 651.82M ops, 10.00M entries: 209.78M inserts, 33.48M updates, 199.78M deletes, 208.79M gets, 0.0% clog
2024/04/15 19:38:10   ... disk space: 13.7G allocated, 1.6G in use, 12.1G reclaimable, 15611.3G available
2024/04/15 19:38:37 [88%] 658.65M ops, 10.00M entries: 212.02M inserts, 33.58M updates, 202.02M deletes, 211.03M gets, 0.0% clog
2024/04/15 19:38:39   ... disk space: 13.8G allocated, 1.6G in use, 12.2G reclaimable, 15611.2G available
2024/04/15 19:39:07 [90%] 664.93M ops, 10.00M entries: 214.08M inserts, 33.68M updates, 204.08M deletes, 213.09M gets, 0.0% clog
2024/04/15 19:39:09   ... disk space: 13.9G allocated, 1.6G in use, 12.3G reclaimable, 15611.1G available
2024/04/15 19:39:37 [91%] 671.31M ops, 10.00M entries: 216.17M inserts, 33.78M updates, 206.17M deletes, 215.19M gets, 0.0% clog
2024/04/15 19:39:40   ... disk space: 14.0G allocated, 1.6G in use, 12.4G reclaimable, 15611.0G available
2024/04/15 19:40:07 [93%] 677.03M ops, 10.00M entries: 218.05M inserts, 33.86M updates, 208.05M deletes, 217.07M gets, 0.0% clog
2024/04/15 19:40:10   ... disk space: 14.1G allocated, 1.6G in use, 12.5G reclaimable, 15610.9G available
2024/04/15 19:40:37 [95%] 682.54M ops, 10.00M entries: 219.86M inserts, 33.95M updates, 209.86M deletes, 218.87M gets, 0.0% clog
2024/04/15 19:40:40   ... disk space: 14.2G allocated, 1.6G in use, 12.6G reclaimable, 15610.8G available
2024/04/15 19:41:07 [96%] 688.63M ops, 10.00M entries: 221.86M inserts, 34.04M updates, 211.86M deletes, 220.88M gets, 0.0% clog
2024/04/15 19:41:10   ... disk space: 14.3G allocated, 1.6G in use, 12.7G reclaimable, 15610.7G available
2024/04/15 19:41:37 [98%] 694.69M ops, 10.00M entries: 223.85M inserts, 34.13M updates, 213.85M deletes, 222.86M gets, 0.0% clog
2024/04/15 19:41:40   ... disk space: 14.4G allocated, 1.6G in use, 12.8G reclaimable, 15610.6G available
2024/04/15 19:42:07 [100%] 700.49M ops, 10.00M entries: 225.75M inserts, 34.21M updates, 215.75M deletes, 224.77M gets, 0.0% clog
2024/04/15 19:42:10   ... disk space: 14.5G allocated, 1.6G in use, 12.9G reclaimable, 15610.5G available
2024/04/15 19:42:10 Verifying contents ...
2024/04/15 19:42:20 SUCCESS! All values matched between reference map and infinimap
2024/04/15 19:42:20 Compacting ...
2024/04/15 19:42:46 New file size is 0.9G
2024/04/15 19:42:46 Verifying contents ...
2024/04/15 19:42:55 SUCCESS! All values matched between reference map and infinimap
```