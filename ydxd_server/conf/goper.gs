#### 配置模块的配置项
## 1. include: 表示配置文件包含其他配置文件，会从include中存在的文件获取配置项的值
##    * include关键字可以出现在配置文件的任意位置，也包括任何其他配置文件的任意位置
##    * 路径值为绝对路径，或者是配置文件相对配置文件夹的相对路径
##    * 如果有多个文件则配置为一个字符串数组
##    * 注意: gson的:是一个键构造字符，因此如果路径中包含:需要用`进行包含或者用%进行转义
## 默认值: 空
## 配置样例: 
# include: # `E:\7.Goper\Goper\conf\host\color.ydxd.club.host.gs`  # 绝对路径配置，用:包住字符串
           # E%:\7.Goper\Goper\conf\host\color.ydxd.club.host.gs   # 绝对路径配置，用%转义:
           # host\color.host.gs   # 相对路径配置，相对于 $(安装路径)\$(配置文件相对安装路径的相对路径) 的相对路径

#### 日志模块的配置项
## 1. log: 与运行时日志相关的配置，配置为一个字典，用{}将各个具体的值包起来
log: {
	## 1. level: 默认输出的日志等级，配置为一个数字，取值和说明如下
	##    * 0 不设置等级，所有日志都输出 
	##    * 1	LevelDebug 输出Debug及以上等级的日志（Debug、Info、Warn、Error、Critical）
	##    * 2	LevelInfo 输出Info及以上等级的日志（Info、Warn、Error、Critical）
	##    * 3	LevelWarn 输出Warn及以上等级的日志（Warn、Error、Critical）
	##    * 4	LevelError 输出Error及以上等级的日志（Error、Critical）
	##    * 5	LevelCritical 输出Critical及以上等级的日志（Critical）
	##    * > 5 所有日志都不输出
	## 默认值: 3
	level: 1

	## 2. tag: 日志输出附带的默认标签，配置为一个字符串，如果不设置可直接注释
	##    * 只有在flag设置了LTag或者LPreTag，tag才会在日志中输出
	## 默认值: 空字符串
	# tag: [goper]

	## 3. flag: 日志中要输出的东西对应的相关的flag标志，配置为一个字符串，如果要输出多个内容，就用|将多个进行或运算
	##    * 注意: |前后不要留空格，因为gson对空格敏感，如果出现空格，后面的值会被gson舍弃，或者用`将字符串包起来
	##    * LDate 输出当地的时区的日期 
	##    * LTime 输出当地的时区的时间
	##    * LMicroseconds 输出当地的时区的时间时，精确到微秒，这个设置了默认LTime也设置
	##    * LNanosceonds 输出当前时间戳的纳秒数，注意这里是输出总纳秒数，可用于在一段代码前后输出查看对应代码运行的时间耗时
	##    * LLongFile 输出调用日志接口所在长文件名称（即包含全路径和文件名）
	##    * LShortFile 输出调用日志接口所在短文件名称（即仅仅是文件名），如果LLongFile和LShortFile同时设置了，将只输出长文件名称
	##    * LUTC 输出时间时，将时间转换为UTC时间，这个要配合LDate或LTime或LMicroseconds使用才有效
	##    * LTag 输出tag字段设置的字符串
	##    * LPreTag 将tag字段设置的字符串输出到最前面
	##    * LLevel 输出日志的等级
	##    * LStdFlags等于LDate|LTime
	## 默认为: LStdFlags|LLevel
	## 配置样例: 
	# flag: LStdFlags|LLevel  # |前后不包含空格，不需要用`包起来
	      #  `LStdFlags | LLevel`   # |前后包含了空格，需要用`包起来
	
	## 4. logfile: 输出的日志文件，填写全路径，配置为一个字符串，如果不设置，goper将默认不输出到日志文件，只输出到标准错误
	# logfile: `E:\7.Goper\Goper\log\goper.log`  # `/Users/wxj/workplace/5.Work/7.Goper/Goper/log/goper.log`
}

## http模块的配置项
## 在goper中，http web服务分为3层，第一层为整个goper服务，整个goper服务支持多个域名（两个域名的域名相同，端口号不同，也是不同域名），一个域名支持多个路由配置
## 对应的，http配置也分为3级，第一级为http下直属的配置项，是整个服务级别的配置项，称为main级别配置项
## 第二级为http下的server关键字下的配置项，是为某一域名服务支持的配置项，称为server级别配置项
## 第三级位server下的location关键字下的配置项，是为某一个路由支持的配置项，称为location级别配置项
## http囊括了所有http模块的配置项，因此，它也是按照模块来设置配置项的
## 实际运行中，goper会将各个模块关注的main、server、location级别配置项分发给各个模块使用，具体如何使用这3个级别的配置项，则由模块的行为自行定义
## main级别是针对所有域名都会用到的配置项，server级别是针对某一个域名或者某一组域名的个性化配置项，而location则是比域名更低一级的路由级别的个性化配置项
## 有些模块没有个性化定制server级别或者location级别的需求，有些模块则不需要main级别的配置项，只需要server或location级别的配置项，所以并非所有模块都是3个级别都有配置项的
## 他们的使用规则也不尽相同，因此需要根据具体模块来确定配置项的具体使用规则
http: {	
	## http下可以用include关键字将配置项独立到其他配置文件中
	## include也可以包含在某个server或者location，那么此时对应配置文件的配置属于对应的server或location级别配置项
	# include: demo/http_demo.gs   # http_demo模块的配置项，这是一个测试模块，正式使用时直接删除即可
	include: host/ydxd.club.gs 

	## server级别的配置项是为一个域名或一组域名定义个性化的配置
	## http下可以有多个server并存，但是不能存在两个或多个server的name一样
	# server: {
		## server级别的服务的名称
		# name: server1

		## location级别的配置项是为一个路由或者一组路由定义个性化的配置
		## 一个server可以有多个location并存，但是不能存在两个或者多个location的name一样
		# location: {
			## loc级别的服务名称
			name: location1
		# }
	# }
}