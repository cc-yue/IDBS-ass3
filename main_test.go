package main

import	(
	"fmt"
	"testing"
	_ "github.com/go-sql-driver/mysql"
)

func TestCreatetable(t *testing.T){
	var L library
	L.connect()
	err:=L.createtables()
	if err!=nil{
		t.Errorf(" Testcreatetable error",)
	}else{
		fmt.Printf("Testcreatetable  pass\n")
	}
}
	

func TestAddstudent(t *testing.T){
	var L library
	L.connect()
	var tests = []struct {
    	a, b string
    }{
        {"101","张一"},
        {"102","张二"},
        {"103","张三"},
        {"104","张四"},
        {"105","张五"},
    }
	var err error
	for num,te:= range tests{
	err=L.addstudent(te.a,te.b)
	if err!=nil{
		t.Errorf(" Testaddstudent case%d error",num)
	}else{
	fmt.Printf("Testaddstudent case%d pass\n",num)
	}}	
}

func TestFindstudent(t *testing.T){
	var L library
	L.connect()
	var tests = []struct {
    	a string
    }{
        {"101"},
        {"102"},
        {"103"},
    }
	var err error
	for num,te:= range tests{
	err=L.findstudent(te.a)
	if err!=nil{
		t.Errorf(" Testfindstudent case%d error",num)
	}else{
	fmt.Printf("\nTestfindstudent case%d pass",num)
	}}	
	fmt.Printf("\n")
	err=L.findstudent("110")
	if err==nil{t.Errorf("Testfindstudent case4 error")
	}else{
	fmt.Printf("Testfindstudent case4 pass\n")
	}
	err=L.findstudent("120")
	if err==nil{t.Errorf(" Testfindstudent case5 error")
	}else{
	fmt.Printf("Testfindstudent case5 pass\n")
	}
}

func TestAddbook(t *testing.T){
	var L library
	L.connect()
	var tests = []struct {
    	a,b string
    }{
        {"数学分析","1111"},
        {"数据库系统教程","2222"},
        {"线性代数","3333"},
		{"大学物理","4444"},
		{"离散数学","5555"},
    }
	var err error
	for num,te:= range tests{
	err=L.addbook(te.a,te.b)
	if err!=nil{
		t.Errorf(" Testaddbook case%d error",num)
	}else{
	fmt.Printf("\nTestaddbook case%d pass",num)
	}}	
}

func TestFindbook(t *testing.T){
	var L library
	L.connect()
	var tests = []struct {
    	a string
    }{
        {"数学分析"},
        {"数据库系统教程"},
        {"线性代数"},
    }
	var m bool
	for num,te:= range tests{
	m=L.findbook(te.a)
	if m!=true{
		t.Errorf(" Testfindbook case%d error",num)
	}else{
	fmt.Printf("\nTestfindbook case%d pass",num)
	}}	
	fmt.Printf("\n")
	m=L.findbook("抽象代数")
	if m==true{t.Errorf("Testfindbook case4 error")
	}else{
	fmt.Printf("Testfindbook case4 pass\n")
	}
	m=L.findbook("120")
	if m==true{t.Errorf(" Testfindbook case5 error")
	}else{
	fmt.Printf("Testfindbook case5 pass\n")
	}
	
}










func TestOvertimebook(t *testing.T){
	var L library
	L.connect()
	err:=L.overtimebook()
	if err!=nil{
		t.Errorf(" Testovertimebook error")
	}else{
	fmt.Printf("\nTestovertimebook pass")
	}
}	

func TestChangebook(t *testing.T){
	var L library
	L.connect()
	var tests = []struct {
    	a string;b string
    }{
        {"1111","数学分析B"},
        {"2222","数据库系统教程进阶版"},
        {"3333","高等代数"},
    }
	var m int
	for num,te:= range tests{
	m=L.changebook(te.a,te.b)
	if m==2{
	fmt.Printf("\nTestchangebook case%d pass",num)
	}else{
	t.Errorf(" Testchangebook case%d error",num)
	}}

	fmt.Printf("\n")
	m=L.changebook("7777","抽象代数")
	if m==1{fmt.Printf("Testfindbook case4 pass\n")
	}else{
	t.Errorf("Testfindbook case4 error")
	}
	m=L.changebook("8888","普通物理")
	if m==1{ fmt.Printf("Testfindbook case5 pass\n")
	}else{
	t.Errorf(" Testfindbook case5 error")
	}
}

func TestStopstudentid(t *testing.T){
	var L library
	L.connect()
	err:=L.stopstudentid()
	if err!=nil{t.Errorf("Teststopstudentid error\n")
	}else{
	fmt.Printf("Teststopstudentid pass\n")
	}
}	

func TestMyinformation(t *testing.T){
	var L library
	L.connect()
	var tests = []struct {
    	a string
    }{
        {"101"},
        {"102"},
        {"103"},
    }
	var m error
	for num,te:= range tests{
	m=L.myinformation(te.a)
	if m==nil{
	fmt.Printf("\nTestmyinformation case%d pass",num)
	}else{
	t.Errorf(" Testmyinformation case%d error",num)
	}}

	fmt.Printf("\n")
	m=L.myinformation("180")
	if m!=nil{fmt.Printf("Testmyinformation case4 pass\n")
	}else{
	t.Errorf("Testmyinformation case4 error")
	}
	m=L.myinformation("888")
	if m!=nil{ fmt.Printf("Testmyinformation case5 pass\n")
	}else{
	t.Errorf(" Testmyinformation case5 error")
	}
}


func TestChangeinformation(t *testing.T){
	var L library
	L.connect()
	var tests = []struct {
    	a string;b string
    }{
        {"101","张十三"},
        {"102","张十四"},
        {"103","张十五"},
    }
	var m error
	for num,te:= range tests{
	m=L.changeinformation(te.a,te.b)
	if m==nil{
	fmt.Printf("\nTestchangeinformation case%d pass",num)
	}else{
	t.Errorf(" Testchangeinformation case%d error",num)
	}}

	fmt.Printf("\n")
	m=L.changeinformation("180","张一百八")
	if m!=nil{fmt.Printf("Testchangeinformation case4 pass\n")
	}else{
	t.Errorf("Testchangeinformation case4 error")
	}
	m=L.changeinformation("888","张八百八")
	if m!=nil{ fmt.Printf("Testchangeinformation case5 pass\n")
	}else{
	t.Errorf(" Testchangeinformation case5 error")
	}
}
	
func TestNowlend(t *testing.T){
	var L library
	L.connect()
	var tests = []struct {
    	a string
    }{
        {"101"},
        {"102"},
        {"103"},
	{"104"},
	{"105"},
    }
	var m error
	for num,te:= range tests{
	m=L.nowlend(te.a)
	if m==nil{
	fmt.Printf("\nTestnowlend case%d pass",num)
	}else{
	t.Errorf(" Testnowlend case%d error",num)
	}}
}


func TestBeforelend(t *testing.T){
	var L library
	L.connect()
	var tests = []struct {
    	a string
    }{
        {"101"},
        {"102"},
        {"103"},
	{"104"},
	{"105"},
    }
	var m error
	for num,te:= range tests{
	m=L.beforelend(te.a)
	if m==nil{
	fmt.Printf("\nTestbeforelend case%d pass\n",num)
	}else{
	t.Errorf(" Testbeforelend case%d error\n",num)
	}}	
}


func TestStudentborrow(t *testing.T){
	var L library
	L.connect()
	var tests = []struct {
    	a string;b string
    }{
        {"101","8888"},
        {"102","2222"},
        {"103","9999"},
	{"104","4444"},
	{"105","5555"},
    }
	var m error
	for num,te:= range tests{
	m=L.borrowstudentbook(te.a,te.b)
	if m==nil{
	fmt.Printf("\nTeststudentborrow case%d pass\n",num)
	}else{
	t.Errorf(" Teststudentborrow case%d error\n",num)
	}}	
}

func TestBorrow(t *testing.T){
	var m error
	var L library
	L.connect()
	m=L.borrow()
	if m==nil{
	fmt.Printf("\nTestborrow  pass\n")
	}else{
	t.Errorf(" Testborrow  error\n")
	}
}	

func TestBorrowmoretime(t *testing.T){
	var L library
	L.connect()
	var tests = []struct {
    	a string;b string
    }{
        {"101","1111"},
        {"102","3333"},
        {"103","4444"},
	{"104","4444"},
	{"105","9999"},
    }
	var m error
	for num,te:= range tests{
	m=L.borrowmoretime(te.a,te.b)
	if m==nil{
	fmt.Printf("\nTestborrowmoretime case%d pass\n",num)
	}else{
	t.Errorf(" Testborrowmoretime case%d error\n",num)
	}}	
}

func TestChangepassword(t *testing.T){
	var L library
	L.connect()
	var tests = []struct {
    	a string;b string
    }{
        {"101","123"},
        {"102","123"},
        {"103","123"},
	{"104","123"},
	{"105","123"},
    }
	var m error
	for num,te:= range tests{
	m=L.changepassword(te.a,te.b)
	if m==nil{
	fmt.Printf("\nTestchangepassword case%d pass\n",num)
	}else{
	t.Errorf(" Testchangepassword case%d error\n",num)
	}}	
}

	
func TestReturnstudentbook(t *testing.T){
var L library
	L.connect()
	var tests = []struct {
    	a string;b string
    }{
        {"101","4444"},
        {"102","2222"},
        {"103","6666"},
	{"104","4444"},
	{"105","5555"},
    }
	var m error
	for num,te:= range tests{
	m=L.returnstudentbook(te.a,te.b)
	if m==nil{
	fmt.Printf("\nTeststudentreturn case%d pass\n",num)
	}else{
	t.Errorf(" Teststudentreturn case%d error\n",num)
	}}	
}


func TestReturn(t *testing.T){
	var m error
	var L library
	L.connect()
	m=L.returnbook()
	if m==nil{
	fmt.Printf("\nTestreturnbook  pass\n")
	}else{
	t.Errorf(" Testreturnbook  error\n")
	}
}



