package models

import "time"

//Transaction - Транзакции(произведенные операции)
type Transaction struct {
	IDTransPluspay   int64     //Идентификатор транзакции(произведенной операции)
	IDService        int64     //Идентификатор услуги СМ.ТАБЛ. service
	PickupPoint      string    //Идентификатор приемного терминала(точки доступа)
	MoneyIn          float32   //Оперируеммая сумма(деньги)
	CurrencyIn       string    //Тип валюты
	MoneyOut         float32   //Оперируеммая сумма(деньги)
	CurrencyOut      string    //Тип валюты
	ComissionPluspay float32   //Комиссия pluspay за услугу
	ComissionContact float32   //Комиссия contact за услугу
	ComissionOther   float32   //Комиссия остальная за услугу
	DateTimeZone     time.Time //Дата и время с учетом часового пояса
	DateTimeContact  string    /*Дата и время системы контакт*/
}

//TransactionsID - ID транзакций систем-участников
type TransactionsID struct {
	IDTransPluspay int64  //Идентификатор транзакции нашей системы
	IDTransContact string //Номер транзакции системы contact
	IDTransBank    string //Номер транзакции оконечного банка
}

//ServiceNum2 - данные для таблицы servicenum2 (2 физ лица и ID транзакции
type ServiceNum2 struct {
	IDTransPluspay    int64     //Идентификатор транзакции
	Sndrlastname      string    //Имя пользователя
	Sndrsecondname    string    //Отчество пользователя
	Sndrfirstname     string    //Фамилия пользователя
	Sndrbirthday      time.Time //Дата рождения
	Sndrbirthplace    string    //Место рождения
	SndrdocType       string    //Тип документа
	SndrdocNumber     string    //Серия и номер документа
	SndrdocIssuedBy   string    //Кем выдан
	SndrdocDepartment string    //Код подразделения
	SndrdocDateOn     time.Time //Дата выдачи
	SndrdocDateOff    time.Time //Дата истечения
	Sndrcitizenship   string    //Гражданство
	Sndrresident      string    //Резидент
	SndradrCountry    string    //Страна
	SndradrIndex      string    //Индекс
	SndradrRegion     string    //Регион
	SndradrCity       string    //Город
	SndradrAdress     string    //Адрес(улица/дом/квартира)
	Sndrphone         string    //Номер телефона
	Rsvrlastname      string    //Фамилия пользователя
	Rsvrfirstname     string    //Имя пользователя
	Rsvrsecondname    string    //Отчество пользователя
	Rsvrbirthday      string    //Дата рождения
	Rsvrbirthplace    string    //Место рождения
	RsvrdocType       string    //Тип документа
	RsvrdocNumber     string    //Серия и номер документа
	RsvrdocIssuedBy   string    //Кем выдан
	RsvrdocDepartment string    //Код подразделения
	RsvrdocDateOn     time.Time //Дата выдачи
	RsvrdocDateOff    time.Time //Дата истечения
	Rsvrcitizenship   string    //Гражданство
	Rsvrresident      string    //Резидент
	RsvradrCountry    string    //Страна
	RsvradrIndex      string    //Индекс
	RsvradrRegion     string    //Регион
	RsvradrCity       string    //Город
	RsvradrAdress     string    //Адрес(улица/дом/квартира)
	Rsvrphone         string    //Номер телефона
	Info              string    //Дополнительная информация
}

//Service - Виды возможннноых операций(услуг)
type Service struct {
	IDService   int64  //Идентификатор услуги
	ServiceName string //Название услуги
	Info        string //Подробное описание услуги
}

//TransactionOfServiceNum2 - структура для выборки полной информации из 2го сервиса
type TransactionOfServiceNum2 struct {
	Trns   Transaction    //Данные из табл. transaction
	TrnsID TransactionsID //Данные из табл. transactionsID
	Users  ServiceNum2    //Данные из табл. servicenum2
}
