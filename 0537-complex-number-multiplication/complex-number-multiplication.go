func complexNumberMultiply(num1 string, num2 string) string {

    // a+bi , c+di
    var a,b,c,d int
    // e+fi = (a+bi)* (c+di) = ac + cbi + adi - bd
    var e,f int

    snum := strings.Split(num1[:len(num1)-1],"+")
    a,err :=strconv.Atoi(snum[0])
    if err != nil{
        fmt.Println(err)
    }
    b,err=strconv.Atoi(snum[1])
    if err != nil{
        fmt.Println(err)
    }
    snum = strings.Split(num2[:len(num2)-1],"+")
    c,err=strconv.Atoi(snum[0])
    if err != nil{
        fmt.Println(err)
    }
    d,err=strconv.Atoi(snum[1])
    if err != nil{
        fmt.Println(err)
    }
    e = a*c -b*d
    f = c*b + a*d
    return fmt.Sprintf("%d+%di",e,f)

}