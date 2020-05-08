package main

import	(
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var(
	user="teacher" 
	password="teacher"  
	database="fdlibrary"
)


func (c *library)createtables()error{
c.db.Exec(`DROP TABLE IF EXISTS  正在借阅表`)
c.db.Exec(`DROP TABLE IF EXISTS  历史借阅表`)
c.db.Exec(`DROP TABLE IF EXISTS  等待审核学生`)

c.db.Exec(`DROP TABLE IF EXISTS  administrator`)
c.db.Exec(`DROP TABLE IF EXISTS  图书`)
c.db.Exec(`DROP TABLE IF EXISTS  学生`)


sq1:=`
create table 图书(
ISBN varchar(20) not null,					   
书名 varchar(50) not null,
馆藏总数 integer not null,
可借数目 integer not null,
primary key(ISBN))`
_,err := c.db.Exec(sq1)
if err!=nil{
return err}

sq2:=`
create table 学生(
学号 varchar(15) not null,
姓名 varchar(10) not null,
借阅次数 integer not null,
是否可借 varchar(5) not null,
password varchar(20) not null,
primary key(学号))`

_,err2 := c.db.Exec(sq2)
if err2!=nil{
return err2}

sq3:=`
create table 正在借阅表(
学号 varchar(15) not null,ISBN varchar(20) not null,
借阅时间 date not null,
最晚归还时间 date not null,
已申请续借次数 integer not null,
primary key (学号,ISBN),
foreign key(学号)references 学生(学号),
foreign key(ISBN)references 图书(ISBN))`

_,err3 := c.db.Exec(sq3)
if err3!=nil{
return err3}

sq4:=`
create table 历史借阅表(
学号 varchar(15) not null,ISBN varchar(20) not null,
借阅时间 date not null,
归还时间 date not null,
primary key (学号,ISBN,借阅时间),
foreign key(学号)references 学生(学号),
foreign key(ISBN)references 图书(ISBN))`

_,err4 := c.db.Exec(sq4)
if err4!=nil{
return err4}

sq5:=`
create table 等待审核学生(
学号 varchar(15) not null,ISBN varchar(20) not null,
申请时间 date not null,借or还 varchar(1) not null,
primary key (学号,ISBN,借or还),
foreign key(学号)references 学生(学号),
foreign key(ISBN)references 图书(ISBN))`

_,err5 := c.db.Exec(sq5)
if err5!=nil{
return err5}

sq6:=`
create table administrator(
工号 varchar(15) not null,
password varchar(20) not null,
primary key(工号))`

_,err6 := c.db.Exec(sq6)
if err6!=nil{
return err6}

sq7:=`INSERT INTO administrator(工号,password)VALUES ('1','12345')`
_,err7:=c.db.Exec(sq7)
if err7!=nil{
return err7}
return nil
}
//以下为管理员相关函数

func printfunction(){					//输出管理员所有权限
	fmt.Println("\n请选择您要进行的操作 0X")
	fmt.Println("Function1:录入学生信息 01")
	fmt.Println("Function2:查看学生信息 02")
	fmt.Println("Function3:查看指定书籍 03")
	fmt.Println("Function4:录入图书信息 04")
	fmt.Println("Function5:查看借阅超期图书 05")
	fmt.Println("Function6:修改图书信息 06")
	fmt.Println("Function7:暂停学生服务 07")
	fmt.Println("Function8:借阅审核 08")
	fmt.Println("Function9:还书管理 09")

	fmt.Println("exit 退出本系统")
}

func (c *library)addstudent(student string,name string) error{  //添加学生账户
     _,err:=c.db.Exec("INSERT INTO 学生(学号,姓名,借阅次数,是否可借,password)VALUES (?,?,0,'Y','12345')",student,name)
    if err != nil {
        fmt.Println("信息不符合系统要求 添加失败！")
        return err
    }  
    return err
}

func (c *library)findstudent(student string)error{				//查找学生信息
	var name string
  	var lendnum int
   	var can string
	err:=c.db.QueryRow("select 姓名,借阅次数,是否可借 from 学生 where 学号=?", student).Scan(&name,&lendnum,&can)
	if err != nil {
    	fmt.Println("没有找到该学生")
		return err
	}else{
	fmt.Printf("\n %s %s  共借阅%d次   是否可借：%s",name,student,lendnum,can) 
	}
	return err
}


func (c *library)findbook(name string) bool{					//查找图书信息
    var allnum int
    var nownum int
	var ISBN string
	rows,err := c.db.Query("SELECT ISBN,馆藏总数,可借数目 FROM 图书 WHERE 书名 = ?", name)
	if err != nil {
    	fmt.Println("查询失败")
		panic(err)
	}else{
		var count int=0
		for rows.Next() {
		var err error
    		if err = rows.Scan(&ISBN,&allnum,&nownum); err == nil {
    			fmt.Printf("ISBN:%s  %s  馆藏总数:%d  可借数目:%d\n",ISBN,name,allnum,nownum)
				count=count+1
			}
		}
		if count==0{
		fmt.Print("没有找到该图书")
		return false
		}
	}
	return true
}

func  (c *library)addbook(name string,ISBN string)error{       //添加图书信息
    err := c.db.QueryRow("SELECT ISBN,书名,馆藏总数,可借数目 FROM 图书 WHERE ISBN = ?", ISBN)
	
	if err != nil {
		_,err:=c.db.Exec("INSERT INTO 图书(ISBN,书名,馆藏总数,可借数目)VALUES (?,?,1,1)",ISBN,name)
		if err!=nil{fmt.Printf("添加失败")
		return err
		}
		fmt.Printf("成功添加!")
		return nil
	}else{
		_,err:=c.db.Exec("UPdate 图书 set 馆藏总数=馆藏总数+1,可借数目=可借数目+1 where ISBN=? ",ISBN)	
		if err!=nil{
		panic(err)
		return err}		
		fmt.Printf("您已经成功添加 %s",name)
		return nil
	}
}


func  (c *library)overtimebook()error{								//查看超期图书
	var ISBN string
	var name string
	var student string
	var stuname string
	rows,err := c.db.Query("SELECT 图书.ISBN,书名,正在借阅表.学号,姓名 FROM 正在借阅表,图书,学生 WHERE 图书.ISBN=正在借阅表.ISBN AND 学生.学号=正在借阅表.学号 AND (TO_DAYS(now())-TO_DAYS(最晚归还时间))>0 ")
	if err != nil {
    	fmt.Println("查询错误")
		return err
	}else{
		for rows.Next() {
    	if err := rows.Scan(&ISBN,&name,&student,&stuname); err == nil {
    	fmt.Printf("ISBN %s  %s\n借阅者 %s,%s",ISBN,name,student,stuname)
		}
		}
	}
	return nil
}

func  (c *library)changebook(ISBN string,name string )int{				//更改图书信息
	err := c.db.QueryRow("SELECT ISBN FROM 图书 WHERE ISBN = ?", ISBN).Scan(&ISBN)
	if err != nil {
		fmt.Println("该图书不存在")
		return 1
	}else{
		_,err:=c.db.Exec("UPdate 图书 set 书名=? where ISBN=? ",name,ISBN)
		if err!=nil{
			return 0
		}
		fmt.Printf("%s 的 新书名为 %s",ISBN,name)
		return 2
	}
}


func  (c *library)stopstudentid()error{							//暂停学生账户服务
	var student string
	rows,err := c.db.Query("select 学号 from 正在借阅表 where (TO_DAYS(now())-TO_DAYS(最晚归还时间))<0 group by 学号 having count(*)>=3 ")
	if err != nil {
        fmt.Println("查询错误")
		return err
	}else{
	fmt.Println("以下几位学生借阅过期图书超过三本")
	for rows.Next() {
    	if err := rows.Scan(&student); err == nil {
    	fmt.Printf("%s\n",student)
		fmt.Println("是否暂停这个学生的帐号 Y")
		var yn string
		fmt.Scanf("%s",&yn)
		if yn=="Y"{
			fmt.Printf("sas")
    		_,err:=c.db.Exec("UPdate 学生 set 是否可借='N'where 学号=?",student)
			if err!=nil{
				return err
				fmt.Printf("已暂停学生%s\n",student)
				return nil
			}
		}
		}
	}
	}
	return nil
}



func  (c *library)borrow()error{				//审核学生借阅信息
var student string
var ISBN string
var name string
rows,err := c.db.Query("select 学号,图书.ISBN,书名 from 等待审核学生,图书 where 图书.ISBN=等待审核学生.ISBN AND 借or还='0' ")
if err != nil {
        fmt.Println("查询失败")
	
}else{
var count int=0
for rows.Next() {
    if err := rows.Scan(&student,&ISBN,&name); err == nil {
	count=count+1    
	_,err:=c.db.Exec("INSERT INTO 正在借阅表(学号,ISBN,借阅时间,最晚归还时间,已申请续借次数)VALUES (?,?,now(),date_add(now(),interval 60 day),0)",student,ISBN)
	if err!=nil{
			return err}
	_,errrr:=c.db.Exec("delete from 等待审核学生 where 学号=? AND ISBN=?",student,ISBN)
	if errrr!=nil{
			return errrr}
	_,errr:=c.db.Exec("UPdate  学生 set 借阅次数=借阅次数+1 where 学号=?",student)	
	if errr!=nil{ 
			return errr}
	}
	fmt.Printf("学生%s借阅图书%s 审核通过\n",student,ISBN)
	}
if count==0{
fmt.Printf("暂无借阅申请")
}}
return nil
}

func  (c *library)returnbook()error{                               //审核还书申请
var student string	
var ISBN string
var lendtime string
rows,err := c.db.Query("select 正在借阅表.学号,正在借阅表.ISBN,借阅时间 from 等待审核学生,正在借阅表 where 正在借阅表.学号=等待审核学生.学号 AND 正在借阅表.ISBN=等待审核学生.ISBN AND 借or还='1' ")
if err != nil {
        fmt.Println("查询失败")
	
}else{
var count int=0
for rows.Next() {
    if err := rows.Scan(&student,&ISBN,&lendtime); err == nil {
	count=count+1    
	_,err:=c.db.Exec("INSERT INTO 历史借阅表(学号,ISBN,借阅时间,归还时间)VALUES (?,?,?,now())",student,ISBN,lendtime)
	if err!=nil{
			return err}
	_,errrr:=c.db.Exec("delete from 等待审核学生 where 学号=? AND ISBN=?",student,ISBN)
	if errrr!=nil{panic(err)
			return errrr}
	fmt.Printf("学生%s归还图书%s 审核通过\n",student,ISBN)
	}
}
if count==0{
fmt.Printf("暂无还书申请")}
}
return nil
}


func librarymanage(){                 //管理员主函数
	var user1 string
    var password1 string
	var student string
	var name string
	var L library
	var ISBN string
	var bookname string
	L.connect()
	for true{
        fmt.Printf("输入exit退出本系统\n")
        fmt.Println("工号：")
        fmt.Scanf("%s",&user1)
        if user1=="exit" {
        return 
        }
	fmt.Println("密码：")
        fmt.Scanf("%s",&password1)
	err:=L.db.QueryRow("select * from administrator where 工号=? AND password=?",user1,password1).Scan(&user1,&password1)
	if err!=nil{
	fmt.Printf("工号或密码错误")
	}else{
	fmt.Printf("登录成功！")
	break
	}}
	fmt.Println("\n**主菜单**\n输入help查询所有功能")
	for true{
	fmt.Printf("\n->")
	var zzz string
	fmt.Scanf("%s",&zzz)
	switch zzz{
	case "exit":return
	case "help":printfunction()
	case "01":{fmt.Println("请输入学号：")
		  fmt.Scanf("%s",&student)
  		  fmt.Println("请输入姓名：")
   		  fmt.Scanf("%s",&name)
		  L.addstudent(student,name)} 
	case "02":{fmt.Println("请输入学号：")
			fmt.Scanf("%s",&student)
			L.findstudent(student)}
	case "03":{fmt.Println("请输入书名：")
				fmt.Scanf("%s",&bookname)
				L.findbook(bookname)}
	case "04":{fmt.Println("请输入书名：")
				fmt.Scanf("%s",&bookname)
				fmt.Println("请输入ISBN：")
				fmt.Scanf("%s",&ISBN)
				L.addbook(bookname,ISBN)}
	case "05":L.overtimebook()


	case "06":{fmt.Println("请输入要修改的图书ISBN编号")
				fmt.Scanf("%s",&ISBN)
				fmt.Println("请输入新的书名：")
				fmt.Scanf("%s",&bookname)
				L.changebook(ISBN,bookname)}
	case "07":L.stopstudentid()
	case "08":L.borrow()


	case "09":L.returnbook()
	default: fmt.Println("抱歉！本系统目前暂不支持此功能！")
	}
	}
}


//以下为学生相关函数
func (c *library)printstufunction(){
        fmt.Println("\n请选择您要进行的操作 0X")
        fmt.Println("Function1:查看个人信息 01")
        fmt.Println("Function2:修改个人信息 02")
        fmt.Println("Function3:目前借阅书籍 03")
        fmt.Println("Function4:历史借阅书籍 04")
        fmt.Println("Function5:查询图书信息 05")
        fmt.Println("Function6:申请借阅书籍 06")
        fmt.Println("Function7:申请归还书籍 07")
        fmt.Println("Function8:申请延长还书期限 08")
	fmt.Println("Function9:更改密码 09")
        fmt.Println("exit 退出本系统")
}

func (c *library)myinformation(student string)error{			//查询个人信息
	var name string
	var lendnum int
	var can string
	err:=c.db.QueryRow("select 姓名,借阅次数,是否可借 from 学生 where 学号=?", student).Scan(&name,&lendnum,&can)
	if err != nil {
    		return err
	}else{
	fmt.Printf("\n %s %s 共借阅%d次 是否可借:%s\n",name,student,lendnum,can) 
	return nil
	}
}

func (c *library)changeinformation(student string,name string)error{   //修改个人信息
 	var nameb string
	err:=c.db.QueryRow("select 姓名 from 学生 where 学号=?", student).Scan(&nameb)
	if err!=nil{
	return err	
		}
	fmt.Printf("修改前个人信息:%s %s \n",nameb,student)
	_,err=c.db.Exec("UPdate 学生 set 姓名=? where 学号=? ",name,student)
	if err!=nil{
	return err}
	fmt.Printf("修改后个人信息:%s %s \n",name,student)
	return nil
}




func (c *library)nowlend(student string)error{				//正在借阅的图书

fmt.Printf("\t\t<学号 %s 目前借阅的图书如下>\n",student)
var ISBN string
var name string
var lendtime string
var returntime string
var moretime int
rows,err := c.db.Query("SELECT 图书.ISBN,图书.书名,正在借阅表.借阅时间,正在借阅表.最晚归还时间,正在借阅表.已申请续借次数 FROM 正在借阅表,图书 WHERE 图书.ISBN=正在借阅表.ISBN AND 正在借阅表.学号=?",student)
if err != nil {
    panic(err)
    fmt.Println("查询错误")
   return nil
}else{
	var count int=0
	for rows.Next() {
    		if err := rows.Scan(&ISBN,&name,&lendtime,&returntime,&moretime); err == nil {
		count=count+1
    		fmt.Printf("ISBN:%s %s 借阅时间:%s 最晚归还时间:%s 已申请续借次数:%d\n",ISBN,name,lendtime,returntime,moretime)
		}
	}
	if count==0{
	fmt.Printf("您目前暂无正在借阅的图书")
	}
}
return nil
}


func (c *library)beforelend(student string)error{				//历史借阅的图书

fmt.Printf("\t\t<学号 %s 历史借阅的图书如下>\n",student)
var ISBN string
var name string
var lendtime string
var returntime string
rows,err := c.db.Query("SELECT 图书.ISBN,图书.书名,历史借阅表.借阅时间,历史借阅表.归还时间  FROM 历史借阅表,图书 WHERE 图书.ISBN=历史借阅表.ISBN AND 学号=? ",student)

if err != nil {
    fmt.Println("系统错误")
    return err
}else{
var count int=0
for rows.Next() {
	if err := rows.Scan(&ISBN,&name,&lendtime,&returntime); err == nil {
 	count=count+1
	fmt.Printf("ISBN:%s %s  借阅时间:%s 归还时间:%s\n",ISBN,name,lendtime,returntime)
	}
}
if count==0{
	fmt.Printf("您目前暂无历史借阅的图书")
}
}
return nil
}


func (c *library)borrowstudentbook(student string,ISBN string )error{                           //向管理员发起借阅申请
var name string
err := c.db.QueryRow("SELECT 图书.书名  FROM 图书 WHERE 图书.ISBN=? ",ISBN).Scan(&name)

if err != nil {
    fmt.Println("本馆暂无该图书！")
    return nil
}else{
 _,err:=c.db.Exec("INSERT INTO 等待审核学生(学号,ISBN,申请时间,借or还)VALUES (?,?,now(),0)",student,ISBN)
if err!=nil{
fmt.Printf("申请借阅失败")
return err}
fmt.Printf("已成功提交 %s %s 的图书借阅申请，请等待管理员审核！",ISBN,name)
}
return nil
}


func (c *library)returnstudentbook(student string,ISBN string)error{			//向管理员发起还书申请
	var name string
	err := c.db.QueryRow("SELECT ISBN FROM 正在借阅表 WHERE 学号=? AND ISBN=? ",student,ISBN).Scan(&ISBN)
	if err != nil {
    		fmt.Println("您没有借过这本书！")
    		return nil
	}else{
 		_,err:=c.db.Exec("INSERT INTO 等待审核学生(学号,ISBN,申请时间,借or还)VALUES (?,?,now(),1)",student,ISBN)
		if err!=nil{
		fmt.Printf("申请归还失败")
		return err}
		fmt.Printf("已成功提交 %s %s 的图书归还申请，请等待管理员审核！",ISBN,name)
	}
	return nil
}


func (c *library)borrowmoretime(student string,ISBN string)error{			//图书续借
	var moretime int
	err:=c.db.QueryRow("SELECT 已申请续借次数 FROM 正在借阅表 WHERE ISBN=? AND 学号=? ",ISBN,student).Scan(&moretime)
	if err != nil {
    		fmt.Println("没有找到你所要延长的图书")
		return nil}
	if moretime>3{
		fmt.Printf("您已经超过可申请续借的最大次数")
		return nil
	}
	_,errr:=c.db.Exec("UPdate 正在借阅表 set 已申请续借次数=已申请续借次数+1,最晚归还时间=date_add(最晚归还时间,interval 10 day) where ISBN=? AND 学号=?",ISBN,student)
	if errr!=nil{
	panic(errr)
	return errr
	}
	fmt.Printf("成功延迟归还时间!")
	return nil
	}

func (c *library)changepassword(student string,newpassword string)error{           //修改密码
	_,err:=c.db.Exec("UPdate 学生 set password=? where 学号=?",newpassword,student)
	if err!=nil{
	return err}
	fmt.Printf("\n成功修改密码！")
	return err
}



func students(){			//学生主函数
	var bookname string
	var L library
	L.connect()
	var user2 string
        var password2 string
	var name string
	var ISBN string
	var newpassword string
	for true{
        fmt.Printf("输入exit退出本系统\n")
        fmt.Println("学号：")
        fmt.Scanf("%s",&user2)
        if user2=="exit" {
        return 
        }
	fmt.Println("密码：")
        fmt.Scanf("%s",&password2)
	err:=L.db.QueryRow("select 学号 from 学生 where 学号=? AND password=?",user2,password2).Scan(&user2)
	if err!=nil{
	fmt.Println("学号或密码错误")
	}else{
	fmt.Println("登录成功！")
	break
	}}
	L.student=user2
	fmt.Println("\n\t\t***欢迎进入复旦大学图书馆管理系统***")
	L.printstufunction()
	for true{
	fmt.Println("\n**主菜单**\n输入help查询所有功能")
	var zzz string
	fmt.Scanf("%s",&zzz)
	switch zzz{
	case "exit":return
	case "help":L.printstufunction()
	case "01":L.myinformation(user2)
	case "02":{fmt.Printf("请输入新的名字")
		   fmt.Scanf("%s",&name)
		L.changeinformation(user2,name)}
	case "03":L.nowlend(user2)
	case "04":L.beforelend(user2)
	case "05":{fmt.Println("请输入书名：")
				fmt.Scanf("%s",&bookname)
				L.findbook(bookname)}
	case "06":{fmt.Println("请输入您要借阅图书的ISBN：")
		fmt.Scanf("%s",&ISBN)
		L.borrowstudentbook(user2,ISBN)}
	case "07":{fmt.Println("请输入您要归还图书的ISBN：")
		fmt.Scanf("%s",&ISBN)
		L.returnstudentbook(user2,ISBN)}
  	case "08":{fmt.Println("请输入您要延长借阅的图书ISBN：")
			fmt.Scanf("%s",&ISBN)
		L.borrowmoretime(user2,ISBN)}
	case "09":{fmt.Printf("请输入新密码:")
		fmt.Scanf("%s",&newpassword)
		L.changepassword(user2,newpassword)}
	default: fmt.Println("抱歉！本系统目前暂不支持此功能！")
	}
	}
}    





//以下为公用部分
type library struct{
	db *sql.DB
	student string
}
func (c *library)connect() {			//连接数据库
	db, err := sql.Open("mysql",fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?charset=utf8",user,password,database))
	if err != nil{
                fmt.Println("连接失败")
        }
	c.db=db
	
}
func main(){
	fmt.Println("\t\t***复旦大学图书馆***")
	fmt.Println("请问您的身份是？\n 学生 1 \n 图书管理员 2")
	var person int
	fmt.Scanf("%d",&person)
	if person==2{
	librarymanage()
	}else {
	students()
	}
  }
