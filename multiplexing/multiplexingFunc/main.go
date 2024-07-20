package main

import (
	"fmt"
	"sync"
)

func multiplexingFunc(channels ...chan int) <-chan int {
	var wg sync.WaitGroup
	// Общий канал, в который будут попадать сообщения от всех источников
	// Иммено его мы и вернем из этой функции для употребления внешним кодом
	multiplexedChan := make(chan int)
	multiplex := func(c <-chan int) {
		defer wg.Done()
		for i := range c {
			// Если поступило сообщение из одного из
			// каналов-источников
			// перенапралвяем его в общий канал
			multiplexedChan <- i
		}
	}
	wg.Add(len(channels))
	for _, c := range channels {
		go multiplex(c)
	}
	// Запускаем горутину, которая закроет канал после того,
	// как в закрывающий канал поступит сигнал о прекращении
	// работы всех
	go func() {
		wg.Wait()
		close(multiplexedChan)
	}()

	return multiplexedChan
}

// Функция разуплотнения каналов
func demultiplexingFunc(dataSourceChan chan int, amount int) []chan int {
	var output = make([]chan int, amount)
	// Обратите внимание: вышеприведённая команда инициализирует слайс,
	// но не инициализирует каждый его элемент. Каждый элемент будет
	// представлять так называемое нулевое значение
	// для данного типа.
	// Так как тип у нас канала — ссылочный, то все элементы
	// будут
	// равны nil
	for i := range output {
		output[i] = make(chan int)
	}
	go func() {
		var wg sync.WaitGroup
		var once sync.Once

		wg.Add(1)
		go func() {
			defer wg.Done()
			// При поступлении сообщений в канал-источник
			// отправляем его в каждый из каналов-потребителей
			for v := range dataSourceChan {
				for _, c := range output {
					c <- v
				}
			}
		}()
		wg.Wait()

		// Закрываем все каналы-потребители одной командой
		once.Do(func() {
			for _, c := range output {
				close(c)
			}
		})
	}()

	return output
}

// Число сообщений от каждого истчника
const messagesAmountPerGoroutine int = 5

func main() {
	dataSourceChans := make([]chan int, 0)

	// Горутина - источник данных
	// Функция создает свой собственный канал
	// и посылает в него пять сообщений
	startDataSource := func(start int) chan int {
		c := make(chan int)
		var wg sync.WaitGroup
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := start; i < start+messagesAmountPerGoroutine; i++ {
				c <- i
			}
		}()
		go func() {
			wg.Wait()
			close(c)
		}()
		return c
	}

	// Запускаем 5 источников
	// Канал от каждого сохраняем в наш специальный буфер
	for i := messagesAmountPerGoroutine; i <= 20; i += messagesAmountPerGoroutine {
		dataSourceChans = append(dataSourceChans, startDataSource(i))
	}

	// Уплотняем канал
	c := multiplexingFunc(dataSourceChans...)

	for data := range c {
		fmt.Println("data: ", data)
	}
	fmt.Println("_______________________________")

	newStartDataSource := func() chan int {
		c := make(chan int)
		go func() {
			var wg sync.WaitGroup
			wg.Add(1)
			go func() {
				defer wg.Done()
				for i := 1; i <= messagesAmountPerGoroutine; i++ {
					c <- i
				}
			}()
			wg.Wait()
			close(c)
		}()
		return c
	}

	var consumers []chan int = demultiplexingFunc(newStartDataSource(), 5)
	c = multiplexingFunc(consumers...)
	// Централизованно получаем сообщения от всех нужных нам
	// источников
	// данных
	for data := range c {
		fmt.Println("data: ", data)
	}
}
