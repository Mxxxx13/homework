package main

import "fmt"

type Author struct {
	Name string				//名字
	VIP bool 				//是否是高贵的带会员
	Icon string 			//头像
	Signature string 		//签名
	Focus int 				//关注人数
}

type Video struct {
	Title string 			//标题
	View int 	     		//播放量
	Comments int		    //弹幕数
	Time string				//时间
	HighestRank string 		//最高排名
}

type RelatedRecommendations struct {		//相关推荐结构体
	Title string							//标题
	Author1 string							//作者
	View int 								//播放量
	Comment int 							//弹幕
}

var s []RelatedRecommendations				//声明相关推荐切片

func main(){
	AA:=Author{
		Name: "哦呼w",
		VIP: true,
		Icon: "不知道怎么表示",
		Signature:"一个做鬼畜的丨 合作请私信 ",
		Focus:1149000,
	}
	VV:=Video{
		Title: "敢 杀 我 的 马？！", 					//标题
		View: 28877000, 	     					//播放量
		Comments: 110000,	    					//弹幕数
		Time: "2020-07-11 12:00:48",				//时间
		HighestRank: "全站排行榜最高第1名 ",			//最高排名
	}
	re1 :=RelatedRecommendations {
		Title: "让子弹飞",
	}
	re2 :=RelatedRecommendations{
		Title: "【春晚鬼畜】赵本山：我就是念诗之王！【改革春风吹满地】",
		Author1: "UP-Sings",
		View: 69657000,
		Comment: 519000,
	}
	re3 :=RelatedRecommendations{
		Title: "the real suger baby已知最高画质",
		Author1: "一嗷垚",
		View: 23682000,
		Comment: 20000,
	}
	s = []RelatedRecommendations{re1,re2,re3}
	fmt.Println(AA)
	fmt.Println(VV)
	fmt.Println(s)
}
