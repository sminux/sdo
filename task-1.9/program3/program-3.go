package main

import "fmt"

/*
Задача 3:
Структура "Банковский счет" и методы для работы с балансом
Описание:
Создайте структуру BankAccount с полями: Owner (владелец счета) и Balance (текущий баланс).
Реализуйте методы:
Deposit(amount float64), увеличивающий баланс.
Withdraw(amount float64), уменьшающий баланс, если хватает средств, иначе выводит сообщение о недостатке средств.

Что нужно сделать:
Объявить структуру и методы.
Создать счет, пополнить его, попытаться снять деньги и вывести итоговый баланс.

*/

// Структура BankAccount
type BankAccount struct {
	Owner   string
	Balance float64
}

// Метод Deposit пополняет счет
func (ba *BankAccount) Deposit(amount float64) {
	if amount <= 0 {
		fmt.Println("Ошибка: сумма пополнения должна быть положительной")
		return
	}
	ba.Balance += amount
	fmt.Printf("Счет пополнен на %.2f. Новый баланс: %.2f\n", amount, ba.Balance)
}

// Метод Withdraw снимает деньги со счета
func (ba *BankAccount) Withdraw(amount float64) {
	if amount <= 0 {
		fmt.Println("Ошибка: сумма снятия должна быть положительной")
		return
	}

	if amount > ba.Balance {
		fmt.Printf("Ошибка: недостаточно средств. Запрошено: %.2f, доступно: %.2f\n",
			amount, ba.Balance)
		return
	}

	ba.Balance -= amount
	fmt.Printf("Со счета снято %.2f. Новый баланс: %.2f\n", amount, ba.Balance)
}

// Метод для получения информации о счете
func (ba BankAccount) GetInfo() string {
	return fmt.Sprintf("Владелец: %s, Баланс: %.2f", ba.Owner, ba.Balance)
}

func main() {
	// Создаем банковский счет
	account := BankAccount{
		Owner:   "Дмитрий В.",
		Balance: 1000.0,
	}

	fmt.Println("Создан банковский счет:")
	fmt.Println(account.GetInfo())

	fmt.Println("\nОперации по счету:")
	// Пополняем счет
	account.Deposit(500.0)

	// Снимаем деньги
	account.Withdraw(200.0)

	// Пытаемся снять больше, чем есть на счете
	account.Withdraw(5000.0)

	// Еще одно пополнение
	account.Deposit(100.0)

	fmt.Println("\nИтоговое состояние счета:")
	fmt.Println(account.GetInfo())
}
