package main

import (
	"fmt"
	"time"
)

/*
Задача 1:
Реализация системы оплаты с использованием интерфейсов

Описание: Создайте интерфейс PaymentProcessor с методом Process(amount float64) string, который возвращает строку с результатом обработки платежа. Реализуйте два типа платежных систем: CreditCard и CryptoWallet, каждый из которых реализует интерфейс PaymentProcessor. В функции main создайте список платежных систем и вызовите метод Process для каждого, выводя результат на экран.

Требования:

    Интерфейс PaymentProcessor.
    Структуры CreditCard и CryptoWallet, реализующие интерфейс.
    В main создайте массив/слайс этих структур и вызовите их методы.
    P.S. Подумайте, какие поля могут быть у каждой структуры

*/

// Интерфейс PaymentProcessor
type PaymentProcessor interface {
	Process(amount float64) string
}

// Структура CreditCard
type CreditCard struct {
	cardHolder string
	cardNumber string    // номер карты
	cardDate   time.Time // дата окончания действия
}

func (card CreditCard) Process(amount float64) string {
	return fmt.Sprintf("Кредитная карта %s:\nПлатеж на сумму %.2f руб. обработан успешно",
		card.cardNumber, amount)
}

// Структура CryptoWallet
type CryptoWallet struct {
	cryptoHash     string // хэш по которуму производится перевод
	cryptoCurrency string // валюта
}

func (wallet CryptoWallet) Process(amount float64) string {
	return fmt.Sprintf("Криптокошелек %s:\nПеревод %.2f %s выполнен",
		wallet.cryptoHash, amount, wallet.cryptoCurrency)
}

func main() {
	fmt.Println("Программа 1")

	paymentProcessors := []PaymentProcessor{
		CreditCard{
			cardNumber: "1234-5678-9012-1111",
			cardHolder: "Дмитрий В.",
			cardDate:   time.Date(2025, time.March, 15, 14, 30, 15, 0, time.UTC),
		},
		CryptoWallet{
			cryptoHash:     "0x1a2b3c4d5e6f7890",
			cryptoCurrency: "BTC",
		},
		CreditCard{
			cardNumber: "1234-5678-9012-7777",
			cardHolder: "Алексей В.",
			cardDate:   time.Date(2025, time.March, 20, 00, 00, 00, 0, time.UTC),
		},
		CryptoWallet{
			cryptoHash:     "0x9f8e7d6c5b4a3f2e1",
			cryptoCurrency: "ETH",
		},
	}

	amounts := []float64{1500.50, 0.025, 750.25, 1.5}

	for i, processor := range paymentProcessors {
		result := processor.Process(amounts[i])
		fmt.Printf("%d. %s\n", i+1, result)
	}

}
