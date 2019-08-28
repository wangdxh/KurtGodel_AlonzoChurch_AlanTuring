package main

import (
	"fmt"
	"github.com/streadway/amqp"
	"os"
	"sync"
)

func main() {
	bmain := true
	if len(os.Args) > 1 {
		bmain = false
	}

	conn, err := amqp.Dial("amqp://wxh:wxh@172.16.119.167:5672/")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()
	fmt.Println("1")

	chn, err := conn.Channel()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer chn.Close()
	fmt.Println("2")

	err = chn.Qos(1, 0, true)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("2.5")

	queanaly, err := chn.QueueDeclare("kd.analysepic", true, false, false, false, amqp.Table{})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(queanaly)
	fmt.Println("3")

	quecapturegbid, err := chn.QueueDeclare("kd.capturegbid", true, false, false, false, amqp.Table{})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(quecapturegbid)
	fmt.Println("4")

	quemain, err := chn.QueueDeclare("kd.main", true, false, false, false, amqp.Table{})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(quemain)
	fmt.Println("5")

	err = chn.ExchangeDeclare("kd.broadcast", "fanout",
		true, false, false, false, amqp.Table{})
	if err != nil {
		fmt.Println(err)
		return
	}

	if bmain {
		chnreturn := make(chan amqp.Return, 100)
		chn.NotifyReturn(chnreturn)
		go func() {
			for {
				select {
				case ret, bok := <-chnreturn:
					if !bok {
						break
					}
					fmt.Println("main get return: ", string(ret.Body))
				}
			}
		}()
		consume("kd.main", chn, "main")

		for {
			var strinfo string
			fmt.Scan(&strinfo)
			err := chn.Publish("kd.broadcast", "keyfanout", true, false,
				amqp.Publishing{
					ContentType:  "text/plain",
					Body:         []byte(strinfo),
					DeliveryMode: amqp.Persistent,
					Expiration:   "30000",
				})
			if err != nil {
				fmt.Println(err)
			}

			err = chn.Publish("", "kd.capturegbid", true, false,
				amqp.Publishing{
					ContentType:  "application/json",
					Body:         []byte(strinfo + "capturegbid"),
					DeliveryMode: amqp.Persistent,
					Expiration:   "30000",
				})
			if err != nil {
				fmt.Println(err)
			}

			err = chn.Publish("", "kd.analysepic", true, false,
				amqp.Publishing{
					ContentType:  "application/json",
					Body:         []byte(strinfo + "analyse"),
					DeliveryMode: amqp.Persistent,
					Expiration:   "30000",
				})
			if err != nil {
				fmt.Println(err)
			}
		}

	} else {
		fmt.Println(6)
		quename := "kd." + os.Args[1]
		quebroad, err := chn.QueueDeclare(quename, true, false, false, false, amqp.Table{})
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(quebroad)

		err = chn.QueueBind(quename, "fanoutkey", "kd.broadcast", false, amqp.Table{})
		if err != nil {
			fmt.Println(err)
			return
		}
		consume(quename, chn, os.Args[1]+"-exe1")
		consume("kd.capturegbid", chn, os.Args[1]+"-capture")
		consume("kd.analysepic", chn, os.Args[1]+"-analyse")

		for {
			var strinfo string
			fmt.Scan(&strinfo)
			err := chn.Publish("", "kd.main", false, false,
				amqp.Publishing{
					ContentType:  "text/plain",
					Body:         []byte(strinfo + "-tomain"),
					DeliveryMode: amqp.Persistent,
					//Expiration:   "30000",
				})
			if err != nil {
				fmt.Println(err)
			}
		}

	}

	var wg sync.WaitGroup
	wg.Add(1)
	wg.Wait()
}

func consume(quename string, chn *amqp.Channel, consumeid string) {
	go func() {
		mainchn, err := chn.Consume(quename, consumeid, false, false, false, false, amqp.Table{})
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("consumeok")
		for {
			select {
			case delivery, bok := <-mainchn:
				if !bok {
					break
				}
				fmt.Println(consumeid, " get: ", string(delivery.Body))
				chn.Ack(delivery.DeliveryTag, false)
			}
		}

	}()
}
