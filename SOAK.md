# Soak Test Result

The following is the soak test dump, you can run it yourself.

The test used an infinimap for 30 minutes, and during that period it performed 682 million random operations. An operation can be: insert, get, update, delete.

The operations were executed in both: an infinimap and in a traditional golang map used as reference.
Each operation was applied to both maps and its behaviour was asserted to be similar. i.e. if not found in the reference map, it should not be in the infinimap,
if an update, the old value should be the same in both maps.

At the end both of the process, both maps were fully compared and found to be exact.

There were 219.94M inserts, 33.95M updates, 209.94M deletes, and 218.94M gets, during the process 14.2G were used, 12.6G could have been reclaim i.e. deleted
records. Effective data stored at the end was 1.6G.

```shell
$  make && demo/soak 30
infinimap: soak-test
2024/04/15 15:11:48 demo/soak: soak test
2024/04/15 15:11:48 running for 30 minute(s).
2024/04/15 15:11:48 (If you want other value, just call this with a numeric parameter.)
2024/04/15 15:12:03 [0%] 14.33M ops, 2.69M entries: 5.18M inserts, 2.49M updates, 2.49M deletes, 4.17M gets, 0.0% clog
2024/04/15 15:12:03   ... disk space: 1.6G allocated, 1.3G in use, 0.3G reclaimable, 15623.4G available
2024/04/15 15:12:18 [1%] 29.99M ops, 5.23M entries: 10.37M inserts, 5.13M updates, 5.13M deletes, 9.36M gets, 0.0% clog
2024/04/15 15:12:18   ... disk space: 2.0G allocated, 1.4G in use, 0.6G reclaimable, 15623.0G available
2024/04/15 15:12:33 [2%] 42.31M ops, 7.27M entries: 14.46M inserts, 7.20M updates, 7.20M deletes, 13.45M gets, 0.0% clog
2024/04/15 15:12:33   ... disk space: 2.4G allocated, 1.5G in use, 0.8G reclaimable, 15622.6G available
2024/04/15 15:12:48 [3%] 55.92M ops, 9.52M entries: 18.99M inserts, 9.47M updates, 9.47M deletes, 17.98M gets, 0.0% clog
2024/04/15 15:12:48   ... disk space: 2.8G allocated, 1.7G in use, 1.1G reclaimable, 15622.2G available
2024/04/15 15:13:03 [4%] 69.43M ops, 10.00M entries: 23.02M inserts, 11.38M updates, 13.02M deletes, 22.01M gets, 0.0% clog
2024/04/15 15:13:03   ... disk space: 3.1G allocated, 1.7G in use, 1.4G reclaimable, 15621.9G available
2024/04/15 15:13:18 [5%] 80.70M ops, 10.00M entries: 26.33M inserts, 12.73M updates, 16.33M deletes, 25.31M gets, 0.0% clog
2024/04/15 15:13:18   ... disk space: 3.4G allocated, 1.7G in use, 1.7G reclaimable, 15621.6G available
2024/04/15 15:13:33 [5%] 92.75M ops, 10.00M entries: 29.92M inserts, 14.00M updates, 19.92M deletes, 28.90M gets, 0.0% clog
2024/04/15 15:13:33   ... disk space: 3.6G allocated, 1.7G in use, 2.0G reclaimable, 15621.4G available
2024/04/15 15:13:48 [6%] 105.32M ops, 10.00M entries: 33.72M inserts, 15.20M updates, 23.72M deletes, 32.70M gets, 0.0% clog
2024/04/15 15:13:48   ... disk space: 3.9G allocated, 1.7G in use, 2.2G reclaimable, 15621.1G available
2024/04/15 15:14:03 [7%] 116.30M ops, 10.00M entries: 37.06M inserts, 16.14M updates, 27.06M deletes, 36.04M gets, 0.0% clog
2024/04/15 15:14:03   ... disk space: 4.1G allocated, 1.7G in use, 2.5G reclaimable, 15620.9G available
2024/04/15 15:14:18 [8%] 127.00M ops, 10.00M entries: 40.34M inserts, 16.99M updates, 30.34M deletes, 39.32M gets, 0.0% clog
2024/04/15 15:14:18   ... disk space: 4.4G allocated, 1.7G in use, 2.7G reclaimable, 15620.6G available
2024/04/15 15:14:33 [9%] 138.66M ops, 10.00M entries: 43.94M inserts, 17.85M updates, 33.94M deletes, 42.93M gets, 0.0% clog
2024/04/15 15:14:33   ... disk space: 4.6G allocated, 1.7G in use, 2.9G reclaimable, 15620.4G available
2024/04/15 15:14:48 [10%] 148.90M ops, 10.00M entries: 47.12M inserts, 18.55M updates, 37.12M deletes, 46.11M gets, 0.0% clog
2024/04/15 15:14:48   ... disk space: 4.8G allocated, 1.7G in use, 3.2G reclaimable, 15620.2G available
2024/04/15 15:15:03 [10%] 158.84M ops, 10.00M entries: 50.22M inserts, 19.19M updates, 40.22M deletes, 49.21M gets, 0.0% clog
2024/04/15 15:15:03   ... disk space: 5.0G allocated, 1.7G in use, 3.4G reclaimable, 15620.0G available
2024/04/15 15:15:18 [11%] 169.10M ops, 10.00M entries: 53.43M inserts, 19.81M updates, 43.43M deletes, 52.43M gets, 0.0% clog
2024/04/15 15:15:18   ... disk space: 5.2G allocated, 1.6G in use, 3.6G reclaimable, 15619.8G available
2024/04/15 15:15:33 [12%] 179.20M ops, 10.00M entries: 56.61M inserts, 20.38M updates, 46.61M deletes, 55.60M gets, 0.0% clog
2024/04/15 15:15:33   ... disk space: 5.4G allocated, 1.6G in use, 3.8G reclaimable, 15619.6G available
2024/04/15 15:15:48 [13%] 188.18M ops, 10.00M entries: 59.44M inserts, 20.87M updates, 49.44M deletes, 58.43M gets, 0.0% clog
2024/04/15 15:15:48   ... disk space: 5.6G allocated, 1.6G in use, 3.9G reclaimable, 15619.4G available
2024/04/15 15:16:03 [14%] 197.18M ops, 10.00M entries: 62.29M inserts, 21.34M updates, 52.29M deletes, 61.27M gets, 0.0% clog
2024/04/15 15:16:03   ... disk space: 5.7G allocated, 1.6G in use, 4.1G reclaimable, 15619.3G available
2024/04/15 15:16:18 [15%] 206.69M ops, 10.00M entries: 65.30M inserts, 21.81M updates, 55.30M deletes, 64.28M gets, 0.0% clog
2024/04/15 15:16:18   ... disk space: 5.9G allocated, 1.6G in use, 4.3G reclaimable, 15619.1G available
2024/04/15 15:16:33 [15%] 215.28M ops, 10.00M entries: 68.02M inserts, 22.22M updates, 58.02M deletes, 67.01M gets, 0.0% clog
2024/04/15 15:16:34   ... disk space: 6.1G allocated, 1.6G in use, 4.4G reclaimable, 15618.9G available
2024/04/15 15:16:48 [16%] 223.47M ops, 10.00M entries: 70.63M inserts, 22.59M updates, 60.63M deletes, 69.61M gets, 0.0% clog
2024/04/15 15:16:48   ... disk space: 6.2G allocated, 1.6G in use, 4.6G reclaimable, 15618.8G available
2024/04/15 15:17:03 [17%] 232.30M ops, 10.00M entries: 73.44M inserts, 22.99M updates, 63.44M deletes, 72.43M gets, 0.0% clog
2024/04/15 15:17:03   ... disk space: 6.4G allocated, 1.6G in use, 4.8G reclaimable, 15618.6G available
2024/04/15 15:17:18 [18%] 240.84M ops, 10.00M entries: 76.17M inserts, 23.35M updates, 66.17M deletes, 75.16M gets, 0.0% clog
2024/04/15 15:17:19   ... disk space: 6.6G allocated, 1.6G in use, 4.9G reclaimable, 15618.4G available
2024/04/15 15:17:33 [19%] 248.30M ops, 10.00M entries: 78.55M inserts, 23.66M updates, 68.55M deletes, 77.54M gets, 0.0% clog
2024/04/15 15:17:34   ... disk space: 6.7G allocated, 1.6G in use, 5.1G reclaimable, 15618.3G available
2024/04/15 15:17:48 [20%] 256.17M ops, 10.00M entries: 81.07M inserts, 23.97M updates, 71.07M deletes, 80.05M gets, 0.0% clog
2024/04/15 15:17:49   ... disk space: 6.8G allocated, 1.6G in use, 5.2G reclaimable, 15618.2G available
2024/04/15 15:18:03 [20%] 264.35M ops, 10.00M entries: 83.69M inserts, 24.29M updates, 73.69M deletes, 82.67M gets, 0.0% clog
2024/04/15 15:18:04   ... disk space: 7.0G allocated, 1.6G in use, 5.4G reclaimable, 15618.0G available
2024/04/15 15:18:18 [21%] 271.65M ops, 10.00M entries: 86.03M inserts, 24.57M updates, 76.03M deletes, 85.02M gets, 0.0% clog
2024/04/15 15:18:19   ... disk space: 7.1G allocated, 1.6G in use, 5.5G reclaimable, 15617.9G available
2024/04/15 15:18:33 [22%] 278.54M ops, 10.00M entries: 88.24M inserts, 24.82M updates, 78.24M deletes, 87.23M gets, 0.0% clog
2024/04/15 15:18:34   ... disk space: 7.2G allocated, 1.6G in use, 5.6G reclaimable, 15617.8G available
2024/04/15 15:18:48 [23%] 286.06M ops, 10.00M entries: 90.66M inserts, 25.09M updates, 80.66M deletes, 89.65M gets, 0.0% clog
2024/04/15 15:18:49   ... disk space: 7.4G allocated, 1.6G in use, 5.8G reclaimable, 15617.6G available
2024/04/15 15:19:03 [24%] 293.61M ops, 10.00M entries: 93.09M inserts, 25.36M updates, 83.09M deletes, 92.08M gets, 0.0% clog
2024/04/15 15:19:04   ... disk space: 7.5G allocated, 1.6G in use, 5.9G reclaimable, 15617.5G available
2024/04/15 15:19:18 [25%] 299.97M ops, 10.00M entries: 95.14M inserts, 25.57M updates, 85.14M deletes, 94.12M gets, 0.0% clog
2024/04/15 15:19:19   ... disk space: 7.6G allocated, 1.6G in use, 6.0G reclaimable, 15617.4G available
2024/04/15 15:19:33 [25%] 306.95M ops, 10.00M entries: 97.39M inserts, 25.81M updates, 87.39M deletes, 96.37M gets, 0.0% clog
2024/04/15 15:19:34   ... disk space: 7.7G allocated, 1.6G in use, 6.1G reclaimable, 15617.3G available
2024/04/15 15:19:48 [26%] 314.10M ops, 10.00M entries: 99.69M inserts, 26.04M updates, 89.69M deletes, 98.68M gets, 0.0% clog
2024/04/15 15:19:49   ... disk space: 7.9G allocated, 1.6G in use, 6.3G reclaimable, 15617.1G available
2024/04/15 15:20:03 [27%] 320.37M ops, 10.00M entries: 101.72M inserts, 26.24M updates, 91.72M deletes, 100.70M gets, 0.0% clog
2024/04/15 15:20:04   ... disk space: 8.0G allocated, 1.6G in use, 6.4G reclaimable, 15617.0G available
2024/04/15 15:20:18 [28%] 326.17M ops, 10.00M entries: 103.59M inserts, 26.42M updates, 93.59M deletes, 102.57M gets, 0.0% clog
2024/04/15 15:20:19   ... disk space: 8.1G allocated, 1.6G in use, 6.5G reclaimable, 15616.9G available
2024/04/15 15:20:33 [29%] 332.92M ops, 10.00M entries: 105.77M inserts, 26.63M updates, 95.77M deletes, 104.75M gets, 0.0% clog
2024/04/15 15:20:34   ... disk space: 8.2G allocated, 1.6G in use, 6.6G reclaimable, 15616.8G available
2024/04/15 15:20:48 [30%] 339.26M ops, 10.00M entries: 107.81M inserts, 26.82M updates, 97.81M deletes, 106.80M gets, 0.0% clog
2024/04/15 15:20:49   ... disk space: 8.3G allocated, 1.6G in use, 6.7G reclaimable, 15616.7G available
2024/04/15 15:21:03 [30%] 344.99M ops, 10.00M entries: 109.67M inserts, 26.99M updates, 99.67M deletes, 108.66M gets, 0.0% clog
2024/04/15 15:21:04   ... disk space: 8.4G allocated, 1.6G in use, 6.8G reclaimable, 15616.6G available
2024/04/15 15:21:18 [31%] 351.41M ops, 10.00M entries: 111.75M inserts, 27.18M updates, 101.75M deletes, 110.73M gets, 0.0% clog
2024/04/15 15:21:19   ... disk space: 8.5G allocated, 1.6G in use, 6.9G reclaimable, 15616.5G available
2024/04/15 15:21:33 [32%] 357.91M ops, 10.00M entries: 113.85M inserts, 27.37M updates, 103.85M deletes, 112.84M gets, 0.0% clog
2024/04/15 15:21:34   ... disk space: 8.7G allocated, 1.6G in use, 7.0G reclaimable, 15616.3G available
2024/04/15 15:21:48 [33%] 363.41M ops, 10.00M entries: 115.64M inserts, 27.52M updates, 105.64M deletes, 114.62M gets, 0.0% clog
2024/04/15 15:21:49   ... disk space: 8.8G allocated, 1.6G in use, 7.1G reclaimable, 15616.2G available
2024/04/15 15:22:03 [34%] 369.39M ops, 10.00M entries: 117.57M inserts, 27.69M updates, 107.57M deletes, 116.55M gets, 0.0% clog
2024/04/15 15:22:04   ... disk space: 8.9G allocated, 1.6G in use, 7.2G reclaimable, 15616.1G available
2024/04/15 15:22:18 [35%] 375.72M ops, 10.00M entries: 119.63M inserts, 27.86M updates, 109.63M deletes, 118.61M gets, 0.0% clog
2024/04/15 15:22:19   ... disk space: 9.0G allocated, 1.6G in use, 7.3G reclaimable, 15616.0G available
2024/04/15 15:22:33 [35%] 381.31M ops, 10.00M entries: 121.44M inserts, 28.01M updates, 111.44M deletes, 120.42M gets, 0.0% clog
2024/04/15 15:22:34   ... disk space: 9.1G allocated, 1.6G in use, 7.4G reclaimable, 15615.9G available
2024/04/15 15:22:48 [36%] 386.73M ops, 10.00M entries: 123.20M inserts, 28.16M updates, 113.20M deletes, 122.18M gets, 0.0% clog
2024/04/15 15:22:49   ... disk space: 9.2G allocated, 1.6G in use, 7.5G reclaimable, 15615.8G available
2024/04/15 15:23:03 [37%] 392.66M ops, 10.00M entries: 125.12M inserts, 28.31M updates, 115.12M deletes, 124.10M gets, 0.0% clog
2024/04/15 15:23:04   ... disk space: 9.3G allocated, 1.6G in use, 7.6G reclaimable, 15615.7G available
2024/04/15 15:23:18 [38%] 398.17M ops, 10.00M entries: 126.91M inserts, 28.45M updates, 116.91M deletes, 125.89M gets, 0.0% clog
2024/04/15 15:23:19   ... disk space: 9.4G allocated, 1.6G in use, 7.7G reclaimable, 15615.6G available
2024/04/15 15:23:33 [39%] 403.27M ops, 10.00M entries: 128.57M inserts, 28.58M updates, 118.57M deletes, 127.54M gets, 0.0% clog
2024/04/15 15:23:34   ... disk space: 9.5G allocated, 1.6G in use, 7.8G reclaimable, 15615.5G available
2024/04/15 15:23:48 [40%] 409.10M ops, 10.00M entries: 130.47M inserts, 28.73M updates, 120.47M deletes, 129.44M gets, 0.0% clog
2024/04/15 15:23:49   ... disk space: 9.6G allocated, 1.6G in use, 7.9G reclaimable, 15615.4G available
2024/04/15 15:24:03 [40%] 414.66M ops, 10.00M entries: 132.27M inserts, 28.87M updates, 122.27M deletes, 131.25M gets, 0.0% clog
2024/04/15 15:24:04   ... disk space: 9.7G allocated, 1.6G in use, 8.0G reclaimable, 15615.3G available
2024/04/15 15:24:18 [41%] 419.51M ops, 10.00M entries: 133.85M inserts, 28.99M updates, 123.85M deletes, 132.83M gets, 0.0% clog
2024/04/15 15:24:19   ... disk space: 9.7G allocated, 1.6G in use, 8.1G reclaimable, 15615.3G available
2024/04/15 15:24:33 [42%] 424.84M ops, 10.00M entries: 135.58M inserts, 29.12M updates, 125.58M deletes, 134.56M gets, 0.0% clog
2024/04/15 15:24:34   ... disk space: 9.8G allocated, 1.6G in use, 8.2G reclaimable, 15615.2G available
2024/04/15 15:24:48 [43%] 430.30M ops, 10.00M entries: 137.36M inserts, 29.25M updates, 127.36M deletes, 136.34M gets, 0.0% clog
2024/04/15 15:24:49   ... disk space: 9.9G allocated, 1.6G in use, 8.3G reclaimable, 15615.1G available
2024/04/15 15:25:03 [44%] 435.09M ops, 10.00M entries: 138.92M inserts, 29.36M updates, 128.92M deletes, 137.90M gets, 0.0% clog
2024/04/15 15:25:04   ... disk space: 10.0G allocated, 1.6G in use, 8.4G reclaimable, 15615.0G available
2024/04/15 15:25:18 [45%] 440.03M ops, 10.00M entries: 140.52M inserts, 29.47M updates, 130.52M deletes, 139.51M gets, 0.0% clog
2024/04/15 15:25:19   ... disk space: 10.1G allocated, 1.6G in use, 8.5G reclaimable, 15614.9G available
2024/04/15 15:25:33 [45%] 445.31M ops, 10.00M entries: 142.24M inserts, 29.60M updates, 132.24M deletes, 141.23M gets, 0.0% clog
2024/04/15 15:25:34   ... disk space: 10.2G allocated, 1.6G in use, 8.6G reclaimable, 15614.8G available
2024/04/15 15:25:48 [46%] 450.17M ops, 10.00M entries: 143.83M inserts, 29.71M updates, 133.83M deletes, 142.81M gets, 0.0% clog
2024/04/15 15:25:49   ... disk space: 10.3G allocated, 1.6G in use, 8.7G reclaimable, 15614.7G available
2024/04/15 15:26:03 [47%] 454.94M ops, 10.00M entries: 145.38M inserts, 29.82M updates, 135.38M deletes, 144.37M gets, 0.0% clog
2024/04/15 15:26:04   ... disk space: 10.4G allocated, 1.6G in use, 8.7G reclaimable, 15614.6G available
2024/04/15 15:26:18 [48%] 460.03M ops, 10.00M entries: 147.04M inserts, 29.93M updates, 137.04M deletes, 146.03M gets, 0.0% clog
2024/04/15 15:26:19   ... disk space: 10.4G allocated, 1.6G in use, 8.8G reclaimable, 15614.6G available
2024/04/15 15:26:33 [49%] 464.80M ops, 10.00M entries: 148.59M inserts, 30.03M updates, 138.59M deletes, 147.58M gets, 0.0% clog
2024/04/15 15:26:34   ... disk space: 10.5G allocated, 1.6G in use, 8.9G reclaimable, 15614.5G available
2024/04/15 15:26:48 [50%] 469.16M ops, 10.00M entries: 150.01M inserts, 30.13M updates, 140.01M deletes, 149.00M gets, 0.0% clog
2024/04/15 15:26:49   ... disk space: 10.6G allocated, 1.6G in use, 9.0G reclaimable, 15614.4G available
2024/04/15 15:27:03 [50%] 474.17M ops, 10.00M entries: 151.65M inserts, 30.24M updates, 141.65M deletes, 150.64M gets, 0.0% clog
2024/04/15 15:27:04   ... disk space: 10.7G allocated, 1.6G in use, 9.1G reclaimable, 15614.3G available
2024/04/15 15:27:18 [51%] 478.98M ops, 10.00M entries: 153.21M inserts, 30.34M updates, 143.21M deletes, 152.21M gets, 0.0% clog
2024/04/15 15:27:19   ... disk space: 10.8G allocated, 1.6G in use, 9.2G reclaimable, 15614.2G available
2024/04/15 15:27:33 [52%] 483.17M ops, 10.00M entries: 154.58M inserts, 30.43M updates, 144.58M deletes, 153.58M gets, 0.0% clog
2024/04/15 15:27:34   ... disk space: 10.8G allocated, 1.6G in use, 9.2G reclaimable, 15614.2G available
2024/04/15 15:27:48 [53%] 487.82M ops, 10.00M entries: 156.10M inserts, 30.53M updates, 146.10M deletes, 155.10M gets, 0.0% clog
2024/04/15 15:27:49   ... disk space: 10.9G allocated, 1.6G in use, 9.3G reclaimable, 15614.1G available
2024/04/15 15:28:03 [54%] 492.49M ops, 10.00M entries: 157.62M inserts, 30.62M updates, 147.62M deletes, 156.62M gets, 0.0% clog
2024/04/15 15:28:04   ... disk space: 11.0G allocated, 1.6G in use, 9.4G reclaimable, 15614.0G available
2024/04/15 15:28:18 [55%] 496.60M ops, 10.00M entries: 158.97M inserts, 30.71M updates, 148.97M deletes, 157.96M gets, 0.0% clog
2024/04/15 15:28:19   ... disk space: 11.1G allocated, 1.6G in use, 9.5G reclaimable, 15613.9G available
2024/04/15 15:28:33 [55%] 501.04M ops, 10.00M entries: 160.42M inserts, 30.80M updates, 150.42M deletes, 159.41M gets, 0.0% clog
2024/04/15 15:28:34   ... disk space: 11.1G allocated, 1.6G in use, 9.5G reclaimable, 15613.9G available
2024/04/15 15:28:48 [56%] 505.63M ops, 10.00M entries: 161.91M inserts, 30.89M updates, 151.91M deletes, 160.91M gets, 0.0% clog
2024/04/15 15:28:49   ... disk space: 11.2G allocated, 1.6G in use, 9.6G reclaimable, 15613.8G available
2024/04/15 15:29:03 [57%] 509.72M ops, 10.00M entries: 163.25M inserts, 30.97M updates, 153.25M deletes, 162.24M gets, 0.0% clog
2024/04/15 15:29:04   ... disk space: 11.3G allocated, 1.6G in use, 9.7G reclaimable, 15613.7G available
2024/04/15 15:29:18 [58%] 513.89M ops, 10.00M entries: 164.61M inserts, 31.06M updates, 154.61M deletes, 163.61M gets, 0.0% clog
2024/04/15 15:29:19   ... disk space: 11.4G allocated, 1.6G in use, 9.8G reclaimable, 15613.6G available
2024/04/15 15:29:33 [59%] 518.40M ops, 10.00M entries: 166.09M inserts, 31.15M updates, 156.09M deletes, 165.08M gets, 0.0% clog
2024/04/15 15:29:34   ... disk space: 11.4G allocated, 1.6G in use, 9.8G reclaimable, 15613.6G available
2024/04/15 15:29:48 [60%] 522.55M ops, 10.00M entries: 167.44M inserts, 31.23M updates, 157.44M deletes, 166.44M gets, 0.0% clog
2024/04/15 15:29:49   ... disk space: 11.5G allocated, 1.6G in use, 9.9G reclaimable, 15613.5G available
2024/04/15 15:30:03 [60%] 526.50M ops, 10.00M entries: 168.73M inserts, 31.30M updates, 158.73M deletes, 167.73M gets, 0.0% clog
2024/04/15 15:30:04   ... disk space: 11.6G allocated, 1.6G in use, 10.0G reclaimable, 15613.4G available
2024/04/15 15:30:18 [61%] 530.82M ops, 10.00M entries: 170.14M inserts, 31.39M updates, 160.14M deletes, 169.14M gets, 0.0% clog
2024/04/15 15:30:19   ... disk space: 11.7G allocated, 1.6G in use, 10.0G reclaimable, 15613.3G available
2024/04/15 15:30:33 [62%] 534.90M ops, 10.00M entries: 171.48M inserts, 31.46M updates, 161.48M deletes, 170.48M gets, 0.0% clog
2024/04/15 15:30:35   ... disk space: 11.7G allocated, 1.6G in use, 10.1G reclaimable, 15613.3G available
2024/04/15 15:30:48 [63%] 538.55M ops, 10.00M entries: 172.67M inserts, 31.53M updates, 162.67M deletes, 171.67M gets, 0.0% clog
2024/04/15 15:30:49   ... disk space: 11.8G allocated, 1.6G in use, 10.2G reclaimable, 15613.2G available
2024/04/15 15:31:03 [64%] 542.79M ops, 10.00M entries: 174.06M inserts, 31.61M updates, 164.06M deletes, 173.06M gets, 0.0% clog
2024/04/15 15:31:04   ... disk space: 11.9G allocated, 1.6G in use, 10.3G reclaimable, 15613.1G available
2024/04/15 15:31:18 [65%] 546.99M ops, 10.00M entries: 175.43M inserts, 31.69M updates, 165.43M deletes, 174.43M gets, 0.0% clog
2024/04/15 15:31:20   ... disk space: 11.9G allocated, 1.6G in use, 10.3G reclaimable, 15613.1G available
2024/04/15 15:31:33 [65%] 550.60M ops, 10.00M entries: 176.61M inserts, 31.76M updates, 166.61M deletes, 175.61M gets, 0.0% clog
2024/04/15 15:31:34   ... disk space: 12.0G allocated, 1.6G in use, 10.4G reclaimable, 15613.0G available
2024/04/15 15:31:48 [66%] 554.53M ops, 10.00M entries: 177.90M inserts, 31.83M updates, 167.90M deletes, 176.90M gets, 0.0% clog
2024/04/15 15:31:49   ... disk space: 12.1G allocated, 1.6G in use, 10.5G reclaimable, 15612.9G available
2024/04/15 15:32:03 [67%] 558.66M ops, 10.00M entries: 179.25M inserts, 31.91M updates, 169.25M deletes, 178.25M gets, 0.0% clog
2024/04/15 15:32:04   ... disk space: 12.1G allocated, 1.6G in use, 10.5G reclaimable, 15612.9G available
2024/04/15 15:32:18 [68%] 562.24M ops, 10.00M entries: 180.42M inserts, 31.97M updates, 170.42M deletes, 179.42M gets, 0.0% clog
2024/04/15 15:32:20   ... disk space: 12.2G allocated, 1.6G in use, 10.6G reclaimable, 15612.8G available
2024/04/15 15:32:33 [69%] 566.10M ops, 10.00M entries: 181.68M inserts, 32.04M updates, 171.68M deletes, 180.68M gets, 0.0% clog
2024/04/15 15:32:34   ... disk space: 12.3G allocated, 1.6G in use, 10.7G reclaimable, 15612.7G available
2024/04/15 15:32:48 [70%] 570.07M ops, 10.00M entries: 182.99M inserts, 32.11M updates, 172.99M deletes, 181.99M gets, 0.0% clog
2024/04/15 15:32:49   ... disk space: 12.3G allocated, 1.6G in use, 10.7G reclaimable, 15612.7G available
2024/04/15 15:33:03 [70%] 573.54M ops, 10.00M entries: 184.12M inserts, 32.18M updates, 174.12M deletes, 183.12M gets, 0.0% clog
2024/04/15 15:33:05   ... disk space: 12.4G allocated, 1.6G in use, 10.8G reclaimable, 15612.6G available
2024/04/15 15:33:18 [71%] 577.23M ops, 10.00M entries: 185.33M inserts, 32.24M updates, 175.33M deletes, 184.33M gets, 0.0% clog
2024/04/15 15:33:19   ... disk space: 12.5G allocated, 1.6G in use, 10.8G reclaimable, 15612.5G available
2024/04/15 15:33:33 [72%] 581.17M ops, 10.00M entries: 186.62M inserts, 32.31M updates, 176.62M deletes, 185.62M gets, 0.0% clog
2024/04/15 15:33:34   ... disk space: 12.5G allocated, 1.6G in use, 10.9G reclaimable, 15612.5G available
2024/04/15 15:33:48 [73%] 584.70M ops, 10.00M entries: 187.77M inserts, 32.37M updates, 177.77M deletes, 186.77M gets, 0.0% clog
2024/04/15 15:33:50   ... disk space: 12.6G allocated, 1.6G in use, 11.0G reclaimable, 15612.4G available
2024/04/15 15:34:03 [74%] 588.18M ops, 10.00M entries: 188.91M inserts, 32.43M updates, 178.91M deletes, 187.92M gets, 0.0% clog
2024/04/15 15:34:04   ... disk space: 12.6G allocated, 1.6G in use, 11.0G reclaimable, 15612.4G available
2024/04/15 15:34:18 [75%] 591.97M ops, 10.00M entries: 190.16M inserts, 32.50M updates, 180.16M deletes, 189.16M gets, 0.0% clog
2024/04/15 15:34:19   ... disk space: 12.7G allocated, 1.6G in use, 11.1G reclaimable, 15612.3G available
2024/04/15 15:34:33 [75%] 595.35M ops, 10.00M entries: 191.26M inserts, 32.56M updates, 181.26M deletes, 190.27M gets, 0.0% clog
2024/04/15 15:34:35   ... disk space: 12.8G allocated, 1.6G in use, 11.1G reclaimable, 15612.2G available
2024/04/15 15:34:48 [76%] 598.69M ops, 10.00M entries: 192.36M inserts, 32.61M updates, 182.36M deletes, 191.36M gets, 0.0% clog
2024/04/15 15:34:50   ... disk space: 12.8G allocated, 1.6G in use, 11.2G reclaimable, 15612.2G available
2024/04/15 15:35:03 [77%] 602.39M ops, 10.00M entries: 193.57M inserts, 32.68M updates, 183.57M deletes, 192.57M gets, 0.0% clog
2024/04/15 15:35:05   ... disk space: 12.9G allocated, 1.6G in use, 11.3G reclaimable, 15612.1G available
2024/04/15 15:35:18 [78%] 605.88M ops, 10.00M entries: 194.71M inserts, 32.74M updates, 184.71M deletes, 193.72M gets, 0.0% clog
2024/04/15 15:35:20   ... disk space: 12.9G allocated, 1.6G in use, 11.3G reclaimable, 15612.1G available
2024/04/15 15:35:33 [79%] 608.95M ops, 10.00M entries: 195.72M inserts, 32.79M updates, 185.72M deletes, 194.72M gets, 0.0% clog
2024/04/15 15:35:35   ... disk space: 13.0G allocated, 1.6G in use, 11.4G reclaimable, 15612.0G available
2024/04/15 15:35:48 [80%] 612.53M ops, 10.00M entries: 196.89M inserts, 32.85M updates, 186.89M deletes, 195.89M gets, 0.0% clog
2024/04/15 15:35:50   ... disk space: 13.1G allocated, 1.6G in use, 11.4G reclaimable, 15611.9G available
2024/04/15 15:36:03 [80%] 616.05M ops, 10.00M entries: 198.05M inserts, 32.91M updates, 188.05M deletes, 197.05M gets, 0.0% clog
2024/04/15 15:36:05   ... disk space: 13.1G allocated, 1.6G in use, 11.5G reclaimable, 15611.9G available
2024/04/15 15:36:18 [81%] 619.07M ops, 10.00M entries: 199.04M inserts, 32.96M updates, 189.04M deletes, 198.04M gets, 0.0% clog
2024/04/15 15:36:20   ... disk space: 13.2G allocated, 1.6G in use, 11.6G reclaimable, 15611.8G available
2024/04/15 15:36:33 [82%] 622.51M ops, 10.00M entries: 200.17M inserts, 33.01M updates, 190.17M deletes, 199.16M gets, 0.0% clog
2024/04/15 15:36:35   ... disk space: 13.2G allocated, 1.6G in use, 11.6G reclaimable, 15611.8G available
2024/04/15 15:36:48 [83%] 626.01M ops, 10.00M entries: 201.32M inserts, 33.07M updates, 191.32M deletes, 200.31M gets, 0.0% clog
2024/04/15 15:36:50   ... disk space: 13.3G allocated, 1.6G in use, 11.7G reclaimable, 15611.7G available
2024/04/15 15:37:03 [84%] 628.94M ops, 10.00M entries: 202.28M inserts, 33.12M updates, 192.28M deletes, 201.27M gets, 0.0% clog
2024/04/15 15:37:05   ... disk space: 13.3G allocated, 1.6G in use, 11.7G reclaimable, 15611.7G available
2024/04/15 15:37:18 [85%] 632.29M ops, 10.00M entries: 203.37M inserts, 33.17M updates, 193.37M deletes, 202.37M gets, 0.0% clog
2024/04/15 15:37:20   ... disk space: 13.4G allocated, 1.6G in use, 11.8G reclaimable, 15611.6G available
2024/04/15 15:37:33 [85%] 635.69M ops, 10.00M entries: 204.49M inserts, 33.23M updates, 194.49M deletes, 203.49M gets, 0.0% clog
2024/04/15 15:37:35   ... disk space: 13.4G allocated, 1.6G in use, 11.8G reclaimable, 15611.6G available
2024/04/15 15:37:48 [86%] 638.60M ops, 10.00M entries: 205.44M inserts, 33.27M updates, 195.44M deletes, 204.44M gets, 0.0% clog
2024/04/15 15:37:50   ... disk space: 13.5G allocated, 1.6G in use, 11.9G reclaimable, 15611.5G available
2024/04/15 15:38:03 [87%] 641.73M ops, 10.00M entries: 206.47M inserts, 33.32M updates, 196.47M deletes, 205.47M gets, 0.0% clog
2024/04/15 15:38:05   ... disk space: 13.5G allocated, 1.6G in use, 11.9G reclaimable, 15611.5G available
2024/04/15 15:38:18 [88%] 645.04M ops, 10.00M entries: 207.55M inserts, 33.38M updates, 197.55M deletes, 206.55M gets, 0.0% clog
2024/04/15 15:38:20   ... disk space: 13.6G allocated, 1.6G in use, 12.0G reclaimable, 15611.4G available
2024/04/15 15:38:33 [89%] 647.89M ops, 10.00M entries: 208.49M inserts, 33.42M updates, 198.49M deletes, 207.49M gets, 0.0% clog
2024/04/15 15:38:35   ... disk space: 13.6G allocated, 1.6G in use, 12.0G reclaimable, 15611.4G available
2024/04/15 15:38:48 [90%] 650.94M ops, 10.00M entries: 209.49M inserts, 33.47M updates, 199.49M deletes, 208.49M gets, 0.0% clog
2024/04/15 15:38:50   ... disk space: 13.7G allocated, 1.6G in use, 12.1G reclaimable, 15611.3G available
2024/04/15 15:39:03 [90%] 654.18M ops, 10.00M entries: 210.56M inserts, 33.52M updates, 200.56M deletes, 209.55M gets, 0.0% clog
2024/04/15 15:39:05   ... disk space: 13.8G allocated, 1.6G in use, 12.1G reclaimable, 15611.2G available
2024/04/15 15:39:18 [91%] 657.04M ops, 10.00M entries: 211.49M inserts, 33.56M updates, 201.49M deletes, 210.49M gets, 0.0% clog
2024/04/15 15:39:20   ... disk space: 13.8G allocated, 1.6G in use, 12.2G reclaimable, 15611.2G available
2024/04/15 15:39:33 [92%] 659.95M ops, 10.00M entries: 212.45M inserts, 33.61M updates, 202.45M deletes, 211.45M gets, 0.0% clog
2024/04/15 15:39:35   ... disk space: 13.9G allocated, 1.6G in use, 12.2G reclaimable, 15611.1G available
2024/04/15 15:39:48 [93%] 663.08M ops, 10.00M entries: 213.47M inserts, 33.66M updates, 203.47M deletes, 212.47M gets, 0.0% clog
2024/04/15 15:39:50   ... disk space: 13.9G allocated, 1.6G in use, 12.3G reclaimable, 15611.1G available
2024/04/15 15:40:03 [94%] 665.94M ops, 10.00M entries: 214.41M inserts, 33.70M updates, 204.41M deletes, 213.41M gets, 0.0% clog
2024/04/15 15:40:05   ... disk space: 14.0G allocated, 1.6G in use, 12.3G reclaimable, 15611.0G available
2024/04/15 15:40:18 [95%] 668.78M ops, 10.00M entries: 215.35M inserts, 33.74M updates, 205.35M deletes, 214.35M gets, 0.0% clog
2024/04/15 15:40:20   ... disk space: 14.0G allocated, 1.6G in use, 12.4G reclaimable, 15611.0G available
2024/04/15 15:40:33 [95%] 671.79M ops, 10.00M entries: 216.33M inserts, 33.79M updates, 206.33M deletes, 215.33M gets, 0.0% clog
2024/04/15 15:40:35   ... disk space: 14.1G allocated, 1.6G in use, 12.4G reclaimable, 15610.9G available
2024/04/15 15:40:48 [96%] 674.54M ops, 10.00M entries: 217.24M inserts, 33.83M updates, 207.24M deletes, 216.24M gets, 0.0% clog
2024/04/15 15:40:50   ... disk space: 14.1G allocated, 1.6G in use, 12.5G reclaimable, 15610.9G available
2024/04/15 15:41:03 [97%] 677.19M ops, 10.00M entries: 218.11M inserts, 33.87M updates, 208.11M deletes, 217.10M gets, 0.0% clog
2024/04/15 15:41:05   ... disk space: 14.1G allocated, 1.6G in use, 12.5G reclaimable, 15610.9G available
2024/04/15 15:41:18 [98%] 680.03M ops, 10.00M entries: 219.04M inserts, 33.91M updates, 209.04M deletes, 218.04M gets, 0.0% clog
2024/04/15 15:41:20   ... disk space: 14.2G allocated, 1.6G in use, 12.6G reclaimable, 15610.8G available
2024/04/15 15:41:33 [99%] 682.79M ops, 10.00M entries: 219.94M inserts, 33.95M updates, 209.94M deletes, 218.94M gets, 0.0% clog
2024/04/15 15:41:35   ... disk space: 14.2G allocated, 1.6G in use, 12.6G reclaimable, 15610.8G available
2024/04/15 15:41:48 Verifying contents ...
2024/04/15 15:41:58 SUCCESS! All values matched between reference map and infinimap
```