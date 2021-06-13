package main

import (
	"fmt"
	"os"
)

func main() {
	// 1
	//fmt.Println("hello world")

	// 2
	//fmt.Println("go" + "lang")
	//fmt.Println("7.0/3.0=", 7.0/3.0)
	//
	//fmt.Println(true && false)
	//fmt.Println(true || false)
	//fmt.Println(!true)

	// 3
	//var a = "initial"
	//fmt.Println(a)
	//
	//var b, c int = 1, 2
	//fmt.Println(b, c)
	//
	//var d = true
	//fmt.Println(d)
	//
	//var e int
	//fmt.Println(e)
	//
	//f := "short"
	//fmt.Println(f)

	// 4
	//fmt.Println(s)
	//
	//const n = 500000000
	//
	//const d = 3e20 / n
	//fmt.Println(d)
	//
	//fmt.Println(int64(d))
	//
	//fmt.Println(math.Sin(n))

	// 5
	//i := 1
	//for i <= 3 {
	//	fmt.Println(i)
	//	i = i + 1
	//}
	//
	//for j := 7; j <=9; j++ {
	//	fmt.Println(j)
	//}
	//
	//for {
	//	fmt.Println("loop")
	//	break
	//}
	//
	//for n:=0; n <= 5; n++ {
	//	if n%2 == 0 {
	//		continue
	//	}
	//	fmt.Println(n)
	//}

	// 6
	//if num := 9; num < 0 {
	//	fmt.Println(num, "is negative")
	//} else if num < 10 {
	//	fmt.Println(num, "less than 10")
	//}

	// 7
	//i := 2
	//fmt.Println("write i as", i)
	//switch i {
	//case 1:
	//	fmt.Println("i is", 1)
	//case 2:
	//	fmt.Println("i is", 2)
	//case 3:
	//	fmt.Println("i is", 3)
	//}
	//
	//switch time.Now().Weekday() {
	//case time.Saturday, time.Sunday:
	//	fmt.Println("today is weekend")
	//default:
	//	fmt.Println("today id workday")
	//}
	//
	//t := time.Now()
	//switch {
	//case t.Hour() < 12:
	//	fmt.Println("It's before noon")
	//default:
	//	fmt.Println("It's after noon")
	//}
	//
	//whatAmI := func(i interface{}) {
	//	switch t := i.(type) {
	//	case bool:
	//		fmt.Println("i am bool")
	//	case int:
	//		fmt.Println("i am int")
	//	case string:
	//		fmt.Println("i am string")
	//	default:
	//		fmt.Println("Don't know type %T\n", t)
	//	}
	//}
	//
	//whatAmI(true)
	//whatAmI(10)
	//whatAmI("balala")

	// 8
	//var a [5]int
	//fmt.Println("emp:", a)
	//
	//a[4] = 100;
	//fmt.Println("a[4]", a[4])
	//
	//fmt.Println("a length", len(a))
	//
	//b := [5]int{1,2,3,4,5}
	//fmt.Println("dcl", b)
	//
	//var twoD [2][3]int
	//for i := 0; i < 2; i++ {
	//	for j :=0; j < 3; j++ {
	//		twoD[i][j] = i + j
	//	}
	//}
	//fmt.Println(twoD)

	// 9
	//s := make([]string, 3)
	//fmt.Println("emp:", s)
	//
	//s[0] = "a"
	//s[1] = "b"
	//s[2] = "c"
	//
	//fmt.Println("s[3]", s[2])
	//fmt.Println("length of s", len(s))
	//
	//s = append(s, "d")
	//s = append(s, "d", "f")
	//fmt.Println("append", s)
	//
	//c := make([]string, len(s))
	//copy(c, s)
	//fmt.Println("cpy", c)
	//
	//l := s[2:5]
	//fmt.Println("s[2:5]", l)
	//
	//l = s[:5]
	//fmt.Println("s[:5]", l)
	//
	//l = s[2:]
	//fmt.Println("s[2:]", l)
	//
	//t := []string{"g", "h", "l"}
	//fmt.Println("t", t)
	//
	//twoD := make([][]int, 3)
	//for i := 0; i < 3; i++ {
	//	innerline := i + 1
	//	twoD[i] = make([]int, innerline)
	//	for j := 0; j < innerline; j++ {
	//		twoD[i][j] = i + j
	//	}
	//}
	//fmt.Println("2d", twoD)

	// 10 map
	//m := make(map[string]int)
	//
	//m["k1"] = 7
	//m["k2"] = 13
	//fmt.Println("map: ", m)
	//
	//v1 := m["k1"]
	//fmt.Println("v1:", v1)
	//
	//fmt.Println("len:", len(m))
	//
	//delete(m, "k2")
	//fmt.Println("m", m)
	//
	//_, prs := m["k2"]
	//fmt.Println("prs:", prs)
	//
	//n := map[string]int{"foo": 1, "bar": 2}
	//fmt.Println("map:", n)

	// 11 range
	//nums := []int{2,3,4}
	//sum := 0
	//for _, num := range nums {
	//	sum += num
	//}
	//fmt.Println("sum:", sum)
	//
	//for i, num := range nums {
	//	if num == 3 {
	//		fmt.Println("index:", i)
	//	}
	//}
	//
	//kvs := map[string]string{"a":"apple", "b":"banana"}
	//for k, v := range kvs {
	//	fmt.Printf("%s -> %s\n", k, v)
	//}
	//
	//for k := range kvs {
	//	fmt.Println("key:", k)
	//}
	//
	//for i, c := range "go" {
	//	fmt.Println(i, c)
	//}

	//res := plus(1, 2)
	//fmt.Println("1+2=", res)
	//
	//res = plusPlus(1,2,3)
	//fmt.Println("1+2+3=", res)

	//a, b := vals()
	//fmt.Println(a)
	//fmt.Println(b)
	//
	//_, c := vals()
	//fmt.Println(c)

	//sum(1,2,3)
	//
	//nums := []int{1,2,3}
	//sum(nums...)

	//nextInt := intSeq()
	//
	//fmt.Println(nextInt())
	//fmt.Println(nextInt())
	//fmt.Println(nextInt())
	//
	//nextInt2 := intSeq()
	//fmt.Println(nextInt2())

	//println(fact(7))

	//i := 1
	//fmt.Println(i)
	//
	//zeroVal(1)
	//fmt.Println(i)
	//
	//ptrVal(&i)
	//fmt.Println(i)
	//
	//fmt.Println(&i)

	//fmt.Println(person{"Bob", 20})
	//
	//fmt.Println(person{name: "Alice", age: 23})
	//
	//fmt.Println(person{name: "Alice"})
	//
	//fmt.Println(&person{name: "young", age: 20})
	//
	//s := person{name: "young", age: 23}
	//fmt.Println(s.name)
	//
	//sp := &s
	//fmt.Println(sp.age)

	//r := rect{width: 5, height: 10}
	//fmt.Printf("%p\n",&r)
	//
	//fmt.Println(r.area())
	//fmt.Println(r.perim())
	//
	//rp := &r
	//fmt.Printf("%p\n", rp)
	//fmt.Println(rp.area())
	//fmt.Println(rp.perim())

	//c := circle{radius: 5}
	//r := rect{width: 3, height: 4}
	//
	//measure(c)
	//measure(r)

	//for _, i := range []int{7, 42} {
	//	if r, e := f1(i); e != nil {
	//		fmt.Println("f1 failed", e)
	//	} else {
	//		fmt.Println("f1 worked", r)
	//	}
	//}
	//
	//for _, i := range []int{7, 42} {
	//	if r, e := f2(i); e != nil {
	//		fmt.Println("f2 failed", e)
	//	} else {
	//		fmt.Println("f2 worked", r)
	//	}
	//}

	//f("direct")
	//
	//go f("go")
	//
	//go func(arg string) {
	//	fmt.Println(arg, ":")
	//}("going")
	//
	//time.Sleep(time.Second*3)
	//fmt.Println("done")

	//messages := make(chan string)
	//
	//go func() {messages <- "ping"}()
	//
	//msg := <-messages
	//fmt.Println(msg)

	//messages := make(chan string, 2)
	//
	//messages <- "buffered"
	//messages <- "channel"
	//
	//fmt.Println(<-messages)
	//fmt.Println(<-messages)

	//done := make(chan bool, 1)
	//go worker(done)
	//
	//<- done

	//pings := make(chan string, 1)
	//pongs := make(chan string, 1)
	//ping(pings, "passed msg")
	//pong(pings, pongs)
	//fmt.Println(<-pongs)

	//c1 := make(chan string)
	//c2 := make(chan string)
	//
	//go func() {
	//	time.Sleep(1 * time.Second)
	//	c1 <- "c1"
	//}()
	//
	//go func() {
	//	time.Sleep(2 * time.Second)
	//	c2 <- "c2"
	//}()
	//
	//for i := 0; i < 2; i++ {
	//	select {
	//	case msg1 := <-c1:
	//		fmt.Println("received", msg1)
	//	case msg2 := <-c2:
	//		fmt.Println("received", msg2)
	//	}
	//}

	//c1 := make(chan string)
	//c2 := make(chan string)
	//
	//go func() {
	//	time.Sleep(2 * time.Second)
	//	c1 <- "result 1"
	//}()
	//
	//select {
	//case msg := <-c1:
	//	fmt.Println(msg)
	//case <-time.After(1 * time.Second):
	//	fmt.Println("timeout 1")
	//}
	//
	//go func() {
	//	time.Sleep(2 * time.Second)
	//	c2 <- "result 2"
	//}()
	//
	//select {
	//case msg := <- c2:
	//	fmt.Println(msg)
	//case <-time.After(3 * time.Second):
	//	fmt.Println("timeout 2")
	//}

	//messages := make(chan string)
	//signals := make(chan string)
	//
	//select {
	//case msg := <- messages:
	//	fmt.Println("received messages", msg)
	//default:
	//	fmt.Println("no msg received")
	//}
	//
	//msg := "hi"
	//select {
	//case messages <- msg:
	//	fmt.Println("sent msg", msg)
	//default:
	//	fmt.Println("no msg sent")
	//}
	//
	//select {
	//case msg := <-messages:
	//	fmt.Println("received msg", msg)
	//case sig := <-signals:
	//	fmt.Println("received sig", sig)
	//default:
	//	fmt.Println("no active")
	//}

	//jobs := make(chan int, 5)
	//done := make(chan bool)
	//
	//go func() {
	//	for {
	//		j, more := <-jobs
	//		if more {
	//			fmt.Println("received job", j)
	//		} else {
	//			fmt.Println("received all job")
	//			done <- true
	//			return
	//		}
	//	}
	//}()
	//
	//for j := 1; j <= 3; j++ {
	//	jobs <- j
	//	fmt.Println("sent job", j)
	//}
	//
	//close(jobs)
	//fmt.Println("sent all jobs")
	//
	//<-done

	//queue := make(chan string, 2)
	//queue <- "one"
	//queue <- "two"
	//close(queue)
	//
	//for elem := range queue {
	//	fmt.Println(elem)
	//}

	//timer1 := time.NewTimer(2 * time.Second)
	//
	//<-timer1.C
	//fmt.Println("timer1 fired")
	//
	//timer2 := time.NewTimer(2 * time.Second)
	//go func() {
	//	<-timer2.C
	//	fmt.Println("timer2 fired")
	//}()
	//stop := timer2.Stop()
	//if stop {
	//	fmt.Println("timer2 stop")
	//}
	//
	//time.Sleep(2 * time.Second)

	//ticker := time.NewTicker(500 * time.Millisecond)
	//done := make(chan bool)
	//
	//go func() {
	//	for {
	//		select {
	//		case <-done:
	//			return
	//		case t := <-ticker.C:
	//			fmt.Println("ticker run", t)
	//		}
	//	}
	//}()
	//
	//time.Sleep(1600 * time.Millisecond)
	//ticker.Stop()
	//done <- true
	//fmt.Println("Ticker stopped")

	//const numjobs = 5
	//
	//jobs := make(chan int, numjobs)
	//results := make(chan int, numjobs)
	//
	//for w := 0; w < 3; w++ {
	//	go worker(w, jobs, results)
	//}
	//
	//for j := 0; j < numjobs; j++ {
	//	jobs <- j
	//}
	//close(jobs)
	//
	//for a := 1; a <= numjobs; a++ {
	//	<-results
	//}

	//var wg sync.WaitGroup
	//
	//for i := 1; i <= 5; i++ {
	//	wg.Add(1)
	//	go worker(i, &wg)
	//}
	//
	//wg.Wait()

	//requests := make(chan int, 5)
	//for i := 1; i <= 5; i++ {
	//	requests <- i
	//}
	//close(requests)
	//
	//limiter := time.Tick(200 * time.Millisecond)
	//
	//for request := range requests {
	//	<-limiter
	//	fmt.Println("request", request, time.Now())
	//}
	//
	//brustylimiter := make(chan time.Time, 3)
	//for i := 1; i < 3; i++ {
	//	brustylimiter <- time.Now()
	//}
	//
	//go func() {
	//	for t := range time.Tick(200 * time.Millisecond) {
	//		brustylimiter <- t
	//	}
	//}()
	//
	//brustyRequest := make(chan int, 5)
	//for i := 1; i <= 5; i++ {
	//	brustyRequest <- i
	//}
	//close(brustyRequest)
	//for br := range brustyRequest {
	//	<-brustylimiter
	//	fmt.Println("brequest", br, time.Now())
	//}

	//var ops uint64
	//
	//var wg sync.WaitGroup
	//
	//for i := 1; i <= 50; i++ {
	//	wg.Add(1)
	//
	//	go func() {
	//		defer wg.Done()
	//		for i := 1; i <= 1000; i++ {
	//			atomic.AddUint64(&ops, 1)
	//		}
	//	}()
	//}
	//
	//wg.Wait()
	//
	//fmt.Println("ops:", ops)

	//var state = make(map[int]int)
	//
	//var mutex = &sync.Mutex{}
	//
	//var readOps uint64
	//var writeOps uint64
	//
	//for r := 1; r <= 100; r++ {
	//	go func() {
	//		total := 0
	//		for {
	//			key := rand.Intn(5)
	//			mutex.Lock()
	//			total += state[key]
	//			mutex.Unlock()
	//			atomic.AddUint64(&readOps, 1)
	//
	//			time.Sleep(time.Millisecond)
	//		}
	//	}()
	//}
	//
	//for w := 1; w <= 10; w++ {
	//	go func() {
	//		key := rand.Intn(5)
	//		val := rand.Intn(100)
	//		mutex.Lock()
	//		state[key] = val
	//		mutex.Unlock()
	//		atomic.AddUint64(&writeOps, 1)
	//
	//		time.Sleep(time.Millisecond)
	//	}()
	//}
	//
	//time.Sleep(time.Second)
	//
	//readOpsFinal := atomic.LoadUint64(&readOps)
	//fmt.Println("readOpsFinal", readOpsFinal)
	//
	//writeOpsFinal := atomic.LoadUint64(&writeOps)
	//fmt.Println("writeOpsFinal", writeOpsFinal)
	//
	//mutex.Lock()
	//fmt.Println("state", state)
	//mutex.Unlock()

	//var readOps uint64
	//var writeOps uint64
	//
	//reads := make(chan readOp)
	//writes := make(chan writeOp)
	//
	//go func() {
	//	state := make(map[int]int)
	//	for {
	//		select {
	//		case read := <-reads:
	//			read.resp <-  state[read.key]
	//		case write := <- writes:
	//			state[write.key] = write.val
	//			write.resp <- true
	//		}
	//	}
	//}()
	//
	//for r := 0; r < 100; r++ {
	//	go func() {
	//		for {
	//			read := readOp{
	//				key: rand.Intn(5),
	//				resp: make(chan int),
	//			}
	//			reads <- read
	//			<-read.resp
	//			atomic.AddUint64(&readOps, 1)
	//			time.Sleep(time.Millisecond)
	//		}
	//	}()
	//}
	//
	//for w := 0; w < 10; w++ {
	//	go func() {
	//		for {
	//			write := writeOp{
	//				key: rand.Intn(5),
	//				val: rand.Intn(100),
	//				resp: make(chan bool),
	//			}
	//			writes <- write
	//			<-write.resp
	//			atomic.AddUint64(&writeOps, 1)
	//			time.Sleep(time.Millisecond)
	//		}
	//	}()
	//}
	//
	//time.Sleep(time.Second)
	//readOpsFinal := atomic.LoadUint64(&readOps)
	//writeOpsFinal := atomic.LoadUint64(&writeOps)
	//fmt.Println("readOps", readOpsFinal)
	//fmt.Println("writeOps", writeOpsFinal)

	//s := []string{"c", "b", "a"}
	//sort.Strings(s)
	//fmt.Println("sorted str", s)
	//
	//n := []int{3, 2, 1}
	//sort.Ints(n)
	//fmt.Println("sorted int", n)
	//
	//fmt.Println(sort.IntsAreSorted(n))

	//fruits := []string{"peach", "banana", "kiwi"}
	////就是int64()
	//sort.Sort(byLength(fruits))
	//fmt.Println(fruits)

	//panic("a problem")
	//_, err := os.Create("/tmp/file")
	//if err != nil {
	//	panic(err)
	//}

	//f := createFile("/tmp/defer.txt")
	//defer closeFile(f)
	//writeFile(f)

	//var strs = []string{"peach", "apple", "pear", "plum"}
	//
	//fmt.Println(Index(strs, "pear"))
	//
	//fmt.Println(Include(strs, "grape"))
	//
	//fmt.Println(Any(strs, func(s string) bool {
	//	return strings.HasPrefix(s, "p")
	//}))
	//
	//fmt.Println(All(strs, func(s string) bool {
	//	return strings.HasPrefix(s, "p")
	//}))
	//
	//fmt.Println(Filter(strs, func(s string) bool {
	//	return strings.Contains(s, "e")
	//}))
	//
	//fmt.Println(Map(strs, func(s string) string {
	//	return strings.ToUpper(s)
	//}))

	//p("Contains:  ", s.Contains("test", "es"))
	//p("Count:     ", s.Count("test", "t"))
	//p("HasPrefix: ", s.HasPrefix("test", "te"))
	//p("HasSuffix: ", s.HasSuffix("test", "st"))
	//p("Index:     ", s.Index("test", "e"))
	//p("Join:      ", s.Join([]string{"a", "b"}, "-"))
	//p("Repeat:    ", s.Repeat("a", 5))
	//p("Replace:   ", s.Replace("foo", "o", "0", -1))
	//p("Replace:   ", s.Replace("foo", "o", "0", 1))
	//p("Split:     ", s.Split("a-b-c-d-e", "-"))
	//p("ToLower:   ", s.ToLower("TEST"))
	//p("ToUpper:   ", s.ToUpper("test"))
	//p()
	//
	//p("Len: ", len("hello"))
	//p("Char:", "hello"[1])

	// 基础类型转json
	//bolB, _ := json.Marshal(true)
	//fmt.Println(string(bolB))
	//
	//intB, _ := json.Marshal(1)
	//fmt.Println(string(intB))
	//
	//fB, _ := json.Marshal(2.34)
	//fmt.Println(string(fB))
	//
	//strB, _ := json.Marshal("gopher")
	//fmt.Println(string(strB))
	//
	//slc := []string{"apple", "peach", "peer"}
	//slcB, _ := json.Marshal(slc)
	//fmt.Println(string(slcB))
	//
	//m := map[string]int{"key1": 1, "key2": 2}
	//mB, _ := json.Marshal(m)
	//fmt.Println(string(mB))
	//
	//res1D := &response1{
	//	Page: 1,
	//	Fruits: []string{"peach", "peer"},
	//}
	//resp1, _ := json.Marshal(res1D)
	//fmt.Println(string(resp1))
	//
	//resp2 := &response2{
	//	Page: 2,
	//	Fruits: []string{"banana", "peach"},
	//}
	//resp2B, _ := json.Marshal(resp2)
	//fmt.Println(string(resp2B))
	//
	//byt := []byte(`{"num":6.13,"strs":["a","b"]}`)
	//
	//var dat map[string]interface{}
	//
	//if err := json.Unmarshal(byt, &dat); err != nil {
	//	panic(err)
	//}
	//fmt.Println(dat)
	//
	//num := dat["num"].(float64)
	//fmt.Println(num)
	//
	//strss := dat["strs"].([]interface{})
	//s := strss[0].(string)
	//fmt.Println(s)
	//
	//str := `{"page":1,"fruits":["apple", "peach"]}`
	//res := response2{}
	//if err := json.Unmarshal([]byte(str), &res); err != nil {
	//	panic(err)
	//}
	//fmt.Println(res.Fruits[0])
	//
	//enc := json.NewEncoder(os.Stdout)
	//d := map[string]int{"apple":5, "lettuce":7}
	//enc.Encode(d)

	//coffee := &Plant{Id: 1, Name: "Coffee"}
	//coffee.Origin = []string{"Ethiopia", "Brazil"}
	//
	//out, _ := xml.MarshalIndent(&coffee, " ", "  ")
	//fmt.Println(string(out))
	//
	//fmt.Println(xml.Header + string(out))
	//
	//var p Plant
	//if err := xml.Unmarshal(out, &p); err != nil {
	//	panic(err)
	//}
	//fmt.Println(p)
	//
	//tomato := &Plant{Id: 81, Name: "Tomato"}
	//tomato.Origin = []string{"Mexico", "California"}
	//
	//type Nesting struct {
	//	XMLName xml.Name `xml:"nesting"`
	//	Plants []*Plant `xml:"parent>child>plant"`
	//}
	//
	//n := &Nesting{}
	//n.Plants = []*Plant{coffee, tomato}
	//
	//out, _ = xml.MarshalIndent(n, " ", " ")
	//fmt.Println(string(out))

	//p := fmt.Println
	//
	//now := time.Now()
	//p(now)
	//
	//then := time.Date(
	//	2021, 11, 11, 11, 11, 11, 651387237, time.Local)
	//p(then)
	//
	//p(then.Year())
	//p(then.Month())
	//p(then.Day())
	//p(then.Hour())
	//p(then.Minute())
	//p(then.Second())
	//p(then.Nanosecond())
	//p(then.Local())
	//p(then.Location())
	//
	//p(then.Weekday())
	//
	//fmt.Println("--->>>")
	//
	//p(then.Before(now))
	//p(then.Equal(now))
	//p(then.After(now))
	//
	//diff := now.Sub(then)
	//p(diff)
	//
	//p(diff.Hours())
	//p(diff.Minutes())
	//p(diff.Seconds())
	//p(diff.Nanoseconds())
	//
	//p(then.Add(diff))
	//p(then.Add(-diff))

	//now := time.Now()
	//secs := now.Unix()
	//nano := now.UnixNano()
	//fmt.Println(now)
	//
	//mils := nano / 1000000
	//fmt.Println(secs)
	//fmt.Println(mils)
	//fmt.Println(nano)
	//
	//fmt.Println(time.Unix(secs, 0))
	//fmt.Println(time.Unix(0, nano))

	//p := fmt.Println
	//t := time.Now()
	//p(t.Format(time.RFC3339))
	//
	//t1, _ := time.Parse(time.RFC3339, "2012-11-01T22:08:41+00:00")
	//fmt.Println(t1)

	//"2006-01-02 15:04:05"
	//p(t.Format("15:04PM"))
	//p(t.Format("3:04PM"))
	//
	//fmt.Println("--->>>")
	//
	//p(t.Format("Mon Jan 2 15:04:05 2006"))
	//p(t.Format("Mon Jan _2 15:04:05 2006"))
	//
	//p(t.Format("2006-01-02T15:04:05.999999-07:00"))
	//
	//form := "3 04 PM"
	//t2, _ := time.Parse(form, "8 41 PM")
	//p(t2)
	//
	//fmt.Printf("%d-%02d-%02dT%02d:%02d:%02d-00:00\n",
	//	t.Year(), t.Month(), t.Day(),
	//	t.Hour(), t.Minute(), t.Second())
	//
	//fmt.Println("--->>>")
	//
	//ansic := "Mon Jan _2 15:04:05 2006"
	//_, e := time.Parse(ansic, "8:41PM")
	//p(e)

	//fmt.Print(rand.Intn(100), ",")
	//fmt.Println(rand.Intn(100))

	//fmt.Println(rand.Float64())
	//
	//fmt.Print((rand.Float64()*5) + 5, ",")
	//fmt.Print((rand.Float64()*5+5), "\n")
	//
	//s1 := rand.NewSource(time.Now().UnixNano())
	//r1 := rand.New(s1)
	//
	//fmt.Print(r1.Intn(100), ", ")
	//fmt.Println(r1.Intn(100))
	//fmt.Println()
	//
	//s2 := rand.NewSource(42)
	//r2 := rand.New(s2)
	//fmt.Print(r2.Intn(100), ", ")
	//fmt.Println(r2.Intn(100))
	//fmt.Println()
	//
	//s3 := rand.NewSource(42)
	//r3 := rand.New(s3)
	//fmt.Print(r3.Intn(100), ",")
	//fmt.Print(r3.Intn(100))
	//fmt.Println()

	//f, _ := strconv.ParseFloat("1.23", 64)
	//fmt.Println(f)
	//
	//i, _ := strconv.ParseInt("123", 0, 64)
	//fmt.Println(i)
	//
	//d, _ := strconv.ParseInt("0x1c8", 0, 64)
	//fmt.Println(d)
	//
	//ui, _ := strconv.ParseUint("789", 0, 64)
	//fmt.Println(ui)
	//
	//k, _ := strconv.Atoi("135")
	//fmt.Println(k)
	//
	//_, e := strconv.Atoi("wat")
	//if e != nil {
	//	fmt.Println(e)
	//}

	//s := "postgres://user:pass@host.com:5432/path?k=v#f"
	//
	//u, err := url.Parse(s)
	//if err != nil {
	//	panic(err)
	//}
	//
	//fmt.Println(u.Scheme)
	//
	//fmt.Println(u.User)
	//fmt.Println(u.User.Username())
	//
	//p, _ := u.User.Password()
	//fmt.Println(p)
	//
	//fmt.Println(u.Host)
	//host, port, _ := net.SplitHostPort(u.Host)
	//fmt.Println(host)
	//fmt.Println(port)
	//
	//fmt.Println("--->>>")
	//
	//fmt.Println(u.Path)
	//fmt.Println(u.Fragment)
	//
	//fmt.Println(u.RawQuery)
	//m, _ := url.ParseQuery(u.RawQuery)
	//fmt.Println(m)
	//fmt.Println(m["k"][0])

	//s := "sha1 this string"
	//
	//h := sha1.New()
	//
	//h.Write([]byte(s))
	//
	//bs := h.Sum(nil)
	//
	//fmt.Println(s)
	//fmt.Printf("%x\n", bs)

	//s := "abc123!?$*&()'-=@~"
	//
	//encs := b64.StdEncoding.EncodeToString([]byte(s))
	//fmt.Println(encs)
	//
	//ds, _ := b64.StdEncoding.DecodeString(encs)
	//fmt.Println(string(ds))
	//
	//urlEncs := b64.URLEncoding.EncodeToString([]byte(s))
	//fmt.Println(urlEncs)
	//
	//urlds, _ := b64.URLEncoding.DecodeString(urlEncs)
	//fmt.Println(string(urlds))

	//file, err := ioutil.ReadFile("/tmp/dat")
	//checkErr(err)
	//fmt.Println(string(file))
	//
	//f, err := os.Open("/tmp/dat")
	//checkErr(err)
	//
	//b1 := make([]byte, 5)
	//n1, err := f.Read(b1)
	//checkErr(err)
	//fmt.Printf("%d bytes: %s\n", n1, string(b1[:n1]))
	//
	//o2, err := f.Seek(6, 0)
	//checkErr(err)
	//b2 := make([]byte, 2)
	//n2, err := f.Read(b2)
	//checkErr(err)
	//fmt.Printf("%d bytes @ %d\n", n2, o2)
	//fmt.Printf("%v\n", string(b2[:n2]))
	//
	//o3, err := f.Seek(6, 0)
	//checkErr(err)
	//b3 := make([]byte, 2)
	//n3, err := io.ReadAtLeast(f, b3, 2)
	//checkErr(err)
	//fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))
	//
	//_, err = f.Seek(0,0)
	//checkErr(err)
	//
	//r4 := bufio.NewReader(f)
	//b4, err := r4.Peek(5)
	//checkErr(err)
	//fmt.Printf("5 bytes: %s\n", string(b4))
	//
	//f.Close()

	//d1 := []byte("hello\ngo\n")
	//err := ioutil.WriteFile("/tmp/dat1", d1, 0644)
	//checkErr(err)
	//
	//f, err := os.Create("/tmp/dat2")
	//checkErr(err)
	//
	//defer f.Close()
	//
	//d2 := []byte{115, 111, 109, 101, 10}
	//n2, err := f.Write(d2)
	//checkErr(err)
	//fmt.Printf("wrote %d bytes\n", n2)
	//
	//n3, err := f.WriteString("writes\n")
	//checkErr(err)
	//fmt.Printf("wrote %d bytes\n", n3)
	//
	//f.Sync()
	//
	//w := bufio.NewWriter(f)
	//n4, err := w.WriteString("buffered\n")
	//checkErr(err)
	//fmt.Printf("wrote %d bytes\n", n4)
	//w.Flush()

	//scanner := bufio.NewScanner(os.Stdin)
	//
	//for scanner.Scan() {
	//
	//	ucl := strings.ToUpper(scanner.Text())
	//	fmt.Println(ucl)
	//}
	//
	//if err := scanner.Err(); err != nil {
	//	fmt.Println(os.Stderr, "error:", err)
	//	os.Exit(1)
	//}

	//p := filepath.Join("dir1", "dir2", "filename")
	//fmt.Println("p:", p)
	//
	//fmt.Println(filepath.Join("dir1//", "filename"))
	//fmt.Println(filepath.Join("dir1/../dir1", "filename"))
	//
	//fmt.Println("Dir(p):", filepath.Dir(p))
	//fmt.Println("Base(p):", filepath.Base(p))
	//
	//fmt.Println(filepath.IsAbs("dir/file"))
	//fmt.Println(filepath.IsAbs("/dir/file"))
	//
	//filename := "config.json"
	//
	//ext := filepath.Ext(filename)
	//fmt.Println(ext)
	//
	//fmt.Println(strings.TrimSuffix(filename, ext))
	//
	//rel, err := filepath.Rel("a/b", "a/b/t/file")
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(rel)
	//
	//rel, err = filepath.Rel("a/b", "a/c/t/file")
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(rel)

	//err := os.Mkdir("subdir", 0755)
	//checkErr(err)
	//
	//defer os.RemoveAll("subdir")
	//
	//createEmptyFile := func(name string) {
	//	d := []byte("")
	//	checkErr(ioutil.WriteFile(name, d, 0644))
	//}
	//
	//createEmptyFile("subdir/file1")
	//fmt.Println("--->>>")
	//
	//err = os.MkdirAll("subdir/parent/child", 0755)
	//checkErr(err)
	//
	//createEmptyFile("subdir/parent/file2")
	//createEmptyFile("subdir/parent/file3")
	//createEmptyFile("subdir/parent/child/file4")
	//fmt.Println("--->>>")
	//
	//dir, err := ioutil.ReadDir("subdir/parent")
	//checkErr(err)
	//
	//for _, entry := range dir {
	//	fmt.Println("", entry.Name(), entry.IsDir())
	//}
	//
	//err = os.Chdir("../../..")
	//checkErr(err)
	//
	//fmt.Println("--->>>")
	//
	//checkErr(filepath.Walk("subdir", visit))

	//f, err := ioutil.TempFile("", "sample")
	//checkErr(err)
	//
	//fmt.Println("Temp file name:", f.Name())
	//
	//defer os.Remove(f.Name())
	//
	//_, err = f.Write([]byte{1, 2, 3, 4})
	//checkErr(err)
	//
	//dname, err := ioutil.TempDir("", "sampledir")
	//checkErr(err)
	//fmt.Println("Temp dir name:", dname)
	//
	//defer os.Remove(dname)
	//
	//fname := filepath.Join(dname, "file1")
	//err = ioutil.WriteFile(fname, []byte("hello world"), 0666)
	//checkErr(err)
	//
	//defer os.RemoveAll(dname)

	//args := os.Args
	//strings := os.Args[1:]
	//
	//fmt.Println("os.Args: ", args)
	//fmt.Println("os.Args[1:]: ", strings)
	//
	//fmt.Println("os.Args[3]", os.Args[3])

	//wordPtr := flag.String("word", "foo", "a string")
	//
	//numPtr := flag.Int("numb", 42, "an int")
	//boolPtr := flag.Bool("fork", false, "a bool")
	//
	//var svar string
	//flag.StringVar(&svar, "svar", "bar", "a string var")
	//
	//flag.Parse()
	//
	//fmt.Println("word:", *wordPtr)
	//fmt.Println("numb:", *numPtr)
	//fmt.Println("fork", *boolPtr)
	//fmt.Println("svar", svar)
	//fmt.Println("tail:", flag.Args())

	//fooCmd := flag.NewFlagSet("foo", flag.ExitOnError)
	//fooEnable := fooCmd.Bool("enable", false, "enable")
	//fooName := fooCmd.String("name", "", "name")
	//
	//barCmd := flag.NewFlagSet("bar", flag.ExitOnError)
	//barLevel := barCmd.Int("level", 0, "level")
	//
	//if len(os.Args) < 2 {
	//	fmt.Println("expected 'foo', or 'bar' subommands")
	//	os.Exit(1)
	//}
	//
	//switch os.Args[1] {
	//case "foo":
	//	fooCmd.Parse(os.Args[2:])
	//	fmt.Println(" subcommand 'foo'")
	//	fmt.Println(" enable:", *fooEnable)
	//	fmt.Println(" fooName:", *fooName)
	//	fmt.Println(" tail:", fooCmd.Args())
	//case "bar":
	//	barCmd.Parse(os.Args[2:])
	//	fmt.Println(" subcommand 'bar'")
	//	fmt.Println(" subcommand 'level'", *barLevel)
	//	fmt.Println(" tail: ", barCmd.Args())
	//default:
	//	fmt.Println("expected 'foo' or 'bar' subcommands")
	//	os.Exit(1)
	//}

	//os.Setenv("FOO", "1")
	//fmt.Println("FOO", os.Getenv("FOO"))
	//fmt.Println("BAR", os.Getenv("BAR"))
	//
	//args := os.Environ()
	//
	//fmt.Println(args)
	//fmt.Println("--->>>")
	//for _, e := range os.Environ() {
	//	pair := strings.SplitN(e, "=", 2)
	//	fmt.Println(pair[0])
	//}

	//resp, err := http.Get("http://gobyexample.com")
	//if err != nil {
	//	panic(err)
	//}
	//defer resp.Body.Close()
	//
	//fmt.Println("Response status:", resp.Status)
	//
	//scanner := bufio.NewScanner(resp.Body)
	//for i := 0; scanner.Scan() && i < 5; i++ {
	//	fmt.Println(scanner.Text())
	//}
	//
	//if err := scanner.Err(); err != nil {
	//	panic(err)
	//}

	//http.HandleFunc("/hello", hello)
	//http.HandleFunc("/headers", headers)
	//
	//http.ListenAndServe(":8080", nil)

	//http.HandleFunc("/hello", hello)
	//http.ListenAndServe(":8080", nil)

	//dateCmd := exec.Command("date")
	//
	//dateOut, err := dateCmd.Output()
	//if err != nil {
	//	panic(err)
	//}
	//
	//fmt.Println("> date")
	//fmt.Println(string(dateOut))
	//
	//grepCmd := exec.Command("grep", "hello")
	//
	//grepIn, _ := grepCmd.StdinPipe()
	//grepOut, _ := grepCmd.StdoutPipe()
	//
	//grepCmd.Start()
	//grepIn.Write([]byte("hello grep\ngoodbye grep"))
	//grepIn.Close()
	//grepBytes, _ := ioutil.ReadAll(grepOut)
	//grepCmd.Wait()
	//
	//fmt.Println("> grep hello")
	//fmt.Println(string(grepBytes))
	//
	//lsCmd := exec.Command("bash", "-c", "ls -a -l -h")
	//lsOut, err := lsCmd.Output()
	//if err != nil {
	//	panic(err)
	//}
	//
	//fmt.Println("> ls -a -l -h")
	//fmt.Println(string(lsOut))

	//binary, lookErr := exec.LookPath("ls")
	//if lookErr != nil {
	//	panic(lookErr)
	//}
	//
	//args := []string{"ls", "-a", "-l", "-h"}
	//
	//env := os.Environ()
	//
	//execErr := syscall.Exec(binary, args, env)
	//if execErr != nil {
	//	panic(execErr)
	//}

	//sigs := make(chan os.Signal, 1)
	//done := make(chan bool)
	//
	//signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	//
	//go func() {
	//	sig := <-sigs
	//	fmt.Println()
	//	fmt.Println(sig)
	//	done <- true
	//}()
	//
	//fmt.Println("awaiting signal")
	//<-done
	//fmt.Println("exiting")

	defer fmt.Println("!")

	os.Exit(3)
}

//func hello(w http.ResponseWriter, r *http.Request) {
//	ctx := r.Context()
//	fmt.Println("server hello start")
//	defer fmt.Println("server hello end")
//
//	select {
//	case <-time.After(10*time.Second):
//		fmt.Fprintf(w, "hello\n")
//	case <-ctx.Done():
//
//		err := ctx.Err()
//		fmt.Println("server: ", err)
//		fmt.Println("--->>>")
//		internalError := http.StatusInternalServerError
//		http.Error(w, err.Error(), internalError)
//	}
//}

//func hello(w http.ResponseWriter, r *http.Request) {
//	fmt.Fprintf(w, "hello \n")
//}
//
//func headers(w http.ResponseWriter, r *http.Request) {
//	for name, headers := range r.Header {
//		for _, header := range headers {
//			fmt.Fprintf(w, "%v: %v\n", name, header)
//		}
//	}
//}

//func IntMin(a, b int) int {
//	if a < b {
//		return a
//	}
//	return b
//}

//func checkErr(err error) {
//	if err != nil {
//		panic(err)
//	}
//}

//func visit(p string, info os.FileInfo, err error) error {
//	if err != nil {
//		return err
//	}
//	fmt.Println("", p, info.Name(), info.IsDir())
//	return nil
//}
//
//func checkErr(e error) {
//	if e != nil {
//		panic(e)
//	}
//}
//func checkErr(e error) {
//	if e != nil {
//		panic(e)
//	}
//}

//type Plant struct {
//	XMLName xml.Name `xml:"plant"`
//	Id int `xml:"id,attr"`
//	Name string `xml:"name"`
//	Origin []string `xml:"origin"`
//}
//
//func (p Plant) String() string {
//	return fmt.Sprintf("id: %v, name: %v, origin: %v", p.Id, p.Name, p.Origin)
//}

//type response1 struct {
//	Page int
//	Fruits []string
//}
//
//type response2 struct {
//	Page int `json:"pages"`
//	Fruits []string `json:"fruits"`
//}

//var p = fmt.Println

//func Index(vs []string, t string) int {
//	for i, v := range vs {
//		if v == t {
//			return i
//		}
//	}
//	return -1
//}
//
//func Include(vs []string, t string) bool {
//	return Index(vs, t) >= 0
//}
//
//func Any(vs []string, f func(string) bool) bool {
//	for _, v := range vs {
//		if f(v) {
//			return true
//		}
//	}
//	return false
//}
//
//func All(vs []string, f func(string) bool) bool {
//	for _, v := range vs {
//		if !f(v) {
//			return false
//		}
//	}
//	return true
//}
//
//func Filter(vs []string, f func(string) bool) []string {
//	// filter 默认长度是0
//	vsf := make([]string, 0)
//	for _, v := range vs {
//		if f(v) {
//			vsf = append(vsf, v)
//		}
//	}
//	return vsf
//}
//
//func Map(vs []string, f func(string) string) []string {
//	// 因为是map所以创建的长度就是原slice长度
//	vsm := make([]string, len(vs))
//	for _, v := range vs {
//		vsm = append(vsm, f(v))
//	}
//	return vsm
//}

//func createFile(p string) *os.File {
//	fmt.Println("creating")
//	f, err := os.Create(p)
//	if err != nil {
//		panic(err)
//	}
//	return f
//}
//
//func writeFile(f *os.File) {
//	fmt.Println("writing")
//	fmt.Fprintln(f, "data")
//}
//
//func closeFile(f *os.File) {
//	fmt.Println("closing")
//	err := f.Close()
//	if err != nil {
//		fmt.Fprintf(os.Stderr, "error%v\n", err)
//		os.Exit(1)
//	}
//}

//type byLength []string
//
//func (s byLength) Len() int {
//	return len(s)
//}
//
//func (s byLength) Swap(i, j int) {
//	s[i], s[j] = s[j], s[i]
//}
//
//func (s byLength) Less(i, j int) bool {
//	return len(s[i]) < len(s[j])
//}

//type readOp struct {
//	key int
//	resp chan int
//}
//
//type writeOp struct {
//	key int
//	val int
//	resp chan bool
//}

//func worker(id int, wg *sync.WaitGroup) {
//	defer wg.Done()
//
//	fmt.Printf("worker %d starting\n", id)
//	time.Sleep(time.Second)
//	fmt.Printf("worker %d done\n", id)
//}

//func worker(id int, jobs <-chan int, results chan<- int) {
//	for j := range jobs {
//		fmt.Println("worker", id, "start job", j)
//		time.Sleep(time.Second)
//		fmt.Println("worker", id, "end job", j)
//		results <- j*2
//	}
//}

//func ping(pings chan<- string, msg string) {
//	pings <- msg
//}
//
//func pong(pings <-chan string, pongs chan<- string)  {
//	msg := <-pings
//	pongs <- msg
//}

//func worker(done chan bool) {
//	fmt.Println("working")
//	time.Sleep(time.Second)
//	fmt.Println("done")
//
//	done <- true
//}

//func f(from string) {
//	for i := 0; i < 6; i++ {
//		fmt.Println(from, ":", i)
//	}
//}
//
//func f1(arg int) (int, error) {
//	if arg == 42 {
//		return -1, errors.New("can't work with 42")
//	}
//	return arg + 3, nil
//}
//
//type argError struct {
//	arg int
//	prob string
//}
//
//func (ae *argError) Error() string {
//	return fmt.Sprintf("%d - %s", ae.arg, ae.prob)
//}
//
//func f2(arg int) (int, error) {
//	if arg == 42 {
//		return -1, &argError{arg: arg, prob: "can't work with 42"}
//	}
//	return arg + 3, nil
//}

//func measure(g geometry) {
//	fmt.Println(g)
//	fmt.Println(g.area())
//	fmt.Println(g.perim())
//}
//
//type geometry interface {
//	area() float64
//	perim() float64
//}
//
//type rect struct {
//	width float64
//	height float64
//}
//
//type circle struct {
//	radius float64
//}
//
//func (r rect) area() float64 {
//	return r.width * r.height
//}
//
//func (r rect) perim() float64 {
//	return r.width * 2 + r.height * 2
//}
//
//func (c circle) area() float64 {
//	return math.Pi*c.radius*c.radius
//}
//
//func (c circle) perim() float64 {
//	return 2*math.Pi*c.radius
//}

//type rect struct {
//	width int
//	height int
//}
//
//func (r *rect) area() int {
//	fmt.Printf("%p\n", &r)
//	return r.width * r.height
//}
//
//func (r rect) perim() int {
//	fmt.Printf("%p\n", &r)
//	return r.width*2 + r.height*2
//}
//
//type person struct {
//	name string
//	age int
//}
//
//func zeroVal(ival int) {
//	ival = 0
//}
//
//func ptrVal(ptrVal *int) {
//	*ptrVal = 0
//}
//
//func fact(n int) int {
//	if n == 0 {
//		return 1
//	}
//	return n*fact(n-1)
//}
//
//func intSeq() func() int {
//	i := 0
//	return func() int {
//		i++
//		return i
//	}
//}
//func sum(nums ...int) {
//	fmt.Println(nums)
//	total := 0
//	for _, num := range nums {
//		total += num
//	}
//	fmt.Println(total)
//}
//
//func vals() (int, int) {
//	return 3, 4
//}
//
//func plus(a int, b int) int {
//	return a + b
//}
//
//func plusPlus(a, b, c int) int {
//	return a + b + c
//}
//
//
//const s string = "constant"
