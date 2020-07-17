package main

func main() {
	config := NewConfig()
	config.Parse()

	printAPIDocs(config)
}
