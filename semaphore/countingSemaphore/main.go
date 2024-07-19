package main

import (
	"fmt"
	"time"
)

// Semaphore - структура семафора подсчета
type Semaphore struct {
	/**
	Семафор - астрактный тип данных,
	в нашем случае в основе его лежит канал
	*/
	sem chan struct{}
	/**
	Время ожидания основных операций с семафором, чтобы не
	блокировать операции с ним навечно
	*/
	timeout time.Duration
}

// Acquire - метод захвата семафора
func (s *Semaphore) Acquire() error {
	select {
	case s.sem <- struct{}{}:
		return nil
	case <-time.After(s.timeout):
		return fmt.Errorf("Не удалось захватить семафор")
	}
}

// Release - метод освобождения семафора
func (s *Semaphore) Release() error {
	select {
	case _ = <-s.sem:
		return nil
	case <-time.After(s.timeout):
		return fmt.Errorf("Не удалось освободить семафор")
	}
}

// NewSemaphore - функция создания семафора
func NewSemaphore(initialCount, maxCount int, timeout time.Duration) *Semaphore {
	sem := make(chan struct{}, maxCount)
	for i := 0; i < initialCount; i++ {
		sem <- struct{}{}
	}
	return &Semaphore{sem, timeout}
}

func main() {
	// Создаем семафор с начальным значением 2 и максимальным значением 3
	sem := NewSemaphore(2, 3, 2*time.Second)

	// Пытаемся захватить семафор 3 раза
	for i := 0; i < 3; i++ {
		go func(i int) {
			if err := sem.Acquire(); err != nil {
				fmt.Printf("Горутина %d: не удалось захватить семафор: %v\n", i, err)
			} else {
				fmt.Printf("Горутина %d: захватила семафор\n", i)
				// Освобожаем семафор после некоторого времени
				time.Sleep(1 * time.Second)
				if err := sem.Release(); err != nil {
					fmt.Printf("Горутина %d: не удалось освободить семаформ: %v\n", i, err)
				} else {
					fmt.Printf("Горутина %d: освободила семафор\n", i)
				}
			}
		}(i)
	}

	// Ждем завершения всех горутин
	time.Sleep(5 * time.Second)
}
