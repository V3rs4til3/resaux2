package main

/*func main() {
	s := make(chan int, 5)
	s <- 42
	go fct1(s)
	go fct2(s)
	go fct3(s)
	go fct4(s)
	go fct5(s)
	go fct6(s)
}

func fct1(s chan int) {
	s <- 2 * <-s
}
func fct2(s chan int) {
	s <- <-s / 3
}

func fct3(s chan int) {
	s <- <-s - 10
}

func fct4(s chan int) {
	if <-s > 0 {
		s <- 0
	}
}

func fct5(s chan int) {
	s <- <-s + 1
}

func fct6(s chan int) {
	fmt.Println(<-s)
}*/

/*func main() {
	var c chan string = make(chan string)
	var i chan int = make(chan int)
	var o chan int = make(chan int)
	for {
		go ping(c, i, o)
		go pong(c, i, o)
		go printer(c)
	}
	var input string
	fmt.Scanln(&input)
}

func ping(c chan string, i chan int, o chan int) {
	if random() == 1 {
		ptsPing := <-i
		s := "Point pour ping, score " + strconv.Itoa(ptsPing) + " - " + strconv.Itoa(<-o)
		c <- s
		i <- ptsPing + 1
	} else {
		c <- "ping"
	}
}

func pong(c chan string, i chan int, o chan int) {
	if random() == 1 {
		ptsPong := <-o
		s := "Point pour pong, score " + strconv.Itoa(<-i) + " - " + strconv.Itoa(ptsPong)
		c <- s
	} else {
		c <- "pong"
	}
}

func printer(c chan string) {
	for {
		msg := <-c
		fmt.Println(msg)
		time.Sleep(time.Second * 10)
	}
}

func random() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(10)
}*/

/*func main() {
	p := Paladin{pv: 0, attaque: 20, defense: 10}
	afficherStatut(&p)
}

type Paladin struct {
	nom     string
	pv      int
	attaque int
	defense int
}

func (p *Paladin) vivant() bool {
	return p.pv > 0
}

func (p *Paladin) defOuOff() string {
	if p.defense > p.attaque {
		return "defensif"
	} else if p.attaque > p.defense {
		return "offensif"
	}
	return "equilibre"

}

type Chevalier interface {
	vivant() bool
	defOuOff() string
}

func afficherStatut(c Chevalier) {
	fmt.Println(c.defOuOff())
	fmt.Println(c.vivant())
}*/

/*var upgrader = websocket.Upgrader{}
var mySocket *websocket.Conn

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/calcul", ws)

	fileServer := http.FileServer(http.Dir("./html"))
	http.Handle("/html/", http.StripPrefix("/html/", fileServer))

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Println("ListendAndServe:", err)
		return
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	content, _ := os.ReadFile("./html/home.html")
	io.WriteString(w, string(content))
}

func ws(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	mySocket = c
	if err != nil {
		return
	}
	var s = make(chan string)
	go listen(c, s)
	go calcul(<-s, s)
	go write(c, s)
}

func listen(c *websocket.Conn, s chan string) {
	for {
		_, msg, err := c.ReadMessage()
		log.Println("recu " + string(msg))
		if err != nil {
			log.Println(err)
			return
		}
		s <- string(msg)
	}
}

func write(c *websocket.Conn, s chan string) {
	for {
		c.WriteMessage(1, []byte(<-s))
	}
}

func calcul(toCalcul string, s chan string) {
	var calcul, _ = strconv.Atoi(toCalcul)
	calcul = calcul * 2
	s <- strconv.Itoa(calcul)
}*/
