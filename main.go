package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"html/template"
	"log"
	"net/http"
)

type user struct {
	Name string
	Age int
}
type city struct {
	id int `json:"id" bson:"id"`
	cityid int `json:"cityid" bson:"cityid"`
	agency_name string `json:"agency_name" bson:"agency_name"`
}

func index(w http.ResponseWriter, r *http.Request) {
	// 往w里写入内容，就会在浏览器里输出
	fmt.Fprintf(w, "Hello golang http!")
}
func list(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql","root:jinsan@tcp(127.0.0.1:3306)/test?charset=utf8")
	if err != nil{
		fmt.Fprintf(w,"数据库链接错误")
		db.Close();
	}
	rows, err2 := db.Query("select id,cityid,agency_name from zg_city_info limit 1,10")
	if err2 != nil {
		print("sql执行失败")
	}
	tem,err := template.ParseFiles("html/list.html")
	if err != nil{
		fmt.Println("读取文件失败,err",err)
		return
	}
	//fmt.Fprintf(w, "这是列表页")
	//head := "id|cityid|agency_name\n"
	//fmt.Fprintf(w,"%s\n",head)

	var citys []city
	for rows.Next() {

		var zg_city_info city
		if err := rows.Scan(&zg_city_info.id,&zg_city_info.cityid,&zg_city_info.agency_name); err != nil {
			log.Fatal(err)
		}
		//fmt.Println(zg_city_info)
		citys = append(citys, zg_city_info)
		//s := "" + strconv.Itoa(id) + "|" + strconv.Itoa(cityid) + "|" + agency_name + "\n"
		//fmt.Printf(" %d\n",  cityid)
		//fmt.Fprintf(w, " %s\n", s) 直接显示

	}
	//fmt.Println(citys)
	//for _,v := range citys{
	//	fmt.Println(v.cityid,v.agency_name)
	//}
	tem.Execute(w,citys)
}


func testhtml(w http.ResponseWriter, r *http.Request){
	tem,err := template.ParseFiles("html/test.html")
	if err != nil{
		fmt.Println("读取文件失败,err",err)
		return
	}
	// 利用给定数据渲染模板，并将结果写入w
	po := user{
		Name:"Jinshengqiang",
		Age:28,
	}
	// 利用给定数据渲染模板，并将结果写入w
	tem.Execute(w,po)
}
func qiepian(w http.ResponseWriter,r *http.Request){
	slice := make([]int,3,5)
	fmt.Println(slice);
}
func main() {
	//db,err = sql.Open("mysql","root:jinsan/test?charset=utf8");
	// 设置路由，如果访问/，则调用index方法
	http.HandleFunc("/", index)
	http.HandleFunc("/list", list)
	http.HandleFunc("/testhtml", testhtml)
	http.HandleFunc("/qiepian",qiepian)
	// 启动web服务，监听9090端口
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}