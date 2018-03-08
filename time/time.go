package main

import (
	"time"
	"fmt"
)

// 通过字符串类型的时间(yyyy-MM-dd HH:mm:SS)得到int64类型的时间
func String2TimeT(pszTime string) int64{
	t2,_ := time.Parse("2006-01-02 15:04:05",pszTime)
	return t2.Unix()
}

// 通过时间戳获得字符串类型的时间(yyyy-MM-dd HH:mm:SS)
func TimeT2String(unixTime int64) string{
	tm := time.Unix(unixTime, 0)
	return tm.Format("2006-01-02 15:04:05")
}

func main1()  {
	now := time.Now()
	local1, err1 := time.LoadLocation("") //等同于"UTC"
	if err1 != nil {
		fmt.Println(err1)
	}
	local2, err2 := time.LoadLocation("Local")//服务器设置的时区
	if err2 != nil {
		fmt.Println(err2)
	}
	local3, err3 := time.LoadLocation("America/Los_Angeles")
	if err3 != nil {
		fmt.Println(err3)
	}

	fmt.Println(now.In(local1))
	fmt.Println(now.In(local2))
	fmt.Println(now.In(local3))
}
func main()  {


	//获取本地location
	toBeCharge := "2018-01-03 09:00:00"                             //待转化为时间戳的字符串 注意 这里的小时和分钟还要秒必须写 因为是跟着模板走的 修改模板的话也可以不写
	timeLayout := "2006-01-02 15:04:05"                             //转化所需模板
	loc, _ := time.LoadLocation("Local")                            //重要：获取时区
	theTime, _ := time.ParseInLocation(timeLayout, toBeCharge, loc) //使用模板在对应时区转化为time.time类型
	sr := theTime.Unix()                                            //转化为时间戳 类型是int64
	fmt.Println(theTime)                                            //打印输出theTime 2015-01-01 15:15:00 +0800 CST
	fmt.Println(sr)                                                 //打印输出时间戳 1420041600

	//时间戳转日期
	dataTimeStr := time.Unix(sr, 0).Format(timeLayout) //设置时间戳 使用模板格式化为日期字符串
	fmt.Println(dataTimeStr)
}
func main2() {
	//time.Time代表一个纳秒精度的时间点
	/*t5 := time.Now()

	fmt.Printf("%+v\n", t5);
	 // t5.Unix() 返回时间戳
	fmt.Println(t5.Unix())

	timeStamp := t5.Unix()

    // 时间戳转化成年月日时分秒格式为"2006-01-02 15:04:05"
	tm := time.Unix(timeStamp, 0)
	timeString := tm.Format("2006-01-02 15:04:05")*/
    //
	timeString := "2018-01-03 09:00:00"
	// 字符串转化为时间戳
	t2,_ := time.Parse("2006-01-02 15:04:05",timeString)
	timeStamp2 := t2.Unix()
    ///


	fmt.Println(timeStamp2)

	tm := time.Unix(timeStamp2, 0)
	timeString2 := tm.Format("2006-01-02 15:04:05")
	fmt.Println(timeString2)
	//返回当前时间
	/*t = time.Now();
	fmt.Printf("%v\n", t);
	//反回所在时区
	fmt.Printf("%v\n", t.Location());
	//返回UTC时间和UTC时区
	fmt.Printf("%v %v\n", t.UTC(), t.UTC().Location());
	//同上，In()返回指定时区的时间
	fmt.Printf("%v %v\n", t.In(time.UTC), t.In(time.UTC).Location());
	//返回当地时区的时间
	fmt.Printf("%v %v\n", t.Local(), t.Local().Location());

	//根据时间戳返回本地时间
	//参数分别表示秒数和纳秒数
	t2 := time.Unix(1487780010, 0);
	fmt.Println(t2);

	//根据指定时间返回time.Time
	//分别指定年，月，日，时，分，秒，纳秒，时区
	t3 := time.Date(2017, time.Month(5), 26, 15, 30, 20, 0, t.Location());
	fmt.Println(t3);

	//格式化输出时间
	t4 := time.Now();
	fmt.Println(t4.Format("2006-01-02 15:04:05"));

	//获取时间信息
	t5 := time.Now();
	//返回日期
	fmt.Println(t5.Date());
	//返回年
	fmt.Println(t5.Year());
	//返回月
	fmt.Println(t5.Month());
	//返回日
	fmt.Println(t5.Day());
	//返回星期
	fmt.Println(t5.Weekday());
	//返回ISO 9601标准下的年份和星期编号
	fmt.Println(t5.ISOWeek());
	//返回时分秒
	fmt.Println(t5.Clock());
	//返回小时
	fmt.Println(t5.Hour());
	//返回分钟
	fmt.Println(t5.Minute());
	//返回秒
	fmt.Println(t5.Second());
	//返回纳秒
	fmt.Println(t5.Nanosecond());
	//返回一年中对应的天
	fmt.Println(t5.YearDay());
	//返回时区
	fmt.Println(t5.Location());
	//返回时区的规范名,时区相对于UTC的时间偏移量(秒)
	fmt.Println(t5.Zone());
	//返回时间戳
	fmt.Println(t5.Unix());
	//返回纳秒时间戳
	fmt.Println(t5.UnixNano());

	//时间的比较与计算
	t6 := time.Now();
	//是否零时时间
	fmt.Println(t6.IsZero());
	//t6时间在t5时间之后，返回真
	fmt.Println(t6.After(t5));
	//t5时间在t6时间之前，返回真
	fmt.Println(t5.Before(t6));
	//时间是否相同
	fmt.Println(t6.Equal(t6));
	//返回t6加上纳秒的时间
	fmt.Println(t6.Add(10000));
	//返回两个时间之差的纳秒数
	fmt.Println(t6.Sub(t5));
	//返回t6加1年，1月，1天的时间
	fmt.Println(t6.AddDate(1, 1, 1));

	//时间的序列化
	t7 := time.Now();
	//序列化二进制
	bin, _ := t7.MarshalBinary();
	//反序列化二进制
	t7.UnmarshalBinary(bin)
	fmt.Println(t7);
	//序列化json
	json, _ := t7.MarshalJSON();
	fmt.Println(string(json));
	//反序列化json
	t7.UnmarshalJSON(json);
	fmt.Println(t7);
	//序列化文本
	txt, _ := t7.MarshalText();
	fmt.Println(string(txt));
	//反序列化文本
	t7.UnmarshalText(txt);
	fmt.Println(t7);
	//gob编码
	gob, _ := t7.GobEncode();
	t7.GobDecode(gob);
	fmt.Println(t7);

	//时间段time.Duration
	dur := time.Duration(6666666600000);
	//返回字符串表示
	fmt.Println(dur.String());
	//返回小时表示
	fmt.Println(dur.Hours());
	//返回分钟表示
	fmt.Println(dur.Minutes());
	//返回秒表示
	fmt.Println(dur.Seconds());
	//返回纳秒表示
	fmt.Println(dur.Nanoseconds());

	//时区time.Location
	//返回时区名
	fmt.Println(time.Local.String());

	//通过地点名和时间偏移量返回时区
	fmt.Println(time.FixedZone("Shanghai", 800));

	//通过给定时区名称，返回时区
	loc, _ := time.LoadLocation("Asia/Shanghai");
	fmt.Println(loc);

	//阻塞当前进程3秒
	time.Sleep(time.Second * 3);

	//定时器time.Timer
	//创建一个1秒后触发定时器
	timer1 := time.NewTimer(time.Second * 1);
	<-timer1.C;
	fmt.Println("timer1 end");

	//1秒后运行函数
	time.AfterFunc(time.Second * 1, func() {
		fmt.Println("wait 1 second");
	});
	time.Sleep(time.Second * 3);

	//打点器time.Ticker
	//创建一个打点器，在固定1秒内重复执行
	ticker := time.NewTicker(time.Second);
	num := 1;
	for {
		if num > 5 {
			//大于5次关闭打点器
			ticker.Stop();
			break;
		}
		//否则从打点器中获取chan
		select {
		case <-ticker.C:
			num++;
			fmt.Println("1 second...");
		}
	}*/
}
