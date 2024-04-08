package main

import "fmt"

func main() {
	//Deadlock()
	//XDeadlock()
	XDeadlock2()
}

func Deadlock() {
	// channelar by default unbufferd channel hisoblanadi! yani unda qiymat ushlab turish uchun buffer bo'lmayi.
	// ushbu channelni bufferi bo'lmaganligi sababli, undan ma'lumot jo'naishimiz uchun, qabul qiluvchi tomon tayyor bo'lishi kerak
	var ch = make(chan int)

	// channelga ma'lumot yozdik
	ch <- 7
	ch <- 5
	ch <- 5

	//
	a := <-ch // deadlock!
	//ushbu holatda channelga yozish va undan qiymatni o'qish o'qish amallari main goroutineda bo'lganligi sababli, channelga ma'lumot yozishfda bu channel blokirovka bo'lib turadi

	fmt.Println(a)
}

// deadlikni oldini olishni 1- yo'li: channelni buffer bilan yaratish
func XDeadlock() {
	// channelar by default unbufferd channel hisoblanadi! yani unda qiymat ushlab turish uchun buffer bo'lmayi.
	// ushbu channelni bufferi bo'lmaganligi sababli, undan ma'lumot jo'naishimiz uchun, qabul qiluvchi tomon tayyor bo'lishi kerak
	var ch = make(chan int, 2)

	// channelga ma'lumot yozdik
	ch <- 7
	ch <- 5
	//ch <- 5

	//
	a := <-ch

	fmt.Println(a, <-ch)

	ch <- 7
	ch <- 5
	//ch <- 5

	//
	a = <-ch

	fmt.Println(a, <-ch)
}

// deadlikni oldini olishni 2- yo'li: channelga yozish amalini goroutine bilan amalga oshirish
func XDeadlock2() {
	// channelar by default unbufferd channel hisoblanadi! yani unda qiymat ushlab turish uchun buffer bo'lmayi.
	// ushbu channelni bufferi bo'lmaganligi sababli, undan ma'lumot jo'naishimiz uchun, qabul qiluvchi tomon tayyor bo'lishi kerak
	var ch = make(chan int)

	// channelga ma'lumot yozdik
	go func() {
		ch <- 7
		ch <- 7
		ch <- 7
	}()

	//
	a := <-ch // deadlock!
	//ushbu holatda channelga yozish va undan qiymatni o'qish o'qish amallari main goroutineda bo'lganligi sababli, channelga ma'lumot yozishfda bu channel blokirovka bo'lib turadi

	fmt.Println(a, <-ch, <-ch)
}
