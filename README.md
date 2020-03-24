# PortScan - Golang
### 如何设置线程数
```
export goroutine_count=20000    // 这里即是设置2w协程
```

### 如何设置超时时间
```
export timeout=5              // 这里即是5秒超时
```

### ip文件怎么读?
在`PortScan_Linux`可执行文件的同级目录放入名称为`ip.txt`, 即可读取.