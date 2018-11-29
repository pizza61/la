# la

Platforma do ładnych ankiet, może być jedna, mogą być dwie, ale nikt nie wie jak je utworzyć!

Docelowo to ma być platforma do ankiet, może być uruchomiona na serwerze specjalnie dla jednej czy kilku ankiet, może być też dla wielu więcej.

## Wymagania
* MongoDB
* Go
* Node.js

## Jak sprawić, aby działało?
### Captcha
Korzystamy z captchy v3, więc należy skonfigurować wszystko dobrze na [stronie recaptchy](https://www.google.com/recaptcha/admin)

Po stworzeniu, dodaniu stron, interesuje nas zakładka **Keys**

(w przyszłości to nie będzie tak wyglądało)

* Site key: w pliku client/src/components/Okno.vue odnajdujemy zmienną sekret (znajduje się zaraz pod `<script>`) i tam między `"` wklejamy Site key
* Secret key: w pliku server.go odnajdujemy zmienną sekret (znajduje się na poczatku pliku, zaraz po `func main()`) i tam między `"` wklejamy Secret key
### Go
* ```go get github.com/labstack/echo/...```
* ```go get github.com/globalsign/mgo/...```

### Kompilowanie vue
* Wchodzimy do `/client`
* `npm install`
* `npm run build`

### Uruchamianie
Aby uruchomić, wystarczy w głównym katalogu wpisać `go run server.go`. Możesz też zbudować i uruchomić W TYM KATALOGU uzyskany plik.

Domyślnie serwer jest uruchomiany na porcie 8080, więc w przeglądarce wystarczy wejść pod adres `http://localhost:8080`
### Przykładowa ankieta
Aby utworzyć przykładową ankietę, wchodzimy w przeglądarce pod adres `/real`. Ankieta została utworzona, teraz możemy wejść pod adres `/ankieta/real` i ją wypełnić.

## Czego nie ma a musi być
* Dodawanie ankiet (będzie mogło być zablokowane)
* Sprawdzanie odpowiedzi

## Screen dla zachęty
![pokazane tu jest](https://i.imgur.com/BpBRLP5.png)
